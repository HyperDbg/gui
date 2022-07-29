{
    files = {
        [[hyperdbg\hprdbghv\code\vmm\vmx\Vmcall.c]]
    },
    values = {
        [[C:\Program Files\Microsoft Visual Studio\2022\Enterprise\VC\Tools\MSVC\14.32.31326\bin\HostX64\x64\cl.exe]],
        {
            "-nologo",
            "-Zi",
            "-FS",
            [[-Fdbuild\windows\x64\debug\compile.hprdbghv.pdb]],
            "-Od",
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
            "-D_WINDLL",
            "-kernel"
        }
    },
    depfiles_cl_json = "{\
    \"Version\": \"1.2\",\
    \"Data\": {\
        \"Source\": \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\hprdbghv\\\\code\\\\vmm\\\\vmx\\\\vmcall.c\",\
        \"ProvidedModule\": \"\",\
        \"Includes\": [\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\hprdbghv\\\\pch.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\km\\\\ntifs.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\km\\\\ntddk.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\km\\\\wdm.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\km\\\\crt\\\\excpt.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\km\\\\crt\\\\crtdefs.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\specstrings.h\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\sal.h\",\
            \"c:\\\\program files\\\\microsoft visual studio\\\\2022\\\\enterprise\\\\vc\\\\tools\\\\msvc\\\\14.32.31326\\\\include\\\\concurrencysal.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\specstrings_strict.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\specstrings_undef.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\driverspecs.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\sdv_driverspecs.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\km\\\\crt\\\\vadefs.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\ntdef.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\km\\\\crt\\\\ctype.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\km\\\\crt\\\\crtdefs.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\winapifamily.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\winpackagefamily.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\kernelspecs.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\basetsd.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\guiddef.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\km\\\\crt\\\\string.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\km\\\\crt\\\\crtdefs.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\sdkddkver.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\ntstatus.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\bugcodes.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\ntiologc.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\km\\\\mce.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\pshpack4.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\poppack.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\pshpack4.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\poppack.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\pshpack1.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\poppack.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\guiddef.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\dpfilter.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\apiset.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\ktmtypes.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\evntprov.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\devpropdef.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\pshpack1.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\poppack.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\pshpack1.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\poppack.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\pshpack1.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\poppack.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\km\\\\ntnls.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\wdf\\\\kmdf\\\\1.11\\\\wdf.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\wdf\\\\kmdf\\\\1.11\\\\wdftypes.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\wdf\\\\kmdf\\\\1.11\\\\wdfglobals.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\wdf\\\\kmdf\\\\1.11\\\\wdffuncenum.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\wdf\\\\kmdf\\\\1.11\\\\wdfstatus.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\wdf\\\\kmdf\\\\1.11\\\\wdfassert.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\wdf\\\\kmdf\\\\1.11\\\\wdfverifier.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\wdf\\\\kmdf\\\\1.11\\\\wdfpool.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\wdf\\\\kmdf\\\\1.11\\\\wdfobject.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\wdf\\\\kmdf\\\\1.11\\\\wdfsync.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\wdf\\\\kmdf\\\\1.11\\\\wdfcore.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\wdf\\\\kmdf\\\\1.11\\\\wdfdriver.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\wdf\\\\kmdf\\\\1.11\\\\wdfqueryinterface.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\wdf\\\\kmdf\\\\1.11\\\\wdfmemory.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\wdf\\\\kmdf\\\\1.11\\\\wdfchildlist.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\wdf\\\\kmdf\\\\1.11\\\\wdffileobject.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\wdf\\\\kmdf\\\\1.11\\\\wdfdevice.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\km\\\\wdmsec.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\wdf\\\\kmdf\\\\1.11\\\\wdfcollection.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\wdf\\\\kmdf\\\\1.11\\\\wdfdpc.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\wdf\\\\kmdf\\\\1.11\\\\wdftimer.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\wdf\\\\kmdf\\\\1.11\\\\wdfworkitem.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\wdf\\\\kmdf\\\\1.11\\\\wdfinterrupt.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\wdf\\\\kmdf\\\\1.11\\\\wdfresource.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\wdf\\\\kmdf\\\\1.11\\\\wdfrequest.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\wdf\\\\kmdf\\\\1.11\\\\wdfiotarget.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\wdf\\\\kmdf\\\\1.11\\\\wdfio.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\wdf\\\\kmdf\\\\1.11\\\\wdffdo.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\wdf\\\\kmdf\\\\1.11\\\\wdfpdo.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\wdf\\\\kmdf\\\\1.11\\\\wdfcontrol.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\wdf\\\\kmdf\\\\1.11\\\\wdfwmi.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\wdf\\\\kmdf\\\\1.11\\\\wdfstring.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\wdf\\\\kmdf\\\\1.11\\\\wdfregistry.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\wdf\\\\kmdf\\\\1.11\\\\wdfdmaenabler.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\wdf\\\\kmdf\\\\1.11\\\\wdfdmatransaction.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\wdf\\\\kmdf\\\\1.11\\\\wdfcommonbuffer.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\wdf\\\\kmdf\\\\1.11\\\\wdfbugcodes.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\wdf\\\\kmdf\\\\1.11\\\\wdfroletypes.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\km\\\\ntstrsafe.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\km\\\\crt\\\\stdio.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\km\\\\crt\\\\crtdefs.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\km\\\\crt\\\\swprintf.inl\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\km\\\\crt\\\\crtdefs.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\km\\\\crt\\\\stdarg.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\km\\\\crt\\\\crtdefs.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\windef.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\shared\\\\minwindef.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\km\\\\crt\\\\intrin.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\km\\\\crt\\\\crtdefs.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\km\\\\crt\\\\setjmp.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\km\\\\crt\\\\crtdefs.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\km\\\\crt\\\\stddef.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\km\\\\crt\\\\crtdefs.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\km\\\\crt\\\\immintrin.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\km\\\\crt\\\\wmmintrin.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\km\\\\crt\\\\nmmintrin.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\km\\\\crt\\\\smmintrin.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\km\\\\crt\\\\tmmintrin.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\km\\\\crt\\\\crtdefs.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\km\\\\crt\\\\pmmintrin.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\km\\\\crt\\\\emmintrin.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\km\\\\crt\\\\crtdefs.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\km\\\\crt\\\\xmmintrin.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\km\\\\crt\\\\crtdefs.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\km\\\\crt\\\\mmintrin.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\km\\\\crt\\\\crtdefs.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\km\\\\crt\\\\malloc.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\km\\\\crt\\\\limits.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\km\\\\crt\\\\crtdefs.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\km\\\\crt\\\\zmmintrin.h\",\
            \"c:\\\\program files (x86)\\\\windows kits\\\\10\\\\include\\\\10.0.22621.0\\\\km\\\\crt\\\\ammintrin.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\dependencies\\\\ia32-doc\\\\out\\\\ia32.h\",\
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
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\include\\\\definition.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\include\\\\configuration.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\hprdbghv\\\\header\\\\platform\\\\crossapi.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\hprdbghv\\\\header\\\\platform\\\\environment.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\hprdbghv\\\\header\\\\platform\\\\metamacros.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\hprdbghv\\\\header\\\\vmm\\\\vmx\\\\vmxbroadcast.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\hprdbghv\\\\header\\\\common\\\\dpc.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\hprdbghv\\\\header\\\\common\\\\lengthdisassemblerengine.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\hprdbghv\\\\header\\\\common\\\\logging.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\hprdbghv\\\\header\\\\memory\\\\memorymapper.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\hprdbghv\\\\header\\\\common\\\\msr.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\hprdbghv\\\\header\\\\debugger\\\\tests\\\\kerneltests.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\hprdbghv\\\\header\\\\memory\\\\poolmanager.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\hprdbghv\\\\header\\\\common\\\\trace.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\hprdbghv\\\\header\\\\debugger\\\\core\\\\debugger.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\hprdbghv\\\\header\\\\debugger\\\\broadcast\\\\dpcroutines.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\hprdbghv\\\\header\\\\misc\\\\inlineasm.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\hprdbghv\\\\header\\\\vmm\\\\ept\\\\vpid.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\hprdbghv\\\\header\\\\vmm\\\\ept\\\\ept.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\hprdbghv\\\\header\\\\common\\\\common.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\hprdbghv\\\\header\\\\vmm\\\\vmx\\\\events.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\hprdbghv\\\\header\\\\debugger\\\\script-engine\\\\scriptengine.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\hprdbghv\\\\header\\\\devices\\\\apic.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\hprdbghv\\\\header\\\\debugger\\\\kernel-level\\\\kd.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\hprdbghv\\\\header\\\\debugger\\\\user-level\\\\ud.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\hprdbghv\\\\header\\\\vmm\\\\vmx\\\\mtf.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\hprdbghv\\\\header\\\\debugger\\\\core\\\\debuggerevents.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\hprdbghv\\\\header\\\\debugger\\\\features\\\\hooks.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\hprdbghv\\\\header\\\\vmm\\\\vmx\\\\counters.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\hprdbghv\\\\header\\\\debugger\\\\transparency\\\\transparency.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\hprdbghv\\\\header\\\\vmm\\\\vmx\\\\idtemulation.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\hprdbghv\\\\header\\\\vmm\\\\ept\\\\invept.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\hprdbghv\\\\header\\\\debugger\\\\broadcast\\\\broadcast.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\hprdbghv\\\\header\\\\vmm\\\\vmx\\\\vmcall.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\hprdbghv\\\\header\\\\vmm\\\\vmx\\\\manageregs.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\hprdbghv\\\\header\\\\vmm\\\\vmx\\\\vmx.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\hprdbghv\\\\header\\\\debugger\\\\commands\\\\breakpointcommands.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\hprdbghv\\\\header\\\\debugger\\\\commands\\\\debuggercommands.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\hprdbghv\\\\header\\\\debugger\\\\commands\\\\extensioncommands.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\hprdbghv\\\\header\\\\debugger\\\\commands\\\\callstack.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\hprdbghv\\\\header\\\\debugger\\\\communication\\\\serialconnection.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\hprdbghv\\\\header\\\\debugger\\\\objects\\\\process.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\hprdbghv\\\\header\\\\debugger\\\\objects\\\\thread.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\hprdbghv\\\\header\\\\components\\\\registers\\\\debugregisters.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\hprdbghv\\\\header\\\\vmm\\\\vmx\\\\hv.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\hprdbghv\\\\header\\\\vmm\\\\vmx\\\\msrhandlers.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\hprdbghv\\\\header\\\\vmm\\\\vmx\\\\protectedhv.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\hprdbghv\\\\header\\\\vmm\\\\vmx\\\\iohandler.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\hprdbghv\\\\header\\\\vmm\\\\vmx\\\\vmxmechanisms.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\hprdbghv\\\\header\\\\debugger\\\\user-level\\\\attaching.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\hprdbghv\\\\header\\\\debugger\\\\core\\\\termination.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\hprdbghv\\\\header\\\\debugger\\\\user-level\\\\useraccess.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\hprdbghv\\\\header\\\\debugger\\\\user-level\\\\threadholder.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\script-eval\\\\header\\\\scriptenginecommondefinitions.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\script-eval\\\\header\\\\scriptengineheader.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\hprdbghv\\\\header\\\\globals\\\\globalvariablemanagement.h\",\
            \"d:\\\\codespace\\\\gui\\\\sdk\\\\hyperdbgdev\\\\hyperdbg\\\\hprdbghv\\\\header\\\\globals\\\\globalvariables.h\"\
        ]\
    }\
}"
}