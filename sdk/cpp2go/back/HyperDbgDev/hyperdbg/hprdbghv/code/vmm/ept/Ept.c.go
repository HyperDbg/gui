package ept


type (
Ept interface{
EptCheckFeatures()(ok bool)//col:58
EptBuildMtrrMap()(ok bool)//col:133
EptGetPml1Entry()(ok bool)//col:195
EptGetPml2Entry()(ok bool)//col:225
EptSplitLargePage()(ok bool)//col:337
EptSetupPML2Entry()(ok bool)//col:432
EptAllocateAndCreateIdentityPageTable()(ok bool)//col:550
EptLogicalProcessorInitialize()(ok bool)//col:607
EptHandlePageHookExit()(ok bool)//col:705
EptHandleEptViolation()(ok bool)//col:739
EptHandleMonitorTrapFlag()(ok bool)//col:755
EptHandleMisconfiguration()(ok bool)//col:774
EptSetPML1AndInvalidateTLB()(ok bool)//col:820
}






)

func NewEpt() { return & ept{} }

func (e *ept)EptCheckFeatures()(ok bool){//col:58































return true
}







func (e *ept)EptBuildMtrrMap()(ok bool){//col:133






























return true
}







func (e *ept)EptGetPml1Entry()(ok bool){//col:195



























return true
}







func (e *ept)EptGetPml2Entry()(ok bool){//col:225














return true
}







func (e *ept)EptSplitLargePage()(ok bool){//col:337

















































return true
}







func (e *ept)EptSetupPML2Entry()(ok bool){//col:432































return true
}







func (e *ept)EptAllocateAndCreateIdentityPageTable()(ok bool){//col:550










































return true
}







func (e *ept)EptLogicalProcessorInitialize()(ok bool){//col:607


















return true
}







func (e *ept)EptHandlePageHookExit()(ok bool){//col:705



























return true
}







func (e *ept)EptHandleEptViolation()(ok bool){//col:739










return true
}







func (e *ept)EptHandleMonitorTrapFlag()(ok bool){//col:755




return true
}







func (e *ept)EptHandleMisconfiguration()(ok bool){//col:774





return true
}







func (e *ept)EptSetPML1AndInvalidateTLB()(ok bool){//col:820


















return true
}









