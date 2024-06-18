#include "pch.h"
extern BOOLEAN g_IsSerialConnectedToRemoteDebuggee;
VOID CommandBcHelp(){
    ShowMessages("bc : clears a breakpoint using breakpoint id.\n\n");
    ShowMessages("syntax : \tbc [BreakpointId (hex)]\n");
    ShowMessages("\n");
    ShowMessages("\t\te.g : bc 0\n");
    ShowMessages("\t\te.g : bc 2\n");
}
VOID CommandBc(vector<string> SplitCommand, string Command){
    UINT64                            BreakpointId;
    DEBUGGEE_BP_LIST_OR_MODIFY_PACKET Request = {0};
    if (SplitCommand.size() != 2){
        ShowMessages("incorrect use of the 'bc'\n\n");
        CommandBcHelp();
        return;
    }
    if (!ConvertStringToUInt64(SplitCommand.at(1), &BreakpointId)){
        ShowMessages("please specify a correct hex value for breakpoint id\n\n");
        CommandBcHelp();
        return;
    }
    if (g_IsSerialConnectedToRemoteDebuggee){
        Request.Request = DEBUGGEE_BREAKPOINT_MODIFICATION_REQUEST_CLEAR;
        Request.BreakpointId = BreakpointId;
        KdSendListOrModifyPacketToDebuggee(&Request);
    }
    else{
        ShowMessages("err, clearing breakpoints is only valid if you connected to "
                     "a debuggee in debugger-mode\n");
    }
}
