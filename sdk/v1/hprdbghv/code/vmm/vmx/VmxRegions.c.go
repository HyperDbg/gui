package vmx

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbghv\code\vmm\vmx\VmxRegions.c.back

type (
	VmxRegions interface {
		VmxAllocateVmxonRegion() (ok bool) //col:55
	}
	vmxRegions struct{}
)

func NewVmxRegions() VmxRegions { return &vmxRegions{} }

func (v *vmxRegions) VmxAllocateVmxonRegion() (ok bool) { //col:55
	/*
	   VmxAllocateVmxonRegion(VIRTUAL_MACHINE_STATE * CurrentGuestState)

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
	   	VmxonRegion = CrsAllocateContiguousZeroedMemory(VmxonSize + ALIGNMENT_PAGE_SIZE);
	   	if (VmxonRegion == NULL)
	   	{
	   	    LogError("Err, couldn't allocate buffer for VMXON region");
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

	   VmxAllocateVmcsRegion(VIRTUAL_MACHINE_STATE * CurrentGuestState)

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
	   	    VmcsRegion = CrsAllocateContiguousZeroedMemory(VmcsSize + ALIGNMENT_PAGE_SIZE);
	   	    if (VmcsRegion == NULL)
	   	    {
	   	        LogError("Err, couldn't allocate Buffer for VMCS region");
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
	   	}
	*/
	return true
}

