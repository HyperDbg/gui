{
    files = {
        [[hyperdbg\dependencies\zydis\src\FormatterATT.c]]
    },
    depfiles_cl_json = "{\
    \"Version\": \"1.2\",\
    \"Data\": {\
        \"Source\": \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\dependencies\\\\zydis\\\\src\\\\formatteratt.c\",\
        \"ProvidedModule\": \"\",\
        \"Includes\": [\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\dependencies\\\\zydis\\\\include\\\\zydis\\\\internal\\\\formatteratt.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\dependencies\\\\zydis\\\\include\\\\zydis\\\\formatter.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\dependencies\\\\zydis\\\\dependencies\\\\zycore\\\\include\\\\zycore\\\\defines.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\ucrt\\\\assert.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\ucrt\\\\corecrt.h\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\vcruntime.h\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\sal.h\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\concurrencysal.h\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\vadefs.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\dependencies\\\\zydis\\\\dependencies\\\\zycore\\\\include\\\\zycore\\\\string.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\include\\\\zycoreexportconfig.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\dependencies\\\\zydis\\\\dependencies\\\\zycore\\\\include\\\\zycore\\\\allocator.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\dependencies\\\\zydis\\\\dependencies\\\\zycore\\\\include\\\\zycore\\\\status.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\dependencies\\\\zydis\\\\dependencies\\\\zycore\\\\include\\\\zycore\\\\types.h\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\stdint.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\ucrt\\\\stddef.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\dependencies\\\\zydis\\\\dependencies\\\\zycore\\\\include\\\\zycore\\\\vector.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\dependencies\\\\zydis\\\\dependencies\\\\zycore\\\\include\\\\zycore\\\\comparison.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\dependencies\\\\zydis\\\\dependencies\\\\zycore\\\\include\\\\zycore\\\\object.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\dependencies\\\\zydis\\\\include\\\\zydis\\\\decodertypes.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\dependencies\\\\zydis\\\\include\\\\zydis\\\\metainfo.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\include\\\\zydisexportconfig.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\dependencies\\\\zydis\\\\include\\\\zydis\\\\generated\\\\enuminstructioncategory.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\dependencies\\\\zydis\\\\include\\\\zydis\\\\generated\\\\enumisaset.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\dependencies\\\\zydis\\\\include\\\\zydis\\\\generated\\\\enumisaext.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\dependencies\\\\zydis\\\\include\\\\zydis\\\\mnemonic.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\dependencies\\\\zydis\\\\include\\\\zydis\\\\shortstring.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\include\\\\zydisexportconfig.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\dependencies\\\\zydis\\\\include\\\\zydis\\\\generated\\\\enummnemonic.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\dependencies\\\\zydis\\\\include\\\\zydis\\\\register.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\dependencies\\\\zydis\\\\include\\\\zydis\\\\sharedtypes.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\dependencies\\\\zydis\\\\include\\\\zydis\\\\generated\\\\enumregister.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\dependencies\\\\zydis\\\\include\\\\zydis\\\\formatterbuffer.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\include\\\\zydisexportconfig.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\dependencies\\\\zydis\\\\include\\\\zydis\\\\status.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\dependencies\\\\zydis\\\\include\\\\zydis\\\\internal\\\\formatterbase.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\dependencies\\\\zydis\\\\include\\\\zydis\\\\internal\\\\string.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\dependencies\\\\zydis\\\\dependencies\\\\zycore\\\\include\\\\zycore\\\\libc.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\ucrt\\\\errno.h\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\stdarg.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\ucrt\\\\stdio.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\ucrt\\\\corecrt_wstdio.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\ucrt\\\\corecrt_stdio_config.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\ucrt\\\\stdlib.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\ucrt\\\\corecrt_malloc.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\ucrt\\\\corecrt_search.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\ucrt\\\\corecrt_wstdlib.h\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\limits.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\ucrt\\\\string.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\ucrt\\\\corecrt_memory.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\ucrt\\\\corecrt_memcpy_s.h\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\vcruntime_string.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\ucrt\\\\corecrt_wstring.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\dependencies\\\\zydis\\\\src\\\\generated\\\\formatterstrings.inc\"\
        ]\
    }\
}",
    values = {
        [[C:\Program Files\Microsoft Visual Studio\2022\Enterprise\VC\Tools\MSVC\14.32.31326\bin\HostX64\x64\cl.exe]],
        {
            "-nologo",
            "-Zi",
            "-FS",
            [[-Fdbuild\windows\x64\debug\zydis.pdb]],
            "-Od",
            [[-Ihyperdbg\dependencies\zydis\src]],
            [[-Ihyperdbg\dependencies\zydis\include]],
            [[-Ihyperdbg\dependencies\zydis\dependencies\zycore\include]],
            [[-Ihyperdbg\include]],
            "-DZYDIS_STATIC_DEFINE"
        }
    }
}