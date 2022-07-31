package transparency


type (
Transparency interface{
TransparentModeRdtscDiffVmexit()(ok bool)//col:63
TransparentModeRdtscVmexitTracing()(ok bool)//col:100
TransparentModeCpuidTimeStampCounter()(ok bool)//col:143
TransparentModeRdtscEmulationDetection()(ok bool)//col:186
TransparentModeCheckHypervisorPresence()(ok bool)//col:217
TransparentModeCheckRdtscpVmexit()(ok bool)//col:248
}

)

func NewTransparency() { return & transparency{} }

func (t *transparency)TransparentModeRdtscDiffVmexit()(ok bool){//col:63










return true
}


func (t *transparency)TransparentModeRdtscVmexitTracing()(ok bool){//col:100








return true
}


func (t *transparency)TransparentModeCpuidTimeStampCounter()(ok bool){//col:143




















return true
}


func (t *transparency)TransparentModeRdtscEmulationDetection()(ok bool){//col:186




















return true
}


func (t *transparency)TransparentModeCheckHypervisorPresence()(ok bool){//col:217















return true
}


func (t *transparency)TransparentModeCheckRdtscpVmexit()(ok bool){//col:248















return true
}




