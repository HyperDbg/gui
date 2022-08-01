package phnt
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\phnt\ntdbg.h.back

const(
_NTDBG_H =  //col:1
DBG_STATUS_CONTROL_C = 1 //col:2
DBG_STATUS_SYSRQ = 2 //col:3
DBG_STATUS_BUGCHECK_FIRST = 3 //col:4
DBG_STATUS_BUGCHECK_SECOND = 4 //col:5
DBG_STATUS_FATAL = 5 //col:6
DBG_STATUS_DEBUG_CONTROL = 6 //col:7
DBG_STATUS_WORKER = 7 //col:8
DEBUG_READ_EVENT = 0x0001 //col:9
DEBUG_PROCESS_ASSIGN = 0x0002 //col:10
DEBUG_SET_INFORMATION = 0x0004 //col:11
DEBUG_QUERY_INFORMATION = 0x0008 //col:12
DEBUG_ALL_ACCESS (STANDARD_RIGHTS_REQUIRED | SYNCHRONIZE | = DEBUG_READ_EVENT | DEBUG_PROCESS_ASSIGN | DEBUG_SET_INFORMATION | DEBUG_QUERY_INFORMATION) //col:13
DEBUG_KILL_ON_CLOSE = 0x1 //col:16
)

const(
    DbgIdle = 1  //col:3
    DbgReplyPending = 2  //col:4
    DbgCreateThreadStateChange = 3  //col:5
    DbgCreateProcessStateChange = 4  //col:6
    DbgExitThreadStateChange = 5  //col:7
    DbgExitProcessStateChange = 6  //col:8
    DbgExceptionStateChange = 7  //col:9
    DbgBreakpointStateChange = 8  //col:10
    DbgSingleStepStateChange = 9  //col:11
    DbgLoadDllStateChange = 10  //col:12
    DbgUnloadDllStateChange = 11  //col:13
)


const(
    DebugObjectUnusedInformation = 1  //col:17
    DebugObjectKillProcessOnExitInformation  = 2  //col:18
    MaxDebugObjectInfoClass = 3  //col:19
)



type DBGKM_EXCEPTION struct{
ExceptionRecord EXCEPTION_RECORD
FirstChance ULONG
}


type DBGKM_CREATE_THREAD struct{
SubSystemKey ULONG
StartAddress PVOID
}


type DBGKM_CREATE_PROCESS struct{
SubSystemKey ULONG
FileHandle HANDLE
BaseOfImage PVOID
DebugInfoFileOffset ULONG
DebugInfoSize ULONG
InitialThread DBGKM_CREATE_THREAD
}


type DBGKM_EXIT_THREAD struct{
ExitStatus NTSTATUS
}


type DBGKM_EXIT_PROCESS struct{
ExitStatus NTSTATUS
}


type DBGKM_LOAD_DLL struct{
FileHandle HANDLE
BaseOfDll PVOID
DebugInfoFileOffset ULONG
DebugInfoSize ULONG
NamePointer PVOID
}


type DBGKM_UNLOAD_DLL struct{
BaseAddress PVOID
}


type DBGUI_CREATE_THREAD struct{
HandleToThread HANDLE
NewThread DBGKM_CREATE_THREAD
}


type DBGUI_CREATE_PROCESS struct{
HandleToProcess HANDLE
HandleToThread HANDLE
NewProcess DBGKM_CREATE_PROCESS
}


type DBGUI_WAIT_STATE_CHANGE struct{
NewState DBG_STATE
AppClientId CLIENT_ID
Union union
Exception DBGKM_EXCEPTION
CreateThread DBGUI_CREATE_THREAD
CreateProcessInfo DBGUI_CREATE_PROCESS
ExitThread DBGKM_EXIT_THREAD
ExitProcess DBGKM_EXIT_PROCESS
LoadDll DBGKM_LOAD_DLL
UnloadDll DBGKM_UNLOAD_DLL
}




