package kdserial
//back\HyperDbgDev\hyperdbg\kdserial\win11sdk.h.back

type (
Win11sdk interface{
__declspec()(ok bool)//col:8
HviGetHypervisorFeatures()(ok bool)//col:12
HviIsHypervisorVendorMicrosoft()(ok bool)//col:16
}
)

func NewWin11sdk() { return & win11sdk{} }

func (w *win11sdk)__declspec()(ok bool){//col:8
/*__declspec(dllexport) void HviGetDebugDeviceOptions();
__declspec(dllexport) void HviGetHypervisorFeatures();
__declspec(dllexport) void HviIsHypervisorVendorMicrosoft();
inline void
HviGetDebugDeviceOptions()
{
    return;
}*/
return true
}

func (w *win11sdk)HviGetHypervisorFeatures()(ok bool){//col:12
/*HviGetHypervisorFeatures()
{
    return;
}*/
return true
}

func (w *win11sdk)HviIsHypervisorVendorMicrosoft()(ok bool){//col:16
/*HviIsHypervisorVendorMicrosoft()
{
    return;
}*/
return true
}



