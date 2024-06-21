//bugfix.h

typedef unsigned short wchar_t;
typedef int bool ;
#define PVOID void*
#define HANDLE void*
#define MAX_PATH 260
typedef unsigned __int64   SIZE_T;
typedef unsigned __int64   time_t;

typedef struct _LIST_ENTRY {
  struct _LIST_ENTRY *Flink;
  struct _LIST_ENTRY *Blink;
} LIST_ENTRY, *PLIST_ENTRY, PRLIST_ENTRY;


//..\..\..\bin\debug\SDK\Headers\BasicTypes.h
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


//..\..\..\bin\debug\SDK\Modules\HyperLog.h
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


//..\..\..\bin\debug\SDK\Modules\VMM.h
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


//..\..\..\bin\debug\SDK\Headers\DataTypes.h
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


//..\..\..\bin\debug\SDK\Headers\Events.h
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


//..\..\..\bin\debug\SDK\Headers\Ioctls.h
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


//..\..\..\bin\debug\SDK\Imports\HyperDbgRevImports.h
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


//..\..\..\bin\debug\SDK\Imports\HyperDbgScriptImports.h
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


//..\..\..\bin\debug\SDK\Imports\HyperDbgVmmImports.h
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


//..\..\..\bin\debug\SDK\Headers\ErrorCodes.h
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


//..\..\..\bin\debug\SDK\Headers\RequestStructures.h
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


//..\..\..\bin\debug\SDK\Imports\HyperDbgHyperLogImports.h
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


//..\..\..\bin\debug\SDK\Imports\HyperDbgHyperLogIntrinsics.h
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


//..\..\..\bin\debug\SDK\Headers\Constants.h
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


//..\..\..\bin\debug\SDK\Headers\Symbols.h
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


//..\..\..\bin\debug\SDK\Imports\HyperDbgCtrlImports.h
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


//..\..\..\bin\debug\SDK\Imports\HyperDbgSymImports.h
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


//..\..\..\bin\debug\SDK\Headers\Connection.h
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


//..\..\..\bin\debug\SDK\Headers\HardwareDebugger.h
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


