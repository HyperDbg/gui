package broadcast

//back\HyperDbgDev\hyperdbg\hprdbghv\code\debugger\broadcast\DpcRoutines.c.back

type (
	DpcRoutines interface {
		*The function that needs to use this featue ()
(ok bool) //col:116
DpcRoutinePerformVirtualization()(ok bool) //col:146
DpcRoutinePerformWriteMsr()(ok bool)       //col:180
DpcRoutinePerformReadMsr()(ok bool) //col:213
DpcRoutinePerformChangeMsrBitmapReadOnSingleCore()(ok bool) //col:241
DpcRoutinePerformChangeMsrBitmapWriteOnSingleCore()(ok bool) //col:269
DpcRoutinePerformEnableRdtscExitingOnSingleCore()(ok bool) //col:298
DpcRoutinePerformEnableRdpmcExitingOnSingleCore()(ok bool) //col:327
DpcRoutinePerformSetExceptionBitmapOnSingleCore()(ok bool) //col:356
DpcRoutinePerformEnableMovToDebugRegistersExiting()(ok bool) //col:385
DpcRoutinePerformEnableMovToControlRegisterExiting()(ok bool) //col:413
DpcRoutinePerformSetExternalInterruptExitingOnSingleCore()(ok bool) //col:442
DpcRoutinePerformEnableEferSyscallHookOnSingleCore()(ok bool)       //col:471
DpcRoutinePerformChangeIoBitmapOnSingleCore()(ok bool) //col:499
DpcRoutineEnableMovToCr3Exiting()(ok bool) //col:530
DpcRoutineDisableMovToCr3Exiting()(ok bool) //col:561
DpcRoutineEnableEferSyscallEvents()(ok bool)  //col:592
DpcRoutineDisableEferSyscallEvents()(ok bool) //col:623
DpcRoutineWriteMsrToAllCores()(ok bool) //col:658
DpcRoutineReadMsrToAllCores()(ok bool) //col:693
DpcRoutineChangeMsrBitmapReadOnAllCores()(ok bool) //col:723
DpcRoutineResetMsrBitmapReadOnAllCores()(ok bool)   //col:754
DpcRoutineChangeMsrBitmapWriteOnAllCores()(ok bool) //col:784
DpcRoutineResetMsrBitmapWriteOnAllCores()(ok bool) //col:815
DpcRoutineEnableRdtscExitingAllCores()(ok bool) //col:846
DpcRoutineDisableRdtscExitingAllCores()(ok bool) //col:877
DpcRoutineDisableRdtscExitingForClearingTscEventsAllCores()(ok bool) //col:910
DpcRoutineDisableMov2DrExitingForClearingDrEventsAllCores()(ok bool) //col:942
DpcRoutineDisableMov2CrExitingForClearingCrEventsAllCores()(ok bool) //col:973
DpcRoutineEnableRdpmcExitingAllCores()(ok bool) //col:1004
DpcRoutineDisableRdpmcExitingAllCores()(ok bool) //col:1035
DpcRoutineSetExceptionBitmapOnAllCores()(ok bool)   //col:1065
DpcRoutineUnsetExceptionBitmapOnAllCores()(ok bool) //col:1095
DpcRoutineResetExceptionBitmapOnlyOnClearingExceptionEventsOnAllCores()(ok bool) //col:1129
DpcRoutineEnableMovDebigRegisterExitingAllCores()(ok bool) //col:1160
DpcRoutineEnableMovControlRegisterExitingAllCores()(ok bool) //col:1190
DpcRoutineDisableMovControlRegisterExitingAllCores()(ok bool) //col:1220
DpcRoutineDisableMovDebigRegisterExitingAllCores()(ok bool)   //col:1251
DpcRoutineSetEnableExternalInterruptExitingOnAllCores()(ok bool) //col:1282
DpcRoutineSetDisableExternalInterruptExitingOnlyOnClearingInterruptEventsOnAllCores()(ok bool) //col:1316
DpcRoutineChangeIoBitmapOnAllCores()(ok bool) //col:1346
DpcRoutineResetIoBitmapOnAllCores()(ok bool)                     //col:1377
DpcRoutineEnableBreakpointOnExceptionBitmapOnAllCores()(ok bool) //col:1408
DpcRoutineDisableBreakpointOnExceptionBitmapOnAllCores()(ok bool) //col:1439
DpcRoutineEnableNmiVmexitOnAllCores()(ok bool) //col:1470
DpcRoutineDisableNmiVmexitOnAllCores()(ok bool) //col:1501
DpcRoutineVmExitAndHaltSystemAllCores()(ok bool)    //col:1532
DpcRoutineEnableDbAndBpExitingOnAllCores()(ok bool) //col:1568
DpcRoutineDisableDbAndBpExitingOnAllCores()(ok bool) //col:1604
DpcRoutineRemoveHookAndInvalidateAllEntriesOnAllCores()(ok bool) //col:1635
DpcRoutineRemoveHookAndInvalidateSingleEntryOnAllCores()(ok bool) //col:1666
DpcRoutineInvalidateEptOnAllCores()(ok bool) //col:1705
DpcRoutineInitializeGuest()(ok bool)         //col:1736
DpcRoutineTerminateGuest()(ok bool) //col:1770
}
)

func NewDpcRoutines() { return &dpcRoutines{} }

func (d *dpcRoutines) * The function that needs to use this featue ()(ok bool) { //col:116
	/* * The function that needs to use this featue (Routine parameter function) should
	    * have the when it ends :
	    *
	    *              SpinlockUnlock(&OneCoreLock);
	    *
	   NTSTATUS
	   DpcRoutineRunTaskOnSingleCore(UINT32 CoreNumber, PVOID Routine, PVOID DeferredContext)
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

func (d *dpcRoutines) DpcRoutinePerformVirtualization() (ok bool) { //col:146
	/*DpcRoutinePerformVirtualization(KDPC * Dpc, PVOID DeferredContext, PVOID SystemArgument1, PVOID SystemArgument2)
	  {
	      VmxPerformVirtualizationOnSpecificCore();
	      KeSignalCallDpcSynchronize(SystemArgument2);
	      KeSignalCallDpcDone(SystemArgument1);
	      return TRUE;
	  }*/
	return true
}

func (d *dpcRoutines) DpcRoutinePerformWriteMsr() (ok bool) { //col:180
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

func (d *dpcRoutines) DpcRoutinePerformReadMsr() (ok bool) { //col:213
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

func (d *dpcRoutines) DpcRoutinePerformChangeMsrBitmapReadOnSingleCore() (ok bool) { //col:241
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

func (d *dpcRoutines) DpcRoutinePerformChangeMsrBitmapWriteOnSingleCore() (ok bool) { //col:269
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

func (d *dpcRoutines) DpcRoutinePerformEnableRdtscExitingOnSingleCore() (ok bool) { //col:298
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

func (d *dpcRoutines) DpcRoutinePerformEnableRdpmcExitingOnSingleCore() (ok bool) { //col:327
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

func (d *dpcRoutines) DpcRoutinePerformSetExceptionBitmapOnSingleCore() (ok bool) { //col:356
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

func (d *dpcRoutines) DpcRoutinePerformEnableMovToDebugRegistersExiting() (ok bool) { //col:385
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

func (d *dpcRoutines) DpcRoutinePerformEnableMovToControlRegisterExiting() (ok bool) { //col:413
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

func (d *dpcRoutines) DpcRoutinePerformSetExternalInterruptExitingOnSingleCore() (ok bool) { //col:442
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

func (d *dpcRoutines) DpcRoutinePerformEnableEferSyscallHookOnSingleCore() (ok bool) { //col:471
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

func (d *dpcRoutines) DpcRoutinePerformChangeIoBitmapOnSingleCore() (ok bool) { //col:499
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

func (d *dpcRoutines) DpcRoutineEnableMovToCr3Exiting() (ok bool) { //col:530
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

func (d *dpcRoutines) DpcRoutineDisableMovToCr3Exiting() (ok bool) { //col:561
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

func (d *dpcRoutines) DpcRoutineEnableEferSyscallEvents() (ok bool) { //col:592
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

func (d *dpcRoutines) DpcRoutineDisableEferSyscallEvents() (ok bool) { //col:623
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

func (d *dpcRoutines) DpcRoutineWriteMsrToAllCores() (ok bool) { //col:658
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

func (d *dpcRoutines) DpcRoutineReadMsrToAllCores() (ok bool) { //col:693
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

func (d *dpcRoutines) DpcRoutineChangeMsrBitmapReadOnAllCores() (ok bool) { //col:723
	/*DpcRoutineChangeMsrBitmapReadOnAllCores(KDPC * Dpc, PVOID DeferredContext, PVOID SystemArgument1, PVOID SystemArgument2)
	  {
	      UNREFERENCED_PARAMETER(Dpc);
	      AsmVmxVmcall(VMCALL_CHANGE_MSR_BITMAP_READ, DeferredContext, 0, 0);
	      KeSignalCallDpcSynchronize(SystemArgument2);
	      KeSignalCallDpcDone(SystemArgument1);
	  }*/
	return true
}

func (d *dpcRoutines) DpcRoutineResetMsrBitmapReadOnAllCores() (ok bool) { //col:754
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

func (d *dpcRoutines) DpcRoutineChangeMsrBitmapWriteOnAllCores() (ok bool) { //col:784
	/*DpcRoutineChangeMsrBitmapWriteOnAllCores(KDPC * Dpc, PVOID DeferredContext, PVOID SystemArgument1, PVOID SystemArgument2)
	  {
	      UNREFERENCED_PARAMETER(Dpc);
	      AsmVmxVmcall(VMCALL_CHANGE_MSR_BITMAP_WRITE, DeferredContext, 0, 0);
	      KeSignalCallDpcSynchronize(SystemArgument2);
	      KeSignalCallDpcDone(SystemArgument1);
	  }*/
	return true
}

func (d *dpcRoutines) DpcRoutineResetMsrBitmapWriteOnAllCores() (ok bool) { //col:815
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

func (d *dpcRoutines) DpcRoutineEnableRdtscExitingAllCores() (ok bool) { //col:846
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

func (d *dpcRoutines) DpcRoutineDisableRdtscExitingAllCores() (ok bool) { //col:877
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

func (d *dpcRoutines) DpcRoutineDisableRdtscExitingForClearingTscEventsAllCores() (ok bool) { //col:910
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

func (d *dpcRoutines) DpcRoutineDisableMov2DrExitingForClearingDrEventsAllCores() (ok bool) { //col:942
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

func (d *dpcRoutines) DpcRoutineDisableMov2CrExitingForClearingCrEventsAllCores() (ok bool) { //col:973
	/*DpcRoutineDisableMov2CrExitingForClearingCrEventsAllCores(KDPC * Dpc, PDEBUGGER_EVENT Event, PVOID SystemArgument1, PVOID SystemArgument2)
	  {
	      UNREFERENCED_PARAMETER(Dpc);
	      AsmVmxVmcall(VMCALL_DISABLE_MOV_TO_CR_EXITING_ONLY_FOR_CR_EVENTS, Event->OptionalParam1, Event->OptionalParam2, 0);
	      KeSignalCallDpcSynchronize(SystemArgument2);
	      KeSignalCallDpcDone(SystemArgument1);
	  }*/
	return true
}

func (d *dpcRoutines) DpcRoutineEnableRdpmcExitingAllCores() (ok bool) { //col:1004
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

func (d *dpcRoutines) DpcRoutineDisableRdpmcExitingAllCores() (ok bool) { //col:1035
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

func (d *dpcRoutines) DpcRoutineSetExceptionBitmapOnAllCores() (ok bool) { //col:1065
	/*DpcRoutineSetExceptionBitmapOnAllCores(KDPC * Dpc, PVOID DeferredContext, PVOID SystemArgument1, PVOID SystemArgument2)
	  {
	      UNREFERENCED_PARAMETER(Dpc);
	      AsmVmxVmcall(VMCALL_SET_EXCEPTION_BITMAP, DeferredContext, 0, 0);
	      KeSignalCallDpcSynchronize(SystemArgument2);
	      KeSignalCallDpcDone(SystemArgument1);
	  }*/
	return true
}

func (d *dpcRoutines) DpcRoutineUnsetExceptionBitmapOnAllCores() (ok bool) { //col:1095
	/*DpcRoutineUnsetExceptionBitmapOnAllCores(KDPC * Dpc, PVOID DeferredContext, PVOID SystemArgument1, PVOID SystemArgument2)
	  {
	      UNREFERENCED_PARAMETER(Dpc);
	      AsmVmxVmcall(VMCALL_UNSET_EXCEPTION_BITMAP, DeferredContext, 0, 0);
	      KeSignalCallDpcSynchronize(SystemArgument2);
	      KeSignalCallDpcDone(SystemArgument1);
	  }*/
	return true
}

func (d *dpcRoutines) DpcRoutineResetExceptionBitmapOnlyOnClearingExceptionEventsOnAllCores() (ok bool) { //col:1129
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

func (d *dpcRoutines) DpcRoutineEnableMovDebigRegisterExitingAllCores() (ok bool) { //col:1160
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

func (d *dpcRoutines) DpcRoutineEnableMovControlRegisterExitingAllCores() (ok bool) { //col:1190
	/*DpcRoutineEnableMovControlRegisterExitingAllCores(KDPC * Dpc, PDEBUGGER_EVENT Event, PVOID SystemArgument1, PVOID SystemArgument2)
	  {
	      UNREFERENCED_PARAMETER(Dpc);
	      AsmVmxVmcall(VMCALL_ENABLE_MOV_TO_CONTROL_REGS_EXITING, Event->OptionalParam1, Event->OptionalParam2, 0);
	      KeSignalCallDpcSynchronize(SystemArgument2);
	      KeSignalCallDpcDone(SystemArgument1);
	  }*/
	return true
}

func (d *dpcRoutines) DpcRoutineDisableMovControlRegisterExitingAllCores() (ok bool) { //col:1220
	/*DpcRoutineDisableMovControlRegisterExitingAllCores(KDPC * Dpc, PDEBUGGER_EVENT Event, PVOID SystemArgument1, PVOID SystemArgument2)
	  {
	      UNREFERENCED_PARAMETER(Dpc);
	      AsmVmxVmcall(VMCALL_DISABLE_MOV_TO_CONTROL_REGS_EXITING, Event->OptionalParam1, Event->OptionalParam2, 0);
	      KeSignalCallDpcSynchronize(SystemArgument2);
	      KeSignalCallDpcDone(SystemArgument1);
	  }*/
	return true
}

func (d *dpcRoutines) DpcRoutineDisableMovDebigRegisterExitingAllCores() (ok bool) { //col:1251
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

func (d *dpcRoutines) DpcRoutineSetEnableExternalInterruptExitingOnAllCores() (ok bool) { //col:1282
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

func (d *dpcRoutines) DpcRoutineSetDisableExternalInterruptExitingOnlyOnClearingInterruptEventsOnAllCores() (ok bool) { //col:1316
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

func (d *dpcRoutines) DpcRoutineChangeIoBitmapOnAllCores() (ok bool) { //col:1346
	/*DpcRoutineChangeIoBitmapOnAllCores(KDPC * Dpc, PVOID DeferredContext, PVOID SystemArgument1, PVOID SystemArgument2)
	  {
	      UNREFERENCED_PARAMETER(Dpc);
	      AsmVmxVmcall(VMCALL_CHANGE_IO_BITMAP, DeferredContext, 0, 0);
	      KeSignalCallDpcSynchronize(SystemArgument2);
	      KeSignalCallDpcDone(SystemArgument1);
	  }*/
	return true
}

func (d *dpcRoutines) DpcRoutineResetIoBitmapOnAllCores() (ok bool) { //col:1377
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

func (d *dpcRoutines) DpcRoutineEnableBreakpointOnExceptionBitmapOnAllCores() (ok bool) { //col:1408
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

func (d *dpcRoutines) DpcRoutineDisableBreakpointOnExceptionBitmapOnAllCores() (ok bool) { //col:1439
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

func (d *dpcRoutines) DpcRoutineEnableNmiVmexitOnAllCores() (ok bool) { //col:1470
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

func (d *dpcRoutines) DpcRoutineDisableNmiVmexitOnAllCores() (ok bool) { //col:1501
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

func (d *dpcRoutines) DpcRoutineVmExitAndHaltSystemAllCores() (ok bool) { //col:1532
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

func (d *dpcRoutines) DpcRoutineEnableDbAndBpExitingOnAllCores() (ok bool) { //col:1568
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

func (d *dpcRoutines) DpcRoutineDisableDbAndBpExitingOnAllCores() (ok bool) { //col:1604
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

func (d *dpcRoutines) DpcRoutineRemoveHookAndInvalidateAllEntriesOnAllCores() (ok bool) { //col:1635
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

func (d *dpcRoutines) DpcRoutineRemoveHookAndInvalidateSingleEntryOnAllCores() (ok bool) { //col:1666
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

func (d *dpcRoutines) DpcRoutineInvalidateEptOnAllCores() (ok bool) { //col:1705
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

func (d *dpcRoutines) DpcRoutineInitializeGuest() (ok bool) { //col:1736
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

func (d *dpcRoutines) DpcRoutineTerminateGuest() (ok bool) { //col:1770
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
