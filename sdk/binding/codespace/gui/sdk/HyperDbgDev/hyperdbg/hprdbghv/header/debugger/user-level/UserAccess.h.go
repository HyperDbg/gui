package user-level
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbghv\header\debugger\user-level\UserAccess.h.back

type PEB_LDR_DATA struct{
Length ULONG
Initialized bool
SsHandle PVOID
ModuleListLoadOrder *list.List
ModuleListMemoryOrder *list.List
ModuleListInitOrder *list.List
}


type RTL_USER_PROCESS_PARAMETERS struct{
Reserved1[16] uint8
Reserved2[10] PVOID
ImagePathName UNICODE_STRING
CommandLine UNICODE_STRING
}


type PEB struct{
Reserved1[2] uint8
BeingDebugged uint8
Reserved2[1] uint8
Reserved3[2] PVOID
Ldr PPEB_LDR_DATA
ProcessParameters PRTL_USER_PROCESS_PARAMETERS
Reserved4[3] PVOID
AtlThunkSListPtr PVOID
Reserved5 PVOID
Reserved6 ULONG
Reserved7 PVOID
Reserved8 ULONG
AtlThunkSListPtr32 ULONG
Reserved9[45] PVOID
Reserved10[96] uint8
PostProcessInitRoutine PPS_POST_PROCESS_INIT_ROUTINE
Reserved11[128] uint8
Reserved12[1] PVOID
SessionId ULONG
}


type PEB32 struct{
InheritedAddressSpace UCHAR
ReadImageFileExecOptions UCHAR
BeingDebugged UCHAR
BitField UCHAR
Mutant ULONG
ImageBaseAddress ULONG
Ldr ULONG
ProcessParameters ULONG
SubSystemData ULONG
ProcessHeap ULONG
FastPebLock ULONG
AtlThunkSListPtr ULONG
IFEOKey ULONG
CrossProcessFlags ULONG
UserSharedInfoPtr ULONG
SystemReserved ULONG
AtlThunkSListPtr32 ULONG
ApiSetMap ULONG
}


type PEB_LDR_DATA32 struct{
Length ULONG
Initialized UCHAR
SsHandle ULONG
InLoadOrderModuleList LIST_ENTRY32
InMemoryOrderModuleList LIST_ENTRY32
InInitializationOrderModuleList LIST_ENTRY32
}


type LDR_DATA_TABLE_ENTRY32 struct{
InLoadOrderLinks LIST_ENTRY32
InMemoryOrderLinks LIST_ENTRY32
InInitializationOrderLinks LIST_ENTRY32
DllBase ULONG
EntryPoint ULONG
SizeOfImage ULONG
FullDllName UNICODE_STRING32
BaseDllName UNICODE_STRING32
Flags ULONG
LoadCount UINT16
TlsIndex UINT16
HashLinks LIST_ENTRY32
TimeDateStamp ULONG
}


type LDR_DATA_TABLE_ENTRY struct{
InLoadOrderModuleList *list.List
InMemoryOrderModuleList *list.List
InInitializationOrderModuleList *list.List
DllBase PVOID
EntryPoint PVOID
SizeOfImage ULONG
FullDllName UNICODE_STRING
BaseDllName UNICODE_STRING
Flags ULONG
LoadCount UINT16
TlsIndex UINT16
HashLinks *list.List
SectionPointer PVOID
CheckSum ULONG
TimeDateStamp ULONG
}




