package memory

//back\HyperDbgDev\hyperdbg\hprdbghv\header\memory\PoolManager.h.back

const (
	MaximumRequestsQueueDepth   = 100 //col:22
	NumberOfPreAllocatedBuffers = 10  //col:23
)

type TRACKING_HOOKED_PAGES uint32

const (
	TRACKING_HOOKED_PAGES           POOL_ALLOCATION_INTENTION = 1 //col:35
	EXEC_TRAMPOLINE                 POOL_ALLOCATION_INTENTION = 2 //col:36
	SPLIT_2MB_PAGING_TO_4KB_PAGE    POOL_ALLOCATION_INTENTION = 3 //col:37
	DETOUR_HOOK_DETAILS             POOL_ALLOCATION_INTENTION = 4 //col:38
	BREAKPOINT_DEFINITION_STRUCTURE POOL_ALLOCATION_INTENTION = 5 //col:39
	PROCESS_THREAD_HOLDER           POOL_ALLOCATION_INTENTION = 6 //col:40
)
