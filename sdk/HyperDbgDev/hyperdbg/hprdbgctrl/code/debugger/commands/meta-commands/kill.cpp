#include "pch.h"
extern ACTIVE_DEBUGGING_PROCESS g_ActiveProcessDebuggingState;
VOID
CommandKillHelp() {
    ShowMessages(".kill : terminates the current running process.\n\n");
    ShowMessages("syntax : \t.kill \n");
}

VOID
CommandKill(vector<string> SplittedCommand, string Command) {
    if (SplittedCommand.size() != 1) {
        ShowMessages("incorrect use of '.kill'\n\n");
        CommandKillHelp();
        return;
    }
    if (!g_ActiveProcessDebuggingState.IsActive) {
        ShowMessages("nothing to terminate!\n");
        return;
    }
    UdKillProcess(g_ActiveProcessDebuggingState.ProcessId);
}
