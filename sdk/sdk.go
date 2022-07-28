package sdk

import (
	"github.com/ddkwork/librarygo/src/bitfield"
	"github.com/ddkwork/librarygo/src/hardwareIndo"
	"github.com/ddkwork/librarygo/src/mycheck"
	"github.com/ddkwork/librarygo/src/mylog"
	"syscall"
)

type (
	Interface interface {
		LoadVmm() (ok bool)
		UnLoadVmm() (ok bool)
		VmxSupportDetection() (ok bool)
	}
	object struct{ handle syscall.Handle }
)

func (o *object) VmxSupportDetection() (ok bool) {
	/*
	   if (g_DeviceHandle){
	       ShowMessages("handle of the driver found, if you use 'load' before, please "
	                    "unload it using 'unload'\n");
	       return 1;
	   }
	*/
	h := hardwareIndo.New()
	if !h.CpuInfo.Get() {
		return
	}
	if h.CpuInfo.Vendor != "GenuineIntel" {
		mylog.Info("", "this program is not designed to run in a non-VT-x environemnt !")
		return
	}
	mylog.Info("", "virtualization technology is vt-x")
	field := bitfield.NewFromUint32(h.CpuInfo.Cpu1.Ecx)
	if !field.Test(5) {
		mylog.Info("", "vmx operation is not supported by your processor")
		return
	}
	mylog.Info("", "vmx operation is supported by your processor")
	//    g_IsVmxOffProcessStart = FALSE;
	return o.DeviceIoControl()
}

func New() Interface { return &object{} }

func (o *object) DeviceName() string { return "HyperdbgHypervisorDevice" }
func (o *object) LinkName() (*uint16, error) {
	return syscall.UTF16PtrFromString(`\\\\.\\` + o.DeviceName())
}
func (o *object) Handle() (ok bool) {
	if o.handle != syscall.InvalidHandle {
		return true //? //todo change to as open stata
	}
	name, err := o.LinkName()
	if !mycheck.Error(err) {
		return
	}
	handle, err := syscall.CreateFile(
		name,
		syscall.GENERIC_READ|syscall.GENERIC_WRITE,
		syscall.FILE_SHARE_READ|syscall.FILE_SHARE_WRITE,
		nil,
		syscall.OPEN_EXISTING,
		syscall.FILE_ATTRIBUTE_NORMAL|syscall.FILE_FLAG_OVERLAPPED,
		0,
	)
	if !mycheck.Error(err) {
		//e := `a device attached to the system is not functioning,vmx feature might be disabled from BIOS or VBS/HVCI is active`
		return mycheck.Error(syscall.GetLastError())
	}
	if handle == syscall.InvalidHandle {
		return mycheck.Error("handle == syscall.InvalidHandle")
	}
	o.handle = handle
	return true
}
func (o *object) DeviceIoControl() (ok bool) {
	return o.Handle()
	//l := list.New() //InitializeListHead(&g_EventTrace);
	//ntdll := syscall.NewLazyDLL("ntdll.dll")
	//ntCreateThread := ntdll.NewProc("NtCreateThread")

	//outBuffer := make([]byte, 528)
	//var bytesReturned uint32
	//if !mycheck.Error(syscall.DeviceIoControl(
	//	handle,
	//	windef.SMART_GET_VERSION,
	//	nil,
	//	0,
	//	&outBuffer[0],
	//	528,
	//	&bytesReturned,
	//	nil,
	//)) {
	//	return
	//}
}

func (o *object) LoadVmm() (ok bool) { return o.VmxSupportDetection() }
func (o *object) UnLoadVmm() (ok bool) {
	mylog.Info("", "start terminating...")
	//remove list ?    UdUninitializeUserDebugger();
	if !o.Handle() {
		return
	}
	if !mycheck.Error(syscall.DeviceIoControl(
		o.handle,
		IOCTL_TERMINATE_VMX,
		nil,
		0,
		nil,
		0,
		nil,
		nil,
	)) {
		return mycheck.Error(syscall.GetLastError())
	}
	return true
}
