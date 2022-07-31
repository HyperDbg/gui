package ept
type (
	Ept interface {
		EptCheckFeatures() (ok bool)                      //col:58
		EptBuildMtrrMap() (ok bool)                       //col:132
		EptGetPml1Entry() (ok bool)                       //col:193
		EptGetPml2Entry() (ok bool)                       //col:222
		EptSplitLargePage() (ok bool)                     //col:333
		EptSetupPML2Entry() (ok bool)                     //col:427
		EptAllocateAndCreateIdentityPageTable() (ok bool) //col:544
		EptLogicalProcessorInitialize() (ok bool)         //col:600
		EptHandlePageHookExit() (ok bool)                 //col:697
		EptHandleEptViolation() (ok bool)                 //col:730
		EptHandleMonitorTrapFlag() (ok bool)              //col:745
		EptHandleMisconfiguration() (ok bool)             //col:763
		EptSetPML1AndInvalidateTLB() (ok bool)            //col:808
	}
)
func NewEpt() { return &ept{} }
func (e *ept) EptCheckFeatures() (ok bool) { //col:58
	return true
}

func (e *ept) EptBuildMtrrMap() (ok bool) { //col:132
	return true
}

func (e *ept) EptGetPml1Entry() (ok bool) { //col:193
	return true
}

func (e *ept) EptGetPml2Entry() (ok bool) { //col:222
	return true
}

func (e *ept) EptSplitLargePage() (ok bool) { //col:333
	return true
}

func (e *ept) EptSetupPML2Entry() (ok bool) { //col:427
	return true
}

func (e *ept) EptAllocateAndCreateIdentityPageTable() (ok bool) { //col:544
	return true
}

func (e *ept) EptLogicalProcessorInitialize() (ok bool) { //col:600
	return true
}

func (e *ept) EptHandlePageHookExit() (ok bool) { //col:697
	return true
}

func (e *ept) EptHandleEptViolation() (ok bool) { //col:730
	return true
}

func (e *ept) EptHandleMonitorTrapFlag() (ok bool) { //col:745
	return true
}

func (e *ept) EptHandleMisconfiguration() (ok bool) { //col:763
	return true
}

func (e *ept) EptSetPML1AndInvalidateTLB() (ok bool) { //col:808
	return true
}

