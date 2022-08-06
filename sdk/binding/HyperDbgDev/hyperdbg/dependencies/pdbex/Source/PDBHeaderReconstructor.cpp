#include "PDBHeaderReconstructor.h"
#include "PDBReconstructorBase.h"

#include <iostream>
#include <numeric> 
#include <string>
#include <map>
#include <set>

#include <cassert>

#include "HyperDbgExport.h"
#include "HyperDbgGlobals.h"


PDBHeaderReconstructor::PDBHeaderReconstructor(
	Settings* VisitorSettings
	)
{
	static Settings DefaultSettings;

	if (VisitorSettings == nullptr)
	{
		VisitorSettings = &DefaultSettings;
	}

	m_Settings = VisitorSettings;
}

void
PDBHeaderReconstructor::Clear()
{
	assert(m_Depth == 0);

	m_AnonymousDataTypeCounter = 0;
	m_PaddingMemberCounter = 0;

	m_UnnamedSymbols.clear();
	m_CorrectedSymbolNames.clear();
	m_VisitedSymbols.clear();
}

const std::string&
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
		}
		else
		{
			CorrectedName += Symbol->Name;
		}

		CorrectedName += m_Settings->SymbolSuffix;

		m_CorrectedSymbolNames[Symbol] = CorrectedName;
	}

	return m_CorrectedSymbolNames[Symbol];
}

bool
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
	}

	return Expand;
}

void
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

	m_Depth += 1;
}

void
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
	}
}

void
PDBHeaderReconstructor::OnEnumField(
	const SYMBOL_ENUM_FIELD* EnumField
	)
{
	WriteIndent();
	Write("%s = ", EnumField->Name);

	WriteVariant(&EnumField->Value);
	Write(",\n");
}

bool
PDBHeaderReconstructor::OnUdt(
	const SYMBOL* Symbol
	)
{
	bool Expand = ShouldExpand(Symbol);

	MarkAsVisited(Symbol);

	if (!Expand)
	{
		std::string CorrectedName = GetCorrectedSymbolName(Symbol);

#ifdef HYPERDBG_CODES

        if (g_ShowInOffestFormat)
        {
			
			
			
            Write("%s", CorrectedName.c_str());
        }
        else
        {
#endif
		WriteConstAndVolatile(Symbol);

		Write("%s %s", PDB::GetUdtKindString(Symbol->u.Udt.Kind), CorrectedName.c_str());

#ifdef HYPERDBG_CODES
        }
#endif
		
		
		
		

		if (m_Depth == 0)
		{
			Write(";\n\n");
		}
	}

	return Expand;
}

void
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

#ifdef HYPERDBG_CODES
    }
#endif

	if (!PDB::IsUnnamedSymbol(Symbol))
	{
		std::string CorrectedName = GetCorrectedSymbolName(Symbol);
		Write(" %s", CorrectedName.c_str());
	}

	Write("\n");

#ifdef HYPERDBG_CODES

    if (g_ShowInOffestFormat)
    {
        
        
        
    }
    else
    {
#endif
	WriteIndent();
	Write("{\n");
#ifdef HYPERDBG_CODES
    }
#endif

	m_Depth += 1;
}

void
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

#ifdef HYPERDBG_CODES
    }
#endif

	
	
	

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

#ifdef HYPERDBG_CODES
        }
#endif

	}

#ifdef HYPERDBG_CODES

    if (g_ShowInOffestFormat)
    {
        
        
        
    }
    else
    {
#endif

	Write(" /* size: 0x%04x */", Symbol->Size);

#ifdef HYPERDBG_CODES
}
#endif

	if (m_Depth == 0)
	{
		Write("\n\n");
	}
}

void
PDBHeaderReconstructor::OnUdtFieldBegin(
	const SYMBOL_UDT_FIELD* UdtField
	)
{
	WriteIndent();

	
	
	

	if (UdtField->Type->Tag != SymTagUDT ||
	    ShouldExpand(UdtField->Type) == false)
	{
		WriteOffset(UdtField, GetParentOffset());
	}

	AppendToTest(UdtField);

	
	
	
	

	m_OffsetStack.push_back(UdtField->Offset);
}

void
PDBHeaderReconstructor::OnUdtFieldEnd(
	const SYMBOL_UDT_FIELD* UdtField
	)
{

	
	
	

	m_OffsetStack.pop_back();
}

void
PDBHeaderReconstructor::OnUdtField(
	const SYMBOL_UDT_FIELD* UdtField,
	UdtFieldDefinitionBase* MemberDefinition
	)
{

	Write("%s", MemberDefinition->GetPrintableDefinition().c_str());

	
	
	

	#ifdef HYPERDBG_CODES

    if (g_ShowInOffestFormat)
    {
		if (UdtField->Bits != 0)
        {
            Write(", Pos %i, %i Bit", UdtField->BitPosition, UdtField->Bits);
        }
    }
    else
    {
#endif

	if (UdtField->Bits != 0)
	{
		Write(" : %i", UdtField->Bits);
	}

	Write(";");

	if (UdtField->Bits != 0)
	{
		Write(" /* bit position: %i */", UdtField->BitPosition);
	}
#ifdef HYPERDBG_CODES
    }
#endif


#ifdef HYPERDBG_CODES

	
	
	
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
                const UINT64 ExtractedBits = BitsFormat;


                
                Write("0y");

				while (CurrentBit < UdtField->Bits)
                {
                    if (BitsFormat & 0x1)
                    {
                        Write("1");
                    }
                    else
                    {
                        Write("0");
                    }

                    CurrentBit++;
                    BitsFormat = BitsFormat >> 1;
                }

				if (UdtField->Bits >= 2)
                {
                    Write(" (0x%llx)", ExtractedBits);
                }

            }
			else if (TempBuffer == NULL)
            {
                Write("(null)");
            }
            else if (UdtField->Type->Size == sizeof(UINT64))
            {
                Write("%s", SymSeparateTo64BitValue(TempBuffer).c_str());
            }
            else
            {
                Write("0x%x",TempBuffer);
            }
        }
        else if (UdtField->Type->Name != NULL && UdtField->Type->Name[0] == '\0')
        {
			
			
			
            if (UdtField->Type->Tag == SymTagArrayType)
            {
                Write(" : (array)");
            }
            else
            {
				
				
				
            }
        }
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
            }
        }


    }
#endif



	Write("\n");
}

void
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

#ifdef HYPERDBG_CODES
    }
#endif

	m_Depth += 1;
}

void
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

#ifdef HYPERDBG_CODES
    }
#endif

	WriteUnnamedDataType(Kind);

#ifdef HYPERDBG_CODES

    if (g_ShowInOffestFormat)
    {
        
        
        
    }
    else
    {
#endif
    Write(";");
	Write(" /* size: 0x%04x */", Size);
    Write("\n");

#ifdef HYPERDBG_CODES
    }
#endif

}

void
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
			Write("%s /* bitfield */\n", PDB::GetUdtKindString(UdtStruct));

			WriteIndent();
			Write("{\n");
#ifdef HYPERDBG_CODES
            }
#endif
			m_Depth += 1;
		}
	}
}

void
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
			Write("}; /* bitfield */\n");

#ifdef HYPERDBG_CODES
            }
#endif
		}
	}
}

void
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
		}

#ifdef HYPERDBG_CODES

        if (g_ShowInOffestFormat)
        {
            
            
            
            Write("\n");
        }
        else
        {
#endif
		Write(";\n");

#ifdef HYPERDBG_CODES
        }
#endif

	}
}

void
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
	}
	else
	{
		Write(
			"%s %s%u",
			PDB::GetBasicTypeString(UdtField->Type), 
			m_Settings->PaddingMemberPrefix.c_str(),
			m_PaddingMemberCounter++
		);
	}

	
	
	

	DWORD Bits = PreviousUdtField
		? UdtField->BitPosition - (PreviousUdtField->BitPosition + PreviousUdtField->Bits)
		: UdtField->BitPosition;

	DWORD BitPosition = PreviousUdtField
		? PreviousUdtField->BitPosition + PreviousUdtField->Bits
		: 0;

	assert(Bits != 0);

	
#ifdef HYPERDBG_CODES

    if (g_ShowInOffestFormat)
    {

    Write(" : Pos %i, %i Bit", BitPosition, Bits);

    }
    else
    {
#endif


	Write(" : %i", Bits);

	Write(";");

	Write(" /* bit position: %i */", BitPosition);

#ifdef HYPERDBG_CODES
    }
#endif

	Write("\n");
}

void
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

#ifdef HYPERDBG_CODES

    if (!g_IsOutputToFile)
    {
        g_MessageHandler(TempBuffer);
    }
    else
    {
#endif
        m_Settings->OutputFile->write(TempBuffer, strlen(TempBuffer));

#ifdef HYPERDBG_CODES
    }
#endif
    }

void
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
        }

#ifdef HYPERDBG_CODES
    }
#endif

 }

void
PDBHeaderReconstructor::WriteVariant(
	const VARIANT* v
	)
{
	switch (v->vt)
	{
		case VT_I1:
			Write("%d", (INT)v->cVal);
			break;

		case VT_UI1:
			Write("0x%x", (UINT)v->cVal);
			break;

		case VT_I2:
			Write("%d", (UINT)v->iVal);
			break;

		case VT_UI2:
			Write("0x%x", (UINT)v->iVal);
			break;

		case VT_INT:
		case VT_I4:
			Write("%d", (UINT)v->lVal);
			break;

		case VT_UINT:
		case VT_UI4:
			Write("0x%x", (UINT)v->lVal);
			break;
	}
}

void
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
				break;

			case UdtUnion:
				Write(" %s", m_Settings->AnonymousUnionPrefix.c_str());
				break;

			default:
				assert(0);
				break;
		}

		if (m_AnonymousDataTypeCounter++ > 0)
		{
			Write("%u", m_AnonymousDataTypeCounter);
		}
	}
}

void
PDBHeaderReconstructor::WriteTypedefBegin(
	const SYMBOL* Symbol
	)
{
	std::string CorrectedName = GetCorrectedSymbolName(Symbol);
	bool UseTypedef = m_Settings->MicrosoftTypedefs && CorrectedName[0] == '_';

	if (UseTypedef && m_Depth == 0)
	{
		Write("typedef ");
	}
}

void
PDBHeaderReconstructor::WriteTypedefEnd(
	const SYMBOL* Symbol
	)
{
	std::string CorrectedName = GetCorrectedSymbolName(Symbol);
	bool UseTypedef = m_Settings->MicrosoftTypedefs && CorrectedName[0] == '_';

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

#ifdef HYPERDBG_CODES
        }
#endif
	}
    }

void
PDBHeaderReconstructor::WriteConstAndVolatile(
	const SYMBOL* Symbol
	)
{
	if (m_Depth != 0)
	{
		
		
		

		if (Symbol->IsConst)
		{
			Write("const ");
		}

		if (Symbol->IsVolatile)
		{
			Write("volatile ");
		}
	}
}

void
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
        }
        else
        {
#endif
            Write("/* 0x%04x */ ", UdtField->Offset + PaddingOffset);

#ifdef HYPERDBG_CODES
        }
#endif

	}
}

bool
PDBHeaderReconstructor::HasBeenVisited(
	const SYMBOL* Symbol
	) const
{
	std::string CorrectedName = GetCorrectedSymbolName(Symbol);
	return m_VisitedSymbols.find(CorrectedName) != m_VisitedSymbols.end();
}

void
PDBHeaderReconstructor::MarkAsVisited(
	const SYMBOL* Symbol
	)
{
	std::string CorrectedName = GetCorrectedSymbolName(Symbol);
	m_VisitedSymbols.insert(CorrectedName);
}

DWORD
PDBHeaderReconstructor::GetParentOffset() const
{
	return std::accumulate(m_OffsetStack.begin(), m_OffsetStack.end(), (DWORD)0);
}

void
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

		
		
		
		
		
		static const char TestDelimiterString[] =
			"\t"
			"printf("
			"\"\\n\""
			");";

		
		
		
		
		
		
		static std::string LastTestedUdt;

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
}

bool
PDBHeaderReconstructor::ShouldExpand(
	const SYMBOL* Symbol
	) const
{
	bool Expand = false;

	switch (m_Settings->MemberStructExpansion)
	{
		default:
		case MemberStructExpansionType::None:
			Expand = m_Depth == 0;
			break;

		case MemberStructExpansionType::InlineUnnamed:
			Expand = m_Depth == 0 || (Symbol->Tag == SymTagUDT && PDB::IsUnnamedSymbol(Symbol));
			break;

		case MemberStructExpansionType::InlineAll:
			Expand = !HasBeenVisited(Symbol);
			break;
	}

	return Expand && Symbol->Size > 0;
}
