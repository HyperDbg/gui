#include "pch.h"
extern std::wstring             g_StartCommandPath;
extern std::wstring             g_StartCommandPathAndArguments;
extern ACTIVE_DEBUGGING_PROCESS g_ActiveProcessDebuggingState;
VOID
CommandRestartHelp() {
    ShowMessages(".restart : restarts the previously started process "
                 "(using '.start' command).\n\n");
    ShowMessages(
        "syntax : \t.restart \n");
}

VOID
CommandRestart(vector<string> SplittedCommand, string Command) {
    if (SplittedCommand.size() != 1) {
        ShowMessages("incorrect use of '.restart'\n\n");
        CommandRestartHelp();
        return;
    }
    if (g_StartCommandPath.empty()) {
        ShowMessages("nothing to restart, did you use the '.start' command before?\n");
        return;
    }
    if (g_ActiveProcessDebuggingState.IsActive) {
        UdKillProcess(g_ActiveProcessDebuggingState.ProcessId);
    }
    if (g_StartCommandPathAndArguments.empty()) {
        UdAttachToProcess(NULL,
                          g_StartCommandPath.c_str(),
                          NULL);
    } else {
        UdAttachToProcess(NULL,
                          g_StartCommandPath.c_str(),
                          (WCHAR *)g_StartCommandPathAndArguments.c_str());
    }
}
