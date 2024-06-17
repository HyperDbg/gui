#include "pch.h"
extern UINT64 g_CpuidAverage;
extern UINT64 g_CpuidStandardDeviation;
extern UINT64 g_CpuidMedian;
extern UINT64 g_RdtscAverage;
extern UINT64 g_RdtscStandardDeviation;
extern UINT64 g_RdtscMedian;
extern BOOLEAN                  g_TransparentResultsMeasured;
extern ACTIVE_DEBUGGING_PROCESS g_ActiveProcessDebuggingState;
VOID CommandHideHelp(){
    ShowMessages("!hide : tries to make HyperDbg transparent from anti-debugging "
                 "and anti-hypervisor methods.\n\n");
    ShowMessages("syntax : \t!hide\n");
    ShowMessages("syntax : \t!hide [pid ProcessId (hex)]\n");
    ShowMessages("syntax : \t!hide [name ProcessName (string)]\n");
    ShowMessages("note : \tprocess names are case sensitive and you can use "
                 "this command multiple times.\n");
    ShowMessages("\n");
    ShowMessages("\t\te.g : !hide\n");
    ShowMessages("\t\te.g : !hide pid b60 \n");
    ShowMessages("\t\te.g : !hide name procexp.exe\n");
}
VOID CommandHide(vector<string> SplitCommand, string Command){
    BOOLEAN                                      Status;
    ULONG                                        ReturnedLength;
    UINT64                                       TargetPid;
    BOOLEAN                                      TrueIfProcessIdAndFalseIfProcessName;
    DEBUGGER_HIDE_AND_TRANSPARENT_DEBUGGER_MODE  HideRequest        = {0};
    PDEBUGGER_HIDE_AND_TRANSPARENT_DEBUGGER_MODE FinalRequestBuffer = 0;
    size_t                                       RequestBufferSize  = 0;
    if (SplitCommand.size() <= 2 && SplitCommand.size() != 1){
        ShowMessages("incorrect use of the '!hide'\n\n");
        CommandHideHelp();
        return;
    }
    if (SplitCommand.size() == 1){
        if (g_ActiveProcessDebuggingState.IsActive){
            TrueIfProcessIdAndFalseIfProcessName = TRUE;
            TargetPid                            = g_ActiveProcessDebuggingState.ProcessId;
        }
        else{
            ShowMessages("you're not attached to any user-mode process, "
                         "please explicitly specify the process id or process name\n");
            return;
        }
    }
    else if (!SplitCommand.at(1).compare("pid")){
        TrueIfProcessIdAndFalseIfProcessName = TRUE;
        if (SplitCommand.size() != 3){
            ShowMessages("incorrect use of the '!hide'\n\n");
            CommandHideHelp();
            return;
        }
        if (!ConvertStringToUInt64(SplitCommand.at(2), &TargetPid)){
            ShowMessages("incorrect process id\n\n");
            return;
        }
    }
    else if (!SplitCommand.at(1).compare("name")){
        TrueIfProcessIdAndFalseIfProcessName = FALSE;
        Trim(Command);
        Command.erase(0, SplitCommand.at(0).size());
        Command.erase(0, 4 + 1);
        Trim(Command);
    }
    else{
        ShowMessages("incorrect use of the '!hide'\n\n");
        CommandHideHelp();
        return;
    }
    if (!g_TransparentResultsMeasured || !g_CpuidAverage ||
        !g_CpuidStandardDeviation || !g_CpuidMedian){
        ShowMessages("the average, median and standard deviation is not measured. "
                     "Did you use '!measure' command?\n");
        return;
    }
    AssertShowMessageReturnStmt(g_DeviceHandle, ASSERT_MESSAGE_DRIVER_NOT_LOADED, AssertReturn);
    HideRequest.IsHide = TRUE;
    HideRequest.CpuidAverage           = g_CpuidAverage;
    HideRequest.CpuidMedian            = g_CpuidMedian;
    HideRequest.CpuidStandardDeviation = g_CpuidStandardDeviation;
    HideRequest.RdtscAverage           = g_RdtscAverage;
    HideRequest.RdtscMedian            = g_RdtscMedian;
    HideRequest.RdtscStandardDeviation = g_RdtscStandardDeviation;
    HideRequest.TrueIfProcessIdAndFalseIfProcessName =
        TrueIfProcessIdAndFalseIfProcessName;
    if (TrueIfProcessIdAndFalseIfProcessName){
        HideRequest.ProcId = (UINT32)TargetPid;
        RequestBufferSize = sizeof(DEBUGGER_HIDE_AND_TRANSPARENT_DEBUGGER_MODE);
    }
    else{
        HideRequest.LengthOfProcessName = (UINT32)Command.size() + 1;
        RequestBufferSize               = sizeof(DEBUGGER_HIDE_AND_TRANSPARENT_DEBUGGER_MODE) + Command.size() + 1;
    }
    FinalRequestBuffer =
        (PDEBUGGER_HIDE_AND_TRANSPARENT_DEBUGGER_MODE)malloc(RequestBufferSize);
    if (FinalRequestBuffer == NULL){
        ShowMessages("insufficient space\n");
        return;
    }
    RtlZeroMemory(FinalRequestBuffer, RequestBufferSize);
    memcpy(FinalRequestBuffer, &HideRequest, sizeof(DEBUGGER_HIDE_AND_TRANSPARENT_DEBUGGER_MODE));
    memcpy(((UINT64 *)((UINT64)FinalRequestBuffer +
                       sizeof(DEBUGGER_HIDE_AND_TRANSPARENT_DEBUGGER_MODE))),
           Command.c_str(),
           Command.size());
    Status = DeviceIoControl(
        g_DeviceHandle,                                             
        IOCTL_DEBUGGER_HIDE_AND_UNHIDE_TO_TRANSPARENT_THE_DEBUGGER, 
        FinalRequestBuffer,                                         
        (DWORD)RequestBufferSize,                                   
        FinalRequestBuffer,                                         
        SIZEOF_DEBUGGER_HIDE_AND_TRANSPARENT_DEBUGGER_MODE,         
        &ReturnedLength,                                            
        NULL                                                        
    );
    if (!Status){
        ShowMessages("ioctl failed with code 0x%x\n", GetLastError());
        free(FinalRequestBuffer);
        return;
    }
    if (FinalRequestBuffer->KernelStatus == DEBUGGER_OPERATION_WAS_SUCCESSFUL){
        ShowMessages("transparent debugging successfully enabled :)\n");
    }
    else if (FinalRequestBuffer->KernelStatus ==
             DEBUGGER_ERROR_UNABLE_TO_HIDE_OR_UNHIDE_DEBUGGER){
        ShowMessages("unable to hide the debugger (transparent-debugging) :(\n");
        free(FinalRequestBuffer);
        return;
    }
    else{
        ShowMessages("unknown error occurred :(\n");
        free(FinalRequestBuffer);
        return;
    }
    free(FinalRequestBuffer);
}
