package commands
type (
	DebuggerCommands interface {
		DebuggerCommandReadMemory() (ok bool)                  //col:44
		DebuggerCommandReadMemoryVmxRoot() (ok bool)           //col:137
		DebuggerReadOrWriteMsr() (ok bool)                     //col:294
		DebuggerCommandEditMemory() (ok bool)                  //col:400
		DebuggerCommandEditMemoryVmxRoot() (ok bool)           //col:501
		PerformSearchAddress() (ok bool)                       //col:789
		SearchAddressWrapper() (ok bool)                       //col:971
		DebuggerCommandSearchMemory() (ok bool)                //col:1077
		DebuggerCommandFlush() (ok bool)                       //col:1096
		DebuggerCommandSignalExecutionState() (ok bool)        //col:1116
		DebuggerCommandSendMessage() (ok bool)                 //col:1138
		DebuggerCommandSendGeneralBufferToDebugger() (ok bool) //col:1160
		DebuggerCommandReservePreallocatedPools() (ok bool)    //col:1217
	}
)
func NewDebuggerCommands() { return &debuggerCommands{} }
func (d *debuggerCommands) DebuggerCommandReadMemory() (ok bool) { //col:44
	return true
}

func (d *debuggerCommands) DebuggerCommandReadMemoryVmxRoot() (ok bool) { //col:137
	return true
}

func (d *debuggerCommands) DebuggerReadOrWriteMsr() (ok bool) { //col:294
	return true
}

func (d *debuggerCommands) DebuggerCommandEditMemory() (ok bool) { //col:400
	return true
}

func (d *debuggerCommands) DebuggerCommandEditMemoryVmxRoot() (ok bool) { //col:501
	return true
}

func (d *debuggerCommands) PerformSearchAddress() (ok bool) { //col:789
	return true
}

func (d *debuggerCommands) SearchAddressWrapper() (ok bool) { //col:971
	return true
}

func (d *debuggerCommands) DebuggerCommandSearchMemory() (ok bool) { //col:1077
	return true
}

func (d *debuggerCommands) DebuggerCommandFlush() (ok bool) { //col:1096
	return true
}

func (d *debuggerCommands) DebuggerCommandSignalExecutionState() (ok bool) { //col:1116
	return true
}

func (d *debuggerCommands) DebuggerCommandSendMessage() (ok bool) { //col:1138
	return true
}

func (d *debuggerCommands) DebuggerCommandSendGeneralBufferToDebugger() (ok bool) { //col:1160
	return true
}

func (d *debuggerCommands) DebuggerCommandReservePreallocatedPools() (ok bool) { //col:1217
	return true
}

