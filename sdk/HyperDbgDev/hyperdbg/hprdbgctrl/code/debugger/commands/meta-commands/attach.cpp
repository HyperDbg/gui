#include "pch.h"
extern BOOLEAN g_IsSerialConnectedToRemoteDebuggee;
extern BOOLEAN g_IsSerialConnectedToRemoteDebugger;
VOID
CommandAttachHelp() {
    ShowMessages(".attach : attaches to debug a thread in VMI Mode.\n\n");
    ShowMessages("syntax : \t.attach [pid ProcessId (hex)]\n");
    ShowMessages("\n");
    ShowMessages("\t\te.g : .attach pid b60 \n");
}

VOID
CommandAttach(vector<string> SplittedCommand, string Command) {
    UINT64  TargetPid = 0;
    BOOLEAN NextIsPid = FALSE;
#if ActivateUserModeDebugger == FALSE
    if (!g_IsSerialConnectedToRemoteDebugger) {
        ShowMessages("The user-mode debugger is still in the beta version and not stable. "
                     "We decided to exclude it from this release and release it in future versions. "
                     "If you want to test the user-mode debugger in VMI Mode, you should build "
                     "HyperDbg with special instructions. \nPlease follow the steps here: "
                     "https://docs.hyperdbg.org/getting-started/build-and-install \n");
        return;
    }
#endif // !ActivateUserModeDebugger
    if (SplittedCommand.size() >= 4) {
        ShowMessages("incorrect use of '.attach'\n\n");
        CommandAttachHelp();
        return;
    }
    if (g_IsSerialConnectedToRemoteDebuggee) {
        ShowMessages("err, '.attach', and '.detach' commands are only usable "
                     "in VMI Mode, you can use the '.process', or the '.thread' "
                     "in Debugger Mode\n");
        return;
    }
    for (auto item : SplittedCommand) {
        if (NextIsPid) {
            NextIsPid = FALSE;
            if (!ConvertStringToUInt64(item, &TargetPid)) {
                ShowMessages("please specify a correct hex value for process id\n\n");
                CommandAttachHelp();
                return;
            }
        } else if (!item.compare("pid")) {
            NextIsPid = TRUE;
        }
    }
    if (TargetPid == 0) {
        ShowMessages("please specify a hex value for process id\n\n");
        CommandAttachHelp();
        return;
    }
    UdAttachToProcess(TargetPid, NULL, NULL);
}
