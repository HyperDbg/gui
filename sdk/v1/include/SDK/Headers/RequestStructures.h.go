package Headers
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\include\SDK\Headers\RequestStructures.h.back

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



type  _DEBUGGER_READ_PAGE_TABLE_ENTRIES_DETAILS struct{
VirtualAddress uint64 //col:16
ProcessId uint32 //col:17
Pml4eVirtualAddress uint64 //col:18
Pml4eValue uint64 //col:19
PdpteVirtualAddress uint64 //col:20
PdpteValue uint64 //col:21
PdeVirtualAddress uint64 //col:22
PdeValue uint64 //col:23
PteVirtualAddress uint64 //col:24
PteValue uint64 //col:25
KernelStatus uint32 //col:26
}


type  _DEBUGGER_VA2PA_AND_PA2VA_COMMANDS struct{
VirtualAddress uint64 //col:24
PhysicalAddress uint64 //col:25
ProcessId uint32 //col:26
IsVirtual2Physical bool //col:27
KernelStatus uint32 //col:28
}


type  _DEBUGGER_DT_COMMAND_OPTIONS struct{
char bool //col:34
SizeOfTypeName uint64 //col:35
Address uint64 //col:36
IsStruct bool //col:37
BufferAddress uintptr //col:38
TargetPid uint32 //col:39
char bool //col:40
}


type  _DEBUGGER_PREALLOC_COMMAND struct{
Type DEBUGGER_PREALLOC_COMMAND_TYPE //col:40
Count uint64 //col:41
KernelStatus uint32 //col:42
}


type  _DEBUGGER_READ_MEMORY struct{
Pid uint32 //col:52
Address uint64 //col:53
Size uint32 //col:54
MemoryType DEBUGGER_READ_MEMORY_TYPE //col:55
ReadingType DEBUGGER_READ_READING_TYPE //col:56
DtDetails PDEBUGGER_DT_COMMAND_OPTIONS //col:57
Style DEBUGGER_SHOW_MEMORY_STYLE //col:58
ReturnLength uint32 //col:59
KernelStatus uint32 //col:60
}


type  _DEBUGGER_FLUSH_LOGGING_BUFFERS struct{
KernelStatus uint32 //col:58
CountOfMessagesThatSetAsReadFromVmxRoot uint32 //col:59
CountOfMessagesThatSetAsReadFromVmxNonRoot uint32 //col:60
}


type  _DEBUGGER_DEBUGGER_TEST_QUERY_BUFFER struct{
RequestIndex uint32 //col:63
KernelStatus uint32 //col:64
}


type  _DEBUGGER_PERFORM_KERNEL_TESTS struct{
KernelStatus uint32 //col:67
}


type  _DEBUGGER_SEND_COMMAND_EXECUTION_FINISHED_SIGNAL struct{
KernelStatus uint32 //col:71
}


type  _DEBUGGEE_KERNEL_AND_USER_TEST_INFORMATION struct{
Value uint64 //col:76
Tag[32] int8 //col:77
}


type  _DEBUGGEE_SEND_GENERAL_PACKET_FROM_DEBUGGEE_TO_DEBUGGER struct{
RequestedAction DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION //col:83
LengthOfBuffer uint32 //col:84
PauseDebuggeeWhenSent bool //col:85
KernelResult uint32 //col:86
}


type  _DEBUGGER_SEND_USERMODE_MESSAGES_TO_DEBUGGER struct{
KernelStatus uint32 //col:88
Length uint32 //col:89
}


type  _DEBUGGER_READ_AND_WRITE_ON_MSR struct{
Msr uint64 //col:96
CoreNumber uint32 //col:97
DebuggerMsrActionType DEBUGGER_MSR_ACTION_TYPE //col:98
ActionType byte //col:99
Value uint64 //col:100
}


type  _DEBUGGER_EDIT_MEMORY struct{
Result uint32 //col:107
Address uint64 //col:108
ProcessId uint32 //col:109
MemoryType DEBUGGER_EDIT_MEMORY_TYPE //col:110
ByteSize DEBUGGER_EDIT_MEMORY_BYTE_SIZE //col:111
CountOf64Chunks uint32 //col:112
FinalStructureSize uint32 //col:113
KernelStatus uint32 //col:114
}


type  _DEBUGGER_SEARCH_MEMORY struct{
Address uint64 //col:117
Length uint64 //col:118
ProcessId uint32 //col:119
MemoryType DEBUGGER_SEARCH_MEMORY_TYPE //col:120
ByteSize DEBUGGER_SEARCH_MEMORY_BYTE_SIZE //col:121
CountOf64Chunks uint32 //col:122
FinalStructureSize uint32 //col:123
}


type  _DEBUGGER_HIDE_AND_TRANSPARENT_DEBUGGER_MODE struct{
IsHide bool //col:133
CpuidAverage uint64 //col:134
CpuidStandardDeviation uint64 //col:135
CpuidMedian uint64 //col:136
RdtscAverage uint64 //col:137
RdtscStandardDeviation uint64 //col:138
RdtscMedian uint64 //col:139
TrueIfProcessIdAndFalseIfProcessName bool //col:140
ProcId uint32 //col:141
LengthOfProcessName uint32 //col:142
DebuggerErrorUnableToHideOrUnhideDebugger DEBUGGER_ERROR_UNABLE_TO_HIDE_OR_UNHIDE_DEBUGGER //col:144
}


type  _DEBUGGER_PREPARE_DEBUGGEE struct{
PortAddress uint32 //col:141
Baudrate uint32 //col:142
NtoskrnlBaseAddress uint64 //col:143
Result uint32 //col:144
OsName[MAXIMUM_CHARACTER_FOR_OS_NAME] int8 //col:145
}


type  _DEBUGGEE_CHANGE_CORE_PACKET struct{
NewCore uint32 //col:146
Result uint32 //col:147
}


type  _DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS struct{
IsStartingNewProcess bool //col:158
ProcessId uint32 //col:159
ThreadId uint32 //col:160
Is32Bit bool //col:161
IsPaused bool //col:162
Action DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS_ACTION_TYPE //col:163
CountOfActiveDebuggingThreadsAndProcesses uint32 //col:164
Token uint64 //col:165
Result uint64 //col:166
}


type  _DEBUGGEE_PROCESS_LIST_NEEDED_DETAILS struct{
PsActiveProcessHead uint64 //col:165
ImageFileNameOffset uint32 //col:166
UniquePidOffset uint32 //col:167
ActiveProcessLinksOffset uint32 //col:168
}


type  _DEBUGGEE_THREAD_LIST_NEEDED_DETAILS struct{
ThreadListHeadOffset uint32 //col:174
ThreadListEntryOffset uint32 //col:175
CidOffset uint32 //col:176
PsActiveProcessHead uint64 //col:177
ActiveProcessLinksOffset uint32 //col:178
Process uint64 //col:179
}


type  _DEBUGGEE_PROCESS_LIST_DETAILS_ENTRY struct{
Eprocess uint64 //col:181
Pid uint32 //col:182
Cr3 uint64 //col:183
ImageFileName[15 uint8 //col:184
}


type  _DEBUGGEE_THREAD_LIST_DETAILS_ENTRY struct{
Eprocess uint64 //col:189
Ethread uint64 //col:190
Pid uint64 //col:191
Tid uint64 //col:192
ImageFileName[15 uint8 //col:193
}


type  _DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS struct{
ProcessListNeededDetails DEBUGGEE_PROCESS_LIST_NEEDED_DETAILS //col:198
ThreadListNeededDetails DEBUGGEE_THREAD_LIST_NEEDED_DETAILS //col:199
QueryType DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_TYPES //col:200
QueryAction DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_ACTIONS //col:201
Count uint32 //col:202
Result uint64 //col:203
}


type  _DEBUGGER_SINGLE_CALLSTACK_FRAME struct{
IsStackAddressValid bool //col:206
IsValidAddress bool //col:207
IsExecutable bool //col:208
Value uint64 //col:209
InstructionBytesOnRip[MAXIMUM_CALL_INSTR_SIZE] uint8 //col:210
}


type  _DEBUGGER_CALLSTACK_REQUEST struct{
Is32Bit bool //col:216
KernelStatus uint32 //col:217
DisplayMethod DEBUGGER_CALLSTACK_DISPLAY_METHOD //col:218
Size uint32 //col:219
FrameCount uint32 //col:220
BaseAddress uint64 //col:221
BufferSize uint64 //col:222
}


type  _USERMODE_DEBUGGING_THREAD_OR_PROCESS_STATE_DETAILS struct{
ProcessId uint32 //col:222
ThreadId uint32 //col:223
IsProcess bool //col:224
}


type  _DEBUGGER_EVENT_ACTION_RUN_SCRIPT_CONFIGURATION struct{
ScriptBuffer uint64 //col:229
ScriptLength uint32 //col:230
ScriptPointer uint32 //col:231
OptionalRequestedBufferSize uint32 //col:232
}


type  _DEBUGGER_EVENT_REQUEST_BUFFER struct{
EnabledRequestBuffer bool //col:235
RequestBufferSize uint32 //col:236
RequstBufferAddress uint64 //col:237
}


type  _DEBUGGER_EVENT_REQUEST_CUSTOM_CODE struct{
CustomCodeBufferSize uint32 //col:241
CustomCodeBufferAddress uintptr //col:242
OptionalRequestedBufferSize uint32 //col:243
}


type  _DEBUGGER_UD_COMMAND_ACTION struct{
ActionType DEBUGGER_UD_COMMAND_ACTION_TYPE //col:249
OptionalParam1 uint64 //col:250
OptionalParam2 uint64 //col:251
OptionalParam3 uint64 //col:252
OptionalParam4 uint64 //col:253
}


type  _DEBUGGER_UD_COMMAND_PACKET struct{
UdAction DEBUGGER_UD_COMMAND_ACTION //col:257
ProcessDebuggingDetailToken uint64 //col:258
TargetThreadId uint32 //col:259
ApplyToAllPausedThreads bool //col:260
Result uint32 //col:261
}


type  _DEBUGGEE_DETAILS_AND_SWITCH_PROCESS_PACKET struct{
ActionType DEBUGGEE_DETAILS_AND_SWITCH_PROCESS_TYPE //col:267
ProcessId uint32 //col:268
Process uint64 //col:269
IsSwitchByClkIntr bool //col:270
ProcessName[16] uint8 //col:271
ProcessListSymDetails DEBUGGEE_PROCESS_LIST_NEEDED_DETAILS //col:272
Result uint32 //col:273
}


type  _DEBUGGEE_DETAILS_AND_SWITCH_THREAD_PACKET struct{
ActionType DEBUGGEE_DETAILS_AND_SWITCH_THREAD_TYPE //col:279
ThreadId uint32 //col:280
ProcessId uint32 //col:281
Thread uint64 //col:282
Process uint64 //col:283
CheckByClockInterrupt bool //col:284
ProcessName[16] uint8 //col:285
ThreadListSymDetails DEBUGGEE_THREAD_LIST_NEEDED_DETAILS //col:286
Result uint32 //col:287
}


type  _DEBUGGEE_STEP_PACKET struct{
StepType DEBUGGER_REMOTE_STEPPING_REQUEST //col:285
IsCurrentInstructionACall bool //col:286
CallLength uint32 //col:287
}


type  _DEBUGGEE_FORMATS_PACKET struct{
Value uint64 //col:290
Result uint32 //col:291
}


type  _DEBUGGEE_SYMBOL_REQUEST_PACKET struct{
ProcessId uint32 //col:294
}


type  _DEBUGGEE_BP_PACKET struct{
Address uint64 //col:302
Pid uint32 //col:303
Tid uint32 //col:304
Core uint32 //col:305
Result uint32 //col:306
}


type  _DEBUGGEE_BP_DESCRIPTOR struct{
BreakpointId uint64 //col:317
BreakpointsList *list.List //col:318
Enabled bool //col:319
Address uint64 //col:320
PhysAddress uint64 //col:321
Pid uint32 //col:322
Tid uint32 //col:323
Core uint32 //col:324
InstructionLength uint16 //col:325
PreviousByte uint8 //col:326
SetRflagsIFBitOnMtf bool //col:327
AvoidReApplyBreakpoint bool //col:328
}


type  _DEBUGGEE_BP_LIST_OR_MODIFY_PACKET struct{
BreakpointId uint64 //col:323
Request DEBUGGEE_BREAKPOINT_MODIFICATION_REQUEST //col:324
Result uint32 //col:325
}


type  _DEBUGGEE_SCRIPT_PACKET struct{
ScriptBufferSize uint32 //col:330
ScriptBufferPointer uint32 //col:331
IsFormat bool //col:332
Result uint32 //col:333
}


type  _DEBUGGEE_RESULT_OF_SEARCH_PACKET struct{
CountOfResults uint32 //col:335
Result uint32 //col:336
}


type  _DEBUGGEE_REGISTER_READ_DESCRIPTION struct{
RegisterID uint32 //col:341
Value uint64 //col:342
KernelStatus uint32 //col:343
}




