package Source
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\pdbex\Source\PDBExtractor.cpp.back

type (
PdbExtractor interface{
		"int_main()(ok bool)//col:36
PDBExtractor::Run()(ok bool)//col:69
PDBExtractor::PrintUsage()(ok bool)//col:110
PDBExtractor::ParseParameters()(ok bool)//col:299
PDBExtractor::OpenPDBFile()(ok bool)//col:306
PDBExtractor::PrintTestHeader()(ok bool)//col:318
PDBExtractor::PrintTestFooter()(ok bool)//col:325
PDBExtractor::PrintPDBHeader()(ok bool)//col:339
PDBExtractor::PrintPDBDeclarations()(ok bool)//col:363
PDBExtractor::PrintPDBDefinitions()(ok bool)//col:401
PDBExtractor::PrintPDBFunctions()(ok bool)//col:419
PDBExtractor::DumpAllSymbols()(ok bool)//col:430
PDBExtractor::DumpOneSymbol()(ok bool)//col:449
PDBExtractor::DumpAllSymbolsOneByOne()(ok bool)//col:485
PDBExtractor::CloseOpenFiles()(ok bool)//col:496
}
pdbExtractor struct{}
)

func NewPdbExtractor()PdbExtractor{ return & pdbExtractor{} }

func (p *pdbExtractor)		"int_main()(ok bool){//col:36
/*		"int main()\n"
		"{\n";
	static const char TEST_FILE_FOOTER[] =
		"\n"
		"\treturn 0;\n"
		"}\n"
		"\n";
	static const char HEADER_FILE_HEADER[] =
		" * PDB file: %s\n"
		" * Image architecture: %s (0x%04x)\n"
		" *\n"
		" * Dumped by pdbex tool v" PDBEX_VERSION_STRING ", by wbenny\n"
	static const char DEFINITIONS_INCLUDE_STDINT[] =
		"#include <stdint.h>";
	static const char DEFINITIONS_PRAGMA_PACK_BEGIN[] =
		"#include <pshpack1.h>";
	static const char DEFINITIONS_PRAGMA_PACK_END[] =
		"#include <poppack.h>";
	static const char* MESSAGE_INVALID_PARAMETERS =
		"Invalid parameters";
	static const char* MESSAGE_FILE_NOT_FOUND =
		"File not found";
	static const char* MESSAGE_SYMBOL_NOT_FOUND =
		"Symbol not found";
	class PDBDumperException
		: public std::runtime_error
	{
		public:
			PDBDumperException(const char* Message)
				: std::runtime_error(Message)
			{
			}
	};
}*/
return true
}

func (p *pdbExtractor)PDBExtractor::Run()(ok bool){//col:69
/*PDBExtractor::Run(
	int argc,
	char** argv
	)
{
	int Result = ERROR_SUCCESS;
	try
	{
		ParseParameters(argc, argv);
		OpenPDBFile();
		PrintTestHeader();
		if (m_Settings.SymbolName == "*")
		{
			DumpAllSymbols();
		}
		else if (m_Settings.SymbolName == "%")
		{
			DumpAllSymbolsOneByOne();
		}
		else
		{
			DumpOneSymbol();
		}
		PrintTestFooter();
	}
	catch (const PDBDumperException& e)
	{
		std::cerr << e.what() << std::endl;
		Result = EXIT_FAILURE;
	}
	CloseOpenFiles();
	return Result;
}*/
return true
}

func (p *pdbExtractor)PDBExtractor::PrintUsage()(ok bool){//col:110
/*PDBExtractor::PrintUsage()
{
	printf("Extracts types and structures from PDB (Program database).\n");
	printf("Version v%s\n", PDBEX_VERSION_STRING);
	printf("\n");
	printf("pdbex <symbol> <path> [-o <filename>] [-t <filename>] [-e <type>]\n");
	printf("                     [-u <prefix>] [-s prefix] [-r prefix] [-g suffix]\n");
	printf("                     [-p] [-x] [-m] [-b] [-d] [-i] [-l]\n");
	printf("\n");
	printf("<symbol>             Symbol name to extract\n");
	printf("                     Use '*' if all symbols should be extracted.\n");
	printf("                     Use '%%' if all symbols should be extracted separately.\n");
	printf("<path>               Path to the PDB file.\n");
	printf(" -o filename         Specifies the output file.                       (stdout)\n");
	printf(" -t filename         Specifies the output test file.                  (off)\n");
	printf(" -e [n,i,a]          Specifies expansion of nested structures/unions. (i)\n");
	printf("                       n = none            Only top-most type is printed.\n");
	printf("                       i = inline unnamed  Unnamed types are nested.\n");
	printf("                       a = inline all      All types are nested.\n");
	printf(" -u prefix           Unnamed union prefix  (in combination with -d).\n");
	printf(" -s prefix           Unnamed struct prefix (in combination with -d).\n");
	printf(" -r prefix           Prefix for all symbols.\n");
	printf(" -g suffix           Suffix for all symbols.\n");
	printf("\n");
	printf("Following options can be explicitly turned off by adding trailing '-'.\n");
	printf("Example: -p-\n");
	printf(" -p                  Create padding members.                          (T)\n");
	printf(" -x                  Show offsets.                                    (T)\n");
	printf(" -m                  Create Microsoft typedefs.                       (T)\n");
	printf(" -b                  Allow bitfields in union.                        (F)\n");
	printf(" -d                  Allow unnamed data types.                        (T)\n");
	printf(" -i                  Use types from stdint.h instead of native types. (F)\n");
	printf(" -j                  Print definitions of referenced types.           (T)\n");
	printf(" -k                  Print header.                                    (T)\n");
	printf(" -n                  Print declarations.                              (T)\n");
	printf(" -l                  Print definitions.                               (T)\n");
	printf(" -f                  Print functions.                                 (F)\n");
	printf(" -z                  Print #pragma pack directives.                   (T)\n");
	printf(" -y                  Sort declarations and definitions.               (F)\n");
	printf("\n");
}*/
return true
}

func (p *pdbExtractor)PDBExtractor::ParseParameters()(ok bool){//col:299
/*PDBExtractor::ParseParameters(
	int argc,
	char** argv
	)
{
	if ( argc == 1 ||
	    (argc == 2 && strcmp(argv[1], "-h") == 0) ||
	    (argc == 2 && strcmp(argv[1], "--help") == 0))
	{
		PrintUsage();
		exit(EXIT_SUCCESS);
	}
	int ArgumentPointer = 0;
	m_Settings.SymbolName = argv[++ArgumentPointer];
	m_Settings.PdbPath = argv[++ArgumentPointer];
	while (++ArgumentPointer < argc)
	{
		const char* CurrentArgument = argv[ArgumentPointer];
		size_t CurrentArgumentLength = strlen(CurrentArgument);
		const char* NextArgument = ArgumentPointer < argc
			? argv[ArgumentPointer + 1]
			: nullptr;
		size_t NextArgumentLength = NextArgument
			? strlen(CurrentArgument)
			: 0;
		if ((CurrentArgumentLength != 2 && CurrentArgumentLength != 3) ||
		    (CurrentArgumentLength == 2 && CurrentArgument[0] != '-') ||
		    (CurrentArgumentLength == 3 && CurrentArgument[0] != '-' && CurrentArgument[2] != '-'))
		{
			throw PDBDumperException(MESSAGE_INVALID_PARAMETERS);
		}
		bool OffSwitch = CurrentArgumentLength == 3 && CurrentArgument[2] == '-';
#ifdef HYPERDBG_CODES
        g_IsOutputToFile = FALSE;
#endif
		switch (CurrentArgument[1])
		{
			case 'o':
				if (!NextArgument)
				{
					throw PDBDumperException(MESSAGE_INVALID_PARAMETERS);
				}
				++ArgumentPointer;
				m_Settings.OutputFilename = NextArgument;
			
#ifdef HYPERDBG_CODES
                g_IsOutputToFile = TRUE;
#endif
				if (m_Settings.SymbolName != "%")
				{
					m_Settings.PdbHeaderReconstructorSettings.OutputFile = new std::ofstream(
						NextArgument,
						std::ios::out
						);
				}
				else
				{
					m_Settings.PdbHeaderReconstructorSettings.OutputFile = nullptr;
				}
				break;
			case 't':
				if (!NextArgument)
				{
					throw PDBDumperException(MESSAGE_INVALID_PARAMETERS);
				}
				++ArgumentPointer;
				m_Settings.TestFilename = NextArgument;
				m_Settings.PdbHeaderReconstructorSettings.TestFile = new std::ofstream(
					m_Settings.TestFilename,
					std::ios::out
					);
				break;
			case 'e':
				if (!NextArgument)
				{
					throw PDBDumperException(MESSAGE_INVALID_PARAMETERS);
				}
				++ArgumentPointer;
				switch (NextArgument[0])
				{
					case 'n':
						m_Settings.PdbHeaderReconstructorSettings.MemberStructExpansion =
							PDBHeaderReconstructor::MemberStructExpansionType::None;
						break;
					case 'i':
						m_Settings.PdbHeaderReconstructorSettings.MemberStructExpansion =
							PDBHeaderReconstructor::MemberStructExpansionType::InlineUnnamed;
						break;
					case 'a':
						m_Settings.PdbHeaderReconstructorSettings.MemberStructExpansion =
							PDBHeaderReconstructor::MemberStructExpansionType::InlineAll;
						break;
					default:
						m_Settings.PdbHeaderReconstructorSettings.MemberStructExpansion =
							PDBHeaderReconstructor::MemberStructExpansionType::InlineUnnamed;
						break;
				}
				break;
			case 'u':
				if (!NextArgument)
				{
					throw PDBDumperException(MESSAGE_INVALID_PARAMETERS);
				}
				++ArgumentPointer;
				m_Settings.PdbHeaderReconstructorSettings.AnonymousUnionPrefix = NextArgument;
				break;
			case 's':
				if (!NextArgument)
				{
					throw PDBDumperException(MESSAGE_INVALID_PARAMETERS);
				}
				++ArgumentPointer;
				m_Settings.PdbHeaderReconstructorSettings.AnonymousStructPrefix = NextArgument;
				break;
			case 'r':
				if (!NextArgument)
				{
					throw PDBDumperException(MESSAGE_INVALID_PARAMETERS);
				}
				++ArgumentPointer;
				m_Settings.PdbHeaderReconstructorSettings.SymbolPrefix = NextArgument;
				break;
			case 'g':
				if (!NextArgument)
				{
					throw PDBDumperException(MESSAGE_INVALID_PARAMETERS);
				}
				++ArgumentPointer;
				m_Settings.PdbHeaderReconstructorSettings.SymbolSuffix = NextArgument;
				break;
			case 'p':
				m_Settings.PdbHeaderReconstructorSettings.CreatePaddingMembers = !OffSwitch;
				break;
			case 'x':
				m_Settings.PdbHeaderReconstructorSettings.ShowOffsets = !OffSwitch;
				break;
			case 'm':
				m_Settings.PdbHeaderReconstructorSettings.MicrosoftTypedefs = !OffSwitch;
				break;
			case 'b':
				m_Settings.PdbHeaderReconstructorSettings.AllowBitFieldsInUnion = !OffSwitch;
				break;
			case 'd':
				m_Settings.PdbHeaderReconstructorSettings.AllowAnonymousDataTypes = !OffSwitch;
				break;
			case 'i':
				m_Settings.UdtFieldDefinitionSettings.UseStdInt = !OffSwitch;
				break;
			case 'j':
				m_Settings.PrintReferencedTypes = !OffSwitch;
				break;
			case 'k':
				m_Settings.PrintHeader = !OffSwitch;
				break;
			case 'n':
				m_Settings.PrintDeclarations = !OffSwitch;
				break;
			case 'l':
				m_Settings.PrintDefinitions = !OffSwitch;
				break;
			case 'f':
				m_Settings.PrintFunctions = !OffSwitch;
				break;
			case 'z':
				m_Settings.PrintPragmaPack = !OffSwitch;
				break;
			case 'y':
				m_Settings.Sort = !OffSwitch;
				break;
			default:
				throw PDBDumperException(MESSAGE_INVALID_PARAMETERS);
		}
	}
	m_HeaderReconstructor = std::make_unique<PDBHeaderReconstructor>(
		&m_Settings.PdbHeaderReconstructorSettings
		);
	m_SymbolVisitor = std::make_unique<PDBSymbolVisitor<UdtFieldDefinition>>(
		m_HeaderReconstructor.get(),
		&m_Settings.UdtFieldDefinitionSettings
		);
	if (m_Settings.Sort)
	{
		m_SymbolSorter = std::make_unique<PDBSymbolSorterAlphabetical>();
	}
	else
	{
		m_SymbolSorter = std::make_unique<PDBSymbolSorter>();
	}
}*/
return true
}

func (p *pdbExtractor)PDBExtractor::OpenPDBFile()(ok bool){//col:306
/*PDBExtractor::OpenPDBFile()
{
	if (m_PDB.Open(m_Settings.PdbPath.c_str()) == FALSE)
	{
		throw PDBDumperException(MESSAGE_FILE_NOT_FOUND);
	}
}*/
return true
}

func (p *pdbExtractor)PDBExtractor::PrintTestHeader()(ok bool){//col:318
/*PDBExtractor::PrintTestHeader()
{
	if (m_Settings.PdbHeaderReconstructorSettings.TestFile != nullptr)
	{
		static char TEST_FILE_HEADER_FORMATTED[16 * 1024];
		sprintf_s(
			TEST_FILE_HEADER_FORMATTED, TEST_FILE_HEADER,
			m_Settings.OutputFilename
			);
		(*m_Settings.PdbHeaderReconstructorSettings.TestFile) << TEST_FILE_HEADER_FORMATTED;
	}
}*/
return true
}

func (p *pdbExtractor)PDBExtractor::PrintTestFooter()(ok bool){//col:325
/*PDBExtractor::PrintTestFooter()
{
	if (m_Settings.PdbHeaderReconstructorSettings.TestFile != nullptr)
	{
		(*m_Settings.PdbHeaderReconstructorSettings.TestFile) << TEST_FILE_FOOTER;
	}
}*/
return true
}

func (p *pdbExtractor)PDBExtractor::PrintPDBHeader()(ok bool){//col:339
/*PDBExtractor::PrintPDBHeader()
{
	if (m_Settings.PrintHeader)
	{
		static const char* const ArchitectureString =
			m_PDB.GetMachineType() == IMAGE_FILE_MACHINE_I386  ? "i386"  :
			m_PDB.GetMachineType() == IMAGE_FILE_MACHINE_AMD64 ? "AMD64" :
			m_PDB.GetMachineType() == IMAGE_FILE_MACHINE_IA64  ? "IA64"  :
			m_PDB.GetMachineType() == IMAGE_FILE_MACHINE_ARMNT ? "ArmNT" :
			m_PDB.GetMachineType() == IMAGE_FILE_MACHINE_ARM64 ? "ARM64" :
			m_PDB.GetMachineType() == IMAGE_FILE_MACHINE_CHPE_X86 ? "CHPE_X86" :
				                                                  "Unknown";
	}
}*/
return true
}

func (p *pdbExtractor)PDBExtractor::PrintPDBDeclarations()(ok bool){//col:363
/*PDBExtractor::PrintPDBDeclarations()
{
	if (m_Settings.PrintDeclarations)
	{
		for (auto&& e : m_SymbolSorter->GetSortedSymbols())
		{
			if (e->Tag == SymTagUDT && !PDB::IsUnnamedSymbol(e))
			{
				*m_Settings.PdbHeaderReconstructorSettings.OutputFile
					<< PDB::GetUdtKindString(e->u.Udt.Kind)
					<< " " << m_HeaderReconstructor->GetCorrectedSymbolName(e) << ";"
					<< std::endl;
			}
			else if (e->Tag == SymTagEnum)
			{
				*m_Settings.PdbHeaderReconstructorSettings.OutputFile
					<< "enum"
					<< " " << m_HeaderReconstructor->GetCorrectedSymbolName(e) << ";"
					<< std::endl;
			}
		}
		*m_Settings.PdbHeaderReconstructorSettings.OutputFile << std::endl;
	}
}*/
return true
}

func (p *pdbExtractor)PDBExtractor::PrintPDBDefinitions()(ok bool){//col:401
/*PDBExtractor::PrintPDBDefinitions()
{
	if (m_Settings.PrintDefinitions)
	{
		if (m_Settings.UdtFieldDefinitionSettings.UseStdInt)
		{
			*m_Settings.PdbHeaderReconstructorSettings.OutputFile
				<< DEFINITIONS_INCLUDE_STDINT
				<< std::endl;
		}
		if (m_Settings.PrintPragmaPack)
		{
			*m_Settings.PdbHeaderReconstructorSettings.OutputFile
				<< DEFINITIONS_PRAGMA_PACK_BEGIN
				<< std::endl;
		}
		for (auto&& e : m_SymbolSorter->GetSortedSymbols())
		{
			bool Expand = true;
			if (m_Settings.PdbHeaderReconstructorSettings.MemberStructExpansion == PDBHeaderReconstructor::MemberStructExpansionType::InlineUnnamed &&
			    e->Tag == SymTagUDT &&
			    PDB::IsUnnamedSymbol(e))
			{
				Expand = false;
			}
			if (Expand)
			{
				m_SymbolVisitor->Run(e);
			}
		}
		if (m_Settings.PrintPragmaPack)
		{
			*m_Settings.PdbHeaderReconstructorSettings.OutputFile
				<< DEFINITIONS_PRAGMA_PACK_END
				<< std::endl;
		}
	}
}*/
return true
}

func (p *pdbExtractor)PDBExtractor::PrintPDBFunctions()(ok bool){//col:419
/*PDBExtractor::PrintPDBFunctions()
{
	if (m_Settings.PrintFunctions)
	{
		*m_Settings.PdbHeaderReconstructorSettings.OutputFile
			<< std::endl;
		for (auto&& e : m_PDB.GetFunctionSet())
		{
			*m_Settings.PdbHeaderReconstructorSettings.OutputFile
				<< e
				<< std::endl;
		}
		*m_Settings.PdbHeaderReconstructorSettings.OutputFile
			<< std::endl;
	}
}*/
return true
}

func (p *pdbExtractor)PDBExtractor::DumpAllSymbols()(ok bool){//col:430
/*PDBExtractor::DumpAllSymbols()
{
	PrintPDBHeader();
	for (auto&& e : m_PDB.GetSymbolMap())
	{
		m_SymbolSorter->Visit(e.second);
	}
	PrintPDBDeclarations();
	PrintPDBDefinitions();
	PrintPDBFunctions();
}*/
return true
}

func (p *pdbExtractor)PDBExtractor::DumpOneSymbol()(ok bool){//col:449
/*PDBExtractor::DumpOneSymbol()
{
	const SYMBOL* Symbol = m_PDB.GetSymbolByName(m_Settings.SymbolName.c_str());
	if (Symbol == nullptr)
	{
		throw PDBDumperException(MESSAGE_SYMBOL_NOT_FOUND);
	}
	PrintPDBHeader();
	if (m_Settings.PrintReferencedTypes &&
	    m_Settings.PdbHeaderReconstructorSettings.MemberStructExpansion != PDBHeaderReconstructor::MemberStructExpansionType::InlineAll)
	{
		m_SymbolSorter->Visit(Symbol);
		PrintPDBDefinitions();
	}
	else
	{
		m_SymbolVisitor->Run(Symbol);
	}
}*/
return true
}

func (p *pdbExtractor)PDBExtractor::DumpAllSymbolsOneByOne()(ok bool){//col:485
/*PDBExtractor::DumpAllSymbolsOneByOne()
{
	for (auto&& e : m_PDB.GetSymbolMap())
	{
		m_SymbolSorter->Visit(e.second);
	}
	auto Symbols = m_SymbolSorter->GetSortedSymbols();
	m_SymbolSorter->Clear();
	std::string OutputDirectory = m_Settings.OutputFilename
		? m_Settings.OutputFilename
		: ".";
	if (OutputDirectory[OutputDirectory.length() - 1] != '\\')
	{
		OutputDirectory += '\\';
	}
	BOOL Result = CreateDirectory(OutputDirectory.c_str(), NULL);
	if (!Result && GetLastError() != ERROR_ALREADY_EXISTS)
	{
		throw PDBDumperException("Cannot create directory");
	}
	for (auto&& e : Symbols)
	{
		if (!PDB::IsUnnamedSymbol(e))
		{
			m_Settings.PdbHeaderReconstructorSettings.OutputFile = new std::ofstream(
				OutputDirectory + std::string(e->Name) + ".h",
				std::ios::out
			);
			m_Settings.SymbolName = e->Name;
			DumpOneSymbol();
			delete m_Settings.PdbHeaderReconstructorSettings.OutputFile;
			m_SymbolSorter->Clear();
		}
	}
	m_Settings.PdbHeaderReconstructorSettings.OutputFile = nullptr;
}*/
return true
}

func (p *pdbExtractor)PDBExtractor::CloseOpenFiles()(ok bool){//col:496
/*PDBExtractor::CloseOpenFiles()
{
	if (m_Settings.TestFilename)
	{
		delete m_Settings.PdbHeaderReconstructorSettings.TestFile;
	}
	if (m_Settings.OutputFilename)
	{
		delete m_Settings.PdbHeaderReconstructorSettings.OutputFile;
	}
}*/
return true
}



