package driver

import (
	"fmt"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

const (
	CommunicationBufferSize = 0x1000
	IoctlReceiveBuffer      = 0x2202004
)

type Handle struct {
	handle syscall.Handle
}

type DriverHandle = Handle

func NewDriverHandle(config *DriverConfig) (*DriverHandle, error) {
	return NewHandle(config.DeviceName)
}

func NewHandle(deviceName string) (*Handle, error) {
	deviceNamePtr, err := syscall.UTF16PtrFromString(deviceName)
	if err != nil {
		return nil, fmt.Errorf("failed to convert device name: %w", err)
	}

	handle, err := windows.CreateFile(
		deviceNamePtr,
		windows.GENERIC_READ|windows.GENERIC_WRITE,
		windows.FILE_SHARE_READ|windows.FILE_SHARE_WRITE,
		nil,
		windows.OPEN_EXISTING,
		windows.FILE_ATTRIBUTE_NORMAL,
		0,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to open device %s: %w", deviceName, err)
	}

	return &Handle{
		handle: syscall.Handle(handle),
	}, nil
}

func (h *Handle) Close() error {
	if h.handle != syscall.InvalidHandle {
		err := windows.CloseHandle(windows.Handle(h.handle))
		h.handle = syscall.InvalidHandle
		return err
	}
	return nil
}

func (h *Handle) IsConnected() bool {
	return h.handle != syscall.InvalidHandle
}

func (h *Handle) SendBuffer(buffer []byte, ioctlCode uint32) error {
	if h.handle == syscall.InvalidHandle {
		return fmt.Errorf("handle is invalid")
	}

	var bytesReturned uint32
	return windows.DeviceIoControl(
		windows.Handle(h.handle),
		ioctlCode,
		unsafe.SliceData(buffer),
		uint32(len(buffer)),
		unsafe.SliceData(buffer),
		uint32(len(buffer)),
		&bytesReturned,
		nil,
	)
}

func (h *Handle) ReceiveBuffer() ([]byte, error) {
	if h.handle == syscall.InvalidHandle {
		return nil, fmt.Errorf("handle is invalid")
	}

	buffer := make([]byte, CommunicationBufferSize)
	var bytesReturned uint32
	err := windows.DeviceIoControl(
		windows.Handle(h.handle),
		IoctlReceiveBuffer,
		nil,
		0,
		unsafe.SliceData(buffer),
		uint32(len(buffer)),
		&bytesReturned,
		nil,
	)
	if err != nil {
		return nil, err
	}

	return buffer[:bytesReturned], nil
}
