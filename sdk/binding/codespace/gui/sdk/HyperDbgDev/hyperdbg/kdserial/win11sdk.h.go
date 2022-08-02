package kdserial



type (
	Win11sdk interface {
		__declspec() (ok bool)                     //col:7
		HviGetHypervisorFeatures() (ok bool)       //col:11
		HviIsHypervisorVendorMicrosoft() (ok bool) //col:15
	}
	win11sdk struct{}
)

func NewWin11sdk() Win11sdk { return &win11sdk{} }

func (w *win11sdk) __declspec() (ok bool) { //col:7










	return true
}


func (w *win11sdk) HviGetHypervisorFeatures() (ok bool) { //col:11







	return true
}


func (w *win11sdk) HviIsHypervisorVendorMicrosoft() (ok bool) { //col:15







	return true
}


