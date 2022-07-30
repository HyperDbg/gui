package vmx

//back\HyperDbgDev\hyperdbg\hprdbghv\code\vmm\vmx\Events.c.back

type (
	Events interface {
		EventInjectInterruption() (ok bool)      //col:37
		EventInjectBreakpoint() (ok bool)        //col:51
		EventInjectGeneralProtection() (ok bool) //col:65
		EventInjectUndefinedOpcode() (ok bool)   //col:81
		EventInjectDebugBreakpoint() (ok bool)   //col:92
		EventInjectPageFault() (ok bool)         //col:123
	}
)

func NewEvents() { return &events{} }

func (e *events) EventInjectInterruption() (ok bool) { //col:37
	/*EventInjectInterruption(INTERRUPT_TYPE InterruptionType, EXCEPTION_VECTORS Vector, BOOLEAN DeliverErrorCode, UINT32 ErrorCode)
	  {
	      INTERRUPT_INFO Inject       = {0};
	      Inject.Fields.Valid         = TRUE;
	      Inject.Fields.InterruptType = InterruptionType;
	      Inject.Fields.Vector        = Vector;
	      Inject.Fields.DeliverCode   = DeliverErrorCode;
	      __vmx_vmwrite(VMCS_CTRL_VMENTRY_INTERRUPTION_INFORMATION_FIELD, Inject.Flags);
	      if (DeliverErrorCode)
	      {
	          __vmx_vmwrite(VMCS_CTRL_VMENTRY_EXCEPTION_ERROR_CODE, ErrorCode);
	      }
	  }*/
	return true
}

func (e *events) EventInjectBreakpoint() (ok bool) { //col:51
	/*EventInjectBreakpoint()
	  {
	      EventInjectInterruption(INTERRUPT_TYPE_SOFTWARE_EXCEPTION, EXCEPTION_VECTOR_BREAKPOINT, FALSE, 0);
	      UINT32 ExitInstrLength;
	      __vmx_vmread(VMCS_VMEXIT_INSTRUCTION_LENGTH, &ExitInstrLength);
	      __vmx_vmwrite(VMCS_CTRL_VMENTRY_INSTRUCTION_LENGTH, ExitInstrLength);
	  }*/
	return true
}

func (e *events) EventInjectGeneralProtection() (ok bool) { //col:65
	/*EventInjectGeneralProtection()
	  {
	      EventInjectInterruption(INTERRUPT_TYPE_HARDWARE_EXCEPTION, EXCEPTION_VECTOR_GENERAL_PROTECTION_FAULT, TRUE, 0);
	      UINT32 ExitInstrLength;
	      __vmx_vmread(VMCS_VMEXIT_INSTRUCTION_LENGTH, &ExitInstrLength);
	      __vmx_vmwrite(VMCS_CTRL_VMENTRY_INSTRUCTION_LENGTH, ExitInstrLength);
	  }*/
	return true
}

func (e *events) EventInjectUndefinedOpcode() (ok bool) { //col:81
	/*EventInjectUndefinedOpcode(UINT32 CurrentProcessorIndex)
	  {
	      EventInjectInterruption(INTERRUPT_TYPE_HARDWARE_EXCEPTION, EXCEPTION_VECTOR_UNDEFINED_OPCODE, FALSE, 0);
	      g_GuestState[CurrentProcessorIndex].IncrementRip = FALSE;
	  }*/
	return true
}

func (e *events) EventInjectDebugBreakpoint() (ok bool) { //col:92
	/*EventInjectDebugBreakpoint()
	  {
	      EventInjectInterruption(INTERRUPT_TYPE_HARDWARE_EXCEPTION, EXCEPTION_VECTOR_DEBUG_BREAKPOINT, FALSE, 0);
	  }*/
	return true
}

func (e *events) EventInjectPageFault() (ok bool) { //col:123
	/*EventInjectPageFault(UINT64 PageFaultAddress)
	  {
	      PAGE_FAULT_ERROR_CODE ErrorCode = {0};
	      __writecr2(PageFaultAddress);
	      ErrorCode.Fields.Fetch    = 0;
	      ErrorCode.Fields.Present  = 0;
	      ErrorCode.Fields.Reserved = 0;
	      ErrorCode.Fields.User     = 0;
	      ErrorCode.Fields.Write    = 0;
	      EventInjectInterruption(INTERRUPT_TYPE_HARDWARE_EXCEPTION, EXCEPTION_VECTOR_PAGE_FAULT, TRUE, ErrorCode.Flags);
	  }*/
	return true
}
