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

        [UnmanagedCallersOnly(EntryPoint = "to_string", CallConvs = new[] { typeof(CallConvCdecl) })]
        public static IntPtr ToString()
        {
            return Marshal.StringToHGlobalAnsi(_all.ToString(""));
        }
    }
}
