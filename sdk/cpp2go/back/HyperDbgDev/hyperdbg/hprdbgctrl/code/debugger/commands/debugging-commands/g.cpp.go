package debugging-commands
type (
G interface{
CommandGHelp()(ok bool)//col:33
CommandGRequest()(ok bool)//col:83
CommandG()(ok bool)//col:103
}

)
func NewG() { return & g{} }
func (g *g)CommandGHelp()(ok bool){//col:33
return true
}

func (g *g)CommandGRequest()(ok bool){//col:83
return true
}

func (g *g)CommandG()(ok bool){//col:103
return true
}

