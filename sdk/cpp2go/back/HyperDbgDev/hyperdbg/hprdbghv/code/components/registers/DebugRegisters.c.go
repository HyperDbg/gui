package registers

//back\HyperDbgDev\hyperdbg\hprdbghv\code\components\registers\DebugRegisters.c.back

type (
	DebugRegisters interface {
		DebugRegistersSet() (ok bool) //col:259
	}
)

func NewDebugRegisters() { return &debugRegisters{} }

func (d *debugRegisters) DebugRegistersSet() (ok bool) { //col:259
	/*DebugRegistersSet(UINT32 DebugRegNum, DEBUG_REGISTER_TYPE ActionType, BOOLEAN ApplyToVmcs, UINT64 TargetAddress)
	  {
	      DR7 Dr7 = {0};
	      if (DebugRegNum >= 4)
	      {
	          return FALSE;
	      }
	      Dr7.Reserved1 = 1;
	      Dr7.LocalExactBreakpoint  = 1;
	      Dr7.GlobalExactBreakpoint = 1;
	      if (DebugRegNum == 0)
	      {
	          __writedr(0, TargetAddress);
	          Dr7.GlobalBreakpoint0 = 1;
	          switch (ActionType)
	          {
	          case BREAK_ON_INSTRUCTION_FETCH:
	              break;
	          case BREAK_ON_WRITE_ONLY:
	              break;
	          case BREAK_ON_IO_READ_OR_WRITE_NOT_SUPPORTED:
	              LogError("Err, I/O access breakpoint by debug regs are not supported");
	              return FALSE;
	              break;
	          case BREAK_ON_READ_AND_WRITE_BUT_NOT_FETCH:
	              break;
	          default:
	              LogError("Err, unknown parameter as debug reg action type");
	              return FALSE;
	              break;
	          }
	      }
	      else if (DebugRegNum == 1)
	      {
	          __writedr(1, TargetAddress);
	          Dr7.GlobalBreakpoint1 = 1;
	          switch (ActionType)
	          {
	          case BREAK_ON_INSTRUCTION_FETCH:
	              break;
	          case BREAK_ON_WRITE_ONLY:
	              break;
	          case BREAK_ON_IO_READ_OR_WRITE_NOT_SUPPORTED:
	              LogError("Err, I/O access breakpoint by debug regs are not supported");
	              return FALSE;
	              break;
	          case BREAK_ON_READ_AND_WRITE_BUT_NOT_FETCH:
	              break;
	          default:
	              LogError("Err, unknown parameter as debug reg action type");
	              return FALSE;
	              break;
	          }
	      }
	      else if (DebugRegNum == 2)
	      {
	          __writedr(2, TargetAddress);
	          Dr7.GlobalBreakpoint2 = 1;
	          switch (ActionType)
	          {
	          case BREAK_ON_INSTRUCTION_FETCH:
	              break;
	          case BREAK_ON_WRITE_ONLY:
	              break;
	          case BREAK_ON_IO_READ_OR_WRITE_NOT_SUPPORTED:
	              LogError("Err, I/O access breakpoint by debug regs are not supported");
	              return FALSE;
	              break;
	          case BREAK_ON_READ_AND_WRITE_BUT_NOT_FETCH:
	              break;
	          default:
	              LogError("Err, unknown parameter as debug reg action type");
	              return FALSE;
	              break;
	          }
	      }
	      else if (DebugRegNum == 3)
	      {
	          __writedr(3, TargetAddress);
	          Dr7.GlobalBreakpoint3 = 1;
	          switch (ActionType)
	          {
	          case BREAK_ON_INSTRUCTION_FETCH:
	              break;
	          case BREAK_ON_WRITE_ONLY:
	              break;
	          case BREAK_ON_IO_READ_OR_WRITE_NOT_SUPPORTED:
	              LogError("Err, I/O access breakpoint by debug regs are not supported");
	              return FALSE;
	              break;
	          case BREAK_ON_READ_AND_WRITE_BUT_NOT_FETCH:
	              break;
	          default:
	              LogError("Err, unknown parameter as debug reg action type");
	              return FALSE;
	              break;
	          }
	      }
	      if (ApplyToVmcs)
	      {
	          __vmx_vmwrite(VMCS_GUEST_DR7, Dr7.AsUInt);
	      }
	      else
	      {
	          __writedr(7, Dr7.AsUInt);
	      }
	      return TRUE;
	  }*/
	return true
}
