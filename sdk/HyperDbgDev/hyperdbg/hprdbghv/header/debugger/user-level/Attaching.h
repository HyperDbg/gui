#pragma once
#define MAX_USER_ACTIONS_FOR_THREADS    3
#define MAX_THREADS_IN_A_PROCESS_HOLDER 100
#define MAX_CR3_IN_A_PROCESS            4
typedef struct _USERMODE_DEBUGGING_PROCESS_DETAILS {
    UINT64     Token;
    BOOLEAN    Enabled;
    PVOID      PebAddressToMonitor;
    UINT32     ActiveThreadId; // active thread
    GUEST_REGS Registers;      // active thread
    UINT64     Context;        // $context
    LIST_ENTRY AttachedProcessList;
    UINT64     UsermodeReservedBuffer;
    UINT64     EntrypointOfMainModule;
    UINT64     BaseAddressOfMainModule;
    PEPROCESS  Eprocess;
    UINT32     ProcessId;
    BOOLEAN    Is32Bit;
    BOOLEAN    IsOnTheStartingPhase;
    BOOLEAN    IsOnThreadInterceptingPhase;
    CR3_TYPE   InterceptedCr3[MAX_CR3_IN_A_PROCESS];
    LIST_ENTRY ThreadsListHead;
} USERMODE_DEBUGGING_PROCESS_DETAILS, *PUSERMODE_DEBUGGING_PROCESS_DETAILS;
BOOLEAN
AttachingInitialize();
BOOLEAN
AttachingCheckPageFaultsWithUserDebugger(UINT32                       CurrentProcessorIndex,
                                         PGUEST_REGS                  GuestRegs,
                                         VMEXIT_INTERRUPT_INFORMATION InterruptExit,
                                         UINT64                       Address,
                                         ULONG                        ErrorCode);
BOOLEAN
AttachingConfigureInterceptingThreads(UINT64 ProcessDebuggingToken, BOOLEAN Enable);
BOOLEAN
AttachingHandleCr3VmexitsForThreadInterception(UINT32 CurrentCoreIndex, CR3_TYPE NewCr3);
VOID
AttachingTargetProcess(PDEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS Request);
VOID
AttachingHandleEntrypointDebugBreak(UINT32 CurrentProcessorIndex, PGUEST_REGS GuestRegs);
VOID
AttachingRemoveAndFreeAllProcessDebuggingDetails();
PUSERMODE_DEBUGGING_PROCESS_DETAILS
AttachingFindProcessDebuggingDetailsByToken(UINT64 Token);
PUSERMODE_DEBUGGING_PROCESS_DETAILS
AttachingFindProcessDebuggingDetailsByProcessId(UINT32 ProcessId);
BOOLEAN
AttachingQueryDetailsOfActiveDebuggingThreadsAndProcesses(PVOID BufferToStoreDetails, UINT32 BufferSize);
