package debugging-commands


type (
Cpu interface{
CommandCpuHelp()(ok bool)//col:25
CommandCpu()(ok bool)//col:45
    static std::string Vendor()(ok bool)//col:238
ReadVendorString()(ok bool)//col:249
ReadCpuDetails()(ok bool)//col:321
}














































)

func NewCpu() { return & cpu{} }

func (c *cpu)CommandCpuHelp()(ok bool){//col:25





return true
}















































func (c *cpu)CommandCpu()(ok bool){//col:45










return true
}















































func (c *cpu)    static std::string Vendor()(ok bool){//col:238










































































































































return true
}















































func (c *cpu)ReadVendorString()(ok bool){//col:249




return true
}















































func (c *cpu)ReadCpuDetails()(ok bool){//col:321






























































return true
}

















































