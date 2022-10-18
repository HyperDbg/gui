package memory
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbghv\header\memory\MemoryMapper.h.back

const(
    PagingLevelPageTable  =  0  //col:3
    PagingLevelPageDirectory = 2  //col:4
    PagingLevelPageDirectoryPointerTable = 3  //col:5
    PagingLevelPageMapLevel4 = 4  //col:6
)


const(
    MEMORY_MAPPER_WRAPPER_READ_PHYSICAL_MEMORY = 1  //col:10
    MEMORY_MAPPER_WRAPPER_READ_VIRTUAL_MEMORY = 2  //col:11
)


const(
    MEMORY_MAPPER_WRAPPER_WRITE_PHYSICAL_MEMORY = 1  //col:15
    MEMORY_MAPPER_WRAPPER_WRITE_VIRTUAL_MEMORY_SAFE = 2  //col:16
    MEMORY_MAPPER_WRAPPER_WRITE_VIRTUAL_MEMORY_UNSAFE = 3  //col:17
)



type  _MEMORY_MAPPER_ADDRESSES struct{
PteVirtualAddress uint64 //col:7
VirualAddress uint64 //col:8
}


type  _PAGE_ENTRY struct{
Union union //col:36
Flags uint64 //col:38
Pml4 PML4E_64 //col:39
PdptLarge PDPTE_1GB_64 //col:40
Pdpt PDPTE_64 //col:41
PdLarge PDE_2MB_64 //col:42
Pd PDE_64 //col:43
Pt PTE_64 //col:44
Struct struct //col:45
Present uint64 //col:47
Write uint64 //col:48
Supervisor uint64 //col:49
PageLevelWriteThrough uint64 //col:50
PageLevelCacheDisable uint64 //col:51
Accessed uint64 //col:52
Dirty uint64 //col:53
LargePage uint64 //col:54
Global uint64 //col:55
Ignored1 uint64 //col:56
PageFrameNumber uint64 //col:57
Reserved1 uint64 //col:58
Ignored2 uint64 //col:59
ProtectionKey uint64 //col:60
ExecuteDisable uint64 //col:61
}


type  _CR3_TYPE struct{
Union union //col:51
Flags uint64 //col:53
Struct struct //col:54
Pcid uint64 //col:56
PageFrameNumber uint64 //col:57
Reserved1 uint64 //col:58
Reserved_2 uint64 //col:59
PcidInvalidate uint64 //col:60
}




