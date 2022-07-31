package vmx


type (
ProtectedHv interface{
ProtectedHvChangeExceptionBitmapWithIntegrityCheck()(ok bool)//col:109
ProtectedHvSetExceptionBitmap()(ok bool)//col:142
ProtectedHvUnsetExceptionBitmap()(ok bool)//col:175
ProtectedHvResetExceptionBitmapToClearEvents()(ok bool)//col:194
ProtectedHvRemoveUndefinedInstructionForDisablingSyscallSysretCommands()(ok bool)//col:223
ProtectedHvApplySetExternalInterruptExiting()(ok bool)//col:315
ProtectedHvSetExternalInterruptExiting()(ok bool)//col:328
ProtectedHvExternalInterruptExitingForDisablingInterruptCommands()(ok bool)//col:340
ProtectedHvSetTscVmexit()(ok bool)//col:418
ProtectedHvSetMovDebugRegsVmexit()(ok bool)//col:497
ProtectedHvSetMovToCrVmexit()(ok bool)//col:539
ProtectedHvSetMovControlRegsVmexit()(ok bool)//col:590
ProtectedHvSetMovToCr3Vmexit()(ok bool)//col:657
ProtectedHvSetRdtscExiting()(ok bool)//col:670
ProtectedHvDisableRdtscExitingForDisablingTscCommands()(ok bool)//col:682
ProtectedHvSetMovDebugRegsExiting()(ok bool)//col:695
ProtectedHvDisableMovDebugRegsExitingForDisablingDrCommands()(ok bool)//col:707
ProtectedHvDisableMovControlRegsExitingForDisablingCrCommands()(ok bool)//col:721
ProtectedHvSetMov2Cr3Exiting()(ok bool)//col:734
ProtectedHvSetMov2CrExiting()(ok bool)//col:749
}
















)

func NewProtectedHv() { return & protectedHv{} }

func (p *protectedHv)ProtectedHvChangeExceptionBitmapWithIntegrityCheck()(ok bool){//col:109



































return true
}

















func (p *protectedHv)ProtectedHvSetExceptionBitmap()(ok bool){//col:142














return true
}

















func (p *protectedHv)ProtectedHvUnsetExceptionBitmap()(ok bool){//col:175














return true
}

















func (p *protectedHv)ProtectedHvResetExceptionBitmapToClearEvents()(ok bool){//col:194





return true
}

















func (p *protectedHv)ProtectedHvRemoveUndefinedInstructionForDisablingSyscallSysretCommands()(ok bool){//col:223







return true
}

















func (p *protectedHv)ProtectedHvApplySetExternalInterruptExiting()(ok bool){//col:315




































return true
}

















func (p *protectedHv)ProtectedHvSetExternalInterruptExiting()(ok bool){//col:328




return true
}

















func (p *protectedHv)ProtectedHvExternalInterruptExitingForDisablingInterruptCommands()(ok bool){//col:340




return true
}

















func (p *protectedHv)ProtectedHvSetTscVmexit()(ok bool){//col:418






























return true
}

















func (p *protectedHv)ProtectedHvSetMovDebugRegsVmexit()(ok bool){//col:497






























return true
}

















func (p *protectedHv)ProtectedHvSetMovToCrVmexit()(ok bool){//col:539





























return true
}

















func (p *protectedHv)ProtectedHvSetMovControlRegsVmexit()(ok bool){//col:590

















return true
}

















func (p *protectedHv)ProtectedHvSetMovToCr3Vmexit()(ok bool){//col:657


























return true
}

















func (p *protectedHv)ProtectedHvSetRdtscExiting()(ok bool){//col:670




return true
}

















func (p *protectedHv)ProtectedHvDisableRdtscExitingForDisablingTscCommands()(ok bool){//col:682




return true
}

















func (p *protectedHv)ProtectedHvSetMovDebugRegsExiting()(ok bool){//col:695




return true
}

















func (p *protectedHv)ProtectedHvDisableMovDebugRegsExitingForDisablingDrCommands()(ok bool){//col:707




return true
}

















func (p *protectedHv)ProtectedHvDisableMovControlRegsExitingForDisablingCrCommands()(ok bool){//col:721




return true
}

















func (p *protectedHv)ProtectedHvSetMov2Cr3Exiting()(ok bool){//col:734




return true
}

















func (p *protectedHv)ProtectedHvSetMov2CrExiting()(ok bool){//col:749




return true
}



















