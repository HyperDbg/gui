package ept


type (
Invept interface{
EptInvept()(ok bool)//col:31
EptInveptSingleContext()(ok bool)//col:47
EptInveptAllContexts()(ok bool)//col:59
}

)

func NewInvept() { return & invept{} }

func (i *invept)EptInvept()(ok bool){//col:31









return true
}


func (i *invept)EptInveptSingleContext()(ok bool){//col:47







return true
}


func (i *invept)EptInveptAllContexts()(ok bool){//col:59




return true
}




