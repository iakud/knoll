using System;
using System.IO;
using System.Runtime.CompilerServices;
using System.Runtime.InteropServices;

namespace kds
{
    public static class Example
    {
        private static All _all = new All(0);
        private static byte[] _lastData = Array.Empty<byte>();

        public static void Initialize()
        {
            try
            {
                SetupChangeListener(_all);
                Console.Out.WriteLine($"Initialized");
            }
            catch (Exception ex)
            {
                Console.Out.WriteLine($"Init error: {ex.Message}");
            }
        }

        static void SetupChangeListener(All all)
        {
            // FIXME: fields onchanged


            all.OnChanged += (sender, e) => Console.Out.WriteLine($"All.Changed\n");
        }

        [UnmanagedCallersOnly(EntryPoint = "apply_sync", CallConvs = new[] { typeof(CallConvCdecl) })]
        public static int ApplySync(IntPtr dataPtr, int length)
        {
            try
            {
                var data = new byte[length];
                Marshal.Copy(dataPtr, data, 0, length);

                _all.ApplySync(data);
                _all.RaiseChanged();
                _all.ClearChanged();

                // Store last data for dump
                _lastData = data;

                Console.Out.WriteLine($"ApplySync: {length} bytes");
                return 0;
            }
            catch (Exception ex)
            {


                Console.Out.WriteLine($"ApplySync error: {ex.Message}, stack: {ex.StackTrace}");
                return 1;
            }
        }

        [UnmanagedCallersOnly(EntryPoint = "dump", CallConvs = new[] { typeof(CallConvCdecl) })]
        public static IntPtr Dump()
        {
            try
            {
                // 遍历 _all 中所有字段并组装成字符串
                var types = _all.Types;
                var sb = new System.Text.StringBuilder();

                // 基础数值类型
                sb.AppendLine($"Int32Val={types.Int32Val}");
                sb.AppendLine($"Int64Val={types.Int64Val}");
                sb.AppendLine($"Uint32Val={types.Uint32Val}");
                sb.AppendLine($"Uint64Val={types.Uint64Val}");
                sb.AppendLine($"Sint32Val={types.Sint32Val}");
                sb.AppendLine($"Sint64Val={types.Sint64Val}");
                sb.AppendLine($"Fixed32Val={types.Fixed32Val}");
                sb.AppendLine($"Fixed64Val={types.Fixed64Val}");
                sb.AppendLine($"Sfixed32Val={types.Sfixed32Val}");
                sb.AppendLine($"Sfixed64Val={types.Sfixed64Val}");

                // 浮点类型
                sb.AppendLine($"FloatVal={types.FloatVal:F6}");
                sb.AppendLine($"DoubleVal={types.DoubleVal:F6}");

                // 布尔和字符串
                sb.AppendLine($"BoolVal={types.BoolVal.ToString().ToLower()}");
                sb.AppendLine($"StringVal={types.StringVal}");
                sb.AppendLine($"BytesVal={Convert.ToBase64String(types.BytesVal)}");

                // 时间类型
                sb.AppendLine($"TimestampVal={(types.TimestampVal.ToUniversalTime() - new DateTime(1970,1,1,0,0,0,DateTimeKind.Utc)).Ticks * 100}");
                sb.AppendLine($"DurationVal={types.DurationVal.Ticks * 100}");

                // 枚举类型
                sb.AppendLine($"EnumVal={(int)types.EnumVal}");

                // 自定义类型
                sb.Append($"ItemData=({types.ItemData.Id},{types.ItemData.Name},{types.ItemData.Count})");

                return Marshal.StringToHGlobalAnsi(sb.ToString());
            }
            catch (Exception ex)
            {
                Console.Out.WriteLine($"Dump error: {ex.Message}, stack: {ex.StackTrace}");
                return Marshal.StringToHGlobalAnsi("");
            }
        }
    }
}
