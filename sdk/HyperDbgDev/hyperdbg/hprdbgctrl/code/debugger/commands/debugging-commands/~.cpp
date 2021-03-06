#include "pch.h"
extern ULONG   g_CurrentRemoteCore;
extern BOOLEAN g_IsSerialConnectedToRemoteDebuggee;
VOID
CommandCoreHelp() {
    ShowMessages("~ : shows and changes the operating processor.\n\n");
    ShowMessages("syntax : \t~\n");
    ShowMessages("syntax : \t~ [CoreNumber (hex)]\n");
    ShowMessages("\n");
    ShowMessages("\t\te.g : ~ \n");
    ShowMessages("\t\te.g : ~ 2 \n");
}

VOID
CommandCore(vector<string> SplittedCommand, string Command) {
    UINT32 TargetCore = 0;
    if (SplittedCommand.size() != 1 && SplittedCommand.size() != 2) {
        ShowMessages("incorrect use of '~'\n\n");
        CommandCoreHelp();
        return;
    }
    if (!g_IsSerialConnectedToRemoteDebuggee) {
        ShowMessages("err, you're not connected to any debuggee\n");
        return;
    }
    if (SplittedCommand.size() == 1) {
        ShowMessages("current processor : 0x%x\n", g_CurrentRemoteCore);
    } else if (SplittedCommand.size() == 2) {
        if (!ConvertStringToUInt32(SplittedCommand.at(1), &TargetCore)) {
            ShowMessages("please specify a correct hex value for the core that you "
                         "want to operate on it\n\n");
            CommandCoreHelp();
            return;
        }
        KdSendSwitchCorePacketToDebuggee(TargetCore);
    }
}
