package kdserial


type (
Uartio interface{
#if defined()(ok bool)//col:55
ReadRegisterWithIndex8()(ok bool)//col:67
WriteRegisterWithIndex16()(ok bool)//col:81
ReadRegisterWithIndex16()(ok bool)//col:93
WriteRegisterWithIndex32()(ok bool)//col:107
ReadRegisterWithIndex32()(ok bool)//col:119
#if defined()(ok bool)//col:135
ReadRegisterWithIndex64()(ok bool)//col:147
#if defined()(ok bool)//col:169
ReadPortWithIndex8()(ok bool)//col:181
WritePortWithIndex16()(ok bool)//col:195
ReadPortWithIndex16()(ok bool)//col:207
WritePortWithIndex32()(ok bool)//col:221
ReadPortWithIndex32()(ok bool)//col:233
UartpSetAccess()(ok bool)//col:373
}

)

func NewUartio() { return & uartio{} }

func (u *uartio)#if defined()(ok bool){//col:55
















return true
}


func (u *uartio)ReadRegisterWithIndex8()(ok bool){//col:67








return true
}


func (u *uartio)WriteRegisterWithIndex16()(ok bool){//col:81










return true
}


func (u *uartio)ReadRegisterWithIndex16()(ok bool){//col:93








return true
}


func (u *uartio)WriteRegisterWithIndex32()(ok bool){//col:107










return true
}


func (u *uartio)ReadRegisterWithIndex32()(ok bool){//col:119








return true
}


func (u *uartio)#if defined()(ok bool){//col:135












return true
}


func (u *uartio)ReadRegisterWithIndex64()(ok bool){//col:147








return true
}


func (u *uartio)#if defined()(ok bool){//col:169












return true
}


func (u *uartio)ReadPortWithIndex8()(ok bool){//col:181








return true
}


func (u *uartio)WritePortWithIndex16()(ok bool){//col:195










return true
}


func (u *uartio)ReadPortWithIndex16()(ok bool){//col:207








return true
}


func (u *uartio)WritePortWithIndex32()(ok bool){//col:221










return true
}


func (u *uartio)ReadPortWithIndex32()(ok bool){//col:233








return true
}


func (u *uartio)UartpSetAccess()(ok bool){//col:373
















































































return true
}




