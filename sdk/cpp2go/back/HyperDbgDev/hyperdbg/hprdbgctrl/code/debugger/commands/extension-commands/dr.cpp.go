package extension-commands
//back\HyperDbgDev\hyperdbg\hprdbgctrl\code\debugger\commands\extension-commands\dr.cpp.back

type (
Dr interface{
CommandDrHelp()(ok bool)//col:33
CommandDr()(ok bool)//col:122
}
)

func NewDr() { return & dr{} }

func (d *dr)CommandDrHelp()(ok bool){//col:33
/*CommandDrHelp()
{
    ShowMessages("!dr : monitors any access to debug registers.\n\n");
    ShowMessages("syntax : \t!dr [pid ProcessId (hex)] [core CoreId (hex)] "
                 "[imm IsImmediate (yesno)] [buffer PreAllocatedBuffer (hex)] "
                 "[script { Script (string) }] [condition { Condition (hex) }] "
                 "[code { Code (hex) }]\n");
    ShowMessages("\n");
    ShowMessages("\t\te.g : !dr\n");
    ShowMessages("\t\te.g : !dr pid 400\n");
    ShowMessages("\t\te.g : !dr core 2 pid 400\n");
}*/
return true
}

func (d *dr)CommandDr()(ok bool){//col:122
/*CommandDr(vector<string> SplittedCommand, string Command)
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
            DEBUG_REGISTERS_ACCESSED,
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
        ShowMessages("incorrect use of '!dr'\n");
        CommandDrHelp();
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
}*/
return true
}



