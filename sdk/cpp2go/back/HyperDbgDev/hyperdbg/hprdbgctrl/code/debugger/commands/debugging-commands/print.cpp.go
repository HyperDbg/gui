package debugging-commands


type (
Print interface{
CommandPrintHelp()(ok bool)//col:36
CommandPrint()(ok bool)//col:132
}

)

func NewPrint() { return & print{} }

func (p *print)CommandPrintHelp()(ok bool){//col:36






return true
}


func (p *print)CommandPrint()(ok bool){//col:132




































return true
}




