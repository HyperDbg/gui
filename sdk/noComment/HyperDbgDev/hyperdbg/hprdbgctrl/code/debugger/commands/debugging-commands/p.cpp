#include "pch.h"
extern BOOLEAN g_IsSerialConnectedToRemoteDebuggee;
extern BOOLEAN g_IsInstrumentingInstructions;
extern ACTIVE_DEBUGGING_PROCESS g_ActiveProcessDebuggingState;
VOID CommandPHelp() {
  ShowMessages(
      "p : executes a single instruction (step) and optionally displays the "
      "resulting values of all registers and flags.\n\n");
  ShowMessages("syntax : \tp\n");
  ShowMessages("syntax : \tp [Count (hex)]\n");
  ShowMessages("syntax : \tpr\n");
  ShowMessages("syntax : \tpr [Count (hex)]\n");
  ShowMessages("\n");
  ShowMessages("\t\te.g : p\n");
  ShowMessages("\t\te.g : pr\n");
  ShowMessages("\t\te.g : pr 1f\n");
}

VOID CommandP(vector<string> SplittedCommand, string Command) {
  UINT32 StepCount;
  DEBUGGER_REMOTE_STEPPING_REQUEST RequestFormat;
  if (SplittedCommand.size() != 1 && SplittedCommand.size() != 2) {
    ShowMessages("incorrect use of 'p'\n\n");
    CommandPHelp();
    return;
  }
  RequestFormat = DEBUGGER_REMOTE_STEPPING_REQUEST_STEP_OVER;
  if (SplittedCommand.size() == 2) {
    if (!ConvertStringToUInt32(SplittedCommand.at(1), &StepCount)) {
      ShowMessages("please specify a correct hex value for [count]\n\n");
      CommandPHelp();
      return;
    }
  } else {
    StepCount = 1;
  }
  if (g_IsSerialConnectedToRemoteDebuggee ||
      g_ActiveProcessDebuggingState.IsActive) {
    if (g_ActiveProcessDebuggingState.IsActive &&
        !g_ActiveProcessDebuggingState.IsPaused) {
      ShowMessages("the target process is running, use the "
                   "'pause' command or press CTRL+C to pause the process\n");
      return;
    }
    g_IsInstrumentingInstructions = TRUE;
    for (size_t i = 0; i < StepCount; i++) {
      if (g_IsSerialConnectedToRemoteDebuggee) {
        KdSendStepPacketToDebuggee(RequestFormat);
      } else {
        UdSendStepPacketToDebuggee(
            g_ActiveProcessDebuggingState.ProcessDebuggingToken,
            g_ActiveProcessDebuggingState.ThreadId, RequestFormat);
      }
      if (!SplittedCommand.at(0).compare("pr")) {
        ShowAllRegisters();
        if (i != StepCount - 1) {
          ShowMessages("\n");
        }
      }
      if (!g_IsInstrumentingInstructions) {
        break;
      }
    }
    g_IsInstrumentingInstructions = FALSE;
  } else {
    ShowMessages("err, stepping (p) is not valid in the current context, you "
                 "should connect to a debuggee\n");
  }
}

