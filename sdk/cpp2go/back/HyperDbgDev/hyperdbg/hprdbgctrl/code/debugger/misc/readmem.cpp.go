package misc
type (
Readmem interface{
HyperDbgReadMemoryAndDisassemble()(ok bool)//col:169
ShowMemoryCommandDB()(ok bool)//col:237
ShowMemoryCommandDC()(ok bool)//col:306
ShowMemoryCommandDD()(ok bool)//col:356
ShowMemoryCommandDQ()(ok bool)//col:409
}

)
func NewReadmem() { return & readmem{} }
func (r *readmem)HyperDbgReadMemoryAndDisassemble()(ok bool){//col:169
return true
}

func (r *readmem)ShowMemoryCommandDB()(ok bool){//col:237
return true
}

func (r *readmem)ShowMemoryCommandDC()(ok bool){//col:306
return true
}

func (r *readmem)ShowMemoryCommandDD()(ok bool){//col:356
return true
}

func (r *readmem)ShowMemoryCommandDQ()(ok bool){//col:409
return true
}

