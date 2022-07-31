package vmx
type (
	VmxRegions interface {
		VmxAllocateVmxonRegion() (ok bool) //col:92
		VmxAllocateVmcsRegion() (ok bool)  //col:158
		VmxAllocateVmmStack() (ok bool)    //col:186
		VmxAllocateMsrBitmap() (ok bool)   //col:217
		VmxAllocateIoBitmaps() (ok bool)   //col:263
	}
)
func NewVmxRegions() { return &vmxRegions{} }
func (v *vmxRegions) VmxAllocateVmxonRegion() (ok bool) { //col:92
	return true
}

func (v *vmxRegions) VmxAllocateVmcsRegion() (ok bool) { //col:158
	return true
}

func (v *vmxRegions) VmxAllocateVmmStack() (ok bool) { //col:186
	return true
}

func (v *vmxRegions) VmxAllocateMsrBitmap() (ok bool) { //col:217
	return true
}

func (v *vmxRegions) VmxAllocateIoBitmaps() (ok bool) { //col:263
	return true
}

