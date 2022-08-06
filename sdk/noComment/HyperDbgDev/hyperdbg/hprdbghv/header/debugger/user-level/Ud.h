#pragma once
BOOLEAN
UdInitializeUserDebugger();
VOID
UdUninitializeUserDebugger();
BOOLEAN
UdCheckAndHandleBreakpointsAndDebugBreaks(UINT32                            CurrentCore,
                                          PGUEST_REGS                       GuestRegs,
                                          DEBUGGEE_PAUSING_REASON           Reason,
                                          PDEBUGGER_TRIGGERED_EVENT_DETAILS EventDetails);
BOOLEAN
UdDispatchUsermodeCommands(PDEBUGGER_UD_COMMAND_PACKET ActionRequest);
BOOLEAN
UdCheckForCommand();
