package ddk

import (
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/ddkwork/golibrary/byteslice"
	"github.com/ddkwork/golibrary/std/mylog"
	"golang.org/x/sys/windows"
)

//go:embed RTCore64.sys
var rtCore64Sys []byte

const (
	RTCORE_SERVICE_NAME = "RTCore64"
	RTCORE_DEVICE_NAME  = "RTCore64"
)

type Driver struct {
	Path string
	Name string
}

func NewDriver(driverPath string) *Driver {
	name := strings.TrimSuffix(filepath.Base(driverPath), filepath.Ext(driverPath))
	return &Driver{Path: driverPath, Name: name}
}

func (d *Driver) withSCManager(fn func(windows.Handle) bool) bool {
	scManager := mylog.Check2(windows.OpenSCManager(nil, nil, windows.SC_MANAGER_ALL_ACCESS))

	defer func() { mylog.Check(windows.CloseServiceHandle(scManager)) }()
	return fn(scManager)
}

func (d *Driver) Install() bool {
	return d.withSCManager(func(sc windows.Handle) bool {
		driverNamePtr := mylog.Check2(windows.UTF16PtrFromString(d.Name))
		serviceExePtr := mylog.Check2(windows.UTF16PtrFromString(d.Path))

		schService := mylog.Check2(windows.CreateService(
			sc,
			driverNamePtr,
			driverNamePtr,
			windows.SERVICE_ALL_ACCESS,
			windows.SERVICE_KERNEL_DRIVER,
			windows.SERVICE_DEMAND_START,
			windows.SERVICE_ERROR_NORMAL,
			serviceExePtr,
			nil, nil, nil, nil, nil,
		))

		if schService != 0 {
			windows.CloseServiceHandle(schService)
		}

		mylog.Success("driver installed successfully")
		return true
	})
}

func (d *Driver) Remove() bool {
	return d.withSCManager(func(handle windows.Handle) bool {
		driverNamePtr := mylog.Check2(windows.UTF16PtrFromString(d.Name))

		schService, e := windows.OpenService(handle, driverNamePtr, windows.SERVICE_ALL_ACCESS)
		if e != nil {
			if e == windows.ERROR_SERVICE_DOES_NOT_EXIST {
				return true
			}
			mylog.Warning("OpenService failed in remove", "error", e)
			return false
		}
		defer func() { mylog.Check(windows.CloseServiceHandle(schService)) }()

		mylog.Check(windows.DeleteService(schService))
		if e != nil && e != windows.ERROR_SERVICE_MARKED_FOR_DELETE {
			mylog.Warning("DeleteService failed", "error", e)
			return false
		}

		mylog.Success("driver removed successfully")
		return true
	})
}

func (d *Driver) Start() bool {
	return d.withSCManager(func(handle windows.Handle) bool {
		driverNamePtr := mylog.Check2(windows.UTF16PtrFromString(d.Name))

		schService, err := windows.OpenService(handle, driverNamePtr, windows.SERVICE_ALL_ACCESS)
		if err != nil {
			if err == windows.ERROR_SERVICE_DOES_NOT_EXIST {
				mylog.Info("service does not exist, trying to install")
				return false
			}
			mylog.Warning("OpenService failed in start", "error", err)
			return false
		}
		defer func() { mylog.Check(windows.CloseServiceHandle(schService)) }()

		mylog.Check(windows.StartService(schService, 0, nil))
		if err != nil {
			if err == windows.ERROR_SERVICE_ALREADY_RUNNING {
				mylog.Info("service is already running")
			} else {
				mylog.Warning(err)
				return false
			}
		} else {
			mylog.Success("driver started successfully")
		}
		return true
	})
}

func (d *Driver) Stop() bool {
	return d.withSCManager(func(handle windows.Handle) bool {
		driverNamePtr := mylog.Check2(windows.UTF16PtrFromString(d.Name))

		schService := mylog.Check2(windows.OpenService(handle, driverNamePtr, windows.SERVICE_ALL_ACCESS))

		defer func() { mylog.Check(windows.CloseServiceHandle(schService)) }()

		var serviceStatus windows.SERVICE_STATUS
		mylog.Check(windows.ControlService(schService, windows.SERVICE_CONTROL_STOP, &serviceStatus))
		mylog.Success("driver stopped successfully")
		return true
	})
}

type RTCore64 struct {
	privilege    *Privilege
	deviceHandle windows.Handle
	driver       *Driver
}

func NewRTCore64() *RTCore64 {
	p := NewPrivilege()
	p.Debug()
	p.LoadDriver()
	return &RTCore64{privilege: p}
}

func (r *RTCore64) Privilege() *Privilege {
	return r.privilege
}

func (r *RTCore64) Load() bool {
	if r.deviceHandle != 0 {
		return true
	}

	tmpPath := filepath.Join(os.TempDir(), "golibrary_RTCore64.sys")
	mylog.Check(os.WriteFile(tmpPath, rtCore64Sys, 0o644))

	r.driver = NewDriver(tmpPath)
	r.driver.Remove()
	if !r.driver.Install() {
		mylog.Warning("install RTCore64 service failed")
		r.driver = nil
		return false
	}
	if !r.driver.Start() {
		mylog.Warning("start RTCore64 service failed")
		r.driver.Remove()
		r.driver = nil
		return false
	}

	namePtr, _ := windows.UTF16PtrFromString(fmt.Sprintf(`\\.\%s`, RTCORE_DEVICE_NAME))
	h, e := windows.CreateFile(namePtr,
		windows.GENERIC_READ|windows.GENERIC_WRITE,
		0, nil, windows.OPEN_EXISTING, 0, 0)
	if e != nil {
		mylog.Warning("CreateFile for RTCore64 failed", "error", e)
		r.driver.Stop()
		r.driver.Remove()
		r.driver = nil
		return false
	}
	r.deviceHandle = h
	mylog.Success("RTCore64 loaded")
	return true
}

func (r *RTCore64) Unload() {
	if r.deviceHandle != 0 && r.deviceHandle != windows.InvalidHandle {
		windows.CloseHandle(r.deviceHandle)
		r.deviceHandle = 0
	}
	if r.driver != nil {
		r.driver.Stop()
		r.driver.Remove()
		r.driver = nil
	}
}

func (r *RTCore64) Handle() windows.Handle {
	return r.deviceHandle
}

func (r *RTCore64) Loaded() bool {
	return r.deviceHandle != 0 && r.deviceHandle != windows.InvalidHandle
}

const (
	RTCORE_IOCTL_READ  = 0x80002048
	RTCORE_IOCTL_WRITE = 0x8000204C
	PAGE_SIZE          = 0x1000
)

type rtcPacket struct {
	pad0  [8]byte
	addr  uint64
	pad1  [8]byte
	size  uint32
	value uint32
	pad3  [16]byte
}

func (r *RTCore64) ReadMemory(addr uint64, buf []byte) {
	for offset := 0; offset < len(buf); {
		readAddr := addr + uint64(offset)
		remaining := len(buf) - offset

		if remaining >= 4 && (readAddr&3) == 0 {
			val := r.readDword(readAddr)

			buf[offset] = byte(val)
			buf[offset+1] = byte(val >> 8)
			buf[offset+2] = byte(val >> 16)
			buf[offset+3] = byte(val >> 24)
			offset += 4
		} else if remaining >= 2 && (readAddr&1) == 0 {
			val := r.readWord(readAddr)

			buf[offset] = byte(val)
			buf[offset+1] = byte(val >> 8)
			offset += 2
		} else {
			val := r.readByte(readAddr)

			buf[offset] = val
			offset += 1
		}
	}
}

func (r *RTCore64) readDword(addr uint64) uint32 {
	var pkt rtcPacket
	pkt.addr = addr
	pkt.size = 4

	pktBytes := byteslice.FromStruct(&pkt)
	var bytesReturned uint32
	mylog.Check(windows.DeviceIoControl(r.deviceHandle, RTCORE_IOCTL_READ,
		byteslice.PtrFromAnySlice(pktBytes), uint32(len(pktBytes)),
		byteslice.PtrFromAnySlice(pktBytes), uint32(len(pktBytes)),
		&bytesReturned, nil))

	return pkt.value
}

func (r *RTCore64) readWord(addr uint64) uint16 {
	var pkt rtcPacket
	pkt.addr = addr
	pkt.size = 2

	pktBytes := byteslice.FromStruct(&pkt)
	var bytesReturned uint32
	mylog.Check(windows.DeviceIoControl(r.deviceHandle, RTCORE_IOCTL_READ,
		byteslice.PtrFromAnySlice(pktBytes), uint32(len(pktBytes)),
		byteslice.PtrFromAnySlice(pktBytes), uint32(len(pktBytes)),
		&bytesReturned, nil))

	return uint16(pkt.value)
}

func (r *RTCore64) readByte(addr uint64) byte {
	var pkt rtcPacket
	pkt.addr = addr
	pkt.size = 1

	pktBytes := byteslice.FromStruct(&pkt)
	var bytesReturned uint32
	mylog.Check(windows.DeviceIoControl(r.deviceHandle, RTCORE_IOCTL_READ,
		byteslice.PtrFromAnySlice(pktBytes), uint32(len(pktBytes)),
		byteslice.PtrFromAnySlice(pktBytes), uint32(len(pktBytes)),
		&bytesReturned, nil))

	return byte(pkt.value)
}

func (r *RTCore64) WriteMemory(addr uint64, buf []byte) error {
	var pkt rtcPacket
	pkt.addr = addr
	pkt.size = uint32(len(buf))
	pkt.value = binaryLittleEndianUint32(buf)

	pktBytes := byteslice.FromStruct(&pkt)
	var bytesReturned uint32
	mylog.Check(windows.DeviceIoControl(r.deviceHandle, RTCORE_IOCTL_WRITE,
		byteslice.PtrFromAnySlice(pktBytes), uint32(len(pktBytes)),
		byteslice.PtrFromAnySlice(pktBytes), uint32(len(pktBytes)),
		&bytesReturned, nil))

	return nil
}

func (r *RTCore64) ReadUint64(addr uint64) uint64 {
	buf := make([]byte, 8)
	r.ReadMemory(addr, buf)
	return binaryLittleEndianUint64(buf)
}

func (r *RTCore64) ReadUint32(addr uint64) uint32 {
	buf := make([]byte, 4)
	r.ReadMemory(addr, buf)
	return binaryLittleEndianUint32(buf)
}

func (r *RTCore64) ReadUint16(addr uint64) uint16 {
	buf := make([]byte, 2)
	r.ReadMemory(addr, buf)
	return binaryLittleEndianUint16(buf)
}

func (r *RTCore64) WriteUint32(addr uint64, val uint32) {
	buf := make([]byte, 4)
	putBinaryLittleEndianUint32(buf, val)
	r.WriteMemory(addr, buf)
}

func binaryLittleEndianUint32(b []byte) uint32 {
	return uint32(b[0]) | uint32(b[1])<<8 | uint32(b[2])<<16 | uint32(b[3])<<24
}

func binaryLittleEndianUint64(b []byte) uint64 {
	return uint64(binaryLittleEndianUint32(b)) | uint64(binaryLittleEndianUint32(b[4:]))<<32
}

func binaryLittleEndianUint16(b []byte) uint16 {
	return uint16(b[0]) | uint16(b[1])<<8
}

func putBinaryLittleEndianUint32(b []byte, v uint32) {
	b[0] = byte(v)
	b[1] = byte(v >> 8)
	b[2] = byte(v >> 16)
	b[3] = byte(v >> 24)
}
