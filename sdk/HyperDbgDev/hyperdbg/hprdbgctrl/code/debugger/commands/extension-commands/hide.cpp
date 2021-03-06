#include "pch.h"
extern UINT64                   g_CpuidAverage;
extern UINT64                   g_CpuidStandardDeviation;
extern UINT64                   g_CpuidMedian;
extern UINT64                   g_RdtscAverage;
extern UINT64                   g_RdtscStandardDeviation;
extern UINT64                   g_RdtscMedian;
extern BOOLEAN                  g_TransparentResultsMeasured;
extern ACTIVE_DEBUGGING_PROCESS g_ActiveProcessDebuggingState;
VOID
CommandHideHelp() {
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

VOID
CommandHide(vector<string> SplittedCommand, string Command) {
    BOOLEAN                                      Status;
    ULONG                                        ReturnedLength;
    UINT64                                       TargetPid;
    BOOLEAN                                      TrueIfProcessIdAndFalseIfProcessName;
    DEBUGGER_HIDE_AND_TRANSPARENT_DEBUGGER_MODE  HideRequest        = {0};
    PDEBUGGER_HIDE_AND_TRANSPARENT_DEBUGGER_MODE FinalRequestBuffer = 0;
    size_t                                       RequestBufferSize  = 0;
    if (SplittedCommand.size() <= 2 && SplittedCommand.size() != 1) {
        ShowMessages("incorrect use of '!hide'\n\n");
        CommandHideHelp();
        return;
    }
    if (SplittedCommand.size() == 1) {
        if (g_ActiveProcessDebuggingState.IsActive) {
            TrueIfProcessIdAndFalseIfProcessName = TRUE;
            TargetPid                            = g_ActiveProcessDebuggingState.ProcessId;
        } else {
            ShowMessages("you're not attached to any user-mode process, "
                         "please explicitly specify the process id or process name\n");
            return;
        }
    } else if (!SplittedCommand.at(1).compare("pid")) {
        TrueIfProcessIdAndFalseIfProcessName = TRUE;
        if (SplittedCommand.size() != 3) {
            ShowMessages("incorrect use of '!hide'\n\n");
            CommandHideHelp();
            return;
        }
        if (!ConvertStringToUInt64(SplittedCommand.at(2), &TargetPid)) {
            ShowMessages("incorrect process id\n\n");
            return;
        }
    } else if (!SplittedCommand.at(1).compare("name")) {
        TrueIfProcessIdAndFalseIfProcessName = FALSE;
        Trim(Command);
        Command.erase(0, 5);
        Command.erase(0, 4 + 1);
        Trim(Command);
    } else {
        ShowMessages("incorrect use of '!hide'\n\n");
        CommandHideHelp();
        return;
    }
    if (!g_TransparentResultsMeasured || !g_CpuidAverage ||
        !g_CpuidStandardDeviation || !g_CpuidMedian) {
        ShowMessages("the average, median and standard deviation is not measured. "
                     "Did you use '!measure' command?\n");
        return;
    }
    AssertShowMessageReturnStmt(g_DeviceHandle, ASSERT_MESSAGE_DRIVER_NOT_LOADED, AssertReturn);
    HideRequest.IsHide                 = TRUE;
    HideRequest.CpuidAverage           = g_CpuidAverage;
    HideRequest.CpuidMedian            = g_CpuidMedian;
    HideRequest.CpuidStandardDeviation = g_CpuidStandardDeviation;
    HideRequest.RdtscAverage           = g_RdtscAverage;
    HideRequest.RdtscMedian            = g_RdtscMedian;
    HideRequest.RdtscStandardDeviation = g_RdtscStandardDeviation;
    HideRequest.TrueIfProcessIdAndFalseIfProcessName =
        TrueIfProcessIdAndFalseIfProcessName;
    if (TrueIfProcessIdAndFalseIfProcessName) {
        HideRequest.ProcId = TargetPid;
        RequestBufferSize  = sizeof(DEBUGGER_HIDE_AND_TRANSPARENT_DEBUGGER_MODE);
    } else {
        HideRequest.LengthOfProcessName = Command.size() + 1;
        RequestBufferSize               = sizeof(DEBUGGER_HIDE_AND_TRANSPARENT_DEBUGGER_MODE) +
                            Command.size() + 1;
    }
    FinalRequestBuffer =
        (PDEBUGGER_HIDE_AND_TRANSPARENT_DEBUGGER_MODE)malloc(RequestBufferSize);
    RtlZeroMemory(FinalRequestBuffer, RequestBufferSize);
    memcpy(FinalRequestBuffer, &HideRequest, sizeof(DEBUGGER_HIDE_AND_TRANSPARENT_DEBUGGER_MODE));
    memcpy(((UINT64 *)((UINT64)FinalRequestBuffer +
                       sizeof(DEBUGGER_HIDE_AND_TRANSPARENT_DEBUGGER_MODE))),
           Command.c_str(),
           Command.size());
    Status = DeviceIoControl(
        g_DeviceHandle,                                             // Handle to device
        IOCTL_DEBUGGER_HIDE_AND_UNHIDE_TO_TRANSPARENT_THE_DEBUGGER, // IO Control
        FinalRequestBuffer,                                         // Input Buffer to driver.
        RequestBufferSize,                                          // Input buffer length
        FinalRequestBuffer,                                         // Output Buffer from driver.
        SIZEOF_DEBUGGER_HIDE_AND_TRANSPARENT_DEBUGGER_MODE,         // Length of output
        &ReturnedLength,                                            // Bytes placed in buffer.
        NULL                                                        // synchronous call
    );
    if (!Status) {
        ShowMessages("ioctl failed with code 0x%x\n", GetLastError());
        free(FinalRequestBuffer);
        return;
    }
    if (FinalRequestBuffer->KernelStatus == DEBUGGER_OPERATION_WAS_SUCCESSFULL) {
        ShowMessages("transparent debugging successfully enabled :)\n");
    } else if (FinalRequestBuffer->KernelStatus ==
               DEBUGGER_ERROR_UNABLE_TO_HIDE_OR_UNHIDE_DEBUGGER) {
        ShowMessages("unable to hide the debugger (transparent-debugging) :(\n");
        free(FinalRequestBuffer);
        return;
    } else {
        ShowMessages("unknown error occurred :(\n");
        free(FinalRequestBuffer);
        return;
    }
    free(FinalRequestBuffer);
}
