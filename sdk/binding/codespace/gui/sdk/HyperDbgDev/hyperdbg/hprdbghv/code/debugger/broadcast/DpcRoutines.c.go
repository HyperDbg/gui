package broadcast
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbghv\code\debugger\broadcast\DpcRoutines.c.back

type (
DpcRoutines interface{
DpcRoutineRunTaskOnSingleCore()(ok bool)//col:30
DpcRoutinePerformVirtualization()(ok bool)//col:37
DpcRoutinePerformWriteMsr()(ok bool)//col:50
DpcRoutinePerformReadMsr()(ok bool)//col:62
DpcRoutinePerformChangeMsrBitmapReadOnSingleCore()(ok bool)//col:70
DpcRoutinePerformChangeMsrBitmapWriteOnSingleCore()(ok bool)//col:78
DpcRoutinePerformEnableRdtscExitingOnSingleCore()(ok bool)//col:87
DpcRoutinePerformEnableRdpmcExitingOnSingleCore()(ok bool)//col:96
DpcRoutinePerformSetExceptionBitmapOnSingleCore()(ok bool)//col:105
DpcRoutinePerformEnableMovToDebugRegistersExiting()(ok bool)//col:114
DpcRoutinePerformEnableMovToControlRegisterExiting()(ok bool)//col:122
DpcRoutinePerformSetExternalInterruptExitingOnSingleCore()(ok bool)//col:131
DpcRoutinePerformEnableEferSyscallHookOnSingleCore()(ok bool)//col:140
DpcRoutinePerformChangeIoBitmapOnSingleCore()(ok bool)//col:148
DpcRoutineEnableMovToCr3Exiting()(ok bool)//col:156
DpcRoutineDisableMovToCr3Exiting()(ok bool)//col:164
DpcRoutineEnableEferSyscallEvents()(ok bool)//col:172
DpcRoutineDisableEferSyscallEvents()(ok bool)//col:180
DpcRoutineWriteMsrToAllCores()(ok bool)//col:191
DpcRoutineReadMsrToAllCores()(ok bool)//col:202
DpcRoutineChangeMsrBitmapReadOnAllCores()(ok bool)//col:209
DpcRoutineResetMsrBitmapReadOnAllCores()(ok bool)//col:217
DpcRoutineChangeMsrBitmapWriteOnAllCores()(ok bool)//col:224
DpcRoutineResetMsrBitmapWriteOnAllCores()(ok bool)//col:232
DpcRoutineEnableRdtscExitingAllCores()(ok bool)//col:240
DpcRoutineDisableRdtscExitingAllCores()(ok bool)//col:248
DpcRoutineDisableRdtscExitingForClearingTscEventsAllCores()(ok bool)//col:256
DpcRoutineDisableMov2DrExitingForClearingDrEventsAllCores()(ok bool)//col:264
DpcRoutineDisableMov2CrExitingForClearingCrEventsAllCores()(ok bool)//col:271
DpcRoutineEnableRdpmcExitingAllCores()(ok bool)//col:279
DpcRoutineDisableRdpmcExitingAllCores()(ok bool)//col:287
DpcRoutineSetExceptionBitmapOnAllCores()(ok bool)//col:294
DpcRoutineUnsetExceptionBitmapOnAllCores()(ok bool)//col:301
DpcRoutineResetExceptionBitmapOnlyOnClearingExceptionEventsOnAllCores()(ok bool)//col:311
DpcRoutineEnableMovDebigRegisterExitingAllCores()(ok bool)//col:319
DpcRoutineEnableMovControlRegisterExitingAllCores()(ok bool)//col:326
DpcRoutineDisableMovControlRegisterExitingAllCores()(ok bool)//col:333
DpcRoutineDisableMovDebigRegisterExitingAllCores()(ok bool)//col:341
DpcRoutineSetEnableExternalInterruptExitingOnAllCores()(ok bool)//col:349
DpcRoutineSetDisableExternalInterruptExitingOnlyOnClearingInterruptEventsOnAllCores()(ok bool)//col:360
DpcRoutineChangeIoBitmapOnAllCores()(ok bool)//col:367
DpcRoutineResetIoBitmapOnAllCores()(ok bool)//col:375
DpcRoutineEnableBreakpointOnExceptionBitmapOnAllCores()(ok bool)//col:383
DpcRoutineDisableBreakpointOnExceptionBitmapOnAllCores()(ok bool)//col:391
DpcRoutineEnableNmiVmexitOnAllCores()(ok bool)//col:399
DpcRoutineDisableNmiVmexitOnAllCores()(ok bool)//col:407
DpcRoutineVmExitAndHaltSystemAllCores()(ok bool)//col:415
DpcRoutineEnableDbAndBpExitingOnAllCores()(ok bool)//col:424
DpcRoutineDisableDbAndBpExitingOnAllCores()(ok bool)//col:433
DpcRoutineRemoveHookAndInvalidateAllEntriesOnAllCores()(ok bool)//col:441
DpcRoutineRemoveHookAndInvalidateSingleEntryOnAllCores()(ok bool)//col:449
DpcRoutineInvalidateEptOnAllCores()(ok bool)//col:463
DpcRoutineInitializeGuest()(ok bool)//col:471
DpcRoutineTerminateGuest()(ok bool)//col:482
}
)

func NewDpcRoutines() { return & dpcRoutines{} }

func (d *dpcRoutines)DpcRoutineRunTaskOnSingleCore()(ok bool){//col:30
/*DpcRoutineRunTaskOnSingleCore(UINT32 CoreNumber, PVOID Routine, PVOID DeferredContext)
{
    PRKDPC Dpc;
    UINT32 ProcessorCount;
    ProcessorCount = KeQueryActiveProcessorCount(0);
    if (CoreNumber >= ProcessorCount)
    {
        return STATUS_INVALID_PARAMETER;
    }
    Dpc = ExAllocatePoolWithTag(NonPagedPool, sizeof(KDPC), POOLTAG);
    if (!Dpc)
    {
        return STATUS_INSUFFICIENT_RESOURCES;
    }
    KeInitializeDpc(Dpc,            
                    Routine,        
                    DeferredContext 
    );
    KeSetTargetProcessorDpc(Dpc, CoreNumber);
    if (!SpinlockTryLock(&OneCoreLock))
    {
        ExFreePoolWithTag(Dpc, POOLTAG);
        return STATUS_UNSUCCESSFUL;
    }
    KeInsertQueueDpc(Dpc, NULL, NULL);
    SpinlockLock(&OneCoreLock);
    SpinlockUnlock(&OneCoreLock);
    ExFreePoolWithTag(Dpc, POOLTAG);
    return STATUS_SUCCESS;
}*/
return true
}

func (d *dpcRoutines)DpcRoutinePerformVirtualization()(ok bool){//col:37
/*DpcRoutinePerformVirtualization(KDPC * Dpc, PVOID DeferredContext, PVOID SystemArgument1, PVOID SystemArgument2)
{
    VmxPerformVirtualizationOnSpecificCore();
    KeSignalCallDpcSynchronize(SystemArgument2);
    KeSignalCallDpcDone(SystemArgument1);
    return TRUE;
}*/
return true
}

func (d *dpcRoutines)DpcRoutinePerformWriteMsr()(ok bool){//col:50
/*DpcRoutinePerformWriteMsr(KDPC * Dpc, PVOID DeferredContext, PVOID SystemArgument1, PVOID SystemArgument2)
{
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
}*/
return true
}

func (d *dpcRoutines)DpcRoutinePerformReadMsr()(ok bool){//col:62
/*DpcRoutinePerformReadMsr(KDPC * Dpc, PVOID DeferredContext, PVOID SystemArgument1, PVOID SystemArgument2)
{
    ULONG                       CurrentCore           = KeGetCurrentProcessorNumber();
    VIRTUAL_MACHINE_STATE *     CurrentVmState        = &g_GuestState[CurrentCore];
    PROCESSOR_DEBUGGING_STATE * CurrentDebuggingState = &CurrentVmState->DebuggingState;
    UNREFERENCED_PARAMETER(Dpc);
    UNREFERENCED_PARAMETER(DeferredContext);
    UNREFERENCED_PARAMETER(SystemArgument1);
    UNREFERENCED_PARAMETER(SystemArgument2);
    CurrentDebuggingState->MsrState.Value = __readmsr(CurrentDebuggingState->MsrState.Msr);
    SpinlockUnlock(&OneCoreLock);
}*/
return true
}

func (d *dpcRoutines)DpcRoutinePerformChangeMsrBitmapReadOnSingleCore()(ok bool){//col:70
/*DpcRoutinePerformChangeMsrBitmapReadOnSingleCore(KDPC * Dpc, PVOID DeferredContext, PVOID SystemArgument1, PVOID SystemArgument2)
{
    UNREFERENCED_PARAMETER(Dpc);
    UNREFERENCED_PARAMETER(SystemArgument1);
    UNREFERENCED_PARAMETER(SystemArgument2);
    AsmVmxVmcall(VMCALL_CHANGE_MSR_BITMAP_READ, DeferredContext, 0, 0);
    SpinlockUnlock(&OneCoreLock);
}*/
return true
}

func (d *dpcRoutines)DpcRoutinePerformChangeMsrBitmapWriteOnSingleCore()(ok bool){//col:78
/*DpcRoutinePerformChangeMsrBitmapWriteOnSingleCore(KDPC * Dpc, PVOID DeferredContext, PVOID SystemArgument1, PVOID SystemArgument2)
{
    UNREFERENCED_PARAMETER(Dpc);
    UNREFERENCED_PARAMETER(SystemArgument1);
    UNREFERENCED_PARAMETER(SystemArgument2);
    AsmVmxVmcall(VMCALL_CHANGE_MSR_BITMAP_WRITE, DeferredContext, 0, 0);
    SpinlockUnlock(&OneCoreLock);
}*/
return true
}

func (d *dpcRoutines)DpcRoutinePerformEnableRdtscExitingOnSingleCore()(ok bool){//col:87
/*DpcRoutinePerformEnableRdtscExitingOnSingleCore(KDPC * Dpc, PVOID DeferredContext, PVOID SystemArgument1, PVOID SystemArgument2)
{
    UNREFERENCED_PARAMETER(Dpc);
    UNREFERENCED_PARAMETER(DeferredContext);
    UNREFERENCED_PARAMETER(SystemArgument1);
    UNREFERENCED_PARAMETER(SystemArgument2);
    AsmVmxVmcall(VMCALL_SET_RDTSC_EXITING, 0, 0, 0);
    SpinlockUnlock(&OneCoreLock);
}*/
return true
}

func (d *dpcRoutines)DpcRoutinePerformEnableRdpmcExitingOnSingleCore()(ok bool){//col:96
/*DpcRoutinePerformEnableRdpmcExitingOnSingleCore(KDPC * Dpc, PVOID DeferredContext, PVOID SystemArgument1, PVOID SystemArgument2)
{
    UNREFERENCED_PARAMETER(Dpc);
    UNREFERENCED_PARAMETER(DeferredContext);
    UNREFERENCED_PARAMETER(SystemArgument1);
    UNREFERENCED_PARAMETER(SystemArgument2);
    AsmVmxVmcall(VMCALL_SET_RDPMC_EXITING, 0, 0, 0);
    SpinlockUnlock(&OneCoreLock);
}*/
return true
}

func (d *dpcRoutines)DpcRoutinePerformSetExceptionBitmapOnSingleCore()(ok bool){//col:105
/*DpcRoutinePerformSetExceptionBitmapOnSingleCore(KDPC * Dpc, PVOID DeferredContext, PVOID SystemArgument1, PVOID SystemArgument2)
{
    UNREFERENCED_PARAMETER(Dpc);
    UNREFERENCED_PARAMETER(DeferredContext);
    UNREFERENCED_PARAMETER(SystemArgument1);
    UNREFERENCED_PARAMETER(SystemArgument2);
    AsmVmxVmcall(VMCALL_SET_EXCEPTION_BITMAP, DeferredContext, 0, 0);
    SpinlockUnlock(&OneCoreLock);
}*/
return true
}

func (d *dpcRoutines)DpcRoutinePerformEnableMovToDebugRegistersExiting()(ok bool){//col:114
/*DpcRoutinePerformEnableMovToDebugRegistersExiting(KDPC * Dpc, PVOID DeferredContext, PVOID SystemArgument1, PVOID SystemArgument2)
{
    UNREFERENCED_PARAMETER(Dpc);
    UNREFERENCED_PARAMETER(DeferredContext);
    UNREFERENCED_PARAMETER(SystemArgument1);
    UNREFERENCED_PARAMETER(SystemArgument2);
    AsmVmxVmcall(VMCALL_ENABLE_MOV_TO_DEBUG_REGS_EXITING, 0, 0, 0);
    SpinlockUnlock(&OneCoreLock);
}*/
return true
}

func (d *dpcRoutines)DpcRoutinePerformEnableMovToControlRegisterExiting()(ok bool){//col:122
/*DpcRoutinePerformEnableMovToControlRegisterExiting(KDPC * Dpc, PDEBUGGER_EVENT Event, PVOID SystemArgument1, PVOID SystemArgument2)
{
    UNREFERENCED_PARAMETER(Dpc);
    UNREFERENCED_PARAMETER(SystemArgument1);
    UNREFERENCED_PARAMETER(SystemArgument2);
    AsmVmxVmcall(VMCALL_ENABLE_MOV_TO_CONTROL_REGS_EXITING, Event->OptionalParam1, Event->OptionalParam2, 0);
    SpinlockUnlock(&OneCoreLock);
}*/
return true
}

func (d *dpcRoutines)DpcRoutinePerformSetExternalInterruptExitingOnSingleCore()(ok bool){//col:131
/*DpcRoutinePerformSetExternalInterruptExitingOnSingleCore(KDPC * Dpc, PVOID DeferredContext, PVOID SystemArgument1, PVOID SystemArgument2)
{
    UNREFERENCED_PARAMETER(Dpc);
    UNREFERENCED_PARAMETER(DeferredContext);
    UNREFERENCED_PARAMETER(SystemArgument1);
    UNREFERENCED_PARAMETER(SystemArgument2);
    AsmVmxVmcall(VMCALL_ENABLE_EXTERNAL_INTERRUPT_EXITING, NULL, 0, 0);
    SpinlockUnlock(&OneCoreLock);
}*/
return true
}

func (d *dpcRoutines)DpcRoutinePerformEnableEferSyscallHookOnSingleCore()(ok bool){//col:140
/*DpcRoutinePerformEnableEferSyscallHookOnSingleCore(KDPC * Dpc, PVOID DeferredContext, PVOID SystemArgument1, PVOID SystemArgument2)
{
    UNREFERENCED_PARAMETER(Dpc);
    UNREFERENCED_PARAMETER(DeferredContext);
    UNREFERENCED_PARAMETER(SystemArgument1);
    UNREFERENCED_PARAMETER(SystemArgument2);
    AsmVmxVmcall(VMCALL_ENABLE_SYSCALL_HOOK_EFER, NULL, 0, 0);
    SpinlockUnlock(&OneCoreLock);
}*/
return true
}

func (d *dpcRoutines)DpcRoutinePerformChangeIoBitmapOnSingleCore()(ok bool){//col:148
/*DpcRoutinePerformChangeIoBitmapOnSingleCore(KDPC * Dpc, PVOID DeferredContext, PVOID SystemArgument1, PVOID SystemArgument2)
{
    UNREFERENCED_PARAMETER(Dpc);
    UNREFERENCED_PARAMETER(SystemArgument1);
    UNREFERENCED_PARAMETER(SystemArgument2);
    AsmVmxVmcall(VMCALL_CHANGE_IO_BITMAP, DeferredContext, 0, 0);
    SpinlockUnlock(&OneCoreLock);
}*/
return true
}

func (d *dpcRoutines)DpcRoutineEnableMovToCr3Exiting()(ok bool){//col:156
/*DpcRoutineEnableMovToCr3Exiting(KDPC * Dpc, PVOID DeferredContext, PVOID SystemArgument1, PVOID SystemArgument2)
{
    UNREFERENCED_PARAMETER(Dpc);
    UNREFERENCED_PARAMETER(DeferredContext);
    AsmVmxVmcall(VMCALL_ENABLE_MOV_TO_CR3_EXITING, 0, 0, 0);
    KeSignalCallDpcSynchronize(SystemArgument2);
    KeSignalCallDpcDone(SystemArgument1);
}*/
return true
}

func (d *dpcRoutines)DpcRoutineDisableMovToCr3Exiting()(ok bool){//col:164
/*DpcRoutineDisableMovToCr3Exiting(KDPC * Dpc, PVOID DeferredContext, PVOID SystemArgument1, PVOID SystemArgument2)
{
    UNREFERENCED_PARAMETER(Dpc);
    UNREFERENCED_PARAMETER(DeferredContext);
    AsmVmxVmcall(VMCALL_DISABLE_MOV_TO_CR3_EXITING, 0, 0, 0);
    KeSignalCallDpcSynchronize(SystemArgument2);
    KeSignalCallDpcDone(SystemArgument1);
}*/
return true
}

func (d *dpcRoutines)DpcRoutineEnableEferSyscallEvents()(ok bool){//col:172
/*DpcRoutineEnableEferSyscallEvents(KDPC * Dpc, PVOID DeferredContext, PVOID SystemArgument1, PVOID SystemArgument2)
{
    UNREFERENCED_PARAMETER(Dpc);
    UNREFERENCED_PARAMETER(DeferredContext);
    AsmVmxVmcall(VMCALL_ENABLE_SYSCALL_HOOK_EFER, 0, 0, 0);
    KeSignalCallDpcSynchronize(SystemArgument2);
    KeSignalCallDpcDone(SystemArgument1);
}*/
return true
}

func (d *dpcRoutines)DpcRoutineDisableEferSyscallEvents()(ok bool){//col:180
/*DpcRoutineDisableEferSyscallEvents(KDPC * Dpc, PVOID DeferredContext, PVOID SystemArgument1, PVOID SystemArgument2)
{
    UNREFERENCED_PARAMETER(Dpc);
    UNREFERENCED_PARAMETER(DeferredContext);
    AsmVmxVmcall(VMCALL_DISABLE_SYSCALL_HOOK_EFER, 0, 0, 0);
    KeSignalCallDpcSynchronize(SystemArgument2);
    KeSignalCallDpcDone(SystemArgument1);
}*/
return true
}

func (d *dpcRoutines)DpcRoutineWriteMsrToAllCores()(ok bool){//col:191
/*DpcRoutineWriteMsrToAllCores(KDPC * Dpc, PVOID DeferredContext, PVOID SystemArgument1, PVOID SystemArgument2)
{
    ULONG                       CurrentCore           = KeGetCurrentProcessorNumber();
    VIRTUAL_MACHINE_STATE *     CurrentVmState        = &g_GuestState[CurrentCore];
    PROCESSOR_DEBUGGING_STATE * CurrentDebuggingState = &CurrentVmState->DebuggingState;
    UNREFERENCED_PARAMETER(Dpc);
    UNREFERENCED_PARAMETER(DeferredContext);
    __writemsr(CurrentDebuggingState->MsrState.Msr, CurrentDebuggingState->MsrState.Value);
    KeSignalCallDpcSynchronize(SystemArgument2);
    KeSignalCallDpcDone(SystemArgument1);
}*/
return true
}

func (d *dpcRoutines)DpcRoutineReadMsrToAllCores()(ok bool){//col:202
/*DpcRoutineReadMsrToAllCores(KDPC * Dpc, PVOID DeferredContext, PVOID SystemArgument1, PVOID SystemArgument2)
{
    ULONG                       CurrentCore           = KeGetCurrentProcessorNumber();
    VIRTUAL_MACHINE_STATE *     CurrentVmState        = &g_GuestState[CurrentCore];
    PROCESSOR_DEBUGGING_STATE * CurrentDebuggingState = &CurrentVmState->DebuggingState;
    UNREFERENCED_PARAMETER(Dpc);
    UNREFERENCED_PARAMETER(DeferredContext);
    CurrentDebuggingState->MsrState.Value = __readmsr(CurrentDebuggingState->MsrState.Msr);
    KeSignalCallDpcSynchronize(SystemArgument2);
    KeSignalCallDpcDone(SystemArgument1);
}*/
return true
}

func (d *dpcRoutines)DpcRoutineChangeMsrBitmapReadOnAllCores()(ok bool){//col:209
/*DpcRoutineChangeMsrBitmapReadOnAllCores(KDPC * Dpc, PVOID DeferredContext, PVOID SystemArgument1, PVOID SystemArgument2)
{
    UNREFERENCED_PARAMETER(Dpc);
    AsmVmxVmcall(VMCALL_CHANGE_MSR_BITMAP_READ, DeferredContext, 0, 0);
    KeSignalCallDpcSynchronize(SystemArgument2);
    KeSignalCallDpcDone(SystemArgument1);
}*/
return true
}

func (d *dpcRoutines)DpcRoutineResetMsrBitmapReadOnAllCores()(ok bool){//col:217
/*DpcRoutineResetMsrBitmapReadOnAllCores(KDPC * Dpc, PVOID DeferredContext, PVOID SystemArgument1, PVOID SystemArgument2)
{
    UNREFERENCED_PARAMETER(Dpc);
    UNREFERENCED_PARAMETER(DeferredContext);
    AsmVmxVmcall(VMCALL_RESET_MSR_BITMAP_READ, NULL, 0, 0);
    KeSignalCallDpcSynchronize(SystemArgument2);
    KeSignalCallDpcDone(SystemArgument1);
}*/
return true
}

func (d *dpcRoutines)DpcRoutineChangeMsrBitmapWriteOnAllCores()(ok bool){//col:224
/*DpcRoutineChangeMsrBitmapWriteOnAllCores(KDPC * Dpc, PVOID DeferredContext, PVOID SystemArgument1, PVOID SystemArgument2)
{
    UNREFERENCED_PARAMETER(Dpc);
    AsmVmxVmcall(VMCALL_CHANGE_MSR_BITMAP_WRITE, DeferredContext, 0, 0);
    KeSignalCallDpcSynchronize(SystemArgument2);
    KeSignalCallDpcDone(SystemArgument1);
}*/
return true
}

func (d *dpcRoutines)DpcRoutineResetMsrBitmapWriteOnAllCores()(ok bool){//col:232
/*DpcRoutineResetMsrBitmapWriteOnAllCores(KDPC * Dpc, PVOID DeferredContext, PVOID SystemArgument1, PVOID SystemArgument2)
{
    UNREFERENCED_PARAMETER(Dpc);
    UNREFERENCED_PARAMETER(DeferredContext);
    AsmVmxVmcall(VMCALL_RESET_MSR_BITMAP_WRITE, NULL, 0, 0);
    KeSignalCallDpcSynchronize(SystemArgument2);
    KeSignalCallDpcDone(SystemArgument1);
}*/
return true
}

func (d *dpcRoutines)DpcRoutineEnableRdtscExitingAllCores()(ok bool){//col:240
/*DpcRoutineEnableRdtscExitingAllCores(KDPC * Dpc, PVOID DeferredContext, PVOID SystemArgument1, PVOID SystemArgument2)
{
    UNREFERENCED_PARAMETER(Dpc);
    UNREFERENCED_PARAMETER(DeferredContext);
    AsmVmxVmcall(VMCALL_SET_RDTSC_EXITING, 0, 0, 0);
    KeSignalCallDpcSynchronize(SystemArgument2);
    KeSignalCallDpcDone(SystemArgument1);
}*/
return true
}

func (d *dpcRoutines)DpcRoutineDisableRdtscExitingAllCores()(ok bool){//col:248
/*DpcRoutineDisableRdtscExitingAllCores(KDPC * Dpc, PVOID DeferredContext, PVOID SystemArgument1, PVOID SystemArgument2)
{
    UNREFERENCED_PARAMETER(Dpc);
    UNREFERENCED_PARAMETER(DeferredContext);
    AsmVmxVmcall(VMCALL_UNSET_RDTSC_EXITING, 0, 0, 0);
    KeSignalCallDpcSynchronize(SystemArgument2);
    KeSignalCallDpcDone(SystemArgument1);
}*/
return true
}

func (d *dpcRoutines)DpcRoutineDisableRdtscExitingForClearingTscEventsAllCores()(ok bool){//col:256
/*DpcRoutineDisableRdtscExitingForClearingTscEventsAllCores(KDPC * Dpc, PVOID DeferredContext, PVOID SystemArgument1, PVOID SystemArgument2)
{
    UNREFERENCED_PARAMETER(Dpc);
    UNREFERENCED_PARAMETER(DeferredContext);
    AsmVmxVmcall(VMCALL_DISABLE_RDTSC_EXITING_ONLY_FOR_TSC_EVENTS, 0, 0, 0);
    KeSignalCallDpcSynchronize(SystemArgument2);
    KeSignalCallDpcDone(SystemArgument1);
}*/
return true
}

func (d *dpcRoutines)DpcRoutineDisableMov2DrExitingForClearingDrEventsAllCores()(ok bool){//col:264
/*DpcRoutineDisableMov2DrExitingForClearingDrEventsAllCores(KDPC * Dpc, PVOID DeferredContext, PVOID SystemArgument1, PVOID SystemArgument2)
{
    UNREFERENCED_PARAMETER(Dpc);
    UNREFERENCED_PARAMETER(DeferredContext);
    AsmVmxVmcall(VMCALL_DISABLE_MOV_TO_HW_DR_EXITING_ONLY_FOR_DR_EVENTS, 0, 0, 0);
    KeSignalCallDpcSynchronize(SystemArgument2);
    KeSignalCallDpcDone(SystemArgument1);
}*/
return true
}

func (d *dpcRoutines)DpcRoutineDisableMov2CrExitingForClearingCrEventsAllCores()(ok bool){//col:271
/*DpcRoutineDisableMov2CrExitingForClearingCrEventsAllCores(KDPC * Dpc, PDEBUGGER_EVENT Event, PVOID SystemArgument1, PVOID SystemArgument2)
{
    UNREFERENCED_PARAMETER(Dpc);
    AsmVmxVmcall(VMCALL_DISABLE_MOV_TO_CR_EXITING_ONLY_FOR_CR_EVENTS, Event->OptionalParam1, Event->OptionalParam2, 0);
    KeSignalCallDpcSynchronize(SystemArgument2);
    KeSignalCallDpcDone(SystemArgument1);
}*/
return true
}

func (d *dpcRoutines)DpcRoutineEnableRdpmcExitingAllCores()(ok bool){//col:279
/*DpcRoutineEnableRdpmcExitingAllCores(KDPC * Dpc, PVOID DeferredContext, PVOID SystemArgument1, PVOID SystemArgument2)
{
    UNREFERENCED_PARAMETER(Dpc);
    UNREFERENCED_PARAMETER(DeferredContext);
    AsmVmxVmcall(VMCALL_SET_RDPMC_EXITING, 0, 0, 0);
    KeSignalCallDpcSynchronize(SystemArgument2);
    KeSignalCallDpcDone(SystemArgument1);
}*/
return true
}

func (d *dpcRoutines)DpcRoutineDisableRdpmcExitingAllCores()(ok bool){//col:287
/*DpcRoutineDisableRdpmcExitingAllCores(KDPC * Dpc, PVOID DeferredContext, PVOID SystemArgument1, PVOID SystemArgument2)
{
    UNREFERENCED_PARAMETER(Dpc);
    UNREFERENCED_PARAMETER(DeferredContext);
    AsmVmxVmcall(VMCALL_UNSET_RDPMC_EXITING, 0, 0, 0);
    KeSignalCallDpcSynchronize(SystemArgument2);
    KeSignalCallDpcDone(SystemArgument1);
}*/
return true
}

func (d *dpcRoutines)DpcRoutineSetExceptionBitmapOnAllCores()(ok bool){//col:294
/*DpcRoutineSetExceptionBitmapOnAllCores(KDPC * Dpc, PVOID DeferredContext, PVOID SystemArgument1, PVOID SystemArgument2)
{
    UNREFERENCED_PARAMETER(Dpc);
    AsmVmxVmcall(VMCALL_SET_EXCEPTION_BITMAP, DeferredContext, 0, 0);
    KeSignalCallDpcSynchronize(SystemArgument2);
    KeSignalCallDpcDone(SystemArgument1);
}*/
return true
}

func (d *dpcRoutines)DpcRoutineUnsetExceptionBitmapOnAllCores()(ok bool){//col:301
/*DpcRoutineUnsetExceptionBitmapOnAllCores(KDPC * Dpc, PVOID DeferredContext, PVOID SystemArgument1, PVOID SystemArgument2)
{
    UNREFERENCED_PARAMETER(Dpc);
    AsmVmxVmcall(VMCALL_UNSET_EXCEPTION_BITMAP, DeferredContext, 0, 0);
    KeSignalCallDpcSynchronize(SystemArgument2);
    KeSignalCallDpcDone(SystemArgument1);
}*/
return true
}

func (d *dpcRoutines)DpcRoutineResetExceptionBitmapOnlyOnClearingExceptionEventsOnAllCores()(ok bool){//col:311
/*DpcRoutineResetExceptionBitmapOnlyOnClearingExceptionEventsOnAllCores(KDPC * Dpc,
                                                                      PVOID  DeferredContext,
                                                                      PVOID  SystemArgument1,
                                                                      PVOID  SystemArgument2)
{
    UNREFERENCED_PARAMETER(Dpc);
    AsmVmxVmcall(VMCALL_RESET_EXCEPTION_BITMAP_ONLY_ON_CLEARING_EXCEPTION_EVENTS, DeferredContext, 0, 0);
    KeSignalCallDpcSynchronize(SystemArgument2);
    KeSignalCallDpcDone(SystemArgument1);
}*/
return true
}

func (d *dpcRoutines)DpcRoutineEnableMovDebigRegisterExitingAllCores()(ok bool){//col:319
/*DpcRoutineEnableMovDebigRegisterExitingAllCores(KDPC * Dpc, PVOID DeferredContext, PVOID SystemArgument1, PVOID SystemArgument2)
{
    UNREFERENCED_PARAMETER(Dpc);
    UNREFERENCED_PARAMETER(DeferredContext);
    AsmVmxVmcall(VMCALL_ENABLE_MOV_TO_DEBUG_REGS_EXITING, 0, 0, 0);
    KeSignalCallDpcSynchronize(SystemArgument2);
    KeSignalCallDpcDone(SystemArgument1);
}*/
return true
}

func (d *dpcRoutines)DpcRoutineEnableMovControlRegisterExitingAllCores()(ok bool){//col:326
/*DpcRoutineEnableMovControlRegisterExitingAllCores(KDPC * Dpc, PDEBUGGER_EVENT Event, PVOID SystemArgument1, PVOID SystemArgument2)
{
    UNREFERENCED_PARAMETER(Dpc);
    AsmVmxVmcall(VMCALL_ENABLE_MOV_TO_CONTROL_REGS_EXITING, Event->OptionalParam1, Event->OptionalParam2, 0);
    KeSignalCallDpcSynchronize(SystemArgument2);
    KeSignalCallDpcDone(SystemArgument1);
}*/
return true
}

func (d *dpcRoutines)DpcRoutineDisableMovControlRegisterExitingAllCores()(ok bool){//col:333
/*DpcRoutineDisableMovControlRegisterExitingAllCores(KDPC * Dpc, PDEBUGGER_EVENT Event, PVOID SystemArgument1, PVOID SystemArgument2)
{
    UNREFERENCED_PARAMETER(Dpc);
    AsmVmxVmcall(VMCALL_DISABLE_MOV_TO_CONTROL_REGS_EXITING, Event->OptionalParam1, Event->OptionalParam2, 0);
    KeSignalCallDpcSynchronize(SystemArgument2);
    KeSignalCallDpcDone(SystemArgument1);
}*/
return true
}

func (d *dpcRoutines)DpcRoutineDisableMovDebigRegisterExitingAllCores()(ok bool){//col:341
/*DpcRoutineDisableMovDebigRegisterExitingAllCores(KDPC * Dpc, PVOID DeferredContext, PVOID SystemArgument1, PVOID SystemArgument2)
{
    UNREFERENCED_PARAMETER(Dpc);
    UNREFERENCED_PARAMETER(DeferredContext);
    AsmVmxVmcall(VMCALL_DISABLE_MOV_TO_DEBUG_REGS_EXITING, 0, 0, 0);
    KeSignalCallDpcSynchronize(SystemArgument2);
    KeSignalCallDpcDone(SystemArgument1);
}*/
return true
}

func (d *dpcRoutines)DpcRoutineSetEnableExternalInterruptExitingOnAllCores()(ok bool){//col:349
/*DpcRoutineSetEnableExternalInterruptExitingOnAllCores(KDPC * Dpc, PVOID DeferredContext, PVOID SystemArgument1, PVOID SystemArgument2)
{
    UNREFERENCED_PARAMETER(Dpc);
    UNREFERENCED_PARAMETER(DeferredContext);
    AsmVmxVmcall(VMCALL_ENABLE_EXTERNAL_INTERRUPT_EXITING, 0, 0, 0);
    KeSignalCallDpcSynchronize(SystemArgument2);
    KeSignalCallDpcDone(SystemArgument1);
}*/
return true
}

func (d *dpcRoutines)DpcRoutineSetDisableExternalInterruptExitingOnlyOnClearingInterruptEventsOnAllCores()(ok bool){//col:360
/*DpcRoutineSetDisableExternalInterruptExitingOnlyOnClearingInterruptEventsOnAllCores(KDPC * Dpc,
                                                                                    PVOID  DeferredContext,
                                                                                    PVOID  SystemArgument1,
                                                                                    PVOID  SystemArgument2)
{
    UNREFERENCED_PARAMETER(Dpc);
    UNREFERENCED_PARAMETER(DeferredContext);
    AsmVmxVmcall(VMCALL_DISABLE_EXTERNAL_INTERRUPT_EXITING_ONLY_TO_CLEAR_INTERRUPT_COMMANDS, 0, 0, 0);
    KeSignalCallDpcSynchronize(SystemArgument2);
    KeSignalCallDpcDone(SystemArgument1);
}*/
return true
}

func (d *dpcRoutines)DpcRoutineChangeIoBitmapOnAllCores()(ok bool){//col:367
/*DpcRoutineChangeIoBitmapOnAllCores(KDPC * Dpc, PVOID DeferredContext, PVOID SystemArgument1, PVOID SystemArgument2)
{
    UNREFERENCED_PARAMETER(Dpc);
    AsmVmxVmcall(VMCALL_CHANGE_IO_BITMAP, DeferredContext, 0, 0);
    KeSignalCallDpcSynchronize(SystemArgument2);
    KeSignalCallDpcDone(SystemArgument1);
}*/
return true
}

func (d *dpcRoutines)DpcRoutineResetIoBitmapOnAllCores()(ok bool){//col:375
/*DpcRoutineResetIoBitmapOnAllCores(KDPC * Dpc, PVOID DeferredContext, PVOID SystemArgument1, PVOID SystemArgument2)
{
    UNREFERENCED_PARAMETER(Dpc);
    UNREFERENCED_PARAMETER(DeferredContext);
    AsmVmxVmcall(VMCALL_RESET_IO_BITMAP, NULL, 0, 0);
    KeSignalCallDpcSynchronize(SystemArgument2);
    KeSignalCallDpcDone(SystemArgument1);
}*/
return true
}

func (d *dpcRoutines)DpcRoutineEnableBreakpointOnExceptionBitmapOnAllCores()(ok bool){//col:383
/*DpcRoutineEnableBreakpointOnExceptionBitmapOnAllCores(KDPC * Dpc, PVOID DeferredContext, PVOID SystemArgument1, PVOID SystemArgument2)
{
    UNREFERENCED_PARAMETER(Dpc);
    UNREFERENCED_PARAMETER(DeferredContext);
    AsmVmxVmcall(VMCALL_SET_EXCEPTION_BITMAP, EXCEPTION_VECTOR_BREAKPOINT, 0, 0);
    KeSignalCallDpcSynchronize(SystemArgument2);
    KeSignalCallDpcDone(SystemArgument1);
}*/
return true
}

func (d *dpcRoutines)DpcRoutineDisableBreakpointOnExceptionBitmapOnAllCores()(ok bool){//col:391
/*DpcRoutineDisableBreakpointOnExceptionBitmapOnAllCores(KDPC * Dpc, PVOID DeferredContext, PVOID SystemArgument1, PVOID SystemArgument2)
{
    UNREFERENCED_PARAMETER(Dpc);
    UNREFERENCED_PARAMETER(DeferredContext);
    AsmVmxVmcall(VMCALL_UNSET_EXCEPTION_BITMAP, EXCEPTION_VECTOR_BREAKPOINT, 0, 0);
    KeSignalCallDpcSynchronize(SystemArgument2);
    KeSignalCallDpcDone(SystemArgument1);
}*/
return true
}

func (d *dpcRoutines)DpcRoutineEnableNmiVmexitOnAllCores()(ok bool){//col:399
/*DpcRoutineEnableNmiVmexitOnAllCores(KDPC * Dpc, PVOID DeferredContext, PVOID SystemArgument1, PVOID SystemArgument2)
{
    UNREFERENCED_PARAMETER(Dpc);
    UNREFERENCED_PARAMETER(DeferredContext);
    AsmVmxVmcall(VMCALL_SET_VM_EXIT_ON_NMIS, NULL, 0, 0);
    KeSignalCallDpcSynchronize(SystemArgument2);
    KeSignalCallDpcDone(SystemArgument1);
}*/
return true
}

func (d *dpcRoutines)DpcRoutineDisableNmiVmexitOnAllCores()(ok bool){//col:407
/*DpcRoutineDisableNmiVmexitOnAllCores(KDPC * Dpc, PVOID DeferredContext, PVOID SystemArgument1, PVOID SystemArgument2)
{
    UNREFERENCED_PARAMETER(Dpc);
    UNREFERENCED_PARAMETER(DeferredContext);
    AsmVmxVmcall(VMCALL_UNSET_VM_EXIT_ON_NMIS, NULL, 0, 0);
    KeSignalCallDpcSynchronize(SystemArgument2);
    KeSignalCallDpcDone(SystemArgument1);
}*/
return true
}

func (d *dpcRoutines)DpcRoutineVmExitAndHaltSystemAllCores()(ok bool){//col:415
/*DpcRoutineVmExitAndHaltSystemAllCores(KDPC * Dpc, PVOID DeferredContext, PVOID SystemArgument1, PVOID SystemArgument2)
{
    UNREFERENCED_PARAMETER(Dpc);
    UNREFERENCED_PARAMETER(DeferredContext);
    AsmVmxVmcall(VMCALL_VM_EXIT_HALT_SYSTEM, 0, 0, 0);
    KeSignalCallDpcSynchronize(SystemArgument2);
    KeSignalCallDpcDone(SystemArgument1);
}*/
return true
}

func (d *dpcRoutines)DpcRoutineEnableDbAndBpExitingOnAllCores()(ok bool){//col:424
/*DpcRoutineEnableDbAndBpExitingOnAllCores(KDPC * Dpc, PVOID DeferredContext, PVOID SystemArgument1, PVOID SystemArgument2)
{
    UNREFERENCED_PARAMETER(Dpc);
    UNREFERENCED_PARAMETER(DeferredContext);
    AsmVmxVmcall(VMCALL_SET_EXCEPTION_BITMAP, EXCEPTION_VECTOR_BREAKPOINT, 0, 0);
    AsmVmxVmcall(VMCALL_SET_EXCEPTION_BITMAP, EXCEPTION_VECTOR_DEBUG_BREAKPOINT, 0, 0);
    KeSignalCallDpcSynchronize(SystemArgument2);
    KeSignalCallDpcDone(SystemArgument1);
}*/
return true
}

func (d *dpcRoutines)DpcRoutineDisableDbAndBpExitingOnAllCores()(ok bool){//col:433
/*DpcRoutineDisableDbAndBpExitingOnAllCores(KDPC * Dpc, PVOID DeferredContext, PVOID SystemArgument1, PVOID SystemArgument2)
{
    UNREFERENCED_PARAMETER(Dpc);
    UNREFERENCED_PARAMETER(DeferredContext);
    AsmVmxVmcall(VMCALL_UNSET_EXCEPTION_BITMAP, EXCEPTION_VECTOR_BREAKPOINT, 0, 0);
    AsmVmxVmcall(VMCALL_UNSET_EXCEPTION_BITMAP, EXCEPTION_VECTOR_DEBUG_BREAKPOINT, 0, 0);
    KeSignalCallDpcSynchronize(SystemArgument2);
    KeSignalCallDpcDone(SystemArgument1);
}*/
return true
}

func (d *dpcRoutines)DpcRoutineRemoveHookAndInvalidateAllEntriesOnAllCores()(ok bool){//col:441
/*DpcRoutineRemoveHookAndInvalidateAllEntriesOnAllCores(KDPC * Dpc, PVOID DeferredContext, PVOID SystemArgument1, PVOID SystemArgument2)
{
    UNREFERENCED_PARAMETER(Dpc);
    UNREFERENCED_PARAMETER(DeferredContext);
    AsmVmxVmcall(VMCALL_UNHOOK_ALL_PAGES, NULL, NULL, NULL);
    KeSignalCallDpcSynchronize(SystemArgument2);
    KeSignalCallDpcDone(SystemArgument1);
}*/
return true
}

func (d *dpcRoutines)DpcRoutineRemoveHookAndInvalidateSingleEntryOnAllCores()(ok bool){//col:449
/*DpcRoutineRemoveHookAndInvalidateSingleEntryOnAllCores(KDPC * Dpc, PVOID DeferredContext, PVOID SystemArgument1, PVOID SystemArgument2)
{
    UNREFERENCED_PARAMETER(Dpc);
    UNREFERENCED_PARAMETER(DeferredContext);
    AsmVmxVmcall(VMCALL_UNHOOK_SINGLE_PAGE, DeferredContext, NULL, NULL);
    KeSignalCallDpcSynchronize(SystemArgument2);
    KeSignalCallDpcDone(SystemArgument1);
}*/
return true
}

func (d *dpcRoutines)DpcRoutineInvalidateEptOnAllCores()(ok bool){//col:463
/*DpcRoutineInvalidateEptOnAllCores(KDPC * Dpc, PVOID DeferredContext, PVOID SystemArgument1, PVOID SystemArgument2)
{
    UNREFERENCED_PARAMETER(Dpc);
    if (DeferredContext == NULL)
    {
        AsmVmxVmcall(VMCALL_INVEPT_ALL_CONTEXTS, NULL, NULL, NULL);
    }
    else
    {
        AsmVmxVmcall(VMCALL_INVEPT_SINGLE_CONTEXT, DeferredContext, NULL, NULL);
    }
    KeSignalCallDpcSynchronize(SystemArgument2);
    KeSignalCallDpcDone(SystemArgument1);
}*/
return true
}

func (d *dpcRoutines)DpcRoutineInitializeGuest()(ok bool){//col:471
/*DpcRoutineInitializeGuest(KDPC * Dpc, PVOID DeferredContext, PVOID SystemArgument1, PVOID SystemArgument2)
{
    UNREFERENCED_PARAMETER(Dpc);
    UNREFERENCED_PARAMETER(DeferredContext);
    AsmVmxSaveState();
    KeSignalCallDpcSynchronize(SystemArgument2);
    KeSignalCallDpcDone(SystemArgument1);
}*/
return true
}

func (d *dpcRoutines)DpcRoutineTerminateGuest()(ok bool){//col:482
/*DpcRoutineTerminateGuest(KDPC * Dpc, PVOID DeferredContext, PVOID SystemArgument1, PVOID SystemArgument2)
{
    UNREFERENCED_PARAMETER(Dpc);
    UNREFERENCED_PARAMETER(DeferredContext);
    if (!VmxTerminate())
    {
        LogError("Err, there were an error terminating vmx");
    }
    KeSignalCallDpcSynchronize(SystemArgument2);
    KeSignalCallDpcDone(SystemArgument1);
}*/
return true
}



