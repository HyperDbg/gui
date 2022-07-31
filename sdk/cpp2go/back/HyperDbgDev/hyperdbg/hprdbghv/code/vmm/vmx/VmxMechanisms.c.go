package vmx
type (
	VmxMechanisms interface {
		VmxMechanismCreateImmediateVmexitByVmxPreemptionTimer() (ok bool)  //col:33
		VmxMechanismDisableImmediateVmexitByVmxPreemptionTimer() (ok bool) //col:50
		VmxMechanismCreateImmediateVmexitBySelfIpi() (ok bool)             //col:65
		VmxMechanismCreateImmediateVmexit() (ok bool)                      //col:97
		VmxMechanismHandleImmediateVmexit() (ok bool)                      //col:118
	}
)
func NewVmxMechanisms() { return &vmxMechanisms{} }
func (v *vmxMechanisms) VmxMechanismCreateImmediateVmexitByVmxPreemptionTimer() (ok bool) { //col:33
	return true
}

func (v *vmxMechanisms) VmxMechanismDisableImmediateVmexitByVmxPreemptionTimer() (ok bool) { //col:50
	return true
}

func (v *vmxMechanisms) VmxMechanismCreateImmediateVmexitBySelfIpi() (ok bool) { //col:65
	return true
}

func (v *vmxMechanisms) VmxMechanismCreateImmediateVmexit() (ok bool) { //col:97
	return true
}

func (v *vmxMechanisms) VmxMechanismHandleImmediateVmexit() (ok bool) { //col:118
	return true
}

