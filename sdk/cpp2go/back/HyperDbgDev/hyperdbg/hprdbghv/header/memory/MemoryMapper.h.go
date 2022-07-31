package memory


const(
PAGE_4KB_OFFSET = ((UINT64)(1 << 12) - 1) //col:21
PAGE_2MB_OFFSET = ((UINT64)(1 << 21) - 1) //col:22
PAGE_4MB_OFFSET = ((UINT64)(1 << 22) - 1) //col:23
PAGE_1GB_OFFSET = ((UINT64)(1 << 30) - 1) //col:24
)

const(
    PagingLevelPageTable  =  0  //col:36
    PagingLevelPageDirectory = 2  //col:37
    PagingLevelPageDirectoryPointerTable = 3  //col:38
    PagingLevelPageMapLevel4 = 4  //col:39
)


const(
    MEMORY_MAPPER_WRAPPER_READ_PHYSICAL_MEMORY = 1  //col:48
    MEMORY_MAPPER_WRAPPER_READ_VIRTUAL_MEMORY = 2  //col:49
)


const(
    MEMORY_MAPPER_WRAPPER_WRITE_PHYSICAL_MEMORY = 1  //col:58
    MEMORY_MAPPER_WRAPPER_WRITE_VIRTUAL_MEMORY_SAFE = 2  //col:59
    MEMORY_MAPPER_WRAPPER_WRITE_VIRTUAL_MEMORY_UNSAFE = 3  //col:60
)




