package debugging-commands


type (
Pause interface{
CommandPauseHelp()(ok bool)//col:32
CommandPauseRequest()(ok bool)//col:62
CommandPause()(ok bool)//col:83
}

)

func NewPause() { return & pause{} }

func (p *pause)CommandPauseHelp()(ok bool){//col:32





return true
}


func (p *pause)CommandPauseRequest()(ok bool){//col:62















return true
}


func (p *pause)CommandPause()(ok bool){//col:83










return true
}




