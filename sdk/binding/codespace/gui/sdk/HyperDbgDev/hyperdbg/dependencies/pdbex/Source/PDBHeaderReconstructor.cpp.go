package Source
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\pdbex\Source\PDBHeaderReconstructor.cpp.back

type (
PdbHeaderReconstructor interface{
PDBHeaderReconstructor::PDBHeaderReconstructor()(ok bool)//col:11
PDBHeaderReconstructor::Clear()(ok bool)//col:531
}
pdbHeaderReconstructor struct{}
)

func NewPdbHeaderReconstructor()PdbHeaderReconstructor{ return & pdbHeaderReconstructor{} }

func (p *pdbHeaderReconstructor)PDBHeaderReconstructor::PDBHeaderReconstructor()(ok bool){//col:11
/*PDBHeaderReconstructor::PDBHeaderReconstructor(
	Settings* VisitorSettings
	)
{
	static Settings DefaultSettings;
	if (VisitorSettings == nullptr)
	{
		VisitorSettings = &DefaultSettings;
	}
	m_Settings = VisitorSettings;
}*/
return true
}

func (p *pdbHeaderReconstructor)PDBHeaderReconstructor::Clear()(ok bool){//col:531
/*PDBHeaderReconstructor::Clear()
{
	assert(m_Depth == 0);
	m_UnnamedSymbols.clear();
	m_CorrectedSymbolNames.clear();
	m_VisitedSymbols.clear();
PDBHeaderReconstructor::GetCorrectedSymbolName(
	const SYMBOL* Symbol
	) const
{
	if (!m_CorrectedSymbolNames.contains(Symbol))
	{
		std::string CorrectedName;
		CorrectedName += m_Settings->SymbolPrefix;
		if (PDB::IsUnnamedSymbol(Symbol))
		{
			if (m_Settings->MicrosoftTypedefs)
			{
				CorrectedName += "_";
			}
			m_UnnamedSymbols.insert(Symbol);
			CorrectedName += m_Settings->UnnamedTypePrefix + std::to_string(m_UnnamedSymbols.size());
PDBHeaderReconstructor::OnEnumType(
	const SYMBOL* Symbol
	)
{
	std::string CorrectedName = GetCorrectedSymbolName(Symbol);
	bool Expand = ShouldExpand(Symbol);
	MarkAsVisited(Symbol);
	if (!Expand)
	{
		Write("enum %s", CorrectedName.c_str());
PDBHeaderReconstructor::OnEnumTypeBegin(
	const SYMBOL* Symbol
	)
{
	std::string CorrectedName = GetCorrectedSymbolName(Symbol);
	WriteTypedefBegin(Symbol);
	Write("enum");
	Write(" %s", CorrectedName.c_str());
	Write("\n");
	WriteIndent();
	Write("{\n");
PDBHeaderReconstructor::OnEnumTypeEnd(
	const SYMBOL* Symbol
	)
{
	m_Depth -= 1;
	WriteIndent();
	Write("}");
	WriteTypedefEnd(Symbol);
	if (m_Depth == 0)
	{
		Write(";\n\n");
PDBHeaderReconstructor::OnEnumField(
	const SYMBOL_ENUM_FIELD* EnumField
	)
{
	WriteIndent();
	Write("%s = ", EnumField->Name);
	WriteVariant(&EnumField->Value);
	Write(",\n");
PDBHeaderReconstructor::OnUdt(
	const SYMBOL* Symbol
	)
{
	bool Expand = ShouldExpand(Symbol);
	MarkAsVisited(Symbol);
	if (!Expand)
	{
		std::string CorrectedName = GetCorrectedSymbolName(Symbol);
        if (g_ShowInOffestFormat)
        {
            Write("%s", CorrectedName.c_str());
		WriteConstAndVolatile(Symbol);
		Write("%s %s", PDB::GetUdtKindString(Symbol->u.Udt.Kind), CorrectedName.c_str());
		if (m_Depth == 0)
		{
			Write(";\n\n");
PDBHeaderReconstructor::OnUdtBegin(
	const SYMBOL* Symbol
	)
{
#ifdef HYPERDBG_CODES
    if (g_ShowInOffestFormat)
    {
    }
    else
    {
#endif
	WriteTypedefBegin(Symbol);
	WriteConstAndVolatile(Symbol);
	Write("%s", PDB::GetUdtKindString(Symbol->u.Udt.Kind));
	if (!PDB::IsUnnamedSymbol(Symbol))
	{
		std::string CorrectedName = GetCorrectedSymbolName(Symbol);
		Write(" %s", CorrectedName.c_str());
	Write("\n");
    if (g_ShowInOffestFormat)
    {
    }
    else
    {
#endif
	WriteIndent();
	Write("{\n");
PDBHeaderReconstructor::OnUdtEnd(
	const SYMBOL* Symbol
	)
{
	m_Depth -= 1;
#ifdef HYPERDBG_CODES
    if (g_ShowInOffestFormat)
    {
    }
    else
    {
#endif
	WriteIndent();
	Write("}");
	WriteTypedefEnd(Symbol);
	if (m_Depth == 0)
	{
#ifdef HYPERDBG_CODES
        if (g_ShowInOffestFormat)
        {
        }
        else
        {
#endif
		Write(";");
    if (g_ShowInOffestFormat)
    {
    }
    else
    {
#endif
	Write("", Symbol->Size);
	if (m_Depth == 0)
	{
		Write("\n\n");
PDBHeaderReconstructor::OnUdtFieldBegin(
	const SYMBOL_UDT_FIELD* UdtField
	)
{
	WriteIndent();
	if (UdtField->Type->Tag != SymTagUDT ||
	    ShouldExpand(UdtField->Type) == false)
	{
		WriteOffset(UdtField, GetParentOffset());
	AppendToTest(UdtField);
	m_OffsetStack.push_back(UdtField->Offset);
PDBHeaderReconstructor::OnUdtFieldEnd(
	const SYMBOL_UDT_FIELD* UdtField
	)
{
	m_OffsetStack.pop_back();
PDBHeaderReconstructor::OnUdtField(
	const SYMBOL_UDT_FIELD* UdtField,
	UdtFieldDefinitionBase* MemberDefinition
	)
{
	Write("%s", MemberDefinition->GetPrintableDefinition().c_str());
    if (g_ShowInOffestFormat)
    {
		if (UdtField->Bits != 0)
        {
            Write(", Pos %i, %i Bit", UdtField->BitPosition, UdtField->Bits);
	if (UdtField->Bits != 0)
	{
		Write(" : %i", UdtField->Bits);
	Write(";");
	if (UdtField->Bits != 0)
	{
		Write("", UdtField->BitPosition);
    if (g_MappingBufferAddress != NULL)
    {
        CHAR * StartAddressOfBuffer = g_MappingBufferAddress + UdtField->Offset;
        UINT64 TempBuffer           = NULL;
        if (sizeof(CHAR) <= UdtField->Type->Size && UdtField->Type->Size <= sizeof(UINT64))
        {
            Write(" : ");
            memcpy(&TempBuffer, StartAddressOfBuffer, UdtField->Type->Size);
			if (UdtField->Bits != 0)
            {
                UINT32  CurrentBit = 0;
                UINT64 BitsFormat = ExtractBits(TempBuffer,
					(UINT64)UdtField->BitPosition,
					(UINT64)(UdtField->BitPosition + UdtField->Bits - 1));
                Write("0y");
				while (CurrentBit < UdtField->Bits)
                {
                    if (BitsFormat & 0x1)
                    {
                        Write("1");
                        Write("0");
				if (UdtField->Bits >= 2)
                {
                    Write(" (0x%llx)", ExtractedBits);
			else if (TempBuffer == NULL)
            {
                Write("(null)");
            else if (UdtField->Type->Size == sizeof(UINT64))
            {
                Write("%s", SymSeparateTo64BitValue(TempBuffer).c_str());
                Write("0x%x",TempBuffer);
        else if (UdtField->Type->Name != NULL && UdtField->Type->Name[0] == '\0')
        {
            if (UdtField->Type->Tag == SymTagArrayType)
            {
                Write(" : (array)");
        else if (UdtField->Type->Name != NULL)
        {
            Write(" : %s", UdtField->Type->Name);
            if (UdtField->Type->Size == sizeof(LIST_ENTRY) &&
                strcmp(UdtField->Type->Name, "_LIST_ENTRY") == 0)
            {
                PLIST_ENTRY ListEntry = (PLIST_ENTRY)(g_MappingBufferAddress + UdtField->Offset);
                Write(" [ %s - %s ]",
					SymSeparateTo64BitValue((UINT64)ListEntry->Flink).c_str(),
                      SymSeparateTo64BitValue((UINT64)ListEntry->Blink).c_str());
	Write("\n");
PDBHeaderReconstructor::OnAnonymousUdtBegin(
	UdtKind Kind,
	const SYMBOL_UDT_FIELD* FirstUdtField
	)
{
#ifdef HYPERDBG_CODES
    if (g_ShowInOffestFormat)
    {
    }
    else
    {
#endif
	WriteIndent();
	Write("%s\n", PDB::GetUdtKindString(Kind));
	WriteIndent();
	Write("{\n");
PDBHeaderReconstructor::OnAnonymousUdtEnd(
	UdtKind Kind,
	const SYMBOL_UDT_FIELD* FirstUdtField,
	const SYMBOL_UDT_FIELD* LastUdtField,
	DWORD Size
	)
{
	m_Depth -= 1;
#ifdef HYPERDBG_CODES
    if (g_ShowInOffestFormat)
    {
    }
    else
    {
#endif
	WriteIndent();
	Write("}");
	WriteUnnamedDataType(Kind);
    if (g_ShowInOffestFormat)
    {
    }
    else
    {
#endif
    Write(";");
	Write("", Size);
    Write("\n");
PDBHeaderReconstructor::OnUdtFieldBitFieldBegin(
	const SYMBOL_UDT_FIELD* FirstUdtFieldBitField,
	const SYMBOL_UDT_FIELD* LastUdtFieldBitField
	)
{
	if (m_Settings->AllowBitFieldsInUnion == false)
	{
		if (FirstUdtFieldBitField != LastUdtFieldBitField)
		{
#ifdef HYPERDBG_CODES
            if (g_ShowInOffestFormat)
            {
            }
            else
            {
#endif
			WriteIndent();
			Write("%s\n", PDB::GetUdtKindString(UdtStruct));
			WriteIndent();
			Write("{\n");
PDBHeaderReconstructor::OnUdtFieldBitFieldEnd(
	const SYMBOL_UDT_FIELD* FirstUdtFieldBitField,
	const SYMBOL_UDT_FIELD* LastUdtFieldBitField
	)
{
	if (m_Settings->AllowBitFieldsInUnion == false)
	{
		if (FirstUdtFieldBitField != LastUdtFieldBitField)
		{
			m_Depth -= 1;
#ifdef HYPERDBG_CODES
            if (g_ShowInOffestFormat)
            {
            }
            else
            {
#endif
			WriteIndent();
			Write("};\n");
PDBHeaderReconstructor::OnPaddingMember(
	const SYMBOL_UDT_FIELD* UdtField,
	BasicType PaddingBasicType,
	DWORD PaddingBasicTypeSize,
	DWORD PaddingSize
	)
{
	if (m_Settings->CreatePaddingMembers)
	{
		WriteIndent();
		WriteOffset(UdtField, -((int)PaddingSize * (int)PaddingBasicTypeSize));
		Write(
			"%s %s%u",
			PDB::GetBasicTypeString(PaddingBasicType, PaddingBasicTypeSize),
			m_Settings->PaddingMemberPrefix.c_str(),
			m_PaddingMemberCounter++
			);
		if (PaddingSize > 1)
		{
			Write("[%u]", PaddingSize);
        if (g_ShowInOffestFormat)
        {
            Write("\n");
		Write(";\n");
PDBHeaderReconstructor::OnPaddingBitFieldField(
	const SYMBOL_UDT_FIELD* UdtField,
	const SYMBOL_UDT_FIELD* PreviousUdtField
	)
{
	WriteIndent();
	WriteOffset(UdtField, GetParentOffset());
	if (m_Settings->BitFieldPaddingMemberPrefix.empty())
	{
		Write(
			"%s",
			PDB::GetBasicTypeString(UdtField->Type) 
		);
		Write(
			"%s %s%u",
			PDB::GetBasicTypeString(UdtField->Type), 
			m_Settings->PaddingMemberPrefix.c_str(),
			m_PaddingMemberCounter++
		);
		? UdtField->BitPosition - (PreviousUdtField->BitPosition + PreviousUdtField->Bits)
		: UdtField->BitPosition;
	DWORD BitPosition = PreviousUdtField
		? PreviousUdtField->BitPosition + PreviousUdtField->Bits
		: 0;
	assert(Bits != 0);
    if (g_ShowInOffestFormat)
    {
    Write(" : Pos %i, %i Bit", BitPosition, Bits);
	Write(" : %i", Bits);
	Write(";");
	Write("", BitPosition);
	Write("\n");
PDBHeaderReconstructor::Write(
	const char* Format,
	...
	)
{
	char TempBuffer[8 * 1024];
	va_list ArgPtr;
	va_start(ArgPtr, Format);
	vsprintf_s(TempBuffer, Format, ArgPtr);
	va_end(ArgPtr);
    if (!g_IsOutputToFile)
    {
        g_MessageHandler(TempBuffer);
        m_Settings->OutputFile->write(TempBuffer, strlen(TempBuffer));
PDBHeaderReconstructor::WriteIndent()
{
#ifdef HYPERDBG_CODES
    if (g_ShowInOffestFormat)
    {
    }
    else
    {
#endif
        for (DWORD i = 0; i < m_Depth; i++)
        {
            Write("  ");
PDBHeaderReconstructor::WriteVariant(
	const VARIANT* v
	)
{
	switch (v->vt)
	{
		case VT_I1:
			Write("%d", (INT)v->cVal);
			Write("0x%x", (UINT)v->cVal);
			Write("%d", (UINT)v->iVal);
			Write("0x%x", (UINT)v->iVal);
			Write("%d", (UINT)v->lVal);
			Write("0x%x", (UINT)v->lVal);
PDBHeaderReconstructor::WriteUnnamedDataType(
	UdtKind Kind
	)
{
	if (m_Settings->AllowAnonymousDataTypes == false)
	{
		switch (Kind)
		{
			case UdtStruct:
			case UdtClass:
				Write(" %s", m_Settings->AnonymousStructPrefix.c_str());
				Write(" %s", m_Settings->AnonymousUnionPrefix.c_str());
				assert(0);
		if (m_AnonymousDataTypeCounter++ > 0)
		{
			Write("%u", m_AnonymousDataTypeCounter);
PDBHeaderReconstructor::WriteTypedefBegin(
	const SYMBOL* Symbol
	)
{
	std::string CorrectedName = GetCorrectedSymbolName(Symbol);
	if (UseTypedef && m_Depth == 0)
	{
		Write("typedef ");
PDBHeaderReconstructor::WriteTypedefEnd(
	const SYMBOL* Symbol
	)
{
	std::string CorrectedName = GetCorrectedSymbolName(Symbol);
	if (UseTypedef && m_Depth == 0)
	{
#ifdef HYPERDBG_CODES
        if (g_ShowInOffestFormat)
        {
        }
        else
        {
#endif
		Write(" %s, *P%s", &CorrectedName[1], &CorrectedName[1]);
PDBHeaderReconstructor::WriteConstAndVolatile(
	const SYMBOL* Symbol
	)
{
	if (m_Depth != 0)
	{
		if (Symbol->IsConst)
		{
			Write("const ");
		if (Symbol->IsVolatile)
		{
			Write("volatile ");
PDBHeaderReconstructor::WriteOffset(
	const SYMBOL_UDT_FIELD* UdtField,
	int PaddingOffset
	)
{
	if (m_Settings->ShowOffsets)
	{
#ifdef HYPERDBG_CODES
        if (g_ShowInOffestFormat)
        {
            Write("  +0x%04x ", UdtField->Offset + PaddingOffset);
            Write( ", UdtField->Offset + PaddingOffset);
PDBHeaderReconstructor::HasBeenVisited(
	const SYMBOL* Symbol
	) const
{
	std::string CorrectedName = GetCorrectedSymbolName(Symbol);
	return m_VisitedSymbols.find(CorrectedName) != m_VisitedSymbols.end();
PDBHeaderReconstructor::MarkAsVisited(
	const SYMBOL* Symbol
	)
{
	std::string CorrectedName = GetCorrectedSymbolName(Symbol);
	m_VisitedSymbols.insert(CorrectedName);
PDBHeaderReconstructor::GetParentOffset() const
{
	return std::accumulate(m_OffsetStack.begin(), m_OffsetStack.end(), (DWORD)0);
PDBHeaderReconstructor::AppendToTest(
	const SYMBOL_UDT_FIELD* UdtField
	)
{
	if (m_Settings->TestFile != nullptr &&
	    m_OffsetStack.empty() &&
	    UdtField->Bits == 0)
	{
		static const char TestFormatString[] =
			"\t"
			"printf("
				"\"[%%c] 0x%%04x - 0x%04x (%s %s.%s)\\n\", "
				"offsetof(%s %s, %s) == %u ? ' ' : '!', "
				"(unsigned)offsetof(%s %s, %s)"
			");";
			"printf("
			"\"\\n\""
			");";
		std::string CorrectedSymbolName = GetCorrectedSymbolName(UdtField->Parent);
		if (CorrectedSymbolName != LastTestedUdt)
		{
			(*m_Settings->TestFile) << TestDelimiterString << std::endl;
		}
		LastTestedUdt = CorrectedSymbolName;
		static char FormattedStringBuffer[4096];
		sprintf_s(
			FormattedStringBuffer,
			TestFormatString,
			UdtField->Offset,
			PDB::GetUdtKindString(UdtField->Parent->u.Udt.Kind),
			CorrectedSymbolName.c_str(),
			UdtField->Name,
			PDB::GetUdtKindString(UdtField->Parent->u.Udt.Kind),
			CorrectedSymbolName.c_str(),
			UdtField->Name,
			UdtField->Offset,
			PDB::GetUdtKindString(UdtField->Parent->u.Udt.Kind),
			CorrectedSymbolName.c_str(),
			UdtField->Name
			);
		(*m_Settings->TestFile) << FormattedStringBuffer << std::endl;
	}
}*/
return true
}



