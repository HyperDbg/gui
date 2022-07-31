package src
type (
Allocator interface{
  Zyan Core Library ()(ok bool)//col:57
static ZyanStatus ZyanAllocatorDefaultReallocate()(ok bool)//col:77
static ZyanStatus ZyanAllocatorDefaultDeallocate()(ok bool)//col:94
ZyanStatus ZyanAllocatorInit()(ok bool)//col:117
ZyanAllocator* ZyanAllocatorDefault()(ok bool)//col:130
}

)
func NewAllocator() { return & allocator{} }
func (a *allocator)  Zyan Core Library ()(ok bool){//col:57
return true
}

func (a *allocator)static ZyanStatus ZyanAllocatorDefaultReallocate()(ok bool){//col:77
return true
}

func (a *allocator)static ZyanStatus ZyanAllocatorDefaultDeallocate()(ok bool){//col:94
return true
}

func (a *allocator)ZyanStatus ZyanAllocatorInit()(ok bool){//col:117
return true
}

func (a *allocator)ZyanAllocator* ZyanAllocatorDefault()(ok bool){//col:130
return true
}

