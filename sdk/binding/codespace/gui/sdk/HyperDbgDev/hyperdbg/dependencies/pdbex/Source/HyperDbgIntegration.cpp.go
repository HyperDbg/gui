package Source

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\pdbex\Source\HyperDbgIntegration.cpp.back

type (
	HyperDbgIntegration interface {
		pdbex_export() (ok bool) //col:20
		ExtractBits() (ok bool)  //col:25
	}
	hyperDbgIntegration struct{}
)

func NewHyperDbgIntegration() HyperDbgIntegration { return &hyperDbgIntegration{} }

func (h *hyperDbgIntegration) pdbex_export() (ok bool) { //col:20
	/*
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

	   pdbex_set_logging_method_export(PVOID handler)

	   	{
	   	    g_MessageHandler = (Callback)handler;
	   	}
	*/
	return true
}

func (h *hyperDbgIntegration) ExtractBits() (ok bool) { //col:25
	/*
	   ExtractBits(UINT64 Orig64Bit, UINT64 From, UINT64 To)

	   	{
	   	    UINT64 Mask = ((1ull << (To - From + 1ull)) - 1ull) << From;
	   	    return (Orig64Bit & Mask) >> From;
	   	}
	*/
	return true
}
