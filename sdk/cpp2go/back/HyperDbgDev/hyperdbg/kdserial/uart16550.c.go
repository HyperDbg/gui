package kdserial


type (
Uart16550 interface{
WritePortWithIndex8()(ok bool)//col:45
ReadPortWithIndex8()(ok bool)//col:57
KdHyperDbgTest()(ok bool)//col:84
KdHyperDbgPrepareDebuggeeConnectionPort()(ok bool)//col:97
KdHyperDbgSendByte()(ok bool)//col:104
KdHyperDbgRecvByte()(ok bool)//col:115
Uart16550SetBaud()(ok bool)//col:227
Uart16550LegacyInitializePort()(ok bool)//col:300
Uart16550InitializePort()(ok bool)//col:379
Uart16550MmInitializePort()(ok bool)//col:427
Uart16550SetBaudCommon()(ok bool)//col:506
Uart16550SetBaud()(ok bool)//col:539
Uart16550GetByte()(ok bool)//col:637
Uart16550PutByte()(ok bool)//col:745
Uart16550RxReady()(ok bool)//col:800
}















)

func NewUart16550() { return & uart16550{} }

func (u *uart16550)WritePortWithIndex8()(ok bool){//col:45










return true
}
















func (u *uart16550)ReadPortWithIndex8()(ok bool){//col:57








return true
}
















func (u *uart16550)KdHyperDbgTest()(ok bool){//col:84














return true
}
















func (u *uart16550)KdHyperDbgPrepareDebuggeeConnectionPort()(ok bool){//col:97









return true
}
















func (u *uart16550)KdHyperDbgSendByte()(ok bool){//col:104




return true
}
















func (u *uart16550)KdHyperDbgRecvByte()(ok bool){//col:115








return true
}
















func (u *uart16550)Uart16550SetBaud()(ok bool){//col:227





























return true
}
















func (u *uart16550)Uart16550LegacyInitializePort()(ok bool){//col:300


































return true
}
















func (u *uart16550)Uart16550InitializePort()(ok bool){//col:379



























return true
}
















func (u *uart16550)Uart16550MmInitializePort()(ok bool){//col:427














return true
}
















func (u *uart16550)Uart16550SetBaudCommon()(ok bool){//col:506
























return true
}
















func (u *uart16550)Uart16550SetBaud()(ok bool){//col:539










return true
}
















func (u *uart16550)Uart16550GetByte()(ok bool){//col:637

















































return true
}
















func (u *uart16550)Uart16550PutByte()(ok bool){//col:745



















































return true
}
















func (u *uart16550)Uart16550RxReady()(ok bool){//col:800



















return true
}


















