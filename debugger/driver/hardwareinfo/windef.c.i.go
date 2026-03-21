package hardwareinfo

import (
	"os"
	"unsafe"
)

type (
	DWORD          = uint32
	BOOL           = int32
	BYTE           = uint8
	WORD           = uint16
	FLOAT          = float32
	PFLOAT         = *float32
	INT            = int32
	UINT           = uint32
	PUINT          = *uint32
	PBOOL          = *int32
	LPBOOL         = *int32
	PBYTE          = *uint8
	LPBYTE         = *uint8
	PINT           = *int32
	LPINT          = *int32
	PWORD          = *uint16
	LPWORD         = *uint16
	LPLONG         = *int32
	PDWORD         = *uint32
	LPDWORD        = *uint32
	LPVOID         = unsafe.Pointer
	LPCVOID        = unsafe.Pointer
	ULONG          = uint32
	PULONG         = *uint32
	USHORT         = uint16
	PUSHORT        = *uint16
	UCHAR          = uint8
	PUCHAR         = *uint8
	CHAR           = int8
	SHORT          = int16
	LONG           = int32
	_cgoa_1_windef struct {
		Xbf_0 uint16
	}
)

type _cgoa_2_windef struct {
	Xbf_0 uint16
}
type _cgoa_3_windef struct {
	Xbf_0 uint16
}
type _cgoa_4_windef struct {
	Xbf_0 uint16
}
type _cgoa_5_windef struct {
	Xbf_0 uint16
}
type _cgoa_6_windef struct {
	Xbf_0 uint16
}
type _cgoa_7_windef struct {
	Xbf_0 uint16
}
type struct__IDINFO struct {
	GenConfig             uint16
	NumCyls               uint16
	Reserved2             uint16
	NumHeads              uint16
	Reserved4             uint16
	Reserved5             uint16
	NumSectorsPerTrack    uint16
	VendorUnique          [3]uint16
	SerialNumber          [20]uint8
	BufferType            uint16
	BufferSize            uint16
	ECCSize               uint16
	FirmwareRev           [8]uint8
	ModelNumber           [40]uint8
	MoreVendorUnique      uint16
	Reserved48            uint16
	Capabilities          _cgoa_1_windef
	Reserved1             uint16
	PIOTiming             uint16
	DMATiming             uint16
	FieldValidity         _cgoa_2_windef
	NumCurCyls            uint16
	NumCurHeads           uint16
	NumCurSectorsPerTrack uint16
	CurSectorsLow         uint16
	CurSectorsHigh        uint16
	MultSectorStuff       _cgoa_3_windef
	dwTotalSectors        uint32
	SingleWordDMA         uint16
	MultiWordDMA          _cgoa_4_windef
	PIOCapacity           _cgoa_5_windef
	MinMultiWordDMACycle  uint16
	RecMultiWordDMACycle  uint16
	MinPIONoFlowCycle     uint16
	MinPOIFlowCycle       uint16
	Reserved69            [11]uint16
	MajorVersion          _cgoa_6_windef
	MinorVersion          uint16
	Reserved82            [6]uint16
	UltraDMA              _cgoa_7_windef
	Reserved89            [167]uint16
}
type (
	IDINFO               = struct__IDINFO
	PIDINFO              = *struct__IDINFO
	struct__DRIVERSTATUS struct {
		DriverError uint8
		IDEError    uint8
		Reserved    [2]uint8
		dwReserved  [2]uint32
	}
)

type (
	DRIVERSTATUS             = struct__DRIVERSTATUS
	PDRIVERSTATUS            = *struct__DRIVERSTATUS
	LPDRIVERSTATUS           = *struct__DRIVERSTATUS
	struct__SENDCMDOUTPARAMS struct {
		BufferSize   uint32
		DriverStatus struct__DRIVERSTATUS
		Buffer       [1]uint8
	}
)

type (
	SENDCMDOUTPARAMS           = struct__SENDCMDOUTPARAMS
	PSENDCMDOUTPARAMS          = *struct__SENDCMDOUTPARAMS
	LPSENDCMDOUTPARAMS         = *struct__SENDCMDOUTPARAMS
	struct__GETVERSIONINPARAMS struct {
		Version       uint8
		Revision      uint8
		Reserved1     uint8
		IDEDeviceMap  uint8
		fCapabilities uint32
		Reserved2     [4]uint32
	}
)

type (
	GETVERSIONINPARAMS   = struct__GETVERSIONINPARAMS
	PGETVERSIONINPARAMS  = *struct__GETVERSIONINPARAMS
	LPGETVERSIONINPARAMS = *struct__GETVERSIONINPARAMS
	struct__IDEREGS      struct {
		FeaturesReg     uint8
		SectorCountReg  uint8
		SectorNumberReg uint8
		CylLowReg       uint8
		CylHighReg      uint8
		DriveHeadReg    uint8
		CommandReg      uint8
		Reserved        uint8
	}
)

type (
	IDEREGS                 = struct__IDEREGS
	PIDEREGS                = *struct__IDEREGS
	LPIDEREGS               = *struct__IDEREGS
	struct__SENDCMDINPARAMS struct {
		BufferSize  uint32
		irDriveRegs struct__IDEREGS
		DriveNumber uint8
		Reserved1   [3]uint8
		Reserved2   [4]uint32
		Buffer      [1]uint8
	}
)

type (
	SENDCMDINPARAMS   = struct__SENDCMDINPARAMS
	PSENDCMDINPARAMS  = *struct__SENDCMDINPARAMS
	LPSENDCMDINPARAMS = *struct__SENDCMDINPARAMS
)

func _cgo_main() int32 {
	return int32(0)
}

func main() {
	os.Exit(int(_cgo_main()))
}
