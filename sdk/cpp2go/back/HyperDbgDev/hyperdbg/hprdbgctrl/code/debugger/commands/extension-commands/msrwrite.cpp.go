package extension-commands


type (
Msrwrite interface{
CommandMsrwriteHelp()(ok bool)//col:33
CommandMsrwrite()(ok bool)//col:162
}






)

func NewMsrwrite() { return & msrwrite{} }

func (m *msrwrite)CommandMsrwriteHelp()(ok bool){//col:33












return true
}







func (m *msrwrite)CommandMsrwrite()(ok bool){//col:162












































































return true
}









