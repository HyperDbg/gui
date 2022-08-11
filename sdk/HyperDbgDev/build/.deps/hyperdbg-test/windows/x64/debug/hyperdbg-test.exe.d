{
    files = {
        [[build\.objs\hyperdbg-test\windows\x64\debug\hyperdbg\hyperdbg-test\code\assembly\asm-test.asm.obj]],
        [[build\.objs\hyperdbg-test\windows\x64\debug\hyperdbg\hyperdbg-test\code\tests\hyperdbg-test.cpp.obj]],
        [[build\.objs\hyperdbg-test\windows\x64\debug\hyperdbg\hyperdbg-test\code\tests\lookup.cpp.obj]],
        [[build\.objs\hyperdbg-test\windows\x64\debug\hyperdbg\hyperdbg-test\code\tests\namedpipe.cpp.obj]],
        [[build\.objs\hyperdbg-test\windows\x64\debug\hyperdbg\hyperdbg-test\code\tests\test-cases.cpp.obj]],
        [[build\.objs\hyperdbg-test\windows\x64\debug\hyperdbg\hyperdbg-test\code\tests\tools.cpp.obj]],
        [[build\windows\x64\debug\zycore.lib]],
        [[build\windows\x64\debug\zydis.lib]],
        [[build\windows\x64\debug\pdbex.lib]],
        [[build\windows\x64\debug\symbol-parser.lib]],
        [[build\windows\x64\debug\script-engine.lib]],
        [[build\windows\x64\debug\hprdbgctrl.lib]]
    },
    values = {
        [[C:\Program Files\Microsoft Visual Studio\2022\Enterprise\VC\Tools\MSVC\14.32.31326\bin\HostX64\x64\link.exe]],
        {
            "-nologo",
            "-dynamicbase",
            "-nxcompat",
            "-machine:x64",
            [[-libpath:build\windows\x64\debug]],
            "-debug",
            [[-pdb:build\windows\x64\debug\hyperdbg-test.pdb]],
            "hprdbgctrl.lib",
            "script-engine.lib",
            "symbol-parser.lib",
            "user32.lib",
            "pdbex.lib",
            "zydis.lib",
            "zycore.lib"
        }
    }
}