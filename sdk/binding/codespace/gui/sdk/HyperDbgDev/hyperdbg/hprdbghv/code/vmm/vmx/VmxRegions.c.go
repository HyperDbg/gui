package vmx



type (
	VmxRegions interface {
		VmxAllocateVmxonRegion() (ok bool) //col:55
	}
	vmxRegions struct{}
)

func NewVmxRegions() VmxRegions { return &vmxRegions{} }

func (v *vmxRegions) VmxAllocateVmxonRegion() (ok bool) { //col:55
































































	return true
}


