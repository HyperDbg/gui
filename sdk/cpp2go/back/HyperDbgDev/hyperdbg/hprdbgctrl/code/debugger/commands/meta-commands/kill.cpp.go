package meta-commands


type (
Kill interface{
CommandKillHelp()(ok bool)//col:30
CommandKill()(ok bool)//col:60
}

)

func NewKill() { return & kill{} }

func (k *kill)CommandKillHelp()(ok bool){//col:30





return true
}


func (k *kill)CommandKill()(ok bool){//col:60















return true
}




