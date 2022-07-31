package meta-commands


type (
Detach interface{
CommandDetachHelp()(ok bool)//col:31
DetachFromProcess()(ok bool)//col:64
CommandDetach()(ok bool)//col:99
}






)

func NewDetach() { return & detach{} }

func (d *detach)CommandDetachHelp()(ok bool){//col:31





return true
}







func (d *detach)DetachFromProcess()(ok bool){//col:64













return true
}







func (d *detach)CommandDetach()(ok bool){//col:99

















return true
}









