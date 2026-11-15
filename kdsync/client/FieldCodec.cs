using System;

namespace Kdsync
{
    internal delegate TValue ValueReader<out TValue>(ref ParseContext ctx);
    internal delegate void ValueWriter<T>(ref WriteContext ctx, T value);

    public sealed class FieldCodec<T>
    {
        internal delegate void InputMerger(ref ParseContext ctx, ref T value);

        internal delegate bool ValuesMerger(ref T value, T other);

        internal delegate void JsonValueWriter(JsonWriter writer, T value);

        private static readonly T DefaultDefault;

        internal ValueWriter<T> ValueWriter { get; }

        internal Func<T, int> ValueSizeCalculator { get; }

        internal ValueReader<T> ValueReader { get; }

        internal InputMerger ValueMerger { get; }

        internal JsonValueWriter JsonWriter { get; }

        internal WireFormat.WireType WireType { get; }

        internal T DefaultValue { get; }

        static FieldCodec()
        {
            if (typeof(T) == typeof(string))
            {
                DefaultDefault = (T)(object)"";
            }
            else if (typeof(T) == typeof(byte[]))
            {
                DefaultDefault = (T)(object)(new byte[0]);
            }
        }

        internal FieldCodec(ValueReader<T> reader, ValueWriter<T> writer, Func<T, int> sizeCalculator, JsonValueWriter jsonWriter, WireFormat.WireType wireType, T defaultValue)
            : this(reader, writer, delegate (ref ParseContext ctx, ref T v)
            {
                v = reader(ref ctx);
            }, sizeCalculator, jsonWriter, wireType, defaultValue)
        {
        }

        internal FieldCodec(ValueReader<T> reader, ValueWriter<T> writer, InputMerger inputMerger, Func<T, int> sizeCalculator, JsonValueWriter jsonWriter, WireFormat.WireType wireType)
            : this(reader, writer, inputMerger, sizeCalculator, jsonWriter, wireType, DefaultDefault)
        {
        }

        internal FieldCodec(ValueReader<T> reader, ValueWriter<T> writer, InputMerger inputMerger, Func<T, int> sizeCalculator, JsonValueWriter jsonWriter, WireFormat.WireType wireType, T defaultValue)
        {
            ValueReader = reader;
            ValueWriter = writer;
            ValueMerger = inputMerger;
            ValueSizeCalculator = sizeCalculator;
            JsonWriter = jsonWriter;
            WireType = wireType;
            DefaultValue = defaultValue;
        }

        public T Read(ref ParseContext ctx)
        {
            return ValueReader(ref ctx);
        }
    }

    public static class FieldCodec
    {
        public static FieldCodec<bool> ForBoolKey()
        {
            return new FieldCodec<bool>(delegate (ref ParseContext ctx)
            {
                return ctx.ReadBool();
            }, delegate (ref WriteContext ctx, bool value)
            {
                ctx.WriteBool(value);
            }, WriteComputeSize.ComputeBoolSize, JsonWriter.WriteName, WireFormat.WireType.Varint, false);
        }

        public static FieldCodec<int> ForIntKey()
        {
            return new FieldCodec<int>(delegate (ref ParseContext ctx)
            {
                return ctx.ReadInt32();
            }, delegate (ref WriteContext ctx, int value)
            {
                ctx.WriteInt32(value);
            }, WriteComputeSize.ComputeInt32Size, JsonWriter.WriteName, WireFormat.WireType.Varint, 0);
        }

        public static FieldCodec<uint> ForUIntKey()
        {
            return new FieldCodec<uint>(delegate (ref ParseContext ctx)
            {
                return ctx.ReadUInt32();
            }, delegate (ref WriteContext ctx, uint value)
            {
                ctx.WriteUInt32(value);
            }, WriteComputeSize.ComputeUInt32Size, JsonWriter.WriteName, WireFormat.WireType.Varint, 0u);
        }

        public static FieldCodec<long> ForLongKey()
        {
            return new FieldCodec<long>(delegate (ref ParseContext ctx)
            {
                return ctx.ReadInt64();
            }, delegate (ref WriteContext ctx, long value)
            {
                ctx.WriteInt64(value);
            }, WriteComputeSize.ComputeInt64Size, JsonWriter.WriteName, WireFormat.WireType.Varint, 0L);
        }

        public static FieldCodec<ulong> ForULongKey()
        {
            return new FieldCodec<ulong>(delegate (ref ParseContext ctx)
            {
                return ctx.ReadUInt64();
            }, delegate (ref WriteContext ctx, ulong value)
            {
                ctx.WriteUInt64(value);
            }, WriteComputeSize.ComputeUInt64Size, JsonWriter.WriteName, WireFormat.WireType.Varint, 0uL);
        }

        public static FieldCodec<string> ForStringKey()
        {
            return new FieldCodec<string>(delegate (ref ParseContext ctx)
            {
                return ctx.ReadString();
            }, delegate (ref WriteContext ctx, string value)
            {
                ctx.WriteString(value);
            }, WriteComputeSize.ComputeStringSize, JsonWriter.WriteName, WireFormat.WireType.LengthDelimited, "");
        }

        // Value Codec

        public static FieldCodec<bool> ForBoolValue()
        {
            return new FieldCodec<bool>(delegate (ref ParseContext ctx)
            {
                return ctx.ReadBool();
            }, delegate (ref WriteContext ctx, bool value)
            {
                ctx.WriteBool(value);
            }, WriteComputeSize.ComputeBoolSize, JsonWriter.WriteBoolValue, WireFormat.WireType.Varint, false);
        }

        public static FieldCodec<int> ForIntValue()
        {
            return new FieldCodec<int>(delegate (ref ParseContext ctx)
            {
                return ctx.ReadInt32();
            }, delegate (ref WriteContext ctx, int value)
            {
                ctx.WriteInt32(value);
            }, WriteComputeSize.ComputeInt32Size, JsonWriter.WriteIntValue, WireFormat.WireType.Varint, 0);
        }

        public static FieldCodec<uint> ForUIntValue()
        {
            return new FieldCodec<uint>(delegate (ref ParseContext ctx)
            {
                return ctx.ReadUInt32();
            }, delegate (ref WriteContext ctx, uint value)
            {
                ctx.WriteUInt32(value);
            }, WriteComputeSize.ComputeUInt32Size, JsonWriter.WriteUIntValue, WireFormat.WireType.Varint, 0u);
        }

        public static FieldCodec<long> ForLongValue()
        {
            return new FieldCodec<long>(delegate (ref ParseContext ctx)
            {
                return ctx.ReadInt64();
            }, delegate (ref WriteContext ctx, long value)
            {
                ctx.WriteInt64(value);
            }, WriteComputeSize.ComputeInt64Size, JsonWriter.WriteLongValue, WireFormat.WireType.Varint, 0L);
        }

        public static FieldCodec<ulong> ForULongValue()
        {
            return new FieldCodec<ulong>(delegate (ref ParseContext ctx)
            {
                return ctx.ReadUInt64();
            }, delegate (ref WriteContext ctx, ulong value)
            {
                ctx.WriteUInt64(value);
            }, WriteComputeSize.ComputeUInt64Size, JsonWriter.WriteULongValue, WireFormat.WireType.Varint, 0uL);
        }

        public static FieldCodec<float> ForFloatValue()
        {
            return new FieldCodec<float>(delegate (ref ParseContext ctx)
            {
                return ctx.ReadFloat();
            }, delegate (ref WriteContext ctx, float value)
            {
                ctx.WriteFloat(value);
            }, WriteComputeSize.ComputeFloatSize, JsonWriter.WriteFloatValue, WireFormat.WireType.Fixed32, 0f);
        }

        public static FieldCodec<double> ForDoubleValue()
        {
            return new FieldCodec<double>(delegate (ref ParseContext ctx)
            {
                return ctx.ReadDouble();
            }, delegate (ref WriteContext ctx, double value)
            {
                ctx.WriteDouble(value);
            }, WriteComputeSize.ComputeDoubleSize, JsonWriter.WriteDoubleValue, WireFormat.WireType.Fixed64, 0.0);
        }

        public static FieldCodec<byte[]> ForBytesValue()
        {
            return new FieldCodec<byte[]>(delegate (ref ParseContext ctx)
            {
                return ctx.ReadBytes();
            }, delegate (ref WriteContext ctx, byte[] value)
            {
                ctx.WriteBytes(value);
            }, WriteComputeSize.ComputeBytesSize, JsonWriter.WriteBytesValue, WireFormat.WireType.LengthDelimited, Array.Empty<byte>());
        }

        public static FieldCodec<string> ForStringValue()
        {
            return new FieldCodec<string>(delegate (ref ParseContext ctx)
            {
                return ctx.ReadString();
            }, delegate (ref WriteContext ctx, string value)
            {
                ctx.WriteString(value);
            }, WriteComputeSize.ComputeStringSize, JsonWriter.WriteStringValue, WireFormat.WireType.LengthDelimited, "");
        }

        public static FieldCodec<Timestamp> ForTimestampValue()
        {
            return new FieldCodec<Timestamp>(delegate (ref ParseContext ctx)
            {
                return ctx.ReadTimestamp();
            }, delegate (ref WriteContext ctx, Timestamp value)
            {
                ctx.WriteTimestamp(value);
            }, WriteComputeSize.ComputeTimestampSize, JsonWriter.WriteTimestampValue, WireFormat.WireType.LengthDelimited, default);
        }

        public static FieldCodec<Duration> ForDurationValue()
        {
            return new FieldCodec<Duration>(delegate (ref ParseContext ctx)
            {
                return ctx.ReadDuration();
            }, delegate (ref WriteContext ctx, Duration value)
            {
                ctx.WriteDuration(value);
            }, WriteComputeSize.ComputeDurationSize, JsonWriter.WriteDurationValue, WireFormat.WireType.LengthDelimited, default);
        }

        public static FieldCodec<Empty> ForEmptyValue()
        {
            return new FieldCodec<Empty>(delegate (ref ParseContext ctx)
            {
                return ctx.ReadEmpty();
            }, delegate (ref WriteContext ctx, Empty value)
            {
                ctx.WriteEmpty(value);
            }, WriteComputeSize.ComputeEmptySize, JsonWriter.WriteEmptyValue, WireFormat.WireType.LengthDelimited, default);
        }

        public static FieldCodec<T> ForEnum<T>() where T : Enum
        {
            return new FieldCodec<T>(delegate (ref ParseContext ctx)
            {
                return (T)(object)ctx.ReadEnum();
            }, delegate (ref WriteContext ctx, T value)
            {
                ctx.WriteEnum(Convert.ToInt32(value));
            }, delegate (T value)
            {
                return WriteComputeSize.ComputeEnumSize(Convert.ToInt32(value));
            }, JsonWriter.WriteEnumValue, WireFormat.WireType.Varint, default);
        }

        public static FieldCodec<T> ForMessage<T>() where T : class, IMessage, new()
        {
            return new FieldCodec<T>(delegate (ref ParseContext ctx)
            {
                // T val = parser.CreateTemplate();
                T val = new T();
                ctx.ReadMessage(val);
                return val;
            }, delegate (ref WriteContext ctx, T value)
            {
                ctx.WriteMessage(value);
            }, delegate (ref ParseContext ctx, ref T v)
            {
                if (v == null)
                {
                    // v = parser.CreateTemplate();
                    v = new T();
                }

                ctx.ReadMessage(v);
            }, delegate (T value)
            {
                return WriteComputeSize.ComputeMessageSize(value);
            }, JsonWriter.WriteValue, WireFormat.WireType.LengthDelimited);
        }
    }
}
