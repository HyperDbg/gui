#include "pch.h"
VOID
SetGuestCsSel(PVMX_SEGMENT_SELECTOR Cs)
{
    __vmx_vmwrite(VMCS_GUEST_CS_SELECTOR, Cs->Selector);
}

VOID
SetGuestCs(PVMX_SEGMENT_SELECTOR Cs)
{
    __vmx_vmwrite(VMCS_GUEST_CS_BASE, Cs->Base);
    __vmx_vmwrite(VMCS_GUEST_CS_LIMIT, Cs->Limit);
    __vmx_vmwrite(VMCS_GUEST_CS_ACCESS_RIGHTS, Cs->Attributes.AsUInt);
    __vmx_vmwrite(VMCS_GUEST_CS_SELECTOR, Cs->Selector);
}

VMX_SEGMENT_SELECTOR
GetGuestCs()
{
    VMX_SEGMENT_SELECTOR Cs;
    __vmx_vmread(VMCS_GUEST_CS_BASE, &Cs.Base);
    __vmx_vmread(VMCS_GUEST_CS_LIMIT, &Cs.Limit);
    __vmx_vmread(VMCS_GUEST_CS_ACCESS_RIGHTS, &Cs.Attributes.AsUInt);
    __vmx_vmread(VMCS_GUEST_CS_SELECTOR, &Cs.Selector);
    return Cs;
}

VOID
SetGuestSsSel(PVMX_SEGMENT_SELECTOR Ss)
{
    __vmx_vmwrite(VMCS_GUEST_SS_SELECTOR, Ss->Selector);
}

VOID
SetGuestSs(PVMX_SEGMENT_SELECTOR Ss)
{
    __vmx_vmwrite(VMCS_GUEST_SS_BASE, Ss->Base);
    __vmx_vmwrite(VMCS_GUEST_SS_LIMIT, Ss->Limit);
    __vmx_vmwrite(VMCS_GUEST_SS_ACCESS_RIGHTS, Ss->Attributes.AsUInt);
    __vmx_vmwrite(VMCS_GUEST_SS_SELECTOR, Ss->Selector);
}

VMX_SEGMENT_SELECTOR
GetGuestSs()
{
    VMX_SEGMENT_SELECTOR Ss;
    __vmx_vmread(VMCS_GUEST_SS_BASE, &Ss.Base);
    __vmx_vmread(VMCS_GUEST_SS_LIMIT, &Ss.Limit);
    __vmx_vmread(VMCS_GUEST_SS_ACCESS_RIGHTS, &Ss.Attributes.AsUInt);
    __vmx_vmread(VMCS_GUEST_SS_SELECTOR, &Ss.Selector);
    return Ss;
}

VOID
SetGuestDsSel(PVMX_SEGMENT_SELECTOR Ds)
{
    __vmx_vmwrite(VMCS_GUEST_DS_SELECTOR, Ds->Selector);
}

VOID
SetGuestDs(PVMX_SEGMENT_SELECTOR Ds)
{
    __vmx_vmwrite(VMCS_GUEST_DS_BASE, Ds->Base);
    __vmx_vmwrite(VMCS_GUEST_DS_LIMIT, Ds->Limit);
    __vmx_vmwrite(VMCS_GUEST_DS_ACCESS_RIGHTS, Ds->Attributes.AsUInt);
    __vmx_vmwrite(VMCS_GUEST_DS_SELECTOR, Ds->Selector);
}

VMX_SEGMENT_SELECTOR
GetGuestDs()
{
    VMX_SEGMENT_SELECTOR Ds;
    __vmx_vmread(VMCS_GUEST_DS_BASE, &Ds.Base);
    __vmx_vmread(VMCS_GUEST_DS_LIMIT, &Ds.Limit);
    __vmx_vmread(VMCS_GUEST_DS_ACCESS_RIGHTS, &Ds.Attributes.AsUInt);
    __vmx_vmread(VMCS_GUEST_DS_SELECTOR, &Ds.Selector);
    return Ds;
}

VOID
SetGuestFsSel(PVMX_SEGMENT_SELECTOR Fs)
{
    __vmx_vmwrite(VMCS_GUEST_FS_SELECTOR, Fs->Selector);
}

VOID
SetGuestFs(PVMX_SEGMENT_SELECTOR Fs)
{
    __vmx_vmwrite(VMCS_GUEST_FS_BASE, Fs->Base);
    __vmx_vmwrite(VMCS_GUEST_FS_LIMIT, Fs->Limit);
    __vmx_vmwrite(VMCS_GUEST_FS_ACCESS_RIGHTS, Fs->Attributes.AsUInt);
    __vmx_vmwrite(VMCS_GUEST_FS_SELECTOR, Fs->Selector);
}

VMX_SEGMENT_SELECTOR
GetGuestFs()
{
    VMX_SEGMENT_SELECTOR Fs;
    __vmx_vmread(VMCS_GUEST_FS_BASE, &Fs.Base);
    __vmx_vmread(VMCS_GUEST_FS_LIMIT, &Fs.Limit);
    __vmx_vmread(VMCS_GUEST_FS_ACCESS_RIGHTS, &Fs.Attributes.AsUInt);
    __vmx_vmread(VMCS_GUEST_FS_SELECTOR, &Fs.Selector);
    return Fs;
}

VOID
SetGuestGsSel(PVMX_SEGMENT_SELECTOR Gs)
{
    __vmx_vmwrite(VMCS_GUEST_GS_SELECTOR, Gs->Selector);
}

VOID
SetGuestGs(PVMX_SEGMENT_SELECTOR Gs)
{
    __vmx_vmwrite(VMCS_GUEST_GS_BASE, Gs->Base);
    __vmx_vmwrite(VMCS_GUEST_GS_LIMIT, Gs->Limit);
    __vmx_vmwrite(VMCS_GUEST_GS_ACCESS_RIGHTS, Gs->Attributes.AsUInt);
    __vmx_vmwrite(VMCS_GUEST_GS_SELECTOR, Gs->Selector);
}

VMX_SEGMENT_SELECTOR
GetGuestGs()
{
    VMX_SEGMENT_SELECTOR Gs;
    __vmx_vmread(VMCS_GUEST_GS_BASE, &Gs.Base);
    __vmx_vmread(VMCS_GUEST_GS_LIMIT, &Gs.Limit);
    __vmx_vmread(VMCS_GUEST_GS_ACCESS_RIGHTS, &Gs.Attributes.AsUInt);
    __vmx_vmread(VMCS_GUEST_GS_SELECTOR, &Gs.Selector);
    return Gs;
}

VOID
SetGuestEsSel(PVMX_SEGMENT_SELECTOR Es)
{
    __vmx_vmwrite(VMCS_GUEST_ES_SELECTOR, Es->Selector);
}

VOID
SetGuestEs(PVMX_SEGMENT_SELECTOR Es)
{
    __vmx_vmwrite(VMCS_GUEST_ES_BASE, Es->Base);
    __vmx_vmwrite(VMCS_GUEST_ES_LIMIT, Es->Limit);
    __vmx_vmwrite(VMCS_GUEST_ES_ACCESS_RIGHTS, Es->Attributes.AsUInt);
    __vmx_vmwrite(VMCS_GUEST_ES_SELECTOR, Es->Selector);
}

VMX_SEGMENT_SELECTOR
GetGuestEs()
{
    VMX_SEGMENT_SELECTOR Es;
    __vmx_vmread(VMCS_GUEST_ES_BASE, &Es.Base);
    __vmx_vmread(VMCS_GUEST_ES_LIMIT, &Es.Limit);
    __vmx_vmread(VMCS_GUEST_ES_ACCESS_RIGHTS, &Es.Attributes.AsUInt);
    __vmx_vmread(VMCS_GUEST_ES_SELECTOR, &Es.Selector);
    return Es;
}

VOID
SetGuestIdtr(UINT64 Idtr)
{
    __vmx_vmwrite(VMCS_GUEST_IDTR_BASE, Idtr);
}

UINT64
GetGuestIdtr()
{
    UINT64 Idtr;
    __vmx_vmread(VMCS_GUEST_IDTR_BASE, &Idtr);
    return Idtr;
}

VOID
SetGuestLdtr(UINT64 Ldtr)
{
    __vmx_vmwrite(VMCS_GUEST_LDTR_BASE, Ldtr);
}

UINT64
GetGuestLdtr()
{
    UINT64 Ldtr;
    __vmx_vmread(VMCS_GUEST_LDTR_BASE, &Ldtr);
    return Ldtr;
}

VOID
SetGuestGdtr(UINT64 Gdtr)
{
    __vmx_vmwrite(VMCS_GUEST_GDTR_BASE, Gdtr);
}

UINT64
GetGuestGdtr()
{
    UINT64 Gdtr;
    __vmx_vmread(VMCS_GUEST_GDTR_BASE, &Gdtr);
    return Gdtr;
}

VOID
SetGuestTr(UINT64 Tr)
{
    __vmx_vmwrite(VMCS_GUEST_TR_BASE, Tr);
}

UINT64
GetGuestTr()
{
    UINT64 Tr;
    __vmx_vmread(VMCS_GUEST_TR_BASE, &Tr);
    return Tr;
}

VOID
SetGuestRFlags(UINT64 RFlags)
{
    __vmx_vmwrite(VMCS_GUEST_RFLAGS, RFlags);
}

UINT64
GetGuestRFlags()
{
    UINT64 RFlags;
    __vmx_vmread(VMCS_GUEST_RFLAGS, &RFlags);
    return RFlags;
}

VOID
SetGuestRIP(UINT64 RIP)
{
    __vmx_vmwrite(VMCS_GUEST_RIP, RIP);
}

VOID
SetGuestRSP(UINT64 RSP)
{
    __vmx_vmwrite(VMCS_GUEST_RSP, RSP);
}

UINT64
GetGuestRIP()
{
    UINT64 RIP;
    __vmx_vmread(VMCS_GUEST_RIP, &RIP);
    return RIP;
}

UINT64
GetGuestCr0()
{
    UINT64 Cr0;
    __vmx_vmread(VMCS_GUEST_CR0, &Cr0);
    return Cr0;
}

UINT64
GetGuestCr2()
{
    UINT64 Cr2;
    Cr2 = __readcr2();
    return Cr2;
}

UINT64
GetGuestCr3()
{
    UINT64 Cr3;
    __vmx_vmread(VMCS_GUEST_CR3, &Cr3);
    return Cr3;
}

UINT64
GetGuestCr4()
{
    UINT64 Cr4;
    __vmx_vmread(VMCS_GUEST_CR4, &Cr4);
    return Cr4;
}

UINT64
GetGuestCr8()
{
    UINT64 Cr8;
    Cr8 = __readcr8();
    return Cr8;
}

VOID
SetGuestCr0(UINT64 Cr0)
{
    __vmx_vmwrite(VMCS_GUEST_CR0, Cr0);
}

VOID
SetGuestCr2(UINT64 Cr2)
{
    __writecr2(Cr2);
}

VOID
SetGuestCr3(UINT64 Cr3)
{
    __vmx_vmwrite(VMCS_GUEST_CR3, Cr3);
}

VOID
SetGuestCr4(UINT64 Cr4)
{
    __vmx_vmwrite(VMCS_GUEST_CR4, Cr4);
}

VOID
SetGuestCr8(UINT64 Cr8)
{
    __writecr8(Cr8);
}

VOID
SetGuestDr0(UINT64 value)
{
    __writedr(0, value);
}

VOID
SetGuestDr1(UINT64 value)
{
    __writedr(1, value);
}

VOID
SetGuestDr2(UINT64 value)
{
    __writedr(2, value);
}

VOID
SetGuestDr3(UINT64 value)
{
    __writedr(3, value);
}

VOID
SetGuestDr6(UINT64 value)
{
    __writedr(6, value);
}

VOID
SetGuestDr7(UINT64 value)
{
    __writedr(7, value);
}

UINT64
GetGuestDr0()
{
    UINT64 Dr0 = 0;
    Dr0        = __readdr(0);
    return Dr0;
}

UINT64
GetGuestDr1()
{
    UINT64 Dr1 = 0;
    Dr1        = __readdr(1);
    return Dr1;
}

UINT64
GetGuestDr2()
{
    UINT64 Dr2 = 0;
    Dr2        = __readdr(2);
    return Dr2;
}

UINT64
GetGuestDr3()
{
    UINT64 Dr3 = 0;
    Dr3        = __readdr(3);
    return Dr3;
}

UINT64
GetGuestDr6()
{
    UINT64 Dr6 = 0;
    Dr6        = __readdr(6);
    return Dr6;
}

UINT64
GetGuestDr7()
{
    UINT64 Dr7 = 0;
    Dr7        = __readdr(7);
    return Dr7;
}

