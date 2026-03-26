package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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

	if err := json.NewDecoder(response.Body).Decode(&resp); err != nil {
		return resp, fmt.Errorf("parse failed: %v", err)
	}

	return resp, nil
}

func main() {
	driverPath := flag.String("driver", `D:\ux\examples\hypedbg\HyperDbg_rust\rust-driver\target\x86_64-pc-windows-msvc\release\hyperdbg_driver.sys`, "Driver path")
	driverName := flag.String("name", "hyperdbg", "Driver name")
	deviceName := flag.String("device", "\\\\.\\hyperdbg", "Device name")
	verbose := flag.Bool("v", true, "Verbose output")
	flag.Parse()

	fmt.Println("=== HyperDbg Communicator ===")
	fmt.Printf("Driver: %s\n", *driverPath)

	fmt.Println("\n=== Step 1: Load Driver ===")
	p := driver.NewWithOptions(*driverPath, *driverName, *deviceName, *verbose)

	if err := p.Install(); err != nil {
		fmt.Printf("Install failed: %v\n", err)
		os.Exit(1)
	}

	if err := p.Start(); err != nil {
		fmt.Printf("Start failed: %v\n", err)
		p.Uninstall()
		os.Exit(1)
	}

	fmt.Println("Driver loaded successfully")

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigChan
		fmt.Println("\n=== Shutting down ===")
		p.Stop()
		p.Uninstall()
		os.Exit(0)
	}()

	fmt.Println("\n=== Step 2: Connect to Driver ===")
	fmt.Println("Testing connection to http://127.0.0.1:50080...")

	time.Sleep(1 * time.Second)

	resp, err := sendHTTPRequest("GET", "/ping", nil)
	if err != nil {
		fmt.Printf("Connection failed: %v\n", err)
		p.Stop()
		p.Uninstall()
		os.Exit(1)
	}
	fmt.Printf("Response: success=%v message=%s\n\n", resp.Success, resp.Message)

	fmt.Println("=== Step 3: Initialize Debugger ===")
	server := NewServer()

	if err := server.Start(50080); err != nil {
		fmt.Printf("Server start failed: %v\n", err)
		p.Stop()
		p.Uninstall()
		os.Exit(1)
	}

	fmt.Println("Debugger initialized successfully")

	server.RegisterCallback(MsgTypeBreakpointEvent, func(event interface{}) {
		fmt.Printf("[Event] Breakpoint: %+v\n", event)
	})

	server.RegisterCallback(MsgTypeExceptionEvent, func(event interface{}) {
		fmt.Printf("[Event] Exception: %+v\n", event)
	})

	server.RegisterCallback(MsgTypeDebugPrintEvent, func(event interface{}) {
		fmt.Printf("[Event] DebugPrint: %+v\n", event)
	})

	fmt.Println("\n=== Step 4: Test Commands ===")

	testCommands := []struct {
		name string
		cmd  Command
	}{
		{"ping", Command{Action: "ping"}},
		{"status", Command{Action: "status"}},
		{"read_memory", Command{Action: "read_memory", Target: "0x1000", Size: 64}},
		{"write_memory", Command{Action: "write_memory", Target: "0x2000", Data: "deadbeef"}},
	}

	for _, tc := range testCommands {
		fmt.Printf("\nTest: %s\n", tc.name)
		data, _ := json.Marshal(tc.cmd)

		resp, err := sendHTTPRequest("POST", "/api", data)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		} else {
			fmt.Printf("Response: success=%v message=%s\n", resp.Success, resp.Message)
		}
	}

	fmt.Println("\n=== Running ===")
	fmt.Println("Press Ctrl+C to exit")

	drivers := server.GetConnectedDrivers()
	if len(drivers) > 0 {
		fmt.Printf("Connected drivers: %v\n", drivers)

		if err := server.InitializeDriver(drivers[0]); err != nil {
			fmt.Printf("Initialize failed: %v\n", err)
		} else {
			fmt.Println("Driver initialized")
		}
	}

	for server.IsConnected() {
		event := server.WaitForEvent(1 * time.Second)
		if event != nil {
			fmt.Printf("Event received: %+v\n", event)
		}
	}

	fmt.Println("\n=== Step 5: Cleanup ===")
	server.Stop()
	p.Stop()
	p.Uninstall()

	fmt.Println("\n=== Test Complete ===")
}
