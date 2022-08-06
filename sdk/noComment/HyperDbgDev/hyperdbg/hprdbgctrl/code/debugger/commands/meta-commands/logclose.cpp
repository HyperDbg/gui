#include "pch.h"
extern BOOLEAN  g_LogOpened;
extern ofstream g_LogOpenFile;
VOID
CommandLogcloseHelp()
{
    ShowMessages(".logclose : closes the previously opened log.\n\n");
    ShowMessages("syntax : \t.logclose\n");
}

VOID
CommandLogclose(vector<string> SplittedCommand, string Command)
{
    if (SplittedCommand.size() != 1)
    {
        ShowMessages("incorrect use of '.logclose'\n\n");
        CommandLogcloseHelp();
        return;
    }
    if (!g_LogOpened)
    {
        ShowMessages("there is no opened log, did you use '.logopen'? \n");
        return;
    }
    time_t    t  = time(NULL);
    struct tm tm = *localtime(&t);
    ShowMessages("log file closed (%d-%02d-%02d "
                 "%02d:%02d:%02d)\n",
                 tm.tm_year + 1900,
                 tm.tm_mon + 1,
                 tm.tm_mday,
                 tm.tm_hour,
                 tm.tm_min,
                 tm.tm_sec);
    g_LogOpenFile.close();
    g_LogOpened = FALSE;
}

