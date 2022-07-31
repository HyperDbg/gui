package Source
//back\HyperDbgDev\hyperdbg\dependencies\pdbex\Source\HyperDbgIntegration.cpp.back

type (
HyperDbgIntegration interface{
#pragma comment()(ok bool)//col:36
pdbex_set_logging_method_export()(ok bool)//col:42
ExtractBits()(ok bool)//col:49
SymSeparateTo64BitValue()(ok bool)//col:62
}
)

func NewHyperDbgIntegration() { return & hyperDbgIntegration{} }

func (h *hyperDbgIntegration)#pragma comment()(ok bool){//col:36
/*#pragma comment(lib, "dbghelp.lib")
int
pdbex_export(int     argc,
             char ** argv,
             bool    is_struct,
             void *  buffer_address)
{
    PDBExtractor Instance;
    if (!is_struct)
    {
        g_ShowInOffestFormat = TRUE;
    }
    else
    {
        g_ShowInOffestFormat = FALSE;
    }
    g_MappingBufferAddress = (CHAR *)buffer_address;
    return Instance.Run(argc, argv);
}*/
return true
}

func (h *hyperDbgIntegration)pdbex_set_logging_method_export()(ok bool){//col:42
/*pdbex_set_logging_method_export(PVOID handler)
{
    g_MessageHandler = (Callback)handler;
}*/
return true
}

func (h *hyperDbgIntegration)ExtractBits()(ok bool){//col:49
/*ExtractBits(UINT64 Orig64Bit, UINT64 From, UINT64 To)
{
    UINT64 Mask = ((1ull << (To - From + 1ull)) - 1ull) << From;
    return (Orig64Bit & Mask) >> From;
}*/
return true
}

func (h *hyperDbgIntegration)SymSeparateTo64BitValue()(ok bool){//col:62
/*SymSeparateTo64BitValue(UINT64 Value)
{
    std::ostringstream OstringStream;
    std::string        Temp;
    OstringStream << std::setw(16) << std::setfill('0') << std::hex << Value;
    Temp = OstringStream.str();
    Temp.insert(8, 1, '`');
    return Temp;
}*/
return true
}


