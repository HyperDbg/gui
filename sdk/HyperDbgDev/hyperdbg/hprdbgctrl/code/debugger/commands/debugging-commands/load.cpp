#include "pch.h"
extern HANDLE  g_IsDriverLoadedSuccessfully;
extern HANDLE  g_DeviceHandle;
extern BOOLEAN g_IsConnectedToHyperDbgLocally;
extern BOOLEAN g_IsDebuggerModulesLoaded;
VOID
CommandLoadHelp() {
    ShowMessages("load : installs the drivers and load the modules.\n\n");
    ShowMessages("syntax : \tload [ModuleName (string)]\n");
    ShowMessages("\n");
    ShowMessages("\t\te.g : load vmm\n");
}

BOOLEAN
CommandLoadVmmModule() {
    BOOL   Status;
    HANDLE hToken;
    Status =
        OpenProcessToken(GetCurrentProcess(), TOKEN_ADJUST_PRIVILEGES, &hToken);
    if (!Status) {
        ShowMessages("err, OpenProcessToken failed (%x)\n", GetLastError());
        return FALSE;
    }
    Status = SetPrivilege(hToken, SE_DEBUG_NAME, TRUE);
    if (!Status) {
        CloseHandle(hToken);
        return FALSE;
    }
    if (HyperDbgInstallVmmDriver() == 1) {
        return FALSE;
    }
    g_IsDriverLoadedSuccessfully = CreateEvent(NULL, FALSE, FALSE, NULL);
    if (HyperDbgLoadVmm() == 1) {
        CloseHandle(g_IsDriverLoadedSuccessfully);
        return FALSE;
    }
    WaitForSingleObject(
        g_IsDriverLoadedSuccessfully,
        INFINITE);
    CloseHandle(g_IsDriverLoadedSuccessfully);
    g_IsDebuggerModulesLoaded = TRUE;
    ShowMessages("vmm module is running...\n");
    return TRUE;
}

VOID
CommandLoad(vector<string> SplittedCommand, string Command) {
    if (SplittedCommand.size() != 2) {
        ShowMessages("incorrect use of 'load'\n\n");
        CommandLoadHelp();
        return;
    }
    if (!g_IsConnectedToHyperDbgLocally) {
        ShowMessages("you're not connected to any instance of HyperDbg, did you "
                     "use '.connect'? \n");
        return;
    }
    if (!SplittedCommand.at(1).compare("vmm")) {
        if (g_DeviceHandle) {
            ShowMessages("handle of the driver found, if you use 'load' before, please "
                         "first unload it then call 'unload'\n");
            return;
        }
        ShowMessages("loading the vmm driver\n");
        if (!CommandLoadVmmModule()) {
            ShowMessages("failed to install or load the driver\n");
            return;
        }
        SymbolLocalReload(GetCurrentProcessId());
    } else {
        ShowMessages("module not found, currently 'vmm' is the only available "
                     "module for HyperDbg\n");
    }
}
