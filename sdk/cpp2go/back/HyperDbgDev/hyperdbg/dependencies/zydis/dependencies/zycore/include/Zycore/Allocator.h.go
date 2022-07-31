package Zycore


const(
ZYCORE_ALLOCATOR_H =  //col:33
)

type (
Allocator interface{
typedef ZyanStatus ()(ok bool)//col:103
ZYCORE_EXPORT ZyanStatus ZyanAllocatorInit()(ok bool)//col:140
}

)

func NewAllocator() { return & allocator{} }

func (a *allocator)typedef ZyanStatus ()(ok bool){//col:103










return true
}


func (a *allocator)ZYCORE_EXPORT ZyanStatus ZyanAllocatorInit()(ok bool){//col:140






return true
}




