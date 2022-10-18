package features
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbghv\header\debugger\features\Hooks.h.back

const(
    SystemModuleInformation          =  11  //col:3
    SystemKernelDebuggerInformation  =  35  //col:4
)



type  _EPT_HOOKS_TEMPORARY_CONTEXT struct{
PhysicalAddress uint64 //col:7
VirtualAddress uint64 //col:8
}


type  _SSDTStruct struct{
* int32 //col:18
pCounterTable uintptr //col:19
#ifdefWin64 #ifdef _WIN64 //col:20
NumberOfServices uint64 //col:21
#else #else //col:22
NumberOfServices uint32 //col:23
#endif #endif //col:24
pArgumentTable PCHAR //col:25
}


type  _HIDDEN_HOOKS_DETOUR_DETAILS struct{
OtherHooksList *list.List //col:24
HookedFunctionAddress uintptr //col:25
ReturnAddress uintptr //col:26
}


type  _SYSTEM_MODULE_ENTRY struct{
Section uintptr //col:37
MappedBase uintptr //col:38
ImageBase uintptr //col:39
ImageSize uint32 //col:40
Flags uint32 //col:41
LoadOrderIndex uint16 //col:42
InitOrderIndex uint16 //col:43
LoadCount uint16 //col:44
OffsetToFileName uint16 //col:45
FullPathName[256] uint8 //col:46
}


type  _SYSTEM_MODULE_INFORMATION struct{
Count uint32 //col:42
Module[0] SYSTEM_MODULE_ENTRY //col:43
}




