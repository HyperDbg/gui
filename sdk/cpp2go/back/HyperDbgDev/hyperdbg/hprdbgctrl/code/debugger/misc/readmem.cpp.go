package misc


type (
Readmem interface{
HyperDbgReadMemoryAndDisassemble()(ok bool)//col:169
ShowMemoryCommandDB()(ok bool)//col:238
ShowMemoryCommandDC()(ok bool)//col:308
ShowMemoryCommandDD()(ok bool)//col:359
ShowMemoryCommandDQ()(ok bool)//col:413
}

)

func NewReadmem() { return & readmem{} }

func (r *readmem)HyperDbgReadMemoryAndDisassemble()(ok bool){//col:169













































































return true
}


func (r *readmem)ShowMemoryCommandDB()(ok bool){//col:238





































return true
}


func (r *readmem)ShowMemoryCommandDC()(ok bool){//col:308






































return true
}


func (r *readmem)ShowMemoryCommandDD()(ok bool){//col:359

























return true
}


func (r *readmem)ShowMemoryCommandDQ()(ok bool){//col:413


























return true
}




