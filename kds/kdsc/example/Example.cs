using System;
using System.IO;
using System.Reflection.Metadata.Ecma335;
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

        [UnmanagedCallersOnly(EntryPoint = "apply_sync", CallConvs = new[] { typeof(System.Runtime.CompilerServices.CallConvCdecl) })]
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

        static void SetupChangeListener(All all)
        {
            // FIXME: fields onchanged


            all.OnChanged += (sender, e) => Console.Out.WriteLine($"All.Changed\n");
        }
    }
}
