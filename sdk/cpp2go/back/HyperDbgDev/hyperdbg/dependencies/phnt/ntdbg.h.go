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
type     DbgIdle uint32
const(
    DbgIdle DBG_STATE = 1  //col:161
    DbgReplyPending DBG_STATE = 2  //col:162
    DbgCreateThreadStateChange DBG_STATE = 3  //col:163
    DbgCreateProcessStateChange DBG_STATE = 4  //col:164
    DbgExitThreadStateChange DBG_STATE = 5  //col:165
    DbgExitProcessStateChange DBG_STATE = 6  //col:166
    DbgExceptionStateChange DBG_STATE = 7  //col:167
    DbgBreakpointStateChange DBG_STATE = 8  //col:168
    DbgSingleStepStateChange DBG_STATE = 9  //col:169
    DbgLoadDllStateChange DBG_STATE = 10  //col:170
    DbgUnloadDllStateChange DBG_STATE = 11  //col:171
)
type     DebugObjectUnusedInformation uint32
const(
    DebugObjectUnusedInformation DEBUGOBJECTINFOCLASS = 1  //col:215
    DebugObjectKillProcessOnExitInformation // s: ULONG DEBUGOBJECTINFOCLASS = 2  //col:216
    MaxDebugObjectInfoClass DEBUGOBJECTINFOCLASS = 3  //col:217
)
type (
Ntdbg interface{
 * Attribution 4.0 International ()(ok bool)//col:117
#define DEBUG_ALL_ACCESS ()(ok bool)//col:218
}

)
func NewNtdbg() { return & ntdbg{} }
func (n *ntdbg) * Attribution 4.0 International ()(ok bool){//col:117
return true
}

func (n *ntdbg)#define DEBUG_ALL_ACCESS ()(ok bool){//col:218
return true
}

