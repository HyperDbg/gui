package Source
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\pdbex\Source\PDB.cpp.back

type (
Pdb interface{
		SymbolModuleBase()(ok bool)//col:19
SymbolModuleBase::SymbolModuleBase()(ok bool)//col:24
SymbolModuleBase::LoadDiaViaCoCreateInstance()(ok bool)//col:34
SymbolModuleBase::LoadDiaViaLoadLibrary()(ok bool)//col:58
SymbolModuleBase::Open()(ok bool)//col:112
SymbolModuleBase::Close()(ok bool)//col:119
SymbolModuleBase::IsOpen()(ok bool)//col:123
		SymbolModule()(ok bool)//col:230
SymbolModule::SymbolModule()(ok bool)//col:233
SymbolModule::~SymbolModule()(ok bool)//col:237
SymbolModule::Open()(ok bool)//col:254
SymbolModule::IsOpen()(ok bool)//col:258
SymbolModule::GetPath()(ok bool)//col:262
SymbolModule::Close()(ok bool)//col:275
SymbolModule::GetMachineType()(ok bool)//col:279
SymbolModule::GetLanguage()(ok bool)//col:283
SymbolModule::GetSymbolName()(ok bool)//col:300
SymbolModule::GetSymbolByName()(ok bool)//col:307
SymbolModule::GetSymbolByTypeId()(ok bool)//col:314
SymbolModule::GetSymbol()(ok bool)//col:336
SymbolModule::BuildSymbolMapFromEnumerator()(ok bool)//col:348
SymbolModule::BuildFunctionSetFromEnumerator()(ok bool)//col:369
SymbolModule::BuildSymbolMap()(ok bool)//col:387
SymbolModule::GetSymbolMap()(ok bool)//col:391
SymbolModule::GetSymbolNameMap()(ok bool)//col:395
SymbolModule::GetFunctionSet()(ok bool)//col:399
SymbolModule::InitSymbol()(ok bool)//col:435
SymbolModule::ProcessSymbolBase()(ok bool)//col:441
SymbolModule::ProcessSymbolEnum()(ok bool)//col:469
SymbolModule::ProcessSymbolTypedef()(ok bool)//col:478
SymbolModule::ProcessSymbolPointer()(ok bool)//col:497
SymbolModule::ProcessSymbolArray()(ok bool)//col:507
SymbolModule::ProcessSymbolFunction()(ok bool)//col:539
SymbolModule::ProcessSymbolFunctionArg()(ok bool)//col:548
SymbolModule::ProcessSymbolUdt()(ok bool)//col:624
void_SymbolModule::DestroySymbol()(ok bool)//col:650
	{_()(ok bool)//col:652
	{_()(ok bool)//col:686
PDB::PDB()(ok bool)//col:690
PDB::PDB()(ok bool)//col:697
PDB::~PDB()(ok bool)//col:701
PDB::Open()(ok bool)//col:707
PDB::IsOpened()(ok bool)//col:711
PDB::GetPath()(ok bool)//col:715
PDB::Close()(ok bool)//col:719
PDB::GetMachineType()(ok bool)//col:723
PDB::GetLanguage()(ok bool)//col:727
PDB::GetSymbolByName()(ok bool)//col:733
PDB::GetSymbolByTypeId()(ok bool)//col:739
PDB::GetSymbolMap()(ok bool)//col:743
PDB::GetSymbolNameMap()(ok bool)//col:747
PDB::GetFunctionSet()(ok bool)//col:751
PDB::GetBasicTypeString()(ok bool)//col:771
PDB::GetBasicTypeString()(ok bool)//col:778
PDB::GetUdtKindString()(ok bool)//col:793
PDB::IsUnnamedSymbol()(ok bool)//col:801
}
pdb struct{}
)

func NewPdb()Pdb{ return & pdb{} }

func (p *pdb)		SymbolModuleBase()(ok bool){//col:19
/*		SymbolModuleBase();
		BOOL
		Open(
			IN const CHAR* Path
			);
		VOID
		Close();
		BOOL
		IsOpen() const;
	private:
		HRESULT
		LoadDiaViaCoCreateInstance();
		HRESULT
		LoadDiaViaLoadLibrary();
	protected:
		CComPtr<IDiaDataSource> m_DataSource;
		CComPtr<IDiaSession>    m_Session;
		CComPtr<IDiaSymbol>     m_GlobalSymbol;
};*/
return true
}

func (p *pdb)SymbolModuleBase::SymbolModuleBase()(ok bool){//col:24
/*SymbolModuleBase::SymbolModuleBase()
{
	HRESULT hr = CoInitialize(nullptr);
	assert(hr == S_OK);
}*/
return true
}

func (p *pdb)SymbolModuleBase::LoadDiaViaCoCreateInstance()(ok bool){//col:34
/*SymbolModuleBase::LoadDiaViaCoCreateInstance()
{
	return CoCreateInstance(
		__uuidof(DiaSource),
		nullptr,
		CLSCTX_INPROC_SERVER,
		__uuidof(IDiaDataSource),
		(void**)& m_DataSource
		);
}*/
return true
}

func (p *pdb)SymbolModuleBase::LoadDiaViaLoadLibrary()(ok bool){//col:58
/*SymbolModuleBase::LoadDiaViaLoadLibrary()
{
	HRESULT Result;
	HMODULE Module = LoadLibrary(TEXT("msdia140.dll"));
	if (!Module)
	{
		Result = HRESULT_FROM_WIN32(GetLastError());
		return Result;
	}
	using PDLLGETCLASSOBJECT_ROUTINE = HRESULT(WINAPI*)(REFCLSID, REFIID, LPVOID);
	auto DllGetClassObject = reinterpret_cast<PDLLGETCLASSOBJECT_ROUTINE>(GetProcAddress(Module, "DllGetClassObject"));
	if (!DllGetClassObject)
	{
		Result = HRESULT_FROM_WIN32(GetLastError());
		return Result;
	}
	CComPtr<IClassFactory> ClassFactory;
	Result = DllGetClassObject(__uuidof(DiaSource), __uuidof(IClassFactory), &ClassFactory);
	if (FAILED(Result))
	{
		return Result;
	}
	return ClassFactory->CreateInstance(nullptr, __uuidof(IDiaDataSource), (void**)& m_DataSource);
}*/
return true
}

func (p *pdb)SymbolModuleBase::Open()(ok bool){//col:112
/*SymbolModuleBase::Open(
	IN const CHAR* Path
	)
{
	HRESULT   Result            = S_OK;
	LPCOLESTR PDBSearchPath     = L"srv*.\\Symbols*https:
	if (FAILED(Result = LoadDiaViaCoCreateInstance()) &&
	    FAILED(Result = LoadDiaViaLoadLibrary()))
	{
		return FALSE;
	}
	int PathUnicodeLength = MultiByteToWideChar(CP_UTF8, 0, Path, -1, NULL, 0);
	auto PathUnicode       = std::make_unique<WCHAR[]>(PathUnicodeLength);
	MultiByteToWideChar(CP_UTF8, 0, Path, -1, PathUnicode.get(), PathUnicodeLength);
	WCHAR FileExtension[8] = { 0 };
	_wsplitpath_s(
		PathUnicode.get(),
		nullptr,
		0,
		nullptr,
		0,
		nullptr,
		0,
		FileExtension,
		_countof(FileExtension));
	if (_wcsicmp(FileExtension, L".pdb") == 0)
	{
		Result = m_DataSource->loadDataFromPdb(PathUnicode.get());
	}
	else
	{
		PDBCallback Callback;
		Callback.AddRef();
		Result = m_DataSource->loadDataForExe(PathUnicode.get(), PDBSearchPath, &Callback);
	}
	if (FAILED(Result))
	{
		goto Error;
	}
	Result = m_DataSource->openSession(&m_Session);
	if (FAILED(Result))
	{
		goto Error;
	}
	Result = m_Session->get_globalScope(&m_GlobalSymbol);
	if (FAILED(Result))
	{
		goto Error;
	}
	return TRUE;
Error:
	Close();
	return FALSE;
}*/
return true
}

func (p *pdb)SymbolModuleBase::Close()(ok bool){//col:119
/*SymbolModuleBase::Close()
{
	m_GlobalSymbol.Release();
	m_Session.Release();
	m_DataSource.Release();
	CoUninitialize();
}*/
return true
}

func (p *pdb)SymbolModuleBase::IsOpen()(ok bool){//col:123
/*SymbolModuleBase::IsOpen() const
{
	return m_DataSource && m_Session && m_GlobalSymbol;
}*/
return true
}

func (p *pdb)		SymbolModule()(ok bool){//col:230
/*		SymbolModule();
		~SymbolModule();
		BOOL
		Open(
			IN const CHAR* Path
			);
		BOOL
		IsOpen() const;
		const CHAR*
		GetPath() const;
		VOID
		Close();
		DWORD
		GetMachineType() const;
		CV_CFL_LANG
		GetLanguage() const;
		SYMBOL*
		GetSymbolByName(
			IN const CHAR* SymbolName
			);
		SYMBOL*
		GetSymbolByTypeId(
			IN DWORD TypeId
			);
		SYMBOL*
		GetSymbol(
			IN IDiaSymbol* DiaSymbol
			);
		CHAR*
		GetSymbolName(
			IN IDiaSymbol* DiaSymbol
			);
		VOID
		BuildSymbolMapFromEnumerator(
			IN IDiaEnumSymbols* DiaSymbolEnumerator
			);
		VOID
		BuildFunctionSetFromEnumerator(
			IN IDiaEnumSymbols* DiaSymbolEnumerator
			);
		VOID
		BuildSymbolMap();
		const SymbolMap&
		GetSymbolMap() const;
		const SymbolNameMap&
		GetSymbolNameMap() const;
		const FunctionSet&
		GetFunctionSet() const;
	private:
		VOID
		InitSymbol(
			IN IDiaSymbol* DiaSymbol,
			IN SYMBOL* Symbol
			);
		VOID
		DestroySymbol(
			IN SYMBOL* Symbol
			);
		VOID
		ProcessSymbolBase(
			IN IDiaSymbol* DiaSymbol,
			IN SYMBOL* Symbol
			);
		VOID
		ProcessSymbolEnum(
			IN IDiaSymbol* DiaSymbol,
			IN SYMBOL* Symbol
			);
		VOID
		ProcessSymbolTypedef(
			IN IDiaSymbol* DiaSymbol,
			IN SYMBOL* Symbol
			);
		VOID
		ProcessSymbolPointer(
			IN IDiaSymbol* DiaSymbol,
			IN SYMBOL* Symbol
			);
		VOID
		ProcessSymbolArray(
			IN IDiaSymbol* DiaSymbol,
			IN SYMBOL* Symbol
			);
		VOID
		ProcessSymbolFunction(
			IN IDiaSymbol* DiaSymbol,
			IN SYMBOL* Symbol
			);
		VOID
		ProcessSymbolFunctionArg(
			IN IDiaSymbol* DiaSymbol,
			IN SYMBOL* Symbol
			);
		VOID
		ProcessSymbolUdt(
			IN IDiaSymbol* DiaSymbol,
			IN SYMBOL* Symbol
			);
	private:
		std::string   m_Path;
		SymbolMap     m_SymbolMap;
		SymbolNameMap m_SymbolNameMap;
		SymbolSet     m_SymbolSet;
		FunctionSet   m_FunctionSet;
		DWORD         m_MachineType = 0;
		CV_CFL_LANG   m_Language = CV_CFL_C;
};*/
return true
}

func (p *pdb)SymbolModule::SymbolModule()(ok bool){//col:233
/*SymbolModule::SymbolModule()
{
}*/
return true
}

func (p *pdb)SymbolModule::~SymbolModule()(ok bool){//col:237
/*SymbolModule::~SymbolModule()
{
	Close();
}*/
return true
}

func (p *pdb)SymbolModule::Open()(ok bool){//col:254
/*SymbolModule::Open(
	IN const CHAR* Path
	)
{
	BOOL Result;
	Result = SymbolModuleBase::Open(Path);
	if (Result == FALSE)
	{
		return FALSE;
	}
	m_GlobalSymbol->get_machineType(&m_MachineType);
	DWORD Language;
	m_GlobalSymbol->get_language(&Language);
	m_Language = static_cast<CV_CFL_LANG>(Language);
	BuildSymbolMap();
	return TRUE;
}*/
return true
}

func (p *pdb)SymbolModule::IsOpen()(ok bool){//col:258
/*SymbolModule::IsOpen() const
{
	return SymbolModuleBase::IsOpen();
}*/
return true
}

func (p *pdb)SymbolModule::GetPath()(ok bool){//col:262
/*SymbolModule::GetPath() const
{
	return m_Path.c_str();
}*/
return true
}

func (p *pdb)SymbolModule::Close()(ok bool){//col:275
/*SymbolModule::Close()
{
	SymbolModuleBase::Close();
	for (auto&& Symbol : m_SymbolSet)
	{
		DestroySymbol(Symbol);
		delete Symbol;
	}
	m_Path.clear();
	m_SymbolMap.clear();
	m_SymbolNameMap.clear();
	m_SymbolSet.clear();
}*/
return true
}

func (p *pdb)SymbolModule::GetMachineType()(ok bool){//col:279
/*SymbolModule::GetMachineType() const
{
	return m_MachineType;
}*/
return true
}

func (p *pdb)SymbolModule::GetLanguage()(ok bool){//col:283
/*SymbolModule::GetLanguage() const
{
	return m_Language;
}*/
return true
}

func (p *pdb)SymbolModule::GetSymbolName()(ok bool){//col:300
/*SymbolModule::GetSymbolName(
	IN IDiaSymbol* DiaSymbol
	)
{
	BSTR SymbolNameBstr;
	if (DiaSymbol->get_name(&SymbolNameBstr) != S_OK)
	{
		return nullptr;
	}
	CHAR*  SymbolNameMb;
	size_t SymbolNameLength;
	SymbolNameLength = (size_t)SysStringLen(SymbolNameBstr) + 1;
	SymbolNameMb = new CHAR[SymbolNameLength];
	wcstombs(SymbolNameMb, SymbolNameBstr, SymbolNameLength);
	SysFreeString(SymbolNameBstr);
	return SymbolNameMb;
}*/
return true
}

func (p *pdb)SymbolModule::GetSymbolByName()(ok bool){//col:307
/*SymbolModule::GetSymbolByName(
	IN const CHAR* SymbolName
	)
{
	auto it = m_SymbolNameMap.find(SymbolName);
	return it == m_SymbolNameMap.end() ? nullptr : it->second;
}*/
return true
}

func (p *pdb)SymbolModule::GetSymbolByTypeId()(ok bool){//col:314
/*SymbolModule::GetSymbolByTypeId(
	IN DWORD TypeId
	)
{
	auto it = m_SymbolMap.find(TypeId);
	return it == m_SymbolMap.end() ? nullptr : it->second;
}*/
return true
}

func (p *pdb)SymbolModule::GetSymbol()(ok bool){//col:336
/*SymbolModule::GetSymbol(
	IN IDiaSymbol* DiaSymbol
	)
{
	DWORD TypeId;
	DiaSymbol->get_symIndexId(&TypeId);
	auto it = m_SymbolMap.find(TypeId);
	if (it != m_SymbolMap.end())
	{
		return it->second;
	}
	SYMBOL* Symbol;
	Symbol = new SYMBOL;
	m_SymbolMap[TypeId] = Symbol;
	m_SymbolSet.insert(Symbol);
	InitSymbol(DiaSymbol, Symbol);
	if (Symbol->Name)
	{
		m_SymbolNameMap[Symbol->Name] = Symbol;
	}
	return Symbol;
}*/
return true
}

func (p *pdb)SymbolModule::BuildSymbolMapFromEnumerator()(ok bool){//col:348
/*SymbolModule::BuildSymbolMapFromEnumerator(
	IN IDiaEnumSymbols* DiaSymbolEnumerator
	)
{
	IDiaSymbol* Result;
	ULONG FetchedSymbolCount = 0;
	while (SUCCEEDED(DiaSymbolEnumerator->Next(1, &Result, &FetchedSymbolCount)) && (FetchedSymbolCount == 1))
	{
		CComPtr<IDiaSymbol> DiaChildSymbol(Result);
		GetSymbol(DiaChildSymbol);
	}
}*/
return true
}

func (p *pdb)SymbolModule::BuildFunctionSetFromEnumerator()(ok bool){//col:369
/*SymbolModule::BuildFunctionSetFromEnumerator(
	IN IDiaEnumSymbols* DiaSymbolEnumerator
	)
{
	IDiaSymbol* Result;
	ULONG FetchedSymbolCount = 0;
	while (SUCCEEDED(DiaSymbolEnumerator->Next(1, &Result, &FetchedSymbolCount)) && (FetchedSymbolCount == 1))
	{
		CComPtr<IDiaSymbol> DiaChildSymbol(Result);
		BOOL IsFunction;
		DiaChildSymbol->get_function(&IsFunction);
		if (IsFunction)
		{
			CHAR* FunctionName = GetSymbolName(DiaChildSymbol);
			DWORD DwordResult;
			DiaChildSymbol->get_symTag(&DwordResult);
			m_FunctionSet.insert(FunctionName);
			delete[] FunctionName;
		}
	}
}*/
return true
}

func (p *pdb)SymbolModule::BuildSymbolMap()(ok bool){//col:387
/*SymbolModule::BuildSymbolMap()
{
	if (CComPtr<IDiaEnumSymbols> DiaSymbolEnumerator;
	    SUCCEEDED(m_GlobalSymbol->findChildren(SymTagPublicSymbol, nullptr, nsNone, &DiaSymbolEnumerator)))
	{
		BuildFunctionSetFromEnumerator(DiaSymbolEnumerator);
	}
	if (CComPtr<IDiaEnumSymbols> DiaSymbolEnumerator;
	    SUCCEEDED(m_GlobalSymbol->findChildren(SymTagEnum, nullptr, nsNone, &DiaSymbolEnumerator)))
	{
		BuildSymbolMapFromEnumerator(DiaSymbolEnumerator);
	}
	if (CComPtr<IDiaEnumSymbols> DiaSymbolEnumerator;
	    SUCCEEDED(m_GlobalSymbol->findChildren(SymTagUDT, nullptr, nsNone, &DiaSymbolEnumerator)))
	{
		BuildSymbolMapFromEnumerator(DiaSymbolEnumerator);
	}
}*/
return true
}

func (p *pdb)SymbolModule::GetSymbolMap()(ok bool){//col:391
/*SymbolModule::GetSymbolMap() const
{
	return m_SymbolMap;
}*/
return true
}

func (p *pdb)SymbolModule::GetSymbolNameMap()(ok bool){//col:395
/*SymbolModule::GetSymbolNameMap() const
{
	return m_SymbolNameMap;
}*/
return true
}

func (p *pdb)SymbolModule::GetFunctionSet()(ok bool){//col:399
/*SymbolModule::GetFunctionSet() const
{
	return m_FunctionSet;
}*/
return true
}

func (p *pdb)SymbolModule::InitSymbol()(ok bool){//col:435
/*SymbolModule::InitSymbol(
	IN IDiaSymbol* DiaSymbol,
	IN SYMBOL* Symbol
	)
{
	DWORD DwordResult;
	ULONGLONG UlonglongResult;
	BOOL BoolResult;
	DiaSymbol->get_symTag(&DwordResult);
	Symbol->Tag = static_cast<enum SymTagEnum>(DwordResult);
	DiaSymbol->get_dataKind(&DwordResult);
	Symbol->DataKind = static_cast<enum DataKind>(DwordResult);
	DiaSymbol->get_baseType(&DwordResult);
	Symbol->BaseType = static_cast<BasicType>(DwordResult);
	DiaSymbol->get_typeId(&DwordResult);
	Symbol->TypeId = DwordResult;
	DiaSymbol->get_length(&UlonglongResult);
	Symbol->Size = static_cast<DWORD>(UlonglongResult);
	DiaSymbol->get_constType(&BoolResult);
	Symbol->IsConst = static_cast<BOOL>(BoolResult);
	DiaSymbol->get_volatileType(&BoolResult);
	Symbol->IsVolatile = static_cast<BOOL>(BoolResult);
	Symbol->Name = GetSymbolName(DiaSymbol);
	switch (Symbol->Tag)
	{
		case SymTagUDT:             ProcessSymbolUdt        (DiaSymbol, Symbol); break;
		case SymTagEnum:            ProcessSymbolEnum       (DiaSymbol, Symbol); break;
		case SymTagFunctionType:    ProcessSymbolFunction   (DiaSymbol, Symbol); break;
		case SymTagPointerType:     ProcessSymbolPointer    (DiaSymbol, Symbol); break;
		case SymTagArrayType:       ProcessSymbolArray      (DiaSymbol, Symbol); break;
		case SymTagBaseType:        ProcessSymbolBase       (DiaSymbol, Symbol); break;
		case SymTagTypedef:         ProcessSymbolTypedef    (DiaSymbol, Symbol); break;
		case SymTagFunctionArgType: ProcessSymbolFunctionArg(DiaSymbol, Symbol); break;
		default:                                                                 break;
	}
}*/
return true
}

func (p *pdb)SymbolModule::ProcessSymbolBase()(ok bool){//col:441
/*SymbolModule::ProcessSymbolBase(
	IN IDiaSymbol* DiaSymbol,
	IN SYMBOL* Symbol
	)
{
}*/
return true
}

func (p *pdb)SymbolModule::ProcessSymbolEnum()(ok bool){//col:469
/*SymbolModule::ProcessSymbolEnum(
	IN IDiaSymbol* DiaSymbol,
	IN SYMBOL* Symbol
	)
{
	CComPtr<IDiaEnumSymbols> DiaSymbolEnumerator;
	if (FAILED(DiaSymbol->findChildren(SymTagNull, nullptr, nsNone, &DiaSymbolEnumerator)))
	{
		return;
	}
	LONG ChildCount;
	DiaSymbolEnumerator->get_Count(&ChildCount);
	Symbol->u.Enum.FieldCount = static_cast<DWORD>(ChildCount);
	Symbol->u.Enum.Fields = new SYMBOL_ENUM_FIELD[ChildCount];
	IDiaSymbol* Result;
	ULONG FetchedSymbolCount = 0;
	DWORD Index = 0;
	while (SUCCEEDED(DiaSymbolEnumerator->Next(1, &Result, &FetchedSymbolCount)) && (FetchedSymbolCount == 1))
	{
		CComPtr<IDiaSymbol> DiaChildSymbol(Result);
		SYMBOL_ENUM_FIELD* EnumValue = &Symbol->u.Enum.Fields[Index];
		EnumValue->Parent = Symbol;
		EnumValue->Name = GetSymbolName(DiaChildSymbol);
		VariantInit(&EnumValue->Value);
		DiaChildSymbol->get_value(&EnumValue->Value);
		Index += 1;
	}
}*/
return true
}

func (p *pdb)SymbolModule::ProcessSymbolTypedef()(ok bool){//col:478
/*SymbolModule::ProcessSymbolTypedef(
	IN IDiaSymbol* DiaSymbol,
	IN SYMBOL* Symbol
	)
{
	CComPtr<IDiaSymbol> DiaTypedefSymbol;
	DiaSymbol->get_type(&DiaTypedefSymbol);
	Symbol->u.Typedef.Type = GetSymbol(DiaTypedefSymbol);
}*/
return true
}

func (p *pdb)SymbolModule::ProcessSymbolPointer()(ok bool){//col:497
/*SymbolModule::ProcessSymbolPointer(
	IN IDiaSymbol* DiaSymbol,
	IN SYMBOL* Symbol
	)
{
	CComPtr<IDiaSymbol> DiaPointerSymbol;
	DiaSymbol->get_type(&DiaPointerSymbol);
	DiaSymbol->get_reference(&Symbol->u.Pointer.IsReference);
	Symbol->u.Pointer.Type = GetSymbol(DiaPointerSymbol);
	if (m_MachineType == 0)
	{
		switch (Symbol->Size)
		{
			case 4:  m_MachineType = IMAGE_FILE_MACHINE_I386;  break;
			case 8:  m_MachineType = IMAGE_FILE_MACHINE_AMD64; break;
			default: m_MachineType = 0; break;
		}
	}
}*/
return true
}

func (p *pdb)SymbolModule::ProcessSymbolArray()(ok bool){//col:507
/*SymbolModule::ProcessSymbolArray(
	IN IDiaSymbol* DiaSymbol,
	IN SYMBOL* Symbol
	)
{
	CComPtr<IDiaSymbol> DiaDataTypeSymbol;
	DiaSymbol->get_type(&DiaDataTypeSymbol);
	Symbol->u.Array.ElementType = GetSymbol(DiaDataTypeSymbol);
	DiaSymbol->get_count(&Symbol->u.Array.ElementCount);
}*/
return true
}

func (p *pdb)SymbolModule::ProcessSymbolFunction()(ok bool){//col:539
/*SymbolModule::ProcessSymbolFunction(
	IN IDiaSymbol* DiaSymbol,
	IN SYMBOL* Symbol
	)
{
	DWORD CallingConvention;
	DiaSymbol->get_callingConvention(&CallingConvention);
	Symbol->u.Function.CallingConvention = static_cast<CV_call_e>(CallingConvention);
	CComPtr<IDiaSymbol> DiaReturnTypeSymbol;
	DiaSymbol->get_type(&DiaReturnTypeSymbol);
	Symbol->u.Function.ReturnType = GetSymbol(DiaReturnTypeSymbol);
	CComPtr<IDiaEnumSymbols> DiaSymbolEnumerator;
	if (FAILED(DiaSymbol->findChildren(SymTagNull, nullptr, nsNone, &DiaSymbolEnumerator)))
	{
		return;
	}
	LONG ChildCount;
	DiaSymbolEnumerator->get_Count(&ChildCount);
	Symbol->u.Function.ArgumentCount = static_cast<DWORD>(ChildCount);
	Symbol->u.Function.Arguments = new SYMBOL*[ChildCount];
	IDiaSymbol* Result;
	ULONG FetchedSymbolCount = 0;
	DWORD Index = 0;
	while (SUCCEEDED(DiaSymbolEnumerator->Next(1, &Result, &FetchedSymbolCount)) && (FetchedSymbolCount == 1))
	{
		CComPtr<IDiaSymbol> DiaChildSymbol(Result);
		SYMBOL* Argument;
		Argument = GetSymbol(DiaChildSymbol);
		Symbol->u.Function.Arguments[Index] = Argument;
		Index += 1;
	}
}*/
return true
}

func (p *pdb)SymbolModule::ProcessSymbolFunctionArg()(ok bool){//col:548
/*SymbolModule::ProcessSymbolFunctionArg(
	IN IDiaSymbol* DiaSymbol,
	IN SYMBOL* Symbol
	)
{
	CComPtr<IDiaSymbol> DiaArgumentTypeSymbol;
	DiaSymbol->get_type(&DiaArgumentTypeSymbol);
	Symbol->u.FunctionArg.Type = GetSymbol(DiaArgumentTypeSymbol);
}*/
return true
}

func (p *pdb)SymbolModule::ProcessSymbolUdt()(ok bool){//col:624
/*SymbolModule::ProcessSymbolUdt(
	IN IDiaSymbol* DiaSymbol,
	IN SYMBOL* Symbol
	)
{
	DWORD Kind;
	DiaSymbol->get_udtKind(&Kind);
	Symbol->u.Udt.Kind = static_cast<UdtKind>(Kind);
	CComPtr<IDiaEnumSymbols> DiaSymbolEnumerator;
	if (FAILED(DiaSymbol->findChildren(SymTagData, nullptr, nsNone, &DiaSymbolEnumerator)))
	{
		return;
	}
	LONG ChildCount;
	DiaSymbolEnumerator->get_Count(&ChildCount);
	Symbol->u.Udt.FieldCount = static_cast<DWORD>(ChildCount);
	Symbol->u.Udt.Fields = new SYMBOL_UDT_FIELD[ChildCount + 1];
	IDiaSymbol* Result;
	ULONG FetchedSymbolCount = 0;
	DWORD Index = 0;
	while (SUCCEEDED(DiaSymbolEnumerator->Next(1, &Result, &FetchedSymbolCount)) && (FetchedSymbolCount == 1))
	{
		CComPtr<IDiaSymbol> DiaChildSymbol(Result);
		SYMBOL_UDT_FIELD* Member = &Symbol->u.Udt.Fields[Index];
		Member->Name = GetSymbolName(DiaChildSymbol);
		Member->Parent = Symbol;
		LONG Offset = 0;
		DiaChildSymbol->get_offset(&Offset);
		Member->Offset = static_cast<DWORD>(Offset);
		ULONGLONG Bits = 0;
		DiaChildSymbol->get_length(&Bits);
		Member->Bits = static_cast<DWORD>(Bits);
		DiaChildSymbol->get_bitPosition(&Member->BitPosition);
		CComPtr<IDiaSymbol> MemberTypeDiaSymbol;
		DiaChildSymbol->get_type(&MemberTypeDiaSymbol);
		Member->Type = GetSymbol(MemberTypeDiaSymbol);
		Index += 1;
	}
	if (Symbol->u.Udt.Kind == UdtStruct && Symbol->u.Udt.FieldCount > 0 && Symbol->u.Udt.Fields[Symbol->u.Udt.FieldCount - 1].Type != nullptr)
	{
		SYMBOL_UDT_FIELD* LastUdtField = &Symbol->u.Udt.Fields[Symbol->u.Udt.FieldCount - 1];
		SYMBOL_UDT_FIELD* PaddingUdtField = &Symbol->u.Udt.Fields[Symbol->u.Udt.FieldCount];
		DWORD PaddingSize = Symbol->Size - (LastUdtField->Offset + LastUdtField->Type->Size);
		if (PaddingSize > 0)
		{
			SYMBOL* PaddingSymbolArrayElement = new SYMBOL;
			PaddingSymbolArrayElement->Tag = SymTagBaseType;
			PaddingSymbolArrayElement->BaseType = !(PaddingSize % 4) ? btLong : btChar;
			PaddingSymbolArrayElement->TypeId = 0;
			PaddingSymbolArrayElement->Size = PaddingSymbolArrayElement->BaseType == btLong ? 4 : 1;
			PaddingSymbolArrayElement->IsConst = FALSE;
			PaddingSymbolArrayElement->IsVolatile = FALSE;
			PaddingSymbolArrayElement->Name = nullptr;
			SYMBOL* PaddingSymbolArray = new SYMBOL;
			PaddingSymbolArray->Tag = SymTagArrayType;
			PaddingSymbolArray->BaseType = btNoType;
			PaddingSymbolArray->TypeId = 0;
			PaddingSymbolArray->Size = PaddingSize;
			PaddingSymbolArray->IsConst = FALSE;
			PaddingSymbolArray->IsVolatile = FALSE;
			PaddingSymbolArray->Name = nullptr;
			PaddingSymbolArray->u.Array.ElementType = PaddingSymbolArrayElement;
			PaddingSymbolArray->u.Array.ElementCount = PaddingSymbolArrayElement->BaseType == btLong ? PaddingSize / 4 : PaddingSize;
			PaddingUdtField->Name = new CHAR[64];
			PaddingUdtField->Type = PaddingSymbolArray;
			PaddingUdtField->Offset = LastUdtField->Offset + LastUdtField->Type->Size;
			PaddingUdtField->Bits = 0;
			PaddingUdtField->BitPosition = 0;
			PaddingUdtField->Parent = Symbol;
			strcpy(PaddingUdtField->Name, "__PADDING__");
			Symbol->u.Udt.FieldCount++;
			m_SymbolSet.insert(PaddingSymbolArray);
			m_SymbolSet.insert(PaddingSymbolArrayElement);
		}
	}
}*/
return true
}

func (p *pdb)void_SymbolModule::DestroySymbol()(ok bool){//col:650
/*void SymbolModule::DestroySymbol(
	IN SYMBOL* Symbol
	)
{
	delete[] Symbol->Name;
	switch (Symbol->Tag)
	{
		case SymTagUDT:
			for (DWORD i = 0; i < Symbol->u.Udt.FieldCount; i++)
			{
				delete[] Symbol->u.Udt.Fields[i].Name;
			}
			delete[] Symbol->u.Udt.Fields;
			break;
		case SymTagEnum:
			for (DWORD i = 0; i < Symbol->u.Enum.FieldCount; i++)
			{
				delete[] Symbol->u.Enum.Fields[i].Name;
			}
			delete[] Symbol->u.Enum.Fields;
			break;
		case SymTagFunctionType:
			delete[] Symbol->u.Function.Arguments;
			break;
	}
}*/
return true
}

func (p *pdb)	{_()(ok bool){//col:652
/*	{ (BasicType)0,   0,  nullptr,            nullptr            },
};*/
return true
}

func (p *pdb)	{_()(ok bool){//col:686
/*	{ (BasicType)0,   0,  nullptr,            nullptr            },
};*/
return true
}

func (p *pdb)PDB::PDB()(ok bool){//col:690
/*PDB::PDB()
{
	m_Impl = new SymbolModule();
}*/
return true
}

func (p *pdb)PDB::PDB()(ok bool){//col:697
/*PDB::PDB(
	IN const CHAR* Path
	)
{
	m_Impl = new SymbolModule();
	m_Impl->Open(Path);
}*/
return true
}

func (p *pdb)PDB::~PDB()(ok bool){//col:701
/*PDB::~PDB()
{
	delete m_Impl;
}*/
return true
}

func (p *pdb)PDB::Open()(ok bool){//col:707
/*PDB::Open(
	IN const CHAR* Path
	)
{
	return m_Impl->Open(Path);
}*/
return true
}

func (p *pdb)PDB::IsOpened()(ok bool){//col:711
/*PDB::IsOpened() const
{
	return m_Impl->IsOpen();
}*/
return true
}

func (p *pdb)PDB::GetPath()(ok bool){//col:715
/*PDB::GetPath() const
{
	return m_Impl->GetPath();
}*/
return true
}

func (p *pdb)PDB::Close()(ok bool){//col:719
/*PDB::Close()
{
	m_Impl->Close();
}*/
return true
}

func (p *pdb)PDB::GetMachineType()(ok bool){//col:723
/*PDB::GetMachineType() const
{
	return m_Impl->GetMachineType();
}*/
return true
}

func (p *pdb)PDB::GetLanguage()(ok bool){//col:727
/*PDB::GetLanguage() const
{
	return m_Impl->GetLanguage();
}*/
return true
}

func (p *pdb)PDB::GetSymbolByName()(ok bool){//col:733
/*PDB::GetSymbolByName(
	IN const CHAR* SymbolName
	)
{
	return m_Impl->GetSymbolByName(SymbolName);
}*/
return true
}

func (p *pdb)PDB::GetSymbolByTypeId()(ok bool){//col:739
/*PDB::GetSymbolByTypeId(
	IN DWORD TypeId
	)
{
	return m_Impl->GetSymbolByTypeId(TypeId);
}*/
return true
}

func (p *pdb)PDB::GetSymbolMap()(ok bool){//col:743
/*PDB::GetSymbolMap() const
{
	return m_Impl->GetSymbolMap();
}*/
return true
}

func (p *pdb)PDB::GetSymbolNameMap()(ok bool){//col:747
/*PDB::GetSymbolNameMap() const
{
	return m_Impl->GetSymbolNameMap();
}*/
return true
}

func (p *pdb)PDB::GetFunctionSet()(ok bool){//col:751
/*PDB::GetFunctionSet() const
{
	return m_Impl->GetFunctionSet();
}*/
return true
}

func (p *pdb)PDB::GetBasicTypeString()(ok bool){//col:771
/*PDB::GetBasicTypeString(
	IN BasicType BaseType,
	IN DWORD Size,
	IN BOOL UseStdInt
	)
{
	BasicTypeMapElement* TypeMap = UseStdInt ? BasicTypeMapStdInt : BasicTypeMapMSVC;
	for (int n = 0; TypeMap[n].BasicTypeString != nullptr; n++)
	{
		if (TypeMap[n].BaseType == BaseType)
		{
			if (TypeMap[n].Length == Size ||
			    TypeMap[n].Length == 0)
			{
				return TypeMap[n].TypeString;
			}
		}
	}
	return nullptr;
}*/
return true
}

func (p *pdb)PDB::GetBasicTypeString()(ok bool){//col:778
/*PDB::GetBasicTypeString(
	IN const SYMBOL* Symbol,
	IN BOOL UseStdInt
	)
{
	return GetBasicTypeString(Symbol->BaseType, Symbol->Size, UseStdInt);
}*/
return true
}

func (p *pdb)PDB::GetUdtKindString()(ok bool){//col:793
/*PDB::GetUdtKindString(
	IN UdtKind Kind
	)
{
	static const CHAR* UdtKindStrings[] = {
		"struct",
		"class",
		"union",
	};
	if (Kind >= UdtStruct && Kind <= UdtUnion)
	{
		return UdtKindStrings[Kind];
	}
	return nullptr;
}*/
return true
}

func (p *pdb)PDB::IsUnnamedSymbol()(ok bool){//col:801
/*PDB::IsUnnamedSymbol(
	const SYMBOL* Symbol
	)
{
	return strstr(Symbol->Name, "<anonymous-") != nullptr ||
	       strstr(Symbol->Name, "<unnamed-") != nullptr ||
	       strstr(Symbol->Name, "__unnamed") != nullptr;
}*/
return true
}



