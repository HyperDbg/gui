#include "pch.h"
extern DEBUGGER_SYNCRONIZATION_EVENTS_STATE
                                            g_KernelSyncronizationObjectsHandleTable[DEBUGGER_MAXIMUM_SYNCRONIZATION_KERNEL_DEBUGGER_OBJECTS];
extern BYTE                                 g_CurrentRunningInstruction[MAXIMUM_INSTR_SIZE];
extern OVERLAPPED                           g_OverlappedIoStructureForReadDebugger;
extern OVERLAPPED                           g_OverlappedIoStructureForWriteDebugger;
extern HANDLE                               g_SerialRemoteComPortHandle;
extern BOOLEAN                              g_IsSerialConnectedToRemoteDebuggee;
extern BOOLEAN                              g_IsDebuggeeRunning;
extern BOOLEAN                              g_IgnoreNewLoggingMessages;
extern BOOLEAN                              g_SharedEventStatus;
extern BOOLEAN                              g_IsRunningInstruction32Bit;
extern ULONG                                g_CurrentRemoteCore;
extern DEBUGGER_EVENT_AND_ACTION_REG_BUFFER g_DebuggeeResultOfRegisteringEvent;
extern DEBUGGER_EVENT_AND_ACTION_REG_BUFFER
              g_DebuggeeResultOfAddingActionsToEvent;
extern UINT64 g_ResultOfEvaluatedExpression;
extern UINT32 g_ErrorStateOfResultOfEvaluatedExpression;
BOOLEAN
ListeningSerialPortInDebugger() {
    PDEBUGGER_PREPARE_DEBUGGEE                  InitPacket;
    PDEBUGGER_REMOTE_PACKET                     TheActualPacket;
    PDEBUGGEE_KD_PAUSED_PACKET                  PausePacket;
    PDEBUGGEE_MESSAGE_PACKET                    MessagePacket;
    PDEBUGGEE_CHANGE_CORE_PACKET                ChangeCorePacket;
    PDEBUGGEE_SCRIPT_PACKET                     ScriptPacket;
    PDEBUGGEE_FORMATS_PACKET                    FormatsPacket;
    PDEBUGGER_EVENT_AND_ACTION_REG_BUFFER       EventAndActionPacket;
    PDEBUGGER_UPDATE_SYMBOL_TABLE               SymbolUpdatePacket;
    PDEBUGGER_MODIFY_EVENTS                     EventModifyAndQueryPacket;
    PDEBUGGEE_SYMBOL_UPDATE_RESULT              SymbolReloadFinishedPacket;
    PDEBUGGEE_DETAILS_AND_SWITCH_PROCESS_PACKET ChangeProcessPacket;
    PDEBUGGEE_RESULT_OF_SEARCH_PACKET           SearchResultsPacket;
    PDEBUGGEE_DETAILS_AND_SWITCH_THREAD_PACKET  ChangeThreadPacket;
    PDEBUGGER_FLUSH_LOGGING_BUFFERS             FlushPacket;
    PDEBUGGER_CALLSTACK_REQUEST                 CallstackPacket;
    PDEBUGGER_SINGLE_CALLSTACK_FRAME            CallstackFramePacket;
    PDEBUGGER_DEBUGGER_TEST_QUERY_BUFFER        TestQueryPacket;
    PDEBUGGEE_REGISTER_READ_DESCRIPTION         ReadRegisterPacket;
    PDEBUGGER_READ_MEMORY                       ReadMemoryPacket;
    PDEBUGGER_EDIT_MEMORY                       EditMemoryPacket;
    PDEBUGGEE_BP_PACKET                         BpPacket;
    PDEBUGGER_READ_PAGE_TABLE_ENTRIES_DETAILS   PtePacket;
    PDEBUGGER_VA2PA_AND_PA2VA_COMMANDS          Va2paPa2vaPacket;
    PDEBUGGEE_BP_LIST_OR_MODIFY_PACKET          ListOrModifyBreakpointPacket;
    PGUEST_REGS                                 Regs;
    PGUEST_EXTRA_REGISTERS                      ExtraRegs;
    unsigned char *                             MemoryBuffer;
    BOOLEAN                                     ShowSignatureWhenDisconnected = FALSE;
StartAgain:
    CHAR   BufferToReceive[MaxSerialPacketSize] = {0};
    UINT32 LengthReceived                       = 0;
    if (!KdReceivePacketFromDebuggee(BufferToReceive, &LengthReceived)) {
        if (LengthReceived == 0 && BufferToReceive[0] == NULL) {
            ShowMessages("\nthe remote connection is closed\n");
            if (g_IsSerialConnectedToRemoteDebuggee) {
                CommandEventsClearAllEventsAndResetTags();
                if (g_IsDebuggeeRunning == FALSE) {
                    ShowSignatureWhenDisconnected = TRUE;
                }
            }
            KdCloseConnection();
            if (ShowSignatureWhenDisconnected) {
                ShowSignatureWhenDisconnected = FALSE;
                ShowMessages("\n");
            }
            return FALSE;
        } else {
            ShowMessages("err, invalid buffer received\n");
            goto StartAgain;
        }
    }
    if (LengthReceived == 1 && BufferToReceive[0] == NULL) {
        goto StartAgain;
    }
    TheActualPacket = (PDEBUGGER_REMOTE_PACKET)BufferToReceive;
    if (TheActualPacket->Indicator == INDICATOR_OF_HYPERDBG_PACKET) {
        if (KdComputeDataChecksum((PVOID)&TheActualPacket->Indicator,
                                  LengthReceived - sizeof(BYTE)) !=
            TheActualPacket->Checksum) {
            ShowMessages("\nerr, checksum is invalid\n");
            goto StartAgain;
        }
        if (TheActualPacket->TypeOfThePacket !=
            DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGEE_TO_DEBUGGER) {
            ShowMessages("\nerr, unknown packet received from the debuggee\n");
            goto StartAgain;
        }
        switch (TheActualPacket->RequestedActionOfThePacket) {
        case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_STARTED:
            InitPacket =
                (DEBUGGER_PREPARE_DEBUGGEE *)(((CHAR *)TheActualPacket) +
                                              sizeof(DEBUGGER_REMOTE_PACKET));
            ShowMessages("connected to debuggee %s\n", InitPacket->OsName);
            DbgReceivedKernelResponse(DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_STARTED_PACKET_RECEIVED);
            break;
        case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_LOGGING_MECHANISM:
            MessagePacket =
                (DEBUGGEE_MESSAGE_PACKET *)(((CHAR *)TheActualPacket) +
                                            sizeof(DEBUGGER_REMOTE_PACKET));
            if (!g_IgnoreNewLoggingMessages) {
                ShowMessages("%s", MessagePacket->Message);
            }
            break;
        case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_PAUSED_AND_CURRENT_INSTRUCTION:
            g_IgnoreNewLoggingMessages = TRUE;
            PausePacket                = (DEBUGGEE_KD_PAUSED_PACKET *)(((CHAR *)TheActualPacket) +
                                                        sizeof(DEBUGGER_REMOTE_PACKET));
            g_IsDebuggeeRunning        = FALSE;
            g_CurrentRemoteCore        = PausePacket->CurrentCore;
            RtlZeroMemory(g_CurrentRunningInstruction, MAXIMUM_INSTR_SIZE);
            memcpy(g_CurrentRunningInstruction, &PausePacket->InstructionBytesOnRip, MAXIMUM_INSTR_SIZE);
            g_IsRunningInstruction32Bit = PausePacket->Is32BitAddress;
            switch (PausePacket->PausingReason) {
            case DEBUGGEE_PAUSING_REASON_DEBUGGEE_SOFTWARE_BREAKPOINT_HIT:
                if (PausePacket->EventTag != NULL) {
                    ShowMessages("breakpoint 0x%x hit\n",
                                 PausePacket->EventTag);
                }
                break;
            case DEBUGGEE_PAUSING_REASON_DEBUGGEE_EVENT_TRIGGERED:
                if (PausePacket->EventTag != NULL) {
                    ShowMessages("event 0x%x triggered\n",
                                 PausePacket->EventTag - DebuggerEventTagStartSeed);
                }
                break;
            case DEBUGGEE_PAUSING_REASON_DEBUGGEE_PROCESS_SWITCHED:
                ShowMessages("switched to the specified process\n");
                break;
            case DEBUGGEE_PAUSING_REASON_DEBUGGEE_THREAD_SWITCHED:
                ShowMessages("switched to the specified thread\n");
                break;
            case DEBUGGEE_PAUSING_REASON_DEBUGGEE_ENTRY_POINT_REACHED:
                ShowMessages("reached to the entrypoint of the user-mode process\n");
                break;
            default:
                break;
            }
            if (PausePacket->PausingReason !=
                DEBUGGEE_PAUSING_REASON_PAUSE_WITHOUT_DISASM) {
                if (PausePacket->ReadInstructionLen != MAXIMUM_INSTR_SIZE) {
                    if (HyperDbgLengthDisassemblerEngine(PausePacket->InstructionBytesOnRip,
                                                         MAXIMUM_INSTR_SIZE,
                                                         PausePacket->Is32BitAddress ? FALSE : TRUE) > PausePacket->ReadInstructionLen) {
                        ShowMessages("oOh, no! there might be a misinterpretation in disassembling the current instruction\n");
                    }
                }
                if (!PausePacket->Is32BitAddress) {
                    HyperDbgDisassembler64(PausePacket->InstructionBytesOnRip,
                                           PausePacket->Rip,
                                           MAXIMUM_INSTR_SIZE,
                                           1,
                                           TRUE,
                                           (PRFLAGS)&PausePacket->Rflags);
                } else {
                    HyperDbgDisassembler32(PausePacket->InstructionBytesOnRip,
                                           PausePacket->Rip,
                                           MAXIMUM_INSTR_SIZE,
                                           1,
                                           TRUE,
                                           (PRFLAGS)&PausePacket->Rflags);
                }
            }
            switch (PausePacket->PausingReason) {
            case DEBUGGEE_PAUSING_REASON_DEBUGGEE_SOFTWARE_BREAKPOINT_HIT:
            case DEBUGGEE_PAUSING_REASON_DEBUGGEE_HARDWARE_DEBUG_REGISTER_HIT:
            case DEBUGGEE_PAUSING_REASON_DEBUGGEE_EVENT_TRIGGERED:
            case DEBUGGEE_PAUSING_REASON_DEBUGGEE_STEPPED:
            case DEBUGGEE_PAUSING_REASON_DEBUGGEE_PROCESS_SWITCHED:
            case DEBUGGEE_PAUSING_REASON_DEBUGGEE_THREAD_SWITCHED:
                DbgReceivedKernelResponse(DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_IS_DEBUGGER_RUNNING);
                break;
            case DEBUGGEE_PAUSING_REASON_DEBUGGEE_ENTRY_POINT_REACHED:
                ShowMessages("\n");
                DbgReceivedKernelResponse(DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_IS_DEBUGGER_RUNNING);
                break;
            case DEBUGGEE_PAUSING_REASON_PAUSE_WITHOUT_DISASM:
                break;
            case DEBUGGEE_PAUSING_REASON_DEBUGGEE_CORE_SWITCHED:
                DbgReceivedKernelResponse(DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_CORE_SWITCHING_RESULT);
                break;
            case DEBUGGEE_PAUSING_REASON_DEBUGGEE_COMMAND_EXECUTION_FINISHED:
                ShowMessages("\n");
                DbgReceivedKernelResponse(DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_DEBUGGEE_FINISHED_COMMAND_EXECUTION);
                break;
            case DEBUGGEE_PAUSING_REASON_REQUEST_FROM_DEBUGGER:
                DbgReceivedKernelResponse(DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_PAUSED_DEBUGGEE_DETAILS);
                break;
            default:
                ShowMessages("err, unknown pausing reason is received\n");
                break;
            }
            break;
        case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_CHANGING_CORE:
            ChangeCorePacket =
                (DEBUGGEE_CHANGE_CORE_PACKET *)(((CHAR *)TheActualPacket) +
                                                sizeof(DEBUGGER_REMOTE_PACKET));
            if (ChangeCorePacket->Result == DEBUGGER_OPERATION_WAS_SUCCESSFULL) {
                ShowMessages("current operating core changed to 0x%x\n",
                             ChangeCorePacket->NewCore);
            } else {
                ShowErrorMessage(ChangeCorePacket->Result);
                DbgReceivedKernelResponse(DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_CORE_SWITCHING_RESULT);
            }
            break;
        case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_CHANGING_PROCESS:
            ChangeProcessPacket =
                (DEBUGGEE_DETAILS_AND_SWITCH_PROCESS_PACKET *)(((CHAR *)TheActualPacket) +
                                                               sizeof(DEBUGGER_REMOTE_PACKET));
            if (ChangeProcessPacket->Result == DEBUGGER_OPERATION_WAS_SUCCESSFULL) {
                if (ChangeProcessPacket->ActionType == DEBUGGEE_DETAILS_AND_SWITCH_PROCESS_GET_PROCESS_DETAILS) {
                    ShowMessages("process id: %x\nprocess (_EPROCESS): %s\nprocess name (16-Byte): %s\n",
                                 ChangeProcessPacket->ProcessId,
                                 SeparateTo64BitValue(ChangeProcessPacket->Process).c_str(),
                                 &ChangeProcessPacket->ProcessName);
                } else if (ChangeProcessPacket->ActionType == DEBUGGEE_DETAILS_AND_SWITCH_PROCESS_PERFORM_SWITCH) {
                    ShowMessages(
                        "press 'g' to continue the debuggee, if the pid or the "
                        "process object address is valid then the debuggee will "
                        "be automatically paused when it attached to the target process\n");
                }
            } else {
                ShowErrorMessage(ChangeProcessPacket->Result);
            }
            DbgReceivedKernelResponse(DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_PROCESS_SWITCHING_RESULT);
            break;
        case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RELOAD_SEARCH_QUERY:
            SearchResultsPacket =
                (DEBUGGEE_RESULT_OF_SEARCH_PACKET *)(((CHAR *)TheActualPacket) +
                                                     sizeof(DEBUGGER_REMOTE_PACKET));
            if (SearchResultsPacket->Result == DEBUGGER_OPERATION_WAS_SUCCESSFULL) {
                if (SearchResultsPacket->CountOfResults == 0) {
                    ShowMessages("not found\n");
                }
            } else {
                ShowErrorMessage(SearchResultsPacket->Result);
            }
            DbgReceivedKernelResponse(DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_SEARCH_QUERY_RESULT);
            break;
        case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_CHANGING_THREAD:
            ChangeThreadPacket =
                (DEBUGGEE_DETAILS_AND_SWITCH_THREAD_PACKET *)(((CHAR *)TheActualPacket) +
                                                              sizeof(DEBUGGER_REMOTE_PACKET));
            if (ChangeThreadPacket->Result == DEBUGGER_OPERATION_WAS_SUCCESSFULL) {
                if (ChangeThreadPacket->ActionType == DEBUGGEE_DETAILS_AND_SWITCH_THREAD_GET_THREAD_DETAILS) {
                    ShowMessages("thread id: %x (pid: %x)\nthread (_ETHREAD): %s\nprocess (_EPROCESS): %s\nprocess name (16-Byte): %s\n",
                                 ChangeThreadPacket->ThreadId,
                                 ChangeThreadPacket->ProcessId,
                                 SeparateTo64BitValue(ChangeThreadPacket->Thread).c_str(),
                                 SeparateTo64BitValue(ChangeThreadPacket->Process).c_str(),
                                 &ChangeThreadPacket->ProcessName);
                } else if (ChangeThreadPacket->ActionType == DEBUGGEE_DETAILS_AND_SWITCH_THREAD_PERFORM_SWITCH) {
                    ShowMessages(
                        "press 'g' to continue the debuggee, if the tid or the "
                        "thread object address is valid then the debuggee will "
                        "be automatically paused when it attached to the target thread\n");
                }
            } else {
                ShowErrorMessage(ChangeThreadPacket->Result);
            }
            DbgReceivedKernelResponse(DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_THREAD_SWITCHING_RESULT);
            break;
        case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_FLUSH:
            FlushPacket =
                (DEBUGGER_FLUSH_LOGGING_BUFFERS *)(((CHAR *)TheActualPacket) +
                                                   sizeof(DEBUGGER_REMOTE_PACKET));
            if (FlushPacket->KernelStatus == DEBUGGER_OPERATION_WAS_SUCCESSFULL) {
                ShowMessages("flushing buffers was successful, total %d messages were "
                             "cleared.\n",
                             FlushPacket->CountOfMessagesThatSetAsReadFromVmxNonRoot +
                                 FlushPacket->CountOfMessagesThatSetAsReadFromVmxRoot);
            } else {
                ShowErrorMessage(FlushPacket->KernelStatus);
            }
            DbgReceivedKernelResponse(DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_FLUSH_RESULT);
            break;
        case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_CALLSTACK:
            CallstackPacket =
                (DEBUGGER_CALLSTACK_REQUEST *)(((CHAR *)TheActualPacket) +
                                               sizeof(DEBUGGER_REMOTE_PACKET));
            CallstackFramePacket =
                (DEBUGGER_SINGLE_CALLSTACK_FRAME *)(((CHAR *)TheActualPacket) +
                                                    sizeof(DEBUGGER_REMOTE_PACKET) +
                                                    sizeof(DEBUGGER_CALLSTACK_REQUEST));
            if (CallstackPacket->KernelStatus == DEBUGGER_OPERATION_WAS_SUCCESSFULL) {
                CallstackShowFrames(CallstackFramePacket,
                                    CallstackPacket->FrameCount,
                                    CallstackPacket->DisplayMethod,
                                    CallstackPacket->Is32Bit);
            } else {
                ShowErrorMessage(CallstackPacket->KernelStatus);
            }
            DbgReceivedKernelResponse(DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_CALLSTACK_RESULT);
            break;
        case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_TEST_QUERY:
            TestQueryPacket =
                (DEBUGGER_DEBUGGER_TEST_QUERY_BUFFER *)(((CHAR *)TheActualPacket) +
                                                        sizeof(DEBUGGER_REMOTE_PACKET));
            if (TestQueryPacket->KernelStatus == DEBUGGER_OPERATION_WAS_SUCCESSFULL) {
            } else {
                ShowErrorMessage(TestQueryPacket->KernelStatus);
            }
            DbgReceivedKernelResponse(DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_TEST_QUERY);
            break;
        case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_RUNNING_SCRIPT:
            ScriptPacket = (DEBUGGEE_SCRIPT_PACKET *)(((CHAR *)TheActualPacket) +
                                                      sizeof(DEBUGGER_REMOTE_PACKET));
            if (ScriptPacket->Result == DEBUGGER_OPERATION_WAS_SUCCESSFULL) {
            } else {
                ShowErrorMessage(ScriptPacket->Result);
            }
            if (ScriptPacket->IsFormat) {
                DbgReceivedKernelResponse(DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_SCRIPT_FORMATS_RESULT);
            }
            DbgReceivedKernelResponse(DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_SCRIPT_RUNNING_RESULT);
            break;
        case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_FORMATS:
            FormatsPacket =
                (DEBUGGEE_FORMATS_PACKET *)(((CHAR *)TheActualPacket) +
                                            sizeof(DEBUGGER_REMOTE_PACKET));
            g_ErrorStateOfResultOfEvaluatedExpression = FormatsPacket->Result;
            g_ResultOfEvaluatedExpression             = FormatsPacket->Value;
            break;
        case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_REGISTERING_EVENT:
            EventAndActionPacket =
                (DEBUGGER_EVENT_AND_ACTION_REG_BUFFER *)(((CHAR *)TheActualPacket) +
                                                         sizeof(
                                                             DEBUGGER_REMOTE_PACKET));
            memcpy(&g_DebuggeeResultOfRegisteringEvent, EventAndActionPacket, sizeof(DEBUGGER_EVENT_AND_ACTION_REG_BUFFER));
            DbgReceivedKernelResponse(DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_REGISTER_EVENT);
            break;
        case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_ADDING_ACTION_TO_EVENT:
            EventAndActionPacket =
                (DEBUGGER_EVENT_AND_ACTION_REG_BUFFER *)(((CHAR *)TheActualPacket) +
                                                         sizeof(
                                                             DEBUGGER_REMOTE_PACKET));
            memcpy(&g_DebuggeeResultOfAddingActionsToEvent, EventAndActionPacket, sizeof(DEBUGGER_EVENT_AND_ACTION_REG_BUFFER));
            DbgReceivedKernelResponse(DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_ADD_ACTION_TO_EVENT);
            break;
        case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_QUERY_AND_MODIFY_EVENT:
            EventModifyAndQueryPacket =
                (DEBUGGER_MODIFY_EVENTS *)(((CHAR *)TheActualPacket) +
                                           sizeof(DEBUGGER_REMOTE_PACKET));
            if (EventModifyAndQueryPacket->KernelStatus !=
                DEBUGGER_OPERATION_WAS_SUCCESSFULL) {
                ShowErrorMessage(EventModifyAndQueryPacket->KernelStatus);
            } else if (EventModifyAndQueryPacket->TypeOfAction ==
                       DEBUGGER_MODIFY_EVENTS_QUERY_STATE) {
                g_SharedEventStatus = EventModifyAndQueryPacket->IsEnabled;
            } else {
                CommandEventsHandleModifiedEvent(EventModifyAndQueryPacket->Tag,
                                                 EventModifyAndQueryPacket);
            }
            DbgReceivedKernelResponse(DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_MODIFY_AND_QUERY_EVENT);
            break;
        case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RELOAD_SYMBOL_FINISHED:
            SymbolReloadFinishedPacket =
                (DEBUGGEE_SYMBOL_UPDATE_RESULT *)(((CHAR *)TheActualPacket) +
                                                  sizeof(DEBUGGER_REMOTE_PACKET));
            if (SymbolReloadFinishedPacket->KernelStatus !=
                DEBUGGER_OPERATION_WAS_SUCCESSFULL) {
                ShowErrorMessage(SymbolReloadFinishedPacket->KernelStatus);
            } else {
                SymbolInitialReload();
            }
            DbgReceivedKernelResponse(DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_SYMBOL_RELOAD);
            break;
        case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_READING_REGISTERS:
            ReadRegisterPacket =
                (DEBUGGEE_REGISTER_READ_DESCRIPTION *)(((CHAR *)TheActualPacket) +
                                                       sizeof(
                                                           DEBUGGER_REMOTE_PACKET));
            if (ReadRegisterPacket->KernelStatus ==
                DEBUGGER_OPERATION_WAS_SUCCESSFULL) {
                if (ReadRegisterPacket->RegisterID == DEBUGGEE_SHOW_ALL_REGISTERS) {
                    Regs          = (GUEST_REGS *)(((CHAR *)ReadRegisterPacket) +
                                          sizeof(DEBUGGEE_REGISTER_READ_DESCRIPTION));
                    ExtraRegs     = (GUEST_EXTRA_REGISTERS *)(((CHAR *)ReadRegisterPacket) +
                                                          sizeof(DEBUGGEE_REGISTER_READ_DESCRIPTION) +
                                                          sizeof(GUEST_REGS));
                    RFLAGS Rflags = {0};
                    Rflags.AsUInt = ExtraRegs->RFLAGS;
                    ShowMessages(
                        "RAX=%016llx RBX=%016llx RCX=%016llx\n"
                        "RDX=%016llx RSI=% 016llx RDI=%016llx\n"
                        "RIP=%016llx RSP=%016llx RBP=%016llx\n"
                        "R8=%016llx  R9=%016llx  R10=%016llx\n"
                        "R11=%016llx R12=%016llx R13=%016llx\n"
                        "R14=%016llx R15=%016llx IOPL=%02x\n"
                        "%s  %s  %s  %s\n%s  %s  %s  %s  \n"
                        "CS %04x SS %04x DS %04x ES %04x FS %04x GS %04x\n"
                        "RFLAGS=%016llx\n",
                        Regs->rax,
                        Regs->rbx,
                        Regs->rcx,
                        Regs->rdx,
                        Regs->rsi,
                        Regs->rdi,
                        ExtraRegs->RIP,
                        Regs->rsp,
                        Regs->rbp,
                        Regs->r8,
                        Regs->r9,
                        Regs->r10,
                        Regs->r11,
                        Regs->r12,
                        Regs->r13,
                        Regs->r14,
                        Regs->r15,
                        Rflags.IoPrivilegeLevel,
                        Rflags.OverflowFlag ? "OF 1" : "OF 0",
                        Rflags.DirectionFlag ? "DF 1" : "DF 0",
                        Rflags.InterruptEnableFlag ? "IF 1" : "IF 0",
                        Rflags.SignFlag ? "SF  1" : "SF  0",
                        Rflags.ZeroFlag ? "ZF 1" : "ZF 0",
                        Rflags.ParityFlag ? "PF 1" : "PF 0",
                        Rflags.CarryFlag ? "CF 1" : "CF 0",
                        Rflags.AuxiliaryCarryFlag ? "AXF 1" : "AXF 0",
                        ExtraRegs->CS,
                        ExtraRegs->SS,
                        ExtraRegs->DS,
                        ExtraRegs->ES,
                        ExtraRegs->FS,
                        ExtraRegs->GS,
                        ExtraRegs->RFLAGS);
                } else {
                    ShowMessages("%s=%016llx\n",
                                 RegistersNames[ReadRegisterPacket->RegisterID],
                                 ReadRegisterPacket->Value);
                }
            } else {
                ShowErrorMessage(ReadRegisterPacket->KernelStatus);
            }
            DbgReceivedKernelResponse(DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_READ_REGISTERS);
            break;
        case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_READING_MEMORY:
            ReadMemoryPacket =
                (DEBUGGER_READ_MEMORY *)(((CHAR *)TheActualPacket) +
                                         sizeof(DEBUGGER_REMOTE_PACKET));
            if (ReadMemoryPacket->KernelStatus ==
                DEBUGGER_OPERATION_WAS_SUCCESSFULL) {
                MemoryBuffer = (unsigned char *)(((CHAR *)TheActualPacket) +
                                                 sizeof(DEBUGGER_REMOTE_PACKET) +
                                                 sizeof(DEBUGGER_READ_MEMORY));
                switch (ReadMemoryPacket->Style) {
                case DEBUGGER_SHOW_COMMAND_DISASSEMBLE64:
                    HyperDbgDisassembler64(MemoryBuffer, ReadMemoryPacket->Address, ReadMemoryPacket->ReturnLength, 0, FALSE, NULL);
                    break;
                case DEBUGGER_SHOW_COMMAND_DISASSEMBLE32:
                    HyperDbgDisassembler32(MemoryBuffer, ReadMemoryPacket->Address, ReadMemoryPacket->ReturnLength, 0, FALSE, NULL);
                    break;
                case DEBUGGER_SHOW_COMMAND_DB:
                    ShowMemoryCommandDB(
                        MemoryBuffer,
                        ReadMemoryPacket->Size,
                        ReadMemoryPacket->Address,
                        ReadMemoryPacket->MemoryType,
                        ReadMemoryPacket->ReturnLength);
                    break;
                case DEBUGGER_SHOW_COMMAND_DC:
                    ShowMemoryCommandDC(
                        MemoryBuffer,
                        ReadMemoryPacket->Size,
                        ReadMemoryPacket->Address,
                        ReadMemoryPacket->MemoryType,
                        ReadMemoryPacket->ReturnLength);
                    break;
                case DEBUGGER_SHOW_COMMAND_DD:
                    ShowMemoryCommandDD(
                        MemoryBuffer,
                        ReadMemoryPacket->Size,
                        ReadMemoryPacket->Address,
                        ReadMemoryPacket->MemoryType,
                        ReadMemoryPacket->ReturnLength);
                    break;
                case DEBUGGER_SHOW_COMMAND_DQ:
                    ShowMemoryCommandDQ(
                        MemoryBuffer,
                        ReadMemoryPacket->Size,
                        ReadMemoryPacket->Address,
                        ReadMemoryPacket->MemoryType,
                        ReadMemoryPacket->ReturnLength);
                    break;
                case DEBUGGER_SHOW_COMMAND_DT:
                    ScriptEngineShowDataBasedOnSymbolTypesWrapper(ReadMemoryPacket->DtDetails->TypeName,
                                                                  ReadMemoryPacket->Address,
                                                                  FALSE,
                                                                  MemoryBuffer,
                                                                  ReadMemoryPacket->DtDetails->AdditionalParameters);
                    break;
                }
            } else {
                ShowErrorMessage(ReadMemoryPacket->KernelStatus);
            }
            DbgReceivedKernelResponse(DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_READ_MEMORY);
            break;
        case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_EDITING_MEMORY:
            EditMemoryPacket =
                (DEBUGGER_EDIT_MEMORY *)(((CHAR *)TheActualPacket) +
                                         sizeof(DEBUGGER_REMOTE_PACKET));
            if (EditMemoryPacket->KernelStatus ==
                DEBUGGER_OPERATION_WAS_SUCCESSFULL) {
            } else {
                ShowErrorMessage(EditMemoryPacket->KernelStatus);
            }
            DbgReceivedKernelResponse(DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_EDIT_MEMORY);
            break;
        case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_BP:
            BpPacket = (DEBUGGEE_BP_PACKET *)(((CHAR *)TheActualPacket) +
                                              sizeof(DEBUGGER_REMOTE_PACKET));
            if (BpPacket->Result == DEBUGGER_OPERATION_WAS_SUCCESSFULL) {
            } else {
                ShowErrorMessage(BpPacket->Result);
            }
            DbgReceivedKernelResponse(DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_BP);
            break;
        case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_PTE:
            PtePacket = (DEBUGGER_READ_PAGE_TABLE_ENTRIES_DETAILS *)(((CHAR *)TheActualPacket) +
                                                                     sizeof(DEBUGGER_REMOTE_PACKET));
            if (PtePacket->KernelStatus == DEBUGGER_OPERATION_WAS_SUCCESSFULL) {
                CommandPteShowResults(PtePacket->VirtualAddress, PtePacket);
            } else {
                ShowErrorMessage(PtePacket->KernelStatus);
            }
            DbgReceivedKernelResponse(DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_PTE_RESULT);
            break;
        case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_VA2PA_AND_PA2VA:
            Va2paPa2vaPacket = (DEBUGGER_VA2PA_AND_PA2VA_COMMANDS *)(((CHAR *)TheActualPacket) +
                                                                     sizeof(DEBUGGER_REMOTE_PACKET));
            if (Va2paPa2vaPacket->KernelStatus == DEBUGGER_OPERATION_WAS_SUCCESSFULL) {
                if (Va2paPa2vaPacket->IsVirtual2Physical) {
                    ShowMessages("%llx\n", Va2paPa2vaPacket->PhysicalAddress);
                } else {
                    ShowMessages("%llx\n", Va2paPa2vaPacket->VirtualAddress);
                }
            } else {
                ShowErrorMessage(Va2paPa2vaPacket->KernelStatus);
            }
            DbgReceivedKernelResponse(DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_VA2PA_AND_PA2VA_RESULT);
            break;
        case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_LIST_OR_MODIFY_BREAKPOINTS:
            ListOrModifyBreakpointPacket =
                (DEBUGGEE_BP_LIST_OR_MODIFY_PACKET *)(((CHAR *)TheActualPacket) +
                                                      sizeof(DEBUGGER_REMOTE_PACKET));
            if (ListOrModifyBreakpointPacket->Result ==
                DEBUGGER_OPERATION_WAS_SUCCESSFULL) {
            } else {
                ShowErrorMessage(ListOrModifyBreakpointPacket->Result);
            }
            DbgReceivedKernelResponse(DEBUGGER_SYNCRONIZATION_OBJECT_KERNEL_DEBUGGER_LIST_OR_MODIFY_BREAKPOINTS);
            break;
        case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_UPDATE_SYMBOL_INFO:
            SymbolUpdatePacket =
                (DEBUGGER_UPDATE_SYMBOL_TABLE *)(((CHAR *)TheActualPacket) +
                                                 sizeof(
                                                     DEBUGGER_REMOTE_PACKET));
            SymbolBuildAndUpdateSymbolTable(&SymbolUpdatePacket->SymbolDetailPacket);
            break;
        default:
            ShowMessages("err, unknown packet action received from the debugger\n");
            break;
        }
    } else {
        ShowMessages("err, invalid packet received\n");
    }
    goto StartAgain;
    return TRUE;
}

BOOLEAN
ListeningSerialPortInDebuggee() {
StartAgain:
    BOOL Status; /* Status */
    char SerialBuffer[MaxSerialPacketSize] = {
        0};                                     /* Buffer to send and receive data */
    DWORD                   EventMask   = 0;    /* Event mask to trigger */
    char                    ReadData    = NULL; /* temperory Character */
    DWORD                   NoBytesRead = 0;    /* Bytes read by ReadFile() */
    UINT32                  Loop        = 0;
    PDEBUGGER_REMOTE_PACKET TheActualPacket =
        (PDEBUGGER_REMOTE_PACKET)SerialBuffer;
    Status = SetCommMask(g_SerialRemoteComPortHandle, EV_RXCHAR);
    if (Status == FALSE) {
    }
    Status = WaitCommEvent(g_SerialRemoteComPortHandle, &EventMask, NULL); /* Wait for the character to be received */
    if (Status == FALSE) {
    }
    do {
        Status = ReadFile(g_SerialRemoteComPortHandle, &ReadData, sizeof(ReadData), &NoBytesRead, NULL);
        if (!(MaxSerialPacketSize > Loop)) {
            ShowMessages("err, a buffer received in debuggee which exceeds the "
                         "buffer limitation\n");
            goto StartAgain;
        }
        SerialBuffer[Loop] = ReadData;
        if (KdCheckForTheEndOfTheBuffer(&Loop, (BYTE *)SerialBuffer)) {
            break;
        }
        ++Loop;
    } while (NoBytesRead > 0);
    if (Loop == 1 && SerialBuffer[0] == NULL) {
        goto StartAgain;
    }
    if (TheActualPacket->Indicator == INDICATOR_OF_HYPERDBG_PACKET) {
        if (KdComputeDataChecksum((PVOID)&TheActualPacket->Indicator,
                                  Loop - sizeof(BYTE)) !=
            TheActualPacket->Checksum) {
            ShowMessages("err checksum is invalid\n");
            goto StartAgain;
        }
        if (TheActualPacket->TypeOfThePacket !=
            DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGER_TO_DEBUGGEE_EXECUTE_ON_USER_MODE) {
            ShowMessages("err, unknown packet received from the debugger\n");
            goto StartAgain;
        }
        switch (TheActualPacket->RequestedActionOfThePacket) {
        case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_USER_MODE_PAUSE:
            if (!DebuggerPauseDebuggee()) {
                ShowMessages("err, debugger tries to pause the debuggee but the "
                             "attempt was unsuccessful\n");
            }
            break;
        case DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_USER_MODE_DO_NOT_READ_ANY_PACKET:
            return TRUE;
            break;
        default:
            ShowMessages("err, unknown packet action received from the debugger\n");
            break;
        }
    } else {
        DebugBreak();
    }
    goto StartAgain;
    return TRUE;
}

DWORD WINAPI
ListeningSerialPauseDebuggerThread(PVOID Param) {
    ListeningSerialPortInDebugger();
    return 0;
}

DWORD WINAPI
ListeningSerialPauseDebuggeeThread(PVOID Param) {
    ListeningSerialPortInDebuggee();
    return 0;
}
