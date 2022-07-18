#include "pch.h"
int
MathPower(int Base, int Exp) {
    int result;
    result = 1;
    for (;;) {
        if (Exp & 1) {
            result *= Base;
        }
        Exp >>= 1;
        if (!Exp) {
            break;
        }
        Base *= Base;
    }
    return result;
}

_Use_decl_annotations_
BOOLEAN
BroadcastToProcessors(ULONG ProcessorNumber, RunOnLogicalCoreFunc Routine) {
    KIRQL OldIrql;
    KeSetSystemAffinityThread((KAFFINITY)(1 << ProcessorNumber));
    OldIrql = KeRaiseIrqlToDpcLevel();
    Routine(ProcessorNumber);
    KeLowerIrql(OldIrql);
    KeRevertToUserAffinityThread();
    return TRUE;
}

int
TestBit(int nth, unsigned long * addr) {
    return (BITMAP_ENTRY(nth, addr) >> BITMAP_SHIFT(nth)) & 1;
}

void
ClearBit(int nth, unsigned long * addr) {
    BITMAP_ENTRY(nth, addr) &= ~(1UL << BITMAP_SHIFT(nth));
}

void
SetBit(int nth, unsigned long * addr) {
    BITMAP_ENTRY(nth, addr) |= (1UL << BITMAP_SHIFT(nth));
}

_Use_decl_annotations_
UINT64
VirtualAddressToPhysicalAddress(_In_ PVOID VirtualAddress) {
    return MmGetPhysicalAddress(VirtualAddress).QuadPart;
}

_Use_decl_annotations_
CR3_TYPE
GetCr3FromProcessId(UINT32 ProcessId) {
    PEPROCESS TargetEprocess;
    CR3_TYPE  ProcessCr3 = {0};
    if (PsLookupProcessByProcessId(ProcessId, &TargetEprocess) != STATUS_SUCCESS) {
        return ProcessCr3;
    }
    NT_KPROCESS * CurrentProcess = (NT_KPROCESS *)(TargetEprocess);
    ProcessCr3.Flags             = CurrentProcess->DirectoryTableBase;
    ObDereferenceObject(TargetEprocess);
    return ProcessCr3;
}

_Use_decl_annotations_
CR3_TYPE
SwitchOnAnotherProcessMemoryLayout(UINT32 ProcessId) {
    UINT64    GuestCr3;
    PEPROCESS TargetEprocess;
    CR3_TYPE  CurrentProcessCr3 = {0};
    if (PsLookupProcessByProcessId(ProcessId, &TargetEprocess) != STATUS_SUCCESS) {
        return CurrentProcessCr3;
    }
    NT_KPROCESS * CurrentProcess = (NT_KPROCESS *)(TargetEprocess);
    GuestCr3                     = CurrentProcess->DirectoryTableBase;
    CurrentProcessCr3.Flags      = __readcr3();
    __writecr3(GuestCr3);
    ObDereferenceObject(TargetEprocess);
    return CurrentProcessCr3;
}

CR3_TYPE
SwitchOnMemoryLayoutOfTargetProcess() {
    CR3_TYPE GuestCr3;
    CR3_TYPE CurrentProcessCr3 = {0};
    GuestCr3.Flags             = GetRunningCr3OnTargetProcess().Flags;
    CurrentProcessCr3.Flags    = __readcr3();
    __writecr3(GuestCr3.Flags);
    return CurrentProcessCr3;
}

_Use_decl_annotations_
CR3_TYPE
SwitchOnAnotherProcessMemoryLayoutByCr3(CR3_TYPE TargetCr3) {
    CR3_TYPE CurrentProcessCr3 = {0};
    CurrentProcessCr3.Flags    = __readcr3();
    __writecr3(TargetCr3.Flags);
    return CurrentProcessCr3;
}

_Use_decl_annotations_
BOOLEAN
GetSegmentDescriptor(PUCHAR GdtBase, UINT16 Selector, PVMX_SEGMENT_SELECTOR SegmentSelector) {
    SEGMENT_DESCRIPTOR_32 * DescriptorTable32;
    SEGMENT_DESCRIPTOR_32 * Descriptor32;
    SEGMENT_SELECTOR        SegSelector = {.AsUInt = Selector};
    if (!SegmentSelector)
        return FALSE;
#define SELECTOR_TABLE_LDT 0x1
#define SELECTOR_TABLE_GDT 0x0
    if ((Selector == 0x0) || (SegSelector.Table != SELECTOR_TABLE_GDT)) {
        return FALSE;
    }
    DescriptorTable32                  = (SEGMENT_DESCRIPTOR_32 *)(GdtBase);
    Descriptor32                       = &DescriptorTable32[SegSelector.Index];
    SegmentSelector->Selector          = Selector;
    SegmentSelector->Limit             = __segmentlimit(Selector);
    SegmentSelector->Base              = (Descriptor32->BaseAddressLow | Descriptor32->BaseAddressMiddle << 16 | Descriptor32->BaseAddressHigh << 24);
    SegmentSelector->Attributes.AsUInt = (AsmGetAccessRights(Selector) >> 8);
    if (SegSelector.Table == 0 && SegSelector.Index == 0) {
        SegmentSelector->Attributes.Unusable = TRUE;
    }
    if ((Descriptor32->Type == SEGMENT_DESCRIPTOR_TYPE_TSS_BUSY) ||
        (Descriptor32->Type == SEGMENT_DESCRIPTOR_TYPE_CALL_GATE)) {
        UINT64 SegmentLimitHigh;
        SegmentLimitHigh      = (*(UINT64 *)((PUCHAR)Descriptor32 + 8));
        SegmentSelector->Base = (SegmentSelector->Base & 0xffffffff) | (SegmentLimitHigh << 32);
    }
    if (SegmentSelector->Attributes.Granularity) {
        SegmentSelector->Limit = (SegmentSelector->Limit << 12) + 0xfff;
    }
    return TRUE;
}

_Use_decl_annotations_
VOID
RestoreToPreviousProcess(CR3_TYPE PreviousProcess) {
    __writecr3(PreviousProcess.Flags);
}

_Use_decl_annotations_
UINT64
PhysicalAddressToVirtualAddressByProcessId(PVOID PhysicalAddress, UINT32 ProcessId) {
    CR3_TYPE         CurrentProcessCr3;
    UINT64           VirtualAddress;
    PHYSICAL_ADDRESS PhysicalAddr;
    CurrentProcessCr3 = SwitchOnAnotherProcessMemoryLayout(ProcessId);
    if (CurrentProcessCr3.Flags == NULL) {
        return NULL;
    }
    PhysicalAddr.QuadPart = PhysicalAddress;
    VirtualAddress        = MmGetVirtualForPhysical(PhysicalAddr);
    RestoreToPreviousProcess(CurrentProcessCr3);
    return VirtualAddress;
}

_Use_decl_annotations_
UINT64
PhysicalAddressToVirtualAddressByCr3(PVOID PhysicalAddress, CR3_TYPE TargetCr3) {
    CR3_TYPE         CurrentProcessCr3;
    UINT64           VirtualAddress;
    PHYSICAL_ADDRESS PhysicalAddr;
    CurrentProcessCr3 = SwitchOnAnotherProcessMemoryLayoutByCr3(TargetCr3);
    if (CurrentProcessCr3.Flags == NULL) {
        return NULL;
    }
    PhysicalAddr.QuadPart = PhysicalAddress;
    VirtualAddress        = MmGetVirtualForPhysical(PhysicalAddr);
    RestoreToPreviousProcess(CurrentProcessCr3);
    return VirtualAddress;
}

_Use_decl_annotations_
UINT64
PhysicalAddressToVirtualAddressOnTargetProcess(PVOID PhysicalAddress) {
    CR3_TYPE GuestCr3;
    GuestCr3.Flags = GetRunningCr3OnTargetProcess().Flags;
    UINT64 Temp    = (UINT64)PhysicalAddress;
    return PhysicalAddressToVirtualAddressByCr3(PhysicalAddress, GuestCr3);
}

CR3_TYPE
GetRunningCr3OnTargetProcess() {
    CR3_TYPE      GuestCr3;
    NT_KPROCESS * CurrentProcess = (NT_KPROCESS *)(PsGetCurrentProcess());
    GuestCr3.Flags               = CurrentProcess->DirectoryTableBase;
    return GuestCr3;
}

_Use_decl_annotations_
UINT64
VirtualAddressToPhysicalAddressByProcessId(PVOID VirtualAddress, UINT32 ProcessId) {
    CR3_TYPE CurrentProcessCr3;
    UINT64   PhysicalAddress;
    CurrentProcessCr3 = SwitchOnAnotherProcessMemoryLayout(ProcessId);
    if (CurrentProcessCr3.Flags == NULL) {
        return NULL;
    }
    PhysicalAddress = MmGetPhysicalAddress(VirtualAddress).QuadPart;
    RestoreToPreviousProcess(CurrentProcessCr3);
    return PhysicalAddress;
}

_Use_decl_annotations_
UINT64
VirtualAddressToPhysicalAddressByProcessCr3(PVOID VirtualAddress, CR3_TYPE TargetCr3) {
    CR3_TYPE CurrentProcessCr3;
    UINT64   PhysicalAddress;
    CurrentProcessCr3 = SwitchOnAnotherProcessMemoryLayoutByCr3(TargetCr3);
    if (CurrentProcessCr3.Flags == NULL) {
        return NULL;
    }
    PhysicalAddress = MmGetPhysicalAddress(VirtualAddress).QuadPart;
    RestoreToPreviousProcess(CurrentProcessCr3);
    return PhysicalAddress;
}

_Use_decl_annotations_
UINT64
VirtualAddressToPhysicalAddressOnTargetProcess(PVOID VirtualAddress) {
    CR3_TYPE CurrentCr3;
    CR3_TYPE GuestCr3;
    UINT64   PhysicalAddress;
    GuestCr3.Flags = GetRunningCr3OnTargetProcess().Flags;
    CurrentCr3     = SwitchOnAnotherProcessMemoryLayoutByCr3(GuestCr3);
    if (CurrentCr3.Flags == NULL) {
        return NULL;
    }
    PhysicalAddress = MmGetPhysicalAddress(VirtualAddress).QuadPart;
    RestoreToPreviousProcess(CurrentCr3);
    return PhysicalAddress;
}

_Use_decl_annotations_
UINT64
PhysicalAddressToVirtualAddress(UINT64 PhysicalAddress) {
    PHYSICAL_ADDRESS PhysicalAddr;
    PhysicalAddr.QuadPart = PhysicalAddress;
    return MmGetVirtualForPhysical(PhysicalAddr);
}

UINT64
FindSystemDirectoryTableBase() {
    NT_KPROCESS * SystemProcess = (NT_KPROCESS *)(PsInitialSystemProcess);
    return SystemProcess->DirectoryTableBase;
}

PCHAR
GetProcessNameFromEprocess(PEPROCESS Eprocess) {
    PCHAR Result = 0;
    Result       = (CHAR *)PsGetProcessImageFileName(Eprocess);
    return Result;
}

BOOLEAN
StartsWith(const char * pre, const char * str) {
    size_t lenpre = strlen(pre),
           lenstr = strlen(str);
    return lenstr < lenpre ? FALSE : memcmp(pre, str, lenpre) == 0;
}

BOOLEAN
IsProcessExist(UINT32 ProcId) {
    PEPROCESS TargetEprocess;
    CR3_TYPE  CurrentProcessCr3 = {0};
    if (PsLookupProcessByProcessId(ProcId, &TargetEprocess) != STATUS_SUCCESS) {
        return FALSE;
    } else {
        ObDereferenceObject(TargetEprocess);
        return TRUE;
    }
}

BOOLEAN
CheckIfAddressIsValidUsingTsx(CHAR * Address) {
    UINT32  Status = 0;
    BOOLEAN Result = FALSE;
    CHAR    TempContent;
    if ((Status = _xbegin()) == _XBEGIN_STARTED) {
        TempContent = *(CHAR *)Address;
        _xend();
        Result = TRUE;
    } else {
        Result = FALSE;
    }
    return Result;
}

VOID
GetCpuid(UINT32 Func, UINT32 SubFunc, int * CpuInfo) {
    __cpuidex(CpuInfo, Func, SubFunc);
}

BOOLEAN
CheckCpuSupportRtm() {
    int     Regs1[4];
    int     Regs2[4];
    BOOLEAN Result;
    GetCpuid(0, 0, Regs1);
    GetCpuid(7, 0, Regs2);
    Result = Regs1[0] >= 0x7 && (Regs2[1] & 0x4800) == 0x4800;
    return Result;
}

UINT32
Getx86VirtualAddressWidth() {
    int Regs[4];
    GetCpuid(CPUID_ADDR_WIDTH, 0, Regs);
    return ((Regs[0] >> 8) & 0x0ff);
}

BOOLEAN
CheckCanonicalVirtualAddress(UINT64 VAddr, PBOOLEAN IsKernelAddress) {
    UINT64 Addr = (UINT64)VAddr;
    UINT64 MaxVirtualAddrLowHalf, MinVirtualAddressHighHalf;
    UINT32 AddrWidth          = g_VirtualAddressWidth;
    MaxVirtualAddrLowHalf     = ((UINT64)1ull << (AddrWidth - 1)) - 1;
    MinVirtualAddressHighHalf = ~MaxVirtualAddrLowHalf;
    if ((Addr > MaxVirtualAddrLowHalf) && (Addr < MinVirtualAddressHighHalf)) {
        *IsKernelAddress = FALSE;
        return FALSE;
    }
    if (MinVirtualAddressHighHalf < Addr) {
        *IsKernelAddress = TRUE;
    } else {
        *IsKernelAddress = FALSE;
    }
    return TRUE;
}

BOOLEAN
CheckMemoryAccessSafety(UINT64 TargetAddress, UINT32 Size) {
    CR3_TYPE GuestCr3;
    UINT64   OriginalCr3;
    BOOLEAN  IsKernelAddress;
    BOOLEAN  Result = FALSE;
    if (!CheckCanonicalVirtualAddress(TargetAddress, &IsKernelAddress)) {
        Result = FALSE;
        goto Return;
    }
    GuestCr3.Flags = GetRunningCr3OnTargetProcess().Flags;
    OriginalCr3    = __readcr3();
    __writecr3(GuestCr3.Flags);
    UINT64 AddressToCheck =
        (CHAR *)TargetAddress + Size - ((CHAR *)PAGE_ALIGN(TargetAddress));
    if (AddressToCheck > PAGE_SIZE) {
        UINT64 ReadSize = AddressToCheck;
        while (Size != 0) {
            ReadSize = (UINT64)PAGE_ALIGN(TargetAddress + PAGE_SIZE) - TargetAddress;
            if (ReadSize == PAGE_SIZE && Size < PAGE_SIZE) {
                ReadSize = Size;
            }
            if (!MemoryMapperCheckIfPageIsPresentByCr3(TargetAddress, GuestCr3)) {
                Result = FALSE;
                goto RestoreCr3;
            }
            Size          = Size - ReadSize;
            TargetAddress = TargetAddress + ReadSize;
        }
    } else {
        if (!MemoryMapperCheckIfPageIsPresentByCr3(TargetAddress, GuestCr3)) {
            Result = FALSE;
            goto RestoreCr3;
        }
    }
    Result = TRUE;
RestoreCr3:
    __writecr3(OriginalCr3);
Return:
    return Result;
}

UINT32
VmxrootCompatibleStrlen(const CHAR * S) {
    CHAR     Temp  = NULL;
    UINT32   Count = 0;
    UINT64   AlignedAddress;
    CR3_TYPE GuestCr3;
    CR3_TYPE OriginalCr3;
    AlignedAddress    = (UINT64)PAGE_ALIGN((UINT64)S);
    GuestCr3.Flags    = GetRunningCr3OnTargetProcess().Flags;
    OriginalCr3.Flags = __readcr3();
    __writecr3(GuestCr3.Flags);
    if (!CheckMemoryAccessSafety(AlignedAddress, sizeof(CHAR))) {
        __writecr3(OriginalCr3.Flags);
        return 0;
    }
    while (TRUE) {
        MemoryMapperReadMemorySafe(S, &Temp, sizeof(CHAR));
        if (Temp != '\0') {
            Count++;
            S++;
        } else {
            __writecr3(OriginalCr3.Flags);
            return Count;
        }
        if (!((UINT64)S & (PAGE_SIZE - 1))) {
            if (!CheckMemoryAccessSafety((UINT64)S, sizeof(CHAR))) {
                __writecr3(OriginalCr3.Flags);
                return 0;
            }
        }
    }
    __writecr3(OriginalCr3.Flags);
}

UINT32
VmxrootCompatibleWcslen(const wchar_t * S) {
    wchar_t  Temp  = NULL;
    UINT32   Count = 0;
    UINT64   AlignedAddress;
    CR3_TYPE GuestCr3;
    CR3_TYPE OriginalCr3;
    AlignedAddress    = (UINT64)PAGE_ALIGN((UINT64)S);
    GuestCr3.Flags    = GetRunningCr3OnTargetProcess().Flags;
    OriginalCr3.Flags = __readcr3();
    __writecr3(GuestCr3.Flags);
    AlignedAddress = (UINT64)PAGE_ALIGN((UINT64)S);
    if (!CheckMemoryAccessSafety(AlignedAddress, sizeof(wchar_t))) {
        __writecr3(OriginalCr3.Flags);
        return 0;
    }
    while (TRUE) {
        MemoryMapperReadMemorySafe(S, &Temp, sizeof(wchar_t));
        if (Temp != '\0\0') {
            Count++;
            S++;
        } else {
            __writecr3(OriginalCr3.Flags);
            return Count;
        }
        if (!((UINT64)S & (PAGE_SIZE - 1))) {
            if (!CheckMemoryAccessSafety((UINT64)S, sizeof(wchar_t))) {
                __writecr3(OriginalCr3.Flags);
                return 0;
            }
        }
    }
    __writecr3(OriginalCr3.Flags);
}

UINT64 *
AllocateInvalidMsrBimap() {
    UINT64 * InvalidMsrBitmap;
    InvalidMsrBitmap = ExAllocatePoolWithTag(NonPagedPool, 0x1000 / 0x8, POOLTAG);
    if (InvalidMsrBitmap == NULL) {
        return NULL;
    }
    RtlZeroMemory(InvalidMsrBitmap, 0x1000 / 0x8);
    for (size_t i = 0; i < 0x1000; ++i) {
        __try {
            __readmsr(i);
        } __except (EXCEPTION_EXECUTE_HANDLER) {
            SetBit(i, InvalidMsrBitmap);
        }
    }
    return InvalidMsrBitmap;
}

_Use_decl_annotations_
NTSTATUS
GetHandleFromProcess(UINT32 ProcessId, PHANDLE Handle) {
    NTSTATUS Status;
    Status = STATUS_SUCCESS;
    OBJECT_ATTRIBUTES ObjAttr;
    CLIENT_ID         Cid;
    InitializeObjectAttributes(&ObjAttr, NULL, 0, NULL, NULL);
    Cid.UniqueProcess = ProcessId;
    Cid.UniqueThread  = (HANDLE)0;
    Status            = ZwOpenProcess(Handle, PROCESS_ALL_ACCESS, &ObjAttr, &Cid);
    return Status;
}

NTSTATUS
UndocumentedNtOpenProcess(
    PHANDLE         ProcessHandle,
    ACCESS_MASK     DesiredAccess,
    HANDLE          ProcessId,
    KPROCESSOR_MODE AccessMode) {
    NTSTATUS     status = STATUS_SUCCESS;
    ACCESS_STATE accessState;
    char         auxData[0x200];
    PEPROCESS    processObject = NULL;
    HANDLE       processHandle = NULL;
    status                     = SeCreateAccessState(
        &accessState,
        auxData,
        DesiredAccess,
        (PGENERIC_MAPPING)((PCHAR)*PsProcessType + 52));
    if (!NT_SUCCESS(status))
        return status;
    accessState.PreviouslyGrantedAccess |= accessState.RemainingDesiredAccess;
    accessState.RemainingDesiredAccess = 0;
    status                             = PsLookupProcessByProcessId(ProcessId, &processObject);
    if (!NT_SUCCESS(status)) {
        SeDeleteAccessState(&accessState);
        return status;
    }
    status = ObOpenObjectByPointer(
        processObject,
        0,
        &accessState,
        0,
        *PsProcessType,
        AccessMode,
        &processHandle);
    SeDeleteAccessState(&accessState);
    ObDereferenceObject(processObject);
    if (NT_SUCCESS(status))
        *ProcessHandle = processHandle;
    return status;
}

_Use_decl_annotations_
BOOLEAN
KillProcess(UINT32 ProcessId, PROCESS_KILL_METHODS KillingMethod) {
    NTSTATUS  Status        = STATUS_SUCCESS;
    HANDLE    ProcessHandle = NULL;
    PEPROCESS Process       = NULL;
    if (ProcessId == NULL) {
        return FALSE;
    }
    switch (KillingMethod) {
    case PROCESS_KILL_METHOD_1:
        Status = GetHandleFromProcess(ProcessId, &ProcessHandle);
        if (!NT_SUCCESS(Status) || ProcessHandle == NULL) {
            return FALSE;
        }
        Status = ZwTerminateProcess(ProcessHandle, 0);
        if (!NT_SUCCESS(Status)) {
            return FALSE;
        }
        break;
    case PROCESS_KILL_METHOD_2:
        UndocumentedNtOpenProcess(
            &ProcessHandle,
            PROCESS_ALL_ACCESS,
            ProcessId,
            KernelMode);
        if (ProcessHandle == NULL) {
            return FALSE;
        }
        Status = ZwTerminateProcess(ProcessHandle, 0);
        if (!NT_SUCCESS(Status)) {
            return FALSE;
        }
        break;
    case PROCESS_KILL_METHOD_3:
        Status = MmUnmapViewOfSection(Process, PsGetProcessSectionBaseAddress(Process));
        ObDereferenceObject(Process);
        break;
    default:
        return FALSE;
        break;
    }
    return TRUE;
}
