package memory

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbghv\header\memory\MemoryMapper.h.back

const (
	PAGE_4KB_OFFSET = ((UINT64)(1<<12) - 1) //col:1
	PAGE_2MB_OFFSET = ((UINT64)(1<<21) - 1) //col:2
	PAGE_4MB_OFFSET = ((UINT64)(1<<22) - 1) //col:3
	PAGE_1GB_OFFSET = ((UINT64)(1<<30) - 1) //col:4
)

const (
	PagingLevelPageTable                 = 0 //col:3
	PagingLevelPageDirectory             = 2 //col:4
	PagingLevelPageDirectoryPointerTable = 3 //col:5
	PagingLevelPageMapLevel4             = 4 //col:6
)

const (
	MEMORY_MAPPER_WRAPPER_READ_PHYSICAL_MEMORY = 1 //col:10
	MEMORY_MAPPER_WRAPPER_READ_VIRTUAL_MEMORY  = 2 //col:11
)

const (
	MEMORY_MAPPER_WRAPPER_WRITE_PHYSICAL_MEMORY       = 1 //col:15
	MEMORY_MAPPER_WRAPPER_WRITE_VIRTUAL_MEMORY_SAFE   = 2 //col:16
	MEMORY_MAPPER_WRAPPER_WRITE_VIRTUAL_MEMORY_UNSAFE = 3 //col:17
)

type MEMORY_MAPPER_ADDRESSES struct {
	PteVirtualAddress uint64 //col:3
	VirualAddress     uint64 //col:4
}

type PAGE_ENTRY struct {
	Union     union        //col:8
	Flags     uint64       //col:10
	Pml4      PML4E_64     //col:11
	PdptLarge PDPTE_1GB_64 //col:12
	Pdpt      PDPTE_64     //col:13
	PdLarge   PDE_2MB_64   //col:14
	Pd        PDE_64       //col:15
	Pt        PTE_64       //col:16
	Struct    struct                 //col:17
		Present               uint64 //col:19
		Write                 uint64 //col:20
		Supervisor            uint64 //col:21
		PageLevelWriteThrough uint64 //col:22
		PageLevelCacheDisable uint64 //col:23
		Accessed              uint64 //col:24
		Dirty                 uint64 //col:25
		LargePage             uint64 //col:26
		Global                uint64 //col:27
		Ignored1              uint64 //col:28
		PageFrameNumber       uint64 //col:29
		Reserved1             uint64 //col:30
		Ignored2              uint64 //col:31
		ProtectionKey         uint64 //col:32
		ExecuteDisable        uint64 //col:33
	}


	type CR3_TYPE struct{
	Union union   //col:39
	Flags uint64  //col:41
	Struct struct //col:42
	Pcid uint64            //col:44
	PageFrameNumber uint64 //col:45
	Reserved1 uint64       //col:46
	Reserved_2 uint64      //col:47
	PcidInvalidate uint64 //col:48
	}
