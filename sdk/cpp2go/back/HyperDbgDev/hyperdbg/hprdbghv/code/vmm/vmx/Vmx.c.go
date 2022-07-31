package vmx


type (
Vmx interface{
VmxCheckVmxSupport()(ok bool)//col:68
VmxInitialize()(ok bool)//col:164
VmxPerformVirtualizationOnAllCores()(ok bool)//col:259
VmxPerformVirtualizationOnSpecificCore()(ok bool)//col:299
VmxFixCr4AndCr0Bits()(ok bool)//col:335
VmxCheckIsOnVmxRoot()(ok bool)//col:365
VmxVirtualizeCurrentSystem()(ok bool)//col:439
VmxTerminate()(ok bool)//col:478
VmxVmptrst()(ok bool)//col:494
VmxClearVmcsState()(ok bool)//col:527
VmxLoadVmcs()(ok bool)//col:550
VmxSetupVmcs()(ok bool)//col:736
VmxVmresume()(ok bool)//col:763
VmxVmxoff()(ok bool)//col:845
VmxReturnStackPointerForVmxoff()(ok bool)//col:857
VmxReturnInstructionPointerForVmxoff()(ok bool)//col:869
VmxPerformTermination()(ok bool)//col:926
}

)

func NewVmx() { return & vmx{} }

func (v *vmx)VmxCheckVmxSupport()(ok bool){//col:68

















return true
}


func (v *vmx)VmxInitialize()(ok bool){//col:164






































return true
}


func (v *vmx)VmxPerformVirtualizationOnAllCores()(ok bool){//col:259













































return true
}


func (v *vmx)VmxPerformVirtualizationOnSpecificCore()(ok bool){//col:299




















return true
}


func (v *vmx)VmxFixCr4AndCr0Bits()(ok bool){//col:335


















return true
}


func (v *vmx)VmxCheckIsOnVmxRoot()(ok bool){//col:365


















return true
}


func (v *vmx)VmxVirtualizeCurrentSystem()(ok bool){//col:439




























return true
}


func (v *vmx)VmxTerminate()(ok bool){//col:478



















return true
}


func (v *vmx)VmxVmptrst()(ok bool){//col:494







return true
}


func (v *vmx)VmxClearVmcsState()(ok bool){//col:527













return true
}


func (v *vmx)VmxLoadVmcs()(ok bool){//col:550











return true
}


func (v *vmx)VmxSetupVmcs()(ok bool){//col:736





























































































return true
}


func (v *vmx)VmxVmresume()(ok bool){//col:763








return true
}


func (v *vmx)VmxVmxoff()(ok bool){//col:845






















return true
}


func (v *vmx)VmxReturnStackPointerForVmxoff()(ok bool){//col:857




return true
}


func (v *vmx)VmxReturnInstructionPointerForVmxoff()(ok bool){//col:869




return true
}


func (v *vmx)VmxPerformTermination()(ok bool){//col:926












return true
}




