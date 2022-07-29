{
    files = {
        [[hyperdbg\hprdbghv\code\assembly\AsmEpt.asm]]
    },
    values = {
        [[C:\Program Files\Microsoft Visual Studio\2022\Enterprise\VC\Tools\MSVC\14.32.31326\bin\HostX64\x64\ml64.exe]],
        {
            "-nologo",
            "-Zi",
            [[-Ihyperdbg\hprdbghv\header]],
            [[-Ihyperdbg\script-eval\header]],
            [[-Ihyperdbg\script-eval]],
            [[-Ihyperdbg\hprdbghv]],
            [[-Ihyperdbg\dependencies]],
            [[-Ihyperdbg\include]],
            [[-IC:\Program Files (x86)\Windows Kits\10\Include\10.0.22621.0\km]],
            [[-IC:\Program Files (x86)\Windows Kits\10\Include\10.0.22621.0\km\crt]],
            [[-IC:\Program Files (x86)\Windows Kits\10\Include\wdf\kmdf\1.11]],
            [[-Ibuild\.gens\hprdbghv\windows\x64\debug\rules\wdk\wpp]],
            [[-Ihyperdbg\kdserial]],
            "-DMANIFEST=NO",
            "-DMAP",
            "-DFAcs",
            "-D_WIN32_WINNT=0x0A00",
            "-DWINVER=0x0A00",
            "-DNTDDI_VERSION=0x0A000000",
            "-D_NT_TARGET_VERSION=0x0A00",
            "-D_WIN64",
            "-D_AMD64_",
            "-DAMD64",
            "-DKMDF_VERSION_MAJOR=1",
            "-DKMDF_VERSION_MINOR=11",
            "-DKMDF_USING_NTSTATUS",
            "-DWIN32_LEAN_AND_MEAN=1",
            "-DWINNT=1",
            "-D_WINDLL"
        }
    }
}