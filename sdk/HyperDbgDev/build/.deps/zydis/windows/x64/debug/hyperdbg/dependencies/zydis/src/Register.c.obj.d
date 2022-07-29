{
    files = {
        [[hyperdbg\dependencies\zydis\src\Register.c]]
    },
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
    },
    depfiles_cl_json = "{\
    \"Version\": \"1.2\",\
    \"Data\": {\
        \"Source\": \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\dependencies\\\\zydis\\\\src\\\\register.c\",\
        \"ProvidedModule\": \"\",\
        \"Includes\": [\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\dependencies\\\\zydis\\\\include\\\\zydis\\\\register.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\dependencies\\\\zydis\\\\dependencies\\\\zycore\\\\include\\\\zycore\\\\defines.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\ucrt\\\\assert.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\ucrt\\\\corecrt.h\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\vcruntime.h\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\sal.h\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\concurrencysal.h\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\vadefs.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\dependencies\\\\zydis\\\\dependencies\\\\zycore\\\\include\\\\zycore\\\\types.h\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\stdint.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\ucrt\\\\stddef.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\dependencies\\\\zydis\\\\include\\\\zydis\\\\sharedtypes.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\dependencies\\\\zydis\\\\include\\\\zydis\\\\shortstring.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\include\\\\zydisexportconfig.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\dependencies\\\\zydis\\\\include\\\\zydis\\\\generated\\\\enumregister.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\dependencies\\\\zydis\\\\src\\\\generated\\\\enumregister.inc\"\
        ]\
    }\
}"
}