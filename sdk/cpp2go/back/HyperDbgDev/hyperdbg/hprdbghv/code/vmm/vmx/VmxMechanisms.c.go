package vmx


type (
VmxMechanisms interface{
VmxMechanismCreateImmediateVmexitByVmxPreemptionTimer()(ok bool)//col:33
VmxMechanismDisableImmediateVmexitByVmxPreemptionTimer()(ok bool)//col:51
VmxMechanismCreateImmediateVmexitBySelfIpi()(ok bool)//col:67
VmxMechanismCreateImmediateVmexit()(ok bool)//col:100
VmxMechanismHandleImmediateVmexit()(ok bool)//col:122
}

)

func NewVmxMechanisms() { return & vmxMechanisms{} }

func (v *vmxMechanisms)VmxMechanismCreateImmediateVmexitByVmxPreemptionTimer()(ok bool){//col:33





return true
}


func (v *vmxMechanisms)VmxMechanismDisableImmediateVmexitByVmxPreemptionTimer()(ok bool){//col:51





return true
}


func (v *vmxMechanisms)VmxMechanismCreateImmediateVmexitBySelfIpi()(ok bool){//col:67




return true
}


func (v *vmxMechanisms)VmxMechanismCreateImmediateVmexit()(ok bool){//col:100






return true
}


func (v *vmxMechanisms)VmxMechanismHandleImmediateVmexit()(ok bool){//col:122





return true
}




