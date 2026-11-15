using System;
using System.Buffers;
using System.Runtime.CompilerServices;
using System.Security;

namespace Kdsync
{
    [SecuritySafeCritical]
    public ref struct WriteContext
    {
        internal Span<byte> buffer;
        internal WriterInternalState state;

        [MethodImpl(MethodImplOptions.AggressiveInlining)]
        internal static void Initialize(ref Span<byte> buffer, ref WriterInternalState state, out WriteContext ctx)
        {
            ctx.buffer = buffer;
            ctx.state = state;
        }

        [MethodImpl(MethodImplOptions.AggressiveInlining)]
        public static void Initialize(IBufferWriter<byte> output, out WriteContext ctx)
        {
            ctx.buffer = default;
            ctx.state = default;
            WriteBufferHelper.Initialize(output, out ctx.state.writeBufferHelper, out ctx.buffer);
            ctx.state.limit = ctx.buffer.Length;
            ctx.state.position = 0;
        }

        [MethodImpl(MethodImplOptions.AggressiveInlining)]
        public static void Initialize(ref Span<byte> buffer, out WriteContext ctx)
        {
            ctx.buffer = buffer;
            ctx.state = default;
            ctx.state.limit = ctx.buffer.Length;
            ctx.state.position = 0;
            WriteBufferHelper.InitializeNonRefreshable(out ctx.state.writeBufferHelper);
        }

        public void WriteDouble(double value) => WritingPrimitives.WriteDouble(ref buffer, ref state, value);

        public void WriteFloat(float value) => WritingPrimitives.WriteFloat(ref buffer, ref state, value);

        public void WriteUInt64(ulong value) => WritingPrimitives.WriteUInt64(ref buffer, ref state, value);

        public void WriteInt64(long value) => WritingPrimitives.WriteInt64(ref buffer, ref state, value);

        public void WriteInt32(int value) => WritingPrimitives.WriteInt32(ref buffer, ref state, value);

        public void WriteFixed64(ulong value) => WritingPrimitives.WriteFixed64(ref buffer, ref state, value);

        public void WriteFixed32(uint value) => WritingPrimitives.WriteFixed32(ref buffer, ref state, value);

        public void WriteBool(bool value) => WritingPrimitives.WriteBool(ref buffer, ref state, value);

        public void WriteString(string value) => WritingPrimitives.WriteString(ref buffer, ref state, value);

        public void WriteMessage(IMessage value) => WritingPrimitivesMessages.WriteMessage(ref this, value);

        public void WriteTimestamp(Timestamp value) => WritingPrimitivesMessages.WriteTimestamp(ref this, value);

        public void WriteDuration(Duration value) => WritingPrimitivesMessages.WriteDuration(ref this, value);

        public void WriteEmpty(Empty value) => WritingPrimitivesMessages.WriteEmpty(ref this, value);

        public void WriteBytes(byte[] value) => WritingPrimitives.WriteBytes(ref buffer, ref state, value);

        public void WriteUInt32(uint value) => WritingPrimitives.WriteUInt32(ref buffer, ref state, value);

        public void WriteEnum(int value) => WritingPrimitives.WriteEnum(ref buffer, ref state, value);

        public void WriteSFixed32(int value) => WritingPrimitives.WriteSFixed32(ref buffer, ref state, value);

        public void WriteSFixed64(long value) => WritingPrimitives.WriteSFixed64(ref buffer, ref state, value);

        public void WriteSInt32(int value) => WritingPrimitives.WriteSInt32(ref buffer, ref state, value);

        public void WriteSInt64(long value) => WritingPrimitives.WriteSInt64(ref buffer, ref state, value);

        public void WriteLength(int length) => WritingPrimitives.WriteLength(ref buffer, ref state, length);

        public void WriteTag(int fieldNumber, WireFormat.WireType type) => WritingPrimitives.WriteTag(ref buffer, ref state, fieldNumber, type);

        public void WriteTag(uint tag) => WritingPrimitives.WriteTag(ref buffer, ref state, tag);

        public void WriteRawTag(byte b1) => WritingPrimitives.WriteRawTag(ref buffer, ref state, b1);

        public void WriteRawTag(byte b1, byte b2) => WritingPrimitives.WriteRawTag(ref buffer, ref state, b1, b2);

        public void WriteRawTag(byte b1, byte b2, byte b3) => WritingPrimitives.WriteRawTag(ref buffer, ref state, b1, b2, b3);

        public void WriteRawTag(byte b1, byte b2, byte b3, byte b4) => WritingPrimitives.WriteRawTag(ref buffer, ref state, b1, b2, b3, b4);

        public void WriteRawTag(byte b1, byte b2, byte b3, byte b4, byte b5) => WritingPrimitives.WriteRawTag(ref buffer, ref state, b1, b2, b3, b4, b5);

        public void Flush() => WriteBufferHelper.Flush(ref buffer, ref state);

        public void CheckNoSpaceLeft() => WriteBufferHelper.CheckNoSpaceLeft(ref state);
    }
}