package Headers
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\include\SDK\Headers\Constants.h.back

const(
VERSION_MAJOR = 0 //col:1
VERSION_MINOR = 2 //col:2
VERSION_PATCH = 0 //col:3
BUILD_YEAR_CH0 = (__DATE__[7]) //col:4
BUILD_YEAR_CH1 = (__DATE__[8]) //col:5
BUILD_YEAR_CH2 = (__DATE__[9]) //col:6
BUILD_YEAR_CH3 = (__DATE__[10]) //col:7
BUILD_MONTH_IS_JAN = (__DATE__[0] == 'J' && __DATE__[1] == 'a' && __DATE__[2] == 'n') //col:8
BUILD_MONTH_IS_FEB = (__DATE__[0] == 'F') //col:9
BUILD_MONTH_IS_MAR = (__DATE__[0] == 'M' && __DATE__[1] == 'a' && __DATE__[2] == 'r') //col:10
BUILD_MONTH_IS_APR = (__DATE__[0] == 'A' && __DATE__[1] == 'p') //col:11
BUILD_MONTH_IS_MAY = (__DATE__[0] == 'M' && __DATE__[1] == 'a' && __DATE__[2] == 'y') //col:12
BUILD_MONTH_IS_JUN = (__DATE__[0] == 'J' && __DATE__[1] == 'u' && __DATE__[2] == 'n') //col:13
BUILD_MONTH_IS_JUL = (__DATE__[0] == 'J' && __DATE__[1] == 'u' && __DATE__[2] == 'l') //col:14
BUILD_MONTH_IS_AUG = (__DATE__[0] == 'A' && __DATE__[1] == 'u') //col:15
BUILD_MONTH_IS_SEP = (__DATE__[0] == 'S') //col:16
BUILD_MONTH_IS_OCT = (__DATE__[0] == 'O') //col:17
BUILD_MONTH_IS_NOV = (__DATE__[0] == 'N') //col:18
BUILD_MONTH_IS_DEC = (__DATE__[0] == 'D') //col:19
BUILD_MONTH_CH0 = ((BUILD_MONTH_IS_OCT || BUILD_MONTH_IS_NOV || BUILD_MONTH_IS_DEC) ? '1' : '0') //col:20
BUILD_MONTH_CH1 = ( (BUILD_MONTH_IS_JAN) ? '1' : (BUILD_MONTH_IS_FEB) ? '2' : (BUILD_MONTH_IS_MAR)   ? '3' : (BUILD_MONTH_IS_APR)   ? '4' : (BUILD_MONTH_IS_MAY)   ? '5' : (BUILD_MONTH_IS_JUN)   ? '6' : (BUILD_MONTH_IS_JUL)   ? '7' : (BUILD_MONTH_IS_AUG)   ? '8' : (BUILD_MONTH_IS_SEP)   ? '9' : (BUILD_MONTH_IS_OCT)   ? '0' : (BUILD_MONTH_IS_NOV)   ? '1' : (BUILD_MONTH_IS_DEC)   ? '2' : '?') //col:22
BUILD_DAY_CH0 = ((__DATE__[4] >= '0') ? (__DATE__[4]) : '0') //col:36
BUILD_DAY_CH1 = (__DATE__[5]) //col:37
BUILD_HOUR_CH0 = (__TIME__[0]) //col:38
BUILD_HOUR_CH1 = (__TIME__[1]) //col:39
BUILD_MIN_CH0 = (__TIME__[3]) //col:40
BUILD_MIN_CH1 = (__TIME__[4]) //col:41
BUILD_SEC_CH0 = (__TIME__[6]) //col:42
BUILD_SEC_CH1 = (__TIME__[7]) //col:43
MaximumPacketsCapacity = 1000 //col:44
MaximumPacketsCapacityPriority = 10 //col:45
PacketChunkSize = 4096 //col:46
UsermodeBufferSize = unsafe.Sizeof(uint32(0)) + PacketChunkSize + 1 //col:47
MaxSerialPacketSize = UsermodeBufferSize + sizeof(DEBUGGER_REMOTE_PACKET) + SERIAL_END_OF_BUFFER_CHARS_COUNT //col:48
LogBufferSize = MaximumPacketsCapacity *(PacketChunkSize + sizeof(BUFFER_HEADER)) //col:51
LogBufferSizePriority = MaximumPacketsCapacityPriority *(PacketChunkSize + sizeof(BUFFER_HEADER)) //col:53
DbgPrintLimitation = 512 //col:55
DebuggerEventTagStartSeed = 0x1000000 //col:56
DebuggerThreadDebuggingTagStartSeed = 0x1000000 //col:57
DebuggerOutputSourceTagStartSeed = 0x1 //col:58
DebuggerOutputSourceMaximumRemoteSourceForSingleEvent = 0x5 //col:59
DEFAULT_PORT = "50000" //col:60
COMMUNICATION_BUFFER_SIZE = PacketChunkSize + 0x100 //col:61
OPERATION_MANDATORY_DEBUGGEE_BIT = (1 << 31) //col:62
OPERATION_LOG_INFO_MESSAGE =          0x1 //col:63
OPERATION_LOG_WARNING_MESSAGE =       0x2 //col:64
OPERATION_LOG_ERROR_MESSAGE =         0x3 //col:65
OPERATION_LOG_NON_IMMEDIATE_MESSAGE = 0x4 //col:66
OPERATION_LOG_WITH_TAG =              0x5 //col:67
OPERATION_COMMAND_FROM_DEBUGGER_CLOSE_AND_UNLOAD_VMM = 0x6 | OPERATION_MANDATORY_DEBUGGEE_BIT //col:68
OPERATION_DEBUGGEE_USER_INPUT =     0x7 | OPERATION_MANDATORY_DEBUGGEE_BIT //col:70
OPERATION_DEBUGGEE_REGISTER_EVENT = 0x8 | OPERATION_MANDATORY_DEBUGGEE_BIT //col:71
OPERATION_DEBUGGEE_ADD_ACTION_TO_EVENT = 0x9 | OPERATION_MANDATORY_DEBUGGEE_BIT //col:72
OPERATION_DEBUGGEE_CLEAR_EVENTS = 0xa | OPERATION_MANDATORY_DEBUGGEE_BIT //col:74
OPERATION_HYPERVISOR_DRIVER_IS_SUCCESSFULLY_LOADED = 0xb | OPERATION_MANDATORY_DEBUGGEE_BIT //col:75
OPERATION_HYPERVISOR_DRIVER_END_OF_IRPS = 0xc | OPERATION_MANDATORY_DEBUGGEE_BIT //col:77
OPERATION_COMMAND_FROM_DEBUGGER_RELOAD_SYMBOL = 0xd | OPERATION_MANDATORY_DEBUGGEE_BIT //col:79
OPERATION_NOTIFICATION_FROM_USER_DEBUGGER_PAUSE = 0xe | OPERATION_MANDATORY_DEBUGGEE_BIT //col:81
MAXIMUM_BREAKPOINTS_WITHOUT_CONTINUE = 50 //col:83
SERIAL_END_OF_BUFFER_CHARS_COUNT = 0x4 //col:84
SERIAL_END_OF_BUFFER_CHAR_1 = 0x00 //col:85
SERIAL_END_OF_BUFFER_CHAR_2 = 0x80 //col:86
SERIAL_END_OF_BUFFER_CHAR_3 = 0xEE //col:87
SERIAL_END_OF_BUFFER_CHAR_4 = 0xFF //col:88
TCP_END_OF_BUFFER_CHARS_COUNT = 0x4 //col:89
TCP_END_OF_BUFFER_CHAR_1 = 0x10 //col:90
TCP_END_OF_BUFFER_CHAR_2 = 0x20 //col:91
TCP_END_OF_BUFFER_CHAR_3 = 0x33 //col:92
TCP_END_OF_BUFFER_CHAR_4 = 0x44 //col:93
MAXIMUM_CHARACTER_FOR_OS_NAME = 256 //col:94
MAXIMUM_INSTR_SIZE = 16 //col:95
MAXIMUM_CALL_INSTR_SIZE = 7 //col:96
MAXIMUM_SUPPORTED_SYMBOLS = 1000 //col:97
MAXIMUM_GUID_AND_AGE_SIZE = 60 //col:98
INDICATOR_OF_HYPERDBG_PACKET = 0x4859504552444247 //col:99
MaximumSearchResults = 0x1000 //col:101
X86_FLAGS_CF =                 (1 << 0) //col:102
X86_FLAGS_PF =                 (1 << 2) //col:103
X86_FLAGS_AF =                 (1 << 4) //col:104
X86_FLAGS_ZF =                 (1 << 6) //col:105
X86_FLAGS_SF =                 (1 << 7) //col:106
X86_FLAGS_TF =                 (1 << 8) //col:107
X86_FLAGS_IF =                 (1 << 9) //col:108
X86_FLAGS_DF =                 (1 << 10) //col:109
X86_FLAGS_OF =                 (1 << 11) //col:110
X86_FLAGS_STATUS_MASK =        (0xfff) //col:111
X86_FLAGS_IOPL_MASK =          (3 << 12) //col:112
X86_FLAGS_IOPL_SHIFT =         (12) //col:113
X86_FLAGS_IOPL_SHIFT_2ND_BIT = (13) //col:114
X86_FLAGS_NT =                 (1 << 14) //col:115
X86_FLAGS_RF =                 (1 << 16) //col:116
X86_FLAGS_VM =                 (1 << 17) //col:117
X86_FLAGS_AC =                 (1 << 18) //col:118
X86_FLAGS_VIF =                (1 << 19) //col:119
X86_FLAGS_VIP =                (1 << 20) //col:120
X86_FLAGS_ID =                 (1 << 21) //col:121
X86_FLAGS_RESERVED_ONES =      0x2 //col:122
X86_FLAGS_RESERVED =           0xffc0802a //col:123
X86_FLAGS_RESERVED_BITS = 0xffc38028 //col:124
X86_FLAGS_FIXED =         0x00000002 //col:125
LOWORD(l) = ((WORD)(l)) //col:126
HIWORD(l) = ((WORD)(((DWORD)(l) >> 16) & 0xFFFF)) //col:127
LOBYTE(w) = ((BYTE)(w)) //col:128
HIBYTE(w) = ((BYTE)(((WORD)(w) >> 8) & 0xFF)) //col:129
MAX_TEMP_COUNT = 128 //col:130
MAX_VAR_COUNT = 512 //col:131
MAX_FUNCTION_NAME_LENGTH = 32 //col:132
DEBUGGER_MODIFY_EVENTS_APPLY_TO_ALL_TAG = 0xffffffffffffffff //col:133
DISASSEMBLY_MAXIMUM_DISTANCE_FROM_OBJECT_NAME = 0xffff //col:134
DEBUGGER_READ_AND_WRITE_ON_MSR_APPLY_ALL_CORES = 0xffffffff //col:135
DEBUGGER_DEBUGGEE_IS_RUNNING_NO_CORE = 0xffffffff //col:136
DEBUGGER_EVENT_APPLY_TO_ALL_CORES = 0xffffffff //col:137
DEBUGGER_EVENT_APPLY_TO_ALL_PROCESSES = 0xffffffff //col:138
DEBUGGER_EVENT_MSR_READ_OR_WRITE_ALL_MSRS = 0xffffffff //col:139
DEBUGGER_EVENT_EXCEPTIONS_ALL_FIRST_32_ENTRIES = 0xffffffff //col:140
DEBUGGER_EVENT_SYSCALL_ALL_SYSRET_OR_SYSCALLS = 0xffffffff //col:141
DEBUGGER_EVENT_ALL_IO_PORTS = 0xffffffff //col:142
DEBUGGEE_BP_APPLY_TO_ALL_CORES = 0xffffffff //col:143
DEBUGGEE_BP_APPLY_TO_ALL_PROCESSES = 0xffffffff //col:144
DEBUGGEE_BP_APPLY_TO_ALL_THREADS = 0xffffffff //col:145
DEBUGGEE_SHOW_ALL_REGISTERS = 0xffffffff //col:146
)


