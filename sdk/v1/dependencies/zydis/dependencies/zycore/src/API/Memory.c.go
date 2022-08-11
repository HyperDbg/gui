package API

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\src\API\Memory.c.back

type (
	Memory interface {
		ZyanU32_ZyanMemoryGetSystemPageSize() (ok bool) //col:29
		ZyanStatus_ZyanMemoryVirtualFree() (ok bool)    //col:45
	}
	memory struct{}
)

func NewMemory() Memory { return &memory{} }

func (m *memory) ZyanU32_ZyanMemoryGetSystemPageSize() (ok bool) { //col:29
	/*
	   ZyanU32 ZyanMemoryGetSystemPageSize()
	   {
	   #if defined(ZYAN_WINDOWS)

	   	SYSTEM_INFO system_info;
	   	GetSystemInfo(&system_info);
	   	return sysconf(_SC_PAGE_SIZE);

	   ZyanU32 ZyanMemoryGetSystemAllocationGranularity()
	   {
	   #if defined(ZYAN_WINDOWS)

	   	SYSTEM_INFO system_info;
	   	GetSystemInfo(&system_info);
	   	return sysconf(_SC_PAGE_SIZE);

	   ZyanStatus ZyanMemoryVirtualProtect(void* address, ZyanUSize size,

	   	ZyanMemoryPageProtection protection)

	   {
	   #if defined(ZYAN_WINDOWS)

	   	DWORD old;
	   	if (!VirtualProtect(address, size, protection, &old))
	   	{
	   	    return ZYAN_STATUS_BAD_SYSTEMCALL;
	   	}

	   #elif defined(ZYAN_POSIX)

	   	if (mprotect(address, size, protection))
	   	{
	   	    return ZYAN_STATUS_BAD_SYSTEMCALL;
	   	}

	   #endif

	   	    return ZYAN_STATUS_SUCCESS;
	   	}
	*/
	return true
}

func (m *memory) ZyanStatus_ZyanMemoryVirtualFree() (ok bool) { //col:45
	/*
	   ZyanStatus ZyanMemoryVirtualFree(void* address, ZyanUSize size)
	   {
	   #if defined(ZYAN_WINDOWS)

	   	ZYAN_UNUSED(size);
	   	if (!VirtualFree(address, 0, MEM_RELEASE))
	   	{
	   	    return ZYAN_STATUS_BAD_SYSTEMCALL;
	   	}

	   #elif defined(ZYAN_POSIX)

	   	if (munmap(address, size))
	   	{
	   	    return ZYAN_STATUS_BAD_SYSTEMCALL;
	   	}

	   #endif

	   	    return ZYAN_STATUS_SUCCESS;
	   	}
	*/
	return true
}

