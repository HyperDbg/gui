package vmx


type (
IoHandler interface{
IoHandleIoVmExits()(ok bool)//col:202
IoHandleIoVmExitsAndDisassemble()(ok bool)//col:250
IoHandleSetIoBitmap()(ok bool)//col:279
IoHandlePerformIoBitmapChange()(ok bool)//col:309
IoHandlePerformIoBitmapReset()(ok bool)//col:327
}






)

func NewIoHandler() { return & ioHandler{} }

func (i *ioHandler)IoHandleIoVmExits()(ok bool){//col:202













































































































return true
}







func (i *ioHandler)IoHandleIoVmExitsAndDisassemble()(ok bool){//col:250














return true
}







func (i *ioHandler)IoHandleSetIoBitmap()(ok bool){//col:279

















return true
}







func (i *ioHandler)IoHandlePerformIoBitmapChange()(ok bool){//col:309














return true
}







func (i *ioHandler)IoHandlePerformIoBitmapReset()(ok bool){//col:327







return true
}









