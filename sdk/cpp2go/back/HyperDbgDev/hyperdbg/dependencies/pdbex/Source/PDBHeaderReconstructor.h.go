package Source
type (
PdbHeaderReconstructor interface{
			Settings()(ok bool)//col:289
}

)
func NewPdbHeaderReconstructor() { return & pdbHeaderReconstructor{} }
func (p *pdbHeaderReconstructor)			Settings()(ok bool){//col:289
return true
}

