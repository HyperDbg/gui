package vmx
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbghv\code\vmm\vmx\CrossVmexits.c.back

type (
CrossVmexits interface{
VmxHandleXsetbv()(ok bool)//col:4
VmxHandleVmxPreemptionTimerVmexit()(ok bool)//col:10
}
crossVmexits struct{}
)

func NewCrossVmexits()CrossVmexits{ return & crossVmexits{} }

func (c *crossVmexits)VmxHandleXsetbv()(ok bool){//col:4
/*VmxHandleXsetbv(UINT32 Reg, UINT64 Value)
{
    _xsetbv(Reg, Value);
}*/
return true
}

func (c *crossVmexits)VmxHandleVmxPreemptionTimerVmexit()(ok bool){//col:10
/*VmxHandleVmxPreemptionTimerVmexit(UINT32 CurrentCoreIndex, PGUEST_REGS GuestRegs)
{
    UNREFERENCED_PARAMETER(GuestRegs);
    LogError("Why vm-exit for VMX preemption timer happened?");
    g_GuestState[CurrentCoreIndex].IncrementRip = FALSE;
}*/
return true
}



