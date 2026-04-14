using System;
using System.Runtime.CompilerServices;
using System.Runtime.InteropServices;
using Kdsync;

namespace Kds
{
    public static class Example
    {
        private static All _all = new All(0);

        static Example()
        {
            Initialize();
            Console.Out.WriteLine($"Initialized");
        }

        public static void Initialize()
        {
            _all.OnChanged += (sender, e) => Console.Out.WriteLine($"All.Changed");
            // types
            _all.Types.OnChanged += (sender, e) => Console.Out.WriteLine($"All.Types.Changed");
            _all.Types.ItemData.OnChanged += (sender, e) => Console.Out.WriteLine($"All.Types.ItemData.Changed");
            // lists
            _all.Lists.OnChanged += (sender, e) => Console.Out.WriteLine($"All.Lists.Changed");
            _all.Lists.Int32List.OnChanged += (sender, e) => Console.Out.WriteLine($"All.Lists.Int32List.Changed");
            _all.Lists.Int64List.OnChanged += (sender, e) => Console.Out.WriteLine($"All.Lists.Int64List.Changed");
            _all.Lists.FloatList.OnChanged += (sender, e) => Console.Out.WriteLine($"All.Lists.FloatList.Changed");
            _all.Lists.DoubleList.OnChanged += (sender, e) => Console.Out.WriteLine($"All.Lists.DoubleList.Changed");
            _all.Lists.BoolList.OnChanged += (sender, e) => Console.Out.WriteLine($"All.Lists.BoolList.Changed");
            _all.Lists.StringList.OnChanged += (sender, e) => Console.Out.WriteLine($"All.Lists.StringList.Changed");
            _all.Lists.TimestampList.OnChanged += (sender, e) => Console.Out.WriteLine($"All.Lists.TimestampList.Changed");
            _all.Lists.DurationList.OnChanged += (sender, e) => Console.Out.WriteLine($"All.Lists.DurationList.Changed");
            _all.Lists.EmptyList.OnChanged += (sender, e) => Console.Out.WriteLine($"All.Lists.EmptyList.Changed");
            _all.Lists.EnumList.OnChanged += (sender, e) => Console.Out.WriteLine($"All.Lists.EnumList.Changed");
            _all.Lists.ItemList.OnChanged += (sender, e) => Console.Out.WriteLine($"All.Lists.ItemList.Changed");
            // maps
            _all.Maps.OnChanged += (sender, e) => Console.Out.WriteLine($"All.Maps.Changed");
            _all.Maps.Int32Int32Map.OnChanged += (sender, e) => Console.Out.WriteLine($"All.Maps.Int32Int32Map.Changed");
            _all.Maps.Int64Int64Map.OnChanged += (sender, e) => Console.Out.WriteLine($"All.Maps.Int64Int64Map.Changed");
            _all.Maps.Uint32Uint32Map.OnChanged += (sender, e) => Console.Out.WriteLine($"All.Maps.Uint32Uint32Map.Changed");
            _all.Maps.Uint64Uint64Map.OnChanged += (sender, e) => Console.Out.WriteLine($"All.Maps.Uint64Uint64Map.Changed");
            _all.Maps.BoolFloatMap.OnChanged += (sender, e) => Console.Out.WriteLine($"All.Maps.BoolFloatMap.Changed");
            _all.Maps.StringDoubleMap.OnChanged += (sender, e) => Console.Out.WriteLine($"All.Maps.StringDoubleMap.Changed");
            _all.Maps.Int32BoolMap.OnChanged += (sender, e) => Console.Out.WriteLine($"All.Maps.Int32BoolMap.Changed");
            _all.Maps.Int64StringMap.OnChanged += (sender, e) => Console.Out.WriteLine($"All.Maps.Int64StringMap.Changed");
            _all.Maps.Uint32BytesMap.OnChanged += (sender, e) => Console.Out.WriteLine($"All.Maps.Uint32BytesMap.Changed");
            _all.Maps.Uint64TimestampMap.OnChanged += (sender, e) => Console.Out.WriteLine($"All.Maps.Uint64TimestampMap.Changed");
            _all.Maps.BoolDurationMap.OnChanged += (sender, e) => Console.Out.WriteLine($"All.Maps.BoolDurationMap.Changed");
            _all.Maps.StringEmptyMap.OnChanged += (sender, e) => Console.Out.WriteLine($"All.Maps.StringEmptyMap.Changed");
            _all.Maps.Int32ItemTypeMap.OnChanged += (sender, e) => Console.Out.WriteLine($"All.Maps.Int32ItemTypeMap.Changed");
            _all.Maps.Int64ItemDataMap.OnChanged += (sender, e) => Console.Out.WriteLine($"All.Maps.Int64ItemDataMap.Changed");
        }

        [UnmanagedCallersOnly(EntryPoint = "merge_from", CallConvs = new[] { typeof(CallConvCdecl) })]
        public static int MergeFrom(IntPtr dataPtr, int length)
        {
            try
            {
                var data = new byte[length];
                Marshal.Copy(dataPtr, data, 0, length);

                _all.MergeFrom(data);
                _all.RaiseChanged();
                _all.ClearChanged();

                Console.Out.WriteLine($"MergeFrom: {length} bytes");
                return 0;
            }
            catch (Exception ex)
            {
                Console.Out.WriteLine($"MergeFrom error: {ex.Message}, stack: {ex.StackTrace}");
                return 1;
            }
        }

        [UnmanagedCallersOnly(EntryPoint = "get_string", CallConvs = new[] { typeof(CallConvCdecl) })]
        public static IntPtr GetString()
        {
            return Marshal.StringToHGlobalAnsi(_all.ToString());
        }
    }
}