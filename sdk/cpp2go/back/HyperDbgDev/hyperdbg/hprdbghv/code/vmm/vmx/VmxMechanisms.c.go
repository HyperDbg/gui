package vmx
//back\HyperDbgDev\hyperdbg\hprdbghv\code\vmm\vmx\VmxMechanisms.c.back

type (
VmxMechanisms interface{
VmxMechanismCreateImmediateVmexitByVmxPreemptionTimer()(ok bool)//col:33
VmxMechanismDisableImmediateVmexitByVmxPreemptionTimer()(ok bool)//col:50
VmxMechanismCreateImmediateVmexitBySelfIpi()(ok bool)//col:65
VmxMechanismCreateImmediateVmexit()(ok bool)//col:97
VmxMechanismHandleImmediateVmexit()(ok bool)//col:118
}
)

func NewVmxMechanisms() { return & vmxMechanisms{} }

func (v *vmxMechanisms)VmxMechanismCreateImmediateVmexitByVmxPreemptionTimer()(ok bool){//col:33
/*VmxMechanismCreateImmediateVmexitByVmxPreemptionTimer()
{
    HvSetVmxPreemptionTimerExiting(TRUE);
    CounterSetPreemptionTimer(0);
}*/
return true
}

func (v *vmxMechanisms)VmxMechanismDisableImmediateVmexitByVmxPreemptionTimer()(ok bool){//col:50
/*VmxMechanismDisableImmediateVmexitByVmxPreemptionTimer()
{
    CounterClearPreemptionTimer();
    HvSetVmxPreemptionTimerExiting(FALSE);
}*/
return true
}

func (v *vmxMechanisms)VmxMechanismCreateImmediateVmexitBySelfIpi()(ok bool){//col:65
/*VmxMechanismCreateImmediateVmexitBySelfIpi()
{
    ApicSelfIpi(IMMEDIATE_VMEXIT_MECHANISM_VECTOR_FOR_SELF_IPI);
}*/
return true
}

func (v *vmxMechanisms)VmxMechanismCreateImmediateVmexit()(ok bool){//col:97
/*VmxMechanismCreateImmediateVmexit(UINT32 CurrentCoreIndex)
{
    g_GuestState[CurrentCoreIndex].WaitForImmediateVmexit = TRUE;
    VmxMechanismCreateImmediateVmexitBySelfIpi();
    HvSetExternalInterruptExiting(TRUE);
}*/
return true
}

func (v *vmxMechanisms)VmxMechanismHandleImmediateVmexit()(ok bool){//col:118
/*VmxMechanismHandleImmediateVmexit(UINT32 CurrentCoreIndex, PGUEST_REGS GuestRegs)
{
    g_GuestState[CurrentCoreIndex].WaitForImmediateVmexit = FALSE;
    HvSetExternalInterruptExiting(FALSE);
}*/
return true
}



