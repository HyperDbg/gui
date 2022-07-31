package phnt


const(
_NTDBG_H =  //col:13
DBG_STATUS_CONTROL_C = 1 //col:38
DBG_STATUS_SYSRQ = 2 //col:39
DBG_STATUS_BUGCHECK_FIRST = 3 //col:40
DBG_STATUS_BUGCHECK_SECOND = 4 //col:41
DBG_STATUS_FATAL = 5 //col:42
DBG_STATUS_DEBUG_CONTROL = 6 //col:43
DBG_STATUS_WORKER = 7 //col:44
DEBUG_READ_EVENT = 0x0001 //col:203
DEBUG_PROCESS_ASSIGN = 0x0002 //col:204
DEBUG_SET_INFORMATION = 0x0004 //col:205
DEBUG_QUERY_INFORMATION = 0x0008 //col:206
DEBUG_ALL_ACCESS (STANDARD_RIGHTS_REQUIRED | SYNCHRONIZE | = DEBUG_READ_EVENT | DEBUG_PROCESS_ASSIGN | DEBUG_SET_INFORMATION | DEBUG_QUERY_INFORMATION) //col:207
DEBUG_KILL_ON_CLOSE = 0x1 //col:211
)

const(
    DbgIdle = 1  //col:161
    DbgReplyPending = 2  //col:162
    DbgCreateThreadStateChange = 3  //col:163
    DbgCreateProcessStateChange = 4  //col:164
    DbgExitThreadStateChange = 5  //col:165
    DbgExitProcessStateChange = 6  //col:166
    DbgExceptionStateChange = 7  //col:167
    DbgBreakpointStateChange = 8  //col:168
    DbgSingleStepStateChange = 9  //col:169
    DbgLoadDllStateChange = 10  //col:170
    DbgUnloadDllStateChange = 11  //col:171
)


const(
    DebugObjectUnusedInformation = 1  //col:215
    DebugObjectKillProcessOnExitInformation // s: ULONG = 2  //col:216
    MaxDebugObjectInfoClass = 3  //col:217
)



type (
Ntdbg interface{
DbgUserBreakPoint()(ok bool)//col:117
}

)

func NewNtdbg() { return & ntdbg{} }

func (n *ntdbg)DbgUserBreakPoint()(ok bool){//col:117





















































































return true
}




