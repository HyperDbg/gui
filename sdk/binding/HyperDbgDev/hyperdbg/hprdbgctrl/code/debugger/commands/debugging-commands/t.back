










#include "pch.h"




extern BOOLEAN                  g_IsSerialConnectedToRemoteDebuggee;
extern BOOLEAN                  g_IsInstrumentingInstructions;
extern ACTIVE_DEBUGGING_PROCESS g_ActiveProcessDebuggingState;






VOID
CommandTHelp()
{
    ShowMessages(
        "t : executes a single instruction (step-in) and optionally displays the "
        "resulting values of all registers and flags.\n\n");

    ShowMessages("syntax : \tt\n");
    ShowMessages("syntax : \tt [Count (hex)]\n");
    ShowMessages("syntax : \ttr\n");
    ShowMessages("syntax : \ttr [Count (hex)]\n");

    ShowMessages("\n");
    ShowMessages("\t\te.g : t\n");
    ShowMessages("\t\te.g : tr\n");
    ShowMessages("\t\te.g : tr 1f\n");
}








VOID
CommandT(vector<string> SplittedCommand, string Command)
{
    UINT32                           StepCount;
    DEBUGGER_REMOTE_STEPPING_REQUEST RequestFormat;

    
    
    
    if (SplittedCommand.size() != 1 && SplittedCommand.size() != 2)
    {
        ShowMessages("incorrect use of 't'\n\n");
        CommandTHelp();
        return;
    }

    
    
    
    RequestFormat = DEBUGGER_REMOTE_STEPPING_REQUEST_STEP_IN;

    
    
    
    if (SplittedCommand.size() == 2)
    {
        if (!ConvertStringToUInt32(SplittedCommand.at(1), &StepCount))
        {
            ShowMessages("please specify a correct hex value for [count]\n\n");
            CommandTHelp();
            return;
        }
    }
    else
    {
        StepCount = 1;
    }

    
    
    
    if (g_IsSerialConnectedToRemoteDebuggee || g_ActiveProcessDebuggingState.IsActive)
    {
        
        
        
        if (g_ActiveProcessDebuggingState.IsActive && !g_ActiveProcessDebuggingState.IsPaused)
        {
            ShowMessages("the target process is running, use the "
                         "'pause' command or press CTRL+C to pause the process\n");
            return;
        }

        
        
        
        g_IsInstrumentingInstructions = TRUE;

        for (size_t i = 0; i < StepCount; i++)
        {
            
            
            
            
            
            

            if (g_IsSerialConnectedToRemoteDebuggee)
            {
                
                
                
                KdSendStepPacketToDebuggee(RequestFormat);
            }
            else
            {
                
                
                
                UdSendStepPacketToDebuggee(g_ActiveProcessDebuggingState.ProcessDebuggingToken,
                                           g_ActiveProcessDebuggingState.ThreadId,
                                           RequestFormat);
            }

            if (!SplittedCommand.at(0).compare("tr"))
            {
                
                
                
                ShowAllRegisters();
                if (i != StepCount - 1)
                {
                    ShowMessages("\n");
                }
            }

            
            
            
            if (!g_IsInstrumentingInstructions)
            {
                break;
            }
        }

        
        
        
        g_IsInstrumentingInstructions = FALSE;
    }
    else
    {
        ShowMessages("err, stepping (t) is not valid in the current context, you "
                     "should connect to a debuggee\n");
    }
}
