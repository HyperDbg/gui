package vmx

//back\HyperDbgDev\hyperdbg\hprdbghv\code\vmm\vmx\CrossVmexits.c.back

type (
	CrossVmexits interface {
		VmxHandleXsetbv() (ok bool)                   //col:25
		VmxHandleVmxPreemptionTimerVmexit() (ok bool) //col:44
	}
)

func NewCrossVmexits() { return &crossVmexits{} }

func (c *crossVmexits) VmxHandleXsetbv() (ok bool) { //col:25
	/*VmxHandleXsetbv(UINT32 Reg, UINT64 Value)
	  {
	      _xsetbv(Reg, Value);
	  }*/
	return true
}

func (c *crossVmexits) VmxHandleVmxPreemptionTimerVmexit() (ok bool) { //col:44
	/*VmxHandleVmxPreemptionTimerVmexit(UINT32 CurrentCoreIndex, PGUEST_REGS GuestRegs)
	  {
	      UNREFERENCED_PARAMETER(GuestRegs);
	      LogError("Why vm-exit for VMX preemption timer happened?");
	      g_GuestState[CurrentCoreIndex].IncrementRip = FALSE;
	  }*/
	return true
}
