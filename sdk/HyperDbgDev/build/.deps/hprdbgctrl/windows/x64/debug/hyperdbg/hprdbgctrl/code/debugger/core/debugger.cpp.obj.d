{
    files = {
        [[hyperdbg\hprdbgctrl\code\debugger\core\debugger.cpp]]
    },
    values = {
        [[C:\Program Files\Microsoft Visual Studio\2022\Enterprise\VC\Tools\MSVC\14.32.31326\bin\HostX64\x64\cl.exe]],
        {
            "/EHsc",
            "-nologo",
            "-Zi",
            "-FS",
            [[-Fdbuild\windows\x64\debug\hprdbgctrl.pdb]],
            "-Od",
            [[-Ihyperdbg\hprdbgctrl]],
            [[-Ihyperdbg\hprdbgctrl\header]],
            [[-Ihyperdbg\dependencies\phnt]],
            [[-Ihyperdbg\script-engine\header]],
            [[-Ihyperdbg\script-engine]],
            [[-Ihyperdbg\symbol-parser\header]],
            [[-Ihyperdbg\symbol-parser]],
            [[-Ihyperdbg\dependencies]],
            [[-Ihyperdbg\dependencies\pdbex\Source]],
            [[-IC:\Program Files\Microsoft Visual Studio\2022\Enterprise\DIA SDK\include]],
            [[-Ihyperdbg\dependencies\zydis\src]],
            [[-Ihyperdbg\dependencies\zydis\include]],
            [[-Ihyperdbg\dependencies\zydis\dependencies\zycore\include]],
            [[-Ihyperdbg\include]]
        }
    },
    depfiles_cl_json = "{\
    \"Version\": \"1.2\",\
    \"Data\": {\
        \"Source\": \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\hprdbgctrl\\\\code\\\\debugger\\\\core\\\\debugger.cpp\",\
        \"ProvidedModule\": \"\",\
        \"Includes\": [\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\hprdbgctrl\\\\pch.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\dependencies\\\\ia32-doc\\\\out\\\\ia32.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\dependencies\\\\phnt\\\\phnt_windows.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\um\\\\windows.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\winapifamily.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\winpackagefamily.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\sdkddkver.h\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\excpt.h\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\vcruntime.h\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\sal.h\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\concurrencysal.h\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\vadefs.h\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\stdarg.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\windef.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\minwindef.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\specstrings.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\specstrings_strict.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\specstrings_undef.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\driverspecs.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\sdv_driverspecs.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\um\\\\winnt.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\ucrt\\\\ctype.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\ucrt\\\\corecrt.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\ucrt\\\\corecrt_wctype.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\kernelspecs.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\basetsd.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\guiddef.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\ucrt\\\\string.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\ucrt\\\\corecrt_memory.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\ucrt\\\\corecrt_memcpy_s.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\ucrt\\\\errno.h\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\vcruntime_string.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\ucrt\\\\corecrt_wstring.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\pshpack4.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\poppack.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\pshpack4.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\poppack.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\pshpack4.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\pshpack2.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\poppack.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\pshpack2.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\pshpack2.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\poppack.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\poppack.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\pshpack8.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\poppack.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\pshpack1.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\poppack.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\pshpack1.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\poppack.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\poppack.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\apiset.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\ktmtypes.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\um\\\\winbase.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\apisetcconv.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\um\\\\minwinbase.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\um\\\\apiquery2.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\um\\\\processenv.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\um\\\\fileapifromapp.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\um\\\\fileapi.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\um\\\\debugapi.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\um\\\\utilapiset.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\um\\\\handleapi.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\um\\\\errhandlingapi.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\um\\\\fibersapi.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\um\\\\namedpipeapi.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\um\\\\profileapi.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\um\\\\heapapi.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\um\\\\ioapiset.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\um\\\\synchapi.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\um\\\\interlockedapi.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\um\\\\processthreadsapi.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\um\\\\sysinfoapi.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\um\\\\memoryapi.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\um\\\\enclaveapi.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\um\\\\threadpoollegacyapiset.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\um\\\\threadpoolapiset.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\um\\\\jobapi.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\um\\\\jobapi2.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\um\\\\wow64apiset.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\um\\\\libloaderapi.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\um\\\\securitybaseapi.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\um\\\\namespaceapi.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\um\\\\systemtopologyapi.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\um\\\\processtopologyapi.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\um\\\\securityappcontainer.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\um\\\\realtimeapiset.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\winerror.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\um\\\\timezoneapi.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\um\\\\wingdi.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\pshpack1.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\poppack.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\pshpack2.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\poppack.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\pshpack2.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\poppack.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\pshpack4.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\poppack.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\pshpack4.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\poppack.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\um\\\\winuser.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\pshpack2.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\poppack.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\tvout.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\um\\\\winnls.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\um\\\\datetimeapi.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\um\\\\stringapiset.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\um\\\\winnls.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\um\\\\wincon.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\um\\\\wincontypes.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\um\\\\consoleapi.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\um\\\\consoleapi2.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\um\\\\consoleapi3.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\um\\\\winver.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\um\\\\verrsrc.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\um\\\\winreg.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\um\\\\reason.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\um\\\\winnetwk.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\wnnc.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\stralign.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\um\\\\winsvc.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\um\\\\mcx.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\um\\\\imm.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\um\\\\ime_cmodes.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\windowsx.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\ntstatus.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\um\\\\winioctl.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\pshpack1.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\poppack.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\pshpack1.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\poppack.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\pshpack1.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\poppack.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\pshpack1.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\poppack.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\pshpack1.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\poppack.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\pshpack1.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\poppack.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\pshpack1.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\poppack.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\dependencies\\\\phnt\\\\phnt.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\dependencies\\\\phnt\\\\phnt_ntdef.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\pshpack4.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\poppack.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\dependencies\\\\phnt\\\\ntnls.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\dependencies\\\\phnt\\\\ntkeapi.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\dependencies\\\\phnt\\\\ntldr.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\dependencies\\\\phnt\\\\ntexapi.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\pshpack4.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\poppack.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\dependencies\\\\phnt\\\\ntbcd.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\dependencies\\\\phnt\\\\ntmmapi.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\dependencies\\\\phnt\\\\ntobapi.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\dependencies\\\\phnt\\\\ntpsapi.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\dependencies\\\\phnt\\\\ntpebteb.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\cfg.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\dependencies\\\\phnt\\\\ntdbg.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\dependencies\\\\phnt\\\\ntioapi.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\dependencies\\\\phnt\\\\ntlpcapi.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\dependencies\\\\phnt\\\\ntpfapi.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\dependencies\\\\phnt\\\\ntpnpapi.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\dependencies\\\\phnt\\\\ntpoapi.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\dependencies\\\\phnt\\\\ntregapi.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\dependencies\\\\phnt\\\\ntrtl.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\dependencies\\\\phnt\\\\ntseapi.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\dependencies\\\\phnt\\\\nttmapi.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\dependencies\\\\phnt\\\\nttp.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\dependencies\\\\phnt\\\\ntxcapi.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\dependencies\\\\phnt\\\\ntwow64.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\um\\\\ntlsa.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\um\\\\lsalookup.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\dependencies\\\\phnt\\\\ntsam.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\pshpack4.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\poppack.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\pshpack4.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\poppack.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\pshpack4.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\poppack.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\pshpack4.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\poppack.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\pshpack4.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\poppack.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\pshpack4.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\poppack.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\dependencies\\\\phnt\\\\ntmisc.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\dependencies\\\\phnt\\\\ntzwapi.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\um\\\\winsock2.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\ws2def.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\inaddr.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\qos.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\um\\\\ws2tcpip.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\ws2ipdef.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\in6addr.h\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\limits.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\strsafe.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\ucrt\\\\stdio.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\ucrt\\\\corecrt_wstdio.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\ucrt\\\\corecrt_stdio_config.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\um\\\\shlobj.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\um\\\\ole2.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\pshpack8.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\um\\\\objbase.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\rpc.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\pshpack8.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\rpcdce.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\rpcdcep.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\um\\\\rpcnsi.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\rpcnterr.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\rpcasync.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\pshpack8.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\poppack.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\poppack.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\rpcndr.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\pshpack8.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\um\\\\rpcnsip.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\rpcsal.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\poppack.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\pshpack8.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\um\\\\combaseapi.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\pshpack8.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\ucrt\\\\stdlib.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\ucrt\\\\corecrt_malloc.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\ucrt\\\\corecrt_search.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\ucrt\\\\stddef.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\ucrt\\\\corecrt_wstdlib.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\wtypesbase.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\guiddef.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\um\\\\unknwnbase.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\um\\\\objidlbase.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\guiddef.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\poppack.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\um\\\\coml2api.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\um\\\\objidl.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\um\\\\unknwn.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\wtypes.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\um\\\\propidlbase.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\um\\\\oaidl.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\um\\\\urlmon.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\um\\\\oleidl.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\um\\\\servprov.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\um\\\\msxml.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\um\\\\propidl.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\poppack.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\um\\\\oleauto.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\pshpack8.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\poppack.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\poppack.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\um\\\\prsht.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\pshpack8.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\poppack.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\um\\\\commctrl.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\um\\\\prsht.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\pshpack1.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\poppack.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\pshpack1.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\poppack.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\pshpack1.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\poppack.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\pshpack1.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\poppack.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\pshpack1.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\poppack.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\um\\\\dpa_dsa.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\um\\\\shtypes.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\pshpack1.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\poppack.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\pshpack1.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\poppack.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\pshpack8.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\poppack.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\pshpack1.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\poppack.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\um\\\\shobjidl.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\um\\\\docobj.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\um\\\\ocidl.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\um\\\\comcat.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\um\\\\prsht.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\um\\\\propsys.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\um\\\\structuredquerycondition.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\um\\\\propkeydef.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\um\\\\objectarray.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\um\\\\shobjidl_core.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\um\\\\sherrors.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\pshpack8.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\poppack.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\pshpack8.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\poppack.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\um\\\\prsht.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\pshpack8.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\poppack.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\pshpack8.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\poppack.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\pshpack8.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\poppack.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\pshpack8.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\poppack.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\pshpack8.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\poppack.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\pshpack1.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\poppack.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\pshpack8.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\poppack.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\um\\\\shlobj_core.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\pshpack1.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\pshpack8.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\poppack.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\pshpack8.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\poppack.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\pshpack8.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\poppack.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\pshpack1.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\poppack.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\pshpack8.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\poppack.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\pshpack8.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\poppack.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\pshpack8.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\poppack.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\poppack.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\pshpack1.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\pshpack8.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\poppack.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\pshpack1.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\poppack.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\pshpack8.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\poppack.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\pshpack8.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\poppack.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\pshpack8.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\poppack.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\pshpack8.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\poppack.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\poppack.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\ucrt\\\\tchar.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\um\\\\tlhelp32.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\um\\\\shlwapi.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\pshpack8.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\poppack.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\um\\\\versionhelpers.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\um\\\\psapi.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\ucrt\\\\time.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\ucrt\\\\corecrt_wtime.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\ucrt\\\\conio.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\ucrt\\\\corecrt_wconio.h\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\intrin.h\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\intrin0.inl.h\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\setjmp.h\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\immintrin.h\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\wmmintrin.h\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\nmmintrin.h\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\smmintrin.h\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\tmmintrin.h\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\pmmintrin.h\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\emmintrin.h\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\xmmintrin.h\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\mmintrin.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\ucrt\\\\malloc.h\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\zmmintrin.h\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\ammintrin.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\ucrt\\\\inttypes.h\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\stdint.h\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\algorithm\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\yvals_core.h\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\xkeycheck.h\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\xmemory\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\cstdint\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\cstdlib\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\ucrt\\\\math.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\ucrt\\\\corecrt_math.h\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\limits\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\cfloat\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\ucrt\\\\float.h\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\climits\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\cwchar\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\cstdio\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\ucrt\\\\wchar.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\ucrt\\\\corecrt_wdirect.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\ucrt\\\\corecrt_wio.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\ucrt\\\\corecrt_share.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\ucrt\\\\corecrt_wprocess.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\ucrt\\\\sys\\\\stat.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\ucrt\\\\sys\\\\types.h\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\intrin0.h\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\isa_availability.h\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\xstddef\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\cstddef\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\xtr1common\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\initializer_list\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\new\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\exception\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\yvals.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\ucrt\\\\crtdbg.h\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\vcruntime_new_debug.h\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\vcruntime_new.h\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\crtdefs.h\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\use_ansi.h\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\type_traits\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\vcruntime_exception.h\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\eh.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\ucrt\\\\corecrt_terminate.h\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\xatomic.h\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\xutility\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\cstring\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\utility\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\string\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\xstring\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\iosfwd\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\cctype\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\vector\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\array\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\bitset\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\iomanip\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\istream\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\ostream\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\ios\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\xlocnum\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\cmath\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\iterator\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\streambuf\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\xiosbase\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\ucrt\\\\share.h\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\system_error\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\__msvc_system_error_abi.hpp\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\cerrno\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\stdexcept\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\xcall_once.h\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\xerrc.h\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\atomic\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\xthreads.h\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\xtimec.h\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\ctime\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\xlocale\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\memory\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\typeinfo\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\vcruntime_typeinfo.h\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\xfacet\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\xlocinfo\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\__msvc_xlocinfo_types.hpp\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\clocale\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\ucrt\\\\locale.h\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\xlocmon\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\xloctime\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\iostream\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\sstream\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\fstream\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\map\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\tuple\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\xtree\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\numeric\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\list\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\locale\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\xlocbuf\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\xlocmes\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\include\\\\sdk\\\\hyperdbgsdk.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\include\\\\sdk\\\\headers\\\\constants.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\include\\\\sdk\\\\headers\\\\basictypes.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\include\\\\sdk\\\\headers\\\\errorcodes.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\include\\\\sdk\\\\headers\\\\connection.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\include\\\\sdk\\\\headers\\\\datatypes.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\include\\\\sdk\\\\headers\\\\ioctls.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\include\\\\sdk\\\\headers\\\\events.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\include\\\\sdk\\\\headers\\\\requeststructures.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\include\\\\sdk\\\\headers\\\\symbols.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\include\\\\sdk\\\\imports\\\\hyperdbgctrlimports.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\include\\\\configuration.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\include\\\\definition.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\script-eval\\\\header\\\\scriptenginecommondefinitions.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\script-eval\\\\header\\\\scriptengineheader.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\include\\\\sdk\\\\imports\\\\hyperdbgscriptimports.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\hprdbgctrl\\\\header\\\\inipp.h\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\functional\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\hprdbgctrl\\\\header\\\\commands.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\hprdbgctrl\\\\header\\\\common.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\hprdbgctrl\\\\header\\\\symbol.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\hprdbgctrl\\\\header\\\\debugger.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\hprdbgctrl\\\\header\\\\script-engine.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\hprdbgctrl\\\\header\\\\exports.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\hprdbgctrl\\\\header\\\\help.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\hprdbgctrl\\\\header\\\\install.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\hprdbgctrl\\\\header\\\\list.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\hprdbgctrl\\\\header\\\\tests.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\hprdbgctrl\\\\header\\\\transparency.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\hprdbgctrl\\\\header\\\\communication.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\hprdbgctrl\\\\header\\\\namedpipe.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\hprdbgctrl\\\\header\\\\forwarding.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\hprdbgctrl\\\\header\\\\kd.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\hprdbgctrl\\\\header\\\\pe-parser.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\hprdbgctrl\\\\header\\\\ud.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\hprdbgctrl\\\\header\\\\objects.h\"\
        ],\
        \"ImportedModules\": [],\
        \"ImportedHeaderUnits\": []\
    }\
}"
}