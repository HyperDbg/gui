package vmx


type (
Vmcall interface{
VmxHypervVmcallHandler()(ok bool)//col:56
VmxHandleVmcallVmExit()(ok bool)//col:105
VmxVmcallHandler()(ok bool)//col:488
VmcallTest()(ok bool)//col:519
}







































)

func NewVmcall() { return & vmcall{} }

func (v *vmcall)VmxHypervVmcallHandler()(ok bool){//col:56























return true
}








































func (v *vmcall)VmxHandleVmcallVmExit()(ok bool){//col:105



























return true
}








































func (v *vmcall)VmxVmcallHandler()(ok bool){//col:488





































































































































































































































































































































return true
}








































func (v *vmcall)VmcallTest()(ok bool){//col:519













return true
}










































