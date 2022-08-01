package memory
//back\HyperDbgDev\hyperdbg\hprdbghv\header\memory\PoolManager.h.back

const(
MaximumRequestsQueueDepth =   100 //col:1
NumberOfPreAllocatedBuffers = 10 //col:2
)

const(
    TRACKING_HOOKED_PAGES = 1  //col:3
    EXEC_TRAMPOLINE = 2  //col:4
    SPLIT_2MB_PAGING_TO_4KB_PAGE = 3  //col:5
    DETOUR_HOOK_DETAILS = 4  //col:6
    BREAKPOINT_DEFINITION_STRUCTURE = 5  //col:7
    PROCESS_THREAD_HOLDER = 6  //col:8
)



type POOL_TABLE struct{
Address uint64
Size SIZE_T
Intention POOL_ALLOCATION_INTENTION
PoolsList LIST_ENTRY
IsBusy bool
ShouldBeFreed bool
AlreadyFreed bool
}


type REQUEST_NEW_ALLOCATION struct{
Size SIZE_T
Count UINT32
Intention POOL_ALLOCATION_INTENTION
}




