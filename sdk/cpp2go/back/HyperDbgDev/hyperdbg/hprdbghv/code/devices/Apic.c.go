package devices
//back\HyperDbgDev\hyperdbg\hprdbghv\code\devices\Apic.c.back

type (
Apic interface{
XApicIcrWrite()(ok bool)//col:26
X2ApicIcrWrite()(ok bool)//col:39
ApicTriggerGenericNmi()(ok bool)//col:57
ApicInitialize()(ok bool)//col:90
ApicUninitialize()(ok bool)//col:105
ApicSelfIpi()(ok bool)//col:127
}
)

func NewApic() { return & apic{} }

func (a *apic)XApicIcrWrite()(ok bool){//col:26
/*XApicIcrWrite(UINT32 Low, UINT32 High)
{
    *(UINT32 *)((uintptr_t)g_ApicBase + ICROffset + 0x10) = High;
    *(UINT32 *)((uintptr_t)g_ApicBase + ICROffset)        = Low;
}*/
return true
}

func (a *apic)X2ApicIcrWrite()(ok bool){//col:39
/*X2ApicIcrWrite(UINT32 Low, UINT32 High)
{
    __writemsr(X2_MSR_BASE + TO_X2(ICROffset), ((UINT64)High << 32) | Low);
}*/
return true
}

func (a *apic)ApicTriggerGenericNmi()(ok bool){//col:57
/*ApicTriggerGenericNmi()
{
    if (g_IsX2Apic)
    {
        X2ApicIcrWrite((4 << 8) | (1 << 14) | (3 << 18), 0);
    }
    else
    {
        XApicIcrWrite((4 << 8) | (1 << 14) | (3 << 18), 0);
    }
}*/
return true
}

func (a *apic)ApicInitialize()(ok bool){//col:90
/*ApicInitialize()
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
}*/
return true
}

func (a *apic)ApicUninitialize()(ok bool){//col:105
/*ApicUninitialize()
{
    if (g_ApicBase)
        MmUnmapIoSpace(g_ApicBase, 0x1000);
}*/
return true
}

func (a *apic)ApicSelfIpi()(ok bool){//col:127
/*ApicSelfIpi(UINT32 Vector)
{
    if (g_IsX2Apic)
    {
        X2ApicIcrWrite(APIC_DEST_SELF | APIC_DEST_PHYSICAL | APIC_DM_FIXED | Vector, 0);
    }
    else
    {
        XApicIcrWrite(APIC_DEST_SELF | APIC_DEST_PHYSICAL | APIC_DM_FIXED | Vector, 0);
    }
}*/
return true
}


