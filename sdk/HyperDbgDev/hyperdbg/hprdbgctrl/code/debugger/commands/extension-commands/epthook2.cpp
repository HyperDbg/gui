#include "pch.h"
VOID
CommandEptHook2Help() {
    ShowMessages("!epthook2 : puts a hidden-hook EPT (detours).\n\n");
    ShowMessages(
        "syntax : \t!epthook2 [Address (hex)] [pid ProcessId (hex)] "
        "[core CoreId (hex)] [imm IsImmediate (yesno)] [buffer PreAllocatedBuffer (hex)] "
        "[script { Script (string) }] [condition { Condition (hex) }] [code { Code (hex) }]\n");
    ShowMessages("\n");
    ShowMessages("\t\te.g : !epthook2 nt!ExAllocatePoolWithTag\n");
    ShowMessages("\t\te.g : !epthook2 nt!ExAllocatePoolWithTag+5\n");
    ShowMessages("\t\te.g : !epthook2 fffff801deadb000\n");
    ShowMessages("\t\te.g : !epthook2 fffff801deadb000 pid 400\n");
    ShowMessages("\t\te.g : !epthook2 fffff801deadb000 core 2 pid 400\n");
}

VOID
CommandEptHook2(vector<string> SplittedCommand, string Command) {
    PDEBUGGER_GENERAL_EVENT_DETAIL     Event                 = NULL;
    PDEBUGGER_GENERAL_ACTION           ActionBreakToDebugger = NULL;
    PDEBUGGER_GENERAL_ACTION           ActionCustomCode      = NULL;
    PDEBUGGER_GENERAL_ACTION           ActionScript          = NULL;
    UINT32                             EventLength;
    UINT32                             ActionBreakToDebuggerLength = 0;
    UINT32                             ActionCustomCodeLength      = 0;
    UINT32                             ActionScriptLength          = 0;
    BOOLEAN                            GetAddress                  = FALSE;
    UINT64                             OptionalParam1              = 0; // Set the target address
    vector<string>                     SplittedCommandCaseSensitive {Split(Command, ' ')};
    UINT32                             IndexInCommandCaseSensitive = 0;
    DEBUGGER_EVENT_PARSING_ERROR_CAUSE EventParsingErrorCause;
    if (SplittedCommand.size() < 2) {
        ShowMessages("incorrect use of '!epthook2'\n");
        CommandEptHook2Help();
        return;
    }
    if (!InterpretGeneralEventAndActionsFields(
            &SplittedCommand,
            &SplittedCommandCaseSensitive,
            HIDDEN_HOOK_EXEC_DETOURS,
            &Event,
            &EventLength,
            &ActionBreakToDebugger,
            &ActionBreakToDebuggerLength,
            &ActionCustomCode,
            &ActionCustomCodeLength,
            &ActionScript,
            &ActionScriptLength,
            &EventParsingErrorCause)) {
        return;
    }
    for (auto Section : SplittedCommand) {
        IndexInCommandCaseSensitive++;
        if (!Section.compare("!epthook2")) {
            continue;
        } else if (!GetAddress) {
            if (!SymbolConvertNameOrExprToAddress(
                    SplittedCommandCaseSensitive.at(IndexInCommandCaseSensitive - 1),
                    &OptionalParam1)) {
                ShowMessages("err, couldn't resolve error at '%s'\n\n",
                             SplittedCommandCaseSensitive.at(IndexInCommandCaseSensitive - 1).c_str());
                CommandEptHook2Help();
                FreeEventsAndActionsMemory(Event, ActionBreakToDebugger, ActionCustomCode, ActionScript);
                return;
            } else {
                GetAddress = TRUE;
            }
        } else {
            ShowMessages("unknown parameter '%s'\n\n", Section.c_str());
            CommandEptHook2Help();
            FreeEventsAndActionsMemory(Event, ActionBreakToDebugger, ActionCustomCode, ActionScript);
            return;
        }
    }
    if (OptionalParam1 == 0) {
        ShowMessages("please choose an address to put the hook on it\n");
        FreeEventsAndActionsMemory(Event, ActionBreakToDebugger, ActionCustomCode, ActionScript);
        return;
    }
    Event->OptionalParam1 = OptionalParam1;
    if (!SendEventToKernel(Event, EventLength)) {
        FreeEventsAndActionsMemory(Event, ActionBreakToDebugger, ActionCustomCode, ActionScript);
        return;
    }
    if (!RegisterActionToEvent(Event,
                               ActionBreakToDebugger,
                               ActionBreakToDebuggerLength,
                               ActionCustomCode,
                               ActionCustomCodeLength,
                               ActionScript,
                               ActionScriptLength)) {
        FreeEventsAndActionsMemory(Event, ActionBreakToDebugger, ActionCustomCode, ActionScript);
        return;
    }
}
