package meta-commands
//back\HyperDbgDev\hyperdbg\hprdbgctrl\code\debugger\commands\meta-commands\switch.cpp.back

type (
SwitchA interface{
CommandSwitchHelp()(ok bool)//col:12
CommandSwitch()(ok bool)//col:63
}
)

func NewSwitchA() { return & switchA{} }

func (s *switchA)CommandSwitchHelp()(ok bool){//col:12
/*CommandSwitchHelp()
{
    ShowMessages(".switch : shows the list of active debugging threads and switches "
                 "between processes and threads in VMI Mode.\n\n");
    ShowMessages("syntax : \t.switch \n");
    ShowMessages("syntax : \t.switch [pid ProcessId (hex)]\n");
    ShowMessages("syntax : \t.switch [tid ThreadId (hex)]\n");
    ShowMessages("\n");
    ShowMessages("\t\te.g : .switch list\n");
    ShowMessages("\t\te.g : .switch pid b60 \n");
    ShowMessages("\t\te.g : .switch tid b60 \n");
}*/
return true
}

func (s *switchA)CommandSwitch()(ok bool){//col:63
/*CommandSwitch(vector<string> SplittedCommand, string Command)
{
    UINT32 PidOrTid = NULL;
    if (SplittedCommand.size() > 3 || SplittedCommand.size() == 2)
    {
        ShowMessages("incorrect use of '.switch'\n\n");
        CommandSwitchHelp();
        return;
    }
    if (g_IsSerialConnectedToRemoteDebuggee)
    {
        ShowMessages("err, the '.switch' command is only usable in VMI Mode, "
                     "you can use the '.process', or the '.thread' "
                     "in Debugger Mode\n");
        return;
    }
    if (SplittedCommand.size() == 1)
    {
        UdShowListActiveDebuggingProcessesAndThreads();
    }
    else if (!SplittedCommand.at(1).compare("pid"))
    {
        if (!ConvertStringToUInt32(SplittedCommand.at(2), &PidOrTid))
        {
            ShowMessages("please specify a correct hex value for process id\n\n");
            return;
        }
        if (UdSetActiveDebuggingThreadByPidOrTid(PidOrTid, FALSE))
        {
            ShowMessages("switched to process id: %x\n", PidOrTid);
        }
    }
    else if (!SplittedCommand.at(1).compare("tid"))
    {
        if (!ConvertStringToUInt32(SplittedCommand.at(2), &PidOrTid))
        {
            ShowMessages("please specify a correct hex value for thread id\n\n");
            return;
        }
        if (UdSetActiveDebuggingThreadByPidOrTid(PidOrTid, TRUE))
        {
            ShowMessages("switched to thread id: %x\n", PidOrTid);
        }
    }
    else
    {
        ShowMessages("err, couldn't resolve error at '%s'\n",
                     SplittedCommand.at(1).c_str());
        return;
    }
}*/
return true
}



