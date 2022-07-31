package vmx


type (
CrossVmexits interface{
VmxHandleXsetbv()(ok bool)//col:25
VmxHandleVmxPreemptionTimerVmexit()(ok bool)//col:45
}

)

func NewCrossVmexits() { return & crossVmexits{} }

func (c *crossVmexits)VmxHandleXsetbv()(ok bool){//col:25




return true
}


func (c *crossVmexits)VmxHandleVmxPreemptionTimerVmexit()(ok bool){//col:45






return true
}




