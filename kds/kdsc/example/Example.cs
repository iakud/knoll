using System;
using System.IO;
using System.Runtime.CompilerServices;
using System.Runtime.InteropServices;

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
                sb.AppendLine($"FloatVal={types.FloatVal:F6}");
                sb.AppendLine($"DoubleVal={types.DoubleVal:F6}");
                sb.AppendLine($"BoolVal={types.BoolVal.ToString().ToLower()}");
                sb.AppendLine($"StringVal={types.StringVal}");
                sb.AppendLine($"BytesVal={Convert.ToBase64String(types.BytesVal)}");
                sb.AppendLine($"TimestampVal={(types.TimestampVal.ToUniversalTime() - new DateTime(1970,1,1,0,0,0,DateTimeKind.Utc)).Ticks * 100}");
                sb.AppendLine($"DurationVal={types.DurationVal.Ticks * 100}");
                sb.AppendLine($"EnumVal={(int)types.EnumVal}");
                sb.AppendLine($"ItemData=({types.ItemData.Id},{types.ItemData.Name},{types.ItemData.Count})");

                // Lists
                sb.AppendLine($"Int32List={DumpInt32List(lists.Int32List)}");
                sb.AppendLine($"Int64List={DumpInt64List(lists.Int64List)}");
                sb.AppendLine($"FloatList={DumpFloatList(lists.FloatList)}");
                sb.AppendLine($"DoubleList={DumpDoubleList(lists.DoubleList)}");
                sb.AppendLine($"BoolList={DumpBoolList(lists.BoolList)}");
                sb.AppendLine($"StringList=[{string.Join(",", lists.StringList)}]");
                sb.AppendLine($"TimestampList={DumpTimestampList(lists.TimestampList)}");
                sb.AppendLine($"DurationList={DumpDurationList(lists.DurationList)}");
                sb.AppendLine($"EmptyList={lists.EmptyList.Count}");
                sb.AppendLine($"EnumList={DumpEnumList(lists.EnumList)}");
                sb.AppendLine($"ItemList={DumpItemDataList(lists.ItemList)}");

                // Maps
                sb.AppendLine($"Int32Int32={DumpInt32Int32Map(maps.Int32Int32)}");
                sb.AppendLine($"Int32String={DumpInt32StringMap(maps.Int32String)}");
                sb.AppendLine($"Int32Timestamp={DumpInt32TimestampMap(maps.Int32Timestamp)}");
                sb.AppendLine($"Int32Duration={DumpInt32DurationMap(maps.Int32Duration)}");
                sb.AppendLine($"Int32Empty={maps.Int32Empty.Count}");
                sb.AppendLine($"Int32Enum={DumpInt32EnumMap(maps.Int32Enum)}");
                sb.AppendLine($"Int32ItemData={DumpInt32ItemDataMap(maps.Int32ItemData)}");
                sb.AppendLine($"Int64Int64={DumpInt64Int64Map(maps.Int64Int64)}");
                sb.AppendLine($"Int64String={DumpInt64StringMap(maps.Int64String)}");
                sb.AppendLine($"Int64Timestamp={DumpInt64TimestampMap(maps.Int64Timestamp)}");
                sb.AppendLine($"Int64Duration={DumpInt64DurationMap(maps.Int64Duration)}");
                sb.AppendLine($"Int64Empty={maps.Int64Empty.Count}");
                sb.AppendLine($"Int64Enum={DumpInt64EnumMap(maps.Int64Enum)}");
                sb.AppendLine($"Int64ItemData={DumpInt64ItemDataMap(maps.Int64ItemData)}");
                sb.AppendLine($"StringInt32={DumpStringInt32Map(maps.StringInt32)}");
                sb.AppendLine($"StringString={DumpStringStringMap(maps.StringString)}");
                sb.AppendLine($"StringTimestamp={DumpStringTimestampMap(maps.StringTimestamp)}");
                sb.AppendLine($"StringDuration={DumpStringDurationMap(maps.StringDuration)}");
                sb.AppendLine($"StringEmpty={maps.StringEmpty.Count}");
                sb.AppendLine($"StringEnum={DumpStringEnumMap(maps.StringEnum)}");
                sb.AppendLine($"StringItemData={DumpStringItemDataMap(maps.StringItemData)}");
                sb.AppendLine($"BoolInt32={DumpBoolInt32Map(maps.BoolInt32)}");
                sb.AppendLine($"BoolString={DumpBoolStringMap(maps.BoolString)}");
                sb.AppendLine($"BoolTimestamp={DumpBoolTimestampMap(maps.BoolTimestamp)}");
                sb.AppendLine($"BoolDuration={DumpBoolDurationMap(maps.BoolDuration)}");
                sb.AppendLine($"BoolEmpty={maps.BoolEmpty.Count}");
                sb.AppendLine($"BoolEnum={DumpBoolEnumMap(maps.BoolEnum)}");
                sb.AppendLine($"BoolItemData={DumpBoolItemDataMap(maps.BoolItemData)}");

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
        static string DumpTimestampList(List<DateTime> l) => l.Count == 0 ? "[]" : $"[{string.Join(",", l.Select(x => (x.ToUniversalTime() - new DateTime(1970,1,1,0,0,0,DateTimeKind.Utc)).Ticks * 100))}]";
        static string DumpDurationList(List<TimeSpan> l) => l.Count == 0 ? "[]" : $"[{string.Join(",", l.Select(x => x.Ticks * 100))}]";
        static string DumpEnumList(List<ItemType> l) => l.Count == 0 ? "[]" : $"[{string.Join(",", l.Select(x => (int)x))}]";
        static string DumpItemDataList(List<ItemData> l) => l.Count == 0 ? "[]" : $"[{string.Join(",", l.Select(x => $"({x.Id},{x.Name},{x.Count})"))}]";

        static string DumpInt32Int32Map(Dictionary<int, int> m) => m.Count == 0 ? "map[]" : $"map[{string.Join(",", m.Select(kv => $"{kv.Key}:{kv.Value}"))}]";
        static string DumpInt32StringMap(Dictionary<int, string> m) => m.Count == 0 ? "map[]" : $"map[{string.Join(",", m.Select(kv => $"{kv.Key}:{kv.Value}"))}]";
        static string DumpInt32TimestampMap(Dictionary<int, DateTime> m) => m.Count == 0 ? "map[]" : $"map[{string.Join(",", m.Select(kv => $"{kv.Key}:{(kv.Value.ToUniversalTime() - new DateTime(1970,1,1,0,0,0,DateTimeKind.Utc)).Ticks * 100}"))}]";
        static string DumpInt32DurationMap(Dictionary<int, TimeSpan> m) => m.Count == 0 ? "map[]" : $"map[{string.Join(",", m.Select(kv => $"{kv.Key}:{kv.Value.Ticks * 100}"))}]";
        static string DumpInt32EnumMap(Dictionary<int, ItemType> m) => m.Count == 0 ? "map[]" : $"map[{string.Join(",", m.Select(kv => $"{kv.Key}:{(int)kv.Value}"))}]";
        static string DumpInt32ItemDataMap(Dictionary<int, ItemData> m) => m.Count == 0 ? "map[]" : $"map[{string.Join(",", m.Select(kv => $"{kv.Key}:({kv.Value.Id},{kv.Value.Name},{kv.Value.Count})"))}]";

        static string DumpInt64Int64Map(Dictionary<long, long> m) => m.Count == 0 ? "map[]" : $"map[{string.Join(",", m.Select(kv => $"{kv.Key}:{kv.Value}"))}]";
        static string DumpInt64StringMap(Dictionary<long, string> m) => m.Count == 0 ? "map[]" : $"map[{string.Join(",", m.Select(kv => $"{kv.Key}:{kv.Value}"))}]";
        static string DumpInt64TimestampMap(Dictionary<long, DateTime> m) => m.Count == 0 ? "map[]" : $"map[{string.Join(",", m.Select(kv => $"{kv.Key}:{(kv.Value.ToUniversalTime() - new DateTime(1970,1,1,0,0,0,DateTimeKind.Utc)).Ticks * 100}"))}]";
        static string DumpInt64DurationMap(Dictionary<long, TimeSpan> m) => m.Count == 0 ? "map[]" : $"map[{string.Join(",", m.Select(kv => $"{kv.Key}:{kv.Value.Ticks * 100}"))}]";
        static string DumpInt64EnumMap(Dictionary<long, ItemType> m) => m.Count == 0 ? "map[]" : $"map[{string.Join(",", m.Select(kv => $"{kv.Key}:{(int)kv.Value}"))}]";
        static string DumpInt64ItemDataMap(Dictionary<long, ItemData> m) => m.Count == 0 ? "map[]" : $"map[{string.Join(",", m.Select(kv => $"{kv.Key}:({kv.Value.Id},{kv.Value.Name},{kv.Value.Count})"))}]";

        static string DumpStringInt32Map(Dictionary<string, int> m) => m.Count == 0 ? "map[]" : $"map[{string.Join(",", m.Select(kv => $"{kv.Key}:{kv.Value}"))}]";
        static string DumpStringStringMap(Dictionary<string, string> m) => m.Count == 0 ? "map[]" : $"map[{string.Join(",", m.Select(kv => $"{kv.Key}:{kv.Value}"))}]";
        static string DumpStringTimestampMap(Dictionary<string, DateTime> m) => m.Count == 0 ? "map[]" : $"map[{string.Join(",", m.Select(kv => $"{kv.Key}:{(kv.Value.ToUniversalTime() - new DateTime(1970,1,1,0,0,0,DateTimeKind.Utc)).Ticks * 100}"))}]";
        static string DumpStringDurationMap(Dictionary<string, TimeSpan> m) => m.Count == 0 ? "map[]" : $"map[{string.Join(",", m.Select(kv => $"{kv.Key}:{kv.Value.Ticks * 100}"))}]";
        static string DumpStringEnumMap(Dictionary<string, ItemType> m) => m.Count == 0 ? "map[]" : $"map[{string.Join(",", m.Select(kv => $"{kv.Key}:{(int)kv.Value}"))}]";
        static string DumpStringItemDataMap(Dictionary<string, ItemData> m) => m.Count == 0 ? "map[]" : $"map[{string.Join(",", m.Select(kv => $"{kv.Key}:({kv.Value.Id},{kv.Value.Name},{kv.Value.Count})"))}]";

        static string DumpBoolInt32Map(Dictionary<bool, int> m) => m.Count == 0 ? "map[]" : $"map[{string.Join(",", m.Select(kv => $"{kv.Key}:{kv.Value}"))}]";
        static string DumpBoolStringMap(Dictionary<bool, string> m) => m.Count == 0 ? "map[]" : $"map[{string.Join(",", m.Select(kv => $"{kv.Key}:{kv.Value}"))}]";
        static string DumpBoolTimestampMap(Dictionary<bool, DateTime> m) => m.Count == 0 ? "map[]" : $"map[{string.Join(",", m.Select(kv => $"{kv.Key}:{(kv.Value.ToUniversalTime() - new DateTime(1970,1,1,0,0,0,DateTimeKind.Utc)).Ticks * 100}"))}]";
        static string DumpBoolDurationMap(Dictionary<bool, TimeSpan> m) => m.Count == 0 ? "map[]" : $"map[{string.Join(",", m.Select(kv => $"{kv.Key}:{kv.Value.Ticks * 100}"))}]";
        static string DumpBoolEnumMap(Dictionary<bool, ItemType> m) => m.Count == 0 ? "map[]" : $"map[{string.Join(",", m.Select(kv => $"{kv.Key}:{(int)kv.Value}"))}]";
        static string DumpBoolItemDataMap(Dictionary<bool, ItemData> m) => m.Count == 0 ? "map[]" : $"map[{string.Join(",", m.Select(kv => $"{kv.Key}:({kv.Value.Id},{kv.Value.Name},{kv.Value.Count})"))}]";
    }
}
