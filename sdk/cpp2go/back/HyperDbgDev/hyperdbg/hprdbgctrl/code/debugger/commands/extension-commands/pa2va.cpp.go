package extension-commands
type (
Pa2va interface{
CommandPa2vaHelp()(ok bool)//col:38
CommandPa2va()(ok bool)//col:208
}

)
func NewPa2va() { return & pa2va{} }
func (p *pa2va)CommandPa2vaHelp()(ok bool){//col:38
return true
}

func (p *pa2va)CommandPa2va()(ok bool){//col:208
return true
}

