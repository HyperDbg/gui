package meta-commands


type (
Listen interface{
CommandListenHelp()(ok bool)//col:43
CommandListen()(ok bool)//col:123
}

)

func NewListen() { return & listen{} }

func (l *listen)CommandListenHelp()(ok bool){//col:43












return true
}


func (l *listen)CommandListen()(ok bool){//col:123














































return true
}




