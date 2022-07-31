package debugging-commands
//back\HyperDbgDev\hyperdbg\hprdbgctrl\code\debugger\commands\debugging-commands\e.cpp.back

type (
E interface{
CommandEditMemoryHelp()(ok bool)//col:49
CommandEditMemory()(ok bool)//col:410
}
)

func NewE() { return & e{} }

func (e *e)CommandEditMemoryHelp()(ok bool){//col:49
/*CommandEditMemoryHelp()
{
    ShowMessages("eb !eb ed !ed eq !eq : edits the memory at specific address \n");
    ShowMessages("eb  Byte and ASCII characters\n");
    ShowMessages("ed  Double-word values (4 bytes)\n");
    ShowMessages("eq  Quad-word values (8 bytes). \n");
    ShowMessages("\n If you want to edit physical (address) memory then add '!' "
                 "at the start of the command\n");
    ShowMessages("syntax : \teb [Address (hex)] [Contents (hex)] [pid ProcessId (hex)]\n");
    ShowMessages("syntax : \ted [Address (hex)] [Contents (hex)] [pid ProcessId (hex)]\n");
    ShowMessages("syntax : \teq [Address (hex)] [Contents (hex)] [pid ProcessId (hex)]\n");
    ShowMessages("\n");
    ShowMessages("\t\te.g : eb fffff8077356f010 90 \n");
    ShowMessages("\t\te.g : eb nt!Kd_DEFAULT_Mask ff ff ff ff \n");
    ShowMessages("\t\te.g : eb fffff8077356f010 90 90 90 90 \n");
    ShowMessages("\t\te.g : !eq 100000 9090909090909090\n");
    ShowMessages("\t\te.g : !eq nt!ExAllocatePoolWithTag+55 9090909090909090\n");
    ShowMessages("\t\te.g : !eq 100000 9090909090909090 9090909090909090 "
                 "9090909090909090 9090909090909090 9090909090909090\n");
}*/
return true
}

func (e *e)CommandEditMemory()(ok bool){//col:410
/*CommandEditMemory(vector<string> SplittedCommand, string Command)
{
    BOOL                 Status;
    UINT64               Address;
    UINT64 *             FinalBuffer;
    vector<UINT64>       ValuesToEdit;
    BOOL                 SetAddress        = FALSE;
    BOOL                 SetValue          = FALSE;
    BOOL                 SetProcId         = FALSE;
    BOOL                 NextIsProcId      = FALSE;
    DEBUGGER_EDIT_MEMORY EditMemoryRequest = {0};
    UINT64               Value             = 0;
    UINT32               ProcId            = 0;
    UINT32               CountOfValues     = 0;
    UINT32               FinalSize         = 0;
    vector<string>       SplittedCommandCaseSensitive {Split(Command, ' ')};
    UINT32               IndexInCommandCaseSensitive = 0;
    if (g_ActiveProcessDebuggingState.IsActive)
    {
        ProcId = g_ActiveProcessDebuggingState.ProcessId;
    }
    if (SplittedCommand.size() <= 2)
    {
        ShowMessages("incorrect use of 'e*'\n\n");
        CommandEditMemoryHelp();
        return;
    }
    for (auto Section : SplittedCommand)
    {
        IndexInCommandCaseSensitive++;
        if (!Section.compare(SplittedCommand.at(0)))
        {
            if (!Section.compare("!eb"))
            {
                EditMemoryRequest.MemoryType = EDIT_PHYSICAL_MEMORY;
                EditMemoryRequest.ByteSize   = EDIT_BYTE;
            }
            else if (!Section.compare("!ed"))
            {
                EditMemoryRequest.MemoryType = EDIT_PHYSICAL_MEMORY;
                EditMemoryRequest.ByteSize   = EDIT_DWORD;
            }
            else if (!Section.compare("!eq"))
            {
                EditMemoryRequest.MemoryType = EDIT_PHYSICAL_MEMORY;
                EditMemoryRequest.ByteSize   = EDIT_QWORD;
            }
            else if (!Section.compare("eb"))
            {
                EditMemoryRequest.MemoryType = EDIT_VIRTUAL_MEMORY;
                EditMemoryRequest.ByteSize   = EDIT_BYTE;
            }
            else if (!Section.compare("ed"))
            {
                EditMemoryRequest.MemoryType = EDIT_VIRTUAL_MEMORY;
                EditMemoryRequest.ByteSize   = EDIT_DWORD;
            }
            else if (!Section.compare("eq"))
            {
                EditMemoryRequest.MemoryType = EDIT_VIRTUAL_MEMORY;
                EditMemoryRequest.ByteSize   = EDIT_QWORD;
            }
            else
            {
                ShowMessages("unknown error happened !\n\n");
                CommandEditMemoryHelp();
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
                CommandEditMemoryHelp();
                return;
            }
            else
            {
                continue;
            }
        }
        if (!SetProcId && !Section.compare("pid"))
        {
            NextIsProcId = TRUE;
            continue;
        }
        if (!SetAddress)
        {
            if (!SymbolConvertNameOrExprToAddress(SplittedCommandCaseSensitive.at(IndexInCommandCaseSensitive - 1),
                                                  &Address))
            {
                ShowMessages("err, couldn't resolve error at '%s'\n\n",
                             SplittedCommandCaseSensitive.at(IndexInCommandCaseSensitive - 1).c_str());
                CommandEditMemoryHelp();
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
            if (EditMemoryRequest.ByteSize == EDIT_BYTE && Section.size() >= 3)
            {
                ShowMessages("please specify a byte (hex) value for 'eb' or '!eb'\n\n");
                return;
            }
            if (EditMemoryRequest.ByteSize == EDIT_DWORD && Section.size() >= 9)
            {
                ShowMessages(
                    "please specify a dword (hex) value for 'ed' or '!ed'\n\n");
                return;
            }
            if (EditMemoryRequest.ByteSize == EDIT_QWORD && Section.size() >= 17)
            {
                ShowMessages(
                    "please specify a qword (hex) value for 'eq' or '!eq'\n\n");
                return;
            }
            if (!ConvertStringToUInt64(Section, &Value))
            {
                ShowMessages("please specify a correct hex value to change the memory "
                             "content\n\n");
                CommandEditMemoryHelp();
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
        ShowMessages("err, you cannot specify 'pid' in the debugger mode\n\n");
        return;
    }
    if (ProcId == 0)
    {
        ProcId = GetCurrentProcessId();
    }
    EditMemoryRequest.ProcessId       = ProcId;
    EditMemoryRequest.Address         = Address;
    EditMemoryRequest.CountOf64Chunks = CountOfValues;
    if (!SetAddress)
    {
        ShowMessages("please specify a correct hex address\n\n");
        CommandEditMemoryHelp();
        return;
    }
    if (!SetValue)
    {
        ShowMessages(
            "please specify a correct hex value as the content to edit\n\n");
        CommandEditMemoryHelp();
        return;
    }
    if (NextIsProcId)
    {
        ShowMessages("please specify a correct hex value as the process id\n\n");
        CommandEditMemoryHelp();
        return;
    }
    FinalSize = (CountOfValues * sizeof(UINT64)) + SIZEOF_DEBUGGER_EDIT_MEMORY;
    EditMemoryRequest.FinalStructureSize = FinalSize;
    FinalBuffer = (UINT64 *)malloc(FinalSize);
    if (!FinalBuffer)
    {
        ShowMessages("unable to allocate memory\n\n");
        return;
    }
    ZeroMemory(FinalBuffer, FinalSize);
    memcpy(FinalBuffer, &EditMemoryRequest, SIZEOF_DEBUGGER_EDIT_MEMORY);
    std::copy(ValuesToEdit.begin(), ValuesToEdit.end(), (UINT64 *)((UINT64)FinalBuffer + SIZEOF_DEBUGGER_EDIT_MEMORY));
    if (g_IsSerialConnectedToRemoteDebuggee)
    {
        KdSendEditMemoryPacketToDebuggee((DEBUGGER_EDIT_MEMORY *)FinalBuffer, FinalSize);
        return;
    }
    AssertShowMessageReturnStmt(g_DeviceHandle, ASSERT_MESSAGE_DRIVER_NOT_LOADED, AssertReturn);
    Status = DeviceIoControl(
    );
    if (!Status)
    {
        free(FinalBuffer);
        ShowMessages("ioctl failed with code 0x%x\n", GetLastError());
        return;
    }
    if (EditMemoryRequest.Result == DEBUGGER_OPERATION_WAS_SUCCESSFULL)
    {
    }
    else if (
        EditMemoryRequest.Result ==
        DEBUGGER_ERROR_EDIT_MEMORY_STATUS_INVALID_ADDRESS_BASED_ON_CURRENT_PROCESS)
    {
        ShowMessages("err, the address is invalid in system process layout\n");
    }
    else if (
        EditMemoryRequest.Result ==
        DEBUGGER_ERROR_EDIT_MEMORY_STATUS_INVALID_ADDRESS_BASED_ON_OTHER_PROCESS)
    {
        ShowMessages("err, the address is invalid based on your specific process id\n");
    }
    else if (EditMemoryRequest.Result ==
             DEBUGGER_ERROR_EDIT_MEMORY_STATUS_INVALID_PARAMETER)
    {
        ShowMessages("invalid parameter\n");
    }
    else
    {
        ShowErrorMessage(EditMemoryRequest.Result);
    }
    free(FinalBuffer);
}*/
return true
}



