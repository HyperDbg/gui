package app


type (
Hprdbgctrl interface{
HyperdbgSetTextMessageCallback()(ok bool)//col:43
ShowMessages()(ok bool)//col:105
ReadIrpBasedBuffer()(ok bool)//col:459
ThreadFunc()(ok bool)//col:479
HyperDbgInstallVmmDriver()(ok bool)//col:515
HyperdbgStopDriver()(ok bool)//col:539
HyperDbgStopVmmDriver()(ok bool)//col:552
HyperdbgUninstallDriver()(ok bool)//col:576
HyperDbgUninstallVmmDriver()(ok bool)//col:589
HyperDbgLoadVmm()(ok bool)//col:694
HyperDbgUnloadVmm()(ok bool)//col:801
}






)

func NewHprdbgctrl() { return & hprdbgctrl{} }

func (h *hprdbgctrl)HyperdbgSetTextMessageCallback()(ok bool){//col:43




return true
}







func (h *hprdbgctrl)ShowMessages()(ok bool){//col:105







































return true
}







func (h *hprdbgctrl)ReadIrpBasedBuffer()(ok bool){//col:459


















































































































































































return true
}







func (h *hprdbgctrl)ThreadFunc()(ok bool){//col:479





return true
}







func (h *hprdbgctrl)HyperDbgInstallVmmDriver()(ok bool){//col:515














return true
}







func (h *hprdbgctrl)HyperdbgStopDriver()(ok bool){//col:539












return true
}







func (h *hprdbgctrl)HyperDbgStopVmmDriver()(ok bool){//col:552




return true
}







func (h *hprdbgctrl)HyperdbgUninstallDriver()(ok bool){//col:576












return true
}







func (h *hprdbgctrl)HyperDbgUninstallVmmDriver()(ok bool){//col:589




return true
}







func (h *hprdbgctrl)HyperDbgLoadVmm()(ok bool){//col:694


































































return true
}







func (h *hprdbgctrl)HyperDbgUnloadVmm()(ok bool){//col:801































return true
}









