#pragma once
#define VERSION_MAJOR      0
#define VERSION_MINOR      2
#define VERSION_PATCH      0
#define BUILD_YEAR_CH0     (__DATE__[7])
#define BUILD_YEAR_CH1     (__DATE__[8])
#define BUILD_YEAR_CH2     (__DATE__[9])
#define BUILD_YEAR_CH3     (__DATE__[10])
#define BUILD_MONTH_IS_JAN (__DATE__[0] == 'J' && __DATE__[1] == 'a' && __DATE__[2] == 'n')
#define BUILD_MONTH_IS_FEB (__DATE__[0] == 'F')
#define BUILD_MONTH_IS_MAR (__DATE__[0] == 'M' && __DATE__[1] == 'a' && __DATE__[2] == 'r')
#define BUILD_MONTH_IS_APR (__DATE__[0] == 'A' && __DATE__[1] == 'p')
#define BUILD_MONTH_IS_MAY (__DATE__[0] == 'M' && __DATE__[1] == 'a' && __DATE__[2] == 'y')
#define BUILD_MONTH_IS_JUN (__DATE__[0] == 'J' && __DATE__[1] == 'u' && __DATE__[2] == 'n')
#define BUILD_MONTH_IS_JUL (__DATE__[0] == 'J' && __DATE__[1] == 'u' && __DATE__[2] == 'l')
#define BUILD_MONTH_IS_AUG (__DATE__[0] == 'A' && __DATE__[1] == 'u')
#define BUILD_MONTH_IS_SEP (__DATE__[0] == 'S')
#define BUILD_MONTH_IS_OCT (__DATE__[0] == 'O')
#define BUILD_MONTH_IS_NOV (__DATE__[0] == 'N')
#define BUILD_MONTH_IS_DEC (__DATE__[0] == 'D')
#define BUILD_MONTH_CH0 \
    ((BUILD_MONTH_IS_OCT || BUILD_MONTH_IS_NOV || BUILD_MONTH_IS_DEC) ? '1' : '0')
#define BUILD_MONTH_CH1                                         \
    (                                                           \
        (BUILD_MONTH_IS_JAN) ? '1' : (BUILD_MONTH_IS_FEB) ? '2' \
                                 : (BUILD_MONTH_IS_MAR)   ? '3' \
                                 : (BUILD_MONTH_IS_APR)   ? '4' \
                                 : (BUILD_MONTH_IS_MAY)   ? '5' \
                                 : (BUILD_MONTH_IS_JUN)   ? '6' \
                                 : (BUILD_MONTH_IS_JUL)   ? '7' \
                                 : (BUILD_MONTH_IS_AUG)   ? '8' \
                                 : (BUILD_MONTH_IS_SEP)   ? '9' \
                                 : (BUILD_MONTH_IS_OCT)   ? '0' \
                                 : (BUILD_MONTH_IS_NOV)   ? '1' \
                                 : (BUILD_MONTH_IS_DEC)   ? '2' \
                                                          : /* error default */ '?')
#define BUILD_DAY_CH0  ((__DATE__[4] >= '0') ? (__DATE__[4]) : '0')
#define BUILD_DAY_CH1  (__DATE__[5])
#define BUILD_HOUR_CH0 (__TIME__[0])
#define BUILD_HOUR_CH1 (__TIME__[1])
#define BUILD_MIN_CH0  (__TIME__[3])
#define BUILD_MIN_CH1  (__TIME__[4])
#define BUILD_SEC_CH0  (__TIME__[6])
#define BUILD_SEC_CH1  (__TIME__[7])
#if VERSION_MAJOR > 100
#    define VERSION_MAJOR_INIT                    \
        ((VERSION_MAJOR / 100) + '0'),            \
            (((VERSION_MAJOR % 100) / 10) + '0'), \
            ((VERSION_MAJOR % 10) + '0')
#elif VERSION_MAJOR > 10
#    define VERSION_MAJOR_INIT        \
        ((VERSION_MAJOR / 10) + '0'), \
            ((VERSION_MAJOR % 10) + '0')
#else
#    define VERSION_MAJOR_INIT \
        (VERSION_MAJOR + '0')
#endif
#if VERSION_MINOR > 100
#    define VERSION_MINOR_INIT                    \
        ((VERSION_MINOR / 100) + '0'),            \
            (((VERSION_MINOR % 100) / 10) + '0'), \
            ((VERSION_MINOR % 10) + '0')
#elif VERSION_MINOR > 10
#    define VERSION_MINOR_INIT        \
        ((VERSION_MINOR / 10) + '0'), \
            ((VERSION_MINOR % 10) + '0')
#else
#    define VERSION_MINOR_INIT \
        (VERSION_MINOR + '0')
#endif
#if VERSION_PATCH > 100
#    define VERSION_PATCH_INIT                    \
        ((VERSION_PATCH / 100) + '0'),            \
            (((VERSION_PATCH % 100) / 10) + '0'), \
            ((VERSION_PATCH % 10) + '0')
#elif VERSION_PATCH > 10
#    define VERSION_PATCH_INIT        \
        ((VERSION_PATCH / 10) + '0'), \
            ((VERSION_PATCH % 10) + '0')
#else
#    define VERSION_PATCH_INIT \
        (VERSION_PATCH + '0')
#endif
#ifndef HYPERDBG_KERNEL_MODE
const unsigned char BuildDateTime[] =
    {
        BUILD_YEAR_CH0,
        BUILD_YEAR_CH1,
        BUILD_YEAR_CH2,
        BUILD_YEAR_CH3,
        '-',
        BUILD_MONTH_CH0,
        BUILD_MONTH_CH1,
        '-',
        BUILD_DAY_CH0,
        BUILD_DAY_CH1,
        ' ',
        BUILD_HOUR_CH0,
        BUILD_HOUR_CH1,
        ':',
        BUILD_MIN_CH0,
        BUILD_MIN_CH1,
        ':',
        BUILD_SEC_CH0,
        BUILD_SEC_CH1,
        '\0'};
const unsigned char CompleteVersion[] =
    {
        'v',
        VERSION_MAJOR_INIT,
        '.',
        VERSION_MINOR_INIT,
        '.',
        VERSION_PATCH_INIT,
        '\0'};
const unsigned char BuildVersion[] =
    {
        BUILD_YEAR_CH0,
        BUILD_YEAR_CH1,
        BUILD_YEAR_CH2,
        BUILD_YEAR_CH3,
        BUILD_MONTH_CH0,
        BUILD_MONTH_CH1,
        BUILD_DAY_CH0,
        BUILD_DAY_CH1,
        '\0'};
#endif // SCRIPT_ENGINE_KERNEL_MODE
#define MaximumPacketsCapacity         1000
#define MaximumPacketsCapacityPriority 10
#define PacketChunkSize                4096 // PAGE_SIZE
#define UsermodeBufferSize             sizeof(UINT32) + PacketChunkSize + 1
#define MaxSerialPacketSize                               \
    UsermodeBufferSize + sizeof(DEBUGGER_REMOTE_PACKET) + \
        SERIAL_END_OF_BUFFER_CHARS_COUNT
#define LogBufferSize \
    MaximumPacketsCapacity *(PacketChunkSize + sizeof(BUFFER_HEADER))
#define LogBufferSizePriority \
    MaximumPacketsCapacityPriority *(PacketChunkSize + sizeof(BUFFER_HEADER))
#define DbgPrintLimitation                                    512
#define DebuggerEventTagStartSeed                             0x1000000
#define DebuggerThreadDebuggingTagStartSeed                   0x1000000
#define DebuggerOutputSourceTagStartSeed                      0x1
#define DebuggerOutputSourceMaximumRemoteSourceForSingleEvent 0x5
#define DEFAULT_PORT                                          "50000"
#define COMMUNICATION_BUFFER_SIZE                             PacketChunkSize + 0x100
#define OPERATION_MANDATORY_DEBUGGEE_BIT                      (1 << 31)
#define OPERATION_LOG_INFO_MESSAGE                            0x1
#define OPERATION_LOG_WARNING_MESSAGE                         0x2
#define OPERATION_LOG_ERROR_MESSAGE                           0x3
#define OPERATION_LOG_NON_IMMEDIATE_MESSAGE                   0x4
#define OPERATION_LOG_WITH_TAG                                0x5
#define OPERATION_COMMAND_FROM_DEBUGGER_CLOSE_AND_UNLOAD_VMM \
    0x6 | OPERATION_MANDATORY_DEBUGGEE_BIT
#define OPERATION_DEBUGGEE_USER_INPUT     0x7 | OPERATION_MANDATORY_DEBUGGEE_BIT
#define OPERATION_DEBUGGEE_REGISTER_EVENT 0x8 | OPERATION_MANDATORY_DEBUGGEE_BIT
#define OPERATION_DEBUGGEE_ADD_ACTION_TO_EVENT \
    0x9 | OPERATION_MANDATORY_DEBUGGEE_BIT
#define OPERATION_DEBUGGEE_CLEAR_EVENTS 0xa | OPERATION_MANDATORY_DEBUGGEE_BIT
#define OPERATION_HYPERVISOR_DRIVER_IS_SUCCESSFULLY_LOADED \
    0xb | OPERATION_MANDATORY_DEBUGGEE_BIT
#define OPERATION_HYPERVISOR_DRIVER_END_OF_IRPS \
    0xc | OPERATION_MANDATORY_DEBUGGEE_BIT
#define OPERATION_COMMAND_FROM_DEBUGGER_RELOAD_SYMBOL \
    0xd | OPERATION_MANDATORY_DEBUGGEE_BIT
#define OPERATION_NOTIFICATION_FROM_USER_DEBUGGER_PAUSE \
    0xe | OPERATION_MANDATORY_DEBUGGEE_BIT
#define MAXIMUM_BREAKPOINTS_WITHOUT_CONTINUE 50
#define SERIAL_END_OF_BUFFER_CHARS_COUNT     0x4
#define SERIAL_END_OF_BUFFER_CHAR_1          0x00
#define SERIAL_END_OF_BUFFER_CHAR_2          0x80
#define SERIAL_END_OF_BUFFER_CHAR_3          0xEE
#define SERIAL_END_OF_BUFFER_CHAR_4          0xFF
#define TCP_END_OF_BUFFER_CHARS_COUNT        0x4
#define TCP_END_OF_BUFFER_CHAR_1             0x10
#define TCP_END_OF_BUFFER_CHAR_2             0x20
#define TCP_END_OF_BUFFER_CHAR_3             0x33
#define TCP_END_OF_BUFFER_CHAR_4             0x44
#define MAXIMUM_CHARACTER_FOR_OS_NAME        256
#define MAXIMUM_INSTR_SIZE                   16
#define MAXIMUM_CALL_INSTR_SIZE              7
#define MAXIMUM_SUPPORTED_SYMBOLS            1000
#define MAXIMUM_GUID_AND_AGE_SIZE            60
#define INDICATOR_OF_HYPERDBG_PACKET \
    0x4859504552444247 // HYPERDBG = 0x4859504552444247
#define MaximumSearchResults                           0x1000
#define X86_FLAGS_CF                                   (1 << 0)
#define X86_FLAGS_PF                                   (1 << 2)
#define X86_FLAGS_AF                                   (1 << 4)
#define X86_FLAGS_ZF                                   (1 << 6)
#define X86_FLAGS_SF                                   (1 << 7)
#define X86_FLAGS_TF                                   (1 << 8)
#define X86_FLAGS_IF                                   (1 << 9)
#define X86_FLAGS_DF                                   (1 << 10)
#define X86_FLAGS_OF                                   (1 << 11)
#define X86_FLAGS_STATUS_MASK                          (0xfff)
#define X86_FLAGS_IOPL_MASK                            (3 << 12)
#define X86_FLAGS_IOPL_SHIFT                           (12)
#define X86_FLAGS_IOPL_SHIFT_2ND_BIT                   (13)
#define X86_FLAGS_NT                                   (1 << 14)
#define X86_FLAGS_RF                                   (1 << 16)
#define X86_FLAGS_VM                                   (1 << 17)
#define X86_FLAGS_AC                                   (1 << 18)
#define X86_FLAGS_VIF                                  (1 << 19)
#define X86_FLAGS_VIP                                  (1 << 20)
#define X86_FLAGS_ID                                   (1 << 21)
#define X86_FLAGS_RESERVED_ONES                        0x2
#define X86_FLAGS_RESERVED                             0xffc0802a
#define X86_FLAGS_RESERVED_BITS                        0xffc38028
#define X86_FLAGS_FIXED                                0x00000002
#define LOWORD(l)                                      ((WORD)(l))
#define HIWORD(l)                                      ((WORD)(((DWORD)(l) >> 16) & 0xFFFF))
#define LOBYTE(w)                                      ((BYTE)(w))
#define HIBYTE(w)                                      ((BYTE)(((WORD)(w) >> 8) & 0xFF))
#define MAX_TEMP_COUNT                                 128
#define MAX_VAR_COUNT                                  512
#define MAX_FUNCTION_NAME_LENGTH                       32
#define DEBUGGER_MODIFY_EVENTS_APPLY_TO_ALL_TAG        0xffffffffffffffff
#define DISASSEMBLY_MAXIMUM_DISTANCE_FROM_OBJECT_NAME  0xffff
#define DEBUGGER_READ_AND_WRITE_ON_MSR_APPLY_ALL_CORES 0xffffffff
#define DEBUGGER_DEBUGGEE_IS_RUNNING_NO_CORE           0xffffffff
#define DEBUGGER_EVENT_APPLY_TO_ALL_CORES              0xffffffff
#define DEBUGGER_EVENT_APPLY_TO_ALL_PROCESSES          0xffffffff
#define DEBUGGER_EVENT_MSR_READ_OR_WRITE_ALL_MSRS      0xffffffff
#define DEBUGGER_EVENT_EXCEPTIONS_ALL_FIRST_32_ENTRIES 0xffffffff
#define DEBUGGER_EVENT_SYSCALL_ALL_SYSRET_OR_SYSCALLS  0xffffffff
#define DEBUGGER_EVENT_ALL_IO_PORTS                    0xffffffff
#define DEBUGGEE_BP_APPLY_TO_ALL_CORES                 0xffffffff
#define DEBUGGEE_BP_APPLY_TO_ALL_PROCESSES             0xffffffff
#define DEBUGGEE_BP_APPLY_TO_ALL_THREADS               0xffffffff
#define DEBUGGEE_SHOW_ALL_REGISTERS                    0xffffffff
