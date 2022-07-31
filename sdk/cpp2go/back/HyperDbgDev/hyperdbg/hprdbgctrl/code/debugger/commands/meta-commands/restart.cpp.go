package meta-commands


type (
Restart interface{
CommandRestartHelp()(ok bool)//col:34
CommandRestart()(ok bool)//col:90
}

)

func NewRestart() { return & restart{} }

func (r *restart)CommandRestartHelp()(ok bool){//col:34







return true
}


func (r *restart)CommandRestart()(ok bool){//col:90






























return true
}




