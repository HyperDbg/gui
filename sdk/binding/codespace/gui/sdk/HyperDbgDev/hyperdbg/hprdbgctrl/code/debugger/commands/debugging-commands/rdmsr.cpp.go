package debugging-commands
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\debugger\commands\debugging-commands\rdmsr.cpp.back

type (
Rdmsr interface{
CommandRdmsrHelp()(ok bool)//col:8
CommandRdmsr()(ok bool)//col:105
}
rdmsr struct{}
)

func NewRdmsr()Rdmsr{ return & rdmsr{} }

func (r *rdmsr)CommandRdmsrHelp()(ok bool){//col:8
/*CommandRdmsrHelp()
{
    ShowMessages("rdmsr : reads a model-specific register (MSR).\n\n");
    ShowMessages("syntax : \trdmsr [Msr (hex)] [core CoreNumber (hex)]\n");
    ShowMessages("\n");
    ShowMessages("\t\te.g : rdmsr c0000082\n");
    ShowMessages("\t\te.g : rdmsr c0000082 core 2\n");
}*/
return true
}

func (r *rdmsr)CommandRdmsr()(ok bool){//col:105
/*CommandRdmsr(vector<string> SplittedCommand, string Command)
{
    BOOL                           Status;
    BOOL                           IsNextCoreId = FALSE;
    BOOL                           SetMsr       = FALSE;
    DEBUGGER_READ_AND_WRITE_ON_MSR MsrReadRequest;
    ULONG                          ReturnedLength;
    UINT64                         Msr;
    UINT32                         CoreNumer = DEBUGGER_READ_AND_WRITE_ON_MSR_APPLY_ALL_CORES;
    SYSTEM_INFO                    SysInfo;
    DWORD                          NumCPU;
    if (SplittedCommand.size() >= 5)
    {
        ShowMessages("incorrect use of 'rdmsr'\n\n");
        CommandRdmsrHelp();
        return;
    }
    for (auto Section : SplittedCommand)
    {
        if (!Section.compare(SplittedCommand.at(0)))
        {
            continue;
        }
        if (IsNextCoreId)
        {
            if (!ConvertStringToUInt32(Section, &CoreNumer))
            {
                ShowMessages("please specify a correct hex value for core id\n\n");
                CommandRdmsrHelp();
                return;
            }
            IsNextCoreId = FALSE;
            continue;
        }
        if (!Section.compare("core"))
        {
            IsNextCoreId = TRUE;
            continue;
        }
        if (SetMsr || !ConvertStringToUInt64(Section, &Msr))
        {
            ShowMessages("please specify a correct hex value to be read\n\n");
            CommandRdmsrHelp();
            return;
        }
        SetMsr = TRUE;
    }
    if (!SetMsr)
    {
        ShowMessages("please specify a correct hex value to be read\n\n");
        CommandRdmsrHelp();
        return;
    }
    if (IsNextCoreId)
    {
        ShowMessages("please specify a correct hex value for core\n\n");
        CommandRdmsrHelp();
        return;
    }
    AssertShowMessageReturnStmt(g_DeviceHandle, ASSERT_MESSAGE_DRIVER_NOT_LOADED, AssertReturn);
    MsrReadRequest.ActionType = DEBUGGER_MSR_READ;
    MsrReadRequest.Msr        = Msr;
    MsrReadRequest.CoreNumber = CoreNumer;
    GetSystemInfo(&SysInfo);
    NumCPU = SysInfo.dwNumberOfProcessors;
    UINT64 * OutputBuffer = (UINT64 *)malloc(sizeof(UINT64) * NumCPU);
    ZeroMemory(OutputBuffer, sizeof(UINT64) * NumCPU);
    Status = DeviceIoControl(
        g_DeviceHandle,                        
        IOCTL_DEBUGGER_READ_OR_WRITE_MSR,      
        &MsrReadRequest,                       
        SIZEOF_DEBUGGER_READ_AND_WRITE_ON_MSR, 
        OutputBuffer,                          
        sizeof(UINT64) * NumCPU,               
        &ReturnedLength,                       
        NULL                                   
    );
    if (!Status)
    {
        ShowMessages("ioctl failed with code (%x), either msr index or core id is invalid\n",
                     GetLastError());
        free(OutputBuffer);
        return;
    }
    if (CoreNumer == DEBUGGER_READ_AND_WRITE_ON_MSR_APPLY_ALL_CORES)
    {
        for (size_t i = 0; i < NumCPU; i++)
        {
            ShowMessages("core : 0x%x - msr[%llx] = %s\n", i, Msr, SeparateTo64BitValue((OutputBuffer[i])).c_str());
        }
    }
    else
    {
        ShowMessages("core : 0x%x - msr[%llx] = %s\n", CoreNumer, Msr, SeparateTo64BitValue((OutputBuffer[0])).c_str());
    }
    free(OutputBuffer);
}*/
return true
}



