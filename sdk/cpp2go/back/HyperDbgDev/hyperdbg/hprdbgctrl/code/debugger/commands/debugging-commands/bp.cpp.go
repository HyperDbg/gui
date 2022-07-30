package debugging-commands
//back\HyperDbgDev\hyperdbg\hprdbgctrl\code\debugger\commands\debugging-commands\bp.cpp.back

type (
Bp interface{
CommandBpHelp()(ok bool)//col:44
CommandBp()(ok bool)//col:215
}
)

func NewBp() { return & bp{} }

func (b *bp)CommandBpHelp()(ok bool){//col:44
/*CommandBpHelp()
{
    ShowMessages("bp : puts a breakpoint (0xcc).\n");
    ShowMessages(
        "Note : 'bp' is not an event, if you want to use an event version "
        "of breakpoints use !epthook or !epthook2 instead. See "
        "documentation for more inforamtion.\n\n");
    ShowMessages("syntax : \tbp [Address (hex)] [pid ProcessId (hex)] [tid ThreadId (hex)] [core CoreId (hex)]\n");
    ShowMessages("\n");
    ShowMessages("\t\te.g : bp nt!ExAllocatePoolWithTag\n");
    ShowMessages("\t\te.g : bp nt!ExAllocatePoolWithTag+5\n");
    ShowMessages("\t\te.g : bp fffff8077356f010\n");
    ShowMessages("\t\te.g : bp fffff8077356f010 pid 0x4\n");
    ShowMessages("\t\te.g : bp fffff8077356f010 tid 0x1000\n");
    ShowMessages("\t\te.g : bp fffff8077356f010 pid 0x4 core 2\n");
}*/
return true
}

func (b *bp)CommandBp()(ok bool){//col:215
/*CommandBp(vector<string> SplittedCommand, string Command)
{
    BOOL IsNextCoreId = FALSE;
    BOOL IsNextPid    = FALSE;
    BOOL IsNextTid    = FALSE;
    BOOLEAN SetCoreId  = FALSE;
    BOOLEAN SetPid     = FALSE;
    BOOLEAN SetTid     = FALSE;
    BOOLEAN SetAddress = FALSE;
    UINT32         Tid       = DEBUGGEE_BP_APPLY_TO_ALL_THREADS;
    UINT32         Pid       = DEBUGGEE_BP_APPLY_TO_ALL_PROCESSES;
    UINT32         CoreNumer = DEBUGGEE_BP_APPLY_TO_ALL_CORES;
    UINT64         Address   = NULL;
    vector<string> SplittedCommandCaseSensitive {Split(Command, ' ')};
    UINT32         IndexInCommandCaseSensitive = 0;
    DEBUGGEE_BP_PACKET BpPacket = {0};
    if (SplittedCommand.size() >= 9)
    {
        ShowMessages("incorrect use of 'bp'\n\n");
        CommandBpHelp();
        return;
    }
    for (auto Section : SplittedCommand)
    {
        IndexInCommandCaseSensitive++;
        if (!Section.compare(SplittedCommand.at(0)))
        {
            continue;
        }
        if (IsNextCoreId)
        {
            if (!ConvertStringToUInt32(Section, &CoreNumer))
            {
                ShowMessages("please specify a correct hex value for core id\n\n");
                CommandBpHelp();
                return;
            }
            IsNextCoreId = FALSE;
            continue;
        }
        if (IsNextPid)
        {
            if (!ConvertStringToUInt32(Section, &Pid))
            {
                ShowMessages("please specify a correct hex value for process id\n\n");
                CommandBpHelp();
                return;
            }
            IsNextPid = FALSE;
            continue;
        }
        if (IsNextTid)
        {
            if (!ConvertStringToUInt32(Section, &Tid))
            {
                ShowMessages("please specify a correct hex value for thread id\n\n");
                CommandBpHelp();
                return;
            }
            IsNextTid = FALSE;
            continue;
        }
        if (!Section.compare("pid"))
        {
            IsNextPid = TRUE;
            continue;
        }
        if (!Section.compare("tid"))
        {
            IsNextTid = TRUE;
            continue;
        }
        if (!Section.compare("core"))
        {
            IsNextCoreId = TRUE;
            continue;
        }
        if (!SetAddress)
        {
            if (!SymbolConvertNameOrExprToAddress(SplittedCommandCaseSensitive.at(IndexInCommandCaseSensitive - 1), &Address))
            {
                ShowMessages("err, couldn't resolve error at '%s'\n\n",
                             SplittedCommandCaseSensitive.at(IndexInCommandCaseSensitive - 1).c_str());
                CommandBpHelp();
                return;
            }
            else
            {
                SetAddress = TRUE;
                continue;
            }
        }
    }
    if (!SetAddress)
    {
        ShowMessages(
            "please specify a correct hex value as the breakpoint address\n\n");
        CommandBpHelp();
        return;
    }
    if (IsNextPid)
    {
        ShowMessages("please specify a correct hex value for process id\n\n");
        CommandBpHelp();
        return;
    }
    if (IsNextCoreId)
    {
        ShowMessages("please specify a correct hex value for core\n\n");
        CommandBpHelp();
        return;
    }
    if (IsNextTid)
    {
        ShowMessages("please specify a correct hex value for thread id\n\n");
        CommandBpHelp();
        return;
    }
    if (!g_IsSerialConnectedToRemoteDebuggee)
    {
        ShowMessages("err, setting breakpoints is not possible when you're not "
                     "connected to a debuggee\n");
        return;
    }
    BpPacket.Address = Address;
    BpPacket.Core    = CoreNumer;
    BpPacket.Pid     = Pid;
    BpPacket.Tid     = Tid;
    KdSendBpPacketToDebuggee(&BpPacket);
}*/
return true
}



