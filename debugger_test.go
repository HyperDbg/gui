package main

import (
	"fmt"
	"testing"
	"time"
	"unsafe"

	"github.com/ddkwork/HyperDbg/debugger"
	"github.com/ddkwork/HyperDbg/debugger/driver"
	"github.com/ddkwork/golibrary/std/stream"
)

func TestFixError(t *testing.T) {
	stream.Fmt(".")
	stream.Fix(".")
}

func TestRustDriverHTTP(t *testing.T) {
	driverPath := `d:\ux\examples\hypedbg\rust-driver\kd\hyperdbg_kd.sys`

	drv := driver.NewWithOptions(driverPath, "hyperdbg", "\\\\.\\hyperdbg", true)
	drv.Install()
	drv.Start()
	defer func() {
		drv.Stop()
		drv.Uninstall()
	}()

	conn := debugger.NewPacket("http://127.0.0.1:50080")

	if err := conn.Connect(); err != nil {
		t.Fatalf("Connect 失败: %v", err)
	}
	defer conn.Disconnect()

	notepadPID, err := conn.StartProcess("notepad.exe")
	if err != nil {
		t.Fatalf("StartProcess 失败: %v", err)
	}
	defer conn.KillProcess(notepadPID)

	time.Sleep(500 * time.Millisecond)

	if err := conn.LoadVmm(); err != nil {
		t.Fatalf("LoadVmm 失败: %v", err)
	}

	if err := conn.AttachProcess(notepadPID); err != nil {
		t.Fatalf("AttachProcess 失败: %v", err)
	}

	if err := conn.SetBreakpoint(0x7FFE0000, debugger.BreakpointSoftware); err != nil {
		t.Fatalf("SetBreakpoint 失败: %v", err)
	}

	if err := conn.Continue(); err != nil {
		t.Fatalf("Continue 失败: %v", err)
	}
}

func TestMultipleDriverInitialization(t *testing.T) {
	driverPath := `d:\ux\examples\hypedbg\rust-driver\kd\hyperdbg_kd.sys`

	iterations := 5
	for i := range iterations {
		drv := driver.NewWithOptions(driverPath, "hyperdbg", "\\\\.\\hyperdbg", true)
		drv.Install()
		drv.Start()

		time.Sleep(200 * time.Millisecond)

		conn := debugger.NewPacket("http://127.0.0.1:50080")
		if err := conn.Connect(); err != nil {
			t.Errorf("第 %d 次初始化失败: %v", i+1, err)
			drv.Stop()
			drv.Uninstall()
			continue
		}

		if !conn.IsConnected() {
			t.Errorf("第 %d 次初始化失败: 驱动未连接", i+1)
			drv.Stop()
			drv.Uninstall()
			continue
		}

		conn.Disconnect()
		drv.Stop()
		drv.Uninstall()
	}
}

func TestHookScript(t *testing.T) {
	packet := debugger.NewPacket("http://127.0.0.1:50080").(*debugger.Packet)
	err := packet.InstallHookScript(&debugger.HookScript{
		ApiName:  "NtDeviceIoControlFile",
		HookType: debugger.HookTypeEPT,
		Filter: &debugger.HookFilter{
			ExcludeSystem: true,
		},
		OnMatch: func(ctx *debugger.HookContext) {
			const (
				SMART_RCV_DRIVE_DATA          = 0x0007c088
				IOCTL_DISK_GET_DRIVE_GEOMETRY = 0x00070000
				IOCTL_STORAGE_QUERY_PROPERTY  = 0x002d1400
				IOCTL_CPUID                   = 0x00220000
				AFD_CONNECT                   = 0x12007
				IOCTL_NDIS_QUERY_GLOBAL_STATS = 0x00170202
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
			}
			swapBytes := func(dst []byte, src []byte) {
				for i := 0; i < len(src); i += 2 {
					if i+1 < len(src) && i < len(dst) {
						dst[i] = src[i+1]
						dst[i+1] = src[i]
					}
				}
			}
			args := ctx.Args.(*debugger.NtDeviceIoControlFileArgs)
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
			}
		},
	})
	if err != nil {
		fmt.Printf("Failed to install hook script: %v\n", err)
		return
	}

	fmt.Println("Hardware spoof hooks installed")
}
