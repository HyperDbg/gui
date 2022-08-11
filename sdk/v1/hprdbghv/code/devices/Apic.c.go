package devices

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbghv\code\devices\Apic.c.back

type (
	Apic interface {
		XApicIcrWrite() (ok bool)  //col:5
		X2ApicIcrWrite() (ok bool) //col:36
	}
	apic struct{}
)

func NewApic() Apic { return &apic{} }

func (a *apic) XApicIcrWrite() (ok bool) { //col:5
	/*
	   XApicIcrWrite(UINT32 Low, UINT32 High)

	   	{
	   	    *(UINT32 *)((uintptr_t)g_ApicBase + ICROffset + 0x10) = High;
	   	    *(UINT32 *)((uintptr_t)g_ApicBase + ICROffset)        = Low;
	   	}
	*/
	return true
}

func (a *apic) X2ApicIcrWrite() (ok bool) { //col:36
	/*
	   X2ApicIcrWrite(UINT32 Low, UINT32 High)

	   	{
	   	    __writemsr(X2_MSR_BASE + TO_X2(ICROffset), ((UINT64)High << 32) | Low);

	   ApicTriggerGenericNmi()

	   	{
	   	    if (g_IsX2Apic)
	   	    {
	   	        X2ApicIcrWrite((4 << 8) | (1 << 14) | (3 << 18), 0);
	   	        XApicIcrWrite((4 << 8) | (1 << 14) | (3 << 18), 0);

	   ApicInitialize()

	   	{
	   	    UINT64           ApicBaseMSR;
	   	    PHYSICAL_ADDRESS PaApicBase;
	   	    ApicBaseMSR = __readmsr(0x1B);
	   	    if (!(ApicBaseMSR & (1 << 11)))
	   	        return FALSE;
	   	    if (ApicBaseMSR & (1 << 10))
	   	    {
	   	        g_IsX2Apic = TRUE;
	   	        return FALSE;
	   	    }
	   	    else
	   	    {
	   	        PaApicBase.QuadPart = ApicBaseMSR & 0xFFFFFF000;
	   	        g_ApicBase          = MmMapIoSpace(PaApicBase, 0x1000, MmNonCached);
	   	        if (!g_ApicBase)
	   	            return FALSE;
	   	        g_IsX2Apic = FALSE;
	   	    }
	   	    return TRUE;
	   	}
	*/
	return true
}

