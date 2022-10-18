package Source
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\pdbex\Source\PDB.cpp.back

type (
Pdb interface{
		SymbolModuleBase()(ok bool)//col:97
		SymbolModule()(ok bool)//col:179
SymbolModule::~SymbolModule()(ok bool)//col:216
SymbolModule::GetLanguage()(ok bool)//col:220
SymbolModule::GetSymbolName()(ok bool)//col:242
SymbolModule::GetSymbolByTypeId()(ok bool)//col:249
SymbolModule::GetSymbol()(ok bool)//col:271
SymbolModule::BuildSymbolMapFromEnumerator()(ok bool)//col:314
SymbolModule::GetSymbolNameMap()(ok bool)//col:318
SymbolModule::GetFunctionSet()(ok bool)//col:322
SymbolModule::InitSymbol()(ok bool)//col:360
SymbolModule::ProcessSymbolEnum()(ok bool)//col:406
SymbolModule::ProcessSymbolArray()(ok bool)//col:530
	{_()(ok bool)//col:532
	{_()(ok bool)//col:566
PDB::PDB()(ok bool)//col:579
PDB::Open()(ok bool)//col:638
PDB::GetBasicTypeString()(ok bool)//col:659
PDB::IsUnnamedSymbol()(ok bool)//col:667
}
pdb struct{}
)

func NewPdb()Pdb{ return & pdb{} }

func (p *pdb)		SymbolModuleBase()(ok bool){//col:97
/*		SymbolModuleBase();
		Open(
			IN const CHAR* Path
			);
		Close();
		IsOpen() const;
	private:
		HRESULT
		LoadDiaViaCoCreateInstance();
		LoadDiaViaLoadLibrary();
SymbolModuleBase::SymbolModuleBase()
{
	HRESULT hr = CoInitialize(nullptr);
	assert(hr == S_OK);
SymbolModuleBase::LoadDiaViaCoCreateInstance()
{
	return CoCreateInstance(
		__uuidof(DiaSource),
		nullptr,
		CLSCTX_INPROC_SERVER,
		__uuidof(IDiaDataSource),
		(void**)& m_DataSource
		);
SymbolModuleBase::LoadDiaViaLoadLibrary()
{
	HRESULT Result;
	HMODULE Module = LoadLibrary(TEXT("msdia140.dll"));
	if (!Module)
	{
		Result = HRESULT_FROM_WIN32(GetLastError());
	using PDLLGETCLASSOBJECT_ROUTINE = HRESULT(WINAPI*)(REFCLSID, REFIID, LPVOID);
	auto DllGetClassObject = reinterpret_cast<PDLLGETCLASSOBJECT_ROUTINE>(GetProcAddress(Module, "DllGetClassObject"));
	if (!DllGetClassObject)
	{
		Result = HRESULT_FROM_WIN32(GetLastError());
	Result = DllGetClassObject(__uuidof(DiaSource), __uuidof(IClassFactory), &ClassFactory);
	if (FAILED(Result))
	{
		return Result;
	}
	return ClassFactory->CreateInstance(nullptr, __uuidof(IDiaDataSource), (void**)& m_DataSource);
SymbolModuleBase::Open(
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
		Callback.AddRef();
		Result = m_DataSource->loadDataForExe(PathUnicode.get(), PDBSearchPath, &Callback);
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
SymbolModuleBase::Close()
{
	m_GlobalSymbol.Release();
	m_Session.Release();
	m_DataSource.Release();
	CoUninitialize();
SymbolModuleBase::IsOpen() const
{
	return m_DataSource && m_Session && m_GlobalSymbol;
}*/
return true
}

func (p *pdb)		SymbolModule()(ok bool){//col:179
/*		SymbolModule();
		~SymbolModule();
		Open(
			IN const CHAR* Path
			);
		IsOpen() const;
		const CHAR*
		GetPath() const;
		VOID
		Close();
		GetMachineType() const;
		CV_CFL_LANG
		GetLanguage() const;
		SYMBOL*
		GetSymbolByName(
			IN const CHAR* SymbolName
			);
		GetSymbolByTypeId(
			IN DWORD TypeId
			);
		GetSymbol(
			IN IDiaSymbol* DiaSymbol
			);
		GetSymbolName(
			IN IDiaSymbol* DiaSymbol
			);
		BuildSymbolMapFromEnumerator(
			IN IDiaEnumSymbols* DiaSymbolEnumerator
			);
		BuildFunctionSetFromEnumerator(
			IN IDiaEnumSymbols* DiaSymbolEnumerator
			);
		BuildSymbolMap();
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
		DestroySymbol(
			IN SYMBOL* Symbol
			);
		ProcessSymbolBase(
			IN IDiaSymbol* DiaSymbol,
			IN SYMBOL* Symbol
			);
		ProcessSymbolEnum(
			IN IDiaSymbol* DiaSymbol,
			IN SYMBOL* Symbol
			);
		ProcessSymbolTypedef(
			IN IDiaSymbol* DiaSymbol,
			IN SYMBOL* Symbol
			);
		ProcessSymbolPointer(
			IN IDiaSymbol* DiaSymbol,
			IN SYMBOL* Symbol
			);
		ProcessSymbolArray(
			IN IDiaSymbol* DiaSymbol,
			IN SYMBOL* Symbol
			);
		ProcessSymbolFunction(
			IN IDiaSymbol* DiaSymbol,
			IN SYMBOL* Symbol
			);
		ProcessSymbolFunctionArg(
			IN IDiaSymbol* DiaSymbol,
			IN SYMBOL* Symbol
			);
		ProcessSymbolUdt(
			IN IDiaSymbol* DiaSymbol,
			IN SYMBOL* Symbol
			);
SymbolModule::SymbolModule()
{
}*/
return true
}

func (p *pdb)SymbolModule::~SymbolModule()(ok bool){//col:216
/*SymbolModule::~SymbolModule()
{
	Close();
SymbolModule::Open(
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
	m_GlobalSymbol->get_language(&Language);
	m_Language = static_cast<CV_CFL_LANG>(Language);
	BuildSymbolMap();
SymbolModule::IsOpen() const
{
	return SymbolModuleBase::IsOpen();
SymbolModule::GetPath() const
{
	return m_Path.c_str();
SymbolModule::Close()
{
	SymbolModuleBase::Close();
	for (auto&& Symbol : m_SymbolSet)
	{
		DestroySymbol(Symbol);
	m_Path.clear();
	m_SymbolMap.clear();
	m_SymbolNameMap.clear();
	m_SymbolSet.clear();
SymbolModule::GetMachineType() const
{
	return m_MachineType;
}*/
return true
}

func (p *pdb)SymbolModule::GetLanguage()(ok bool){//col:220
/*SymbolModule::GetLanguage() const
{
	return m_Language;
}*/
return true
}

func (p *pdb)SymbolModule::GetSymbolName()(ok bool){//col:242
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
SymbolModule::GetSymbolByName(
	IN const CHAR* SymbolName
	)
{
	auto it = m_SymbolNameMap.find(SymbolName);
	return it == m_SymbolNameMap.end() ? nullptr : it->second;
}*/
return true
}

func (p *pdb)SymbolModule::GetSymbolByTypeId()(ok bool){//col:249
/*SymbolModule::GetSymbolByTypeId(
	IN DWORD TypeId
	)
{
	auto it = m_SymbolMap.find(TypeId);
	return it == m_SymbolMap.end() ? nullptr : it->second;
}*/
return true
}

func (p *pdb)SymbolModule::GetSymbol()(ok bool){//col:271
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

func (p *pdb)SymbolModule::BuildSymbolMapFromEnumerator()(ok bool){//col:314
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
SymbolModule::BuildFunctionSetFromEnumerator(
	IN IDiaEnumSymbols* DiaSymbolEnumerator
	)
{
	IDiaSymbol* Result;
	ULONG FetchedSymbolCount = 0;
	while (SUCCEEDED(DiaSymbolEnumerator->Next(1, &Result, &FetchedSymbolCount)) && (FetchedSymbolCount == 1))
	{
		CComPtr<IDiaSymbol> DiaChildSymbol(Result);
		DiaChildSymbol->get_function(&IsFunction);
		if (IsFunction)
		{
			CHAR* FunctionName = GetSymbolName(DiaChildSymbol);
			DiaChildSymbol->get_symTag(&DwordResult);
			m_FunctionSet.insert(FunctionName);
SymbolModule::BuildSymbolMap()
{
	if (CComPtr<IDiaEnumSymbols> DiaSymbolEnumerator;
	    SUCCEEDED(m_GlobalSymbol->findChildren(SymTagPublicSymbol, nullptr, nsNone, &DiaSymbolEnumerator)))
	{
		BuildFunctionSetFromEnumerator(DiaSymbolEnumerator);
	if (CComPtr<IDiaEnumSymbols> DiaSymbolEnumerator;
	    SUCCEEDED(m_GlobalSymbol->findChildren(SymTagEnum, nullptr, nsNone, &DiaSymbolEnumerator)))
	{
		BuildSymbolMapFromEnumerator(DiaSymbolEnumerator);
	if (CComPtr<IDiaEnumSymbols> DiaSymbolEnumerator;
	    SUCCEEDED(m_GlobalSymbol->findChildren(SymTagUDT, nullptr, nsNone, &DiaSymbolEnumerator)))
	{
		BuildSymbolMapFromEnumerator(DiaSymbolEnumerator);
SymbolModule::GetSymbolMap() const
{
	return m_SymbolMap;
}*/
return true
}

func (p *pdb)SymbolModule::GetSymbolNameMap()(ok bool){//col:318
/*SymbolModule::GetSymbolNameMap() const
{
	return m_SymbolNameMap;
}*/
return true
}

func (p *pdb)SymbolModule::GetFunctionSet()(ok bool){//col:322
/*SymbolModule::GetFunctionSet() const
{
	return m_FunctionSet;
}*/
return true
}

func (p *pdb)SymbolModule::InitSymbol()(ok bool){//col:360
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
SymbolModule::ProcessSymbolBase(
	IN IDiaSymbol* DiaSymbol,
	IN SYMBOL* Symbol
	)
{
}*/
return true
}

func (p *pdb)SymbolModule::ProcessSymbolEnum()(ok bool){//col:406
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
	while (SUCCEEDED(DiaSymbolEnumerator->Next(1, &Result, &FetchedSymbolCount)) && (FetchedSymbolCount == 1))
	{
		CComPtr<IDiaSymbol> DiaChildSymbol(Result);
		EnumValue->Name = GetSymbolName(DiaChildSymbol);
		VariantInit(&EnumValue->Value);
		DiaChildSymbol->get_value(&EnumValue->Value);
SymbolModule::ProcessSymbolTypedef(
	IN IDiaSymbol* DiaSymbol,
	IN SYMBOL* Symbol
	)
{
	CComPtr<IDiaSymbol> DiaTypedefSymbol;
	DiaSymbol->get_type(&DiaTypedefSymbol);
	Symbol->u.Typedef.Type = GetSymbol(DiaTypedefSymbol);
SymbolModule::ProcessSymbolPointer(
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

func (p *pdb)SymbolModule::ProcessSymbolArray()(ok bool){//col:530
/*SymbolModule::ProcessSymbolArray(
	IN IDiaSymbol* DiaSymbol,
	IN SYMBOL* Symbol
	)
{
	CComPtr<IDiaSymbol> DiaDataTypeSymbol;
	DiaSymbol->get_type(&DiaDataTypeSymbol);
	Symbol->u.Array.ElementType = GetSymbol(DiaDataTypeSymbol);
	DiaSymbol->get_count(&Symbol->u.Array.ElementCount);
SymbolModule::ProcessSymbolFunction(
	IN IDiaSymbol* DiaSymbol,
	IN SYMBOL* Symbol
	)
{
	DWORD CallingConvention;
	DiaSymbol->get_callingConvention(&CallingConvention);
	Symbol->u.Function.CallingConvention = static_cast<CV_call_e>(CallingConvention);
	DiaSymbol->get_type(&DiaReturnTypeSymbol);
	Symbol->u.Function.ReturnType = GetSymbol(DiaReturnTypeSymbol);
	if (FAILED(DiaSymbol->findChildren(SymTagNull, nullptr, nsNone, &DiaSymbolEnumerator)))
	{
		return;
	}
	LONG ChildCount;
	DiaSymbolEnumerator->get_Count(&ChildCount);
	Symbol->u.Function.ArgumentCount = static_cast<DWORD>(ChildCount);
	while (SUCCEEDED(DiaSymbolEnumerator->Next(1, &Result, &FetchedSymbolCount)) && (FetchedSymbolCount == 1))
	{
		CComPtr<IDiaSymbol> DiaChildSymbol(Result);
		Argument = GetSymbol(DiaChildSymbol);
SymbolModule::ProcessSymbolFunctionArg(
	IN IDiaSymbol* DiaSymbol,
	IN SYMBOL* Symbol
	)
{
	CComPtr<IDiaSymbol> DiaArgumentTypeSymbol;
	DiaSymbol->get_type(&DiaArgumentTypeSymbol);
	Symbol->u.FunctionArg.Type = GetSymbol(DiaArgumentTypeSymbol);
SymbolModule::ProcessSymbolUdt(
	IN IDiaSymbol* DiaSymbol,
	IN SYMBOL* Symbol
	)
{
	DWORD Kind;
	DiaSymbol->get_udtKind(&Kind);
	Symbol->u.Udt.Kind = static_cast<UdtKind>(Kind);
	if (FAILED(DiaSymbol->findChildren(SymTagData, nullptr, nsNone, &DiaSymbolEnumerator)))
	{
		return;
	}
	LONG ChildCount;
	DiaSymbolEnumerator->get_Count(&ChildCount);
	Symbol->u.Udt.FieldCount = static_cast<DWORD>(ChildCount);
	while (SUCCEEDED(DiaSymbolEnumerator->Next(1, &Result, &FetchedSymbolCount)) && (FetchedSymbolCount == 1))
	{
		CComPtr<IDiaSymbol> DiaChildSymbol(Result);
		Member->Name = GetSymbolName(DiaChildSymbol);
		DiaChildSymbol->get_offset(&Offset);
		Member->Offset = static_cast<DWORD>(Offset);
		DiaChildSymbol->get_length(&Bits);
		Member->Bits = static_cast<DWORD>(Bits);
		DiaChildSymbol->get_bitPosition(&Member->BitPosition);
		DiaChildSymbol->get_type(&MemberTypeDiaSymbol);
		Member->Type = GetSymbol(MemberTypeDiaSymbol);
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
			m_SymbolSet.insert(PaddingSymbolArray);
			m_SymbolSet.insert(PaddingSymbolArrayElement);
void SymbolModule::DestroySymbol(
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

func (p *pdb)	{_()(ok bool){//col:532
/*	{ (BasicType)0,   0,  nullptr,            nullptr            },
};*/
return true
}

func (p *pdb)	{_()(ok bool){//col:566
/*	{ (BasicType)0,   0,  nullptr,            nullptr            },
};*/
return true
}

func (p *pdb)PDB::PDB()(ok bool){//col:579
/*PDB::PDB()
{
	m_Impl = new SymbolModule();
PDB::PDB(
	IN const CHAR* Path
	)
{
	m_Impl = new SymbolModule();
	m_Impl->Open(Path);
PDB::~PDB()
{
	delete m_Impl;
}*/
return true
}

func (p *pdb)PDB::Open()(ok bool){//col:638
/*PDB::Open(
	IN const CHAR* Path
	)
{
	return m_Impl->Open(Path);
PDB::IsOpened() const
{
	return m_Impl->IsOpen();
PDB::GetPath() const
{
	return m_Impl->GetPath();
PDB::Close()
{
	m_Impl->Close();
PDB::GetMachineType() const
{
	return m_Impl->GetMachineType();
PDB::GetLanguage() const
{
	return m_Impl->GetLanguage();
PDB::GetSymbolByName(
	IN const CHAR* SymbolName
	)
{
	return m_Impl->GetSymbolByName(SymbolName);
PDB::GetSymbolByTypeId(
	IN DWORD TypeId
	)
{
	return m_Impl->GetSymbolByTypeId(TypeId);
PDB::GetSymbolMap() const
{
	return m_Impl->GetSymbolMap();
PDB::GetSymbolNameMap() const
{
	return m_Impl->GetSymbolNameMap();
PDB::GetFunctionSet() const
{
	return m_Impl->GetFunctionSet();
PDB::GetBasicTypeString(
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

func (p *pdb)PDB::GetBasicTypeString()(ok bool){//col:659
/*PDB::GetBasicTypeString(
	IN const SYMBOL* Symbol,
	IN BOOL UseStdInt
	)
{
	return GetBasicTypeString(Symbol->BaseType, Symbol->Size, UseStdInt);
PDB::GetUdtKindString(
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

func (p *pdb)PDB::IsUnnamedSymbol()(ok bool){//col:667
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



