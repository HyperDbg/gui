package transparency
type (
Transparency interface{
TransparentModeRdtscDiffVmexit()(ok bool)//col:63
TransparentModeRdtscVmexitTracing()(ok bool)//col:99
TransparentModeCpuidTimeStampCounter()(ok bool)//col:141
TransparentModeRdtscEmulationDetection()(ok bool)//col:183
TransparentModeCheckHypervisorPresence()(ok bool)//col:213
TransparentModeCheckRdtscpVmexit()(ok bool)//col:243
}

)
func NewTransparency() { return & transparency{} }
func (t *transparency)TransparentModeRdtscDiffVmexit()(ok bool){//col:63
return true
}

func (t *transparency)TransparentModeRdtscVmexitTracing()(ok bool){//col:99
return true
}

func (t *transparency)TransparentModeCpuidTimeStampCounter()(ok bool){//col:141
return true
}

func (t *transparency)TransparentModeRdtscEmulationDetection()(ok bool){//col:183
return true
}

func (t *transparency)TransparentModeCheckHypervisorPresence()(ok bool){//col:213
return true
}

func (t *transparency)TransparentModeCheckRdtscpVmexit()(ok bool){//col:243
return true
}

