package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"time"
	"unsafe"

	"github.com/ddkwork/HyperDbg/debugger"
	"github.com/ddkwork/HyperDbg/debugger/driver"
	"github.com/ddkwork/golibrary/std/mylog"
	"github.com/ddkwork/golibrary/std/stream"
)

const (
	SMART_RCV_DRIVE_DATA          = 0x0007c088
	IOCTL_DISK_GET_DRIVE_GEOMETRY = 0x00070000
	IOCTL_STORAGE_QUERY_PROPERTY  = 0x002d1400
	IOCTL_CPUID                   = 0x00220000
	AFD_CONNECT                   = 0x12007
	IOCTL_NDIS_QUERY_GLOBAL_STATS = 0x00170202
)

func TestFixError(t *testing.T) {
	stream.Fmt(".")
	stream.Fix(".")
}

func TestRustDriverHTTP(t *testing.T) {
	driverPath := `d:\ux\examples\hypedbg\rust-driver\kd\hyperdbg_kd.sys`

	t.Log("=== 测试 Rust 驱动 HTTP API ===")

	p := driver.NewWithOptions(driverPath, "hyperdbg", "\\\\.\\hyperdbg", true)

	t.Log("安装驱动...")
	p.Install()

	t.Log("启动驱动...")
	p.Start()
	defer func() {
		t.Log("停止驱动...")
		p.Stop()
		p.Uninstall()
		t.Log("驱动已卸载")
	}()

	t.Log("等待 HTTP 服务就绪...")
	baseURL := "http://127.0.0.1:50080"
	client := &http.Client{Timeout: 5 * time.Second}

	for range 10 {
		req, _ := json.Marshal(map[string]string{"action": "ping"})
		resp, err := client.Post(baseURL+"/api", "application/json", bytes.NewReader(req))
		if err == nil {
			resp.Body.Close()
			t.Log("HTTP 服务已就绪")
			break
		}
		time.Sleep(100 * time.Millisecond)
	}

	t.Log("\n1. 测试 load_vmm...")
	loadReq, _ := json.Marshal(map[string]string{"action": "load_vmm"})
	resp, err := client.Post(baseURL+"/api", "application/json", bytes.NewReader(loadReq))
	if err != nil {
		t.Fatalf("HTTP 请求失败: %v", err)
	}
	defer resp.Body.Close()

	var result map[string]any
	json.NewDecoder(resp.Body).Decode(&result)
	t.Logf("   响应: %+v", result)

	t.Log("\n2. 测试 get_version...")
	verReq, _ := json.Marshal(map[string]string{"action": "get_version"})
	resp2, err := client.Post(baseURL+"/api", "application/json", bytes.NewReader(verReq))
	if err != nil {
		t.Logf("   失败: %v", err)
	} else {
		defer resp2.Body.Close()
		var verResult map[string]any
		json.NewDecoder(resp2.Body).Decode(&verResult)
		t.Logf("   响应: %+v", verResult)
	}

	t.Log("\n3. 测试 pause...")
	pauseReq, _ := json.Marshal(map[string]string{"action": "pause"})
	resp3, err := client.Post(baseURL+"/api", "application/json", bytes.NewReader(pauseReq))
	if err != nil {
		t.Logf("   失败: %v", err)
	} else {
		defer resp3.Body.Close()
		var pauseResult map[string]any
		json.NewDecoder(resp3.Body).Decode(&pauseResult)
		t.Logf("   响应: %+v", pauseResult)
	}

	t.Log("\n4. 测试 continue...")
	contReq, _ := json.Marshal(map[string]string{"action": "continue"})
	resp4, err := client.Post(baseURL+"/api", "application/json", bytes.NewReader(contReq))
	if err != nil {
		t.Logf("   失败: %v", err)
	} else {
		defer resp4.Body.Close()
		var contResult map[string]any
		json.NewDecoder(resp4.Body).Decode(&contResult)
		t.Logf("   响应: %+v", contResult)
	}

	t.Log("\n5. 测试 set_breakpoint...")
	bpReq, _ := json.Marshal(map[string]any{
		"action":  "set_breakpoint",
		"address": "0x7FFE0000",
		"type":    0,
	})
	resp5, err := client.Post(baseURL+"/api", "application/json", bytes.NewReader(bpReq))
	if err != nil {
		t.Logf("   失败: %v", err)
	} else {
		defer resp5.Body.Close()
		var bpResult map[string]any
		json.NewDecoder(resp5.Body).Decode(&bpResult)
		t.Logf("   响应: %+v", bpResult)
	}

	t.Log("\n=== 测试完成 ===")
}

func TestPacketAPI(t *testing.T) {
	driverPath := `d:\ux\examples\hypedbg\rust-driver\kd\hyperdbg_kd.sys`

	t.Log("=== 测试 Packet API ===")

	p := driver.NewWithOptions(driverPath, "hyperdbg", "\\\\.\\hyperdbg", true)

	t.Log("安装驱动...")
	p.Install()

	t.Log("启动驱动...")
	p.Start()
	defer func() {
		t.Log("停止驱动...")
		p.Stop()
		p.Uninstall()
		t.Log("驱动已卸载")
	}()

	t.Log("等待 HTTP 服务就绪...")
	time.Sleep(500 * time.Millisecond)

	packet := debugger.NewPacket()

	t.Log("连接 HTTP 服务...")
	if err := packet.Start(); err != nil {
		t.Fatalf("连接失败: %v", err)
	}
	defer packet.Stop()

	t.Log("连接成功!")

	t.Log("加载 VMM...")
	if err := packet.LoadVmm(); err != nil {
		t.Logf("LoadVmm: %v", err)
	}

	t.Log("发送 pause...")
	if err := packet.Pause(); err != nil {
		t.Logf("Pause 失败: %v", err)
	}

	t.Log("发送 continue...")
	if err := packet.Continue(); err != nil {
		t.Logf("Continue 失败: %v", err)
	}

	t.Log("=== 测试完成 ===")
}

func TestNotepadDebugging(t *testing.T) {
	dbg := debugger.NewUserDebug()
	defer dbg.UnloadDriver()

	mylog.Call(func() {
		dbg.StartProcess("c:\\windows\\system32\\notepad.exe")
		activeProcess := dbg.GetActiveDebuggingProcess()
		entryPoint := activeProcess.Rip
		if entryPoint == 0 {
			t.Error("无法获取入口点地址")
			return
		}
		mylog.Hex(entryPoint)
		dbg.SetBreakpoint(entryPoint)
		dbg.Continue()
		dbg.StepInto()
		dbg.ReadMemory(entryPoint, 16)
		dbg.WriteMemory(entryPoint, []byte{0x90, 0x90})
		dbg.KillProcess(activeProcess.ProcessId)
	})

	t.Log("驱动加载/卸载测试通过")
}

func TestMultipleDriverInitialization(t *testing.T) {
	t.Log("=== 多次驱动初始化测试 ===")
	t.Log("此测试验证驱动能否多次成功加载而不BSOD")

	iterations := 5
	for i := range iterations {
		t.Logf("=== 第 %d 次初始化 ===", i+1)

		dbg := debugger.NewUserDebug()

		if !dbg.IsConnected() {
			t.Errorf("第 %d 次初始化失败: 驱动未连接", i+1)
			continue
		}

		t.Logf("第 %d 次初始化成功: 驱动已连接", i+1)

		t.Logf("注意: 第 %d 次不卸载驱动，继续下一次初始化", i+1)
	}

	t.Log("=== 多次初始化测试完成 ===")
	t.Logf("共完成 %d 次初始化", iterations)
}

func TestHookScript(t *testing.T) {
	packet := debugger.NewPacket().(*debugger.Packet)
	err := packet.InstallHookScript(&debugger.HookScript{
		ApiName:  "NtDeviceIoControlFile",
		HookType: debugger.HookTypeEPT,
		Filter: &debugger.HookFilter{
			ExcludeSystem: true,
		},
		OnMatch: func(ctx *debugger.HookContext) {
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

func jsonBytes(data []byte) *jsonReader {
	return &jsonReader{data: data}
}

type jsonReader struct {
	data []byte
	pos  int
}

func (r *jsonReader) Read(p []byte) (n int, err error) {
	if r.pos >= len(r.data) {
		return 0, fmt.Errorf("EOF")
	}
	n = copy(p, r.data[r.pos:])
	r.pos += n
	return n, nil
}
