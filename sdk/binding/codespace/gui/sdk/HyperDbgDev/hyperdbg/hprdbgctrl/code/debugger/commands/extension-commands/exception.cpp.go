package extension-commands
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\debugger\commands\extension-commands\exception.cpp.back

type (
	Exception interface {
		CommandExceptionHelp() (ok bool) //col:15
		CommandException() (ok bool)     //col:99
	}
	exception struct{}
)

func NewException() Exception { return &exception{} }

func (e *exception) CommandExceptionHelp() (ok bool) { //col:15
	/*CommandExceptionHelp()
	  {
	      ShowMessages("!exception : monitors the first 32 entry of IDT (starting from "
	                   "zero).\n\n");
	      ShowMessages(
	          "syntax : \t!exception [IdtIndex (hex)] [pid ProcessId (hex)] "
	          "[core CoreId (hex)] [imm IsImmediate (yesno)] [buffer PreAllocatedBuffer (hex)] "
	          "[script { Script (string) }] [condition { Condition (hex) }] [code { Code (hex) }]\n");
	      ShowMessages("\nnote: monitoring page-faults (entry 0xe) is implemented differently.\n");
	      ShowMessages("\n");
	      ShowMessages("\t\te.g : !exception\n");
	      ShowMessages("\t\te.g : !exception 0xe\n");
	      ShowMessages("\t\te.g : !exception pid 400\n");
	      ShowMessages("\t\te.g : !exception core 2 pid 400\n");
	  }*/
	return true
}

func (e *exception) CommandException() (ok bool) { //col:99
	/*CommandException(vector<string> SplittedCommand, string Command)
	  {
	      PDEBUGGER_GENERAL_EVENT_DETAIL     Event                 = NULL;
	      PDEBUGGER_GENERAL_ACTION           ActionBreakToDebugger = NULL;
	      PDEBUGGER_GENERAL_ACTION           ActionCustomCode      = NULL;
	      PDEBUGGER_GENERAL_ACTION           ActionScript          = NULL;
	      UINT32                             EventLength;
	      UINT32                             ActionBreakToDebuggerLength = 0;
	      UINT32                             ActionCustomCodeLength      = 0;
	      UINT32                             ActionScriptLength          = 0;
	      UINT64                             SpecialTarget               = DEBUGGER_EVENT_EXCEPTIONS_ALL_FIRST_32_ENTRIES;
	      BOOLEAN                            GetEntry                    = FALSE;
	      vector<string>                     SplittedCommandCaseSensitive {Split(Command, ' ')};
	      DEBUGGER_EVENT_PARSING_ERROR_CAUSE EventParsingErrorCause;
	      if (!InterpretGeneralEventAndActionsFields(
	              &SplittedCommand,
	              &SplittedCommandCaseSensitive,
	              EXCEPTION_OCCURRED,
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
	          if (!Section.compare("!exception"))
	          {
	              continue;
	          }
	          else if (!GetEntry)
	          {
	              if (!ConvertStringToUInt64(Section, &SpecialTarget))
	              {
	                  ShowMessages("unknown parameter '%s'\n\n", Section.c_str());
	                  CommandExceptionHelp();
	                  FreeEventsAndActionsMemory(Event, ActionBreakToDebugger, ActionCustomCode, ActionScript);
	                  return;
	              }
	              else
	              {
	                  if (SpecialTarget >= 31)
	                  {
	                      ShowMessages("the entry should be between 0x0 to 0x1f or first 32 "
	                                   "entries'\n\n");
	                      CommandExceptionHelp();
	                      FreeEventsAndActionsMemory(Event, ActionBreakToDebugger, ActionCustomCode, ActionScript);
	                      return;
	                  }
	                  GetEntry = TRUE;
	              }
	          }
	          else
	          {
	              ShowMessages("unknown parameter '%s'\n\n", Section.c_str());
	              CommandExceptionHelp();
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
	  }*/
	return true
}
