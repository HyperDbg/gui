package platform


type (
CrossApi interface{
CrsAllocateContiguousZeroedMemory()(ok bool)//col:31
CrsAllocateNonPagedPool()(ok bool)//col:42
}






)

func NewCrossApi() { return & crossApi{} }

func (c *crossApi)CrsAllocateContiguousZeroedMemory()(ok bool){//col:31









return true
}







func (c *crossApi)CrsAllocateNonPagedPool()(ok bool){//col:42







return true
}









