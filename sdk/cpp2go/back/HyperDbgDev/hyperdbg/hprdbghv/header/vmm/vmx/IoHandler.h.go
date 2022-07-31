package vmx


const(
typedef enum = 1  //col:26
    AccessOut  =  0  //col:28
    AccessIn   =  1  //col:29
)


const(
typedef enum = 1  //col:36
    OpEncodingDx   =  0  //col:38
    OpEncodingImm  =  1  //col:39
)



type (
IoHandler interface{
__inbyte()(ok bool)//col:54
__inword()(ok bool)//col:65
__indword()(ok bool)//col:76
__inbytestring()(ok bool)//col:87
__inwordstring()(ok bool)//col:98
__indwordstring()(ok bool)//col:109
__outbyte()(ok bool)//col:120
__outword()(ok bool)//col:131
__outdword()(ok bool)//col:142
__outbytestring()(ok bool)//col:153
__outwordstring()(ok bool)//col:164
__outdwordstring()(ok bool)//col:175
}






)

func NewIoHandler() { return & ioHandler{} }

func (i *ioHandler)__inbyte()(ok bool){//col:54







return true
}







func (i *ioHandler)__inword()(ok bool){//col:65







return true
}







func (i *ioHandler)__indword()(ok bool){//col:76







return true
}







func (i *ioHandler)__inbytestring()(ok bool){//col:87







return true
}







func (i *ioHandler)__inwordstring()(ok bool){//col:98







return true
}







func (i *ioHandler)__indwordstring()(ok bool){//col:109







return true
}







func (i *ioHandler)__outbyte()(ok bool){//col:120







return true
}







func (i *ioHandler)__outword()(ok bool){//col:131







return true
}







func (i *ioHandler)__outdword()(ok bool){//col:142







return true
}







func (i *ioHandler)__outbytestring()(ok bool){//col:153







return true
}







func (i *ioHandler)__outwordstring()(ok bool){//col:164







return true
}







func (i *ioHandler)__outdwordstring()(ok bool){//col:175







return true
}









