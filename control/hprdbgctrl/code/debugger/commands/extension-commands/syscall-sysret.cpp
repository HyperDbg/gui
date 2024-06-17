#include "pch.h"
VOID CommandSyscallHelp(){
    ShowMessages("!syscall : monitors and hooks all execution of syscall "
                 "instructions (by accessing memory and checking for instructions).\n\n");
    ShowMessages("!syscall2 : monitors and hooks all execution of syscall "
                 "instructions (by emulating all #UDs).\n\n");
    ShowMessages("syntax : \t!syscall [SyscallNumber (hex)] [pid ProcessId (hex)] [core CoreId (hex)] "
                 "[imm IsImmediate (yesno)] [sc EnableShortCircuiting (onoff)] [stage CallingStage (prepostall)] "
                 "[buffer PreAllocatedBuffer (hex)] [script { Script (string) }] [condition { Condition (hex) }] "
                 "[code { Code (hex) }] [output {OutputName (string)}]\n");
    ShowMessages("syntax : \t!syscall2 [SyscallNumber (hex)] [pid ProcessId (hex)] [core CoreId (hex)] "
                 "[imm IsImmediate (yesno)] [sc EnableShortCircuiting (onoff)] [stage CallingStage (prepostall)] "
                 "[buffer PreAllocatedBuffer (hex)] [script { Script (string) }] [condition { Condition (hex) }] "
                 "[code { Code (hex) }] [output {OutputName (string)}]\n");
    ShowMessages("\n");
    ShowMessages("\t\te.g : !syscall\n");
    ShowMessages("\t\te.g : !syscall2\n");
    ShowMessages("\t\te.g : !syscall 0x55\n");
    ShowMessages("\t\te.g : !syscall2 0x55\n");
    ShowMessages("\t\te.g : !syscall 0x55 pid 400\n");
    ShowMessages("\t\te.g : !syscall 0x55 core 2 pid 400\n");
    ShowMessages("\t\te.g : !syscall2 0x55 core 2 pid 400\n");
}
VOID CommandSysretHelp(){
    ShowMessages("!sysret : monitors and hooks all execution of sysret "
                 "instructions (by accessing memory and checking for instructions).\n\n");
    ShowMessages("!sysret2 : monitors and hooks all execution of sysret "
                 "instructions (by emulating all #UDs).\n\n");
    ShowMessages("syntax : \t!sysret [pid ProcessId (hex)] [core CoreId (hex)] "
                 "[imm IsImmediate (yesno)] [sc EnableShortCircuiting (onoff)] [buffer PreAllocatedBuffer (hex)] "
                 "[script { Script (string) }] [condition { Condition (hex) }] [code { Code (hex) }]\n");
    ShowMessages("\n");
    ShowMessages("\t\te.g : !sysret\n");
    ShowMessages("\t\te.g : !sysret2\n");
    ShowMessages("\t\te.g : !sysret pid 400\n");
    ShowMessages("\t\te.g : !sysret2 pid 400\n");
    ShowMessages("\t\te.g : !sysret core 2 pid 400\n");
    ShowMessages("\t\te.g : !sysret2 core 2 pid 400\n");
}
VOID CommandSyscallAndSysret(vector<string> SplitCommand, string Command){
    PDEBUGGER_GENERAL_EVENT_DETAIL     Event                 = NULL;
    PDEBUGGER_GENERAL_ACTION           ActionBreakToDebugger = NULL;
    PDEBUGGER_GENERAL_ACTION           ActionCustomCode      = NULL;
    PDEBUGGER_GENERAL_ACTION           ActionScript          = NULL;
    UINT32                             EventLength;
    UINT32                             ActionBreakToDebuggerLength = 0;
    UINT32                             ActionCustomCodeLength      = 0;
    UINT32                             ActionScriptLength          = 0;
    UINT64                             SpecialTarget               = DEBUGGER_EVENT_SYSCALL_ALL_SYSRET_OR_SYSCALLS;
    BOOLEAN                            GetSyscallNumber            = FALSE;
    vector<string>                     SplitCommandCaseSensitive {Split(Command, ' ')};
    DEBUGGER_EVENT_PARSING_ERROR_CAUSE EventParsingErrorCause;
    string                             Cmd;
    Cmd = SplitCommand.at(0);
    if (!Cmd.compare("!syscall") || !Cmd.compare("!syscall2")){
        if (!InterpretGeneralEventAndActionsFields(
                &SplitCommand,
                &SplitCommandCaseSensitive,
                SYSCALL_HOOK_EFER_SYSCALL,
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
    }
    else{
        if (!InterpretGeneralEventAndActionsFields(
                &SplitCommand,
                &SplitCommandCaseSensitive,
                SYSCALL_HOOK_EFER_SYSRET,
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
    }
    if (!Cmd.compare("!syscall") || !Cmd.compare("!syscall2")){
        for (auto Section : SplitCommand){
            if (!Section.compare("!syscall") ||
                !Section.compare("!syscall2") ||
                !Section.compare("!sysret") ||
                !Section.compare("!sysret2")){
                continue;
            }
            else if (!GetSyscallNumber){
                if (!ConvertStringToUInt64(Section, &SpecialTarget)){
                    ShowMessages("unknown parameter '%s'\n\n", Section.c_str());
                    if (!Cmd.compare("!syscall") || !Cmd.compare("!syscall2")){
                        CommandSyscallHelp();
                    }
                    else{
                        CommandSysretHelp();
                    }
                    FreeEventsAndActionsMemory(Event, ActionBreakToDebugger, ActionCustomCode, ActionScript);
                    return;
                }
                else{
                    GetSyscallNumber = TRUE;
                }
            }
            else{
                ShowMessages("unknown parameter '%s'\n\n", Section.c_str());
                if (!Cmd.compare("!syscall") || !Cmd.compare("!syscall2")){
                    CommandSyscallHelp();
                }
                else{
                    CommandSysretHelp();
                }
                FreeEventsAndActionsMemory(Event, ActionBreakToDebugger, ActionCustomCode, ActionScript);
                return;
            }
        }
        Event->Options.OptionalParam1 = SpecialTarget;
    }
    if (!Cmd.compare("!syscall2") || !Cmd.compare("!sysret2")){
        Event->Options.OptionalParam2 = DEBUGGER_EVENT_SYSCALL_SYSRET_HANDLE_ALL_UD;
    }
    else{
        Event->Options.OptionalParam2 = DEBUGGER_EVENT_SYSCALL_SYSRET_SAFE_ACCESS_MEMORY;
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
