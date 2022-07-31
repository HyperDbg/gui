package extension-commands


type (
Vmcall interface{
CommandVmcallHelp()(ok bool)//col:33
CommandVmcall()(ok bool)//col:123
}
















)

func NewVmcall() { return & vmcall{} }

func (v *vmcall)CommandVmcallHelp()(ok bool){//col:33












return true
}

















func (v *vmcall)CommandVmcall()(ok bool){//col:123




















































return true
}



















