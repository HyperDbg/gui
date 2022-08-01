package core

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\debugger\core\break-control.cpp.back

type (
	BreakControl interface {
		BreakController() (ok bool) //col:80
	}
	breakControl struct{}
)

func NewBreakControl() BreakControl { return &breakControl{} }

func (b *breakControl) BreakController() (ok bool) { //col:80
	/*
	   BreakController(DWORD CtrlType)

	   	{
	   	    switch (CtrlType)
	   	    {
	   	    case CTRL_BREAK_EVENT:
	   	    case CTRL_C_EVENT:
	   	        if (g_IgnorePauseRequests)
	   	        {
	   	            return TRUE;
	   	        }
	   	        if (g_IsExecutingSymbolLoadingRoutines)
	   	        {
	   	            g_IsExecutingSymbolLoadingRoutines = FALSE;
	   	            ScriptEngineSymbolAbortLoadingWrapper();
	   	        }
	   	        if (g_IsSerialConnectedToRemoteDebuggee)
	   	        {
	   	            if (g_IsInstrumentingInstructions)
	   	            {
	   	                g_IsInstrumentingInstructions = FALSE;
	   	            }
	   	            else
	   	            {
	   	                KdBreakControlCheckAndPauseDebugger();
	   	            }
	   	        }
	   	        else if (!g_IsDebuggerModulesLoaded && !g_IsConnectedToRemoteDebuggee)
	   	        {
	   	        }
	   	        else
	   	        {
	   	            if (g_IsInstrumentingInstructions)
	   	            {
	   	                g_IsInstrumentingInstructions = FALSE;
	   	            }
	   	            else
	   	            {
	   	                g_BreakPrintingOutput = TRUE;
	   	                if (g_IsConnectedToRemoteDebuggee)
	   	                {
	   	                    RemoteConnectionSendCommand("pause", strlen("pause") + 1);
	   	                }
	   	                Sleep(300);
	   	                if (!g_IsConnectedToRemoteDebuggee)
	   	                {
	   	                    if (g_AutoUnpause)
	   	                    {
	   	                        ShowMessages(
	   	                            "\npausing...\nauto-unpause mode is enabled, "
	   	                            "debugger will automatically continue when you run a new "
	   	                            "event command, if you want to change this behaviour then "
	   	                            "run run 'settings autounpause off'\n\n");
	   	                    }
	   	                    else
	   	                    {
	   	                        ShowMessages(
	   	                            "\npausing...\nauto-unpause mode is disabled, you "
	   	                            "should run 'g' when you want to continue, otherwise run "
	   	                            "'settings "
	   	                            "autounpause on'\n\n");
	   	                    }
	   	                    HyperDbgShowSignature();
	   	                    if (g_ActiveProcessDebuggingState.IsActive)
	   	                    {
	   	                        UdPauseProcess(g_ActiveProcessDebuggingState.ProcessDebuggingToken);
	   	                    }
	   	                }
	   	            }
	   	        }
	   	        return TRUE;
	   	    case CTRL_CLOSE_EVENT:
	   	        return TRUE;
	   	    case CTRL_LOGOFF_EVENT:
	   	        return FALSE;
	   	    case CTRL_SHUTDOWN_EVENT:
	   	        return FALSE;
	   	    default:
	   	        return FALSE;
	   	    }
	   	}
	*/
	return true
}

