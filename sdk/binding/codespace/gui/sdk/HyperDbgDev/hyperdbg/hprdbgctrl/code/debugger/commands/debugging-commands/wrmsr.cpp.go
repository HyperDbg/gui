package debugging-commands
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\debugger\commands\debugging-commands\wrmsr.cpp.back

type (
Wrmsr interface{
CommandWrmsrHelp()(ok bool)//col:10
CommandWrmsr()(ok bool)//col:119
}
)

func NewWrmsr() { return & wrmsr{} }

func (w *wrmsr)CommandWrmsrHelp()(ok bool){//col:10
/*CommandWrmsrHelp()
{
    ShowMessages("wrmsr : writes on a model-specific register (MSR).\n\n");
    ShowMessages("syntax : \twrmsr [Msr (hex)] [Value (hex)] [core CoreNumber (hex)]\n");
    ShowMessages("\n");
    ShowMessages("\t\te.g : wrmsr c0000082 fffff8077356f010\n");
    ShowMessages("\t\te.g : wrmsr c0000082 fffff8077356f010 core 2\n");
}*/
return true
}

func (w *wrmsr)CommandWrmsr()(ok bool){//col:119
/*CommandWrmsr(vector<string> SplittedCommand, string Command)
{
    BOOL                           Status;
    BOOL                           IsNextCoreId = FALSE;
    BOOL                           SetMsr       = FALSE;
    BOOL                           SetValue     = FALSE;
    DEBUGGER_READ_AND_WRITE_ON_MSR MsrWriteRequest;
    UINT64                         Msr;
    UINT64                         Value     = 0;
    UINT32                         CoreNumer = DEBUGGER_READ_AND_WRITE_ON_MSR_APPLY_ALL_CORES;
    if (SplittedCommand.size() >= 6)
    {
        ShowMessages("incorrect use of 'wrmsr'\n\n");
        CommandWrmsrHelp();
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
                CommandWrmsrHelp();
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
        if (!SetMsr)
        {
            if (!ConvertStringToUInt64(Section, &Msr))
            {
                ShowMessages("please specify a correct hex value to be read\n\n");
                CommandWrmsrHelp();
                return;
            }
            else
            {
                SetMsr = TRUE;
                continue;
            }
        }
        if (SetMsr)
        {
            if (!SymbolConvertNameOrExprToAddress(Section, &Value))
            {
                ShowMessages(
                    "please specify a correct hex value or an expression to put on the msr\n\n");
                CommandWrmsrHelp();
                return;
            }
            else
            {
                SetValue = TRUE;
                continue;
            }
        }
    }
    if (!SetMsr)
    {
        ShowMessages("please specify a correct hex value to write\n\n");
        CommandWrmsrHelp();
        return;
    }
    if (!SetValue)
    {
        ShowMessages("please specify a correct hex value to put on msr\n\n");
        CommandWrmsrHelp();
        return;
    }
    if (IsNextCoreId)
    {
        ShowMessages("please specify a correct hex value for core\n\n");
        CommandWrmsrHelp();
        return;
    }
    AssertShowMessageReturnStmt(g_DeviceHandle, ASSERT_MESSAGE_DRIVER_NOT_LOADED, AssertReturn);
    MsrWriteRequest.ActionType = DEBUGGER_MSR_WRITE;
    MsrWriteRequest.Msr        = Msr;
    MsrWriteRequest.CoreNumber = CoreNumer;
    MsrWriteRequest.Value      = Value;
    Status = DeviceIoControl(
        g_DeviceHandle,                        
        IOCTL_DEBUGGER_READ_OR_WRITE_MSR,      
        &MsrWriteRequest,                      
        SIZEOF_DEBUGGER_READ_AND_WRITE_ON_MSR, 
        NULL,                                  
        NULL,                                  
        NULL,                                  
        NULL                                   
    );
    if (!Status)
    {
        ShowMessages("ioctl failed with code (%x), either msr index or core id is invalid\n",
                     GetLastError());
        return;
    }
    ShowMessages("\n");
}*/
return true
}



