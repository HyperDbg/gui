package meta-commands


type (
Status interface{
CommandStatusHelp()(ok bool)//col:40
CommandStatus()(ok bool)//col:107
}

)

func NewStatus() { return & status{} }

func (s *status)CommandStatusHelp()(ok bool){//col:40









return true
}


func (s *status)CommandStatus()(ok bool){//col:107





































return true
}




