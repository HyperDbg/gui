package vmx
type (
	ProtectedHv interface {
		ProtectedHvChangeExceptionBitmapWithIntegrityCheck() (ok bool)                     //col:109
		ProtectedHvSetExceptionBitmap() (ok bool)                                          //col:141
		ProtectedHvUnsetExceptionBitmap() (ok bool)                                        //col:173
		ProtectedHvResetExceptionBitmapToClearEvents() (ok bool)                           //col:191
		ProtectedHvRemoveUndefinedInstructionForDisablingSyscallSysretCommands() (ok bool) //col:219
		ProtectedHvApplySetExternalInterruptExiting() (ok bool)                            //col:310
		ProtectedHvSetExternalInterruptExiting() (ok bool)                                 //col:322
		ProtectedHvExternalInterruptExitingForDisablingInterruptCommands() (ok bool)       //col:333
		ProtectedHvSetTscVmexit() (ok bool)                                                //col:410
		ProtectedHvSetMovDebugRegsVmexit() (ok bool)                                       //col:488
		ProtectedHvSetMovToCrVmexit() (ok bool)                                            //col:529
		ProtectedHvSetMovControlRegsVmexit() (ok bool)                                     //col:579
		ProtectedHvSetMovToCr3Vmexit() (ok bool)                                           //col:645
		ProtectedHvSetRdtscExiting() (ok bool)                                             //col:657
		ProtectedHvDisableRdtscExitingForDisablingTscCommands() (ok bool)                  //col:668
		ProtectedHvSetMovDebugRegsExiting() (ok bool)                                      //col:680
		ProtectedHvDisableMovDebugRegsExitingForDisablingDrCommands() (ok bool)            //col:691
		ProtectedHvDisableMovControlRegsExitingForDisablingCrCommands() (ok bool)          //col:704
		ProtectedHvSetMov2Cr3Exiting() (ok bool)                                           //col:716
		ProtectedHvSetMov2CrExiting() (ok bool)                                            //col:730
	}
)
func NewProtectedHv() { return &protectedHv{} }
func (p *protectedHv) ProtectedHvChangeExceptionBitmapWithIntegrityCheck() (ok bool) { //col:109
	return true
}

func (p *protectedHv) ProtectedHvSetExceptionBitmap() (ok bool) { //col:141
	return true
}

func (p *protectedHv) ProtectedHvUnsetExceptionBitmap() (ok bool) { //col:173
	return true
}

func (p *protectedHv) ProtectedHvResetExceptionBitmapToClearEvents() (ok bool) { //col:191
	return true
}

func (p *protectedHv) ProtectedHvRemoveUndefinedInstructionForDisablingSyscallSysretCommands() (ok bool) { //col:219
	return true
}

func (p *protectedHv) ProtectedHvApplySetExternalInterruptExiting() (ok bool) { //col:310
	return true
}

func (p *protectedHv) ProtectedHvSetExternalInterruptExiting() (ok bool) { //col:322
	return true
}

func (p *protectedHv) ProtectedHvExternalInterruptExitingForDisablingInterruptCommands() (ok bool) { //col:333
	return true
}

func (p *protectedHv) ProtectedHvSetTscVmexit() (ok bool) { //col:410
	return true
}

func (p *protectedHv) ProtectedHvSetMovDebugRegsVmexit() (ok bool) { //col:488
	return true
}

func (p *protectedHv) ProtectedHvSetMovToCrVmexit() (ok bool) { //col:529
	return true
}

func (p *protectedHv) ProtectedHvSetMovControlRegsVmexit() (ok bool) { //col:579
	return true
}

func (p *protectedHv) ProtectedHvSetMovToCr3Vmexit() (ok bool) { //col:645
	return true
}

func (p *protectedHv) ProtectedHvSetRdtscExiting() (ok bool) { //col:657
	return true
}

func (p *protectedHv) ProtectedHvDisableRdtscExitingForDisablingTscCommands() (ok bool) { //col:668
	return true
}

func (p *protectedHv) ProtectedHvSetMovDebugRegsExiting() (ok bool) { //col:680
	return true
}

func (p *protectedHv) ProtectedHvDisableMovDebugRegsExitingForDisablingDrCommands() (ok bool) { //col:691
	return true
}

func (p *protectedHv) ProtectedHvDisableMovControlRegsExitingForDisablingCrCommands() (ok bool) { //col:704
	return true
}

func (p *protectedHv) ProtectedHvSetMov2Cr3Exiting() (ok bool) { //col:716
	return true
}

func (p *protectedHv) ProtectedHvSetMov2CrExiting() (ok bool) { //col:730
	return true
}

