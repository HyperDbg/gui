#include "pch.h"
PVOID
CrsAllocateContiguousZeroedMemory(_In_ SIZE_T NumberOfBytes) {
    PVOID            Result          = NULL;
    PHYSICAL_ADDRESS MaxPhysicalAddr = {.QuadPart = MAXULONG64};
    Result                           = MmAllocateContiguousMemory(NumberOfBytes, MaxPhysicalAddr);
    if (Result != NULL)
        RtlSecureZeroMemory(Result, NumberOfBytes);
    return Result;
}

PVOID
CrsAllocateNonPagedPool(SIZE_T NumberOfBytes) {
    PVOID Result = ExAllocatePoolWithTag(NonPagedPool, NumberOfBytes, POOLTAG);
    if (Result != NULL)
        RtlSecureZeroMemory(Result, NumberOfBytes);
    return Result;
}
