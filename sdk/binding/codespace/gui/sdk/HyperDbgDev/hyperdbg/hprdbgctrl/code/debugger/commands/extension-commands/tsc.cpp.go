package extension_commands

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\debugger\commands\extension-commands\tsc.cpp.back

type (
	Tsc interface {
		CommandTscHelp() (ok bool) //col:12
		CommandTsc() (ok bool)     //col:64
	}
	tsc struct{}
)

func NewTsc() Tsc { return &tsc{} }

func (t *tsc) CommandTscHelp() (ok bool) { //col:12
	/*
	   CommandTscHelp()

	   	{
	   	    ShowMessages("!tsc : monitors execution of rdtsc/rdtscp instructions.\n\n");
	   	    ShowMessages("syntax : \t!tsc [pid ProcessId (hex)] [core CoreId (hex)] "
	   	                 "[imm IsImmediate (yesno)] [buffer PreAllocatedBuffer (hex)] "
	   	                 "[script { Script (string) }] [condition { Condition (hex) }] "
	   	                 "[code { Code (hex) }]\n");
	   	    ShowMessages("\n");
	   	    ShowMessages("\t\te.g : !tsc\n");
	   	    ShowMessages("\t\te.g : !tsc pid 400\n");
	   	    ShowMessages("\t\te.g : !tsc core 2 pid 400\n");
	   	}
	*/
	return true
}

func (t *tsc) CommandTsc() (ok bool) { //col:64
	/*
	   CommandTsc(vector<string> SplittedCommand, string Command)

	   	{
	   	    PDEBUGGER_GENERAL_EVENT_DETAIL     Event                 = NULL;
	   	    PDEBUGGER_GENERAL_ACTION           ActionBreakToDebugger = NULL;
	   	    PDEBUGGER_GENERAL_ACTION           ActionCustomCode      = NULL;
	   	    PDEBUGGER_GENERAL_ACTION           ActionScript          = NULL;
	   	    UINT32                             EventLength;
	   	    UINT32                             ActionBreakToDebuggerLength = 0;
	   	    UINT32                             ActionCustomCodeLength      = 0;
	   	    UINT32                             ActionScriptLength          = 0;
	   	    vector<string>                     SplittedCommandCaseSensitive {Split(Command, ' ')};
	   	    DEBUGGER_EVENT_PARSING_ERROR_CAUSE EventParsingErrorCause;
	   	    if (!InterpretGeneralEventAndActionsFields(
	   	            &SplittedCommand,
	   	            &SplittedCommandCaseSensitive,
	   	            TSC_INSTRUCTION_EXECUTION,
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
	   	    if (SplittedCommand.size() > 1)
	   	    {
	   	        ShowMessages("incorrect use of '!tsc'\n");
	   	        CommandTscHelp();
	   	        FreeEventsAndActionsMemory(Event, ActionBreakToDebugger, ActionCustomCode, ActionScript);
	   	        return;
	   	    }
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

