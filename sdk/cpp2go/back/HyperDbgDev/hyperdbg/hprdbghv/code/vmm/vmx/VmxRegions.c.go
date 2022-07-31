package vmx
//back\HyperDbgDev\hyperdbg\hprdbghv\code\vmm\vmx\VmxRegions.c.back

type (
VmxRegions interface{
VmxAllocateVmxonRegion()(ok bool)//col:92
VmxAllocateVmcsRegion()(ok bool)//col:158
VmxAllocateVmmStack()(ok bool)//col:186
VmxAllocateMsrBitmap()(ok bool)//col:217
VmxAllocateIoBitmaps()(ok bool)//col:263
}
)

func NewVmxRegions() { return & vmxRegions{} }

func (v *vmxRegions)VmxAllocateVmxonRegion()(ok bool){//col:92
/*VmxAllocateVmxonRegion(VIRTUAL_MACHINE_STATE * CurrentGuestState)
{
    IA32_VMX_BASIC_REGISTER VmxBasicMsr = {0};
    SIZE_T                  VmxonSize;
    UINT8                   VmxonStatus;
    UINT8 *                 VmxonRegion;
    UINT64                  VmxonRegionPhysicalAddr;
    UINT64                  AlignedVmxonRegion;
    UINT64                  AlignedVmxonRegionPhysicalAddr;
#ifdef ENV_WINDOWS
    if (KeGetCurrentIrql() > DISPATCH_LEVEL)
        KeRaiseIrqlToDpcLevel();
    VmxonSize   = 2 * VMXON_SIZE;
    VmxonRegion = CrsAllocateContiguousZeroedMemory(VmxonSize + ALIGNMENT_PAGE_SIZE);
    if (VmxonRegion == NULL)
    {
        LogError("Err, couldn't allocate buffer for VMXON region");
        return FALSE;
    }
    VmxonRegionPhysicalAddr = VirtualAddressToPhysicalAddress(VmxonRegion);
    AlignedVmxonRegion = (BYTE *)((ULONG_PTR)(VmxonRegion + ALIGNMENT_PAGE_SIZE - 1) & ~(ALIGNMENT_PAGE_SIZE - 1));
    LogDebugInfo("VMXON Region Address : %llx", AlignedVmxonRegion);
    AlignedVmxonRegionPhysicalAddr = (BYTE *)((ULONG_PTR)(VmxonRegionPhysicalAddr + ALIGNMENT_PAGE_SIZE - 1) & ~(ALIGNMENT_PAGE_SIZE - 1));
    LogDebugInfo("VMXON Region Physical Address : %llx", AlignedVmxonRegionPhysicalAddr);
    VmxBasicMsr.AsUInt = __readmsr(IA32_VMX_BASIC);
    LogDebugInfo("Revision Identifier (IA32_VMX_BASIC - MSR 0x480) : 0x%x", VmxBasicMsr.VmcsRevisionId);
    *(UINT64 *)AlignedVmxonRegion = VmxBasicMsr.VmcsRevisionId;
    VmxonStatus = __vmx_on(&AlignedVmxonRegionPhysicalAddr);
    if (VmxonStatus)
    {
        LogError("Err, executing vmxon instruction failed with status : %d", VmxonStatus);
        return FALSE;
    }
    CurrentGuestState->VmxonRegionPhysicalAddress = AlignedVmxonRegionPhysicalAddr;
    CurrentGuestState->VmxonRegionVirtualAddress = VmxonRegion;
    return TRUE;
}*/
return true
}

func (v *vmxRegions)VmxAllocateVmcsRegion()(ok bool){//col:158
/*VmxAllocateVmcsRegion(VIRTUAL_MACHINE_STATE * CurrentGuestState)
{
    IA32_VMX_BASIC_REGISTER VmxBasicMsr = {0};
    SIZE_T                  VmcsSize;
    UINT8 *                 VmcsRegion;
    UINT64                  VmcsPhysicalAddr;
    UINT64                  AlignedVmcsRegion;
    UINT64                  AlignedVmcsRegionPhysicalAddr;
#ifdef ENV_WINDOWS
    if (KeGetCurrentIrql() > DISPATCH_LEVEL)
        KeRaiseIrqlToDpcLevel();
    VmcsSize   = 2 * VMCS_SIZE;
    VmcsRegion = CrsAllocateContiguousZeroedMemory(VmcsSize + ALIGNMENT_PAGE_SIZE);
    if (VmcsRegion == NULL)
    {
        LogError("Err, couldn't allocate Buffer for VMCS region");
        return FALSE;
    }
    VmcsPhysicalAddr = VirtualAddressToPhysicalAddress(VmcsRegion);
    AlignedVmcsRegion = (BYTE *)((ULONG_PTR)(VmcsRegion + ALIGNMENT_PAGE_SIZE - 1) & ~(ALIGNMENT_PAGE_SIZE - 1));
    LogDebugInfo("VMCS region address : %llx", AlignedVmcsRegion);
    AlignedVmcsRegionPhysicalAddr = (BYTE *)((ULONG_PTR)(VmcsPhysicalAddr + ALIGNMENT_PAGE_SIZE - 1) & ~(ALIGNMENT_PAGE_SIZE - 1));
    LogDebugInfo("VMCS region physical address : %llx", AlignedVmcsRegionPhysicalAddr);
    VmxBasicMsr.AsUInt = __readmsr(IA32_VMX_BASIC);
    LogDebugInfo("Revision Identifier (IA32_VMX_BASIC - MSR 0x480) : 0x%x", VmxBasicMsr.VmcsRevisionId);
    *(UINT64 *)AlignedVmcsRegion = VmxBasicMsr.VmcsRevisionId;
    CurrentGuestState->VmcsRegionPhysicalAddress = AlignedVmcsRegionPhysicalAddr;
    CurrentGuestState->VmcsRegionVirtualAddress = VmcsRegion;
    return TRUE;
}*/
return true
}

func (v *vmxRegions)VmxAllocateVmmStack()(ok bool){//col:186
/*VmxAllocateVmmStack(_In_ INT ProcessorID)
{
    VIRTUAL_MACHINE_STATE * CurrentVmState = &g_GuestState[ProcessorID];
    CurrentVmState->VmmStack = ExAllocatePoolWithTag(NonPagedPool, VMM_STACK_SIZE, POOLTAG);
    if (CurrentVmState->VmmStack == NULL)
    {
        LogError("Err, insufficient memory in allocationg vmm stack");
        return FALSE;
    }
    RtlZeroMemory(CurrentVmState->VmmStack, VMM_STACK_SIZE);
    LogDebugInfo("VMM Stack for logical processor : 0x%llx", CurrentVmState->VmmStack);
    return TRUE;
}*/
return true
}

func (v *vmxRegions)VmxAllocateMsrBitmap()(ok bool){//col:217
/*VmxAllocateMsrBitmap(_In_ INT ProcessorID)
{
    VIRTUAL_MACHINE_STATE * CurrentVmState = &g_GuestState[ProcessorID];
    CurrentVmState->MsrBitmapVirtualAddress = ExAllocatePoolWithTag(NonPagedPool, PAGE_SIZE, POOLTAG);
    if (CurrentVmState->MsrBitmapVirtualAddress == NULL)
    {
        LogError("Err, insufficient memory in allocationg MSR Bitmaps");
        return FALSE;
    }
    RtlZeroMemory(CurrentVmState->MsrBitmapVirtualAddress, PAGE_SIZE);
    CurrentVmState->MsrBitmapPhysicalAddress = VirtualAddressToPhysicalAddress(CurrentVmState->MsrBitmapVirtualAddress);
    LogDebugInfo("MSR Bitmap virtual address  : 0x%llx", CurrentVmState->MsrBitmapVirtualAddress);
    LogDebugInfo("MSR Bitmap physical address : 0x%llx", CurrentVmState->MsrBitmapPhysicalAddress);
    return TRUE;
}*/
return true
}

func (v *vmxRegions)VmxAllocateIoBitmaps()(ok bool){//col:263
/*VmxAllocateIoBitmaps(_In_ INT ProcessorID)
{
    VIRTUAL_MACHINE_STATE * CurrentVmState = &g_GuestState[ProcessorID];
    if (CurrentVmState->IoBitmapVirtualAddressA == NULL)
    {
        LogError("Err, insufficient memory in allocationg I/O Bitmaps A");
        return FALSE;
    }
    RtlZeroMemory(CurrentVmState->IoBitmapVirtualAddressA, PAGE_SIZE);
    CurrentVmState->IoBitmapPhysicalAddressA = VirtualAddressToPhysicalAddress(CurrentVmState->IoBitmapVirtualAddressA);
    LogDebugInfo("I/O Bitmap A Virtual Address  : 0x%llx", CurrentVmState->IoBitmapVirtualAddressA);
    LogDebugInfo("I/O Bitmap A Physical Address : 0x%llx", CurrentVmState->IoBitmapPhysicalAddressA);
    if (CurrentVmState->IoBitmapVirtualAddressB == NULL)
    {
        LogError("Err, insufficient memory in allocationg I/O Bitmaps B");
        return FALSE;
    }
    RtlZeroMemory(CurrentVmState->IoBitmapVirtualAddressB, PAGE_SIZE);
    CurrentVmState->IoBitmapPhysicalAddressB = VirtualAddressToPhysicalAddress(CurrentVmState->IoBitmapVirtualAddressB);
    LogDebugInfo("I/O Bitmap B virtual address  : 0x%llx", CurrentVmState->IoBitmapVirtualAddressB);
    LogDebugInfo("I/O Bitmap B physical address : 0x%llx", CurrentVmState->IoBitmapPhysicalAddressB);
    return TRUE;
}*/
return true
}



