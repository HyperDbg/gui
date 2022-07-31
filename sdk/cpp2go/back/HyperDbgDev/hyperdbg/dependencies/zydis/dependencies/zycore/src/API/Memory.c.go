package API
//back\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\src\API\Memory.c.back

type (
Memory interface{
  Zyan Core Library ()(ok bool)//col:59
ZyanU32 ZyanMemoryGetSystemAllocationGranularity()(ok bool)//col:75
ZyanStatus ZyanMemoryVirtualProtect()(ok bool)//col:102
ZyanStatus ZyanMemoryVirtualFree()(ok bool)//col:124
}
)

func NewMemory() { return & memory{} }

func (m *memory)  Zyan Core Library ()(ok bool){//col:59
/*  Zyan Core Library (Zycore-C)
  Original Author : Florian Bernd
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
#include <Zycore/API/Memory.h>
#if   defined(ZYAN_WINDOWS)
#elif defined(ZYAN_POSIX)
#   include <unistd.h>
#else
#   error "Unsupported platform detected"
#endif
ZyanU32 ZyanMemoryGetSystemPageSize()
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

func (m *memory)ZyanU32 ZyanMemoryGetSystemAllocationGranularity()(ok bool){//col:75
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

func (m *memory)ZyanStatus ZyanMemoryVirtualProtect()(ok bool){//col:102
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

func (m *memory)ZyanStatus ZyanMemoryVirtualFree()(ok bool){//col:124
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


