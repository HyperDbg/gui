package Source
type (
Pdb interface{
		SymbolModuleBase()(ok bool)//col:43
SymbolModuleBase::SymbolModuleBase()(ok bool)//col:54
SymbolModuleBase::LoadDiaViaCoCreateInstance()(ok bool)//col:66
SymbolModuleBase::LoadDiaViaLoadLibrary()(ok bool)//col:98
SymbolModuleBase::Open()(ok bool)//col:198
SymbolModuleBase::Close()(ok bool)//col:208
SymbolModuleBase::IsOpen()(ok bool)//col:214
		SymbolModule()(ok bool)//col:359
SymbolModule::SymbolModule()(ok bool)//col:364
SymbolModule::~SymbolModule()(ok bool)//col:369
SymbolModule::Open()(ok bool)//col:394
SymbolModule::IsOpen()(ok bool)//col:400
SymbolModule::GetPath()(ok bool)//col:406
SymbolModule::Close()(ok bool)//col:423
SymbolModule::GetMachineType()(ok bool)//col:429
SymbolModule::GetLanguage()(ok bool)//col:435
SymbolModule::GetSymbolName()(ok bool)//col:473
SymbolModule::GetSymbolByName()(ok bool)//col:482
SymbolModule::GetSymbolByTypeId()(ok bool)//col:491
SymbolModule::GetSymbol()(ok bool)//col:521
SymbolModule::BuildSymbolMapFromEnumerator()(ok bool)//col:537
SymbolModule::BuildFunctionSetFromEnumerator()(ok bool)//col:566
SymbolModule::BuildSymbolMap()(ok bool)//col:588
SymbolModule::GetSymbolMap()(ok bool)//col:594
SymbolModule::GetSymbolNameMap()(ok bool)//col:600
SymbolModule::GetFunctionSet()(ok bool)//col:606
SymbolModule::InitSymbol()(ok bool)//col:653
SymbolModule::ProcessSymbolBase()(ok bool)//col:662
SymbolModule::ProcessSymbolEnum()(ok bool)//col:701
SymbolModule::ProcessSymbolTypedef()(ok bool)//col:714
SymbolModule::ProcessSymbolPointer()(ok bool)//col:745
SymbolModule::ProcessSymbolArray()(ok bool)//col:759
SymbolModule::ProcessSymbolFunction()(ok bool)//col:816
SymbolModule::ProcessSymbolFunctionArg()(ok bool)//col:828
SymbolModule::ProcessSymbolUdt()(ok bool)//col:931
void SymbolModule::DestroySymbol()(ok bool)//col:963
	{ ()(ok bool)//col:1010
	{ ()(ok bool)//col:1045
PDB::PDB()(ok bool)//col:1050
PDB::PDB()(ok bool)//col:1058
PDB::~PDB()(ok bool)//col:1063
PDB::Open()(ok bool)//col:1071
PDB::IsOpened()(ok bool)//col:1077
PDB::GetPath()(ok bool)//col:1083
PDB::Close()(ok bool)//col:1089
PDB::GetMachineType()(ok bool)//col:1095
PDB::GetLanguage()(ok bool)//col:1101
PDB::GetSymbolByName()(ok bool)//col:1109
PDB::GetSymbolByTypeId()(ok bool)//col:1117
PDB::GetSymbolMap()(ok bool)//col:1123
PDB::GetSymbolNameMap()(ok bool)//col:1129
PDB::GetFunctionSet()(ok bool)//col:1135
PDB::GetBasicTypeString()(ok bool)//col:1159
PDB::GetBasicTypeString()(ok bool)//col:1168
PDB::GetUdtKindString()(ok bool)//col:1187
PDB::IsUnnamedSymbol()(ok bool)//col:1197
}

)
func NewPdb() { return & pdb{} }
func (p *pdb)		SymbolModuleBase()(ok bool){//col:43
return true
}

func (p *pdb)SymbolModuleBase::SymbolModuleBase()(ok bool){//col:54
return true
}

func (p *pdb)SymbolModuleBase::LoadDiaViaCoCreateInstance()(ok bool){//col:66
return true
}

func (p *pdb)SymbolModuleBase::LoadDiaViaLoadLibrary()(ok bool){//col:98
return true
}

func (p *pdb)SymbolModuleBase::Open()(ok bool){//col:198
return true
}

func (p *pdb)SymbolModuleBase::Close()(ok bool){//col:208
return true
}

func (p *pdb)SymbolModuleBase::IsOpen()(ok bool){//col:214
return true
}

func (p *pdb)		SymbolModule()(ok bool){//col:359
return true
}

func (p *pdb)SymbolModule::SymbolModule()(ok bool){//col:364
return true
}

func (p *pdb)SymbolModule::~SymbolModule()(ok bool){//col:369
return true
}

func (p *pdb)SymbolModule::Open()(ok bool){//col:394
return true
}

func (p *pdb)SymbolModule::IsOpen()(ok bool){//col:400
return true
}

func (p *pdb)SymbolModule::GetPath()(ok bool){//col:406
return true
}

func (p *pdb)SymbolModule::Close()(ok bool){//col:423
return true
}

func (p *pdb)SymbolModule::GetMachineType()(ok bool){//col:429
return true
}

func (p *pdb)SymbolModule::GetLanguage()(ok bool){//col:435
return true
}

func (p *pdb)SymbolModule::GetSymbolName()(ok bool){//col:473
return true
}

func (p *pdb)SymbolModule::GetSymbolByName()(ok bool){//col:482
return true
}

func (p *pdb)SymbolModule::GetSymbolByTypeId()(ok bool){//col:491
return true
}

func (p *pdb)SymbolModule::GetSymbol()(ok bool){//col:521
return true
}

func (p *pdb)SymbolModule::BuildSymbolMapFromEnumerator()(ok bool){//col:537
return true
}

func (p *pdb)SymbolModule::BuildFunctionSetFromEnumerator()(ok bool){//col:566
return true
}

func (p *pdb)SymbolModule::BuildSymbolMap()(ok bool){//col:588
return true
}

func (p *pdb)SymbolModule::GetSymbolMap()(ok bool){//col:594
return true
}

func (p *pdb)SymbolModule::GetSymbolNameMap()(ok bool){//col:600
return true
}

func (p *pdb)SymbolModule::GetFunctionSet()(ok bool){//col:606
return true
}

func (p *pdb)SymbolModule::InitSymbol()(ok bool){//col:653
return true
}

func (p *pdb)SymbolModule::ProcessSymbolBase()(ok bool){//col:662
return true
}

func (p *pdb)SymbolModule::ProcessSymbolEnum()(ok bool){//col:701
return true
}

func (p *pdb)SymbolModule::ProcessSymbolTypedef()(ok bool){//col:714
return true
}

func (p *pdb)SymbolModule::ProcessSymbolPointer()(ok bool){//col:745
return true
}

func (p *pdb)SymbolModule::ProcessSymbolArray()(ok bool){//col:759
return true
}

func (p *pdb)SymbolModule::ProcessSymbolFunction()(ok bool){//col:816
return true
}

func (p *pdb)SymbolModule::ProcessSymbolFunctionArg()(ok bool){//col:828
return true
}

func (p *pdb)SymbolModule::ProcessSymbolUdt()(ok bool){//col:931
return true
}

func (p *pdb)void SymbolModule::DestroySymbol()(ok bool){//col:963
return true
}

func (p *pdb)	{ ()(ok bool){//col:1010
return true
}

func (p *pdb)	{ ()(ok bool){//col:1045
return true
}

func (p *pdb)PDB::PDB()(ok bool){//col:1050
return true
}

func (p *pdb)PDB::PDB()(ok bool){//col:1058
return true
}

func (p *pdb)PDB::~PDB()(ok bool){//col:1063
return true
}

func (p *pdb)PDB::Open()(ok bool){//col:1071
return true
}

func (p *pdb)PDB::IsOpened()(ok bool){//col:1077
return true
}

func (p *pdb)PDB::GetPath()(ok bool){//col:1083
return true
}

func (p *pdb)PDB::Close()(ok bool){//col:1089
return true
}

func (p *pdb)PDB::GetMachineType()(ok bool){//col:1095
return true
}

func (p *pdb)PDB::GetLanguage()(ok bool){//col:1101
return true
}

func (p *pdb)PDB::GetSymbolByName()(ok bool){//col:1109
return true
}

func (p *pdb)PDB::GetSymbolByTypeId()(ok bool){//col:1117
return true
}

func (p *pdb)PDB::GetSymbolMap()(ok bool){//col:1123
return true
}

func (p *pdb)PDB::GetSymbolNameMap()(ok bool){//col:1129
return true
}

func (p *pdb)PDB::GetFunctionSet()(ok bool){//col:1135
return true
}

func (p *pdb)PDB::GetBasicTypeString()(ok bool){//col:1159
return true
}

func (p *pdb)PDB::GetBasicTypeString()(ok bool){//col:1168
return true
}

func (p *pdb)PDB::GetUdtKindString()(ok bool){//col:1187
return true
}

func (p *pdb)PDB::IsUnnamedSymbol()(ok bool){//col:1197
return true
}

