package vmx
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbghv\code\vmm\vmx\ManageRegs.c.back

type (
ManageRegs interface{
SetGuestCsSel()(ok bool)//col:4
SetGuestCs()(ok bool)//col:11
GetGuestCs()(ok bool)//col:20
SetGuestSsSel()(ok bool)//col:24
SetGuestSs()(ok bool)//col:31
GetGuestSs()(ok bool)//col:40
SetGuestDsSel()(ok bool)//col:44
SetGuestDs()(ok bool)//col:51
GetGuestDs()(ok bool)//col:60
SetGuestFsSel()(ok bool)//col:64
SetGuestFs()(ok bool)//col:71
GetGuestFs()(ok bool)//col:80
SetGuestGsSel()(ok bool)//col:84
SetGuestGs()(ok bool)//col:91
GetGuestGs()(ok bool)//col:100
SetGuestEsSel()(ok bool)//col:104
SetGuestEs()(ok bool)//col:111
GetGuestEs()(ok bool)//col:120
SetGuestIdtr()(ok bool)//col:124
GetGuestIdtr()(ok bool)//col:130
SetGuestLdtr()(ok bool)//col:134
GetGuestLdtr()(ok bool)//col:140
SetGuestGdtr()(ok bool)//col:144
GetGuestGdtr()(ok bool)//col:150
SetGuestTr()(ok bool)//col:154
GetGuestTr()(ok bool)//col:160
SetGuestRFlags()(ok bool)//col:164
GetGuestRFlags()(ok bool)//col:170
SetGuestRIP()(ok bool)//col:174
SetGuestRSP()(ok bool)//col:178
GetGuestRIP()(ok bool)//col:184
GetGuestCr0()(ok bool)//col:190
GetGuestCr2()(ok bool)//col:196
GetGuestCr3()(ok bool)//col:202
GetGuestCr4()(ok bool)//col:208
GetGuestCr8()(ok bool)//col:214
SetGuestCr0()(ok bool)//col:218
SetGuestCr2()(ok bool)//col:222
SetGuestCr3()(ok bool)//col:226
SetGuestCr4()(ok bool)//col:230
SetGuestCr8()(ok bool)//col:234
SetGuestDr0()(ok bool)//col:238
SetGuestDr1()(ok bool)//col:242
SetGuestDr2()(ok bool)//col:246
SetGuestDr3()(ok bool)//col:250
SetGuestDr6()(ok bool)//col:254
SetGuestDr7()(ok bool)//col:258
GetGuestDr0()(ok bool)//col:264
GetGuestDr1()(ok bool)//col:270
GetGuestDr2()(ok bool)//col:276
GetGuestDr3()(ok bool)//col:282
GetGuestDr6()(ok bool)//col:288
GetGuestDr7()(ok bool)//col:294
}
)

func NewManageRegs() { return & manageRegs{} }

func (m *manageRegs)SetGuestCsSel()(ok bool){//col:4
/*SetGuestCsSel(PVMX_SEGMENT_SELECTOR Cs)
{
    __vmx_vmwrite(VMCS_GUEST_CS_SELECTOR, Cs->Selector);
}*/
return true
}

func (m *manageRegs)SetGuestCs()(ok bool){//col:11
/*SetGuestCs(PVMX_SEGMENT_SELECTOR Cs)
{
    __vmx_vmwrite(VMCS_GUEST_CS_BASE, Cs->Base);
    __vmx_vmwrite(VMCS_GUEST_CS_LIMIT, Cs->Limit);
    __vmx_vmwrite(VMCS_GUEST_CS_ACCESS_RIGHTS, Cs->Attributes.AsUInt);
    __vmx_vmwrite(VMCS_GUEST_CS_SELECTOR, Cs->Selector);
}*/
return true
}

func (m *manageRegs)GetGuestCs()(ok bool){//col:20
/*GetGuestCs()
{
    VMX_SEGMENT_SELECTOR Cs;
    __vmx_vmread(VMCS_GUEST_CS_BASE, &Cs.Base);
    __vmx_vmread(VMCS_GUEST_CS_LIMIT, &Cs.Limit);
    __vmx_vmread(VMCS_GUEST_CS_ACCESS_RIGHTS, &Cs.Attributes.AsUInt);
    __vmx_vmread(VMCS_GUEST_CS_SELECTOR, &Cs.Selector);
    return Cs;
}*/
return true
}

func (m *manageRegs)SetGuestSsSel()(ok bool){//col:24
/*SetGuestSsSel(PVMX_SEGMENT_SELECTOR Ss)
{
    __vmx_vmwrite(VMCS_GUEST_SS_SELECTOR, Ss->Selector);
}*/
return true
}

func (m *manageRegs)SetGuestSs()(ok bool){//col:31
/*SetGuestSs(PVMX_SEGMENT_SELECTOR Ss)
{
    __vmx_vmwrite(VMCS_GUEST_SS_BASE, Ss->Base);
    __vmx_vmwrite(VMCS_GUEST_SS_LIMIT, Ss->Limit);
    __vmx_vmwrite(VMCS_GUEST_SS_ACCESS_RIGHTS, Ss->Attributes.AsUInt);
    __vmx_vmwrite(VMCS_GUEST_SS_SELECTOR, Ss->Selector);
}*/
return true
}

func (m *manageRegs)GetGuestSs()(ok bool){//col:40
/*GetGuestSs()
{
    VMX_SEGMENT_SELECTOR Ss;
    __vmx_vmread(VMCS_GUEST_SS_BASE, &Ss.Base);
    __vmx_vmread(VMCS_GUEST_SS_LIMIT, &Ss.Limit);
    __vmx_vmread(VMCS_GUEST_SS_ACCESS_RIGHTS, &Ss.Attributes.AsUInt);
    __vmx_vmread(VMCS_GUEST_SS_SELECTOR, &Ss.Selector);
    return Ss;
}*/
return true
}

func (m *manageRegs)SetGuestDsSel()(ok bool){//col:44
/*SetGuestDsSel(PVMX_SEGMENT_SELECTOR Ds)
{
    __vmx_vmwrite(VMCS_GUEST_DS_SELECTOR, Ds->Selector);
}*/
return true
}

func (m *manageRegs)SetGuestDs()(ok bool){//col:51
/*SetGuestDs(PVMX_SEGMENT_SELECTOR Ds)
{
    __vmx_vmwrite(VMCS_GUEST_DS_BASE, Ds->Base);
    __vmx_vmwrite(VMCS_GUEST_DS_LIMIT, Ds->Limit);
    __vmx_vmwrite(VMCS_GUEST_DS_ACCESS_RIGHTS, Ds->Attributes.AsUInt);
    __vmx_vmwrite(VMCS_GUEST_DS_SELECTOR, Ds->Selector);
}*/
return true
}

func (m *manageRegs)GetGuestDs()(ok bool){//col:60
/*GetGuestDs()
{
    VMX_SEGMENT_SELECTOR Ds;
    __vmx_vmread(VMCS_GUEST_DS_BASE, &Ds.Base);
    __vmx_vmread(VMCS_GUEST_DS_LIMIT, &Ds.Limit);
    __vmx_vmread(VMCS_GUEST_DS_ACCESS_RIGHTS, &Ds.Attributes.AsUInt);
    __vmx_vmread(VMCS_GUEST_DS_SELECTOR, &Ds.Selector);
    return Ds;
}*/
return true
}

func (m *manageRegs)SetGuestFsSel()(ok bool){//col:64
/*SetGuestFsSel(PVMX_SEGMENT_SELECTOR Fs)
{
    __vmx_vmwrite(VMCS_GUEST_FS_SELECTOR, Fs->Selector);
}*/
return true
}

func (m *manageRegs)SetGuestFs()(ok bool){//col:71
/*SetGuestFs(PVMX_SEGMENT_SELECTOR Fs)
{
    __vmx_vmwrite(VMCS_GUEST_FS_BASE, Fs->Base);
    __vmx_vmwrite(VMCS_GUEST_FS_LIMIT, Fs->Limit);
    __vmx_vmwrite(VMCS_GUEST_FS_ACCESS_RIGHTS, Fs->Attributes.AsUInt);
    __vmx_vmwrite(VMCS_GUEST_FS_SELECTOR, Fs->Selector);
}*/
return true
}

func (m *manageRegs)GetGuestFs()(ok bool){//col:80
/*GetGuestFs()
{
    VMX_SEGMENT_SELECTOR Fs;
    __vmx_vmread(VMCS_GUEST_FS_BASE, &Fs.Base);
    __vmx_vmread(VMCS_GUEST_FS_LIMIT, &Fs.Limit);
    __vmx_vmread(VMCS_GUEST_FS_ACCESS_RIGHTS, &Fs.Attributes.AsUInt);
    __vmx_vmread(VMCS_GUEST_FS_SELECTOR, &Fs.Selector);
    return Fs;
}*/
return true
}

func (m *manageRegs)SetGuestGsSel()(ok bool){//col:84
/*SetGuestGsSel(PVMX_SEGMENT_SELECTOR Gs)
{
    __vmx_vmwrite(VMCS_GUEST_GS_SELECTOR, Gs->Selector);
}*/
return true
}

func (m *manageRegs)SetGuestGs()(ok bool){//col:91
/*SetGuestGs(PVMX_SEGMENT_SELECTOR Gs)
{
    __vmx_vmwrite(VMCS_GUEST_GS_BASE, Gs->Base);
    __vmx_vmwrite(VMCS_GUEST_GS_LIMIT, Gs->Limit);
    __vmx_vmwrite(VMCS_GUEST_GS_ACCESS_RIGHTS, Gs->Attributes.AsUInt);
    __vmx_vmwrite(VMCS_GUEST_GS_SELECTOR, Gs->Selector);
}*/
return true
}

func (m *manageRegs)GetGuestGs()(ok bool){//col:100
/*GetGuestGs()
{
    VMX_SEGMENT_SELECTOR Gs;
    __vmx_vmread(VMCS_GUEST_GS_BASE, &Gs.Base);
    __vmx_vmread(VMCS_GUEST_GS_LIMIT, &Gs.Limit);
    __vmx_vmread(VMCS_GUEST_GS_ACCESS_RIGHTS, &Gs.Attributes.AsUInt);
    __vmx_vmread(VMCS_GUEST_GS_SELECTOR, &Gs.Selector);
    return Gs;
}*/
return true
}

func (m *manageRegs)SetGuestEsSel()(ok bool){//col:104
/*SetGuestEsSel(PVMX_SEGMENT_SELECTOR Es)
{
    __vmx_vmwrite(VMCS_GUEST_ES_SELECTOR, Es->Selector);
}*/
return true
}

func (m *manageRegs)SetGuestEs()(ok bool){//col:111
/*SetGuestEs(PVMX_SEGMENT_SELECTOR Es)
{
    __vmx_vmwrite(VMCS_GUEST_ES_BASE, Es->Base);
    __vmx_vmwrite(VMCS_GUEST_ES_LIMIT, Es->Limit);
    __vmx_vmwrite(VMCS_GUEST_ES_ACCESS_RIGHTS, Es->Attributes.AsUInt);
    __vmx_vmwrite(VMCS_GUEST_ES_SELECTOR, Es->Selector);
}*/
return true
}

func (m *manageRegs)GetGuestEs()(ok bool){//col:120
/*GetGuestEs()
{
    VMX_SEGMENT_SELECTOR Es;
    __vmx_vmread(VMCS_GUEST_ES_BASE, &Es.Base);
    __vmx_vmread(VMCS_GUEST_ES_LIMIT, &Es.Limit);
    __vmx_vmread(VMCS_GUEST_ES_ACCESS_RIGHTS, &Es.Attributes.AsUInt);
    __vmx_vmread(VMCS_GUEST_ES_SELECTOR, &Es.Selector);
    return Es;
}*/
return true
}

func (m *manageRegs)SetGuestIdtr()(ok bool){//col:124
/*SetGuestIdtr(UINT64 Idtr)
{
    __vmx_vmwrite(VMCS_GUEST_IDTR_BASE, Idtr);
}*/
return true
}

func (m *manageRegs)GetGuestIdtr()(ok bool){//col:130
/*GetGuestIdtr()
{
    UINT64 Idtr;
    __vmx_vmread(VMCS_GUEST_IDTR_BASE, &Idtr);
    return Idtr;
}*/
return true
}

func (m *manageRegs)SetGuestLdtr()(ok bool){//col:134
/*SetGuestLdtr(UINT64 Ldtr)
{
    __vmx_vmwrite(VMCS_GUEST_LDTR_BASE, Ldtr);
}*/
return true
}

func (m *manageRegs)GetGuestLdtr()(ok bool){//col:140
/*GetGuestLdtr()
{
    UINT64 Ldtr;
    __vmx_vmread(VMCS_GUEST_LDTR_BASE, &Ldtr);
    return Ldtr;
}*/
return true
}

func (m *manageRegs)SetGuestGdtr()(ok bool){//col:144
/*SetGuestGdtr(UINT64 Gdtr)
{
    __vmx_vmwrite(VMCS_GUEST_GDTR_BASE, Gdtr);
}*/
return true
}

func (m *manageRegs)GetGuestGdtr()(ok bool){//col:150
/*GetGuestGdtr()
{
    UINT64 Gdtr;
    __vmx_vmread(VMCS_GUEST_GDTR_BASE, &Gdtr);
    return Gdtr;
}*/
return true
}

func (m *manageRegs)SetGuestTr()(ok bool){//col:154
/*SetGuestTr(UINT64 Tr)
{
    __vmx_vmwrite(VMCS_GUEST_TR_BASE, Tr);
}*/
return true
}

func (m *manageRegs)GetGuestTr()(ok bool){//col:160
/*GetGuestTr()
{
    UINT64 Tr;
    __vmx_vmread(VMCS_GUEST_TR_BASE, &Tr);
    return Tr;
}*/
return true
}

func (m *manageRegs)SetGuestRFlags()(ok bool){//col:164
/*SetGuestRFlags(UINT64 RFlags)
{
    __vmx_vmwrite(VMCS_GUEST_RFLAGS, RFlags);
}*/
return true
}

func (m *manageRegs)GetGuestRFlags()(ok bool){//col:170
/*GetGuestRFlags()
{
    UINT64 RFlags;
    __vmx_vmread(VMCS_GUEST_RFLAGS, &RFlags);
    return RFlags;
}*/
return true
}

func (m *manageRegs)SetGuestRIP()(ok bool){//col:174
/*SetGuestRIP(UINT64 RIP)
{
    __vmx_vmwrite(VMCS_GUEST_RIP, RIP);
}*/
return true
}

func (m *manageRegs)SetGuestRSP()(ok bool){//col:178
/*SetGuestRSP(UINT64 RSP)
{
    __vmx_vmwrite(VMCS_GUEST_RSP, RSP);
}*/
return true
}

func (m *manageRegs)GetGuestRIP()(ok bool){//col:184
/*GetGuestRIP()
{
    UINT64 RIP;
    __vmx_vmread(VMCS_GUEST_RIP, &RIP);
    return RIP;
}*/
return true
}

func (m *manageRegs)GetGuestCr0()(ok bool){//col:190
/*GetGuestCr0()
{
    UINT64 Cr0;
    __vmx_vmread(VMCS_GUEST_CR0, &Cr0);
    return Cr0;
}*/
return true
}

func (m *manageRegs)GetGuestCr2()(ok bool){//col:196
/*GetGuestCr2()
{
    UINT64 Cr2;
    Cr2 = __readcr2();
    return Cr2;
}*/
return true
}

func (m *manageRegs)GetGuestCr3()(ok bool){//col:202
/*GetGuestCr3()
{
    UINT64 Cr3;
    __vmx_vmread(VMCS_GUEST_CR3, &Cr3);
    return Cr3;
}*/
return true
}

func (m *manageRegs)GetGuestCr4()(ok bool){//col:208
/*GetGuestCr4()
{
    UINT64 Cr4;
    __vmx_vmread(VMCS_GUEST_CR4, &Cr4);
    return Cr4;
}*/
return true
}

func (m *manageRegs)GetGuestCr8()(ok bool){//col:214
/*GetGuestCr8()
{
    UINT64 Cr8;
    Cr8 = __readcr8();
    return Cr8;
}*/
return true
}

func (m *manageRegs)SetGuestCr0()(ok bool){//col:218
/*SetGuestCr0(UINT64 Cr0)
{
    __vmx_vmwrite(VMCS_GUEST_CR0, Cr0);
}*/
return true
}

func (m *manageRegs)SetGuestCr2()(ok bool){//col:222
/*SetGuestCr2(UINT64 Cr2)
{
    __writecr2(Cr2);
}*/
return true
}

func (m *manageRegs)SetGuestCr3()(ok bool){//col:226
/*SetGuestCr3(UINT64 Cr3)
{
    __vmx_vmwrite(VMCS_GUEST_CR3, Cr3);
}*/
return true
}

func (m *manageRegs)SetGuestCr4()(ok bool){//col:230
/*SetGuestCr4(UINT64 Cr4)
{
    __vmx_vmwrite(VMCS_GUEST_CR4, Cr4);
}*/
return true
}

func (m *manageRegs)SetGuestCr8()(ok bool){//col:234
/*SetGuestCr8(UINT64 Cr8)
{
    __writecr8(Cr8);
}*/
return true
}

func (m *manageRegs)SetGuestDr0()(ok bool){//col:238
/*SetGuestDr0(UINT64 value)
{
    __writedr(0, value);
}*/
return true
}

func (m *manageRegs)SetGuestDr1()(ok bool){//col:242
/*SetGuestDr1(UINT64 value)
{
    __writedr(1, value);
}*/
return true
}

func (m *manageRegs)SetGuestDr2()(ok bool){//col:246
/*SetGuestDr2(UINT64 value)
{
    __writedr(2, value);
}*/
return true
}

func (m *manageRegs)SetGuestDr3()(ok bool){//col:250
/*SetGuestDr3(UINT64 value)
{
    __writedr(3, value);
}*/
return true
}

func (m *manageRegs)SetGuestDr6()(ok bool){//col:254
/*SetGuestDr6(UINT64 value)
{
    __writedr(6, value);
}*/
return true
}

func (m *manageRegs)SetGuestDr7()(ok bool){//col:258
/*SetGuestDr7(UINT64 value)
{
    __writedr(7, value);
}*/
return true
}

func (m *manageRegs)GetGuestDr0()(ok bool){//col:264
/*GetGuestDr0()
{
    UINT64 Dr0 = 0;
    Dr0        = __readdr(0);
    return Dr0;
}*/
return true
}

func (m *manageRegs)GetGuestDr1()(ok bool){//col:270
/*GetGuestDr1()
{
    UINT64 Dr1 = 0;
    Dr1        = __readdr(1);
    return Dr1;
}*/
return true
}

func (m *manageRegs)GetGuestDr2()(ok bool){//col:276
/*GetGuestDr2()
{
    UINT64 Dr2 = 0;
    Dr2        = __readdr(2);
    return Dr2;
}*/
return true
}

func (m *manageRegs)GetGuestDr3()(ok bool){//col:282
/*GetGuestDr3()
{
    UINT64 Dr3 = 0;
    Dr3        = __readdr(3);
    return Dr3;
}*/
return true
}

func (m *manageRegs)GetGuestDr6()(ok bool){//col:288
/*GetGuestDr6()
{
    UINT64 Dr6 = 0;
    Dr6        = __readdr(6);
    return Dr6;
}*/
return true
}

func (m *manageRegs)GetGuestDr7()(ok bool){//col:294
/*GetGuestDr7()
{
    UINT64 Dr7 = 0;
    Dr7        = __readdr(7);
    return Dr7;
}*/
return true
}



