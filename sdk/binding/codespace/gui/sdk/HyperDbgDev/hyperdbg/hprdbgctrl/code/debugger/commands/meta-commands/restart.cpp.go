package meta-commands
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\debugger\commands\meta-commands\restart.cpp.back

type (
Restart interface{
CommandRestartHelp()(ok bool)//col:7
CommandRestart()(ok bool)//col:37
}
)

func NewRestart() { return & restart{} }

func (r *restart)CommandRestartHelp()(ok bool){//col:7
/*CommandRestartHelp()
{
    ShowMessages(".restart : restarts the previously started process "
                 "(using '.start' command).\n\n");
    ShowMessages(
        "syntax : \t.restart \n");
}*/
return true
}

func (r *restart)CommandRestart()(ok bool){//col:37
/*CommandRestart(vector<string> SplittedCommand, string Command)
{
    if (SplittedCommand.size() != 1)
    {
        ShowMessages("incorrect use of '.restart'\n\n");
        CommandRestartHelp();
        return;
    }
    if (g_StartCommandPath.empty())
    {
        ShowMessages("nothing to restart, did you use the '.start' command before?\n");
        return;
    }
    if (g_ActiveProcessDebuggingState.IsActive)
    {
        UdKillProcess(g_ActiveProcessDebuggingState.ProcessId);
    }
    if (g_StartCommandPathAndArguments.empty())
    {
        UdAttachToProcess(NULL,
                          g_StartCommandPath.c_str(),
                          NULL);
    }
    else
    {
        UdAttachToProcess(NULL,
                          g_StartCommandPath.c_str(),
                          (WCHAR *)g_StartCommandPathAndArguments.c_str());
    }
}*/
return true
}



