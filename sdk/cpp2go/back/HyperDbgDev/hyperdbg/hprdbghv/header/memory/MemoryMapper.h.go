package memory
const (
	PAGE_4KB_OFFSET = ((UINT64)(1<<12) - 1) //col:21
	PAGE_2MB_OFFSET = ((UINT64)(1<<21) - 1) //col:22
	PAGE_4MB_OFFSET = ((UINT64)(1<<22) - 1) //col:23
	PAGE_1GB_OFFSET = ((UINT64)(1<<30) - 1) //col:24
)
type PagingLevelPageTable =
0 uint32
const (
	PagingLevelPageTable                 PAGING_LEVEL = 0 //col:36
	PagingLevelPageDirectory             PAGING_LEVEL = 2 //col:37
	PagingLevelPageDirectoryPointerTable PAGING_LEVEL = 3 //col:38
	PagingLevelPageMapLevel4             PAGING_LEVEL = 4 //col:39
)
type MEMORY_MAPPER_WRAPPER_READ_PHYSICAL_MEMORY uint32
const (
	MEMORY_MAPPER_WRAPPER_READ_PHYSICAL_MEMORY MEMORY_MAPPER_WRAPPER_FOR_MEMORY_READ = 1 //col:48
	MEMORY_MAPPER_WRAPPER_READ_VIRTUAL_MEMORY  MEMORY_MAPPER_WRAPPER_FOR_MEMORY_READ = 2 //col:49
)
type MEMORY_MAPPER_WRAPPER_WRITE_PHYSICAL_MEMORY uint32
const (
	MEMORY_MAPPER_WRAPPER_WRITE_PHYSICAL_MEMORY       MEMORY_MAPPER_WRAPPER_FOR_MEMORY_WRITE = 1 //col:58
	MEMORY_MAPPER_WRAPPER_WRITE_VIRTUAL_MEMORY_SAFE   MEMORY_MAPPER_WRAPPER_FOR_MEMORY_WRITE = 2 //col:59
	MEMORY_MAPPER_WRAPPER_WRITE_VIRTUAL_MEMORY_UNSAFE MEMORY_MAPPER_WRAPPER_FOR_MEMORY_WRITE = 3 //col:60
)
type (
	MemoryMapper interface{
#define PAGE_4KB_OFFSET ()
(ok bool) //col:40
}

)
func NewMemoryMapper() { return &memoryMapper{} }
func (m *memoryMapper) #define PAGE_4KB_OFFSET ()(ok bool) { //col:40
	return true
}

