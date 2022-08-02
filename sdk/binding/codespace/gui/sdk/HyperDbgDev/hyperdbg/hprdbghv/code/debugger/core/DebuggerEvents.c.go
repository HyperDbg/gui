package core



type (
	DebuggerEvents interface {
		DebuggerEventEnableEferOnAllProcessors() (ok bool) //col:30
	}
	debuggerEvents struct{}
)

func NewDebuggerEvents() DebuggerEvents { return &debuggerEvents{} }

func (d *debuggerEvents) DebuggerEventEnableEferOnAllProcessors() (ok bool) { //col:30









































	return true
}


