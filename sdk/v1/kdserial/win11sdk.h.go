package kdserial

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\kdserial\win11sdk.h.back

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
	/*
	   __declspec(dllexport) void HviGetDebugDeviceOptions();
	   __declspec(dllexport) void HviGetHypervisorFeatures();
	   __declspec(dllexport) void HviIsHypervisorVendorMicrosoft();
	   HviGetDebugDeviceOptions()

	   	{
	   	    return;
	   	}
	*/
	return true
}

func (w *win11sdk) HviGetHypervisorFeatures() (ok bool) { //col:11
	/*
	   HviGetHypervisorFeatures()

	   	{
	   	    return;
	   	}
	*/
	return true
}

func (w *win11sdk) HviIsHypervisorVendorMicrosoft() (ok bool) { //col:15
	/*
	   HviIsHypervisorVendorMicrosoft()

	   	{
	   	    return;
	   	}
	*/
	return true
}

