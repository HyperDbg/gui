package Source
//back\HyperDbgDev\hyperdbg\dependencies\pdbex\Source\PDBSymbolSorterBase.h.back

type (
PdbSymbolSorterBase interface{
		GetSortedSymbols()(ok bool)//col:29
}
)

func NewPdbSymbolSorterBase() { return & pdbSymbolSorterBase{} }

func (p *pdbSymbolSorterBase)		GetSortedSymbols()(ok bool){//col:29
/*		GetSortedSymbols() = 0;
		virtual
		ImageArchitecture
		GetImageArchitecture() const = 0;
		virtual
		void
		Clear() = 0;
};*/
return true
}



