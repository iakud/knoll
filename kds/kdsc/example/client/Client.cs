using System;
using System.IO;
using System.Runtime.InteropServices;

namespace kds
{
    /// <summary>
    /// AOT编译的客户端库
    /// </summary>
    public static class Client
    {
        private static PlayerSync? _player;
        private static string _output = "";

        // 导出函数
        [UnmanagedCallersOnly(EntryPoint = "client_init", CallConvs = new[] { typeof(System.Runtime.CompilerServices.CallConvCdecl) })]
        public static int Init(long playerId)
        {
            try
            {
                _player = new PlayerSync(playerId);
                SetupChangeListener(_player);
                _output = $"Initialized player {playerId}\n";
                return 0;
            }
            catch (Exception ex)
            {
                _output = $"Init error: {ex.Message}\n";
                return 1;
            }
        }

        [UnmanagedCallersOnly(EntryPoint = "client_apply_sync", CallConvs = new[] { typeof(System.Runtime.CompilerServices.CallConvCdecl) })]
        public static int ApplySync(IntPtr dataPtr, int length)
        {
            try
            {
                if (_player == null)
                {
                    _output = "Error: not initialized\n";
                    return 1;
                }

                var data = new byte[length];
                Marshal.Copy(dataPtr, data, 0, length);

                _player.ApplySync(data);
                _player.RaiseChanged();
                _player.ClearChanged();

                _output = $"Applied sync: {length} bytes\n";
                return 0;
            }
            catch (Exception ex)
            {
                _output = $"ApplySync error: {ex.Message}\n";
                return 1;
            }
        }

        [UnmanagedCallersOnly(EntryPoint = "client_apply_sync_file", CallConvs = new[] { typeof(System.Runtime.CompilerServices.CallConvCdecl) })]
        public static int ApplySyncFile(IntPtr filePathPtr)
        {
            try
            {
                string? filePath = Marshal.PtrToStringAnsi(filePathPtr);
                if (string.IsNullOrEmpty(filePath))
                {
                    _output = "Error: empty file path\n";
                    return 1;
                }

                if (!File.Exists(filePath))
                {
                    _output = $"Error: file not found: {filePath}\n";
                    return 1;
                }

                var data = File.ReadAllBytes(filePath);
                _output = $"Read file: {filePath}, {data.Length} bytes\n";

                return ApplySyncInternal(data, data.Length);
            }
            catch (Exception ex)
            {
                _output = $"Error: {ex.Message}\n";
                return 1;
            }
        }

        [UnmanagedCallersOnly(EntryPoint = "client_get_info", CallConvs = new[] { typeof(System.Runtime.CompilerServices.CallConvCdecl) })]
        public static int GetInfo()
        {
            try
            {
                if (_player == null)
                {
                    _output = "Error: not initialized\n";
                    return 1;
                }

                _output += $"PlayerID: {_player.Id}\n";
                _output += $"Name: {_player.Info.Name}\n";
                _output += $"Level: {_player.Info.Level}\n";
                _output += $"IsNew: {_player.Info.IsNew}\n";

                _output += "Currencies:\n";
                foreach (var kvp in _player.Bag.Currencies)
                {
                    _output += $"  ID={kvp.Key}: {kvp.Value}\n";
                }

                _output += "Items:\n";
                foreach (var kvp in _player.Bag.Items)
                {
                    _output += $"  ID={kvp.Key}: {kvp.Value}\n";
                }

                _output += "Heroes:\n";
                foreach (var kvp in _player.Hero.Heroes)
                {
                    _output += $"  ID={kvp.Key}: Lv={kvp.Value.Level} Star={kvp.Value.Star} Exp={kvp.Value.Exp}\n";
                }

                return 0;
            }
            catch (Exception ex)
            {
                _output = $"GetInfo error: {ex.Message}\n";
                return 1;
            }
        }

        [UnmanagedCallersOnly(EntryPoint = "client_get_output", CallConvs = new[] { typeof(System.Runtime.CompilerServices.CallConvCdecl) })]
        public static IntPtr GetOutput()
        {
            return Marshal.StringToHGlobalAnsi(_output);
        }

        [UnmanagedCallersOnly(EntryPoint = "client_free_output", CallConvs = new[] { typeof(System.Runtime.CompilerServices.CallConvCdecl) })]
        public static void FreeOutput(IntPtr ptr)
        {
            if (ptr != IntPtr.Zero)
            {
                Marshal.FreeHGlobal(ptr);
            }
        }

        [UnmanagedCallersOnly(EntryPoint = "client_get_output_len", CallConvs = new[] { typeof(System.Runtime.CompilerServices.CallConvCdecl) })]
        public static int GetOutputLen()
        {
            return _output.Length;
        }

        [UnmanagedCallersOnly(EntryPoint = "client_clear_output", CallConvs = new[] { typeof(System.Runtime.CompilerServices.CallConvCdecl) })]
        public static void ClearOutput()
        {
            _output = "";
        }

        // 内部方法，调用ApplySyncInternal
        static int ApplySyncInternal(byte[] data, int length)
        {
            try
            {
                if (_player == null)
                {
                    _output = "Error: not initialized\n";
                    return 1;
                }

                _player.ApplySync(data);
                _player.RaiseChanged();
                _player.ClearChanged();

                _output += $"Applied sync: {length} bytes\n";
                return 0;
            }
            catch (Exception ex)
            {
                _output = $"ApplySync error: {ex.Message}\n";
                return 1;
            }
        }

        static void SetupChangeListener(PlayerSync player)
        {
            player.OnChanged += (sender, e) =>
            {
                _output += "[Event] Player changed!\n";
                if (e.Info) _output += "  - Info changed\n";
                if (e.Hero) _output += "  - Hero changed\n";
                if (e.Bag) _output += "  - Bag changed\n";
            };

            player.Info.OnChanged += (sender, e) =>
            {
                _output += "[Event] PlayerInfo changed:\n";
                if (e.Name) _output += $"  - Name: {player.Info.Name}\n";
                if (e.Level) _output += $"  - Level: {player.Info.Level}\n";
                if (e.IsNew) _output += $"  - IsNew: {player.Info.IsNew}\n";
            };

            player.Hero.OnChanged += (sender, e) =>
            {
                _output += "[Event] PlayerHero changed:\n";
                if (e.Heroes)
                {
                    foreach (var kvp in player.Hero.Heroes)
                    {
                        _output += $"  ID={kvp.Key}: Lv={kvp.Value.Level}\n";
                    }
                }
            };

            player.Bag.OnChanged += (sender, e) =>
            {
                _output += "[Event] PlayerBag changed:\n";
                if (e.Items)
                {
                    foreach (var kvp in player.Bag.Items)
                    {
                        _output += $"  Item ID={kvp.Key}: {kvp.Value}\n";
                    }
                }
                if (e.Currencies)
                {
                    foreach (var kvp in player.Bag.Currencies)
                    {
                        _output += $"  Currency ID={kvp.Key}: {kvp.Value}\n";
                    }
                }
            };
        }
    }
}
