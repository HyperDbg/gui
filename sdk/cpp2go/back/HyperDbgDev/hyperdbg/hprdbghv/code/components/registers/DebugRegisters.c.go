package registers
type (
	DebugRegisters interface {
		DebugRegistersSet() (ok bool) //col:259
	}
)
func NewDebugRegisters() { return &debugRegisters{} }
func (d *debugRegisters) DebugRegistersSet() (ok bool) { //col:259
	return true
}

