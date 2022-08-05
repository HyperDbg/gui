package script_engine_wrapper

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\debugger\script-engine-wrapper\symbol.cpp.back

type (
	Symbol interface {
		SymbolInitialReload() (ok bool) //col:136
	}
	symbol struct{}
)

func NewSymbol() Symbol { return &symbol{} }

func (s *symbol) SymbolInitialReload() (ok bool) { //col:136
	/*
	   SymbolInitialReload()

	   	{
	   	    ShowMessages("interpreting symbols and creating symbol maps\n");
	   	    SymbolLoadOrDownloadSymbols(FALSE, TRUE);

	   SymbolLocalReload(UINT32 UserProcessId)

	   	{
	   	    ShowMessages("interpreting symbols and creating symbol maps\n");
	   	    SymbolBuildSymbolTable(&g_SymbolTable, &g_SymbolTableSize, UserProcessId, FALSE);
	   	    return SymbolLoadOrDownloadSymbols(FALSE, TRUE);

	   SymbolPrepareDebuggerWithSymbolInfo(UINT32 UserProcessId)

	   	{
	   	    SymbolBuildSymbolTable(&g_SymbolTable, &g_SymbolTableSize, UserProcessId, TRUE);

	   SymbolCreateDisassemblerMapCallback(UINT64       Address,

	   	char *       ModuleName,
	   	char *       ObjectName,
	   	unsigned int ObjectSize)

	   	{
	   	    LOCAL_FUNCTION_DESCRIPTION LocalFunctionDescription = {};
	   	    string                     FinalModuleName          = "";
	   	    if (ObjectSize == 0)
	   	    {
	   	        ObjectSize = DISASSEMBLY_MAXIMUM_DISTANCE_FROM_OBJECT_NAME;
	   	    }
	   	    if (ModuleName != NULL)
	   	    {
	   	        FinalModuleName += std::string(ModuleName) + "!";
	   	    }
	   	    if (ObjectName != NULL)
	   	    {
	   	        FinalModuleName += std::string(ObjectName);
	   	    LocalFunctionDescription.ObjectName = std::move(FinalModuleName);

	   SymbolCreateDisassemblerSymbolMap()

	   	{
	   	    g_DisassemblerSymbolMap.clear();
	   	    ScriptEngineCreateSymbolTableForDisassemblerWrapper(SymbolCreateDisassemblerMapCallback);

	   SymbolShowFunctionNameBasedOnAddress(UINT64 Address, PUINT64 UsedBaseAddress)

	   	{
	   	    std::map<UINT64, LOCAL_FUNCTION_DESCRIPTION>::iterator Low, Prev;
	   	    UINT64                                                 Pos = Address;
	   	    if (!g_AddressConversion)
	   	    {
	   	        return FALSE;
	   	    }
	   	    if (!g_DisassemblerSymbolMap.empty())
	   	    {
	   	        Low = g_DisassemblerSymbolMap.lower_bound(Pos);
	   	        if (Low == g_DisassemblerSymbolMap.end())
	   	        {
	   	            return FALSE;
	   	        }
	   	        else if (Low == g_DisassemblerSymbolMap.begin() && Low->first > Address)
	   	        {
	   	            return FALSE;
	   	        }
	   	        else if (Low->first == Address)
	   	        {
	   	            if (*UsedBaseAddress != Address)
	   	            {
	   	                ShowMessages("%s", Low->second.ObjectName.c_str());
	   	            Prev        = std::prev(Low);
	   	            if (Prev->second.ObjectSize >= Diff)
	   	            {
	   	                if (*UsedBaseAddress != Prev->first)
	   	                {
	   	                    ShowMessages("%s+0x%x", Prev->second.ObjectName.c_str(), Diff);
	   	            else if (DISASSEMBLY_MAXIMUM_DISTANCE_FROM_OBJECT_NAME >= Diff)
	   	            {
	   	                if (*UsedBaseAddress != Prev->first)
	   	                {
	   	                    ShowMessages("%s+0x%x+0x%x", Prev->second.ObjectName.c_str(), Diff, Diff - Prev->second.ObjectSize);

	   SymbolBuildAndShowSymbolTable()

	   	{
	   	    if (g_SymbolTable == NULL || g_SymbolTableSize == NULL)
	   	    {
	   	        ShowMessages("err, symbol table is empty. please use '.sym reload' "
	   	                     "to build the symbol table\n");
	   	    for (size_t i = 0; i < g_SymbolTableSize / sizeof(MODULE_SYMBOL_DETAIL); i++)
	   	        ShowMessages("is pdb details available? : %s\n", g_SymbolTable[i].IsSymbolDetailsFound ? "true" : "false");
	   	        ShowMessages("is pdb a path instead of module name? : %s\n", g_SymbolTable[i].IsLocalSymbolPath ? "true" : "false");
	   	        ShowMessages("base address : %llx\n", g_SymbolTable[i].BaseAddress);
	   	        ShowMessages("file path : %s\n", g_SymbolTable[i].FilePath);
	   	        ShowMessages("guid and age : %s\n", g_SymbolTable[i].ModuleSymbolGuidAndAge);
	   	        ShowMessages("module symbol path/name : %s\n", g_SymbolTable[i].ModuleSymbolPath);
	   	        ShowMessages("========================================================================\n");

	   SymbolLoadOrDownloadSymbols(BOOLEAN IsDownload, BOOLEAN SilentLoad)

	   	{
	   	    string  SymbolServer;
	   	    BOOLEAN Result = FALSE;
	   	    if (!CommandSettingsGetValueFromConfigFile("SymbolServer", SymbolServer))
	   	    {
	   	        ShowMessages("please configure the symbol path (use '.help .sympath' for more information)\n");
	   	    if (g_SymbolTable == NULL || g_SymbolTableSize == NULL)
	   	    {
	   	        ShowMessages("symbol table is empty, please use '.sym reload' to build a symbol table\n");
	   	    Result                             = ScriptEngineSymbolInitLoadWrapper(g_SymbolTable,
	   	                                               g_SymbolTableSize,
	   	                                               IsDownload,
	   	                                               SymbolServer.c_str(),
	   	                                               SilentLoad);
	   	    SymbolCreateDisassemblerSymbolMap();

	   SymbolConvertNameOrExprToAddress(const string & TextToConvert, PUINT64 Result)

	   	{
	   	    BOOLEAN IsFound  = FALSE;
	   	    BOOLEAN HasError = NULL;
	   	    UINT64  Address  = NULL;
	   	    if (!ConvertStringToUInt64(TextToConvert, &Address))
	   	    {
	   	        string ConstTextToConvert = TextToConvert;
	   	        Address                   = ScriptEngineConvertNameToAddressWrapper(ConstTextToConvert.c_str(), &IsFound);
	   	        if (!IsFound)
	   	        {
	   	            Address = ScriptEngineEvalSingleExpression(TextToConvert, &HasError);
	   	            if (HasError)
	   	            {
	   	                IsFound = FALSE;
	   	            }
	   	            else
	   	            {
	   	                IsFound = TRUE;
	   	            }
	   	        }
	   	        else
	   	        {
	   	            IsFound = TRUE;
	   	        }
	   	    }
	   	    else
	   	    {
	   	        IsFound = TRUE;
	   	    }
	   	    if (IsFound)
	   	    {
	   	        *Result = Address;
	   	    }
	   	    return IsFound;
	   	}
	*/
	return true
}

