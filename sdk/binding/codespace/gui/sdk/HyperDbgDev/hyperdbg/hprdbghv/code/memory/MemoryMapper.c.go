package memory

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbghv\code\memory\MemoryMapper.c.back

type (
	MemoryMapper interface {
		MemoryMapperGetIndex() (ok bool)                             //col:6
		MemoryMapperGetOffset() (ok bool)                            //col:12
		MemoryMapperGetPteVa() (ok bool)                             //col:84
		MemoryMapperCheckIfPageIsNxBitSetByCr3() (ok bool)           //col:97
		MemoryMapperCheckIfPageIsNxBitSetOnTargetProcess() (ok bool) //col:441
		MemoryMapperFreeMemoryInTargetProcess() (ok bool)            //col:477
		MemoryMapperMapPhysicalAddressToPte() (ok bool)              //col:506
	}
	memoryMapper struct{}
)

func NewMemoryMapper() MemoryMapper { return &memoryMapper{} }

func (m *memoryMapper) MemoryMapperGetIndex() (ok bool) { //col:6
	/*
	   MemoryMapperGetIndex(PAGING_LEVEL Level, UINT64 Va)

	   	{
	   	    UINT64 Result = Va;
	   	    Result >>= 12 + Level * 9;
	   	    return Result;
	   	}
	*/
	return true
}

func (m *memoryMapper) MemoryMapperGetOffset() (ok bool) { //col:12
	/*
	   MemoryMapperGetOffset(PAGING_LEVEL Level, UINT64 Va)

	   	{
	   	    UINT64 Result = MemoryMapperGetIndex(Level, Va);
	   	    Result &= (1 << 9) - 1;
	   	    return Result;
	   	}
	*/
	return true
}

func (m *memoryMapper) MemoryMapperGetPteVa() (ok bool) { //col:84
	/*
	   MemoryMapperGetPteVa(PVOID Va, PAGING_LEVEL Level)

	   	{
	   	    CR3_TYPE Cr3;
	   	    Cr3.Flags = __readcr3();
	   	    return MemoryMapperGetPteVaWithoutSwitchingByCr3(Va, Level, Cr3);

	   MemoryMapperGetPteVaByCr3(PVOID Va, PAGING_LEVEL Level, CR3_TYPE TargetCr3)

	   	{
	   	    PPAGE_ENTRY PageEntry         = NULL;
	   	    CR3_TYPE    CurrentProcessCr3 = {0};
	   	    CurrentProcessCr3 = SwitchOnAnotherProcessMemoryLayoutByCr3(TargetCr3);
	   	    PageEntry = MemoryMapperGetPteVaWithoutSwitchingByCr3(Va, Level, TargetCr3);
	   	    RestoreToPreviousProcess(CurrentProcessCr3);

	   MemoryMapperGetPteVaWithoutSwitchingByCr3(PVOID Va, PAGING_LEVEL Level, CR3_TYPE TargetCr3)

	   	{
	   	    CR3_TYPE Cr3;
	   	    UINT64   TempCr3;
	   	    PUINT64  Cr3Va;
	   	    PUINT64  PdptVa;
	   	    PUINT64  PdVa;
	   	    PUINT64  PtVa;
	   	    UINT32   Offset;
	   	    Cr3.Flags = TargetCr3.Flags;
	   	    TempCr3 = Cr3.Fields.PageFrameNumber << 12;
	   	    Cr3Va = PhysicalAddressToVirtualAddress(TempCr3);
	   	    if (Cr3Va == NULL)
	   	    {
	   	        return NULL;
	   	    }
	   	    Offset = MemoryMapperGetOffset(PagingLevelPageMapLevel4, Va);
	   	    if (!Pml4e->Fields.Present || Level == PagingLevelPageMapLevel4)
	   	    {
	   	        return Pml4e;
	   	    }
	   	    PdptVa = PhysicalAddressToVirtualAddress(Pml4e->Fields.PageFrameNumber << 12);
	   	    if (PdptVa == NULL)
	   	    {
	   	        return NULL;
	   	    }
	   	    Offset = MemoryMapperGetOffset(PagingLevelPageDirectoryPointerTable, Va);
	   	    if (!Pdpte->Fields.Present || Pdpte->Fields.LargePage || Level == PagingLevelPageDirectoryPointerTable)
	   	    {
	   	        return Pdpte;
	   	    }
	   	    PdVa = PhysicalAddressToVirtualAddress(Pdpte->Fields.PageFrameNumber << 12);
	   	    if (PdVa == NULL)
	   	    {
	   	        return NULL;
	   	    }
	   	    Offset = MemoryMapperGetOffset(PagingLevelPageDirectory, Va);
	   	    if (!Pde->Fields.Present || Pde->Fields.LargePage || Level == PagingLevelPageDirectory)
	   	    {
	   	        return Pde;
	   	    }
	   	    PtVa = PhysicalAddressToVirtualAddress(Pde->Fields.PageFrameNumber << 12);
	   	    if (PtVa == NULL)
	   	    {
	   	        return NULL;
	   	    }
	   	    Offset = MemoryMapperGetOffset(PagingLevelPageTable, Va);

	   MemoryMapperCheckIfPageIsPresentByCr3(PVOID Va, CR3_TYPE TargetCr3)

	   	{
	   	    PPAGE_ENTRY PageEntry;
	   	    PageEntry = MemoryMapperGetPteVaByCr3(Va, PagingLevelPageTable, TargetCr3);
	   	    if (PageEntry != NULL && PageEntry->Fields.Present)
	   	    {
	   	        return TRUE;
	   	    }
	   	    else
	   	    {
	   	        return FALSE;
	   	    }
	   	}
	*/
	return true
}

func (m *memoryMapper) MemoryMapperCheckIfPageIsNxBitSetByCr3() (ok bool) { //col:97
	/*
	   MemoryMapperCheckIfPageIsNxBitSetByCr3(PVOID Va, CR3_TYPE TargetCr3)

	   	{
	   	    PPAGE_ENTRY PageEntry;
	   	    PageEntry = MemoryMapperGetPteVaByCr3(Va, PagingLevelPageTable, TargetCr3);
	   	    if (PageEntry != NULL && !PageEntry->Fields.ExecuteDisable)
	   	    {
	   	        return TRUE;
	   	    }
	   	    else
	   	    {
	   	        return FALSE;
	   	    }
	   	}
	*/
	return true
}

func (m *memoryMapper) MemoryMapperCheckIfPageIsNxBitSetOnTargetProcess() (ok bool) { //col:441
	/*
	   MemoryMapperCheckIfPageIsNxBitSetOnTargetProcess(PVOID Va)

	   	{
	   	    BOOLEAN     Result;
	   	    CR3_TYPE    GuestCr3;
	   	    PPAGE_ENTRY PageEntry;
	   	    CR3_TYPE    CurrentProcessCr3 = {0};
	   	    GuestCr3.Flags = GetRunningCr3OnTargetProcess().Flags;
	   	    CurrentProcessCr3 = SwitchOnAnotherProcessMemoryLayoutByCr3(GuestCr3);
	   	    PageEntry = MemoryMapperGetPteVa(Va, PagingLevelPageTable);
	   	    if (PageEntry != NULL && !PageEntry->Fields.ExecuteDisable)
	   	    {
	   	        Result = TRUE;
	   	    }
	   	    else
	   	    {
	   	        Result = FALSE;
	   	    }
	   	    RestoreToPreviousProcess(CurrentProcessCr3);

	   MemoryMapperMapReservedPageRange(SIZE_T Size)

	   	{
	   	    return MmAllocateMappingAddress(Size, POOLTAG);

	   MemoryMapperUnmapReservedPageRange(PVOID VirtualAddress)

	   	{
	   	    MmFreeMappingAddress(VirtualAddress, POOLTAG);

	   MemoryMapperGetPte(PVOID VirtualAddress)

	   	{
	   	    return MemoryMapperGetPteVa(VirtualAddress, PagingLevelPageTable);

	   MemoryMapperGetPteByCr3(PVOID VirtualAddress, CR3_TYPE TargetCr3)

	   	{
	   	    return MemoryMapperGetPteVaByCr3(VirtualAddress, PagingLevelPageTable, TargetCr3);

	   MemoryMapperMapPageAndGetPte(PUINT64 PteAddress)

	   	{
	   	    UINT64 Va;
	   	    UINT64 Pte;
	   	    Va = MemoryMapperMapReservedPageRange(PAGE_SIZE);
	   	    Pte = MemoryMapperGetPte(Va);

	   MemoryMapperInitialize()

	   	{
	   	    UINT64                  TempPte;
	   	    UINT32                  ProcessorCount = KeQueryActiveProcessorCount(0);
	   	    for (size_t i = 0; i < ProcessorCount; i++)
	   	    {
	   	        CurrentVmState = &g_GuestState[i];
	   	        CurrentVmState->MemoryMapper.VirualAddress     = MemoryMapperMapPageAndGetPte(&TempPte);

	   MemoryMapperUninitialize()

	   	{
	   	    UINT32                  ProcessorCount = KeQueryActiveProcessorCount(0);
	   	    for (size_t i = 0; i < ProcessorCount; i++)
	   	    {
	   	        CurrentVmState = &g_GuestState[i];
	   	        if (CurrentVmState->MemoryMapper.VirualAddress != NULL)
	   	            MemoryMapperUnmapReservedPageRange(CurrentVmState->MemoryMapper.VirualAddress);

	   MemoryMapperReadMemorySafeByPte(PHYSICAL_ADDRESS PaAddressToRead,

	   	PVOID            BufferToSaveMemory,
	   	SIZE_T           SizeToRead,
	   	UINT64           PteVaAddress,
	   	UINT64           MappingVa,
	   	BOOLEAN          InvalidateVpids)

	   	{
	   	    PVOID       Va = MappingVa;
	   	    PVOID       NewAddress;
	   	    PAGE_ENTRY  PageEntry;
	   	    PPAGE_ENTRY Pte = PteVaAddress;
	   	    PageEntry.Flags = Pte->Flags;
	   	    PageEntry.Fields.Present = 1;
	   	    PageEntry.Fields.Write = 1;
	   	    PageEntry.Fields.Global = 1;
	   	    PageEntry.Fields.PageFrameNumber = PaAddressToRead.QuadPart >> 12;
	   	    Pte->Flags = PageEntry.Flags;
	   	    __invlpg(Va);
	   	    if (InvalidateVpids)
	   	    {
	   	    }
	   	    NewAddress = (PVOID)((UINT64)Va + (PAGE_4KB_OFFSET & (PaAddressToRead.QuadPart)));
	   	    memcpy(BufferToSaveMemory, NewAddress, SizeToRead);

	   MemoryMapperWriteMemorySafeByPte(PVOID            SourceVA,

	   	PHYSICAL_ADDRESS PaAddressToWrite,
	   	SIZE_T           SizeToWrite,
	   	UINT64           PteVaAddress,
	   	UINT64           MappingVa,
	   	BOOLEAN          InvalidateVpids)

	   	{
	   	    PVOID       Va = MappingVa;
	   	    PVOID       NewAddress;
	   	    PAGE_ENTRY  PageEntry;
	   	    PPAGE_ENTRY Pte = PteVaAddress;
	   	    PageEntry.Flags = Pte->Flags;
	   	    PageEntry.Fields.Present = 1;
	   	    PageEntry.Fields.Write = 1;
	   	    PageEntry.Fields.Global = 1;
	   	    PageEntry.Fields.PageFrameNumber = PaAddressToWrite.QuadPart >> 12;
	   	    Pte->Flags = PageEntry.Flags;
	   	    __invlpg(Va);
	   	    if (InvalidateVpids)
	   	    {
	   	    }
	   	    NewAddress = (PVOID)((UINT64)Va + (PAGE_4KB_OFFSET & (PaAddressToWrite.QuadPart)));
	   	    memcpy(NewAddress, SourceVA, SizeToWrite);

	   MemoryMapperReadMemorySafeByPhysicalAddressWrapperAddressMaker(

	   	MEMORY_MAPPER_WRAPPER_FOR_MEMORY_WRITE TypeOfRead,
	   	UINT64                                 AddressToRead)

	   	{
	   	    PHYSICAL_ADDRESS PhysicalAddress = {0};
	   	    switch (TypeOfRead)
	   	    {
	   	    case MEMORY_MAPPER_WRAPPER_READ_PHYSICAL_MEMORY:
	   	        PhysicalAddress.QuadPart = AddressToRead;
	   	        break;
	   	    case MEMORY_MAPPER_WRAPPER_READ_VIRTUAL_MEMORY:
	   	        PhysicalAddress.QuadPart = VirtualAddressToPhysicalAddress(AddressToRead);

	   MemoryMapperReadMemorySafeByPhysicalAddressWrapper(

	   	MEMORY_MAPPER_WRAPPER_FOR_MEMORY_WRITE TypeOfRead,
	   	UINT64                                 AddressToRead,
	   	UINT64                                 BufferToSaveMemory,
	   	SIZE_T                                 SizeToRead)

	   	{
	   	    ULONG                   ProcessorIndex = KeGetCurrentProcessorNumber();
	   	    if (CurrentVmState->MemoryMapper.VirualAddress == NULL ||
	   	        CurrentVmState->MemoryMapper.PteVirtualAddress == NULL)
	   	    {
	   	        return FALSE;
	   	    }
	   	    AddressToCheck = (CHAR *)AddressToRead + SizeToRead - ((CHAR *)PAGE_ALIGN(AddressToRead));
	   	    if (AddressToCheck > PAGE_SIZE)
	   	    {
	   	        UINT64 ReadSize = AddressToCheck;
	   	        while (SizeToRead != 0)
	   	        {
	   	            ReadSize = (UINT64)PAGE_ALIGN(AddressToRead + PAGE_SIZE) - AddressToRead;
	   	            if (ReadSize == PAGE_SIZE && SizeToRead < PAGE_SIZE)
	   	            {
	   	                ReadSize = SizeToRead;
	   	            }
	   	            PhysicalAddress.QuadPart = MemoryMapperReadMemorySafeByPhysicalAddressWrapperAddressMaker(TypeOfRead,
	   	                                                                                                      AddressToRead);
	   	            if (!MemoryMapperReadMemorySafeByPte(
	   	                    PhysicalAddress,
	   	                    BufferToSaveMemory,
	   	                    ReadSize,
	   	                    CurrentVmState->MemoryMapper.PteVirtualAddress,
	   	                    CurrentVmState->MemoryMapper.VirualAddress,
	   	                    CurrentVmState->IsOnVmxRootMode))
	   	            {
	   	                return FALSE;
	   	            }
	   	            SizeToRead         = SizeToRead - ReadSize;
	   	            AddressToRead      = AddressToRead + ReadSize;
	   	            BufferToSaveMemory = BufferToSaveMemory + ReadSize;
	   	        }
	   	        return TRUE;
	   	    }
	   	    else
	   	    {
	   	        PhysicalAddress.QuadPart = MemoryMapperReadMemorySafeByPhysicalAddressWrapperAddressMaker(TypeOfRead,
	   	                                                                                                  AddressToRead);
	   	        return MemoryMapperReadMemorySafeByPte(
	   	            PhysicalAddress,
	   	            BufferToSaveMemory,
	   	            SizeToRead,
	   	            CurrentVmState->MemoryMapper.PteVirtualAddress,
	   	            CurrentVmState->MemoryMapper.VirualAddress,
	   	            CurrentVmState->IsOnVmxRootMode);

	   MemoryMapperReadMemorySafeByPhysicalAddress(UINT64 PaAddressToRead,

	   	UINT64 BufferToSaveMemory,
	   	SIZE_T SizeToRead)

	   	{
	   	    return MemoryMapperReadMemorySafeByPhysicalAddressWrapper(MEMORY_MAPPER_WRAPPER_READ_PHYSICAL_MEMORY,
	   	                                                              PaAddressToRead,
	   	                                                              BufferToSaveMemory,
	   	                                                              SizeToRead);

	   MemoryMapperReadMemorySafe(UINT64 VaAddressToRead, PVOID BufferToSaveMemory, SIZE_T SizeToRead)

	   	{
	   	    return MemoryMapperReadMemorySafeByPhysicalAddressWrapper(MEMORY_MAPPER_WRAPPER_READ_VIRTUAL_MEMORY,
	   	                                                              VaAddressToRead,
	   	                                                              BufferToSaveMemory,
	   	                                                              SizeToRead);

	   MemoryMapperReadMemorySafeOnTargetProcess(UINT64 VaAddressToRead, PVOID BufferToSaveMemory, SIZE_T SizeToRead)

	   	{
	   	    CR3_TYPE GuestCr3;
	   	    CR3_TYPE OriginalCr3;
	   	    BOOLEAN  Result;
	   	    GuestCr3.Flags = GetRunningCr3OnTargetProcess().Flags;
	   	    OriginalCr3.Flags = __readcr3();
	   	    __writecr3(GuestCr3.Flags);
	   	    Result = MemoryMapperReadMemorySafe(VaAddressToRead, BufferToSaveMemory, SizeToRead);
	   	    __writecr3(OriginalCr3.Flags);

	   MemoryMapperWriteMemorySafeOnTargetProcess(UINT64 Destination, PVOID Source, SIZE_T Size)

	   	{
	   	    CR3_TYPE GuestCr3;
	   	    CR3_TYPE OriginalCr3;
	   	    BOOLEAN  Result;
	   	    GuestCr3.Flags = GetRunningCr3OnTargetProcess().Flags;
	   	    OriginalCr3.Flags = __readcr3();
	   	    __writecr3(GuestCr3.Flags);
	   	    Result = MemoryMapperWriteMemorySafe(Destination, Source, Size, GuestCr3);
	   	    __writecr3(OriginalCr3.Flags);

	   MemoryMapperWriteMemorySafeWrapperAddressMaker(MEMORY_MAPPER_WRAPPER_FOR_MEMORY_WRITE TypeOfWrite,

	   	UINT64                                 DestinationAddr,
	   	PCR3_TYPE                              TargetProcessCr3,
	   	UINT32                                 TargetProcessId)

	   	{
	   	    PHYSICAL_ADDRESS PhysicalAddress = {0};
	   	    switch (TypeOfWrite)
	   	    {
	   	    case MEMORY_MAPPER_WRAPPER_WRITE_PHYSICAL_MEMORY:
	   	        PhysicalAddress.QuadPart = DestinationAddr;
	   	        break;
	   	    case MEMORY_MAPPER_WRAPPER_WRITE_VIRTUAL_MEMORY_UNSAFE:
	   	        if (TargetProcessId == NULL)
	   	        {
	   	            PhysicalAddress.QuadPart = VirtualAddressToPhysicalAddress(DestinationAddr);
	   	            PhysicalAddress.QuadPart = VirtualAddressToPhysicalAddressByProcessId(DestinationAddr, TargetProcessId);
	   	        if (TargetProcessCr3 == NULL || TargetProcessCr3->Flags == NULL)
	   	        {
	   	            PhysicalAddress.QuadPart = VirtualAddressToPhysicalAddress(DestinationAddr);
	   	            PhysicalAddress.QuadPart = VirtualAddressToPhysicalAddressByProcessCr3(DestinationAddr, *TargetProcessCr3);

	   MemoryMapperWriteMemorySafeWrapper(MEMORY_MAPPER_WRAPPER_FOR_MEMORY_WRITE TypeOfWrite,

	   	UINT64                                 DestinationAddr,
	   	UINT64                                 Source,
	   	SIZE_T                                 SizeToWrite,
	   	PCR3_TYPE                              TargetProcessCr3,
	   	UINT32                                 TargetProcessId)

	   	{
	   	    ULONG                   ProcessorIndex = KeGetCurrentProcessorNumber();
	   	    if (CurrentVmState->MemoryMapper.VirualAddress == NULL ||
	   	        CurrentVmState->MemoryMapper.PteVirtualAddress == NULL)
	   	    {
	   	        return FALSE;
	   	    }
	   	    AddressToCheck = (CHAR *)DestinationAddr + SizeToWrite - ((CHAR *)PAGE_ALIGN(DestinationAddr));
	   	    if (AddressToCheck > PAGE_SIZE)
	   	    {
	   	        UINT64 WriteSize = AddressToCheck;
	   	        while (SizeToWrite != 0)
	   	        {
	   	            WriteSize = (UINT64)PAGE_ALIGN(DestinationAddr + PAGE_SIZE) - DestinationAddr;
	   	            if (WriteSize == PAGE_SIZE && SizeToWrite < PAGE_SIZE)
	   	            {
	   	                WriteSize = SizeToWrite;
	   	            }
	   	            PhysicalAddress.QuadPart = MemoryMapperWriteMemorySafeWrapperAddressMaker(TypeOfWrite,
	   	                                                                                      DestinationAddr,
	   	                                                                                      TargetProcessCr3,
	   	                                                                                      TargetProcessId);
	   	            if (!MemoryMapperWriteMemorySafeByPte(
	   	                    Source,
	   	                    PhysicalAddress,
	   	                    WriteSize,
	   	                    CurrentVmState->MemoryMapper.PteVirtualAddress,
	   	                    CurrentVmState->MemoryMapper.VirualAddress,
	   	                    CurrentVmState->IsOnVmxRootMode))
	   	            {
	   	                return FALSE;
	   	            }
	   	            SizeToWrite     = SizeToWrite - WriteSize;
	   	            DestinationAddr = DestinationAddr + WriteSize;
	   	            Source          = Source + WriteSize;
	   	        }
	   	        return TRUE;
	   	    }
	   	    else
	   	    {
	   	        PhysicalAddress.QuadPart = MemoryMapperWriteMemorySafeWrapperAddressMaker(TypeOfWrite,
	   	                                                                                  DestinationAddr,
	   	                                                                                  TargetProcessCr3,
	   	                                                                                  TargetProcessId);
	   	        return MemoryMapperWriteMemorySafeByPte(
	   	            Source,
	   	            PhysicalAddress,
	   	            SizeToWrite,
	   	            CurrentVmState->MemoryMapper.PteVirtualAddress,
	   	            CurrentVmState->MemoryMapper.VirualAddress,
	   	            CurrentVmState->IsOnVmxRootMode);

	   MemoryMapperWriteMemorySafe(UINT64   Destination,

	   	PVOID    Source,
	   	SIZE_T   SizeToWrite,
	   	CR3_TYPE TargetProcessCr3)

	   	{
	   	    return MemoryMapperWriteMemorySafeWrapper(MEMORY_MAPPER_WRAPPER_WRITE_VIRTUAL_MEMORY_SAFE,
	   	                                              Destination,
	   	                                              Source,
	   	                                              SizeToWrite,
	   	                                              &TargetProcessCr3,
	   	                                              NULL);

	   MemoryMapperWriteMemoryUnsafe(UINT64 Destination, PVOID Source, SIZE_T SizeToWrite, UINT32 TargetProcessId)

	   	{
	   	    return MemoryMapperWriteMemorySafeWrapper(MEMORY_MAPPER_WRAPPER_WRITE_VIRTUAL_MEMORY_UNSAFE,
	   	                                              Destination,
	   	                                              Source,
	   	                                              SizeToWrite,
	   	                                              NULL,
	   	                                              TargetProcessId);

	   MemoryMapperWriteMemorySafeByPhysicalAddress(UINT64 DestinationPa,

	   	UINT64 Source,
	   	SIZE_T SizeToWrite)

	   	{
	   	    return MemoryMapperWriteMemorySafeWrapper(MEMORY_MAPPER_WRAPPER_WRITE_PHYSICAL_MEMORY,
	   	                                              DestinationPa,
	   	                                              Source,
	   	                                              SizeToWrite,
	   	                                              NULL,
	   	                                              NULL);

	   MemoryMapperReserveUsermodeAddressInTargetProcess(UINT32 ProcessId, BOOLEAN Allocate)

	   	{
	   	    NTSTATUS   Status;
	   	    PVOID      AllocPtr  = NULL;
	   	    SIZE_T     AllocSize = PAGE_SIZE;
	   	    PEPROCESS  SourceProcess;
	   	    KAPC_STATE State = {0};
	   	    if (PsGetCurrentProcessId() != ProcessId)
	   	    {
	   	        if (PsLookupProcessByProcessId(ProcessId, &SourceProcess) != STATUS_SUCCESS)
	   	        {
	   	            return NULL;
	   	        }
	   	        __try
	   	        {
	   	            KeStackAttachProcess(SourceProcess, &State);
	   	            Status = ZwAllocateVirtualMemory(
	   	                NtCurrentProcess(),
	   	                &AllocPtr,
	   	                NULL,
	   	                &AllocSize,
	   	                Allocate ? MEM_COMMIT : MEM_RESERVE,
	   	                PAGE_EXECUTE_READWRITE);
	   	            KeUnstackDetachProcess(&State);
	   	            ObDereferenceObject(SourceProcess);
	   	        __except (EXCEPTION_EXECUTE_HANDLER)
	   	        {
	   	            KeUnstackDetachProcess(&State);
	   	            ObDereferenceObject(SourceProcess);
	   	        Status = ZwAllocateVirtualMemory(
	   	            NtCurrentProcess(),
	   	            &AllocPtr,
	   	            NULL,
	   	            &AllocSize,
	   	            Allocate ? MEM_COMMIT : MEM_RESERVE,
	   	            PAGE_EXECUTE_READWRITE);
	   	    if (!NT_SUCCESS(Status))
	   	    {
	   	        return NULL;
	   	    }
	   	    return AllocPtr;
	   	}
	*/
	return true
}

func (m *memoryMapper) MemoryMapperFreeMemoryInTargetProcess() (ok bool) { //col:477
	/*
	   MemoryMapperFreeMemoryInTargetProcess(UINT32 ProcessId,

	   	PVOID  BaseAddress)

	   	{
	   	    NTSTATUS   Status;
	   	    SIZE_T     AllocSize = PAGE_SIZE;
	   	    PEPROCESS  SourceProcess;
	   	    KAPC_STATE State = {0};
	   	    if (PsGetCurrentProcessId() != ProcessId)
	   	    {
	   	        if (PsLookupProcessByProcessId(ProcessId, &SourceProcess) != STATUS_SUCCESS)
	   	        {
	   	            return FALSE;
	   	        }
	   	        __try
	   	        {
	   	            KeStackAttachProcess(SourceProcess, &State);
	   	            Status = ZwFreeVirtualMemory(NtCurrentProcess(),
	   	                                         &BaseAddress,
	   	                                         &AllocSize,
	   	                                         MEM_RELEASE);
	   	            KeUnstackDetachProcess(&State);
	   	            ObDereferenceObject(SourceProcess);
	   	        __except (EXCEPTION_EXECUTE_HANDLER)
	   	        {
	   	            KeUnstackDetachProcess(&State);
	   	            ObDereferenceObject(SourceProcess);
	   	        Status = ZwFreeVirtualMemory(NtCurrentProcess(),
	   	                                     &BaseAddress,
	   	                                     &AllocSize,
	   	                                     MEM_RELEASE);
	   	    if (!NT_SUCCESS(Status))
	   	    {
	   	        return FALSE;
	   	    }
	   	    return TRUE;
	   	}
	*/
	return true
}

func (m *memoryMapper) MemoryMapperMapPhysicalAddressToPte() (ok bool) { //col:506
	/*
	   MemoryMapperMapPhysicalAddressToPte(PHYSICAL_ADDRESS PhysicalAddress,

	   	PVOID            TargetProcessVirtualAddress,
	   	CR3_TYPE         TargetProcessKernelCr3)

	   	{
	   	    PPAGE_ENTRY PreviousPteEntry;
	   	    PAGE_ENTRY  PageEntry;
	   	    CR3_TYPE    CurrentProcessCr3;
	   	    PreviousPteEntry = MemoryMapperGetPteVaByCr3(TargetProcessVirtualAddress, PagingLevelPageTable, TargetProcessKernelCr3);
	   	    CurrentProcessCr3 = SwitchOnAnotherProcessMemoryLayoutByCr3(TargetProcessKernelCr3);
	   	    __invlpg(TargetProcessVirtualAddress);
	   	    RestoreToPreviousProcess(CurrentProcessCr3);

	   MemoryMapperSetSupervisorBitWithoutSwitchingByCr3(PVOID Va, BOOLEAN Set, PAGING_LEVEL Level, CR3_TYPE TargetCr3)

	   	{
	   	    PPAGE_ENTRY Pml = NULL;
	   	    Pml = MemoryMapperGetPteVaWithoutSwitchingByCr3(Va, Level, TargetCr3);
	   	    if (!Pml)
	   	    {
	   	        return FALSE;
	   	    }
	   	    if (Set)
	   	    {
	   	        Pml->Fields.Supervisor = 1;
	   	    }
	   	    else
	   	    {
	   	        Pml->Fields.Supervisor = 0;
	   	    }
	   	    return TRUE;
	   	}
	*/
	return true
}

