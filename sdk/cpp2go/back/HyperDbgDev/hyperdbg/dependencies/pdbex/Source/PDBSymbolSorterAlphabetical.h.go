package Source
type (
PdbSymbolSorterAlphabetical interface{
		GetSortedSymbols()(ok bool)//col:161
}

)
func NewPdbSymbolSorterAlphabetical() { return & pdbSymbolSorterAlphabetical{} }
func (p *pdbSymbolSorterAlphabetical)		GetSortedSymbols()(ok bool){//col:161
return true
}

