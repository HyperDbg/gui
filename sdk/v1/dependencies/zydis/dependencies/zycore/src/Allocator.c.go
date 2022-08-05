package src

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\zydis\dependencies\zycore\src\Allocator.c.back

type (
	Allocator interface {
		static_ZyanStatus_ZyanAllocatorDefaultAllocate() (ok bool)   //col:15
		static_ZyanStatus_ZyanAllocatorDefaultReallocate() (ok bool) //col:31
		static_ZyanStatus_ZyanAllocatorDefaultDeallocate() (ok bool) //col:54
		ZyanAllocatorPtr_ZyanAllocatorDefault() (ok bool)            //col:64
	}
	allocator struct{}
)

func NewAllocator() Allocator { return &allocator{} }

func (a *allocator) static_ZyanStatus_ZyanAllocatorDefaultAllocate() (ok bool) { //col:15
	/*
	   static ZyanStatus ZyanAllocatorDefaultAllocate(ZyanAllocator* allocator, void** p,

	   	ZyanUSize element_size, ZyanUSize n)

	   	{
	   	    ZYAN_ASSERT(allocator);
	   	    ZYAN_ASSERT(p);
	   	    ZYAN_ASSERT(element_size);
	   	    ZYAN_ASSERT(n);
	   	    ZYAN_UNUSED(allocator);
	   	    *p = ZYAN_MALLOC(element_size * n);
	   	    if (!*p)
	   	    {
	   	        return ZYAN_STATUS_NOT_ENOUGH_MEMORY;
	   	    }
	   	    return ZYAN_STATUS_SUCCESS;
	   	}
	*/
	return true
}

func (a *allocator) static_ZyanStatus_ZyanAllocatorDefaultReallocate() (ok bool) { //col:31
	/*
	   static ZyanStatus ZyanAllocatorDefaultReallocate(ZyanAllocator* allocator, void** p,

	   	ZyanUSize element_size, ZyanUSize n)

	   	{
	   	    ZYAN_ASSERT(allocator);
	   	    ZYAN_ASSERT(p);
	   	    ZYAN_ASSERT(element_size);
	   	    ZYAN_ASSERT(n);
	   	    ZYAN_UNUSED(allocator);
	   	    void* const x = ZYAN_REALLOC(*p, element_size * n);
	   	    if (!x)
	   	    {
	   	        return ZYAN_STATUS_NOT_ENOUGH_MEMORY;
	   	    }
	   	    *p = x;
	   	    return ZYAN_STATUS_SUCCESS;
	   	}
	*/
	return true
}

func (a *allocator) static_ZyanStatus_ZyanAllocatorDefaultDeallocate() (ok bool) { //col:54
	/*
	   static ZyanStatus ZyanAllocatorDefaultDeallocate(ZyanAllocator* allocator, void* p,

	   	ZyanUSize element_size, ZyanUSize n)

	   	{
	   	    ZYAN_ASSERT(allocator);
	   	    ZYAN_ASSERT(p);
	   	    ZYAN_ASSERT(element_size);
	   	    ZYAN_ASSERT(n);
	   	    ZYAN_UNUSED(allocator);
	   	    ZYAN_UNUSED(element_size);
	   	    ZYAN_UNUSED(n);
	   	    ZYAN_FREE(p);

	   ZyanStatus ZyanAllocatorInit(ZyanAllocator* allocator, ZyanAllocatorAllocate allocate,

	   	ZyanAllocatorAllocate reallocate, ZyanAllocatorDeallocate deallocate)

	   	{
	   	    if (!allocator || !allocate || !reallocate || !deallocate)
	   	    {
	   	        return ZYAN_STATUS_INVALID_ARGUMENT;
	   	    }
	   	    allocator->allocate   = allocate;
	   	    allocator->reallocate = reallocate;
	   	    allocator->deallocate = deallocate;
	   	    return ZYAN_STATUS_SUCCESS;
	   	}
	*/
	return true
}

func (a *allocator) ZyanAllocatorPtr_ZyanAllocatorDefault() (ok bool) { //col:64
	/*
	   ZyanAllocator* ZyanAllocatorDefault(void)

	   	{
	   	    static ZyanAllocator allocator =
	   	    {
	   	        &ZyanAllocatorDefaultAllocate,
	   	        &ZyanAllocatorDefaultReallocate,
	   	        &ZyanAllocatorDefaultDeallocate
	   	    };
	   	    return &allocator;
	   	}
	*/
	return true
}

