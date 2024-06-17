#include "pch.h"
VOID CommandIoinHelp(){
    ShowMessages("!ioin : detects the execution of IN (I/O instructions) "
                 "instructions.\n\n");
    ShowMessages("syntax : \t!ioin [Port (hex)] [pid ProcessId (hex)] [core CoreId (hex)] "
                 "[imm IsImmediate (yesno)] [sc EnableShortCircuiting (onoff)] [stage CallingStage (prepostall)] "
                 "[buffer PreAllocatedBuffer (hex)] [script { Script (string) }] [condition { Condition (hex) }] "
                 "[code { Code (hex) }] [output {OutputName (string)}]\n");
    ShowMessages("\n");
    ShowMessages("\t\te.g : !ioin\n");
    ShowMessages("\t\te.g : !ioin 0x64\n");
    ShowMessages("\t\te.g : !ioin pid 400\n");
    ShowMessages("\t\te.g : !ioin core 2 pid 400\n");
}
VOID CommandIoin(vector<string> SplitCommand, string Command){
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
    vector<string>                     SplitCommandCaseSensitive {Split(Command, ' ')};
    DEBUGGER_EVENT_PARSING_ERROR_CAUSE EventParsingErrorCause;
    if (!InterpretGeneralEventAndActionsFields(
            &SplitCommand,
            &SplitCommandCaseSensitive,
            IN_INSTRUCTION_EXECUTION,
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
        if (!Section.compare("!ioin")){
            continue;
        }
        else if (!GetPort){
            if (!ConvertStringToUInt64(Section, &SpecialTarget)){
                ShowMessages("unknown parameter '%s'\n\n", Section.c_str());
                CommandIoinHelp();
                FreeEventsAndActionsMemory(Event, ActionBreakToDebugger, ActionCustomCode, ActionScript);
                return;
            }
            else{
                GetPort = TRUE;
            }
        }
        else{
            ShowMessages("unknown parameter '%s'\n\n", Section.c_str());
            CommandIoinHelp();
            FreeEventsAndActionsMemory(Event, ActionBreakToDebugger, ActionCustomCode, ActionScript);
            return;
        }
    }
    Event->Options.OptionalParam1 = SpecialTarget;
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
