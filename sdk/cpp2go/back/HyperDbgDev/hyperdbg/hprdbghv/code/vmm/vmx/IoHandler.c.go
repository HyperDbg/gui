package vmx
type (
	IoHandler interface {
		IoHandleIoVmExits() (ok bool)               //col:202
		IoHandleIoVmExitsAndDisassemble() (ok bool) //col:249
		IoHandleSetIoBitmap() (ok bool)             //col:277
		IoHandlePerformIoBitmapChange() (ok bool)   //col:306
		IoHandlePerformIoBitmapReset() (ok bool)    //col:323
	}
)
func NewIoHandler() { return &ioHandler{} }
func (i *ioHandler) IoHandleIoVmExits() (ok bool) { //col:202
	return true
}

func (i *ioHandler) IoHandleIoVmExitsAndDisassemble() (ok bool) { //col:249
	return true
}

func (i *ioHandler) IoHandleSetIoBitmap() (ok bool) { //col:277
	return true
}

func (i *ioHandler) IoHandlePerformIoBitmapChange() (ok bool) { //col:306
	return true
}

func (i *ioHandler) IoHandlePerformIoBitmapReset() (ok bool) { //col:323
	return true
}

