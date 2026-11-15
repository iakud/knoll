using System;

namespace Kdsync
{
    public static class WriteComputeSize
    {
        private const int LittleEndian64Size = 8;
        private const int LittleEndian32Size = 4;

        internal const int DoubleSize = LittleEndian64Size;
        internal const int FloatSize = LittleEndian32Size;
        internal const int BoolSize = 1;

        public static int ComputeDoubleSize(double value)
        {
            return DoubleSize;
        }

        public static int ComputeFloatSize(float value)
        {
            return FloatSize;
        }

        public static int ComputeUInt64Size(ulong value)
        {
            return ComputeRawVarint64Size(value);
        }

        public static int ComputeInt64Size(long value)
        {
            return ComputeRawVarint64Size((ulong)value);
        }

        public static int ComputeInt32Size(int value)
        {
            if (value >= 0)
            {
                return ComputeRawVarint32Size((uint)value);
            }
            else
            {
                // Must sign-extend.
                return 10;
            }
        }

        public static int ComputeFixed64Size(ulong value)
        {
            return LittleEndian64Size;
        }

        public static int ComputeFixed32Size(uint value)
        {
            return LittleEndian32Size;
        }

        public static int ComputeBoolSize(bool value)
        {
            return BoolSize;
        }

        public static int ComputeStringSize(String value)
        {
            int byteArraySize = WritingPrimitives.Utf8Encoding.GetByteCount(value);
            return ComputeLengthSize(byteArraySize) + byteArraySize;
        }

        public static int ComputeMessageSize(IMessage value)
        {
            int size = value.CalculateSize();
            return ComputeLengthSize(size) + size;
        }

        public static int ComputeBytesSize(byte[] value)
        {
            return ComputeLengthSize(value.Length) + value.Length;
        }

        public static int ComputeUInt32Size(uint value)
        {
            return ComputeRawVarint32Size(value);
        }

        public static int ComputeEnumSize(int value)
        {
            // Currently just a pass-through, but it's nice to separate it logically.
            return ComputeInt32Size(value);
        }

        public static int ComputeSFixed32Size(int value)
        {
            return LittleEndian32Size;
        }

        public static int ComputeSFixed64Size(long value)
        {
            return LittleEndian64Size;
        }

        public static int ComputeSInt32Size(int value)
        {
            return ComputeRawVarint32Size(WritingPrimitives.EncodeZigZag32(value));
        }

        public static int ComputeSInt64Size(long value)
        {
            return ComputeRawVarint64Size(WritingPrimitives.EncodeZigZag64(value));
        }

        public static int ComputeLengthSize(int length)
        {
            return ComputeRawVarint32Size((uint)length);
        }

        public static int ComputeRawVarint32Size(uint value)
        {
            if ((value & (0xffffffff << 7)) == 0)
            {
                return 1;
            }
            if ((value & (0xffffffff << 14)) == 0)
            {
                return 2;
            }
            if ((value & (0xffffffff << 21)) == 0)
            {
                return 3;
            }
            if ((value & (0xffffffff << 28)) == 0)
            {
                return 4;
            }
            return 5;
        }

        public static int ComputeRawVarint64Size(ulong value)
        {
            if ((value & (0xffffffffffffffffL << 7)) == 0)
            {
                return 1;
            }
            if ((value & (0xffffffffffffffffL << 14)) == 0)
            {
                return 2;
            }
            if ((value & (0xffffffffffffffffL << 21)) == 0)
            {
                return 3;
            }
            if ((value & (0xffffffffffffffffL << 28)) == 0)
            {
                return 4;
            }
            if ((value & (0xffffffffffffffffL << 35)) == 0)
            {
                return 5;
            }
            if ((value & (0xffffffffffffffffL << 42)) == 0)
            {
                return 6;
            }
            if ((value & (0xffffffffffffffffL << 49)) == 0)
            {
                return 7;
            }
            if ((value & (0xffffffffffffffffL << 56)) == 0)
            {
                return 8;
            }
            if ((value & (0xffffffffffffffffL << 63)) == 0)
            {
                return 9;
            }
            return 10;
        }

        public static int ComputeTagSize(int fieldNumber)
        {
            return ComputeRawVarint32Size(WireFormat.MakeTag(fieldNumber, 0));
        }

        public static int ComputeTimestampContentSize(Timestamp value)
        {
            return ComputeTagSize(Timestamp.SecondsFieldNumber) + ComputeInt64Size(value.Seconds)
                 + ComputeTagSize(Timestamp.NanosFieldNumber) + ComputeInt32Size(value.Nanos);
        }

        public static int ComputeTimestampSize(Timestamp value)
        {
            int contentSize = ComputeTimestampContentSize(value);
            return ComputeLengthSize(contentSize) + contentSize;
        }

        public static int ComputeDurationContentSize(Duration value)
        {
            return ComputeTagSize(Duration.SecondsFieldNumber) + ComputeInt64Size(value.Seconds)
                 + ComputeTagSize(Duration.NanosFieldNumber) + ComputeInt32Size(value.Nanos);
        }

        public static int ComputeDurationSize(Duration value)
        {
            int contentSize = ComputeDurationContentSize(value);
            return ComputeLengthSize(contentSize) + contentSize;
        }

        public static int ComputeEmptySize(Empty value)
        {
            return ComputeLengthSize(0);
        }
    }
}