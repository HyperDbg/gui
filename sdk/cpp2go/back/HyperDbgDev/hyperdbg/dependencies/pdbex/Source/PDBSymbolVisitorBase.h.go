package Source
type (
PdbSymbolVisitorBase interface{
		~PDBSymbolVisitorBase()(ok bool)//col:258
}

)
func NewPdbSymbolVisitorBase() { return & pdbSymbolVisitorBase{} }
func (p *pdbSymbolVisitorBase)		~PDBSymbolVisitorBase()(ok bool){//col:258
return true
}

