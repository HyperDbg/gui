package Source
//back\HyperDbgDev\hyperdbg\dependencies\pdbex\Source\PDB.h.back

const(
_CRT_SECURE_NO_WARNINGS =  //col:2
)


type SYMBOL_ENUM_FIELD struct{
Name CHAR*
Value VARIANT
Parent SYMBOL*
}


type SYMBOL_UDT_FIELD struct{
Name CHAR*
Type SYMBOL*
Offset DWORD
Bits DWORD
BitPosition DWORD
Parent SYMBOL*
}


type SYMBOL_ENUM struct{
FieldCount DWORD
Fields SYMBOL_ENUM_FIELD*
}


type SYMBOL_TYPEDEF struct{
Type SYMBOL*
}


type SYMBOL_POINTER struct{
Type SYMBOL*
IsReference BOOL
}


type SYMBOL_ARRAY struct{
ElementType SYMBOL*
ElementCount DWORD
}


type SYMBOL_FUNCTION struct{
ReturnType SYMBOL*
CallingConvention CV_call_e
ArgumentCount DWORD
Arguments SYMBOL**
}


type SYMBOL_FUNCTIONARG struct{
Type SYMBOL*
}


type SYMBOL_UDT struct{
Kind UdtKind
FieldCount DWORD
Fields SYMBOL_UDT_FIELD*
}



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



