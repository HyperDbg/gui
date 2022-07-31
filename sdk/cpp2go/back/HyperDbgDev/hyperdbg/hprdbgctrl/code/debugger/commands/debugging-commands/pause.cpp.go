package debugging-commands
//back\HyperDbgDev\hyperdbg\hprdbgctrl\code\debugger\commands\debugging-commands\pause.cpp.back

type (
Pause interface{
CommandPauseHelp()(ok bool)//col:32
CommandPauseRequest()(ok bool)//col:61
CommandPause()(ok bool)//col:81
}
)

func NewPause() { return & pause{} }

func (p *pause)CommandPauseHelp()(ok bool){//col:32
/*CommandPauseHelp()
{
    ShowMessages("pause : pauses the kernel events.\n\n");
    ShowMessages("syntax : \tpause \n");
}*/
return true
}

func (p *pause)CommandPauseRequest()(ok bool){//col:61
/*CommandPauseRequest()
{
    g_BreakPrintingOutput = TRUE;
    ShowMessages("pausing...\n");
    if (g_IsConnectedToRemoteDebuggee)
    {
        RemoteConnectionSendCommand("pause", strlen("pause") + 1);
    }
    else if (g_ActiveProcessDebuggingState.IsActive && UdPauseProcess(g_ActiveProcessDebuggingState.ProcessDebuggingToken))
    {
        ShowMessages("please keep interacting with the process until all the "
                     "threads are intercepted and halted; whenever you execute "
                     "the first command, the thread interception will be stopped\n");
    }
}*/
return true
}

func (p *pause)CommandPause()(ok bool){//col:81
/*CommandPause(vector<string> SplittedCommand, string Command)
{
    if (SplittedCommand.size() != 1)
    {
        ShowMessages("incorrect use of 'pause'\n\n");
        CommandPauseHelp();
        return;
    }
    CommandPauseRequest();
}*/
return true
}


