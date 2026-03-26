package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/ddkwork/HyperDbg/debugger/driver"
)

func main() {
	driverPath := flag.String("driver", `D:\ux\examples\hypedbg\HyperDbg_rust\rust-driver\target\x86_64-pc-windows-msvc\release\hyperdbg_driver.sys`, "Driver path")
	driverName := flag.String("name", "hyperdbg", "Driver name")
	deviceName := flag.String("device", "\\\\.\\hyperdbg", "Device name")
	verbose := flag.Bool("v", true, "Verbose output")
	flag.Parse()

	fmt.Println("=== HyperDbg Communicator ===")
	fmt.Printf("Driver: %s\n", *driverPath)

	p := driver.NewWithOptions(*driverPath, *driverName, *deviceName, *verbose)
	p.Install()
	p.Start()
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

	packet := NewPacket()
	if err := packet.Start(); err != nil {
		fmt.Printf("Packet start failed: %v\n", err)
		p.Stop()
		p.Uninstall()
		os.Exit(1)
	}
	defer packet.Stop()

	fmt.Println("Debugger initialized successfully")

	packet.RegisterCallback(MsgTypeBreakpointEvent, func(event interface{}) {
		fmt.Printf("[Event] Breakpoint: %+v\n", event)
	})

	packet.RegisterCallback(MsgTypeExceptionEvent, func(event interface{}) {
		fmt.Printf("[Event] Exception: %+v\n", event)
	})

	packet.RegisterCallback(MsgTypeDebugPrintEvent, func(event interface{}) {
		fmt.Printf("[Event] DebugPrint: %+v\n", event)
	})

	if err := packet.Initialize(); err != nil {
		fmt.Printf("Initialize failed: %v\n", err)
	} else {
		fmt.Println("Driver initialized")
	}

	fmt.Println("Press Ctrl+C to exit")

	for packet.IsConnected() {
		packet.WaitForEvent(1e9)
	}

	p.Stop()
	p.Uninstall()
	fmt.Println("=== Test Complete ===")
}
