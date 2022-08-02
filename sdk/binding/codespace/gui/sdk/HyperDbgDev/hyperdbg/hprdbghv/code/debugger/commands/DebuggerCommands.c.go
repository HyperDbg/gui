package commands



type (
	DebuggerCommands interface {
		DebuggerCommandReadMemory() (ok bool) //col:59
	}
	debuggerCommands struct{}
)

func NewDebuggerCommands() DebuggerCommands { return &debuggerCommands{} }

func (d *debuggerCommands) DebuggerCommandReadMemory() (ok bool) { //col:59
































































	return true
}


