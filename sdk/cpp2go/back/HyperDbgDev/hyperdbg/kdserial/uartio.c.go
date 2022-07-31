package kdserial
type (
Uartio interface{
Copyright ()(ok bool)//col:55
ReadRegisterWithIndex8()(ok bool)//col:66
WriteRegisterWithIndex16()(ok bool)//col:79
ReadRegisterWithIndex16()(ok bool)//col:90
WriteRegisterWithIndex32()(ok bool)//col:103
ReadRegisterWithIndex32()(ok bool)//col:114
#if defined()(ok bool)//col:129
ReadRegisterWithIndex64()(ok bool)//col:140
#if defined()(ok bool)//col:161
ReadPortWithIndex8()(ok bool)//col:172
WritePortWithIndex16()(ok bool)//col:185
ReadPortWithIndex16()(ok bool)//col:196
WritePortWithIndex32()(ok bool)//col:209
ReadPortWithIndex32()(ok bool)//col:220
UartpSetAccess()(ok bool)//col:359
}

)
func NewUartio() { return & uartio{} }
func (u *uartio)Copyright ()(ok bool){//col:55
return true
}

func (u *uartio)ReadRegisterWithIndex8()(ok bool){//col:66
return true
}

func (u *uartio)WriteRegisterWithIndex16()(ok bool){//col:79
return true
}

func (u *uartio)ReadRegisterWithIndex16()(ok bool){//col:90
return true
}

func (u *uartio)WriteRegisterWithIndex32()(ok bool){//col:103
return true
}

func (u *uartio)ReadRegisterWithIndex32()(ok bool){//col:114
return true
}

func (u *uartio)#if defined()(ok bool){//col:129
return true
}

func (u *uartio)ReadRegisterWithIndex64()(ok bool){//col:140
return true
}

func (u *uartio)#if defined()(ok bool){//col:161
return true
}

func (u *uartio)ReadPortWithIndex8()(ok bool){//col:172
return true
}

func (u *uartio)WritePortWithIndex16()(ok bool){//col:185
return true
}

func (u *uartio)ReadPortWithIndex16()(ok bool){//col:196
return true
}

func (u *uartio)WritePortWithIndex32()(ok bool){//col:209
return true
}

func (u *uartio)ReadPortWithIndex32()(ok bool){//col:220
return true
}

func (u *uartio)UartpSetAccess()(ok bool){//col:359
return true
}

