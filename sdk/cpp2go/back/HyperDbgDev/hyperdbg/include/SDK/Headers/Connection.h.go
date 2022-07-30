package Headers
//back\HyperDbgDev\hyperdbg\include\SDK\Headers\Connection.h.back

type     // uint32
const(
    // DEBUGGEE_PAUSING_REASON = 1  //col:21
    // For both kernel & user debugger DEBUGGEE_PAUSING_REASON = 2  //col:22
    // DEBUGGEE_PAUSING_REASON = 3  //col:23
    DEBUGGEE_PAUSING_REASON_NOT_PAUSED  DEBUGGEE_PAUSING_REASON =  0  //col:24
    DEBUGGEE_PAUSING_REASON_PAUSE_WITHOUT_DISASM DEBUGGEE_PAUSING_REASON = 5  //col:25
    DEBUGGEE_PAUSING_REASON_REQUEST_FROM_DEBUGGER DEBUGGEE_PAUSING_REASON = 6  //col:26
    DEBUGGEE_PAUSING_REASON_DEBUGGEE_STEPPED DEBUGGEE_PAUSING_REASON = 7  //col:27
    DEBUGGEE_PAUSING_REASON_DEBUGGEE_SOFTWARE_BREAKPOINT_HIT DEBUGGEE_PAUSING_REASON = 8  //col:28
    DEBUGGEE_PAUSING_REASON_DEBUGGEE_HARDWARE_DEBUG_REGISTER_HIT DEBUGGEE_PAUSING_REASON = 9  //col:29
    DEBUGGEE_PAUSING_REASON_DEBUGGEE_CORE_SWITCHED DEBUGGEE_PAUSING_REASON = 10  //col:30
    DEBUGGEE_PAUSING_REASON_DEBUGGEE_PROCESS_SWITCHED DEBUGGEE_PAUSING_REASON = 11  //col:31
    DEBUGGEE_PAUSING_REASON_DEBUGGEE_THREAD_SWITCHED DEBUGGEE_PAUSING_REASON = 12  //col:32
    DEBUGGEE_PAUSING_REASON_DEBUGGEE_COMMAND_EXECUTION_FINISHED DEBUGGEE_PAUSING_REASON = 13  //col:33
    DEBUGGEE_PAUSING_REASON_DEBUGGEE_EVENT_TRIGGERED DEBUGGEE_PAUSING_REASON = 14  //col:34
    DEBUGGEE_PAUSING_REASON_DEBUGGEE_ENTRY_POINT_REACHED DEBUGGEE_PAUSING_REASON = 15  //col:35
    // DEBUGGEE_PAUSING_REASON = 16  //col:37
    // Only for user-debugger DEBUGGEE_PAUSING_REASON = 17  //col:38
    // DEBUGGEE_PAUSING_REASON = 18  //col:39
    DEBUGGEE_PAUSING_REASON_DEBUGGEE_GENERAL_DEBUG_BREAK DEBUGGEE_PAUSING_REASON = 19  //col:40
    DEBUGGEE_PAUSING_REASON_DEBUGGEE_GENERAL_THREAD_INTERCEPTED DEBUGGEE_PAUSING_REASON = 20  //col:41
    // DEBUGGEE_PAUSING_REASON = 21  //col:43
    // Only used for hardware debugging DEBUGGEE_PAUSING_REASON = 22  //col:44
    // DEBUGGEE_PAUSING_REASON = 23  //col:45
    DEBUGGEE_PAUSING_REASON_HARDWARE_BASED_DEBUGGEE_GENERAL_BREAK DEBUGGEE_PAUSING_REASON = 24  //col:46
)


type     // uint32
const(
    // DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION = 1  //col:57
    // Debugger to debuggee (user-mode execution) DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION = 2  //col:58
    // DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION = 3  //col:59
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_USER_MODE_PAUSE  DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION =  1  //col:60
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_USER_MODE_DO_NOT_READ_ANY_PACKET DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION = 5  //col:61
    // DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION = 6  //col:63
    // Debugger to debuggee (vmx-root mode execution) DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION = 7  //col:64
    // DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION = 8  //col:65
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_MODE_STEP DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION = 9  //col:66
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_MODE_CONTINUE DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION = 10  //col:67
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_MODE_CLOSE_AND_UNLOAD_DEBUGGEE DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION = 11  //col:68
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_MODE_CHANGE_CORE DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION = 12  //col:69
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_MODE_FLUSH_BUFFERS DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION = 13  //col:70
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_MODE_CALLSTACK DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION = 14  //col:71
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_MODE_TEST_QUERY DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION = 15  //col:72
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_MODE_CHANGE_PROCESS DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION = 16  //col:73
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_MODE_CHANGE_THREAD DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION = 17  //col:74
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_RUN_SCRIPT DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION = 18  //col:75
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_USER_INPUT_BUFFER DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION = 19  //col:76
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_SEARCH_QUERY DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION = 20  //col:77
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_REGISTER_EVENT DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION = 21  //col:78
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_ADD_ACTION_TO_EVENT DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION = 22  //col:79
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_QUERY_AND_MODIFY_EVENT DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION = 23  //col:80
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_READ_REGISTERS DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION = 24  //col:81
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_READ_MEMORY DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION = 25  //col:82
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_EDIT_MEMORY DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION = 26  //col:83
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_BP DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION = 27  //col:84
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_LIST_OR_MODIFY_BREAKPOINTS DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION = 28  //col:85
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_SYMBOL_RELOAD DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION = 29  //col:86
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_QUERY_PA2VA_AND_VA2PA DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION = 30  //col:87
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_SYMBOL_QUERY_PTE DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION = 31  //col:88
    // DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION = 32  //col:90
    // Debuggee to debugger DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION = 33  //col:91
    // DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION = 34  //col:92
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_NO_ACTION  DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION =  0  //col:93
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_STARTED DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION = 36  //col:94
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_LOGGING_MECHANISM DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION = 37  //col:95
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_PAUSED_AND_CURRENT_INSTRUCTION DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION = 38  //col:96
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_CHANGING_CORE DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION = 39  //col:97
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_CHANGING_PROCESS DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION = 40  //col:98
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_CHANGING_THREAD DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION = 41  //col:99
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_RUNNING_SCRIPT DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION = 42  //col:100
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_FORMATS DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION = 43  //col:101
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_FLUSH DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION = 44  //col:102
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_CALLSTACK DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION = 45  //col:103
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_TEST_QUERY DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION = 46  //col:104
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_REGISTERING_EVENT DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION = 47  //col:105
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_ADDING_ACTION_TO_EVENT DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION = 48  //col:106
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_QUERY_AND_MODIFY_EVENT DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION = 49  //col:107
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_READING_REGISTERS DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION = 50  //col:108
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_READING_MEMORY DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION = 51  //col:109
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_EDITING_MEMORY DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION = 52  //col:110
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_BP DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION = 53  //col:111
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_LIST_OR_MODIFY_BREAKPOINTS DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION = 54  //col:112
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_UPDATE_SYMBOL_INFO DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION = 55  //col:113
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RELOAD_SYMBOL_FINISHED DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION = 56  //col:114
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RELOAD_SEARCH_QUERY DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION = 57  //col:115
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_PTE DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION = 58  //col:116
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_VA2PA_AND_PA2VA DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION = 59  //col:117
    // DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION = 60  //col:119
    // hardware debuggee to debugger DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION = 61  //col:120
    // DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION = 62  //col:121
    // DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION = 63  //col:123
    // hardware debugger to debuggee DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION = 64  //col:124
    // DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION = 65  //col:125
)


type     // uint32
const(
    // DEBUGGER_REMOTE_PACKET_TYPE = 1  //col:136
    // Debugger to debuggee (vmx-root) DEBUGGER_REMOTE_PACKET_TYPE = 2  //col:137
    // DEBUGGER_REMOTE_PACKET_TYPE = 3  //col:138
    DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGER_TO_DEBUGGEE_EXECUTE_ON_VMX_ROOT  DEBUGGER_REMOTE_PACKET_TYPE =  1  //col:139
    // DEBUGGER_REMOTE_PACKET_TYPE = 5  //col:141
    // Debugger to debuggee (user-mode) DEBUGGER_REMOTE_PACKET_TYPE = 6  //col:142
    // DEBUGGER_REMOTE_PACKET_TYPE = 7  //col:143
    DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGER_TO_DEBUGGEE_EXECUTE_ON_USER_MODE DEBUGGER_REMOTE_PACKET_TYPE = 8  //col:144
    // DEBUGGER_REMOTE_PACKET_TYPE = 9  //col:146
    // Debuggee to debugger (user-mode and kernel-mode vmx-root mode) DEBUGGER_REMOTE_PACKET_TYPE = 10  //col:147
    // DEBUGGER_REMOTE_PACKET_TYPE = 11  //col:148
    DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGEE_TO_DEBUGGER DEBUGGER_REMOTE_PACKET_TYPE = 12  //col:149
    // DEBUGGER_REMOTE_PACKET_TYPE = 13  //col:151
    // Debugger to debuggee (hardware) DEBUGGER_REMOTE_PACKET_TYPE = 14  //col:152
    // DEBUGGER_REMOTE_PACKET_TYPE = 15  //col:153
    DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGER_TO_DEBUGGEE_HARDWARE_LEVEL  DEBUGGER_REMOTE_PACKET_TYPE =  1  //col:154
    // DEBUGGER_REMOTE_PACKET_TYPE = 17  //col:156
    // Debuggee to debugger (hardware) DEBUGGER_REMOTE_PACKET_TYPE = 18  //col:157
    // DEBUGGER_REMOTE_PACKET_TYPE = 19  //col:158
    DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGEE_TO_DEBUGGER_HARDWARE_LEVEL DEBUGGER_REMOTE_PACKET_TYPE = 20  //col:159
)




