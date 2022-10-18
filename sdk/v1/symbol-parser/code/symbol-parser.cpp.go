package code

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\symbol-parser\code\symbol-parser.cpp.back

type (
	SymbolParser interface {
		SymSetTextMessageCallback() (ok bool)    //col:60
		SymGetFieldOffsetFromModule() (ok bool)  //col:103
		SymGetDataTypeSizeFromModule() (ok bool) //col:119
		SymLoadFileSymbol() (ok bool)            //col:184
		SymUnloadModuleSymbol() (ok bool)        //col:253
		SymGetFieldOffset() (ok bool)            //col:344
		SymSeparateTo64BitValue() (ok bool)      //col:373
	}
	symbolParser struct{}
)

func NewSymbolParser() SymbolParser { return &symbolParser{} }

func (s *symbolParser) SymSetTextMessageCallback() (ok bool) { //col:60
	/*
	   SymSetTextMessageCallback(PVOID handler)

	   	{
	   	    g_MessageHandler = (Callback)handler;
	   	    pdbex_set_logging_method_export(handler);

	   ShowMessages(const char * Fmt, ...)

	   	{
	   	    va_list ArgList;
	   	    va_list Args;
	   	    if (g_MessageHandler == NULL)
	   	    {
	   	        va_start(Args, Fmt);
	   	        vprintf(Fmt, Args);
	   	        va_end(Args);
	   	        va_start(ArgList, Fmt);
	   	        int sprintfresult = vsprintf_s(TempMessage, Fmt, ArgList);
	   	        va_end(ArgList);
	   	        if (sprintfresult != -1)
	   	        {
	   	            g_MessageHandler(TempMessage);

	   SymGetModuleBaseFromSearchMask(const char * SearchMask, BOOLEAN SetModuleNameGlobally)

	   	{
	   	    string Token;
	   	    char   ModuleName[256] = {0};
	   	    int    Index                  = 0;
	   	    char   Ch                     = NULL;
	   	    if (!g_IsLoadedModulesInitialized || SearchMask == NULL)
	   	    {
	   	        return NULL;
	   	    }
	   	    string SearchMaskString(SearchMask);
	   	    if (SearchMaskString.find(Delimiter) != std::string::npos)
	   	    {
	   	        Token = SearchMaskString.substr(0, SearchMaskString.find(Delimiter));
	   	        strcpy(ModuleName, Token.c_str());
	   	        if (ModuleName[0] == '\0')
	   	        {
	   	            return NULL;
	   	        }
	   	        while (ModuleName[Index])
	   	        {
	   	            Ch = ModuleName[Index];
	   	            ModuleName[Index] = tolower(Ch);
	   	        if (strcmp(ModuleName, "ntkrnlmp") == 0 || strcmp(ModuleName, "ntoskrnl") == 0 ||
	   	            strcmp(ModuleName, "ntkrpamp") == 0 || strcmp(ModuleName, "ntkrnlpa") == 0)
	   	        {
	   	            RtlZeroMemory(ModuleName, _MAX_FNAME);
	   	        RtlZeroMemory(ModuleName, _MAX_FNAME);
	   	    for (auto item : g_LoadedModules)
	   	    {
	   	        if (strcmp((const char *)item->ModuleName, ModuleName) == 0)
	   	        {
	   	            if (SetModuleNameGlobally)
	   	            {
	   	                g_CurrentModuleName = (char *)item->ModuleName;
	   	            }
	   	            return item;
	   	        }
	   	    }
	   	    return NULL;
	   	}
	*/
	return true
}

func (s *symbolParser) SymGetFieldOffsetFromModule() (ok bool) { //col:103
	/*
	   SymGetFieldOffsetFromModule(UINT64 Base, WCHAR * TypeName, WCHAR * FieldName, UINT32 * FieldOffset)

	   	{
	   	    BOOLEAN Found = FALSE;
	   	    const DWORD SizeOfStruct =
	   	        sizeof(SYMBOL_INFOW) + ((MAX_SYM_NAME - 1) * sizeof(wchar_t));
	   	    auto    SymbolInfo = PSYMBOL_INFOW(SymbolInfoBuffer);
	   	    SymbolInfo->SizeOfStruct = sizeof(SYMBOL_INFOW);
	   	    if (!SymGetTypeFromNameW(GetCurrentProcess(), Base, TypeName, SymbolInfo))
	   	    {
	   	        return FALSE;
	   	    }
	   	    const ULONG TypeIndex     = SymbolInfo->TypeIndex;
	   	    DWORD       ChildrenCount = 0;
	   	    if (!SymGetTypeInfo(GetCurrentProcess(), Base, TypeIndex, TI_GET_CHILDRENCOUNT, &ChildrenCount))
	   	    {
	   	        return FALSE;
	   	    }
	   	    auto FindChildrenParamsBacking = std::make_unique<uint8_t[]>(
	   	        sizeof(_TI_FINDCHILDREN_PARAMS) + ((ChildrenCount - 1) * sizeof(ULONG)));
	   	        (_TI_FINDCHILDREN_PARAMS *)FindChildrenParamsBacking.get();
	   	    if (!SymGetTypeInfo(GetCurrentProcess(), Base, TypeIndex, TI_FINDCHILDREN, FindChildrenParams))
	   	    {
	   	        return FALSE;
	   	    }
	   	    for (DWORD ChildIdx = 0; ChildIdx < ChildrenCount; ChildIdx++)
	   	    {
	   	        const ULONG ChildId   = FindChildrenParams->ChildId[ChildIdx];
	   	        WCHAR *     ChildName = nullptr;
	   	        SymGetTypeInfo(GetCurrentProcess(), Base, ChildId, TI_GET_SYMNAME, &ChildName);
	   	        SymGetTypeInfo(GetCurrentProcess(), Base, ChildId, TI_GET_LENGTH, &ChildSize);
	   	        if (wcscmp(ChildName, FieldName) == 0)
	   	        {
	   	            const IMAGEHLP_SYMBOL_TYPE_INFO Info =
	   	                (ChildSize == 1) ? TI_GET_BITPOSITION : TI_GET_OFFSET;
	   	            SymGetTypeInfo(GetCurrentProcess(), Base, ChildId, Info, FieldOffset);
	   	        LocalFree(ChildName);
	   	        if (Found)
	   	        {
	   	            break;
	   	        }
	   	    }
	   	    return Found;
	   	}
	*/
	return true
}

func (s *symbolParser) SymGetDataTypeSizeFromModule() (ok bool) { //col:119
	/*
	   SymGetDataTypeSizeFromModule(UINT64 Base, WCHAR * TypeName, UINT64 * TypeSize)

	   	{
	   	    const DWORD SizeOfStruct =
	   	        sizeof(SYMBOL_INFOW) + ((MAX_SYM_NAME - 1) * sizeof(wchar_t));
	   	    auto    SymbolInfo = PSYMBOL_INFOW(SymbolInfoBuffer);
	   	    SymbolInfo->SizeOfStruct = sizeof(SYMBOL_INFOW);
	   	    if (!SymGetTypeFromNameW(GetCurrentProcess(), Base, TypeName, SymbolInfo))
	   	    {
	   	        return FALSE;
	   	    }
	   	    if (!SymGetTypeInfo(GetCurrentProcess(), Base, SymbolInfo->TypeIndex, TI_GET_LENGTH, TypeSize))
	   	    {
	   	        return FALSE;
	   	    }
	   	    return TRUE;
	   	}
	*/
	return true
}

func (s *symbolParser) SymLoadFileSymbol() (ok bool) { //col:184
	/*
	   SymLoadFileSymbol(UINT64 BaseAddress, const char * PdbFileName)

	   	{
	   	    BOOL                          Ret                             = FALSE;
	   	    DWORD                         Options                         = 0;
	   	    DWORD                         FileSize                        = 0;
	   	    int                           Index                           = 0;
	   	    char                          Ch                              = NULL;
	   	    char                          ModuleName[256]          = {0};
	   	    char                          AlternateModuleName[256] = {0};
	   	    PSYMBOL_LOADED_MODULE_DETAILS ModuleDetails                   = NULL;
	   	    Options = SymGetOptions();
	   	    SymSetOptions(Options);
	   	    Ret = SymInitialize(
	   	        GetCurrentProcess(),
	   	        NULL,
	   	        FALSE
	   	    );
	   	    if (!Ret)
	   	    {
	   	        ShowMessages("err, symbol init failed (%x)\n",
	   	                     GetLastError());
	   	    if (!SymGetFileParams(PdbFileName, FileSize))
	   	    {
	   	        ShowMessages("err, cannot obtain file parameters (internal error)\n");
	   	    _splitpath(PdbFileName, NULL, NULL, ModuleName, NULL);
	   	    strcpy(AlternateModuleName, ModuleName);
	   	    while (ModuleName[Index])
	   	    {
	   	        Ch = ModuleName[Index];
	   	        ModuleName[Index] = tolower(Ch);
	   	    if (strcmp(ModuleName, ("ntkrnlmp")) == 0 || strcmp(ModuleName, ("ntoskrnl")) == 0 ||
	   	        strcmp(ModuleName, ("ntkrpamp")) == 0 || strcmp(ModuleName, ("ntkrnlpa")) == 0)
	   	    {
	   	        RtlZeroMemory(ModuleName, _MAX_FNAME);
	   	        RtlZeroMemory(g_NtModuleName, _MAX_FNAME);
	   	        strcpy(g_NtModuleName, AlternateModuleName);
	   	    ModuleDetails = (SYMBOL_LOADED_MODULE_DETAILS *)malloc(sizeof(SYMBOL_LOADED_MODULE_DETAILS));
	   	    if (ModuleDetails == NULL)
	   	    {
	   	        ShowMessages("err, allocating buffer for storing symbol details (%x)\n",
	   	                     GetLastError());
	   	    RtlZeroMemory(ModuleDetails, sizeof(SYMBOL_LOADED_MODULE_DETAILS));
	   	    ModuleDetails->ModuleBase = SymLoadModule64(
	   	        GetCurrentProcess(),
	   	        NULL,
	   	        PdbFileName,
	   	        NULL,
	   	        BaseAddress,
	   	        FileSize
	   	    );
	   	    if (ModuleDetails->ModuleBase == NULL)
	   	    {
	   	        free(ModuleDetails);
	   	    ShowMessages("loading symbols for: %s\n", PdbFilePath);
	   	    ShowMessages("load address: %I64x\n", ModuleDetails.ModuleBase);
	   	    SymShowSymbolInfo(ModuleDetails.ModuleBase);
	   	    strcpy((char *)ModuleDetails->ModuleName, ModuleName);
	   	    strcpy((char *)ModuleDetails->PdbFilePath, PdbFileName);
	   	    g_LoadedModules.push_back(ModuleDetails);
	   	    if (!g_IsLoadedModulesInitialized)
	   	    {
	   	        g_IsLoadedModulesInitialized = TRUE;
	   	    }
	   	    return 0;
	   	}
	*/
	return true
}

func (s *symbolParser) SymUnloadModuleSymbol() (ok bool) { //col:253
	/*
	   SymUnloadModuleSymbol(char * ModuleName)

	   	{
	   	    BOOLEAN OneModuleFound = FALSE;
	   	    BOOL    Ret            = FALSE;
	   	    UINT32  Index          = 0;
	   	    for (auto item : g_LoadedModules)
	   	    {
	   	        Index++;
	   	        if (strcmp(item->ModuleName, ModuleName) == 0)
	   	        {
	   	            Ret = SymUnloadModule64(GetCurrentProcess(), item->ModuleBase);
	   	            if (!Ret)
	   	            {
	   	                ShowMessages("err, unload symbol failed (%x)\n",
	   	                             GetLastError());
	   	            free(item);
	   	    if (!OneModuleFound)
	   	    {
	   	        return -1;
	   	    }
	   	    std::vector<PSYMBOL_LOADED_MODULE_DETAILS>::iterator it = g_LoadedModules.begin();
	   	    std::advance(it, --Index);
	   	    g_LoadedModules.erase(it);

	   SymUnloadAllSymbols()

	   	{
	   	    BOOL Ret = FALSE;
	   	    for (auto item : g_LoadedModules)
	   	    {
	   	        Ret = SymUnloadModule64(GetCurrentProcess(), item->ModuleBase);
	   	        if (!Ret)
	   	        {
	   	            ShowMessages("err, unload symbol failed (%x)\n",
	   	                         GetLastError());
	   	        free(item);
	   	    g_LoadedModules.clear();
	   	    Ret = SymCleanup(GetCurrentProcess());
	   	    if (!Ret)
	   	    {
	   	        ShowMessages("err, symbol cleanup failed (%x)\n", GetLastError());

	   SymConvertNameToAddress(const char * FunctionOrVariableName, PBOOLEAN WasFound)

	   	{
	   	    BOOLEAN      Found   = FALSE;
	   	    UINT64       Address = NULL;
	   	    UINT64       Buffer[(sizeof(SYMBOL_INFO) + MAX_SYM_NAME * sizeof(CHAR) + sizeof(UINT64) - 1) / sizeof(UINT64)];
	   	    PSYMBOL_INFO Symbol = (PSYMBOL_INFO)Buffer;
	   	    string       name   = FunctionOrVariableName;
	   	    *WasFound = FALSE;
	   	    Symbol->SizeOfStruct = sizeof(SYMBOL_INFO);
	   	    if (strlen(FunctionOrVariableName) >= 4 &&
	   	        tolower(FunctionOrVariableName[0]) == 'n' &&
	   	        tolower(FunctionOrVariableName[1]) == 't' &&
	   	        tolower(FunctionOrVariableName[2]) == '!')
	   	    {
	   	        FunctionOrVariableName += 3;
	   	        name = "ntkrnlmp!";
	   	        name += FunctionOrVariableName;
	   	    }
	   	    if (SymFromName(GetCurrentProcess(), name.c_str(), Symbol))
	   	    {
	   	        Found   = TRUE;
	   	        Address = Symbol->Address;
	   	    }
	   	    else
	   	    {
	   	        Found = FALSE;
	   	    }
	   	    *WasFound = Found;
	   	    return Address;
	   	}
	*/
	return true
}

func (s *symbolParser) SymGetFieldOffset() (ok bool) { //col:344
	/*
	   SymGetFieldOffset(CHAR * TypeName, CHAR * FieldName, UINT32 * FieldOffset)

	   	{
	   	    BOOL                          Ret        = FALSE;
	   	    UINT32                        Index      = 0;
	   	    PSYMBOL_LOADED_MODULE_DETAILS SymbolInfo = NULL;
	   	    BOOLEAN                       Result     = FALSE;
	   	    SymbolInfo = SymGetModuleBaseFromSearchMask(TypeName, TRUE);
	   	    if (SymbolInfo == NULL)
	   	    {
	   	        return FALSE;
	   	    }
	   	    while (TypeName[Index] != '\0')
	   	    {
	   	        if (TypeName[Index] == '!')
	   	        {
	   	            TypeName = (CHAR *)(TypeName + Index + 1);
	   	    const size_t TypeNameSize = strlen(TypeName) + 1;
	   	    WCHAR *      TypeNameW    = (WCHAR *)malloc(sizeof(wchar_t) * TypeNameSize);
	   	    RtlZeroMemory(TypeNameW, sizeof(wchar_t) * TypeNameSize);
	   	    mbstowcs(TypeNameW, TypeName, TypeNameSize);
	   	    const size_t FieldNameSize = strlen(FieldName) + 1;
	   	    WCHAR *      FieldNameW    = (WCHAR *)malloc(sizeof(wchar_t) * FieldNameSize);
	   	    RtlZeroMemory(FieldNameW, sizeof(wchar_t) * FieldNameSize);
	   	    mbstowcs(FieldNameW, FieldName, FieldNameSize);
	   	    Result = SymGetFieldOffsetFromModule(SymbolInfo->ModuleBase, TypeNameW, FieldNameW, FieldOffset);
	   	    free(TypeNameW);
	   	    free(FieldNameW);

	   SymGetDataTypeSize(CHAR * TypeName, UINT64 * TypeSize)

	   	{
	   	    BOOL                          Ret        = FALSE;
	   	    UINT32                        Index      = 0;
	   	    PSYMBOL_LOADED_MODULE_DETAILS SymbolInfo = NULL;
	   	    BOOLEAN                       Result     = FALSE;
	   	    SymbolInfo = SymGetModuleBaseFromSearchMask(TypeName, TRUE);
	   	    if (SymbolInfo == NULL)
	   	    {
	   	        return FALSE;
	   	    }
	   	    while (TypeName[Index] != '\0')
	   	    {
	   	        if (TypeName[Index] == '!')
	   	        {
	   	            TypeName = (CHAR *)(TypeName + Index + 1);
	   	    const size_t TypeNameSize = strlen(TypeName) + 1;
	   	    WCHAR *      TypeNameW    = (WCHAR *)malloc(sizeof(wchar_t) * TypeNameSize);
	   	    RtlZeroMemory(TypeNameW, sizeof(wchar_t) * TypeNameSize);
	   	    mbstowcs(TypeNameW, TypeName, TypeNameSize);
	   	    Result = SymGetDataTypeSizeFromModule(SymbolInfo->ModuleBase, TypeNameW, TypeSize);
	   	    free(TypeNameW);

	   SymSearchSymbolForMask(const char * SearchMask)

	   	{
	   	    BOOL                          Ret        = FALSE;
	   	    PSYMBOL_LOADED_MODULE_DETAILS SymbolInfo = NULL;
	   	    SymbolInfo = SymGetModuleBaseFromSearchMask(SearchMask, TRUE);
	   	    if (SymbolInfo == NULL)
	   	    {
	   	        return -1;
	   	    }
	   	    Ret = SymEnumSymbols(
	   	        GetCurrentProcess(),
	   	        SymbolInfo->ModuleBase,
	   	        SearchMask,
	   	        SymDisplayMaskSymbolsCallback,
	   	        NULL
	   	    );
	   	    if (!Ret)
	   	    {
	   	        ShowMessages("err, symbol enum failed (%x)\n",
	   	                     GetLastError());

	   SymCreateSymbolTableForDisassembler(void * CallbackFunction)

	   	{
	   	    BOOL    Ret    = FALSE;
	   	    BOOLEAN Result = TRUE;
	   	    g_SymbolMapForDisassembler = (SymbolMapCallback)CallbackFunction;
	   	    for (auto item : g_LoadedModules)
	   	    {
	   	        g_CurrentModuleName = (char *)item->ModuleName;
	   	        Ret = SymEnumSymbols(
	   	            GetCurrentProcess(),
	   	            item->BaseAddress,
	   	            NULL,
	   	            SymDeliverDisassemblerSymbolMapCallback,
	   	            NULL
	   	        );
	   	        if (!Ret)
	   	        {
	   	            Result = FALSE;
	   	        }
	   	    }
	   	    return Result;
	   	}
	*/
	return true
}

func (s *symbolParser) SymSeparateTo64BitValue() (ok bool) { //col:373
	/*
	   SymSeparateTo64BitValue(UINT64 Value)

	   	{
	   	    ostringstream OstringStream;
	   	    string        Temp;
	   	    OstringStream << setw(16) << setfill('0') << hex << Value;
	   	    Temp = OstringStream.str();
	   	    Temp.insert(8, 1, '`');

	   SymGetFileParams(const char * FileName, DWORD & FileSize)

	   	{
	   	    if (FileName == 0)
	   	    {
	   	        return FALSE;
	   	    }
	   	    char FileExt[_MAX_EXT] = {0};
	   	    _splitpath(FileName, NULL, NULL, NULL, FileExt);
	   	    if (strcmp(FileExt, (".pdb")) == 0 || strcmp(FileExt, (".PDB")) == 0)
	   	    {
	   	        if (!SymGetFileSize(FileName, FileSize))
	   	        {
	   	            return FALSE;
	   	        }
	   	    }
	   	    else
	   	    {
	   	        FileSize = 0;
	   	        return FALSE;
	   	    }
	   	    return TRUE;
	   	}
	*/
	return true
}

