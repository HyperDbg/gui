package Source
type (
PdbSymbolSorter interface{
		GetSortedSymbols()(ok bool)//col:146
}

)
func NewPdbSymbolSorter() { return & pdbSymbolSorter{} }
func (p *pdbSymbolSorter)		GetSortedSymbols()(ok bool){//col:146
return true
}

