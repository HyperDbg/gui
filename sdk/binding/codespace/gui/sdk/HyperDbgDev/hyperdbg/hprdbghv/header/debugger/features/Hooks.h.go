package features


const(
    SystemModuleInformation          =  11  //col:3
    SystemKernelDebuggerInformation  =  35  //col:4
)



type EPT_HOOKS_TEMPORARY_CONTEXT struct{
PhysicalAddress uint64 //col:3
VirtualAddress uint64 //col:4
}



type SSDTStruct struct{
* LONG //col:8
pCounterTable PVOID //col:9
#ifdefWin64 #ifdef _WIN64 //col:10
NumberOfServices uint64 //col:11
#else #else //col:12
NumberOfServices uint32 //col:13
#endif #endif //col:14
pArgumentTable PCHAR //col:15
}



type HIDDEN_HOOKS_DETOUR_DETAILS struct{
OtherHooksList *list.List //col:19
HookedFunctionAddress PVOID //col:20
ReturnAddress PVOID //col:21
}



type SYSTEM_MODULE_ENTRY struct{
Section HANDLE //col:25
MappedBase PVOID //col:26
ImageBase PVOID //col:27
ImageSize uint32 //col:28
Flags uint32 //col:29
LoadOrderIndex uint16 //col:30
InitOrderIndex uint16 //col:31
LoadCount uint16 //col:32
OffsetToFileName uint16 //col:33
FullPathName[256] uint8 //col:34
}



type SYSTEM_MODULE_INFORMATION struct{
Count uint32 //col:38
Module[0] SYSTEM_MODULE_ENTRY //col:39
}





