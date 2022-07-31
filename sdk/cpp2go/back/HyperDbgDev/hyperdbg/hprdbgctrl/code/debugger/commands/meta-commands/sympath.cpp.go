package meta-commands
type (
Sympath interface{
CommandSympathHelp()(ok bool)//col:34
CommandSympath()(ok bool)//col:131
}

)
func NewSympath() { return & sympath{} }
func (s *sympath)CommandSympathHelp()(ok bool){//col:34
return true
}

func (s *sympath)CommandSympath()(ok bool){//col:131
return true
}

