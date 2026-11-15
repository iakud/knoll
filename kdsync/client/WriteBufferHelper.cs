using System;
using System.Buffers;
using System.Runtime.CompilerServices;
using System.Security;

namespace Kdsync
{
    [SecuritySafeCritical]
    internal struct WriteBufferHelper
    {
        private IBufferWriter<byte> bufferWriter;

        [MethodImpl(MethodImplOptions.AggressiveInlining)]
        public static void Initialize(IBufferWriter<byte> bufferWriter, out WriteBufferHelper instance, out Span<byte> buffer)
        {
            instance.bufferWriter = bufferWriter;
            buffer = default;  // TODO: initialize the initial buffer so that the first write is not via slowpath.
        }

        [MethodImpl(MethodImplOptions.AggressiveInlining)]
        public static void InitializeNonRefreshable(out WriteBufferHelper instance)
        {
            instance.bufferWriter = null;
        }

        [MethodImpl(MethodImplOptions.AggressiveInlining)]
        public static void CheckNoSpaceLeft(ref WriterInternalState state)
        {
            if (GetSpaceLeft(ref state) != 0)
            {
                throw new InvalidOperationException("Did not write as much data as expected.");
            }
        }

        [MethodImpl(MethodImplOptions.AggressiveInlining)]
        public static int GetSpaceLeft(ref WriterInternalState state)
        {
            if (state.writeBufferHelper.bufferWriter == null)
            {
                return state.limit - state.position;
            }
            else
            {
                throw new InvalidOperationException(
                    "SpaceLeft can only be called on WriteContext that are " +
                        "writing to a flat array or when writing to a single span.");
            }
        }

        [MethodImpl(MethodImplOptions.NoInlining)]
        public static void RefreshBuffer(ref Span<byte> buffer, ref WriterInternalState state)
        {
            if (state.writeBufferHelper.bufferWriter != null)
            {
                // commit the bytes and get a new buffer to write to.
                state.writeBufferHelper.bufferWriter.Advance(state.position);
                state.position = 0;
                buffer = state.writeBufferHelper.bufferWriter.GetSpan();
                state.limit = buffer.Length;
            }
            else
            {
                // We're writing to a single buffer.
                throw new OutOfSpaceException();
            }
        }

        [MethodImpl(MethodImplOptions.AggressiveInlining)]
        public static void Flush(ref Span<byte> buffer, ref WriterInternalState state)
        {
            if (state.writeBufferHelper.bufferWriter != null)
            {
                // calling Advance invalidates the current buffer and we must not continue writing to it,
                // so we set the current buffer to point to an empty Span. If any subsequent writes happen,
                // the first subsequent write will trigger refresing of the buffer.
                state.writeBufferHelper.bufferWriter.Advance(state.position);
                state.position = 0;
                state.limit = 0;
                buffer = default;  // invalidate the current buffer
            }
        }
    }
}