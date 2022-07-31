package broadcast
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
	return true
}

func (d *dpcRoutines) DpcRoutinePerformVirtualization() (ok bool) { //col:146
	return true
}

func (d *dpcRoutines) DpcRoutinePerformWriteMsr() (ok bool) { //col:180
	return true
}

func (d *dpcRoutines) DpcRoutinePerformReadMsr() (ok bool) { //col:213
	return true
}

func (d *dpcRoutines) DpcRoutinePerformChangeMsrBitmapReadOnSingleCore() (ok bool) { //col:241
	return true
}

func (d *dpcRoutines) DpcRoutinePerformChangeMsrBitmapWriteOnSingleCore() (ok bool) { //col:269
	return true
}

func (d *dpcRoutines) DpcRoutinePerformEnableRdtscExitingOnSingleCore() (ok bool) { //col:298
	return true
}

func (d *dpcRoutines) DpcRoutinePerformEnableRdpmcExitingOnSingleCore() (ok bool) { //col:327
	return true
}

func (d *dpcRoutines) DpcRoutinePerformSetExceptionBitmapOnSingleCore() (ok bool) { //col:356
	return true
}

func (d *dpcRoutines) DpcRoutinePerformEnableMovToDebugRegistersExiting() (ok bool) { //col:385
	return true
}

func (d *dpcRoutines) DpcRoutinePerformEnableMovToControlRegisterExiting() (ok bool) { //col:413
	return true
}

func (d *dpcRoutines) DpcRoutinePerformSetExternalInterruptExitingOnSingleCore() (ok bool) { //col:442
	return true
}

func (d *dpcRoutines) DpcRoutinePerformEnableEferSyscallHookOnSingleCore() (ok bool) { //col:471
	return true
}

func (d *dpcRoutines) DpcRoutinePerformChangeIoBitmapOnSingleCore() (ok bool) { //col:499
	return true
}

func (d *dpcRoutines) DpcRoutineEnableMovToCr3Exiting() (ok bool) { //col:530
	return true
}

func (d *dpcRoutines) DpcRoutineDisableMovToCr3Exiting() (ok bool) { //col:561
	return true
}

func (d *dpcRoutines) DpcRoutineEnableEferSyscallEvents() (ok bool) { //col:592
	return true
}

func (d *dpcRoutines) DpcRoutineDisableEferSyscallEvents() (ok bool) { //col:623
	return true
}

func (d *dpcRoutines) DpcRoutineWriteMsrToAllCores() (ok bool) { //col:658
	return true
}

func (d *dpcRoutines) DpcRoutineReadMsrToAllCores() (ok bool) { //col:693
	return true
}

func (d *dpcRoutines) DpcRoutineChangeMsrBitmapReadOnAllCores() (ok bool) { //col:723
	return true
}

func (d *dpcRoutines) DpcRoutineResetMsrBitmapReadOnAllCores() (ok bool) { //col:754
	return true
}

func (d *dpcRoutines) DpcRoutineChangeMsrBitmapWriteOnAllCores() (ok bool) { //col:784
	return true
}

func (d *dpcRoutines) DpcRoutineResetMsrBitmapWriteOnAllCores() (ok bool) { //col:815
	return true
}

func (d *dpcRoutines) DpcRoutineEnableRdtscExitingAllCores() (ok bool) { //col:846
	return true
}

func (d *dpcRoutines) DpcRoutineDisableRdtscExitingAllCores() (ok bool) { //col:877
	return true
}

func (d *dpcRoutines) DpcRoutineDisableRdtscExitingForClearingTscEventsAllCores() (ok bool) { //col:910
	return true
}

func (d *dpcRoutines) DpcRoutineDisableMov2DrExitingForClearingDrEventsAllCores() (ok bool) { //col:942
	return true
}

func (d *dpcRoutines) DpcRoutineDisableMov2CrExitingForClearingCrEventsAllCores() (ok bool) { //col:973
	return true
}

func (d *dpcRoutines) DpcRoutineEnableRdpmcExitingAllCores() (ok bool) { //col:1004
	return true
}

func (d *dpcRoutines) DpcRoutineDisableRdpmcExitingAllCores() (ok bool) { //col:1035
	return true
}

func (d *dpcRoutines) DpcRoutineSetExceptionBitmapOnAllCores() (ok bool) { //col:1065
	return true
}

func (d *dpcRoutines) DpcRoutineUnsetExceptionBitmapOnAllCores() (ok bool) { //col:1095
	return true
}

func (d *dpcRoutines) DpcRoutineResetExceptionBitmapOnlyOnClearingExceptionEventsOnAllCores() (ok bool) { //col:1129
	return true
}

func (d *dpcRoutines) DpcRoutineEnableMovDebigRegisterExitingAllCores() (ok bool) { //col:1160
	return true
}

func (d *dpcRoutines) DpcRoutineEnableMovControlRegisterExitingAllCores() (ok bool) { //col:1190
	return true
}

func (d *dpcRoutines) DpcRoutineDisableMovControlRegisterExitingAllCores() (ok bool) { //col:1220
	return true
}

func (d *dpcRoutines) DpcRoutineDisableMovDebigRegisterExitingAllCores() (ok bool) { //col:1251
	return true
}

func (d *dpcRoutines) DpcRoutineSetEnableExternalInterruptExitingOnAllCores() (ok bool) { //col:1282
	return true
}

func (d *dpcRoutines) DpcRoutineSetDisableExternalInterruptExitingOnlyOnClearingInterruptEventsOnAllCores() (ok bool) { //col:1316
	return true
}

func (d *dpcRoutines) DpcRoutineChangeIoBitmapOnAllCores() (ok bool) { //col:1346
	return true
}

func (d *dpcRoutines) DpcRoutineResetIoBitmapOnAllCores() (ok bool) { //col:1377
	return true
}

func (d *dpcRoutines) DpcRoutineEnableBreakpointOnExceptionBitmapOnAllCores() (ok bool) { //col:1408
	return true
}

func (d *dpcRoutines) DpcRoutineDisableBreakpointOnExceptionBitmapOnAllCores() (ok bool) { //col:1439
	return true
}

func (d *dpcRoutines) DpcRoutineEnableNmiVmexitOnAllCores() (ok bool) { //col:1470
	return true
}

func (d *dpcRoutines) DpcRoutineDisableNmiVmexitOnAllCores() (ok bool) { //col:1501
	return true
}

func (d *dpcRoutines) DpcRoutineVmExitAndHaltSystemAllCores() (ok bool) { //col:1532
	return true
}

func (d *dpcRoutines) DpcRoutineEnableDbAndBpExitingOnAllCores() (ok bool) { //col:1568
	return true
}

func (d *dpcRoutines) DpcRoutineDisableDbAndBpExitingOnAllCores() (ok bool) { //col:1604
	return true
}

func (d *dpcRoutines) DpcRoutineRemoveHookAndInvalidateAllEntriesOnAllCores() (ok bool) { //col:1635
	return true
}

func (d *dpcRoutines) DpcRoutineRemoveHookAndInvalidateSingleEntryOnAllCores() (ok bool) { //col:1666
	return true
}

func (d *dpcRoutines) DpcRoutineInvalidateEptOnAllCores() (ok bool) { //col:1705
	return true
}

func (d *dpcRoutines) DpcRoutineInitializeGuest() (ok bool) { //col:1736
	return true
}

func (d *dpcRoutines) DpcRoutineTerminateGuest() (ok bool) { //col:1770
	return true
}

