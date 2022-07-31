package Headers
const(
VERSION_MAJOR = 0 //col:19
VERSION_MINOR = 2 //col:20
VERSION_PATCH = 0 //col:21
BUILD_YEAR_CH0 = (__DATE__[7]) //col:27
BUILD_YEAR_CH1 = (__DATE__[8]) //col:28
BUILD_YEAR_CH2 = (__DATE__[9]) //col:29
BUILD_YEAR_CH3 = (__DATE__[10]) //col:30
BUILD_MONTH_IS_JAN = (__DATE__[0] == 'J' && __DATE__[1] == 'a' && __DATE__[2] == 'n') //col:32
BUILD_MONTH_IS_FEB = (__DATE__[0] == 'F') //col:33
BUILD_MONTH_IS_MAR = (__DATE__[0] == 'M' && __DATE__[1] == 'a' && __DATE__[2] == 'r') //col:34
BUILD_MONTH_IS_APR = (__DATE__[0] == 'A' && __DATE__[1] == 'p') //col:35
BUILD_MONTH_IS_MAY = (__DATE__[0] == 'M' && __DATE__[1] == 'a' && __DATE__[2] == 'y') //col:36
BUILD_MONTH_IS_JUN = (__DATE__[0] == 'J' && __DATE__[1] == 'u' && __DATE__[2] == 'n') //col:37
BUILD_MONTH_IS_JUL = (__DATE__[0] == 'J' && __DATE__[1] == 'u' && __DATE__[2] == 'l') //col:38
BUILD_MONTH_IS_AUG = (__DATE__[0] == 'A' && __DATE__[1] == 'u') //col:39
BUILD_MONTH_IS_SEP = (__DATE__[0] == 'S') //col:40
BUILD_MONTH_IS_OCT = (__DATE__[0] == 'O') //col:41
BUILD_MONTH_IS_NOV = (__DATE__[0] == 'N') //col:42
BUILD_MONTH_IS_DEC = (__DATE__[0] == 'D') //col:43
BUILD_MONTH_CH0 = ((BUILD_MONTH_IS_OCT || BUILD_MONTH_IS_NOV || BUILD_MONTH_IS_DEC) ? '1' : '0') //col:45
BUILD_MONTH_CH1 = ( (BUILD_MONTH_IS_JAN) ? '1' : (BUILD_MONTH_IS_FEB) ? '2' : (BUILD_MONTH_IS_MAR)   ? '3' : (BUILD_MONTH_IS_APR)   ? '4' : (BUILD_MONTH_IS_MAY)   ? '5' : (BUILD_MONTH_IS_JUN)   ? '6' : (BUILD_MONTH_IS_JUL)   ? '7' : (BUILD_MONTH_IS_AUG)   ? '8' : (BUILD_MONTH_IS_SEP)   ? '9' : (BUILD_MONTH_IS_OCT)   ? '0' : (BUILD_MONTH_IS_NOV)   ? '1' : (BUILD_MONTH_IS_DEC)   ? '2' : /* error default */ '?') //col:48
BUILD_DAY_CH0 = ((__DATE__[4] >= '0') ? (__DATE__[4]) : '0') //col:63
BUILD_DAY_CH1 = (__DATE__[5]) //col:64
BUILD_HOUR_CH0 = (__TIME__[0]) //col:70
BUILD_HOUR_CH1 = (__TIME__[1]) //col:71
BUILD_MIN_CH0 = (__TIME__[3]) //col:73
BUILD_MIN_CH1 = (__TIME__[4]) //col:74
BUILD_SEC_CH0 = (__TIME__[6]) //col:76
BUILD_SEC_CH1 = (__TIME__[7]) //col:77
MaximumPacketsCapacity = 1000 //col:198
MaximumPacketsCapacityPriority = 10 //col:204
PacketChunkSize = 4096 // PAGE_SIZE //col:209
UsermodeBufferSize = unsafe.Sizeof(uint32(0)) + PacketChunkSize + 1 //col:217
MaxSerialPacketSize = UsermodeBufferSize + sizeof(DEBUGGER_REMOTE_PACKET) + SERIAL_END_OF_BUFFER_CHARS_COUNT //col:225
LogBufferSize = MaximumPacketsCapacity *(PacketChunkSize + sizeof(BUFFER_HEADER)) //col:233
LogBufferSizePriority = MaximumPacketsCapacityPriority *(PacketChunkSize + sizeof(BUFFER_HEADER)) //col:240
DbgPrintLimitation = 512 //col:248
DebuggerEventTagStartSeed = 0x1000000 //col:255
DebuggerThreadDebuggingTagStartSeed = 0x1000000 //col:262
DebuggerOutputSourceTagStartSeed = 0x1 //col:269
DebuggerOutputSourceMaximumRemoteSourceForSingleEvent = 0x5 //col:276
DEFAULT_PORT = "50000" //col:287
COMMUNICATION_BUFFER_SIZE = PacketChunkSize + 0x100 //col:294
OPERATION_MANDATORY_DEBUGGEE_BIT = (1 << 31) //col:305
OPERATION_LOG_INFO_MESSAGE =          0x1 //col:312
OPERATION_LOG_WARNING_MESSAGE =       0x2 //col:313
OPERATION_LOG_ERROR_MESSAGE =         0x3 //col:314
OPERATION_LOG_NON_IMMEDIATE_MESSAGE = 0x4 //col:315
OPERATION_LOG_WITH_TAG =              0x5 //col:316
OPERATION_COMMAND_FROM_DEBUGGER_CLOSE_AND_UNLOAD_VMM = 0x6 | OPERATION_MANDATORY_DEBUGGEE_BIT //col:318
OPERATION_DEBUGGEE_USER_INPUT =     0x7 | OPERATION_MANDATORY_DEBUGGEE_BIT //col:320
OPERATION_DEBUGGEE_REGISTER_EVENT = 0x8 | OPERATION_MANDATORY_DEBUGGEE_BIT //col:321
OPERATION_DEBUGGEE_ADD_ACTION_TO_EVENT = 0x9 | OPERATION_MANDATORY_DEBUGGEE_BIT //col:322
OPERATION_DEBUGGEE_CLEAR_EVENTS = 0xa | OPERATION_MANDATORY_DEBUGGEE_BIT //col:324
OPERATION_HYPERVISOR_DRIVER_IS_SUCCESSFULLY_LOADED = 0xb | OPERATION_MANDATORY_DEBUGGEE_BIT //col:325
OPERATION_HYPERVISOR_DRIVER_END_OF_IRPS = 0xc | OPERATION_MANDATORY_DEBUGGEE_BIT //col:327
OPERATION_COMMAND_FROM_DEBUGGER_RELOAD_SYMBOL = 0xd | OPERATION_MANDATORY_DEBUGGEE_BIT //col:329
OPERATION_NOTIFICATION_FROM_USER_DEBUGGER_PAUSE = 0xe | OPERATION_MANDATORY_DEBUGGEE_BIT //col:332
MAXIMUM_BREAKPOINTS_WITHOUT_CONTINUE = 50 //col:343
SERIAL_END_OF_BUFFER_CHARS_COUNT = 0x4 //col:352
SERIAL_END_OF_BUFFER_CHAR_1 = 0x00 //col:358
SERIAL_END_OF_BUFFER_CHAR_2 = 0x80 //col:359
SERIAL_END_OF_BUFFER_CHAR_3 = 0xEE //col:360
SERIAL_END_OF_BUFFER_CHAR_4 = 0xFF //col:361
TCP_END_OF_BUFFER_CHARS_COUNT = 0x4 //col:366
TCP_END_OF_BUFFER_CHAR_1 = 0x10 //col:372
TCP_END_OF_BUFFER_CHAR_2 = 0x20 //col:373
TCP_END_OF_BUFFER_CHAR_3 = 0x33 //col:374
TCP_END_OF_BUFFER_CHAR_4 = 0x44 //col:375
MAXIMUM_CHARACTER_FOR_OS_NAME = 256 //col:385
MAXIMUM_INSTR_SIZE = 16 //col:394
MAXIMUM_CALL_INSTR_SIZE = 7 //col:399
MAXIMUM_SUPPORTED_SYMBOLS = 1000 //col:409
MAXIMUM_GUID_AND_AGE_SIZE = 60 //col:417
INDICATOR_OF_HYPERDBG_PACKET = 0x4859504552444247 // HYPERDBG = 0x4859504552444247 //col:423
MaximumSearchResults = 0x1000 //col:435
X86_FLAGS_CF =                 (1 << 0) //col:445
X86_FLAGS_PF =                 (1 << 2) //col:446
X86_FLAGS_AF =                 (1 << 4) //col:447
X86_FLAGS_ZF =                 (1 << 6) //col:448
X86_FLAGS_SF =                 (1 << 7) //col:449
X86_FLAGS_TF =                 (1 << 8) //col:450
X86_FLAGS_IF =                 (1 << 9) //col:451
X86_FLAGS_DF =                 (1 << 10) //col:452
X86_FLAGS_OF =                 (1 << 11) //col:453
X86_FLAGS_STATUS_MASK =        (0xfff) //col:454
X86_FLAGS_IOPL_MASK =          (3 << 12) //col:455
X86_FLAGS_IOPL_SHIFT =         (12) //col:456
X86_FLAGS_IOPL_SHIFT_2ND_BIT = (13) //col:457
X86_FLAGS_NT =                 (1 << 14) //col:458
X86_FLAGS_RF =                 (1 << 16) //col:459
X86_FLAGS_VM =                 (1 << 17) //col:460
X86_FLAGS_AC =                 (1 << 18) //col:461
X86_FLAGS_VIF =                (1 << 19) //col:462
X86_FLAGS_VIP =                (1 << 20) //col:463
X86_FLAGS_ID =                 (1 << 21) //col:464
X86_FLAGS_RESERVED_ONES =      0x2 //col:465
X86_FLAGS_RESERVED =           0xffc0802a //col:466
X86_FLAGS_RESERVED_BITS = 0xffc38028 //col:468
X86_FLAGS_FIXED =         0x00000002 //col:469
LOWORD(l) = ((WORD)(l)) //col:471
HIWORD(l) = ((WORD)(((DWORD)(l) >> 16) & 0xFFFF)) //col:472
LOBYTE(w) = ((BYTE)(w)) //col:473
HIBYTE(w) = ((BYTE)(((WORD)(w) >> 8) & 0xFF)) //col:474
MAX_TEMP_COUNT = 128 //col:476
MAX_VAR_COUNT = 512 //col:480
MAX_FUNCTION_NAME_LENGTH = 32 //col:482
DEBUGGER_MODIFY_EVENTS_APPLY_TO_ALL_TAG = 0xffffffffffffffff //col:492
DISASSEMBLY_MAXIMUM_DISTANCE_FROM_OBJECT_NAME = 0xffff //col:499
DEBUGGER_READ_AND_WRITE_ON_MSR_APPLY_ALL_CORES = 0xffffffff //col:505
DEBUGGER_DEBUGGEE_IS_RUNNING_NO_CORE = 0xffffffff //col:511
DEBUGGER_EVENT_APPLY_TO_ALL_CORES = 0xffffffff //col:517
DEBUGGER_EVENT_APPLY_TO_ALL_PROCESSES = 0xffffffff //col:523
DEBUGGER_EVENT_MSR_READ_OR_WRITE_ALL_MSRS = 0xffffffff //col:529
DEBUGGER_EVENT_EXCEPTIONS_ALL_FIRST_32_ENTRIES = 0xffffffff //col:535
DEBUGGER_EVENT_SYSCALL_ALL_SYSRET_OR_SYSCALLS = 0xffffffff //col:541
DEBUGGER_EVENT_ALL_IO_PORTS = 0xffffffff //col:547
DEBUGGEE_BP_APPLY_TO_ALL_CORES = 0xffffffff //col:553
DEBUGGEE_BP_APPLY_TO_ALL_PROCESSES = 0xffffffff //col:559
DEBUGGEE_BP_APPLY_TO_ALL_THREADS = 0xffffffff //col:565
DEBUGGEE_SHOW_ALL_REGISTERS = 0xffffffff //col:571
)
