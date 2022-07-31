package app
type (
Hprdbgctrl interface{
HyperdbgSetTextMessageCallback()(ok bool)//col:43
ShowMessages()(ok bool)//col:104
ReadIrpBasedBuffer()(ok bool)//col:457
ThreadFunc()(ok bool)//col:476
HyperDbgInstallVmmDriver()(ok bool)//col:511
HyperdbgStopDriver()(ok bool)//col:534
HyperDbgStopVmmDriver()(ok bool)//col:546
HyperdbgUninstallDriver()(ok bool)//col:569
HyperDbgUninstallVmmDriver()(ok bool)//col:581
HyperDbgLoadVmm()(ok bool)//col:685
HyperDbgUnloadVmm()(ok bool)//col:791
}

)
func NewHprdbgctrl() { return & hprdbgctrl{} }
func (h *hprdbgctrl)HyperdbgSetTextMessageCallback()(ok bool){//col:43
return true
}

func (h *hprdbgctrl)ShowMessages()(ok bool){//col:104
return true
}

func (h *hprdbgctrl)ReadIrpBasedBuffer()(ok bool){//col:457
return true
}

func (h *hprdbgctrl)ThreadFunc()(ok bool){//col:476
return true
}

func (h *hprdbgctrl)HyperDbgInstallVmmDriver()(ok bool){//col:511
return true
}

func (h *hprdbgctrl)HyperdbgStopDriver()(ok bool){//col:534
return true
}

func (h *hprdbgctrl)HyperDbgStopVmmDriver()(ok bool){//col:546
return true
}

func (h *hprdbgctrl)HyperdbgUninstallDriver()(ok bool){//col:569
return true
}

func (h *hprdbgctrl)HyperDbgUninstallVmmDriver()(ok bool){//col:581
return true
}

func (h *hprdbgctrl)HyperDbgLoadVmm()(ok bool){//col:685
return true
}

func (h *hprdbgctrl)HyperDbgUnloadVmm()(ok bool){//col:791
return true
}

