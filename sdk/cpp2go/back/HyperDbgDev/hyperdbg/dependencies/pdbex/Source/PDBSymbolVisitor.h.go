package Source
type (
PdbSymbolVisitor interface{
		PDBSymbolVisitor()(ok bool)//col:372
}

)
func NewPdbSymbolVisitor() { return & pdbSymbolVisitor{} }
func (p *pdbSymbolVisitor)		PDBSymbolVisitor()(ok bool){//col:372
return true
}

