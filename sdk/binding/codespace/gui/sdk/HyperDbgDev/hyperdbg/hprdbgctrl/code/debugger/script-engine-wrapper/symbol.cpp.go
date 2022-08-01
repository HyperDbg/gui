package script-engine-wrapper
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\debugger\script-engine-wrapper\symbol.cpp.back

type (
Symbol interface{
SymbolInitialReload()(ok bool)//col:5
SymbolLocalReload()(ok bool)//col:11
SymbolPrepareDebuggerWithSymbolInfo()(ok bool)//col:15
SymbolCreateDisassemblerMapCallback()(ok bool)//col:38
SymbolCreateDisassemblerSymbolMap()(ok bool)//col:44
SymbolShowFunctionNameBasedOnAddress()(ok bool)//col:101
SymbolBuildAndShowSymbolTable()(ok bool)//col:120
SymbolLoadOrDownloadSymbols()(ok bool)//col:144
SymbolConvertNameOrExprToAddress()(ok bool)//col:180
SymbolDeleteSymTable()(ok bool)//col:195
SymbolBuildSymbolTable()(ok bool)//col:433
SymbolBuildAndUpdateSymbolTable()(ok bool)//col:457
SymbolReloadSymbolTableInDebuggerMode()(ok bool)//col:470
}
symbol struct{}
)

func NewSymbol()Symbol{ return & symbol{} }

func (s *symbol)SymbolInitialReload()(ok bool){//col:5
/*SymbolInitialReload()
{
    ShowMessages("interpreting symbols and creating symbol maps\n");
    SymbolLoadOrDownloadSymbols(FALSE, TRUE);
}*/
return true
}

func (s *symbol)SymbolLocalReload()(ok bool){//col:11
/*SymbolLocalReload(UINT32 UserProcessId)
{
    ShowMessages("interpreting symbols and creating symbol maps\n");
    SymbolBuildSymbolTable(&g_SymbolTable, &g_SymbolTableSize, UserProcessId, FALSE);
    return SymbolLoadOrDownloadSymbols(FALSE, TRUE);
}*/
return true
}

func (s *symbol)SymbolPrepareDebuggerWithSymbolInfo()(ok bool){//col:15
/*SymbolPrepareDebuggerWithSymbolInfo(UINT32 UserProcessId)
{
    SymbolBuildSymbolTable(&g_SymbolTable, &g_SymbolTableSize, UserProcessId, TRUE);
}*/
return true
}

func (s *symbol)SymbolCreateDisassemblerMapCallback()(ok bool){//col:38
/*SymbolCreateDisassemblerMapCallback(UINT64       Address,
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
    }
    LocalFunctionDescription.ObjectName = std::move(FinalModuleName);
    LocalFunctionDescription.ObjectSize = ObjectSize;
    g_DisassemblerSymbolMap[Address] = LocalFunctionDescription;
}*/
return true
}

func (s *symbol)SymbolCreateDisassemblerSymbolMap()(ok bool){//col:44
/*SymbolCreateDisassemblerSymbolMap()
{
    g_DisassemblerSymbolMap.clear();
    ScriptEngineCreateSymbolTableForDisassemblerWrapper(SymbolCreateDisassemblerMapCallback);
    return TRUE;
}*/
return true
}

func (s *symbol)SymbolShowFunctionNameBasedOnAddress()(ok bool){//col:101
/*SymbolShowFunctionNameBasedOnAddress(UINT64 Address, PUINT64 UsedBaseAddress)
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
                *UsedBaseAddress = Address;
                return TRUE;
            }
            return FALSE;
        }
        else
        {
            Prev        = std::prev(Low);
            UINT64 Diff = Address - Prev->first;
            if (Prev->second.ObjectSize >= Diff)
            {
                if (*UsedBaseAddress != Prev->first)
                {
                    ShowMessages("%s+0x%x", Prev->second.ObjectName.c_str(), Diff);
                    *UsedBaseAddress = Prev->first;
                    return TRUE;
                }
                return FALSE;
            }
            else if (DISASSEMBLY_MAXIMUM_DISTANCE_FROM_OBJECT_NAME >= Diff)
            {
                if (*UsedBaseAddress != Prev->first)
                {
                    ShowMessages("%s+0x%x+0x%x", Prev->second.ObjectName.c_str(), Diff, Diff - Prev->second.ObjectSize);
                    *UsedBaseAddress = Prev->first;
                    return TRUE;
                }
                return FALSE;
            }
        }
    }
    return FALSE;
}*/
return true
}

func (s *symbol)SymbolBuildAndShowSymbolTable()(ok bool){//col:120
/*SymbolBuildAndShowSymbolTable()
{
    if (g_SymbolTable == NULL || g_SymbolTableSize == NULL)
    {
        ShowMessages("err, symbol table is empty. please use '.sym reload' "
                     "to build the symbol table\n");
        return;
    }
    for (size_t i = 0; i < g_SymbolTableSize / sizeof(MODULE_SYMBOL_DETAIL); i++)
    {
        ShowMessages("is pdb details available? : %s\n", g_SymbolTable[i].IsSymbolDetailsFound ? "true" : "false");
        ShowMessages("is pdb a path instead of module name? : %s\n", g_SymbolTable[i].IsLocalSymbolPath ? "true" : "false");
        ShowMessages("base address : %llx\n", g_SymbolTable[i].BaseAddress);
        ShowMessages("file path : %s\n", g_SymbolTable[i].FilePath);
        ShowMessages("guid and age : %s\n", g_SymbolTable[i].ModuleSymbolGuidAndAge);
        ShowMessages("module symbol path/name : %s\n", g_SymbolTable[i].ModuleSymbolPath);
        ShowMessages("========================================================================\n");
    }
}*/
return true
}

func (s *symbol)SymbolLoadOrDownloadSymbols()(ok bool){//col:144
/*SymbolLoadOrDownloadSymbols(BOOLEAN IsDownload, BOOLEAN SilentLoad)
{
    string  SymbolServer;
    BOOLEAN Result = FALSE;
    if (!CommandSettingsGetValueFromConfigFile("SymbolServer", SymbolServer))
    {
        ShowMessages("please configure the symbol path (use '.help .sympath' for more information)\n");
        return FALSE;
    }
    if (g_SymbolTable == NULL || g_SymbolTableSize == NULL)
    {
        ShowMessages("symbol table is empty, please use '.sym reload' to build a symbol table\n");
        return FALSE;
    }
    g_IsExecutingSymbolLoadingRoutines = TRUE;
    Result                             = ScriptEngineSymbolInitLoadWrapper(g_SymbolTable,
                                               g_SymbolTableSize,
                                               IsDownload,
                                               SymbolServer.c_str(),
                                               SilentLoad);
    SymbolCreateDisassemblerSymbolMap();
    g_IsExecutingSymbolLoadingRoutines = FALSE;
    return Result;
}*/
return true
}

func (s *symbol)SymbolConvertNameOrExprToAddress()(ok bool){//col:180
/*SymbolConvertNameOrExprToAddress(const string & TextToConvert, PUINT64 Result)
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
}*/
return true
}

func (s *symbol)SymbolDeleteSymTable()(ok bool){//col:195
/*SymbolDeleteSymTable()
{
    if (g_SymbolTable != NULL)
    {
        free(g_SymbolTable);
        g_SymbolTable             = NULL;
        g_SymbolTableSize         = NULL;
        g_SymbolTableCurrentIndex = 0;
        return TRUE;
    }
    else
    {
        return FALSE;
    }
}*/
return true
}

func (s *symbol)SymbolBuildSymbolTable()(ok bool){//col:433
/*SymbolBuildSymbolTable(PMODULE_SYMBOL_DETAIL * BufferToStoreDetails,
                       PUINT32                 StoredLength,
                       UINT32                  UserProcessId,
                       BOOLEAN                 SendOverSerial)
{
    PRTL_PROCESS_MODULES            ModuleInfo;
    NTSTATUS                        NtStatus;
    BOOLEAN                         Status;
    ULONG                           ReturnedLength;
    PMODULE_SYMBOL_DETAIL           ModuleSymDetailArray                              = NULL;
    char                            SystemRoot[MAX_PATH]                              = {0};
    char                            ModuleSymbolPath[MAX_PATH]                        = {0};
    char                            TempPath[MAX_PATH]                                = {0};
    char                            ModuleSymbolGuidAndAge[MAXIMUM_GUID_AND_AGE_SIZE] = {0};
    BOOLEAN                         IsSymbolPdbDetailAvailable                        = FALSE;
    BOOLEAN                         IsFreeUsermodeModulesBuffer                       = FALSE;
    ULONG                           SysModuleInfoBufferSize                           = 0;
    UINT32                          ModuleDetailsSize                                 = 0;
    UINT32                          ModulesCount                                      = 0;
    PUSERMODE_LOADED_MODULE_DETAILS ModuleDetailsRequest                              = NULL;
    PUSERMODE_LOADED_MODULE_SYMBOLS Modules                                           = NULL;
    USERMODE_LOADED_MODULE_DETAILS  ModuleCountRequest                                = {0};
    SymbolDeleteSymTable();
    if (GetSystemDirectoryA(SystemRoot, MAX_PATH) == NULL)
    {
        ShowMessages("err, unable to get system directory (%x)\n",
                     GetLastError());
        return FALSE;
    }
    string SystemRootString(SystemRoot);
    transform(SystemRootString.begin(),
              SystemRootString.end(),
              SystemRootString.begin(),
              [](unsigned char c) { return std::tolower(c); });
    Replace(SystemRootString, "\\system32", "");
    NtStatus = NtQuerySystemInformation(SystemModuleInformation, NULL, NULL, &SysModuleInfoBufferSize);
    ModuleInfo = (PRTL_PROCESS_MODULES)VirtualAlloc(
        NULL,
        SysModuleInfoBufferSize,
        MEM_COMMIT | MEM_RESERVE,
        PAGE_READWRITE);
    if (!ModuleInfo)
    {
        ShowMessages("err, unable to allocate memory for module list (%x)\n",
                     GetLastError());
        return FALSE;
    }
    if (!NT_SUCCESS(
            NtStatus = NtQuerySystemInformation(SystemModuleInformation,
                                                ModuleInfo,
                                                SysModuleInfoBufferSize,
                                                NULL)))
    {
        ShowMessages("err, unable to query module list (%#x)\n", NtStatus);
        VirtualFree(ModuleInfo, 0, MEM_RELEASE);
        return FALSE;
    }
    do
    {
        if (!g_DeviceHandle)
        {
            break;
        }
        ModuleCountRequest.ProcessId        = UserProcessId;
        ModuleCountRequest.OnlyCountModules = TRUE;
        Status = DeviceIoControl(
            g_DeviceHandle,                         
            IOCTL_GET_USER_MODE_MODULE_DETAILS,     
            &ModuleCountRequest,                    
            sizeof(USERMODE_LOADED_MODULE_DETAILS), 
            &ModuleCountRequest,                    
            sizeof(USERMODE_LOADED_MODULE_DETAILS), 
            &ReturnedLength,                        
            NULL                                    
        );
        if (!Status)
        {
            break;
        }
        if (ModuleCountRequest.Result == DEBUGGER_OPERATION_WAS_SUCCESSFULL)
        {
            ModulesCount = ModuleCountRequest.ModulesCount;
            ModuleDetailsSize = sizeof(USERMODE_LOADED_MODULE_DETAILS) +
                                (ModuleCountRequest.ModulesCount * sizeof(USERMODE_LOADED_MODULE_SYMBOLS));
            ModuleDetailsRequest = (PUSERMODE_LOADED_MODULE_DETAILS)malloc(ModuleDetailsSize);
            if (ModuleDetailsRequest == NULL)
            {
                break;
            }
            RtlZeroMemory(ModuleDetailsRequest, ModuleDetailsSize);
            ModuleDetailsRequest->ProcessId        = UserProcessId;
            ModuleDetailsRequest->OnlyCountModules = FALSE;
            Status = DeviceIoControl(
                g_DeviceHandle,                         
                IOCTL_GET_USER_MODE_MODULE_DETAILS,     
                ModuleDetailsRequest,                   
                sizeof(USERMODE_LOADED_MODULE_DETAILS), 
                ModuleDetailsRequest,                   
                ModuleDetailsSize,                      
                &ReturnedLength,                        
                NULL                                    
            );
            if (!Status)
            {
                free(ModuleDetailsRequest);
                break;
            }
            if (ModuleCountRequest.Result == DEBUGGER_OPERATION_WAS_SUCCESSFULL)
            {
                Modules = (PUSERMODE_LOADED_MODULE_SYMBOLS)((UINT64)ModuleDetailsRequest +
                                                            sizeof(USERMODE_LOADED_MODULE_DETAILS));
                if (Modules == NULL)
                {
                    ModulesCount = 0;
                    free(ModuleDetailsRequest);
                    break;
                }
            }
            else
            {
                ShowErrorMessage(ModuleCountRequest.Result);
                free(ModuleDetailsRequest);
                break;
            }
            IsFreeUsermodeModulesBuffer = TRUE;
        }
        else
        {
            ShowErrorMessage(ModuleCountRequest.Result);
            break;
        }
    } while (FALSE);
    ModuleSymDetailArray = (PMODULE_SYMBOL_DETAIL)malloc(
        (ModuleInfo->NumberOfModules + ModulesCount) * sizeof(MODULE_SYMBOL_DETAIL));
    if (ModuleSymDetailArray == NULL)
    {
        ShowMessages("err, unable to allocate memory for module list (%x)\n",
                     GetLastError());
        if (IsFreeUsermodeModulesBuffer)
        {
            free(ModuleDetailsRequest);
        }
        return FALSE;
    }
    RtlZeroMemory(ModuleSymDetailArray, (ModuleInfo->NumberOfModules + ModulesCount) * sizeof(MODULE_SYMBOL_DETAIL));
    if (ModulesCount != 0)
    {
        for (int i = 0; i < ModulesCount; i++)
        {
            RtlZeroMemory(ModuleSymbolPath, sizeof(ModuleSymbolPath));
            RtlZeroMemory(TempPath, sizeof(TempPath));
            RtlZeroMemory(ModuleSymbolGuidAndAge, sizeof(ModuleSymbolGuidAndAge));
            wcstombs(TempPath, Modules[i].FilePath, MAX_PATH);
            if (ScriptEngineConvertFileToPdbFileAndGuidAndAgeDetailsWrapper(
                    TempPath,
                    ModuleSymbolPath,
                    ModuleSymbolGuidAndAge))
            {
                IsSymbolPdbDetailAvailable = TRUE;
            }
            else
            {
                IsSymbolPdbDetailAvailable = FALSE;
            }
            ModuleSymDetailArray[i].BaseAddress = Modules[i].BaseAddress;
            ModuleSymDetailArray[i].IsUserMode  = TRUE;
            memcpy(ModuleSymDetailArray[i].FilePath, TempPath, strlen(TempPath));
            if (IsSymbolPdbDetailAvailable)
            {
                ModuleSymDetailArray[i].IsSymbolDetailsFound = TRUE;
                memcpy(ModuleSymDetailArray[i].ModuleSymbolGuidAndAge, ModuleSymbolGuidAndAge, MAXIMUM_GUID_AND_AGE_SIZE);
                memcpy(ModuleSymDetailArray[i].ModuleSymbolPath, ModuleSymbolPath, MAX_PATH);
                string ModuleSymbolPathString(ModuleSymbolPath);
                if (ModuleSymbolPathString.find(":\\") != std::string::npos)
                    ModuleSymDetailArray[i].IsLocalSymbolPath = TRUE;
                else
                    ModuleSymDetailArray[i].IsLocalSymbolPath = FALSE;
            }
            else
            {
                ModuleSymDetailArray[i].IsSymbolDetailsFound = FALSE;
            }
            if (SendOverSerial)
            {
                KdSendSymbolDetailPacket(&ModuleSymDetailArray[i], i, ModuleInfo->NumberOfModules + ModulesCount);
            }
        }
    }
    for (int i = 0; i < ModuleInfo->NumberOfModules; i++)
    {
        UINT32 IndexInSymbolBuffer = ModulesCount + i;
        auto PathName = ModuleInfo->Modules[i].FullPathName + ModuleInfo->Modules[i].OffsetToFileName;
        RtlZeroMemory(ModuleSymbolPath, sizeof(ModuleSymbolPath));
        RtlZeroMemory(ModuleSymbolGuidAndAge, sizeof(ModuleSymbolGuidAndAge));
        string ModuleFullPath((const char *)ModuleInfo->Modules[i].FullPathName);
        if (ModuleFullPath.rfind("\\SystemRoot\\", 0) == 0)
        {
            Replace(ModuleFullPath, "\\SystemRoot", SystemRootString);
        }
        if (ScriptEngineConvertFileToPdbFileAndGuidAndAgeDetailsWrapper(ModuleFullPath.c_str(), ModuleSymbolPath, ModuleSymbolGuidAndAge))
        {
            IsSymbolPdbDetailAvailable = TRUE;
        }
        else
        {
            IsSymbolPdbDetailAvailable = FALSE;
        }
        ModuleSymDetailArray[IndexInSymbolBuffer].BaseAddress = (UINT64)ModuleInfo->Modules[i].ImageBase;
        memcpy(ModuleSymDetailArray[IndexInSymbolBuffer].FilePath, ModuleFullPath.c_str(), ModuleFullPath.size());
        if (IsSymbolPdbDetailAvailable)
        {
            ModuleSymDetailArray[IndexInSymbolBuffer].IsSymbolDetailsFound = TRUE;
            memcpy(ModuleSymDetailArray[IndexInSymbolBuffer].ModuleSymbolGuidAndAge, ModuleSymbolGuidAndAge, MAXIMUM_GUID_AND_AGE_SIZE);
            memcpy(ModuleSymDetailArray[IndexInSymbolBuffer].ModuleSymbolPath, ModuleSymbolPath, MAX_PATH);
            string ModuleSymbolPathString(ModuleSymbolPath);
            if (ModuleSymbolPathString.find(":\\") != std::string::npos)
                ModuleSymDetailArray[IndexInSymbolBuffer].IsLocalSymbolPath = TRUE;
            else
                ModuleSymDetailArray[IndexInSymbolBuffer].IsLocalSymbolPath = FALSE;
        }
        else
        {
            ModuleSymDetailArray[IndexInSymbolBuffer].IsSymbolDetailsFound = FALSE;
        }
        if (SendOverSerial)
        {
            KdSendSymbolDetailPacket(&ModuleSymDetailArray[IndexInSymbolBuffer], i, ModuleInfo->NumberOfModules + ModulesCount);
        }
    }
    *BufferToStoreDetails = ModuleSymDetailArray;
    *StoredLength         = (ModuleInfo->NumberOfModules + ModulesCount) * sizeof(MODULE_SYMBOL_DETAIL);
    VirtualFree(ModuleInfo, 0, MEM_RELEASE);
    if (IsFreeUsermodeModulesBuffer)
    {
        free(ModuleDetailsRequest);
    }
    return TRUE;
}*/
return true
}

func (s *symbol)SymbolBuildAndUpdateSymbolTable()(ok bool){//col:457
/*SymbolBuildAndUpdateSymbolTable(PMODULE_SYMBOL_DETAIL SymbolDetail)
{
    if (g_SymbolTableCurrentIndex >= MAXIMUM_SUPPORTED_SYMBOLS)
    {
        ShowMessages("err, the symbol table buffer is full, unable to add new symbol\n");
        return FALSE;
    }
    if (g_SymbolTable == NULL)
    {
        g_SymbolTable = (PMODULE_SYMBOL_DETAIL)malloc(MAXIMUM_SUPPORTED_SYMBOLS * sizeof(MODULE_SYMBOL_DETAIL));
        if (g_SymbolTable == NULL)
        {
            ShowMessages("err, unable to allocate memory for module list (%x)\n",
                         GetLastError());
            return FALSE;
        }
        g_SymbolTableCurrentIndex = 0;
        RtlZeroMemory(g_SymbolTable, MAXIMUM_SUPPORTED_SYMBOLS * sizeof(MODULE_SYMBOL_DETAIL));
    }
    memcpy(&g_SymbolTable[g_SymbolTableCurrentIndex], SymbolDetail, sizeof(MODULE_SYMBOL_DETAIL));
    g_SymbolTableCurrentIndex++;
    g_SymbolTableSize = g_SymbolTableCurrentIndex * sizeof(MODULE_SYMBOL_DETAIL);
    return TRUE;
}*/
return true
}

func (s *symbol)SymbolReloadSymbolTableInDebuggerMode()(ok bool){//col:470
/*SymbolReloadSymbolTableInDebuggerMode(UINT32 ProcessId)
{
    SymbolDeleteSymTable();
    if (KdSendSymbolReloadPacketToDebuggee(ProcessId))
    {
        ShowMessages("symbol table updated successfully\n");
        return TRUE;
    }
    else
    {
        return FALSE;
    }
}*/
return true
}



