package Source
type (
HyperDbgIntegration interface{
#pragma comment()(ok bool)//col:36
pdbex_set_logging_method_export()(ok bool)//col:42
ExtractBits()(ok bool)//col:49
SymSeparateTo64BitValue()(ok bool)//col:62
}

)
func NewHyperDbgIntegration() { return & hyperDbgIntegration{} }
func (h *hyperDbgIntegration)#pragma comment()(ok bool){//col:36
return true
}

func (h *hyperDbgIntegration)pdbex_set_logging_method_export()(ok bool){//col:42
return true
}

func (h *hyperDbgIntegration)ExtractBits()(ok bool){//col:49
return true
}

func (h *hyperDbgIntegration)SymSeparateTo64BitValue()(ok bool){//col:62
return true
}

