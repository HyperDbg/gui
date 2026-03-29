package debugger

import (
	"fmt"
	"unsafe"
)

type HookFilter struct {
	ProcessNames  []string
	ProcessIds    []uint32
	ExcludeSystem bool
}

type HookScriptHandler func(ctx *HookContext)

type HookScript struct {
	ApiName  string
	HookType HookType
	Filter   *HookFilter
	OnMatch  HookScriptHandler
}

type HookContext struct {
	Args   interface{}
	Return int32
}

const (
	SMART_RCV_DRIVE_DATA          = 0x0007c088
	IOCTL_DISK_GET_DRIVE_GEOMETRY = 0x00070000
	IOCTL_STORAGE_QUERY_PROPERTY  = 0x002d1400
	IOCTL_CPUID                   = 0x00220000
	AFD_CONNECT                   = 0x12007
	IOCTL_NDIS_QUERY_GLOBAL_STATS = 0x00170202
)

func Example_HookScript() {
	packet := NewPacket().(*Packet)
	err := packet.InstallHookScript(&HookScript{
		ApiName:  "NtDeviceIoControlFile",
		HookType: HookTypeEPT,
		Filter: &HookFilter{
			ExcludeSystem: true,
		},
		OnMatch: func(ctx *HookContext) {
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
			}
			swapBytes := func(dst []byte, src []byte) {
				for i := 0; i < len(src); i += 2 {
					if i+1 < len(src) && i < len(dst) {
						dst[i] = src[i+1]
						dst[i+1] = src[i]
					}
				}
			}
			args := ctx.Args.(*NtDeviceIoControlFileArgs)
			switch args.IoControlCode {
			case SMART_RCV_DRIVE_DATA, IOCTL_STORAGE_QUERY_PROPERTY:
				if args.OutputBuffer != 0 && args.OutputBufferLength >= 512 {
					info := (*IDINFO)(unsafe.Pointer(args.OutputBuffer))
					model := "Hitachi HTS545050A7E380"
					serial := "TEA55A3Q2NTK8R"
					swapBytes(info.SSerialNumber[:], []byte(serial))
					swapBytes(info.SModelNumber[:], []byte(model))
				}
				ctx.Return = 0

			case IOCTL_NDIS_QUERY_GLOBAL_STATS:
				mac := []byte{0x00, 0x11, 0x22, 0x33, 0x44, 0x55} //todo解码网卡buffer
				args.WriteOutputBytes(mac)
				ctx.Return = 0
			}
		},
	})

	packet.RegisterCpuidCallback(func(event any) {
		cpuidEvent, ok := event.(*CpuidEvent)
		if !ok {
			return
		}

		if cpuidEvent.Leaf == 1 {
			cpuidEvent.EAX = 0x000306A9
			cpuidEvent.ECX = 0x3DBAE3BF
			cpuidEvent.EDX = 0xBFEBFBFF
		}

		if cpuidEvent.Leaf == 0 {
			cpuidEvent.EBX = 0x756E6547
			cpuidEvent.ECX = 0x6C65746E
			cpuidEvent.EDX = 0x49656E69
		}
	})

	if err != nil {
		fmt.Printf("Failed to install hook script: %v\n", err)
		return
	}

	fmt.Println("Hardware spoof hooks installed")
}

func Example_HookIopXxxControlFile() {
	packet := NewPacket().(*Packet)

	err := packet.InstallHookScript(&HookScript{
		ApiName:  "IopXxxControlFile",
		HookType: HookTypeInline,
		Filter: &HookFilter{
			ExcludeSystem: true,
		},
		OnMatch: func(ctx *HookContext) {
			args := ctx.Args.(*IopXxxControlFileArgs)
			if args.IoControlCode == AFD_CONNECT {
				args.WriteOutputString("00:11:22:33:44:55")
				ctx.Return = 0
			}
		},
	})
	if err != nil {
		fmt.Printf("Failed to install IopXxxControlFile hook: %v\n", err)
		return
	}

	fmt.Println("IopXxxControlFile hook installed")
}

func Example_HookCPUDID() {
	packet := NewPacket().(*Packet)

	err := packet.InstallHookScript(&HookScript{
		ApiName:  "Cpuid",
		HookType: HookTypeEPT,
		OnMatch: func(ctx *HookContext) {
			args := ctx.Args.(*CpuidArgs)
			if args.Leaf == 0x01 {
				args.EAX = 0x000306A9
				args.ECX = 0x3DBAE3BF
				args.EDX = 0xBFEBFBFF
			}
			if args.Leaf == 0x00 {
				args.EBX = 0x756E6547
				args.ECX = 0x6C65746E
				args.EDX = 0x49656E69
			}
		},
	})
	if err != nil {
		fmt.Printf("Failed to install CPUID hook: %v\n", err)
		return
	}

	fmt.Println("CPUID hook installed")
}
