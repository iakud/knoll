using System.Runtime.CompilerServices;
using System.Runtime.InteropServices;
using Google.Protobuf.WellKnownTypes;

namespace kds
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
            _all.Types.OnChanged += (sender, e) => Console.Out.WriteLine($"All.Types.Changed");
            _all.Types.ItemData.OnChanged += (sender, e) => Console.Out.WriteLine($"All.Types.ItemData.Changed");
            _all.Lists.OnChanged += (sender, e) =>Console.Out.WriteLine($"All.Lists.Changed");
            _all.Maps.OnChanged += (sender, e) => Console.Out.WriteLine($"All.Maps.Changed");
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

        [UnmanagedCallersOnly(EntryPoint = "get_string", CallConvs = new[] { typeof(CallConvCdecl) })]
        public static IntPtr GetString()
        {
            return Marshal.StringToHGlobalAnsi(_all.ToString(""));
        }
    }
}
