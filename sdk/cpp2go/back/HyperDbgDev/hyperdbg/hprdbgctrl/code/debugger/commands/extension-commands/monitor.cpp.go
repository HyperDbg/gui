package extension-commands
type (
Monitor interface{
CommandMonitorHelp()(ok bool)//col:35
CommandMonitor()(ok bool)//col:243
}

)
func NewMonitor() { return & monitor{} }
func (m *monitor)CommandMonitorHelp()(ok bool){//col:35
return true
}

func (m *monitor)CommandMonitor()(ok bool){//col:243
return true
}

