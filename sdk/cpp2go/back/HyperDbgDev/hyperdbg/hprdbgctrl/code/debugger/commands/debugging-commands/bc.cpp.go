package debugging-commands
//back\HyperDbgDev\hyperdbg\hprdbgctrl\code\debugger\commands\debugging-commands\bc.cpp.back

type (
Bc interface{
CommandBcHelp()(ok bool)//col:34
CommandBc()(ok bool)//col:95
}
)

func NewBc() { return & bc{} }

func (b *bc)CommandBcHelp()(ok bool){//col:34
/*CommandBcHelp()
{
    ShowMessages("bc : clears a breakpoint using breakpoint id.\n\n");
    ShowMessages("syntax : \tbc [BreakpointId (hex)]\n");
    ShowMessages("\n");
    ShowMessages("\t\te.g : bc 0\n");
    ShowMessages("\t\te.g : bc 2\n");
}*/
return true
}

func (b *bc)CommandBc()(ok bool){//col:95
/*CommandBc(vector<string> SplittedCommand, string Command)
{
    UINT64                            BreakpointId;
    DEBUGGEE_BP_LIST_OR_MODIFY_PACKET Request = {0};
    if (SplittedCommand.size() != 2)
    {
        ShowMessages("incorrect use of 'bc'\n\n");
        CommandBcHelp();
        return;
    }
    if (!ConvertStringToUInt64(SplittedCommand.at(1), &BreakpointId))
    {
        ShowMessages("please specify a correct hex value for breakpoint id\n\n");
        CommandBcHelp();
        return;
    }
    if (g_IsSerialConnectedToRemoteDebuggee)
    {
        Request.Request = DEBUGGEE_BREAKPOINT_MODIFICATION_REQUEST_CLEAR;
        Request.BreakpointId = BreakpointId;
        KdSendListOrModifyPacketToDebuggee(&Request);
    }
    else
    {
        ShowMessages("err, clearing breakpoints is only valid if you connected to "
                     "a debuggee in debugger-mode\n");
    }
}*/
return true
}



