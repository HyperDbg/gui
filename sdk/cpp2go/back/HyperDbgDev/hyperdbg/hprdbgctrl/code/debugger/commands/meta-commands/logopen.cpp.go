package meta-commands
//back\HyperDbgDev\hyperdbg\hprdbgctrl\code\debugger\commands\meta-commands\logopen.cpp.back

type (
Logopen interface{
CommandLogopenHelp()(ok bool)//col:33
CommandLogopen()(ok bool)//col:110
LogopenSaveToFile()(ok bool)//col:122
}
)

func NewLogopen() { return & logopen{} }

func (l *logopen)CommandLogopenHelp()(ok bool){//col:33
/*CommandLogopenHelp()
{
    ShowMessages(".logopen : saves commands and results in a file.\n\n");
    ShowMessages("syntax : \t.logopen [FilePath (string)]\n");
}*/
return true
}

func (l *logopen)CommandLogopen()(ok bool){//col:110
/*CommandLogopen(vector<string> SplittedCommand, string Command)
{
    if (SplittedCommand.size() == 1)
    {
        ShowMessages("please specify a file\n");
        CommandLogopenHelp();
        return;
    }
    if (g_LogOpened)
    {
        ShowMessages("log was opened previously, you have the close it first "
                     "(using .logclose)\n");
        return;
    }
    Trim(Command);
    Command.erase(0, 8);
    Trim(Command);
    g_LogOpenFile.open(Command.c_str());
    if (g_LogOpenFile.is_open())
    {
        g_LogOpened = TRUE;
        time_t    t  = time(NULL);
        struct tm tm = *localtime(&t);
        ShowMessages("save commands and results into file : %s (%d-%02d-%02d "
                     "%02d:%02d:%02d)\n",
                     Command.c_str(),
                     tm.tm_year + 1900,
                     tm.tm_mon + 1,
                     tm.tm_mday,
                     tm.tm_hour,
                     tm.tm_min,
                     tm.tm_sec);
    }
    else
    {
        ShowMessages("unable to open file : %s\n", Command.c_str());
        return;
    }
}*/
return true
}

func (l *logopen)LogopenSaveToFile()(ok bool){//col:122
/*LogopenSaveToFile(const char * Text)
{
    g_LogOpenFile << Text;
}*/
return true
}


