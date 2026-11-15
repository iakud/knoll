using System.Runtime.CompilerServices;
using System.Security;

namespace Kdsync
{
    /// <summary>
    /// Writing messages / groups.
    /// </summary>
    [SecuritySafeCritical]
    internal static class WritingPrimitivesMessages
    {
        [MethodImpl(MethodImplOptions.AggressiveInlining)]
        public static void WriteMessage(ref WriteContext ctx, IMessage value)
        {
            WritingPrimitives.WriteLength(ref ctx.buffer, ref ctx.state, value.CalculateSize());
            WriteRawMessage(ref ctx, value);
        }

        [MethodImpl(MethodImplOptions.AggressiveInlining)]
        public static void WriteRawMessage(ref WriteContext ctx, IMessage message)
        {
            message.WriteTo(ref ctx);
        }

        [MethodImpl(MethodImplOptions.AggressiveInlining)]
        public static void WriteTimestamp(ref WriteContext ctx, Timestamp value)
        {
            ctx.WriteLength(WriteComputeSize.ComputeTimestampContentSize(value));
            ctx.WriteTag(Timestamp.SecondsFieldNumber, WireFormat.WireType.Varint);
            ctx.WriteInt64(value.Seconds);
            ctx.WriteTag(Timestamp.NanosFieldNumber, WireFormat.WireType.Varint);
            ctx.WriteInt32(value.Nanos);
        }

        [MethodImpl(MethodImplOptions.AggressiveInlining)]
        public static void WriteDuration(ref WriteContext ctx, Duration value)
        {
            ctx.WriteLength(WriteComputeSize.ComputeDurationContentSize(value));
            ctx.WriteTag(Duration.SecondsFieldNumber, WireFormat.WireType.Varint);
            ctx.WriteInt64(value.Seconds);
            ctx.WriteTag(Duration.NanosFieldNumber, WireFormat.WireType.Varint);
            ctx.WriteInt32(value.Nanos);
        }

        [MethodImpl(MethodImplOptions.AggressiveInlining)]
        public static void WriteEmpty(ref WriteContext ctx, Empty value)
        {
            ctx.WriteLength(0);
        }
    }
}