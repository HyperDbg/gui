#include "pch.h"
_Use_decl_annotations_
UINT64
MemoryMapperGetIndex(PAGING_LEVEL Level, UINT64 Va) {
    UINT64 Result = Va;
    Result >>= 12 + Level * 9;
    return Result;
}

_Use_decl_annotations_
UINT32
MemoryMapperGetOffset(PAGING_LEVEL Level, UINT64 Va) {
    UINT64 Result = MemoryMapperGetIndex(Level, Va);
    Result &= (1 << 9) - 1; // 0x1ff
    return Result;
}

_Use_decl_annotations_
PPAGE_ENTRY
MemoryMapperGetPteVa(PVOID Va, PAGING_LEVEL Level) {
    CR3_TYPE Cr3;
    Cr3.Flags = __readcr3();
    return MemoryMapperGetPteVaWithoutSwitchingByCr3(Va, Level, Cr3);
}

_Use_decl_annotations_
PPAGE_ENTRY
MemoryMapperGetPteVaByCr3(PVOID Va, PAGING_LEVEL Level, CR3_TYPE TargetCr3) {
    PPAGE_ENTRY PageEntry         = NULL;
    CR3_TYPE    CurrentProcessCr3 = {0};
    CurrentProcessCr3             = SwitchOnAnotherProcessMemoryLayoutByCr3(TargetCr3);
    PageEntry                     = MemoryMapperGetPteVaWithoutSwitchingByCr3(Va, Level, TargetCr3);
    RestoreToPreviousProcess(CurrentProcessCr3);
    return PageEntry;
}

_Use_decl_annotations_
PPAGE_ENTRY
MemoryMapperGetPteVaWithoutSwitchingByCr3(PVOID Va, PAGING_LEVEL Level, CR3_TYPE TargetCr3) {
    CR3_TYPE Cr3;
    UINT64   TempCr3;
    PUINT64  Cr3Va;
    PUINT64  PdptVa;
    PUINT64  PdVa;
    PUINT64  PtVa;
    UINT32   Offset;
    Cr3.Flags = TargetCr3.Flags;
    TempCr3   = Cr3.Fields.PageFrameNumber << 12;
    Cr3Va     = PhysicalAddressToVirtualAddress(TempCr3);
    if (Cr3Va == NULL) {
        return NULL;
    }
    Offset            = MemoryMapperGetOffset(PagingLevelPageMapLevel4, Va);
    PPAGE_ENTRY Pml4e = &Cr3Va[Offset];
    if (!Pml4e->Fields.Present || Level == PagingLevelPageMapLevel4) {
        return Pml4e;
    }
    PdptVa = PhysicalAddressToVirtualAddress(Pml4e->Fields.PageFrameNumber << 12);
    if (PdptVa == NULL) {
        return NULL;
    }
    Offset            = MemoryMapperGetOffset(PagingLevelPageDirectoryPointerTable, Va);
    PPAGE_ENTRY Pdpte = &PdptVa[Offset];
    if (!Pdpte->Fields.Present || Pdpte->Fields.LargePage || Level == PagingLevelPageDirectoryPointerTable) {
        return Pdpte;
    }
    PdVa = PhysicalAddressToVirtualAddress(Pdpte->Fields.PageFrameNumber << 12);
    if (PdVa == NULL) {
        return NULL;
    }
    Offset          = MemoryMapperGetOffset(PagingLevelPageDirectory, Va);
    PPAGE_ENTRY Pde = &PdVa[Offset];
    if (!Pde->Fields.Present || Pde->Fields.LargePage || Level == PagingLevelPageDirectory) {
        return Pde;
    }
    PtVa = PhysicalAddressToVirtualAddress(Pde->Fields.PageFrameNumber << 12);
    if (PtVa == NULL) {
        return NULL;
    }
    Offset         = MemoryMapperGetOffset(PagingLevelPageTable, Va);
    PPAGE_ENTRY Pt = &PtVa[Offset];
    return Pt;
}

_Use_decl_annotations_
BOOLEAN
MemoryMapperCheckIfPageIsPresentByCr3(PVOID Va, CR3_TYPE TargetCr3) {
    PPAGE_ENTRY PageEntry;
    PageEntry = MemoryMapperGetPteVaByCr3(Va, PagingLevelPageTable, TargetCr3);
    if (PageEntry != NULL && PageEntry->Fields.Present) {
        return TRUE;
    } else {
        return FALSE;
    }
}

_Use_decl_annotations_
BOOLEAN
MemoryMapperCheckIfPageIsNxBitSetByCr3(PVOID Va, CR3_TYPE TargetCr3) {
    PPAGE_ENTRY PageEntry;
    PageEntry = MemoryMapperGetPteVaByCr3(Va, PagingLevelPageTable, TargetCr3);
    if (PageEntry != NULL && !PageEntry->Fields.ExecuteDisable) {
        return TRUE;
    } else {
        return FALSE;
    }
}

_Use_decl_annotations_
BOOLEAN
MemoryMapperCheckIfPageIsNxBitSetOnTargetProcess(PVOID Va) {
    BOOLEAN     Result;
    CR3_TYPE    GuestCr3;
    PPAGE_ENTRY PageEntry;
    CR3_TYPE    CurrentProcessCr3 = {0};
    GuestCr3.Flags                = GetRunningCr3OnTargetProcess().Flags;
    CurrentProcessCr3             = SwitchOnAnotherProcessMemoryLayoutByCr3(GuestCr3);
    PageEntry                     = MemoryMapperGetPteVa(Va, PagingLevelPageTable);
    if (PageEntry != NULL && !PageEntry->Fields.ExecuteDisable) {
        Result = TRUE;
    } else {
        Result = FALSE;
    }
    RestoreToPreviousProcess(CurrentProcessCr3);
    return Result;
}

_Use_decl_annotations_
PVOID
MemoryMapperMapReservedPageRange(SIZE_T Size) {
    return MmAllocateMappingAddress(Size, POOLTAG);
}

_Use_decl_annotations_
VOID
MemoryMapperUnmapReservedPageRange(PVOID VirtualAddress) {
    MmFreeMappingAddress(VirtualAddress, POOLTAG);
}

_Use_decl_annotations_
PVOID
MemoryMapperGetPte(PVOID VirtualAddress) {
    return MemoryMapperGetPteVa(VirtualAddress, PagingLevelPageTable);
}

_Use_decl_annotations_
PVOID
MemoryMapperGetPteByCr3(PVOID VirtualAddress, CR3_TYPE TargetCr3) {
    return MemoryMapperGetPteVaByCr3(VirtualAddress, PagingLevelPageTable, TargetCr3);
}

_Use_decl_annotations_
PVOID
MemoryMapperMapPageAndGetPte(PUINT64 PteAddress) {
    UINT64 Va;
    UINT64 Pte;
    Va          = MemoryMapperMapReservedPageRange(PAGE_SIZE);
    Pte         = MemoryMapperGetPte(Va);
    *PteAddress = Pte;
    return Va;
}

VOID
MemoryMapperInitialize() {
    UINT64                  TempPte;
    UINT32                  ProcessorCount = KeQueryActiveProcessorCount(0);
    VIRTUAL_MACHINE_STATE * CurrentVmState = NULL;
    for (size_t i = 0; i < ProcessorCount; i++) {
        CurrentVmState                                 = &g_GuestState[i];
        CurrentVmState->MemoryMapper.VirualAddress     = MemoryMapperMapPageAndGetPte(&TempPte);
        CurrentVmState->MemoryMapper.PteVirtualAddress = TempPte;
    }
}

VOID
MemoryMapperUninitialize() {
    UINT32                  ProcessorCount = KeQueryActiveProcessorCount(0);
    VIRTUAL_MACHINE_STATE * CurrentVmState = NULL;
    for (size_t i = 0; i < ProcessorCount; i++) {
        CurrentVmState = &g_GuestState[i];
        if (CurrentVmState->MemoryMapper.VirualAddress != NULL)
            MemoryMapperUnmapReservedPageRange(CurrentVmState->MemoryMapper.VirualAddress);
        CurrentVmState->MemoryMapper.VirualAddress     = NULL;
        CurrentVmState->MemoryMapper.PteVirtualAddress = NULL;
    }
}

_Use_decl_annotations_
BOOLEAN
MemoryMapperReadMemorySafeByPte(PHYSICAL_ADDRESS PaAddressToRead,
                                PVOID            BufferToSaveMemory,
                                SIZE_T           SizeToRead,
                                UINT64           PteVaAddress,
                                UINT64           MappingVa,
                                BOOLEAN          InvalidateVpids) {
    PVOID       Va = MappingVa;
    PVOID       NewAddress;
    PAGE_ENTRY  PageEntry;
    PPAGE_ENTRY Pte                  = PteVaAddress;
    PageEntry.Flags                  = Pte->Flags;
    PageEntry.Fields.Present         = 1;
    PageEntry.Fields.Write           = 1;
    PageEntry.Fields.Global          = 1;
    PageEntry.Fields.PageFrameNumber = PaAddressToRead.QuadPart >> 12;
    Pte->Flags                       = PageEntry.Flags;
    __invlpg(Va);
    if (InvalidateVpids) {
    }
    NewAddress = (PVOID)((UINT64)Va + (PAGE_4KB_OFFSET & (PaAddressToRead.QuadPart)));
    memcpy(BufferToSaveMemory, NewAddress, SizeToRead);
    Pte->Flags = NULL;
    return TRUE;
}

_Use_decl_annotations_
BOOLEAN
MemoryMapperWriteMemorySafeByPte(PVOID            SourceVA,
                                 PHYSICAL_ADDRESS PaAddressToWrite,
                                 SIZE_T           SizeToWrite,
                                 UINT64           PteVaAddress,
                                 UINT64           MappingVa,
                                 BOOLEAN          InvalidateVpids) {
    PVOID       Va = MappingVa;
    PVOID       NewAddress;
    PAGE_ENTRY  PageEntry;
    PPAGE_ENTRY Pte                  = PteVaAddress;
    PageEntry.Flags                  = Pte->Flags;
    PageEntry.Fields.Present         = 1;
    PageEntry.Fields.Write           = 1;
    PageEntry.Fields.Global          = 1;
    PageEntry.Fields.PageFrameNumber = PaAddressToWrite.QuadPart >> 12;
    Pte->Flags                       = PageEntry.Flags;
    __invlpg(Va);
    if (InvalidateVpids) {
    }
    NewAddress = (PVOID)((UINT64)Va + (PAGE_4KB_OFFSET & (PaAddressToWrite.QuadPart)));
    memcpy(NewAddress, SourceVA, SizeToWrite);
    Pte->Flags = NULL;
    return TRUE;
}

_Use_decl_annotations_
UINT64
MemoryMapperReadMemorySafeByPhysicalAddressWrapperAddressMaker(
    MEMORY_MAPPER_WRAPPER_FOR_MEMORY_WRITE TypeOfRead,
    UINT64                                 AddressToRead) {
    PHYSICAL_ADDRESS PhysicalAddress = {0};
    switch (TypeOfRead) {
    case MEMORY_MAPPER_WRAPPER_READ_PHYSICAL_MEMORY:
        PhysicalAddress.QuadPart = AddressToRead;
        break;
    case MEMORY_MAPPER_WRAPPER_READ_VIRTUAL_MEMORY:
        PhysicalAddress.QuadPart = VirtualAddressToPhysicalAddress(AddressToRead);
        break;
    default:
        return NULL;
        break;
    }
    return PhysicalAddress.QuadPart;
}

_Use_decl_annotations_
BOOLEAN
MemoryMapperReadMemorySafeByPhysicalAddressWrapper(
    MEMORY_MAPPER_WRAPPER_FOR_MEMORY_WRITE TypeOfRead,
    UINT64                                 AddressToRead,
    UINT64                                 BufferToSaveMemory,
    SIZE_T                                 SizeToRead) {
    ULONG                   ProcessorIndex = KeGetCurrentProcessorNumber();
    UINT64                  AddressToCheck;
    PHYSICAL_ADDRESS        PhysicalAddress;
    VIRTUAL_MACHINE_STATE * CurrentVmState = &g_GuestState[ProcessorIndex];
    if (CurrentVmState->MemoryMapper.VirualAddress == NULL ||
        CurrentVmState->MemoryMapper.PteVirtualAddress == NULL) {
        return FALSE;
    }
    AddressToCheck = (CHAR *)AddressToRead + SizeToRead - ((CHAR *)PAGE_ALIGN(AddressToRead));
    if (AddressToCheck > PAGE_SIZE) {
        UINT64 ReadSize = AddressToCheck;
        while (SizeToRead != 0) {
            ReadSize = (UINT64)PAGE_ALIGN(AddressToRead + PAGE_SIZE) - AddressToRead;
            if (ReadSize == PAGE_SIZE && SizeToRead < PAGE_SIZE) {
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
                    CurrentVmState->IsOnVmxRootMode)) {
                return FALSE;
            }
            SizeToRead         = SizeToRead - ReadSize;
            AddressToRead      = AddressToRead + ReadSize;
            BufferToSaveMemory = BufferToSaveMemory + ReadSize;
        }
        return TRUE;
    } else {
        PhysicalAddress.QuadPart = MemoryMapperReadMemorySafeByPhysicalAddressWrapperAddressMaker(TypeOfRead,
                                                                                                  AddressToRead);
        return MemoryMapperReadMemorySafeByPte(
            PhysicalAddress,
            BufferToSaveMemory,
            SizeToRead,
            CurrentVmState->MemoryMapper.PteVirtualAddress,
            CurrentVmState->MemoryMapper.VirualAddress,
            CurrentVmState->IsOnVmxRootMode);
    }
}

_Use_decl_annotations_
BOOLEAN
MemoryMapperReadMemorySafeByPhysicalAddress(UINT64 PaAddressToRead,
                                            UINT64 BufferToSaveMemory,
                                            SIZE_T SizeToRead) {
    return MemoryMapperReadMemorySafeByPhysicalAddressWrapper(MEMORY_MAPPER_WRAPPER_READ_PHYSICAL_MEMORY,
                                                              PaAddressToRead,
                                                              BufferToSaveMemory,
                                                              SizeToRead);
}

_Use_decl_annotations_
BOOLEAN
MemoryMapperReadMemorySafe(UINT64 VaAddressToRead, PVOID BufferToSaveMemory, SIZE_T SizeToRead) {
    return MemoryMapperReadMemorySafeByPhysicalAddressWrapper(MEMORY_MAPPER_WRAPPER_READ_VIRTUAL_MEMORY,
                                                              VaAddressToRead,
                                                              BufferToSaveMemory,
                                                              SizeToRead);
}

_Use_decl_annotations_
BOOLEAN
MemoryMapperReadMemorySafeOnTargetProcess(UINT64 VaAddressToRead, PVOID BufferToSaveMemory, SIZE_T SizeToRead) {
    CR3_TYPE GuestCr3;
    CR3_TYPE OriginalCr3;
    BOOLEAN  Result;
    GuestCr3.Flags    = GetRunningCr3OnTargetProcess().Flags;
    OriginalCr3.Flags = __readcr3();
    __writecr3(GuestCr3.Flags);
    Result = MemoryMapperReadMemorySafe(VaAddressToRead, BufferToSaveMemory, SizeToRead);
    __writecr3(OriginalCr3.Flags);
    return Result;
}

_Use_decl_annotations_
BOOLEAN
MemoryMapperWriteMemorySafeOnTargetProcess(UINT64 Destination, PVOID Source, SIZE_T Size) {
    CR3_TYPE GuestCr3;
    CR3_TYPE OriginalCr3;
    BOOLEAN  Result;
    GuestCr3.Flags    = GetRunningCr3OnTargetProcess().Flags;
    OriginalCr3.Flags = __readcr3();
    __writecr3(GuestCr3.Flags);
    Result = MemoryMapperWriteMemorySafe(Destination, Source, Size, GuestCr3);
    __writecr3(OriginalCr3.Flags);
    return Result;
}

_Use_decl_annotations_
UINT64
MemoryMapperWriteMemorySafeWrapperAddressMaker(MEMORY_MAPPER_WRAPPER_FOR_MEMORY_WRITE TypeOfWrite,
                                               UINT64                                 DestinationAddr,
                                               PCR3_TYPE                              TargetProcessCr3,
                                               UINT32                                 TargetProcessId) {
    PHYSICAL_ADDRESS PhysicalAddress = {0};
    switch (TypeOfWrite) {
    case MEMORY_MAPPER_WRAPPER_WRITE_PHYSICAL_MEMORY:
        PhysicalAddress.QuadPart = DestinationAddr;
        break;
    case MEMORY_MAPPER_WRAPPER_WRITE_VIRTUAL_MEMORY_UNSAFE:
        if (TargetProcessId == NULL) {
            PhysicalAddress.QuadPart = VirtualAddressToPhysicalAddress(DestinationAddr);
        } else {
            PhysicalAddress.QuadPart = VirtualAddressToPhysicalAddressByProcessId(DestinationAddr, TargetProcessId);
        }
        break;
    case MEMORY_MAPPER_WRAPPER_WRITE_VIRTUAL_MEMORY_SAFE:
        if (TargetProcessCr3 == NULL || TargetProcessCr3->Flags == NULL) {
            PhysicalAddress.QuadPart = VirtualAddressToPhysicalAddress(DestinationAddr);
        } else {
            PhysicalAddress.QuadPart = VirtualAddressToPhysicalAddressByProcessCr3(DestinationAddr, *TargetProcessCr3);
        }
        break;
    default:
        return NULL;
        break;
    }
    return PhysicalAddress.QuadPart;
}

_Use_decl_annotations_
BOOLEAN
MemoryMapperWriteMemorySafeWrapper(MEMORY_MAPPER_WRAPPER_FOR_MEMORY_WRITE TypeOfWrite,
                                   UINT64                                 DestinationAddr,
                                   UINT64                                 Source,
                                   SIZE_T                                 SizeToWrite,
                                   PCR3_TYPE                              TargetProcessCr3,
                                   UINT32                                 TargetProcessId) {
    ULONG                   ProcessorIndex = KeGetCurrentProcessorNumber();
    UINT64                  AddressToCheck;
    PHYSICAL_ADDRESS        PhysicalAddress;
    VIRTUAL_MACHINE_STATE * CurrentVmState = &g_GuestState[ProcessorIndex];
    if (CurrentVmState->MemoryMapper.VirualAddress == NULL ||
        CurrentVmState->MemoryMapper.PteVirtualAddress == NULL) {
        return FALSE;
    }
    AddressToCheck = (CHAR *)DestinationAddr + SizeToWrite - ((CHAR *)PAGE_ALIGN(DestinationAddr));
    if (AddressToCheck > PAGE_SIZE) {
        UINT64 WriteSize = AddressToCheck;
        while (SizeToWrite != 0) {
            WriteSize = (UINT64)PAGE_ALIGN(DestinationAddr + PAGE_SIZE) - DestinationAddr;
            if (WriteSize == PAGE_SIZE && SizeToWrite < PAGE_SIZE) {
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
                    CurrentVmState->IsOnVmxRootMode)) {
                return FALSE;
            }
            SizeToWrite     = SizeToWrite - WriteSize;
            DestinationAddr = DestinationAddr + WriteSize;
            Source          = Source + WriteSize;
        }
        return TRUE;
    } else {
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
    }
}

_Use_decl_annotations_
BOOLEAN
MemoryMapperWriteMemorySafe(UINT64   Destination,
                            PVOID    Source,
                            SIZE_T   SizeToWrite,
                            CR3_TYPE TargetProcessCr3) {
    return MemoryMapperWriteMemorySafeWrapper(MEMORY_MAPPER_WRAPPER_WRITE_VIRTUAL_MEMORY_SAFE,
                                              Destination,
                                              Source,
                                              SizeToWrite,
                                              &TargetProcessCr3,
                                              NULL);
}

_Use_decl_annotations_
BOOLEAN
MemoryMapperWriteMemoryUnsafe(UINT64 Destination, PVOID Source, SIZE_T SizeToWrite, UINT32 TargetProcessId) {
    return MemoryMapperWriteMemorySafeWrapper(MEMORY_MAPPER_WRAPPER_WRITE_VIRTUAL_MEMORY_UNSAFE,
                                              Destination,
                                              Source,
                                              SizeToWrite,
                                              NULL,
                                              TargetProcessId);
}

_Use_decl_annotations_
BOOLEAN
MemoryMapperWriteMemorySafeByPhysicalAddress(UINT64 DestinationPa,
                                             UINT64 Source,
                                             SIZE_T SizeToWrite) {
    return MemoryMapperWriteMemorySafeWrapper(MEMORY_MAPPER_WRAPPER_WRITE_PHYSICAL_MEMORY,
                                              DestinationPa,
                                              Source,
                                              SizeToWrite,
                                              NULL,
                                              NULL);
}

_Use_decl_annotations_
UINT64
MemoryMapperReserveUsermodeAddressInTargetProcess(UINT32 ProcessId, BOOLEAN Allocate) {
    NTSTATUS   Status;
    PVOID      AllocPtr  = NULL;
    SIZE_T     AllocSize = PAGE_SIZE;
    PEPROCESS  SourceProcess;
    KAPC_STATE State = {0};
    if (PsGetCurrentProcessId() != ProcessId) {
        if (PsLookupProcessByProcessId(ProcessId, &SourceProcess) != STATUS_SUCCESS) {
            return NULL;
        }
        __try {
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
        } __except (EXCEPTION_EXECUTE_HANDLER) {
            KeUnstackDetachProcess(&State);
            ObDereferenceObject(SourceProcess);
            return NULL;
        }
    } else {
        Status = ZwAllocateVirtualMemory(
            NtCurrentProcess(),
            &AllocPtr,
            NULL,
            &AllocSize,
            Allocate ? MEM_COMMIT : MEM_RESERVE,
            PAGE_EXECUTE_READWRITE);
    }
    if (!NT_SUCCESS(Status)) {
        return NULL;
    }
    return AllocPtr;
}

_Use_decl_annotations_
BOOLEAN
MemoryMapperFreeMemoryInTargetProcess(UINT32 ProcessId,
                                      PVOID  BaseAddress) {
    NTSTATUS   Status;
    SIZE_T     AllocSize = PAGE_SIZE;
    PEPROCESS  SourceProcess;
    KAPC_STATE State = {0};
    if (PsGetCurrentProcessId() != ProcessId) {
        if (PsLookupProcessByProcessId(ProcessId, &SourceProcess) != STATUS_SUCCESS) {
            return FALSE;
        }
        __try {
            KeStackAttachProcess(SourceProcess, &State);
            Status = ZwFreeVirtualMemory(NtCurrentProcess(),
                                         &BaseAddress,
                                         &AllocSize,
                                         MEM_RELEASE);
            KeUnstackDetachProcess(&State);
            ObDereferenceObject(SourceProcess);
        } __except (EXCEPTION_EXECUTE_HANDLER) {
            KeUnstackDetachProcess(&State);
            ObDereferenceObject(SourceProcess);
            return FALSE;
        }
    } else {
        Status = ZwFreeVirtualMemory(NtCurrentProcess(),
                                     &BaseAddress,
                                     &AllocSize,
                                     MEM_RELEASE);
    }
    if (!NT_SUCCESS(Status)) {
        return FALSE;
    }
    return TRUE;
}

_Use_decl_annotations_
VOID
MemoryMapperMapPhysicalAddressToPte(PHYSICAL_ADDRESS PhysicalAddress,
                                    PVOID            TargetProcessVirtualAddress,
                                    CR3_TYPE         TargetProcessKernelCr3) {
    PPAGE_ENTRY PreviousPteEntry;
    PAGE_ENTRY  PageEntry;
    CR3_TYPE    CurrentProcessCr3;
    PreviousPteEntry                 = MemoryMapperGetPteVaByCr3(TargetProcessVirtualAddress, PagingLevelPageTable, TargetProcessKernelCr3);
    CurrentProcessCr3                = SwitchOnAnotherProcessMemoryLayoutByCr3(TargetProcessKernelCr3);
    PageEntry.Flags                  = PreviousPteEntry->Flags;
    PageEntry.Fields.Present         = 1;
    PageEntry.Fields.Supervisor      = 1;
    PageEntry.Fields.Write           = 1;
    PageEntry.Fields.Global          = 1;
    PageEntry.Fields.PageFrameNumber = PhysicalAddress.QuadPart >> 12;
    PreviousPteEntry->Flags          = PageEntry.Flags;
    __invlpg(TargetProcessVirtualAddress);
    RestoreToPreviousProcess(CurrentProcessCr3);
}

_Use_decl_annotations_
BOOLEAN
MemoryMapperSetSupervisorBitWithoutSwitchingByCr3(PVOID Va, BOOLEAN Set, PAGING_LEVEL Level, CR3_TYPE TargetCr3) {
    PPAGE_ENTRY Pml = NULL;
    Pml             = MemoryMapperGetPteVaWithoutSwitchingByCr3(Va, Level, TargetCr3);
    if (!Pml) {
        return FALSE;
    }
    if (Set) {
        Pml->Fields.Supervisor = 1;
    } else {
        Pml->Fields.Supervisor = 0;
    }
    return TRUE;
}
