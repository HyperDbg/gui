package src


type (
Allocator interface{
static ZyanStatus ZyanAllocatorDefaultAllocate()(ok bool)//col:57
static ZyanStatus ZyanAllocatorDefaultReallocate()(ok bool)//col:78
static ZyanStatus ZyanAllocatorDefaultDeallocate()(ok bool)//col:96
ZyanStatus ZyanAllocatorInit()(ok bool)//col:120
ZyanAllocator* ZyanAllocatorDefault()(ok bool)//col:134
}

)

func NewAllocator() { return & allocator{} }

func (a *allocator)static ZyanStatus ZyanAllocatorDefaultAllocate()(ok bool){//col:57















return true
}


func (a *allocator)static ZyanStatus ZyanAllocatorDefaultReallocate()(ok bool){//col:78
















return true
}


func (a *allocator)static ZyanStatus ZyanAllocatorDefaultDeallocate()(ok bool){//col:96













return true
}


func (a *allocator)ZyanStatus ZyanAllocatorInit()(ok bool){//col:120












return true
}


func (a *allocator)ZyanAllocator* ZyanAllocatorDefault()(ok bool){//col:134










return true
}




