#include "pch.h"
NTSTATUS
MemoryManagerReadProcessMemoryNormal(HANDLE                    PID,
                                     PVOID                     Address,
                                     DEBUGGER_READ_MEMORY_TYPE MemType,
                                     PVOID                     UserBuffer,
                                     SIZE_T                    Size,
                                     PSIZE_T                   ReturnSize) {
    PEPROCESS        SourceProcess;
    MM_COPY_ADDRESS  CopyAddress         = {0};
    KAPC_STATE       State               = {0};
    PHYSICAL_ADDRESS TempPhysicalAddress = {0};
    if (PsGetCurrentProcessId() != PID && MemType == DEBUGGER_READ_VIRTUAL_ADDRESS) {
        if (PsLookupProcessByProcessId(PID, &SourceProcess) != STATUS_SUCCESS) {
            return STATUS_UNSUCCESSFUL;
        }
        __try {
            KeStackAttachProcess(SourceProcess, &State);
            TempPhysicalAddress = MmGetPhysicalAddress(Address);
            KeUnstackDetachProcess(&State);
            CopyAddress.PhysicalAddress.QuadPart = TempPhysicalAddress.QuadPart;
            MmCopyMemory(UserBuffer, CopyAddress, Size, MM_COPY_MEMORY_PHYSICAL, ReturnSize);
            ObDereferenceObject(SourceProcess);
            return STATUS_SUCCESS;
        } __except (EXCEPTION_EXECUTE_HANDLER) {
            KeUnstackDetachProcess(&State);
            ObDereferenceObject(SourceProcess);
            return STATUS_UNSUCCESSFUL;
        }
    } else {
        __try {
            if (MemType == DEBUGGER_READ_VIRTUAL_ADDRESS) {
                CopyAddress.VirtualAddress = Address;
                MmCopyMemory(UserBuffer, CopyAddress, Size, MM_COPY_MEMORY_VIRTUAL, ReturnSize);
            } else if (MemType == DEBUGGER_READ_PHYSICAL_ADDRESS) {
                CopyAddress.PhysicalAddress.QuadPart = Address;
                MmCopyMemory(UserBuffer, CopyAddress, Size, MM_COPY_MEMORY_PHYSICAL, ReturnSize);
            } else {
                return STATUS_UNSUCCESSFUL;
            }
            return STATUS_SUCCESS;
        } __except (EXCEPTION_EXECUTE_HANDLER) {
            return STATUS_ACCESS_DENIED;
        }
    }
}
