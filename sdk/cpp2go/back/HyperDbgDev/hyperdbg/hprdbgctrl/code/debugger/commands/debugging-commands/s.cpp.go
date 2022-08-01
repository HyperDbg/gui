package debugging-commands
//back\HyperDbgDev\hyperdbg\hprdbgctrl\code\debugger\commands\debugging-commands\s.cpp.back

type (
S interface{
CommandSearchMemoryHelp()(ok bool)//col:24
CommandSearchSendRequest()(ok bool)//col:63
CommandSearchMemory()(ok bool)//col:307
}
)

func NewS() { return & s{} }

func (s *s)CommandSearchMemoryHelp()(ok bool){//col:24
/*CommandSearchMemoryHelp()
{
    ShowMessages("sb !sb sd !sd sq !sq : searches a contiguous memory for a "
                 "special byte pattern\n");
    ShowMessages("sb  Byte and ASCII characters\n");
    ShowMessages("sd  Double-word values (4 bytes)\n");
    ShowMessages("sq  Quad-word values (8 bytes). \n");
    ShowMessages(
        "\n If you want to search in physical (address) memory then add '!' "
        "at the start of the command\n");
    ShowMessages("syntax : \tsb [StartAddress (hex)] [l Length (hex)] [BytePattern (hex)] [pid ProcessId (hex)]\n");
    ShowMessages("syntax : \tsd [StartAddress (hex)] [l Length (hex)] [BytePattern (hex)] [pid ProcessId (hex)]\n");
    ShowMessages("syntax : \tsq [StartAddress (hex)] [l Length (hex)] [BytePattern (hex)] [pid ProcessId (hex)]\n");
    ShowMessages("\n");
    ShowMessages("\t\te.g : sb nt!ExAllocatePoolWithTag 90 85 95 l ffff \n");
    ShowMessages("\t\te.g : sb nt!ExAllocatePoolWithTag+5 90 85 95 l ffff \n");
    ShowMessages("\t\te.g : sb fffff8077356f010 90 85 95 l ffff \n");
    ShowMessages("\t\te.g : sd fffff8077356f010 90423580 l ffff pid 1c0 \n");
    ShowMessages("\t\te.g : !sq 100000 9090909090909090 l ffff\n");
    ShowMessages("\t\te.g : !sq 100000 9090909090909090 9090909090909090 "
                 "9090909090909090 l ffffff\n");
}*/
return true
}

func (s *s)CommandSearchSendRequest()(ok bool){//col:63
/*CommandSearchSendRequest(UINT64 * BufferToSendAsIoctl, UINT32 BufferToSendAsIoctlSize)
{
    BOOL    Status;
    UINT64  CurrentValue;
    PUINT64 ResultsBuffer = NULL;
    ResultsBuffer = (PUINT64)malloc(MaximumSearchResults * sizeof(UINT64));
    ZeroMemory(ResultsBuffer, MaximumSearchResults * sizeof(UINT64));
    Status =
        DeviceIoControl(g_DeviceHandle,               
                        IOCTL_DEBUGGER_SEARCH_MEMORY, 
                        BufferToSendAsIoctl,          
                        BufferToSendAsIoctlSize,      
                        ResultsBuffer,                
                        MaximumSearchResults *
                            sizeof(UINT64), 
                        NULL,               
                        NULL                
        );
    if (!Status)
    {
        ShowMessages("ioctl failed with code 0x%x\n", GetLastError());
        free(ResultsBuffer);
        return;
    }
    for (size_t i = 0; i < MaximumSearchResults; i++)
    {
        CurrentValue = ResultsBuffer[i];
        if (CurrentValue == NULL)
        {
            if (i == 0)
            {
                ShowMessages("not found\n");
            }
            break;
        }
        ShowMessages("%llx\n", CurrentValue);
    }
    free(ResultsBuffer);
}*/
return true
}

func (s *s)CommandSearchMemory()(ok bool){//col:307
/*CommandSearchMemory(vector<string> SplittedCommand, string Command)
{
    BOOL                   SetAddress          = FALSE;
    BOOL                   SetValue            = FALSE;
    BOOL                   SetProcId           = FALSE;
    BOOL                   NextIsProcId        = FALSE;
    BOOL                   SetLength           = FALSE;
    BOOL                   NextIsLength        = FALSE;
    DEBUGGER_SEARCH_MEMORY SearchMemoryRequest = {0};
    UINT64                 Address;
    UINT64                 Value         = 0;
    UINT64                 Length        = 0;
    UINT32                 ProcId        = 0;
    UINT32                 CountOfValues = 0;
    UINT32                 FinalSize     = 0;
    UINT64 *               FinalBuffer   = NULL;
    vector<UINT64>         ValuesToEdit;
    vector<string>         SplittedCommandCaseSensitive {Split(Command, ' ')};
    UINT32                 IndexInCommandCaseSensitive = 0;
    if (g_ActiveProcessDebuggingState.IsActive)
    {
        ProcId = g_ActiveProcessDebuggingState.ProcessId;
    }
    if (SplittedCommand.size() <= 4)
    {
        ShowMessages("incorrect use of 's*'\n\n");
        CommandSearchMemoryHelp();
        return;
    }
    for (auto Section : SplittedCommand)
    {
        IndexInCommandCaseSensitive++;
        if (!Section.compare(SplittedCommand.at(0)))
        {
            if (!Section.compare("!sb"))
            {
                SearchMemoryRequest.MemoryType = SEARCH_PHYSICAL_MEMORY;
                SearchMemoryRequest.ByteSize   = SEARCH_BYTE;
            }
            else if (!Section.compare("!sd"))
            {
                SearchMemoryRequest.MemoryType = SEARCH_PHYSICAL_MEMORY;
                SearchMemoryRequest.ByteSize   = SEARCH_DWORD;
            }
            else if (!Section.compare("!sq"))
            {
                SearchMemoryRequest.MemoryType = SEARCH_PHYSICAL_MEMORY;
                SearchMemoryRequest.ByteSize   = SEARCH_QWORD;
            }
            else if (!Section.compare("sb"))
            {
                SearchMemoryRequest.MemoryType = SEARCH_VIRTUAL_MEMORY;
                SearchMemoryRequest.ByteSize   = SEARCH_BYTE;
            }
            else if (!Section.compare("sd"))
            {
                SearchMemoryRequest.MemoryType = SEARCH_VIRTUAL_MEMORY;
                SearchMemoryRequest.ByteSize   = SEARCH_DWORD;
            }
            else if (!Section.compare("sq"))
            {
                SearchMemoryRequest.MemoryType = SEARCH_VIRTUAL_MEMORY;
                SearchMemoryRequest.ByteSize   = SEARCH_QWORD;
            }
            else
            {
                ShowMessages("unknown error happened !\n\n");
                CommandSearchMemoryHelp();
                return;
            }
            continue;
        }
        if (NextIsProcId)
        {
            NextIsProcId = FALSE;
            if (!ConvertStringToUInt32(Section, &ProcId))
            {
                ShowMessages("please specify a correct hex prcoess id\n\n");
                CommandSearchMemoryHelp();
                return;
            }
            else
            {
                continue;
            }
        }
        if (NextIsLength)
        {
            NextIsLength = FALSE;
            if (!ConvertStringToUInt64(Section, &Length))
            {
                ShowMessages("please specify a correct hex length\n\n");
                CommandSearchMemoryHelp();
                return;
            }
            else
            {
                SetLength = TRUE;
                continue;
            }
        }
        if (!SetProcId && !Section.compare("pid"))
        {
            NextIsProcId = TRUE;
            continue;
        }
        if (!SetLength && !Section.compare("l"))
        {
            NextIsLength = TRUE;
            continue;
        }
        if (!SetAddress)
        {
            if (!SymbolConvertNameOrExprToAddress(SplittedCommandCaseSensitive.at(IndexInCommandCaseSensitive - 1), &Address))
            {
                ShowMessages("err, couldn't resolve error at '%s'\n\n",
                             SplittedCommandCaseSensitive.at(IndexInCommandCaseSensitive - 1).c_str());
                CommandSearchMemoryHelp();
                return;
            }
            else
            {
                SetAddress = TRUE;
                continue;
            }
        }
        if (SetAddress)
        {
            if (Section.rfind("0x", 0) == 0 || Section.rfind("0X", 0) == 0 ||
                Section.rfind("\\x", 0) == 0 || Section.rfind("\\X", 0) == 0)
            {
                Section = Section.erase(0, 2);
            }
            else if (Section.rfind('x', 0) == 0 || Section.rfind('X', 0) == 0)
            {
                Section = Section.erase(0, 1);
            }
            Section.erase(remove(Section.begin(), Section.end(), '`'), Section.end());
            if (SearchMemoryRequest.ByteSize == SEARCH_BYTE && Section.size() >= 3)
            {
                ShowMessages("please specify a byte (hex) value for 'sb' or '!sb'\n\n");
                return;
            }
            if (SearchMemoryRequest.ByteSize == SEARCH_DWORD && Section.size() >= 9)
            {
                ShowMessages(
                    "please specify a dword (hex) value for 'sd' or '!sd'\n\n");
                return;
            }
            if (SearchMemoryRequest.ByteSize == SEARCH_QWORD &&
                Section.size() >= 17)
            {
                ShowMessages(
                    "please specify a qword (hex) value for 'sq' or '!sq'\n\n");
                return;
            }
            if (!ConvertStringToUInt64(Section, &Value))
            {
                ShowMessages("please specify a correct hex value to search in the "
                             "memory content\n\n");
                CommandSearchMemoryHelp();
                return;
            }
            else
            {
                ValuesToEdit.push_back(Value);
                CountOfValues++;
                if (!SetValue)
                {
                    SetValue = TRUE;
                }
                continue;
            }
        }
    }
    if (g_IsSerialConnectedToRemoteDebuggee && ProcId != 0)
    {
        ShowMessages("err, you cannot specify 'pid' in the debugger mode\n");
        return;
    }
    if (ProcId == 0)
    {
        ProcId = GetCurrentProcessId();
    }
    SearchMemoryRequest.ProcessId       = ProcId;
    SearchMemoryRequest.Address         = Address;
    SearchMemoryRequest.CountOf64Chunks = CountOfValues;
    if (!SetAddress)
    {
        ShowMessages("please specify a correct hex address\n\n");
        CommandSearchMemoryHelp();
        return;
    }
    if (!SetValue)
    {
        ShowMessages(
            "please specify a correct hex value as the content to search\n\n");
        CommandSearchMemoryHelp();
        return;
    }
    if (!SetLength)
    {
        ShowMessages("please specify a correct hex value as the length\n\n");
        CommandSearchMemoryHelp();
        return;
    }
    if (NextIsProcId)
    {
        ShowMessages("please specify a correct hex value as the process id\n\n");
        CommandSearchMemoryHelp();
        return;
    }
    if (NextIsLength)
    {
        ShowMessages("please specify a correct hex length\n\n");
        CommandSearchMemoryHelp();
        return;
    }
    FinalSize = (CountOfValues * sizeof(UINT64)) + SIZEOF_DEBUGGER_SEARCH_MEMORY;
    SearchMemoryRequest.FinalStructureSize = FinalSize;
    SearchMemoryRequest.Length = Length;
    if (!g_IsSerialConnectedToRemoteDebuggee)
    {
        AssertShowMessageReturnStmt(g_DeviceHandle, ASSERT_MESSAGE_DRIVER_NOT_LOADED, AssertReturn);
    }
    FinalBuffer = (UINT64 *)malloc(FinalSize);
    if (!FinalBuffer)
    {
        ShowMessages("unable to allocate memory\n\n");
        return;
    }
    ZeroMemory(FinalBuffer, FinalSize);
    memcpy(FinalBuffer, &SearchMemoryRequest, SIZEOF_DEBUGGER_SEARCH_MEMORY);
    std::copy(ValuesToEdit.begin(), ValuesToEdit.end(), (UINT64 *)((UINT64)FinalBuffer + SIZEOF_DEBUGGER_SEARCH_MEMORY));
    if (g_IsSerialConnectedToRemoteDebuggee)
    {
        KdSendSearchRequestPacketToDebuggee(FinalBuffer, FinalSize);
    }
    else
    {
        CommandSearchSendRequest(FinalBuffer, FinalSize);
    }
    free(FinalBuffer);
}*/
return true
}



