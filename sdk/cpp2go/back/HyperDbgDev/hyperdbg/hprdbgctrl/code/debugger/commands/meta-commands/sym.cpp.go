package meta-commands
type (
Sym interface{
CommandSymHelp()(ok bool)//col:44
CommandSym()(ok bool)//col:279
}

)
func NewSym() { return & sym{} }
func (s *sym)CommandSymHelp()(ok bool){//col:44
return true
}

func (s *sym)CommandSym()(ok bool){//col:279
return true
}

