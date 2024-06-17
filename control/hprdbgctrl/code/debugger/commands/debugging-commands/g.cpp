#include "pch.h"
extern BOOLEAN                  g_BreakPrintingOutput;
extern BOOLEAN                  g_IsConnectedToRemoteDebuggee;
extern BOOLEAN                  g_IsSerialConnectedToRemoteDebuggee;
extern ACTIVE_DEBUGGING_PROCESS g_ActiveProcessDebuggingState;
VOID CommandGHelp(){
    ShowMessages("g : continues debuggee or continues processing kernel messages.\n\n");
    ShowMessages("syntax : \tg \n");
}
VOID CommandGRequest(){
    if (g_IsSerialConnectedToRemoteDebuggee){
        KdBreakControlCheckAndContinueDebugger();
    }
    else{
        g_BreakPrintingOutput = FALSE;
        if (g_IsConnectedToRemoteDebuggee){
            RemoteConnectionSendCommand("g", (UINT32)strlen("g") + 1);
        }
        else if (g_ActiveProcessDebuggingState.IsActive){
            if (g_ActiveProcessDebuggingState.IsPaused){
                UdContinueDebuggee(g_ActiveProcessDebuggingState.ProcessDebuggingToken);
                g_ActiveProcessDebuggingState.IsPaused = FALSE;
            }
            else{
                ShowMessages("err, target process is already running\n");
            }
        }
    }
}
VOID CommandG(vector<string> SplitCommand, string Command){
    if (SplitCommand.size() != 1){
        ShowMessages("incorrect use of the 'g'\n\n");
        CommandGHelp();
        return;
    }
    CommandGRequest();
}
