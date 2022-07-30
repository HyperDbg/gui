package Source
//back\HyperDbgDev\hyperdbg\dependencies\pdbex\Source\PDB.h.back

const(
_CRT_SECURE_NO_WARNINGS =  //col:2
)

type (
Pdb interface{
		PDB()(ok bool)//col:443
}
)

func NewPdb() { return & pdb{} }

func (p *pdb)		PDB()(ok bool){//col:443
/*		PDB();
		PDB(
			IN const CHAR* Path
			);
		~PDB();
		BOOL
		Open(
			IN const CHAR* Path
			);
		BOOL
		IsOpened() const;
		const CHAR*
		GetPath() const;
		VOID
		Close();
		DWORD
		GetMachineType() const;
		CV_CFL_LANG
		GetLanguage() const;
		const SYMBOL*
		GetSymbolByName(
			IN const CHAR* SymbolName
			);
		const SYMBOL*
		GetSymbolByTypeId(
			IN DWORD TypeId
			);
		const SymbolMap&
		GetSymbolMap() const;
		const SymbolNameMap&
		GetSymbolNameMap() const;
		const FunctionSet&
		GetFunctionSet() const;
		static
		const CHAR*
		GetBasicTypeString(
			IN BasicType BaseType,
			IN DWORD Size,
			IN BOOL UseStdInt = FALSE
			);
		static
		const CHAR*
		GetBasicTypeString(
			IN const SYMBOL* Symbol,
			IN BOOL UseStdInt = FALSE
			);
		static
		const CHAR*
		GetUdtKindString(
			IN UdtKind Kind
			);
		static
		BOOL
		IsUnnamedSymbol(
			const SYMBOL* Symbol
			);
	private:
		SymbolModule* m_Impl;
};*/
return true
}



