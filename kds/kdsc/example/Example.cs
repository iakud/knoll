using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Runtime.CompilerServices;
using System.Runtime.InteropServices;
using Google.Protobuf.WellKnownTypes;

namespace kds
{
    public static class Example
    {
        private static All _all = new All(0);

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
                var types = _all.Types;
                var lists = _all.Lists;
                var maps = _all.Maps;
                var sb = new System.Text.StringBuilder();

                // Types
                sb.Append($"Int32Val={types.Int32Val}");
                sb.Append(", ");
                sb.Append($"Int64Val={types.Int64Val}");
                sb.Append(", ");
                sb.Append($"Uint32Val={types.Uint32Val}");
                sb.Append(", ");
                sb.Append($"Uint64Val={types.Uint64Val}");
                sb.Append(", ");
                sb.Append($"Sint32Val={types.Sint32Val}");
                sb.Append(", ");
                sb.Append($"Sint64Val={types.Sint64Val}");
                sb.Append(", ");
                sb.Append($"Fixed32Val={types.Fixed32Val}");
                sb.Append(", ");
                sb.Append($"Fixed64Val={types.Fixed64Val}");
                sb.Append(", ");
                sb.Append($"Sfixed32Val={types.Sfixed32Val}");
                sb.Append(", ");
                sb.Append($"Sfixed64Val={types.Sfixed64Val}");
                sb.Append(", ");
                sb.Append($"FloatVal={types.FloatVal:F6}");
                sb.Append(", ");
                sb.Append($"DoubleVal={types.DoubleVal:F6}");
                sb.Append(", ");
                sb.Append($"BoolVal={types.BoolVal.ToString().ToLower()}");
                sb.Append(", ");
                sb.Append($"StringVal={types.StringVal}");
                sb.Append(", ");
                sb.Append($"BytesVal={Convert.ToBase64String(types.BytesVal)}");
                sb.Append(", ");
                sb.Append($"TimestampVal={types.TimestampVal.Seconds * 1000000000 + types.TimestampVal.Nanos}");
                sb.Append(", ");
                sb.Append($"DurationVal={types.DurationVal.Seconds * Duration.NanosecondsPerSecond + types.DurationVal.Nanos}");
                sb.Append(", ");
                sb.Append($"EnumVal={(int)types.EnumVal}");
                sb.Append(", ");
                sb.Append($"ItemData=({types.ItemData.Id},{types.ItemData.Name},{types.ItemData.Count})");
                sb.Append(", ");

                // Lists
                sb.Append($"Int32List={DumpInt32List(lists.Int32List)}");
                sb.Append(", ");
                sb.Append($"Int64List={DumpInt64List(lists.Int64List)}");
                sb.Append(", ");
                sb.Append($"FloatList={DumpFloatList(lists.FloatList)}");
                sb.Append(", ");
                sb.Append($"DoubleList={DumpDoubleList(lists.DoubleList)}");
                sb.Append(", ");
                sb.Append($"BoolList={DumpBoolList(lists.BoolList)}");
                sb.Append(", ");
                sb.Append($"StringList=[{string.Join(",", lists.StringList)}]");
                sb.Append(", ");
                sb.Append($"TimestampList={DumpTimestampList(lists.TimestampList)}");
                sb.Append(", ");
                sb.Append($"DurationList={DumpDurationList(lists.DurationList)}");
                sb.Append(", ");
                sb.Append($"EmptyList={lists.EmptyList.Count}");
                sb.Append(", ");
                sb.Append($"EnumList={DumpEnumList(lists.EnumList)}");
                sb.Append(", ");
                sb.Append($"ItemList={DumpItemDataList(lists.ItemList)}");
                sb.Append(", ");

                // Maps
                sb.Append($"Int32Int32={DumpInt32Int32Map(maps.Int32Int32)}");
                sb.Append(", ");
                sb.Append($"Int32String={DumpInt32StringMap(maps.Int32String)}");
                sb.Append(", ");
                sb.Append($"Int32Timestamp={DumpInt32TimestampMap(maps.Int32Timestamp)}");
                sb.Append(", ");
                sb.Append($"Int32Duration={DumpInt32DurationMap(maps.Int32Duration)}");
                sb.Append(", ");
                sb.Append($"Int32Empty={maps.Int32Empty.Count}");
                sb.Append(", ");
                sb.Append($"Int32Enum={DumpInt32EnumMap(maps.Int32Enum)}");
                sb.Append(", ");
                sb.Append($"Int32ItemData={DumpInt32ItemDataMap(maps.Int32ItemData)}");
                sb.Append(", ");
                sb.Append($"Int64Int64={DumpInt64Int64Map(maps.Int64Int64)}");
                sb.Append(", ");
                sb.Append($"Int64String={DumpInt64StringMap(maps.Int64String)}");
                sb.Append(", ");
                sb.Append($"Int64Timestamp={DumpInt64TimestampMap(maps.Int64Timestamp)}");
                sb.Append(", ");
                sb.Append($"Int64Duration={DumpInt64DurationMap(maps.Int64Duration)}");
                sb.Append(", ");
                sb.Append($"Int64Empty={maps.Int64Empty.Count}");
                sb.Append(", ");
                sb.Append($"Int64Enum={DumpInt64EnumMap(maps.Int64Enum)}");
                sb.Append(", ");
                sb.Append($"Int64ItemData={DumpInt64ItemDataMap(maps.Int64ItemData)}");
                sb.Append(", ");
                sb.Append($"StringInt32={DumpStringInt32Map(maps.StringInt32)}");
                sb.Append(", ");
                sb.Append($"StringString={DumpStringStringMap(maps.StringString)}");
                sb.Append(", ");
                sb.Append($"StringTimestamp={DumpStringTimestampMap(maps.StringTimestamp)}");
                sb.Append(", ");
                sb.Append($"StringDuration={DumpStringDurationMap(maps.StringDuration)}");
                sb.Append(", ");
                sb.Append($"StringEmpty={maps.StringEmpty.Count}");
                sb.Append(", ");
                sb.Append($"StringEnum={DumpStringEnumMap(maps.StringEnum)}");
                sb.Append(", ");
                sb.Append($"StringItemData={DumpStringItemDataMap(maps.StringItemData)}");
                sb.Append(", ");
                sb.Append($"BoolInt32={DumpBoolInt32Map(maps.BoolInt32)}");
                sb.Append(", ");
                sb.Append($"BoolString={DumpBoolStringMap(maps.BoolString)}");
                sb.Append(", ");
                sb.Append($"BoolTimestamp={DumpBoolTimestampMap(maps.BoolTimestamp)}");
                sb.Append(", ");
                sb.Append($"BoolDuration={DumpBoolDurationMap(maps.BoolDuration)}");
                sb.Append(", ");
                sb.Append($"BoolEmpty={maps.BoolEmpty.Count}");
                sb.Append(", ");
                sb.Append($"BoolEnum={DumpBoolEnumMap(maps.BoolEnum)}");
                sb.Append(", ");
                sb.Append($"BoolItemData={DumpBoolItemDataMap(maps.BoolItemData)}");

                return Marshal.StringToHGlobalAnsi(sb.ToString());
            }
            catch (Exception ex)
            {
                Console.Out.WriteLine($"Dump error: {ex.Message}, stack: {ex.StackTrace}");
                return Marshal.StringToHGlobalAnsi("");
            }
        }

        static string DumpInt32List(List<int> l) => l.Count == 0 ? "[]" : $"[{string.Join(",", l)}]";
        static string DumpInt64List(List<long> l) => l.Count == 0 ? "[]" : $"[{string.Join(",", l)}]";
        static string DumpFloatList(List<float> l) => l.Count == 0 ? "[]" : $"[{string.Join(",", l.Select(x => x.ToString("F6")))}]";
        static string DumpDoubleList(List<double> l) => l.Count == 0 ? "[]" : $"[{string.Join(",", l.Select(x => x.ToString("F6")))}]";
        static string DumpBoolList(List<bool> l) => l.Count == 0 ? "[]" : $"[{string.Join(",", l.Select(x => x.ToString().ToLower()))}]";
        static string DumpTimestampList(List<Timestamp> l) => l.Count == 0 ? "[]" : $"[{string.Join(",", l.Select(x => x.Seconds* 1000000000 + x.Nanos))}]";
        static string DumpDurationList(List<Duration> l) => l.Count == 0 ? "[]" : $"[{string.Join(",", l.Select(x => x.Seconds * Duration.NanosecondsPerSecond + x.Nanos))}]";
        static string DumpEnumList(List<ItemType> l) => l.Count == 0 ? "[]" : $"[{string.Join(",", l.Select(x => (int)x))}]";
        static string DumpItemDataList(List<ItemData> l) => l.Count == 0 ? "[]" : $"[{string.Join(",", l.Select(x => $"({x.Id},{x.Name},{x.Count})"))}]";

        static string DumpInt32Int32Map(Dictionary<int, int> m) => m.Count == 0 ? "map[]" : $"map[{string.Join(",", m.Select(kv => $"{kv.Key}:{kv.Value}"))}]";
        static string DumpInt32StringMap(Dictionary<int, string> m) => m.Count == 0 ? "map[]" : $"map[{string.Join(",", m.Select(kv => $"{kv.Key}:{kv.Value}"))}]";
        static string DumpInt32TimestampMap(Dictionary<int, Timestamp> m) => m.Count == 0 ? "map[]" : $"map[{string.Join(",", m.Select(kv => $"{kv.Key}:{kv.Value.Seconds * 1000000000 + kv.Value.Nanos}"))}]";
        static string DumpInt32DurationMap(Dictionary<int, Duration> m) => m.Count == 0 ? "map[]" : $"map[{string.Join(",", m.Select(kv => $"{kv.Key}:{kv.Value.Seconds * Duration.NanosecondsPerSecond + kv.Value.Nanos}"))}]";
        static string DumpInt32EnumMap(Dictionary<int, ItemType> m) => m.Count == 0 ? "map[]" : $"map[{string.Join(",", m.Select(kv => $"{kv.Key}:{(int)kv.Value}"))}]";
        static string DumpInt32ItemDataMap(Dictionary<int, ItemData> m) => m.Count == 0 ? "map[]" : $"map[{string.Join(",", m.Select(kv => $"{kv.Key}:({kv.Value.Id},{kv.Value.Name},{kv.Value.Count})"))}]";

        static string DumpInt64Int64Map(Dictionary<long, long> m) => m.Count == 0 ? "map[]" : $"map[{string.Join(",", m.Select(kv => $"{kv.Key}:{kv.Value}"))}]";
        static string DumpInt64StringMap(Dictionary<long, string> m) => m.Count == 0 ? "map[]" : $"map[{string.Join(",", m.Select(kv => $"{kv.Key}:{kv.Value}"))}]";
        static string DumpInt64TimestampMap(Dictionary<long, Timestamp> m) => m.Count == 0 ? "map[]" : $"map[{string.Join(",", m.Select(kv => $"{kv.Key}:{kv.Value.Seconds * 1000000000 + kv.Value.Nanos}"))}]";
        static string DumpInt64DurationMap(Dictionary<long, Duration> m) => m.Count == 0 ? "map[]" : $"map[{string.Join(",", m.Select(kv => $"{kv.Key}:{kv.Value.Seconds * Duration.NanosecondsPerSecond + kv.Value.Nanos}"))}]";
        static string DumpInt64EnumMap(Dictionary<long, ItemType> m) => m.Count == 0 ? "map[]" : $"map[{string.Join(",", m.Select(kv => $"{kv.Key}:{(int)kv.Value}"))}]";
        static string DumpInt64ItemDataMap(Dictionary<long, ItemData> m) => m.Count == 0 ? "map[]" : $"map[{string.Join(",", m.Select(kv => $"{kv.Key}:({kv.Value.Id},{kv.Value.Name},{kv.Value.Count})"))}]";

        static string DumpStringInt32Map(Dictionary<string, int> m) => m.Count == 0 ? "map[]" : $"map[{string.Join(",", m.Select(kv => $"{kv.Key}:{kv.Value}"))}]";
        static string DumpStringStringMap(Dictionary<string, string> m) => m.Count == 0 ? "map[]" : $"map[{string.Join(",", m.Select(kv => $"{kv.Key}:{kv.Value}"))}]";
        static string DumpStringTimestampMap(Dictionary<string, Timestamp> m) => m.Count == 0 ? "map[]" : $"map[{string.Join(",", m.Select(kv => $"{kv.Key}:{kv.Value.Seconds * 1000000000 + kv.Value.Nanos}"))}]";
        static string DumpStringDurationMap(Dictionary<string, Duration> m) => m.Count == 0 ? "map[]" : $"map[{string.Join(",", m.Select(kv => $"{kv.Key}:{kv.Value.Seconds * Duration.NanosecondsPerSecond + kv.Value.Nanos}"))}]";
        static string DumpStringEnumMap(Dictionary<string, ItemType> m) => m.Count == 0 ? "map[]" : $"map[{string.Join(",", m.Select(kv => $"{kv.Key}:{(int)kv.Value}"))}]";
        static string DumpStringItemDataMap(Dictionary<string, ItemData> m) => m.Count == 0 ? "map[]" : $"map[{string.Join(",", m.Select(kv => $"{kv.Key}:({kv.Value.Id},{kv.Value.Name},{kv.Value.Count})"))}]";

        static string DumpBoolInt32Map(Dictionary<bool, int> m) => m.Count == 0 ? "map[]" : $"map[{string.Join(",", m.Select(kv => $"{kv.Key}:{kv.Value}"))}]";
        static string DumpBoolStringMap(Dictionary<bool, string> m) => m.Count == 0 ? "map[]" : $"map[{string.Join(",", m.Select(kv => $"{kv.Key}:{kv.Value}"))}]";
        static string DumpBoolTimestampMap(Dictionary<bool, Timestamp> m) => m.Count == 0 ? "map[]" : $"map[{string.Join(",", m.Select(kv => $"{kv.Key}:{kv.Value.Seconds * 1000000000 + kv.Value.Nanos}"))}]";
        static string DumpBoolDurationMap(Dictionary<bool, Duration> m) => m.Count == 0 ? "map[]" : $"map[{string.Join(",", m.Select(kv => $"{kv.Key}:{kv.Value.Seconds * Duration.NanosecondsPerSecond + kv.Value.Nanos}"))}]";
        static string DumpBoolEnumMap(Dictionary<bool, ItemType> m) => m.Count == 0 ? "map[]" : $"map[{string.Join(",", m.Select(kv => $"{kv.Key}:{(int)kv.Value}"))}]";
        static string DumpBoolItemDataMap(Dictionary<bool, ItemData> m) => m.Count == 0 ? "map[]" : $"map[{string.Join(",", m.Select(kv => $"{kv.Key}:({kv.Value.Id},{kv.Value.Name},{kv.Value.Count})"))}]";
    }
}
