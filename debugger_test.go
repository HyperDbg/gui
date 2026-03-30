package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/ddkwork/HyperDbg/debugger"
	"github.com/ddkwork/HyperDbg/debugger/driver"
)

func TestRustDriverHTTP(t *testing.T) {
	driverPath := `D:\ux\examples\hypedbg\HyperDbg_rust\rust-driver\kd\hyperdbg_kd.sys`

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
	driverPath := `D:\ux\examples\hypedbg\HyperDbg_rust\rust-driver\kd\hyperdbg_kd.sys`

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
	time.Sleep(500 * time.Millisecond) // ListenAndServe 是同步的，服务已在运行

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
