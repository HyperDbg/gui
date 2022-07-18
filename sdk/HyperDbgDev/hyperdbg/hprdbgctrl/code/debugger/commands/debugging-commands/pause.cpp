#include "pch.h"
extern BOOLEAN                  g_BreakPrintingOutput;
extern BOOLEAN                  g_IsConnectedToRemoteDebuggee;
extern ACTIVE_DEBUGGING_PROCESS g_ActiveProcessDebuggingState;
VOID
CommandPauseHelp() {
    ShowMessages("pause : pauses the kernel events.\n\n");
    ShowMessages("syntax : \tpause \n");
}

VOID
CommandPauseRequest() {
    g_BreakPrintingOutput = TRUE;
    ShowMessages("pausing...\n");
    if (g_IsConnectedToRemoteDebuggee) {
        RemoteConnectionSendCommand("pause", strlen("pause") + 1);
    } else if (g_ActiveProcessDebuggingState.IsActive && UdPauseProcess(g_ActiveProcessDebuggingState.ProcessDebuggingToken)) {
        ShowMessages("please keep interacting with the process until all the "
                     "threads are intercepted and halted; whenever you execute "
                     "the first command, the thread interception will be stopped\n");
    }
}

VOID
CommandPause(vector<string> SplittedCommand, string Command) {
    if (SplittedCommand.size() != 1) {
        ShowMessages("incorrect use of 'pause'\n\n");
        CommandPauseHelp();
        return;
    }
    CommandPauseRequest();
}
