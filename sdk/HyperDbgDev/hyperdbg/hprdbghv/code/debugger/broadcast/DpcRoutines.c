#include "pch.h"
volatile LONG OneCoreLock;
NTSTATUS
DpcRoutineRunTaskOnSingleCore(UINT32 CoreNumber, PVOID Routine, PVOID DeferredContext) {
    PRKDPC Dpc;
    UINT32 ProcessorCount;
    ProcessorCount = KeQueryActiveProcessorCount(0);
    if (CoreNumber >= ProcessorCount) {
        return STATUS_INVALID_PARAMETER;
    }
    Dpc = ExAllocatePoolWithTag(NonPagedPool, sizeof(KDPC), POOLTAG);
    if (!Dpc) {
        return STATUS_INSUFFICIENT_RESOURCES;
    }
    KeInitializeDpc(Dpc,            // Dpc
                    Routine,        // DeferredRoutine
                    DeferredContext // DeferredContext
    );
    KeSetTargetProcessorDpc(Dpc, CoreNumber);
    if (!SpinlockTryLock(&OneCoreLock)) {
        ExFreePoolWithTag(Dpc, POOLTAG);
        return STATUS_UNSUCCESSFUL;
    }
    KeInsertQueueDpc(Dpc, NULL, NULL);
    SpinlockLock(&OneCoreLock);
    SpinlockUnlock(&OneCoreLock);
    ExFreePoolWithTag(Dpc, POOLTAG);
    return STATUS_SUCCESS;
}

BOOLEAN
DpcRoutinePerformVirtualization(KDPC * Dpc, PVOID DeferredContext, PVOID SystemArgument1, PVOID SystemArgument2) {
    VmxPerformVirtualizationOnSpecificCore();
    KeSignalCallDpcSynchronize(SystemArgument2);
    KeSignalCallDpcDone(SystemArgument1);
    return TRUE;
}

VOID
DpcRoutinePerformWriteMsr(KDPC * Dpc, PVOID DeferredContext, PVOID SystemArgument1, PVOID SystemArgument2) {
    ULONG                       CurrentProcessorIndex = 0;
    ULONG                       CurrentCore           = KeGetCurrentProcessorNumber();
    VIRTUAL_MACHINE_STATE *     CurrentVmState        = &g_GuestState[CurrentCore];
    PROCESSOR_DEBUGGING_STATE * CurrentDebuggingState = &CurrentVmState->DebuggingState;
    UNREFERENCED_PARAMETER(Dpc);
    UNREFERENCED_PARAMETER(DeferredContext);
    UNREFERENCED_PARAMETER(SystemArgument1);
    UNREFERENCED_PARAMETER(SystemArgument2);
    __writemsr(CurrentDebuggingState->MsrState.Msr, CurrentDebuggingState->MsrState.Value);
    SpinlockUnlock(&OneCoreLock);
}

VOID
DpcRoutinePerformReadMsr(KDPC * Dpc, PVOID DeferredContext, PVOID SystemArgument1, PVOID SystemArgument2) {
    ULONG                       CurrentCore           = KeGetCurrentProcessorNumber();
    VIRTUAL_MACHINE_STATE *     CurrentVmState        = &g_GuestState[CurrentCore];
    PROCESSOR_DEBUGGING_STATE * CurrentDebuggingState = &CurrentVmState->DebuggingState;
    UNREFERENCED_PARAMETER(Dpc);
    UNREFERENCED_PARAMETER(DeferredContext);
    UNREFERENCED_PARAMETER(SystemArgument1);
    UNREFERENCED_PARAMETER(SystemArgument2);
    CurrentDebuggingState->MsrState.Value = __readmsr(CurrentDebuggingState->MsrState.Msr);
    SpinlockUnlock(&OneCoreLock);
}

VOID
DpcRoutinePerformChangeMsrBitmapReadOnSingleCore(KDPC * Dpc, PVOID DeferredContext, PVOID SystemArgument1, PVOID SystemArgument2) {
    UNREFERENCED_PARAMETER(Dpc);
    UNREFERENCED_PARAMETER(SystemArgument1);
    UNREFERENCED_PARAMETER(SystemArgument2);
    AsmVmxVmcall(VMCALL_CHANGE_MSR_BITMAP_READ, DeferredContext, 0, 0);
    SpinlockUnlock(&OneCoreLock);
}

VOID
DpcRoutinePerformChangeMsrBitmapWriteOnSingleCore(KDPC * Dpc, PVOID DeferredContext, PVOID SystemArgument1, PVOID SystemArgument2) {
    UNREFERENCED_PARAMETER(Dpc);
    UNREFERENCED_PARAMETER(SystemArgument1);
    UNREFERENCED_PARAMETER(SystemArgument2);
    AsmVmxVmcall(VMCALL_CHANGE_MSR_BITMAP_WRITE, DeferredContext, 0, 0);
    SpinlockUnlock(&OneCoreLock);
}

VOID
DpcRoutinePerformEnableRdtscExitingOnSingleCore(KDPC * Dpc, PVOID DeferredContext, PVOID SystemArgument1, PVOID SystemArgument2) {
    UNREFERENCED_PARAMETER(Dpc);
    UNREFERENCED_PARAMETER(DeferredContext);
    UNREFERENCED_PARAMETER(SystemArgument1);
    UNREFERENCED_PARAMETER(SystemArgument2);
    AsmVmxVmcall(VMCALL_SET_RDTSC_EXITING, 0, 0, 0);
    SpinlockUnlock(&OneCoreLock);
}

VOID
DpcRoutinePerformEnableRdpmcExitingOnSingleCore(KDPC * Dpc, PVOID DeferredContext, PVOID SystemArgument1, PVOID SystemArgument2) {
    UNREFERENCED_PARAMETER(Dpc);
    UNREFERENCED_PARAMETER(DeferredContext);
    UNREFERENCED_PARAMETER(SystemArgument1);
    UNREFERENCED_PARAMETER(SystemArgument2);
    AsmVmxVmcall(VMCALL_SET_RDPMC_EXITING, 0, 0, 0);
    SpinlockUnlock(&OneCoreLock);
}

VOID
DpcRoutinePerformSetExceptionBitmapOnSingleCore(KDPC * Dpc, PVOID DeferredContext, PVOID SystemArgument1, PVOID SystemArgument2) {
    UNREFERENCED_PARAMETER(Dpc);
    UNREFERENCED_PARAMETER(DeferredContext);
    UNREFERENCED_PARAMETER(SystemArgument1);
    UNREFERENCED_PARAMETER(SystemArgument2);
    AsmVmxVmcall(VMCALL_SET_EXCEPTION_BITMAP, DeferredContext, 0, 0);
    SpinlockUnlock(&OneCoreLock);
}

VOID
DpcRoutinePerformEnableMovToDebugRegistersExiting(KDPC * Dpc, PVOID DeferredContext, PVOID SystemArgument1, PVOID SystemArgument2) {
    UNREFERENCED_PARAMETER(Dpc);
    UNREFERENCED_PARAMETER(DeferredContext);
    UNREFERENCED_PARAMETER(SystemArgument1);
    UNREFERENCED_PARAMETER(SystemArgument2);
    AsmVmxVmcall(VMCALL_ENABLE_MOV_TO_DEBUG_REGS_EXITING, 0, 0, 0);
    SpinlockUnlock(&OneCoreLock);
}

VOID
DpcRoutinePerformEnableMovToControlRegisterExiting(KDPC * Dpc, PDEBUGGER_EVENT Event, PVOID SystemArgument1, PVOID SystemArgument2) {
    UNREFERENCED_PARAMETER(Dpc);
    UNREFERENCED_PARAMETER(SystemArgument1);
    UNREFERENCED_PARAMETER(SystemArgument2);
    AsmVmxVmcall(VMCALL_ENABLE_MOV_TO_CONTROL_REGS_EXITING, Event->OptionalParam1, Event->OptionalParam2, 0);
    SpinlockUnlock(&OneCoreLock);
}

VOID
DpcRoutinePerformSetExternalInterruptExitingOnSingleCore(KDPC * Dpc, PVOID DeferredContext, PVOID SystemArgument1, PVOID SystemArgument2) {
    UNREFERENCED_PARAMETER(Dpc);
    UNREFERENCED_PARAMETER(DeferredContext);
    UNREFERENCED_PARAMETER(SystemArgument1);
    UNREFERENCED_PARAMETER(SystemArgument2);
    AsmVmxVmcall(VMCALL_ENABLE_EXTERNAL_INTERRUPT_EXITING, NULL, 0, 0);
    SpinlockUnlock(&OneCoreLock);
}

VOID
DpcRoutinePerformEnableEferSyscallHookOnSingleCore(KDPC * Dpc, PVOID DeferredContext, PVOID SystemArgument1, PVOID SystemArgument2) {
    UNREFERENCED_PARAMETER(Dpc);
    UNREFERENCED_PARAMETER(DeferredContext);
    UNREFERENCED_PARAMETER(SystemArgument1);
    UNREFERENCED_PARAMETER(SystemArgument2);
    AsmVmxVmcall(VMCALL_ENABLE_SYSCALL_HOOK_EFER, NULL, 0, 0);
    SpinlockUnlock(&OneCoreLock);
}

VOID
DpcRoutinePerformChangeIoBitmapOnSingleCore(KDPC * Dpc, PVOID DeferredContext, PVOID SystemArgument1, PVOID SystemArgument2) {
    UNREFERENCED_PARAMETER(Dpc);
    UNREFERENCED_PARAMETER(SystemArgument1);
    UNREFERENCED_PARAMETER(SystemArgument2);
    AsmVmxVmcall(VMCALL_CHANGE_IO_BITMAP, DeferredContext, 0, 0);
    SpinlockUnlock(&OneCoreLock);
}

VOID
DpcRoutineEnableMovToCr3Exiting(KDPC * Dpc, PVOID DeferredContext, PVOID SystemArgument1, PVOID SystemArgument2) {
    UNREFERENCED_PARAMETER(Dpc);
    UNREFERENCED_PARAMETER(DeferredContext);
    AsmVmxVmcall(VMCALL_ENABLE_MOV_TO_CR3_EXITING, 0, 0, 0);
    KeSignalCallDpcSynchronize(SystemArgument2);
    KeSignalCallDpcDone(SystemArgument1);
}

VOID
DpcRoutineDisableMovToCr3Exiting(KDPC * Dpc, PVOID DeferredContext, PVOID SystemArgument1, PVOID SystemArgument2) {
    UNREFERENCED_PARAMETER(Dpc);
    UNREFERENCED_PARAMETER(DeferredContext);
    AsmVmxVmcall(VMCALL_DISABLE_MOV_TO_CR3_EXITING, 0, 0, 0);
    KeSignalCallDpcSynchronize(SystemArgument2);
    KeSignalCallDpcDone(SystemArgument1);
}

VOID
DpcRoutineEnableEferSyscallEvents(KDPC * Dpc, PVOID DeferredContext, PVOID SystemArgument1, PVOID SystemArgument2) {
    UNREFERENCED_PARAMETER(Dpc);
    UNREFERENCED_PARAMETER(DeferredContext);
    AsmVmxVmcall(VMCALL_ENABLE_SYSCALL_HOOK_EFER, 0, 0, 0);
    KeSignalCallDpcSynchronize(SystemArgument2);
    KeSignalCallDpcDone(SystemArgument1);
}

VOID
DpcRoutineDisableEferSyscallEvents(KDPC * Dpc, PVOID DeferredContext, PVOID SystemArgument1, PVOID SystemArgument2) {
    UNREFERENCED_PARAMETER(Dpc);
    UNREFERENCED_PARAMETER(DeferredContext);
    AsmVmxVmcall(VMCALL_DISABLE_SYSCALL_HOOK_EFER, 0, 0, 0);
    KeSignalCallDpcSynchronize(SystemArgument2);
    KeSignalCallDpcDone(SystemArgument1);
}

VOID
DpcRoutineWriteMsrToAllCores(KDPC * Dpc, PVOID DeferredContext, PVOID SystemArgument1, PVOID SystemArgument2) {
    ULONG                       CurrentCore           = KeGetCurrentProcessorNumber();
    VIRTUAL_MACHINE_STATE *     CurrentVmState        = &g_GuestState[CurrentCore];
    PROCESSOR_DEBUGGING_STATE * CurrentDebuggingState = &CurrentVmState->DebuggingState;
    UNREFERENCED_PARAMETER(Dpc);
    UNREFERENCED_PARAMETER(DeferredContext);
    __writemsr(CurrentDebuggingState->MsrState.Msr, CurrentDebuggingState->MsrState.Value);
    KeSignalCallDpcSynchronize(SystemArgument2);
    KeSignalCallDpcDone(SystemArgument1);
}

VOID
DpcRoutineReadMsrToAllCores(KDPC * Dpc, PVOID DeferredContext, PVOID SystemArgument1, PVOID SystemArgument2) {
    ULONG                       CurrentCore           = KeGetCurrentProcessorNumber();
    VIRTUAL_MACHINE_STATE *     CurrentVmState        = &g_GuestState[CurrentCore];
    PROCESSOR_DEBUGGING_STATE * CurrentDebuggingState = &CurrentVmState->DebuggingState;
    UNREFERENCED_PARAMETER(Dpc);
    UNREFERENCED_PARAMETER(DeferredContext);
    CurrentDebuggingState->MsrState.Value = __readmsr(CurrentDebuggingState->MsrState.Msr);
    KeSignalCallDpcSynchronize(SystemArgument2);
    KeSignalCallDpcDone(SystemArgument1);
}

VOID
DpcRoutineChangeMsrBitmapReadOnAllCores(KDPC * Dpc, PVOID DeferredContext, PVOID SystemArgument1, PVOID SystemArgument2) {
    UNREFERENCED_PARAMETER(Dpc);
    AsmVmxVmcall(VMCALL_CHANGE_MSR_BITMAP_READ, DeferredContext, 0, 0);
    KeSignalCallDpcSynchronize(SystemArgument2);
    KeSignalCallDpcDone(SystemArgument1);
}

VOID
DpcRoutineResetMsrBitmapReadOnAllCores(KDPC * Dpc, PVOID DeferredContext, PVOID SystemArgument1, PVOID SystemArgument2) {
    UNREFERENCED_PARAMETER(Dpc);
    UNREFERENCED_PARAMETER(DeferredContext);
    AsmVmxVmcall(VMCALL_RESET_MSR_BITMAP_READ, NULL, 0, 0);
    KeSignalCallDpcSynchronize(SystemArgument2);
    KeSignalCallDpcDone(SystemArgument1);
}

VOID
DpcRoutineChangeMsrBitmapWriteOnAllCores(KDPC * Dpc, PVOID DeferredContext, PVOID SystemArgument1, PVOID SystemArgument2) {
    UNREFERENCED_PARAMETER(Dpc);
    AsmVmxVmcall(VMCALL_CHANGE_MSR_BITMAP_WRITE, DeferredContext, 0, 0);
    KeSignalCallDpcSynchronize(SystemArgument2);
    KeSignalCallDpcDone(SystemArgument1);
}

VOID
DpcRoutineResetMsrBitmapWriteOnAllCores(KDPC * Dpc, PVOID DeferredContext, PVOID SystemArgument1, PVOID SystemArgument2) {
    UNREFERENCED_PARAMETER(Dpc);
    UNREFERENCED_PARAMETER(DeferredContext);
    AsmVmxVmcall(VMCALL_RESET_MSR_BITMAP_WRITE, NULL, 0, 0);
    KeSignalCallDpcSynchronize(SystemArgument2);
    KeSignalCallDpcDone(SystemArgument1);
}

VOID
DpcRoutineEnableRdtscExitingAllCores(KDPC * Dpc, PVOID DeferredContext, PVOID SystemArgument1, PVOID SystemArgument2) {
    UNREFERENCED_PARAMETER(Dpc);
    UNREFERENCED_PARAMETER(DeferredContext);
    AsmVmxVmcall(VMCALL_SET_RDTSC_EXITING, 0, 0, 0);
    KeSignalCallDpcSynchronize(SystemArgument2);
    KeSignalCallDpcDone(SystemArgument1);
}

VOID
DpcRoutineDisableRdtscExitingAllCores(KDPC * Dpc, PVOID DeferredContext, PVOID SystemArgument1, PVOID SystemArgument2) {
    UNREFERENCED_PARAMETER(Dpc);
    UNREFERENCED_PARAMETER(DeferredContext);
    AsmVmxVmcall(VMCALL_UNSET_RDTSC_EXITING, 0, 0, 0);
    KeSignalCallDpcSynchronize(SystemArgument2);
    KeSignalCallDpcDone(SystemArgument1);
}

VOID
DpcRoutineDisableRdtscExitingForClearingTscEventsAllCores(KDPC * Dpc, PVOID DeferredContext, PVOID SystemArgument1, PVOID SystemArgument2) {
    UNREFERENCED_PARAMETER(Dpc);
    UNREFERENCED_PARAMETER(DeferredContext);
    AsmVmxVmcall(VMCALL_DISABLE_RDTSC_EXITING_ONLY_FOR_TSC_EVENTS, 0, 0, 0);
    KeSignalCallDpcSynchronize(SystemArgument2);
    KeSignalCallDpcDone(SystemArgument1);
}

VOID
DpcRoutineDisableMov2DrExitingForClearingDrEventsAllCores(KDPC * Dpc, PVOID DeferredContext, PVOID SystemArgument1, PVOID SystemArgument2) {
    UNREFERENCED_PARAMETER(Dpc);
    UNREFERENCED_PARAMETER(DeferredContext);
    AsmVmxVmcall(VMCALL_DISABLE_MOV_TO_HW_DR_EXITING_ONLY_FOR_DR_EVENTS, 0, 0, 0);
    KeSignalCallDpcSynchronize(SystemArgument2);
    KeSignalCallDpcDone(SystemArgument1);
}

VOID
DpcRoutineDisableMov2CrExitingForClearingCrEventsAllCores(KDPC * Dpc, PDEBUGGER_EVENT Event, PVOID SystemArgument1, PVOID SystemArgument2) {
    UNREFERENCED_PARAMETER(Dpc);
    AsmVmxVmcall(VMCALL_DISABLE_MOV_TO_CR_EXITING_ONLY_FOR_CR_EVENTS, Event->OptionalParam1, Event->OptionalParam2, 0);
    KeSignalCallDpcSynchronize(SystemArgument2);
    KeSignalCallDpcDone(SystemArgument1);
}

VOID
DpcRoutineEnableRdpmcExitingAllCores(KDPC * Dpc, PVOID DeferredContext, PVOID SystemArgument1, PVOID SystemArgument2) {
    UNREFERENCED_PARAMETER(Dpc);
    UNREFERENCED_PARAMETER(DeferredContext);
    AsmVmxVmcall(VMCALL_SET_RDPMC_EXITING, 0, 0, 0);
    KeSignalCallDpcSynchronize(SystemArgument2);
    KeSignalCallDpcDone(SystemArgument1);
}

VOID
DpcRoutineDisableRdpmcExitingAllCores(KDPC * Dpc, PVOID DeferredContext, PVOID SystemArgument1, PVOID SystemArgument2) {
    UNREFERENCED_PARAMETER(Dpc);
    UNREFERENCED_PARAMETER(DeferredContext);
    AsmVmxVmcall(VMCALL_UNSET_RDPMC_EXITING, 0, 0, 0);
    KeSignalCallDpcSynchronize(SystemArgument2);
    KeSignalCallDpcDone(SystemArgument1);
}

VOID
DpcRoutineSetExceptionBitmapOnAllCores(KDPC * Dpc, PVOID DeferredContext, PVOID SystemArgument1, PVOID SystemArgument2) {
    UNREFERENCED_PARAMETER(Dpc);
    AsmVmxVmcall(VMCALL_SET_EXCEPTION_BITMAP, DeferredContext, 0, 0);
    KeSignalCallDpcSynchronize(SystemArgument2);
    KeSignalCallDpcDone(SystemArgument1);
}

VOID
DpcRoutineUnsetExceptionBitmapOnAllCores(KDPC * Dpc, PVOID DeferredContext, PVOID SystemArgument1, PVOID SystemArgument2) {
    UNREFERENCED_PARAMETER(Dpc);
    AsmVmxVmcall(VMCALL_UNSET_EXCEPTION_BITMAP, DeferredContext, 0, 0);
    KeSignalCallDpcSynchronize(SystemArgument2);
    KeSignalCallDpcDone(SystemArgument1);
}

VOID
DpcRoutineResetExceptionBitmapOnlyOnClearingExceptionEventsOnAllCores(KDPC * Dpc,
                                                                      PVOID  DeferredContext,
                                                                      PVOID  SystemArgument1,
                                                                      PVOID  SystemArgument2) {
    UNREFERENCED_PARAMETER(Dpc);
    AsmVmxVmcall(VMCALL_RESET_EXCEPTION_BITMAP_ONLY_ON_CLEARING_EXCEPTION_EVENTS, DeferredContext, 0, 0);
    KeSignalCallDpcSynchronize(SystemArgument2);
    KeSignalCallDpcDone(SystemArgument1);
}

VOID
DpcRoutineEnableMovDebigRegisterExitingAllCores(KDPC * Dpc, PVOID DeferredContext, PVOID SystemArgument1, PVOID SystemArgument2) {
    UNREFERENCED_PARAMETER(Dpc);
    UNREFERENCED_PARAMETER(DeferredContext);
    AsmVmxVmcall(VMCALL_ENABLE_MOV_TO_DEBUG_REGS_EXITING, 0, 0, 0);
    KeSignalCallDpcSynchronize(SystemArgument2);
    KeSignalCallDpcDone(SystemArgument1);
}

VOID
DpcRoutineEnableMovControlRegisterExitingAllCores(KDPC * Dpc, PDEBUGGER_EVENT Event, PVOID SystemArgument1, PVOID SystemArgument2) {
    UNREFERENCED_PARAMETER(Dpc);
    AsmVmxVmcall(VMCALL_ENABLE_MOV_TO_CONTROL_REGS_EXITING, Event->OptionalParam1, Event->OptionalParam2, 0);
    KeSignalCallDpcSynchronize(SystemArgument2);
    KeSignalCallDpcDone(SystemArgument1);
}

VOID
DpcRoutineDisableMovControlRegisterExitingAllCores(KDPC * Dpc, PDEBUGGER_EVENT Event, PVOID SystemArgument1, PVOID SystemArgument2) {
    UNREFERENCED_PARAMETER(Dpc);
    AsmVmxVmcall(VMCALL_DISABLE_MOV_TO_CONTROL_REGS_EXITING, Event->OptionalParam1, Event->OptionalParam2, 0);
    KeSignalCallDpcSynchronize(SystemArgument2);
    KeSignalCallDpcDone(SystemArgument1);
}

VOID
DpcRoutineDisableMovDebigRegisterExitingAllCores(KDPC * Dpc, PVOID DeferredContext, PVOID SystemArgument1, PVOID SystemArgument2) {
    UNREFERENCED_PARAMETER(Dpc);
    UNREFERENCED_PARAMETER(DeferredContext);
    AsmVmxVmcall(VMCALL_DISABLE_MOV_TO_DEBUG_REGS_EXITING, 0, 0, 0);
    KeSignalCallDpcSynchronize(SystemArgument2);
    KeSignalCallDpcDone(SystemArgument1);
}

VOID
DpcRoutineSetEnableExternalInterruptExitingOnAllCores(KDPC * Dpc, PVOID DeferredContext, PVOID SystemArgument1, PVOID SystemArgument2) {
    UNREFERENCED_PARAMETER(Dpc);
    UNREFERENCED_PARAMETER(DeferredContext);
    AsmVmxVmcall(VMCALL_ENABLE_EXTERNAL_INTERRUPT_EXITING, 0, 0, 0);
    KeSignalCallDpcSynchronize(SystemArgument2);
    KeSignalCallDpcDone(SystemArgument1);
}

VOID
DpcRoutineSetDisableExternalInterruptExitingOnlyOnClearingInterruptEventsOnAllCores(KDPC * Dpc,
                                                                                    PVOID  DeferredContext,
                                                                                    PVOID  SystemArgument1,
                                                                                    PVOID  SystemArgument2) {
    UNREFERENCED_PARAMETER(Dpc);
    UNREFERENCED_PARAMETER(DeferredContext);
    AsmVmxVmcall(VMCALL_DISABLE_EXTERNAL_INTERRUPT_EXITING_ONLY_TO_CLEAR_INTERRUPT_COMMANDS, 0, 0, 0);
    KeSignalCallDpcSynchronize(SystemArgument2);
    KeSignalCallDpcDone(SystemArgument1);
}

VOID
DpcRoutineChangeIoBitmapOnAllCores(KDPC * Dpc, PVOID DeferredContext, PVOID SystemArgument1, PVOID SystemArgument2) {
    UNREFERENCED_PARAMETER(Dpc);
    AsmVmxVmcall(VMCALL_CHANGE_IO_BITMAP, DeferredContext, 0, 0);
    KeSignalCallDpcSynchronize(SystemArgument2);
    KeSignalCallDpcDone(SystemArgument1);
}

VOID
DpcRoutineResetIoBitmapOnAllCores(KDPC * Dpc, PVOID DeferredContext, PVOID SystemArgument1, PVOID SystemArgument2) {
    UNREFERENCED_PARAMETER(Dpc);
    UNREFERENCED_PARAMETER(DeferredContext);
    AsmVmxVmcall(VMCALL_RESET_IO_BITMAP, NULL, 0, 0);
    KeSignalCallDpcSynchronize(SystemArgument2);
    KeSignalCallDpcDone(SystemArgument1);
}

VOID
DpcRoutineEnableBreakpointOnExceptionBitmapOnAllCores(KDPC * Dpc, PVOID DeferredContext, PVOID SystemArgument1, PVOID SystemArgument2) {
    UNREFERENCED_PARAMETER(Dpc);
    UNREFERENCED_PARAMETER(DeferredContext);
    AsmVmxVmcall(VMCALL_SET_EXCEPTION_BITMAP, EXCEPTION_VECTOR_BREAKPOINT, 0, 0);
    KeSignalCallDpcSynchronize(SystemArgument2);
    KeSignalCallDpcDone(SystemArgument1);
}

VOID
DpcRoutineDisableBreakpointOnExceptionBitmapOnAllCores(KDPC * Dpc, PVOID DeferredContext, PVOID SystemArgument1, PVOID SystemArgument2) {
    UNREFERENCED_PARAMETER(Dpc);
    UNREFERENCED_PARAMETER(DeferredContext);
    AsmVmxVmcall(VMCALL_UNSET_EXCEPTION_BITMAP, EXCEPTION_VECTOR_BREAKPOINT, 0, 0);
    KeSignalCallDpcSynchronize(SystemArgument2);
    KeSignalCallDpcDone(SystemArgument1);
}

VOID
DpcRoutineEnableNmiVmexitOnAllCores(KDPC * Dpc, PVOID DeferredContext, PVOID SystemArgument1, PVOID SystemArgument2) {
    UNREFERENCED_PARAMETER(Dpc);
    UNREFERENCED_PARAMETER(DeferredContext);
    AsmVmxVmcall(VMCALL_SET_VM_EXIT_ON_NMIS, NULL, 0, 0);
    KeSignalCallDpcSynchronize(SystemArgument2);
    KeSignalCallDpcDone(SystemArgument1);
}

VOID
DpcRoutineDisableNmiVmexitOnAllCores(KDPC * Dpc, PVOID DeferredContext, PVOID SystemArgument1, PVOID SystemArgument2) {
    UNREFERENCED_PARAMETER(Dpc);
    UNREFERENCED_PARAMETER(DeferredContext);
    AsmVmxVmcall(VMCALL_UNSET_VM_EXIT_ON_NMIS, NULL, 0, 0);
    KeSignalCallDpcSynchronize(SystemArgument2);
    KeSignalCallDpcDone(SystemArgument1);
}

VOID
DpcRoutineVmExitAndHaltSystemAllCores(KDPC * Dpc, PVOID DeferredContext, PVOID SystemArgument1, PVOID SystemArgument2) {
    UNREFERENCED_PARAMETER(Dpc);
    UNREFERENCED_PARAMETER(DeferredContext);
    AsmVmxVmcall(VMCALL_VM_EXIT_HALT_SYSTEM, 0, 0, 0);
    KeSignalCallDpcSynchronize(SystemArgument2);
    KeSignalCallDpcDone(SystemArgument1);
}

VOID
DpcRoutineEnableDbAndBpExitingOnAllCores(KDPC * Dpc, PVOID DeferredContext, PVOID SystemArgument1, PVOID SystemArgument2) {
    UNREFERENCED_PARAMETER(Dpc);
    UNREFERENCED_PARAMETER(DeferredContext);
    AsmVmxVmcall(VMCALL_SET_EXCEPTION_BITMAP, EXCEPTION_VECTOR_BREAKPOINT, 0, 0);
    AsmVmxVmcall(VMCALL_SET_EXCEPTION_BITMAP, EXCEPTION_VECTOR_DEBUG_BREAKPOINT, 0, 0);
    KeSignalCallDpcSynchronize(SystemArgument2);
    KeSignalCallDpcDone(SystemArgument1);
}

VOID
DpcRoutineDisableDbAndBpExitingOnAllCores(KDPC * Dpc, PVOID DeferredContext, PVOID SystemArgument1, PVOID SystemArgument2) {
    UNREFERENCED_PARAMETER(Dpc);
    UNREFERENCED_PARAMETER(DeferredContext);
    AsmVmxVmcall(VMCALL_UNSET_EXCEPTION_BITMAP, EXCEPTION_VECTOR_BREAKPOINT, 0, 0);
    AsmVmxVmcall(VMCALL_UNSET_EXCEPTION_BITMAP, EXCEPTION_VECTOR_DEBUG_BREAKPOINT, 0, 0);
    KeSignalCallDpcSynchronize(SystemArgument2);
    KeSignalCallDpcDone(SystemArgument1);
}

VOID
DpcRoutineRemoveHookAndInvalidateAllEntriesOnAllCores(KDPC * Dpc, PVOID DeferredContext, PVOID SystemArgument1, PVOID SystemArgument2) {
    UNREFERENCED_PARAMETER(Dpc);
    UNREFERENCED_PARAMETER(DeferredContext);
    AsmVmxVmcall(VMCALL_UNHOOK_ALL_PAGES, NULL, NULL, NULL);
    KeSignalCallDpcSynchronize(SystemArgument2);
    KeSignalCallDpcDone(SystemArgument1);
}

VOID
DpcRoutineRemoveHookAndInvalidateSingleEntryOnAllCores(KDPC * Dpc, PVOID DeferredContext, PVOID SystemArgument1, PVOID SystemArgument2) {
    UNREFERENCED_PARAMETER(Dpc);
    UNREFERENCED_PARAMETER(DeferredContext);
    AsmVmxVmcall(VMCALL_UNHOOK_SINGLE_PAGE, DeferredContext, NULL, NULL);
    KeSignalCallDpcSynchronize(SystemArgument2);
    KeSignalCallDpcDone(SystemArgument1);
}

VOID
DpcRoutineInvalidateEptOnAllCores(KDPC * Dpc, PVOID DeferredContext, PVOID SystemArgument1, PVOID SystemArgument2) {
    UNREFERENCED_PARAMETER(Dpc);
    if (DeferredContext == NULL) {
        AsmVmxVmcall(VMCALL_INVEPT_ALL_CONTEXTS, NULL, NULL, NULL);
    } else {
        AsmVmxVmcall(VMCALL_INVEPT_SINGLE_CONTEXT, DeferredContext, NULL, NULL);
    }
    KeSignalCallDpcSynchronize(SystemArgument2);
    KeSignalCallDpcDone(SystemArgument1);
}

VOID
DpcRoutineInitializeGuest(KDPC * Dpc, PVOID DeferredContext, PVOID SystemArgument1, PVOID SystemArgument2) {
    UNREFERENCED_PARAMETER(Dpc);
    UNREFERENCED_PARAMETER(DeferredContext);
    AsmVmxSaveState();
    KeSignalCallDpcSynchronize(SystemArgument2);
    KeSignalCallDpcDone(SystemArgument1);
}

VOID
DpcRoutineTerminateGuest(KDPC * Dpc, PVOID DeferredContext, PVOID SystemArgument1, PVOID SystemArgument2) {
    UNREFERENCED_PARAMETER(Dpc);
    UNREFERENCED_PARAMETER(DeferredContext);
    if (!VmxTerminate()) {
        LogError("Err, there were an error terminating vmx");
    }
    KeSignalCallDpcSynchronize(SystemArgument2);
    KeSignalCallDpcDone(SystemArgument1);
}
