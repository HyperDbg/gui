package Headers


const(
SIZEOF_DEBUGGER_READ_PAGE_TABLE_ENTRIES_DETAILS = sizeof(DEBUGGER_READ_PAGE_TABLE_ENTRIES_DETAILS) //col:14
SIZEOF_DEBUGGER_VA2PA_AND_PA2VA_COMMANDS = sizeof(DEBUGGER_VA2PA_AND_PA2VA_COMMANDS) //col:46
SIZEOF_DEBUGGER_DT_COMMAND_OPTIONS = sizeof(DEBUGGER_DT_COMMAND_OPTIONS) //col:66
SIZEOF_DEBUGGER_PREALLOC_COMMAND = sizeof(DEBUGGER_PREALLOC_COMMAND) //col:98
SIZEOF_DEBUGGER_READ_MEMORY = sizeof(DEBUGGER_READ_MEMORY) //col:116
SIZEOF_DEBUGGER_FLUSH_LOGGING_BUFFERS = sizeof(DEBUGGER_FLUSH_LOGGING_BUFFERS) //col:175
SIZEOF_DEBUGGER_TEST_QUERY_BUFFER = sizeof(DEBUGGER_TEST_QUERY_BUFFER) //col:193
SIZEOF_DEBUGGER_PERFORM_KERNEL_TESTS = sizeof(DEBUGGER_PERFORM_KERNEL_TESTS) //col:210
SIZEOF_DEBUGGER_SEND_COMMAND_EXECUTION_FINISHED_SIGNAL = sizeof(DEBUGGER_SEND_COMMAND_EXECUTION_FINISHED_SIGNAL) //col:226
SIZEOF_DEBUGGEE_KERNEL_AND_USER_TEST_INFORMATION = sizeof(DEBUGGEE_KERNEL_AND_USER_TEST_INFORMATION) //col:243
SIZEOF_DEBUGGEE_SEND_GENERAL_PACKET_FROM_DEBUGGEE_TO_DEBUGGER = sizeof(DEBUGGEE_SEND_GENERAL_PACKET_FROM_DEBUGGEE_TO_DEBUGGER) //col:261
SIZEOF_DEBUGGER_SEND_USERMODE_MESSAGES_TO_DEBUGGER = sizeof(DEBUGGER_SEND_USERMODE_MESSAGES_TO_DEBUGGER) //col:285
SIZEOF_DEBUGGER_READ_AND_WRITE_ON_MSR = sizeof(DEBUGGER_READ_AND_WRITE_ON_MSR) //col:307
SIZEOF_DEBUGGER_EDIT_MEMORY = sizeof(DEBUGGER_EDIT_MEMORY) //col:339
SIZEOF_DEBUGGER_SEARCH_MEMORY = sizeof(DEBUGGER_SEARCH_MEMORY) //col:382
SIZEOF_DEBUGGER_HIDE_AND_TRANSPARENT_DEBUGGER_MODE = sizeof(DEBUGGER_HIDE_AND_TRANSPARENT_DEBUGGER_MODE) //col:427
SIZEOF_DEBUGGER_PREPARE_DEBUGGEE = sizeof(DEBUGGER_PREPARE_DEBUGGEE) //col:461
SIZEOF_DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS = sizeof(DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS) //col:493
SIZEOF_DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS = sizeof(DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS) //col:533
SIZEOF_DEBUGGER_CALLSTACK_REQUEST = sizeof(DEBUGGER_CALLSTACK_REQUEST) //col:655
SIZEOF_USERMODE_DEBUGGING_THREAD_OR_PROCESS_STATE_DETAILS = sizeof(USERMODE_DEBUGGING_THREAD_OR_PROCESS_STATE_DETAILS) //col:691
SIZEOF_DEBUGGEE_DETAILS_AND_SWITCH_PROCESS_PACKET = sizeof(DEBUGGEE_DETAILS_AND_SWITCH_PROCESS_PACKET) //col:827
SIZEOF_DEBUGGEE_DETAILS_AND_SWITCH_THREAD_PACKET = sizeof(DEBUGGEE_DETAILS_AND_SWITCH_THREAD_PACKET) //col:865
)

const(
    DEBUGGER_PREALLOC_COMMAND_TYPE_MONITOR = 1  //col:94
    DEBUGGER_PREALLOC_COMMAND_TYPE_THREAD_INTERCEPTION = 2  //col:95
)


const(
    READ_FROM_KERNEL = 1  //col:124
    READ_FROM_VMX_ROOT = 2  //col:125
)


const(
    DEBUGGER_READ_PHYSICAL_ADDRESS = 1  //col:134
    DEBUGGER_READ_VIRTUAL_ADDRESS = 2  //col:135
)


const(
    DEBUGGER_SHOW_COMMAND_DT  =  1  //col:145
    DEBUGGER_SHOW_COMMAND_DISASSEMBLE64 = 2  //col:146
    DEBUGGER_SHOW_COMMAND_DISASSEMBLE32 = 3  //col:147
    DEBUGGER_SHOW_COMMAND_DB = 4  //col:148
    DEBUGGER_SHOW_COMMAND_DC = 5  //col:149
    DEBUGGER_SHOW_COMMAND_DQ = 6  //col:150
    DEBUGGER_SHOW_COMMAND_DD = 7  //col:151
)


const(
    DEBUGGER_MSR_READ = 1  //col:316
    DEBUGGER_MSR_WRITE = 2  //col:317
)


const(
    EDIT_PHYSICAL_MEMORY = 1  //col:347
    EDIT_VIRTUAL_MEMORY = 2  //col:348
)


const(
    EDIT_BYTE = 1  //col:357
    EDIT_DWORD = 2  //col:358
    EDIT_QWORD = 3  //col:359
)


const(
    SEARCH_PHYSICAL_MEMORY = 1  //col:390
    SEARCH_VIRTUAL_MEMORY = 2  //col:391
    SEARCH_PHYSICAL_FROM_VIRTUAL_MEMORY = 3  //col:392
)


const(
    SEARCH_BYTE = 1  //col:402
    SEARCH_DWORD = 2  //col:403
    SEARCH_QWORD = 3  //col:404
)


const(
    DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS_ACTION_ATTACH = 1  //col:502
    DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS_ACTION_DETACH = 2  //col:503
    DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS_ACTION_REMOVE_HOOKS = 3  //col:504
    DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS_ACTION_KILL_PROCESS = 4  //col:505
    DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS_ACTION_PAUSE_PROCESS = 5  //col:506
    DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS_ACTION_SWITCH_BY_PROCESS_OR_THREAD = 6  //col:507
    DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS_ACTION_QUERY_COUNT_OF_ACTIVE_DEBUGGING_THREADS = 7  //col:508
)


const(
    DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_QUERY_PROCESS_COUNT    =  1  //col:542
    DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_QUERY_THREAD_COUNT     =  2  //col:543
    DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_QUERY_PROCESS_LIST     =  3  //col:544
    DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_QUERY_THREAD_LIST      =  4  //col:545
    DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_QUERY_CURRENT_PROCESS  =  5  //col:546
    DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_QUERY_CURRENT_THREAD   =  6  //col:547
)


const(
    DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_ACTION_SHOW_INSTANTLY      =  1  //col:557
    DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_ACTION_QUERY_COUNT         =  2  //col:558
    DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_ACTION_QUERY_SAVE_DETAILS  =  3  //col:559
)


const(
    DEBUGGER_CALLSTACK_DISPLAY_METHOD_WITHOUT_PARAMS = 1  //col:664
    DEBUGGER_CALLSTACK_DISPLAY_METHOD_WITH_PARAMS = 2  //col:665
)


const(
    DEBUGGER_UD_COMMAND_ACTION_TYPE_NONE  =  0  //col:752
    DEBUGGER_UD_COMMAND_ACTION_TYPE_PAUSE = 2  //col:753
    DEBUGGER_UD_COMMAND_ACTION_TYPE_CONTINUE = 3  //col:754
    DEBUGGER_UD_COMMAND_ACTION_TYPE_REGULAR_STEP = 4  //col:755
)


const(
    DEBUGGEE_DETAILS_AND_SWITCH_PROCESS_GET_PROCESS_DETAILS = 1  //col:797
    DEBUGGEE_DETAILS_AND_SWITCH_PROCESS_GET_PROCESS_LIST = 2  //col:798
    DEBUGGEE_DETAILS_AND_SWITCH_PROCESS_PERFORM_SWITCH = 3  //col:799
)


const(
    DEBUGGEE_DETAILS_AND_SWITCH_THREAD_PERFORM_SWITCH = 1  //col:837
    DEBUGGEE_DETAILS_AND_SWITCH_THREAD_GET_THREAD_DETAILS = 2  //col:838
    DEBUGGEE_DETAILS_AND_SWITCH_THREAD_GET_THREAD_LIST = 3  //col:839
)


const(
    DEBUGGER_REMOTE_STEPPING_REQUEST_STEP_OVER = 1  //col:878
    DEBUGGER_REMOTE_STEPPING_REQUEST_STEP_IN = 2  //col:879
    DEBUGGER_REMOTE_STEPPING_REQUEST_INSTRUMENTATION_STEP_IN = 3  //col:880
)


const(
    DEBUGGEE_BREAKPOINT_MODIFICATION_REQUEST_LIST_BREAKPOINTS = 1  //col:973
    DEBUGGEE_BREAKPOINT_MODIFICATION_REQUEST_ENABLE = 2  //col:974
    DEBUGGEE_BREAKPOINT_MODIFICATION_REQUEST_DISABLE = 3  //col:975
    DEBUGGEE_BREAKPOINT_MODIFICATION_REQUEST_CLEAR = 4  //col:976
)


const(
    DEBUGGER_CONDITIONAL_JUMP_STATUS_ERROR  =  0  //col:1002
    DEBUGGER_CONDITIONAL_JUMP_STATUS_NOT_CONDITIONAL_JUMP = 2  //col:1003
    DEBUGGER_CONDITIONAL_JUMP_STATUS_JUMP_IS_TAKEN = 3  //col:1004
    DEBUGGER_CONDITIONAL_JUMP_STATUS_JUMP_IS_NOT_TAKEN = 4  //col:1005
)



type (
RequestStructures interface{
    sizeof()(ok bool)//col:40
    sizeof()(ok bool)//col:61
    sizeof()(ok bool)//col:83
    sizeof()(ok bool)//col:111
    sizeof()(ok bool)//col:188
    sizeof()(ok bool)//col:205
    sizeof()(ok bool)//col:221
    sizeof()(ok bool)//col:237
    sizeof()(ok bool)//col:255
    sizeof()(ok bool)//col:279
    sizeof()(ok bool)//col:301
    sizeof()(ok bool)//col:318
    sizeof()(ok bool)//col:455
    sizeof()(ok bool)//col:510
    sizeof()(ok bool)//col:549
    sizeof()(ok bool)//col:667
    sizeof()(ok bool)//col:700
    sizeof()(ok bool)//col:841
    sizeof()(ok bool)//col:882
}





)

func NewRequestStructures() { return & requestStructures{} }

func (r *requestStructures)    sizeof()(ok bool){//col:40















return true
}






func (r *requestStructures)    sizeof()(ok bool){//col:61









return true
}






func (r *requestStructures)    sizeof()(ok bool){//col:83











return true
}






func (r *requestStructures)    sizeof()(ok bool){//col:111







return true
}






func (r *requestStructures)    sizeof()(ok bool){//col:188







return true
}






func (r *requestStructures)    sizeof()(ok bool){//col:205






return true
}






func (r *requestStructures)    sizeof()(ok bool){//col:221





return true
}






func (r *requestStructures)    sizeof()(ok bool){//col:237





return true
}






func (r *requestStructures)    sizeof()(ok bool){//col:255






return true
}






func (r *requestStructures)    sizeof()(ok bool){//col:279








return true
}






func (r *requestStructures)    sizeof()(ok bool){//col:301






return true
}






func (r *requestStructures)    sizeof()(ok bool){//col:318






return true
}






func (r *requestStructures)    sizeof()(ok bool){//col:455














return true
}






func (r *requestStructures)    sizeof()(ok bool){//col:510











return true
}






func (r *requestStructures)    sizeof()(ok bool){//col:549










return true
}






func (r *requestStructures)    sizeof()(ok bool){//col:667






return true
}






func (r *requestStructures)    sizeof()(ok bool){//col:700







return true
}






func (r *requestStructures)    sizeof()(ok bool){//col:841







return true
}






func (r *requestStructures)    sizeof()(ok bool){//col:882







return true
}








