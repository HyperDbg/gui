package misc


type (
Callstack interface{
CallstackReturnAddressToCallingAddress()(ok bool)//col:199
CallstackShowFrames()(ok bool)//col:329
}

)

func NewCallstack() { return & callstack{} }

func (c *callstack)CallstackReturnAddressToCallingAddress()(ok bool){//col:199










































































return true
}


func (c *callstack)CallstackShowFrames()(ok bool){//col:329
























































































return true
}




