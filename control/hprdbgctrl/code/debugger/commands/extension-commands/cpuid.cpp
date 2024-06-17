#include "pch.h"
VOID CommandCpuidHelp(){
    ShowMessages("!cpuid : monitors execution of a special cpuid index or all "
                 "cpuids instructions.\n\n");
    ShowMessages("syntax : \t!cpuid [Eax (hex)] [pid ProcessId (hex)] [core CoreId (hex)] "
                 "[imm IsImmediate (yesno)] [sc EnableShortCircuiting (onoff)] [stage CallingStage (prepostall)] "
                 "[buffer PreAllocatedBuffer (hex)] [script { Script (string) }] [condition { Condition (hex) }] "
                 "[code { Code (hex) }] [output {OutputName (string)}]\n");
    ShowMessages("\n");
    ShowMessages("\t\te.g : !cpuid\n");
    ShowMessages("\t\te.g : !cpuid 1\n");
    ShowMessages("\t\te.g : !cpuid pid 400\n");
    ShowMessages("\t\te.g : !cpuid core 2 pid 400\n");
}
VOID CommandCpuid(vector<string> SplitCommand, string Command){
    PDEBUGGER_GENERAL_EVENT_DETAIL     Event                 = NULL;
    PDEBUGGER_GENERAL_ACTION           ActionBreakToDebugger = NULL;
    PDEBUGGER_GENERAL_ACTION           ActionCustomCode      = NULL;
    PDEBUGGER_GENERAL_ACTION           ActionScript          = NULL;
    BOOLEAN                            GetEax                = FALSE;
    UINT32                             EventLength;
    UINT64                             SpecialTarget               = 0;
    UINT32                             ActionBreakToDebuggerLength = 0;
    UINT32                             ActionCustomCodeLength      = 0;
    UINT32                             ActionScriptLength          = 0;
    vector<string>                     SplitCommandCaseSensitive {Split(Command, ' ')};
    DEBUGGER_EVENT_PARSING_ERROR_CAUSE EventParsingErrorCause;
    if (!InterpretGeneralEventAndActionsFields(
            &SplitCommand,
            &SplitCommandCaseSensitive,
            CPUID_INSTRUCTION_EXECUTION,
            &Event,
            &EventLength,
            &ActionBreakToDebugger,
            &ActionBreakToDebuggerLength,
            &ActionCustomCode,
            &ActionCustomCodeLength,
            &ActionScript,
            &ActionScriptLength,
            &EventParsingErrorCause)){
        return;
    }
    for (auto Section : SplitCommand){
        if (!Section.compare("!cpuid")){
            continue;
        }
        else if (!GetEax){
            if (!ConvertStringToUInt64(Section, &SpecialTarget)){
                ShowMessages("unknown parameter '%s'\n\n", Section.c_str());
                CommandCpuidHelp();
                FreeEventsAndActionsMemory(Event, ActionBreakToDebugger, ActionCustomCode, ActionScript);
                return;
            }
            else{
                GetEax = TRUE;
            }
        }
        else{
            ShowMessages("unknown parameter '%s'\n\n", Section.c_str());
            CommandCpuidHelp();
            FreeEventsAndActionsMemory(Event, ActionBreakToDebugger, ActionCustomCode, ActionScript);
            return;
        }
    }
    Event->Options.OptionalParam1 = GetEax;
    if (GetEax){
        Event->Options.OptionalParam2 = SpecialTarget;
    }
    if (!SendEventToKernel(Event, EventLength)){
        FreeEventsAndActionsMemory(Event, ActionBreakToDebugger, ActionCustomCode, ActionScript);
        return;
    }
    if (!RegisterActionToEvent(Event,
                               ActionBreakToDebugger,
                               ActionBreakToDebuggerLength,
                               ActionCustomCode,
                               ActionCustomCodeLength,
                               ActionScript,
                               ActionScriptLength)){
        FreeEventsAndActionsMemory(Event, ActionBreakToDebugger, ActionCustomCode, ActionScript);
        return;
    }
}
