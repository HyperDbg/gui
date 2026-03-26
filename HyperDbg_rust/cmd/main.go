package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ddkwork/HyperDbg_rust/debugger"
)

func main() {
	port := flag.Int("port", protocol.DefaultPort, "Server port")
	processID := flag.Uint("pid", 0, "Process ID to attach")
	flag.Parse()

	dbg := debugger.NewDebugger()

	fmt.Printf("Starting HyperDbg communicator on port %d...\n", *port)
	if err := dbg.Start(*port); err != nil {
		fmt.Printf("Failed to start: %v\n", err)
		os.Exit(1)
	}
	defer dbg.Stop()

	fmt.Println("Waiting for driver connection...")
	if err := dbg.WaitForDriver(30 * time.Second); err != nil {
		fmt.Printf("Failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Driver connected, initializing...")
	if err := dbg.Initialize(); err != nil {
		fmt.Printf("Failed to initialize: %v\n", err)
		os.Exit(1)
	}
	defer dbg.Terminate()

	if *processID > 0 {
		fmt.Printf("Attaching to process %d...\n", *processID)
		if err := dbg.AttachProcess(uint32(*processID)); err != nil {
			fmt.Printf("Failed to attach: %v\n", err)
			os.Exit(1)
		}
		defer dbg.DetachProcess()
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("Debugger ready. Press Ctrl+C to exit.")

	for {
		select {
		case <-sigChan:
			fmt.Println("\nShutting down...")
			return
		default:
			event := dbg.WaitForEvent(100 * time.Millisecond)
			if event != nil {
				handleEvent(dbg, event)
			}
		}
	}
}

func handleEvent(dbg *debugger.Debugger, event *protocol.Message) {
	switch event.Header.Type {
	case protocol.MsgTypeBreakpointEvent:
		fmt.Println("\n=== Breakpoint Hit ===")
		regs, err := dbg.ReadRegisters()
		if err == nil {
			fmt.Printf("RIP: 0x%016x\n", regs.RIP)
			fmt.Printf("RSP: 0x%016x\n", regs.RSP)
			fmt.Printf("RAX: 0x%016x\n", regs.RAX)
		}
		fmt.Println("Press Enter to continue...")
		fmt.Scanln()
		dbg.Continue()

	case protocol.MsgTypeExceptionEvent:
		fmt.Println("\n=== Exception ===")
		dbg.Continue()

	case protocol.MsgTypeDebugPrintEvent:
		// Debug prints are handled by the debugger callbacks
	}
}
