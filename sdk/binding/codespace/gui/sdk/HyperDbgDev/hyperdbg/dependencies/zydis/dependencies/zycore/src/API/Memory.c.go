package API
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\src\API\Memory.c.back

type (
Memory interface{
ZyanU32_ZyanMemoryGetSystemPageSize()(ok bool)//col:10
ZyanU32_ZyanMemoryGetSystemAllocationGranularity()(ok bool)//col:20
ZyanStatus_ZyanMemoryVirtualProtect()(ok bool)//col:37
ZyanStatus_ZyanMemoryVirtualFree()(ok bool)//col:53
}
)

func NewMemory() { return & memory{} }

func (m *memory)ZyanU32_ZyanMemoryGetSystemPageSize()(ok bool){//col:10
/*ZyanU32 ZyanMemoryGetSystemPageSize()
{
#if defined(ZYAN_WINDOWS)
    SYSTEM_INFO system_info;
    GetSystemInfo(&system_info);
    return system_info.dwPageSize;
#elif defined(ZYAN_POSIX)
    return sysconf(_SC_PAGE_SIZE);
#endif
}*/
return true
}

func (m *memory)ZyanU32_ZyanMemoryGetSystemAllocationGranularity()(ok bool){//col:20
/*ZyanU32 ZyanMemoryGetSystemAllocationGranularity()
{
#if defined(ZYAN_WINDOWS)
    SYSTEM_INFO system_info;
    GetSystemInfo(&system_info);
    return system_info.dwAllocationGranularity;
#elif defined(ZYAN_POSIX)
    return sysconf(_SC_PAGE_SIZE);
#endif
}*/
return true
}

func (m *memory)ZyanStatus_ZyanMemoryVirtualProtect()(ok bool){//col:37
/*ZyanStatus ZyanMemoryVirtualProtect(void* address, ZyanUSize size, 
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
}*/
return true
}

func (m *memory)ZyanStatus_ZyanMemoryVirtualFree()(ok bool){//col:53
/*ZyanStatus ZyanMemoryVirtualFree(void* address, ZyanUSize size)
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
}*/
return true
}



