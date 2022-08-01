package debugging-commands
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\debugger\commands\debugging-commands\g.cpp.back

type (
G interface{
CommandGHelp()(ok bool)//col:5
CommandGRequest()(ok bool)//col:32
CommandG()(ok bool)//col:42
}
)

func NewG() { return & g{} }

func (g *g)CommandGHelp()(ok bool){//col:5
/*CommandGHelp()
{
    ShowMessages("g : continues debuggee or continues processing kernel messages.\n\n");
    ShowMessages("syntax : \tg \n");
}*/
return true
}

func (g *g)CommandGRequest()(ok bool){//col:32
/*CommandGRequest()
{
    if (g_IsSerialConnectedToRemoteDebuggee)
    {
        KdBreakControlCheckAndContinueDebugger();
    }
    else
    {
        g_BreakPrintingOutput = FALSE;
        if (g_IsConnectedToRemoteDebuggee)
        {
            RemoteConnectionSendCommand("g", strlen("g") + 1);
        }
        else if (g_ActiveProcessDebuggingState.IsActive)
        {
            if (g_ActiveProcessDebuggingState.IsPaused)
            {
                UdContinueDebuggee(g_ActiveProcessDebuggingState.ProcessDebuggingToken);
                g_ActiveProcessDebuggingState.IsPaused = FALSE;
            }
            else
            {
                ShowMessages("err, target process is already running\n");
            }
        }
    }
}*/
return true
}

func (g *g)CommandG()(ok bool){//col:42
/*CommandG(vector<string> SplittedCommand, string Command)
{
    if (SplittedCommand.size() != 1)
    {
        ShowMessages("incorrect use of 'g'\n\n");
        CommandGHelp();
        return;
    }
    CommandGRequest();
}*/
return true
}



