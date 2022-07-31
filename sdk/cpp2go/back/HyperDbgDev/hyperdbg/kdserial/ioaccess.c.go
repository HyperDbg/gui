package kdserial


type (
Ioaccess interface{
#if defined()(ok bool)//col:25
ReadPort16()(ok bool)//col:33
ReadPort32()(ok bool)//col:41
WritePort8()(ok bool)//col:50
WritePort16()(ok bool)//col:59
WritePort32()(ok bool)//col:68
ReadRegister8()(ok bool)//col:76
ReadRegister16()(ok bool)//col:84
ReadRegister32()(ok bool)//col:92
WriteRegister8()(ok bool)//col:101
WriteRegister16()(ok bool)//col:110
WriteRegister32()(ok bool)//col:119
#if defined()(ok bool)//col:187
}






)

func NewIoaccess() { return & ioaccess{} }

func (i *ioaccess)#if defined()(ok bool){//col:25







return true
}







func (i *ioaccess)ReadPort16()(ok bool){//col:33





return true
}







func (i *ioaccess)ReadPort32()(ok bool){//col:41





return true
}







func (i *ioaccess)WritePort8()(ok bool){//col:50






return true
}







func (i *ioaccess)WritePort16()(ok bool){//col:59






return true
}







func (i *ioaccess)WritePort32()(ok bool){//col:68






return true
}







func (i *ioaccess)ReadRegister8()(ok bool){//col:76





return true
}







func (i *ioaccess)ReadRegister16()(ok bool){//col:84





return true
}







func (i *ioaccess)ReadRegister32()(ok bool){//col:92





return true
}







func (i *ioaccess)WriteRegister8()(ok bool){//col:101






return true
}







func (i *ioaccess)WriteRegister16()(ok bool){//col:110






return true
}







func (i *ioaccess)WriteRegister32()(ok bool){//col:119






return true
}







func (i *ioaccess)#if defined()(ok bool){//col:187













































return true
}









