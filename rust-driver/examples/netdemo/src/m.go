package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
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

func main() {
	fmt.Println("=== Step 1: Load Driver ===")
	p := driver.NewWithOptions(`D:\\ux\\examples\\hypedbg\\rust-driver\\examples\\netdemo\\netdemo.sys`, "netDemo", "\\\\.\\netDemo", true)

	p.Install()
	p.Start()

	fmt.Println("\n=== Step 2: TCP Client Test ===")
	addr := "127.0.0.1:50080"
	fmt.Printf("Connecting to %s...\n", addr)

	conn, err := net.DialTimeout("tcp", addr, 5*time.Second)
	if err != nil {
		fmt.Printf("Dial failed: %v\n", err)
		p.Stop()
		p.Uninstall()
		return
	}
	defer conn.Close()
	fmt.Println("Connected!")

	testCases := []Command{
		{Action: "ping"},
		{Action: "status"},
		{Action: "read_memory", Target: "0x1000", Size: 64},
		{Action: "write_memory", Target: "0x2000", Data: "deadbeef"},
		{Action: "read_memory"},
		{Action: "write_memory"},
		{Action: "unknown_command"},
	}

	for _, cmd := range testCases {
		data, _ := json.Marshal(cmd)
		data = append(data, '\n')

		_, err = conn.Write(data)
		if err != nil {
			fmt.Printf("Send failed: %v\n", err)
			continue
		}
		fmt.Printf("Sent: %s", data)

		reader := bufio.NewReader(conn)
		respData, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Read failed: %v\n", err)
			continue
		}

		var resp Response
		if err := json.Unmarshal([]byte(respData), &resp); err != nil {
			fmt.Printf("Parse failed: %v\n", err)
			continue
		}
		fmt.Printf("Response: success=%v message=%s\n\n", resp.Success, resp.Message)
	}

	fmt.Println("\n=== Step 3: Unload Driver ===")
	p.Stop()
	p.Uninstall()

	fmt.Println("\n=== Test Complete ===")
}
