#include "pch.h"
VOID
CommandCrwriteHelp() {
    ShowMessages("!crwrite : monitors modification of control registers (CR0 / CR4).\n\n");
    ShowMessages("syntax : \t!crwrite [Cr (hex)] [mask Mask (hex)] [pid ProcessId (hex)] "
                 "[core CoreId (hex)] [imm IsImmediate (yesno)] [buffer PreAllocatedBuffer (hex)] "
                 "[script { Script (string) }] [condition { Condition (hex) }] [code { Code (hex) }]\n");
    ShowMessages("\n");
    ShowMessages("\t\te.g : !crwrite 0\n");
    ShowMessages("\t\te.g : !crwrite 0 0x10000\n");
    ShowMessages("\t\te.g : !crwrite 4 pid 400\n");
    ShowMessages("\t\te.g : !crwrite 4 core 2 pid 400\n");
}

VOID
CommandCrwrite(vector<string> SplittedCommand, string Command) {
    PDEBUGGER_GENERAL_EVENT_DETAIL     Event                 = NULL;
    PDEBUGGER_GENERAL_ACTION           ActionBreakToDebugger = NULL;
    PDEBUGGER_GENERAL_ACTION           ActionCustomCode      = NULL;
    PDEBUGGER_GENERAL_ACTION           ActionScript          = NULL;
    UINT32                             EventLength;
    UINT32                             ActionBreakToDebuggerLength = 0;
    UINT32                             ActionCustomCodeLength      = 0;
    UINT32                             ActionScriptLength          = 0;
    UINT64                             TargetRegister              = 1;
    UINT64                             MaskRegister                = 0xFFFFFFFFFFFFFFFF;
    BOOLEAN                            GetRegister                 = FALSE;
    BOOLEAN                            GetMask                     = FALSE;
    vector<string>                     SplittedCommandCaseSensitive {Split(Command, ' ')};
    DEBUGGER_EVENT_PARSING_ERROR_CAUSE EventParsingErrorCause;
    if (SplittedCommand.size() < 2) {
        ShowMessages("incorrect use of '!crwrite'\n");
        CommandCrwriteHelp();
        return;
    }
    if (!InterpretGeneralEventAndActionsFields(
            &SplittedCommand,
            &SplittedCommandCaseSensitive,
            CONTROL_REGISTER_MODIFIED,
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
        if (!Section.compare("!crwrite")) {
            continue;
        } else if (!GetRegister) {
            if (!ConvertStringToUInt64(Section, &TargetRegister)) {
                ShowMessages("unknown parameter '%s'\n\n", Section.c_str());
                CommandCrwriteHelp();
                FreeEventsAndActionsMemory(Event, ActionBreakToDebugger, ActionCustomCode, ActionScript);
                return;
            } else {
                GetRegister = TRUE;
            }
        } else if (!GetMask) {
            if (!ConvertStringToUInt64(Section, &MaskRegister)) {
                ShowMessages("unknown parameter '%s'\n\n", Section.c_str());
                CommandCrwriteHelp();
                FreeEventsAndActionsMemory(Event, ActionBreakToDebugger, ActionCustomCode, ActionScript);
                return;
            } else {
                GetMask = TRUE;
            }
        } else {
            ShowMessages("unknown parameter '%s'\n\n", Section.c_str());
            CommandCrwriteHelp();
            FreeEventsAndActionsMemory(Event, ActionBreakToDebugger, ActionCustomCode, ActionScript);
            return;
        }
    }
    if (TargetRegister != VMX_EXIT_QUALIFICATION_REGISTER_CR0 && TargetRegister != VMX_EXIT_QUALIFICATION_REGISTER_CR4) {
        ShowMessages("please choose either 0 for cr0 or 4 for cr4\n");
        CommandCrwriteHelp();
        FreeEventsAndActionsMemory(Event, ActionBreakToDebugger, ActionCustomCode, ActionScript);
        return;
    }
    Event->OptionalParam1 = TargetRegister;

    Event->OptionalParam2 = MaskRegister;
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
