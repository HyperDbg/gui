package extension_commands

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\debugger\commands\extension-commands\interrupt.cpp.back

type (
	Interrupt interface {
		CommandInterruptHelp() (ok bool) //col:12
		CommandInterrupt() (ok bool)     //col:105
	}
	interrupt struct{}
)

func NewInterrupt() Interrupt { return &interrupt{} }

func (i *interrupt) CommandInterruptHelp() (ok bool) { //col:12
	/*
	   CommandInterruptHelp()

	   	{
	   	    ShowMessages("syntax : \t[IdtIndex (hex)] [pid ProcessId (hex)] "
	   	                 "[core CoreId (hex)] [imm IsImmediate (yesno)] [buffer PreAllocatedBuffer (hex)] "
	   	                 "[script { Script (string) }] [condition { Condition (hex) }] [code { Code (hex) }]\n");
	   	    ShowMessages("\nnote : The index should be greater than 0x20 (32) and less "
	   	                 "than 0xFF (255) - starting from zero.\n");
	   	    ShowMessages("\n");
	   	    ShowMessages("\t\te.g : !interrupt 0x2f\n");
	   	    ShowMessages("\t\te.g : !interrupt 0x2f pid 400\n");
	   	    ShowMessages("\t\te.g : !interrupt 0x2f core 2 pid 400\n");
	   	}
	*/
	return true
}

func (i *interrupt) CommandInterrupt() (ok bool) { //col:105
	/*
	   CommandInterrupt(vector<string> SplittedCommand, string Command)

	   	{
	   	    PDEBUGGER_GENERAL_EVENT_DETAIL     Event                 = NULL;
	   	    PDEBUGGER_GENERAL_ACTION           ActionBreakToDebugger = NULL;
	   	    PDEBUGGER_GENERAL_ACTION           ActionCustomCode      = NULL;
	   	    PDEBUGGER_GENERAL_ACTION           ActionScript          = NULL;
	   	    UINT32                             EventLength;
	   	    UINT32                             ActionBreakToDebuggerLength = 0;
	   	    UINT32                             ActionCustomCodeLength      = 0;
	   	    UINT32                             ActionScriptLength          = 0;
	   	    UINT64                             SpecialTarget               = 0;
	   	    BOOLEAN                            GetEntry                    = FALSE;
	   	    vector<string>                     SplittedCommandCaseSensitive {Split(Command, ' ')};
	   	    DEBUGGER_EVENT_PARSING_ERROR_CAUSE EventParsingErrorCause;
	   	    if (!InterpretGeneralEventAndActionsFields(
	   	            &SplittedCommand,
	   	            &SplittedCommandCaseSensitive,
	   	            EXTERNAL_INTERRUPT_OCCURRED,
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
	   	        if (!Section.compare("!interrupt"))
	   	        {
	   	            continue;
	   	        }
	   	        else if (!GetEntry)
	   	        {
	   	            if (!ConvertStringToUInt64(Section, &SpecialTarget))
	   	            {
	   	                ShowMessages("unknown parameter '%s'\n\n", Section.c_str());
	   	                CommandInterruptHelp();
	   	                FreeEventsAndActionsMemory(Event, ActionBreakToDebugger, ActionCustomCode, ActionScript);
	   	                return;
	   	            }
	   	            else
	   	            {
	   	                if (!(SpecialTarget >= 32 && SpecialTarget <= 0xff))
	   	                {
	   	                    ShowMessages("the entry should be between 0x20 to 0xFF or the "
	   	                                 "entries between 32 to 255\n\n");
	   	                    CommandInterruptHelp();
	   	                    FreeEventsAndActionsMemory(Event, ActionBreakToDebugger, ActionCustomCode, ActionScript);
	   	                    return;
	   	                }
	   	                GetEntry = TRUE;
	   	            }
	   	        }
	   	        else
	   	        {
	   	            ShowMessages("unknown parameter '%s'\n\n", Section.c_str());
	   	            CommandInterruptHelp();
	   	            FreeEventsAndActionsMemory(Event, ActionBreakToDebugger, ActionCustomCode, ActionScript);
	   	            return;
	   	        }
	   	    }
	   	    if (SpecialTarget == 0)
	   	    {
	   	        ShowMessages("please specify an interrupt index to monitor, HyperDbg "
	   	                     "doesn't support to trigger events on all interrupts because "
	   	                     "it's not reasonable and make the system unresponsive\n");
	   	        CommandInterruptHelp();
	   	        FreeEventsAndActionsMemory(Event, ActionBreakToDebugger, ActionCustomCode, ActionScript);
	   	        return;
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

