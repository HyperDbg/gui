package ept

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbghv\code\vmm\ept\Ept.c.back

type (
	Ept interface {
		EptCheckFeatures() (ok bool)                      //col:74
		EptGetPml2Entry() (ok bool)                       //col:88
		EptSplitLargePage() (ok bool)                     //col:152
		EptAllocateAndCreateIdentityPageTable() (ok bool) //col:200
	}
	ept struct{}
)

func NewEpt() Ept { return &ept{} }

func (e *ept) EptCheckFeatures() (ok bool) { //col:74
	/*
	   EptCheckFeatures()

	   	{
	   	    IA32_VMX_EPT_VPID_CAP_REGISTER VpidRegister;
	   	    IA32_MTRR_DEF_TYPE_REGISTER    MTRRDefType;
	   	    VpidRegister.AsUInt = __readmsr(IA32_VMX_EPT_VPID_CAP);
	   	    MTRRDefType.AsUInt  = __readmsr(IA32_MTRR_DEF_TYPE);
	   	    if (!VpidRegister.PageWalkLength4 || !VpidRegister.MemoryTypeWriteBack || !VpidRegister.Pde2MbPages)
	   	    {
	   	        return FALSE;
	   	    }
	   	    if (!VpidRegister.AdvancedVmexitEptViolationsInformation)
	   	    {
	   	        LogDebugInfo("The processor doesn't report advanced VM-exit information for EPT violations");
	   	    if (!VpidRegister.ExecuteOnlyPages)
	   	    {
	   	        g_ExecuteOnlySupport = FALSE;
	   	        LogDebugInfo("The processor doesn't support execute-only pages, execute hooks won't work as they're on this feature in our design");
	   	    if (!MTRRDefType.MtrrEnable)
	   	    {
	   	        LogError("Err, MTRR dynamic ranges are not supported");
	   	    LogDebugInfo("All EPT features are present");

	   EptBuildMtrrMap()

	   	{
	   	    IA32_MTRR_CAPABILITIES_REGISTER MTRRCap;
	   	    IA32_MTRR_PHYSBASE_REGISTER     CurrentPhysBase;
	   	    IA32_MTRR_PHYSMASK_REGISTER     CurrentPhysMask;
	   	    PMTRR_RANGE_DESCRIPTOR          Descriptor;
	   	    ULONG                           CurrentRegister;
	   	    ULONG                           NumberOfBitsInMask;
	   	    MTRRCap.AsUInt = __readmsr(IA32_MTRR_CAPABILITIES);
	   	    for (CurrentRegister = 0; CurrentRegister < MTRRCap.VariableRangeCount; CurrentRegister++)
	   	    {
	   	        CurrentPhysBase.AsUInt = __readmsr(IA32_MTRR_PHYSBASE0 + (CurrentRegister * 2));
	   	        CurrentPhysMask.AsUInt = __readmsr(IA32_MTRR_PHYSMASK0 + (CurrentRegister * 2));
	   	        if (CurrentPhysMask.Valid)
	   	        {
	   	            Descriptor = &g_EptState->MemoryRanges[g_EptState->NumberOfEnabledMemoryRanges++];
	   	            Descriptor->PhysicalBaseAddress = CurrentPhysBase.PageFrameNumber * PAGE_SIZE;
	   	            _BitScanForward64(&NumberOfBitsInMask, CurrentPhysMask.PageFrameNumber * PAGE_SIZE);
	   	            Descriptor->PhysicalEndAddress = Descriptor->PhysicalBaseAddress + ((1ULL << NumberOfBitsInMask) - 1ULL);
	   	            Descriptor->MemoryType = (UCHAR)CurrentPhysBase.Type;
	   	            if (Descriptor->MemoryType == MEMORY_TYPE_WRITE_BACK)
	   	            {
	   	                g_EptState->NumberOfEnabledMemoryRanges--;
	   	            }
	   	            LogDebugInfo("MTRR Range: Base=0x%llx End=0x%llx Type=0x%x", Descriptor->PhysicalBaseAddress, Descriptor->PhysicalEndAddress, Descriptor->MemoryType);
	   	    LogDebugInfo("Total MTRR ranges committed: 0x%x", g_EptState->NumberOfEnabledMemoryRanges);

	   EptGetPml1Entry(PVMM_EPT_PAGE_TABLE EptPageTable, SIZE_T PhysicalAddress)

	   	{
	   	    SIZE_T            Directory, DirectoryPointer, PML4Entry;
	   	    PEPT_PML2_ENTRY   PML2;
	   	    PEPT_PML1_ENTRY   PML1;
	   	    PEPT_PML2_POINTER PML2Pointer;
	   	    Directory        = ADDRMASK_EPT_PML2_INDEX(PhysicalAddress);
	   	    DirectoryPointer = ADDRMASK_EPT_PML3_INDEX(PhysicalAddress);
	   	    PML4Entry        = ADDRMASK_EPT_PML4_INDEX(PhysicalAddress);
	   	    if (PML4Entry > 0)
	   	    {
	   	        return NULL;
	   	    }
	   	    PML2 = &EptPageTable->PML2[DirectoryPointer][Directory];
	   	    if (PML2->LargePage)
	   	    {
	   	        return NULL;
	   	    }
	   	    PML2Pointer = (PEPT_PML2_POINTER)PML2;
	   	    PML1 = (PEPT_PML1_ENTRY)PhysicalAddressToVirtualAddress((PVOID)(PML2Pointer->PageFrameNumber * PAGE_SIZE));
	   	    if (!PML1)
	   	    {
	   	        return NULL;
	   	    }
	   	    PML1 = &PML1[ADDRMASK_EPT_PML1_INDEX(PhysicalAddress)];
	   	    return PML1;
	   	}
	*/
	return true
}

func (e *ept) EptGetPml2Entry() (ok bool) { //col:88
	/*
	   EptGetPml2Entry(PVMM_EPT_PAGE_TABLE EptPageTable, SIZE_T PhysicalAddress)

	   	{
	   	    SIZE_T          Directory, DirectoryPointer, PML4Entry;
	   	    PEPT_PML2_ENTRY PML2;
	   	    Directory        = ADDRMASK_EPT_PML2_INDEX(PhysicalAddress);
	   	    DirectoryPointer = ADDRMASK_EPT_PML3_INDEX(PhysicalAddress);
	   	    PML4Entry        = ADDRMASK_EPT_PML4_INDEX(PhysicalAddress);
	   	    if (PML4Entry > 0)
	   	    {
	   	        return NULL;
	   	    }
	   	    PML2 = &EptPageTable->PML2[DirectoryPointer][Directory];
	   	    return PML2;
	   	}
	*/
	return true
}

func (e *ept) EptSplitLargePage() (ok bool) { //col:152
	/*
	   EptSplitLargePage(PVMM_EPT_PAGE_TABLE EptPageTable,

	   	PVOID               PreAllocatedBuffer,
	   	SIZE_T              PhysicalAddress,
	   	ULONG               CoreIndex)

	   	{
	   	    PVMM_EPT_DYNAMIC_SPLIT NewSplit;
	   	    EPT_PML1_ENTRY         EntryTemplate;
	   	    SIZE_T                 EntryIndex;
	   	    PEPT_PML2_ENTRY        TargetEntry;
	   	    EPT_PML2_POINTER       NewPointer;
	   	    TargetEntry = EptGetPml2Entry(EptPageTable, PhysicalAddress);
	   	    if (!TargetEntry)
	   	    {
	   	        LogError("Err, an invalid physical address passed");
	   	    if (!TargetEntry->LargePage)
	   	    {
	   	        PoolManagerFreePool(PreAllocatedBuffer);
	   	    NewSplit = (PVMM_EPT_DYNAMIC_SPLIT)PreAllocatedBuffer;
	   	    if (!NewSplit)
	   	    {
	   	        LogError("Err, failed to allocate dynamic split memory");
	   	    RtlZeroMemory(NewSplit, sizeof(VMM_EPT_DYNAMIC_SPLIT));
	   	    __stosq((SIZE_T *)&NewSplit->PML1[0], EntryTemplate.AsUInt, VMM_EPT_PML1E_COUNT);
	   	    for (EntryIndex = 0; EntryIndex < VMM_EPT_PML1E_COUNT; EntryIndex++)
	   	    {
	   	        NewSplit->PML1[EntryIndex].PageFrameNumber = ((TargetEntry->PageFrameNumber * SIZE_2_MB) / PAGE_SIZE) + EntryIndex;
	   	    }
	   	    NewPointer.AsUInt          = 0;
	   	    NewPointer.WriteAccess     = 1;
	   	    NewPointer.ReadAccess      = 1;
	   	    NewPointer.ExecuteAccess   = 1;
	   	    NewPointer.PageFrameNumber = (SIZE_T)VirtualAddressToPhysicalAddress(&NewSplit->PML1[0]) / PAGE_SIZE;
	   	    RtlCopyMemory(TargetEntry, &NewPointer, sizeof(NewPointer));

	   EptSetupPML2Entry(PEPT_PML2_ENTRY NewEntry, SIZE_T PageFrameNumber)

	   	{
	   	    SIZE_T                  AddressOfPage;
	   	    SIZE_T                  CurrentMtrrRange;
	   	    SIZE_T                  TargetMemoryType;
	   	    MTRR_RANGE_DESCRIPTOR * CurrentMemoryRange = NULL;
	   	    NewEntry->PageFrameNumber = PageFrameNumber;
	   	    AddressOfPage = PageFrameNumber * SIZE_2_MB;
	   	    if (PageFrameNumber == 0)
	   	    {
	   	        NewEntry->MemoryType = MEMORY_TYPE_UNCACHEABLE;
	   	        return;
	   	    }
	   	    TargetMemoryType = MEMORY_TYPE_WRITE_BACK;
	   	    for (CurrentMtrrRange = 0; CurrentMtrrRange < g_EptState->NumberOfEnabledMemoryRanges; CurrentMtrrRange++)
	   	    {
	   	        CurrentMemoryRange = &g_EptState->MemoryRanges[CurrentMtrrRange];
	   	        if (AddressOfPage <= CurrentMemoryRange->PhysicalEndAddress)
	   	        {
	   	            if ((AddressOfPage + SIZE_2_MB - 1) >= CurrentMemoryRange->PhysicalBaseAddress)
	   	            {
	   	                TargetMemoryType = CurrentMemoryRange->MemoryType;
	   	                if (TargetMemoryType == MEMORY_TYPE_UNCACHEABLE)
	   	                {
	   	                    break;
	   	                }
	   	            }
	   	        }
	   	    }
	   	    NewEntry->MemoryType = TargetMemoryType;
	   	}
	*/
	return true
}

func (e *ept) EptAllocateAndCreateIdentityPageTable() (ok bool) { //col:200
	/*
	   EptAllocateAndCreateIdentityPageTable()

	   	{
	   	    PVMM_EPT_PAGE_TABLE PageTable;
	   	    EPT_PML3_POINTER    RWXTemplate;
	   	    EPT_PML2_ENTRY      PML2EntryTemplate;
	   	    SIZE_T              EntryGroupIndex;
	   	    SIZE_T              EntryIndex;
	   	    PVOID Output;
	   	    PageTable = CrsAllocateContiguousZeroedMemory(sizeof(VMM_EPT_PAGE_TABLE));
	   	    if (PageTable == NULL)
	   	    {
	   	        LogError("Err, failed to allocate memory for PageTable");
	   	    PageTable->PML4[0].PageFrameNumber = (SIZE_T)VirtualAddressToPhysicalAddress(&PageTable->PML3[0]) / PAGE_SIZE;
	   	    PageTable->PML4[0].ReadAccess      = 1;
	   	    PageTable->PML4[0].WriteAccess     = 1;
	   	    PageTable->PML4[0].ExecuteAccess   = 1;
	   	    RWXTemplate.AsUInt = 0;
	   	    RWXTemplate.ReadAccess    = 1;
	   	    RWXTemplate.WriteAccess   = 1;
	   	    RWXTemplate.ExecuteAccess = 1;
	   	    __stosq((SIZE_T *)&PageTable->PML3[0], RWXTemplate.AsUInt, VMM_EPT_PML3E_COUNT);
	   	    for (EntryIndex = 0; EntryIndex < VMM_EPT_PML3E_COUNT; EntryIndex++)
	   	    {
	   	        PageTable->PML3[EntryIndex].PageFrameNumber = (SIZE_T)VirtualAddressToPhysicalAddress(&PageTable->PML2[EntryIndex][0]) / PAGE_SIZE;
	   	    }
	   	    PML2EntryTemplate.AsUInt = 0;
	   	    PML2EntryTemplate.WriteAccess   = 1;
	   	    PML2EntryTemplate.ReadAccess    = 1;
	   	    PML2EntryTemplate.ExecuteAccess = 1;
	   	    PML2EntryTemplate.LargePage = 1;
	   	    __stosq((SIZE_T *)&PageTable->PML2[0], PML2EntryTemplate.AsUInt, VMM_EPT_PML3E_COUNT * VMM_EPT_PML2E_COUNT);
	   	    for (EntryGroupIndex = 0; EntryGroupIndex < VMM_EPT_PML3E_COUNT; EntryGroupIndex++)
	   	    {
	   	        for (EntryIndex = 0; EntryIndex < VMM_EPT_PML2E_COUNT; EntryIndex++)
	   	        {
	   	            EptSetupPML2Entry(&PageTable->PML2[EntryGroupIndex][EntryIndex], (EntryGroupIndex * VMM_EPT_PML2E_COUNT) + EntryIndex);

	   EptLogicalProcessorInitialize()

	   	{
	   	    PVMM_EPT_PAGE_TABLE PageTable;
	   	    EPT_POINTER         EPTP = {0};
	   	    PageTable = EptAllocateAndCreateIdentityPageTable();
	   	    if (!PageTable)
	   	    {
	   	        LogError("Err, unable to allocate memory for EPT");
	   	    EPTP.PageFrameNumber = (SIZE_T)VirtualAddressToPhysicalAddress(&PageTable->PML4) / PAGE_SIZE;
	   	    g_EptState->EptPointer = EPTP;
	   	    return TRUE;
	   	}
	*/
	return true
}

