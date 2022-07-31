package kernel-level
type (
KernelListening interface{
ListeningSerialPortInDebugger()(ok bool)//col:1106
ListeningSerialPortInDebuggee()(ok bool)//col:1277
ListeningSerialPauseDebuggerThread()(ok bool)//col:1294
ListeningSerialPauseDebuggeeThread()(ok bool)//col:1311
}

)
func NewKernelListening() { return & kernelListening{} }
func (k *kernelListening)ListeningSerialPortInDebugger()(ok bool){//col:1106
return true
}

func (k *kernelListening)ListeningSerialPortInDebuggee()(ok bool){//col:1277
return true
}

func (k *kernelListening)ListeningSerialPauseDebuggerThread()(ok bool){//col:1294
return true
}

func (k *kernelListening)ListeningSerialPauseDebuggeeThread()(ok bool){//col:1311
return true
}

