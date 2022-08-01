package vmx
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbghv\code\vmm\vmx\Vmx.c.back

type (
Vmx interface{
VmxCheckVmxSupport()(ok bool)//col:17
VmxInitialize()(ok bool)//col:55
VmxPerformVirtualizationOnAllCores()(ok bool)//col:100
VmxPerformVirtualizationOnSpecificCore()(ok bool)//col:120
VmxFixCr4AndCr0Bits()(ok bool)//col:138
VmxCheckIsOnVmxRoot()(ok bool)//col:156
VmxVirtualizeCurrentSystem()(ok bool)//col:184
VmxTerminate()(ok bool)//col:203
VmxVmptrst()(ok bool)//col:210
VmxClearVmcsState()(ok bool)//col:223
VmxLoadVmcs()(ok bool)//col:234
VmxSetupVmcs()(ok bool)//col:327
VmxVmresume()(ok bool)//col:335
VmxVmxoff()(ok bool)//col:359
VmxReturnStackPointerForVmxoff()(ok bool)//col:363
VmxReturnInstructionPointerForVmxoff()(ok bool)//col:367
VmxPerformTermination()(ok bool)//col:379
}
vmx struct{}
)

func NewVmx()Vmx{ return & vmx{} }

func (v *vmx)VmxCheckVmxSupport()(ok bool){//col:17
/*VmxCheckVmxSupport()
{
    CPUID                         Data              = {0};
    IA32_FEATURE_CONTROL_REGISTER FeatureControlMsr = {0};
    __cpuid((int *)&Data, 1);
    if (!_bittest((const LONG *)&Data.ecx, 5))
    {
        return FALSE;
    }
    FeatureControlMsr.AsUInt = __readmsr(IA32_FEATURE_CONTROL);
    if (FeatureControlMsr.EnableVmxOutsideSmx == FALSE)
    {
        LogError("Err, you should enable vt-x from BIOS");
        return FALSE;
    }
    return TRUE;
}*/
return true
}

func (v *vmx)VmxInitialize()(ok bool){//col:55
/*VmxInitialize()
{
    ULONG LogicalProcessorsCount;
    if (!VmxPerformVirtualizationOnAllCores())
    {
        return FALSE;
    }
    LogicalProcessorsCount = KeQueryActiveProcessorCount(0);
    for (size_t ProcessorID = 0; ProcessorID < LogicalProcessorsCount; ProcessorID++)
    {
        if (!VmxAllocateVmmStack(ProcessorID))
        {
            return FALSE;
        }
        if (!VmxAllocateMsrBitmap(ProcessorID))
        {
            return FALSE;
        }
        if (!VmxAllocateIoBitmaps(ProcessorID))
        {
            return FALSE;
        }
    }
    g_MsrBitmapInvalidMsrs = AllocateInvalidMsrBimap();
    if (g_MsrBitmapInvalidMsrs == NULL)
    {
        return FALSE;
    }
    KeGenericCallDpc(DpcRoutineInitializeGuest, 0x0);
    if (AsmVmxVmcall(VMCALL_TEST, 0x22, 0x333, 0x4444) == STATUS_SUCCESS)
    {
        return TRUE;
    }
    else
    {
        return FALSE;
    }
}*/
return true
}

func (v *vmx)VmxPerformVirtualizationOnAllCores()(ok bool){//col:100
/*VmxPerformVirtualizationOnAllCores()
{
    int       ProcessorCount;
    KAFFINITY AffinityMask;
    if (!VmxCheckVmxSupport())
    {
        LogError("Err, VMX is not supported in this machine");
        return FALSE;
    }
    PAGED_CODE();
    g_EptState = ExAllocatePoolWithTag(NonPagedPool, sizeof(EPT_STATE), POOLTAG);
    if (!g_EptState)
    {
        LogError("Err, insufficient memory");
        return FALSE;
    }
    RtlZeroMemory(g_EptState, sizeof(EPT_STATE));
    InitializeListHead(&g_EptState->HookedPagesList);
    if (!EptCheckFeatures())
    {
        LogError("Err, your processor doesn't support all EPT features");
        return FALSE;
    }
    else
    {
        LogDebugInfo("Your processor supports all EPT features");
        if (!EptBuildMtrrMap())
        {
            LogError("Err, could not build MTRR memory map");
            return FALSE;
        }
        LogDebugInfo("MTRR memory map built successfully");
    }
    if (!PoolManagerInitialize())
    {
        LogError("Err, could not initialize pool manager");
        return FALSE;
    }
    if (!EptLogicalProcessorInitialize())
    {
        return FALSE;
    }
    BroadcastVmxVirtualizationAllCores();
    return TRUE;
}*/
return true
}

func (v *vmx)VmxPerformVirtualizationOnSpecificCore()(ok bool){//col:120
/*VmxPerformVirtualizationOnSpecificCore()
{
    ULONG                   CurrentProcessorNumber = KeGetCurrentProcessorNumber();
    VIRTUAL_MACHINE_STATE * CurrentVmState         = &g_GuestState[CurrentProcessorNumber];
    LogDebugInfo("Allocating vmx regions for logical core %d", CurrentProcessorNumber);
    AsmEnableVmxOperation();
    VmxFixCr4AndCr0Bits();
    LogDebugInfo("VMX-Operation enabled successfully");
    if (!VmxAllocateVmxonRegion(CurrentVmState))
    {
        LogError("Err, allocating memory for vmxon region was not successfull");
        return FALSE;
    }
    if (!VmxAllocateVmcsRegion(CurrentVmState))
    {
        LogError("Err, allocating memory for vmcs region was not successfull");
        return FALSE;
    }
    return TRUE;
}*/
return true
}

func (v *vmx)VmxFixCr4AndCr0Bits()(ok bool){//col:138
/*VmxFixCr4AndCr0Bits()
{
    CR_FIXED CrFixed = {0};
    CR4      Cr4     = {0};
    CR0      Cr0     = {0};
    CrFixed.Flags = __readmsr(IA32_VMX_CR0_FIXED0);
    Cr0.AsUInt    = __readcr0();
    Cr0.AsUInt |= CrFixed.Fields.Low;
    CrFixed.Flags = __readmsr(IA32_VMX_CR0_FIXED1);
    Cr0.AsUInt &= CrFixed.Fields.Low;
    __writecr0(Cr0.AsUInt);
    CrFixed.Flags = __readmsr(IA32_VMX_CR4_FIXED0);
    Cr4.AsUInt    = __readcr4();
    Cr4.AsUInt |= CrFixed.Fields.Low;
    CrFixed.Flags = __readmsr(IA32_VMX_CR4_FIXED1);
    Cr4.AsUInt &= CrFixed.Fields.Low;
    __writecr4(Cr4.AsUInt);
}*/
return true
}

func (v *vmx)VmxCheckIsOnVmxRoot()(ok bool){//col:156
/*VmxCheckIsOnVmxRoot()
{
    UINT64 VmcsLink = 0;
    __try
    {
        if (!__vmx_vmread(VMCS_GUEST_VMCS_LINK_POINTER, &VmcsLink))
        {
            if (VmcsLink != 0)
            {
                return TRUE;
            }
        }
    }
    __except (1)
    {
    }
    return FALSE;
}*/
return true
}

func (v *vmx)VmxVirtualizeCurrentSystem()(ok bool){//col:184
/*VmxVirtualizeCurrentSystem(PVOID GuestStack)
{
    UINT64                  ErrorCode      = 0;
    ULONG                   ProcessorID    = KeGetCurrentProcessorNumber();
    VIRTUAL_MACHINE_STATE * CurrentVmState = &g_GuestState[ProcessorID];
    LogDebugInfo("Virtualizing current system (logical core : 0x%x)", ProcessorID);
    if (!VmxClearVmcsState(CurrentVmState))
    {
        LogError("Err, failed to clear vmcs");
        return FALSE;
    }
    if (!VmxLoadVmcs(CurrentVmState))
    {
        LogError("Err, failed to load vmcs");
        return FALSE;
    }
    LogDebugInfo("Setting up VMCS for current logical core");
    VmxSetupVmcs(CurrentVmState, GuestStack);
    LogDebugInfo("Executing VMLAUNCH on logical core %d", ProcessorID);
    CurrentVmState->HasLaunched = TRUE;
    __vmx_vmlaunch();
    CurrentVmState->HasLaunched = FALSE;
    __vmx_vmread(VMCS_VM_INSTRUCTION_ERROR, &ErrorCode);
    LogError("Err, unable to execute VMLAUNCH, status : 0x%llx", ErrorCode);
    __vmx_off();
    LogError("Err, VMXOFF Executed Successfully but it was because of an error");
    return FALSE;
}*/
return true
}

func (v *vmx)VmxTerminate()(ok bool){//col:203
/*VmxTerminate()
{
    NTSTATUS                Status           = STATUS_SUCCESS;
    ULONG                   CurrentCoreIndex = KeGetCurrentProcessorNumber();
    VIRTUAL_MACHINE_STATE * CurrentVmState   = &g_GuestState[CurrentCoreIndex];
    Status = AsmVmxVmcall(VMCALL_VMXOFF, NULL, NULL, NULL);
    if (Status == STATUS_SUCCESS)
    {
        LogDebugInfo("VMX terminated on logical core %d\n", CurrentCoreIndex);
        MmFreeContiguousMemory(CurrentVmState->VmxonRegionVirtualAddress);
        MmFreeContiguousMemory(CurrentVmState->VmcsRegionVirtualAddress);
        ExFreePoolWithTag(CurrentVmState->VmmStack, POOLTAG);
        ExFreePoolWithTag(CurrentVmState->MsrBitmapVirtualAddress, POOLTAG);
        ExFreePoolWithTag(CurrentVmState->IoBitmapVirtualAddressA, POOLTAG);
        ExFreePoolWithTag(CurrentVmState->IoBitmapVirtualAddressB, POOLTAG);
        return TRUE;
    }
    return FALSE;
}*/
return true
}

func (v *vmx)VmxVmptrst()(ok bool){//col:210
/*VmxVmptrst()
{
    PHYSICAL_ADDRESS VmcsPhysicalAddr;
    VmcsPhysicalAddr.QuadPart = 0;
    __vmx_vmptrst((unsigned __int64 *)&VmcsPhysicalAddr);
    LogDebugInfo("VMPTRST result : %llx", VmcsPhysicalAddr);
}*/
return true
}

func (v *vmx)VmxClearVmcsState()(ok bool){//col:223
/*VmxClearVmcsState(VIRTUAL_MACHINE_STATE * CurrentGuestState)
{
    UINT8 VmclearStatus;
    VmclearStatus = __vmx_vmclear(&CurrentGuestState->VmcsRegionPhysicalAddress);
    LogDebugInfo("VMCS VMCLEAR status : 0x%x", VmclearStatus);
    if (VmclearStatus)
    {
        LogDebugInfo("VMCS failed to clear, status : 0x%x", VmclearStatus);
        __vmx_off();
        return FALSE;
    }
    return TRUE;
}*/
return true
}

func (v *vmx)VmxLoadVmcs()(ok bool){//col:234
/*VmxLoadVmcs(VIRTUAL_MACHINE_STATE * CurrentGuestState)
{
    int VmptrldStatus;
    VmptrldStatus = __vmx_vmptrld(&CurrentGuestState->VmcsRegionPhysicalAddress);
    if (VmptrldStatus)
    {
        LogDebugInfo("VMCS failed to load, status : 0x%x", VmptrldStatus);
        return FALSE;
    }
    return TRUE;
}*/
return true
}

func (v *vmx)VmxSetupVmcs()(ok bool){//col:327
/*VmxSetupVmcs(VIRTUAL_MACHINE_STATE * CurrentGuestState, PVOID GuestStack)
{
    ULONG                   CpuBasedVmExecControls;
    ULONG                   SecondaryProcBasedVmExecControls;
    PVOID                   HostRsp;
    UINT64                  GdtBase         = 0;
    VMX_SEGMENT_SELECTOR    SegmentSelector = {0};
    IA32_VMX_BASIC_REGISTER VmxBasicMsr     = {0};
    VmxBasicMsr.AsUInt = __readmsr(IA32_VMX_BASIC);
    __vmx_vmwrite(VMCS_HOST_ES_SELECTOR, AsmGetEs() & 0xF8);
    __vmx_vmwrite(VMCS_HOST_CS_SELECTOR, AsmGetCs() & 0xF8);
    __vmx_vmwrite(VMCS_HOST_SS_SELECTOR, AsmGetSs() & 0xF8);
    __vmx_vmwrite(VMCS_HOST_DS_SELECTOR, AsmGetDs() & 0xF8);
    __vmx_vmwrite(VMCS_HOST_FS_SELECTOR, AsmGetFs() & 0xF8);
    __vmx_vmwrite(VMCS_HOST_GS_SELECTOR, AsmGetGs() & 0xF8);
    __vmx_vmwrite(VMCS_HOST_TR_SELECTOR, AsmGetTr() & 0xF8);
    __vmx_vmwrite(VMCS_GUEST_VMCS_LINK_POINTER, ~0ULL);
    __vmx_vmwrite(VMCS_GUEST_DEBUGCTL, __readmsr(IA32_DEBUGCTL) & 0xFFFFFFFF);
    __vmx_vmwrite(VMCS_GUEST_DEBUGCTL_HIGH, __readmsr(IA32_DEBUGCTL) >> 32);
    __vmx_vmwrite(VMCS_CTRL_TSC_OFFSET, 0);
    __vmx_vmwrite(VMCS_CTRL_PAGEFAULT_ERROR_CODE_MASK, 0);
    __vmx_vmwrite(VMCS_CTRL_PAGEFAULT_ERROR_CODE_MATCH, 0);
    __vmx_vmwrite(VMCS_CTRL_VMEXIT_MSR_STORE_COUNT, 0);
    __vmx_vmwrite(VMCS_CTRL_VMEXIT_MSR_LOAD_COUNT, 0);
    __vmx_vmwrite(VMCS_CTRL_VMENTRY_MSR_LOAD_COUNT, 0);
    __vmx_vmwrite(VMCS_CTRL_VMENTRY_INTERRUPTION_INFORMATION_FIELD, 0);
    GdtBase = AsmGetGdtBase();
    HvFillGuestSelectorData((PVOID)GdtBase, ES, AsmGetEs());
    HvFillGuestSelectorData((PVOID)GdtBase, CS, AsmGetCs());
    HvFillGuestSelectorData((PVOID)GdtBase, SS, AsmGetSs());
    HvFillGuestSelectorData((PVOID)GdtBase, DS, AsmGetDs());
    HvFillGuestSelectorData((PVOID)GdtBase, FS, AsmGetFs());
    HvFillGuestSelectorData((PVOID)GdtBase, GS, AsmGetGs());
    HvFillGuestSelectorData((PVOID)GdtBase, LDTR, AsmGetLdtr());
    HvFillGuestSelectorData((PVOID)GdtBase, TR, AsmGetTr());
    __vmx_vmwrite(VMCS_GUEST_FS_BASE, __readmsr(IA32_FS_BASE));
    __vmx_vmwrite(VMCS_GUEST_GS_BASE, __readmsr(IA32_GS_BASE));
    CpuBasedVmExecControls = HvAdjustControls(CPU_BASED_ACTIVATE_IO_BITMAP | CPU_BASED_ACTIVATE_MSR_BITMAP | CPU_BASED_ACTIVATE_SECONDARY_CONTROLS,
                                              VmxBasicMsr.VmxControls ? IA32_VMX_TRUE_PROCBASED_CTLS : IA32_VMX_PROCBASED_CTLS);
    __vmx_vmwrite(VMCS_CTRL_PROCESSOR_BASED_VM_EXECUTION_CONTROLS, CpuBasedVmExecControls);
    LogDebugInfo("CPU Based VM Exec Controls (Based on %s) : 0x%x",
                 VmxBasicMsr.VmxControls ? "IA32_VMX_TRUE_PROCBASED_CTLS" : "IA32_VMX_PROCBASED_CTLS",
                 CpuBasedVmExecControls);
    SecondaryProcBasedVmExecControls = HvAdjustControls(CPU_BASED_CTL2_RDTSCP |
                                                            CPU_BASED_CTL2_ENABLE_EPT | CPU_BASED_CTL2_ENABLE_INVPCID |
                                                            CPU_BASED_CTL2_ENABLE_XSAVE_XRSTORS | CPU_BASED_CTL2_ENABLE_VPID,
                                                        IA32_VMX_PROCBASED_CTLS2);
    __vmx_vmwrite(VMCS_CTRL_SECONDARY_PROCESSOR_BASED_VM_EXECUTION_CONTROLS, SecondaryProcBasedVmExecControls);
    LogDebugInfo("Secondary Proc Based VM Exec Controls (IA32_VMX_PROCBASED_CTLS2) : 0x%x", SecondaryProcBasedVmExecControls);
    __vmx_vmwrite(VMCS_CTRL_PIN_BASED_VM_EXECUTION_CONTROLS, HvAdjustControls(0, VmxBasicMsr.VmxControls ? IA32_VMX_TRUE_PINBASED_CTLS : IA32_VMX_PINBASED_CTLS));
    __vmx_vmwrite(VMCS_CTRL_PRIMARY_VMEXIT_CONTROLS, HvAdjustControls(VM_EXIT_HOST_ADDR_SPACE_SIZE, VmxBasicMsr.VmxControls ? IA32_VMX_TRUE_EXIT_CTLS : IA32_VMX_EXIT_CTLS));
    __vmx_vmwrite(VMCS_CTRL_VMENTRY_CONTROLS, HvAdjustControls(VM_ENTRY_IA32E_MODE, VmxBasicMsr.VmxControls ? IA32_VMX_TRUE_ENTRY_CTLS : IA32_VMX_ENTRY_CTLS));
    __vmx_vmwrite(VMCS_CTRL_CR0_GUEST_HOST_MASK, 0);
    __vmx_vmwrite(VMCS_CTRL_CR4_GUEST_HOST_MASK, 0);
    __vmx_vmwrite(VMCS_CTRL_CR0_READ_SHADOW, 0);
    __vmx_vmwrite(VMCS_CTRL_CR4_READ_SHADOW, 0);
    __vmx_vmwrite(VMCS_GUEST_CR0, __readcr0());
    __vmx_vmwrite(VMCS_GUEST_CR3, __readcr3());
    __vmx_vmwrite(VMCS_GUEST_CR4, __readcr4());
    __vmx_vmwrite(VMCS_GUEST_DR7, 0x400);
    __vmx_vmwrite(VMCS_HOST_CR0, __readcr0());
    __vmx_vmwrite(VMCS_HOST_CR4, __readcr4());
    __vmx_vmwrite(VMCS_HOST_CR3, FindSystemDirectoryTableBase());
    __vmx_vmwrite(VMCS_GUEST_GDTR_BASE, AsmGetGdtBase());
    __vmx_vmwrite(VMCS_GUEST_IDTR_BASE, AsmGetIdtBase());
    __vmx_vmwrite(VMCS_GUEST_GDTR_LIMIT, AsmGetGdtLimit());
    __vmx_vmwrite(VMCS_GUEST_IDTR_LIMIT, AsmGetIdtLimit());
    __vmx_vmwrite(VMCS_GUEST_RFLAGS, AsmGetRflags());
    __vmx_vmwrite(VMCS_GUEST_SYSENTER_CS, __readmsr(IA32_SYSENTER_CS));
    __vmx_vmwrite(VMCS_GUEST_SYSENTER_EIP, __readmsr(IA32_SYSENTER_EIP));
    __vmx_vmwrite(VMCS_GUEST_SYSENTER_ESP, __readmsr(IA32_SYSENTER_ESP));
    GetSegmentDescriptor((PUCHAR)AsmGetGdtBase(), AsmGetTr(), &SegmentSelector);
    __vmx_vmwrite(VMCS_HOST_TR_BASE, SegmentSelector.Base);
    __vmx_vmwrite(VMCS_HOST_FS_BASE, __readmsr(IA32_FS_BASE));
    __vmx_vmwrite(VMCS_HOST_GS_BASE, __readmsr(IA32_GS_BASE));
    __vmx_vmwrite(VMCS_HOST_GDTR_BASE, AsmGetGdtBase());
    __vmx_vmwrite(VMCS_HOST_IDTR_BASE, AsmGetIdtBase());
    __vmx_vmwrite(VMCS_HOST_SYSENTER_CS, __readmsr(IA32_SYSENTER_CS));
    __vmx_vmwrite(VMCS_HOST_SYSENTER_EIP, __readmsr(IA32_SYSENTER_EIP));
    __vmx_vmwrite(VMCS_HOST_SYSENTER_ESP, __readmsr(IA32_SYSENTER_ESP));
    __vmx_vmwrite(VMCS_CTRL_MSR_BITMAP_ADDRESS, CurrentGuestState->MsrBitmapPhysicalAddress);
    __vmx_vmwrite(VMCS_CTRL_IO_BITMAP_A_ADDRESS, CurrentGuestState->IoBitmapPhysicalAddressA);
    __vmx_vmwrite(VMCS_CTRL_IO_BITMAP_B_ADDRESS, CurrentGuestState->IoBitmapPhysicalAddressB);
    __vmx_vmwrite(VMCS_CTRL_EPT_POINTER, g_EptState->EptPointer.AsUInt);
    __vmx_vmwrite(VIRTUAL_PROCESSOR_ID, VPID_TAG);
    __vmx_vmwrite(VMCS_GUEST_RSP, (UINT64)GuestStack);
    __vmx_vmwrite(VMCS_GUEST_RIP, (UINT64)AsmVmxRestoreState);
    HostRsp = (UINT64)CurrentGuestState->VmmStack + VMM_STACK_SIZE - 1;
    HostRsp = ((PVOID)((ULONG_PTR)(HostRsp) & ~(16 - 1)));
    __vmx_vmwrite(VMCS_HOST_RSP, HostRsp);
    __vmx_vmwrite(VMCS_HOST_RIP, (UINT64)AsmVmexitHandler);
    return TRUE;
}*/
return true
}

func (v *vmx)VmxVmresume()(ok bool){//col:335
/*VmxVmresume()
{
    UINT64 ErrorCode = 0;
    __vmx_vmresume();
    __vmx_vmread(VMCS_VM_INSTRUCTION_ERROR, &ErrorCode);
    __vmx_off();
    LogError("Err,  in executing VMRESUME , status : 0x%llx", ErrorCode);
}*/
return true
}

func (v *vmx)VmxVmxoff()(ok bool){//col:359
/*VmxVmxoff()
{
    ULONG  CurrentProcessorIndex = 0;
    UINT64 GuestRSP              = 0; 
    UINT64 GuestRIP              = 0; 
    UINT64 GuestCr3              = 0;
    UINT64 ExitInstructionLength = 0;
    CurrentProcessorIndex                  = KeGetCurrentProcessorNumber();
    VIRTUAL_MACHINE_STATE * CurrentVmState = &g_GuestState[CurrentProcessorIndex];
    __vmx_vmread(VMCS_GUEST_CR3, &GuestCr3);
    __writecr3(GuestCr3);
    __vmx_vmread(VMCS_GUEST_RIP, &GuestRIP);
    __vmx_vmread(VMCS_GUEST_RSP, &GuestRSP);
    __vmx_vmread(VMCS_VMEXIT_INSTRUCTION_LENGTH, &ExitInstructionLength);
    GuestRIP += ExitInstructionLength;
    CurrentVmState->VmxoffState.GuestRip = GuestRIP;
    CurrentVmState->VmxoffState.GuestRsp = GuestRSP;
    CurrentVmState->VmxoffState.IsVmxoffExecuted = TRUE;
    HvRestoreRegisters();
    VmxClearVmcsState(CurrentVmState);
    __vmx_off();
    CurrentVmState->HasLaunched = FALSE;
    __writecr4(__readcr4() & (~X86_CR4_VMXE));
}*/
return true
}

func (v *vmx)VmxReturnStackPointerForVmxoff()(ok bool){//col:363
/*VmxReturnStackPointerForVmxoff()
{
    return g_GuestState[KeGetCurrentProcessorNumber()].VmxoffState.GuestRsp;
}*/
return true
}

func (v *vmx)VmxReturnInstructionPointerForVmxoff()(ok bool){//col:367
/*VmxReturnInstructionPointerForVmxoff()
{
    return g_GuestState[KeGetCurrentProcessorNumber()].VmxoffState.GuestRip;
}*/
return true
}

func (v *vmx)VmxPerformTermination()(ok bool){//col:379
/*VmxPerformTermination()
{
    LogDebugInfo("Terminating VMX ...\n");
    TransparentUnhideDebugger();
    EptHookUnHookAll();
    KeGenericCallDpc(DpcRoutineTerminateGuest, 0x0);
    ExFreePoolWithTag(g_MsrBitmapInvalidMsrs, POOLTAG);
    MmFreeContiguousMemory(g_EptState->EptPageTable);
    ExFreePoolWithTag(g_EptState, POOLTAG);
    PoolManagerUninitialize();
    LogDebugInfo("VMX operation turned off successfully");
}*/
return true
}



