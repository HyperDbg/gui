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
ExceptionRecord EXCEPTION_RECORD //col:3
FirstChance uint32 //col:4
}


type DBGKM_CREATE_THREAD struct{
SubSystemKey uint32 //col:8
StartAddress PVOID //col:9
}


type DBGKM_CREATE_PROCESS struct{
SubSystemKey uint32 //col:13
FileHandle HANDLE //col:14
BaseOfImage PVOID //col:15
DebugInfoFileOffset uint32 //col:16
DebugInfoSize uint32 //col:17
InitialThread DBGKM_CREATE_THREAD //col:18
}


type DBGKM_EXIT_THREAD struct{
ExitStatus NTSTATUS //col:22
}


type DBGKM_EXIT_PROCESS struct{
ExitStatus NTSTATUS //col:26
}


type DBGKM_LOAD_DLL struct{
FileHandle HANDLE //col:30
BaseOfDll PVOID //col:31
DebugInfoFileOffset uint32 //col:32
DebugInfoSize uint32 //col:33
NamePointer PVOID //col:34
}


type DBGKM_UNLOAD_DLL struct{
BaseAddress PVOID //col:38
}


type DBGUI_CREATE_THREAD struct{
HandleToThread HANDLE //col:42
NewThread DBGKM_CREATE_THREAD //col:43
}


type DBGUI_CREATE_PROCESS struct{
HandleToProcess HANDLE //col:47
HandleToThread HANDLE //col:48
NewProcess DBGKM_CREATE_PROCESS //col:49
}


type DBGUI_WAIT_STATE_CHANGE struct{
NewState DBG_STATE //col:53
AppClientId CLIENT_ID //col:54
Union union //col:55
Exception DBGKM_EXCEPTION //col:57
CreateThread DBGUI_CREATE_THREAD //col:58
CreateProcessInfo DBGUI_CREATE_PROCESS //col:59
ExitThread DBGKM_EXIT_THREAD //col:60
ExitProcess DBGKM_EXIT_PROCESS //col:61
LoadDll DBGKM_LOAD_DLL //col:62
UnloadDll DBGKM_UNLOAD_DLL //col:63
}




