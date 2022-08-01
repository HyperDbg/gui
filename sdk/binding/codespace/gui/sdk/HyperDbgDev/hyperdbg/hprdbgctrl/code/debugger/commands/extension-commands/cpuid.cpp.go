package extension-commands
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\debugger\commands\extension-commands\cpuid.cpp.back

type (
Cpuid interface{
CommandCpuidHelp()(ok bool)//col:13
CommandCpuid()(ok bool)//col:65
}
cpuid struct{}
)

func NewCpuid()Cpuid{ return & cpuid{} }

func (c *cpuid)CommandCpuidHelp()(ok bool){//col:13
/*CommandCpuidHelp()
{
    ShowMessages("!cpuid : monitors execution of a special cpuid index or all "
                 "cpuids instructions.\n\n");
    ShowMessages("syntax : \t!cpuid [pid ProcessId (hex)] [core CoreId (hex)] "
                 "[imm IsImmediate (yesno)] [buffer PreAllocatedBuffer (hex)] "
                 "[script { Script (string) }] [condition { Condition (hex) }] "
                 "[code { Code (hex) }]\n");
    ShowMessages("\n");
    ShowMessages("\t\te.g : !cpuid\n");
    ShowMessages("\t\te.g : !cpuid pid 400\n");
    ShowMessages("\t\te.g : !cpuid core 2 pid 400\n");
}*/
return true
}

func (c *cpuid)CommandCpuid()(ok bool){//col:65
/*CommandCpuid(vector<string> SplittedCommand, string Command)
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
            CPUID_INSTRUCTION_EXECUTION,
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
        ShowMessages("incorrect use of '!cpuid'\n");
        CommandCpuidHelp();
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



