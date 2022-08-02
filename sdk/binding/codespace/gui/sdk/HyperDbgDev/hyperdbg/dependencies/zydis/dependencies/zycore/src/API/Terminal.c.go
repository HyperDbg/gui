package API

type (
	Terminal interface {
		ZyanStatus_ZyanTerminalEnableVT100() (ok bool) //col:35
		ZyanStatus_ZyanTerminalIsTTY() (ok bool)       //col:79
	}
	terminal struct{}
)

func NewTerminal() Terminal { return &terminal{} }

func (t *terminal) ZyanStatus_ZyanTerminalEnableVT100() (ok bool) { //col:35

	return true
}

func (t *terminal) ZyanStatus_ZyanTerminalIsTTY() (ok bool) { //col:79

	return true
}
