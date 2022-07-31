package app


type (
Dllmain interface{
DllMain()(ok bool)//col:34
}
















)

func NewDllmain() { return & dllmain{} }

func (d *dllmain)DllMain()(ok bool){//col:34












return true
}



















