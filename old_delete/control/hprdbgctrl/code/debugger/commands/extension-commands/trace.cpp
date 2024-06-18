#include "pch.h"
VOID CommandTraceHelp(){
    ShowMessages("!trace : traces the execution of user-mode/kernel-mode instructions.\n\n");
    ShowMessages("syntax : \t!trace [TraceType (string)] [pid ProcessId (hex)] [core CoreId (hex)] [imm IsImmediate (yesno)] "
                 "[sc EnableShortCircuiting (onoff)] [buffer PreAllocatedBuffer (hex)] [script { Script (string) }] "
                 "[condition { Condition (hex) }] [code { Code (hex) }] [output {OutputName (string)}]\n");
    ShowMessages("\n");
    ShowMessages("\t\te.g : !trace step-out\n");
    ShowMessages("\t\te.g : !trace step-in pid 1c0\n");
    ShowMessages("\t\te.g : !trace instrument-step core 2 pid 400\n");
    ShowMessages("\n");
    ShowMessages("valid trace types: \n");
    ShowMessages("\tstep-in : single step-in (regular)\n");
    ShowMessages("\tstep-out : single step-out (regular)\n");
    ShowMessages("\tinstrument-step : single step-in (instrumentation)\n");
}
VOID CommandTrace(vector<string> SplitCommand, string Command){
    PDEBUGGER_GENERAL_EVENT_DETAIL     Event                 = NULL;
    PDEBUGGER_GENERAL_ACTION           ActionBreakToDebugger = NULL;
    PDEBUGGER_GENERAL_ACTION           ActionCustomCode      = NULL;
    PDEBUGGER_GENERAL_ACTION           ActionScript          = NULL;
    UINT32                             EventLength;
    UINT32                             ActionBreakToDebuggerLength = 0;
    UINT32                             ActionCustomCodeLength      = 0;
    UINT32                             ActionScriptLength          = 0;
    vector<string>                     SplitCommandCaseSensitive {Split(Command, ' ')};
    DEBUGGER_EVENT_PARSING_ERROR_CAUSE EventParsingErrorCause;
    BOOLEAN                            SetTraceType = FALSE;
    DEBUGGER_EVENT_TRACE_TYPE          TargetTrace  = DEBUGGER_EVENT_TRACE_TYPE_INVALID;
    if (!InterpretGeneralEventAndActionsFields(
            &SplitCommand,
            &SplitCommandCaseSensitive,
            TRAP_EXECUTION_INSTRUCTION_TRACE,
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
    if (SplitCommand.size() > 2){
        ShowMessages("incorrect use of the '!trace'\n");
        CommandTraceHelp();
        FreeEventsAndActionsMemory(Event, ActionBreakToDebugger, ActionCustomCode, ActionScript);
        return;
    }
    for (auto Section : SplitCommand){
        if (!Section.compare("!trace")){
            continue;
        }
        else if ((!Section.compare("step-in") || !Section.compare("stepin") || !Section.compare("step")) && !SetTraceType){
            TargetTrace  = DEBUGGER_EVENT_TRACE_TYPE_STEP_IN;
            SetTraceType = TRUE;
        }
        else if ((!Section.compare("step-out") || !Section.compare("stepout")) && !SetTraceType){
            TargetTrace  = DEBUGGER_EVENT_TRACE_TYPE_STEP_OUT;
            SetTraceType = TRUE;
        }
        else if ((!Section.compare("step-instrument") || !Section.compare("instrument-step") ||
                  !Section.compare("instrumentstep") ||
                  !Section.compare("instrument-step-in")) &&
                 !SetTraceType){
            TargetTrace  = DEBUGGER_EVENT_TRACE_TYPE_INSTRUMENTATION_STEP_IN;
            SetTraceType = TRUE;
        }
        else{
            ShowMessages("err, couldn't resolve error at '%s'\n\n",
                         Section.c_str());
            CommandTraceHelp();
            FreeEventsAndActionsMemory(Event, ActionBreakToDebugger, ActionCustomCode, ActionScript);
        }
    }
    if (!SetTraceType){
        ShowMessages("please specify the trace type\n");
        FreeEventsAndActionsMemory(Event, ActionBreakToDebugger, ActionCustomCode, ActionScript);
        return;
    }
    Event->Options.OptionalParam1 = (UINT64)SetTraceType;
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
