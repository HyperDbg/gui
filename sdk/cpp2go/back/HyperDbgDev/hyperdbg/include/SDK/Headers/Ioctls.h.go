package Headers
//back\HyperDbgDev\hyperdbg\include\SDK\Headers\Ioctls.h.back

const(
IOCTL_REGISTER_EVENT = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x800, METHOD_BUFFERED, FILE_ANY_ACCESS) //col:22
IOCTL_RETURN_IRP_PENDING_PACKETS_AND_DISALLOW_IOCTL = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x801, METHOD_BUFFERED, FILE_ANY_ACCESS) //col:29
IOCTL_TERMINATE_VMX = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x802, METHOD_BUFFERED, FILE_ANY_ACCESS) //col:36
IOCTL_DEBUGGER_READ_MEMORY = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x803, METHOD_BUFFERED, FILE_ANY_ACCESS) //col:43
IOCTL_DEBUGGER_READ_OR_WRITE_MSR = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x804, METHOD_BUFFERED, FILE_ANY_ACCESS) //col:50
IOCTL_DEBUGGER_READ_PAGE_TABLE_ENTRIES_DETAILS = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x805, METHOD_BUFFERED, FILE_ANY_ACCESS) //col:57
IOCTL_DEBUGGER_REGISTER_EVENT = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x806, METHOD_BUFFERED, FILE_ANY_ACCESS) //col:64
IOCTL_DEBUGGER_ADD_ACTION_TO_EVENT = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x807, METHOD_BUFFERED, FILE_ANY_ACCESS) //col:71
IOCTL_DEBUGGER_HIDE_AND_UNHIDE_TO_TRANSPARENT_THE_DEBUGGER = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x808, METHOD_BUFFERED, FILE_ANY_ACCESS) //col:78
IOCTL_DEBUGGER_VA2PA_AND_PA2VA_COMMANDS = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x809, METHOD_BUFFERED, FILE_ANY_ACCESS) //col:85
IOCTL_DEBUGGER_EDIT_MEMORY = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x80a, METHOD_BUFFERED, FILE_ANY_ACCESS) //col:92
IOCTL_DEBUGGER_SEARCH_MEMORY = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x80b, METHOD_BUFFERED, FILE_ANY_ACCESS) //col:99
IOCTL_DEBUGGER_MODIFY_EVENTS = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x80c, METHOD_BUFFERED, FILE_ANY_ACCESS) //col:106
IOCTL_DEBUGGER_FLUSH_LOGGING_BUFFERS = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x80d, METHOD_BUFFERED, FILE_ANY_ACCESS) //col:113
IOCTL_DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x80e, METHOD_BUFFERED, FILE_ANY_ACCESS) //col:120
IOCTL_DEBUGGER_PRINT = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x80f, METHOD_BUFFERED, FILE_ANY_ACCESS) //col:128
IOCTL_PREPARE_DEBUGGEE = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x810, METHOD_BUFFERED, FILE_ANY_ACCESS) //col:135
IOCTL_PAUSE_PACKET_RECEIVED = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x811, METHOD_BUFFERED, FILE_ANY_ACCESS) //col:142
IOCTL_SEND_SIGNAL_EXECUTION_IN_DEBUGGEE_FINISHED = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x812, METHOD_BUFFERED, FILE_ANY_ACCESS) //col:149
IOCTL_SEND_USERMODE_MESSAGES_TO_DEBUGGER = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x813, METHOD_BUFFERED, FILE_ANY_ACCESS) //col:156
IOCTL_SEND_GENERAL_BUFFER_FROM_DEBUGGEE_TO_DEBUGGER = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x814, METHOD_BUFFERED, FILE_ANY_ACCESS) //col:163
IOCTL_SEND_GET_KERNEL_SIDE_TEST_INFORMATION = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x815, METHOD_BUFFERED, FILE_ANY_ACCESS) //col:170
IOCTL_PERFROM_KERNEL_SIDE_TESTS = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x816, METHOD_BUFFERED, FILE_ANY_ACCESS) //col:177
IOCTL_RESERVE_PRE_ALLOCATED_POOLS = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x817, METHOD_BUFFERED, FILE_ANY_ACCESS) //col:184
IOCTL_SEND_USER_DEBUGGER_COMMANDS = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x818, METHOD_BUFFERED, FILE_ANY_ACCESS) //col:191
IOCTL_GET_DETAIL_OF_ACTIVE_THREADS_AND_PROCESSES = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x819, METHOD_BUFFERED, FILE_ANY_ACCESS) //col:198
IOCTL_GET_USER_MODE_MODULE_DETAILS = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x81a, METHOD_BUFFERED, FILE_ANY_ACCESS) //col:205
IOCTL_QUERY_COUNT_OF_ACTIVE_PROCESSES_OR_THREADS = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x81b, METHOD_BUFFERED, FILE_ANY_ACCESS) //col:212
IOCTL_GET_LIST_OF_THREADS_AND_PROCESSES = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x81c, METHOD_BUFFERED, FILE_ANY_ACCESS) //col:219
IOCTL_QUERY_CURRENT_PROCESS = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x81d, METHOD_BUFFERED, FILE_ANY_ACCESS) //col:226
IOCTL_QUERY_CURRENT_THREAD = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x81e, METHOD_BUFFERED, FILE_ANY_ACCESS) //col:233
)


