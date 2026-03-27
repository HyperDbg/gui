package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ddkwork/HyperDbg/debugger/driver"
)

// Request matches the Rust-generated Request structure
type Request struct {
	Action       string  `json:"action"`
	ProcessID    *uint32 `json:"process_id,omitempty"`
	Address      *uint64 `json:"address,omitempty"`
	BreakpointID *uint64 `json:"breakpoint_id,omitempty"`
	Size         *uint32 `json:"size,omitempty"`
	Type         *int32  `json:"type,omitempty"`
	Data         []byte  `json:"data,omitempty"`
}

// Response matches the Rust-generated Response structure
type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func sendHTTPRequest(method, path string, body []byte) (Response, error) {
	var resp Response
	url := fmt.Sprintf("http://127.0.0.1:50080%s", path)

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return resp, fmt.Errorf("create request failed: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Host", "127.0.0.1")

	fmt.Printf("Sending %s %s\n", method, path)
	if len(body) > 0 {
		fmt.Printf("Body: %s\n", string(body))
	}

	response, err := client.Do(req)
	if err != nil {
		return resp, fmt.Errorf("send failed: %v", err)
	}
	defer response.Body.Close()

	fmt.Printf("Status: %s\n", response.Status)
	for key, values := range response.Header {
		for _, value := range values {
			fmt.Printf("Header: %s: %s\n", key, value)
		}
	}

	if err := json.NewDecoder(response.Body).Decode(&resp); err != nil {
		return resp, fmt.Errorf("parse failed: %v", err)
	}

	return resp, nil
}

func main() {
	fmt.Println("=== Step 1: Load Driver ===")
	p := driver.NewWithOptions(`D:\\ux\\examples\\hypedbg\\HyperDbg_rust\\rust-driver\\examples\\netdemo\\netdemo.sys`, "netDemo", "\\\\.\\netDemo", true)

	p.Install()
	p.Start()

	fmt.Println("\n=== Step 2: HTTP Client Test ===")
	fmt.Println("Testing connection to http://127.0.0.1:50080...")

	fmt.Println("\n=== Test 1: API endpoint with generated NoOpDebugger ===")

	testCases := []struct {
		req     Request
		comment string
	}{
		{Request{Action: "initialize"}, "initialize action"},
		{Request{Action: "terminate"}, "terminate action"},
		{Request{Action: "attach_process", ProcessID: new(uint32(1234))}, "attach_process action"},
		{Request{Action: "detach_process"}, "detach_process action"},
		{Request{Action: "pause"}, "pause action"},
		{Request{Action: "continue_"}, "continue action"},
		{Request{Action: "step_into"}, "step_into action"},
		{Request{Action: "read_memory", Address: new(uint64(0x1000)), Size: new(uint32(64))}, "read_memory action"},
		{Request{Action: "write_memory", Address: new(uint64(0x2000)), Data: []byte{0xDE, 0xAD, 0xBE, 0xEF}}, "write_memory action"},
		{Request{Action: "get_registers"}, "get_registers action"},
		{Request{Action: "set_registers"}, "set_registers action"},
		{Request{Action: "unknown_command"}, "unknown action"},
	}

	for _, tc := range testCases {
		fmt.Printf("Test: %s\n", tc.comment)
		data, _ := json.Marshal(tc.req)

		resp, err := sendHTTPRequest("POST", "/api", data)
		if err != nil {
			fmt.Printf("Error: %v\n\n", err)
		} else {
			fmt.Printf("Response: success=%v message=%s\n\n", resp.Success, resp.Message)
		}
	}

	fmt.Println("\n=== Test 2: Invalid route ===")

	resp, err := sendHTTPRequest("GET", "/invalid", []byte{})
	if err != nil {
		fmt.Printf("Error: %v\n\n", err)
	} else {
		fmt.Printf("Response: success=%v message=%s\n\n", resp.Success, resp.Message)
	}

	// 等待请求处理完成
	fmt.Println("\n=== Step 3: Waiting for requests to complete ===")
	fmt.Println("Waiting for driver to process requests...")
	time.Sleep(3 * time.Second)

	fmt.Println("\n=== Step 4: Unload Driver ===")
	p.Stop()
	p.Uninstall()

	fmt.Println("\n=== Test Complete ===")
}
