package Source
//back\HyperDbgDev\hyperdbg\dependencies\pdbex\Source\PDBExtractor.h.back

const(
IMAGE_FILE_MACHINE_CHPE_X86 = 0x3A64 //col:12
PDBEX_VERSION_MAJOR = 0 //col:15
PDBEX_VERSION_MINOR = 18 //col:16
PDBEX_VERSION_STRING = "0.18" //col:18
)

type (
PdbExtractor interface{
		int Run()(ok bool)//col:98
}
)

func NewPdbExtractor() { return & pdbExtractor{} }

func (p *pdbExtractor)		int Run()(ok bool){//col:98
/*		int Run(
			int argc,
			char** argv
			);
	private:
		void
		PrintUsage();
		void
		ParseParameters(
			int argc,
			char** argv
			);
		void
		OpenPDBFile();
		void
		PrintTestHeader();
		void
		PrintTestFooter();
		void
		PrintPDBHeader();
		void
		PrintPDBDeclarations();
		void
		PrintPDBDefinitions();
		void
		PrintPDBFunctions();
		void
		DumpAllSymbols();
		void
		DumpOneSymbol();
		void
		DumpAllSymbolsOneByOne();
		void
		CloseOpenFiles();
	private:
		PDB m_PDB;
		Settings m_Settings;
		std::unique_ptr<PDBSymbolSorterBase> m_SymbolSorter;
		std::unique_ptr<PDBHeaderReconstructor> m_HeaderReconstructor;
		std::unique_ptr<PDBSymbolVisitor<UdtFieldDefinition>> m_SymbolVisitor;
};*/
return true
}


