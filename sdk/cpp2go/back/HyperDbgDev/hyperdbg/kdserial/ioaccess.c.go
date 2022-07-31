package kdserial
type (
Ioaccess interface{
Copyright ()(ok bool)//col:25
ReadPort16()(ok bool)//col:32
ReadPort32()(ok bool)//col:39
WritePort8()(ok bool)//col:47
WritePort16()(ok bool)//col:55
WritePort32()(ok bool)//col:63
ReadRegister8()(ok bool)//col:70
ReadRegister16()(ok bool)//col:77
ReadRegister32()(ok bool)//col:84
WriteRegister8()(ok bool)//col:92
WriteRegister16()(ok bool)//col:100
WriteRegister32()(ok bool)//col:108
#if defined()(ok bool)//col:175
}

)
func NewIoaccess() { return & ioaccess{} }
func (i *ioaccess)Copyright ()(ok bool){//col:25
return true
}

func (i *ioaccess)ReadPort16()(ok bool){//col:32
return true
}

func (i *ioaccess)ReadPort32()(ok bool){//col:39
return true
}

func (i *ioaccess)WritePort8()(ok bool){//col:47
return true
}

func (i *ioaccess)WritePort16()(ok bool){//col:55
return true
}

func (i *ioaccess)WritePort32()(ok bool){//col:63
return true
}

func (i *ioaccess)ReadRegister8()(ok bool){//col:70
return true
}

func (i *ioaccess)ReadRegister16()(ok bool){//col:77
return true
}

func (i *ioaccess)ReadRegister32()(ok bool){//col:84
return true
}

func (i *ioaccess)WriteRegister8()(ok bool){//col:92
return true
}

func (i *ioaccess)WriteRegister16()(ok bool){//col:100
return true
}

func (i *ioaccess)WriteRegister32()(ok bool){//col:108
return true
}

func (i *ioaccess)#if defined()(ok bool){//col:175
return true
}

