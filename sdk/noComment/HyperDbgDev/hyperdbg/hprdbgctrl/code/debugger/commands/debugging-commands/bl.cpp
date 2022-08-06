#include "pch.h"
extern BOOLEAN g_IsSerialConnectedToRemoteDebuggee;
VOID
CommandBlHelp()
{
    ShowMessages("bl : lists all the enabled and disabled breakpoints.\n\n");
    ShowMessages("syntax : \tbl\n");
}

VOID
CommandBl(vector<string> SplittedCommand, string Command)
{
    DEBUGGEE_BP_LIST_OR_MODIFY_PACKET Request = {0};
    if (SplittedCommand.size() != 1)
    {
        ShowMessages("incorrect use of 'bl'\n\n");
        CommandBlHelp();
        return;
    }
    if (g_IsSerialConnectedToRemoteDebuggee)
    {
        Request.Request = DEBUGGEE_BREAKPOINT_MODIFICATION_REQUEST_LIST_BREAKPOINTS;
        KdSendListOrModifyPacketToDebuggee(&Request);
        ShowMessages("\n");
    }
    else
    {
        ShowMessages("err, listing breakpoints is only valid if you connected to "
                     "a debuggee in debugger-mode\n");
    }
}

