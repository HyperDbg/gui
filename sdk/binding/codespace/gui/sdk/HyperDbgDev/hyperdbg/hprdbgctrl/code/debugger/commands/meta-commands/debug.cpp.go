package meta_commands



type (
	Debug interface {
		CommandDebugHelp() (ok bool)         //col:32
		CommandDebugCheckComPort() (ok bool) //col:56
	}
	debug struct{}
)

func NewDebug() Debug { return &debug{} }

func (d *debug) CommandDebugHelp() (ok bool) { //col:32





































	return true
}


func (d *debug) CommandDebugCheckComPort() (ok bool) { //col:56



























	return true
}


