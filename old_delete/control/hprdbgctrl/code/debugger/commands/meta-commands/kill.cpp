#include "pch.h"
extern UINT32                   g_ProcessIdOfLatestStartingProcess;
extern ACTIVE_DEBUGGING_PROCESS g_ActiveProcessDebuggingState;
VOID CommandKillHelp(){
    ShowMessages(".kill : terminates the current running process.\n\n");
    ShowMessages("syntax : \t.kill \n");
}
VOID CommandKill(vector<string> SplitCommand, string Command){
    if (SplitCommand.size() != 1){
        ShowMessages("incorrect use of the '.kill'\n\n");
        CommandKillHelp();
        return;
    }
    if (g_ActiveProcessDebuggingState.IsActive){
        if (!UdKillProcess(g_ActiveProcessDebuggingState.ProcessId)){
            ShowMessages("process does not exists, is it already terminated?\n");
        }
    }
    else if (g_ProcessIdOfLatestStartingProcess != NULL){
        if (!UdKillProcess(g_ProcessIdOfLatestStartingProcess)){
            ShowMessages("process does not exists, is it already terminated?\n");
        }
        g_ProcessIdOfLatestStartingProcess = NULL;
    }
    else{
        ShowMessages("nothing to terminate!\n");
        return;
    }
}
