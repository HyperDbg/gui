#include "pch.h"
extern BOOLEAN g_IsSerialConnectedToRemoteDebuggee;
extern BOOLEAN g_IsSerialConnectedToRemoteDebugger;
VOID CommandAttachHelp(){
    ShowMessages(".attach : attaches to debug a thread in VMI Mode.\n\n");
    ShowMessages("syntax : \t.attach [pid ProcessId (hex)]\n");
    ShowMessages("\n");
    ShowMessages("\t\te.g : .attach pid b60 \n");
}
VOID CommandAttach(vector<string> SplitCommand, string Command){
    UINT32  TargetPid = 0;
    BOOLEAN NextIsPid = FALSE;
#if ActivateUserModeDebugger == FALSE
    if (!g_IsSerialConnectedToRemoteDebugger){
        ShowMessages("the user-mode debugger in VMI Mode is still in the beta version and not stable. "
                     "we decided to exclude it from this release and release it in future versions. "
                     "if you want to test the user-mode debugger in VMI Mode, you should build "
                     "HyperDbg with special instructions. But attaching/switching to other processes\n"
                     "are fully supported in the Debugger Mode.\n"
                     "(it's not recommended to use it in VMI Mode yet!)\n");
        return;
    }
#endif 
    if (SplitCommand.size() >= 4){
        ShowMessages("incorrect use of the '.attach'\n\n");
        CommandAttachHelp();
        return;
    }
    if (g_IsSerialConnectedToRemoteDebuggee){
        ShowMessages("err, '.attach', and '.detach' commands are only usable "
                     "in VMI Mode, you can use the '.process', or the '.thread' "
                     "in Debugger Mode\n");
        return;
    }
    for (auto item : SplitCommand){
        if (NextIsPid){
            NextIsPid = FALSE;
            if (!ConvertStringToUInt32(item, &TargetPid)){
                ShowMessages("please specify a correct hex value for process id\n\n");
                CommandAttachHelp();
                return;
            }
        }
        else if (!item.compare("pid")){
            NextIsPid = TRUE;
        }
    }
    if (TargetPid == 0){
        ShowMessages("please specify a hex value for process id\n\n");
        CommandAttachHelp();
        return;
    }
    UdAttachToProcess(TargetPid, NULL, NULL, FALSE);
}
