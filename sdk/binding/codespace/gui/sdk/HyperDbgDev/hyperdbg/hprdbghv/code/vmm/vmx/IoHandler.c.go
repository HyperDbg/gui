package vmx



type (
	IoHandler interface {
		IoHandleIoVmExits() (ok bool) //col:73
	}
	ioHandler struct{}
)

func NewIoHandler() IoHandler { return &ioHandler{} }

func (i *ioHandler) IoHandleIoVmExits() (ok bool) { //col:73












































































	return true
}


