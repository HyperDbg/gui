package debugging-commands
//back\HyperDbgDev\hyperdbg\hprdbgctrl\code\debugger\commands\debugging-commands\be.cpp.back

type (
Be interface{
CommandBeHelp()(ok bool)//col:8
CommandBe()(ok bool)//col:36
}
)

func NewBe() { return & be{} }

func (b *be)CommandBeHelp()(ok bool){//col:8
/*CommandBeHelp()
{
    ShowMessages("be : enables a breakpoint using breakpoint id.\n\n");
    ShowMessages("syntax : \tbe [BreakpointId (hex)]\n");
    ShowMessages("\n");
    ShowMessages("\t\te.g : be 0\n");
    ShowMessages("\t\te.g : be 2\n");
}*/
return true
}

func (b *be)CommandBe()(ok bool){//col:36
/*CommandBe(vector<string> SplittedCommand, string Command)
{
    UINT64                            BreakpointId;
    DEBUGGEE_BP_LIST_OR_MODIFY_PACKET Request = {0};
    if (SplittedCommand.size() != 2)
    {
        ShowMessages("incorrect use of 'be'\n\n");
        CommandBeHelp();
        return;
    }
    if (!ConvertStringToUInt64(SplittedCommand.at(1), &BreakpointId))
    {
        ShowMessages("please specify a correct hex value for breakpoint id\n\n");
        CommandBeHelp();
        return;
    }
    if (g_IsSerialConnectedToRemoteDebuggee)
    {
        Request.Request = DEBUGGEE_BREAKPOINT_MODIFICATION_REQUEST_ENABLE;
        Request.BreakpointId = BreakpointId;
        KdSendListOrModifyPacketToDebuggee(&Request);
    }
    else
    {
        ShowMessages("err, enabling breakpoints is only valid if you connected to "
                     "a debuggee in debugger-mode\n");
    }
}*/
return true
}



