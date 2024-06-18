#include "pch.h"
extern BOOLEAN                  g_IsSerialConnectedToRemoteDebuggee;
extern BOOLEAN                  g_IsInstrumentingInstructions;
extern ACTIVE_DEBUGGING_PROCESS g_ActiveProcessDebuggingState;
extern BYTE                     g_CurrentRunningInstruction[MAXIMUM_INSTR_SIZE];
extern BOOLEAN                  g_IsRunningInstruction32Bit;
VOID CommandGuHelp(){
    ShowMessages(
        "gu : executes a single instruction (step-out) and optionally displays the "
        "resulting values of all registers and flags.\n\n");
    ShowMessages("syntax : \tgu\n");
    ShowMessages("syntax : \tgu [Count (hex)]\n");
    ShowMessages("\n");
    ShowMessages("\t\te.g : gu\n");
    ShowMessages("\t\te.g : gu 10000\n");
}
VOID CommandGu(vector<string> SplitCommand, string Command){
    UINT32                           StepCount;
    DEBUGGER_REMOTE_STEPPING_REQUEST RequestFormat;
    BOOLEAN                          BreakOnNextInstruction = FALSE;
    if (SplitCommand.size() != 1 && SplitCommand.size() != 2){
        ShowMessages("incorrect use of the 'gu'\n\n");
        CommandGuHelp();
        return;
    }
    RequestFormat = DEBUGGER_REMOTE_STEPPING_REQUEST_STEP_OVER_FOR_GU;
    if (SplitCommand.size() == 2){
        if (!ConvertStringToUInt32(SplitCommand.at(1), &StepCount)){
            ShowMessages("please specify a correct hex value for [count]\n\n");
            CommandGuHelp();
            return;
        }
    }
    else{
        StepCount = DEBUGGER_REMOTE_TRACKING_DEFAULT_COUNT_OF_STEPPING;
    }
    if (g_IsSerialConnectedToRemoteDebuggee || g_ActiveProcessDebuggingState.IsActive){
        if (g_ActiveProcessDebuggingState.IsActive && !g_ActiveProcessDebuggingState.IsPaused){
            ShowMessages("the target process is running, use the "
                         "'pause' command or press CTRL+C to pause the process\n");
            return;
        }
        g_IsInstrumentingInstructions = TRUE;
        for (size_t i = 0; i < StepCount; i++){
            if (HyperDbgCheckWhetherTheCurrentInstructionIsRet(
                    g_CurrentRunningInstruction,
                    MAXIMUM_INSTR_SIZE,
                    g_IsRunningInstruction32Bit ? FALSE : TRUE 
                    )){
                BreakOnNextInstruction = TRUE;
                RequestFormat = DEBUGGER_REMOTE_STEPPING_REQUEST_STEP_OVER_FOR_GU_LAST_INSTRUCTION;
            }
            if (g_IsSerialConnectedToRemoteDebuggee){
                KdSendStepPacketToDebuggee(RequestFormat);
            }
            else{
                UdSendStepPacketToDebuggee(g_ActiveProcessDebuggingState.ProcessDebuggingToken,
                                           g_ActiveProcessDebuggingState.ThreadId,
                                           RequestFormat);
            }
            if (!g_IsInstrumentingInstructions){
                break;
            }
            if (BreakOnNextInstruction){
                break;
            }
        }
        g_IsInstrumentingInstructions = FALSE;
    }
    else{
        ShowMessages("err, going up (gu) is not valid in the current context, you "
                     "should connect to a debuggee\n");
    }
}
