package extension-commands


type (
Crwrite interface{
CommandCrwriteHelp()(ok bool)//col:33
CommandCrwrite()(ok bool)//col:200
}
















)

func NewCrwrite() { return & crwrite{} }

func (c *crwrite)CommandCrwriteHelp()(ok bool){//col:33












return true
}

















func (c *crwrite)CommandCrwrite()(ok bool){//col:200











































































































return true
}



















