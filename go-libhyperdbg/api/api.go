// Package api — api.go
//
// API defines the complete HyperDbg debugger interface in a single file.
// Every method corresponds to a HyperDbg command or lifecycle operation;
// the concrete implementation is *Debugger (see debugger.go and the
// commands_*.go files). This interface exists so the framework's full
// capability can be seen at a glance, without reading every source file.
//
// Methods are grouped by function. Private helpers (e.g. dispatchPacket)
// are intentionally omitted. The compile-time assertion at the bottom of
// this file guarantees *Debugger stays in sync with the interface.
package api

import (
	"time"

	"github.com/hyperdbg/go-libhyperdbg/debugger/commands"
	"github.com/hyperdbg/go-libhyperdbg/debugger/core"
	"github.com/hyperdbg/go-libhyperdbg/symbolparser"
)

// API is the complete HyperDbg debugger interface. *Debugger (see
// debugger.go) is the concrete implementation; the compile-time check at
// the bottom of this file guarantees it implements every method.
type API interface {
	// === Lifecycle ===
	// Device, VMM, process and session lifecycle. Most of these mirror the
	// meta command set (.connect / load / unload / .start / attach / ...).

	// Connect opens the HyperDbg device by name ("local" for local debug).
	Connect(target string) error
	// LoadVMM installs+starts the VMM driver and sends IOCTL_INIT_VMM.
	LoadVMM(driverPath string) error
	// UnloadVMM stops the VMM and removes the driver service.
	UnloadVMM() error
	// Close releases all resources held by the debugger.
	Close() error
	// StartProcess launches a process for debugging; the VMM intercepts it
	// via the debug-port callback. Caller owns the returned Process handle.
	StartProcess(exePath string) (core.Process, error)
	// Attach attaches to an already-running process by pid.
	Attach(pid uint32) error
	// Debug starts a new process in debug mode (waits for symbols).
	Debug(exePath string) error
	// Detach detaches from the current target without terminating it.
	Detach() error
	// Disconnect disconnects from the local/remote debug target.
	Disconnect() error
	// Restart restarts the current debug target.
	Restart() error
	// Kill terminates the target process by pid.
	Kill(pid uint32) error
	// Listen listens for an incoming remote-debug connection.
	Listen(ip string, port int) error
	// StartMessagePump spawns the goroutine that drains kernel packets via
	// the IRP-based channel. Must be called after LoadVMM; the returned pump
	// must be Stop-ed before UnloadVMM.
	StartMessagePump() (*core.MessagePump, error)
	// SetOnPaused registers a callback fired from the pump goroutine whenever
	// the debuggee pauses (bp hit, single-step, manual Pause, OEP, ...).
	SetOnPaused(fn func())
	// Exit signals the CLI loop to exit (returns metacmds.ErrExit).
	Exit() error

	// === Execution Control ===
	// Resume / step / run-to-address primitives.

	// Continue resumes the debugged process (g / go).
	Continue() error
	// Pause halts the debugged process.
	Pause() error
	// StepOver single-steps over calls (p).
	StepOver() error
	// StepOut runs until the current function returns (execute till return).
	StepOut() error
	// TraceInto single-steps into calls (t).
	TraceInto() error
	// Gg runs until the given address (go until address).
	Gg(addr uint64) error
	// Gu runs until the current function returns (go until return).
	Gu() error

	// === Registers & Memory ===
	// Read/write registers and target memory, search and evaluate.

	// ReadMemory reads size bytes from pid's address space at addr.
	ReadMemory(addr uint64, pid uint32, size uint32) ([]byte, error)
	// DumpMem dumps size bytes from the currently attached target.
	DumpMem(addr uint64, size uint32) ([]byte, error)
	// EditMem writes data to the target address (e).
	EditMem(addr uint64, data []byte) error
	// Register reads a single register and returns its numeric value (r <reg>).
	Register(reg string) (uint64, error)
	// AllRegisters reads all registers and returns formatted text (r).
	AllRegisters() (string, error)
	// SetRegister writes a value to a register (r <reg> <val>).
	SetRegister(reg string, val uint64) error
	// Search scans [addr, addr+size) for pattern, returning match addresses (s).
	Search(addr uint64, size uint32, pattern []byte) ([]uint64, error)
	// Eval evaluates an expression and returns the 64-bit result (eval).
	Eval(expr string) (uint64, error)

	// === Disassembly ===
	// Assemble and Unassemble (a / u).

	// Assemble assembles one instruction at addr (a <addr> <instr>).
	Assemble(addr uint64, instr string) error
	// Unassemble disassembles count instructions at addr (u <addr> <count>).
	Unassemble(addr uint64, count uint32) (string, error)

	// === Breakpoints ===
	// Software breakpoint management (bp / bc / bd / be / bl).

	// BpSet sets a software breakpoint at addr, returns its tag (bp).
	BpSet(addr uint64) (uint64, error)
	// BpClear removes the breakpoint with the given tag (bc).
	BpClear(tag uint64) error
	// BpDisable disables a breakpoint without removing it (bd).
	BpDisable(tag uint64) error
	// BpEnable re-enables a disabled breakpoint (be).
	BpEnable(tag uint64) error
	// BpList lists all breakpoints (bl).
	BpList() ([]Breakpoint, error)

	// === Stack & Modules ===
	// Call stack, module list and symbol examination (k / lm / x).

	// K walks the call stack, returning up to count frames (k <count>).
	K(count uint32) ([]CallFrame, error)
	// Lm lists the modules loaded by the current target (lm).
	Lm() ([]Module, error)
	// Examine lists symbols matching a wildcard pattern (x <pattern>).
	Examine(pattern string) ([]Module, error)

	// === CPU & MSR ===
	// CPU info, core switching, MSR and TSC access (cpu / core / rdmsr / wrmsr / tsc).

	// Cpu returns CPU info (vendor / brand / cores / frequency).
	Cpu() (CpuInfo, error)
	// Core switches the active logical core (core <id>).
	Core(coreId uint32) error
	// Rdmsr reads a Model-Specific Register (rdmsr <msr>).
	Rdmsr(msr uint32) (uint64, error)
	// Wrmsr writes a Model-Specific Register (wrmsr <msr> <val>).
	Wrmsr(msr uint32, val uint64) error
	// Tsc reads the Time-Stamp Counter (tsc).
	Tsc() (uint64, error)

	// === Address Translation ===
	// Virtual/physical address conversion and page-table walk (pa2va / va2pa / pte).

	// Pa2Va translates a physical address to virtual (pa2va <pa> <pid>).
	Pa2Va(pa uint64, pid uint32) (uint64, error)
	// Va2Pa translates a virtual address to physical (va2pa <va> <pid>).
	Va2Pa(va uint64, pid uint32) (uint64, error)
	// Pte returns the last-level PTE for va (pte <va>).
	Pte(va uint64) (uint64, error)

	// === EPT Hooks ===
	// Execution hooks via EPT (epthook family).

	// EptHook registers an EPT execution hook at hookAddress with a Go callback.
	EptHook(hookAddress uint64, callbackSrc string) (uint64, error)
	// EptHookForProcess registers an EPT hook scoped to a specific pid
	// (needed for WOW64 targets whose DLL addresses differ).
	EptHookForProcess(hookAddress uint64, pid uint32, callbackSrc string) (uint64, error)
	// EptHookSymbol registers an EPT hook at the address of a "mod!sym"
	// string; requires a SymbolResolver injected via WithSymbolResolver.
	EptHookSymbol(symbol string, callbackSrc string) (uint64, error)
	// EptHook2 registers a detour-style EPT hook (no #VMEXIT, epthook2).
	EptHook2(hookAddress uint64, callbackSrc string) (uint64, error)

	// === Monitor Hooks ===
	// Memory access monitoring (monitor r/w/e).

	// MonitorReadForProcess monitors read access to [addrStart, addrEnd) for pid.
	MonitorReadForProcess(addrStart, addrEnd uint64, pid uint32, callbackSrc string) (uint64, error)
	// MonitorWrite monitors write access to [addrStart, addrEnd) for pid.
	MonitorWrite(addrStart, addrEnd uint64, pid uint32, callbackSrc string) (uint64, error)
	// MonitorExec monitors execution at addr for pid.
	MonitorExec(addr uint64, pid uint32, callbackSrc string) (uint64, error)

	// === Syscall Hooks ===
	// SYSCALL/SYSRET interception.

	// SyscallHook hooks every SYSCALL instruction entry (!syscall).
	SyscallHook(callbackSrc string) (uint64, error)
	// SysretHook hooks every SYSRET instruction (!sysret); applies to all
	// processes — prefer SysretHookForProcess for single-process targets.
	SysretHook(callbackSrc string) (uint64, error)
	// SysretHookForProcess scopes a SYSRET hook to a single pid (recommended).
	SysretHookForProcess(pid uint32, callbackSrc string) (uint64, error)

	// === Other Hooks ===
	// Misc VM-exit hook families (!cpuid / !crwrite / !dr / !exception / ...).

	// CpuidHook hooks every CPUID instruction (!cpuid).
	CpuidHook(callbackSrc string) (uint64, error)
	// CrwriteHook hooks writes to control register cr (!crwrite <cr>).
	CrwriteHook(cr uint32, callbackSrc string) (uint64, error)
	// DrHook hooks debug-register access (!dr).
	DrHook(callbackSrc string) (uint64, error)
	// ExceptionHook hooks an exception vector (!exception <vector>).
	ExceptionHook(vector uint32, callbackSrc string) (uint64, error)
	// InterruptHook hooks a hardware interrupt vector (!interrupt <vector>).
	InterruptHook(vector uint32, callbackSrc string) (uint64, error)
	// IoInHook hooks IN instructions on an I/O port (!ioin <port>).
	IoInHook(port uint16, callbackSrc string) (uint64, error)
	// IoOutHook hooks OUT instructions on an I/O port (!ioout <port>).
	IoOutHook(port uint16, callbackSrc string) (uint64, error)
	// ModeHook hooks user/kernel mode switches (!mode).
	ModeHook(callbackSrc string) (uint64, error)
	// MsrReadHook hooks RDMSR for an MSR (!msrread <msr>; 0 = all).
	MsrReadHook(msr uint32, callbackSrc string) (uint64, error)
	// MsrWriteHook hooks WRMSR for an MSR (!msrwrite <msr>).
	MsrWriteHook(msr uint32, callbackSrc string) (uint64, error)
	// VmcallHook hooks the VMCALL instruction (!vmcall).
	VmcallHook(callbackSrc string) (uint64, error)
	// XsetbvHook hooks the XSETBV instruction (!xsetbv).
	XsetbvHook(callbackSrc string) (uint64, error)

	// === Event Management ===
	// Enable / disable / clear registered hook events (events c/d/e).

	// Events lists the tags of currently registered events.
	Events() ([]uint64, error)
	// ClearEvent removes the event with the given tag (events c <tag>).
	ClearEvent(tag uint64) error
	// DisableEvent temporarily disables an event, keeping its config (events d).
	DisableEvent(tag uint64) error
	// EnableEvent re-enables a previously disabled event (events e).
	EnableEvent(tag uint64) error

	// === Hardware Debug ===
	// FPGA-based hwdbg subsystem (hw / hw_clk).

	// Hw starts/queries the hardware debug device (hw).
	Hw() error
	// HwClk configures the hardware debug device clock (hw_clk).
	HwClk() error

	// === Debugger Info ===
	// Status, settings, symbol resolution and the command registry.

	// Status returns the current debugger state (status).
	Status() (core.DebuggerState, error)
	// Settings returns the current debugger settings (settings).
	Settings() (Settings, error)
	// SymbolResolver returns the injected symbol resolver (nil if none).
	SymbolResolver() symbolparser.Resolver
	// Commands returns the underlying command registry (for help UIs / dispatch).
	Commands() *commands.Registry
	// Help lists commands or prints detailed help for cmdName (help [name]).
	Help(cmdName string) error
	// Sym resolves a symbol name to an address (sym <name>); needs a resolver.
	Sym(name string) (uint64, error)
	// SymPath sets the symbol search path (sympath <path>); needs a resolver.
	SymPath(path string) error

	// === Prealloc ===
	// Pre-allocation / pre-activation of kernel buffers and breakpoints.

	// Preactivate pre-activates breakpoints so the next Continue fires immediately.
	Preactivate() error
	// Prealloc pre-allocates kernel event buffers of the given count (prealloc).
	Prealloc(size uint64) error

	// === IO ===
	// Direct I/O port access (i).

	// IoIn reads one byte from an I/O port (i <port>).
	IoIn(port uint16) (byte, error)

	// === Type Display ===
	// Struct-aware memory display (dt).

	// Dt displays memory at addr formatted as the named struct type (dt <type> <addr>).
	Dt(typeName string, addr uint64) (string, error)

	// === Misc ===
	// Output, scripting, platform introspection and assorted commands.

	// Sleep waits for dur; convenience for scripts that let the target run first.
	Sleep(dur time.Duration)
	// Printf writes formatted output to the debugger's Output sink.
	Printf(format string, args ...any) error
	// Exec parses and runs a command line (e.g. "g", "load vmm", ".logopen x").
	Exec(cmdLine string) error
	// Output evaluates expr and logs the result without printing to console (output).
	Output(expr string) error
	// Print prints an expression's value with formatting (print <expr>).
	Print(expr string) (string, error)
	// Flush flushes kernel logging buffers to user space (flush).
	Flush() error
	// Test runs the internal self-test suite (test).
	Test() error
	// ClearScreen writes the ANSI clear-screen sequence to output (cls / clear).
	ClearScreen() error
	// Apic prints local APIC status (!apic).
	Apic() error
	// Hide hides the VMM from hypervisor detection (!hide).
	Hide() error
	// Unhide reverses !hide, making the hypervisor visible again (!unhide).
	Unhide() error
	// Idt returns the IDT entry address for the given vector (!idt <vector>).
	Idt(vector uint32) (uint64, error)
	// Ioapic prints I/O APIC status (!ioapic).
	Ioapic() error
	// Lbr enables Last Branch Record capture (!lbr).
	Lbr() error
	// LbrDump dumps the captured LBR branches (!lbrdump).
	LbrDump() error
	// Measure measures VM-exit overhead for anti-debug timing assessment (!measure).
	Measure() error
	// PciCam reads PCI config via Configuration Access Method (!pcicam).
	PciCam(bus, device, function, offset uint32) (uint32, error)
	// PciTree prints the PCI device tree (!pcitree).
	PciTree() error
	// Pmc reads a Performance Monitoring Counter (!pmc <counter>).
	Pmc(counter uint32) (uint64, error)
	// Pt configures Intel Processor Trace (!pt).
	Pt() error
	// Rev triggers the reversing-machine memory reconstruction (!rev).
	Rev() (uint32, error)
	// Smi triggers a System Management Interrupt (!smi).
	Smi() error
	// Trace enables Intel PT tracing (!trace).
	Trace() error
	// Track tracks memory access patterns (!track).
	Track() error
	// Dump writes a minidump of the target process to path (dump <path>).
	Dump(path string) error
	// Formats displays an expression in hex/dec/oct/bin (formats <expr>).
	Formats(expr string) error
	// PageIn forces the page containing addr into physical memory (pagein <addr>).
	PageIn(addr uint64) error
	// Pe parses and prints a PE file's header/sections/imports/exports (pe <path>).
	Pe(path string) error
	// Process lists all system processes (process).
	Process() error
	// Script executes a .ds script file (script <path>).
	Script(path string) error
	// Switch switches the current debug target to pid (switch <pid>).
	Switch(pid uint32) error
	// Thread lists the threads of the current target (thread).
	Thread() error
}

// Compile-time guarantee that *Debugger implements every method of API.
// If a method is added/renamed/removed on either side, the build breaks
// here — which is the point: this file is the single source of truth for
// the framework's public surface.
var _ API = (*Debugger)(nil)
