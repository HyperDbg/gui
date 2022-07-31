package vmx
type (
	IdtEmulation interface {
		IdtEmulationReInjectInterruptOrException() (ok bool)             //col:46
		IdtEmulationHandlePageFaults() (ok bool)                         //col:116
		IdtEmulationHandleExceptionAndNmi() (ok bool)                    //col:347
		IdtEmulationInjectInterruptWhenInterruptWindowIsOpen() (ok bool) //col:386
		IdtEmulationCheckProcessOrThreadChange() (ok bool)               //col:432
		IdtEmulationHandleExternalInterrupt() (ok bool)                  //col:578
		IdtEmulationHandleNmiWindowExiting() (ok bool)                   //col:591
		IdtEmulationHandleInterruptWindowExiting() (ok bool)             //col:665
	}
)
func NewIdtEmulation() { return &idtEmulation{} }
func (i *idtEmulation) IdtEmulationReInjectInterruptOrException() (ok bool) { //col:46
	return true
}

func (i *idtEmulation) IdtEmulationHandlePageFaults() (ok bool) { //col:116
	return true
}

func (i *idtEmulation) IdtEmulationHandleExceptionAndNmi() (ok bool) { //col:347
	return true
}

func (i *idtEmulation) IdtEmulationInjectInterruptWhenInterruptWindowIsOpen() (ok bool) { //col:386
	return true
}

func (i *idtEmulation) IdtEmulationCheckProcessOrThreadChange() (ok bool) { //col:432
	return true
}

func (i *idtEmulation) IdtEmulationHandleExternalInterrupt() (ok bool) { //col:578
	return true
}

func (i *idtEmulation) IdtEmulationHandleNmiWindowExiting() (ok bool) { //col:591
	return true
}

func (i *idtEmulation) IdtEmulationHandleInterruptWindowExiting() (ok bool) { //col:665
	return true
}

