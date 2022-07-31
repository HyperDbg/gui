package vmx
type (
	Hv interface {
		HvAdjustControls() (ok bool)               //col:31
		HvSetGuestSelector() (ok bool)             //col:58
		HvHandleCpuid() (ok bool)                  //col:158
		HvHandleControlRegisterAccess() (ok bool)  //col:284
		HvFillGuestSelectorData() (ok bool)        //col:313
		HvResumeToNextInstruction() (ok bool)      //col:333
		HvSetMonitorTrapFlag() (ok bool)           //col:364
		HvSetLoadDebugControls() (ok bool)         //col:395
		HvSetSaveDebugControls() (ok bool)         //col:426
		HvRestoreRegisters() (ok bool)             //col:470
		HvSetPmcVmexit() (ok bool)                 //col:502
		HvSetMovControlRegsExiting() (ok bool)     //col:517
		HvSetMovToCr3Vmexit() (ok bool)            //col:530
		HvWriteExceptionBitmap() (ok bool)         //col:547
		HvReadExceptionBitmap() (ok bool)          //col:566
		HvSetInterruptWindowExiting() (ok bool)    //col:600
		HvSetNmiWindowExiting() (ok bool)          //col:634
		HvHandleMovDebugRegister() (ok bool)       //col:842
		HvSetNmiExiting() (ok bool)                //col:878
		HvSetVmxPreemptionTimerExiting() (ok bool) //col:909
		HvSetExceptionBitmap() (ok bool)           //col:925
		HvUnsetExceptionBitmap() (ok bool)         //col:941
		HvSetExternalInterruptExiting() (ok bool)  //col:956
		HvSetRdtscExiting() (ok bool)              //col:968
		HvSetMovDebugRegsExiting() (ok bool)       //col:980
	}
)
func NewHv() { return &hv{} }
func (h *hv) HvAdjustControls() (ok bool) { //col:31
	return true
}

func (h *hv) HvSetGuestSelector() (ok bool) { //col:58
	return true
}

func (h *hv) HvHandleCpuid() (ok bool) { //col:158
	return true
}

func (h *hv) HvHandleControlRegisterAccess() (ok bool) { //col:284
	return true
}

func (h *hv) HvFillGuestSelectorData() (ok bool) { //col:313
	return true
}

func (h *hv) HvResumeToNextInstruction() (ok bool) { //col:333
	return true
}

func (h *hv) HvSetMonitorTrapFlag() (ok bool) { //col:364
	return true
}

func (h *hv) HvSetLoadDebugControls() (ok bool) { //col:395
	return true
}

func (h *hv) HvSetSaveDebugControls() (ok bool) { //col:426
	return true
}

func (h *hv) HvRestoreRegisters() (ok bool) { //col:470
	return true
}

func (h *hv) HvSetPmcVmexit() (ok bool) { //col:502
	return true
}

func (h *hv) HvSetMovControlRegsExiting() (ok bool) { //col:517
	return true
}

func (h *hv) HvSetMovToCr3Vmexit() (ok bool) { //col:530
	return true
}

func (h *hv) HvWriteExceptionBitmap() (ok bool) { //col:547
	return true
}

func (h *hv) HvReadExceptionBitmap() (ok bool) { //col:566
	return true
}

func (h *hv) HvSetInterruptWindowExiting() (ok bool) { //col:600
	return true
}

func (h *hv) HvSetNmiWindowExiting() (ok bool) { //col:634
	return true
}

func (h *hv) HvHandleMovDebugRegister() (ok bool) { //col:842
	return true
}

func (h *hv) HvSetNmiExiting() (ok bool) { //col:878
	return true
}

func (h *hv) HvSetVmxPreemptionTimerExiting() (ok bool) { //col:909
	return true
}

func (h *hv) HvSetExceptionBitmap() (ok bool) { //col:925
	return true
}

func (h *hv) HvUnsetExceptionBitmap() (ok bool) { //col:941
	return true
}

func (h *hv) HvSetExternalInterruptExiting() (ok bool) { //col:956
	return true
}

func (h *hv) HvSetRdtscExiting() (ok bool) { //col:968
	return true
}

func (h *hv) HvSetMovDebugRegsExiting() (ok bool) { //col:980
	return true
}

