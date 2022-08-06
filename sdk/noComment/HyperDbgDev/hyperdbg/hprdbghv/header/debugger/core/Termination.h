#pragma once
VOID TerminateExternalInterruptEvent(PDEBUGGER_EVENT Event);
VOID TerminateHiddenHookReadAndWriteEvent(PDEBUGGER_EVENT Event);
VOID TerminateHiddenHookExecCcEvent(PDEBUGGER_EVENT Event);
VOID TerminateHiddenHookExecDetoursEvent(PDEBUGGER_EVENT Event);
VOID TerminateRdmsrExecutionEvent(PDEBUGGER_EVENT Event);
VOID TerminateWrmsrExecutionEvent(PDEBUGGER_EVENT Event);
VOID TerminateExceptionEvent(PDEBUGGER_EVENT Event);
VOID TerminateInInstructionExecutionEvent(PDEBUGGER_EVENT Event);
VOID TerminateOutInstructionExecutionEvent(PDEBUGGER_EVENT Event);
VOID TerminateSyscallHookEferEvent(PDEBUGGER_EVENT Event);
VOID TerminateSysretHookEferEvent(PDEBUGGER_EVENT Event);
VOID TerminateVmcallExecutionEvent(PDEBUGGER_EVENT Event);
VOID TerminateTscEvent(PDEBUGGER_EVENT Event);
VOID TerminatePmcEvent(PDEBUGGER_EVENT Event);
VOID TerminateDebugRegistersEvent(PDEBUGGER_EVENT Event);
VOID TerminateCpuidExecutionEvent(PDEBUGGER_EVENT Event);
VOID TerminateControlRegistersEvent(PDEBUGGER_EVENT Event);
