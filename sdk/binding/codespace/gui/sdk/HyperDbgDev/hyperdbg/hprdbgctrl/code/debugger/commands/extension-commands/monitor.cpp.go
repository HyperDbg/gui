package extension_commands

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\debugger\commands\extension-commands\monitor.cpp.back

type (
	Monitor interface {
		CommandMonitorHelp() (ok bool) //col:14
		CommandMonitor() (ok bool)     //col:148
	}
	monitor struct{}
)

func NewMonitor() Monitor { return &monitor{} }

func (m *monitor) CommandMonitorHelp() (ok bool) { //col:14
	/*
	   CommandMonitorHelp()

	   	{
	   	    ShowMessages("!monitor : monitors address range for read and writes.\n\n");
	   	    ShowMessages("syntax : \t!monitor [Mode (string)] [FromAddress (hex)] "
	   	                 "[ToAddress (hex)] [pid ProcessId (hex)] [core CoreId (hex)] "
	   	                 "[imm IsImmediate (yesno)] [buffer PreAllocatedBuffer (hex)] "
	   	                 "[script { Script (string) }] [condition { Condition (hex) }] "
	   	                 "[code { Code (hex) }]\n");
	   	    ShowMessages("\n");
	   	    ShowMessages("\t\te.g : !monitor rw fffff801deadb000 fffff801deadbfff\n");
	   	    ShowMessages("\t\te.g : !monitor rw nt!Kd_DEFAULT_Mask Kd_DEFAULT_Mask+5\n");
	   	    ShowMessages("\t\te.g : !monitor r fffff801deadb000 fffff801deadbfff pid 400\n");
	   	    ShowMessages("\t\te.g : !monitor w fffff801deadb000 fffff801deadbfff core 2 pid 400\n");
	   	}
	*/
	return true
}

func (m *monitor) CommandMonitor() (ok bool) { //col:148
	/*
	   CommandMonitor(vector<string> SplittedCommand, string Command)

	   	{
	   	    PDEBUGGER_GENERAL_EVENT_DETAIL     Event                 = NULL;
	   	    PDEBUGGER_GENERAL_ACTION           ActionBreakToDebugger = NULL;
	   	    PDEBUGGER_GENERAL_ACTION           ActionCustomCode      = NULL;
	   	    PDEBUGGER_GENERAL_ACTION           ActionScript          = NULL;
	   	    UINT32                             EventLength;
	   	    UINT32                             ActionBreakToDebuggerLength = 0;
	   	    UINT32                             ActionCustomCodeLength      = 0;
	   	    UINT32                             ActionScriptLength          = 0;
	   	    UINT64                             TargetAddress;
	   	    UINT64                             OptionalParam1 = 0;
	   	    UINT64                             OptionalParam2 = 0;
	   	    BOOLEAN                            SetFrom        = FALSE;
	   	    BOOLEAN                            SetTo          = FALSE;
	   	    BOOLEAN                            SetMode        = FALSE;
	   	    vector<string>                     SplittedCommandCaseSensitive {Split(Command, ' ')};
	   	    UINT32                             IndexInCommandCaseSensitive = 0;
	   	    DEBUGGER_EVENT_PARSING_ERROR_CAUSE EventParsingErrorCause;
	   	    if (SplittedCommand.size() < 4)
	   	    {
	   	        ShowMessages("incorrect use of '!monitor'\n");
	   	        CommandMonitorHelp();
	   	        return;
	   	    }
	   	    if (!InterpretGeneralEventAndActionsFields(
	   	            &SplittedCommand,
	   	            &SplittedCommandCaseSensitive,
	   	            HIDDEN_HOOK_READ_AND_WRITE,
	   	            &Event,
	   	            &EventLength,
	   	            &ActionBreakToDebugger,
	   	            &ActionBreakToDebuggerLength,
	   	            &ActionCustomCode,
	   	            &ActionCustomCodeLength,
	   	            &ActionScript,
	   	            &ActionScriptLength,
	   	            &EventParsingErrorCause))
	   	    {
	   	        return;
	   	    }
	   	    for (auto Section : SplittedCommand)
	   	    {
	   	        IndexInCommandCaseSensitive++;
	   	        if (!Section.compare("!monitor"))
	   	        {
	   	            continue;
	   	        }
	   	        else if (!Section.compare("r") && !SetMode)
	   	        {
	   	            Event->EventType = HIDDEN_HOOK_READ;
	   	            SetMode          = TRUE;
	   	        }
	   	        else if (!Section.compare("w") && !SetMode)
	   	        {
	   	            Event->EventType = HIDDEN_HOOK_WRITE;
	   	            SetMode          = TRUE;
	   	        }
	   	        else if ((!Section.compare("rw") || !Section.compare("wr")) && !SetMode)
	   	        {
	   	            Event->EventType = HIDDEN_HOOK_READ_AND_WRITE;
	   	            SetMode          = TRUE;
	   	        }
	   	        else
	   	        {
	   	            if (!SetFrom)
	   	            {
	   	                if (!SymbolConvertNameOrExprToAddress(
	   	                        SplittedCommandCaseSensitive.at(IndexInCommandCaseSensitive - 1),
	   	                        &OptionalParam1))
	   	                {
	   	                    ShowMessages("err, couldn't resolve error at '%s'\n\n",
	   	                                 SplittedCommandCaseSensitive.at(IndexInCommandCaseSensitive - 1).c_str());
	   	                    CommandMonitorHelp();
	   	                    FreeEventsAndActionsMemory(Event, ActionBreakToDebugger, ActionCustomCode, ActionScript);
	   	                    return;
	   	                }
	   	                SetFrom = TRUE;
	   	            }
	   	            else if (!SetTo)
	   	            {
	   	                if (!SymbolConvertNameOrExprToAddress(
	   	                        SplittedCommandCaseSensitive.at(IndexInCommandCaseSensitive - 1),
	   	                        &OptionalParam2))
	   	                {
	   	                    ShowMessages("err, couldn't resolve error at '%s'\n\n",
	   	                                 SplittedCommandCaseSensitive.at(IndexInCommandCaseSensitive - 1).c_str());
	   	                    CommandMonitorHelp();
	   	                    FreeEventsAndActionsMemory(Event, ActionBreakToDebugger, ActionCustomCode, ActionScript);
	   	                    return;
	   	                }
	   	                SetTo = TRUE;
	   	            }
	   	            else
	   	            {
	   	                ShowMessages("unknown parameter '%s'\n\n", Section.c_str());
	   	                CommandMonitorHelp();
	   	                FreeEventsAndActionsMemory(Event, ActionBreakToDebugger, ActionCustomCode, ActionScript);
	   	                return;
	   	            }
	   	        }
	   	    }
	   	    if (OptionalParam1 > OptionalParam2)
	   	    {
	   	        ShowMessages("please choose the 'from' value first, then choose the 'to' "
	   	                     "value\n");
	   	        FreeEventsAndActionsMemory(Event, ActionBreakToDebugger, ActionCustomCode, ActionScript);
	   	        return;
	   	    }
	   	    if (!SetMode)
	   	    {
	   	        ShowMessages("please specify the attribute(s) that you want to monitor (r, w, rw)\n");
	   	        FreeEventsAndActionsMemory(Event, ActionBreakToDebugger, ActionCustomCode, ActionScript);
	   	        return;
	   	    }
	   	    Event->OptionalParam1 = OptionalParam1;
	   	    Event->OptionalParam2 = OptionalParam2;
	   	    if (!SendEventToKernel(Event, EventLength))
	   	    {
	   	        FreeEventsAndActionsMemory(Event, ActionBreakToDebugger, ActionCustomCode, ActionScript);
	   	        return;
	   	    }
	   	    if (!RegisterActionToEvent(Event,
	   	                               ActionBreakToDebugger,
	   	                               ActionBreakToDebuggerLength,
	   	                               ActionCustomCode,
	   	                               ActionCustomCodeLength,
	   	                               ActionScript,
	   	                               ActionScriptLength))
	   	    {
	   	        FreeEventsAndActionsMemory(Event, ActionBreakToDebugger, ActionCustomCode, ActionScript);
	   	        return;
	   	    }
	   	}
	*/
	return true
}

