package hardwareinfo

import (
	"syscall"
	"unsafe"

	"github.com/ddkwork/golibrary/std/mylog"
)

const (
	IOCTL_STORAGE_QUERY_PROPERTY  = 0x2D1400
	IOCTL_DISK_GET_DRIVE_GEOMETRY = 0x70000
)

type HDiskInfo struct {
	module   []byte   // 40 型号
	firmware [8]byte  // 固件版本
	serialno [20]byte // 序列号
	capacity uint32   // 容量
}

func (h HDiskInfo) String() {
	type info struct {
		Module   string
		Firmware string
		Serialno string
		Capacity uint32
	}

	i := info{
		Module:   string(h.module),
		Firmware: string(h.firmware[:]),
		Serialno: string(h.serialno[:]),
		Capacity: h.capacity,
	}
	mylog.MarshalJson("capacity", i)
}

type STORAGE_PROPERTY_QUERY struct {
	PropertyId           uint32
	QueryType            uint32
	AdditionalParameters [1]byte
}

type STORAGE_DEVICE_DESCRIPTOR struct {
	Version               uint32
	Size                  uint32
	DeviceType            byte
	DeviceTypeModifier    byte
	RemovableMedia        bool
	CommandQueueing       bool
	VendorIdOffset        uint32
	ProductIdOffset       uint32
	ProductRevisionOffset uint32
	SerialNumberOffset    uint32
	BusType               byte
	RawPropertiesLength   uint32
}

type DISK_GEOMETRY struct {
	Cylinders         int64
	MediaType         uint32
	TracksPerCylinder uint32
	SectorsPerTrack   uint32
	BytesPerSector    uint32
}

func readHarddiskInfo(pinfo *HDiskInfo) {
	hDevice := mylog.Check2(syscall.CreateFile(
		syscall.StringToUTF16Ptr(`\\.\PhysicalDrive0`),
		syscall.GENERIC_READ|syscall.GENERIC_WRITE,
		syscall.FILE_SHARE_READ|syscall.FILE_SHARE_WRITE,
		nil,
		syscall.OPEN_EXISTING,
		0,
		0,
	))
	defer syscall.CloseHandle(hDevice)
	query := STORAGE_PROPERTY_QUERY{
		PropertyId: 0,
		QueryType:  0,
	}
	buffer := make([]byte, 1024)
	var bytesReturned uint32
	mylog.Check(syscall.DeviceIoControl(
		hDevice,
		IOCTL_STORAGE_QUERY_PROPERTY,
		(*byte)(unsafe.Pointer(&query)),
		uint32(unsafe.Sizeof(query)),
		&buffer[0],
		uint32(len(buffer)),
		&bytesReturned,
		nil,
	))
	deviceDescriptor := (*STORAGE_DEVICE_DESCRIPTOR)(unsafe.Pointer(&buffer[0]))
	if deviceDescriptor.SerialNumberOffset != 0 {
		copy(pinfo.serialno[:], buffer[deviceDescriptor.SerialNumberOffset:])
	}
	if deviceDescriptor.ProductIdOffset != 0 {
		all := buffer[deviceDescriptor.ProductIdOffset:]
		end := 0
		size := 0
		for i, b := range all {
			if b == 0 { // c string end
				end = i + int(deviceDescriptor.ProductIdOffset)
				all = buffer[deviceDescriptor.ProductIdOffset:end]
				size = i
				break
			}
		}
		pinfo.module = make([]byte, size)
		copy(pinfo.module[:size], all)
	}
	if deviceDescriptor.ProductRevisionOffset != 0 {
		copy(pinfo.firmware[:], buffer[deviceDescriptor.ProductRevisionOffset:])
	}
	geom := DISK_GEOMETRY{}
	mylog.Check(syscall.DeviceIoControl(
		hDevice,
		IOCTL_DISK_GET_DRIVE_GEOMETRY,
		nil,
		0,
		(*byte)(unsafe.Pointer(&geom)),
		uint32(unsafe.Sizeof(geom)),
		&bytesReturned,
		nil,
	))
	pinfo.capacity = uint32(geom.Cylinders) * geom.TracksPerCylinder * geom.SectorsPerTrack * geom.BytesPerSector / (1024 * 1024)
}

func nvme() {
	var hddInfo HDiskInfo
	readHarddiskInfo(&hddInfo)
	hddInfo.String()
}
