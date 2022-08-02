package meta_commands



type (
	Logopen interface {
		CommandLogopenHelp() (ok bool) //col:37
	}
	logopen struct{}
)

func NewLogopen() Logopen { return &logopen{} }

func (l *logopen) CommandLogopenHelp() (ok bool) { //col:37












































	return true
}


