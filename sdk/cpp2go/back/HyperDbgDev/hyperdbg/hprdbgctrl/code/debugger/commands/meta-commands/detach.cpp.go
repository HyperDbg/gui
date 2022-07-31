package meta-commands
type (
Detach interface{
CommandDetachHelp()(ok bool)//col:31
DetachFromProcess()(ok bool)//col:63
CommandDetach()(ok bool)//col:97
}

)
func NewDetach() { return & detach{} }
func (d *detach)CommandDetachHelp()(ok bool){//col:31
return true
}

func (d *detach)DetachFromProcess()(ok bool){//col:63
return true
}

func (d *detach)CommandDetach()(ok bool){//col:97
return true
}

