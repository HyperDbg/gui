package kdserial
type (
Win11sdk interface{
__declspec()(ok bool)//col:11
HviGetHypervisorFeatures()(ok bool)//col:16
HviIsHypervisorVendorMicrosoft()(ok bool)//col:21
}

)
func NewWin11sdk() { return & win11sdk{} }
func (w *win11sdk)__declspec()(ok bool){//col:11
return true
}

func (w *win11sdk)HviGetHypervisorFeatures()(ok bool){//col:16
return true
}

func (w *win11sdk)HviIsHypervisorVendorMicrosoft()(ok bool){//col:21
return true
}

