package debugging-commands
//back\HyperDbgDev\hyperdbg\hprdbgctrl\code\debugger\commands\debugging-commands\bd.cpp.back

type (
Bd interface{
CommandBdHelp()(ok bool)//col:34
CommandBd()(ok bool)//col:95
}
)

func NewBd() { return & bd{} }

func (b *bd)CommandBdHelp()(ok bool){//col:34
/*CommandBdHelp()
{
    ShowMessages("bd : disables a breakpoint using breakpoint id.\n\n");
    ShowMessages("syntax : \tbd [BreakpointId (hex)]\n");
    ShowMessages("\n");
    ShowMessages("\t\te.g : bd 0\n");
    ShowMessages("\t\te.g : bd 2\n");
}*/
return true
}

func (b *bd)CommandBd()(ok bool){//col:95
/*CommandBd(vector<string> SplittedCommand, string Command)
{
    UINT64                            BreakpointId;
    DEBUGGEE_BP_LIST_OR_MODIFY_PACKET Request = {0};
    if (SplittedCommand.size() != 2)
    {
        ShowMessages("incorrect use of 'bd'\n\n");
        CommandBdHelp();
        return;
    }
    if (!ConvertStringToUInt64(SplittedCommand.at(1), &BreakpointId))
    {
        ShowMessages("please specify a correct hex value for breakpoint id\n\n");
        CommandBdHelp();
        return;
    }
    if (g_IsSerialConnectedToRemoteDebuggee)
    {
        Request.Request = DEBUGGEE_BREAKPOINT_MODIFICATION_REQUEST_DISABLE;
        Request.BreakpointId = BreakpointId;
        KdSendListOrModifyPacketToDebuggee(&Request);
    }
    else
    {
        ShowMessages("err, disabling breakpoints is only valid if you connected to "
                     "a debuggee in debugger-mode\n");
    }
}*/
return true
}



