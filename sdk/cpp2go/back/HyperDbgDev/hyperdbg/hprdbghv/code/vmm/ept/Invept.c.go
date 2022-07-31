package ept
type (
	Invept interface {
		EptInvept() (ok bool)              //col:31
		EptInveptSingleContext() (ok bool) //col:46
		EptInveptAllContexts() (ok bool)   //col:57
	}
)
func NewInvept() { return &invept{} }
func (i *invept) EptInvept() (ok bool) { //col:31
	return true
}

func (i *invept) EptInveptSingleContext() (ok bool) { //col:46
	return true
}

func (i *invept) EptInveptAllContexts() (ok bool) { //col:57
	return true
}

