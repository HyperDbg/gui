package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/ddkwork/HyperDbg/debugger"
	"github.com/ddkwork/HyperDbg/debugger/driver"
)

func main() {
	driverPath := flag.String("driver", `d:\ux\examples\hypedbg\rust-driver\kd\hyperdbg_kd.sys`, "Driver path")
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

	packet := debugger.NewPacket("http://127.0.0.1:50080")
	if err := packet.Connect(); err != nil {
		fmt.Printf("Connect failed: %v\n", err)
		p.Stop()
		p.Uninstall()
		os.Exit(1)
	}
	defer packet.Disconnect()

	fmt.Println("Debugger initialized successfully")

	packet.RegisterCallback(debugger.MsgTypeBreakpointEvent, func(event any) {
		fmt.Printf("[Event] Breakpoint: %+v\n", event)
	})

	packet.RegisterCallback(debugger.MsgTypeExceptionEvent, func(event any) {
		fmt.Printf("[Event] Exception: %+v\n", event)
	})

	packet.RegisterCallback(debugger.MsgTypeDebugPrintEvent, func(event any) {
		fmt.Printf("[Event] DebugPrint: %+v\n", event)
	})

	if err := packet.LoadVmm(); err != nil {
		fmt.Printf("LoadVmm failed: %v\n", err)
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
