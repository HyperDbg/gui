package src

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

	return true
}

func (a *allocator) static_ZyanStatus_ZyanAllocatorDefaultReallocate() (ok bool) { //col:31

	return true
}

func (a *allocator) static_ZyanStatus_ZyanAllocatorDefaultDeallocate() (ok bool) { //col:54

	return true
}

func (a *allocator) ZyanAllocatorPtr_ZyanAllocatorDefault() (ok bool) { //col:64

	return true
}
