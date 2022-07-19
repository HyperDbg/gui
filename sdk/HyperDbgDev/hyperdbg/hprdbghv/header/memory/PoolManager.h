#pragma once
#define MaximumRequestsQueueDepth   100
#define NumberOfPreAllocatedBuffers 10
typedef enum _POOL_ALLOCATION_INTENTION {
    TRACKING_HOOKED_PAGES,
    EXEC_TRAMPOLINE,
    SPLIT_2MB_PAGING_TO_4KB_PAGE,
    DETOUR_HOOK_DETAILS,
    BREAKPOINT_DEFINITION_STRUCTURE,
    PROCESS_THREAD_HOLDER,
} POOL_ALLOCATION_INTENTION;
typedef struct _POOL_TABLE {
    UINT64                    Address; // Should be the start of the list as we compute it as the start address
    SIZE_T                    Size;
    POOL_ALLOCATION_INTENTION Intention;
    LIST_ENTRY                PoolsList;
    BOOLEAN                   IsBusy;
    BOOLEAN                   ShouldBeFreed;
    BOOLEAN                   AlreadyFreed;
} POOL_TABLE, *PPOOL_TABLE;
typedef struct _REQUEST_NEW_ALLOCATION {
    SIZE_T                    Size;
    UINT32                    Count;
    POOL_ALLOCATION_INTENTION Intention;
} REQUEST_NEW_ALLOCATION, *PREQUEST_NEW_ALLOCATION;
REQUEST_NEW_ALLOCATION * g_RequestNewAllocation;
volatile LONG            LockForRequestAllocation;
volatile LONG            LockForReadingPool;
BOOLEAN                  g_IsNewRequestForAllocationReceived;
BOOLEAN                  g_IsNewRequestForDeAllocation;
LIST_ENTRY               g_ListOfAllocatedPoolsHead;
static BOOLEAN
            PlmgrAllocateRequestNewAllocation(SIZE_T NumberOfBytes);
static VOID PlmgrFreeRequestNewAllocation(VOID);
BOOLEAN
PoolManagerInitialize();
BOOLEAN
PoolManagerCheckAndPerformAllocationAndDeallocation();
BOOLEAN
PoolManagerRequestAllocation(SIZE_T Size, UINT32 Count, POOL_ALLOCATION_INTENTION Intention);
UINT64
PoolManagerRequestPool(POOL_ALLOCATION_INTENTION Intention, BOOLEAN RequestNewPool, UINT32 Size);
VOID
PoolManagerUninitialize();
BOOLEAN
PoolManagerFreePool(UINT64 AddressToFree);
