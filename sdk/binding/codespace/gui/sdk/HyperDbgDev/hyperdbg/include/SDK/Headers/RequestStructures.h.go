package Headers


const(
    DEBUGGER_PREALLOC_COMMAND_TYPE_MONITOR = 1  //col:3
    DEBUGGER_PREALLOC_COMMAND_TYPE_THREAD_INTERCEPTION = 2  //col:4
)


const(
    READ_FROM_KERNEL = 1  //col:8
    READ_FROM_VMX_ROOT = 2  //col:9
)


const(
    DEBUGGER_READ_PHYSICAL_ADDRESS = 1  //col:13
    DEBUGGER_READ_VIRTUAL_ADDRESS = 2  //col:14
)


const(
    DEBUGGER_SHOW_COMMAND_DT  =  1  //col:18
    DEBUGGER_SHOW_COMMAND_DISASSEMBLE64 = 2  //col:19
    DEBUGGER_SHOW_COMMAND_DISASSEMBLE32 = 3  //col:20
    DEBUGGER_SHOW_COMMAND_DB = 4  //col:21
    DEBUGGER_SHOW_COMMAND_DC = 5  //col:22
    DEBUGGER_SHOW_COMMAND_DQ = 6  //col:23
    DEBUGGER_SHOW_COMMAND_DD = 7  //col:24
)


const(
    DEBUGGER_MSR_READ = 1  //col:28
    DEBUGGER_MSR_WRITE = 2  //col:29
)


const(
    EDIT_PHYSICAL_MEMORY = 1  //col:33
    EDIT_VIRTUAL_MEMORY = 2  //col:34
)


const(
    EDIT_BYTE = 1  //col:38
    EDIT_DWORD = 2  //col:39
    EDIT_QWORD = 3  //col:40
)


const(
    SEARCH_PHYSICAL_MEMORY = 1  //col:44
    SEARCH_VIRTUAL_MEMORY = 2  //col:45
    SEARCH_PHYSICAL_FROM_VIRTUAL_MEMORY = 3  //col:46
)


const(
    SEARCH_BYTE = 1  //col:50
    SEARCH_DWORD = 2  //col:51
    SEARCH_QWORD = 3  //col:52
)


const(
    DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS_ACTION_ATTACH = 1  //col:56
    DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS_ACTION_DETACH = 2  //col:57
    DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS_ACTION_REMOVE_HOOKS = 3  //col:58
    DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS_ACTION_KILL_PROCESS = 4  //col:59
    DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS_ACTION_PAUSE_PROCESS = 5  //col:60
    DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS_ACTION_SWITCH_BY_PROCESS_OR_THREAD = 6  //col:61
    DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS_ACTION_QUERY_COUNT_OF_ACTIVE_DEBUGGING_THREADS = 7  //col:62
)


const(
    DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_QUERY_PROCESS_COUNT    =  1  //col:66
    DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_QUERY_THREAD_COUNT     =  2  //col:67
    DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_QUERY_PROCESS_LIST     =  3  //col:68
    DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_QUERY_THREAD_LIST      =  4  //col:69
    DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_QUERY_CURRENT_PROCESS  =  5  //col:70
    DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_QUERY_CURRENT_THREAD   =  6  //col:71
)


const(
    DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_ACTION_SHOW_INSTANTLY      =  1  //col:75
    DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_ACTION_QUERY_COUNT         =  2  //col:76
    DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_ACTION_QUERY_SAVE_DETAILS  =  3  //col:77
)


const(
    DEBUGGER_CALLSTACK_DISPLAY_METHOD_WITHOUT_PARAMS = 1  //col:81
    DEBUGGER_CALLSTACK_DISPLAY_METHOD_WITH_PARAMS = 2  //col:82
)


const(
    DEBUGGER_UD_COMMAND_ACTION_TYPE_NONE  =  0  //col:86
    DEBUGGER_UD_COMMAND_ACTION_TYPE_PAUSE = 2  //col:87
    DEBUGGER_UD_COMMAND_ACTION_TYPE_CONTINUE = 3  //col:88
    DEBUGGER_UD_COMMAND_ACTION_TYPE_REGULAR_STEP = 4  //col:89
)


const(
    DEBUGGEE_DETAILS_AND_SWITCH_PROCESS_GET_PROCESS_DETAILS = 1  //col:93
    DEBUGGEE_DETAILS_AND_SWITCH_PROCESS_GET_PROCESS_LIST = 2  //col:94
    DEBUGGEE_DETAILS_AND_SWITCH_PROCESS_PERFORM_SWITCH = 3  //col:95
)


const(
    DEBUGGEE_DETAILS_AND_SWITCH_THREAD_PERFORM_SWITCH = 1  //col:99
    DEBUGGEE_DETAILS_AND_SWITCH_THREAD_GET_THREAD_DETAILS = 2  //col:100
    DEBUGGEE_DETAILS_AND_SWITCH_THREAD_GET_THREAD_LIST = 3  //col:101
)


const(
    DEBUGGER_REMOTE_STEPPING_REQUEST_STEP_OVER = 1  //col:105
    DEBUGGER_REMOTE_STEPPING_REQUEST_STEP_IN = 2  //col:106
    DEBUGGER_REMOTE_STEPPING_REQUEST_INSTRUMENTATION_STEP_IN = 3  //col:107
)


const(
    DEBUGGEE_BREAKPOINT_MODIFICATION_REQUEST_LIST_BREAKPOINTS = 1  //col:111
    DEBUGGEE_BREAKPOINT_MODIFICATION_REQUEST_ENABLE = 2  //col:112
    DEBUGGEE_BREAKPOINT_MODIFICATION_REQUEST_DISABLE = 3  //col:113
    DEBUGGEE_BREAKPOINT_MODIFICATION_REQUEST_CLEAR = 4  //col:114
)


const(
    DEBUGGER_CONDITIONAL_JUMP_STATUS_ERROR  =  0  //col:118
    DEBUGGER_CONDITIONAL_JUMP_STATUS_NOT_CONDITIONAL_JUMP = 2  //col:119
    DEBUGGER_CONDITIONAL_JUMP_STATUS_JUMP_IS_TAKEN = 3  //col:120
    DEBUGGER_CONDITIONAL_JUMP_STATUS_JUMP_IS_NOT_TAKEN = 4  //col:121
)



type DEBUGGER_READ_PAGE_TABLE_ENTRIES_DETAILS struct{
VirtualAddress uint64 //col:3
ProcessId uint32 //col:4
Pml4eVirtualAddress uint64 //col:5
Pml4eValue uint64 //col:6
PdpteVirtualAddress uint64 //col:7
PdpteValue uint64 //col:8
PdeVirtualAddress uint64 //col:9
PdeValue uint64 //col:10
PteVirtualAddress uint64 //col:11
PteValue uint64 //col:12
KernelStatus uint32 //col:13
}



type DEBUGGER_VA2PA_AND_PA2VA_COMMANDS struct{
VirtualAddress uint64 //col:17
PhysicalAddress uint64 //col:18
ProcessId uint32 //col:19
IsVirtual2Physical bool //col:20
KernelStatus uint32 //col:21
}



type DEBUGGER_DT_COMMAND_OPTIONS struct{
char bool //col:25
SizeOfTypeName uint64 //col:26
Address uint64 //col:27
IsStruct bool //col:28
BufferAddress PVOID //col:29
TargetPid uint32 //col:30
char bool //col:31
}



type DEBUGGER_PREALLOC_COMMAND struct{
Type DEBUGGER_PREALLOC_COMMAND_TYPE //col:35
Count uint64 //col:36
KernelStatus uint32 //col:37
}



type DEBUGGER_READ_MEMORY struct{
Pid uint32 //col:41
Address uint64 //col:42
Size uint32 //col:43
MemoryType DEBUGGER_READ_MEMORY_TYPE //col:44
ReadingType DEBUGGER_READ_READING_TYPE //col:45
DtDetails PDEBUGGER_DT_COMMAND_OPTIONS //col:46
Style DEBUGGER_SHOW_MEMORY_STYLE //col:47
ReturnLength uint32 //col:48
KernelStatus uint32 //col:49
}



type DEBUGGER_FLUSH_LOGGING_BUFFERS struct{
KernelStatus uint32 //col:53
CountOfMessagesThatSetAsReadFromVmxRoot uint32 //col:54
CountOfMessagesThatSetAsReadFromVmxNonRoot uint32 //col:55
}



type DEBUGGER_DEBUGGER_TEST_QUERY_BUFFER struct{
RequestIndex uint32 //col:59
KernelStatus uint32 //col:60
}



type DEBUGGER_PERFORM_KERNEL_TESTS struct{
KernelStatus uint32 //col:64
}



type DEBUGGER_SEND_COMMAND_EXECUTION_FINISHED_SIGNAL struct{
KernelStatus uint32 //col:68
}



type DEBUGGEE_KERNEL_AND_USER_TEST_INFORMATION struct{
Value uint64 //col:72
Tag[32] int8 //col:73
}



type DEBUGGEE_SEND_GENERAL_PACKET_FROM_DEBUGGEE_TO_DEBUGGER struct{
RequestedAction DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION //col:77
LengthOfBuffer uint32 //col:78
PauseDebuggeeWhenSent bool //col:79
KernelResult uint32 //col:80
}



type DEBUGGER_SEND_USERMODE_MESSAGES_TO_DEBUGGER struct{
KernelStatus uint32 //col:84
Length uint32 //col:85
}



type DEBUGGER_READ_AND_WRITE_ON_MSR struct{
Msr uint64 //col:89
CoreNumber uint32 //col:90
DebuggerMsrActionType DEBUGGER_MSR_ACTION_TYPE //col:91
ActionType byte //col:92
Value uint64 //col:93
}



type DEBUGGER_EDIT_MEMORY struct{
Result uint32 //col:97
Address uint64 //col:98
ProcessId uint32 //col:99
MemoryType DEBUGGER_EDIT_MEMORY_TYPE //col:100
ByteSize DEBUGGER_EDIT_MEMORY_BYTE_SIZE //col:101
CountOf64Chunks uint32 //col:102
FinalStructureSize uint32 //col:103
KernelStatus uint32 //col:104
}



type DEBUGGER_SEARCH_MEMORY struct{
Address uint64 //col:108
Length uint64 //col:109
ProcessId uint32 //col:110
MemoryType DEBUGGER_SEARCH_MEMORY_TYPE //col:111
ByteSize DEBUGGER_SEARCH_MEMORY_BYTE_SIZE //col:112
CountOf64Chunks uint32 //col:113
FinalStructureSize uint32 //col:114
}



type DEBUGGER_HIDE_AND_TRANSPARENT_DEBUGGER_MODE struct{
IsHide bool //col:118
CpuidAverage uint64 //col:119
CpuidStandardDeviation uint64 //col:120
CpuidMedian uint64 //col:121
RdtscAverage uint64 //col:122
RdtscStandardDeviation uint64 //col:123
RdtscMedian uint64 //col:124
TrueIfProcessIdAndFalseIfProcessName bool //col:125
ProcId uint32 //col:126
LengthOfProcessName uint32 //col:127
DebuggerErrorUnableToHideOrUnhideDebugger DEBUGGER_ERROR_UNABLE_TO_HIDE_OR_UNHIDE_DEBUGGER //col:129
}



type DEBUGGER_PREPARE_DEBUGGEE struct{
PortAddress uint32 //col:134
Baudrate uint32 //col:135
NtoskrnlBaseAddress uint64 //col:136
Result uint32 //col:137
OsName[MAXIMUM_CHARACTER_FOR_OS_NAME] int8 //col:138
}



type DEBUGGEE_CHANGE_CORE_PACKET struct{
NewCore uint32 //col:142
Result uint32 //col:143
}



type DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS struct{
IsStartingNewProcess bool //col:147
ProcessId uint32 //col:148
ThreadId uint32 //col:149
Is32Bit bool //col:150
IsPaused bool //col:151
Action DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS_ACTION_TYPE //col:152
CountOfActiveDebuggingThreadsAndProcesses uint32 //col:153
Token uint64 //col:154
Result uint64 //col:155
}



type DEBUGGEE_PROCESS_LIST_NEEDED_DETAILS struct{
PsActiveProcessHead uint64 //col:159
ImageFileNameOffset uint32 //col:160
UniquePidOffset uint32 //col:161
ActiveProcessLinksOffset uint32 //col:162
}



type DEBUGGEE_THREAD_LIST_NEEDED_DETAILS struct{
ThreadListHeadOffset uint32 //col:166
ThreadListEntryOffset uint32 //col:167
CidOffset uint32 //col:168
PsActiveProcessHead uint64 //col:169
ActiveProcessLinksOffset uint32 //col:170
Process uint64 //col:171
}



type DEBUGGEE_PROCESS_LIST_DETAILS_ENTRY struct{
Eprocess uint64 //col:175
Pid uint32 //col:176
Cr3 uint64 //col:177
ImageFileName[15 uint8 //col:178
}



type DEBUGGEE_THREAD_LIST_DETAILS_ENTRY struct{
Eprocess uint64 //col:182
Ethread uint64 //col:183
Pid uint64 //col:184
Tid uint64 //col:185
ImageFileName[15 uint8 //col:186
}



type DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS struct{
ProcessListNeededDetails DEBUGGEE_PROCESS_LIST_NEEDED_DETAILS //col:190
ThreadListNeededDetails DEBUGGEE_THREAD_LIST_NEEDED_DETAILS //col:191
QueryType DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_TYPES //col:192
QueryAction DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_ACTIONS //col:193
Count uint32 //col:194
Result uint64 //col:195
}



type DEBUGGER_SINGLE_CALLSTACK_FRAME struct{
IsStackAddressValid bool //col:199
IsValidAddress bool //col:200
IsExecutable bool //col:201
Value uint64 //col:202
BYTE // //col:203
}



type DEBUGGER_CALLSTACK_REQUEST struct{
Is32Bit bool //col:207
KernelStatus uint32 //col:208
DisplayMethod DEBUGGER_CALLSTACK_DISPLAY_METHOD //col:209
Size uint32 //col:210
FrameCount uint32 //col:211
BaseAddress uint64 //col:212
BufferSize uint64 //col:213
}



type USERMODE_DEBUGGING_THREAD_OR_PROCESS_STATE_DETAILS struct{
ProcessId uint32 //col:217
ThreadId uint32 //col:218
IsProcess bool //col:219
}



type DEBUGGER_EVENT_ACTION_RUN_SCRIPT_CONFIGURATION struct{
ScriptBuffer uint64 //col:223
ScriptLength uint32 //col:224
ScriptPointer uint32 //col:225
OptionalRequestedBufferSize uint32 //col:226
}



type DEBUGGER_EVENT_REQUEST_BUFFER struct{
EnabledRequestBuffer bool //col:230
RequestBufferSize uint32 //col:231
RequstBufferAddress uint64 //col:232
}



type DEBUGGER_EVENT_REQUEST_CUSTOM_CODE struct{
CustomCodeBufferSize uint32 //col:236
CustomCodeBufferAddress PVOID //col:237
OptionalRequestedBufferSize uint32 //col:238
}



type DEBUGGER_UD_COMMAND_ACTION struct{
ActionType DEBUGGER_UD_COMMAND_ACTION_TYPE //col:242
OptionalParam1 uint64 //col:243
OptionalParam2 uint64 //col:244
OptionalParam3 uint64 //col:245
OptionalParam4 uint64 //col:246
}



type DEBUGGER_UD_COMMAND_PACKET struct{
UdAction DEBUGGER_UD_COMMAND_ACTION //col:250
ProcessDebuggingDetailToken uint64 //col:251
TargetThreadId uint32 //col:252
ApplyToAllPausedThreads bool //col:253
Result uint32 //col:254
}



type DEBUGGEE_DETAILS_AND_SWITCH_PROCESS_PACKET struct{
ActionType DEBUGGEE_DETAILS_AND_SWITCH_PROCESS_TYPE //col:258
ProcessId uint32 //col:259
Process uint64 //col:260
IsSwitchByClkIntr bool //col:261
ProcessName[16] uint8 //col:262
ProcessListSymDetails DEBUGGEE_PROCESS_LIST_NEEDED_DETAILS //col:263
Result uint32 //col:264
}



type DEBUGGEE_DETAILS_AND_SWITCH_THREAD_PACKET struct{
ActionType DEBUGGEE_DETAILS_AND_SWITCH_THREAD_TYPE //col:268
ThreadId uint32 //col:269
ProcessId uint32 //col:270
Thread uint64 //col:271
Process uint64 //col:272
CheckByClockInterrupt bool //col:273
ProcessName[16] uint8 //col:274
ThreadListSymDetails DEBUGGEE_THREAD_LIST_NEEDED_DETAILS //col:275
Result uint32 //col:276
}



type DEBUGGEE_STEP_PACKET struct{
StepType DEBUGGER_REMOTE_STEPPING_REQUEST //col:280
BOOLEAN // //col:281
CallLength uint32 //col:282
}



type DEBUGGEE_FORMATS_PACKET struct{
Value uint64 //col:286
Result uint32 //col:287
}



type DEBUGGEE_SYMBOL_REQUEST_PACKET struct{
ProcessId uint32 //col:291
}



type DEBUGGEE_BP_PACKET struct{
Address uint64 //col:295
Pid uint32 //col:296
Tid uint32 //col:297
Core uint32 //col:298
Result uint32 //col:299
}



type DEBUGGEE_BP_DESCRIPTOR struct{
BreakpointId uint64 //col:303
BreakpointsList *list.List //col:304
Enabled bool //col:305
Address uint64 //col:306
PhysAddress uint64 //col:307
Pid uint32 //col:308
Tid uint32 //col:309
Core uint32 //col:310
UINT16 // //col:311
PreviousByte uint8 //col:312
SetRflagsIFBitOnMtf bool //col:313
AvoidReApplyBreakpoint bool //col:314
}



type DEBUGGEE_BP_LIST_OR_MODIFY_PACKET struct{
BreakpointId uint64 //col:318
Request DEBUGGEE_BREAKPOINT_MODIFICATION_REQUEST //col:319
Result uint32 //col:320
}



type DEBUGGEE_SCRIPT_PACKET struct{
ScriptBufferSize uint32 //col:324
ScriptBufferPointer uint32 //col:325
IsFormat bool //col:326
Result uint32 //col:327
}



type DEBUGGEE_RESULT_OF_SEARCH_PACKET struct{
CountOfResults uint32 //col:331
Result uint32 //col:332
}



type DEBUGGEE_REGISTER_READ_DESCRIPTION struct{
RegisterID uint32 //col:336
Value uint64 //col:337
KernelStatus uint32 //col:338
}





