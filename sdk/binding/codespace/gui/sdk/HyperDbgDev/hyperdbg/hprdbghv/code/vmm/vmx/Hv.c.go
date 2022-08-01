package vmx
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbghv\code\vmm\vmx\Hv.c.back

type (
Hv interface{
HvAdjustControls()(ok bool)//col:8
HvSetGuestSelector()(ok bool)//col:22
HvHandleCpuid()(ok bool)//col:49
HvHandleControlRegisterAccess()(ok bool)//col:123
HvFillGuestSelectorData()(ok bool)//col:138
HvResumeToNextInstruction()(ok bool)//col:148
HvSetMonitorTrapFlag()(ok bool)//col:162
HvSetLoadDebugControls()(ok bool)//col:176
HvSetSaveDebugControls()(ok bool)//col:190
HvRestoreRegisters()(ok bool)//col:209
HvSetPmcVmexit()(ok bool)//col:223
HvSetMovControlRegsExiting()(ok bool)//col:227
HvSetMovToCr3Vmexit()(ok bool)//col:231
HvWriteExceptionBitmap()(ok bool)//col:235
HvReadExceptionBitmap()(ok bool)//col:241
HvSetInterruptWindowExiting()(ok bool)//col:255
HvSetNmiWindowExiting()(ok bool)//col:269
HvHandleMovDebugRegister()(ok bool)//col:377
HvSetNmiExiting()(ok bool)//col:396
HvSetVmxPreemptionTimerExiting()(ok bool)//col:410
HvSetExceptionBitmap()(ok bool)//col:414
HvUnsetExceptionBitmap()(ok bool)//col:418
HvSetExternalInterruptExiting()(ok bool)//col:422
HvSetRdtscExiting()(ok bool)//col:426
HvSetMovDebugRegsExiting()(ok bool)//col:430
}
hv struct{}
)

func NewHv()Hv{ return & hv{} }

func (h *hv)HvAdjustControls()(ok bool){//col:8
/*HvAdjustControls(ULONG Ctl, ULONG Msr)
{
    MSR MsrValue = {0};
    MsrValue.Flags = __readmsr(Msr);
    Ctl &= MsrValue.Fields.High;
    Ctl |= MsrValue.Fields.Low; 
    return Ctl;
}*/
return true
}

func (h *hv)HvSetGuestSelector()(ok bool){//col:22
/*HvSetGuestSelector(PVOID GdtBase, ULONG SegmentRegister, UINT16 Selector)
{
    VMX_SEGMENT_SELECTOR SegmentSelector = {0};
    GetSegmentDescriptor(GdtBase, Selector, &SegmentSelector);
    if (Selector == 0x0)
    {
        SegmentSelector.Attributes.Unusable = TRUE;
    }
    __vmx_vmwrite(VMCS_GUEST_ES_SELECTOR + SegmentRegister * 2, Selector);
    __vmx_vmwrite(VMCS_GUEST_ES_LIMIT + SegmentRegister * 2, SegmentSelector.Limit);
    __vmx_vmwrite(VMCS_GUEST_ES_ACCESS_RIGHTS + SegmentRegister * 2, SegmentSelector.Attributes.AsUInt);
    __vmx_vmwrite(VMCS_GUEST_ES_BASE + SegmentRegister * 2, SegmentSelector.Base);
    return TRUE;
}*/
return true
}

func (h *hv)HvHandleCpuid()(ok bool){//col:49
/*HvHandleCpuid(PGUEST_REGS RegistersState)
{
    INT32  CpuInfo[4];
    ULONG  Mode    = 0;
    UINT64 Context = 0;
    if (g_UserDebuggerState && UdCheckForCommand())
    {
        return;
    }
    Context = RegistersState->rax;
    __cpuidex(CpuInfo, (INT32)RegistersState->rax, (INT32)RegistersState->rcx);
    if (!g_TransparentMode)
    {
        if (RegistersState->rax == CPUID_PROCESSOR_AND_PROCESSOR_FEATURE_IDENTIFIERS)
        {
            CpuInfo[2] |= HYPERV_HYPERVISOR_PRESENT_BIT;
        }
    }
    RegistersState->rax = CpuInfo[0];
    RegistersState->rbx = CpuInfo[1];
    RegistersState->rcx = CpuInfo[2];
    RegistersState->rdx = CpuInfo[3];
    if (g_TriggerEventForCpuids)
    {
        DebuggerTriggerEvents(CPUID_INSTRUCTION_EXECUTION, RegistersState, Context);
    }
}*/
return true
}

func (h *hv)HvHandleControlRegisterAccess()(ok bool){//col:123
/*HvHandleControlRegisterAccess(PGUEST_REGS GuestState, UINT32 ProcessorIndex)
{
    ULONG                           ExitQualification = 0;
    VMX_EXIT_QUALIFICATION_MOV_CR * CrExitQualification;
    UINT64 *                        RegPtr;
    UINT64                          NewCr3;
    CR3_TYPE                        NewCr3Reg;
    VIRTUAL_MACHINE_STATE *         CurrentVmState = &g_GuestState[ProcessorIndex];
    __vmx_vmread(VMCS_EXIT_QUALIFICATION, &ExitQualification);
    CrExitQualification = (VMX_EXIT_QUALIFICATION_MOV_CR *)&ExitQualification;
    RegPtr = (UINT64 *)&GuestState->rax + CrExitQualification->GeneralPurposeRegister;
    switch (CrExitQualification->AccessType)
    {
    case VMX_EXIT_QUALIFICATION_ACCESS_MOV_TO_CR:
    {
        switch (CrExitQualification->ControlRegister)
        {
        case VMX_EXIT_QUALIFICATION_REGISTER_CR0:
            __vmx_vmwrite(VMCS_GUEST_CR0, *RegPtr);
            __vmx_vmwrite(VMCS_CTRL_CR0_READ_SHADOW, *RegPtr);
            DebuggerTriggerEvents(CONTROL_REGISTER_MODIFIED, GuestState, VMX_EXIT_QUALIFICATION_REGISTER_CR0);
            break;
        case VMX_EXIT_QUALIFICATION_REGISTER_CR3:
            NewCr3          = (*RegPtr & ~(1ULL << 63));
            NewCr3Reg.Flags = NewCr3;
            __vmx_vmwrite(VMCS_GUEST_CR3, NewCr3Reg.Flags);
            VpidInvvpidSingleContext(VPID_TAG);
            if (CurrentVmState->DebuggingState.ThreadOrProcessTracingDetails.IsWatingForMovCr3VmExits)
            {
                ProcessHandleProcessChange(ProcessorIndex, GuestState);
            }
            if (g_CheckPageFaultsAndMov2Cr3VmexitsWithUserDebugger)
            {
                AttachingHandleCr3VmexitsForThreadInterception(ProcessorIndex, NewCr3Reg);
            }
            break;
        case VMX_EXIT_QUALIFICATION_REGISTER_CR4:
            __vmx_vmwrite(VMCS_GUEST_CR4, *RegPtr);
            __vmx_vmwrite(VMCS_CTRL_CR4_READ_SHADOW, *RegPtr);
            DebuggerTriggerEvents(CONTROL_REGISTER_MODIFIED, GuestState, VMX_EXIT_QUALIFICATION_REGISTER_CR4);
            break;
        default:
            LogWarning("Unsupported register 0x%x in handling control registers access",
                       CrExitQualification->ControlRegister);
            break;
        }
    }
    break;
    case VMX_EXIT_QUALIFICATION_ACCESS_MOV_FROM_CR:
    {
        switch (CrExitQualification->ControlRegister)
        {
        case VMX_EXIT_QUALIFICATION_REGISTER_CR0:
            __vmx_vmread(VMCS_GUEST_CR0, RegPtr);
            break;
        case VMX_EXIT_QUALIFICATION_REGISTER_CR3:
            __vmx_vmread(VMCS_GUEST_CR3, RegPtr);
            break;
        case VMX_EXIT_QUALIFICATION_REGISTER_CR4:
            __vmx_vmread(VMCS_GUEST_CR4, RegPtr);
            break;
        default:
            LogWarning("Unsupported register 0x%x in handling control registers access",
                       CrExitQualification->ControlRegister);
            break;
        }
    }
    break;
    default:
        LogWarning("Unsupported operation 0x%x in handling control registers access",
                   CrExitQualification->AccessType);
        break;
    }
}*/
return true
}

func (h *hv)HvFillGuestSelectorData()(ok bool){//col:138
/*HvFillGuestSelectorData(PVOID GdtBase, ULONG SegmentRegister, UINT16 Selector)
{
    VMX_SEGMENT_SELECTOR SegmentSelector = {0};
    GetSegmentDescriptor(GdtBase, Selector, &SegmentSelector);
    if (Selector == 0x0)
    {
        SegmentSelector.Attributes.Unusable = TRUE;
    }
    SegmentSelector.Attributes.Reserved1 = 0;
    SegmentSelector.Attributes.Reserved2 = 0;
    __vmx_vmwrite(VMCS_GUEST_ES_SELECTOR + SegmentRegister * 2, Selector);
    __vmx_vmwrite(VMCS_GUEST_ES_LIMIT + SegmentRegister * 2, SegmentSelector.Limit);
    __vmx_vmwrite(VMCS_GUEST_ES_ACCESS_RIGHTS + SegmentRegister * 2, SegmentSelector.Attributes.AsUInt);
    __vmx_vmwrite(VMCS_GUEST_ES_BASE + SegmentRegister * 2, SegmentSelector.Base);
}*/
return true
}

func (h *hv)HvResumeToNextInstruction()(ok bool){//col:148
/*HvResumeToNextInstruction()
{
    UINT64 ResumeRIP             = NULL;
    UINT64 CurrentRIP            = NULL;
    size_t ExitInstructionLength = 0;
    __vmx_vmread(VMCS_GUEST_RIP, &CurrentRIP);
    __vmx_vmread(VMCS_VMEXIT_INSTRUCTION_LENGTH, &ExitInstructionLength);
    ResumeRIP = CurrentRIP + ExitInstructionLength;
    __vmx_vmwrite(VMCS_GUEST_RIP, ResumeRIP);
}*/
return true
}

func (h *hv)HvSetMonitorTrapFlag()(ok bool){//col:162
/*HvSetMonitorTrapFlag(BOOLEAN Set)
{
    ULONG CpuBasedVmExecControls = 0;
    __vmx_vmread(VMCS_CTRL_PROCESSOR_BASED_VM_EXECUTION_CONTROLS, &CpuBasedVmExecControls);
    if (Set)
    {
        CpuBasedVmExecControls |= CPU_BASED_MONITOR_TRAP_FLAG;
    }
    else
    {
        CpuBasedVmExecControls &= ~CPU_BASED_MONITOR_TRAP_FLAG;
    }
    __vmx_vmwrite(VMCS_CTRL_PROCESSOR_BASED_VM_EXECUTION_CONTROLS, CpuBasedVmExecControls);
}*/
return true
}

func (h *hv)HvSetLoadDebugControls()(ok bool){//col:176
/*HvSetLoadDebugControls(BOOLEAN Set)
{
    ULONG VmentryControls = 0;
    __vmx_vmread(VMCS_CTRL_VMENTRY_CONTROLS, &VmentryControls);
    if (Set)
    {
        VmentryControls |= VM_ENTRY_LOAD_DEBUG_CONTROLS;
    }
    else
    {
        VmentryControls &= ~VM_ENTRY_LOAD_DEBUG_CONTROLS;
    }
    __vmx_vmwrite(VMCS_CTRL_VMENTRY_CONTROLS, VmentryControls);
}*/
return true
}

func (h *hv)HvSetSaveDebugControls()(ok bool){//col:190
/*HvSetSaveDebugControls(BOOLEAN Set)
{
    ULONG VmexitControls = 0;
    __vmx_vmread(VMCS_CTRL_PRIMARY_VMEXIT_CONTROLS, &VmexitControls);
    if (Set)
    {
        VmexitControls |= VM_EXIT_SAVE_DEBUG_CONTROLS;
    }
    else
    {
        VmexitControls &= ~VM_EXIT_SAVE_DEBUG_CONTROLS;
    }
    __vmx_vmwrite(VMCS_CTRL_PRIMARY_VMEXIT_CONTROLS, VmexitControls);
}*/
return true
}

func (h *hv)HvRestoreRegisters()(ok bool){//col:209
/*HvRestoreRegisters()
{
    UINT64 FsBase;
    UINT64 GsBase;
    UINT64 GdtrBase;
    UINT64 GdtrLimit;
    UINT64 IdtrBase;
    UINT64 IdtrLimit;
    __vmx_vmread(VMCS_GUEST_FS_BASE, &FsBase);
    __writemsr(IA32_FS_BASE, FsBase);
    __vmx_vmread(VMCS_GUEST_GS_BASE, &GsBase);
    __writemsr(IA32_GS_BASE, GsBase);
    __vmx_vmread(VMCS_GUEST_GDTR_BASE, &GdtrBase);
    __vmx_vmread(VMCS_GUEST_GDTR_LIMIT, &GdtrLimit);
    AsmReloadGdtr(GdtrBase, GdtrLimit);
    __vmx_vmread(VMCS_GUEST_IDTR_BASE, &IdtrBase);
    __vmx_vmread(VMCS_GUEST_IDTR_LIMIT, &IdtrLimit);
    AsmReloadIdtr(IdtrBase, IdtrLimit);
}*/
return true
}

func (h *hv)HvSetPmcVmexit()(ok bool){//col:223
/*HvSetPmcVmexit(BOOLEAN Set)
{
    ULONG CpuBasedVmExecControls = 0;
    __vmx_vmread(VMCS_CTRL_PROCESSOR_BASED_VM_EXECUTION_CONTROLS, &CpuBasedVmExecControls);
    if (Set)
    {
        CpuBasedVmExecControls |= CPU_BASED_RDPMC_EXITING;
    }
    else
    {
        CpuBasedVmExecControls &= ~CPU_BASED_RDPMC_EXITING;
    }
    __vmx_vmwrite(VMCS_CTRL_PROCESSOR_BASED_VM_EXECUTION_CONTROLS, CpuBasedVmExecControls);
}*/
return true
}

func (h *hv)HvSetMovControlRegsExiting()(ok bool){//col:227
/*HvSetMovControlRegsExiting(BOOLEAN Set, UINT64 ControlRegister, UINT64 MaskRegister)
{
    ProtectedHvSetMov2CrExiting(Set, ControlRegister, MaskRegister);
}*/
return true
}

func (h *hv)HvSetMovToCr3Vmexit()(ok bool){//col:231
/*HvSetMovToCr3Vmexit(BOOLEAN Set)
{
    ProtectedHvSetMov2Cr3Exiting(Set);
}*/
return true
}

func (h *hv)HvWriteExceptionBitmap()(ok bool){//col:235
/*HvWriteExceptionBitmap(UINT32 BitmapMask)
{
    __vmx_vmwrite(VMCS_CTRL_EXCEPTION_BITMAP, BitmapMask);
}*/
return true
}

func (h *hv)HvReadExceptionBitmap()(ok bool){//col:241
/*HvReadExceptionBitmap()
{
    UINT32 ExceptionBitmap = 0;
    __vmx_vmread(VMCS_CTRL_EXCEPTION_BITMAP, &ExceptionBitmap);
    return ExceptionBitmap;
}*/
return true
}

func (h *hv)HvSetInterruptWindowExiting()(ok bool){//col:255
/*HvSetInterruptWindowExiting(BOOLEAN Set)
{
    ULONG CpuBasedVmExecControls = 0;
    __vmx_vmread(VMCS_CTRL_PROCESSOR_BASED_VM_EXECUTION_CONTROLS, &CpuBasedVmExecControls);
    if (Set)
    {
        CpuBasedVmExecControls |= CPU_BASED_VIRTUAL_INTR_PENDING;
    }
    else
    {
        CpuBasedVmExecControls &= ~CPU_BASED_VIRTUAL_INTR_PENDING;
    }
    __vmx_vmwrite(VMCS_CTRL_PROCESSOR_BASED_VM_EXECUTION_CONTROLS, CpuBasedVmExecControls);
}*/
return true
}

func (h *hv)HvSetNmiWindowExiting()(ok bool){//col:269
/*HvSetNmiWindowExiting(BOOLEAN Set)
{
    ULONG CpuBasedVmExecControls = 0;
    __vmx_vmread(VMCS_CTRL_PROCESSOR_BASED_VM_EXECUTION_CONTROLS, &CpuBasedVmExecControls);
    if (Set)
    {
        CpuBasedVmExecControls |= CPU_BASED_VIRTUAL_NMI_PENDING;
    }
    else
    {
        CpuBasedVmExecControls &= ~CPU_BASED_VIRTUAL_NMI_PENDING;
    }
    __vmx_vmwrite(VMCS_CTRL_PROCESSOR_BASED_VM_EXECUTION_CONTROLS, CpuBasedVmExecControls);
}*/
return true
}

func (h *hv)HvHandleMovDebugRegister()(ok bool){//col:377
/*HvHandleMovDebugRegister(UINT32 ProcessorIndex, PGUEST_REGS Regs)
{
    VMX_EXIT_QUALIFICATION_MOV_DR ExitQualification;
    CR4                           Cr4;
    DR7                           Dr7;
    VMX_SEGMENT_SELECTOR          Cs;
    UINT64 *                      GpRegs         = Regs;
    VIRTUAL_MACHINE_STATE *       CurrentVmState = &g_GuestState[ProcessorIndex];
    __vmx_vmread(VMCS_EXIT_QUALIFICATION, &ExitQualification);
    UINT64 GpRegister = GpRegs[ExitQualification.GeneralPurposeRegister];
    Cs = GetGuestCs();
    if (Cs.Attributes.DescriptorPrivilegeLevel != 0)
    {
        EventInjectGeneralProtection();
        CurrentVmState->IncrementRip = FALSE;
        return;
    }
    __vmx_vmread(VMCS_GUEST_CR4, &Cr4);
    if (ExitQualification.DebugRegister == 4 || ExitQualification.DebugRegister == 5)
    {
        if (Cr4.DebuggingExtensions)
        {
            EventInjectUndefinedOpcode(ProcessorIndex);
            return;
        }
        else
        {
            ExitQualification.DebugRegister += 2;
        }
    }
    __vmx_vmread(VMCS_GUEST_DR7, &Dr7);
    if (Dr7.GeneralDetect)
    {
        DR6 Dr6 = {
            .AsUInt                      = __readdr(6),
            .BreakpointCondition         = 0,
            .DebugRegisterAccessDetected = TRUE};
        __writedr(6, Dr6.AsUInt);
        Dr7.GeneralDetect = FALSE;
        __vmx_vmwrite(VMCS_GUEST_DR7, Dr7.AsUInt);
        EventInjectDebugBreakpoint();
        CurrentVmState->IncrementRip = FALSE;
        return;
    }
    if (ExitQualification.DirectionOfAccess == VMX_EXIT_QUALIFICATION_DIRECTION_MOV_TO_DR &&
        (ExitQualification.DebugRegister == VMX_EXIT_QUALIFICATION_REGISTER_DR6 ||
         ExitQualification.DebugRegister == VMX_EXIT_QUALIFICATION_REGISTER_DR7) &&
        (GpRegister >> 32) != 0)
    {
        EventInjectGeneralProtection();
        CurrentVmState->IncrementRip = FALSE;
        return;
    }
    switch (ExitQualification.DirectionOfAccess)
    {
    case VMX_EXIT_QUALIFICATION_DIRECTION_MOV_TO_DR:
        switch (ExitQualification.DebugRegister)
        {
        case VMX_EXIT_QUALIFICATION_REGISTER_DR0:
            __writedr(VMX_EXIT_QUALIFICATION_REGISTER_DR0, GpRegister);
            break;
        case VMX_EXIT_QUALIFICATION_REGISTER_DR1:
            __writedr(VMX_EXIT_QUALIFICATION_REGISTER_DR1, GpRegister);
            break;
        case VMX_EXIT_QUALIFICATION_REGISTER_DR2:
            __writedr(VMX_EXIT_QUALIFICATION_REGISTER_DR2, GpRegister);
            break;
        case VMX_EXIT_QUALIFICATION_REGISTER_DR3:
            __writedr(VMX_EXIT_QUALIFICATION_REGISTER_DR3, GpRegister);
            break;
        case VMX_EXIT_QUALIFICATION_REGISTER_DR6:
            __writedr(VMX_EXIT_QUALIFICATION_REGISTER_DR6, GpRegister);
            break;
        case VMX_EXIT_QUALIFICATION_REGISTER_DR7:
            __writedr(VMX_EXIT_QUALIFICATION_REGISTER_DR7, GpRegister);
            break;
        default:
            break;
        }
        break;
    case VMX_EXIT_QUALIFICATION_DIRECTION_MOV_FROM_DR:
        switch (ExitQualification.DebugRegister)
        {
        case VMX_EXIT_QUALIFICATION_REGISTER_DR0:
            GpRegister = __readdr(VMX_EXIT_QUALIFICATION_REGISTER_DR0);
            break;
        case VMX_EXIT_QUALIFICATION_REGISTER_DR1:
            GpRegister = __readdr(VMX_EXIT_QUALIFICATION_REGISTER_DR1);
            break;
        case VMX_EXIT_QUALIFICATION_REGISTER_DR2:
            GpRegister = __readdr(VMX_EXIT_QUALIFICATION_REGISTER_DR2);
            break;
        case VMX_EXIT_QUALIFICATION_REGISTER_DR3:
            GpRegister = __readdr(VMX_EXIT_QUALIFICATION_REGISTER_DR3);
            break;
        case VMX_EXIT_QUALIFICATION_REGISTER_DR6:
            GpRegister = __readdr(VMX_EXIT_QUALIFICATION_REGISTER_DR6);
            break;
        case VMX_EXIT_QUALIFICATION_REGISTER_DR7:
            GpRegister = __readdr(VMX_EXIT_QUALIFICATION_REGISTER_DR7);
            break;
        default:
            break;
        }
    default:
        break;
    }
}*/
return true
}

func (h *hv)HvSetNmiExiting()(ok bool){//col:396
/*HvSetNmiExiting(BOOLEAN Set)
{
    ULONG PinBasedControls = 0;
    ULONG VmExitControls   = 0;
    __vmx_vmread(VMCS_CTRL_PIN_BASED_VM_EXECUTION_CONTROLS, &PinBasedControls);
    __vmx_vmread(VMCS_CTRL_PRIMARY_VMEXIT_CONTROLS, &VmExitControls);
    if (Set)
    {
        PinBasedControls |= PIN_BASED_VM_EXECUTION_CONTROLS_NMI_EXITING;
        VmExitControls |= VM_EXIT_ACK_INTR_ON_EXIT;
    }
    else
    {
        PinBasedControls &= ~PIN_BASED_VM_EXECUTION_CONTROLS_NMI_EXITING;
        VmExitControls &= ~VM_EXIT_ACK_INTR_ON_EXIT;
    }
    __vmx_vmwrite(VMCS_CTRL_PIN_BASED_VM_EXECUTION_CONTROLS, PinBasedControls);
    __vmx_vmwrite(VMCS_CTRL_PRIMARY_VMEXIT_CONTROLS, VmExitControls);
}*/
return true
}

func (h *hv)HvSetVmxPreemptionTimerExiting()(ok bool){//col:410
/*HvSetVmxPreemptionTimerExiting(BOOLEAN Set)
{
    ULONG PinBasedControls = 0;
    __vmx_vmread(VMCS_CTRL_PIN_BASED_VM_EXECUTION_CONTROLS, &PinBasedControls);
    if (Set)
    {
        PinBasedControls |= PIN_BASED_VM_EXECUTION_CONTROLS_ACTIVE_VMX_TIMER;
    }
    else
    {
        PinBasedControls &= ~PIN_BASED_VM_EXECUTION_CONTROLS_ACTIVE_VMX_TIMER;
    }
    __vmx_vmwrite(VMCS_CTRL_PIN_BASED_VM_EXECUTION_CONTROLS, PinBasedControls);
}*/
return true
}

func (h *hv)HvSetExceptionBitmap()(ok bool){//col:414
/*HvSetExceptionBitmap(UINT32 IdtIndex)
{
    ProtectedHvSetExceptionBitmap(IdtIndex);
}*/
return true
}

func (h *hv)HvUnsetExceptionBitmap()(ok bool){//col:418
/*HvUnsetExceptionBitmap(UINT32 IdtIndex)
{
    ProtectedHvUnsetExceptionBitmap(IdtIndex);
}*/
return true
}

func (h *hv)HvSetExternalInterruptExiting()(ok bool){//col:422
/*HvSetExternalInterruptExiting(BOOLEAN Set)
{
    ProtectedHvSetExternalInterruptExiting(Set);
}*/
return true
}

func (h *hv)HvSetRdtscExiting()(ok bool){//col:426
/*HvSetRdtscExiting(BOOLEAN Set)
{
    ProtectedHvSetRdtscExiting(Set);
}*/
return true
}

func (h *hv)HvSetMovDebugRegsExiting()(ok bool){//col:430
/*HvSetMovDebugRegsExiting(BOOLEAN Set)
{
    ProtectedHvSetMovDebugRegsExiting(Set);
}*/
return true
}



