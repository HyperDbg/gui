package sdk

import (
	"github.com/ddkwork/librarygo/src/mycheck"
	"github.com/ddkwork/librarygo/src/stream/tool"
	"github.com/ddkwork/librarygo/src/windef"
	"syscall"
)

// todo need fix long type as int64 error
// add define fomr clang ast
type (
	Interface interface {
		LoadVmm() (ok bool)
		UnLoadVmm() (ok bool)
	}
	object struct {
	}
)

func (o *object) DeviceName() string { return "HyperdbgHypervisorDevice" }
func (o *object) LinkName() (*uint16, error) {
	return syscall.UTF16PtrFromString(`\\\\.\\` + o.DeviceName())
}
func (o *object) Handle() (handle syscall.Handle, err error) {
	name, err := o.LinkName()
	if !mycheck.Error(err) {
		return
	}
	return syscall.CreateFile(
		name,
		syscall.GENERIC_READ|syscall.GENERIC_WRITE,
		syscall.FILE_SHARE_READ|syscall.FILE_SHARE_WRITE,
		nil,
		syscall.OPEN_EXISTING,
		syscall.FILE_ATTRIBUTE_NORMAL|syscall.FILE_FLAG_OVERLAPPED,
		0,
	)
}
func (o *object) DeviceIoControl() (ok bool) {
	handle, err := o.Handle()
	if !mycheck.Error(err) {
		switch err {
		//todo get ERROR_ACCESS_DENIED error type number ? administrator
		// ERROR_GEN_FAILURE
		/*
		 ShowMessages("err, a device attached to the system is not functioning\n"
		                         "vmx feature might be disabled from BIOS or VBS/HVCI is active\n");
		*/
		// default "err, CreateFile failed with (%x)\n", ErrorNum);
		}
		return
	}
	if handle == syscall.InvalidHandle {

	}

	//so now start translate src/cppkit/gui/sdk/HyperDbgDev/hyperdbg/include/SDK/Headers/Constants.h
	// for set outBufferSize etc
	//
	//the verson generated was so big,maybe call this is small
	//src/cppkit/gui/sdk/HyperDbgDev/hyperdbg/include/SDK/Headers/windefTest/windef.c.i.go
	verSion := tool.VerSion()
	verSion.SetMajor(2)

	outBuffer := make([]byte, 528)
	var bytesReturned uint32
	if !mycheck.Error(syscall.DeviceIoControl(
		handle,
		windef.SMART_GET_VERSION,
		nil,
		0,
		&outBuffer[0],
		528,
		&bytesReturned,
		nil,
	)) {
		return
	}
	//todo return event buffer
	// add iocode
	return true
}

func (o *object) LoadVmm() (ok bool) {
	//TODO implement me
	panic("implement me")
}

func (o *object) UnLoadVmm() (ok bool) {
	//TODO implement me
	panic("implement me")
}

func New() Interface { return &object{} }
