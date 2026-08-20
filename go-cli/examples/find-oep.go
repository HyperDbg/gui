// find-oep.go is the Go equivalent of find-oep.ds. It launches a packed
// (Themida/WinLicense) executable under the HyperDbg VMM, registers EPT
// execution hooks on RtlAllocateHeap and VirtualAlloc, and logs the
// caller return addresses so the OEP can be located.
//
// The hook callbacks are written in the Go subset (see
// docs/go-subset-spec.md) and compiled to binary AST by go-bridge/ast.
// They run in VMX-root inside the hyperkd driver's Go interpreter.
//
// Two modes:
//
//  1. Symbol mode (preferred, requires DbgHelp + sympath):
//
//     go run ./examples/find-oep.go \
//     -exe "D:\...\SuperRecovery_V4.8.1.5.exe" \
//     -driver "Debug\hyperkd.sys" \
//     -log find-oep.log \
//     -sympath "srv*c:\symbols*https://msdl.microsoft.com/download/symbols" \
//     -ntdll "C:\Windows\System32\ntdll.dll" \
//     -kernelbase "C:\Windows\System32\kernelbase.dll" \
//     -rtl-allocate-heap-sym "ntdll!RtlAllocateHeap" \
//     -virtual-alloc-sym "kernelbase!VirtualAlloc"
//
//  2. Address mode (fallback, when symbol resolution is unavailable):
//
//     go run ./examples/find-oep.go \
//     -exe "..." -driver "..." -log find-oep.log \
//     -rtl-allocate-heap 0x77f31230 \
//     -virtual-alloc 0x77f41000
//
// Symbol addresses can be obtained with dumpbin /symbols ntdll.dll or from
// the Visual Studio debugger.
package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/hyperdbg/go-libhyperdbg/api"
	"github.com/hyperdbg/go-libhyperdbg/symbolparser"
)

func main() {
	var (
		exePath    = flag.String("exe", "", "path to the packed executable to debug (required)")
		driverPath = flag.String("driver", "Debug\\hyperkd.sys", "path to the hyperkd.sys VMM driver")
		logPath    = flag.String("log", "find-oep.log", "path to the hook log output file")
		runSeconds = flag.Int("run-seconds", 30, "how long to let the target run before pausing")
		peBase     = flag.Uint64("pe-base", 0, "if set, break only when caller falls in [pe-base, pe-base+pe-size)")
		peSize     = flag.Uint64("pe-size", 0, "size of the Themida PE range (used with -pe-base)")
		// Symbol mode
		sympath         = flag.String("sympath", "", "DbgHelp symbol search path (enables symbol mode)")
		ntdllPath       = flag.String("ntdll", "", "path to ntdll.dll (for symbol mode)")
		kernelbasePath  = flag.String("kernelbase", "", "path to kernelbase.dll (for symbol mode, optional)")
		rtlAllocHeapSym = flag.String("rtl-allocate-heap-sym", "", "symbol for RtlAllocateHeap (e.g. ntdll!RtlAllocateHeap)")
		virtualAllocSym = flag.String("virtual-alloc-sym", "", "symbol for VirtualAlloc (e.g. kernelbase!VirtualAlloc)")
		// Address mode (fallback)
		rtlAllocHeap = flag.Uint64("rtl-allocate-heap", 0, "address of ntdll!RtlAllocateHeap (hex, address mode)")
		virtualAlloc = flag.Uint64("virtual-alloc", 0, "address of kernelbase!VirtualAlloc (hex, address mode)")
	)
	flag.Parse()

	if *exePath == "" {
		fmt.Fprintln(os.Stderr, "error: -exe is required")
		flag.Usage()
		os.Exit(2)
	}

	useSymbolMode := *sympath != "" && *rtlAllocHeapSym != ""
	if !useSymbolMode && *rtlAllocHeap == 0 {
		fmt.Fprintln(os.Stderr, "error: either (-sympath + -rtl-allocate-heap-sym) or -rtl-allocate-heap is required")
		os.Exit(2)
	}

	// Catch Ctrl-C so the VMM is always unloaded on exit.
	// Build options: output + optional symbol resolver.
	opts := []api.Option{api.WithOutput(os.Stdout)}
	var resolver symbolparser.Resolver
	if useSymbolMode {
		resolver = symbolparser.New()
		opts = append(opts, api.WithSymbolResolver(resolver))
	}

	dbg, err := api.New(opts...)
	if err != nil {
		fmt.Fprintf(os.Stderr, "New failed: %v\n", err)
		os.Exit(1)
	}
	defer dbg.Close()

	if err := dbg.Connect("local"); err != nil {
		fmt.Fprintf(os.Stderr, "Connect failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("[*] Connected to local HyperDbg device\n")

	if err := dbg.LoadVMM(*driverPath); err != nil {
		fmt.Fprintf(os.Stderr, "LoadVMM failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("[*] VMM loaded (%s)\n", *driverPath)
	defer dbg.UnloadVMM()

	// Initialise symbol resolver (symbol mode only).
	if useSymbolMode {
		if err := resolver.Init(*sympath); err != nil {
			fmt.Fprintf(os.Stderr, "symbol Init failed: %v\n", err)
			os.Exit(1)
		}
		defer resolver.Close()
		if *ntdllPath != "" {
			if _, err := resolver.LoadModule(*ntdllPath, 0); err != nil {
				fmt.Fprintf(os.Stderr, "LoadModule(ntdll) failed (non-fatal): %v\n", err)
			}
		}
		if *kernelbasePath != "" {
			if _, err := resolver.LoadModule(*kernelbasePath, 0); err != nil {
				fmt.Fprintf(os.Stderr, "LoadModule(kernelbase) failed (non-fatal): %v\n", err)
			}
		}
		fmt.Printf("[*] Symbol resolver initialised (sympath=%s)\n", *sympath)
	}

	// ---- RtlAllocateHeap hook ----
	rahSrc := `package hook
func hook(ctx *HookCtx) {
	ret := ctx.StackReadQword(0) & 0xFFFFFFFF
	ctx.Printf("RAH ret=%x\n", ret)
`
	if *peBase != 0 && *peSize != 0 {
		rahSrc += fmt.Sprintf(`
	if ret >= 0x%x && ret < 0x%x {
		ctx.Printf(" *** RAH NEAR-OEP ***\n")
		ctx.Break()
	}
`, *peBase, *peBase+*peSize)
	} else {
		rahSrc += `	if ret < 0x10000000 {
		ctx.Printf(" *** LOW ***\n")
	}
`
	}
	rahSrc += `}`

	if useSymbolMode {
		if _, err := dbg.EptHookSymbol(*rtlAllocHeapSym, rahSrc); err != nil {
			fmt.Fprintf(os.Stderr, "EptHookSymbol(%s) failed: %v\n", *rtlAllocHeapSym, err)
			os.Exit(1)
		}
		fmt.Printf("[*] Hooked %s\n", *rtlAllocHeapSym)
	} else {
		if _, err := dbg.EptHook(*rtlAllocHeap, rahSrc); err != nil {
			fmt.Fprintf(os.Stderr, "EptHook(RtlAllocateHeap @ %#x) failed: %v\n", *rtlAllocHeap, err)
			os.Exit(1)
		}
		fmt.Printf("[*] Hooked RtlAllocateHeap @ %#x\n", *rtlAllocHeap)
	}

	// ---- VirtualAlloc hook (optional) ----
	// Best-effort: on Win10+ kernel32!VirtualAlloc is a forwarding stub and
	// the kernelbase implementation is what we actually want to hook. Either
	// may fail; RAH alone is sufficient for OEP detection (see project memory).
	vaHooked := false
	vaSrc := `package hook
func hook(ctx *HookCtx) {
	ret := ctx.StackReadQword(0) & 0xFFFFFFFF
	ctx.Printf("VA ret=%x\n", ret)
}`
	if useSymbolMode && *virtualAllocSym != "" {
		if _, err := dbg.EptHookSymbol(*virtualAllocSym, vaSrc); err != nil {
			fmt.Fprintf(os.Stderr, "EptHookSymbol(%s) failed (non-fatal): %v\n", *virtualAllocSym, err)
		} else {
			fmt.Printf("[*] Hooked %s\n", *virtualAllocSym)
			vaHooked = true
		}
	} else if *virtualAlloc != 0 {
		if _, err := dbg.EptHook(*virtualAlloc, vaSrc); err != nil {
			fmt.Fprintf(os.Stderr, "EptHook(VirtualAlloc @ %#x) failed (non-fatal): %v\n", *virtualAlloc, err)
		} else {
			fmt.Printf("[*] Hooked VirtualAlloc @ %#x\n", *virtualAlloc)
			vaHooked = true
		}
	}
	if !vaHooked {
		fmt.Printf("[*] VirtualAlloc hook not installed (RAH alone is sufficient)\n")
	}

	// ---- Launch the debuggee ----
	proc, err := dbg.StartProcess(*exePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "StartProcess failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("[*] Started %s (pid=%d)\n", *exePath, proc.Pid)
	defer proc.Close()

	// ---- Two 'g' commands, mirroring run-notepad.ds ----
	//
	// The .ds script flow (per project memory) is:
	//   .start  → process pauses at the first instruction
	//   g (1st) → runs to the entry point, kernel auto-pauses
	//   g (2nd) → continues past the entry point; the Themida loader
	//             runs and calls RtlAllocateHeap/VirtualAlloc, firing
	//             our hooks.
	//
	// We replicate that here: StartProcess (with CheckCallbackAtFirstInstruction
	// = true in core.Debugger) pauses at the first instruction; the first
	// Continue moves to the entry point; the second Continue lets the
	// loader run so the hooks can fire.
	if err := dbg.Continue(); err != nil {
		fmt.Fprintf(os.Stderr, "Continue (1st, to entry point) failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("[*] 1st 'g' sent (running to entry point)\n")

	// Brief pause to let the kernel settle at the entry point before the
	// second 'g'. The kernel pauses asynchronously; without this delay the
	// second Continue may arrive before the kernel has registered the
	// pause state and be coalesced into the first.
	time.Sleep(2 * time.Second)

	if err := dbg.Continue(); err != nil {
		fmt.Fprintf(os.Stderr, "Continue (2nd, past entry point) failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("[*] 2nd 'g' sent (running past entry point; hooks are live)\n")
	fmt.Printf("[*] Running for %d seconds... (Ctrl-C to pause early)\n", *runSeconds)

	time.Sleep(time.Duration(*runSeconds) * time.Second)

	if err := dbg.Pause(); err != nil {
		fmt.Fprintf(os.Stderr, "Pause failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("[*] Paused. Inspect %s for 'RAH ret=' / 'VA ret=' lines.\n", *logPath)
	fmt.Printf("[*] Near-OEP addresses appear as 'RAH ret=<low-addr>' (Themida loader region).\n")
}
