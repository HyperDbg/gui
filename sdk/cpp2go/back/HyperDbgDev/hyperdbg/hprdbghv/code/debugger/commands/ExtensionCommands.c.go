package commands
type (
	ExtensionCommands interface {
		ExtensionCommandVa2paAndPa2va() (ok bool)                                                      //col:157
		ExtensionCommandPte() (ok bool)                                                                //col:285
		ExtensionCommandChangeAllMsrBitmapReadAllCores() (ok bool)                                     //col:300
		ExtensionCommandResetChangeAllMsrBitmapReadAllCores() (ok bool)                                //col:313
		ExtensionCommandChangeAllMsrBitmapWriteAllCores() (ok bool)                                    //col:327
		ExtensionCommandResetAllMsrBitmapWriteAllCores() (ok bool)                                     //col:340
		ExtensionCommandEnableRdtscExitingAllCores() (ok bool)                                         //col:354
		ExtensionCommandDisableRdtscExitingAllCores() (ok bool)                                        //col:367
		ExtensionCommandDisableRdtscExitingForClearingEventsAllCores() (ok bool)                       //col:380
		ExtensionCommandDisableMov2ControlRegsExitingForClearingEventsAllCores() (ok bool)             //col:394
		ExtensionCommandDisableMov2DebugRegsExitingForClearingEventsAllCores() (ok bool)               //col:407
		ExtensionCommandEnableRdpmcExitingAllCores() (ok bool)                                         //col:421
		ExtensionCommandDisableRdpmcExitingAllCores() (ok bool)                                        //col:434
		ExtensionCommandSetExceptionBitmapAllCores() (ok bool)                                         //col:450
		ExtensionCommandUnsetExceptionBitmapAllCores() (ok bool)                                       //col:466
		ExtensionCommandResetExceptionBitmapAllCores() (ok bool)                                       //col:479
		ExtensionCommandEnableMovControlRegisterExitingAllCores() (ok bool)                            //col:494
		ExtensionCommandDisableMovToControlRegistersExitingAllCores() (ok bool)                        //col:508
		ExtensionCommandEnableMovDebugRegistersExitingAllCores() (ok bool)                             //col:523
		ExtensionCommandDisableMovDebugRegistersExitingAllCores() (ok bool)                            //col:536
		ExtensionCommandSetExternalInterruptExitingAllCores() (ok bool)                                //col:550
		ExtensionCommandUnsetExternalInterruptExitingOnlyOnClearingInterruptEventsAllCores() (ok bool) //col:563
		ExtensionCommandIoBitmapChangeAllCores() (ok bool)                                             //col:577
		ExtensionCommandIoBitmapResetAllCores() (ok bool)                                              //col:590
	}
)
func NewExtensionCommands() { return &extensionCommands{} }
func (e *extensionCommands) ExtensionCommandVa2paAndPa2va() (ok bool) { //col:157
	return true
}

func (e *extensionCommands) ExtensionCommandPte() (ok bool) { //col:285
	return true
}

func (e *extensionCommands) ExtensionCommandChangeAllMsrBitmapReadAllCores() (ok bool) { //col:300
	return true
}

func (e *extensionCommands) ExtensionCommandResetChangeAllMsrBitmapReadAllCores() (ok bool) { //col:313
	return true
}

func (e *extensionCommands) ExtensionCommandChangeAllMsrBitmapWriteAllCores() (ok bool) { //col:327
	return true
}

func (e *extensionCommands) ExtensionCommandResetAllMsrBitmapWriteAllCores() (ok bool) { //col:340
	return true
}

func (e *extensionCommands) ExtensionCommandEnableRdtscExitingAllCores() (ok bool) { //col:354
	return true
}

func (e *extensionCommands) ExtensionCommandDisableRdtscExitingAllCores() (ok bool) { //col:367
	return true
}

func (e *extensionCommands) ExtensionCommandDisableRdtscExitingForClearingEventsAllCores() (ok bool) { //col:380
	return true
}

func (e *extensionCommands) ExtensionCommandDisableMov2ControlRegsExitingForClearingEventsAllCores() (ok bool) { //col:394
	return true
}

func (e *extensionCommands) ExtensionCommandDisableMov2DebugRegsExitingForClearingEventsAllCores() (ok bool) { //col:407
	return true
}

func (e *extensionCommands) ExtensionCommandEnableRdpmcExitingAllCores() (ok bool) { //col:421
	return true
}

func (e *extensionCommands) ExtensionCommandDisableRdpmcExitingAllCores() (ok bool) { //col:434
	return true
}

func (e *extensionCommands) ExtensionCommandSetExceptionBitmapAllCores() (ok bool) { //col:450
	return true
}

func (e *extensionCommands) ExtensionCommandUnsetExceptionBitmapAllCores() (ok bool) { //col:466
	return true
}

func (e *extensionCommands) ExtensionCommandResetExceptionBitmapAllCores() (ok bool) { //col:479
	return true
}

func (e *extensionCommands) ExtensionCommandEnableMovControlRegisterExitingAllCores() (ok bool) { //col:494
	return true
}

func (e *extensionCommands) ExtensionCommandDisableMovToControlRegistersExitingAllCores() (ok bool) { //col:508
	return true
}

func (e *extensionCommands) ExtensionCommandEnableMovDebugRegistersExitingAllCores() (ok bool) { //col:523
	return true
}

func (e *extensionCommands) ExtensionCommandDisableMovDebugRegistersExitingAllCores() (ok bool) { //col:536
	return true
}

func (e *extensionCommands) ExtensionCommandSetExternalInterruptExitingAllCores() (ok bool) { //col:550
	return true
}

func (e *extensionCommands) ExtensionCommandUnsetExternalInterruptExitingOnlyOnClearingInterruptEventsAllCores() (ok bool) { //col:563
	return true
}

func (e *extensionCommands) ExtensionCommandIoBitmapChangeAllCores() (ok bool) { //col:577
	return true
}

func (e *extensionCommands) ExtensionCommandIoBitmapResetAllCores() (ok bool) { //col:590
	return true
}

