package features
//back\HyperDbgDev\hyperdbg\hprdbghv\header\debugger\features\Hooks.h.back

const(
IS_SYSRET_INSTRUCTION(Code) = (*((PUINT8)(Code) + 0) == 0x48 && *((PUINT8)(Code) + 1) == 0x0F && *((PUINT8)(Code) + 2) == 0x07) //col:1
IS_SYSCALL_INSTRUCTION(Code) = (*((PUINT8)(Code) + 0) == 0x0F && *((PUINT8)(Code) + 1) == 0x05) //col:5
IMAGE_DOS_SIGNATURE =    0x5A4D //col:8
IMAGE_OS2_SIGNATURE =    0x454E //col:9
IMAGE_OS2_SIGNATURE_LE = 0x454C //col:10
IMAGE_VXD_SIGNATURE =    0x454C //col:11
IMAGE_NT_SIGNATURE =     0x00004550 //col:12
)

const(
    SystemModuleInformation          =  11  //col:3
    SystemKernelDebuggerInformation  =  35  //col:4
)



type EPT_HOOKS_TEMPORARY_CONTEXT struct{
PhysicalAddress uint64
VirtualAddress uint64
}


type SSDTStruct struct{
* LONG
pCounterTable PVOID
#ifdefWin64 #ifdef _WIN64
NumberOfServices uint64
#else #else
NumberOfServices ULONG
#endif #endif
pArgumentTable PCHAR
}


type HIDDEN_HOOKS_DETOUR_DETAILS struct{
OtherHooksList LIST_ENTRY
HookedFunctionAddress PVOID
ReturnAddress PVOID
}


type SYSTEM_MODULE_ENTRY struct{
Section HANDLE
MappedBase PVOID
ImageBase PVOID
ImageSize ULONG
Flags ULONG
LoadOrderIndex UINT16
InitOrderIndex UINT16
LoadCount UINT16
OffsetToFileName UINT16
FullPathName[256] UCHAR
}


type SYSTEM_MODULE_INFORMATION struct{
Count ULONG
Module[0] SYSTEM_MODULE_ENTRY
}




