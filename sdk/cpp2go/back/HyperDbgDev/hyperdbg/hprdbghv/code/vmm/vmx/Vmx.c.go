package vmx
type (
	Vmx interface {
		VmxCheckVmxSupport() (ok bool)                     //col:68
		VmxInitialize() (ok bool)                          //col:163
		VmxPerformVirtualizationOnAllCores() (ok bool)     //col:257
		VmxPerformVirtualizationOnSpecificCore() (ok bool) //col:296
		VmxFixCr4AndCr0Bits() (ok bool)                    //col:331
		VmxCheckIsOnVmxRoot() (ok bool)                    //col:360
		VmxVirtualizeCurrentSystem() (ok bool)             //col:433
		VmxTerminate() (ok bool)                           //col:471
		VmxVmptrst() (ok bool)                             //col:486
		VmxClearVmcsState() (ok bool)                      //col:518
		VmxLoadVmcs() (ok bool)                            //col:540
		VmxSetupVmcs() (ok bool)                           //col:725
		VmxVmresume() (ok bool)                            //col:751
		VmxVmxoff() (ok bool)                              //col:832
		VmxReturnStackPointerForVmxoff() (ok bool)         //col:843
		VmxReturnInstructionPointerForVmxoff() (ok bool)   //col:854
		VmxPerformTermination() (ok bool)                  //col:910
	}
)
func NewVmx() { return &vmx{} }
func (v *vmx) VmxCheckVmxSupport() (ok bool) { //col:68
	return true
}

func (v *vmx) VmxInitialize() (ok bool) { //col:163
	return true
}

func (v *vmx) VmxPerformVirtualizationOnAllCores() (ok bool) { //col:257
	return true
}

func (v *vmx) VmxPerformVirtualizationOnSpecificCore() (ok bool) { //col:296
	return true
}

func (v *vmx) VmxFixCr4AndCr0Bits() (ok bool) { //col:331
	return true
}

func (v *vmx) VmxCheckIsOnVmxRoot() (ok bool) { //col:360
	return true
}

func (v *vmx) VmxVirtualizeCurrentSystem() (ok bool) { //col:433
	return true
}

func (v *vmx) VmxTerminate() (ok bool) { //col:471
	return true
}

func (v *vmx) VmxVmptrst() (ok bool) { //col:486
	return true
}

func (v *vmx) VmxClearVmcsState() (ok bool) { //col:518
	return true
}

func (v *vmx) VmxLoadVmcs() (ok bool) { //col:540
	return true
}

func (v *vmx) VmxSetupVmcs() (ok bool) { //col:725
	return true
}

func (v *vmx) VmxVmresume() (ok bool) { //col:751
	return true
}

func (v *vmx) VmxVmxoff() (ok bool) { //col:832
	return true
}

func (v *vmx) VmxReturnStackPointerForVmxoff() (ok bool) { //col:843
	return true
}

func (v *vmx) VmxReturnInstructionPointerForVmxoff() (ok bool) { //col:854
	return true
}

func (v *vmx) VmxPerformTermination() (ok bool) { //col:910
	return true
}

