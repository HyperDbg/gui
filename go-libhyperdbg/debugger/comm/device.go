package comm

import (
	"context"
	"fmt"
	"unsafe"

	"golang.org/x/sys/windows"
)

// Device wraps a HANDLE to the HyperDbg VMM driver. It is the Go equivalent of
// g_DeviceHandle + PlatformOpenDevice/PlatformDeviceIoControl in the C
// libhyperdbg. A Device is not safe for concurrent use; callers serialising
// IOCTLs themselves may share one, but the higher-level Debugger will expose a
// goroutine-safe façade.
type Device struct {
	handle windows.Handle
	name   string
}

// Open opens the HyperDbg device at the given Win32 path (e.g. DeviceName).
// The access flags mirror PlatformOpenDevice exactly:
// GENERIC_READ|GENERIC_WRITE, FILE_SHARE_READ|FILE_SHARE_WRITE,
// OPEN_EXISTING, FILE_ATTRIBUTE_NORMAL.
//
// Open requires administrator privileges; ERROR_ACCESS_DENIED is mapped to a
// descriptive error to aid diagnosis.
func Open(ctx context.Context, devicePath string) (*Device, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	namePtr, err := windows.UTF16PtrFromString(devicePath)
	if err != nil {
		return nil, err
	}
	h, err := windows.CreateFile(
		namePtr,
		windows.GENERIC_READ|windows.GENERIC_WRITE,
		windows.FILE_SHARE_READ|windows.FILE_SHARE_WRITE,
		nil,
		windows.OPEN_EXISTING,
		windows.FILE_ATTRIBUTE_NORMAL,
		0)
	if err != nil {
		if err == windows.ERROR_ACCESS_DENIED {
			return nil, fmt.Errorf("access denied opening %q: run as administrator", devicePath)
		}
		// ERROR_GEN_FAILURE (31) is what hyperkd.sys's IRP_MJ_CREATE handler
		// returns when VMX/VT-x is unavailable. The C++ libhyperdbg maps this
		// to the same diagnostic (libhyperdbg.cpp:341):
		//   "vmx feature might be disabled from BIOS or VBS/HVCI is active"
		if err == windows.ERROR_GEN_FAILURE {
			return nil, fmt.Errorf("CreateFile(%q) failed: %w"+
				"\n  VMX feature might be disabled from BIOS or VBS/HVCI is active."+
				"\n  Fixes:"+
				"\n    1. Enable VT-x/Intel VT in BIOS/UEFI."+
				"\n    2. Disable VBS & HVCI: gpedit.msc ->"+
				" Computer Configuration -> Administrative Templates ->"+
				" System -> Device Guard -> 'Turn On Virtualization Based Security' = Disabled,"+
				" then reboot."+
				"\n    3. Disable Hyper-V (optional): bcdedit /set hypervisorlaunchtype off,"+
				" then reboot."+
				"\n    4. Verify with 'msinfo32' -> 'Virtualization-based security' should be 'Not enabled'.",
				devicePath, err)
		}
		return nil, fmt.Errorf("CreateFile(%q) failed: %w", devicePath, err)
	}
	return &Device{handle: h, name: devicePath}, nil
}

// OpenDefault opens the standard HyperDbg device (\\.\HyperDbgDebuggerDevice).
// This is the convenience entry point used by Connect("local").
func OpenDefault(ctx context.Context) (*Device, error) {
	return Open(ctx, DeviceName)
}

// Close releases the device handle. It is safe to call on an already-closed
// Device; subsequent calls are no-ops.
func (d *Device) Close() error {
	if d.handle == 0 || d.handle == windows.InvalidHandle {
		return nil
	}
	err := windows.CloseHandle(d.handle)
	d.handle = 0
	return err
}

// Handle returns the underlying Windows HANDLE. It is intended only for
// advanced callers that need to pass the handle to OS APIs not yet wrapped by
// this package (e.g. overlapped I/O for event IRPs).
func (d *Device) Handle() windows.Handle { return d.handle }

// Name returns the device path the handle was opened with.
func (d *Device) Name() string { return d.name }

// Ioctl sends a DeviceIoControl request to the driver.
//
// inBuf is the input buffer (may be nil for IOCTLs that take no input); outBuf
// is the output buffer (may be nil). When inBuf and outBuf alias the same
// memory (the common METHOD_BUFFERED case) pass the same slice for both and
// the driver will overwrite it in place. The number of bytes returned by the
// driver is reported via bytesReturned.
//
// ctx is checked for cancellation before the syscall; DeviceIoControl itself
// is synchronous and cannot be interrupted mid-call.
func (d *Device) Ioctl(ctx context.Context, code uint32, inBuf, outBuf []byte) (bytesReturned uint32, err error) {
	if err := ctx.Err(); err != nil {
		return 0, err
	}
	if d.handle == 0 || d.handle == windows.InvalidHandle {
		return 0, fmt.Errorf("Ioctl on closed device")
	}
	var inPtr, outPtr *byte
	if len(inBuf) > 0 {
		inPtr = (*byte)(unsafe.Pointer(&inBuf[0]))
	}
	if len(outBuf) > 0 {
		outPtr = (*byte)(unsafe.Pointer(&outBuf[0]))
	}
	err = windows.DeviceIoControl(
		d.handle, code,
		inPtr, uint32(len(inBuf)),
		outPtr, uint32(len(outBuf)),
		&bytesReturned, nil)
	if err != nil {
		return bytesReturned, fmt.Errorf("DeviceIoControl(0x%x) failed: %w", code, err)
	}
	return bytesReturned, nil
}

// IoctlStruct is a typed convenience over Ioctl for the common case where the
// input and output buffers are Go structs (or pointers to structs) whose
// memory layout matches the C ABI. The caller is responsible for ensuring
// layout compatibility — typically by constructing the value from the
// go-libhyperdbg/types package.
//
// inVal must be a pointer to a struct (or a slice with a backing array);
// outVal must be a pointer to a struct into which the driver writes. Either
// may be nil. The returned n is the number of bytes the driver wrote.
func (d *Device) IoctlStruct(ctx context.Context, code uint32, inVal, outVal unsafe.Pointer, inSize, outSize uint32) (uint32, error) {
	if err := ctx.Err(); err != nil {
		return 0, err
	}
	if d.handle == 0 || d.handle == windows.InvalidHandle {
		return 0, fmt.Errorf("IoctlStruct on closed device")
	}
	var n uint32
	err := windows.DeviceIoControl(
		d.handle, code,
		(*byte)(inVal), inSize,
		(*byte)(outVal), outSize,
		&n, nil)
	if err != nil {
		return n, fmt.Errorf("DeviceIoControl(0x%x) failed: %w", code, err)
	}
	return n, nil
}
