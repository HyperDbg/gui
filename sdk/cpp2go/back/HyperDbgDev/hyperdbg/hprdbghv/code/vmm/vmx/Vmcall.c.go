package vmx
type (
	Vmcall interface {
		VmxHypervVmcallHandler() (ok bool) //col:56
		VmxHandleVmcallVmExit() (ok bool)  //col:104
		VmxVmcallHandler() (ok bool)       //col:486
		VmcallTest() (ok bool)             //col:516
	}
)
func NewVmcall() { return &vmcall{} }
func (v *vmcall) VmxHypervVmcallHandler() (ok bool) { //col:56
	return true
}

func (v *vmcall) VmxHandleVmcallVmExit() (ok bool) { //col:104
	return true
}

func (v *vmcall) VmxVmcallHandler() (ok bool) { //col:486
	return true
}

func (v *vmcall) VmcallTest() (ok bool) { //col:516
	return true
}

