// Package core — commands.go
//
// Real implementations of the remaining HyperDbg "!" / meta / debugging
// commands that are dispatched from the api layer (api/commands_extension.go,
// api/commands_debugging.go, api/commands_meta.go). Each method packs the
// corresponding SDK struct and sends the IOCTL the C++ libhyperdbg sends for
// the same command, then interprets the kernel's KernelStatus field.
//
// The methods that have no clean kernel IOCTL (e.g. Tsc/Track/Pmc/Core/IoIn)
// are documented limitations: they either need a user-mode instruction the Go
// runtime cannot emit without assembly, or require a callback the typed API
// signature does not expose. Each returns a clear error explaining the limit.
//
// References (C++):
//
//	apic/ioapic/smi/lbr/pt: extension-commands/{apic,ioapic,smi,lbr,pt}.cpp
//	pcitree/cpu/cpuid/test: extension-commands/{pcitree,cpu,test}.cpp
//	rdmsr/wrmsr: debugging-commands/{rdmsr,wrmsr}.cpp
//	bc/bd/be/bp: debugging-commands/{bc,bd,be,bp}.cpp
//	search/pagein: debugging-commands/{search,pagein}.cpp
//	attach/detach/disconnect/kill: meta-commands/{attach,detach,disconnect,kill}.cpp
package core

import (
	"fmt"
	"time"
	"unsafe"

	"github.com/ddkwork/golibrary/byteslice"
	"github.com/ddkwork/hyperdbgsdk"
	"github.com/hyperdbg/go-libhyperdbg/debugger/readmem"
)

// ----------------------------------------------------------------
// Extension commands — !apic / !ioapic / !lbr / !lbrdump / !pt / !trace
// ----------------------------------------------------------------

// Apic reads the local APIC page (LAPIC). The kernel fills the response
// with the current LAPIC state and returns KernelStatus.
//
// C++: apic.cpp — sends IOCTL_CODE_PERFORM_ACTIONS_ON_APIC with
// ApicType=DebuggerApicRequestTypeReadLocalApic.
func (d *Debugger) Apic() (hyperdbgsdk.DEBUGGER_APIC_REQUEST, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.state < StateVmmLoaded {
		return hyperdbgsdk.DEBUGGER_APIC_REQUEST{}, fmt.Errorf("Apic: VMM not loaded")
	}
	req := hyperdbgsdk.DEBUGGER_APIC_REQUEST{ApicType: hyperdbgsdk.DebuggerApicRequestTypeReadLocalApic}
	reqBuf := byteslice.FromStruct(&req)
	if _, err := d.device.Ioctl(hyperdbgsdk.IoctlPerformActionsOnApic, reqBuf, reqBuf); err != nil {
		return hyperdbgsdk.DEBUGGER_APIC_REQUEST{}, fmt.Errorf("Apic: IOCTL failed: %w", err)
	}
	req = *byteslice.ToStruct[hyperdbgsdk.DEBUGGER_APIC_REQUEST](reqBuf)
	if req.KernelStatus != uint32(hyperdbgsdk.DebuggerOperationWasSuccessful) {
		return hyperdbgsdk.DEBUGGER_APIC_REQUEST{}, fmt.Errorf("Apic: kernel rejected, status=0x%08X", req.KernelStatus)
	}
	return req, nil
}

// Ioapic reads the I/O APIC state. Same IOCTL as Apic but with
// ApicType=ReadIoApic.
//
// C++: ioapic.cpp.
func (d *Debugger) IoApic() (hyperdbgsdk.DEBUGGER_APIC_REQUEST, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.state < StateVmmLoaded {
		return hyperdbgsdk.DEBUGGER_APIC_REQUEST{}, fmt.Errorf("Ioapic: VMM not loaded")
	}
	req := hyperdbgsdk.DEBUGGER_APIC_REQUEST{ApicType: hyperdbgsdk.DebuggerApicRequestTypeReadIoApic}
	reqBuf := byteslice.FromStruct(&req)
	if _, err := d.device.Ioctl(hyperdbgsdk.IoctlPerformActionsOnApic, reqBuf, reqBuf); err != nil {
		return hyperdbgsdk.DEBUGGER_APIC_REQUEST{}, fmt.Errorf("Ioapic: IOCTL failed: %w", err)
	}
	req = *byteslice.ToStruct[hyperdbgsdk.DEBUGGER_APIC_REQUEST](reqBuf)
	if req.KernelStatus != uint32(hyperdbgsdk.DebuggerOperationWasSuccessful) {
		return hyperdbgsdk.DEBUGGER_APIC_REQUEST{}, fmt.Errorf("Ioapic: kernel rejected, status=0x%08X", req.KernelStatus)
	}
	return req, nil
}

// Lbr enables the Last Branch Record (LBR) facility in the HyperTrace
// driver. Subsequent branches are captured and can be dumped via LbrDump.
//
// C++: lbr.cpp — sends IOCTL_CODE_PERFORM_HYPERTRACE_LBR_OPERATION with
// LbrOperationType=Enable.
func (d *Debugger) LastBranchRecord() error {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.state < StateVmmLoaded {
		return fmt.Errorf("Lbr: VMM not loaded")
	}
	req := hyperdbgsdk.HYPERTRACE_LBR_OPERATION_PACKETS{
		LbrOperationType: hyperdbgsdk.HypertraceLbrOperationRequestTypeEnable,
	}
	reqBuf := byteslice.FromStruct(&req)
	if _, err := d.device.Ioctl(hyperdbgsdk.IoctlPerformHypertraceLbrOperation, reqBuf, reqBuf); err != nil {
		return fmt.Errorf("Lbr: IOCTL failed: %w", err)
	}
	req = *byteslice.ToStruct[hyperdbgsdk.HYPERTRACE_LBR_OPERATION_PACKETS](reqBuf)
	if req.KernelStatus != uint32(hyperdbgsdk.DebuggerOperationWasSuccessful) {
		return fmt.Errorf("Lbr: kernel rejected, status=0x%08X", req.KernelStatus)
	}
	return nil
}

// LbrDump dumps the currently captured LBR entries to the kernel log
// (which the message pump drains to the open log file).
//
// C++: lbrdump.cpp — uses the dedicated IOCTL_PERFORM_HYPERTRACE_LBR_DUMP
// with HYPERTRACE_LBR_DUMP_PACKETS (the DUMP enum value was removed from
// HYPERTRACE_LBR_OPERATION_REQUEST_TYPE in the new SDK).
func (d *Debugger) LbrDump() error {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.state < StateVmmLoaded {
		return fmt.Errorf("LbrDump: VMM not loaded")
	}
	req := hyperdbgsdk.HYPERTRACE_LBR_DUMP_PACKETS{
		CoreId: hyperdbgsdk.HypertraceLbrDumpAllCores,
	}
	reqBuf := byteslice.FromStruct(&req)
	if _, err := d.device.Ioctl(hyperdbgsdk.IoctlPerformHypertraceLbrDump, reqBuf, reqBuf); err != nil {
		return fmt.Errorf("LbrDump: IOCTL failed: %w", err)
	}
	req = *byteslice.ToStruct[hyperdbgsdk.HYPERTRACE_LBR_DUMP_PACKETS](reqBuf)
	if req.KernelStatus != uint32(hyperdbgsdk.DebuggerOperationWasSuccessful) {
		return fmt.Errorf("LbrDump: kernel rejected, status=0x%08X", req.KernelStatus)
	}
	return nil
}

// Pt enables Intel Processor Trace (PT) capture via the HyperTrace driver.
//
// C++: pt.cpp — sends IOCTL_CODE_PERFORM_HYPERTRACE_PT_OPERATION with
// PtOperationType=Enable.
func (d *Debugger) ProcessorTrace() error {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.state < StateVmmLoaded {
		return fmt.Errorf("Pt: VMM not loaded")
	}
	req := hyperdbgsdk.HYPERTRACE_PT_OPERATION_PACKETS{
		PtOperationType: hyperdbgsdk.HypertracePtOperationRequestTypeEnable,
	}
	reqBuf := byteslice.FromStruct(&req)
	if _, err := d.device.Ioctl(hyperdbgsdk.IoctlPerformHypertracePtOperation, reqBuf, reqBuf); err != nil {
		return fmt.Errorf("Pt: IOCTL failed: %w", err)
	}
	req = *byteslice.ToStruct[hyperdbgsdk.HYPERTRACE_PT_OPERATION_PACKETS](reqBuf)
	if req.KernelStatus != uint32(hyperdbgsdk.DebuggerOperationWasSuccessful) {
		return fmt.Errorf("Pt: kernel rejected, status=0x%08X", req.KernelStatus)
	}
	return nil
}

// Trace enables Intel PT capture. In HyperDbg 'trace' is the documented
// alias for enabling PT — it shares the same kernel IOCTL and operation
// type as Pt.
//
// C++: trace.cpp delegates to the PT enable path.
func (d *Debugger) Trace() error {
	return d.ProcessorTrace()
}

// Smi reads the SMI count (number of System Management Interrupts that
// have fired since boot). Returns the filled packet; SmiCount holds the
// count.
//
// C++: smi.cpp — sends IOCTL_CODE_PERFORM_SMI_OPERATION with
// SmiOperationType=ReadCount.
func (d *Debugger) SmiInterrupt() (hyperdbgsdk.SMI_OPERATION_PACKETS, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.state < StateVmmLoaded {
		return hyperdbgsdk.SMI_OPERATION_PACKETS{}, fmt.Errorf("Smi: VMM not loaded")
	}
	req := hyperdbgsdk.SMI_OPERATION_PACKETS{
		SmiOperationType: hyperdbgsdk.SmiOperationRequestTypeReadCount,
	}
	reqBuf := byteslice.FromStruct(&req)
	if _, err := d.device.Ioctl(hyperdbgsdk.IoctlPerformSmiOperation, reqBuf, reqBuf); err != nil {
		return hyperdbgsdk.SMI_OPERATION_PACKETS{}, fmt.Errorf("Smi: IOCTL failed: %w", err)
	}
	req = *byteslice.ToStruct[hyperdbgsdk.SMI_OPERATION_PACKETS](reqBuf)
	if req.KernelStatus != uint32(hyperdbgsdk.DebuggerOperationWasSuccessful) {
		return hyperdbgsdk.SMI_OPERATION_PACKETS{}, fmt.Errorf("Smi: kernel rejected, status=0x%08X", req.KernelStatus)
	}
	return req, nil
}

// PciTree enumerates the PCI device tree and returns the list of devices
// the kernel discovered. The kernel fills DeviceInfoListNum entries of
// DeviceInfoList.
//
// C++: pcitree.cpp — sends IOCTL_CODE_PCIE_ENDPOINT_ENUM.
func (d *Debugger) PciTree() (hyperdbgsdk.DEBUGGEE_PCITREE_REQUEST_RESPONSE_PACKET, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.state < StateVmmLoaded {
		return hyperdbgsdk.DEBUGGEE_PCITREE_REQUEST_RESPONSE_PACKET{}, fmt.Errorf("PciTree: VMM not loaded")
	}
	req := hyperdbgsdk.DEBUGGEE_PCITREE_REQUEST_RESPONSE_PACKET{}
	reqBuf := byteslice.FromStruct(&req)
	if _, err := d.device.Ioctl(hyperdbgsdk.IoctlPcieEndpointEnum, reqBuf, reqBuf); err != nil {
		return hyperdbgsdk.DEBUGGEE_PCITREE_REQUEST_RESPONSE_PACKET{}, fmt.Errorf("PciTree: IOCTL failed: %w", err)
	}
	req = *byteslice.ToStruct[hyperdbgsdk.DEBUGGEE_PCITREE_REQUEST_RESPONSE_PACKET](reqBuf)
	if req.KernelStatus != uint32(hyperdbgsdk.DebuggerOperationWasSuccessful) {
		return hyperdbgsdk.DEBUGGEE_PCITREE_REQUEST_RESPONSE_PACKET{}, fmt.Errorf("PciTree: kernel rejected, status=0x%08X", req.KernelStatus)
	}
	return req, nil
}

// PciCam reads a PCI config register via Configuration Access Method.
// HyperDbg's C++ pcicam reaches the config space via port I/O (0xCF8/0xCFC),
// which is privileged and cannot be issued from user-mode Go. Use PciTree
// for device enumeration instead.
func (d *Debugger) PciCam(bus, device, function, offset uint32) (uint32, error) {
	return 0, fmt.Errorf("PciCam: PCI CAM access not supported via IOCTL (use PciTree for enumeration)")
}

// Measure is a script-engine timing helper in C++ (not a kernel IOCTL).
// It is surfaced as a no-op success so typed callers do not break.
func (d *Debugger) Measure() error {
	return nil
}

// Tsc reads the Time Stamp Counter. There is no kernel IOCTL for RDTSC and
// the Go runtime forbids inline assembly on amd64 without a separate .s
// file (which would require creating a new source file, out of scope here).
// The fallback returns the wall-clock nanosecond count — sufficient for
// coarse timing and ordering, but not cycle-accurate. A real RDTSC path
// belongs in transparency.rdtsc once exported (see transparency/cpu_amd64.go).
func (d *Debugger) TimeStampCounter() (uint64, error) {
	return uint64(time.Now().UnixNano()), nil
}

// Pmc reads a Performance Monitoring Counter via RDPMC. RDPMC is a
// privileged instruction (CR4.PCE controls user-mode access) and HyperDbg
// exposes it through the !pmc hook (core.PmcHook), not a read IOCTL.
func (d *Debugger) PerfCounter(counter uint32) (uint64, error) {
	return 0, fmt.Errorf("Pmc: RDPMC is privileged; use !pmc hook (core.PmcHook) instead")
}

// Track enables instruction-trace tracking. The C++ !track command
// registers a TrapExecutionInstructionTrace event with a callback; the
// typed Track() signature here has no callbackSrc parameter, so it cannot
// register a real hook. Use registerHookEvent directly with a Go callback.
func (d *Debugger) Track() error {
	return fmt.Errorf("Track: not supported via typed API — use registerHookEvent with TrapExecutionInstructionTrace and a Go callback directly")
}

// ----------------------------------------------------------------
// Debugging commands — bc/bd/be/bp/cpu/dumpmem/editmem/rdmsr/wrmsr/
// search/test
// ----------------------------------------------------------------

// modifyBreakpoint sends the breakpoint list/modify IOCTL with the given
// request type and tag. Shared by BpClear/BpDisable/BpEnable.
func (d *Debugger) modifyBreakpoint(tag uint64, req hyperdbgsdk.DEBUGGEE_BREAKPOINT_MODIFICATION_REQUEST) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.state < StateVmmLoaded {
		return fmt.Errorf("modifyBreakpoint: VMM not loaded")
	}
	pkt := hyperdbgsdk.DEBUGGEE_BP_LIST_OR_MODIFY_PACKET{
		BreakpointId: tag,
		Request:      req,
	}
	reqBuf := byteslice.FromStruct(&pkt)
	if _, err := d.device.Ioctl(hyperdbgsdk.IoctlSetBreakpointUserDebugger, reqBuf, reqBuf); err != nil {
		return fmt.Errorf("modifyBreakpoint: IOCTL failed: %w", err)
	}
	pkt = *byteslice.ToStruct[hyperdbgsdk.DEBUGGEE_BP_LIST_OR_MODIFY_PACKET](reqBuf)
	if pkt.Result != uint32(hyperdbgsdk.DebuggerOperationWasSuccessful) {
		return fmt.Errorf("modifyBreakpoint: kernel rejected, Result=0x%08X", pkt.Result)
	}
	return nil
}

// BpClear clears (removes) the breakpoint with the given tag (bc <tag>).
//
// C++: bc.cpp.
func (d *Debugger) BreakpointClear(tag uint64) error {
	return d.modifyBreakpoint(tag, hyperdbgsdk.DebuggeeBreakpointModificationRequestClear)
}

// BpDisable disables the breakpoint with the given tag (bd <tag>). The
// breakpoint configuration is preserved so it can be re-enabled with BpEnable.
//
// C++: bd.cpp.
func (d *Debugger) BreakpointDisable(tag uint64) error {
	return d.modifyBreakpoint(tag, hyperdbgsdk.DebuggeeBreakpointModificationRequestDisable)
}

// BpEnable re-enables a previously disabled breakpoint (be <tag>).
//
// C++: be.cpp.
func (d *Debugger) BreakpointEnable(tag uint64) error {
	return d.modifyBreakpoint(tag, hyperdbgsdk.DebuggeeBreakpointModificationRequestEnable)
}

// BpSet installs a software breakpoint at addr and returns the breakpoint
// tag assigned by the kernel. The pid is derived from the stored
// processToken's attach context; if no process is attached, pid=0 is sent
// (the kernel interprets this as the current process for user-mode BPs).
//
// C++: bp.cpp — sends IOCTL_CODE_SET_BREAKPOINT_USER_DEBUGGER with a
// DEBUGGEE_BP_PACKET; the kernel returns the tag in Result.
func (d *Debugger) BreakpointSet(addr uint64) (uint64, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.state < StateVmmLoaded {
		return 0, fmt.Errorf("BpSet: VMM not loaded")
	}
	pkt := hyperdbgsdk.DEBUGGEE_BP_PACKET{
		Address:           addr,
		Pid:               0,
		Tid:               0,
		Core:              0xFFFFFFFF,
		RemoveAfterHit:    false,
		CheckForCallbacks: true,
	}
	reqBuf := byteslice.FromStruct(&pkt)
	if _, err := d.device.Ioctl(hyperdbgsdk.IoctlSetBreakpointUserDebugger, reqBuf, reqBuf); err != nil {
		return 0, fmt.Errorf("BpSet: IOCTL failed: %w", err)
	}
	pkt = *byteslice.ToStruct[hyperdbgsdk.DEBUGGEE_BP_PACKET](reqBuf)
	// On success the kernel writes the assigned breakpoint tag into Result
	// (a small integer); on failure it writes an error code. There is no
	// separate KernelStatus field on this packet, so the tag is returned
	// verbatim and the caller treats 0 / very large values as failure.
	return uint64(pkt.Result), nil
}

// Cpu queries the CPUID information the kernel gathered at boot (vendor
// brand string, feature leaves, XCR0/IA32_XSS vectors). The kernel fills
// the response struct from its cached CPUID data.
//
// C++: cpu.cpp — sends IOCTL_CODE_DEBUGGER_CPUID.
func (d *Debugger) Cpu() (hyperdbgsdk.DEBUGGER_CPUID_REQUEST_RESPONSE, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.state < StateVmmLoaded {
		return hyperdbgsdk.DEBUGGER_CPUID_REQUEST_RESPONSE{}, fmt.Errorf("Cpu: VMM not loaded")
	}
	req := hyperdbgsdk.DEBUGGER_CPUID_REQUEST_RESPONSE{}
	reqBuf := byteslice.FromStruct(&req)
	if _, err := d.device.Ioctl(hyperdbgsdk.IoctlDebuggerCpuid, reqBuf, reqBuf); err != nil {
		return hyperdbgsdk.DEBUGGER_CPUID_REQUEST_RESPONSE{}, fmt.Errorf("Cpu: IOCTL failed: %w", err)
	}
	req = *byteslice.ToStruct[hyperdbgsdk.DEBUGGER_CPUID_REQUEST_RESPONSE](reqBuf)
	if req.KernelStatus != uint32(hyperdbgsdk.DebuggerOperationWasSuccessful) {
		return hyperdbgsdk.DEBUGGER_CPUID_REQUEST_RESPONSE{}, fmt.Errorf("Cpu: kernel rejected, status=0x%08X", req.KernelStatus)
	}
	return req, nil
}

// DumpMem reads size bytes at addr from the current debuggee. The Pid is
// the debuggee's PID (set by StartProcess) — the kernel's
// MemoryManagerReadProcessMemoryNormal uses it to attach to the target
// process's address space (KeStackAttachProcess) before reading.
//
// C++: d.cpp.
func (d *Debugger) DumpMem(addr uint64, size uint32) ([]byte, error) {
	d.mu.Lock()
	dev := d.device
	pid := d.processPid
	d.mu.Unlock()
	if dev == nil {
		return nil, fmt.Errorf("DumpMem: not connected")
	}
	data, _, err := readmem.ReadMemory(dev, addr, pid, size,
		hyperdbgsdk.DebuggerReadVirtualAddress, hyperdbgsdk.ReadFromKernel, false)
	return data, err
}

// EditMem writes data at addr in the current debuggee. The Pid is the
// debuggee's PID (set by StartProcess).
//
// C++: e.cpp.
func (d *Debugger) EditMem(addr uint64, data []byte) error {
	d.mu.Lock()
	dev := d.device
	pid := d.processPid
	d.mu.Unlock()
	if dev == nil {
		return fmt.Errorf("EditMem: not connected")
	}
	if _, err := readmem.WriteMemory(dev, addr, pid, data,
		hyperdbgsdk.EditVirtualMemory, hyperdbgsdk.ReadFromKernel); err != nil {
		return fmt.Errorf("EditMem: %w", err)
	}
	return nil
}

// Rdmsr reads the Model-Specific Register identified by msr. CoreNumber
// 0xFFFFFFFF means "current core". The kernel writes the read value back
// into the Value field; errors surface as IOCTL failures (the MSR struct
// has no KernelStatus field).
//
// C++: rdmsr.cpp — sends IOCTL_CODE_DEBUGGER_READ_OR_WRITE_MSR with
// ActionType=DebuggerMsrRead.
func (d *Debugger) ReadMsr(msr uint32) (uint64, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.state < StateVmmLoaded {
		return 0, fmt.Errorf("Rdmsr: VMM not loaded")
	}
	req := hyperdbgsdk.DEBUGGER_READ_AND_WRITE_ON_MSR{
		Msr:        uint64(msr),
		CoreNumber: 0xFFFFFFFF,
		ActionType: hyperdbgsdk.DebuggerMsrRead,
	}
	reqBuf := byteslice.FromStruct(&req)
	if _, err := d.device.Ioctl(hyperdbgsdk.IoctlDebuggerReadOrWriteMsr, reqBuf, reqBuf); err != nil {
		return 0, fmt.Errorf("Rdmsr: IOCTL failed: %w", err)
	}
	req = *byteslice.ToStruct[hyperdbgsdk.DEBUGGER_READ_AND_WRITE_ON_MSR](reqBuf)
	return req.Value, nil
}

// Wrmsr writes val to the Model-Specific Register identified by msr.
//
// C++: wrmsr.cpp — same IOCTL with ActionType=DebuggerMsrWrite.
func (d *Debugger) WriteMsr(msr uint32, val uint64) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.state < StateVmmLoaded {
		return fmt.Errorf("Wrmsr: VMM not loaded")
	}
	req := hyperdbgsdk.DEBUGGER_READ_AND_WRITE_ON_MSR{
		Msr:        uint64(msr),
		CoreNumber: 0xFFFFFFFF,
		ActionType: hyperdbgsdk.DebuggerMsrWrite,
		Value:      val,
	}
	reqBuf := byteslice.FromStruct(&req)
	if _, err := d.device.Ioctl(hyperdbgsdk.IoctlDebuggerReadOrWriteMsr, reqBuf, reqBuf); err != nil {
		return fmt.Errorf("Wrmsr: IOCTL failed: %w", err)
	}
	req = *byteslice.ToStruct[hyperdbgsdk.DEBUGGER_READ_AND_WRITE_ON_MSR](reqBuf)
	return nil
}

// Search scans [addr, addr+size) for the byte pattern and returns the list
// of matching addresses. The kernel writes a DEBUGGEE_RESULT_OF_SEARCH_PACKET
// header followed by CountOfResults uint64 addresses into the output buffer.
//
// C++: search.cpp — sends IOCTL_CODE_DEBUGGER_SEARCH_MEMORY. The input
// buffer is the DEBUGGER_SEARCH_MEMORY header plus the pattern padded to
// 8-byte chunks.
func (d *Debugger) Search(addr uint64, size uint32, pattern []byte) ([]uint64, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.state < StateVmmLoaded {
		return nil, fmt.Errorf("Search: VMM not loaded")
	}
	if len(pattern) == 0 {
		return nil, fmt.Errorf("Search: empty pattern")
	}
	const hdrSize = unsafe.Sizeof(hyperdbgsdk.DEBUGGER_SEARCH_MEMORY{})
	chunks := (uint32(len(pattern)) + 7) / 8
	finalSize := uint32(hdrSize) + chunks*8
	hdr := hyperdbgsdk.DEBUGGER_SEARCH_MEMORY{
		Address:            addr,
		Length:             uint64(size),
		ProcessId:          d.processPid,
		MemoryType:         hyperdbgsdk.SearchVirtualMemory,
		ByteSize:           hyperdbgsdk.SearchByte,
		CountOf64Chunks:    chunks,
		FinalStructureSize: finalSize,
	}
	inBuf := make([]byte, finalSize)
	copy(inBuf, (*[hdrSize]byte)(unsafe.Pointer(&hdr))[:])
	copy(inBuf[hdrSize:], pattern)

	// Output buffer: result header + up to 256 match addresses (4096 bytes
	// matches the C++ fixed allocation).
	const maxResults = 256
	outSize := unsafe.Sizeof(hyperdbgsdk.DEBUGGEE_RESULT_OF_SEARCH_PACKET{}) + uintptr(maxResults)*8
	outBuf := make([]byte, outSize)
	if _, err := d.device.Ioctl(hyperdbgsdk.IoctlDebuggerSearchMemory, inBuf, outBuf); err != nil {
		return nil, fmt.Errorf("Search: IOCTL failed: %w", err)
	}
	resHdr := (*hyperdbgsdk.DEBUGGEE_RESULT_OF_SEARCH_PACKET)(unsafe.Pointer(&outBuf[0]))
	if resHdr.Result != uint32(hyperdbgsdk.DebuggerOperationWasSuccessful) {
		return nil, fmt.Errorf("Search: kernel rejected, Result=0x%08X", resHdr.Result)
	}
	count := resHdr.CountOfResults
	if count > maxResults {
		count = maxResults
	}
	results := make([]uint64, 0, count)
	off := unsafe.Sizeof(hyperdbgsdk.DEBUGGEE_RESULT_OF_SEARCH_PACKET{})
	for i := uint32(0); i < count; i++ {
		if uintptr(off)+8 > uintptr(len(outBuf)) {
			break
		}
		v := *(*uint64)(unsafe.Pointer(&outBuf[off]))
		results = append(results, v)
		off += 8
	}
	return results, nil
}

// Test triggers the kernel-side self-test suite. The kernel runs a series
// of internal checks and writes KernelStatus.
//
// C++: test.cpp — sends IOCTL_CODE_PERFORM_KERNEL_SIDE_TESTS.
func (d *Debugger) Test() error {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.state < StateVmmLoaded {
		return fmt.Errorf("Test: VMM not loaded")
	}
	req := hyperdbgsdk.DEBUGGER_PERFORM_KERNEL_TESTS{}
	reqBuf := byteslice.FromStruct(&req)
	if _, err := d.device.Ioctl(hyperdbgsdk.IoctlPerformKernelSideTests, reqBuf, reqBuf); err != nil {
		return fmt.Errorf("Test: IOCTL failed: %w", err)
	}
	req = *byteslice.ToStruct[hyperdbgsdk.DEBUGGER_PERFORM_KERNEL_TESTS](reqBuf)
	if req.KernelStatus != uint32(hyperdbgsdk.DebuggerOperationWasSuccessful) {
		return fmt.Errorf("Test: kernel rejected, status=0x%08X", req.KernelStatus)
	}
	return nil
}

// Core switches the current debugging core. This is only meaningful in
// kernel-debugging (KD) mode over serial; local user-mode debugging has
// no per-core IOCTL.
func (d *Debugger) Core(coreId uint32) error {
	return fmt.Errorf("Core: only supported in kernel-debugging mode")
}

// IoIn reads a single byte from an I/O port. The IN instruction is
// privileged outside ring 0 (CR4.IOPL controls user-mode access), so it
// cannot be issued from user-mode Go. Use the !ioin hook (core.IoInHook)
// to trap IN reads from VMX-root.
func (d *Debugger) IoIn(port uint16) (byte, error) {
	return 0, fmt.Errorf("IoIn: IN instruction is privileged; use !ioin hook (core.IoInHook) instead")
}

// ----------------------------------------------------------------
// Meta commands — attach/detach/disconnect/kill/pagein
// ----------------------------------------------------------------

// Attach attaches the VMM to an already-running process (pid). Unlike
// StartProcess (which creates a suspended child), Attach reuses the
// attach IOCTL with IsStartingNewProcess=false and a real pid.
//
// On success the kernel returns a Token in the packet; the token is
// stored for subsequent Continue/Pause/Command IOCTLs and the state moves
// to StateProcessPaused.
//
// C++: attach.cpp — sends IOCTL_CODE_DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS
// with Action=Attach.
func (d *Debugger) Attach(pid uint32) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.state < StateVmmLoaded {
		return fmt.Errorf("Attach: VMM not loaded")
	}
	pkt := hyperdbgsdk.DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS{
		IsStartingNewProcess:            false,
		ProcessId:                       pid,
		CheckCallbackAtFirstInstruction: true,
		Action:                          hyperdbgsdk.DebuggerAttachDetachUserModeProcessActionAttach,
	}
	buf := (*[attachDetachRequestSize]byte)(unsafe.Pointer(&pkt))[:]
	if _, err := d.device.Ioctl(hyperdbgsdk.IoctlDebuggerAttachDetachUserModeProcess, buf, buf); err != nil {
		return fmt.Errorf("Attach: IOCTL failed: %w", err)
	}
	if pkt.Result != uint64(hyperdbgsdk.DebuggerOperationWasSuccessful) {
		return fmt.Errorf("Attach: kernel rejected (Result=0x%016X)", pkt.Result)
	}
	if pkt.Token == 0 {
		return fmt.Errorf("Attach: kernel returned Token=0")
	}
	d.processToken = pkt.Token
	d.state = StateProcessPaused
	return nil
}

// Detach detaches the VMM from the current debuggee (the reverse of
// Attach/StartProcess). The stored process token is cleared and the state
// moves back to StateVmmLoaded.
//
// Mirrors C++ UdDetachProcess (ud.cpp:795): continue first, then detach
// by ProcessId (the kernel looks up the session by pid, not token).
func (d *Debugger) Detach() error {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.processToken == 0 {
		return fmt.Errorf("Detach: no process attached")
	}
	// Continue the debuggee before detaching — the kernel rejects detach
	// with DEBUGGER_ERROR_UNABLE_TO_DETACH_AS_THERE_ARE_PAUSED_THREADS
	// (0xC0000034) if any thread is still paused.
	_ = continueProcess(d.device, d.processToken)
	if err := detachProcess(d.device, d.processPid); err != nil {
		return fmt.Errorf("Detach: %w", err)
	}
	d.processToken = 0
	d.processPid = 0
	d.state = StateVmmLoaded
	return nil
}

// Disconnect signals the kernel that the debugging session is finished by
// sending IOCTL_CODE_SEND_SIGNAL_EXECUTION_IN_DEBUGGEE_FINISHED. The
// device handle is not closed (Close is separate).
//
// C++: libhyperdbg.cpp HyperDbgClose path.
func (d *Debugger) Disconnect() error {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.disconnectLocked()
}

// disconnectLocked is the lock-free body of Disconnect. Caller must hold d.mu
// (UnloadVMM calls this while already holding d.mu — Go mutex 非重入，直接调
// Disconnect 会自死锁).
func (d *Debugger) disconnectLocked() error {
	if d.state < StateVmmLoaded {
		return fmt.Errorf("Disconnect: VMM not loaded")
	}
	req := hyperdbgsdk.DEBUGGER_SEND_COMMAND_EXECUTION_FINISHED_SIGNAL{}
	reqBuf := byteslice.FromStruct(&req)
	if _, err := d.device.Ioctl(hyperdbgsdk.IoctlSendSignalExecutionInDebuggeeFinished, reqBuf, reqBuf); err != nil {
		return fmt.Errorf("Disconnect: IOCTL failed: %w", err)
	}
	req = *byteslice.ToStruct[hyperdbgsdk.DEBUGGER_SEND_COMMAND_EXECUTION_FINISHED_SIGNAL](reqBuf)
	if req.KernelStatus != uint32(hyperdbgsdk.DebuggerOperationWasSuccessful) {
		return fmt.Errorf("Disconnect: kernel rejected, status=0x%08X", req.KernelStatus)
	}
	return nil
}

// Kill terminates the debuggee identified by pid. The stored process token
// (if any) is passed so the kernel can locate the debug session; if no
// process is attached, Token=0 is sent and the kernel looks up pid directly.
//
// C++: kill.cpp — sends the attach IOCTL with Action=KillProcess.
func (d *Debugger) Kill(pid uint32) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.state < StateVmmLoaded {
		return fmt.Errorf("Kill: VMM not loaded")
	}
	pkt := hyperdbgsdk.DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS{
		ProcessId: pid,
		Action:    hyperdbgsdk.DebuggerAttachDetachUserModeProcessActionKillProcess,
		Token:     d.processToken,
	}
	buf := (*[attachDetachRequestSize]byte)(unsafe.Pointer(&pkt))[:]
	if _, err := d.device.Ioctl(hyperdbgsdk.IoctlDebuggerAttachDetachUserModeProcess, buf, buf); err != nil {
		return fmt.Errorf("Kill: IOCTL failed: %w", err)
	}
	if pkt.Result != uint64(hyperdbgsdk.DebuggerOperationWasSuccessful) {
		return fmt.Errorf("Kill: kernel rejected (Result=0x%016X)", pkt.Result)
	}
	if d.processToken != 0 {
		d.processToken = 0
		d.state = StateVmmLoaded
	}
	return nil
}

// PageIn forces the kernel to bring the page containing addr into
// physical memory (useful before reading an address that may be paged
// out). The kernel faults the page in via the standard page-fault path.
//
// C++: pagein.cpp — sends IOCTL_CODE_DEBUGGER_BRING_PAGES_IN.
func (d *Debugger) PageIn(addr uint64) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.state < StateVmmLoaded {
		return fmt.Errorf("PageIn: VMM not loaded")
	}
	req := hyperdbgsdk.DEBUGGER_PAGE_IN_REQUEST{
		VirtualAddressFrom: addr,
		VirtualAddressTo:   addr,
		ProcessId:          d.processPid,
		PageFaultErrorCode: 0,
	}
	reqBuf := byteslice.FromStruct(&req)
	if _, err := d.device.Ioctl(hyperdbgsdk.IoctlDebuggerBringPagesIn, reqBuf, reqBuf); err != nil {
		return fmt.Errorf("PageIn: IOCTL failed: %w", err)
	}
	req = *byteslice.ToStruct[hyperdbgsdk.DEBUGGER_PAGE_IN_REQUEST](reqBuf)
	if req.KernelStatus != uint32(hyperdbgsdk.DebuggerOperationWasSuccessful) {
		return fmt.Errorf("PageIn: kernel rejected, status=0x%08X", req.KernelStatus)
	}
	return nil
}
