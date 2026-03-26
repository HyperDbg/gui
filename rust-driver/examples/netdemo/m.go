package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ddkwork/HyperDbg/debugger/driver"
)

type Command struct {
	Action string `json:"action"`
	Target string `json:"target,omitempty"`
	Size   int    `json:"size,omitempty"`
	Data   string `json:"data,omitempty"`
}

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
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
	p := driver.NewWithOptions(`D:\\ux\\examples\\hypedbg\\rust-driver\\examples\\netdemo\\netdemo.sys`, "netDemo", "\\\\.\\netDemo", true)

	p.Install()
	p.Start()

	fmt.Println("\n=== Step 2: HTTP Client Test ===")
	fmt.Println("Testing connection to http://127.0.0.1:50080...")

	fmt.Println("\n=== Test 1: Direct route handlers ===")

	resp, err := sendHTTPRequest("GET", "/ping", []byte{})
	if err != nil {
		fmt.Printf("Error: %v\n\n", err)
	} else {
		fmt.Printf("Response: success=%v message=%s\n\n", resp.Success, resp.Message)
	}

	resp, err = sendHTTPRequest("GET", "/status", []byte{})
	if err != nil {
		fmt.Printf("Error: %v\n\n", err)
	} else {
		fmt.Printf("Response: success=%v message=%s\n\n", resp.Success, resp.Message)
	}

	readCmd := Command{Action: "read_memory", Target: "0x1000", Size: 64}
	readData, _ := json.Marshal(readCmd)
	resp, err = sendHTTPRequest("POST", "/read_memory", readData)
	if err != nil {
		fmt.Printf("Error: %v\n\n", err)
	} else {
		fmt.Printf("Response: success=%v message=%s\n\n", resp.Success, resp.Message)
	}

	writeCmd := Command{Action: "write_memory", Target: "0x2000", Data: "deadbeef"}
	writeData, _ := json.Marshal(writeCmd)
	resp, err = sendHTTPRequest("POST", "/write_memory", writeData)
	if err != nil {
		fmt.Printf("Error: %v\n\n", err)
	} else {
		fmt.Printf("Response: success=%v message=%s\n\n", resp.Success, resp.Message)
	}

	fmt.Println("\n=== Test 2: API endpoint with action-based routing ===")

	testCases := []struct {
		cmd     Command
		comment string
	}{
		{Command{Action: "ping"}, "ping action"},
		{Command{Action: "status"}, "status action"},
		{Command{Action: "read_memory", Target: "0x1000", Size: 64}, "read_memory action"},
		{Command{Action: "write_memory", Target: "0x2000", Data: "deadbeef"}, "write_memory action"},
		{Command{Action: "read_memory"}, "read_memory without params"},
		{Command{Action: "write_memory"}, "write_memory without params"},
		{Command{Action: "unknown_command"}, "unknown action"},
	}

	for _, tc := range testCases {
		fmt.Printf("Test: %s\n", tc.comment)
		data, _ := json.Marshal(tc.cmd)

		resp, err := sendHTTPRequest("POST", "/api", data)
		if err != nil {
			fmt.Printf("Error: %v\n\n", err)
		} else {
			fmt.Printf("Response: success=%v message=%s\n\n", resp.Success, resp.Message)
		}
	}

	fmt.Println("\n=== Test 3: Invalid route ===")

	resp, err = sendHTTPRequest("GET", "/invalid", []byte{})
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
