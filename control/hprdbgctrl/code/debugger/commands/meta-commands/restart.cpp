#include "pch.h"
extern UINT32                   g_ProcessIdOfLatestStartingProcess;
extern std::wstring             g_StartCommandPath;
extern std::wstring             g_StartCommandPathAndArguments;
extern ACTIVE_DEBUGGING_PROCESS g_ActiveProcessDebuggingState;
;
VOID CommandRestartHelp(){
    ShowMessages(".restart : restarts the previously started process "
                 "(using '.start' command).\n\n");
    ShowMessages(
        "syntax : \t.restart \n");
}
VOID CommandRestart(vector<string> SplitCommand, string Command){
    if (SplitCommand.size() != 1){
        ShowMessages("incorrect use of the '.restart'\n\n");
        CommandRestartHelp();
        return;
    }
    if (g_StartCommandPath.empty()){
        ShowMessages("nothing to restart, did you use the '.start' command before?\n");
        return;
    }
    if (g_ActiveProcessDebuggingState.IsActive){
        UdKillProcess(g_ActiveProcessDebuggingState.ProcessId);
    }
    else if (g_ProcessIdOfLatestStartingProcess != NULL){
        UdKillProcess(g_ProcessIdOfLatestStartingProcess);
        g_ProcessIdOfLatestStartingProcess = NULL;
    }
    if (g_StartCommandPathAndArguments.empty()){
        UdAttachToProcess(NULL,
                          g_StartCommandPath.c_str(),
                          NULL,
                          FALSE);
    }
    else{
        UdAttachToProcess(NULL,
                          g_StartCommandPath.c_str(),
                          (WCHAR *)g_StartCommandPathAndArguments.c_str(),
                          FALSE);
    }
}
