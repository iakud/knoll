using System.Runtime.InteropServices;

namespace Kds;

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
        _all.Maps.Int32Int32.OnChanged += (sender, e) => Console.Out.WriteLine($"All.Maps.Int32Int32.Changed");
        _all.Maps.Int32String.OnChanged += (sender, e) => Console.Out.WriteLine($"All.Maps.Int32String.Changed");
        _all.Maps.Int32Timestamp.OnChanged += (sender, e) => Console.Out.WriteLine($"All.Maps.Int32Timestamp.Changed");
        _all.Maps.Int32Duration.OnChanged += (sender, e) => Console.Out.WriteLine($"All.Maps.Int32Duration.Changed");
        _all.Maps.Int32Empty.OnChanged += (sender, e) => Console.Out.WriteLine($"All.Maps.Int32Empty.Changed");
        _all.Maps.Int32Enum.OnChanged += (sender, e) => Console.Out.WriteLine($"All.Maps.Int32Enum.Changed");
        _all.Maps.Int32ItemData.OnChanged += (sender, e) => Console.Out.WriteLine($"All.Maps.Int32ItemData.Changed");
        _all.Maps.Int64Int64.OnChanged += (sender, e) => Console.Out.WriteLine($"All.Maps.Int64Int64.Changed");
        _all.Maps.Int64String.OnChanged += (sender, e) => Console.Out.WriteLine($"All.Maps.Int64String.Changed");
        _all.Maps.Int64Timestamp.OnChanged += (sender, e) => Console.Out.WriteLine($"All.Maps.Int64Timestamp.Changed");
        _all.Maps.Int64Duration.OnChanged += (sender, e) => Console.Out.WriteLine($"All.Maps.Int64Duration.Changed");
        _all.Maps.Int64Empty.OnChanged += (sender, e) => Console.Out.WriteLine($"All.Maps.Int64Empty.Changed");
        _all.Maps.Int64Enum.OnChanged += (sender, e) => Console.Out.WriteLine($"All.Maps.Int64Enum.Changed");
        _all.Maps.Int64ItemData.OnChanged += (sender, e) => Console.Out.WriteLine($"All.Maps.Int64ItemData.Changed");
        _all.Maps.StringInt32.OnChanged += (sender, e) => Console.Out.WriteLine($"All.Maps.StringInt32.Changed");
        _all.Maps.StringString.OnChanged += (sender, e) => Console.Out.WriteLine($"All.Maps.StringString.Changed");
        _all.Maps.StringTimestamp.OnChanged += (sender, e) => Console.Out.WriteLine($"All.Maps.StringTimestamp.Changed");
        _all.Maps.StringDuration.OnChanged += (sender, e) => Console.Out.WriteLine($"All.Maps.StringDuration.Changed");
        _all.Maps.StringEmpty.OnChanged += (sender, e) => Console.Out.WriteLine($"All.Maps.StringEmpty.Changed");
        _all.Maps.StringEnum.OnChanged += (sender, e) => Console.Out.WriteLine($"All.Maps.StringEnum.Changed");
        _all.Maps.StringItemData.OnChanged += (sender, e) => Console.Out.WriteLine($"All.Maps.StringItemData.Changed");
        _all.Maps.BoolInt32.OnChanged += (sender, e) => Console.Out.WriteLine($"All.Maps.BoolInt32.Changed");
        _all.Maps.BoolString.OnChanged += (sender, e) => Console.Out.WriteLine($"All.Maps.BoolString.Changed");
        _all.Maps.BoolTimestamp.OnChanged += (sender, e) => Console.Out.WriteLine($"All.Maps.BoolTimestamp.Changed");
        _all.Maps.BoolDuration.OnChanged += (sender, e) => Console.Out.WriteLine($"All.Maps.BoolDuration.Changed");
        _all.Maps.BoolEmpty.OnChanged += (sender, e) => Console.Out.WriteLine($"All.Maps.BoolEmpty.Changed");
        _all.Maps.BoolEnum.OnChanged += (sender, e) => Console.Out.WriteLine($"All.Maps.BoolEnum.Changed");
        _all.Maps.BoolItemData.OnChanged += (sender, e) => Console.Out.WriteLine($"All.Maps.BoolItemData.Changed");
    }

#if NET10_0_OR_GREATER
    [System.Runtime.InteropServices.UnmanagedCallersOnly(EntryPoint = "merge_from", CallConvs = new[] { typeof(System.Runtime.CompilerServices.CallConvCdecl) })]
#endif
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

#if NET10_0_OR_GREATER
    [System.Runtime.InteropServices.UnmanagedCallersOnly(EntryPoint = "get_string", CallConvs = new[] { typeof(System.Runtime.CompilerServices.CallConvCdecl) })]
#endif
    public static IntPtr GetString()
    {
        return Marshal.StringToHGlobalAnsi(_all.ToString(""));
    }
}
