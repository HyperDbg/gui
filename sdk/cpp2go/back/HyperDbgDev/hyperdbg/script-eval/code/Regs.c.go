package code
type (
Regs interface{
GetRegValue()(ok bool)//col:951
SetRegValue()(ok bool)//col:1974
}

)
func NewRegs() { return & regs{} }
func (r *regs)GetRegValue()(ok bool){//col:951
return true
}

func (r *regs)SetRegValue()(ok bool){//col:1974
return true
}

