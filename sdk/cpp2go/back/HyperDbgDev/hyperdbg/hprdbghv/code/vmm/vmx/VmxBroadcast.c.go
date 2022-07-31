package vmx
type (
	VmxBroadcast interface {
		VmxBroadcastHandleNmiCallback() (ok bool)   //col:59
		VmxBroadcastNmi() (ok bool)                 //col:111
		VmxBroadcastHandleKdDebugBreaks() (ok bool) //col:169
		VmxBroadcastNmiHandler() (ok bool)          //col:233
	}
)
func NewVmxBroadcast() { return &vmxBroadcast{} }
func (v *vmxBroadcast) VmxBroadcastHandleNmiCallback() (ok bool) { //col:59
	return true
}

func (v *vmxBroadcast) VmxBroadcastNmi() (ok bool) { //col:111
	return true
}

func (v *vmxBroadcast) VmxBroadcastHandleKdDebugBreaks() (ok bool) { //col:169
	return true
}

func (v *vmxBroadcast) VmxBroadcastNmiHandler() (ok bool) { //col:233
	return true
}

