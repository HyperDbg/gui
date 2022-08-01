package extension_commands

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\debugger\commands\extension-commands\msrwrite.cpp.back

type (
	Msrwrite interface {
		CommandMsrwriteHelp() (ok bool) //col:12
		CommandMsrwrite() (ok bool)     //col:88
	}
	msrwrite struct{}
)

func NewMsrwrite() Msrwrite { return &msrwrite{} }

func (m *msrwrite) CommandMsrwriteHelp() (ok bool) { //col:12
	/*
	   CommandMsrwriteHelp()

	   	{
	   	    ShowMessages("!msrwrite : detects the execution of wrmsr instructions.\n\n");
	   	    ShowMessages("syntax : \t!msrwrite [Msr (hex)] [pid ProcessId (hex)] "
	   	                 "[core CoreId (hex)] [imm IsImmediate (yesno)] [buffer PreAllocatedBuffer (hex)] "
	   	                 "[script { Script (string) }] [condition { Condition (hex) }] [code { Code (hex) }]\n");
	   	    ShowMessages("\n");
	   	    ShowMessages("\t\te.g : !msrwrite\n");
	   	    ShowMessages("\t\te.g : !msrwrite 0xc0000082\n");
	   	    ShowMessages("\t\te.g : !msrwrite pid 400\n");
	   	    ShowMessages("\t\te.g : !msrwrite core 2 pid 400\n");
	   	}
	*/
	return true
}

func (m *msrwrite) CommandMsrwrite() (ok bool) { //col:88
	/*
	   CommandMsrwrite(vector<string> SplittedCommand, string Command)

	   	{
	   	    PDEBUGGER_GENERAL_EVENT_DETAIL     Event                 = NULL;
	   	    PDEBUGGER_GENERAL_ACTION           ActionBreakToDebugger = NULL;
	   	    PDEBUGGER_GENERAL_ACTION           ActionCustomCode      = NULL;
	   	    PDEBUGGER_GENERAL_ACTION           ActionScript          = NULL;
	   	    UINT32                             EventLength;
	   	    UINT32                             ActionBreakToDebuggerLength = 0;
	   	    UINT32                             ActionCustomCodeLength      = 0;
	   	    UINT32                             ActionScriptLength          = 0;
	   	    UINT64                             SpecialTarget               = DEBUGGER_EVENT_MSR_READ_OR_WRITE_ALL_MSRS;
	   	    BOOLEAN                            GetAddress                  = FALSE;
	   	    vector<string>                     SplittedCommandCaseSensitive {Split(Command, ' ')};
	   	    DEBUGGER_EVENT_PARSING_ERROR_CAUSE EventParsingErrorCause;
	   	    if (!InterpretGeneralEventAndActionsFields(
	   	            &SplittedCommand,
	   	            &SplittedCommandCaseSensitive,
	   	            WRMSR_INSTRUCTION_EXECUTION,
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
	   	        if (!Section.compare("!msrwrite"))
	   	        {
	   	            continue;
	   	        }
	   	        else if (!GetAddress)
	   	        {
	   	            if (!ConvertStringToUInt64(Section, &SpecialTarget))
	   	            {
	   	                ShowMessages("unknown parameter '%s'\n\n", Section.c_str());
	   	                CommandMsrwriteHelp();
	   	                FreeEventsAndActionsMemory(Event, ActionBreakToDebugger, ActionCustomCode, ActionScript);
	   	                return;
	   	            }
	   	            else
	   	            {
	   	                GetAddress = TRUE;
	   	            }
	   	        }
	   	        else
	   	        {
	   	            ShowMessages("unknown parameter '%s'\n\n", Section.c_str());
	   	            CommandMsrwriteHelp();
	   	            FreeEventsAndActionsMemory(Event, ActionBreakToDebugger, ActionCustomCode, ActionScript);
	   	            return;
	   	        }
	   	    }
	   	    Event->OptionalParam1 = SpecialTarget;
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

