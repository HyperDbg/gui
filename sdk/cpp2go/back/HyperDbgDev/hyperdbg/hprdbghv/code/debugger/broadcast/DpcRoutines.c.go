package broadcast


type (
DpcRoutines interface{
DpcRoutineRunTaskOnSingleCore()(ok bool)//col:116
DpcRoutinePerformVirtualization()(ok bool)//col:147
DpcRoutinePerformWriteMsr()(ok bool)//col:182
DpcRoutinePerformReadMsr()(ok bool)//col:216
DpcRoutinePerformChangeMsrBitmapReadOnSingleCore()(ok bool)//col:245
DpcRoutinePerformChangeMsrBitmapWriteOnSingleCore()(ok bool)//col:274
DpcRoutinePerformEnableRdtscExitingOnSingleCore()(ok bool)//col:304
DpcRoutinePerformEnableRdpmcExitingOnSingleCore()(ok bool)//col:334
DpcRoutinePerformSetExceptionBitmapOnSingleCore()(ok bool)//col:364
DpcRoutinePerformEnableMovToDebugRegistersExiting()(ok bool)//col:394
DpcRoutinePerformEnableMovToControlRegisterExiting()(ok bool)//col:423
DpcRoutinePerformSetExternalInterruptExitingOnSingleCore()(ok bool)//col:453
DpcRoutinePerformEnableEferSyscallHookOnSingleCore()(ok bool)//col:483
DpcRoutinePerformChangeIoBitmapOnSingleCore()(ok bool)//col:512
DpcRoutineEnableMovToCr3Exiting()(ok bool)//col:544
DpcRoutineDisableMovToCr3Exiting()(ok bool)//col:576
DpcRoutineEnableEferSyscallEvents()(ok bool)//col:608
DpcRoutineDisableEferSyscallEvents()(ok bool)//col:640
DpcRoutineWriteMsrToAllCores()(ok bool)//col:676
DpcRoutineReadMsrToAllCores()(ok bool)//col:712
DpcRoutineChangeMsrBitmapReadOnAllCores()(ok bool)//col:743
DpcRoutineResetMsrBitmapReadOnAllCores()(ok bool)//col:775
DpcRoutineChangeMsrBitmapWriteOnAllCores()(ok bool)//col:806
DpcRoutineResetMsrBitmapWriteOnAllCores()(ok bool)//col:838
DpcRoutineEnableRdtscExitingAllCores()(ok bool)//col:870
DpcRoutineDisableRdtscExitingAllCores()(ok bool)//col:902
DpcRoutineDisableRdtscExitingForClearingTscEventsAllCores()(ok bool)//col:936
DpcRoutineDisableMov2DrExitingForClearingDrEventsAllCores()(ok bool)//col:969
DpcRoutineDisableMov2CrExitingForClearingCrEventsAllCores()(ok bool)//col:1001
DpcRoutineEnableRdpmcExitingAllCores()(ok bool)//col:1033
DpcRoutineDisableRdpmcExitingAllCores()(ok bool)//col:1065
DpcRoutineSetExceptionBitmapOnAllCores()(ok bool)//col:1096
DpcRoutineUnsetExceptionBitmapOnAllCores()(ok bool)//col:1127
DpcRoutineResetExceptionBitmapOnlyOnClearingExceptionEventsOnAllCores()(ok bool)//col:1162
DpcRoutineEnableMovDebigRegisterExitingAllCores()(ok bool)//col:1194
DpcRoutineEnableMovControlRegisterExitingAllCores()(ok bool)//col:1225
DpcRoutineDisableMovControlRegisterExitingAllCores()(ok bool)//col:1256
DpcRoutineDisableMovDebigRegisterExitingAllCores()(ok bool)//col:1288
DpcRoutineSetEnableExternalInterruptExitingOnAllCores()(ok bool)//col:1320
DpcRoutineSetDisableExternalInterruptExitingOnlyOnClearingInterruptEventsOnAllCores()(ok bool)//col:1355
DpcRoutineChangeIoBitmapOnAllCores()(ok bool)//col:1386
DpcRoutineResetIoBitmapOnAllCores()(ok bool)//col:1418
DpcRoutineEnableBreakpointOnExceptionBitmapOnAllCores()(ok bool)//col:1450
DpcRoutineDisableBreakpointOnExceptionBitmapOnAllCores()(ok bool)//col:1482
DpcRoutineEnableNmiVmexitOnAllCores()(ok bool)//col:1514
DpcRoutineDisableNmiVmexitOnAllCores()(ok bool)//col:1546
DpcRoutineVmExitAndHaltSystemAllCores()(ok bool)//col:1578
DpcRoutineEnableDbAndBpExitingOnAllCores()(ok bool)//col:1615
DpcRoutineDisableDbAndBpExitingOnAllCores()(ok bool)//col:1652
DpcRoutineRemoveHookAndInvalidateAllEntriesOnAllCores()(ok bool)//col:1684
DpcRoutineRemoveHookAndInvalidateSingleEntryOnAllCores()(ok bool)//col:1716
DpcRoutineInvalidateEptOnAllCores()(ok bool)//col:1756
DpcRoutineInitializeGuest()(ok bool)//col:1788
DpcRoutineTerminateGuest()(ok bool)//col:1823
}






)

func NewDpcRoutines() { return & dpcRoutines{} }

func (d *dpcRoutines)DpcRoutineRunTaskOnSingleCore()(ok bool){//col:116



























return true
}







func (d *dpcRoutines)DpcRoutinePerformVirtualization()(ok bool){//col:147







return true
}







func (d *dpcRoutines)DpcRoutinePerformWriteMsr()(ok bool){//col:182













return true
}







func (d *dpcRoutines)DpcRoutinePerformReadMsr()(ok bool){//col:216












return true
}







func (d *dpcRoutines)DpcRoutinePerformChangeMsrBitmapReadOnSingleCore()(ok bool){//col:245








return true
}







func (d *dpcRoutines)DpcRoutinePerformChangeMsrBitmapWriteOnSingleCore()(ok bool){//col:274








return true
}







func (d *dpcRoutines)DpcRoutinePerformEnableRdtscExitingOnSingleCore()(ok bool){//col:304









return true
}







func (d *dpcRoutines)DpcRoutinePerformEnableRdpmcExitingOnSingleCore()(ok bool){//col:334









return true
}







func (d *dpcRoutines)DpcRoutinePerformSetExceptionBitmapOnSingleCore()(ok bool){//col:364









return true
}







func (d *dpcRoutines)DpcRoutinePerformEnableMovToDebugRegistersExiting()(ok bool){//col:394









return true
}







func (d *dpcRoutines)DpcRoutinePerformEnableMovToControlRegisterExiting()(ok bool){//col:423








return true
}







func (d *dpcRoutines)DpcRoutinePerformSetExternalInterruptExitingOnSingleCore()(ok bool){//col:453









return true
}







func (d *dpcRoutines)DpcRoutinePerformEnableEferSyscallHookOnSingleCore()(ok bool){//col:483









return true
}







func (d *dpcRoutines)DpcRoutinePerformChangeIoBitmapOnSingleCore()(ok bool){//col:512








return true
}







func (d *dpcRoutines)DpcRoutineEnableMovToCr3Exiting()(ok bool){//col:544








return true
}







func (d *dpcRoutines)DpcRoutineDisableMovToCr3Exiting()(ok bool){//col:576








return true
}







func (d *dpcRoutines)DpcRoutineEnableEferSyscallEvents()(ok bool){//col:608








return true
}







func (d *dpcRoutines)DpcRoutineDisableEferSyscallEvents()(ok bool){//col:640








return true
}







func (d *dpcRoutines)DpcRoutineWriteMsrToAllCores()(ok bool){//col:676











return true
}







func (d *dpcRoutines)DpcRoutineReadMsrToAllCores()(ok bool){//col:712











return true
}







func (d *dpcRoutines)DpcRoutineChangeMsrBitmapReadOnAllCores()(ok bool){//col:743







return true
}







func (d *dpcRoutines)DpcRoutineResetMsrBitmapReadOnAllCores()(ok bool){//col:775








return true
}







func (d *dpcRoutines)DpcRoutineChangeMsrBitmapWriteOnAllCores()(ok bool){//col:806







return true
}







func (d *dpcRoutines)DpcRoutineResetMsrBitmapWriteOnAllCores()(ok bool){//col:838








return true
}







func (d *dpcRoutines)DpcRoutineEnableRdtscExitingAllCores()(ok bool){//col:870








return true
}







func (d *dpcRoutines)DpcRoutineDisableRdtscExitingAllCores()(ok bool){//col:902








return true
}







func (d *dpcRoutines)DpcRoutineDisableRdtscExitingForClearingTscEventsAllCores()(ok bool){//col:936








return true
}







func (d *dpcRoutines)DpcRoutineDisableMov2DrExitingForClearingDrEventsAllCores()(ok bool){//col:969








return true
}







func (d *dpcRoutines)DpcRoutineDisableMov2CrExitingForClearingCrEventsAllCores()(ok bool){//col:1001







return true
}







func (d *dpcRoutines)DpcRoutineEnableRdpmcExitingAllCores()(ok bool){//col:1033








return true
}







func (d *dpcRoutines)DpcRoutineDisableRdpmcExitingAllCores()(ok bool){//col:1065








return true
}







func (d *dpcRoutines)DpcRoutineSetExceptionBitmapOnAllCores()(ok bool){//col:1096







return true
}







func (d *dpcRoutines)DpcRoutineUnsetExceptionBitmapOnAllCores()(ok bool){//col:1127







return true
}







func (d *dpcRoutines)DpcRoutineResetExceptionBitmapOnlyOnClearingExceptionEventsOnAllCores()(ok bool){//col:1162










return true
}







func (d *dpcRoutines)DpcRoutineEnableMovDebigRegisterExitingAllCores()(ok bool){//col:1194








return true
}







func (d *dpcRoutines)DpcRoutineEnableMovControlRegisterExitingAllCores()(ok bool){//col:1225







return true
}







func (d *dpcRoutines)DpcRoutineDisableMovControlRegisterExitingAllCores()(ok bool){//col:1256







return true
}







func (d *dpcRoutines)DpcRoutineDisableMovDebigRegisterExitingAllCores()(ok bool){//col:1288








return true
}







func (d *dpcRoutines)DpcRoutineSetEnableExternalInterruptExitingOnAllCores()(ok bool){//col:1320








return true
}







func (d *dpcRoutines)DpcRoutineSetDisableExternalInterruptExitingOnlyOnClearingInterruptEventsOnAllCores()(ok bool){//col:1355











return true
}







func (d *dpcRoutines)DpcRoutineChangeIoBitmapOnAllCores()(ok bool){//col:1386







return true
}







func (d *dpcRoutines)DpcRoutineResetIoBitmapOnAllCores()(ok bool){//col:1418








return true
}







func (d *dpcRoutines)DpcRoutineEnableBreakpointOnExceptionBitmapOnAllCores()(ok bool){//col:1450








return true
}







func (d *dpcRoutines)DpcRoutineDisableBreakpointOnExceptionBitmapOnAllCores()(ok bool){//col:1482








return true
}







func (d *dpcRoutines)DpcRoutineEnableNmiVmexitOnAllCores()(ok bool){//col:1514








return true
}







func (d *dpcRoutines)DpcRoutineDisableNmiVmexitOnAllCores()(ok bool){//col:1546








return true
}







func (d *dpcRoutines)DpcRoutineVmExitAndHaltSystemAllCores()(ok bool){//col:1578








return true
}







func (d *dpcRoutines)DpcRoutineEnableDbAndBpExitingOnAllCores()(ok bool){//col:1615









return true
}







func (d *dpcRoutines)DpcRoutineDisableDbAndBpExitingOnAllCores()(ok bool){//col:1652









return true
}







func (d *dpcRoutines)DpcRoutineRemoveHookAndInvalidateAllEntriesOnAllCores()(ok bool){//col:1684








return true
}







func (d *dpcRoutines)DpcRoutineRemoveHookAndInvalidateSingleEntryOnAllCores()(ok bool){//col:1716








return true
}







func (d *dpcRoutines)DpcRoutineInvalidateEptOnAllCores()(ok bool){//col:1756














return true
}







func (d *dpcRoutines)DpcRoutineInitializeGuest()(ok bool){//col:1788








return true
}







func (d *dpcRoutines)DpcRoutineTerminateGuest()(ok bool){//col:1823











return true
}









