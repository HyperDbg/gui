#include "pch.h"
extern HANDLE g_DeviceHandle;
VOID
CommandExitHelp() {
    ShowMessages(
        "exit : unloads and uninstalls the drivers and closes the debugger.\n\n");
    ShowMessages("syntax : \texit\n");
}

VOID
CommandExit(vector<string> SplittedCommand, string Command) {
    if (SplittedCommand.size() != 1) {
        ShowMessages("incorrect use of 'exit'\n\n");
        CommandExitHelp();
        return;
    }
    if (g_DeviceHandle) {
        HyperDbgUnloadVmm();
    }
    exit(0);
}
