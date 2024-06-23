//bugfix.h
#ifndef size_t
#define size_t int
#endif

#ifndef wchar_t
typedef int rune;
#define wchar_t rune
#endif

#ifndef _WINT_T_DEFINED_
 #define _WINT_T_DEFINED_
 typedef int wint_t;
#endif

#undef WCHAR_MIN
#undef WCHAR_MAX
#define WCHAR_MIN   0
#define WCHAR_MAX   65535

//typedef unsigned short wchar_t;

typedef int bool ;
#define PVOID void*
#define HANDLE void*
#define PIRP void*//todo
#define PDEVICE_OBJECT void*//todo
#define PSYMBOL_BUFFER void*//todo
#define PSYMBOL void*//todo
#define MAX_PATH 260
typedef unsigned __int64   SIZE_T,*PSIZE_T;
typedef unsigned __int64   time_t;
typedef unsigned __int64   NTSTATUS;
typedef char *  va_list;

typedef struct _LIST_ENTRY {
  struct _LIST_ENTRY *Flink;
  struct _LIST_ENTRY *Blink;
} LIST_ENTRY, *PLIST_ENTRY, PRLIST_ENTRY;

#ifndef _In_
#define _In_
#endif

#ifndef _Out_
#define _Out_
#endif

#ifndef _Inout_
#define _Inout_
#endif

#ifndef _Out_writes_bytes_
#define _Out_writes_bytes_(x)
#endif

#ifndef _In_reads_
#define _In_reads_(x)
#endif

#ifndef _In_reads_bytes_
#define _In_reads_bytes_(x)
#endif

/*
typedef struct _IRP {
  CSHORT                    Type;
  USHORT                    Size;
  PMDL                      MdlAddress;
  ULONG                     Flags;
  union {
    struct _IRP     *MasterIrp;
    __volatile LONG IrpCount;
    PVOID           SystemBuffer;
  } AssociatedIrp;
  LIST_ENTRY                ThreadListEntry;
  IO_STATUS_BLOCK           IoStatus;
  KPROCESSOR_MODE           RequestorMode;
  BOOLEAN                   PendingReturned;
  CHAR                      StackCount;
  CHAR                      CurrentLocation;
  BOOLEAN                   Cancel;
  KIRQL                     CancelIrql;
  CCHAR                     ApcEnvironment;
  UCHAR                     AllocationFlags;
  union {
    PIO_STATUS_BLOCK UserIosb;
    PVOID            IoRingContext;
  };
  PKEVENT                   UserEvent;
  union {
    struct {
      union {
        PIO_APC_ROUTINE UserApcRoutine;
        PVOID           IssuingProcess;
      };
      union {
        PVOID                 UserApcContext;
#if ...
        _IORING_OBJECT        *IoRing;
#else
        struct _IORING_OBJECT *IoRing;
#endif
      };
    } AsynchronousParameters;
    LARGE_INTEGER AllocationSize;
  } Overlay;
  __volatile PDRIVER_CANCEL CancelRoutine;
  PVOID                     UserBuffer;
  union {
    struct {
      union {
        KDEVICE_QUEUE_ENTRY DeviceQueueEntry;
        struct {
          PVOID DriverContext[4];
        };
      };
      PETHREAD     Thread;
      PCHAR        AuxiliaryBuffer;
      struct {
        LIST_ENTRY ListEntry;
        union {
          struct _IO_STACK_LOCATION *CurrentStackLocation;
          ULONG                     PacketType;
        };
      };
      PFILE_OBJECT OriginalFileObject;
    } Overlay;
    KAPC  Apc;
    PVOID CompletionKey;
  } Tail;
} IRP;
*/


//..\bin\debug\SDK\Headers\BasicTypes.h
/**
 * @file BasicTypes.h
 * @author Sina Karvandi (sina@hyperdbg.org)
 * @brief HyperDbg's SDK Headers For Basic Datatypes
 * @details This file contains definitions of basic datatypes
 * @version 0.2
 * @date 2022-06-28
 *
 * @copyright This project is released under the GNU Public License v3.
 *
 */
#pragma once

#pragma warning(disable : 4201) // Suppress nameless struct/union warning

//////////////////////////////////////////////////
//               Basic Datatypes                //
//////////////////////////////////////////////////

typedef unsigned long long QWORD;
typedef unsigned __int64   UINT64, *PUINT64;
typedef unsigned long      DWORD;
typedef int                BOOL;
typedef unsigned char      BYTE;
typedef unsigned short     WORD;
typedef int                INT;
typedef unsigned int       UINT;
typedef unsigned int *     PUINT;
typedef unsigned __int64   ULONG64, *PULONG64;
typedef unsigned __int64   DWORD64, *PDWORD64;
typedef char               CHAR;
typedef wchar_t            WCHAR;
#define VOID void

typedef unsigned char  UCHAR;
typedef unsigned short USHORT;
typedef unsigned long  ULONG;

typedef UCHAR     BOOLEAN;  // winnt
typedef BOOLEAN * PBOOLEAN; // winnt

typedef signed char      INT8, *PINT8;
typedef signed short     INT16, *PINT16;
typedef signed int       INT32, *PINT32;
typedef signed __int64   INT64, *PINT64;
typedef unsigned char    UINT8, *PUINT8;
typedef unsigned short   UINT16, *PUINT16;
typedef unsigned int     UINT32, *PUINT32;
typedef unsigned __int64 UINT64, *PUINT64;

#define NULL_ZERO   0
#define NULL64_ZERO 0ull

#define FALSE 0
#define TRUE  1

#define UPPER_56_BITS                  0xffffffffffffff00
#define UPPER_48_BITS                  0xffffffffffff0000
#define UPPER_32_BITS                  0xffffffff00000000
#define LOWER_32_BITS                  0x00000000ffffffff
#define LOWER_16_BITS                  0x000000000000ffff
#define LOWER_8_BITS                   0x00000000000000ff
#define SECOND_LOWER_8_BITS            0x000000000000ff00
#define UPPER_48_BITS_AND_LOWER_8_BITS 0xffffffffffff00ff

//
// DO NOT FUCKING TOUCH THIS STRUCTURE WITHOUT COORDINATION WITH SINA
//
typedef struct GUEST_REGS
{
    //
    // DO NOT FUCKING TOUCH THIS STRUCTURE WITHOUT COORDINATION WITH SINA
    //

    UINT64 rax; // 0x00
    UINT64 rcx; // 0x08
    UINT64 rdx; // 0x10
    UINT64 rbx; // 0x18
    UINT64 rsp; // 0x20
    UINT64 rbp; // 0x28
    UINT64 rsi; // 0x30
    UINT64 rdi; // 0x38
    UINT64 r8;  // 0x40
    UINT64 r9;  // 0x48
    UINT64 r10; // 0x50
    UINT64 r11; // 0x58
    UINT64 r12; // 0x60
    UINT64 r13; // 0x68
    UINT64 r14; // 0x70
    UINT64 r15; // 0x78

    //
    // DO NOT FUCKING TOUCH THIS STRUCTURE WITHOUT COORDINATION WITH SINA
    //

} GUEST_REGS, *PGUEST_REGS;

/**
 * @brief struct for extra registers
 *
 */
typedef struct GUEST_EXTRA_REGISTERS
{
    UINT16 CS;
    UINT16 DS;
    UINT16 FS;
    UINT16 GS;
    UINT16 ES;
    UINT16 SS;
    UINT64 RFLAGS;
    UINT64 RIP;
} GUEST_EXTRA_REGISTERS, *PGUEST_EXTRA_REGISTERS;

/**
 * @brief List of different variables
 */
typedef struct _SCRIPT_ENGINE_VARIABLES_LIST
{
    UINT64 * TempList;
    UINT64 * GlobalVariablesList;
    UINT64 * LocalVariablesList;

} SCRIPT_ENGINE_VARIABLES_LIST, *PSCRIPT_ENGINE_VARIABLES_LIST;

/**
 * @brief CR3 Structure
 *
 */
typedef struct _CR3_TYPE
{
    union
    {
        UINT64 Flags;

        struct
        {
            UINT64 Pcid : 12;
            UINT64 PageFrameNumber : 36;
            UINT64 Reserved1 : 12;
            UINT64 Reserved_2 : 3;
            UINT64 PcidInvalidate : 1;
        } Fields;
    };
} CR3_TYPE, *PCR3_TYPE;


//..\bin\debug\SDK\Headers\Connection.h
/**
 * @file Connection.h
 * @author Sina Karvandi (sina@hyperdbg.org)
 * @brief HyperDbg's SDK Headers For Native Structures, Enums and Constants
 * @details These datatypes are used in all devices like HDL (FPGAs)
 * @version 0.2
 * @date 2022-07-14
 *
 * @copyright This project is released under the GNU Public License v3.
 *
 */
#pragma once

/**
 * @brief enum for reasons why debuggee is paused
 *
 */
typedef enum _DEBUGGEE_PAUSING_REASON
{

    //
    // For both kernel & user debugger
    //
    DEBUGGEE_PAUSING_REASON_NOT_PAUSED = 0,
    DEBUGGEE_PAUSING_REASON_PAUSE,
    DEBUGGEE_PAUSING_REASON_REQUEST_FROM_DEBUGGER,
    DEBUGGEE_PAUSING_REASON_DEBUGGEE_STEPPED,
    DEBUGGEE_PAUSING_REASON_DEBUGGEE_TRACKING_STEPPED,
    DEBUGGEE_PAUSING_REASON_DEBUGGEE_SOFTWARE_BREAKPOINT_HIT,
    DEBUGGEE_PAUSING_REASON_DEBUGGEE_HARDWARE_DEBUG_REGISTER_HIT,
    DEBUGGEE_PAUSING_REASON_DEBUGGEE_CORE_SWITCHED,
    DEBUGGEE_PAUSING_REASON_DEBUGGEE_PROCESS_SWITCHED,
    DEBUGGEE_PAUSING_REASON_DEBUGGEE_THREAD_SWITCHED,
    DEBUGGEE_PAUSING_REASON_DEBUGGEE_COMMAND_EXECUTION_FINISHED,
    DEBUGGEE_PAUSING_REASON_DEBUGGEE_EVENT_TRIGGERED,
    DEBUGGEE_PAUSING_REASON_DEBUGGEE_STARTING_MODULE_LOADED,

    //
    // Only for user-debugger
    //
    DEBUGGEE_PAUSING_REASON_DEBUGGEE_GENERAL_DEBUG_BREAK,
    DEBUGGEE_PAUSING_REASON_DEBUGGEE_GENERAL_THREAD_INTERCEPTED,

    //
    // Only used for hardware debugging
    //
    DEBUGGEE_PAUSING_REASON_HARDWARE_BASED_DEBUGGEE_GENERAL_BREAK,

} DEBUGGEE_PAUSING_REASON;

/**
 * @brief enum for requested action for HyperDbg packet
 *
 */
typedef enum _DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION
{

    //
    // Debugger to debuggee (user-mode execution)
    //
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_USER_MODE_PAUSE = 1,
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_USER_MODE_DO_NOT_READ_ANY_PACKET,
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_USER_MODE_DEBUGGER_VERSION,

    //
    // Debuggee to debugger (user-mode execution)
    //
    DEBUGGER_REMOTE_PACKET_PING_AND_SEND_SUPPORTED_VERSION,

    //
    // Debugger to debuggee (vmx-root mode execution)
    //
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_MODE_STEP,
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_MODE_CONTINUE,
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_MODE_CLOSE_AND_UNLOAD_DEBUGGEE,
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_MODE_CHANGE_CORE,
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_MODE_FLUSH_BUFFERS,
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_MODE_CALLSTACK,
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_MODE_TEST_QUERY,
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_MODE_CHANGE_PROCESS,
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_MODE_CHANGE_THREAD,
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_RUN_SCRIPT,
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_USER_INPUT_BUFFER,
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_SEARCH_QUERY,
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_REGISTER_EVENT,
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_ADD_ACTION_TO_EVENT,
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_QUERY_AND_MODIFY_EVENT,
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_READ_REGISTERS,
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_READ_MEMORY,
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_EDIT_MEMORY,
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_BP,
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_LIST_OR_MODIFY_BREAKPOINTS,
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_SYMBOL_RELOAD,
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_QUERY_PA2VA_AND_VA2PA,
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_SYMBOL_QUERY_PTE,
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_SET_SHORT_CIRCUITING_STATE,
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_INJECT_PAGE_FAULT,

    //
    // Debuggee to debugger
    //
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_NO_ACTION,
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_STARTED,
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_LOGGING_MECHANISM,
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_PAUSED_AND_CURRENT_INSTRUCTION,
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_CHANGING_CORE,
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_CHANGING_PROCESS,
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_CHANGING_THREAD,
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_RUNNING_SCRIPT,
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_FORMATS,
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_FLUSH,
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_CALLSTACK,
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_TEST_QUERY,
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_REGISTERING_EVENT,
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_ADDING_ACTION_TO_EVENT,
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_QUERY_AND_MODIFY_EVENT,
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_SHORT_CIRCUITING_EVENT,
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_READING_REGISTERS,
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_READING_MEMORY,
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_EDITING_MEMORY,
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_BP,
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_SHORT_CIRCUITING_STATE,
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_LIST_OR_MODIFY_BREAKPOINTS,
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_UPDATE_SYMBOL_INFO,
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RELOAD_SYMBOL_FINISHED,
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RELOAD_SEARCH_QUERY,
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_PTE,
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_VA2PA_AND_PA2VA,
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_BRINGING_PAGES_IN,

    //
    // hardware debuggee to debugger
    //

    //
    // hardware debugger to debuggee
    //

} DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION;

/**
 * @brief enum for different packet types in HyperDbg packets
 * @warning used in hwdbg
 *
 */
typedef enum _DEBUGGER_REMOTE_PACKET_TYPE
{

    //
    // Debugger to debuggee (vmx-root)
    //
    DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGER_TO_DEBUGGEE_EXECUTE_ON_VMX_ROOT = 1,

    //
    // Debugger to debuggee (user-mode)
    //
    DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGER_TO_DEBUGGEE_EXECUTE_ON_USER_MODE = 2,

    //
    // Debuggee to debugger (user-mode and kernel-mode, vmx-root mode)
    //
    DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGEE_TO_DEBUGGER = 3,

    //
    // Debugger to debuggee (hardware), used in hwdbg
    //
    DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGER_TO_DEBUGGEE_HARDWARE_LEVEL = 4,

    //
    // Debuggee to debugger (hardware), used in hwdbg
    //
    DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGEE_TO_DEBUGGER_HARDWARE_LEVEL = 5,

} DEBUGGER_REMOTE_PACKET_TYPE;

/**
 * @brief The structure of remote packets in HyperDbg
 *
 */
typedef struct _DEBUGGER_REMOTE_PACKET
{
    BYTE                                    Checksum;
    UINT64                                  Indicator; /* Shows the type of the packet */
    DEBUGGER_REMOTE_PACKET_TYPE             TypeOfThePacket;
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION RequestedActionOfThePacket;

} DEBUGGER_REMOTE_PACKET, *PDEBUGGER_REMOTE_PACKET;


//..\bin\debug\SDK\Headers\Constants.h
/**
 * @file Constants.h
 * @author Sina Karvandi (sina@hyperdbg.org)
 * @brief HyperDbg's SDK constants
 * @details This file contains definitions of constants
 * used in HyperDbg
 * @version 0.2
 * @date 2022-06-24
 *
 * @copyright This project is released under the GNU Public License v3.
 *
 */
#pragma once

//////////////////////////////////////////////////
//			 Version Information                //
//////////////////////////////////////////////////

#define VERSION_MAJOR 1
#define VERSION_MINOR 0
#define VERSION_PATCH 0

//
// Example of __DATE__ string: "Jul 27 2012"
//                              01234567890

#define BUILD_YEAR_CH0 (__DATE__[7])
#define BUILD_YEAR_CH1 (__DATE__[8])
#define BUILD_YEAR_CH2 (__DATE__[9])
#define BUILD_YEAR_CH3 (__DATE__[10])

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

#define BUILD_DAY_CH0 ((__DATE__[4] >= '0') ? (__DATE__[4]) : '0')
#define BUILD_DAY_CH1 (__DATE__[5])

//
// Example of __TIME__ string: "21:06:19"
//                              01234567

#define BUILD_HOUR_CH0 (__TIME__[0])
#define BUILD_HOUR_CH1 (__TIME__[1])

#define BUILD_MIN_CH0 (__TIME__[3])
#define BUILD_MIN_CH1 (__TIME__[4])

#define BUILD_SEC_CH0 (__TIME__[6])
#define BUILD_SEC_CH1 (__TIME__[7])

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

const unsigned char BuildDateTime[] = {
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

const unsigned char CompleteVersion[] = {
    'v',
    VERSION_MAJOR_INIT,
    '.',
    VERSION_MINOR_INIT,
    '.',
    VERSION_PATCH_INIT,
    '\0'};

const unsigned char BuildVersion[] = {
    BUILD_YEAR_CH0,
    BUILD_YEAR_CH1,
    BUILD_YEAR_CH2,
    BUILD_YEAR_CH3,
    BUILD_MONTH_CH0,
    BUILD_MONTH_CH1,
    BUILD_DAY_CH0,
    BUILD_DAY_CH1,
    '.',
    BUILD_HOUR_CH0,
    BUILD_HOUR_CH1,
    BUILD_MIN_CH0,
    BUILD_MIN_CH1,

    '\0'};

const unsigned char BuildSignature[] = {
    VERSION_MAJOR_INIT,
    '.',
    VERSION_MINOR_INIT,
    '.',
    VERSION_PATCH_INIT,
    '-',
    BUILD_YEAR_CH0,
    BUILD_YEAR_CH1,
    BUILD_YEAR_CH2,
    BUILD_YEAR_CH3,
    BUILD_MONTH_CH0,
    BUILD_MONTH_CH1,
    BUILD_DAY_CH0,
    BUILD_DAY_CH1,
    '.',
    BUILD_HOUR_CH0,
    BUILD_HOUR_CH1,
    BUILD_MIN_CH0,
    BUILD_MIN_CH1,

    '\0'};

#endif // SCRIPT_ENGINE_KERNEL_MODE

//////////////////////////////////////////////////
//				Message Tracing                 //
//////////////////////////////////////////////////

/**
 * @brief Default buffer count of packets for message tracing
 * @details number of packets storage for regular buffers
 */
#define MaximumPacketsCapacity 1000

/**
 * @brief Default buffer count of packets for message tracing
 * @details number of packets storage for priority buffers
 */
#define MaximumPacketsCapacityPriority 50

/**
 * @brief Size of normal OS (processor) pages
 */
#define NORMAL_PAGE_SIZE 4096 // PAGE_SIZE

/**
 * @brief Size of each packet
 */
#define PacketChunkSize NORMAL_PAGE_SIZE

/**
 * @brief size of user-mode buffer
 * @details Because of operation code at the start of the
 * buffer + 1 for null-termminating
 *
 */
#define UsermodeBufferSize sizeof(UINT32) + PacketChunkSize + 1

/**
 * @brief size of buffer for serial
 * @details the maximum packet size for sending over serial
 *
 */
#define MaxSerialPacketSize 10 * NORMAL_PAGE_SIZE

/**
 * @brief Final storage size of message tracing
 *
 */
#define LogBufferSize \
    MaximumPacketsCapacity *(PacketChunkSize + sizeof(BUFFER_HEADER))

/**
 * @brief Final storage size of message tracing
 *
 */
#define LogBufferSizePriority \
    MaximumPacketsCapacityPriority *(PacketChunkSize + sizeof(BUFFER_HEADER))

/**
 * @brief limitation of Windows DbgPrint message size
 * @details currently is not functional
 *
 */
#define DbgPrintLimitation 512

/**
 * @brief The seeds that user-mode codes use as the starter
 * of their events' tag
 *
 */
#define DebuggerEventTagStartSeed 0x1000000

/**
 * @brief The seeds that user-mode thread detail token start with it
 * @details This seed should not start with zero (0), otherwise it's
 * interpreted as error
 */
#define DebuggerThreadDebuggingTagStartSeed 0x1000000

/**
 * @brief The seeds that user-mode codes use as the starter
 * of their output source tag
 *
 */
#define DebuggerOutputSourceTagStartSeed 0x1

/**
 * @brief Determines how many sources a debugger can have for
 * a single event
 *
 */
#define DebuggerOutputSourceMaximumRemoteSourceForSingleEvent 0x5

/**
 * @brief The size of each chunk of memory used in the 'memcpy' function
 * of the script engine for transferring buffers in the VMX-root mode
 *
 */
#define DebuggerScriptEngineMemcpyMovingBufferSize 64

//////////////////////////////////////////////////
//                   EPT Hook                   //
//////////////////////////////////////////////////

/**
 * @brief Maximum number of initial pre-allocated EPT hooks
 *
 */
#define MAXIMUM_NUMBER_OF_INITIAL_PREALLOCATED_EPT_HOOKS 5

//////////////////////////////////////////////////
//             Instant Event Configs            //
//////////////////////////////////////////////////

/**
 * @brief Maximum number of (regular) instant events that are pre-allocated
 *
 */
#define MAXIMUM_REGULAR_INSTANT_EVENTS 20

/**
 * @brief Maximum number of (big) instant events that are pre-allocated
 *
 */
#define MAXIMUM_BIG_INSTANT_EVENTS 0

/**
 * @brief Pre-allocated size for a regular event + conditions buffer
 *
 */
#define REGULAR_INSTANT_EVENT_CONDITIONAL_BUFFER sizeof(DEBUGGER_EVENT) + 100

/**
 * @brief Pre-allocated size for a big event + conditions buffer
 *
 */
#define BIG_INSTANT_EVENT_CONDITIONAL_BUFFER sizeof(DEBUGGER_EVENT) + PAGE_SIZE

/**
 * @brief Pre-allocated size for a regular action + custom code or script buffer
 *
 */
#define REGULAR_INSTANT_EVENT_ACTION_BUFFER sizeof(DEBUGGER_EVENT_ACTION) + (PAGE_SIZE * 2)

/**
 * @brief Pre-allocated size for a big action + custom code or script buffer
 *
 */
#define BIG_INSTANT_EVENT_ACTION_BUFFER sizeof(DEBUGGER_EVENT_ACTION) + MaxSerialPacketSize

/**
 * @brief Pre-allocated size for a regular requested safe buffer
 *
 */
#define REGULAR_INSTANT_EVENT_REQUESTED_SAFE_BUFFER PAGE_SIZE

/**
 * @brief Pre-allocated size for a big requested safe buffer
 *
 */
#define BIG_INSTANT_EVENT_REQUESTED_SAFE_BUFFER MaxSerialPacketSize

//////////////////////////////////////////////////
//               Remote Connection              //
//////////////////////////////////////////////////

/**
 * @brief default port of HyperDbg for listening by
 * debuggee (server, guest)
 *
 */
#define DEFAULT_PORT "50000"

/**
 * @brief Packet size for TCP connections
 * @details Note that we might add something to the kernel buffers
 * that's why we add 0x100 to it
 */
#define COMMUNICATION_BUFFER_SIZE PacketChunkSize + 0x100

//////////////////////////////////////////////////
//             VMCALL Numbers                  //
//////////////////////////////////////////////////

/**
 * @brief The start number of VMCALL number allowed to be
 * used by top-level drivers
 *
 */
#define TOP_LEVEL_DRIVERS_VMCALL_STARTING_NUMBER 0x00000200

/**
 * @brief The start number of VMCALL number allowed to be
 * used by top-level drivers
 *
 */
#define TOP_LEVEL_DRIVERS_VMCALL_ENDING_NUMBER TOP_LEVEL_DRIVERS_VMCALL_STARTING_NUMBER + 0x100

//////////////////////////////////////////////////
//             Operation Codes                  //
//////////////////////////////////////////////////

/**
 * @brief If a operation use this bit in its Operation code,
 * then it means that the operation should be performed
 * mandatorily in debuggee and should not be sent to the debugger
 */
#define OPERATION_MANDATORY_DEBUGGEE_BIT (1 << 31)

/**
 * @brief Message logs id that comes from kernel-mode to
 * user-mode
 * @details Message area >= 0x5
 */
#define OPERATION_LOG_INFO_MESSAGE          1U
#define OPERATION_LOG_WARNING_MESSAGE       2U
#define OPERATION_LOG_ERROR_MESSAGE         3U
#define OPERATION_LOG_NON_IMMEDIATE_MESSAGE 4U
#define OPERATION_LOG_WITH_TAG              5U

#define OPERATION_COMMAND_FROM_DEBUGGER_CLOSE_AND_UNLOAD_VMM \
    6U | OPERATION_MANDATORY_DEBUGGEE_BIT
#define OPERATION_DEBUGGEE_USER_INPUT     7U | OPERATION_MANDATORY_DEBUGGEE_BIT
#define OPERATION_DEBUGGEE_REGISTER_EVENT 8U | OPERATION_MANDATORY_DEBUGGEE_BIT
#define OPERATION_DEBUGGEE_ADD_ACTION_TO_EVENT \
    9 | OPERATION_MANDATORY_DEBUGGEE_BIT
#define OPERATION_DEBUGGEE_CLEAR_EVENTS                            10U | OPERATION_MANDATORY_DEBUGGEE_BIT
#define OPERATION_DEBUGGEE_CLEAR_EVENTS_WITHOUT_NOTIFYING_DEBUGGER 11U | OPERATION_MANDATORY_DEBUGGEE_BIT
#define OPERATION_HYPERVISOR_DRIVER_IS_SUCCESSFULLY_LOADED \
    12U | OPERATION_MANDATORY_DEBUGGEE_BIT
#define OPERATION_HYPERVISOR_DRIVER_END_OF_IRPS \
    13U | OPERATION_MANDATORY_DEBUGGEE_BIT
#define OPERATION_COMMAND_FROM_DEBUGGER_RELOAD_SYMBOL \
    14U | OPERATION_MANDATORY_DEBUGGEE_BIT

#define OPERATION_NOTIFICATION_FROM_USER_DEBUGGER_PAUSE \
    15U | OPERATION_MANDATORY_DEBUGGEE_BIT

//////////////////////////////////////////////////
//       Breakpoints & Debug Breakpoints        //
//////////////////////////////////////////////////

/**
 * @brief maximum number of buffers to be allocated for a single
 * breakpoint
 */
#define MAXIMUM_BREAKPOINTS_WITHOUT_CONTINUE 100

/**
 * @brief maximum number of thread/process ids to be allocated for a simultaneous
 * debugging
 * @details it shows the maximum number of threads/processes that HyperDbg sets
 * trap flag for them
 *
 */
#define MAXIMUM_NUMBER_OF_THREAD_INFORMATION_FOR_TRAPS 200

//////////////////////////////////////////////////
//          Pool tags used in HyperDbg          //
//////////////////////////////////////////////////

/**
 * @brief Pool tag
 *
 */
#define POOLTAG 0x48444247 // [H]yper[DBG] (HDBG)

//////////////////////////////////////////////////
//            End of Buffer Detection           //
//////////////////////////////////////////////////

/**
 * @brief count of characters for serial end of buffer
 */
#define SERIAL_END_OF_BUFFER_CHARS_COUNT 0x4

/**
 * @brief characters of the buffer that we set at the end of
 * buffers for serial
 */
#define SERIAL_END_OF_BUFFER_CHAR_1 0x00
#define SERIAL_END_OF_BUFFER_CHAR_2 0x80
#define SERIAL_END_OF_BUFFER_CHAR_3 0xEE
#define SERIAL_END_OF_BUFFER_CHAR_4 0xFF

/**
 * @brief count of characters for tcp end of buffer
 */
#define TCP_END_OF_BUFFER_CHARS_COUNT 0x4

/**
 * @brief characters of the buffer that we set at the end of
 * buffers for tcp
 */
#define TCP_END_OF_BUFFER_CHAR_1 0x10
#define TCP_END_OF_BUFFER_CHAR_2 0x20
#define TCP_END_OF_BUFFER_CHAR_3 0x33
#define TCP_END_OF_BUFFER_CHAR_4 0x44

//////////////////////////////////////////////////
//                 Name of OS                    //
//////////////////////////////////////////////////

/**
 * @brief maximum name for OS name buffer
 *
 */
#define MAXIMUM_CHARACTER_FOR_OS_NAME 256

//////////////////////////////////////////////////
//              Processor Details               //
//////////////////////////////////////////////////

/**
 * @brief maximum instruction size in Intel
 */
#define MAXIMUM_INSTR_SIZE 16

/**
 * @brief maximum size for call instruction in Intel
 */
#define MAXIMUM_CALL_INSTR_SIZE 7

//////////////////////////////////////////////////
//              Symbols Details                 //
//////////////////////////////////////////////////

/**
 * @brief maximum supported modules to load
 * their symbol information
 */
#define MAXIMUM_SUPPORTED_SYMBOLS 1000

/**
 * @brief maximum size for GUID and Age of PE
 * @detail It seems that 33 bytes is enough but let's
 * have more space because there might be sth that we
 * missed :)
 */
#define MAXIMUM_GUID_AND_AGE_SIZE 60

//////////////////////////////////////////////////
//            Debuggee Communication            //
//////////////////////////////////////////////////

/**
 * @brief constant indicator of a HyperDbg packet
 * @warning used in hwdbg
 *
 */
#define INDICATOR_OF_HYPERDBG_PACKET \
    0x4859504552444247 // HYPERDBG = 0x4859504552444247

//////////////////////////////////////////////////
//               Command Details                //
//////////////////////////////////////////////////

/**
 * @brief maximum results that will be returned by !s* s*
 * command
 *
 */
#define MaximumSearchResults 0x1000

//////////////////////////////////////////////////
//                 Script Engine                //
//////////////////////////////////////////////////

/**
 * @brief EFLAGS/RFLAGS
 *
 */
#define X86_FLAGS_CF                 (1 << 0)
#define X86_FLAGS_PF                 (1 << 2)
#define X86_FLAGS_AF                 (1 << 4)
#define X86_FLAGS_ZF                 (1 << 6)
#define X86_FLAGS_SF                 (1 << 7)
#define X86_FLAGS_TF                 (1 << 8)
#define X86_FLAGS_IF                 (1 << 9)
#define X86_FLAGS_DF                 (1 << 10)
#define X86_FLAGS_OF                 (1 << 11)
#define X86_FLAGS_STATUS_MASK        (0xfff)
#define X86_FLAGS_IOPL_MASK          (3 << 12)
#define X86_FLAGS_IOPL_SHIFT         (12)
#define X86_FLAGS_IOPL_SHIFT_2ND_BIT (13)
#define X86_FLAGS_NT                 (1 << 14)
#define X86_FLAGS_RF                 (1 << 16)
#define X86_FLAGS_VM                 (1 << 17)
#define X86_FLAGS_AC                 (1 << 18)
#define X86_FLAGS_VIF                (1 << 19)
#define X86_FLAGS_VIP                (1 << 20)
#define X86_FLAGS_ID                 (1 << 21)
#define X86_FLAGS_RESERVED_ONES      0x2
#define X86_FLAGS_RESERVED           0xffc0802a

#define X86_FLAGS_RESERVED_BITS 0xffc38028
#define X86_FLAGS_FIXED         0x00000002

#ifndef LOWORD
#    define LOWORD(l) ((WORD)(l))
#endif // !LOWORD

#ifndef HIWORD
#    define HIWORD(l) ((WORD)(((DWORD)(l) >> 16) & 0xFFFF))
#endif // !HIWORD

#ifndef LOBYTE
#    define LOBYTE(w) ((BYTE)(w))
#endif // !LOBYTE

#ifndef HIBYTE
#    define HIBYTE(w) ((BYTE)(((WORD)(w) >> 8) & 0xFF))
#endif // !HIBYTE

#define MAX_TEMP_COUNT 128

#define MAX_STACK_BUFFER_COUNT 128

// TODO: Extract number of variables from input of ScriptEngine
// and allocate variableList Dynamically.
#define MAX_VAR_COUNT 512

#define MAX_FUNCTION_NAME_LENGTH 32

//////////////////////////////////////////////////
//                  Debugger                    //
//////////////////////////////////////////////////

/**
 * @brief Apply event modifications to all tags
 *
 */
#define DEBUGGER_MODIFY_EVENTS_APPLY_TO_ALL_TAG 0xffffffffffffffff

/**
 * @brief Maximum length for a function (to be used in showing distance
 * from symbol functions in the 'u' command)
 *
 */
#define DISASSEMBLY_MAXIMUM_DISTANCE_FROM_OBJECT_NAME 0xffff

/**
 * @brief Read and write MSRs to all cores
 *
 */
#define DEBUGGER_READ_AND_WRITE_ON_MSR_APPLY_ALL_CORES 0xffffffff

/**
 * @brief Apply the event to all the cores
 *
 */
#define DEBUGGER_DEBUGGEE_IS_RUNNING_NO_CORE 0xffffffff

/**
 * @brief Apply the event to all the cores
 *
 */
#define DEBUGGER_EVENT_APPLY_TO_ALL_CORES 0xffffffff

/**
 * @brief Apply the event to all the processes
 *
 */
#define DEBUGGER_EVENT_APPLY_TO_ALL_PROCESSES 0xffffffff

/**
 * @brief Apply to all Model Specific Registers
 *
 */
#define DEBUGGER_EVENT_MSR_READ_OR_WRITE_ALL_MSRS 0xffffffff

/**
 * @brief Apply to all first 32 exceptions
 *
 */
#define DEBUGGER_EVENT_EXCEPTIONS_ALL_FIRST_32_ENTRIES 0xffffffff

/**
 * @brief Apply to all syscalls and sysrets
 *
 */
#define DEBUGGER_EVENT_SYSCALL_ALL_SYSRET_OR_SYSCALLS 0xffffffff

/**
 * @brief Apply to all I/O ports
 *
 */
#define DEBUGGER_EVENT_ALL_IO_PORTS 0xffffffff

/**
 * @brief The constant to apply to all cores for bp command
 *
 */
#define DEBUGGEE_BP_APPLY_TO_ALL_CORES 0xffffffff

/**
 * @brief The constant to apply to all processes for bp command
 *
 */
#define DEBUGGEE_BP_APPLY_TO_ALL_PROCESSES 0xffffffff

/**
 * @brief The constant to apply to all threads for bp command
 *
 */
#define DEBUGGEE_BP_APPLY_TO_ALL_THREADS 0xffffffff

/**
 * @brief for reading all registers in r command.
 *
 */
#define DEBUGGEE_SHOW_ALL_REGISTERS 0xffffffff


//..\bin\debug\SDK\Headers\DataTypes.h
/**
 * @file DataTypes.h
 * @author Sina Karvandi (sina@hyperdbg.org)
 * @brief HyperDbg's SDK data type definitions
 * @details This file contains definitions of structures, enums, etc.
 * used in HyperDbg
 * @version 0.2
 * @date 2022-06-22
 *
 * @copyright This project is released under the GNU Public License v3.
 *
 */
#pragma once

//////////////////////////////////////////////////
//               Memory Stages                  //
//////////////////////////////////////////////////

/**
 * @brief Different levels of paging
 *
 */
typedef enum _PAGING_LEVEL
{
    PagingLevelPageTable = 0,
    PagingLevelPageDirectory,
    PagingLevelPageDirectoryPointerTable,
    PagingLevelPageMapLevel4
} PAGING_LEVEL;

//////////////////////////////////////////////////
//                 Pool Manager      			//
//////////////////////////////////////////////////

/**
 * @brief Inum of intentions for buffers (buffer tag)
 *
 */
typedef enum _POOL_ALLOCATION_INTENTION
{
    TRACKING_HOOKED_PAGES,
    EXEC_TRAMPOLINE,
    SPLIT_2MB_PAGING_TO_4KB_PAGE,
    DETOUR_HOOK_DETAILS,
    BREAKPOINT_DEFINITION_STRUCTURE,
    PROCESS_THREAD_HOLDER,

    //
    // Instant event buffers
    //
    INSTANT_REGULAR_EVENT_BUFFER,
    INSTANT_BIG_EVENT_BUFFER,
    INSTANT_REGULAR_EVENT_ACTION_BUFFER,
    INSTANT_BIG_EVENT_ACTION_BUFFER,

    //
    // Use for request safe buffers of the event
    //
    INSTANT_REGULAR_SAFE_BUFFER_FOR_EVENTS,
    INSTANT_BIG_SAFE_BUFFER_FOR_EVENTS,

} POOL_ALLOCATION_INTENTION;

//////////////////////////////////////////////////
//	   	Debug Registers Modifications 	    	//
//////////////////////////////////////////////////

typedef enum _DEBUG_REGISTER_TYPE
{
    BREAK_ON_INSTRUCTION_FETCH,
    BREAK_ON_WRITE_ONLY,
    BREAK_ON_IO_READ_OR_WRITE_NOT_SUPPORTED,
    BREAK_ON_READ_AND_WRITE_BUT_NOT_FETCH
} DEBUG_REGISTER_TYPE;

//////////////////////////////////////////////////
//              Execution Stages                //
//////////////////////////////////////////////////

typedef enum _VMX_EXECUTION_MODE
{
    VmxExecutionModeNonRoot = FALSE,
    VmxExecutionModeRoot    = TRUE
} VMX_EXECUTION_MODE;

/**
 * @brief Type of calling the event
 *
 */
typedef enum _VMM_CALLBACK_EVENT_CALLING_STAGE_TYPE
{
    VMM_CALLBACK_CALLING_STAGE_INVALID_EVENT_EMULATION = 0,
    VMM_CALLBACK_CALLING_STAGE_PRE_EVENT_EMULATION     = 1,
    VMM_CALLBACK_CALLING_STAGE_POST_EVENT_EMULATION    = 2,
    VMM_CALLBACK_CALLING_STAGE_ALL_EVENT_EMULATION     = 3

} VMM_CALLBACK_EVENT_CALLING_STAGE_TYPE;

/**
 * @brief enum to query different process and thread interception mechanisms
 *
 */
typedef enum _DEBUGGER_THREAD_PROCESS_TRACING
{

    DEBUGGER_THREAD_PROCESS_TRACING_INTERCEPT_CLOCK_INTERRUPTS_FOR_THREAD_CHANGE,
    DEBUGGER_THREAD_PROCESS_TRACING_INTERCEPT_CLOCK_INTERRUPTS_FOR_PROCESS_CHANGE,
    DEBUGGER_THREAD_PROCESS_TRACING_INTERCEPT_CLOCK_DEBUG_REGISTER_INTERCEPTION,
    DEBUGGER_THREAD_PROCESS_TRACING_INTERCEPT_CLOCK_WAITING_FOR_MOV_CR3_VM_EXITS,

} DEBUGGER_THREAD_PROCESS_TRACING;

//////////////////////////////////////////////////
//            Callback Definitions              //
//////////////////////////////////////////////////

/**
 * @brief Callback type that can be used to be used
 * as a custom ShowMessages function
 *
 */
typedef int (*Callback)(const char * Text);

//////////////////////////////////////////////////
//                Communications                //
//////////////////////////////////////////////////

/**
 * @brief The structure of user-input packet in HyperDbg
 *
 */
typedef struct _DEBUGGEE_USER_INPUT_PACKET
{
    UINT32  CommandLen;
    BOOLEAN IgnoreFinishedSignal;
    UINT32  Result;

    //
    // The user's input is here
    //

} DEBUGGEE_USER_INPUT_PACKET, *PDEBUGGEE_USER_INPUT_PACKET;

/**
 * @brief The structure of user-input packet in HyperDbg
 *
 */
typedef struct _DEBUGGEE_EVENT_AND_ACTION_HEADER_FOR_REMOTE_PACKET
{
    UINT32 Length;

    //
    // The buffer for event and action is here
    //

} DEBUGGEE_EVENT_AND_ACTION_HEADER_FOR_REMOTE_PACKET,
    *PDEBUGGEE_EVENT_AND_ACTION_HEADER_FOR_REMOTE_PACKET;

//////////////////////////////////////////////////
//                  Pausing                    //
//////////////////////////////////////////////////

#define SIZEOF_DEBUGGER_PAUSE_PACKET_RECEIVED \
    sizeof(DEBUGGER_PAUSE_PACKET_RECEIVED)

/**
 * @brief request to pause and halt the system
 *
 */
typedef struct _DEBUGGER_PAUSE_PACKET_RECEIVED
{
    UINT32 Result; // Result from kernel

} DEBUGGER_PAUSE_PACKET_RECEIVED, *PDEBUGGER_PAUSE_PACKET_RECEIVED;

/* ==============================================================================================
 */

/**
 * @brief The structure of detail of a triggered event in HyperDbg
 * @details This structure is also used for transferring breakpoint ids, RIP as the context, etc.
 *
 */
typedef struct _DEBUGGER_TRIGGERED_EVENT_DETAILS
{
    UINT64                                Tag; /* in breakpoints Tag is breakpoint id, not event tag */
    PVOID                                 Context;
    VMM_CALLBACK_EVENT_CALLING_STAGE_TYPE Stage;

} DEBUGGER_TRIGGERED_EVENT_DETAILS, *PDEBUGGER_TRIGGERED_EVENT_DETAILS;

/* ==============================================================================================
 */

/**
 * @brief The structure of pausing packet in kHyperDbg
 *
 */
typedef struct _DEBUGGEE_KD_PAUSED_PACKET
{
    UINT64                                Rip;
    BOOLEAN                               IsProcessorOn32BitMode; // if true shows that the address should be interpreted in 32-bit mode
    BOOLEAN                               IgnoreDisassembling;    // if check if diassembling should be ignored or not
    DEBUGGEE_PAUSING_REASON               PausingReason;
    ULONG                                 CurrentCore;
    UINT64                                EventTag;
    VMM_CALLBACK_EVENT_CALLING_STAGE_TYPE EventCallingStage;
    UINT64                                Rflags;
    BYTE                                  InstructionBytesOnRip[MAXIMUM_INSTR_SIZE];
    UINT16                                ReadInstructionLen;

} DEBUGGEE_KD_PAUSED_PACKET, *PDEBUGGEE_KD_PAUSED_PACKET;

/* ==============================================================================================
 */

/**
 * @brief The structure of pausing packet in uHyperDbg
 *
 */
typedef struct _DEBUGGEE_UD_PAUSED_PACKET
{
    UINT64                                Rip;
    UINT64                                ProcessDebuggingToken;
    BOOLEAN                               Is32Bit; // if true shows that the address should be interpreted in 32-bit mode
    DEBUGGEE_PAUSING_REASON               PausingReason;
    UINT32                                ProcessId;
    UINT32                                ThreadId;
    UINT64                                Rflags;
    UINT64                                EventTag;
    VMM_CALLBACK_EVENT_CALLING_STAGE_TYPE EventCallingStage;
    BYTE                                  InstructionBytesOnRip[MAXIMUM_INSTR_SIZE];
    UINT16                                ReadInstructionLen;
    GUEST_REGS                            GuestRegs;

} DEBUGGEE_UD_PAUSED_PACKET, *PDEBUGGEE_UD_PAUSED_PACKET;

//////////////////////////////////////////////////
//            Message Tracing Enums             //
//////////////////////////////////////////////////

/**
 * @brief Type of transferring buffer between user-to-kernel
 *
 */
typedef enum _NOTIFY_TYPE
{
    IRP_BASED,
    EVENT_BASED
} NOTIFY_TYPE;

//////////////////////////////////////////////////
//                  Structures                  //
//////////////////////////////////////////////////

/**
 * @brief The structure of message packet in HyperDbg
 *
 */
typedef struct _DEBUGGEE_MESSAGE_PACKET
{
    UINT32 OperationCode;
    CHAR   Message[PacketChunkSize];

} DEBUGGEE_MESSAGE_PACKET, *PDEBUGGEE_MESSAGE_PACKET;

/**
 * @brief Used to register event for transferring buffer between user-to-kernel
 *
 */
typedef struct _REGISTER_NOTIFY_BUFFER
{
    NOTIFY_TYPE Type;
    HANDLE      hEvent;

} REGISTER_NOTIFY_BUFFER, *PREGISTER_NOTIFY_BUFFER;

//////////////////////////////////////////////////
//                 Direct VMCALL                //
//////////////////////////////////////////////////

/**
 * @brief Used for sending direct VMCALLs on the VMX root-mode
 *
 */
typedef struct _DIRECT_VMCALL_PARAMETERS
{
    UINT64 OptionalParam1;
    UINT64 OptionalParam2;
    UINT64 OptionalParam3;

} DIRECT_VMCALL_PARAMETERS, *PDIRECT_VMCALL_PARAMETERS;

//////////////////////////////////////////////////
//                  EPT Hook                    //
//////////////////////////////////////////////////

/**
 * @brief different type of memory addresses
 *
 */
typedef enum _DEBUGGER_HOOK_MEMORY_TYPE
{
    DEBUGGER_MEMORY_HOOK_VIRTUAL_ADDRESS,
    DEBUGGER_MEMORY_HOOK_PHYSICAL_ADDRESS
} DEBUGGER_HOOK_MEMORY_TYPE;

/**
 * @brief Temporary $context used in some EPT hook commands
 *
 */
typedef struct _EPT_HOOKS_CONTEXT
{
    UINT64 HookingTag; // This is same as the event tag
    UINT64 PhysicalAddress;
    UINT64 VirtualAddress;
} EPT_HOOKS_CONTEXT, *PEPT_HOOKS_CONTEXT;

/**
 * @brief Setting details for EPT Hooks (!monitor)
 *
 */
typedef struct _EPT_HOOKS_ADDRESS_DETAILS_FOR_MEMORY_MONITOR
{
    UINT64                    StartAddress;
    UINT64                    EndAddress;
    BOOLEAN                   SetHookForRead;
    BOOLEAN                   SetHookForWrite;
    BOOLEAN                   SetHookForExec;
    DEBUGGER_HOOK_MEMORY_TYPE MemoryType;
    UINT64                    Tag;

} EPT_HOOKS_ADDRESS_DETAILS_FOR_MEMORY_MONITOR, *PEPT_HOOKS_ADDRESS_DETAILS_FOR_MEMORY_MONITOR;

/**
 * @brief Setting details for EPT Hooks (!epthook2)
 *
 */
typedef struct _EPT_HOOKS_ADDRESS_DETAILS_FOR_EPTHOOK2
{
    PVOID TargetAddress;
    PVOID HookFunction;

} EPT_HOOKS_ADDRESS_DETAILS_FOR_EPTHOOK2, *PEPT_HOOKS_ADDRESS_DETAILS_FOR_EPTHOOK2;

/**
 * @brief Details of unhooking single EPT hooks
 *
 */
typedef struct _EPT_SINGLE_HOOK_UNHOOKING_DETAILS
{
    BOOLEAN                     CallerNeedsToRestoreEntryAndInvalidateEpt;
    BOOLEAN                     RemoveBreakpointInterception;
    SIZE_T                      PhysicalAddress;
    UINT64 /* EPT_PML1_ENTRY */ OriginalEntry;

} EPT_SINGLE_HOOK_UNHOOKING_DETAILS, *PEPT_SINGLE_HOOK_UNHOOKING_DETAILS;

//////////////////////////////////////////////////
//                 Segment Types                //
//////////////////////////////////////////////////

/**
 * @brief Describe segment selector in VMX
 * @details This structure is copied from ia32.h to the SDK to
 * be used as a data type for functions
 *
 */
typedef union
{
    struct
    {
        /**
         * [Bits 3:0] Segment type.
         */
        UINT32 Type : 4;

        /**
         * [Bit 4] S - Descriptor type (0 = system; 1 = code or data).
         */
        UINT32 DescriptorType : 1;

        /**
         * [Bits 6:5] DPL - Descriptor privilege level.
         */
        UINT32 DescriptorPrivilegeLevel : 2;

        /**
         * [Bit 7] P - Segment present.
         */
        UINT32 Present : 1;

        UINT32 Reserved1 : 4;

        /**
         * [Bit 12] AVL - Available for use by system software.
         */
        UINT32 AvailableBit : 1;

        /**
         * [Bit 13] Reserved (except for CS). L - 64-bit mode active (for CS only).
         */
        UINT32 LongMode : 1;

        /**
         * [Bit 14] D/B - Default operation size (0 = 16-bit segment; 1 = 32-bit segment).
         */
        UINT32 DefaultBig : 1;

        /**
         * [Bit 15] G - Granularity.
         */
        UINT32 Granularity : 1;
        /**
         * [Bit 16] Segment unusable (0 = usable; 1 = unusable).
         */
        UINT32 Unusable : 1;
        UINT32 Reserved2 : 15;
    };

    UINT32 AsUInt;
} VMX_SEGMENT_ACCESS_RIGHTS_TYPE;

/**
 * @brief Segment selector
 *
 */
typedef struct _VMX_SEGMENT_SELECTOR
{
    UINT16                         Selector;
    VMX_SEGMENT_ACCESS_RIGHTS_TYPE Attributes;
    UINT32                         Limit;
    UINT64                         Base;
} VMX_SEGMENT_SELECTOR, *PVMX_SEGMENT_SELECTOR;


//..\bin\debug\SDK\Headers\ErrorCodes.h
/**
 * @file ErrorCodes.h
 * @author Sina Karvandi (sina@hyperdbg.org)
 * @brief HyperDbg's SDK Error codes
 * @details This file contains definitions of error codes used in HyperDbg
 * @version 0.2
 * @date 2022-06-24
 *
 * @copyright This project is released under the GNU Public License v3.
 *
 */
#pragma once

//////////////////////////////////////////////////
//		    	  Success Codes                 //
//////////////////////////////////////////////////

/**
 * @brief General value to indicate that the operation or
 * request was successful
 *
 */
#define DEBUGGER_OPERATION_WAS_SUCCESSFUL 0xFFFFFFFF

//////////////////////////////////////////////////
//		    	   Error Codes                  //
//////////////////////////////////////////////////

/**
 * @brief error, the tag not exist
 *
 */
#define DEBUGGER_ERROR_TAG_NOT_EXISTS 0xc0000000

/**
 * @brief error, invalid type of action
 *
 */
#define DEBUGGER_ERROR_INVALID_ACTION_TYPE 0xc0000001

/**
 * @brief error, the action buffer size is invalid
 *
 */
#define DEBUGGER_ERROR_ACTION_BUFFER_SIZE_IS_ZERO 0xc0000002

/**
 * @brief error, the event type is unknown
 *
 */
#define DEBUGGER_ERROR_EVENT_TYPE_IS_INVALID 0xc0000003

/**
 * @brief error, enable to create event
 *
 */
#define DEBUGGER_ERROR_UNABLE_TO_CREATE_EVENT 0xc0000004

/**
 * @brief error, invalid address specified for debugger
 *
 */
#define DEBUGGER_ERROR_INVALID_ADDRESS 0xc0000005

/**
 * @brief error, the core id is invalid
 *
 */
#define DEBUGGER_ERROR_INVALID_CORE_ID 0xc0000006

/**
 * @brief error, the index is greater than 32 in !exception command
 *
 */
#define DEBUGGER_ERROR_EXCEPTION_INDEX_EXCEED_FIRST_32_ENTRIES 0xc0000007

/**
 * @brief error, the index for !interrupt command is not between 32 to 256
 *
 */
#define DEBUGGER_ERROR_INTERRUPT_INDEX_IS_NOT_VALID 0xc0000008

/**
 * @brief error, unable to hide the debugger and enter to transparent-mode
 *
 */
#define DEBUGGER_ERROR_UNABLE_TO_HIDE_OR_UNHIDE_DEBUGGER 0xc0000009

/**
 * @brief error, the debugger is already in transparent-mode
 *
 */
#define DEBUGGER_ERROR_DEBUGGER_ALREADY_UHIDE 0xc000000a

/**
 * @brief error, invalid parameters in !e* e* commands
 *
 */
#define DEBUGGER_ERROR_EDIT_MEMORY_STATUS_INVALID_PARAMETER 0xc000000b

/**
 * @brief error, an invalid address is specified based on current cr3
 * in !e* or e* commands
 *
 */
#define DEBUGGER_ERROR_EDIT_MEMORY_STATUS_INVALID_ADDRESS_BASED_ON_CURRENT_PROCESS \
    0xc000000c

/**
 * @brief error, an invalid address is specified based on anotehr process's cr3
 * in !e* or e* commands
 *
 */
#define DEBUGGER_ERROR_EDIT_MEMORY_STATUS_INVALID_ADDRESS_BASED_ON_OTHER_PROCESS \
    0xc000000d

/**
 * @brief error, invalid tag for 'events' command (tag id is unknown for kernel)
 *
 */
#define DEBUGGER_ERROR_MODIFY_EVENTS_INVALID_TAG 0xc000000e

/**
 * @brief error, type of action (enable/disable/clear) is wrong
 *
 */
#define DEBUGGER_ERROR_MODIFY_EVENTS_INVALID_TYPE_OF_ACTION 0xc000000f

/**
 * @brief error, invalid parameters steppings actions
 *
 */
#define DEBUGGER_ERROR_STEPPING_INVALID_PARAMETER 0xc0000010

/**
 * @brief error, thread is invalid (not found) or disabled in
 * stepping (step-in & step-out) requests
 *
 */
#define DEBUGGER_ERROR_STEPPINGS_EITHER_THREAD_NOT_FOUND_OR_DISABLED 0xc0000011

/**
 * @brief error, baud rate is invalid
 *
 */
#define DEBUGGER_ERROR_PREPARING_DEBUGGEE_INVALID_BAUDRATE 0xc0000012

/**
 * @brief error, serial port address is invalid
 *
 */
#define DEBUGGER_ERROR_PREPARING_DEBUGGEE_INVALID_SERIAL_PORT 0xc0000013

/**
 * @brief error, invalid core selected in changing core in remote debuggee
 *
 */
#define DEBUGGER_ERROR_PREPARING_DEBUGGEE_INVALID_CORE_IN_REMOTE_DEBUGGE \
    0xc0000014

/**
 * @brief error, invalid process selected in changing process in remote debuggee
 *
 */
#define DEBUGGER_ERROR_PREPARING_DEBUGGEE_UNABLE_TO_SWITCH_TO_NEW_PROCESS \
    0xc0000015

/**
 * @brief error, unable to run script in remote debuggee
 *
 */
#define DEBUGGER_ERROR_PREPARING_DEBUGGEE_TO_RUN_SCRIPT 0xc0000016

/**
 * @brief error, invalid register number
 *
 */
#define DEBUGGER_ERROR_INVALID_REGISTER_NUMBER 0xc0000017

/**
 * @brief error, maximum pools were used without continuing debuggee
 *
 */
#define DEBUGGER_ERROR_MAXIMUM_BREAKPOINT_WITHOUT_CONTINUE 0xc0000018

/**
 * @brief error, breakpoint already exists on the target address
 *
 */
#define DEBUGGER_ERROR_BREAKPOINT_ALREADY_EXISTS_ON_THE_ADDRESS 0xc0000019

/**
 * @brief error, breakpoint id not found
 *
 */
#define DEBUGGER_ERROR_BREAKPOINT_ID_NOT_FOUND 0xc000001a

/**
 * @brief error, breakpoint already disabled
 *
 */
#define DEBUGGER_ERROR_BREAKPOINT_ALREADY_DISABLED 0xc000001b

/**
 * @brief error, breakpoint already enabled
 *
 */
#define DEBUGGER_ERROR_BREAKPOINT_ALREADY_ENABLED 0xc000001c

/**
 * @brief error, memory type is invalid
 *
 */
#define DEBUGGER_ERROR_MEMORY_TYPE_INVALID 0xc000001d

/**
 * @brief error, the process id is invalid
 *
 */
#define DEBUGGER_ERROR_INVALID_PROCESS_ID 0xc000001e

/**
 * @brief error, for event specific reasons the event is not
 * applied
 *
 */
#define DEBUGGER_ERROR_EVENT_IS_NOT_APPLIED 0xc000001f

/**
 * @brief error, for process switch or process details, invalid parameter
 *
 */
#define DEBUGGER_ERROR_DETAILS_OR_SWITCH_PROCESS_INVALID_PARAMETER 0xc0000020

/**
 * @brief error, for thread switch or thread details, invalid parameter
 *
 */
#define DEBUGGER_ERROR_DETAILS_OR_SWITCH_THREAD_INVALID_PARAMETER 0xc0000021

/**
 * @brief error, maximum breakpoint for a single page is hit
 *
 */
#define DEBUGGER_ERROR_MAXIMUM_BREAKPOINT_FOR_A_SINGLE_PAGE_IS_HIT 0xc0000022

/**
 * @brief error, there is no pre-allocated buffer
 *
 */
#define DEBUGGER_ERROR_PRE_ALLOCATED_BUFFER_IS_EMPTY 0xc0000023

/**
 * @brief error, in the EPT handler, it could not split the 2MB pages to
 * 512 entries of 4 KB pages
 *
 */
#define DEBUGGER_ERROR_EPT_COULD_NOT_SPLIT_THE_LARGE_PAGE_TO_4KB_PAGES 0xc0000024

/**
 * @brief error, failed to get PML1 entry of the target address
 *
 */
#define DEBUGGER_ERROR_EPT_FAILED_TO_GET_PML1_ENTRY_OF_TARGET_ADDRESS 0xc0000025

/**
 * @brief error, multiple EPT Hooks or Monitors are applied on a single page
 *
 */
#define DEBUGGER_ERROR_EPT_MULTIPLE_HOOKS_IN_A_SINGLE_PAGE 0xc0000026

/**
 * @brief error, could not build the EPT Hook
 *
 */
#define DEBUGGER_ERROR_COULD_NOT_BUILD_THE_EPT_HOOK 0xc0000027

/**
 * @brief error, could not find the type of allocation
 *
 */
#define DEBUGGER_ERROR_COULD_NOT_FIND_ALLOCATION_TYPE 0xc0000028

/**
 * @brief error, could not find the index of test query
 *
 */
#define DEBUGGER_ERROR_INVALID_TEST_QUERY_INDEX 0xc0000029

/**
 * @brief error, failed to attach to the target user-mode process
 *
 */
#define DEBUGGER_ERROR_UNABLE_TO_ATTACH_TO_TARGET_USER_MODE_PROCESS 0xc000002a

/**
 * @brief error, failed to remove hooks as entrypoint is not reached yet
 * @details The caller of this functionality should keep sending the previous
 * IOCTL until the hook is remove successfully
 *
 */
#define DEBUGGER_ERROR_UNABLE_TO_REMOVE_HOOKS_ENTRYPOINT_NOT_REACHED 0xc000002b

/**
 * @brief error, could not remove the previous hook
 *
 */
#define DEBUGGER_ERROR_UNABLE_TO_REMOVE_HOOKS 0xc000002c

/**
 * @brief error, the needed routines for debugging is not initialized
 *
 */
#define DEBUGGER_ERROR_FUNCTIONS_FOR_INITIALIZING_PEB_ADDRESSES_ARE_NOT_INITIALIZED 0xc000002d

/**
 * @brief error, unable to get 32-bit or 64-bit of the target process
 *
 */
#define DEBUGGER_ERROR_UNABLE_TO_DETECT_32_BIT_OR_64_BIT_PROCESS 0xc000002e

/**
 * @brief error, unable to kill the target process
 *
 */
#define DEBUGGER_ERROR_UNABLE_TO_KILL_THE_PROCESS 0xc000002f

/**
 * @brief error, invalid thread debugging token
 *
 */
#define DEBUGGER_ERROR_INVALID_THREAD_DEBUGGING_TOKEN 0xc0000030

/**
 * @brief error, unable to pause the process's threads
 *
 */
#define DEBUGGER_ERROR_UNABLE_TO_PAUSE_THE_PROCESS_THREADS 0xc0000031

/**
 * @brief error, user debugger already attached to this process
 *
 */
#define DEBUGGER_ERROR_UNABLE_TO_ATTACH_TO_AN_ALREADY_ATTACHED_PROCESS 0xc0000032

/**
 * @brief error, the user debugger is not attached to the target process
 *
 */
#define DEBUGGER_ERROR_THE_USER_DEBUGGER_NOT_ATTACHED_TO_THE_PROCESS 0xc0000033

/**
 * @brief error, cannot detach from the process as there are paused threads
 *
 */
#define DEBUGGER_ERROR_UNABLE_TO_DETACH_AS_THERE_ARE_PAUSED_THREADS 0xc0000034

/**
 * @brief error, cannot switch to new thread as the process id or thread id is not found
 *
 */
#define DEBUGGER_ERROR_UNABLE_TO_SWITCH_PROCESS_ID_OR_THREAD_ID_IS_INVALID 0xc0000035

/**
 * @brief error, cannot switch to new thread the process doesn't contain an active thread
 *
 */
#define DEBUGGER_ERROR_UNABLE_TO_SWITCH_THERE_IS_NO_THREAD_ON_THE_PROCESS 0xc0000036

/**
 * @brief error, unable to get modules
 *
 */
#define DEBUGGER_ERROR_UNABLE_TO_GET_MODULES_OF_THE_PROCESS 0xc0000037

/**
 * @brief error, unable to get the callstack
 *
 */
#define DEBUGGER_ERROR_UNABLE_TO_GET_CALLSTACK 0xc0000038

/**
 * @brief error, unable to query count of processes or threads
 *
 */
#define DEBUGGER_ERROR_UNABLE_TO_QUERY_COUNT_OF_PROCESSES_OR_THREADS 0xc0000039

/**
 * @brief error, using short-circuiting event with post-event mode is
 * not supported in HyperDbg
 *
 */
#define DEBUGGER_ERROR_USING_SHORT_CIRCUITING_EVENT_WITH_POST_EVENT_MODE_IS_FORBIDDEDN 0xc000003a

/**
 * @brief error, unknown test query is received
 *
 */
#define DEBUGGER_ERROR_UNKNOWN_TEST_QUERY_RECEIVED 0xc000003b

/**
 * @brief error, for reading from memory in case of invalid parameters
 *
 */
#define DEBUGGER_ERROR_READING_MEMORY_INVALID_PARAMETER 0xc000003c

/**
 * @brief error, the list of threads/process trap flag is full
 *
 */
#define DEBUGGER_ERROR_THE_TRAP_FLAG_LIST_IS_FULL 0xc000003d

/**
 * @brief error, unable to kill the target process. process does not exists
 *
 */
#define DEBUGGER_ERROR_UNABLE_TO_KILL_THE_PROCESS_DOES_NOT_EXISTS 0xc000003e

/**
 * @brief error, the execution mode is incorrect
 *
 */
#define DEBUGGER_ERROR_MODE_EXECUTION_IS_INVALID 0xc000003f

/**
 * @brief error, the process id cannot be specified while the debugger is in VMX-root mode
 *
 */
#define DEBUGGER_ERROR_PROCESS_ID_CANNOT_BE_SPECIFIED_WHILE_APPLYING_EVENT_FROM_VMX_ROOT_MODE 0xc0000040

/**
 * @brief error, the preallocated buffer is not enough for storing event+conditional buffer
 *
 */
#define DEBUGGER_ERROR_INSTANT_EVENT_PREALLOCATED_BUFFER_IS_NOT_ENOUGH_FOR_EVENT_AND_CONDITIONALS 0xc0000041

/**
 * @brief error, the regular preallocated buffer not found
 *
 */
#define DEBUGGER_ERROR_INSTANT_EVENT_REGULAR_PREALLOCATED_BUFFER_NOT_FOUND 0xc0000042

/**
 * @brief error, the big preallocated buffer not found
 *
 */
#define DEBUGGER_ERROR_INSTANT_EVENT_BIG_PREALLOCATED_BUFFER_NOT_FOUND 0xc0000043

/**
 * @brief error, enable to create action (cannot allocate buffer)
 *
 */
#define DEBUGGER_ERROR_UNABLE_TO_CREATE_ACTION_CANNOT_ALLOCATE_BUFFER 0xc0000044

/**
 * @brief error, the regular preallocated buffer not found (for action)
 *
 */
#define DEBUGGER_ERROR_INSTANT_EVENT_ACTION_REGULAR_PREALLOCATED_BUFFER_NOT_FOUND 0xc0000045

/**
 * @brief error, the big preallocated buffer not found (for action)
 *
 */
#define DEBUGGER_ERROR_INSTANT_EVENT_ACTION_BIG_PREALLOCATED_BUFFER_NOT_FOUND 0xc0000046

/**
 * @brief error, the preallocated buffer is not enough for storing action buffer
 *
 */
#define DEBUGGER_ERROR_INSTANT_EVENT_PREALLOCATED_BUFFER_IS_NOT_ENOUGH_FOR_ACTION_BUFFER 0xc0000047

/**
 * @brief error, the requested optional buffer is bigger than send/receive stack of the debugger
 *
 */
#define DEBUGGER_ERROR_INSTANT_EVENT_REQUESTED_OPTIONAL_BUFFER_IS_BIGGER_THAN_DEBUGGERS_SEND_RECEIVE_STACK 0xc0000048

/**
 * @brief error, the requested safe buffer does not exist (regular)
 *
 */
#define DEBUGGER_ERROR_INSTANT_EVENT_REGULAR_REQUESTED_SAFE_BUFFER_NOT_FOUND 0xc0000049

/**
 * @brief error, the requested safe buffer does not exists (big)
 *
 */
#define DEBUGGER_ERROR_INSTANT_EVENT_BIG_REQUESTED_SAFE_BUFFER_NOT_FOUND 0xc000004a

/**
 * @brief error, the preallocated buffer is not enough for storing safe requested buffer
 *
 */
#define DEBUGGER_ERROR_INSTANT_EVENT_PREALLOCATED_BUFFER_IS_NOT_ENOUGH_FOR_REQUESTED_SAFE_BUFFER 0xc000004b

/**
 * @brief error, enable to create requested safe buffer (cannot allocate buffer)
 *
 */
#define DEBUGGER_ERROR_UNABLE_TO_ALLOCATE_REQUESTED_SAFE_BUFFER 0xc000004c

/**
 * @brief error, could not find the type of preactivation
 *
 */
#define DEBUGGER_ERROR_COULD_NOT_FIND_PREACTIVATION_TYPE 0xc000004d

/**
 * @brief error, the mode exec trap is not already initialized
 *
 */
#define DEBUGGER_ERROR_THE_MODE_EXEC_TRAP_IS_NOT_INITIALIZED 0xc000004e

/**
 * @brief error, the target event(s) is/are disabled but cannot clear them because the buffer of the user-mode
 * priority is full
 *
 */
#define DEBUGGER_ERROR_THE_TARGET_EVENT_IS_DISABLED_BUT_CANNOT_BE_CLEARED_PRIRITY_BUFFER_IS_FULL 0xc000004f

/**
 * @brief error, not all cores are locked (probably due to a race condition in HyperDbg) in
 * instant-event mechanism
 *
 */
#define DEBUGGER_ERROR_NOT_ALL_CORES_ARE_LOCKED_FOR_APPLYING_INSTANT_EVENT 0xc0000050

/**
 * @brief error, switching to the target core is not possible because core is not locked
 * (probably due to a race condition in HyperDbg)
 *
 */
#define DEBUGGER_ERROR_TARGET_SWITCHING_CORE_IS_NOT_LOCKED 0xc0000051

/**
 * @brief error, invalid physical address
 *
 */
#define DEBUGGER_ERROR_INVALID_PHYSICAL_ADDRESS 0xc0000052

//
// WHEN YOU ADD ANYTHING TO THIS LIST OF ERRORS, THEN
// MAKE SURE TO ADD AN ERROR MESSAGE TO ShowErrorMessage(UINT32 Error)
// FUNCTION
//


//..\bin\debug\SDK\Headers\Events.h
/**
 * @file Events.h
 * @author Sina Karvandi (sina@hyperdbg.org)
 * @brief HyperDbg's SDK Headers for Events
 * @details This file contains definitions of event datatypes
 * @version 0.2
 * @date 2022-06-28
 *
 * @copyright This project is released under the GNU Public License v3.
 *
 */
#pragma once

//////////////////////////////////////////////////
//               System Events                  //
//////////////////////////////////////////////////

/**
 * @brief Exceptions enum
 *
 */
typedef enum _EXCEPTION_VECTORS
{
    EXCEPTION_VECTOR_DIVIDE_ERROR,
    EXCEPTION_VECTOR_DEBUG_BREAKPOINT,
    EXCEPTION_VECTOR_NMI,
    EXCEPTION_VECTOR_BREAKPOINT,
    EXCEPTION_VECTOR_OVERFLOW,
    EXCEPTION_VECTOR_BOUND_RANGE_EXCEEDED,
    EXCEPTION_VECTOR_UNDEFINED_OPCODE,
    EXCEPTION_VECTOR_NO_MATH_COPROCESSOR,
    EXCEPTION_VECTOR_DOUBLE_FAULT,
    EXCEPTION_VECTOR_RESERVED0,
    EXCEPTION_VECTOR_INVALID_TASK_SEGMENT_SELECTOR,
    EXCEPTION_VECTOR_SEGMENT_NOT_PRESENT,
    EXCEPTION_VECTOR_STACK_SEGMENT_FAULT,
    EXCEPTION_VECTOR_GENERAL_PROTECTION_FAULT,
    EXCEPTION_VECTOR_PAGE_FAULT,
    EXCEPTION_VECTOR_RESERVED1,
    EXCEPTION_VECTOR_MATH_FAULT,
    EXCEPTION_VECTOR_ALIGNMENT_CHECK,
    EXCEPTION_VECTOR_MACHINE_CHECK,
    EXCEPTION_VECTOR_SIMD_FLOATING_POINT_NUMERIC_ERROR,
    EXCEPTION_VECTOR_VIRTUAL_EXCEPTION,
    EXCEPTION_VECTOR_RESERVED2,
    EXCEPTION_VECTOR_RESERVED3,
    EXCEPTION_VECTOR_RESERVED4,
    EXCEPTION_VECTOR_RESERVED5,
    EXCEPTION_VECTOR_RESERVED6,
    EXCEPTION_VECTOR_RESERVED7,
    EXCEPTION_VECTOR_RESERVED8,
    EXCEPTION_VECTOR_RESERVED9,
    EXCEPTION_VECTOR_RESERVED10,
    EXCEPTION_VECTOR_RESERVED11,
    EXCEPTION_VECTOR_RESERVED12,

    //
    // NT (Windows) specific exception vectors.
    //
    APC_INTERRUPT   = 31,
    DPC_INTERRUPT   = 47,
    CLOCK_INTERRUPT = 209,
    IPI_INTERRUPT   = 225,
    PMI_INTERRUPT   = 254,

} EXCEPTION_VECTORS;

//////////////////////////////////////////////////
//			     Callback Enums                 //
//////////////////////////////////////////////////

/**
 * @brief The status of triggering events
 *
 */
typedef enum _VMM_CALLBACK_TRIGGERING_EVENT_STATUS_TYPE
{
    VMM_CALLBACK_TRIGGERING_EVENT_STATUS_SUCCESSFUL_NO_INITIALIZED = 0,
    VMM_CALLBACK_TRIGGERING_EVENT_STATUS_SUCCESSFUL                = 0,
    VMM_CALLBACK_TRIGGERING_EVENT_STATUS_SUCCESSFUL_IGNORE_EVENT   = 1,
    VMM_CALLBACK_TRIGGERING_EVENT_STATUS_DEBUGGER_NOT_ENABLED      = 2,
    VMM_CALLBACK_TRIGGERING_EVENT_STATUS_INVALID_EVENT_TYPE        = 3,

} VMM_CALLBACK_TRIGGERING_EVENT_STATUS_TYPE;

//////////////////////////////////////////////////
//               Event Details                  //
//////////////////////////////////////////////////

/**
 * @brief enum to show type of all HyperDbg events
 *
 */
typedef enum _VMM_EVENT_TYPE_ENUM
{

    //
    // EPT Memory Monitoring Events
    //
    HIDDEN_HOOK_READ_AND_WRITE_AND_EXECUTE,
    HIDDEN_HOOK_READ_AND_WRITE,
    HIDDEN_HOOK_READ_AND_EXECUTE,
    HIDDEN_HOOK_WRITE_AND_EXECUTE,
    HIDDEN_HOOK_READ,
    HIDDEN_HOOK_WRITE,
    HIDDEN_HOOK_EXECUTE,

    //
    // EPT Hook Events
    //
    HIDDEN_HOOK_EXEC_DETOURS,
    HIDDEN_HOOK_EXEC_CC,

    //
    // System-call Events
    //
    SYSCALL_HOOK_EFER_SYSCALL,
    SYSCALL_HOOK_EFER_SYSRET,

    //
    // CPUID Instruction Execution Events
    //
    CPUID_INSTRUCTION_EXECUTION,

    //
    // Model-Specific Registers (MSRs) Reads/Modifications Events
    //
    RDMSR_INSTRUCTION_EXECUTION,
    WRMSR_INSTRUCTION_EXECUTION,

    //
    // PMIO Events
    //
    IN_INSTRUCTION_EXECUTION,
    OUT_INSTRUCTION_EXECUTION,

    //
    // Interrupts/Exceptions/Faults Events
    //
    EXCEPTION_OCCURRED,
    EXTERNAL_INTERRUPT_OCCURRED,

    //
    // Debug Registers Events
    //
    DEBUG_REGISTERS_ACCESSED,

    //
    // Timing & Performance Events
    //
    TSC_INSTRUCTION_EXECUTION,
    PMC_INSTRUCTION_EXECUTION,

    //
    // VMCALL Instruction Execution Events
    //
    VMCALL_INSTRUCTION_EXECUTION,

    //
    // Control Registers Events
    //
    CONTROL_REGISTER_MODIFIED,
    CONTROL_REGISTER_READ,
    CONTROL_REGISTER_3_MODIFIED,

    //
    // Execution Trap Events
    //
    TRAP_EXECUTION_MODE_CHANGED,
    TRAP_EXECUTION_INSTRUCTION_TRACE,

} VMM_EVENT_TYPE_ENUM;

/**
 * @brief Type of Actions
 *
 */
typedef enum _DEBUGGER_EVENT_ACTION_TYPE_ENUM
{
    BREAK_TO_DEBUGGER,
    RUN_SCRIPT,
    RUN_CUSTOM_CODE

} DEBUGGER_EVENT_ACTION_TYPE_ENUM;

/**
 * @brief Type of handling !syscall or !sysret
 *
 */
typedef enum _DEBUGGER_EVENT_SYSCALL_SYSRET_TYPE
{
    DEBUGGER_EVENT_SYSCALL_SYSRET_SAFE_ACCESS_MEMORY = 0,
    DEBUGGER_EVENT_SYSCALL_SYSRET_HANDLE_ALL_UD      = 1,

} DEBUGGER_EVENT_SYSCALL_SYSRET_TYPE;

#define SIZEOF_DEBUGGER_MODIFY_EVENTS sizeof(DEBUGGER_MODIFY_EVENTS)

/**
 * @brief Type of mode change traps
 *
 */
typedef enum _DEBUGGER_EVENT_MODE_TYPE
{
    DEBUGGER_EVENT_MODE_TYPE_USER_MODE_AND_KERNEL_MODE = 1,
    DEBUGGER_EVENT_MODE_TYPE_USER_MODE                 = 3,
    DEBUGGER_EVENT_MODE_TYPE_KERNEL_MODE               = 0,
    DEBUGGER_EVENT_MODE_TYPE_INVALID                   = 0xffffffff,

} DEBUGGER_EVENT_MODE_TYPE;

/**
 * @brief Type of tracing events
 *
 */
typedef enum _DEBUGGER_EVENT_TRACE_TYPE
{
    DEBUGGER_EVENT_TRACE_TYPE_INVALID                 = 0,
    DEBUGGER_EVENT_TRACE_TYPE_STEP_IN                 = 1,
    DEBUGGER_EVENT_TRACE_TYPE_STEP_OUT                = 2,
    DEBUGGER_EVENT_TRACE_TYPE_INSTRUMENTATION_STEP_IN = 3,

} DEBUGGER_EVENT_TRACE_TYPE;

/**
 * @brief different types of modifying events request (enable/disable/clear)
 *
 */
typedef enum _DEBUGGER_MODIFY_EVENTS_TYPE
{
    DEBUGGER_MODIFY_EVENTS_QUERY_STATE,
    DEBUGGER_MODIFY_EVENTS_ENABLE,
    DEBUGGER_MODIFY_EVENTS_DISABLE,
    DEBUGGER_MODIFY_EVENTS_CLEAR,
} DEBUGGER_MODIFY_EVENTS_TYPE;

/**
 * @brief request for modifying events (enable/disable/clear)
 *
 */
typedef struct _DEBUGGER_MODIFY_EVENTS
{
    UINT64 Tag;          // Tag of the target event that we want to modify
    UINT64 KernelStatus; // Kernel put the status in this field
    DEBUGGER_MODIFY_EVENTS_TYPE
    TypeOfAction;      // Determines what's the action (enable | disable | clear)
    BOOLEAN IsEnabled; // Determines what's the action (enable | disable | clear)

} DEBUGGER_MODIFY_EVENTS, *PDEBUGGER_MODIFY_EVENTS;

/**
 * @brief request for performing a short-circuiting event
 *
 */
typedef struct _DEBUGGER_SHORT_CIRCUITING_EVENT
{
    UINT64  KernelStatus;      // Kernel put the status in this field
    BOOLEAN IsShortCircuiting; // Determines whether to perform short circuting (on | off)

} DEBUGGER_SHORT_CIRCUITING_EVENT, *PDEBUGGER_SHORT_CIRCUITING_EVENT;

//////////////////////////////////////////////////
//                Event Options                 //
//////////////////////////////////////////////////

/**
 * @brief request for performing a short-circuiting event
 *
 */
typedef struct _DEBUGGER_EVENT_OPTIONS
{
    UINT64 OptionalParam1; // Optional parameter
    UINT64 OptionalParam2; // Optional parameter
    UINT64 OptionalParam3; // Optional parameter
    UINT64 OptionalParam4; // Optional parameter
    UINT64 OptionalParam5; // Optional parameter
    UINT64 OptionalParam6; // Optional parameter

} DEBUGGER_EVENT_OPTIONS, *PDEBUGGER_EVENT_OPTIONS;

//////////////////////////////////////////////////
//    Enums For Event And Debugger Resources    //
//////////////////////////////////////////////////

/**
 * @brief Things to consider when applying resources
 *
 */
typedef enum _PROTECTED_HV_RESOURCES_PASSING_OVERS
{
    //
    // for exception bitmap
    //
    PASSING_OVER_NONE                                  = 0,
    PASSING_OVER_UD_EXCEPTIONS_FOR_SYSCALL_SYSRET_HOOK = 1,
    PASSING_OVER_EXCEPTION_EVENTS,

    //
    // for external interupts-exitings
    //
    PASSING_OVER_INTERRUPT_EVENTS,

    //
    // for external rdtsc/p exitings
    //
    PASSING_OVER_TSC_EVENTS,

    //
    // for external mov to hardware debug registers exitings
    //
    PASSING_OVER_MOV_TO_HW_DEBUG_REGS_EVENTS,

    //
    // for external mov to control registers exitings
    //
    PASSING_OVER_MOV_TO_CONTROL_REGS_EVENTS,

} PROTECTED_HV_RESOURCES_PASSING_OVERS;

/**
 * @brief Type of protected (multi-used) resources
 *
 */
typedef enum _PROTECTED_HV_RESOURCES_TYPE
{
    PROTECTED_HV_RESOURCES_EXCEPTION_BITMAP,

    PROTECTED_HV_RESOURCES_EXTERNAL_INTERRUPT_EXITING,

    PROTECTED_HV_RESOURCES_RDTSC_RDTSCP_EXITING,

    PROTECTED_HV_RESOURCES_MOV_TO_DEBUG_REGISTER_EXITING,

    PROTECTED_HV_RESOURCES_MOV_CONTROL_REGISTER_EXITING,

    PROTECTED_HV_RESOURCES_MOV_TO_CR3_EXITING,

} PROTECTED_HV_RESOURCES_TYPE;

//////////////////////////////////////////////////
//               Event Details                  //
//////////////////////////////////////////////////

/**
 * @brief Each command is like the following struct, it also used for
 * tracing works in user mode and sending it to the kernl mode
 * @details THIS IS NOT WHAT HYPERDBG SAVES FOR EVENTS IN KERNEL-MODE
 */
typedef struct _DEBUGGER_GENERAL_EVENT_DETAIL
{
    LIST_ENTRY
    CommandsEventList; // Linked-list of commands list (used for tracing purpose
                       // in user mode)

    time_t CreationTime; // Date of creating this event

    UINT32 CoreId; // determines the core index to apply this event to, if it's
                   // 0xffffffff means that we have to apply it to all cores

    UINT32 ProcessId; // determines the process id to apply this to
                      // only that 0xffffffff means that we have to
                      // apply it to all processes

    BOOLEAN IsEnabled;

    BOOLEAN EnableShortCircuiting; // indicates whether the short-circuiting event
                                   // is enabled or not for this event

    VMM_CALLBACK_EVENT_CALLING_STAGE_TYPE EventStage; // reveals the calling stage of the event
    // (whether it's a all- pre- or post- event)

    BOOLEAN HasCustomOutput; // Shows whether this event has a custom output
                             // source or not

    UINT64
    OutputSourceTags
        [DebuggerOutputSourceMaximumRemoteSourceForSingleEvent]; // tags of
                                                                 // multiple
                                                                 // sources which
                                                                 // can be used to
                                                                 // send the event
                                                                 // results of
                                                                 // scripts to
                                                                 // remote sources

    UINT32 CountOfActions;

    UINT64              Tag; // is same as operation code
    VMM_EVENT_TYPE_ENUM EventType;

    DEBUGGER_EVENT_OPTIONS Options;

    PVOID CommandStringBuffer;

    UINT32 ConditionBufferSize;

} DEBUGGER_GENERAL_EVENT_DETAIL, *PDEBUGGER_GENERAL_EVENT_DETAIL;

/**
 * @brief Each event can have multiple actions
 * @details THIS STRUCTURE IS ONLY USED IN USER MODE
 * WE USE SEPARATE STRUCTURE FOR ACTIONS IN
 * KERNEL MODE
 */
typedef struct _DEBUGGER_GENERAL_ACTION
{
    UINT64                          EventTag;
    DEBUGGER_EVENT_ACTION_TYPE_ENUM ActionType;
    BOOLEAN                         ImmediateMessagePassing;
    UINT32                          PreAllocatedBuffer;

    UINT32 CustomCodeBufferSize;
    UINT32 ScriptBufferSize;
    UINT32 ScriptBufferPointer;

} DEBUGGER_GENERAL_ACTION, *PDEBUGGER_GENERAL_ACTION;

/**
 * @brief Status of register buffers
 *
 */
typedef struct _DEBUGGER_EVENT_AND_ACTION_RESULT
{
    BOOLEAN IsSuccessful;
    UINT32  Error; // If IsSuccessful was, FALSE

} DEBUGGER_EVENT_AND_ACTION_RESULT, *PDEBUGGER_EVENT_AND_ACTION_RESULT;

#define SIZEOF_REGISTER_EVENT sizeof(REGISTER_NOTIFY_BUFFER)


//..\bin\debug\SDK\Headers\HardwareDebugger.h
/**
 * @file HardwareDebugger.h
 * @author Sina Karvandi (sina@hyperdbg.org)
 * @brief HyperDbg's Hardware Debugger (hwdbg) types and constants
 * @details This file contains definitions of hwdbg elements
 * used in HyperDbg
 * @version 0.9
 * @date 2024-04-28
 *
 * @copyright This project is released under the GNU Public License v3.
 *
 */
#pragma once

//////////////////////////////////////////////////
//                 Definitions                  //
//////////////////////////////////////////////////

/**
 * @brief Initial debuggee to debugger offset
 *
 */
#define DEFAULT_INITIAL_DEBUGGEE_TO_DEBUGGER_OFFSET 0x200

/**
 * @brief Initial debugger to debuggee offset
 *
 */
#define DEFAULT_INITIAL_DEBUGGER_TO_DEBUGGEE_OFFSET 0x0

//////////////////////////////////////////////////
//                   Enums                      //
//////////////////////////////////////////////////

/**
 * @brief Different action of hwdbg
 * @warning This file should be changed along with hwdbg files
 *
 */
typedef enum _HWDBG_ACTION_ENUMS
{
    hwdbgActionSendInstanceInfo      = 1,
    hwdbgActionConfigureScriptBuffer = 2,

} HWDBG_ACTION_ENUMS;

/**
 * @brief Different responses come from hwdbg
 * @warning This file should be changed along with hwdbg files
 *
 */
typedef enum _HWDBG_RESPONSE_ENUMS
{
    hwdbgResponseSuccessOrErrorMessage = 1,
    hwdbgResponseInstanceInfo          = 2,

} HWDBG_RESPONSE_ENUMS;

/**
 * @brief Different success or error codes in hwdbg
 * @warning This file should be changed along with hwdbg files
 *
 */
typedef enum _HWDBG_SUCCESS_OR_ERROR_ENUMS
{
    hwdbgOperationWasSuccessful = 0x7FFFFFFF,
    hwdbgErrorInvalidPacket     = 1,

} HWDBG_SUCCESS_OR_ERROR_ENUMS;

//////////////////////////////////////////////////
//                   Structures                 //
//////////////////////////////////////////////////

/**
 * @brief The structure of port information (each item) in hwdbg
 *
 */
typedef struct _HWDBG_PORT_INFORMATION_ITEMS
{
    UINT32 PortSize;

} HWDBG_PORT_INFORMATION_ITEMS, *PHWDBG_PORT_INFORMATION_ITEMS;

/**
 * @brief The structure of script capabilities information in hwdbg
 *
 */
typedef struct _HWDBG_INSTANCE_INFORMATION
{
    //
    // ANY ADDITION TO THIS STRUCTURE SHOULD BE SYNCHRONIZED WITH SCALA AND INSTANCE INFO SENDER MODULE
    //
    UINT32 version;                                    // Target version of HyperDbg (same as hwdbg)
    UINT32 maximumNumberOfStages;                      // Number of stages that this instance of hwdbg supports (NumberOfSupportedStages == 0 means script engine is disabled)
    UINT32 scriptVariableLength;                       // maximum length of variables (and other script elements)
    UINT32 maximumNumberOfSupportedGetScriptOperators; // Maximum supported GET operators in a single func
    UINT32 maximumNumberOfSupportedSetScriptOperators; // Maximum supported SET operators in a single func
    UINT32 sharedMemorySize;                           // Size of shared memory
    UINT32 debuggerAreaOffset;                         // The memory offset of debugger
    UINT32 debuggeeAreaOffset;                         // The memory offset of debuggee
    UINT32 numberOfPins;                               // Number of pins
    UINT32 numberOfPorts;                              // Number of ports

    //
    // ANY ADDITION TO THIS STRUCTURE SHOULD BE SYNCHRONIZED WITH SCALA AND INSTANCE INFO SENDER MODULE
    //

    struct _HWDBG_SCRIPT_CAPABILITIES
    {
        //
        // ANY ADDITION TO THIS MASK SHOULD BE ADDED TO HwdbgInterpreterShowScriptCapabilities
        // and HwdbgInterpreterCheckScriptBufferWithScriptCapabilities as well Scala file
        //
        UINT64 func_or : 1;
        UINT64 func_xor : 1;
        UINT64 func_and : 1;
        UINT64 func_asr : 1;
        UINT64 func_asl : 1;
        UINT64 func_add : 1;
        UINT64 func_sub : 1;
        UINT64 func_mul : 1;
        UINT64 func_div : 1;
        UINT64 func_mod : 1;
        UINT64 func_gt : 1;
        UINT64 func_lt : 1;
        UINT64 func_egt : 1;
        UINT64 func_elt : 1;
        UINT64 func_equal : 1;
        UINT64 func_neq : 1;
        UINT64 func_jmp : 1;
        UINT64 func_jz : 1;
        UINT64 func_jnz : 1;
        UINT64 func_mov : 1;
        UINT64 func_printf : 1;

        //
        // ANY ADDITION TO THIS MASK SHOULD BE ADDED TO HwdbgInterpreterShowScriptCapabilities
        // and HwdbgInterpreterCheckScriptBufferWithScriptCapabilities as well Scala file
        //

    } scriptCapabilities;

    UINT32 bramAddrWidth; // BRAM address width
    UINT32 bramDataWidth; // BRAM data width

    //
    // Here the details of port arrangements are located (HWDBG_PORT_INFORMATION_ITEMS)
    // As the following type:
    //   HWDBG_PORT_INFORMATION_ITEMS portsConfiguration[numberOfPorts]   ; Port arrangement
    //

} HWDBG_INSTANCE_INFORMATION, *PHWDBG_INSTANCE_INFORMATION;

/**
 * @brief The structure of script buffer in hwdbg
 *
 */
typedef struct _HWDBG_SCRIPT_BUFFER
{
    UINT32 scriptNumberOfSymbols; // Number of symbols in the script

    //
    // Here the script buffer is located
    //
    // UINT8 scriptBuffer[scriptNumberOfSymbols]; // The script buffer
    //

} HWDBG_SCRIPT_BUFFER, *PHWDBG_SCRIPT_BUFFER;


//..\bin\debug\SDK\Headers\Ioctls.h
/**
 * @file Ioctls.h
 * @author Sina Karvandi (sina@hyperdbg.org)
 * @brief HyperDbg's SDK IOCTL codes
 * @details This file contains definitions of IOCTLs used in HyperDbg
 * @version 0.2
 * @date 2022-06-24
 *
 * @copyright This project is released under the GNU Public License v3.
 *
 */
#pragma once

//////////////////////////////////////////////////
//                 Definitions                  //
//////////////////////////////////////////////////

//
// The following controls are mainly defined in <winioctl.h>
//

//
// Macro definition for defining IOCTL and FSCTL function control codes.  Note
// that function codes 0-2047 are reserved for Microsoft Corporation, and
// 2048-4095 are reserved for customers.
//
#ifndef CTL_CODE

#    define CTL_CODE(DeviceType, Function, Method, Access) ( \
        ((DeviceType) << 16) | ((Access) << 14) | ((Function) << 2) | (Method))

#endif // ! CTL_CODE

#ifndef FILE_ANY_ACCESS

#    define FILE_ANY_ACCESS 0

#endif // !FILE_ANY_ACCESS

//
// Define the method codes for how buffers are passed for I/O and FS controls
//

#ifndef METHOD_BUFFERED

#    define METHOD_BUFFERED 0

#endif // !METHOD_BUFFERED

#ifndef FILE_DEVICE_UNKNOWN

#    define FILE_DEVICE_UNKNOWN 0x00000022

#endif // !FILE_DEVICE_UNKNOWN

//////////////////////////////////////////////////
//                   IOCTLs                     //
//////////////////////////////////////////////////

/**
 * @brief ioctl, register a new event
 *
 */
#define IOCTL_REGISTER_EVENT \
    CTL_CODE(FILE_DEVICE_UNKNOWN, 0x800, METHOD_BUFFERED, FILE_ANY_ACCESS)

/**
 * @brief ioctl, irp pending mechanism for reading from message tracing buffers
 *
 */
#define IOCTL_RETURN_IRP_PENDING_PACKETS_AND_DISALLOW_IOCTL \
    CTL_CODE(FILE_DEVICE_UNKNOWN, 0x801, METHOD_BUFFERED, FILE_ANY_ACCESS)

/**
 * @brief ioctl, to terminate vmx and exit form debugger
 *
 */
#define IOCTL_TERMINATE_VMX \
    CTL_CODE(FILE_DEVICE_UNKNOWN, 0x802, METHOD_BUFFERED, FILE_ANY_ACCESS)

/**
 * @brief ioctl, request to read memory
 *
 */
#define IOCTL_DEBUGGER_READ_MEMORY \
    CTL_CODE(FILE_DEVICE_UNKNOWN, 0x803, METHOD_BUFFERED, FILE_ANY_ACCESS)

/**
 * @brief ioctl, request to read or write on a special MSR
 *
 */
#define IOCTL_DEBUGGER_READ_OR_WRITE_MSR \
    CTL_CODE(FILE_DEVICE_UNKNOWN, 0x804, METHOD_BUFFERED, FILE_ANY_ACCESS)

/**
 * @brief ioctl, request to read page table entries
 *
 */
#define IOCTL_DEBUGGER_READ_PAGE_TABLE_ENTRIES_DETAILS \
    CTL_CODE(FILE_DEVICE_UNKNOWN, 0x805, METHOD_BUFFERED, FILE_ANY_ACCESS)

/**
 * @brief ioctl, register an event
 *
 */
#define IOCTL_DEBUGGER_REGISTER_EVENT \
    CTL_CODE(FILE_DEVICE_UNKNOWN, 0x806, METHOD_BUFFERED, FILE_ANY_ACCESS)

/**
 * @brief ioctl, add action to event
 *
 */
#define IOCTL_DEBUGGER_ADD_ACTION_TO_EVENT \
    CTL_CODE(FILE_DEVICE_UNKNOWN, 0x807, METHOD_BUFFERED, FILE_ANY_ACCESS)

/**
 * @brief ioctl, request to enable or disable transparent-mode
 *
 */
#define IOCTL_DEBUGGER_HIDE_AND_UNHIDE_TO_TRANSPARENT_THE_DEBUGGER \
    CTL_CODE(FILE_DEVICE_UNKNOWN, 0x808, METHOD_BUFFERED, FILE_ANY_ACCESS)

/**
 * @brief ioctl, for !va2pa and !pa2va commands
 *
 */
#define IOCTL_DEBUGGER_VA2PA_AND_PA2VA_COMMANDS \
    CTL_CODE(FILE_DEVICE_UNKNOWN, 0x809, METHOD_BUFFERED, FILE_ANY_ACCESS)

/**
 * @brief ioctl, request to edit virtual and physical memory
 *
 */
#define IOCTL_DEBUGGER_EDIT_MEMORY \
    CTL_CODE(FILE_DEVICE_UNKNOWN, 0x80a, METHOD_BUFFERED, FILE_ANY_ACCESS)

/**
 * @brief ioctl, request to search virtual and physical memory
 *
 */
#define IOCTL_DEBUGGER_SEARCH_MEMORY \
    CTL_CODE(FILE_DEVICE_UNKNOWN, 0x80b, METHOD_BUFFERED, FILE_ANY_ACCESS)

/**
 * @brief ioctl, request to modify an event (enable/disable/clear)
 *
 */
#define IOCTL_DEBUGGER_MODIFY_EVENTS \
    CTL_CODE(FILE_DEVICE_UNKNOWN, 0x80c, METHOD_BUFFERED, FILE_ANY_ACCESS)

/**
 * @brief ioctl, flush the kernel buffers
 *
 */
#define IOCTL_DEBUGGER_FLUSH_LOGGING_BUFFERS \
    CTL_CODE(FILE_DEVICE_UNKNOWN, 0x80d, METHOD_BUFFERED, FILE_ANY_ACCESS)

/**
 * @brief ioctl, attach or detach user-mode processes
 *
 */
#define IOCTL_DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS \
    CTL_CODE(FILE_DEVICE_UNKNOWN, 0x80e, METHOD_BUFFERED, FILE_ANY_ACCESS)

/**
 * @brief ioctl, print states (Deprecated)
 *
 *
 */
#define IOCTL_DEBUGGER_PRINT \
    CTL_CODE(FILE_DEVICE_UNKNOWN, 0x80f, METHOD_BUFFERED, FILE_ANY_ACCESS)

/**
 * @brief ioctl, prepare debuggee
 *
 */
#define IOCTL_PREPARE_DEBUGGEE \
    CTL_CODE(FILE_DEVICE_UNKNOWN, 0x810, METHOD_BUFFERED, FILE_ANY_ACCESS)

/**
 * @brief ioctl, pause and halt the system
 *
 */
#define IOCTL_PAUSE_PACKET_RECEIVED \
    CTL_CODE(FILE_DEVICE_UNKNOWN, 0x811, METHOD_BUFFERED, FILE_ANY_ACCESS)

/**
 * @brief ioctl, send a signal that execution of command finished
 *
 */
#define IOCTL_SEND_SIGNAL_EXECUTION_IN_DEBUGGEE_FINISHED \
    CTL_CODE(FILE_DEVICE_UNKNOWN, 0x812, METHOD_BUFFERED, FILE_ANY_ACCESS)

/**
 * @brief ioctl, send user-mode messages to the debugger
 *
 */
#define IOCTL_SEND_USERMODE_MESSAGES_TO_DEBUGGER \
    CTL_CODE(FILE_DEVICE_UNKNOWN, 0x813, METHOD_BUFFERED, FILE_ANY_ACCESS)

/**
 * @brief ioctl, send general buffer from debuggee to debugger
 *
 */
#define IOCTL_SEND_GENERAL_BUFFER_FROM_DEBUGGEE_TO_DEBUGGER \
    CTL_CODE(FILE_DEVICE_UNKNOWN, 0x814, METHOD_BUFFERED, FILE_ANY_ACCESS)

/**
 * @brief ioctl, to perform kernel-side tests
 *
 */
#define IOCTL_PERFROM_KERNEL_SIDE_TESTS \
    CTL_CODE(FILE_DEVICE_UNKNOWN, 0x815, METHOD_BUFFERED, FILE_ANY_ACCESS)

/**
 * @brief ioctl, to reserve pre-allocated pools
 *
 */
#define IOCTL_RESERVE_PRE_ALLOCATED_POOLS \
    CTL_CODE(FILE_DEVICE_UNKNOWN, 0x816, METHOD_BUFFERED, FILE_ANY_ACCESS)

/**
 * @brief ioctl, to send user debugger commands
 *
 */
#define IOCTL_SEND_USER_DEBUGGER_COMMANDS \
    CTL_CODE(FILE_DEVICE_UNKNOWN, 0x817, METHOD_BUFFERED, FILE_ANY_ACCESS)

/**
 * @brief ioctl, to get active threads/processes that are debugging
 *
 */
#define IOCTL_GET_DETAIL_OF_ACTIVE_THREADS_AND_PROCESSES \
    CTL_CODE(FILE_DEVICE_UNKNOWN, 0x818, METHOD_BUFFERED, FILE_ANY_ACCESS)

/**
 * @brief ioctl, to get user mode modules details
 *
 */
#define IOCTL_GET_USER_MODE_MODULE_DETAILS \
    CTL_CODE(FILE_DEVICE_UNKNOWN, 0x819, METHOD_BUFFERED, FILE_ANY_ACCESS)

/**
 * @brief ioctl, query count of active threads or processes
 *
 */
#define IOCTL_QUERY_COUNT_OF_ACTIVE_PROCESSES_OR_THREADS \
    CTL_CODE(FILE_DEVICE_UNKNOWN, 0x81a, METHOD_BUFFERED, FILE_ANY_ACCESS)

/**
 * @brief ioctl, to get list threads/processes
 *
 */
#define IOCTL_GET_LIST_OF_THREADS_AND_PROCESSES \
    CTL_CODE(FILE_DEVICE_UNKNOWN, 0x81b, METHOD_BUFFERED, FILE_ANY_ACCESS)

/**
 * @brief ioctl, query the current process details
 *
 */
#define IOCTL_QUERY_CURRENT_PROCESS \
    CTL_CODE(FILE_DEVICE_UNKNOWN, 0x81c, METHOD_BUFFERED, FILE_ANY_ACCESS)

/**
 * @brief ioctl, query the current thread details
 *
 */
#define IOCTL_QUERY_CURRENT_THREAD \
    CTL_CODE(FILE_DEVICE_UNKNOWN, 0x81d, METHOD_BUFFERED, FILE_ANY_ACCESS)

/**
 * @brief ioctl, request service from the reversing machine
 *
 */
#define IOCTL_REQUEST_REV_MACHINE_SERVICE \
    CTL_CODE(FILE_DEVICE_UNKNOWN, 0x81e, METHOD_BUFFERED, FILE_ANY_ACCESS)

/**
 * @brief ioctl, request to bring pages in
 *
 */
#define IOCTL_DEBUGGER_BRING_PAGES_IN \
    CTL_CODE(FILE_DEVICE_UNKNOWN, 0x81f, METHOD_BUFFERED, FILE_ANY_ACCESS)

/**
 * @brief ioctl, to preactivate a functionality
 *
 */
#define IOCTL_PREACTIVATE_FUNCTIONALITY \
    CTL_CODE(FILE_DEVICE_UNKNOWN, 0x820, METHOD_BUFFERED, FILE_ANY_ACCESS)


//..\bin\debug\SDK\Headers\RequestStructures.h
/**
 * @file RequestStructures.h
 * @author Sina Karvandi (sina@hyperdbg.org)
 * @brief HyperDbg's SDK Headers Request Packets
 * @details This file contains definitions of request packets (enums, structs)
 * @version 0.2
 * @date 2022-06-28
 *
 * @copyright This project is released under the GNU Public License v3.
 *
 */
#pragma once

#define SIZEOF_DEBUGGER_READ_PAGE_TABLE_ENTRIES_DETAILS \
    sizeof(DEBUGGER_READ_PAGE_TABLE_ENTRIES_DETAILS)

/**
 * @brief request for !pte command
 *
 */
typedef struct _DEBUGGER_READ_PAGE_TABLE_ENTRIES_DETAILS
{
    UINT64 VirtualAddress;
    UINT32 ProcessId;

    UINT64 Pml4eVirtualAddress;
    UINT64 Pml4eValue;

    UINT64 PdpteVirtualAddress;
    UINT64 PdpteValue;

    UINT64 PdeVirtualAddress;
    UINT64 PdeValue;

    UINT64 PteVirtualAddress;
    UINT64 PteValue;

    UINT32 KernelStatus;

} DEBUGGER_READ_PAGE_TABLE_ENTRIES_DETAILS,
    *PDEBUGGER_READ_PAGE_TABLE_ENTRIES_DETAILS;

/* ==============================================================================================
 */

#define SIZEOF_DEBUGGER_VA2PA_AND_PA2VA_COMMANDS \
    sizeof(DEBUGGER_VA2PA_AND_PA2VA_COMMANDS)

/**
 * @brief requests for !va2pa and !pa2va commands
 *
 */
typedef struct _DEBUGGER_VA2PA_AND_PA2VA_COMMANDS
{
    UINT64  VirtualAddress;
    UINT64  PhysicalAddress;
    UINT32  ProcessId;
    BOOLEAN IsVirtual2Physical;
    UINT32  KernelStatus;

} DEBUGGER_VA2PA_AND_PA2VA_COMMANDS, *PDEBUGGER_VA2PA_AND_PA2VA_COMMANDS;

/* ==============================================================================================
 */
#define SIZEOF_DEBUGGER_PAGE_IN_REQUEST \
    sizeof(DEBUGGER_PAGE_IN_REQUEST)

/**
 * @brief requests for the '.pagein' command
 *
 */
typedef struct _DEBUGGER_PAGE_IN_REQUEST
{
    UINT64 VirtualAddressFrom;
    UINT64 VirtualAddressTo;
    UINT32 ProcessId;
    UINT32 PageFaultErrorCode;
    UINT32 KernelStatus;

} DEBUGGER_PAGE_IN_REQUEST, *PDEBUGGER_PAGE_IN_REQUEST;

/* ==============================================================================================
 */

/**
 * @brief different modes of reconstruct requests
 *
 */
typedef enum _REVERSING_MACHINE_RECONSTRUCT_MEMORY_MODE
{
    REVERSING_MACHINE_RECONSTRUCT_MEMORY_MODE_UNKNOWN = 0,
    REVERSING_MACHINE_RECONSTRUCT_MEMORY_MODE_USER_MODE,
    REVERSING_MACHINE_RECONSTRUCT_MEMORY_MODE_KERNEL_MODE,
} REVERSING_MACHINE_RECONSTRUCT_MEMORY_MODE;

/**
 * @brief different types of reconstruct requests
 *
 */
typedef enum _REVERSING_MACHINE_RECONSTRUCT_MEMORY_TYPE
{
    REVERSING_MACHINE_RECONSTRUCT_MEMORY_TYPE_UNKNOWN = 0,
    REVERSING_MACHINE_RECONSTRUCT_MEMORY_TYPE_RECONSTRUCT,
    REVERSING_MACHINE_RECONSTRUCT_MEMORY_TYPE_PATTERN,
} REVERSING_MACHINE_RECONSTRUCT_MEMORY_TYPE;

#define SIZEOF_REVERSING_MACHINE_RECONSTRUCT_MEMORY_REQUEST \
    sizeof(REVERSING_MACHINE_RECONSTRUCT_MEMORY_REQUEST)

/**
 * @brief requests for !rev command
 *
 */
typedef struct _REVERSING_MACHINE_RECONSTRUCT_MEMORY_REQUEST
{
    UINT32                                    ProcessId;
    UINT32                                    Size;
    REVERSING_MACHINE_RECONSTRUCT_MEMORY_MODE Mode;
    REVERSING_MACHINE_RECONSTRUCT_MEMORY_TYPE Type;
    UINT32                                    KernelStatus;

} REVERSING_MACHINE_RECONSTRUCT_MEMORY_REQUEST, *PREVERSING_MACHINE_RECONSTRUCT_MEMORY_REQUEST;

/* ==============================================================================================
 */

#define SIZEOF_DEBUGGER_DT_COMMAND_OPTIONS \
    sizeof(DEBUGGER_DT_COMMAND_OPTIONS)

/**
 * @brief requests options for dt and struct command
 *
 */
typedef struct _DEBUGGER_DT_COMMAND_OPTIONS
{
    const char * TypeName;
    UINT64       SizeOfTypeName;
    UINT64       Address;
    BOOLEAN      IsStruct;
    PVOID        BufferAddress;
    UINT32       TargetPid;
    const char * AdditionalParameters;

} DEBUGGER_DT_COMMAND_OPTIONS, *PDEBUGGER_DT_COMMAND_OPTIONS;

/* ==============================================================================================
 */

/**
 * @brief different types of prealloc requests
 *
 */
typedef enum _DEBUGGER_PREALLOC_COMMAND_TYPE
{
    DEBUGGER_PREALLOC_COMMAND_TYPE_THREAD_INTERCEPTION,
    DEBUGGER_PREALLOC_COMMAND_TYPE_MONITOR,
    DEBUGGER_PREALLOC_COMMAND_TYPE_EPTHOOK,
    DEBUGGER_PREALLOC_COMMAND_TYPE_EPTHOOK2,
    DEBUGGER_PREALLOC_COMMAND_TYPE_REGULAR_EVENT,
    DEBUGGER_PREALLOC_COMMAND_TYPE_BIG_EVENT,
    DEBUGGER_PREALLOC_COMMAND_TYPE_REGULAR_SAFE_BUFFER,
    DEBUGGER_PREALLOC_COMMAND_TYPE_BIG_SAFE_BUFFER,

} DEBUGGER_PREALLOC_COMMAND_TYPE;

#define SIZEOF_DEBUGGER_PREALLOC_COMMAND \
    sizeof(DEBUGGER_PREALLOC_COMMAND)

/**
 * @brief requests for the 'prealloc' command
 *
 */
typedef struct _DEBUGGER_PREALLOC_COMMAND
{
    DEBUGGER_PREALLOC_COMMAND_TYPE Type;
    UINT32                         Count;
    UINT32                         KernelStatus;

} DEBUGGER_PREALLOC_COMMAND, *PDEBUGGER_PREALLOC_COMMAND;

/* ==============================================================================================
 */

/**
 * @brief different types of preactivate requests
 *
 */
typedef enum _DEBUGGER_PREACTIVATE_COMMAND_TYPE
{
    DEBUGGER_PREACTIVATE_COMMAND_TYPE_MODE,

} DEBUGGER_PREACTIVATE_COMMAND_TYPE;

#define SIZEOF_DEBUGGER_PREACTIVATE_COMMAND \
    sizeof(DEBUGGER_PREACTIVATE_COMMAND)

/**
 * @brief requests for the 'preactivate' command
 *
 */
typedef struct _DEBUGGER_PREACTIVATE_COMMAND
{
    DEBUGGER_PREACTIVATE_COMMAND_TYPE Type;
    UINT32                            KernelStatus;

} DEBUGGER_PREACTIVATE_COMMAND, *PDEBUGGER_PREACTIVATE_COMMAND;

/* ==============================================================================================
 */

#define SIZEOF_DEBUGGER_READ_MEMORY sizeof(DEBUGGER_READ_MEMORY)

/**
 * @brief different types of reading memory
 *
 */
typedef enum _DEBUGGER_READ_READING_TYPE
{
    READ_FROM_KERNEL,
    READ_FROM_VMX_ROOT
} DEBUGGER_READ_READING_TYPE;

/**
 * @brief different type of addresses
 *
 */
typedef enum _DEBUGGER_READ_MEMORY_TYPE
{
    DEBUGGER_READ_PHYSICAL_ADDRESS,
    DEBUGGER_READ_VIRTUAL_ADDRESS
} DEBUGGER_READ_MEMORY_TYPE;

/**
 * @brief the way that debugger should show
 * the details of memory or disassemble them
 *
 */
typedef enum _DEBUGGER_SHOW_MEMORY_STYLE
{
    DEBUGGER_SHOW_COMMAND_DT = 1,
    DEBUGGER_SHOW_COMMAND_DISASSEMBLE64,
    DEBUGGER_SHOW_COMMAND_DISASSEMBLE32,
    DEBUGGER_SHOW_COMMAND_DB,
    DEBUGGER_SHOW_COMMAND_DC,
    DEBUGGER_SHOW_COMMAND_DQ,
    DEBUGGER_SHOW_COMMAND_DD,
    DEBUGGER_SHOW_COMMAND_DUMP
} DEBUGGER_SHOW_MEMORY_STYLE;

/**
 * @brief request for reading virtual and physical memory
 *
 */
typedef struct _DEBUGGER_READ_MEMORY
{
    UINT32                       Pid; // Read from cr3 of what process
    UINT64                       Address;
    UINT32                       Size;
    BOOLEAN                      IsForDisasm;    // Debugger sets whether the read memory is for diassembler or not
    BOOLEAN                      Is32BitAddress; // Debuggee sets the status of address
    DEBUGGER_READ_MEMORY_TYPE    MemoryType;
    DEBUGGER_READ_READING_TYPE   ReadingType;
    PDEBUGGER_DT_COMMAND_OPTIONS DtDetails;
    DEBUGGER_SHOW_MEMORY_STYLE   Style;        // not used in local debugging
    UINT32                       ReturnLength; // not used in local debugging
    UINT32                       KernelStatus; // not used in local debugging

    //
    // Here is the target buffer (actual memory)
    //

} DEBUGGER_READ_MEMORY, *PDEBUGGER_READ_MEMORY;

/* ==============================================================================================
 */

#define SIZEOF_DEBUGGER_FLUSH_LOGGING_BUFFERS \
    sizeof(DEBUGGER_FLUSH_LOGGING_BUFFERS)

/**
 * @brief request for flushing buffers
 *
 */
typedef struct _DEBUGGER_FLUSH_LOGGING_BUFFERS
{
    UINT32 KernelStatus;
    UINT32 CountOfMessagesThatSetAsReadFromVmxRoot;
    UINT32 CountOfMessagesThatSetAsReadFromVmxNonRoot;

} DEBUGGER_FLUSH_LOGGING_BUFFERS, *PDEBUGGER_FLUSH_LOGGING_BUFFERS;

/* ==============================================================================================
 */

#define SIZEOF_DEBUGGER_TEST_QUERY_BUFFER \
    sizeof(DEBUGGER_TEST_QUERY_BUFFER)

/**
 * @brief test query used for test purposed
 *
 */
typedef enum _DEBUGGER_TEST_QUERY_STATE
{
    TEST_QUERY_HALTING_CORE_STATUS                                          = 1,  // Query constant to show detail of halting of core
    TEST_QUERY_PREALLOCATED_POOL_STATE                                      = 2,  // Query pre-allocated pool state
    TEST_QUERY_TRAP_STATE                                                   = 3,  // Query trap state
    TEST_BREAKPOINT_TURN_OFF_BPS                                            = 4,  // Turn off the breakpoints (#BP)
    TEST_BREAKPOINT_TURN_ON_BPS                                             = 5,  // Turn on the breakpoints (#BP)
    TEST_BREAKPOINT_TURN_OFF_BPS_AND_EVENTS_FOR_COMMANDS_IN_REMOTE_COMPUTER = 6,  // Turn off the breakpoints and events for executing the commands in the remote computer
    TEST_BREAKPOINT_TURN_ON_BPS_AND_EVENTS_FOR_COMMANDS_IN_REMOTE_COMPUTER  = 7,  // Turn on the breakpoints and events for executing the commands in the remote computer
    TEST_SETTING_TARGET_TASKS_ON_HALTED_CORES_SYNCHRONOUS                   = 8,  // For testing synchronized event
    TEST_SETTING_TARGET_TASKS_ON_HALTED_CORES_ASYNCHRONOUS                  = 9,  // For testing unsynchronized event
    TEST_SETTING_TARGET_TASKS_ON_TARGET_HALTED_CORES                        = 10, // Send the task to the halted core
    TEST_BREAKPOINT_TURN_OFF_DBS                                            = 11, // Turn off the debug breaks (#DB)
    TEST_BREAKPOINT_TURN_ON_DBS                                             = 12, // Turn on the debug breaks (#DB)

} DEBUGGER_TEST_QUERY_STATE;

/**
 * @brief request for test query buffers
 *
 */
typedef struct _DEBUGGER_DEBUGGER_TEST_QUERY_BUFFER
{
    DEBUGGER_TEST_QUERY_STATE RequestType;
    UINT64                    Context;
    UINT32                    KernelStatus;

} DEBUGGER_DEBUGGER_TEST_QUERY_BUFFER, *PDEBUGGER_DEBUGGER_TEST_QUERY_BUFFER;

/* ==============================================================================================
 */

#define SIZEOF_DEBUGGER_PERFORM_KERNEL_TESTS \
    sizeof(DEBUGGER_PERFORM_KERNEL_TESTS)

/**
 * @brief request performing kernel tests
 *
 */
typedef struct _DEBUGGER_PERFORM_KERNEL_TESTS
{
    UINT32 KernelStatus;

} DEBUGGER_PERFORM_KERNEL_TESTS, *PDEBUGGER_PERFORM_KERNEL_TESTS;

/* ==============================================================================================
 */

#define SIZEOF_DEBUGGER_SEND_COMMAND_EXECUTION_FINISHED_SIGNAL \
    sizeof(DEBUGGER_SEND_COMMAND_EXECUTION_FINISHED_SIGNAL)

/**
 * @brief request for send a signal that command execution finished
 *
 */
typedef struct _DEBUGGER_SEND_COMMAND_EXECUTION_FINISHED_SIGNAL
{
    UINT32 KernelStatus;

} DEBUGGER_SEND_COMMAND_EXECUTION_FINISHED_SIGNAL,
    *PDEBUGGER_SEND_COMMAND_EXECUTION_FINISHED_SIGNAL;

/* ==============================================================================================
 */

#define SIZEOF_DEBUGGEE_SEND_GENERAL_PACKET_FROM_DEBUGGEE_TO_DEBUGGER \
    sizeof(DEBUGGEE_SEND_GENERAL_PACKET_FROM_DEBUGGEE_TO_DEBUGGER)

/**
 * @brief request for send general packets from debuggee to debugger
 *
 */
typedef struct _DEBUGGEE_SEND_GENERAL_PACKET_FROM_DEBUGGEE_TO_DEBUGGER
{
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION RequestedAction;
    UINT32                                  LengthOfBuffer;
    BOOLEAN                                 PauseDebuggeeWhenSent;
    UINT32                                  KernelResult;

    //
    // The buffer for the general packet is here
    //

} DEBUGGEE_SEND_GENERAL_PACKET_FROM_DEBUGGEE_TO_DEBUGGER,
    *PDEBUGGEE_SEND_GENERAL_PACKET_FROM_DEBUGGEE_TO_DEBUGGER;

/* ==============================================================================================
 */

#define SIZEOF_DEBUGGER_SEND_USERMODE_MESSAGES_TO_DEBUGGER \
    sizeof(DEBUGGER_SEND_USERMODE_MESSAGES_TO_DEBUGGER)

/**
 * @brief request for send a user-mode message to debugger
 *
 */
typedef struct _DEBUGGER_SEND_USERMODE_MESSAGES_TO_DEBUGGER
{
    UINT32 KernelStatus;
    UINT32 Length;

    //
    // Here is the messages
    //

} DEBUGGER_SEND_USERMODE_MESSAGES_TO_DEBUGGER,
    *PDEBUGGER_SEND_USERMODE_MESSAGES_TO_DEBUGGER;

/* ==============================================================================================
 */

#define SIZEOF_DEBUGGER_READ_AND_WRITE_ON_MSR \
    sizeof(DEBUGGER_READ_AND_WRITE_ON_MSR)

/**
 * @brief different types of actions on MSRs
 *
 */
typedef enum _DEBUGGER_MSR_ACTION_TYPE
{
    DEBUGGER_MSR_READ,
    DEBUGGER_MSR_WRITE
} DEBUGGER_MSR_ACTION_TYPE;

/**
 * @brief request to read or write on MSRs
 *
 */
typedef struct _DEBUGGER_READ_AND_WRITE_ON_MSR
{
    UINT64 Msr;        // It's actually a 32-Bit value but let's not mess with a register
    UINT32 CoreNumber; // specifies the core to execute wrmsr or read the msr
                       // (DEBUGGER_READ_AND_WRITE_ON_MSR_APPLY_ALL_CORES mean all
                       // the cores)
    DEBUGGER_MSR_ACTION_TYPE
    ActionType; // Detects whether user needs wrmsr or rdmsr
    UINT64 Value;

} DEBUGGER_READ_AND_WRITE_ON_MSR, *PDEBUGGER_READ_AND_WRITE_ON_MSR;

/* ==============================================================================================
 */

#define SIZEOF_DEBUGGER_EDIT_MEMORY sizeof(DEBUGGER_EDIT_MEMORY)

/**
 * @brief different type of addresses for editing memory
 *
 */
typedef enum _DEBUGGER_EDIT_MEMORY_TYPE
{
    EDIT_PHYSICAL_MEMORY,
    EDIT_VIRTUAL_MEMORY
} DEBUGGER_EDIT_MEMORY_TYPE;

/**
 * @brief size of editing memory
 *
 */
typedef enum _DEBUGGER_EDIT_MEMORY_BYTE_SIZE
{
    EDIT_BYTE,
    EDIT_DWORD,
    EDIT_QWORD
} DEBUGGER_EDIT_MEMORY_BYTE_SIZE;

/**
 * @brief request for edit virtual and physical memory
 *
 */
typedef struct _DEBUGGER_EDIT_MEMORY
{
    UINT32                         Result;     // Result from kernel
    UINT64                         Address;    // Target address to modify
    UINT32                         ProcessId;  // specifies the process id
    DEBUGGER_EDIT_MEMORY_TYPE      MemoryType; // Type of memory
    DEBUGGER_EDIT_MEMORY_BYTE_SIZE ByteSize;   // Modification size
    UINT32                         CountOf64Chunks;
    UINT32                         FinalStructureSize;
    UINT32                         KernelStatus; // not used in local debugging

} DEBUGGER_EDIT_MEMORY, *PDEBUGGER_EDIT_MEMORY;

/* ==============================================================================================
 */

#define SIZEOF_DEBUGGER_SEARCH_MEMORY sizeof(DEBUGGER_SEARCH_MEMORY)

/**
 * @brief different types of address for searching on memory
 *
 */
typedef enum _DEBUGGER_SEARCH_MEMORY_TYPE
{
    SEARCH_PHYSICAL_MEMORY,
    SEARCH_VIRTUAL_MEMORY,
    SEARCH_PHYSICAL_FROM_VIRTUAL_MEMORY,

} DEBUGGER_SEARCH_MEMORY_TYPE;

/**
 * @brief different sizes on searching memory
 *
 */
typedef enum _DEBUGGER_SEARCH_MEMORY_BYTE_SIZE
{
    SEARCH_BYTE,
    SEARCH_DWORD,
    SEARCH_QWORD

} DEBUGGER_SEARCH_MEMORY_BYTE_SIZE;

/**
 * @brief request for searching memory
 *
 */
typedef struct _DEBUGGER_SEARCH_MEMORY
{
    UINT64                           Address;    // Target address to start searching
    UINT64                           Length;     // Length of bytes to search
    UINT32                           ProcessId;  // specifies the process id
    DEBUGGER_SEARCH_MEMORY_TYPE      MemoryType; // Type of memory
    DEBUGGER_SEARCH_MEMORY_BYTE_SIZE ByteSize;   // Modification size
    UINT32                           CountOf64Chunks;
    UINT32                           FinalStructureSize;

} DEBUGGER_SEARCH_MEMORY, *PDEBUGGER_SEARCH_MEMORY;

/* ==============================================================================================
 */

#define SIZEOF_DEBUGGER_HIDE_AND_TRANSPARENT_DEBUGGER_MODE \
    sizeof(DEBUGGER_HIDE_AND_TRANSPARENT_DEBUGGER_MODE)

/**
 * @brief request for enable or disable transparent-mode
 *
 */
typedef struct _DEBUGGER_HIDE_AND_TRANSPARENT_DEBUGGER_MODE
{
    BOOLEAN IsHide;

    UINT64 CpuidAverage;
    UINT64 CpuidStandardDeviation;
    UINT64 CpuidMedian;

    UINT64 RdtscAverage;
    UINT64 RdtscStandardDeviation;
    UINT64 RdtscMedian;

    BOOLEAN TrueIfProcessIdAndFalseIfProcessName;
    UINT32  ProcId;
    UINT32  LengthOfProcessName; // in the case of !hide name xxx, this parameter
                                 // shows the length of xxx

    UINT64 KernelStatus; /* DEBUGGER_OPERATION_WAS_SUCCESSFUL ,
                          DEBUGGER_ERROR_UNABLE_TO_HIDE_OR_UNHIDE_DEBUGGER
                          */

} DEBUGGER_HIDE_AND_TRANSPARENT_DEBUGGER_MODE,
    *PDEBUGGER_HIDE_AND_TRANSPARENT_DEBUGGER_MODE;

/* ==============================================================================================
 */

#define SIZEOF_DEBUGGER_PREPARE_DEBUGGEE sizeof(DEBUGGER_PREPARE_DEBUGGEE)

/**
 * @brief request to make this computer to a debuggee
 *
 */
typedef struct _DEBUGGER_PREPARE_DEBUGGEE
{
    UINT32 PortAddress;
    UINT32 Baudrate;
    UINT64 NtoskrnlBaseAddress;
    UINT32 Result; // Result from the kernel
    CHAR   OsName[MAXIMUM_CHARACTER_FOR_OS_NAME];

} DEBUGGER_PREPARE_DEBUGGEE, *PDEBUGGER_PREPARE_DEBUGGEE;

/* ==============================================================================================
 */

/**
 * @brief The structure of changing core packet in HyperDbg
 *
 */
typedef struct _DEBUGGEE_CHANGE_CORE_PACKET
{
    UINT32 NewCore;
    UINT32 Result;

} DEBUGGEE_CHANGE_CORE_PACKET, *PDEBUGGEE_CHANGE_CORE_PACKET;

/* ==============================================================================================
 */
#define SIZEOF_DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS \
    sizeof(DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS)

/**
 * @brief different actions of switchings
 *
 */
typedef enum _DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS_ACTION_TYPE
{
    DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS_ACTION_ATTACH,
    DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS_ACTION_DETACH,
    DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS_ACTION_REMOVE_HOOKS,
    DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS_ACTION_KILL_PROCESS,
    DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS_ACTION_PAUSE_PROCESS,
    DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS_ACTION_SWITCH_BY_PROCESS_OR_THREAD,
    DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS_ACTION_QUERY_COUNT_OF_ACTIVE_DEBUGGING_THREADS,

} DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS_ACTION_TYPE;

/**
 * @brief request for attaching user-mode process
 *
 */
typedef struct _DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS
{
    BOOLEAN                                              IsStartingNewProcess;
    UINT32                                               ProcessId;
    UINT32                                               ThreadId;
    BOOLEAN                                              CheckCallbackAtFirstInstruction;
    BOOLEAN                                              Is32Bit;
    BOOLEAN                                              IsPaused; // used in switching to threads
    DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS_ACTION_TYPE Action;
    UINT32                                               CountOfActiveDebuggingThreadsAndProcesses; // used in showing the list of active threads/processes
    UINT64                                               Token;
    UINT64                                               Result;

} DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS,
    *PDEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS;

/* ==============================================================================================
 */
#define SIZEOF_DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS \
    sizeof(DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS)

/**
 * @brief different type of process or thread queries
 *
 */
typedef enum _DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_TYPES
{
    DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_QUERY_PROCESS_COUNT   = 1,
    DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_QUERY_THREAD_COUNT    = 2,
    DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_QUERY_PROCESS_LIST    = 3,
    DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_QUERY_THREAD_LIST     = 4,
    DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_QUERY_CURRENT_PROCESS = 5,
    DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_QUERY_CURRENT_THREAD  = 6,

} DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_TYPES;

/**
 * @brief different actions on showing or querying list of process or threads
 *
 */
typedef enum _DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_ACTIONS
{
    DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_ACTION_SHOW_INSTANTLY     = 1,
    DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_ACTION_QUERY_COUNT        = 2,
    DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_ACTION_QUERY_SAVE_DETAILS = 3,

} DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_ACTIONS;

/**
 * @brief The structure of needed information to get the details
 * of the process from nt!_EPROCESS and location of needed variables
 *
 */
typedef struct _DEBUGGEE_PROCESS_LIST_NEEDED_DETAILS
{
    UINT64 PsActiveProcessHead;      // nt!PsActiveProcessHead
    ULONG  ImageFileNameOffset;      // nt!_EPROCESS.ImageFileName
    ULONG  UniquePidOffset;          // nt!_EPROCESS.UniqueProcessId
    ULONG  ActiveProcessLinksOffset; // nt!_EPROCESS.ActiveProcessLinks

} DEBUGGEE_PROCESS_LIST_NEEDED_DETAILS, *PDEBUGGEE_PROCESS_LIST_NEEDED_DETAILS;

/**
 * @brief The structure of needed information to get the details
 * of the thread from nt!_ETHREAD and location of needed variables
 *
 */
typedef struct _DEBUGGEE_THREAD_LIST_NEEDED_DETAILS
{
    UINT32 ThreadListHeadOffset;     // nt!_EPROCESS.ThreadListHead
    UINT32 ThreadListEntryOffset;    // nt!_ETHREAD.ThreadListEntry
    UINT32 CidOffset;                // nt!_ETHREAD.Cid
    UINT64 PsActiveProcessHead;      // nt!PsActiveProcessHead
    ULONG  ActiveProcessLinksOffset; // nt!_EPROCESS.ActiveProcessLinks
    UINT64 Process;

} DEBUGGEE_THREAD_LIST_NEEDED_DETAILS, *PDEBUGGEE_THREAD_LIST_NEEDED_DETAILS;

/**
 * @brief The structure showing list of processes (details of each
 * entry)
 *
 */
typedef struct _DEBUGGEE_PROCESS_LIST_DETAILS_ENTRY
{
    UINT64 Eprocess;
    UINT32 ProcessId;
    UINT64 Cr3;
    UCHAR  ImageFileName[15 + 1];

} DEBUGGEE_PROCESS_LIST_DETAILS_ENTRY, *PDEBUGGEE_PROCESS_LIST_DETAILS_ENTRY;

/**
 * @brief The structure showing list of threads (details of each
 * entry)
 *
 */
typedef struct _DEBUGGEE_THREAD_LIST_DETAILS_ENTRY
{
    UINT64 Eprocess;
    UINT64 Ethread;
    UINT32 ProcessId;
    UINT32 ThreadId;
    UCHAR  ImageFileName[15 + 1];

} DEBUGGEE_THREAD_LIST_DETAILS_ENTRY, *PDEBUGGEE_THREAD_LIST_DETAILS_ENTRY;

/**
 * @brief request for query count of active processes and threads
 *
 */
typedef struct _DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS
{
    DEBUGGEE_PROCESS_LIST_NEEDED_DETAILS               ProcessListNeededDetails;
    DEBUGGEE_THREAD_LIST_NEEDED_DETAILS                ThreadListNeededDetails;
    DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_TYPES   QueryType;
    DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_ACTIONS QueryAction;
    UINT32                                             Count;
    UINT64                                             Result;

} DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS,
    *PDEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS;

/* ==============================================================================================
 */

/**
 * @brief The structure for saving the callstack frame of one parameter
 *
 */
typedef struct _DEBUGGER_SINGLE_CALLSTACK_FRAME
{
    BOOLEAN IsStackAddressValid;
    BOOLEAN IsValidAddress;
    BOOLEAN IsExecutable;
    UINT64  Value;
    BYTE    InstructionBytesOnRip[MAXIMUM_CALL_INSTR_SIZE];

} DEBUGGER_SINGLE_CALLSTACK_FRAME, *PDEBUGGER_SINGLE_CALLSTACK_FRAME;

#define SIZEOF_DEBUGGER_CALLSTACK_REQUEST \
    sizeof(DEBUGGER_CALLSTACK_REQUEST)

/**
 * @brief callstack showing method
 *
 */
typedef enum _DEBUGGER_CALLSTACK_DISPLAY_METHOD
{
    DEBUGGER_CALLSTACK_DISPLAY_METHOD_WITHOUT_PARAMS,
    DEBUGGER_CALLSTACK_DISPLAY_METHOD_WITH_PARAMS,

} DEBUGGER_CALLSTACK_DISPLAY_METHOD;

/**
 * @brief request for callstack frames
 *
 */
typedef struct _DEBUGGER_CALLSTACK_REQUEST
{
    BOOLEAN                           Is32Bit;
    UINT32                            KernelStatus;
    DEBUGGER_CALLSTACK_DISPLAY_METHOD DisplayMethod;
    UINT32                            Size;
    UINT32                            FrameCount;
    UINT64                            BaseAddress;
    UINT64                            BufferSize;

    //
    // Here is the size of stack frames
    //

} DEBUGGER_CALLSTACK_REQUEST, *PDEBUGGER_CALLSTACK_REQUEST;

/* ==============================================================================================
 */
#define SIZEOF_USERMODE_DEBUGGING_THREAD_OR_PROCESS_STATE_DETAILS \
    sizeof(USERMODE_DEBUGGING_THREAD_OR_PROCESS_STATE_DETAILS)

typedef struct _USERMODE_DEBUGGING_THREAD_OR_PROCESS_STATE_DETAILS
{
    UINT32  ProcessId;
    UINT32  ThreadId;
    BOOLEAN IsProcess;

} USERMODE_DEBUGGING_THREAD_OR_PROCESS_STATE_DETAILS, *PUSERMODE_DEBUGGING_THREAD_OR_PROCESS_STATE_DETAILS;

/* ==============================================================================================
 */

/**
 * @brief Used for run the script
 *
 */
typedef struct _DEBUGGER_EVENT_ACTION_RUN_SCRIPT_CONFIGURATION
{
    UINT64 ScriptBuffer;
    UINT32 ScriptLength;
    UINT32 ScriptPointer;
    UINT32 OptionalRequestedBufferSize;

} DEBUGGER_EVENT_ACTION_RUN_SCRIPT_CONFIGURATION,
    *PDEBUGGER_EVENT_ACTION_RUN_SCRIPT_CONFIGURATION;

/**
 * @brief used in the case of requesting a "request buffer"
 *
 */
typedef struct _DEBUGGER_EVENT_REQUEST_BUFFER
{
    BOOLEAN EnabledRequestBuffer;
    UINT32  RequestBufferSize;
    UINT64  RequstBufferAddress;

} DEBUGGER_EVENT_REQUEST_BUFFER, *PDEBUGGER_EVENT_REQUEST_BUFFER;

/**
 * @brief used in the case of custom code requests to the debugger
 *
 */
typedef struct _DEBUGGER_EVENT_REQUEST_CUSTOM_CODE
{
    UINT32 CustomCodeBufferSize;
    PVOID  CustomCodeBufferAddress;
    UINT32 OptionalRequestedBufferSize;

} DEBUGGER_EVENT_REQUEST_CUSTOM_CODE, *PDEBUGGER_EVENT_REQUEST_CUSTOM_CODE;

/* ==============================================================================================
 */

/**
 * @brief User-mode debugging actions
 *
 */
typedef enum _DEBUGGER_UD_COMMAND_ACTION_TYPE
{
    DEBUGGER_UD_COMMAND_ACTION_TYPE_NONE = 0,
    DEBUGGER_UD_COMMAND_ACTION_TYPE_PAUSE,
    DEBUGGER_UD_COMMAND_ACTION_TYPE_CONTINUE,
    DEBUGGER_UD_COMMAND_ACTION_TYPE_REGULAR_STEP,

} DEBUGGER_UD_COMMAND_ACTION_TYPE;

/**
 * @brief Description of user-mode debugging actions
 *
 */
typedef struct _DEBUGGER_UD_COMMAND_ACTION
{
    DEBUGGER_UD_COMMAND_ACTION_TYPE ActionType;
    UINT64                          OptionalParam1;
    UINT64                          OptionalParam2;
    UINT64                          OptionalParam3;
    UINT64                          OptionalParam4;

} DEBUGGER_UD_COMMAND_ACTION, *PDEBUGGER_UD_COMMAND_ACTION;

/**
 * @brief The structure of command packet in uHyperDbg
 *
 */
typedef struct _DEBUGGER_UD_COMMAND_PACKET
{
    DEBUGGER_UD_COMMAND_ACTION UdAction;
    UINT64                     ProcessDebuggingDetailToken;
    UINT32                     TargetThreadId;
    BOOLEAN                    ApplyToAllPausedThreads;
    UINT32                     Result;

} DEBUGGER_UD_COMMAND_PACKET, *PDEBUGGER_UD_COMMAND_PACKET;

/* ==============================================================================================
 */

/**
 * @brief Debugger process switch and process details
 *
 */
typedef enum _DEBUGGEE_DETAILS_AND_SWITCH_PROCESS_TYPE
{

    DEBUGGEE_DETAILS_AND_SWITCH_PROCESS_GET_PROCESS_DETAILS,
    DEBUGGEE_DETAILS_AND_SWITCH_PROCESS_GET_PROCESS_LIST,
    DEBUGGEE_DETAILS_AND_SWITCH_PROCESS_PERFORM_SWITCH,

} DEBUGGEE_DETAILS_AND_SWITCH_PROCESS_TYPE;

/**
 * @brief The structure of changing process and show process
 * packet in HyperDbg
 *
 */
typedef struct _DEBUGGEE_DETAILS_AND_SWITCH_PROCESS_PACKET
{
    DEBUGGEE_DETAILS_AND_SWITCH_PROCESS_TYPE ActionType;
    UINT32                                   ProcessId;
    UINT64                                   Process;
    BOOLEAN                                  IsSwitchByClkIntr;
    UCHAR                                    ProcessName[16];
    DEBUGGEE_PROCESS_LIST_NEEDED_DETAILS     ProcessListSymDetails;
    UINT32                                   Result;

} DEBUGGEE_DETAILS_AND_SWITCH_PROCESS_PACKET, *PDEBUGGEE_DETAILS_AND_SWITCH_PROCESS_PACKET;

/* ==============================================================================================
 */

/**
 * @brief Debugger size of DEBUGGEE_DETAILS_AND_SWITCH_PROCESS_PACKET
 *
 */
#define SIZEOF_DEBUGGEE_DETAILS_AND_SWITCH_PROCESS_PACKET \
    sizeof(DEBUGGEE_DETAILS_AND_SWITCH_PROCESS_PACKET)

/**
 * @brief Debugger thread switch and thread details
 *
 */
typedef enum _DEBUGGEE_DETAILS_AND_SWITCH_THREAD_TYPE
{

    DEBUGGEE_DETAILS_AND_SWITCH_THREAD_PERFORM_SWITCH,
    DEBUGGEE_DETAILS_AND_SWITCH_THREAD_GET_THREAD_DETAILS,
    DEBUGGEE_DETAILS_AND_SWITCH_THREAD_GET_THREAD_LIST,

} DEBUGGEE_DETAILS_AND_SWITCH_THREAD_TYPE;

/**
 * @brief The structure of changing thead and show thread
 * packet in HyperDbg
 */
typedef struct _DEBUGGEE_DETAILS_AND_SWITCH_THREAD_PACKET
{
    DEBUGGEE_DETAILS_AND_SWITCH_THREAD_TYPE ActionType;
    UINT32                                  ThreadId;
    UINT32                                  ProcessId;
    UINT64                                  Thread;
    UINT64                                  Process;
    BOOLEAN                                 CheckByClockInterrupt;
    UCHAR                                   ProcessName[16];
    DEBUGGEE_THREAD_LIST_NEEDED_DETAILS     ThreadListSymDetails;
    UINT32                                  Result;

} DEBUGGEE_DETAILS_AND_SWITCH_THREAD_PACKET, *PDEBUGGEE_DETAILS_AND_SWITCH_THREAD_PACKET;

/**
 * @brief Debugger size of DEBUGGEE_DETAILS_AND_SWITCH_THREAD_PACKET
 *
 */
#define SIZEOF_DEBUGGEE_DETAILS_AND_SWITCH_THREAD_PACKET \
    sizeof(DEBUGGEE_DETAILS_AND_SWITCH_THREAD_PACKET)

/* ==============================================================================================
 */

/**
 * @brief stepping and tracking types
 *
 */
typedef enum _DEBUGGER_REMOTE_STEPPING_REQUEST
{
    DEBUGGER_REMOTE_STEPPING_REQUEST_STEP_IN,
    DEBUGGER_REMOTE_STEPPING_REQUEST_INSTRUMENTATION_STEP_IN,
    DEBUGGER_REMOTE_STEPPING_REQUEST_INSTRUMENTATION_STEP_IN_FOR_TRACKING,

    DEBUGGER_REMOTE_STEPPING_REQUEST_STEP_OVER,
    DEBUGGER_REMOTE_STEPPING_REQUEST_STEP_OVER_FOR_GU,
    DEBUGGER_REMOTE_STEPPING_REQUEST_STEP_OVER_FOR_GU_LAST_INSTRUCTION,

} DEBUGGER_REMOTE_STEPPING_REQUEST;

/**
 * @brief The structure of stepping packet in HyperDbg
 *
 */
typedef struct _DEBUGGEE_STEP_PACKET
{
    DEBUGGER_REMOTE_STEPPING_REQUEST StepType;

    //
    // Only in the case of call instructions
    // the 'p' command
    //
    BOOLEAN IsCurrentInstructionACall;
    UINT32  CallLength;

} DEBUGGEE_STEP_PACKET, *PDEBUGGEE_STEP_PACKET;

/**
 * @brief default number of instructions used in tracking and stepping
 *
 */
#define DEBUGGER_REMOTE_TRACKING_DEFAULT_COUNT_OF_STEPPING 0xffffffff

/* ==============================================================================================
 */

/**
 * @brief The structure of .formats result packet in HyperDbg
 *
 */
typedef struct _DEBUGGEE_FORMATS_PACKET
{
    UINT64 Value;
    UINT32 Result;

} DEBUGGEE_FORMATS_PACKET, *PDEBUGGEE_FORMATS_PACKET;

/* ==============================================================================================
 */

/**
 * @brief The structure of .sym reload packet in HyperDbg
 *
 */
typedef struct _DEBUGGEE_SYMBOL_REQUEST_PACKET
{
    UINT32 ProcessId;

} DEBUGGEE_SYMBOL_REQUEST_PACKET, *PDEBUGGEE_SYMBOL_REQUEST_PACKET;

/* ==============================================================================================
 */

/**
 * @brief The structure of bp command packet in HyperDbg
 *
 */
typedef struct _DEBUGGEE_BP_PACKET
{
    UINT64  Address;
    UINT32  Pid;
    UINT32  Tid;
    UINT32  Core;
    BOOLEAN RemoveAfterHit;
    BOOLEAN CheckForCallbacks;
    UINT32  Result;

} DEBUGGEE_BP_PACKET, *PDEBUGGEE_BP_PACKET;

/**
 * @brief breakpoint modification types
 *
 */
typedef enum _DEBUGGEE_BREAKPOINT_MODIFICATION_REQUEST
{

    DEBUGGEE_BREAKPOINT_MODIFICATION_REQUEST_LIST_BREAKPOINTS,
    DEBUGGEE_BREAKPOINT_MODIFICATION_REQUEST_ENABLE,
    DEBUGGEE_BREAKPOINT_MODIFICATION_REQUEST_DISABLE,
    DEBUGGEE_BREAKPOINT_MODIFICATION_REQUEST_CLEAR,

} DEBUGGEE_BREAKPOINT_MODIFICATION_REQUEST;

/**
 * @brief The structure of breakpoint modification requests packet in HyperDbg
 *
 */
typedef struct _DEBUGGEE_BP_LIST_OR_MODIFY_PACKET
{
    UINT64                                   BreakpointId;
    DEBUGGEE_BREAKPOINT_MODIFICATION_REQUEST Request;
    UINT32                                   Result;

} DEBUGGEE_BP_LIST_OR_MODIFY_PACKET, *PDEBUGGEE_BP_LIST_OR_MODIFY_PACKET;

/* ==============================================================================================
 */

/**
 * @brief Whether a jump is taken or not taken
 *
 */
typedef enum _DEBUGGER_CONDITIONAL_JUMP_STATUS
{

    DEBUGGER_CONDITIONAL_JUMP_STATUS_ERROR = 0,
    DEBUGGER_CONDITIONAL_JUMP_STATUS_NOT_CONDITIONAL_JUMP,
    DEBUGGER_CONDITIONAL_JUMP_STATUS_JUMP_IS_TAKEN,
    DEBUGGER_CONDITIONAL_JUMP_STATUS_JUMP_IS_NOT_TAKEN,

} DEBUGGER_CONDITIONAL_JUMP_STATUS;

/* ==============================================================================================
 */

/**
 * @brief The structure of script packet in HyperDbg
 *
 */
typedef struct _DEBUGGEE_SCRIPT_PACKET
{
    UINT32  ScriptBufferSize;
    UINT32  ScriptBufferPointer;
    BOOLEAN IsFormat;
    UINT32  Result;

    //
    // The script buffer is here
    //

} DEBUGGEE_SCRIPT_PACKET, *PDEBUGGEE_SCRIPT_PACKET;

/* ==============================================================================================
 */

/**
 * @brief The structure of result of search packet in HyperDbg
 *
 */
typedef struct _DEBUGGEE_RESULT_OF_SEARCH_PACKET
{
    UINT32 CountOfResults;
    UINT32 Result;

} DEBUGGEE_RESULT_OF_SEARCH_PACKET, *PDEBUGGEE_RESULT_OF_SEARCH_PACKET;

/* ==============================================================================================
 */

/**
 * @brief Register Descriptor Structure to use in r command.
 *
 */
typedef struct _DEBUGGEE_REGISTER_READ_DESCRIPTION
{
    UINT32 RegisterID; // the number is from REGS_ENUM
    UINT64 Value;
    UINT32 KernelStatus;

} DEBUGGEE_REGISTER_READ_DESCRIPTION, *PDEBUGGEE_REGISTER_READ_DESCRIPTION;

/* ==============================================================================================
 */


//..\bin\debug\SDK\Headers\Symbols.h
/**
 * @file Symbols.h
 * @author Sina Karvandi (sina@hyperdbg.org)
 * @brief HyperDbg's SDK Header Files For Symbol Parsing
 * @details This file contains definitions of symbol parsers
 * @version 0.2
 * @date 2022-06-24
 *
 * @copyright This project is released under the GNU Public License v3.
 *
 */
#pragma once

//////////////////////////////////////////////////
//              Symbols Details                 //
//////////////////////////////////////////////////

/**
 * @brief structures for sending and saving details
 * about each module and symbols details
 *
 */
typedef struct _MODULE_SYMBOL_DETAIL
{
    BOOLEAN IsSymbolDetailsFound; // TRUE if the details of symbols found, FALSE if not found
    BOOLEAN IsLocalSymbolPath;    // TRUE if the ModuleSymbolPath is a real path
                                  // and FALSE if ModuleSymbolPath is just a module name
    BOOLEAN IsSymbolPDBAvaliable; // TRUE if the module's pdb is available(if exists in the sympath)
    BOOLEAN IsUserMode;           // TRUE if the module is a user-mode module
    BOOLEAN Is32Bit;              // TRUE if the module is a 32-bit
    UINT64  BaseAddress;
    char    FilePath[MAX_PATH];
    char    ModuleSymbolPath[MAX_PATH];
    char    ModuleSymbolGuidAndAge[MAXIMUM_GUID_AND_AGE_SIZE];

} MODULE_SYMBOL_DETAIL, *PMODULE_SYMBOL_DETAIL;

typedef struct _USERMODE_LOADED_MODULE_SYMBOLS
{
    UINT64  BaseAddress;
    UINT64  Entrypoint;
    wchar_t FilePath[MAX_PATH];

} USERMODE_LOADED_MODULE_SYMBOLS, *PUSERMODE_LOADED_MODULE_SYMBOLS;

typedef struct _USERMODE_LOADED_MODULE_DETAILS
{
    UINT32  ProcessId;
    BOOLEAN OnlyCountModules;
    BOOLEAN Is32Bit;
    UINT32  ModulesCount;
    UINT32  Result;

    //
    // Here is a list of USERMODE_LOADED_MODULE_SYMBOLS (appended)
    //

} USERMODE_LOADED_MODULE_DETAILS, *PUSERMODE_LOADED_MODULE_DETAILS;

/**
 * @brief Callback type that should be used to add
 * list of Addresses to ObjectNames
 *
 */
typedef VOID (*SymbolMapCallback)(UINT64 Address, char * ModuleName, char * ObjectName, unsigned int ObjectSize);

/**
 * @brief request to add new symbol detail or update a previous
 * symbol table entry
 *
 */
typedef struct _DEBUGGER_UPDATE_SYMBOL_TABLE
{
    UINT32               TotalSymbols;
    UINT32               CurrentSymbolIndex;
    MODULE_SYMBOL_DETAIL SymbolDetailPacket;

} DEBUGGER_UPDATE_SYMBOL_TABLE, *PDEBUGGER_UPDATE_SYMBOL_TABLE;

/*
==============================================================================================
 */

/**
 * @brief request that shows, symbol reload process is finished
 *
 */
typedef struct _DEBUGGEE_SYMBOL_UPDATE_RESULT
{
    UINT64 KernelStatus; // Kernel put the status in this field

} DEBUGGEE_SYMBOL_UPDATE_RESULT, *PDEBUGGEE_SYMBOL_UPDATE_RESULT;

/*
==============================================================================================
 */


//..\bin\debug\SDK\Modules\HyperLog.h
/**
 * @file HyperLog.h
 * @author Sina Karvandi (sina@hyperdbg.org)
 * @brief HyperDbg's SDK for HyperLog project
 * @details This file contains definitions of HyperLog routines
 * @version 0.2
 * @date 2023-01-15
 *
 * @copyright This project is released under the GNU Public License v3.
 *
 */
#pragma once

//////////////////////////////////////////////////
//			     Callback Types                 //
//////////////////////////////////////////////////

/**
 * @brief A function that checks whether the current operation
 * is on vmx-root mode or not
 *
 */
typedef BOOLEAN (*CHECK_VMX_OPERATION)();

/**
 * @brief A function that checks whether the immediate message
 * sending is needed or not
 *
 */
typedef BOOLEAN (*CHECK_IMMEDIATE_MESSAGE_SENDING)(UINT32 OperationCode);

/**
 * @brief A function that sends immediate messages
 *
 */
typedef BOOLEAN (*SEND_IMMEDIATE_MESSAGE)(CHAR * OptionalBuffer,
                                          UINT32 OptionalBufferLength,
                                          UINT32 OperationCode);

//////////////////////////////////////////////////
//			   Callback Structure               //
//////////////////////////////////////////////////

/**
 * @brief Prototype of each function needed by message tracer
 *
 */
typedef struct _MESSAGE_TRACING_CALLBACKS
{
    CHECK_VMX_OPERATION             VmxOperationCheck;
    CHECK_IMMEDIATE_MESSAGE_SENDING CheckImmediateMessageSending;
    SEND_IMMEDIATE_MESSAGE          SendImmediateMessage;

} MESSAGE_TRACING_CALLBACKS, *PMESSAGE_TRACING_CALLBACKS;


//..\bin\debug\SDK\Modules\VMM.h
/**
 * @file VMM.h
 * @author Sina Karvandi (sina@hyperdbg.org)
 * @brief HyperDbg's SDK for VMM project
 * @details This file contains definitions of HyperLog routines
 * @version 0.2
 * @date 2023-01-15
 *
 * @copyright This project is released under the GNU Public License v3.
 *
 */
#pragma once

//////////////////////////////////////////////////
//			     Callback Types                 //
//////////////////////////////////////////////////

/**
 * @brief A function from the message tracer that send the inputs to the
 * queue of the messages
 *
 */
typedef BOOLEAN (*LOG_CALLBACK_PREPARE_AND_SEND_MESSAGE_TO_QUEUE)(UINT32       OperationCode,
                                                                  BOOLEAN      IsImmediateMessage,
                                                                  BOOLEAN      ShowCurrentSystemTime,
                                                                  BOOLEAN      Priority,
                                                                  const char * Fmt,
                                                                  va_list      ArgList);

/**
 * @brief A function that sends the messages to message tracer buffers
 *
 */
typedef BOOLEAN (*LOG_CALLBACK_SEND_MESSAGE_TO_QUEUE)(UINT32 OperationCode, BOOLEAN IsImmediateMessage, CHAR * LogMessage, UINT32 BufferLen, BOOLEAN Priority);

/**
 * @brief A function that sends the messages to message tracer buffers
 *
 */
typedef BOOLEAN (*LOG_CALLBACK_SEND_BUFFER)(_In_ UINT32                          OperationCode,
                                            _In_reads_bytes_(BufferLength) PVOID Buffer,
                                            _In_ UINT32                          BufferLength,
                                            _In_ BOOLEAN                         Priority);

/**
 * @brief A function that checks whether the priority or regular buffer is full or not
 *
 */
typedef BOOLEAN (*LOG_CALLBACK_CHECK_IF_BUFFER_IS_FULL)(BOOLEAN Priority);

/**
 * @brief A function that handles trigger events
 *
 */
typedef VMM_CALLBACK_TRIGGERING_EVENT_STATUS_TYPE (*VMM_CALLBACK_TRIGGER_EVENTS)(VMM_EVENT_TYPE_ENUM                   EventType,
                                                                                 VMM_CALLBACK_EVENT_CALLING_STAGE_TYPE CallingStage,
                                                                                 PVOID                                 Context,
                                                                                 BOOLEAN *                             PostEventRequired,
                                                                                 GUEST_REGS *                          Regs);

/**
 * @brief A function that checks and handles breakpoints
 *
 */
typedef BOOLEAN (*DEBUGGING_CALLBACK_HANDLE_BREAKPOINT_EXCEPTION)(UINT32 CoreId);

/**
 * @brief A function that checks and handles debug breakpoints
 *
 */
typedef BOOLEAN (*DEBUGGING_CALLBACK_HANDLE_DEBUG_BREAKPOINT_EXCEPTION)(UINT32 CoreId);

/**
 * @brief Check for page-faults in user-debugger
 *
 */
typedef BOOLEAN (*DEBUGGING_CALLBACK_CONDITIONAL_PAGE_FAULT_EXCEPTION)(UINT32 CoreId,
                                                                       UINT64 Address,
                                                                       UINT32 PageFaultErrorCode);

/**
 * @brief Check for commands in user-debugger
 *
 */
typedef BOOLEAN (*UD_CHECK_FOR_COMMAND)();

/**
 * @brief Handle registered MTF callback
 *
 */
typedef VOID (*VMM_CALLBACK_REGISTERED_MTF_HANDLER)(UINT32 CoreId);

/**
 * @brief Check for user-mode access for loaded module details
 *
 */
typedef BOOLEAN (*VMM_CALLBACK_RESTORE_EPT_STATE)(UINT32 CoreId);

/**
 * @brief Check for unhandled EPT violations
 *
 */
typedef BOOLEAN (*VMM_CALLBACK_CHECK_UNHANDLED_EPT_VIOLATION)(UINT32 CoreId, UINT64 ViolationQualification, UINT64 GuestPhysicalAddr);

/**
 * @brief Handle cr3 process change callbacks
 *
 */
typedef VOID (*INTERCEPTION_CALLBACK_TRIGGER_CR3_CHANGE)(UINT32 CoreId);

/**
 * @brief Check for process or thread change callback
 *
 */
typedef BOOLEAN (*INTERCEPTION_CALLBACK_TRIGGER_CLOCK_AND_IPI)(_In_ UINT32 CoreId);

/**
 * @brief Check to handle cr3 events for thread interception
 *
 */
typedef BOOLEAN (*ATTACHING_HANDLE_CR3_EVENTS_FOR_THREAD_INTERCEPTION)(UINT32 CoreId, CR3_TYPE NewCr3);

/**
 * @brief Check and handle reapplying breakpoint
 *
 */
typedef BOOLEAN (*BREAKPOINT_CHECK_AND_HANDLE_REAPPLYING_BREAKPOINT)(UINT32 CoreId);

/**
 * @brief Handle NMI broadcast
 *
 */
typedef VOID (*VMM_CALLBACK_NMI_BROADCAST_REQUEST_HANDLER)(UINT32 CoreId, BOOLEAN IsOnVmxNmiHandler);

/**
 * @brief Check and handle NMI callbacks
 *
 */
typedef BOOLEAN (*KD_CHECK_AND_HANDLE_NMI_CALLBACK)(UINT32 CoreId);

/**
 * @brief Set the top-level driver's error status
 *
 */
typedef VOID (*VMM_CALLBACK_SET_LAST_ERROR)(UINT32 LastError);

/**
 * @brief Check and modify the protected resources of the hypervisor
 *
 */
typedef BOOLEAN (*VMM_CALLBACK_QUERY_TERMINATE_PROTECTED_RESOURCE)(UINT32                               CoreId,
                                                                   PROTECTED_HV_RESOURCES_TYPE          ResourceType,
                                                                   PVOID                                Context,
                                                                   PROTECTED_HV_RESOURCES_PASSING_OVERS PassOver);

/**
 * @brief Query debugger thread or process tracing details by core ID
 *
 */
typedef BOOLEAN (*KD_QUERY_DEBUGGER_THREAD_OR_PROCESS_TRACING_DETAILS_BY_CORE_ID)(UINT32                          CoreId,
                                                                                  DEBUGGER_THREAD_PROCESS_TRACING TracingType);
/**
 * @brief Handler of debugger specific VMCALLs
 *
 */
typedef BOOLEAN (*VMM_CALLBACK_VMCALL_HANDLER)(UINT32 CoreId,
                                               UINT64 VmcallNumber,
                                               UINT64 OptionalParam1,
                                               UINT64 OptionalParam2,
                                               UINT64 OptionalParam3);

//////////////////////////////////////////////////
//			   Callback Structure               //
//////////////////////////////////////////////////

/**
 * @brief Prototype of each function needed by VMM module
 *
 */
typedef struct _VMM_CALLBACKS
{
    //
    // Log (Hyperlog) callbacks
    //
    LOG_CALLBACK_PREPARE_AND_SEND_MESSAGE_TO_QUEUE LogCallbackPrepareAndSendMessageToQueueWrapper; // Fixed
    LOG_CALLBACK_SEND_MESSAGE_TO_QUEUE             LogCallbackSendMessageToQueue;                  // Fixed
    LOG_CALLBACK_SEND_BUFFER                       LogCallbackSendBuffer;                          // Fixed
    LOG_CALLBACK_CHECK_IF_BUFFER_IS_FULL           LogCallbackCheckIfBufferIsFull;                 // Fixed

    //
    // VMM callbacks
    //
    VMM_CALLBACK_TRIGGER_EVENTS                     VmmCallbackTriggerEvents;                   // Fixed
    VMM_CALLBACK_SET_LAST_ERROR                     VmmCallbackSetLastError;                    // Fixed
    VMM_CALLBACK_VMCALL_HANDLER                     VmmCallbackVmcallHandler;                   // Fixed
    VMM_CALLBACK_NMI_BROADCAST_REQUEST_HANDLER      VmmCallbackNmiBroadcastRequestHandler;      // Fixed
    VMM_CALLBACK_QUERY_TERMINATE_PROTECTED_RESOURCE VmmCallbackQueryTerminateProtectedResource; // Fixed
    VMM_CALLBACK_RESTORE_EPT_STATE                  VmmCallbackRestoreEptState;                 // Fixed
    VMM_CALLBACK_CHECK_UNHANDLED_EPT_VIOLATION      VmmCallbackCheckUnhandledEptViolations;     // Fixed

    //
    // Debugging callbacks
    //
    DEBUGGING_CALLBACK_HANDLE_BREAKPOINT_EXCEPTION       DebuggingCallbackHandleBreakpointException;      // Fixed
    DEBUGGING_CALLBACK_HANDLE_DEBUG_BREAKPOINT_EXCEPTION DebuggingCallbackHandleDebugBreakpointException; // Fixed
    DEBUGGING_CALLBACK_CONDITIONAL_PAGE_FAULT_EXCEPTION  DebuggingCallbackConditionalPageFaultException;  // Fixed

    //
    // Interception callbacks
    //
    INTERCEPTION_CALLBACK_TRIGGER_CR3_CHANGE InterceptionCallbackTriggerCr3ProcessChange; // Fixed

    //
    // Callbacks to be removed
    //
    BREAKPOINT_CHECK_AND_HANDLE_REAPPLYING_BREAKPOINT              BreakpointCheckAndHandleReApplyingBreakpoint;
    UD_CHECK_FOR_COMMAND                                           UdCheckForCommand;
    KD_CHECK_AND_HANDLE_NMI_CALLBACK                               KdCheckAndHandleNmiCallback;
    VMM_CALLBACK_REGISTERED_MTF_HANDLER                            VmmCallbackRegisteredMtfHandler; // Fixed but not good
    INTERCEPTION_CALLBACK_TRIGGER_CLOCK_AND_IPI                    DebuggerCheckProcessOrThreadChange;
    ATTACHING_HANDLE_CR3_EVENTS_FOR_THREAD_INTERCEPTION            AttachingHandleCr3VmexitsForThreadInterception;
    KD_QUERY_DEBUGGER_THREAD_OR_PROCESS_TRACING_DETAILS_BY_CORE_ID KdQueryDebuggerQueryThreadOrProcessTracingDetailsByCoreId;

} VMM_CALLBACKS, *PVMM_CALLBACKS;


//..\bin\debug\SDK\Imports\HyperDbgCtrlImports.h
/**
 * @file HyperDbgCtrlImports.h
 * @author Sina Karvandi (sina@hyperdbg.org)
 * @brief Headers relating exported functions from controller interface
 * @version 0.2
 * @date 2023-02-02
 *
 * @copyright This project is released under the GNU Public License v3.
 *
 */
#pragma once

#ifdef HYPERDBG_HPRDBGCTRL
#    define IMPORT_EXPORT_CTRL __declspec(dllexport)
#else
#    define IMPORT_EXPORT_CTRL __declspec(dllimport)
#endif

//
// Header file of HPRDBGCTRL
// Imports
//
#ifdef __cplusplus
extern "C" {
#endif

//
// Support Detection
//
IMPORT_EXPORT_CTRL bool HyperDbgVmxSupportDetection();
IMPORT_EXPORT_CTRL void HyperDbgReadVendorString(char *);

//
// VMM Module
//
IMPORT_EXPORT_CTRL int HyperDbgLoadVmm();
IMPORT_EXPORT_CTRL int HyperDbgUnloadVmm();
IMPORT_EXPORT_CTRL int HyperDbgInstallVmmDriver();
IMPORT_EXPORT_CTRL int HyperDbgUninstallVmmDriver();
IMPORT_EXPORT_CTRL int HyperDbgStopVmmDriver();

//
// General imports
//
IMPORT_EXPORT_CTRL int HyperDbgInterpreter(char * Command);
IMPORT_EXPORT_CTRL void HyperDbgShowSignature();
IMPORT_EXPORT_CTRL void HyperDbgSetTextMessageCallback(Callback handler);
IMPORT_EXPORT_CTRL int HyperDbgScriptReadFileAndExecuteCommandline(int argc, char * argv[]);
IMPORT_EXPORT_CTRL bool HyperDbgContinuePreviousCommand();
IMPORT_EXPORT_CTRL bool HyperDbgCheckMultilineCommand(char * CurrentCommand, bool Reset);

#ifdef __cplusplus
}
#endif


//..\bin\debug\SDK\Imports\HyperDbgHyperLogImports.h
/**
 * @file HyperDbgHyperLogImports.h
 * @author Sina Karvandi (sina@hyperdbg.org)
 * @brief Headers relating exported functions from hyperlog project
 * @version 0.1
 * @date 2023-01-15
 *
 * @copyright This project is released under the GNU Public License v3.
 *
 */
#pragma once

#ifdef HYPERDBG_HYPER_LOG
#    define IMPORT_EXPORT_HYPERLOG __declspec(dllexport)
#else
#    define IMPORT_EXPORT_HYPERLOG __declspec(dllimport)
#endif

//////////////////////////////////////////////////
//				   Functions					//
//////////////////////////////////////////////////

IMPORT_EXPORT_HYPERLOG BOOLEAN
LogInitialize(MESSAGE_TRACING_CALLBACKS * MsgTracingCallbacks);

IMPORT_EXPORT_HYPERLOG VOID
LogUnInitialize();

IMPORT_EXPORT_HYPERLOG UINT32
LogMarkAllAsRead(BOOLEAN IsVmxRoot);

IMPORT_EXPORT_HYPERLOG BOOLEAN
LogCallbackPrepareAndSendMessageToQueue(UINT32       OperationCode,
                                        BOOLEAN      IsImmediateMessage,
                                        BOOLEAN      ShowCurrentSystemTime,
                                        BOOLEAN      Priority,
                                        const char * Fmt,
                                        ...);

IMPORT_EXPORT_HYPERLOG BOOLEAN
LogCallbackPrepareAndSendMessageToQueueWrapper(UINT32       OperationCode,
                                               BOOLEAN      IsImmediateMessage,
                                               BOOLEAN      ShowCurrentSystemTime,
                                               BOOLEAN      Priority,
                                               const char * Fmt,
                                               va_list      ArgList);

IMPORT_EXPORT_HYPERLOG BOOLEAN
LogCallbackSendBuffer(_In_ UINT32                          OperationCode,
                      _In_reads_bytes_(BufferLength) PVOID Buffer,
                      _In_ UINT32                          BufferLength,
                      _In_ BOOLEAN                         Priority);

IMPORT_EXPORT_HYPERLOG BOOLEAN
LogCallbackCheckIfBufferIsFull(BOOLEAN Priority);

IMPORT_EXPORT_HYPERLOG BOOLEAN
LogCallbackSendMessageToQueue(UINT32 OperationCode, BOOLEAN IsImmediateMessage, CHAR * LogMessage, UINT32 BufferLen, BOOLEAN Priority);

IMPORT_EXPORT_HYPERLOG NTSTATUS
LogRegisterEventBasedNotification(PDEVICE_OBJECT DeviceObject, PIRP Irp);

IMPORT_EXPORT_HYPERLOG NTSTATUS
LogRegisterIrpBasedNotification(PDEVICE_OBJECT DeviceObject, PIRP Irp);


//..\bin\debug\SDK\Imports\HyperDbgHyperLogIntrinsics.h
/**
 * @file HyperDbgHyperLogIntrinsics.h
 * @author Sina Karvandi (sina@hyperdbg.org)
 * @brief Headers relating exported functions from hyperlog project
 * @version 0.1
 * @date 2023-01-15
 *
 * @copyright This project is released under the GNU Public License v3.
 *
 */
#pragma once

//////////////////////////////////////////////////
//					Enums						//
//////////////////////////////////////////////////

/**
 * @brief Types of log messages
 *
 */
typedef enum _LOG_TYPE
{
    LOG_INFO,
    LOG_WARNING,
    LOG_ERROR

} LOG_TYPE;

//////////////////////////////////////////////////
//					Logging						//
//////////////////////////////////////////////////

/**
 * @brief Define log variables
 *
 */
#if UseDbgPrintInsteadOfUsermodeMessageTracking
/* Use DbgPrint */
#    define Logformat, ...)                           \
        DbgPrint("[+] Information (%s:%d) | " format "\n", \
                 __func__,                                 \
                 __LINE__,                                 \
                 __VA_ARGS__)

#    define LogWarning(format, ...)                    \
        DbgPrint("[-] Warning (%s:%d) | " format "\n", \
                 __func__,                             \
                 __LINE__,                             \
                 __VA_ARGS__)

#    define LogError(format, ...)                    \
        DbgPrint("[!] Error (%s:%d) | " format "\n", \
                 __func__,                           \
                 __LINE__,                           \
                 __VA_ARGS__);                       \
        DbgBreakPoint()

/**
 * @brief Log without any prefix
 *
 */
#    define Log(format, ...) \
        DbgPrint(format, __VA_ARGS__)

#else

/**
 * @brief Log, general
 *
 */
#    define LogInfo(format, ...)                                                          \
        LogCallbackPrepareAndSendMessageToQueue(OPERATION_LOG_INFO_MESSAGE,               \
                                                UseImmediateMessaging,                    \
                                                ShowSystemTimeOnDebugMessages,            \
                                                FALSE,                                    \
                                                "[+] Information (%s:%d) | " format "\n", \
                                                __func__,                                 \
                                                __LINE__,                                 \
                                                __VA_ARGS__)

/**
 * @brief Log in the case of priority message
 *
 */
#    define LogInfoPriority(format, ...)                                                  \
        LogCallbackPrepareAndSendMessageToQueue(OPERATION_LOG_INFO_MESSAGE,               \
                                                TRUE,                                     \
                                                ShowSystemTimeOnDebugMessages,            \
                                                TRUE,                                     \
                                                "[+] Information (%s:%d) | " format "\n", \
                                                __func__,                                 \
                                                __LINE__,                                 \
                                                __VA_ARGS__)

/**
 * @brief Log in the case of warning
 *
 */
#    define LogWarning(format, ...)                                                   \
        LogCallbackPrepareAndSendMessageToQueue(OPERATION_LOG_WARNING_MESSAGE,        \
                                                UseImmediateMessaging,                \
                                                ShowSystemTimeOnDebugMessages,        \
                                                TRUE,                                 \
                                                "[-] Warning (%s:%d) | " format "\n", \
                                                __func__,                             \
                                                __LINE__,                             \
                                                __VA_ARGS__)

/**
 * @brief Log in the case of error
 *
 */
#    define LogError(format, ...)                                                   \
        LogCallbackPrepareAndSendMessageToQueue(OPERATION_LOG_ERROR_MESSAGE,        \
                                                UseImmediateMessaging,              \
                                                ShowSystemTimeOnDebugMessages,      \
                                                TRUE,                               \
                                                "[!] Error (%s:%d) | " format "\n", \
                                                __func__,                           \
                                                __LINE__,                           \
                                                __VA_ARGS__);                       \
        if (DebugMode)                                                              \
        DbgBreakPoint()

/**
 * @brief Log without any prefix
 *
 */
#    define Log(format, ...)                                                \
        LogCallbackPrepareAndSendMessageToQueue(OPERATION_LOG_INFO_MESSAGE, \
                                                TRUE,                       \
                                                FALSE,                      \
                                                FALSE,                      \
                                                format,                     \
                                                __VA_ARGS__)

/**
 * @brief Log without any prefix and bypass the stack
 * problem (getting two temporary stacks in preparing phase)
 *
 */
#    define LogSimpleWithTag(tag, isimmdte, buffer, len) \
        LogCallbackSendMessageToQueue(tag,               \
                                      isimmdte,          \
                                      buffer,            \
                                      len,               \
                                      FALSE)

#endif // UseDbgPrintInsteadOfUsermodeMessageTracking

/**
 * @brief Log, initialize boot information and debug information
 *
 */
#define LogDebugInfo(format, ...)                                                     \
    if (DebugMode)                                                                    \
    LogCallbackPrepareAndSendMessageToQueue(OPERATION_LOG_INFO_MESSAGE,               \
                                            UseImmediateMessaging,                    \
                                            ShowSystemTimeOnDebugMessages,            \
                                            FALSE,                                    \
                                            "[+] Information (%s:%d) | " format "\n", \
                                            __func__,                                 \
                                            __LINE__,                                 \
                                            __VA_ARGS__)


//..\bin\debug\SDK\Imports\HyperDbgRevImports.h
/**
 * @file HyperDbgRevImports.h
 * @author Sina Karvandi (sina@hyperdbg.org)
 * @brief Headers relating exported functions from reversing machine interface
 * @version 0.2
 * @date 2023-02-02
 *
 * @copyright This project is released under the GNU Public License v3.
 *
 */
#pragma once

//
// Header file of hpr
// Imports
//
#ifdef __cplusplus
extern "C" {
#endif

//
// Reversing Machine Module
//
__declspec(dllimport) int ReversingMachineStart();
__declspec(dllimport) int ReversingMachineStop();

#ifdef __cplusplus
}
#endif


//..\bin\debug\SDK\Imports\HyperDbgScriptImports.h
/**
 * @file HyperDbgScriptImports.h
 * @author Sina Karvandi (sina@hyperdbg.org)
 * @brief Headers relating exported functions from script engine
 * @version 0.2
 * @date 2023-02-02
 *
 * @copyright This project is released under the GNU Public License v3.
 *
 */
#pragma once

//
// Header file of script-engine
// Imports
//
#ifdef __cplusplus
extern "C" {
#endif

//
// Script engine
//
__declspec(dllimport) PSYMBOL_BUFFER
ScriptEngineParse(char * str);
__declspec(dllimport) void
PrintSymbolBuffer(const PSYMBOL_BUFFER SymbolBuffer);
__declspec(dllimport) void
PrintSymbol(PSYMBOL Symbol);
__declspec(dllimport) void
RemoveSymbolBuffer(PSYMBOL_BUFFER SymbolBuffer);
__declspec(dllimport) BOOLEAN
FuncGetNumberOfOperands(UINT64 FuncType, UINT32 * NumberOfGetOperands, UINT32 * NumberOfSetOperands);
__declspec(dllimport) BOOLEAN
ScriptEngineSetHwdbgInstanceInfo(HWDBG_INSTANCE_INFORMATION * InstancInfo);

;

//
// pdb parser
//
__declspec(dllimport) VOID
ScriptEngineSetTextMessageCallback(PVOID Handler);
__declspec(dllimport) VOID
ScriptEngineSymbolAbortLoading();
__declspec(dllimport) UINT64
ScriptEngineConvertNameToAddress(const char * FunctionOrVariableName, PBOOLEAN WasFound);
__declspec(dllimport) UINT32
ScriptEngineLoadFileSymbol(UINT64 BaseAddress, const char * PdbFileName, const char * CustomModuleName);
__declspec(dllimport) UINT32
ScriptEngineUnloadAllSymbols();
__declspec(dllimport) UINT32
ScriptEngineUnloadModuleSymbol(char * ModuleName);
__declspec(dllimport) UINT32
ScriptEngineSearchSymbolForMask(const char * SearchMask);
__declspec(dllimport) BOOLEAN
ScriptEngineGetFieldOffset(CHAR * TypeName, CHAR * FieldName, UINT32 * FieldOffset);
__declspec(dllimport) BOOLEAN
ScriptEngineGetDataTypeSize(CHAR * TypeName, UINT64 * TypeSize);
__declspec(dllimport) BOOLEAN
ScriptEngineCreateSymbolTableForDisassembler(void * CallbackFunction);
__declspec(dllimport) BOOLEAN
ScriptEngineConvertFileToPdbPath(const char * LocalFilePath, char * ResultPath);
__declspec(dllimport) BOOLEAN
ScriptEngineConvertFileToPdbFileAndGuidAndAgeDetails(const char * LocalFilePath, char * PdbFilePath, char * GuidAndAgeDetails, BOOLEAN Is32BitModule);
__declspec(dllimport) BOOLEAN
ScriptEngineSymbolInitLoad(PVOID BufferToStoreDetails, UINT32 StoredLength, BOOLEAN DownloadIfAvailable, const char * SymbolPath, BOOLEAN IsSilentLoad);
__declspec(dllimport) BOOLEAN
ScriptEngineShowDataBasedOnSymbolTypes(const char * TypeName, UINT64 Address, BOOLEAN IsStruct, PVOID BufferAddress, const char * AdditionalParameters);

#ifdef __cplusplus
}
#endif


//..\bin\debug\SDK\Imports\HyperDbgSymImports.h
/**
 * @file HyperDbgSymImports.h
 * @author Sina Karvandi (sina@hyperdbg.org)
 * @brief Headers relating exported functions from symbol parser
 * @version 0.2
 * @date 2023-02-02
 *
 * @copyright This project is released under the GNU Public License v3.
 *
 */
#pragma once

//
// Header file of symbol-parser
// Imports
//
#ifdef __cplusplus
extern "C" {
#endif

__declspec(dllimport) VOID
    SymSetTextMessageCallback(PVOID Handler);
__declspec(dllimport) VOID
    SymbolAbortLoading();
__declspec(dllimport) UINT64
    SymConvertNameToAddress(const char * FunctionOrVariableName, PBOOLEAN WasFound);
__declspec(dllimport) UINT32
    SymLoadFileSymbol(UINT64 BaseAddress, const char * PdbFileName, const char * CustomModuleName);
__declspec(dllimport) UINT32
    SymUnloadAllSymbols();
__declspec(dllimport) UINT32
    SymUnloadModuleSymbol(char * ModuleName);
__declspec(dllimport) UINT32
    SymSearchSymbolForMask(const char * SearchMask);
__declspec(dllimport) BOOLEAN
    SymGetFieldOffset(CHAR * TypeName, CHAR * FieldName, UINT32 * FieldOffset);
__declspec(dllimport) BOOLEAN
    SymGetDataTypeSize(CHAR * TypeName, UINT64 * TypeSize);
__declspec(dllimport) BOOLEAN
    SymCreateSymbolTableForDisassembler(void * CallbackFunction);
__declspec(dllimport) BOOLEAN
    SymConvertFileToPdbPath(const char * LocalFilePath, char * ResultPath);
__declspec(dllimport) BOOLEAN
    SymConvertFileToPdbFileAndGuidAndAgeDetails(const char * LocalFilePath,
                                                char *       PdbFilePath,
                                                char *       GuidAndAgeDetails,
                                                BOOLEAN      Is32BitModule);
__declspec(dllimport) BOOLEAN
    SymbolInitLoad(PVOID        BufferToStoreDetails,
                   UINT32       StoredLength,
                   BOOLEAN      DownloadIfAvailable,
                   const char * SymbolPath,
                   BOOLEAN      IsSilentLoad);
__declspec(dllimport) BOOLEAN
    SymShowDataBasedOnSymbolTypes(const char * TypeName,
                                  UINT64       Address,
                                  BOOLEAN      IsStruct,
                                  PVOID        BufferAddress,
                                  const char * AdditionalParameters);
__declspec(dllimport) BOOLEAN
    SymQuerySizeof(_In_ const char * StructNameOrTypeName, _Out_ UINT32 * SizeOfField);
__declspec(dllimport) BOOLEAN
    SymCastingQueryForFiledsAndTypes(_In_ const char * StructName,
                                     _In_ const char * FiledOfStructName,
                                     _Out_ PBOOLEAN    IsStructNamePointerOrNot,
                                     _Out_ PBOOLEAN    IsFiledOfStructNamePointerOrNot,
                                     _Out_ char **     NewStructOrTypeName,
                                     _Out_ UINT32 *    OffsetOfFieldFromTop,
                                     _Out_ UINT32 *    SizeOfField);

#ifdef __cplusplus
}
#endif


//..\bin\debug\SDK\Imports\HyperDbgVmmImports.h
/**
 * @file HyperDbgVmmImports.h
 * @author Sina Karvandi (sina@hyperdbg.org)
 * @brief Headers relating exported functions from hypervisor
 * @version 0.1
 * @date 2022-12-09
 *
 * @copyright This project is released under the GNU Public License v3.
 *
 */
#pragma once

#ifdef HYPERDBG_VMM
#    define IMPORT_EXPORT_VMM __declspec(dllexport)
#else
#    define IMPORT_EXPORT_VMM __declspec(dllimport)
#endif

//////////////////////////////////////////////////
//                 VM Functions 	    		//
//////////////////////////////////////////////////

IMPORT_EXPORT_VMM NTSTATUS
VmFuncVmxVmcall(unsigned long long VmcallNumber,
                unsigned long long OptionalParam1,
                unsigned long long OptionalParam2,
                unsigned long long OptionalParam3);

IMPORT_EXPORT_VMM VOID
VmFuncPerformRipIncrement(UINT32 CoreId);

IMPORT_EXPORT_VMM VOID
VmFuncSuppressRipIncrement(UINT32 CoreId);

IMPORT_EXPORT_VMM VOID
VmFuncChangeMtfUnsettingState(UINT32 CoreId, BOOLEAN Set);

IMPORT_EXPORT_VMM VOID
VmFuncChangeIgnoreOneMtfState(UINT32 CoreId, BOOLEAN Set);

IMPORT_EXPORT_VMM VOID
VmFuncSetMonitorTrapFlag(BOOLEAN Set);

IMPORT_EXPORT_VMM VOID
VmFuncSetRflagTrapFlag(BOOLEAN Set);

IMPORT_EXPORT_VMM VOID
VmFuncRegisterMtfBreak(UINT32 CoreId);

IMPORT_EXPORT_VMM VOID
VmFuncUnRegisterMtfBreak(UINT32 CoreId);

IMPORT_EXPORT_VMM VOID
VmFuncSetLoadDebugControls(BOOLEAN Set);

IMPORT_EXPORT_VMM VOID
VmFuncSetSaveDebugControls(BOOLEAN Set);

IMPORT_EXPORT_VMM VOID
VmFuncSetPmcVmexit(BOOLEAN Set);

IMPORT_EXPORT_VMM VOID
VmFuncSetMovControlRegsExiting(BOOLEAN Set, UINT64 ControlRegister, UINT64 MaskRegister);

IMPORT_EXPORT_VMM VOID
VmFuncSetMovToCr3Vmexit(UINT32 CoreId, BOOLEAN Set);

IMPORT_EXPORT_VMM VOID
VmFuncWriteExceptionBitmap(UINT32 BitmapMask);

IMPORT_EXPORT_VMM VOID
VmFuncSetInterruptWindowExiting(BOOLEAN Set);

IMPORT_EXPORT_VMM VOID
VmFuncSetNmiWindowExiting(BOOLEAN Set);

IMPORT_EXPORT_VMM VOID
VmFuncSetNmiExiting(BOOLEAN Set);

IMPORT_EXPORT_VMM VOID
VmFuncSetExceptionBitmap(UINT32 CoreId, UINT32 IdtIndex);

IMPORT_EXPORT_VMM VOID
VmFuncUnsetExceptionBitmap(UINT32 CoreId, UINT32 IdtIndex);

IMPORT_EXPORT_VMM VOID
VmFuncSetExternalInterruptExiting(UINT32 CoreId, BOOLEAN Set);

IMPORT_EXPORT_VMM VOID
VmFuncSetRdtscExiting(UINT32 CoreId, BOOLEAN Set);

IMPORT_EXPORT_VMM VOID
VmFuncSetMovDebugRegsExiting(UINT32 CoreId, BOOLEAN Set);

IMPORT_EXPORT_VMM VOID
VmFuncInjectPendingExternalInterrupts(UINT32 CoreId);

IMPORT_EXPORT_VMM VOID
VmFuncSetRflags(UINT64 Rflags);

IMPORT_EXPORT_VMM VOID
VmFuncSetRip(UINT64 Rip);

IMPORT_EXPORT_VMM VOID
VmFuncSetTriggerEventForVmcalls(BOOLEAN Set);

IMPORT_EXPORT_VMM VOID
VmFuncSetTriggerEventForCpuids(BOOLEAN Set);

IMPORT_EXPORT_VMM VOID
VmFuncSetInterruptibilityState(UINT64 InterruptibilityState);

IMPORT_EXPORT_VMM VOID
VmFuncCheckAndEnableExternalInterrupts(UINT32 CoreId);

IMPORT_EXPORT_VMM VOID
VmFuncDisableExternalInterruptsAndInterruptWindow(UINT32 CoreId);

IMPORT_EXPORT_VMM VOID
VmFuncEventInjectPageFaultWithCr2(UINT32 CoreId, UINT64 Address, UINT32 PageFaultCode);

IMPORT_EXPORT_VMM VOID
VmFuncEventInjectPageFaultRangeAddress(UINT32 CoreId,
                                       UINT64 AddressFrom,
                                       UINT64 AddressTo,
                                       UINT32 PageFaultCode);

IMPORT_EXPORT_VMM VOID
VmFuncEventInjectInterruption(UINT32  InterruptionType,
                              UINT32  Vector,
                              BOOLEAN DeliverErrorCode,
                              UINT32  ErrorCode);

IMPORT_EXPORT_VMM VOID
VmFuncVmxBroadcastInitialize();

IMPORT_EXPORT_VMM VOID
VmFuncVmxBroadcastUninitialize();

IMPORT_EXPORT_VMM VOID
VmFuncEventInjectBreakpoint();

IMPORT_EXPORT_VMM VOID
VmFuncInvalidateEptSingleContext(UINT32 CoreId);

IMPORT_EXPORT_VMM VOID
VmFuncInvalidateEptAllContexts();

IMPORT_EXPORT_VMM VOID
VmFuncUninitVmm();

IMPORT_EXPORT_VMM VOID
VmFuncEnableMtfAndChangeExternalInterruptState(UINT32 CoreId);

IMPORT_EXPORT_VMM VOID
VmFuncEnableAndCheckForPreviousExternalInterrupts(UINT32 CoreId);

IMPORT_EXPORT_VMM UINT16
VmFuncGetCsSelector();

IMPORT_EXPORT_VMM UINT32
VmFuncReadExceptionBitmap();

IMPORT_EXPORT_VMM UINT64
VmFuncGetLastVmexitRip(UINT32 CoreId);

IMPORT_EXPORT_VMM UINT64
VmFuncGetRflags();

IMPORT_EXPORT_VMM UINT64
VmFuncGetRip();

IMPORT_EXPORT_VMM UINT64
VmFuncGetInterruptibilityState();

IMPORT_EXPORT_VMM UINT64
VmFuncClearSteppingBits(UINT64 Interruptibility);

IMPORT_EXPORT_VMM BOOLEAN
VmFuncInitVmm(VMM_CALLBACKS * VmmCallbacks);

IMPORT_EXPORT_VMM UINT32
VmFuncVmxCompatibleStrlen(const CHAR * s);

IMPORT_EXPORT_VMM UINT32
VmFuncVmxCompatibleWcslen(const wchar_t * s);

IMPORT_EXPORT_VMM BOOLEAN
VmFuncNmiBroadcastRequest(UINT32 CoreId);

IMPORT_EXPORT_VMM BOOLEAN
VmFuncNmiBroadcastInvalidateEptSingleContext(UINT32 CoreId);

IMPORT_EXPORT_VMM BOOLEAN
VmFuncNmiBroadcastInvalidateEptAllContexts(UINT32 CoreId);

IMPORT_EXPORT_VMM BOOLEAN
VmFuncVmxGetCurrentExecutionMode();

IMPORT_EXPORT_VMM BOOLEAN
VmFuncQueryModeExecTrap();

IMPORT_EXPORT_VMM INT32
VmFuncVmxCompatibleStrcmp(const CHAR * Address1, const CHAR * Address2);

IMPORT_EXPORT_VMM INT32
VmFuncVmxCompatibleStrncmp(const CHAR * Address1, const CHAR * Address2, SIZE_T Num);

IMPORT_EXPORT_VMM INT32
VmFuncVmxCompatibleWcscmp(const wchar_t * Address1, const wchar_t * Address2);

IMPORT_EXPORT_VMM INT32
VmFuncVmxCompatibleWcsncmp(const wchar_t * Address1, const wchar_t * Address2, SIZE_T Num);

IMPORT_EXPORT_VMM INT32
VmFuncVmxCompatibleMemcmp(const CHAR * Address1, const CHAR * Address2, size_t Count);

//////////////////////////////////////////////////
//            Configuration Functions 	   		//
//////////////////////////////////////////////////

IMPORT_EXPORT_VMM VOID
ConfigureEnableMovToCr3ExitingOnAllProcessors();

IMPORT_EXPORT_VMM VOID
ConfigureDisableMovToCr3ExitingOnAllProcessors();

IMPORT_EXPORT_VMM VOID
ConfigureEnableEferSyscallEventsOnAllProcessors();

IMPORT_EXPORT_VMM VOID
ConfigureDisableEferSyscallEventsOnAllProcessors();

IMPORT_EXPORT_VMM VOID
ConfigureSetExternalInterruptExitingOnSingleCore(UINT32 TargetCoreId);

IMPORT_EXPORT_VMM VOID
ConfigureEnableRdtscExitingOnSingleCore(UINT32 TargetCoreId);

IMPORT_EXPORT_VMM VOID
ConfigureEnableRdpmcExitingOnSingleCore(UINT32 TargetCoreId);

IMPORT_EXPORT_VMM VOID
ConfigureEnableMovToDebugRegistersExitingOnSingleCore(UINT32 TargetCoreId);

IMPORT_EXPORT_VMM VOID
ConfigureSetExceptionBitmapOnSingleCore(UINT32 TargetCoreId, UINT32 BitMask);

IMPORT_EXPORT_VMM VOID
ConfigureEnableMovToControlRegisterExitingOnSingleCore(UINT32 TargetCoreId, DEBUGGER_EVENT_OPTIONS * BroadcastingOption);

IMPORT_EXPORT_VMM VOID
ConfigureChangeMsrBitmapWriteOnSingleCore(UINT32 TargetCoreId, UINT64 MsrMask);

IMPORT_EXPORT_VMM VOID
ConfigureChangeMsrBitmapReadOnSingleCore(UINT32 TargetCoreId, UINT64 MsrMask);

IMPORT_EXPORT_VMM VOID
ConfigureChangeIoBitmapOnSingleCore(UINT32 TargetCoreId, UINT64 Port);

IMPORT_EXPORT_VMM VOID
ConfigureEnableEferSyscallHookOnSingleCore(UINT32 TargetCoreId);

IMPORT_EXPORT_VMM VOID
ConfigureSetEferSyscallOrSysretHookType(DEBUGGER_EVENT_SYSCALL_SYSRET_TYPE SyscallHookType);

IMPORT_EXPORT_VMM VOID
ConfigureDirtyLoggingInitializeOnAllProcessors();

IMPORT_EXPORT_VMM VOID
ConfigureDirtyLoggingUninitializeOnAllProcessors();

IMPORT_EXPORT_VMM VOID
ConfigureModeBasedExecHookUninitializeOnAllProcessors();

IMPORT_EXPORT_VMM VOID
ConfigureUninitializeExecTrapOnAllProcessors();

IMPORT_EXPORT_VMM BOOLEAN
ConfigureInitializeExecTrapOnAllProcessors();

IMPORT_EXPORT_VMM BOOLEAN
ConfigureEptHook(PVOID TargetAddress, UINT32 ProcessId);

IMPORT_EXPORT_VMM BOOLEAN
ConfigureEptHookFromVmxRoot(PVOID TargetAddress);

IMPORT_EXPORT_VMM BOOLEAN
ConfigureEptHook2(UINT32 CoreId,
                  PVOID  TargetAddress,
                  PVOID  HookFunction,
                  UINT32 ProcessId);

IMPORT_EXPORT_VMM BOOLEAN
ConfigureEptHook2FromVmxRoot(UINT32 CoreId,
                             PVOID  TargetAddress,
                             PVOID  HookFunction);

IMPORT_EXPORT_VMM BOOLEAN
ConfigureEptHookMonitor(UINT32                                         CoreId,
                        EPT_HOOKS_ADDRESS_DETAILS_FOR_MEMORY_MONITOR * HookingDetails,
                        UINT32                                         ProcessId);

IMPORT_EXPORT_VMM BOOLEAN
ConfigureEptHookMonitorFromVmxRoot(UINT32                                         CoreId,
                                   EPT_HOOKS_ADDRESS_DETAILS_FOR_MEMORY_MONITOR * MemoryAddressDetails);

IMPORT_EXPORT_VMM BOOLEAN
ConfigureEptHookModifyInstructionFetchState(UINT32  CoreId,
                                            PVOID   PhysicalAddress,
                                            BOOLEAN IsUnset);

IMPORT_EXPORT_VMM BOOLEAN
ConfigureEptHookModifyPageReadState(UINT32  CoreId,
                                    PVOID   PhysicalAddress,
                                    BOOLEAN IsUnset);

IMPORT_EXPORT_VMM BOOLEAN
ConfigureEptHookModifyPageWriteState(UINT32  CoreId,
                                     PVOID   PhysicalAddress,
                                     BOOLEAN IsUnset);

IMPORT_EXPORT_VMM BOOLEAN
ConfigureEptHookUnHookSingleAddress(UINT64 VirtualAddress,
                                    UINT64 PhysAddress,
                                    UINT32 ProcessId);

IMPORT_EXPORT_VMM BOOLEAN
ConfigureEptHookUnHookSingleAddressFromVmxRoot(UINT64                              VirtualAddress,
                                               UINT64                              PhysAddress,
                                               EPT_SINGLE_HOOK_UNHOOKING_DETAILS * TargetUnhookingDetails);

IMPORT_EXPORT_VMM VOID
ConfigureEptHookAllocateExtraHookingPagesForMemoryMonitorsAndExecEptHooks(UINT32 Count);

IMPORT_EXPORT_VMM VOID
ConfigureEptHookReservePreallocatedPoolsForEptHooks(UINT32 Count);

IMPORT_EXPORT_VMM BOOLEAN
ConfigureExecTrapAddProcessToWatchingList(UINT32 ProcessId);

IMPORT_EXPORT_VMM BOOLEAN
ConfigureExecTrapRemoveProcessFromWatchingList(UINT32 ProcessId);

//////////////////////////////////////////////////
//           Direct VMCALL Functions 	   		//
//////////////////////////////////////////////////

IMPORT_EXPORT_VMM NTSTATUS
DirectVmcallTest(UINT32 CoreId, DIRECT_VMCALL_PARAMETERS * DirectVmcallOptions);

IMPORT_EXPORT_VMM NTSTATUS
DirectVmcallPerformVmcall(UINT32 CoreId, UINT64 VmcallNumber, DIRECT_VMCALL_PARAMETERS * DirectVmcallOptions);

IMPORT_EXPORT_VMM NTSTATUS
DirectVmcallChangeMsrBitmapRead(UINT32 CoreId, DIRECT_VMCALL_PARAMETERS * DirectVmcallOptions);

IMPORT_EXPORT_VMM NTSTATUS
DirectVmcallChangeMsrBitmapWrite(UINT32 CoreId, DIRECT_VMCALL_PARAMETERS * DirectVmcallOptions);

IMPORT_EXPORT_VMM NTSTATUS
DirectVmcallChangeIoBitmap(UINT32 CoreId, DIRECT_VMCALL_PARAMETERS * DirectVmcallOptions);

IMPORT_EXPORT_VMM NTSTATUS
DirectVmcallEnableRdpmcExiting(UINT32 CoreId, DIRECT_VMCALL_PARAMETERS * DirectVmcallOptions);

IMPORT_EXPORT_VMM NTSTATUS
DirectVmcallEnableRdtscpExiting(UINT32 CoreId, DIRECT_VMCALL_PARAMETERS * DirectVmcallOptions);

IMPORT_EXPORT_VMM NTSTATUS
DirectVmcallEnableMov2DebugRegsExiting(UINT32 CoreId, DIRECT_VMCALL_PARAMETERS * DirectVmcallOptions);

IMPORT_EXPORT_VMM NTSTATUS
DirectVmcallSetExceptionBitmap(UINT32 CoreId, DIRECT_VMCALL_PARAMETERS * DirectVmcallOptions);

IMPORT_EXPORT_VMM NTSTATUS
DirectVmcallEnableExternalInterruptExiting(UINT32 CoreId, DIRECT_VMCALL_PARAMETERS * DirectVmcallOptions);

IMPORT_EXPORT_VMM NTSTATUS
DirectVmcallEnableMovToCrExiting(UINT32 CoreId, DIRECT_VMCALL_PARAMETERS * DirectVmcallOptions);

IMPORT_EXPORT_VMM NTSTATUS
DirectVmcallEnableEferSyscall(UINT32 CoreId, DIRECT_VMCALL_PARAMETERS * DirectVmcallOptions);

IMPORT_EXPORT_VMM NTSTATUS
DirectVmcallSetHiddenBreakpointHook(UINT32 CoreId, DIRECT_VMCALL_PARAMETERS * DirectVmcallOptions);

IMPORT_EXPORT_VMM NTSTATUS
DirectVmcallInvalidateEptAllContexts(UINT32 CoreId, DIRECT_VMCALL_PARAMETERS * DirectVmcallOptions);

IMPORT_EXPORT_VMM NTSTATUS
DirectVmcallInvalidateSingleContext(UINT32 CoreId, DIRECT_VMCALL_PARAMETERS * DirectVmcallOptions);

IMPORT_EXPORT_VMM NTSTATUS
DirectVmcallUnsetExceptionBitmap(UINT32 CoreId, DIRECT_VMCALL_PARAMETERS * DirectVmcallOptions);

IMPORT_EXPORT_VMM NTSTATUS
DirectVmcallUnhookSinglePage(UINT32 CoreId, DIRECT_VMCALL_PARAMETERS * DirectVmcallOptions);

IMPORT_EXPORT_VMM NTSTATUS
DirectVmcallSetDisableExternalInterruptExitingOnlyOnClearingInterruptEvents(UINT32 CoreId, DIRECT_VMCALL_PARAMETERS * DirectVmcallOptions);

IMPORT_EXPORT_VMM NTSTATUS
DirectVmcallResetMsrBitmapRead(UINT32 CoreId, DIRECT_VMCALL_PARAMETERS * DirectVmcallOptions);

IMPORT_EXPORT_VMM NTSTATUS
DirectVmcallResetMsrBitmapWrite(UINT32 CoreId, DIRECT_VMCALL_PARAMETERS * DirectVmcallOptions);

IMPORT_EXPORT_VMM NTSTATUS
DirectVmcallResetExceptionBitmapOnlyOnClearingExceptionEvents(UINT32 CoreId, DIRECT_VMCALL_PARAMETERS * DirectVmcallOptions);

IMPORT_EXPORT_VMM NTSTATUS
DirectVmcallResetIoBitmap(UINT32 CoreId, DIRECT_VMCALL_PARAMETERS * DirectVmcallOptions);

IMPORT_EXPORT_VMM NTSTATUS
DirectVmcallDisableRdtscExitingForClearingTscEvents(UINT32 CoreId, DIRECT_VMCALL_PARAMETERS * DirectVmcallOptions);

IMPORT_EXPORT_VMM NTSTATUS
DirectVmcallDisableRdpmcExiting(UINT32 CoreId, DIRECT_VMCALL_PARAMETERS * DirectVmcallOptions);

IMPORT_EXPORT_VMM NTSTATUS
DirectVmcallDisableEferSyscallEvents(UINT32 CoreId, DIRECT_VMCALL_PARAMETERS * DirectVmcallOptions);

IMPORT_EXPORT_VMM NTSTATUS
DirectVmcallDisableMov2DrExitingForClearingDrEvents(UINT32 CoreId, DIRECT_VMCALL_PARAMETERS * DirectVmcallOptions);

IMPORT_EXPORT_VMM NTSTATUS
DirectVmcallDisableMov2CrExitingForClearingCrEvents(UINT32 CoreId, DIRECT_VMCALL_PARAMETERS * DirectVmcallOptions);

//////////////////////////////////////////////////
//                 Disassembler 	    		//
//////////////////////////////////////////////////

IMPORT_EXPORT_VMM BOOLEAN
DisassemblerShowInstructionsInVmxNonRootMode(PVOID Address, UINT32 Length, BOOLEAN Is32Bit);

IMPORT_EXPORT_VMM BOOLEAN
DisassemblerShowOneInstructionInVmxNonRootMode(PVOID Address, UINT64 ActualRip, BOOLEAN Is32Bit);

IMPORT_EXPORT_VMM UINT32
DisassemblerShowOneInstructionInVmxRootMode(PVOID Address, BOOLEAN Is32Bit);

//////////////////////////////////////////////////
//                General Functions 	   		//
//////////////////////////////////////////////////

// ----------------------------------------------------------------------------
// Exported Interfaces For Virtual Addresses
//

IMPORT_EXPORT_VMM UINT64
VirtualAddressToPhysicalAddress(_In_ PVOID VirtualAddress);

IMPORT_EXPORT_VMM UINT64
VirtualAddressToPhysicalAddressByProcessId(_In_ PVOID  VirtualAddress,
                                           _In_ UINT32 ProcessId);

IMPORT_EXPORT_VMM UINT64
VirtualAddressToPhysicalAddressByProcessCr3(_In_ PVOID    VirtualAddress,
                                            _In_ CR3_TYPE TargetCr3);

IMPORT_EXPORT_VMM UINT64
VirtualAddressToPhysicalAddressOnTargetProcess(_In_ PVOID VirtualAddress);

// ----------------------------------------------------------------------------
// Exported Interfaces For Physical Addresses
//
IMPORT_EXPORT_VMM UINT64
PhysicalAddressToVirtualAddress(_In_ UINT64 PhysicalAddress);

IMPORT_EXPORT_VMM UINT64
PhysicalAddressToVirtualAddressByProcessId(_In_ PVOID PhysicalAddress, _In_ UINT32 ProcessId);

IMPORT_EXPORT_VMM UINT64
PhysicalAddressToVirtualAddressByCr3(_In_ PVOID PhysicalAddress, _In_ CR3_TYPE TargetCr3);

IMPORT_EXPORT_VMM UINT64
PhysicalAddressToVirtualAddressOnTargetProcess(_In_ PVOID PhysicalAddress);

// ----------------------------------------------------------------------------
// Exported Interfaces For Layout Switching Functions
//
IMPORT_EXPORT_VMM CR3_TYPE
SwitchToProcessMemoryLayout(_In_ UINT32 ProcessId);

IMPORT_EXPORT_VMM CR3_TYPE
SwitchToCurrentProcessMemoryLayout();

IMPORT_EXPORT_VMM CR3_TYPE
SwitchToProcessMemoryLayoutByCr3(_In_ CR3_TYPE TargetCr3);

IMPORT_EXPORT_VMM VOID
SwitchToPreviousProcess(_In_ CR3_TYPE PreviousProcess);

// ----------------------------------------------------------------------------
// Exported Interfaces For Check Validity of Addresses
//
IMPORT_EXPORT_VMM BOOLEAN
CheckAddressValidityUsingTsx(CHAR * Address);

IMPORT_EXPORT_VMM BOOLEAN
CheckAccessValidityAndSafety(UINT64 TargetAddress, UINT32 Size);

IMPORT_EXPORT_VMM BOOLEAN
CheckAddressPhysical(UINT64 PAddr);

IMPORT_EXPORT_VMM UINT32
CheckAddressMaximumInstructionLength(PVOID Address);

// ----------------------------------------------------------------------------
// Exported Interfaces For Layout Functions
//
IMPORT_EXPORT_VMM CR3_TYPE
LayoutGetCurrentProcessCr3();

IMPORT_EXPORT_VMM CR3_TYPE
LayoutGetExactGuestProcessCr3();

//////////////////////////////////////////////////
//         Memory Management Functions 	   		//
//////////////////////////////////////////////////

// ----------------------------------------------------------------------------
// PTE-related Functions
//

IMPORT_EXPORT_VMM PVOID
MemoryMapperGetPteVa(_In_ PVOID        Va,
                     _In_ PAGING_LEVEL Level);

IMPORT_EXPORT_VMM PVOID
MemoryMapperGetPteVaByCr3(_In_ PVOID        Va,
                          _In_ PAGING_LEVEL Level,
                          _In_ CR3_TYPE     TargetCr3);

IMPORT_EXPORT_VMM PVOID
MemoryMapperGetPteVaWithoutSwitchingByCr3(_In_ PVOID        Va,
                                          _In_ PAGING_LEVEL Level,
                                          _In_ CR3_TYPE     TargetCr3);

IMPORT_EXPORT_VMM PVOID
MemoryMapperGetPteVaOnTargetProcess(_In_ PVOID        Va,
                                    _In_ PAGING_LEVEL Level);

IMPORT_EXPORT_VMM PVOID
MemoryMapperSetExecuteDisableToPteOnTargetProcess(_In_ PVOID   Va,
                                                  _In_ BOOLEAN Set);

IMPORT_EXPORT_VMM BOOLEAN
MemoryMapperCheckPteIsPresentOnTargetProcess(PVOID        Va,
                                             PAGING_LEVEL Level);

// ----------------------------------------------------------------------------
// Reading Memory Functions
//
IMPORT_EXPORT_VMM BOOLEAN
MemoryMapperReadMemorySafe(_In_ UINT64   VaAddressToRead,
                           _Inout_ PVOID BufferToSaveMemory,
                           _In_ SIZE_T   SizeToRead);

IMPORT_EXPORT_VMM BOOLEAN
MemoryMapperReadMemorySafeByPhysicalAddress(_In_ UINT64    PaAddressToRead,
                                            _Inout_ UINT64 BufferToSaveMemory,
                                            _In_ SIZE_T    SizeToRead);

IMPORT_EXPORT_VMM BOOLEAN
MemoryMapperReadMemorySafeOnTargetProcess(_In_ UINT64   VaAddressToRead,
                                          _Inout_ PVOID BufferToSaveMemory,
                                          _In_ SIZE_T   SizeToRead);

// ----------------------------------------------------------------------------
// Disassembler Functions
//
IMPORT_EXPORT_VMM UINT32
DisassemblerLengthDisassembleEngine(PVOID Address, BOOLEAN Is32Bit);

IMPORT_EXPORT_VMM UINT32
DisassemblerLengthDisassembleEngineInVmxRootOnTargetProcess(PVOID Address, BOOLEAN Is32Bit);

// ----------------------------------------------------------------------------
// Writing Memory Functions
//
IMPORT_EXPORT_VMM BOOLEAN
MemoryMapperWriteMemorySafe(_Inout_ UINT64 Destination,
                            _In_ PVOID     Source,
                            _In_ SIZE_T    SizeToWrite,
                            _In_ CR3_TYPE  TargetProcessCr3);

IMPORT_EXPORT_VMM BOOLEAN
MemoryMapperWriteMemorySafeOnTargetProcess(_Inout_ UINT64 Destination,
                                           _In_ PVOID     Source,
                                           _In_ SIZE_T    Size);

IMPORT_EXPORT_VMM BOOLEAN
MemoryMapperWriteMemorySafeByPhysicalAddress(_Inout_ UINT64 DestinationPa,
                                             _In_ UINT64    Source,
                                             _In_ SIZE_T    SizeToWrite);

IMPORT_EXPORT_VMM BOOLEAN
MemoryMapperWriteMemoryUnsafe(_Inout_ UINT64 Destination,
                              _In_ PVOID     Source,
                              _In_ SIZE_T    SizeToWrite,
                              _In_ UINT32    TargetProcessId);

// ----------------------------------------------------------------------------
// Reserving Memory Functions
//
IMPORT_EXPORT_VMM UINT64
MemoryMapperReserveUsermodeAddressOnTargetProcess(_In_ UINT32  ProcessId,
                                                  _In_ BOOLEAN Allocate);

IMPORT_EXPORT_VMM BOOLEAN
MemoryMapperFreeMemoryOnTargetProcess(_In_ UINT32   ProcessId,
                                      _Inout_ PVOID BaseAddress);

// ----------------------------------------------------------------------------
// Miscellaneous Memory Functions
//
IMPORT_EXPORT_VMM BOOLEAN
MemoryMapperSetSupervisorBitWithoutSwitchingByCr3(_In_ PVOID        Va,
                                                  _In_ BOOLEAN      Set,
                                                  _In_ PAGING_LEVEL Level,
                                                  _In_ CR3_TYPE     TargetCr3);

IMPORT_EXPORT_VMM BOOLEAN
MemoryMapperCheckIfPageIsNxBitSetOnTargetProcess(_In_ PVOID Va);

IMPORT_EXPORT_VMM BOOLEAN
MemoryMapperCheckIfPdeIsLargePageOnTargetProcess(_In_ PVOID Va);

//////////////////////////////////////////////////
//				Memory Manager		    		//
//////////////////////////////////////////////////

IMPORT_EXPORT_VMM BOOLEAN
MemoryManagerReadProcessMemoryNormal(HANDLE PID, PVOID Address, DEBUGGER_READ_MEMORY_TYPE MemType, PVOID UserBuffer, SIZE_T Size, PSIZE_T ReturnSize);

//////////////////////////////////////////////////
//                 Pool Manager     	   		//
//////////////////////////////////////////////////

IMPORT_EXPORT_VMM BOOLEAN
PoolManagerCheckAndPerformAllocationAndDeallocation();

IMPORT_EXPORT_VMM BOOLEAN
PoolManagerRequestAllocation(SIZE_T Size, UINT32 Count, POOL_ALLOCATION_INTENTION Intention);

IMPORT_EXPORT_VMM UINT64
PoolManagerRequestPool(POOL_ALLOCATION_INTENTION Intention, BOOLEAN RequestNewPool, UINT32 Size);

IMPORT_EXPORT_VMM BOOLEAN
PoolManagerFreePool(UINT64 AddressToFree);

IMPORT_EXPORT_VMM VOID
PoolManagerShowPreAllocatedPools();

//////////////////////////////////////////////////
//          VMX Registers Modification  		//
//////////////////////////////////////////////////

IMPORT_EXPORT_VMM VOID
SetGuestCsSel(PVMX_SEGMENT_SELECTOR Cs);

IMPORT_EXPORT_VMM VOID
SetGuestCs(PVMX_SEGMENT_SELECTOR Cs);

IMPORT_EXPORT_VMM VMX_SEGMENT_SELECTOR
GetGuestCs();

IMPORT_EXPORT_VMM VOID
SetGuestSsSel(PVMX_SEGMENT_SELECTOR Ss);

IMPORT_EXPORT_VMM VOID
SetGuestSs(PVMX_SEGMENT_SELECTOR Ss);

IMPORT_EXPORT_VMM VMX_SEGMENT_SELECTOR
GetGuestSs();

IMPORT_EXPORT_VMM VOID
SetGuestDsSel(PVMX_SEGMENT_SELECTOR Ds);

IMPORT_EXPORT_VMM VOID
SetGuestDs(PVMX_SEGMENT_SELECTOR Ds);

IMPORT_EXPORT_VMM VMX_SEGMENT_SELECTOR
GetGuestDs();

IMPORT_EXPORT_VMM VOID
SetGuestFsSel(PVMX_SEGMENT_SELECTOR Fs);

IMPORT_EXPORT_VMM VOID
SetGuestFs(PVMX_SEGMENT_SELECTOR Fs);

IMPORT_EXPORT_VMM VMX_SEGMENT_SELECTOR
GetGuestFs();

IMPORT_EXPORT_VMM VOID
SetGuestGsSel(PVMX_SEGMENT_SELECTOR Gs);

IMPORT_EXPORT_VMM VOID
SetGuestGs(PVMX_SEGMENT_SELECTOR Gs);

IMPORT_EXPORT_VMM VMX_SEGMENT_SELECTOR
GetGuestGs();

IMPORT_EXPORT_VMM VOID
SetGuestEsSel(PVMX_SEGMENT_SELECTOR Es);

IMPORT_EXPORT_VMM VOID
SetGuestEs(PVMX_SEGMENT_SELECTOR Es);

IMPORT_EXPORT_VMM VMX_SEGMENT_SELECTOR
GetGuestEs();

IMPORT_EXPORT_VMM VOID
SetGuestIdtr(UINT64 Idtr);

IMPORT_EXPORT_VMM UINT64
GetGuestIdtr();

IMPORT_EXPORT_VMM VOID
SetGuestLdtr(UINT64 Ldtr);

IMPORT_EXPORT_VMM UINT64
GetGuestLdtr();

IMPORT_EXPORT_VMM VOID
SetGuestGdtr(UINT64 Gdtr);

IMPORT_EXPORT_VMM UINT64
GetGuestGdtr();

IMPORT_EXPORT_VMM VOID
SetGuestTr(UINT64 Tr);

IMPORT_EXPORT_VMM UINT64
GetGuestTr();

IMPORT_EXPORT_VMM VOID
SetGuestRFlags(UINT64 RFlags);

IMPORT_EXPORT_VMM UINT64
GetGuestRFlags();

IMPORT_EXPORT_VMM VOID
SetGuestRIP(UINT64 RIP);

IMPORT_EXPORT_VMM VOID
SetGuestRSP(UINT64 RSP);

IMPORT_EXPORT_VMM UINT64
GetGuestRIP();

IMPORT_EXPORT_VMM UINT64
GetGuestCr0();

IMPORT_EXPORT_VMM UINT64
GetGuestCr2();

IMPORT_EXPORT_VMM UINT64
GetGuestCr3();

IMPORT_EXPORT_VMM UINT64
GetGuestCr4();

IMPORT_EXPORT_VMM UINT64
GetGuestCr8();

IMPORT_EXPORT_VMM VOID
SetGuestCr0(UINT64 Cr0);

IMPORT_EXPORT_VMM VOID
SetGuestCr2(UINT64 Cr2);

IMPORT_EXPORT_VMM VOID
SetGuestCr3(UINT64 Cr3);

IMPORT_EXPORT_VMM VOID
SetGuestCr4(UINT64 Cr4);

IMPORT_EXPORT_VMM VOID
SetGuestCr8(UINT64 Cr8);

IMPORT_EXPORT_VMM UINT64
GetGuestDr0();

IMPORT_EXPORT_VMM UINT64
GetGuestDr1();

IMPORT_EXPORT_VMM UINT64
GetGuestDr2();

IMPORT_EXPORT_VMM UINT64
GetGuestDr3();

IMPORT_EXPORT_VMM UINT64
GetGuestDr6();

IMPORT_EXPORT_VMM UINT64
GetGuestDr7();

IMPORT_EXPORT_VMM VOID
SetGuestDr0(UINT64 value);

IMPORT_EXPORT_VMM VOID
SetGuestDr1(UINT64 value);

IMPORT_EXPORT_VMM VOID
SetGuestDr2(UINT64 value);

IMPORT_EXPORT_VMM VOID
SetGuestDr3(UINT64 value);

IMPORT_EXPORT_VMM VOID
SetGuestDr6(UINT64 value);

IMPORT_EXPORT_VMM VOID
SetGuestDr7(UINT64 value);

IMPORT_EXPORT_VMM BOOLEAN
SetDebugRegisters(UINT32 DebugRegNum, DEBUG_REGISTER_TYPE ActionType, BOOLEAN ApplyToVmcs, UINT64 TargetAddress);

//////////////////////////////////////////////////
//              Transparent Mode        		//
//////////////////////////////////////////////////

IMPORT_EXPORT_VMM NTSTATUS
TransparentHideDebugger(PDEBUGGER_HIDE_AND_TRANSPARENT_DEBUGGER_MODE Measurements);

IMPORT_EXPORT_VMM NTSTATUS
TransparentUnhideDebugger();

//////////////////////////////////////////////////
//     Non-internal Broadcasting Functions    	//
//////////////////////////////////////////////////

IMPORT_EXPORT_VMM VOID
BroadcastEnableBreakpointExitingOnExceptionBitmapAllCores();

IMPORT_EXPORT_VMM VOID
BroadcastDisableBreakpointExitingOnExceptionBitmapAllCores();

IMPORT_EXPORT_VMM VOID
BroadcastEnableDbAndBpExitingAllCores();

IMPORT_EXPORT_VMM VOID
BroadcastDisableDbAndBpExitingAllCores();

IMPORT_EXPORT_VMM VOID
BroadcastEnableRdtscExitingAllCores();

IMPORT_EXPORT_VMM VOID
BroadcastDisableRdtscExitingAllCores();

IMPORT_EXPORT_VMM VOID
BroadcastChangeAllMsrBitmapReadAllCores(UINT64 BitmapMask);

IMPORT_EXPORT_VMM VOID
BroadcastResetChangeAllMsrBitmapReadAllCores();

IMPORT_EXPORT_VMM VOID
BroadcastChangeAllMsrBitmapWriteAllCores(UINT64 BitmapMask);

IMPORT_EXPORT_VMM VOID
BroadcastResetAllMsrBitmapWriteAllCores();

IMPORT_EXPORT_VMM VOID
BroadcastDisableRdtscExitingForClearingEventsAllCores();

IMPORT_EXPORT_VMM VOID
BroadcastDisableMov2ControlRegsExitingForClearingEventsAllCores(PDEBUGGER_EVENT_OPTIONS BroadcastingOption);

IMPORT_EXPORT_VMM VOID
BroadcastDisableMov2DebugRegsExitingForClearingEventsAllCores();

IMPORT_EXPORT_VMM VOID
BroadcastEnableRdpmcExitingAllCores();

IMPORT_EXPORT_VMM VOID
BroadcastDisableRdpmcExitingAllCores();

IMPORT_EXPORT_VMM VOID
BroadcastSetExceptionBitmapAllCores(UINT64 ExceptionIndex);

IMPORT_EXPORT_VMM VOID
BroadcastUnsetExceptionBitmapAllCores(UINT64 ExceptionIndex);

IMPORT_EXPORT_VMM VOID
BroadcastResetExceptionBitmapAllCores();

IMPORT_EXPORT_VMM VOID
BroadcastEnableMovControlRegisterExitingAllCores(PDEBUGGER_EVENT_OPTIONS BroadcastingOption);

IMPORT_EXPORT_VMM VOID
BroadcastDisableMovToControlRegistersExitingAllCores(PDEBUGGER_EVENT_OPTIONS BroadcastingOption);

IMPORT_EXPORT_VMM VOID
BroadcastEnableMovDebugRegistersExitingAllCores();

IMPORT_EXPORT_VMM VOID
BroadcastDisableMovDebugRegistersExitingAllCores();

IMPORT_EXPORT_VMM VOID
BroadcastSetExternalInterruptExitingAllCores();

IMPORT_EXPORT_VMM VOID
BroadcastUnsetExternalInterruptExitingOnlyOnClearingInterruptEventsAllCores();

IMPORT_EXPORT_VMM VOID
BroadcastIoBitmapChangeAllCores(UINT64 Port);

IMPORT_EXPORT_VMM VOID
BroadcastIoBitmapResetAllCores();

IMPORT_EXPORT_VMM VOID
BroadcastEnableMovToCr3ExitingOnAllProcessors();

IMPORT_EXPORT_VMM VOID
BroadcastDisableMovToCr3ExitingOnAllProcessors();

IMPORT_EXPORT_VMM VOID
BroadcastEnableEferSyscallEventsOnAllProcessors();

IMPORT_EXPORT_VMM VOID
BroadcastDisableEferSyscallEventsOnAllProcessors();


//..\bin\debug\headerAll\commands.h
/**
 * @file commands.h
 * @author Sina Karvandi (sina@hyperdbg.org)
 * @author Alee Amini (alee@hyperdbg.org)
 * @brief The hyperdbg command interpreter and driver connector
 * @details
 * @version 0.1
 * @date 2020-04-11
 *
 * @copyright This project is released under the GNU Public License v3.
 *
 */
#pragma once

using namespace std;

//////////////////////////////////////////////////
//                    Externs                   //
//////////////////////////////////////////////////

extern HANDLE g_DeviceHandle;

//////////////////////////////////////////////////
//                  Settings                    //
//////////////////////////////////////////////////

VOID
CommandSettingsLoadDefaultValuesFromConfigFile();

VOID
CommandSettingsSetValueFromConfigFile(std::string OptionName, std::string OptionValue);

BOOLEAN
CommandSettingsGetValueFromConfigFile(std::string OptionName, std::string & OptionValue);

//////////////////////////////////////////////////
//                  Functions                   //
//////////////////////////////////////////////////

int
ReadCpuDetails();

VOID
ShowMessages(const char * Fmt, ...);

string
SeparateTo64BitValue(UINT64 Value);

void
ShowMemoryCommandDB(unsigned char * OutputBuffer, UINT32 Size, UINT64 Address, DEBUGGER_READ_MEMORY_TYPE MemoryType, UINT64 Length);

void
ShowMemoryCommandDD(unsigned char * OutputBuffer, UINT32 Size, UINT64 Address, DEBUGGER_READ_MEMORY_TYPE MemoryType, UINT64 Length);

void
ShowMemoryCommandDC(unsigned char * OutputBuffer, UINT32 Size, UINT64 Address, DEBUGGER_READ_MEMORY_TYPE MemoryType, UINT64 Length);

void
ShowMemoryCommandDQ(unsigned char * OutputBuffer, UINT32 Size, UINT64 Address, DEBUGGER_READ_MEMORY_TYPE MemoryType, UINT64 Length);

VOID
CommandPteShowResults(UINT64 TargetVa, PDEBUGGER_READ_PAGE_TABLE_ENTRIES_DETAILS PteRead);

DEBUGGER_CONDITIONAL_JUMP_STATUS
HyperDbgIsConditionalJumpTaken(unsigned char * BufferToDisassemble,
                               UINT64          BuffLength,
                               RFLAGS          Rflags,
                               BOOLEAN         Isx86_64);

int
HyperDbgDisassembler64(unsigned char * BufferToDisassemble,
                       UINT64          BaseAddress,
                       UINT64          Size,
                       UINT32          MaximumInstrDecoded,
                       BOOLEAN         ShowBranchIsTakenOrNot,
                       PRFLAGS         Rflags);

int
HyperDbgDisassembler32(unsigned char * BufferToDisassemble,
                       UINT64          BaseAddress,
                       UINT64          Size,
                       UINT32          MaximumInstrDecoded,
                       BOOLEAN         ShowBranchIsTakenOrNot,
                       PRFLAGS         Rflags);

UINT32
HyperDbgLengthDisassemblerEngine(
    unsigned char * BufferToDisassemble,
    UINT64          BuffLength,
    BOOLEAN         Isx86_64);

BOOLEAN
HyperDbgCheckWhetherTheCurrentInstructionIsCall(
    unsigned char * BufferToDisassemble,
    UINT64          BuffLength,
    BOOLEAN         Isx86_64,
    PUINT32         CallLength);

BOOLEAN
HyperDbgCheckWhetherTheCurrentInstructionIsCallOrRet(
    unsigned char * BufferToDisassemble,
    UINT64          CurrentRip,
    UINT32          BuffLength,
    BOOLEAN         Isx86_64,
    PBOOLEAN        IsRet);

BOOLEAN
HyperDbgCheckWhetherTheCurrentInstructionIsRet(
    unsigned char * BufferToDisassemble,
    UINT64          BuffLength,
    BOOLEAN         Isx86_64);

VOID
HyperDbgReadMemoryAndDisassemble(DEBUGGER_SHOW_MEMORY_STYLE   Style,
                                 UINT64                       Address,
                                 DEBUGGER_READ_MEMORY_TYPE    MemoryType,
                                 DEBUGGER_READ_READING_TYPE   ReadingType,
                                 UINT32                       Pid,
                                 UINT32                       Size,
                                 PDEBUGGER_DT_COMMAND_OPTIONS DtDetails);

VOID
InitializeCommandsDictionary();

VOID
InitializeDebugger();

VOID
CommandDumpSaveIntoFile(PVOID Buffer, UINT32 Length);

//////////////////////////////////////////////////
//              Type of Commands                //
//////////////////////////////////////////////////

/**
 * @brief Command's function type
 *
 */
typedef VOID (*CommandFuncType)(vector<string> SplitCommand, string Command);

/**
 * @brief Command's help function type
 *
 */
typedef VOID (*CommandHelpFuncType)();

/**
 * @brief Details of each command
 *
 */
typedef struct _COMMAND_DETAIL
{
    CommandFuncType     CommandFunction;
    CommandHelpFuncType CommandHelpFunction;
    UINT64              CommandAttrib;

} COMMAND_DETAIL, *PCOMMAND_DETAIL;

/**
 * @brief Type saving commands and mapping to command string
 *
 */
typedef std::map<std::string, COMMAND_DETAIL> CommandType;

/**
 * @brief Different attributes of commands
 *
 */
#define DEBUGGER_COMMAND_ATTRIBUTE_EVENT \
    0x1 | DEBUGGER_COMMAND_ATTRIBUTE_LOCAL_COMMAND_IN_DEBUGGER_MODE | DEBUGGER_COMMAND_ATTRIBUTE_LOCAL_CASE_SENSITIVE
#define DEBUGGER_COMMAND_ATTRIBUTE_LOCAL_COMMAND_IN_DEBUGGER_MODE     0x2
#define DEBUGGER_COMMAND_ATTRIBUTE_LOCAL_COMMAND_IN_REMOTE_CONNECTION 0x4
#define DEBUGGER_COMMAND_ATTRIBUTE_LOCAL_CASE_SENSITIVE               0x8
#define DEBUGGER_COMMAND_ATTRIBUTE_REPEAT_ON_ENTER                    0x10
#define DEBUGGER_COMMAND_ATTRIBUTE_WONT_STOP_DEBUGGER_AGAIN           0x20
#define DEBUGGER_COMMAND_ATTRIBUTE_HWDBG                              0x40

/**
 * @brief Absolute local commands
 *
 */
#define DEBUGGER_COMMAND_ATTRIBUTE_ABSOLUTE_LOCAL               \
    DEBUGGER_COMMAND_ATTRIBUTE_LOCAL_COMMAND_IN_DEBUGGER_MODE | \
        DEBUGGER_COMMAND_ATTRIBUTE_LOCAL_COMMAND_IN_REMOTE_CONNECTION

/**
 * @brief Command's attributes
 *
 */
#define DEBUGGER_COMMAND_CLEAR_ATTRIBUTES \
    DEBUGGER_COMMAND_ATTRIBUTE_ABSOLUTE_LOCAL

#define DEBUGGER_COMMAND_HELP_ATTRIBUTES \
    DEBUGGER_COMMAND_ATTRIBUTE_ABSOLUTE_LOCAL

#define DEBUGGER_COMMAND_CONNECT_ATTRIBUTES \
    DEBUGGER_COMMAND_ATTRIBUTE_ABSOLUTE_LOCAL

#define DEBUGGER_COMMAND_LISTEN_ATTRIBUTES \
    DEBUGGER_COMMAND_ATTRIBUTE_ABSOLUTE_LOCAL

#define DEBUGGER_COMMAND_G_ATTRIBUTES DEBUGGER_COMMAND_ATTRIBUTE_ABSOLUTE_LOCAL | DEBUGGER_COMMAND_ATTRIBUTE_REPEAT_ON_ENTER

#define DEBUGGER_COMMAND_ATTACH_ATTRIBUTES DEBUGGER_COMMAND_ATTRIBUTE_LOCAL_COMMAND_IN_DEBUGGER_MODE

#define DEBUGGER_COMMAND_DETACH_ATTRIBUTES DEBUGGER_COMMAND_ATTRIBUTE_LOCAL_COMMAND_IN_DEBUGGER_MODE

#define DEBUGGER_COMMAND_SWITCH_ATTRIBUTES DEBUGGER_COMMAND_ATTRIBUTE_LOCAL_COMMAND_IN_DEBUGGER_MODE

#define DEBUGGER_COMMAND_START_ATTRIBUTES DEBUGGER_COMMAND_ATTRIBUTE_WONT_STOP_DEBUGGER_AGAIN

#define DEBUGGER_COMMAND_RESTART_ATTRIBUTES DEBUGGER_COMMAND_ATTRIBUTE_WONT_STOP_DEBUGGER_AGAIN

#define DEBUGGER_COMMAND_KILL_ATTRIBUTES NULL

#define DEBUGGER_COMMAND_PROCESS_ATTRIBUTES \
    DEBUGGER_COMMAND_ATTRIBUTE_LOCAL_COMMAND_IN_DEBUGGER_MODE

#define DEBUGGER_COMMAND_THREAD_ATTRIBUTES \
    DEBUGGER_COMMAND_ATTRIBUTE_LOCAL_COMMAND_IN_DEBUGGER_MODE

#define DEBUGGER_COMMAND_SLEEP_ATTRIBUTES \
    DEBUGGER_COMMAND_ATTRIBUTE_ABSOLUTE_LOCAL

#define DEBUGGER_COMMAND_EVENTS_ATTRIBUTES \
    DEBUGGER_COMMAND_ATTRIBUTE_LOCAL_COMMAND_IN_DEBUGGER_MODE

#define DEBUGGER_COMMAND_SETTINGS_ATTRIBUTES \
    DEBUGGER_COMMAND_ATTRIBUTE_LOCAL_COMMAND_IN_DEBUGGER_MODE

#define DEBUGGER_COMMAND_DISCONNECT_ATTRIBUTES \
    DEBUGGER_COMMAND_ATTRIBUTE_ABSOLUTE_LOCAL

#define DEBUGGER_COMMAND_DEBUG_ATTRIBUTES \
    DEBUGGER_COMMAND_ATTRIBUTE_ABSOLUTE_LOCAL | DEBUGGER_COMMAND_ATTRIBUTE_LOCAL_CASE_SENSITIVE

#define DEBUGGER_COMMAND_DOT_STATUS_ATTRIBUTES \
    DEBUGGER_COMMAND_ATTRIBUTE_ABSOLUTE_LOCAL

#define DEBUGGER_COMMAND_STATUS_ATTRIBUTES NULL

#define DEBUGGER_COMMAND_LOAD_ATTRIBUTES NULL

#define DEBUGGER_COMMAND_EXIT_ATTRIBUTES \
    DEBUGGER_COMMAND_ATTRIBUTE_ABSOLUTE_LOCAL

#define DEBUGGER_COMMAND_FLUSH_ATTRIBUTES \
    DEBUGGER_COMMAND_ATTRIBUTE_LOCAL_COMMAND_IN_DEBUGGER_MODE

#define DEBUGGER_COMMAND_PAUSE_ATTRIBUTES \
    DEBUGGER_COMMAND_ATTRIBUTE_ABSOLUTE_LOCAL

#define DEBUGGER_COMMAND_UNLOAD_ATTRIBUTES NULL

#define DEBUGGER_COMMAND_SCRIPT_ATTRIBUTES \
    DEBUGGER_COMMAND_ATTRIBUTE_ABSOLUTE_LOCAL | DEBUGGER_COMMAND_ATTRIBUTE_LOCAL_CASE_SENSITIVE

#define DEBUGGER_COMMAND_OUTPUT_ATTRIBUTES \
    DEBUGGER_COMMAND_ATTRIBUTE_LOCAL_COMMAND_IN_DEBUGGER_MODE | DEBUGGER_COMMAND_ATTRIBUTE_LOCAL_CASE_SENSITIVE

#define DEBUGGER_COMMAND_PRINT_ATTRIBUTES \
    DEBUGGER_COMMAND_ATTRIBUTE_LOCAL_COMMAND_IN_DEBUGGER_MODE | DEBUGGER_COMMAND_ATTRIBUTE_LOCAL_CASE_SENSITIVE | DEBUGGER_COMMAND_ATTRIBUTE_REPEAT_ON_ENTER

#define DEBUGGER_COMMAND_EVAL_ATTRIBUTES \
    DEBUGGER_COMMAND_ATTRIBUTE_LOCAL_COMMAND_IN_DEBUGGER_MODE | DEBUGGER_COMMAND_ATTRIBUTE_LOCAL_CASE_SENSITIVE | DEBUGGER_COMMAND_ATTRIBUTE_REPEAT_ON_ENTER

#define DEBUGGER_COMMAND_LOGOPEN_ATTRIBUTES \
    DEBUGGER_COMMAND_ATTRIBUTE_ABSOLUTE_LOCAL | DEBUGGER_COMMAND_ATTRIBUTE_LOCAL_CASE_SENSITIVE

#define DEBUGGER_COMMAND_LOGCLOSE_ATTRIBUTES \
    DEBUGGER_COMMAND_ATTRIBUTE_ABSOLUTE_LOCAL

#define DEBUGGER_COMMAND_TEST_ATTRIBUTES DEBUGGER_COMMAND_ATTRIBUTE_LOCAL_COMMAND_IN_DEBUGGER_MODE

#define DEBUGGER_COMMAND_CPU_ATTRIBUTES NULL

#define DEBUGGER_COMMAND_WRMSR_ATTRIBUTES NULL

#define DEBUGGER_COMMAND_RDMSR_ATTRIBUTES NULL

#define DEBUGGER_COMMAND_VA2PA_ATTRIBUTES DEBUGGER_COMMAND_ATTRIBUTE_LOCAL_CASE_SENSITIVE | DEBUGGER_COMMAND_ATTRIBUTE_LOCAL_COMMAND_IN_DEBUGGER_MODE

#define DEBUGGER_COMMAND_PA2VA_ATTRIBUTES DEBUGGER_COMMAND_ATTRIBUTE_LOCAL_CASE_SENSITIVE | DEBUGGER_COMMAND_ATTRIBUTE_LOCAL_COMMAND_IN_DEBUGGER_MODE

#define DEBUGGER_COMMAND_FORMATS_ATTRIBUTES \
    DEBUGGER_COMMAND_ATTRIBUTE_LOCAL_COMMAND_IN_DEBUGGER_MODE | DEBUGGER_COMMAND_ATTRIBUTE_LOCAL_CASE_SENSITIVE | DEBUGGER_COMMAND_ATTRIBUTE_REPEAT_ON_ENTER

#define DEBUGGER_COMMAND_PTE_ATTRIBUTES DEBUGGER_COMMAND_ATTRIBUTE_LOCAL_CASE_SENSITIVE | DEBUGGER_COMMAND_ATTRIBUTE_LOCAL_COMMAND_IN_DEBUGGER_MODE

#define DEBUGGER_COMMAND_CORE_ATTRIBUTES \
    DEBUGGER_COMMAND_ATTRIBUTE_LOCAL_COMMAND_IN_DEBUGGER_MODE

#define DEBUGGER_COMMAND_MONITOR_ATTRIBUTES DEBUGGER_COMMAND_ATTRIBUTE_EVENT

#define DEBUGGER_COMMAND_VMCALL_ATTRIBUTES DEBUGGER_COMMAND_ATTRIBUTE_EVENT

#define DEBUGGER_COMMAND_EPTHOOK_ATTRIBUTES DEBUGGER_COMMAND_ATTRIBUTE_EVENT

#define DEBUGGER_COMMAND_EPTHOOK2_ATTRIBUTES DEBUGGER_COMMAND_ATTRIBUTE_EVENT

#define DEBUGGER_COMMAND_CPUID_ATTRIBUTES DEBUGGER_COMMAND_ATTRIBUTE_EVENT

#define DEBUGGER_COMMAND_MSRREAD_ATTRIBUTES DEBUGGER_COMMAND_ATTRIBUTE_EVENT

#define DEBUGGER_COMMAND_MSRWRITE_ATTRIBUTES DEBUGGER_COMMAND_ATTRIBUTE_EVENT

#define DEBUGGER_COMMAND_TSC_ATTRIBUTES DEBUGGER_COMMAND_ATTRIBUTE_EVENT

#define DEBUGGER_COMMAND_PMC_ATTRIBUTES DEBUGGER_COMMAND_ATTRIBUTE_EVENT

#define DEBUGGER_COMMAND_CRWRITE_ATTRIBUTES DEBUGGER_COMMAND_ATTRIBUTE_EVENT

#define DEBUGGER_COMMAND_DR_ATTRIBUTES DEBUGGER_COMMAND_ATTRIBUTE_EVENT

#define DEBUGGER_COMMAND_IOIN_ATTRIBUTES DEBUGGER_COMMAND_ATTRIBUTE_EVENT

#define DEBUGGER_COMMAND_IOOUT_ATTRIBUTES DEBUGGER_COMMAND_ATTRIBUTE_EVENT

#define DEBUGGER_COMMAND_EXCEPTION_ATTRIBUTES DEBUGGER_COMMAND_ATTRIBUTE_EVENT

#define DEBUGGER_COMMAND_INTERRUPT_ATTRIBUTES DEBUGGER_COMMAND_ATTRIBUTE_EVENT

#define DEBUGGER_COMMAND_SYSCALL_ATTRIBUTES DEBUGGER_COMMAND_ATTRIBUTE_EVENT

#define DEBUGGER_COMMAND_SYSRET_ATTRIBUTES DEBUGGER_COMMAND_ATTRIBUTE_EVENT

#define DEBUGGER_COMMAND_MODE_ATTRIBUTES DEBUGGER_COMMAND_ATTRIBUTE_EVENT

#define DEBUGGER_COMMAND_TRACE_ATTRIBUTES DEBUGGER_COMMAND_ATTRIBUTE_EVENT

#define DEBUGGER_COMMAND_HIDE_ATTRIBUTES DEBUGGER_COMMAND_ATTRIBUTE_LOCAL_CASE_SENSITIVE

#define DEBUGGER_COMMAND_UNHIDE_ATTRIBUTES NULL

#define DEBUGGER_COMMAND_MEASURE_ATTRIBUTES NULL

#define DEBUGGER_COMMAND_LM_ATTRIBUTES NULL

#define DEBUGGER_COMMAND_P_ATTRIBUTES \
    DEBUGGER_COMMAND_ATTRIBUTE_LOCAL_COMMAND_IN_DEBUGGER_MODE | DEBUGGER_COMMAND_ATTRIBUTE_REPEAT_ON_ENTER

#define DEBUGGER_COMMAND_T_ATTRIBUTES \
    DEBUGGER_COMMAND_ATTRIBUTE_LOCAL_COMMAND_IN_DEBUGGER_MODE | DEBUGGER_COMMAND_ATTRIBUTE_REPEAT_ON_ENTER

#define DEBUGGER_COMMAND_I_ATTRIBUTES \
    DEBUGGER_COMMAND_ATTRIBUTE_LOCAL_COMMAND_IN_DEBUGGER_MODE | DEBUGGER_COMMAND_ATTRIBUTE_REPEAT_ON_ENTER

#define DEBUGGER_COMMAND_D_AND_U_ATTRIBUTES \
    DEBUGGER_COMMAND_ATTRIBUTE_LOCAL_COMMAND_IN_DEBUGGER_MODE | DEBUGGER_COMMAND_ATTRIBUTE_LOCAL_CASE_SENSITIVE | DEBUGGER_COMMAND_ATTRIBUTE_REPEAT_ON_ENTER

#define DEBUGGER_COMMAND_E_ATTRIBUTES \
    DEBUGGER_COMMAND_ATTRIBUTE_LOCAL_COMMAND_IN_DEBUGGER_MODE | DEBUGGER_COMMAND_ATTRIBUTE_LOCAL_CASE_SENSITIVE

#define DEBUGGER_COMMAND_S_ATTRIBUTES \
    DEBUGGER_COMMAND_ATTRIBUTE_LOCAL_COMMAND_IN_DEBUGGER_MODE | DEBUGGER_COMMAND_ATTRIBUTE_LOCAL_CASE_SENSITIVE

#define DEBUGGER_COMMAND_R_ATTRIBUTES \
    DEBUGGER_COMMAND_ATTRIBUTE_LOCAL_COMMAND_IN_DEBUGGER_MODE | DEBUGGER_COMMAND_ATTRIBUTE_LOCAL_CASE_SENSITIVE | DEBUGGER_COMMAND_ATTRIBUTE_REPEAT_ON_ENTER

#define DEBUGGER_COMMAND_BP_ATTRIBUTES \
    DEBUGGER_COMMAND_ATTRIBUTE_LOCAL_COMMAND_IN_DEBUGGER_MODE | DEBUGGER_COMMAND_ATTRIBUTE_LOCAL_CASE_SENSITIVE

#define DEBUGGER_COMMAND_BE_ATTRIBUTES \
    DEBUGGER_COMMAND_ATTRIBUTE_LOCAL_COMMAND_IN_DEBUGGER_MODE

#define DEBUGGER_COMMAND_BD_ATTRIBUTES \
    DEBUGGER_COMMAND_ATTRIBUTE_LOCAL_COMMAND_IN_DEBUGGER_MODE

#define DEBUGGER_COMMAND_BC_ATTRIBUTES \
    DEBUGGER_COMMAND_ATTRIBUTE_LOCAL_COMMAND_IN_DEBUGGER_MODE

#define DEBUGGER_COMMAND_BL_ATTRIBUTES \
    DEBUGGER_COMMAND_ATTRIBUTE_LOCAL_COMMAND_IN_DEBUGGER_MODE

#define DEBUGGER_COMMAND_SYMPATH_ATTRIBUTES \
    DEBUGGER_COMMAND_ATTRIBUTE_LOCAL_COMMAND_IN_DEBUGGER_MODE | DEBUGGER_COMMAND_ATTRIBUTE_LOCAL_CASE_SENSITIVE

#define DEBUGGER_COMMAND_SYM_ATTRIBUTES \
    DEBUGGER_COMMAND_ATTRIBUTE_LOCAL_COMMAND_IN_DEBUGGER_MODE | DEBUGGER_COMMAND_ATTRIBUTE_LOCAL_CASE_SENSITIVE

#define DEBUGGER_COMMAND_X_ATTRIBUTES \
    DEBUGGER_COMMAND_ATTRIBUTE_LOCAL_COMMAND_IN_DEBUGGER_MODE | DEBUGGER_COMMAND_ATTRIBUTE_LOCAL_CASE_SENSITIVE

#define DEBUGGER_COMMAND_PREALLOC_ATTRIBUTES NULL

#define DEBUGGER_COMMAND_PREACTIVATE_ATTRIBUTES NULL

#define DEBUGGER_COMMAND_K_ATTRIBUTES \
    DEBUGGER_COMMAND_ATTRIBUTE_LOCAL_COMMAND_IN_DEBUGGER_MODE

#define DEBUGGER_COMMAND_DT_ATTRIBUTES \
    DEBUGGER_COMMAND_ATTRIBUTE_LOCAL_COMMAND_IN_DEBUGGER_MODE | DEBUGGER_COMMAND_ATTRIBUTE_LOCAL_CASE_SENSITIVE | DEBUGGER_COMMAND_ATTRIBUTE_REPEAT_ON_ENTER

#define DEBUGGER_COMMAND_STRUCT_ATTRIBUTES \
    DEBUGGER_COMMAND_ATTRIBUTE_LOCAL_COMMAND_IN_DEBUGGER_MODE | DEBUGGER_COMMAND_ATTRIBUTE_LOCAL_CASE_SENSITIVE

#define DEBUGGER_COMMAND_PE_ATTRIBUTES NULL

// #define DEBUGGER_COMMAND_REV_ATTRIBUTES DEBUGGER_COMMAND_ATTRIBUTE_WONT_STOP_DEBUGGER_AGAIN
#define DEBUGGER_COMMAND_REV_ATTRIBUTES NULL

#define DEBUGGER_COMMAND_TRACK_ATTRIBUTES DEBUGGER_COMMAND_ATTRIBUTE_LOCAL_COMMAND_IN_DEBUGGER_MODE

#define DEBUGGER_COMMAND_PAGEIN_ATTRIBUTES DEBUGGER_COMMAND_ATTRIBUTE_LOCAL_CASE_SENSITIVE | DEBUGGER_COMMAND_ATTRIBUTE_LOCAL_COMMAND_IN_DEBUGGER_MODE

#define DEBUGGER_COMMAND_DUMP_ATTRIBUTES \
    DEBUGGER_COMMAND_ATTRIBUTE_LOCAL_COMMAND_IN_DEBUGGER_MODE | DEBUGGER_COMMAND_ATTRIBUTE_LOCAL_CASE_SENSITIVE

#define DEBUGGER_COMMAND_GU_ATTRIBUTES \
    DEBUGGER_COMMAND_ATTRIBUTE_LOCAL_COMMAND_IN_DEBUGGER_MODE | DEBUGGER_COMMAND_ATTRIBUTE_REPEAT_ON_ENTER

#define DEBUGGER_COMMAND_HWDBG_HW_CLK_ATTRIBUTES DEBUGGER_COMMAND_ATTRIBUTE_HWDBG

//////////////////////////////////////////////////
//             Command Functions                //
//////////////////////////////////////////////////

VOID
CommandTest(vector<string> SplitCommand, string Command);

VOID
CommandClearScreen(vector<string> SplitCommand, string Command);

VOID
CommandReadMemoryAndDisassembler(vector<string> SplitCommand,
                                 string         Command);

VOID
CommandConnect(vector<string> SplitCommand, string Command);

VOID
CommandLoad(vector<string> SplitCommand, string Command);

VOID
CommandUnload(vector<string> SplitCommand, string Command);

VOID
CommandScript(vector<string> SplitCommand, string Command);

VOID
CommandCpu(vector<string> SplitCommand, string Command);

VOID
CommandExit(vector<string> SplitCommand, string Command);

VOID
CommandDisconnect(vector<string> SplitCommand, string Command);

VOID
CommandFormats(vector<string> SplitCommand, string Command);

VOID
CommandRdmsr(vector<string> SplitCommand, string Command);

VOID
CommandWrmsr(vector<string> SplitCommand, string Command);

VOID
CommandPte(vector<string> SplitCommand, string Command);

VOID
CommandMonitor(vector<string> SplitCommand, string Command);

VOID
CommandSyscallAndSysret(vector<string> SplitCommand, string Command);

VOID
CommandEptHook(vector<string> SplitCommand, string Command);

VOID
CommandEptHook2(vector<string> SplitCommand, string Command);

VOID
CommandCpuid(vector<string> SplitCommand, string Command);

VOID
CommandMsrread(vector<string> SplitCommand, string Command);

VOID
CommandMsrwrite(vector<string> SplitCommand, string Command);

VOID
CommandTsc(vector<string> SplitCommand, string Command);

VOID
CommandPmc(vector<string> SplitCommand, string Command);

VOID
CommandException(vector<string> SplitCommand, string Command);

VOID
CommandCrwrite(vector<string> SplitCommand, string Command);

VOID
CommandDr(vector<string> SplitCommand, string Command);

VOID
CommandInterrupt(vector<string> SplitCommand, string Command);

VOID
CommandIoin(vector<string> SplitCommand, string Command);

VOID
CommandIoout(vector<string> SplitCommand, string Command);

VOID
CommandVmcall(vector<string> SplitCommand, string Command);

VOID
CommandMode(vector<string> SplitCommand, string Command);

VOID
CommandTrace(vector<string> SplitCommand, string Command);

VOID
CommandHide(vector<string> SplitCommand, string Command);

VOID
CommandUnhide(vector<string> SplitCommand, string Command);

VOID
CommandLogopen(vector<string> SplitCommand, string Command);

VOID
CommandLogclose(vector<string> SplitCommand, string Command);

VOID
CommandVa2pa(vector<string> SplitCommand, string Command);

VOID
CommandPa2va(vector<string> SplitCommand, string Command);

VOID
CommandEvents(vector<string> SplitCommand, string Command);

VOID
CommandG(vector<string> SplitCommand, string Command);

VOID
CommandLm(vector<string> SplitCommand, string Command);

VOID
CommandSleep(vector<string> SplitCommand, string Command);

VOID
CommandEditMemory(vector<string> SplitCommand, string Command);

VOID
CommandSearchMemory(vector<string> SplitCommand, string Command);

VOID
CommandMeasure(vector<string> SplitCommand, string Command);

VOID
CommandSettings(vector<string> SplitCommand, string Command);

VOID
CommandFlush(vector<string> SplitCommand, string Command);

VOID
CommandPause(vector<string> SplitCommand, string Command);

VOID
CommandListen(vector<string> SplitCommand, string Command);

VOID
CommandStatus(vector<string> SplitCommand, string Command);

VOID
CommandAttach(vector<string> SplitCommand, string Command);

VOID
CommandDetach(vector<string> SplitCommand, string Command);

VOID
CommandStart(vector<string> SplitCommand, string Command);

VOID
CommandRestart(vector<string> SplitCommand, string Command);

VOID
CommandSwitch(vector<string> SplitCommand, string Command);

VOID
CommandKill(vector<string> SplitCommand, string Command);

VOID
CommandT(vector<string> SplitCommand, string Command);

VOID
CommandI(vector<string> SplitCommand, string Command);

VOID
CommandPrint(vector<string> SplitCommand, string Command);

VOID
CommandOutput(vector<string> SplitCommand, string Command);

VOID
CommandDebug(vector<string> SplitCommand, string Command);

VOID
CommandP(vector<string> SplitCommand, string Command);

VOID
CommandCore(vector<string> SplitCommand, string Command);

VOID
CommandProcess(vector<string> SplitCommand, string Command);

VOID
CommandThread(vector<string> SplitCommand, string Command);

VOID
CommandEval(vector<string> SplitCommand, string Command);

VOID
CommandR(vector<string> SplitCommand, string Command);

VOID
CommandBp(vector<string> SplitCommand, string Command);

VOID
CommandBl(vector<string> SplitCommand, string Command);

VOID
CommandBe(vector<string> SplitCommand, string Command);

VOID
CommandBd(vector<string> SplitCommand, string Command);

VOID
CommandBc(vector<string> SplitCommand, string Command);

VOID
CommandSympath(vector<string> SplitCommand, string Command);

VOID
CommandSym(vector<string> SplitCommand, string Command);

VOID
CommandX(vector<string> SplitCommand, string Command);

VOID
CommandPrealloc(vector<string> SplitCommand, string Command);

VOID
CommandPreactivate(vector<string> SplitCommand, string Command);

VOID
CommandDtAndStruct(vector<string> SplitCommand, string Command);

VOID
CommandK(vector<string> SplitCommand, string Command);

VOID
CommandPe(vector<string> SplitCommand, string Command);

VOID
CommandRev(vector<string> SplitCommand, string Command);

VOID
CommandTrack(vector<string> SplitCommand, string Command);

VOID
CommandPagein(vector<string> SplitCommand, string Command);

VOID
CommandDump(vector<string> SplitCommand, string Command);

VOID
CommandGu(vector<string> SplitCommand, string Command);

//
// hwdbg commands
//
VOID
CommandHwClk(vector<string> SplitCommand, string Command);


//..\bin\debug\headerAll\common.h
/**
 * @file common.h
 * @author Sina Karvandi (sina@hyperdbg.org)
 * @brief header for HyperDbg's general functions for reading and converting and
 * etc
 * @details
 * @version 0.1
 * @date 2020-05-27
 *
 * @copyright This project is released under the GNU Public License v3.
 *
 */
#pragma once

//////////////////////////////////////////////////
//			    	 Definitions                //
//////////////////////////////////////////////////

#define AssertReturn return;

#define AssertReturnFalse return FALSE;

#define AssertReturnOne return 1;

#define ASSERT_MESSAGE_DRIVER_NOT_LOADED "handle of the driver not found, probably the driver is not loaded. Did you use 'load' command?\n"

#define ASSERT_MESSAGE_BUILD_SIGNATURE_DOESNT_MATCH "the handshaking process was successful; however, there is a mismatch between " \
                                                    "the version/build of the debuggee and the debugger. please use the same "      \
                                                    "version/build for both the debuggee and debugger\n"

#define ASSERT_MESSAGE_CANNOT_SPECIFY_PID "err, since HyperDbg won't context-switch to keep the system in a halted state, "                       \
                                          "you cannot specify 'pid' for this command in the debugger mode. You can switch to the target process " \
                                          "memory layout using the '.process' or the '.thread' command. After that, you can use "                 \
                                          "this command without specifying the process ID. Alternatively, you can modify the current "            \
                                          "CR3 register to achieve the same functionality\n"

#define AssertReturnStmt(expr, stmt, rc) \
    do                                   \
    {                                    \
        if (expr)                        \
        {                                \
            /* likely */                 \
        }                                \
        else                             \
        {                                \
            stmt;                        \
            rc;                          \
        }                                \
    } while (0)

#define AssertShowMessageReturnStmt(expr, message, rc) \
    do                                                 \
    {                                                  \
        if (expr)                                      \
        {                                              \
            /* likely */                               \
        }                                              \
        else                                           \
        {                                              \
            ShowMessages(message);                     \
            rc;                                        \
        }                                              \
    } while (0)

/**
 * @brief Size of each page (4096 bytes)
 *
 */
#define PAGE_SIZE 0x1000

/**
 * @brief Aligning a page
 *
 */
#define PAGE_ALIGN(Va) ((PVOID)((ULONG_PTR)(Va) & ~(PAGE_SIZE - 1)))

/**
 * @brief Cpuid to get virtual address width
 *
 */
#define CPUID_ADDR_WIDTH 0x80000008

//////////////////////////////////////////////////
//			  Assembly Functions                //
//////////////////////////////////////////////////

#ifdef __cplusplus
extern "C" {
#endif

extern bool
AsmVmxSupportDetection();

#ifdef __cplusplus
}
#endif

//////////////////////////////////////////////////
//			    	 Spinlocks                  //
//////////////////////////////////////////////////

void
SpinlockLock(volatile LONG * Lock);

void
SpinlockLockWithCustomWait(volatile LONG * Lock, unsigned MaximumWait);

void
SpinlockUnlock(volatile LONG * Lock);

//////////////////////////////////////////////////
//			    	 Functions                  //
//////////////////////////////////////////////////

VOID
PrintBits(const UINT32 size, const void * ptr);

BOOL
Replace(std::string & str, const std::string & from, const std::string & to);

VOID
ReplaceAll(string & str, const string & from, const string & to);

const vector<string>
Split(const string & s, const char & c);

BOOLEAN
IsNumber(const string & str);

vector<string>
SplitIp(const string & str, char delim);

BOOLEAN
IsHexNotation(const string & s);

vector<char>
HexToBytes(const string & hex);

BOOLEAN
ConvertStringToUInt64(string TextToConvert, PUINT64 Result);

BOOLEAN
ConvertStringToUInt32(string TextToConvert, PUINT32 Result);

BOOLEAN
HasEnding(string const & fullString, string const & ending);

BOOLEAN
ValidateIP(const string & ip);

BOOL
SetPrivilege(HANDLE  Token,          // access token handle
             LPCTSTR Privilege,      // name of privilege to enable/disable
             BOOL    EnablePrivilege // to enable or disable privilege
);

void
Trim(std::string & s);

std::string
RemoveSpaces(std::string str);

BOOLEAN
IsFileExistA(const char * FileName);

BOOLEAN
IsFileExistW(const wchar_t * FileName);

VOID
GetConfigFilePath(PWCHAR ConfigPath);

VOID
StringToWString(std::wstring & ws, const std::string & s);

VOID
SplitPathAndArgs(std::vector<std::string> & Qargs, const std::string & Command);

size_t
FindCaseInsensitive(std::string Input, std::string ToSearch, size_t Pos);

size_t
FindCaseInsensitiveW(std::wstring Input, std::wstring ToSearch, size_t Pos);

char *
ConvertStringVectorToCharPointerArray(const std::string & s);

std::vector<std::string>
ListDirectory(const std::string & Directory, const std::string & Extension);

BOOLEAN
IsEmptyString(char * Text);

VOID
CommonCpuidInstruction(UINT32 Func, UINT32 SubFunc, int * CpuInfo);

BOOLEAN
CheckCpuSupportRtm();

UINT32
Getx86VirtualAddressWidth();

BOOLEAN
CheckAccessValidityAndSafety(UINT64 TargetAddress, UINT32 Size);


//..\bin\debug\headerAll\communication.h
/**
 * @file tcpcommunication.h
 * @author Sina Karvandi (sina@hyperdbg.org)
 * @brief Communication over TCP (header)
 * @details
 * @version 0.1
 * @date 2020-08-21
 *
 * @copyright This project is released under the GNU Public License v3.
 *
 */
#pragma once

//////////////////////////////////////////
//			Serial Constants             //
//////////////////////////////////////////

#define COM1_PORT 0x03F8
#define COM2_PORT 0x02F8
#define COM3_PORT 0x03E8
#define COM4_PORT 0x02E8

//////////////////////////////////////////
//			   	Server 		            //
//////////////////////////////////////////

int
CommunicationServerCreateServerAndWaitForClient(PCSTR    Port,
                                                SOCKET * ClientSocketArg,
                                                SOCKET * ListenSocketArg);

int
CommunicationServerReceiveMessage(SOCKET ClientSocket, char * recvbuf, int recvbuflen);

int
CommunicationServerSendMessage(SOCKET ClientSocket, const char * sendbuf, int length);

int
CommunicationServerShutdownAndCleanupConnection(SOCKET ClientSocket,
                                                SOCKET ListenSocket);

//////////////////////////////////////////
//                Client                //
//////////////////////////////////////////

int
CommunicationClientConnectToServer(PCSTR Ip, PCSTR Port, SOCKET * ConnectSocketArg);

int
CommunicationClientSendMessage(SOCKET ConnectSocket, const char * sendbuf, int buflen);

int
CommunicationClientShutdownConnection(SOCKET ConnectSocket);

int
CommunicationClientReceiveMessage(SOCKET ConnectSocket, CHAR * RecvBuf, UINT32 MaxBuffLen, PUINT32 BuffLenRecvd);

int
CommunicationClientCleanup(SOCKET ConnectSocket);

//////////////////////////////////////////
//     Handle Remote Connection         //
//////////////////////////////////////////

VOID
RemoteConnectionListen(PCSTR Port);

VOID
RemoteConnectionConnect(PCSTR Ip, PCSTR Port);

int
RemoteConnectionSendCommand(const char * sendbuf, int len);

int
RemoteConnectionSendResultsToHost(const char * sendbuf, int len);

int
RemoteConnectionCloseTheConnectionWithDebuggee();


//..\bin\debug\headerAll\debugger.h
/**
 * @file debugger.h
 * @author Sina Karvandi (sina@hyperdbg.org)
 * @brief General debugger functions
 * @details
 * @version 0.1
 * @date 2020-05-27
 *
 * @copyright This project is released under the GNU Public License v3.
 *
 */
#pragma once

//////////////////////////////////////////////////
//		Debugger Synchronization Objects        //
//////////////////////////////////////////////////

/**
 * @brief maximum number of event handles in kernel-debugger
 */
#define DEBUGGER_MAXIMUM_SYNCRONIZATION_KERNEL_DEBUGGER_OBJECTS 0x40

/**
 * @brief An event to show whether the debugger is running
 * or not in kernel-debugger
 *
 */

//
// Kernel-debugger
//
#define DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_IS_DEBUGGER_RUNNING                 0x0
#define DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_STARTED_PACKET_RECEIVED             0x1
#define DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_PAUSED_DEBUGGEE_DETAILS             0x2
#define DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_CORE_SWITCHING_RESULT               0x3
#define DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_PROCESS_SWITCHING_RESULT            0x4
#define DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_THREAD_SWITCHING_RESULT             0x5
#define DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_SCRIPT_RUNNING_RESULT               0x6
#define DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_SCRIPT_FORMATS_RESULT               0x7
#define DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_DEBUGGEE_FINISHED_COMMAND_EXECUTION 0x8
#define DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_FLUSH_RESULT                        0x9
#define DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_REGISTER_EVENT                      0xa
#define DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_ADD_ACTION_TO_EVENT                 0xb
#define DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_MODIFY_AND_QUERY_EVENT              0xc
#define DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_READ_REGISTERS                      0xd
#define DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_BP                                  0xe
#define DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_LIST_OR_MODIFY_BREAKPOINTS          0xf
#define DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_READ_MEMORY                         0x10
#define DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_EDIT_MEMORY                         0x11
#define DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_SYMBOL_RELOAD                       0x12
#define DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_TEST_QUERY                          0x13
#define DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_CALLSTACK_RESULT                    0x14
#define DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_SEARCH_QUERY_RESULT                 0x15
#define DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_VA2PA_AND_PA2VA_RESULT              0x16
#define DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_PTE_RESULT                          0x17
#define DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_SHORT_CIRCUITING_EVENT_STATE        0x18
#define DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_PAGE_IN_STATE                       0x19

//////////////////////////////////////////////////
//               Event Details                  //
//////////////////////////////////////////////////

/**
 * @brief Reason for error in parsing commands
 *
 */
typedef enum _DEBUGGER_EVENT_PARSING_ERROR_CAUSE
{
    DEBUGGER_EVENT_PARSING_ERROR_CAUSE_SUCCESSFUL_NO_ERROR                          = 0,
    DEBUGGER_EVENT_PARSING_ERROR_CAUSE_SCRIPT_SYNTAX_ERROR                          = 1,
    DEBUGGER_EVENT_PARSING_ERROR_CAUSE_NO_INPUT                                     = 2,
    DEBUGGER_EVENT_PARSING_ERROR_CAUSE_MAXIMUM_INPUT_REACHED                        = 3,
    DEBUGGER_EVENT_PARSING_ERROR_CAUSE_OUTPUT_NAME_NOT_FOUND                        = 4,
    DEBUGGER_EVENT_PARSING_ERROR_CAUSE_OUTPUT_SOURCE_ALREADY_CLOSED                 = 5,
    DEBUGGER_EVENT_PARSING_ERROR_CAUSE_ALLOCATION_ERROR                             = 6,
    DEBUGGER_EVENT_PARSING_ERROR_CAUSE_FORMAT_ERROR                                 = 7,
    DEBUGGER_EVENT_PARSING_ERROR_CAUSE_ATTEMPT_TO_BREAK_ON_VMI_MODE                 = 8,
    DEBUGGER_EVENT_PARSING_ERROR_CAUSE_IMMEDIATE_MESSAGING_IN_EVENT_FORWARDING_MODE = 9,
    DEBUGGER_EVENT_PARSING_ERROR_CAUSE_USING_SHORT_CIRCUITING_IN_POST_EVENTS        = 10,

} DEBUGGER_EVENT_PARSING_ERROR_CAUSE,
    *PDEBUGGER_EVENT_PARSING_ERROR_CAUSE;

/**
 * @brief Maximum number of event handles in user-debugger
 */
#define DEBUGGER_MAXIMUM_SYNCRONIZATION_USER_DEBUGGER_OBJECTS 0x40

/**
 * @brief An event to show whether the debugger is running
 * or not in user-debugger
 *
 */

//
// User-debugger
//
#define DEBUGGER_SYNCRONIZATION_OBJECT_USER_DEBUGGER_IS_DEBUGGER_RUNNING 0x30

//////////////////////////////////////////////////
//            	   Event Details                //
//////////////////////////////////////////////////

/**
 * @brief In debugger holds the state of events
 *
 */
typedef struct _DEBUGGER_SYNCRONIZATION_EVENTS_STATE
{
    HANDLE  EventHandle;
    BOOLEAN IsOnWaitingState;
} DEBUGGER_SYNCRONIZATION_EVENTS_STATE, *PDEBUGGER_SYNCRONIZATION_EVENTS_STATE;

//////////////////////////////////////////////////
//				    Functions                   //
//////////////////////////////////////////////////

VOID
InterpreterRemoveComments(char * CommandText);

BOOLEAN
ShowErrorMessage(UINT32 Error);

BOOLEAN
IsConnectedToAnyInstanceOfDebuggerOrDebuggee();

BOOLEAN
IsTagExist(UINT64 Tag);

UINT64
DebuggerGetNtoskrnlBase();

BOOLEAN
DebuggerPauseDebuggee();

BOOLEAN
InterpretConditionsAndCodes(vector<string> * SplitCommand,
                            vector<string> * SplitCommandCaseSensitive,
                            BOOLEAN          IsConditionBuffer,
                            PUINT64          BufferAddress,
                            PUINT32          BufferLength);

VOID
FreeEventsAndActionsMemory(PDEBUGGER_GENERAL_EVENT_DETAIL Event,
                           PDEBUGGER_GENERAL_ACTION       ActionBreakToDebugger,
                           PDEBUGGER_GENERAL_ACTION       ActionCustomCode,
                           PDEBUGGER_GENERAL_ACTION       ActionScript);

BOOLEAN
SendEventToKernel(PDEBUGGER_GENERAL_EVENT_DETAIL Event,
                  UINT32                         EventBufferLength);

BOOLEAN
RegisterActionToEvent(PDEBUGGER_GENERAL_EVENT_DETAIL Event,
                      PDEBUGGER_GENERAL_ACTION       ActionBreakToDebugger,
                      UINT32                         ActionBreakToDebuggerLength,
                      PDEBUGGER_GENERAL_ACTION       ActionCustomCode,
                      UINT32                         ActionCustomCodeLength,
                      PDEBUGGER_GENERAL_ACTION       ActionScript,
                      UINT32                         ActionScriptLength);

BOOLEAN
InterpretGeneralEventAndActionsFields(
    vector<string> *                    SplitCommand,
    vector<string> *                    SplitCommandCaseSensitive,
    VMM_EVENT_TYPE_ENUM                 EventType,
    PDEBUGGER_GENERAL_EVENT_DETAIL *    EventDetailsToFill,
    PUINT32                             EventBufferLength,
    PDEBUGGER_GENERAL_ACTION *          ActionDetailsToFillBreakToDebugger,
    PUINT32                             ActionBufferLengthBreakToDebugger,
    PDEBUGGER_GENERAL_ACTION *          ActionDetailsToFillCustomCode,
    PUINT32                             ActionBufferLengthCustomCode,
    PDEBUGGER_GENERAL_ACTION *          ActionDetailsToFillScript,
    PUINT32                             ActionBufferLengthScript,
    PDEBUGGER_EVENT_PARSING_ERROR_CAUSE ReasonForErrorInParsing);

BOOLEAN
CallstackReturnAddressToCallingAddress(UCHAR * ReturnAddress, PUINT32 IndexOfCallFromReturnAddress);

VOID
CallstackShowFrames(PDEBUGGER_SINGLE_CALLSTACK_FRAME  CallstackFrames,
                    UINT32                            FrameCount,
                    DEBUGGER_CALLSTACK_DISPLAY_METHOD DisplayMethod,
                    BOOLEAN                           Is32Bit);

UINT64
GetNewDebuggerEventTag();

DWORD WINAPI
ListeningSerialPauseDebuggeeThread(PVOID Param);

DWORD WINAPI
ListeningSerialPauseDebuggerThread(PVOID Param);

VOID
LogopenSaveToFile(const char * Text);

BOOL
BreakController(DWORD CtrlType);

VOID
CommandEventsShowEvents();

BOOLEAN
CommandEventsModifyAndQueryEvents(UINT64                      Tag,
                                  DEBUGGER_MODIFY_EVENTS_TYPE TypeOfAction);

VOID
CommandEventsHandleModifiedEvent(
    UINT64                  Tag,
    PDEBUGGER_MODIFY_EVENTS ModifyEventRequest);

VOID
CommandEventsClearAllEventsAndResetTags();

VOID
CommandFlushRequestFlush();

UINT64
GetCommandAttributes(const string & FirstCommand);

VOID
DetachFromProcess();

BOOLEAN
CommandLoadVmmModule();

VOID
ShowAllRegisters();

VOID
CommandPauseRequest();

VOID
CommandGRequest();

VOID
CommandTrackHandleReceivedInstructions(unsigned char * BufferToDisassemble,
                                       UINT32          BuffLength,
                                       BOOLEAN         Isx86_64,
                                       UINT64          RipAddress);

VOID
CommandTrackHandleReceivedCallInstructions(const char * NameOfFunctionFromSymbols,
                                           UINT64       ComputedAbsoluteAddress);

VOID
CommandTrackHandleReceivedRetInstructions(UINT64 CurrentRip);


//..\bin\debug\headerAll\forwarding.h
/**
 * @file forwarding.h
 * @author Sina Karvandi (sina@hyperdbg.org)
 * @brief Headers for event source forwarding
 * @details
 * @version 0.1
 * @date 2020-11-16
 *
 * @copyright This project is released under the GNU Public License v3.
 *
 */
#pragma once

//////////////////////////////////////////
//           Forwarding Types           //
//////////////////////////////////////////

/**
 * @brief maximum characters for event forwarding source names
 *
 */
typedef void (*hyperdbg_event_forwarding_t)(const char *, unsigned int);

//////////////////////////////////////////
//     Output Source Forwarding         //
//////////////////////////////////////////

/**
 * @brief maximum characters for event forwarding source names
 *
 */
#define MAXIMUM_CHARACTERS_FOR_EVENT_FORWARDING_NAME 50

/**
 * @brief event forwarding type
 *
 */
typedef enum _DEBUGGER_EVENT_FORWARDING_TYPE
{
    EVENT_FORWARDING_NAMEDPIPE,
    EVENT_FORWARDING_FILE,
    EVENT_FORWARDING_TCP,
    EVENT_FORWARDING_MODULE,

} DEBUGGER_EVENT_FORWARDING_TYPE;

/**
 * @brief event forwarding states
 *
 */
typedef enum _DEBUGGER_EVENT_FORWARDING_STATE
{
    EVENT_FORWARDING_STATE_NOT_OPENED,
    EVENT_FORWARDING_STATE_OPENED,
    EVENT_FORWARDING_CLOSED,

} DEBUGGER_EVENT_FORWARDING_STATE;

/**
 * @brief output source status
 *
 * @details this enum is used as the result returned from
 * the functions that work with opening and closing sources
 */
typedef enum _DEBUGGER_OUTPUT_SOURCE_STATUS
{
    DEBUGGER_OUTPUT_SOURCE_STATUS_SUCCESSFULLY_OPENED,
    DEBUGGER_OUTPUT_SOURCE_STATUS_SUCCESSFULLY_CLOSED,
    DEBUGGER_OUTPUT_SOURCE_STATUS_ALREADY_OPENED,
    DEBUGGER_OUTPUT_SOURCE_STATUS_ALREADY_CLOSED,
    DEBUGGER_OUTPUT_SOURCE_STATUS_UNKNOWN_ERROR,

} DEBUGGER_OUTPUT_SOURCE_STATUS;

/**
 * @brief structures hold the detail of event forwarding
 *
 */
typedef struct _DEBUGGER_EVENT_FORWARDING
{
    DEBUGGER_EVENT_FORWARDING_TYPE  Type;
    DEBUGGER_EVENT_FORWARDING_STATE State;
    VOID *                          Handle;
    SOCKET                          Socket;
    HMODULE                         Module;
    UINT64                          OutputUniqueTag;
    LIST_ENTRY
    OutputSourcesList; // Linked-list of output sources list
    CHAR Name[MAXIMUM_CHARACTERS_FOR_EVENT_FORWARDING_NAME];

} DEBUGGER_EVENT_FORWARDING, *PDEBUGGER_EVENT_FORWARDING;

//////////////////////////////////////////
//              Functions	            //
//////////////////////////////////////////

UINT64
ForwardingGetNewOutputSourceTag();

DEBUGGER_OUTPUT_SOURCE_STATUS
ForwardingOpenOutputSource(PDEBUGGER_EVENT_FORWARDING SourceDescriptor);

DEBUGGER_OUTPUT_SOURCE_STATUS
ForwardingCloseOutputSource(PDEBUGGER_EVENT_FORWARDING SourceDescriptor);

BOOLEAN
ForwardingCheckAndPerformEventForwarding(UINT32 OperationCode,
                                         CHAR * Message,
                                         UINT32 MessageLength);

BOOLEAN
ForwardingWriteToFile(HANDLE FileHandle, CHAR * Message, UINT32 MessageLength);

BOOLEAN
ForwardingSendToNamedPipe(HANDLE NamedPipeHandle, CHAR * Message, UINT32 MessageLength);

BOOLEAN
ForwardingSendToTcpSocket(SOCKET TcpSocket, CHAR * Message, UINT32 MessageLength);

VOID *
ForwardingCreateOutputSource(DEBUGGER_EVENT_FORWARDING_TYPE SourceType,
                             const string &                 Description,
                             SOCKET *                       Socket,
                             HMODULE *                      Module);


//..\bin\debug\headerAll\globals.h
/**
 * @file globals.h
 * @author Sina Karvandi (sina@hyperdbg.org)
 * @brief Global Variables for user-mode interface
 * @details
 * @version 0.1
 * @date 2020-07-13
 *
 * @copyright This project is released under the GNU Public License v3.
 *
 */
#pragma once

//////////////////////////////////////////////////
//	            Feature Indicators              //
//////////////////////////////////////////////////

/**
 * @brief check for RTM support
 *
 */
BOOLEAN g_RtmSupport = FALSE;

/**
 * @brief Virtual address width for x86 processors
 *
 */
UINT32 g_VirtualAddressWidth = 0;

//////////////////////////////////////////////////
//	         Interpreter Variables              //
//////////////////////////////////////////////////

/**
 * @brief shows whether the interpreter is currently on a string or not
 */
BOOLEAN g_IsInterpreterOnString = FALSE;

/**
 * @brief Is interpreter encountered a back slash at previous run
 */
BOOLEAN g_IsInterpreterPreviousCharacterABackSlash = FALSE;

/**
 * @brief Keeps the trace of curly brackets in the interpreter
 */
UINT32 g_InterpreterCountOfOpenCurlyBrackets = 0;

//////////////////////////////////////////////////
//		 Remote and Local Connection            //
//////////////////////////////////////////////////

/**
 * @brief the buffer that we set at the end of buffers for tcp connection
 */
BYTE g_EndOfBufferCheckTcp[TCP_END_OF_BUFFER_CHARS_COUNT] = {
    TCP_END_OF_BUFFER_CHAR_1,
    TCP_END_OF_BUFFER_CHAR_2,
    TCP_END_OF_BUFFER_CHAR_3,
    TCP_END_OF_BUFFER_CHAR_4};

/**
 * @brief Shows whether the user is allowed to use 'load' command
 * to load modules locally in VMI (virtual machine introspection) mode
 *
 */
BOOLEAN g_IsConnectedToHyperDbgLocally = FALSE;

/**
 * @brief Shows whether the current debugger is the host and
 * connected to a remote debuggee (guest)
 *
 */
BOOLEAN g_IsConnectedToRemoteDebuggee = FALSE;

/**
 * @brief  Shows whether the current system is a guest (debuggee)
 * and a remote debugger is connected to this system
 *
 */
BOOLEAN g_IsConnectedToRemoteDebugger = FALSE;

/**
 * @brief The socket object of host debugger (not debuggee)
 * it is because in HyperDbg, debuggee is server and debugger
 * is a client
 *
 */
SOCKET g_ClientConnectSocket = {0};

/**
 * @brief The socket object of guest debuggee (not debugger)
 * it is because in HyperDbg, debugger is client and debuggee
 * is a server
 *
 */
SOCKET g_SeverSocket = {0};

/**
 * @brief Server in debuggee needs an extra socket
 *
 */
SOCKET g_ServerListenSocket = {0};

/**
 * @brief In debugger (not debuggee), we save the port of server
 * debuggee in this variable to use it later e.g, in signature
 *
 */
string g_ServerPort = "";

/**
 * @brief In debugger (not debuggee), we save the port of server
 * debuggee in this variable to use it later e.g, in signature
 *
 */
string g_ServerIp = "";

/**
 * @brief In debugger (not debuggee), we save the ip of server
 * debuggee in this variable to use it later e.g, in signature
 *
 */
HANDLE g_RemoteDebuggeeListeningThread = NULL;

/**
 * @brief Handle to show that if the debugger is loaded successfully
 *
 */
HANDLE g_IsDriverLoadedSuccessfully = NULL;

/**
 * @brief Handle to if the end of the message received (for showing
 * signature)
 *
 */
HANDLE g_EndOfMessageReceivedEvent = NULL;

/**
 * @brief variable to keep track if the end of the message received
 * (for showing signature)
 *
 */
BOOLEAN g_IsEndOfMessageReceived = FALSE;

/**
 * @brief In both debuggee and debugger we save the state of
 * the closed connection to avoid double close
 *
 */
BOOLEAN g_SerialConnectionAlreadyClosed = FALSE;

/**
 * @brief Show whether the pause request (CTRL+C or CTRL+BREAK)
 * should be ignored or not
 *
 */
BOOLEAN g_IgnorePauseRequests = FALSE;

//////////////////////////////////////////////////
//		 User Debugging Variables             //
//////////////////////////////////////////////////

/**
 * @brief Whether the user debugger is initialized or not
 */
BOOLEAN g_IsUserDebuggerInitialized = FALSE;

/**
 * @brief In debugger (not debuggee), we save the handle
 * of the user-mode listening thread for pauses here for user debugger
 *
 */

DEBUGGER_SYNCRONIZATION_EVENTS_STATE g_UserSyncronizationObjectsHandleTable
    [DEBUGGER_MAXIMUM_SYNCRONIZATION_USER_DEBUGGER_OBJECTS] = {0};

//////////////////////////////////////////////////
//		 Serial Debugging Variables             //
//////////////////////////////////////////////////

/**
 * @brief the buffer that we set at the end of buffers for serial
 */
BYTE g_EndOfBufferCheckSerial[SERIAL_END_OF_BUFFER_CHARS_COUNT] = {
    SERIAL_END_OF_BUFFER_CHAR_1,
    SERIAL_END_OF_BUFFER_CHAR_2,
    SERIAL_END_OF_BUFFER_CHAR_3,
    SERIAL_END_OF_BUFFER_CHAR_4};

/**
 * @brief In debugger (not debuggee), we save the handle
 * of the user-mode listening thread for pauses here for kernel debugger
 *
 */

DEBUGGER_SYNCRONIZATION_EVENTS_STATE g_KernelSyncronizationObjectsHandleTable
    [DEBUGGER_MAXIMUM_SYNCRONIZATION_KERNEL_DEBUGGER_OBJECTS] = {0};

/**
 * @brief Current executing instructions
 *
 */
BYTE g_CurrentRunningInstruction[MAXIMUM_INSTR_SIZE] = {0};

/**
 * @brief whether the Current executing instructions is 32-bit or 64 bit
 *
 */
BOOLEAN g_IsRunningInstruction32Bit = FALSE;

/**
 * @brief In debuggee and debugger, we save the handle
 * of the user-mode listening thread for pauses here
 *
 */
HANDLE g_SerialListeningThreadHandle = NULL;

/**
 * @brief In debugger (not debuggee), we save the handle
 * of the user-mode listening thread for remote system here
 *
 */
HANDLE g_SerialRemoteComPortHandle = NULL;

/**
 * @brief Shows if the debugger was connected to
 * remote debuggee over (A remote guest)
 *
 */
BOOLEAN g_IsSerialConnectedToRemoteDebuggee = FALSE;

/**
 * @brief Shows if the debugger was connected to
 * remote debugger (A remote host)
 *
 */
BOOLEAN g_IsSerialConnectedToRemoteDebugger = FALSE;

/**
 * @brief Shows if the debuggee is in the handshake phase or not
 *
 */
BOOLEAN g_IsDebuggeeInHandshakingPhase = FALSE;

/**
 * @brief Shows if the debuggee is running or not
 *
 */
BOOLEAN g_IsDebuggeeRunning = FALSE;

/**
 * @brief Shows if the debugger should show debuggee's messages
 * or not
 *
 */
BOOLEAN g_IgnoreNewLoggingMessages = FALSE;

/**
 * @brief Current core that the debuggee is debugging
 *
 */
ULONG g_CurrentRemoteCore = DEBUGGER_DEBUGGEE_IS_RUNNING_NO_CORE;

/**
 * @brief Shows if the debugger is connected to the
 * guest using named pipe
 *
 */
BOOLEAN g_IsDebuggerConntectedToNamedPipe = FALSE;

/**
 * @brief An event to make sure that the user won't give any command in debuggee
 * and all the commands are coming from just the debugger
 *
 */
HANDLE g_DebuggeeStopCommandEventHandle = NULL;

/**
 * @brief Holds the result of registering events from the remote debuggee
 *
 */
DEBUGGER_EVENT_AND_ACTION_RESULT g_DebuggeeResultOfRegisteringEvent = {0};

/**
 * @brief Holds the result of adding action to events from the remote debuggee
 *
 */
DEBUGGER_EVENT_AND_ACTION_RESULT g_DebuggeeResultOfAddingActionsToEvent = {
    0};

/**
 * @brief This is an OVERLAPPED structure for managing simultaneous
 * read and writes for debugger (in current design debuggee is not needed
 * to write simultaneously but it's needed for write)
 *
 */
OVERLAPPED g_OverlappedIoStructureForReadDebugger  = {0};
OVERLAPPED g_OverlappedIoStructureForWriteDebugger = {0};

OVERLAPPED g_OverlappedIoStructureForReadDebuggee = {0};

/**
 * @brief Shows whether the queried event is enabled or disabled
 *
 */
BOOLEAN g_SharedEventStatus = FALSE;

//////////////////////////////////////////////////
//				 Global Variables               //
//////////////////////////////////////////////////

/**
 * @brief Shows whether the previous command should be
 * continued or not
 *
 */
BOOLEAN g_ShouldPreviousCommandBeContinued;

/**
 * @brief List of command and attributes
 *
 */
CommandType g_CommandsList;

/**
 * @brief Holder of global variables for script engine
 *
 */
UINT64 * g_ScriptGlobalVariables;

/**
 * @brief Holder of local variables for script engine
 *
 */
UINT64 * g_ScriptLocalVariables;

/**
 * @brief Holder of temp variables for script engine
 *
 */
UINT64 * g_ScriptTempVariables;

/**
 * @brief Is list of command initialized
 *
 */
BOOLEAN g_IsCommandListInitialized = FALSE;

/**
 * @brief this variable is used to indicate that modules
 * are loaded so we make sure to later use a trace of
 * loading in 'unload' command (used in Debugger VMM)
 *
 */
BOOLEAN g_IsDebuggerModulesLoaded = FALSE;

/**
 * @brief State of active debugging thread
 *
 */
ACTIVE_DEBUGGING_PROCESS g_ActiveProcessDebuggingState = {0};

/**
 * @brief The process id of the latest starting process
 *
 */
UINT32 g_ProcessIdOfLatestStartingProcess = NULL;

/**
 * @brief This variable holds the trace and generate numbers
 * for new tags of events
 *
 */
UINT64 g_EventTag = DebuggerEventTagStartSeed;

/**
 * @brief This variable holds the trace and generate numbers
 * for unique tag of the output resources
 *
 */
UINT64 g_OutputSourceTag = DebuggerOutputSourceTagStartSeed;

/**
 * @brief it shows whether the debugger started using
 * events or not or in other words, is g_EventTrace
 * initialized with a variable or it is empty
 *
 */
BOOLEAN g_EventTraceInitialized = FALSE;

/**
 * @brief Holds a list of events in kernel and the state of events
 * and the commands to show the state of each command (disabled/enabled)
 *
 * @details this list is not have any relation with the things that HyperDbg
 * holds for each event in the kernel
 *
 */
LIST_ENTRY g_EventTrace = {0};

/**
 * @brief it shows whether the debugger started using
 * output sources or not or in other words, is g_OutputSources
 * initialized with a variable or it is empty
 *
 */
BOOLEAN g_OutputSourcesInitialized = FALSE;

/**
 * @brief Holds a list of output sources created by output command
 *
 * @details user-mode events and output sources are two separate things
 * in HyperDbg
 *
 */
LIST_ENTRY g_OutputSources = {0};

/**
 * @brief Holds the location driver to install it
 *
 */
TCHAR g_DriverLocation[MAX_PATH] = {0};

/**
 * @brief Holds the location test-hyperdbg.exe
 *
 */
TCHAR g_TestLocation[MAX_PATH] = {0};

/**
 * @brief The handler for ShowMessages function
 * this is because the user might choose not to use
 * printf and instead use his/her handler for showing
 * messages
 *
 */
Callback g_MessageHandler = 0;

/**
 * @brief Shows whether the vmxoff process start or not
 *
 */
BOOLEAN g_IsVmxOffProcessStart;

/**
 * @brief Holds the global handle of device which is used
 * to send the request to the kernel by IOCTL, this
 * handle is not used for IRP Pending of message tracing
 * this handle is used in KD VMM
 *
 */
HANDLE g_DeviceHandle;

/**
 * @brief Shows whether the '.logopen' command is executed
 * and the log file is open or not
 *
 */
BOOLEAN g_LogOpened = FALSE;

/**
 * @brief The object of log file ('.logopen' command)
 *
 */
ofstream g_LogOpenFile;

/**
 * @brief Shows whether the target is executing a script
 * form '.script' command or executing script by an
 * argument
 *
 */
BOOLEAN g_ExecutingScript = FALSE;

/**
 * @brief Shows whether the pause command or CTRL+C
 * or CTRL+Break is executed or not
 *
 */
BOOLEAN g_BreakPrintingOutput = FALSE;

/**
 * @brief Executing symbol reloading or downloading
 * routines
 *
 */
BOOLEAN g_IsExecutingSymbolLoadingRoutines = FALSE;

/**
 * @brief Symbol table for disassembler
 *
 */
std::map<UINT64, LOCAL_FUNCTION_DESCRIPTION> g_DisassemblerSymbolMap;

/**
 * @brief Shows whether the user executed and mesaured '!measure'
 * command or not, it is because we want to use these measurements
 * later in '!hide' command
 *
 */
BOOLEAN g_TransparentResultsMeasured = FALSE;

/**
 * @brief The average calculated from the measurements of cpuid
 * '!measure' command
 */
UINT64 g_CpuidAverage = 0;

/**
 * @brief The standard deviation calculated from the measurements of cpuid
 * '!measure' command
 */
UINT64 g_CpuidStandardDeviation = 0;

/**
 * @brief The median calculated from the measurements of cpuid
 * '!measure' command
 */
UINT64 g_CpuidMedian = 0;

/**
 * @brief The average calculated from the measurements of rdtsc/p
 * '!measure' command
 */
UINT64 g_RdtscAverage = 0;

/**
 * @brief The standard deviation calculated from the measurements
 *  of rdtsc/p '!measure' command
 */
UINT64 g_RdtscStandardDeviation = 0;

/**
 * @brief The median calculated from the measurements of rdtsc/p
 * '!measure' command
 */
UINT64 g_RdtscMedian = 0;

/**
 * @brief Shows whether the user is running 't', 'p', or 'i' command
 */
BOOLEAN g_IsInstrumentingInstructions = FALSE;

//////////////////////////////////////////////////
//			     	 Settings			        //
//////////////////////////////////////////////////

/**
 * @brief Whether auto-unpause mode is enabled or not enabled
 * @details it is enabled by default
 *
 */
BOOLEAN g_AutoUnpause = TRUE;

/**
 * @brief Whether converting addresses to object names or not
 * @details it is enabled by default
 *
 */
BOOLEAN g_AddressConversion = TRUE;

/**
 * @brief Whether auto-flush mode is enabled or not enabled
 * @details it is disabled by default
 *
 */
BOOLEAN g_AutoFlush = FALSE;

/**
 * @brief Shows the syntax used in !u !u2 u u2 commands
 * @details INTEL = 1, ATT = 2, MASM = 3
 *
 */
UINT32 g_DisassemblerSyntax = 1;

//////////////////////////////////////////////////
//			   	 Symbol Table			        //
//////////////////////////////////////////////////

/**
 * @brief The buffer that stores the details of
 * symbol table
 *
 */
PMODULE_SYMBOL_DETAIL g_SymbolTable = NULL;

/**
 * @brief The buffer that stores size of the
 * details of symbol table
 *
 */
UINT32 g_SymbolTableSize = NULL;

/**
 * @brief The index to hold the track of
 * added symbols
 *
 */
UINT32 g_SymbolTableCurrentIndex = NULL;

/**
 * @brief Result of the expression that is evaluated in the
 * debuggee
 *
 */
UINT64 g_ResultOfEvaluatedExpression = NULL;

/**
 * @brief Shows the state of the evaluation of expression which
 * whether contains error or not
 *
 */
UINT32 g_ErrorStateOfResultOfEvaluatedExpression = NULL;

//////////////////////////////////////////////////
//			 User mode Debugging		        //
//////////////////////////////////////////////////

/**
 * @brief the start path used in .start command
 *
 */
std::wstring g_StartCommandPath = L"";

/**
 * @brief the start arguments used in .start command
 *
 */
std::wstring g_StartCommandPathAndArguments = L"";

//////////////////////////////////////////////////
//			 Script engine tests		        //
//////////////////////////////////////////////////

/**
 * @brief global variable to save the result of script-engine statement
 * tests
 *
 */
UINT64 g_CurrentExprEvalResult;

/**
 * @brief global variable to detect if there was an error in the result
 *  of script-engine statement tests
 *
 */
BOOLEAN g_CurrentExprEvalResultHasError;

//////////////////////////////////////////////////
//				      hwdbg                     //
//////////////////////////////////////////////////

/**
 * @brief Instance information of the current hwdbg debuggee
 *
 */
HWDBG_INSTANCE_INFORMATION g_HwdbgInstanceInfo;

/**
 * @brief Shows whether the instance info is valid (received) or not
 *
 */
BOOLEAN g_HwdbgInstanceInfoIsValid;

/**
 * @brief Ports configuration of hwdbg
 *
 */
std::vector<UINT32> g_HwdbgPortConfiguration;


//..\bin\debug\headerAll\help.h
/**
 * @file help.h
 * @author Sina Karvandi (sina@hyperdbg.org)
 * @brief help of commands header
 * @details
 * @version 0.1
 * @date 2020-05-27
 *
 * @copyright This project is released under the GNU Public License v3.
 *
 */
#pragma once

//////////////////////////////////////////////////
//					Help commands               //
//////////////////////////////////////////////////

VOID
CommandHelpHelp();

VOID
CommandReadMemoryAndDisassemblerHelp();

VOID
CommandConnectHelp();

VOID
CommandDisconnectHelp();

VOID
CommandExitHelp();

VOID
CommandCpuHelp();

VOID
CommandUnloadHelp();

VOID
CommandLoadHelp();

VOID
CommandConnectHelp();

VOID
CommandScriptHelp();

VOID
CommandFormatsHelp();

VOID
CommandRdmsrHelp();

VOID
CommandWrmsrHelp();

VOID
CommandPteHelp();

VOID
CommandMonitorHelp();

VOID
CommandSyscallHelp();

VOID
CommandSysretHelp();

VOID
CommandEptHookHelp();

VOID
CommandEptHook2Help();

VOID
CommandCpuidHelp();

VOID
CommandMsrreadHelp();

VOID
CommandMsrwriteHelp();

VOID
CommandTscHelp();

VOID
CommandPmcHelp();

VOID
CommandExceptionHelp();

VOID
CommandCrwriteHelp();

VOID
CommandDrHelp();

VOID
CommandInterruptHelp();

VOID
CommandIooutHelp();

VOID
CommandIoinHelp();

VOID
CommandVmcallHelp();

VOID
CommandModeHelp();

VOID
CommandTraceHelp();

VOID
CommandHideHelp();

VOID
CommandUnhideHelp();

VOID
CommandTestHelp();

VOID
CommandLogopenHelp();

VOID
CommandLogcloseHelp();

VOID
CommandVa2paHelp();

VOID
CommandPa2vaHelp();

VOID
CommandEventsHelp();

VOID
CommandGHelp();

VOID
CommandClearScreenHelp();

VOID
CommandSleepHelp();

VOID
CommandEditMemoryHelp();

VOID
CommandSearchMemoryHelp();

VOID
CommandMeasureHelp();

VOID
CommandLmHelp();

VOID
CommandSettingsHelp();

VOID
CommandFlushHelp();

VOID
CommandPauseHelp();

VOID
CommandListenHelp();

VOID
CommandStatusHelp();

VOID
CommandAttachHelp();

VOID
CommandDetachHelp();

VOID
CommandStartHelp();

VOID
CommandRestartHelp();

VOID
CommandSwitchHelp();

VOID
CommandKillHelp();

VOID
CommandTHelp();

VOID
CommandIHelp();

VOID
CommandPrintHelp();

VOID
CommandOutputHelp();

VOID
CommandDebugHelp();

VOID
CommandPHelp();

VOID
CommandCoreHelp();

VOID
CommandProcessHelp();

VOID
CommandThreadHelp();

VOID
CommandEvalHelp();

VOID
CommandRHelp();

VOID
CommandBpHelp();

VOID
CommandBlHelp();

VOID
CommandBeHelp();

VOID
CommandBdHelp();

VOID
CommandBcHelp();

VOID
CommandSympathHelp();

VOID
CommandSymHelp();

VOID
CommandXHelp();

VOID
CommandPreallocHelp();

VOID
CommandPreactivateHelp();

VOID
CommandDtHelp();

VOID
CommandStructHelp();

VOID
CommandKHelp();

VOID
CommandPeHelp();

VOID
CommandRevHelp();

VOID
CommandTrackHelp();

VOID
CommandPageinHelp();

VOID
CommandDumpHelp();

VOID
CommandGuHelp();

//
// hwdbg commands
//
VOID
CommandHwClkHelp();


//..\bin\debug\headerAll\hwdbg-interpreter.h
/**
 * @file hwdbg-interpreter.h
 * @author Sina Karvandi (sina@hyperdbg.org)
 * @brief Headers for the interpreter of hwdbg packets and requests
 * @details
 * @version 1.0
 * @date 2024-06-11
 *
 * @copyright This project is released under the GNU Public License v3.
 *
 */
#pragma once

//////////////////////////////////////////////////
//				   Definitions                  //
//////////////////////////////////////////////////

/**
 * @brief Path to read the sample of the instance info
 *
 */
#define HWDBG_TEST_READ_INSTANCE_INFO_PATH "..\\..\\..\\..\\hwdbg\\sim\\hwdbg\\DebuggerModuleTestingBRAM\\bram_instance_info.txt"

/**
 * @brief Path to write the sample of the script buffer
 *
 */
#define HWDBG_TEST_WRITE_SCRIPT_BUFFER_PATH "..\\..\\..\\..\\hwdbg\\src\\test\\bram\\script_buffer.hex.txt"

/**
 * @brief Path to write the sample of the instance info requests
 *
 */
#define HWDBG_TEST_WRITE_INSTANCE_INFO_PATH "..\\..\\..\\..\\hwdbg\\src\\test\\bram\\instance_info.hex.txt"

//////////////////////////////////////////////////
//				    Functions                   //
//////////////////////////////////////////////////

BOOLEAN
HwdbgInterpretPacket(PVOID BufferReceived, UINT32 LengthReceived);

BOOLEAN
HwdbgInterpreterFillFileFromMemory(
    HWDBG_INSTANCE_INFORMATION * InstanceInfo,
    const TCHAR *                FileName,
    UINT32 *                     MemoryBuffer,
    size_t                       BufferSize,
    HWDBG_ACTION_ENUMS           RequestedAction);

BOOLEAN
HwdbgInterpreterFillMemoryFromFile(const TCHAR * FileName, UINT32 * MemoryBuffer, size_t BufferSize);

BOOLEAN
HwdbgInterpreterConvertSymbolToHwdbgShortSymbolBuffer(
    HWDBG_INSTANCE_INFORMATION * InstanceInfo,
    SYMBOL *                     SymbolBuffer,
    size_t                       SymbolBufferLength,
    UINT32                       NumberOfStages,
    HWDBG_SHORT_SYMBOL **        NewShortSymbolBuffer,
    size_t *                     NewBufferSize);

BOOLEAN
HwdbgInterpreterCompressBuffer(UINT64 * Buffer,
                               size_t   BufferLength,
                               UINT32   ScriptVariableLength,
                               UINT32   BramDataWidth,
                               size_t * NewBufferSize,
                               size_t * NumberOfBytesPerChunk);

VOID
HwdbgInterpreterShowScriptCapabilities(HWDBG_INSTANCE_INFORMATION * InstanceInfo);

BOOLEAN
HwdbgInterpreterCheckScriptBufferWithScriptCapabilities(HWDBG_INSTANCE_INFORMATION * InstanceInfo,
                                                        PVOID                        ScriptBuffer,
                                                        UINT32                       CountOfScriptSymbolChunks,
                                                        UINT32 *                     NumberOfStages,
                                                        UINT32 *                     NumberOfOperands);

BOOLEAN
HwdbgInterpreterSendPacketAndBufferToHwdbg(HWDBG_INSTANCE_INFORMATION * InstanceInfo,
                                           const TCHAR *                FileName,
                                           DEBUGGER_REMOTE_PACKET_TYPE  PacketType,
                                           HWDBG_ACTION_ENUMS           RequestedAction,
                                           CHAR *                       Buffer,
                                           UINT32                       BufferLength);

BOOLEAN
HwdbgInterpreterSendScriptPacket(HWDBG_INSTANCE_INFORMATION * InstanceInfo,
                                 const TCHAR *                FileName,
                                 UINT32                       NumberOfSymbols,
                                 HWDBG_SHORT_SYMBOL *         Buffer,
                                 UINT32                       BufferLength);


//..\bin\debug\headerAll\inipp.h
/*
MIT License

Copyright (c) 2017-2020 Matthias C. M. Troffaes

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

#pragma once

#include <algorithm>
#include <cctype>
#include <cstring>
#include <functional>
#include <iostream>
#include <list>
#include <vector>
#include <locale>
#include <map>
#include <memory>
#include <sstream>
#include <string>

namespace inipp {

namespace detail {

// trim functions based on http://stackoverflow.com/a/217605

template <class CharT>
inline void
ltrim(std::basic_string<CharT> & s, const std::locale & loc)
{
    s.erase(s.begin(),
            std::find_if(s.begin(), s.end(), [&loc](CharT ch) { return !std::isspace(ch, loc); }));
}

template <class CharT>
inline void
rtrim(std::basic_string<CharT> & s, const std::locale & loc)
{
    s.erase(std::find_if(s.rbegin(), s.rend(), [&loc](CharT ch) { return !std::isspace(ch, loc); }).base(),
            s.end());
}

template <class CharT, class UnaryPredicate>
inline void
rtrim2(std::basic_string<CharT> & s, UnaryPredicate pred)
{
    s.erase(std::find_if(s.begin(), s.end(), pred), s.end());
}

// string replacement function based on http://stackoverflow.com/a/3418285

template <class CharT>
inline bool
replace(std::basic_string<CharT> & str, const std::basic_string<CharT> & from, const std::basic_string<CharT> & to)
{
    auto   changed   = false;
    size_t start_pos = 0;
    while ((start_pos = str.find(from, start_pos)) != std::basic_string<CharT>::npos)
    {
        str.replace(start_pos, from.length(), to);
        start_pos += to.length();
        changed = true;
    }
    return changed;
}

} // namespace detail

template <typename CharT, typename T>
inline bool
extract(const std::basic_string<CharT> & value, T & dst)
{
    CharT                           c;
    std::basic_istringstream<CharT> is {value};
    T                               result;
    if ((is >> std::boolalpha >> result) && !(is >> c))
    {
        dst = result;
        return true;
    }
    else
    {
        return false;
    }
}

template <typename CharT>
inline bool
extract(const std::basic_string<CharT> & value, std::basic_string<CharT> & dst)
{
    dst = value;
    return true;
}

template <typename CharT, typename T>
inline bool
get_value(const std::map<std::basic_string<CharT>, std::basic_string<CharT>> & sec, const std::basic_string<CharT> & key, T & dst)
{
    const auto it = sec.find(key);
    if (it == sec.end())
        return false;
    return extract(it->second, dst);
}

template <typename CharT, typename T>
inline bool
get_value(const std::map<std::basic_string<CharT>, std::basic_string<CharT>> & sec, const CharT * key, T & dst)
{
    return get_value(sec, std::basic_string<CharT>(key), dst);
}

template <class CharT>
class Format
{
public:
    // used for generating
    const CharT char_section_start;
    const CharT char_section_end;
    const CharT char_assign;
    const CharT char_comment;

    // used for parsing
    virtual bool is_section_start(CharT ch) const { return ch == char_section_start; }
    virtual bool is_section_end(CharT ch) const { return ch == char_section_end; }
    virtual bool is_assign(CharT ch) const { return ch == char_assign; }
    virtual bool is_comment(CharT ch) const { return ch == char_comment; }

    // used for interpolation
    const CharT char_interpol;
    const CharT char_interpol_start;
    const CharT char_interpol_sep;
    const CharT char_interpol_end;

    Format(CharT section_start, CharT section_end, CharT assign, CharT comment, CharT interpol, CharT interpol_start, CharT interpol_sep, CharT interpol_end) :
        char_section_start(section_start), char_section_end(section_end), char_assign(assign), char_comment(comment), char_interpol(interpol), char_interpol_start(interpol_start), char_interpol_sep(interpol_sep), char_interpol_end(interpol_end) { }

    Format() :
        Format('[', ']', '=', ';', '$', '{', ':', '}') { }

    const std::basic_string<CharT> local_symbol(const std::basic_string<CharT> & name) const
    {
        return char_interpol + (char_interpol_start + name + char_interpol_end);
    }

    const std::basic_string<CharT> global_symbol(const std::basic_string<CharT> & sec_name, const std::basic_string<CharT> & name) const
    {
        return local_symbol(sec_name + char_interpol_sep + name);
    }
};

template <class CharT>
class Ini
{
public:
    using String   = std::basic_string<CharT>;
    using Section  = std::map<String, String>;
    using Sections = std::map<String, Section>;

    Sections                       sections;
    std::list<String>              errors;
    std::shared_ptr<Format<CharT>> format;

    static const int max_interpolation_depth = 10;

    Ini() :
        format(std::make_shared<Format<CharT>>()) {};
    Ini(std::shared_ptr<Format<CharT>> fmt) :
        format(fmt) {};

    void generate(std::basic_ostream<CharT> & os) const
    {
        for (auto const & sec : sections)
        {
            os << format->char_section_start << sec.first << format->char_section_end << std::endl;
            for (auto const & val : sec.second)
            {
                os << val.first << format->char_assign << val.second << std::endl;
            }
            os << std::endl;
        }
    }

    void parse(std::basic_istream<CharT> & is)
    {
        String            line;
        String            section;
        const std::locale loc {"C"};
        while (std::getline(is, line))
        {
            detail::ltrim(line, loc);
            detail::rtrim(line, loc);
            const auto length = line.length();
            if (length > 0)
            {
                const auto   pos   = std::find_if(line.begin(), line.end(), [this](CharT ch) { return format->is_assign(ch); });
                const auto & front = line.front();
                if (format->is_comment(front))
                {
                }
                else if (format->is_section_start(front))
                {
                    if (format->is_section_end(line.back()))
                        section = line.substr(1, length - 2);
                    else
                        errors.push_back(line);
                }
                else if (pos != line.begin() && pos != line.end())
                {
                    String variable(line.begin(), pos);
                    String value(pos + 1, line.end());
                    detail::rtrim(variable, loc);
                    detail::ltrim(value, loc);
                    auto & sec = sections[section];
                    if (sec.find(variable) == sec.end())
                        sec.emplace(variable, value);
                    else
                        errors.push_back(line);
                }
                else
                {
                    errors.push_back(line);
                }
            }
        }
    }

    void interpolate()
    {
        int  global_iteration = 0;
        auto changed          = false;
        // replace each "${variable}" by "${section:variable}"
        for (auto & sec : sections)
            replace_symbols(local_symbols(sec.first, sec.second), sec.second);
        // replace each "${section:variable}" by its value
        do
        {
            changed         = false;
            const auto syms = global_symbols();
            for (auto & sec : sections)
                changed |= replace_symbols(syms, sec.second);
        } while (changed && (max_interpolation_depth > global_iteration++));
    }

    void default_section(const Section & sec)
    {
        for (auto & sec2 : sections)
            for (const auto & val : sec)
                sec2.second.insert(val);
    }

    void strip_trailing_comments()
    {
        const std::locale loc {"C"};
        for (auto & sec : sections)
            for (auto & val : sec.second)
            {
                detail::rtrim2(val.second, [this](CharT ch) { return format->is_comment(ch); });
                detail::rtrim(val.second, loc);
            }
    }

    void clear()
    {
        sections.clear();
        errors.clear();
    }

private:
    using Symbols = std::vector<std::pair<String, String>>;

    const Symbols local_symbols(const String & sec_name, const Section & sec) const
    {
        Symbols result;
        for (const auto & val : sec)
            result.emplace_back(format->local_symbol(val.first), format->global_symbol(sec_name, val.first));
        return result;
    }

    const Symbols global_symbols() const
    {
        Symbols result;
        for (const auto & sec : sections)
            for (const auto & val : sec.second)
                result.emplace_back(format->global_symbol(sec.first, val.first), val.second);
        return result;
    }

    bool replace_symbols(const Symbols & syms, Section & sec) const
    {
        auto changed = false;
        for (auto & sym : syms)
            for (auto & val : sec)
                changed |= detail::replace(val.second, sym.first, sym.second);
        return changed;
    }
};

} // namespace inipp


//..\bin\debug\headerAll\install.h
/**
 * @file hprdbgctrl.h
 * @author Sina Karvandi (sina@hyperdbg.org)
 * @brief Main interface to connect applications to driver headers
 * @details
 * @version 0.1
 * @date 2020-04-11
 *
 * @copyright This project is released under the GNU Public License v3.
 *
 */
#pragma once

//
// The following ifdef block is the standard way of creating macros which make
// exporting from a DLL simpler. All files within this DLL are compiled with the
// HPRDBGCTRL_EXPORTS symbol defined on the command line. This symbol should not
// be defined on any project that uses this DLL. This way any other project
// whose source files include this file see HPRDBGCTRL_API functions as being
// imported from a DLL, whereas this DLL sees symbols defined with this macro as
// being exported.
//

#ifdef HPRDBGCTRL_EXPORTS
#    define HPRDBGCTRL_API __declspec(dllexport)
#else
#    define HPRDBGCTRL_API __declspec(dllimport)
#endif

//////////////////////////////////////////////////
//					Installer                   //
//////////////////////////////////////////////////

#define DRIVER_FUNC_INSTALL 0x01
#define DRIVER_FUNC_STOP    0x02
#define DRIVER_FUNC_REMOVE  0x03

BOOLEAN
ManageDriver(_In_ LPCTSTR DriverName, _In_ LPCTSTR ServiceName, _In_ UINT16 Function);

BOOLEAN
SetupPathForFileName(const CHAR *                                  FileName,
                     _Inout_updates_bytes_all_(BufferLength) PCHAR FileLocation,
                     ULONG                                         BufferLength,
                     BOOLEAN                                       CheckFileExists);


//..\bin\debug\headerAll\kd.h
/**
 * @file kd.h
 * @author Sina Karvandi (sina@hyperdbg.org)
 * @brief routines for remote kernel debugging
 * @details
 * @version 0.1
 * @date 2020-12-22
 *
 * @copyright This project is released under the GNU Public License v3.
 *
 */
#pragma once

//////////////////////////////////////////////////
//		            Definitions                 //
//////////////////////////////////////////////////

#define DbgWaitForKernelResponse(KernelSyncObjectId)                       \
    do                                                                     \
    {                                                                      \
        DEBUGGER_SYNCRONIZATION_EVENTS_STATE * SyncronizationObject =      \
            &g_KernelSyncronizationObjectsHandleTable[KernelSyncObjectId]; \
                                                                           \
        SyncronizationObject->IsOnWaitingState = TRUE;                     \
        WaitForSingleObject(SyncronizationObject->EventHandle, INFINITE);  \
    } while (FALSE);

#define DbgReceivedKernelResponse(KernelSyncObjectId)                      \
    do                                                                     \
    {                                                                      \
        DEBUGGER_SYNCRONIZATION_EVENTS_STATE * SyncronizationObject =      \
            &g_KernelSyncronizationObjectsHandleTable[KernelSyncObjectId]; \
                                                                           \
        SyncronizationObject->IsOnWaitingState = FALSE;                    \
        SetEvent(SyncronizationObject->EventHandle);                       \
    } while (FALSE);

//////////////////////////////////////////////////
//		    Display Windows Details             //
//////////////////////////////////////////////////

struct HKeyHolder
{
private:
    HKEY m_Key;

public:
    HKeyHolder() :
        m_Key(nullptr) { }

    HKeyHolder(const HKeyHolder &)             = delete;
    HKeyHolder & operator=(const HKeyHolder &) = delete;

    ~HKeyHolder()
    {
        if (m_Key != nullptr)
            RegCloseKey(m_Key);
    }

    operator HKEY() const { return m_Key; }

    HKEY * operator&() { return &m_Key; }
};

//////////////////////////////////////////////////
//			    	 Functions                  //
//////////////////////////////////////////////////

BOOLEAN
KdCommandPacketToDebuggee(
    DEBUGGER_REMOTE_PACKET_TYPE             PacketType,
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION RequestedAction);

BOOLEAN
KdCommandPacketAndBufferToDebuggee(
    DEBUGGER_REMOTE_PACKET_TYPE             PacketType,
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION RequestedAction,
    CHAR *                                  Buffer,
    UINT32                                  BufferLength);

BOOLEAN
KdPrepareSerialConnectionToRemoteSystem(HANDLE  SerialHandle,
                                        BOOLEAN IsNamedPipe);

BOOLEAN
KdPrepareAndConnectDebugPort(const char * PortName, DWORD Baudrate, UINT32 Port, BOOLEAN IsPreparing, BOOLEAN IsNamedPipe);

BOOLEAN
KdSendPacketToDebuggee(const CHAR * Buffer, UINT32 Length, BOOLEAN SendEndOfBuffer);

BOOLEAN
KdReceivePacketFromDebuggee(CHAR * BufferToSave, UINT32 * LengthReceived);

BOOLEAN
KdReceivePacketFromDebugger(CHAR * BufferToSave, UINT32 * LengthReceived);

BOOLEAN
KdCheckForTheEndOfTheBuffer(PUINT32 CurrentLoopIndex, BYTE * Buffer);

BOOLEAN
KdSendSwitchCorePacketToDebuggee(UINT32 NewCore);

BOOLEAN
KdSendShortCircuitingEventToDebuggee(BOOLEAN IsEnabled);

BOOLEAN
KdSendEventQueryAndModifyPacketToDebuggee(
    UINT64                      Tag,
    DEBUGGER_MODIFY_EVENTS_TYPE TypeOfAction,
    BOOLEAN *                   IsEnabled);

BOOLEAN
KdSendFlushPacketToDebuggee();

BOOLEAN
KdSendCallStackPacketToDebuggee(UINT64                            BaseAddress,
                                UINT32                            Size,
                                DEBUGGER_CALLSTACK_DISPLAY_METHOD DisplayMethod,
                                BOOLEAN                           Is32Bit);

BOOLEAN
KdSendTestQueryPacketToDebuggee(DEBUGGER_TEST_QUERY_STATE Type);

BOOLEAN
KdSendTestQueryPacketWithContextToDebuggee(DEBUGGER_TEST_QUERY_STATE Type, UINT64 Context);

BOOLEAN
KdSendSymbolReloadPacketToDebuggee(UINT32 ProcessId);

BOOLEAN KdSendReadRegisterPacketToDebuggee(PDEBUGGEE_REGISTER_READ_DESCRIPTION);

BOOLEAN
KdSendReadMemoryPacketToDebuggee(PDEBUGGER_READ_MEMORY ReadMem);

BOOLEAN
KdSendEditMemoryPacketToDebuggee(PDEBUGGER_EDIT_MEMORY EditMem, UINT32 Size);

PDEBUGGER_EVENT_AND_ACTION_RESULT
KdSendRegisterEventPacketToDebuggee(PDEBUGGER_GENERAL_EVENT_DETAIL Event,
                                    UINT32                         EventBufferLength);

PDEBUGGER_EVENT_AND_ACTION_RESULT
KdSendAddActionToEventPacketToDebuggee(PDEBUGGER_GENERAL_ACTION GeneralAction,
                                       UINT32                   GeneralActionLength);

BOOLEAN
KdSendSwitchProcessPacketToDebuggee(DEBUGGEE_DETAILS_AND_SWITCH_PROCESS_TYPE ActionType,
                                    UINT32                                   NewPid,
                                    UINT64                                   NewProcess,
                                    BOOLEAN                                  SetChangeByClockInterrupt,
                                    PDEBUGGEE_PROCESS_LIST_NEEDED_DETAILS    SymDetailsForProcessList);

BOOLEAN
KdSendSwitchThreadPacketToDebuggee(DEBUGGEE_DETAILS_AND_SWITCH_THREAD_TYPE ActionType,
                                   UINT32                                  NewTid,
                                   UINT64                                  NewThread,
                                   BOOLEAN                                 CheckByClockInterrupt,
                                   PDEBUGGEE_THREAD_LIST_NEEDED_DETAILS    SymDetailsForThreadList);

BOOLEAN
KdSendBpPacketToDebuggee(PDEBUGGEE_BP_PACKET BpPacket);

BOOLEAN
KdSendVa2paAndPa2vaPacketToDebuggee(PDEBUGGER_VA2PA_AND_PA2VA_COMMANDS Va2paAndPa2vaPacket);

BOOLEAN
KdSendPtePacketToDebuggee(PDEBUGGER_READ_PAGE_TABLE_ENTRIES_DETAILS PtePacket);

BOOLEAN
KdSendPageinPacketToDebuggee(PDEBUGGER_PAGE_IN_REQUEST PageinPacket);

BOOLEAN
KdSendListOrModifyPacketToDebuggee(
    PDEBUGGEE_BP_LIST_OR_MODIFY_PACKET ListOrModifyPacket);

BOOLEAN
KdSendScriptPacketToDebuggee(UINT64 BufferAddress, UINT32 BufferLength, UINT32 Pointer, BOOLEAN IsFormat);

BOOLEAN
KdSendUserInputPacketToDebuggee(const char * Sendbuf, int Len, BOOLEAN IgnoreBreakingAgain);

BOOLEAN
KdSendSearchRequestPacketToDebuggee(UINT64 * SearchRequestBuffer, UINT32 SearchRequestBufferSize);

BOOLEAN
KdSendStepPacketToDebuggee(DEBUGGER_REMOTE_STEPPING_REQUEST StepRequestType);

BYTE
KdComputeDataChecksum(PVOID Buffer, UINT32 Length);

BOOLEAN
KdRegisterEventInDebuggee(PDEBUGGER_GENERAL_EVENT_DETAIL EventRegBuffer,
                          UINT32                         Length);

BOOLEAN
KdAddActionToEventInDebuggee(PDEBUGGER_GENERAL_ACTION ActionAddingBuffer,
                             UINT32                   Length);

BOOLEAN
KdSendModifyEventInDebuggee(PDEBUGGER_MODIFY_EVENTS ModifyEvent, BOOLEAN SendTheResultBackToDebugger);

BOOLEAN
KdSendGeneralBuffersFromDebuggeeToDebugger(
    DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION RequestedAction,
    PVOID                                   Buffer,
    UINT32                                  BufferLength,
    BOOLEAN                                 PauseDebuggeeWhenSent);

BOOLEAN
KdCloseConnection();

BOOLEAN
KdReloadSymbolsInDebuggee(BOOLEAN PauseDebuggee, UINT32 UserProcessId);

BOOLEAN
KdSendResponseOfThePingPacket();

VOID
KdUninitializeConnection();

VOID
KdSendUsermodePrints(CHAR * Input, UINT32 Length);

VOID
KdSendSymbolDetailPacket(PMODULE_SYMBOL_DETAIL SymbolDetailPacket,
                         UINT32                CurrentSymbolInfoIndex,
                         UINT32                TotalSymbols);

VOID
KdHandleUserInputInDebuggee(DEBUGGEE_USER_INPUT_PACKET * Descriptor);

VOID
KdTheRemoteSystemIsRunning();

VOID
KdBreakControlCheckAndPauseDebugger();

VOID
KdBreakControlCheckAndContinueDebugger();

VOID
KdSetStatusAndWaitForPause();


//..\bin\debug\headerAll\list.h
/**
 * @file list.h
 * @author Sina Karvandi (sina@hyperdbg.org)
 * @brief The list working functions headers
 * @details
 * @version 0.1
 * @date 2020-04-11
 *
 * @copyright This project is released under the GNU Public License v3.
 *
 */
#pragma once

/*
#define CONTAINING_RECORD(address, type, field) \
    ((type *)((char *)(address) - (unsigned long)(&((type *)0)->field)))
*/


//..\bin\debug\headerAll\namedpipe.h
/**
 * @file namedpipe.h
 * @author Sina Karvandi (sina@hyperdbg.org)
 * @brief Named pipe communication headers
 * @details
 * @version 0.1
 * @date 2020-04-11
 *
 * @copyright This project is released under the GNU Public License v3.
 *
 */
#pragma once

////////////////////////////////////////////////////////////////////////////
//                            Server Side                                 //
////////////////////////////////////////////////////////////////////////////

HANDLE
NamedPipeServerCreatePipe(LPCSTR PipeName, UINT32 OutputBufferSize, UINT32 InputBufferSize);

BOOLEAN
NamedPipeServerWaitForClientConntection(HANDLE PipeHandle);

UINT32
NamedPipeServerReadClientMessage(HANDLE PipeHandle, char * BufferToSave, int MaximumReadBufferLength);

BOOLEAN
NamedPipeServerSendMessageToClient(HANDLE PipeHandle,
                                   char * BufferToSend,
                                   int    BufferSize);

VOID
NamedPipeServerCloseHandle(HANDLE PipeHandle);

////////////////////////////////////////////////////////////////////////////
//                            Client Side                                 //
////////////////////////////////////////////////////////////////////////////

HANDLE
NamedPipeClientCreatePipe(LPCSTR PipeName);

HANDLE
NamedPipeClientCreatePipeOverlappedIo(LPCSTR PipeName);

BOOLEAN
NamedPipeClientSendMessage(HANDLE PipeHandle, char * BufferToSend, int BufferSize);

UINT32
NamedPipeClientReadMessage(HANDLE PipeHandle, char * BufferToRead, int MaximumSizeOfBuffer);

VOID
NamedPipeClientClosePipe(HANDLE PipeHandle);


//..\bin\debug\headerAll\objects.h
/**
 * @file objects.h
 * @author Sina Karvandi (sina@hyperdbg.org)
 * @brief Header for routines related to objects
 * @details
 * @version 0.1
 * @date 2022-05-06
 *
 * @copyright This project is released under the GNU Public License v3.
 *
 */
#pragma once

//////////////////////////////////////////
//				Functions 		     	//
//////////////////////////////////////////

BOOLEAN
ObjectShowProcessesOrThreadList(BOOLEAN                               IsProcess,
                                PDEBUGGEE_PROCESS_LIST_NEEDED_DETAILS SymDetailsForProcessList,
                                UINT64                                Eprocess,
                                PDEBUGGEE_THREAD_LIST_NEEDED_DETAILS  SymDetailsForThreadList);

BOOLEAN
ObjectShowProcessesOrThreadDetails(BOOLEAN IsProcess);


//..\bin\debug\headerAll\pe-parser.h
/**
 * @file pe-parser.h
 * @author Sina Karvandi (sina@hyperdbg.org)
 * @brief Header for Portable Executable parser
 * @details
 * @version 0.1
 * @date 2021-12-26
 *
 * @copyright This project is released under the GNU Public License v3.
 *
 */
#pragma once

//////////////////////////////////////////////////
//					  Functions                 //
//////////////////////////////////////////////////

BOOLEAN
PeShowSectionInformationAndDump(const WCHAR * AddressOfFile, const CHAR * SectionToShow, BOOLEAN Is32Bit);

BOOLEAN
PeIsPE32BitOr64Bit(const WCHAR * AddressOfFile, PBOOLEAN Is32Bit);


//..\bin\debug\headerAll\rev-ctrl.h
/**
 * @file rev-ctrl.h
 * @author Sina Karvandi (sina@hyperdbg.org)
 * @brief headers for controller of the reversing machine's module
 * @details
 * @version 0.2
 * @date 2023-03-23
 *
 * @copyright This project is released under the GNU Public License v3.
 *
 */
#pragma once

//////////////////////////////////////////////////
//            	    Functions                   //
//////////////////////////////////////////////////

BOOLEAN
RevRequestService(REVERSING_MACHINE_RECONSTRUCT_MEMORY_REQUEST * RevRequest);


//..\bin\debug\headerAll\script-engine.h
/**
 * @file script-engine.h
 * @author Sina Karvandi (sina@hyperdbg.org)
 * @brief General script-engine functions and wrappers
 * @details
 * @version 0.1
 * @date 2021-09-23
 *
 * @copyright This project is released under the GNU Public License v3.
 *
 */
#pragma once

//////////////////////////////////////////////////
//    Pdb Parser Wrapper (from script-engine)   //
//////////////////////////////////////////////////
UINT64
ScriptEngineConvertNameToAddressWrapper(const char * FunctionOrVariableName, PBOOLEAN WasFound);

UINT32
ScriptEngineLoadFileSymbolWrapper(UINT64 BaseAddress, const char * PdbFileName, const char * CustomModuleName);

VOID
ScriptEngineSetTextMessageCallbackWrapper(PVOID Handler);

UINT32
ScriptEngineUnloadAllSymbolsWrapper();

UINT32
ScriptEngineUnloadModuleSymbolWrapper(char * ModuleName);

UINT32
ScriptEngineSearchSymbolForMaskWrapper(const char * SearchMask);

BOOLEAN
ScriptEngineGetFieldOffsetWrapper(CHAR * TypeName, CHAR * FieldName, UINT32 * FieldOffset);

BOOLEAN
ScriptEngineGetDataTypeSizeWrapper(CHAR * TypeName, UINT64 * TypeSize);

BOOLEAN
ScriptEngineCreateSymbolTableForDisassemblerWrapper(void * CallbackFunction);

BOOLEAN
ScriptEngineConvertFileToPdbPathWrapper(const char * LocalFilePath, char * ResultPath);

BOOLEAN
ScriptEngineConvertFileToPdbFileAndGuidAndAgeDetailsWrapper(const char * LocalFilePath,
                                                            char *       PdbFilePath,
                                                            char *       GuidAndAgeDetails,
                                                            BOOLEAN      Is32BitModule);

BOOLEAN
ScriptEngineSymbolInitLoadWrapper(PMODULE_SYMBOL_DETAIL BufferToStoreDetails,
                                  UINT32                StoredLength,
                                  BOOLEAN               DownloadIfAvailable,
                                  const char *          SymbolPath,
                                  BOOLEAN               IsSilentLoad);

BOOLEAN
ScriptEngineShowDataBasedOnSymbolTypesWrapper(
    const char * TypeName,
    UINT64       Address,
    BOOLEAN      IsStruct,
    PVOID        BufferAddress,
    const char * AdditionalParameters);

VOID
ScriptEngineSymbolAbortLoadingWrapper();

//////////////////////////////////////////////////
//          Script Engine Wrapper               //
//////////////////////////////////////////////////

VOID
ScriptEngineWrapperTestParser(const string & Expr);

BOOLEAN
ScriptAutomaticStatementsTestWrapper(const string & Expr, UINT64 ExpectationValue, BOOLEAN ExceptError);

PVOID
ScriptEngineParseWrapper(char * Expr, BOOLEAN ShowErrorMessageIfAny);

VOID
PrintSymbolBufferWrapper(PVOID SymbolBuffer);

UINT64
ScriptEngineWrapperGetHead(PVOID SymbolBuffer);

UINT32
ScriptEngineWrapperGetSize(PVOID SymbolBuffer);

UINT32
ScriptEngineWrapperGetPointer(PVOID SymbolBuffer);

VOID
ScriptEngineWrapperRemoveSymbolBuffer(PVOID SymbolBuffer);

BOOLEAN
ScriptEngineFuncNumberOfOperands(UINT64 FuncType, UINT32 * NumberOfGetOperands, UINT32 * NumberOfSetOperands);

UINT64
ScriptEngineEvalUInt64StyleExpressionWrapper(const string & Expr, PBOOLEAN HasError);

//////////////////////////////////////////////////
//          Script Engine Functions             //
//////////////////////////////////////////////////

UINT64
ScriptEngineEvalSingleExpression(string Expr, PBOOLEAN HasError);


//..\bin\debug\headerAll\symbol.h
/**
 * @file symbol.h
 * @author Sina Karvandi (sina@hyperdbg.org)
 * @brief Symbol related functions header
 * @details
 * @version 0.1
 * @date 2021-06-09
 *
 * @copyright This project is released under the GNU Public License v3.
 *
 */
#pragma once

//////////////////////////////////////////////////
//			        Structures		            //
//////////////////////////////////////////////////

/**
 * @brief Save the local function symbols' description
 *
 */
typedef struct _LOCAL_FUNCTION_DESCRIPTION
{
    std::string ObjectName;
    UINT32      ObjectSize;

} LOCAL_FUNCTION_DESCRIPTION, *PLOCAL_FUNCTION_DESCRIPTION;

//////////////////////////////////////////////////
//			    	    Pdbex                   //
//////////////////////////////////////////////////

#define PDBEX_DEFAULT_CONFIGURATION "-j- -k- -e n -i"

//////////////////////////////////////////////////
//			 For symbol (pdb) parsing		    //
//////////////////////////////////////////////////

VOID
SymbolBuildAndShowSymbolTable();

BOOLEAN
SymbolShowFunctionNameBasedOnAddress(UINT64 Address, PUINT64 UsedBaseAddress);

BOOLEAN
SymbolLoadOrDownloadSymbols(BOOLEAN IsDownload, BOOLEAN SilentLoad);

BOOLEAN
SymbolConvertNameOrExprToAddress(const string & TextToConvert, PUINT64 Result);

BOOLEAN
SymbolDeleteSymTable();

BOOLEAN
SymbolBuildSymbolTable(PMODULE_SYMBOL_DETAIL * BufferToStoreDetails,
                       PUINT32                 StoredLength,
                       UINT32                  UserProcessId,
                       BOOLEAN                 SendOverSerial);

BOOLEAN
SymbolBuildAndUpdateSymbolTable(PMODULE_SYMBOL_DETAIL SymbolDetail);

VOID
SymbolInitialReload();

BOOLEAN
SymbolLocalReload(UINT32 UserProcessId);

VOID
SymbolPrepareDebuggerWithSymbolInfo(UINT32 UserProcessId);

BOOLEAN
SymbolReloadSymbolTableInDebuggerMode(UINT32 ProcessId);


//..\bin\debug\headerAll\tests.h
/**
 * @file tests.h
 * @author Sina Karvandi (sina@hyperdbg.org)
 * @brief headers for test functions
 * @details
 * @version 0.1
 * @date 2020-05-27
 *
 * @copyright This project is released under the GNU Public License v3.
 *
 */
#pragma once

//////////////////////////////////////////////////
//					Constants                   //
//////////////////////////////////////////////////

/**
 * @brief exe name of test process
 *
 */
#define TEST_PROCESS_NAME "hyperdbg-test.exe"

//////////////////////////////////////////////////
//					Functions                   //
//////////////////////////////////////////////////

BOOLEAN
CreateProcessAndOpenPipeConnection(PHANDLE ConnectionPipeHandle,
                                   PHANDLE ThreadHandle,
                                   PHANDLE ProcessHandle);
VOID
CloseProcessAndClosePipeConnection(HANDLE ConnectionPipeHandle,
                                   HANDLE ThreadHandle,
                                   HANDLE ProcessHandle);


//..\bin\debug\headerAll\transparency.h
/**
 * @file transparency.h
 * @author Sina Karvandi (sina@hyperdbg.org)
 * @brief headers for test functions
 * @details
 * @version 0.1
 * @date 2020-07-30
 *
 * @copyright This project is released under the GNU Public License v3.
 *
 */
#pragma once

//////////////////////////////////////////////////
//				  Definitions                   //
//////////////////////////////////////////////////

/**
 * @brief Number of tests for each instruction sets
 * @details used to generate test cases for rdts+cpuid+rdtsc
 * and rdtsc+rdtsc commands
 */
#define TestCount 1000

//////////////////////////////////////////////////
//				    Functions                   //
//////////////////////////////////////////////////

void
GuassianGenerateRandom(vector<double> Data, UINT64 * AverageOfData, UINT64 * StandardDeviationOfData, UINT64 * MedianOfData);

BOOLEAN
TransparentModeCheckHypervisorPresence(UINT64 * Average,
                                       UINT64 * StandardDeviation,
                                       UINT64 * Median);

BOOLEAN
TransparentModeCheckRdtscpVmexit(UINT64 * Average,
                                 UINT64 * StandardDeviation,
                                 UINT64 * Median);

double
Randn(double mu, double sigma);

double
Median(vector<double> Cases);

unsigned long long
TransparentModeRdtscDiffVmexit();

unsigned long long
TransparentModeRdtscVmexitTracing();


//..\bin\debug\headerAll\ud.h
/**
 * @file ud.h
 * @author Sina Karvandi (sina@hyperdbg.org)
 * @brief headers for user-mode debugging routines
 * @details
 * @version 0.1
 * @date 2021-12-28
 *
 * @copyright This project is released under the GNU Public License v3.
 *
 */
#pragma once

//////////////////////////////////////////////////
//		            Definitions                 //
//////////////////////////////////////////////////

#define DbgWaitForUserResponse(UserSyncObjectId)                          \
    do                                                                    \
    {                                                                     \
        DEBUGGER_SYNCRONIZATION_EVENTS_STATE * SyncronizationObject =     \
            &g_UserSyncronizationObjectsHandleTable[UserSyncObjectId];    \
                                                                          \
        SyncronizationObject->IsOnWaitingState = TRUE;                    \
        WaitForSingleObject(SyncronizationObject->EventHandle, INFINITE); \
    } while (FALSE);

#define DbgReceivedUserResponse(UserSyncObjectId)                      \
    do                                                                 \
    {                                                                  \
        DEBUGGER_SYNCRONIZATION_EVENTS_STATE * SyncronizationObject =  \
            &g_UserSyncronizationObjectsHandleTable[UserSyncObjectId]; \
                                                                       \
        SyncronizationObject->IsOnWaitingState = FALSE;                \
        SetEvent(SyncronizationObject->EventHandle);                   \
    } while (FALSE);

//////////////////////////////////////////////////
//            	    Structures                  //
//////////////////////////////////////////////////

/**
 * @brief structures related to current thread debugging
 * state
 *
 */
typedef struct _ACTIVE_DEBUGGING_PROCESS
{
    BOOLEAN    IsActive;
    UINT64     ProcessDebuggingToken;
    UINT32     ProcessId;
    UINT32     ThreadId;
    BOOLEAN    IsPaused;
    GUEST_REGS Registers; // thread registers
    UINT64     Context;   // $context
    BOOLEAN    Is32Bit;
} ACTIVE_DEBUGGING_PROCESS, *PACTIVE_DEBUGGING_PROCESS;

//////////////////////////////////////////////////
//            	    Functions                  //
//////////////////////////////////////////////////

VOID
UdInitializeUserDebugger();

VOID
UdUninitializeUserDebugger();

VOID
UdRemoveActiveDebuggingProcess(BOOLEAN DontSwitchToNewProcess);

VOID
UdHandleUserDebuggerPausing(PDEBUGGEE_UD_PAUSED_PACKET PausePacket);

VOID
UdContinueDebuggee(UINT64 ProcessDetailToken);

VOID
UdSendStepPacketToDebuggee(UINT64 ThreadDetailToken, UINT32 TargetThreadId, DEBUGGER_REMOTE_STEPPING_REQUEST StepType);

VOID
UdSetActiveDebuggingProcess(UINT64  DebuggingId,
                            UINT32  ProcessId,
                            UINT32  ThreadId,
                            BOOLEAN Is32Bit,
                            BOOLEAN IsPaused);
BOOLEAN
UdSetActiveDebuggingThreadByPidOrTid(UINT32 TargetPidOrTid, BOOLEAN IsTid);

BOOLEAN
UdSetActiveDebuggingThreadByPidOrTid(UINT32 TargetPidOrTid, BOOLEAN IsTid);

BOOLEAN
UdShowListActiveDebuggingProcessesAndThreads();

BOOL
UdListProcessThreads(DWORD OwnerPID);

BOOLEAN
UdCheckThreadByProcessId(DWORD Pid, DWORD Tid);

BOOLEAN
UdAttachToProcess(UINT32        TargetPid,
                  const WCHAR * TargetFileAddress,
                  WCHAR *       CommandLine,
                  BOOLEAN       RunCallbackAtTheFirstInstruction);

BOOLEAN
UdKillProcess(UINT32 TargetPid);

BOOLEAN
UdDetachProcess(UINT32 TargetPid, UINT64 ProcessDetailToken);

BOOLEAN
UdPauseProcess(UINT64 ProcessDebuggingToken);


