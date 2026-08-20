// unpacker_runner is a standalone command-line tool that runs the Themida
// unpacker against a target executable.
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/hyperdbg/go-libhyperdbg/debugger/themida"
)

func main() {
	var (
		driverPath     = flag.String("driver", "", "path to hyperkd.sys (required)")
		exePath        = flag.String("exe", "", "path to packed executable (required)")
		logPath        = flag.String("log", "API_Logger.txt", "path to API Logger output")
		runSeconds     = flag.Int("run-seconds", 30, "seconds to let target run")
		kernel32Path   = flag.String("kernel32", "", "path to kernel32.dll (for PE export resolution)")
		ntdllPath      = flag.String("ntdll", "", "path to ntdll.dll (for PE export resolution)")
		kernelbasePath = flag.String("kernelbase", "", "path to kernelbase.dll (for PE export resolution)")
	)
	flag.Parse()

	if *driverPath == "" || *exePath == "" {
		fmt.Fprintln(os.Stderr, "error: -driver and -exe are required")
		flag.Usage()
		os.Exit(2)
	}

	cfg := themida.UnpackerConfig{
		DriverPath:     *driverPath,
		ExePath:        *exePath,
		LogPath:        *logPath,
		RunSeconds:     *runSeconds,
		Kernel32Path:   *kernel32Path,
		NtdllPath:      *ntdllPath,
		KernelbasePath: *kernelbasePath,
	}

	u := themida.NewUnpacker(cfg)
	result, err := u.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %v\n", err)
		os.Exit(1)
	}

	fmt.Fprintf(os.Stderr, "\n[*] Result:\n")
	fmt.Fprintf(os.Stderr, "  Log: %s\n", result.LogPath)
	fmt.Fprintf(os.Stderr, "  API calls: %d\n", len(result.APICalls))
	fmt.Fprintf(os.Stderr, "  OEP hints: %d\n", len(result.OEPHints))
	if result.SetEventEntry != nil {
		fmt.Fprintf(os.Stderr, "  SetEvent VM entry: %s\n", result.SetEventEntry.Format())
	}
	if result.IOMarkerAddress != 0 {
		fmt.Fprintf(os.Stderr, "  I/O Marker: 0x%X\n", result.IOMarkerAddress)
	}
	if result.WLSection.BaseAddress != 0 {
		fmt.Fprintf(os.Stderr, "  WL Section: 0x%X size=0x%X\n", result.WLSection.BaseAddress, result.WLSection.RegionSize)
	}
	if result.VMType != 0 {
		vmName := "Unknown"
		switch result.VMType {
		case 0:
			vmName = "Old CISC"
		case 1:
			vmName = "New CISC"
		case 3:
			vmName = "RISC"
		}
		fmt.Fprintf(os.Stderr, "  VM Type: %s\n", vmName)
	}
}
