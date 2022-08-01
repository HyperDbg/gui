package memory
//back\HyperDbgDev\hyperdbg\hprdbghv\header\memory\MemoryMapper.h.back

const(
PAGE_4KB_OFFSET = ((UINT64)(1 << 12) - 1) //col:1
PAGE_2MB_OFFSET = ((UINT64)(1 << 21) - 1) //col:2
PAGE_4MB_OFFSET = ((UINT64)(1 << 22) - 1) //col:3
PAGE_1GB_OFFSET = ((UINT64)(1 << 30) - 1) //col:4
)

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



type MEMORY_MAPPER_ADDRESSES struct{
PteVirtualAddress uint64
VirualAddress uint64
}


type PAGE_ENTRY struct{
Union union
Flags uint64
Pml4 PML4E_64
PdptLarge PDPTE_1GB_64
Pdpt PDPTE_64
PdLarge PDE_2MB_64
Pd PDE_64
Pt PTE_64
Struct struct
Present uint64
Write uint64
Supervisor uint64
PageLevelWriteThrough uint64
PageLevelCacheDisable uint64
Accessed uint64
Dirty uint64
LargePage uint64
Global uint64
Ignored1 uint64
PageFrameNumber uint64
Reserved1 uint64
Ignored2 uint64
ProtectionKey uint64
ExecuteDisable uint64
}


type CR3_TYPE struct{
Union union
Flags uint64
Struct struct
Pcid uint64
PageFrameNumber uint64
Reserved1 uint64
Reserved_2 uint64
PcidInvalidate uint64
}




