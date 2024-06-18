#include "pch.h"
extern HANDLE  g_DeviceHandle;
extern BOOLEAN g_IsConnectedToHyperDbgLocally;
extern BOOLEAN g_IsSerialConnectedToRemoteDebuggee;
VOID CommandExitHelp(){
    ShowMessages(
        "exit : unloads and uninstalls the drivers and closes the debugger.\n\n");
    ShowMessages("syntax : \texit\n");
}
VOID CommandExit(vector<string> SplitCommand, string Command){
    if (SplitCommand.size() != 1){
        ShowMessages("incorrect use of the 'exit'\n\n");
        CommandExitHelp();
        return;
    }
    if (g_IsConnectedToHyperDbgLocally){
        if (g_DeviceHandle){
            HyperDbgUnloadVmm();
        }
    }
    else if (g_IsSerialConnectedToRemoteDebuggee){
        KdCloseConnection();
    }
    exit(0);
}
