package vmx


type (
Hv interface{
HvAdjustControls()(ok bool)//col:31
HvSetGuestSelector()(ok bool)//col:59
HvHandleCpuid()(ok bool)//col:160
HvHandleControlRegisterAccess()(ok bool)//col:287
HvFillGuestSelectorData()(ok bool)//col:317
HvResumeToNextInstruction()(ok bool)//col:338
HvSetMonitorTrapFlag()(ok bool)//col:370
HvSetLoadDebugControls()(ok bool)//col:402
HvSetSaveDebugControls()(ok bool)//col:434
HvRestoreRegisters()(ok bool)//col:479
HvSetPmcVmexit()(ok bool)//col:512
HvSetMovControlRegsExiting()(ok bool)//col:528
HvSetMovToCr3Vmexit()(ok bool)//col:542
HvWriteExceptionBitmap()(ok bool)//col:560
HvReadExceptionBitmap()(ok bool)//col:580
HvSetInterruptWindowExiting()(ok bool)//col:615
HvSetNmiWindowExiting()(ok bool)//col:650
HvHandleMovDebugRegister()(ok bool)//col:859
HvSetNmiExiting()(ok bool)//col:896
HvSetVmxPreemptionTimerExiting()(ok bool)//col:928
HvSetExceptionBitmap()(ok bool)//col:945
HvUnsetExceptionBitmap()(ok bool)//col:962
HvSetExternalInterruptExiting()(ok bool)//col:978
HvSetRdtscExiting()(ok bool)//col:991
HvSetMovDebugRegsExiting()(ok bool)//col:1004
}






)

func NewHv() { return & hv{} }

func (h *hv)HvAdjustControls()(ok bool){//col:31






return true
}







func (h *hv)HvSetGuestSelector()(ok bool){//col:59














return true
}







func (h *hv)HvHandleCpuid()(ok bool){//col:160



























return true
}







func (h *hv)HvHandleControlRegisterAccess()(ok bool){//col:287










































































return true
}







func (h *hv)HvFillGuestSelectorData()(ok bool){//col:317















return true
}







func (h *hv)HvResumeToNextInstruction()(ok bool){//col:338










return true
}







func (h *hv)HvSetMonitorTrapFlag()(ok bool){//col:370














return true
}







func (h *hv)HvSetLoadDebugControls()(ok bool){//col:402














return true
}







func (h *hv)HvSetSaveDebugControls()(ok bool){//col:434














return true
}







func (h *hv)HvRestoreRegisters()(ok bool){//col:479



















return true
}







func (h *hv)HvSetPmcVmexit()(ok bool){//col:512














return true
}







func (h *hv)HvSetMovControlRegsExiting()(ok bool){//col:528




return true
}







func (h *hv)HvSetMovToCr3Vmexit()(ok bool){//col:542




return true
}







func (h *hv)HvWriteExceptionBitmap()(ok bool){//col:560




return true
}







func (h *hv)HvReadExceptionBitmap()(ok bool){//col:580






return true
}







func (h *hv)HvSetInterruptWindowExiting()(ok bool){//col:615














return true
}







func (h *hv)HvSetNmiWindowExiting()(ok bool){//col:650














return true
}







func (h *hv)HvHandleMovDebugRegister()(ok bool){//col:859












































































































return true
}







func (h *hv)HvSetNmiExiting()(ok bool){//col:896



















return true
}







func (h *hv)HvSetVmxPreemptionTimerExiting()(ok bool){//col:928














return true
}







func (h *hv)HvSetExceptionBitmap()(ok bool){//col:945




return true
}







func (h *hv)HvUnsetExceptionBitmap()(ok bool){//col:962




return true
}







func (h *hv)HvSetExternalInterruptExiting()(ok bool){//col:978




return true
}







func (h *hv)HvSetRdtscExiting()(ok bool){//col:991




return true
}







func (h *hv)HvSetMovDebugRegsExiting()(ok bool){//col:1004




return true
}









