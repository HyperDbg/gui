//go:build windows

package main

import (
	"fmt"
	"syscall"
	"unsafe"

	"github.com/ddkwork/golibrary/std/mylog"
)

const (
	SMART_GET_VERSION    = 0xc08d0020
	SMART_RCV_DRIVE_DATA = 0x7c088
)

type IDINFO struct {
	WGenConfig          uint16
	WNumCyls            uint16
	WReserved2          uint16
	WNumHeads           uint16
	WReserved4          uint16
	WReserved5          uint16
	WNumSectorsPerTrack uint16
	WVendorUnique       [3]uint16
	SSerialNumber       [20]byte
	WBufferType         uint16
	WBufferSize         uint16
	WECCSize            uint16
	SFirmwareRev        [8]byte
	SModelNumber        [40]byte
	WMoreVendorUnique   uint16
	WReserved48         uint16
	WCapabilities       struct {
		Reserved1 uint16 `struc:"bits=8"`
		DMA       uint16 `struc:"bits=1"`
		LBA       uint16 `struc:"bits=1"`
		DisIORDY  uint16 `struc:"bits=1"`
		IORDY     uint16 `struc:"bits=1"`
		SoftReset uint16 `struc:"bits=1"`
		Overlap   uint16 `struc:"bits=1"`
		Queue     uint16 `struc:"bits=1"`
		InlDMA    uint16 `struc:"bits=1"`
	}
	WReserved1     uint16
	WPIOTiming     uint16
	WDMATiming     uint16
	WFieldValidity struct {
		CHSNumber   uint16 `struc:"bits=1"`
		CycleNumber uint16 `struc:"bits=1"`
		UnltraDMA   uint16 `struc:"bits=1"`
		Reserved    uint16 `struc:"bits=13"`
	}
	WNumCurCyls            uint16
	WNumCurHeads           uint16
	WNumCurSectorsPerTrack uint16
	WCurSectorsLow         uint16
	WCurSectorsHigh        uint16
	WMultSectorStuff       struct {
		CurNumber uint16 `struc:"bits=8"`
		Multi     uint16 `struc:"bits=1"`
		Reserved1 uint16 `struc:"bits=7"`
	}
	DwTotalSectors uint32
	WSingleWordDMA uint16
	WMultiWordDMA  struct {
		Mode0     uint16 `struc:"bits=1"`
		Mode1     uint16 `struc:"bits=1"`
		Mode2     uint16 `struc:"bits=1"`
		Reserved1 uint16 `struc:"bits=5"`
		Mode0Sel  uint16 `struc:"bits=1"`
		Mode1Sel  uint16 `struc:"bits=1"`
		Mode2Sel  uint16 `struc:"bits=1"`
		Reserved2 uint16 `struc:"bits=5"`
	}
	WPIOCapacity struct {
		AdvPOIModes uint16 `struc:"bits=8"`
		Reserved    uint16 `struc:"bits=8"`
	}
	WMinMultiWordDMACycle uint16
	WRecMultiWordDMACycle uint16
	WMinPIONoFlowCycle    uint16
	WMinPOIFlowCycle      uint16
	WReserved69           [11]uint16
	WMajorVersion         struct {
		Reserved1 uint16 `struc:"bits=1"`
		ATA1      uint16 `struc:"bits=1"`
		ATA2      uint16 `struc:"bits=1"`
		ATA3      uint16 `struc:"bits=1"`
		ATA4      uint16 `struc:"bits=1"`
		ATA5      uint16 `struc:"bits=1"`
		ATA6      uint16 `struc:"bits=1"`
		ATA7      uint16 `struc:"bits=1"`
		ATA8      uint16 `struc:"bits=1"`
		ATA9      uint16 `struc:"bits=1"`
		ATA10     uint16 `struc:"bits=1"`
		ATA11     uint16 `struc:"bits=1"`
		ATA12     uint16 `struc:"bits=1"`
		ATA13     uint16 `struc:"bits=1"`
		ATA14     uint16 `struc:"bits=1"`
		Reserved2 uint16 `struc:"bits=1"`
	}
	WMinorVersion uint16
	WReserved82   [6]uint16
	WUltraDMA     struct {
		Mode0    uint16 `struc:"bits=1"`
		Mode1    uint16 `struc:"bits=1"`
		Mode2    uint16 `struc:"bits=1"`
		Mode3    uint16 `struc:"bits=1"`
		Mode4    uint16 `struc:"bits=1"`
		Mode5    uint16 `struc:"bits=1"`
		Mode6    uint16 `struc:"bits=1"`
		Mode7    uint16 `struc:"bits=1"`
		Mode0Sel uint16 `struc:"bits=1"`
		Mode1Sel uint16 `struc:"bits=1"`
		Mode2Sel uint16 `struc:"bits=1"`
		Mode3Sel uint16 `struc:"bits=1"`
		Mode4Sel uint16 `struc:"bits=1"`
		Mode5Sel uint16 `struc:"bits=1"`
		Mode6Sel uint16 `struc:"bits=1"`
		Mode7Sel uint16 `struc:"bits=1"`
	}
	WReserved89 [167]uint16
}

func hexdump(title string, data []byte) {
	fmt.Println(title)
	for i, b := range data {
		if i%16 == 0 {
			if i != 0 {
				fmt.Println()
			}
			fmt.Printf("%08X  ", i)
		}
		fmt.Printf("%02X ", b)
	}
	fmt.Println()
}

func exchangeChar(in, out []byte) {
	if len(in)%2 != 0 {
		return
	}
	for i := 0; i < len(in); i += 2 {
		out[i], out[i+1] = in[i+1], in[i]
	}
}

func main() {
	mylog.Call(ssd)
}

func ssd() {
	hDevice := mylog.Check2(syscall.CreateFile(
		syscall.StringToUTF16Ptr("\\\\.\\PhysicalDrive0"),
		syscall.GENERIC_READ|syscall.GENERIC_WRITE,
		syscall.FILE_SHARE_READ|syscall.FILE_SHARE_WRITE,
		nil,
		syscall.OPEN_EXISTING,
		0,
		0,
	))

	defer func() { mylog.Check(syscall.CloseHandle(hDevice)) }()

	var getVersion struct {
		BCmdSupport   uint8
		BIDEDeviceMap uint8
		BSmartSupport uint8
	}

	var bytesReturned uint32
	mylog.Check(syscall.DeviceIoControl(
		hDevice,
		SMART_GET_VERSION,
		nil, 0,
		(*byte)(unsafe.Pointer(&getVersion)), uint32(unsafe.Sizeof(getVersion)),
		&bytesReturned,
		nil,
	))

	inBuffer := struct {
		IRDriveRegs struct {
			BFeaturesReg     uint8
			BSectorCountReg  uint8
			BSectorNumberReg uint8
			BSectorLowReg    uint8
			BSectorHighReg   uint8
			BDriveHeadReg    uint8
			BCommandReg      uint8
			BReserved        uint8
		}
		BDirection uint8
		BReserved  [3]uint8
	}{
		IRDriveRegs: struct {
			BFeaturesReg     uint8
			BSectorCountReg  uint8
			BSectorNumberReg uint8
			BSectorLowReg    uint8
			BSectorHighReg   uint8
			BDriveHeadReg    uint8
			BCommandReg      uint8
			BReserved        uint8
		}{
			BCommandReg: func() uint8 {
				if getVersion.BIDEDeviceMap&0x10 != 0 {
					return 0xA1 // ATAPI_IDENTIFY_DEVICE
				}
				return 0xEC // IDENTIFY_DEVICE
			}(),
		},
	}

	outBuffer := make([]byte, 512+32)
	mylog.Check(syscall.DeviceIoControl(
		hDevice,
		SMART_RCV_DRIVE_DATA,
		(*byte)(unsafe.Pointer(&inBuffer)), uint32(unsafe.Sizeof(inBuffer))-1,
		(*byte)(unsafe.Pointer(&outBuffer[0])), uint32(len(outBuffer)),
		&bytesReturned,
		nil,
	))

	outPtr := (*IDINFO)(unsafe.Pointer(&outBuffer[0]))

	var diskID, diskModel, sSerial, sModel [40]byte
	exchangeChar(outPtr.SSerialNumber[:], diskID[:])
	exchangeChar(outPtr.SModelNumber[:], diskModel[:])
	exchangeChar(outPtr.SSerialNumber[:], sSerial[:])
	exchangeChar(outPtr.SModelNumber[:], sModel[:])

	fmt.Printf("Disk ID: %s\n", diskID)
	fmt.Printf("Disk Model: %s\n", diskModel)
}
