package vmx


type (
VmxBroadcast interface{
VmxBroadcastHandleNmiCallback()(ok bool)//col:59
VmxBroadcastNmi()(ok bool)//col:112
VmxBroadcastHandleKdDebugBreaks()(ok bool)//col:171
VmxBroadcastNmiHandler()(ok bool)//col:236
}
















)

func NewVmxBroadcast() { return & vmxBroadcast{} }

func (v *vmxBroadcast)VmxBroadcastHandleNmiCallback()(ok bool){//col:59













return true
}

















func (v *vmxBroadcast)VmxBroadcastNmi()(ok bool){//col:112






















return true
}

















func (v *vmxBroadcast)VmxBroadcastHandleKdDebugBreaks()(ok bool){//col:171














return true
}

















func (v *vmxBroadcast)VmxBroadcastNmiHandler()(ok bool){//col:236




























return true
}



















