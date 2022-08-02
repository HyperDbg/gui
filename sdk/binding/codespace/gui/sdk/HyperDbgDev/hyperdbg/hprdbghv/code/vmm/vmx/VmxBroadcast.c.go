package vmx



type (
	VmxBroadcast interface {
		VmxBroadcastHandleNmiCallback() (ok bool) //col:13
	}
	vmxBroadcast struct{}
)

func NewVmxBroadcast() VmxBroadcast { return &vmxBroadcast{} }

func (v *vmxBroadcast) VmxBroadcastHandleNmiCallback() (ok bool) { //col:13
















	return true
}


