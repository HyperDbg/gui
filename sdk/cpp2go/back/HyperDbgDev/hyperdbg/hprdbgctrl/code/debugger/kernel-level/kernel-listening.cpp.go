package kernel-level


type (
KernelListening interface{
ListeningSerialPortInDebugger()(ok bool)//col:1106
ListeningSerialPortInDebuggee()(ok bool)//col:1278
ListeningSerialPauseDebuggerThread()(ok bool)//col:1296
ListeningSerialPauseDebuggeeThread()(ok bool)//col:1314
}













































)

func NewKernelListening() { return & kernelListening{} }

func (k *kernelListening)ListeningSerialPortInDebugger()(ok bool){//col:1106



































































































































































































































































































































































































































































































































































































































































return true
}














































func (k *kernelListening)ListeningSerialPortInDebuggee()(ok bool){//col:1278









































































return true
}














































func (k *kernelListening)ListeningSerialPauseDebuggerThread()(ok bool){//col:1296





return true
}














































func (k *kernelListening)ListeningSerialPauseDebuggeeThread()(ok bool){//col:1314





return true
}
















































