#include <Zycore/Allocator.h>
#include <Zycore/LibC.h>
#ifndef ZYAN_NO_LIBC
static ZyanStatus ZyanAllocatorDefaultAllocate(ZyanAllocator *allocator,
                                               void **p, ZyanUSize element_size,
                                               ZyanUSize n) {
  ZYAN_ASSERT(allocator);
  ZYAN_ASSERT(p);
  ZYAN_ASSERT(element_size);
  ZYAN_ASSERT(n);
  ZYAN_UNUSED(allocator);
  *p = ZYAN_MALLOC(element_size * n);
  if (!*p) {
    return ZYAN_STATUS_NOT_ENOUGH_MEMORY;
  }
  return ZYAN_STATUS_SUCCESS;
}

static ZyanStatus ZyanAllocatorDefaultReallocate(ZyanAllocator *allocator,
                                                 void **p,
                                                 ZyanUSize element_size,
                                                 ZyanUSize n) {
  ZYAN_ASSERT(allocator);
  ZYAN_ASSERT(p);
  ZYAN_ASSERT(element_size);
  ZYAN_ASSERT(n);
  ZYAN_UNUSED(allocator);
  void *const x = ZYAN_REALLOC(*p, element_size * n);
  if (!x) {
    return ZYAN_STATUS_NOT_ENOUGH_MEMORY;
  }
  *p = x;
  return ZYAN_STATUS_SUCCESS;
}

static ZyanStatus ZyanAllocatorDefaultDeallocate(ZyanAllocator *allocator,
                                                 void *p,
                                                 ZyanUSize element_size,
                                                 ZyanUSize n) {
  ZYAN_ASSERT(allocator);
  ZYAN_ASSERT(p);
  ZYAN_ASSERT(element_size);
  ZYAN_ASSERT(n);
  ZYAN_UNUSED(allocator);
  ZYAN_UNUSED(element_size);
  ZYAN_UNUSED(n);
  ZYAN_FREE(p);
  return ZYAN_STATUS_SUCCESS;
}

#endif 
ZyanStatus ZyanAllocatorInit(ZyanAllocator *allocator,
                             ZyanAllocatorAllocate allocate,
                             ZyanAllocatorAllocate reallocate,
                             ZyanAllocatorDeallocate deallocate) {
  if (!allocator || !allocate || !reallocate || !deallocate) {
    return ZYAN_STATUS_INVALID_ARGUMENT;
  }
  allocator->allocate = allocate;
  allocator->reallocate = reallocate;
  allocator->deallocate = deallocate;
  return ZYAN_STATUS_SUCCESS;
}

#ifndef ZYAN_NO_LIBC
ZyanAllocator *ZyanAllocatorDefault(void) {
  static ZyanAllocator allocator = {&ZyanAllocatorDefaultAllocate,
                                    &ZyanAllocatorDefaultReallocate,
                                    &ZyanAllocatorDefaultDeallocate};
  return &allocator;
}

#endif
