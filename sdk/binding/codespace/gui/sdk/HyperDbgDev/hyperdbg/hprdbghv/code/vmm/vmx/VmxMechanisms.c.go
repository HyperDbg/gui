package vmx

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbghv\code\vmm\vmx\VmxMechanisms.c.back

type (
	VmxMechanisms interface {
		VmxMechanismCreateImmediateVmexitByVmxPreemptionTimer() (ok bool)  //col:5
		VmxMechanismDisableImmediateVmexitByVmxPreemptionTimer() (ok bool) //col:10
		VmxMechanismCreateImmediateVmexitBySelfIpi() (ok bool)             //col:14
		VmxMechanismCreateImmediateVmexit() (ok bool)                      //col:20
		VmxMechanismHandleImmediateVmexit() (ok bool)                      //col:25
	}
	vmxMechanisms struct{}
)

func NewVmxMechanisms() VmxMechanisms { return &vmxMechanisms{} }

func (v *vmxMechanisms) VmxMechanismCreateImmediateVmexitByVmxPreemptionTimer() (ok bool) { //col:5
	/*
	   VmxMechanismCreateImmediateVmexitByVmxPreemptionTimer()

	   	{
	   	    HvSetVmxPreemptionTimerExiting(TRUE);
	   	    CounterSetPreemptionTimer(0);
	   	}
	*/
	return true
}

func (v *vmxMechanisms) VmxMechanismDisableImmediateVmexitByVmxPreemptionTimer() (ok bool) { //col:10
	/*
	   VmxMechanismDisableImmediateVmexitByVmxPreemptionTimer()

	   	{
	   	    CounterClearPreemptionTimer();
	   	    HvSetVmxPreemptionTimerExiting(FALSE);
	   	}
	*/
	return true
}

func (v *vmxMechanisms) VmxMechanismCreateImmediateVmexitBySelfIpi() (ok bool) { //col:14
	/*
	   VmxMechanismCreateImmediateVmexitBySelfIpi()

	   	{
	   	    ApicSelfIpi(IMMEDIATE_VMEXIT_MECHANISM_VECTOR_FOR_SELF_IPI);
	   	}
	*/
	return true
}

func (v *vmxMechanisms) VmxMechanismCreateImmediateVmexit() (ok bool) { //col:20
	/*
	   VmxMechanismCreateImmediateVmexit(UINT32 CurrentCoreIndex)

	   	{
	   	    g_GuestState[CurrentCoreIndex].WaitForImmediateVmexit = TRUE;
	   	    VmxMechanismCreateImmediateVmexitBySelfIpi();
	   	    HvSetExternalInterruptExiting(TRUE);
	   	}
	*/
	return true
}

func (v *vmxMechanisms) VmxMechanismHandleImmediateVmexit() (ok bool) { //col:25
	/*
	   VmxMechanismHandleImmediateVmexit(UINT32 CurrentCoreIndex, PGUEST_REGS GuestRegs)

	   	{
	   	    g_GuestState[CurrentCoreIndex].WaitForImmediateVmexit = FALSE;
	   	    HvSetExternalInterruptExiting(FALSE);
	   	}
	*/
	return true
}

