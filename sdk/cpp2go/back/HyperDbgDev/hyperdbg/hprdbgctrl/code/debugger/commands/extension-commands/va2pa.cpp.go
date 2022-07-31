package extension-commands
type (
Va2pa interface{
CommandVa2paHelp()(ok bool)//col:39
CommandVa2pa()(ok bool)//col:210
}

)
func NewVa2pa() { return & va2pa{} }
func (v *va2pa)CommandVa2paHelp()(ok bool){//col:39
return true
}

func (v *va2pa)CommandVa2pa()(ok bool){//col:210
return true
}

