package extension-commands
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\debugger\commands\extension-commands\ioin.cpp.back

type (
Ioin interface{
CommandIoinHelp()(ok bool)//col:13
CommandIoin()(ok bool)//col:89
}
ioin struct{}
)

func NewIoin()Ioin{ return & ioin{} }

func (i *ioin)CommandIoinHelp()(ok bool){//col:13
/*CommandIoinHelp()
{
    ShowMessages("!ioin : detects the execution of IN (I/O instructions) "
                 "instructions.\n\n");
    ShowMessages("syntax : \t!ioin [Port (hex)] [pid ProcessId (hex)] [core CoreId (hex)] "
                 "[imm IsImmediate (yesno)] [buffer PreAllocatedBuffer (hex)] [script { Script (string) }] "
                 "[condition { Condition (hex) }] [code { Code (hex) }]\n");
    ShowMessages("\n");
    ShowMessages("\t\te.g : !ioin\n");
    ShowMessages("\t\te.g : !ioin 0x64\n");
    ShowMessages("\t\te.g : !ioin pid 400\n");
    ShowMessages("\t\te.g : !ioin core 2 pid 400\n");
}*/
return true
}

func (i *ioin)CommandIoin()(ok bool){//col:89
/*CommandIoin(vector<string> SplittedCommand, string Command)
{
    PDEBUGGER_GENERAL_EVENT_DETAIL     Event                 = NULL;
    PDEBUGGER_GENERAL_ACTION           ActionBreakToDebugger = NULL;
    PDEBUGGER_GENERAL_ACTION           ActionCustomCode      = NULL;
    PDEBUGGER_GENERAL_ACTION           ActionScript          = NULL;
    UINT32                             EventLength;
    UINT32                             ActionBreakToDebuggerLength = 0;
    UINT32                             ActionCustomCodeLength      = 0;
    UINT32                             ActionScriptLength          = 0;
    UINT64                             SpecialTarget               = DEBUGGER_EVENT_ALL_IO_PORTS;
    BOOLEAN                            GetPort                     = FALSE;
    vector<string>                     SplittedCommandCaseSensitive {Split(Command, ' ')};
    DEBUGGER_EVENT_PARSING_ERROR_CAUSE EventParsingErrorCause;
    if (!InterpretGeneralEventAndActionsFields(
            &SplittedCommand,
            &SplittedCommandCaseSensitive,
            IN_INSTRUCTION_EXECUTION,
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
        if (!Section.compare("!ioin"))
        {
            continue;
        }
        else if (!GetPort)
        {
            if (!ConvertStringToUInt64(Section, &SpecialTarget))
            {
                ShowMessages("unknown parameter '%s'\n\n", Section.c_str());
                CommandIoinHelp();
                FreeEventsAndActionsMemory(Event, ActionBreakToDebugger, ActionCustomCode, ActionScript);
                return;
            }
            else
            {
                GetPort = TRUE;
            }
        }
        else
        {
            ShowMessages("unknown parameter '%s'\n\n", Section.c_str());
            CommandIoinHelp();
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



