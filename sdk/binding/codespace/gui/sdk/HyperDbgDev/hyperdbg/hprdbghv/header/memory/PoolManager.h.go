package memory

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbghv\header\memory\PoolManager.h.back

const (
	TRACKING_HOOKED_PAGES           = 1 //col:3
	EXEC_TRAMPOLINE                 = 2 //col:4
	SPLIT_2MB_PAGING_TO_4KB_PAGE    = 3 //col:5
	DETOUR_HOOK_DETAILS             = 4 //col:6
	BREAKPOINT_DEFINITION_STRUCTURE = 5 //col:7
	PROCESS_THREAD_HOLDER           = 6 //col:8
)

type _POOL_TABLE struct {
	Address       uint64                    //col:12
	Size          int64                     //col:13
	Intention     POOL_ALLOCATION_INTENTION //col:14
	PoolsList     *list.List                //col:15
	IsBusy        bool                      //col:16
	ShouldBeFreed bool                      //col:17
	AlreadyFreed  bool                      //col:18
}

type _REQUEST_NEW_ALLOCATION struct {
	Size      int64                     //col:18
	Count     uint32                    //col:19
	Intention POOL_ALLOCATION_INTENTION //col:20
}

