package debugging-commands
//back\HyperDbgDev\hyperdbg\hprdbgctrl\code\debugger\commands\debugging-commands\sleep.cpp.back

type (
Sleep interface{
CommandSleepHelp()(ok bool)//col:27
CommandSleep()(ok bool)//col:56
}
)

func NewSleep() { return & sleep{} }

func (s *sleep)CommandSleepHelp()(ok bool){//col:27
/*CommandSleepHelp()
{
    ShowMessages("sleep : sleeps the debugger; this command is used in scripts, it doesn't breaks "
                 "the debugger but the debugger still shows the buffers received "
                 "from kernel.\n\n");
    ShowMessages("syntax : \tsleep [MillisecondsTime (hex)]\n");
}*/
return true
}

func (s *sleep)CommandSleep()(ok bool){//col:56
/*CommandSleep(vector<string> SplittedCommand, string Command)
{
    UINT32 MillisecondsTime = 0;
    if (SplittedCommand.size() != 2)
    {
        ShowMessages("incorrect use of 'sleep'\n\n");
        CommandSleepHelp();
        return;
    }
    if (!ConvertStringToUInt32(SplittedCommand.at(1), &MillisecondsTime))
    {
        ShowMessages(
            "please specify a correct hex value for time (milliseconds)\n\n");
        CommandSleepHelp();
        return;
    }
    Sleep(MillisecondsTime);
}*/
return true
}



