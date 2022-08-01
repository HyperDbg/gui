package vmx

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbghv\code\vmm\vmx\Vmexit.c.back

type (
	Vmexit interface {
		VmxVmexitHandler() (ok bool) //col:216
	}
	vmexit struct{}
)

func NewVmexit() Vmexit { return &vmexit{} }

func (v *vmexit) VmxVmexitHandler() (ok bool) { //col:216
	/*
	   VmxVmexitHandler(_Inout_ PGUEST_REGS GuestRegs)

	   	{
	   	    VMEXIT_INTERRUPT_INFORMATION          InterruptExit         = {0};
	   	    VMX_EXIT_QUALIFICATION_IO_INSTRUCTION IoQualification       = {0};
	   	    RFLAGS                                Flags                 = {0};
	   	    UINT64                                GuestPhysicalAddr     = 0;
	   	    UINT64                                GuestRsp              = 0;
	   	    UINT64                                GuestRip              = 0;
	   	    ULONG                                 ExitReason            = 0;
	   	    ULONG                                 ExitQualification     = 0;
	   	    ULONG                                 Rflags                = 0;
	   	    ULONG                                 EcxReg                = 0;
	   	    ULONG                                 ExitInstructionLength = 0;
	   	    ULONG                                 CurrentProcessorIndex = 0;
	   	    BOOLEAN                               Result                = FALSE;
	   	    BOOLEAN                               ShouldEmulateRdtscp   = TRUE;
	   	    VIRTUAL_MACHINE_STATE *               CurrentGuestState     = NULL;
	   	    CurrentProcessorIndex = KeGetCurrentProcessorNumber();
	   	    CurrentGuestState     = &g_GuestState[CurrentProcessorIndex];
	   	    CurrentGuestState->IsOnVmxRootMode = TRUE;
	   	    __vmx_vmread(VMCS_EXIT_REASON, &ExitReason);
	   	    ExitReason &= 0xffff;
	   	    if (g_TransparentMode)
	   	    {
	   	        ShouldEmulateRdtscp = TransparentModeStart(GuestRegs, CurrentProcessorIndex, ExitReason);
	   	    }
	   	    CurrentGuestState->IncrementRip = TRUE;
	   	    __vmx_vmread(VMCS_GUEST_RIP, &GuestRip);
	   	    CurrentGuestState->LastVmexitRip = GuestRip;
	   	    __vmx_vmread(VMCS_GUEST_RSP, &GuestRsp);
	   	    GuestRegs->rsp = GuestRsp;
	   	    __vmx_vmread(VMCS_EXIT_QUALIFICATION, &ExitQualification);
	   	    switch (ExitReason)
	   	    {
	   	    case VMX_EXIT_REASON_TRIPLE_FAULT:
	   	    {
	   	        LogError("Err, triple fault error occurred");
	   	        break;
	   	    }
	   	    case VMX_EXIT_REASON_EXECUTE_VMCLEAR:
	   	    case VMX_EXIT_REASON_EXECUTE_VMPTRLD:
	   	    case VMX_EXIT_REASON_EXECUTE_VMPTRST:
	   	    case VMX_EXIT_REASON_EXECUTE_VMREAD:
	   	    case VMX_EXIT_REASON_EXECUTE_VMRESUME:
	   	    case VMX_EXIT_REASON_EXECUTE_VMWRITE:
	   	    case VMX_EXIT_REASON_EXECUTE_VMXOFF:
	   	    case VMX_EXIT_REASON_EXECUTE_VMXON:
	   	    case VMX_EXIT_REASON_EXECUTE_VMLAUNCH:
	   	    {
	   	        EventInjectUndefinedOpcode(CurrentProcessorIndex);
	   	        break;
	   	    }
	   	    case VMX_EXIT_REASON_EXECUTE_INVEPT:
	   	    case VMX_EXIT_REASON_EXECUTE_INVVPID:
	   	    case VMX_EXIT_REASON_EXECUTE_GETSEC:
	   	    case VMX_EXIT_REASON_EXECUTE_INVD:
	   	    {
	   	        EventInjectUndefinedOpcode(CurrentProcessorIndex);
	   	        break;
	   	    }
	   	    case VMX_EXIT_REASON_MOV_CR:
	   	    {
	   	        HvHandleControlRegisterAccess(GuestRegs, CurrentProcessorIndex);
	   	        break;
	   	    }
	   	    case VMX_EXIT_REASON_EXECUTE_RDMSR:
	   	    {
	   	        EcxReg = GuestRegs->rcx & 0xffffffff;
	   	        MsrHandleRdmsrVmexit(GuestRegs);
	   	        DebuggerTriggerEvents(RDMSR_INSTRUCTION_EXECUTION, GuestRegs, EcxReg);
	   	        break;
	   	    }
	   	    case VMX_EXIT_REASON_EXECUTE_WRMSR:
	   	    {
	   	        EcxReg = GuestRegs->rcx & 0xffffffff;
	   	        MsrHandleWrmsrVmexit(GuestRegs);
	   	        DebuggerTriggerEvents(WRMSR_INSTRUCTION_EXECUTION, GuestRegs, EcxReg);
	   	        break;
	   	    }
	   	    case VMX_EXIT_REASON_EXECUTE_CPUID:
	   	    {
	   	        HvHandleCpuid(GuestRegs);
	   	        break;
	   	    }
	   	    case VMX_EXIT_REASON_EXECUTE_IO_INSTRUCTION:
	   	    {
	   	        __vmx_vmread(VMCS_EXIT_QUALIFICATION, &IoQualification);
	   	        __vmx_vmread(VMCS_GUEST_RFLAGS, &Flags);
	   	        IoHandleIoVmExits(GuestRegs, IoQualification, Flags);
	   	        if (IoQualification.DirectionOfAccess == AccessIn)
	   	        {
	   	            DebuggerTriggerEvents(IN_INSTRUCTION_EXECUTION, GuestRegs, IoQualification.PortNumber);
	   	        }
	   	        else if (IoQualification.DirectionOfAccess == AccessOut)
	   	        {
	   	            DebuggerTriggerEvents(OUT_INSTRUCTION_EXECUTION, GuestRegs, IoQualification.PortNumber);
	   	        }
	   	        break;
	   	    }
	   	    case VMX_EXIT_REASON_EPT_VIOLATION:
	   	    {
	   	        __vmx_vmread(VMCS_GUEST_PHYSICAL_ADDRESS, &GuestPhysicalAddr);
	   	        if (EptHandleEptViolation(GuestRegs, ExitQualification, GuestPhysicalAddr) == FALSE)
	   	        {
	   	            LogError("Err, there were errors in handling EPT violation");
	   	        }
	   	        break;
	   	    }
	   	    case VMX_EXIT_REASON_EPT_MISCONFIGURATION:
	   	    {
	   	        __vmx_vmread(VMCS_GUEST_PHYSICAL_ADDRESS, &GuestPhysicalAddr);
	   	        EptHandleMisconfiguration(GuestPhysicalAddr);
	   	        break;
	   	    }
	   	    case VMX_EXIT_REASON_EXECUTE_VMCALL:
	   	    {
	   	        VmxHandleVmcallVmExit(CurrentProcessorIndex, GuestRegs);
	   	        break;
	   	    }
	   	    case VMX_EXIT_REASON_EXCEPTION_OR_NMI:
	   	    {
	   	        __vmx_vmread(VMCS_VMEXIT_INTERRUPTION_INFORMATION, &InterruptExit);
	   	        IdtEmulationHandleExceptionAndNmi(CurrentProcessorIndex, InterruptExit, GuestRegs);
	   	        break;
	   	    }
	   	    case VMX_EXIT_REASON_EXTERNAL_INTERRUPT:
	   	    {
	   	        __vmx_vmread(VMCS_VMEXIT_INTERRUPTION_INFORMATION, &InterruptExit);
	   	        IdtEmulationHandleExternalInterrupt(CurrentProcessorIndex, InterruptExit, GuestRegs);
	   	        break;
	   	    }
	   	    case VMX_EXIT_REASON_INTERRUPT_WINDOW:
	   	    {
	   	        IdtEmulationHandleInterruptWindowExiting(CurrentProcessorIndex);
	   	        break;
	   	    }
	   	    case VMX_EXIT_REASON_NMI_WINDOW:
	   	    {
	   	        IdtEmulationHandleNmiWindowExiting(CurrentProcessorIndex, GuestRegs);
	   	        break;
	   	    }
	   	    case VMX_EXIT_REASON_MONITOR_TRAP_FLAG:
	   	    {
	   	        MtfHandleVmexit(CurrentProcessorIndex, GuestRegs);
	   	        break;
	   	    }
	   	    case VMX_EXIT_REASON_EXECUTE_HLT:
	   	    {
	   	        break;
	   	    }
	   	    case VMX_EXIT_REASON_EXECUTE_RDTSC:
	   	    {
	   	        if (ShouldEmulateRdtscp)
	   	        {
	   	            CounterEmulateRdtsc(GuestRegs);
	   	            DebuggerTriggerEvents(TSC_INSTRUCTION_EXECUTION, GuestRegs, FALSE);
	   	        }
	   	        break;
	   	    }
	   	    case VMX_EXIT_REASON_EXECUTE_RDTSCP:
	   	    {
	   	        if (ShouldEmulateRdtscp)
	   	        {
	   	            CounterEmulateRdtscp(GuestRegs);
	   	            DebuggerTriggerEvents(TSC_INSTRUCTION_EXECUTION, GuestRegs, TRUE);
	   	        }
	   	        break;
	   	    }
	   	    case VMX_EXIT_REASON_EXECUTE_RDPMC:
	   	    {
	   	        CounterEmulateRdpmc(GuestRegs);
	   	        DebuggerTriggerEvents(PMC_INSTRUCTION_EXECUTION, GuestRegs, NULL);
	   	        break;
	   	    }
	   	    case VMX_EXIT_REASON_MOV_DR:
	   	    {
	   	        if (!CurrentGuestState->DebuggingState.ThreadOrProcessTracingDetails.DebugRegisterInterceptionState)
	   	        {
	   	            HvHandleMovDebugRegister(CurrentProcessorIndex, GuestRegs);
	   	            DebuggerTriggerEvents(DEBUG_REGISTERS_ACCESSED, GuestRegs, NULL);
	   	        }
	   	        break;
	   	    }
	   	    case VMX_EXIT_REASON_EXECUTE_XSETBV:
	   	    {
	   	        EcxReg = GuestRegs->rcx & 0xffffffff;
	   	        VmxHandleXsetbv(EcxReg, GuestRegs->rdx << 32 | GuestRegs->rax);
	   	        break;
	   	    }
	   	    case VMX_EXIT_REASON_VMX_PREEMPTION_TIMER_EXPIRED:
	   	    {
	   	        VmxHandleVmxPreemptionTimerVmexit(CurrentProcessorIndex, GuestRegs);
	   	        break;
	   	    }
	   	    default:
	   	    {
	   	        LogError("Err, unknown vmexit, reason : 0x%llx", ExitReason);
	   	        break;
	   	    }
	   	    }
	   	    if (!CurrentGuestState->VmxoffState.IsVmxoffExecuted && CurrentGuestState->IncrementRip)
	   	    {
	   	        HvResumeToNextInstruction();
	   	    }
	   	    CurrentGuestState->IsOnVmxRootMode = FALSE;
	   	    if (CurrentGuestState->VmxoffState.IsVmxoffExecuted)
	   	        Result = TRUE;
	   	    if (g_TransparentMode)
	   	    {
	   	        if (ExitReason != VMX_EXIT_REASON_EXECUTE_RDTSC && ExitReason != VMX_EXIT_REASON_EXECUTE_RDTSCP && ExitReason != VMX_EXIT_REASON_EXECUTE_CPUID)
	   	        {
	   	            __writemsr(MSR_IA32_TIME_STAMP_COUNTER, CurrentGuestState->TransparencyState.PreviousTimeStampCounter);
	   	        }
	   	    }
	   	    return Result;
	   	}
	*/
	return true
}

