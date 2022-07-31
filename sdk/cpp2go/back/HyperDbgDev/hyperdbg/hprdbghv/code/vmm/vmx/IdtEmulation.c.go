package vmx


type (
IdtEmulation interface{
IdtEmulationReInjectInterruptOrException()(ok bool)//col:46
IdtEmulationHandlePageFaults()(ok bool)//col:117
IdtEmulationHandleExceptionAndNmi()(ok bool)//col:349
IdtEmulationInjectInterruptWhenInterruptWindowIsOpen()(ok bool)//col:389
IdtEmulationCheckProcessOrThreadChange()(ok bool)//col:436
IdtEmulationHandleExternalInterrupt()(ok bool)//col:583
IdtEmulationHandleNmiWindowExiting()(ok bool)//col:597
IdtEmulationHandleInterruptWindowExiting()(ok bool)//col:672
}

)

func NewIdtEmulation() { return & idtEmulation{} }

func (i *idtEmulation)IdtEmulationReInjectInterruptOrException()(ok bool){//col:46










return true
}


func (i *idtEmulation)IdtEmulationHandlePageFaults()(ok bool){//col:117

























return true
}


func (i *idtEmulation)IdtEmulationHandleExceptionAndNmi()(ok bool){//col:349




























































































return true
}


func (i *idtEmulation)IdtEmulationInjectInterruptWhenInterruptWindowIsOpen()(ok bool){//col:389
















return true
}


func (i *idtEmulation)IdtEmulationCheckProcessOrThreadChange()(ok bool){//col:436






















return true
}


func (i *idtEmulation)IdtEmulationHandleExternalInterrupt()(ok bool){//col:583

















































return true
}


func (i *idtEmulation)IdtEmulationHandleNmiWindowExiting()(ok bool){//col:597




return true
}


func (i *idtEmulation)IdtEmulationHandleInterruptWindowExiting()(ok bool){//col:672





























return true
}




