package meta-commands
//back\HyperDbgDev\hyperdbg\hprdbgctrl\code\debugger\commands\meta-commands\kill.cpp.back

type (
Kill interface{
CommandKillHelp()(ok bool)//col:30
CommandKill()(ok bool)//col:59
}
)

func NewKill() { return & kill{} }

func (k *kill)CommandKillHelp()(ok bool){//col:30
/*CommandKillHelp()
{
    ShowMessages(".kill : terminates the current running process.\n\n");
    ShowMessages("syntax : \t.kill \n");
}*/
return true
}

func (k *kill)CommandKill()(ok bool){//col:59
/*CommandKill(vector<string> SplittedCommand, string Command)
{
    if (SplittedCommand.size() != 1)
    {
        ShowMessages("incorrect use of '.kill'\n\n");
        CommandKillHelp();
        return;
    }
    if (!g_ActiveProcessDebuggingState.IsActive)
    {
        ShowMessages("nothing to terminate!\n");
        return;
    }
    UdKillProcess(g_ActiveProcessDebuggingState.ProcessId);
}*/
return true
}



