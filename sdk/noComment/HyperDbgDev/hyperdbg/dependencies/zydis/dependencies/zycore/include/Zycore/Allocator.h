#ifndef ZYCORE_ALLOCATOR_H
#define ZYCORE_ALLOCATOR_H
#include <ZycoreExportConfig.h>
#include <Zycore/Status.h>
#include <Zycore/Types.h>
#ifdef __cplusplus
extern "C" {
#endif
struct ZyanAllocator_;
typedef ZyanStatus (*ZyanAllocatorAllocate)(struct ZyanAllocator_* allocator, void** p,
    ZyanUSize element_size, ZyanUSize n);
typedef ZyanStatus (*ZyanAllocatorDeallocate)(struct ZyanAllocator_* allocator, void* p,
    ZyanUSize element_size, ZyanUSize n);
typedef struct ZyanAllocator_
{
    ZyanAllocatorAllocate allocate;
    ZyanAllocatorAllocate reallocate;
    ZyanAllocatorDeallocate deallocate;
} ZyanAllocator;
ZYCORE_EXPORT ZyanStatus ZyanAllocatorInit(ZyanAllocator* allocator, ZyanAllocatorAllocate allocate,
    ZyanAllocatorAllocate reallocate, ZyanAllocatorDeallocate deallocate);
#ifndef ZYAN_NO_LIBC
ZYCORE_EXPORT ZYAN_REQUIRES_LIBC ZyanAllocator* ZyanAllocatorDefault(void);
#endif 
#ifdef __cplusplus
}

#endif
#endif 
