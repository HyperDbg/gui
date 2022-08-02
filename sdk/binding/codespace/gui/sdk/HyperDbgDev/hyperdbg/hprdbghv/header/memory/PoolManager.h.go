package memory



const (
	TRACKING_HOOKED_PAGES           = 1 //col:3
	EXEC_TRAMPOLINE                 = 2 //col:4
	SPLIT_2MB_PAGING_TO_4KB_PAGE    = 3 //col:5
	DETOUR_HOOK_DETAILS             = 4 //col:6
	BREAKPOINT_DEFINITION_STRUCTURE = 5 //col:7
	PROCESS_THREAD_HOLDER           = 6 //col:8
)

type POOL_TABLE struct {
	Address       uint64                    //col:3
	Size          SIZE_T                    //col:4
	Intention     POOL_ALLOCATION_INTENTION //col:5
	PoolsList     *list.List                //col:6
	IsBusy        bool                      //col:7
	ShouldBeFreed bool                      //col:8
	AlreadyFreed  bool                      //col:9
}


type REQUEST_NEW_ALLOCATION struct {
	Size      SIZE_T                    //col:13
	Count     uint32                    //col:14
	Intention POOL_ALLOCATION_INTENTION //col:15
}


