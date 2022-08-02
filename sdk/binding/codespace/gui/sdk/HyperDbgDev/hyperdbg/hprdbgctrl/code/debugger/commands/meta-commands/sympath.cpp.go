package meta_commands



type (
	Sympath interface {
		CommandSympathHelp() (ok bool) //col:9
	}
	sympath struct{}
)

func NewSympath() Sympath { return &sympath{} }

func (s *sympath) CommandSympathHelp() (ok bool) { //col:9












	return true
}


