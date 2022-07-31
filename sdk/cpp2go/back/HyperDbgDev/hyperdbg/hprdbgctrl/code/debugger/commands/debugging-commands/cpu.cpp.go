package debugging-commands
type (
Cpu interface{
CommandCpuHelp()(ok bool)//col:25
CommandCpu()(ok bool)//col:44
    static std::string Vendor()(ok bool)//col:236
ReadVendorString()(ok bool)//col:247
ReadCpuDetails()(ok bool)//col:318
}

)
func NewCpu() { return & cpu{} }
func (c *cpu)CommandCpuHelp()(ok bool){//col:25
return true
}

func (c *cpu)CommandCpu()(ok bool){//col:44
return true
}

func (c *cpu)    static std::string Vendor()(ok bool){//col:236
return true
}

func (c *cpu)ReadVendorString()(ok bool){//col:247
return true
}

func (c *cpu)ReadCpuDetails()(ok bool){//col:318
return true
}

