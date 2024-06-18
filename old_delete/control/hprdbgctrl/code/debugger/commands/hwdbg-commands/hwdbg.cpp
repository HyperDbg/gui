#include "pch.h"
VOID CommandHwdbgHelp(){
    ShowMessages("!hwdbg : performs actions related to hwdbg hardware debugging.\n\n");
    ShowMessages("syntax : \t!hwdbg  [script { Script (string) }]\n");
    ShowMessages("\n");
    ShowMessages("\t\te.g : !hwdbg script { @hw_pin1 = 0; }\n");
}
VOID CommandHwdbg(vector<string> SplitCommand, string Command){
    PDEBUGGER_GENERAL_EVENT_DETAIL     Event                 = NULL;
    PDEBUGGER_GENERAL_ACTION           ActionBreakToDebugger = NULL;
    PDEBUGGER_GENERAL_ACTION           ActionCustomCode      = NULL;
    PDEBUGGER_GENERAL_ACTION           ActionScript          = NULL;
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
            (VMM_EVENT_TYPE_ENUM)NULL, 
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
    ShowMessages("hwdbg script buffer (size=%d, stages=%d, flip-flops=%d):\n",
                 ActionScript->ScriptBufferSize,
                 ActionScript->ScriptBufferSize / 32,
                 ActionScript->ScriptBufferSize * 8);
    CHAR * ScriptBuffer = (CHAR *)((UINT64)ActionScript + sizeof(DEBUGGER_GENERAL_ACTION));
    for (size_t i = 0; i < ActionScript->ScriptBufferSize; i++){
        ShowMessages("%02x ", ScriptBuffer[i]);
    }
    FreeEventsAndActionsMemory(Event, ActionBreakToDebugger, ActionCustomCode, ActionScript);
}
