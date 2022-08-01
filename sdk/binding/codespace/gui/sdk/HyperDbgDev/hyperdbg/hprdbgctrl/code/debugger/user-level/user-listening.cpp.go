package user-level
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\debugger\user-level\user-listening.cpp.back

type (
	UserListening interface {
		UdHandleUserDebuggerPausing() (ok bool) //col:54
	}
	userListening struct{}
)

func NewUserListening() UserListening { return &userListening{} }

func (u *userListening) UdHandleUserDebuggerPausing() (ok bool) { //col:54
	/*UdHandleUserDebuggerPausing(PDEBUGGEE_UD_PAUSED_PACKET PausePacket)
	  {
	      UdSetActiveDebuggingProcess(PausePacket->ProcessDebuggingToken,
	                                  PausePacket->ProcessId,
	                                  PausePacket->ThreadId,
	                                  PausePacket->Is32Bit,
	                                  TRUE);
	      switch (PausePacket->PausingReason)
	      {
	      case DEBUGGEE_PAUSING_REASON_DEBUGGEE_ENTRY_POINT_REACHED:
	          ShowMessages("reached to the entrypoint of the main module\n");
	          break;
	      case DEBUGGEE_PAUSING_REASON_DEBUGGEE_GENERAL_THREAD_INTERCEPTED:
	          ShowMessages("\nthread: %x from process: %x intercepted\n",
	                       PausePacket->ThreadId,
	                       PausePacket->ProcessId);
	          break;
	      default:
	          break;
	      }
	      if (PausePacket->ReadInstructionLen != MAXIMUM_INSTR_SIZE)
	      {
	          if (HyperDbgLengthDisassemblerEngine(PausePacket->InstructionBytesOnRip,
	                                               MAXIMUM_INSTR_SIZE,
	                                               PausePacket->Is32Bit ? FALSE : TRUE) > PausePacket->ReadInstructionLen)
	          {
	              ShowMessages("oOh, no! there might be a misinterpretation in disassembling the current instruction\n");
	          }
	      }
	      if (!PausePacket->Is32Bit)
	      {
	          HyperDbgDisassembler64(PausePacket->InstructionBytesOnRip,
	                                 PausePacket->Rip,
	                                 MAXIMUM_INSTR_SIZE,
	                                 1,
	                                 TRUE,
	                                 (PRFLAGS)&PausePacket->Rflags);
	      }
	      else
	      {
	          HyperDbgDisassembler32(PausePacket->InstructionBytesOnRip,
	                                 PausePacket->Rip,
	                                 MAXIMUM_INSTR_SIZE,
	                                 1,
	                                 TRUE,
	                                 (PRFLAGS)&PausePacket->Rflags);
	      }
	      if (g_UserSyncronizationObjectsHandleTable
	              [DEBUGGER_SYNCRONIZATION_OBJECT_USER_DEBUGGER_IS_DEBUGGER_RUNNING]
	                  .IsOnWaitingState == TRUE)
	      {
	          DbgReceivedUserResponse(DEBUGGER_SYNCRONIZATION_OBJECT_USER_DEBUGGER_IS_DEBUGGER_RUNNING);
	      }
	  }*/
	return true
}
