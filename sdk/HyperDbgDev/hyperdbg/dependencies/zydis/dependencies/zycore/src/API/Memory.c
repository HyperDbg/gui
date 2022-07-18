#include <Zycore/API/Memory.h>
#if defined(ZYAN_WINDOWS)
#elif defined(ZYAN_POSIX)
#    include <unistd.h>
#else
#    error "Unsupported platform detected"
#endif
ZyanU32
ZyanMemoryGetSystemPageSize() {
#if defined(ZYAN_WINDOWS)
    SYSTEM_INFO system_info;
    GetSystemInfo(&system_info);
    return system_info.dwPageSize;
#elif defined(ZYAN_POSIX)
    return sysconf(_SC_PAGE_SIZE);
#endif
}

ZyanU32
ZyanMemoryGetSystemAllocationGranularity() {
#if defined(ZYAN_WINDOWS)
    SYSTEM_INFO system_info;
    GetSystemInfo(&system_info);
    return system_info.dwAllocationGranularity;
#elif defined(ZYAN_POSIX)
    return sysconf(_SC_PAGE_SIZE);
#endif
}

ZyanStatus
ZyanMemoryVirtualProtect(void * address, ZyanUSize size, ZyanMemoryPageProtection protection) {
#if defined(ZYAN_WINDOWS)
    DWORD old;
    if (!VirtualProtect(address, size, protection, &old)) {
        return ZYAN_STATUS_BAD_SYSTEMCALL;
    }
#elif defined(ZYAN_POSIX)
    if (mprotect(address, size, protection)) {
        return ZYAN_STATUS_BAD_SYSTEMCALL;
    }
#endif
    return ZYAN_STATUS_SUCCESS;
}

ZyanStatus
ZyanMemoryVirtualFree(void * address, ZyanUSize size) {
#if defined(ZYAN_WINDOWS)
    ZYAN_UNUSED(size);
    if (!VirtualFree(address, 0, MEM_RELEASE)) {
        return ZYAN_STATUS_BAD_SYSTEMCALL;
    }
#elif defined(ZYAN_POSIX)
    if (munmap(address, size)) {
        return ZYAN_STATUS_BAD_SYSTEMCALL;
    }
#endif
    return ZYAN_STATUS_SUCCESS;
}
