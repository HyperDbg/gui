#pragma once
PVOID
CrsAllocateContiguousZeroedMemory(_In_ SIZE_T NumberOfBytes);
PVOID
CrsAllocateNonPagedPool(SIZE_T NumberOfBytes);
