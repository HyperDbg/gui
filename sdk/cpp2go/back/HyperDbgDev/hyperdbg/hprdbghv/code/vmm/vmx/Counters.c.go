package vmx
//back\HyperDbgDev\hyperdbg\hprdbghv\code\vmm\vmx\Counters.c.back

type (
Counters interface{
CounterEmulateRdtsc()(ok bool)//col:6
CounterEmulateRdtscp()(ok bool)//col:14
CounterEmulateRdpmc()(ok bool)//col:22
CounterSetPreemptionTimer()(ok bool)//col:26
CounterClearPreemptionTimer()(ok bool)//col:30
}
)

func NewCounters() { return & counters{} }

func (c *counters)CounterEmulateRdtsc()(ok bool){//col:6
/*CounterEmulateRdtsc(PGUEST_REGS GuestRegs)
{
    UINT64 Tsc     = __rdtsc();
    GuestRegs->rax = 0x00000000ffffffff & Tsc;
    GuestRegs->rdx = 0x00000000ffffffff & (Tsc >> 32);
}*/
return true
}

func (c *counters)CounterEmulateRdtscp()(ok bool){//col:14
/*CounterEmulateRdtscp(PGUEST_REGS GuestRegs)
{
    int    Aux     = 0;
    UINT64 Tsc     = __rdtscp(&Aux);
    GuestRegs->rax = 0x00000000ffffffff & Tsc;
    GuestRegs->rdx = 0x00000000ffffffff & (Tsc >> 32);
    GuestRegs->rcx = 0x00000000ffffffff & Aux;
}*/
return true
}

func (c *counters)CounterEmulateRdpmc()(ok bool){//col:22
/*CounterEmulateRdpmc(PGUEST_REGS GuestRegs)
{
    ULONG EcxReg = 0;
    EcxReg         = GuestRegs->rcx & 0xffffffff;
    UINT64 Pmc     = __readpmc(EcxReg);
    GuestRegs->rax = 0x00000000ffffffff & Pmc;
    GuestRegs->rdx = 0x00000000ffffffff & (Pmc >> 32);
}*/
return true
}

func (c *counters)CounterSetPreemptionTimer()(ok bool){//col:26
/*CounterSetPreemptionTimer(UINT32 TimerValue)
{
    __vmx_vmwrite(VMCS_GUEST_VMX_PREEMPTION_TIMER_VALUE, TimerValue);
}*/
return true
}

func (c *counters)CounterClearPreemptionTimer()(ok bool){//col:30
/*CounterClearPreemptionTimer()
{
    __vmx_vmwrite(VMCS_GUEST_VMX_PREEMPTION_TIMER_VALUE, NULL);
}*/
return true
}



