package vmx


type (
VmxRegions interface{
VmxAllocateVmxonRegion()(ok bool)//col:92
VmxAllocateVmcsRegion()(ok bool)//col:159
VmxAllocateVmmStack()(ok bool)//col:188
VmxAllocateMsrBitmap()(ok bool)//col:220
VmxAllocateIoBitmaps()(ok bool)//col:267
}
















)

func NewVmxRegions() { return & vmxRegions{} }

func (v *vmxRegions)VmxAllocateVmxonRegion()(ok bool){//col:92





































return true
}

















func (v *vmxRegions)VmxAllocateVmcsRegion()(ok bool){//col:159






























return true
}

















func (v *vmxRegions)VmxAllocateVmmStack()(ok bool){//col:188













return true
}

















func (v *vmxRegions)VmxAllocateMsrBitmap()(ok bool){//col:220















return true
}

















func (v *vmxRegions)VmxAllocateIoBitmaps()(ok bool){//col:267























return true
}



















