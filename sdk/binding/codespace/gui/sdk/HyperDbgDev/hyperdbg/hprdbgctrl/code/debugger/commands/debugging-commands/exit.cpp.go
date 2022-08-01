package debugging-commands
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\debugger\commands\debugging-commands\exit.cpp.back

type (
Exit interface{
CommandExitHelp()(ok bool)//col:6
CommandExit()(ok bool)//col:20
}
exit struct{}
)

func NewExit()Exit{ return & exit{} }

func (e *exit)CommandExitHelp()(ok bool){//col:6
/*CommandExitHelp()
{
    ShowMessages(
        "exit : unloads and uninstalls the drivers and closes the debugger.\n\n");
    ShowMessages("syntax : \texit\n");
}*/
return true
}

func (e *exit)CommandExit()(ok bool){//col:20
/*CommandExit(vector<string> SplittedCommand, string Command)
{
    if (SplittedCommand.size() != 1)
    {
        ShowMessages("incorrect use of 'exit'\n\n");
        CommandExitHelp();
        return;
    }
    if (g_DeviceHandle)
    {
        HyperDbgUnloadVmm();
    }
    exit(0);
}*/
return true
}



