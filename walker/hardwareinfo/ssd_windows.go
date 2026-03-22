//go:build windows

package hardwareinfo

import (
	"fmt"
	"reflect"
	"strconv"
	"syscall"
	"unsafe"

	"github.com/ddkwork/ddk/windef"
	"github.com/ddkwork/golibrary/std/mylog"
	"github.com/ddkwork/golibrary/std/stream"
)

type ssdInfo struct {
	SerialNumber string
	ModelNumber  string
	Version      string
}

var (
	kernel32, _             = syscall.LoadLibrary("kernel32.dll")
	globalMemoryStatusEx, _ = syscall.GetProcAddress(kernel32, "GlobalMemoryStatusEx")
)

const (
	IDENTIFY_BUFFER_SIZE = 512
	ID_CMD               = 0xEC
	ATAPI_ID_CMD         = 0xA1
	SMART_CMD            = 0xB0

	DFP_GET_VERSION        = 0x00074080
	DFP_SEND_DRIVE_COMMAND = 0x0007c084
	DFP_RECEIVE_DRIVE_DATA = 0x0007c088
)

func (s *ssdInfo) Get() (ok bool) {
	path := fmt.Sprintf("\\\\.\\PhysicalDrive%d", 0)
	fromString := mylog.Check2(syscall.UTF16PtrFromString(path))

	handle := mylog.Check2(syscall.CreateFile(
		fromString,
		syscall.GENERIC_READ|syscall.GENERIC_WRITE,
		syscall.FILE_SHARE_READ|syscall.FILE_SHARE_WRITE,
		nil,
		syscall.OPEN_EXISTING,
		syscall.FILE_ATTRIBUTE_NORMAL,
		0,
	))
	// if (hDevice == INVALID_HANDLE_VALUE)
	outBuffer := make([]byte, 528)
	var bytesReturned uint32
	mylog.Check(syscall.DeviceIoControl( // todo check mapid return is 1 ?
		handle,
		windef.SMART_GET_VERSION,
		nil,
		0,
		&outBuffer[0],
		528,
		// IDENTIFY_BUFFER_SIZE,
		&bytesReturned,
		nil,
	))
	getVersionInParam := (*struct__GETVERSIONINPARAMS)(unsafe.Pointer(&outBuffer[0]))
	mylog.MarshalJson(*getVersionInParam)
	if getVersionInParam.IDEDeviceMap == 1 { // ? <=0
	}
	var BytesReturned uint32
	sendcmdinparams := struct__SENDCMDINPARAMS{
		BufferSize: 0,
		irDriveRegs: struct__IDEREGS{
			FeaturesReg:     0,
			SectorCountReg:  0,
			SectorNumberReg: 0,
			CylLowReg:       0,
			CylHighReg:      0,
			DriveHeadReg:    0,
			CommandReg:      ID_CMD,
			Reserved:        0,
		},
		DriveNumber: 0,
		Reserved1:   [3]uint8{},
		Reserved2:   [4]uint32{},
		Buffer:      [1]uint8{},
	}
	// marshal := mylog.Check2(cstruct.Marshal(&sendcmdinparams))

	// mylog.HexDump("sendcmdinparams", marshal)
	mylog.Info("unsafe.Sizeof(sendcmdinparams)", unsafe.Sizeof(sendcmdinparams))
	// mylog.Info("len(sendcmdinparams)", len(marshal))

	fnStructToBytes := func() (b []byte) { // more than big one c memory,because memory align
		header := reflect.SliceHeader{
			Data: uintptr(unsafe.Pointer(&sendcmdinparams)),
			Len:  int(unsafe.Sizeof(sendcmdinparams)),
			Cap:  int(unsafe.Sizeof(sendcmdinparams)),
		}
		return *(*[]byte)(unsafe.Pointer(&header))
	}
	mylog.HexDump(fnStructToBytes())

	mylog.Hex(windef.SMART_RCV_DRIVE_DATA)
	// mylog.HexDump("marshal", marshal)
	mylog.Check(syscall.DeviceIoControl(
		handle,
		windef.SMART_RCV_DRIVE_DATA,
		(*byte)(unsafe.Pointer(&sendcmdinparams)),
		// &marshal[0],
		32,
		&outBuffer[0],
		528,
		&BytesReturned,
		nil,
	))
	outParams_ := (*struct__SENDCMDOUTPARAMS)(unsafe.Pointer(&outBuffer[0]))
	b := outParams_.Buffer[:]
	mylog.HexDump(b)

	info := (*struct__IDINFO)(unsafe.Pointer(&b[0]))
	SerialNumber := stream.SwapAdjacent(stream.NewBuffer(info.SerialNumber[:]).String())
	ModelNumber := stream.SwapAdjacent(stream.NewBuffer(info.ModelNumber[:]).String())
	FirmwareRev := stream.SwapAdjacent(stream.NewBuffer(info.FirmwareRev[:]).String())

	mylog.Info(strconv.Quote(SerialNumber.String()))
	mylog.Info(strconv.Quote(ModelNumber.String()))
	mylog.Info(strconv.Quote(FirmwareRev.String()))
	*s = ssdInfo{
		SerialNumber: SerialNumber.String(),
		ModelNumber:  ModelNumber.String(),
		Version:      FirmwareRev.String(),
	}
	return true
}

// https://github.com/gioui/gio/blob/main/internal/byteslice/byteslice.go
func Struct(s any) []byte {
	v := reflect.ValueOf(s)
	sz := int(v.Elem().Type().Size())
	return unsafe.Slice((*byte)(unsafe.Pointer(v.Pointer())), sz)
}

// https://github.com/alkemir/winsmart-go
