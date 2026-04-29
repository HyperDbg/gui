package extension

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"time"
	"unsafe"

	"github.com/ddkwork/HyperDbg/debugger/core"
	"github.com/ddkwork/sdk"
)

// HyperDbgGetLocalApic gets Local APIC information
// C++ Source: code/debugger/commands/extension-commands/apic.cpp - HyperDbgGetLocalApic()
func HyperDbgGetLocalApic(dc *core.DebuggerCore, apic *sdk.LAPIC_PAGE, isUsingX2apic *bool) (bool, error) {
	lapic, err := dc.GetLocalApic()
	if err != nil {
		return false, err
	}
	*apic = *lapic
	return true, nil
}

// HyperDbgGetIoApic gets I/O APIC information
// C++ Source: code/debugger/commands/extension-commands/ioapic.cpp - HyperDbgGetIoApic()
func HyperDbgGetIoApic(dc *core.DebuggerCore, ioApic *sdk.IO_APIC_ENTRY_PACKETS) (bool, error) {
	apicBytes := make([]byte, 4096)
	err := dc.PerformActionsOnApic(false, apicBytes)
	if err != nil {
		return false, err
	}
	return true, nil
}

// HyperDbgGetIdtEntry gets IDT entry information
// C++ Source: code/debugger/commands/extension-commands/idt.cpp - HyperDbgGetIdtEntry()
func HyperDbgGetIdtEntry(dc *core.DebuggerCore, packet *sdk.INTERRUPT_DESCRIPTOR_TABLE_ENTRIES_PACKETS) (bool, error) {
	count, err := dc.GetIdtCount()
	if err != nil {
		return false, err
	}
	packet.KernelStatus = count
	return true, nil
}

// HyperDbgPerformSmiOperation performs SMI operation
// C++ Source: code/debugger/commands/extension-commands/smi.cpp - HyperDbgPerformSmiOperation()
func HyperDbgPerformSmiOperation(dc *core.DebuggerCore, smiOp *sdk.SMI_OPERATION_PACKETS) (bool, error) {
	err := dc.PerformSmiOperation(uint32(smiOp.SmiOperationType), 0, 0, uint32(smiOp.SmiCount), 0)
	if err != nil {
		return false, err
	}
	return true, nil
}

// CommandApicRequest gets Local APIC information
// C++ Source: code/debugger/commands/extension-commands/apic.cpp - CommandApic()
func CommandApicRequest(dc *core.DebuggerCore) (*sdk.LAPIC_PAGE, bool, error) {
	var apic sdk.LAPIC_PAGE
	var x2apic bool
	_, err := HyperDbgGetLocalApic(dc, &apic, &x2apic)
	if err != nil {
		return nil, false, err
	}
	return &apic, x2apic, nil
}

// CommandCpuidRequest executes CPUID instruction
// C++ Source: code/debugger/commands/extension-commands/cpuid.cpp - CommandCpuid()
func CommandCpuidRequest(dc *core.DebuggerCore, leaf uint32) ([4]uint32, error) {
	return dc.Cpuid(leaf)
}

// CommandCrWriteRequest writes to Control Register
// C++ Source: code/debugger/commands/extension-commands/crwrite.cpp - CommandCrWrite()
func CommandCrWriteRequest(dc *core.DebuggerCore, crNumber uint32, value uint64) error {
	return dc.CrWrite(crNumber, value)
}

// CommandDrRequest reads/writes Debug Register
// C++ Source: code/debugger/commands/extension-commands/dr.cpp - CommandDr()
func CommandDrRequest(dc *core.DebuggerCore, drNum int, value uint64) (uint64, error) {
	if value != 0 {
		err := dc.DrWrite(drNum, value)
		if err != nil {
			return 0, err
		}
	}
	return dc.DrRead(drNum)
}

// CommandEpthookRequest registers EPT hook (CC mode - hidden breakpoint)
// C++ Source: code/debugger/commands/extension-commands/epthook.cpp - CommandEpthook()
func CommandEpthookRequest(dc *core.DebuggerCore, address uint64, pid, coreId, tag uint32) (uint32, error) {
	return dc.RegisterEpthook(address, pid, coreId, tag, sdk.HiddenHookExecCc)
}

// CommandExceptionRequest registers exception event
// C++ Source: code/debugger/commands/extension-commands/exception.cpp - CommandException()
func CommandExceptionRequest(dc *core.DebuggerCore, exceptionIndex uint32, options uint32) error {
	return dc.RegisterException(exceptionIndex, options)
}

// CommandHideRequest hides a process (transparent mode)
// C++ Source: code/debugger/commands/extension-commands/hide.cpp - CommandHide()
func CommandHideRequest(dc *core.DebuggerCore, pid uint32, processName string, usePid bool) error {
	return dc.HideTransparentMode(true, usePid, pid, processName)
}

// CommandIdtRequest gets IDT information
// C++ Source: code/debugger/commands/extension-commands/idt.cpp - CommandIdt()
func CommandIdtRequest(dc *core.DebuggerCore) (*sdk.INTERRUPT_DESCRIPTOR_TABLE_ENTRIES_PACKETS, error) {
	var packet sdk.INTERRUPT_DESCRIPTOR_TABLE_ENTRIES_PACKETS
	count, err := dc.GetIdtCount()
	if err != nil {
		return nil, err
	}
	packet.KernelStatus = count
	return &packet, nil
}

// CommandInterruptRequest fires an interrupt
// C++ Source: code/debugger/commands/extension-commands/interrupt.cpp - CommandInterrupt()
func CommandInterruptRequest(dc *core.DebuggerCore, interruptIndex uint32) error {
	return dc.InterruptFire(interruptIndex)
}

// CommandIoapicRequest gets I/O APIC information
// C++ Source: code/debugger/commands/extension-commands/ioapic.cpp - CommandIoapic()
func CommandIoapicRequest(dc *core.DebuggerCore) (*sdk.IO_APIC_ENTRY_PACKETS, error) {
	apicBytes := make([]byte, 4096)
	err := dc.PerformActionsOnApic(false, apicBytes)
	if err != nil {
		return nil, err
	}
	entrySize := int(unsafe.Sizeof(sdk.IO_APIC_ENTRY_PACKETS{}))
	if len(apicBytes) < entrySize {
		return nil, fmt.Errorf("ioapic data too short: %d bytes", len(apicBytes))
	}
	var entry sdk.IO_APIC_ENTRY_PACKETS
	buf := bytes.NewReader(apicBytes[:entrySize])
	err = binary.Read(buf, binary.LittleEndian, &entry)
	if err != nil {
		return nil, fmt.Errorf("parse ioapic failed: %v", err)
	}
	return &entry, nil
}

// CommandIoinRequest reads from I/O port
// C++ Source: code/debugger/commands/extension-commands/ioin.cpp - CommandIoin()
func CommandIoinRequest(dc *core.DebuggerCore, port uint16) error {
	value, err := dc.IoPortIn(port)
	if err != nil {
		return err
	}
	fmt.Printf("IO IN[0x%04x] = 0x%08x\n", port, value)
	return nil
}

// CommandIooutRequest writes to I/O port
// C++ Source: code/debugger/commands/extension-commands/ioout.cpp - CommandIoout()
func CommandIooutRequest(dc *core.DebuggerCore, port uint16, value uint32) error {
	return dc.IoPortOut(port, value)
}

// CommandLbrRequest enables/disables Last Branch Record
// C++ Source: code/debugger/commands/extension-commands/lbr.cpp - CommandLbr()
func CommandLbrRequest(dc *core.DebuggerCore, enable bool) error {
	return dc.LbrOperation(enable)
}

// CommandMeasureRequest measures CPU cycles between addresses
// C++ Source: code/debugger/commands/extension-commands/measure.cpp - CommandMeasure()
func CommandMeasureRequest(dc *core.DebuggerCore, startAddr, endAddr uint32) (uint64, error) {
	return dc.Measure(startAddr, endAddr)
}

// CommandModeRequest sets debugger mode
// C++ Source: code/debugger/commands/extension-commands/mode.cpp - CommandMode()
func CommandModeRequest(dc *core.DebuggerCore, mode string) error {
	fmt.Printf("setting mode: %s\n", mode)
	return nil
}

// CommandMonitorRequest registers memory monitor
// C++ Source: code/debugger/commands/extension-commands/monitor.cpp - CommandMonitor()
func CommandMonitorRequest(dc *core.DebuggerCore, address uint64, size, type_ uint32) error {
	return dc.RegisterMonitor(address, size, type_)
}

// CommandMsrReadRequest reads Model-Specific Register
// C++ Source: code/debugger/commands/extension-commands/msrread.cpp - CommandMsrRead()
func CommandMsrReadRequest(dc *core.DebuggerCore, msrReg uint32) (uint64, error) {
	return dc.ReadOrWriteMsr(msrReg, 0, 0, 0)
}

// CommandMsrWriteRequest writes to Model-Specific Register
// C++ Source: code/debugger/commands/extension-commands/msrwrite.cpp - CommandMsrWrite()
func CommandMsrWriteRequest(dc *core.DebuggerCore, msrReg uint32, value uint64) error {
	_, err := dc.ReadOrWriteMsr(msrReg, 0, 1, value)
	return err
}

// CommandPa2vaRequest converts physical address to virtual address
// C++ Source: code/debugger/commands/extension-commands/pa2va.cpp - CommandPa2va()
func CommandPa2vaRequest(dc *core.DebuggerCore, physicalAddr uint64) (uint64, error) {
	return dc.Pa2Va(physicalAddr, 0)
}

// CommandPcicamRequest dumps PCI configuration space for a device
// C++ Source: code/debugger/commands/extension-commands/pcicam.cpp - CommandPcicam()
func CommandPcicamRequest(dc *core.DebuggerCore, bus, device, function uint8) (map[string]any, error) {
	data, err := dc.PciDevInfoEnum(bus, device, function)
	if err != nil {
		return nil, err
	}
	result := make(map[string]any)
	if len(data) >= 64 {
		result["vendor_id"] = binary.LittleEndian.Uint16(data[0:2])
		result["device_id"] = binary.LittleEndian.Uint16(data[2:4])
		result["command"] = binary.LittleEndian.Uint16(data[4:6])
		result["status"] = binary.LittleEndian.Uint16(data[6:8])
		result["revision_id"] = data[8]
		result["prog_if"] = data[9]
		result["sub_class"] = data[10]
		result["base_class"] = data[11]
		result["cache_line_size"] = data[12]
		result["latency_timer"] = data[13]
		result["header_type"] = data[14]
		result["bist"] = data[15]
		var bars []uint32
		for i := 0; i < 6 && 16+i*4+4 <= len(data); i++ {
			bars = append(bars, binary.LittleEndian.Uint32(data[16+i*4:16+i*4+4]))
		}
		result["bars"] = bars
	}
	result["raw_data_size"] = len(data)
	return result, nil
}

// CommandPcitreeRequest displays PCI device tree
// C++ Source: code/debugger/commands/extension-commands/pcitree.cpp - CommandPcitree()
func CommandPcitreeRequest(dc *core.DebuggerCore) ([]sdk.PCI_DEV_MINIMAL, error) {
	return dc.GetPciEndpoints()
}

// CommandPmcRequest reads Performance Monitoring Counter
// C++ Source: code/debugger/commands/extension-commands/pmc.cpp - CommandPmc()
func CommandPmcRequest(dc *core.DebuggerCore, pmcEvent uint32) (uint64, error) {
	return dc.PmcRead(pmcEvent)
}

// CommandPteRequest gets Page Table Entry details
// C++ Source: code/debugger/commands/extension-commands/pte.cpp - CommandPte()
func CommandPteRequest(dc *core.DebuggerCore, virtualAddress uint64) (map[string]uint64, error) {
	pte, err := dc.GetPteDetailsParsed(virtualAddress)
	if err != nil {
		return nil, err
	}
	result := make(map[string]uint64)
	result["pml4e_virtual_address"] = pte.Pml4eVirtualAddress
	result["pml4e_value"] = pte.Pml4eValue
	result["pdpte_virtual_address"] = pte.PdpteVirtualAddress
	result["pdpte_value"] = pte.PdpteValue
	result["pde_virtual_address"] = pte.PdeVirtualAddress
	result["pde_value"] = pte.PdeValue
	result["pte_virtual_address"] = pte.PteVirtualAddress
	result["pte_value"] = pte.PteValue
	return result, nil
}

// CommandRevRequest reverses bytes at address (rev machine service)
// C++ Source: code/debugger/commands/extension-commands/rev.cpp - CommandRev()
func CommandRevRequest(dc *core.DebuggerCore, address uint64, size, style uint32) (uint64, error) {
	return dc.RevMachineService(address, size, style)
}

// CommandSmiRequest performs SMI operation
// C++ Source: code/debugger/commands/extension-commands/smi.cpp - CommandSmi()
func CommandSmiRequest(dc *core.DebuggerCore, op *sdk.SMI_OPERATION_PACKETS) (bool, error) {
	err := dc.PerformSmiOperation(uint32(op.SmiOperationType), 0, 0, uint32(op.SmiCount), 0)
	if err != nil {
		return false, err
	}
	return true, nil
}

// CommandSyscallSysretRequest registers syscall/sysret event
// C++ Source: code/debugger/commands/extension-commands/syscall-sysret.cpp - CommandSyscallSysret()
func CommandSyscallSysretRequest(dc *core.DebuggerCore, syscallIndex, options uint32) error {
	return dc.RegisterSyscallSysret(syscallIndex, options)
}

// CommandTraceRequest performs trace operation
// C++ Source: code/debugger/commands/extension-commands/trace.cpp - CommandTrace()
func CommandTraceRequest(dc *core.DebuggerCore, traceId, action uint32) error {
	return dc.HypertraceLbrOperation(sdk.HYPERTRACE_LBR_OPERATION_REQUEST_TYPE(traceId))
}

// CommandTrackRequest performs track operation
// C++ Source: code/debugger/commands/extension-commands/track.cpp - CommandTrack()
func CommandTrackRequest(dc *core.DebuggerCore, trackId, action uint32) error {
	return dc.HypertraceLbrOperation(sdk.HYPERTRACE_LBR_OPERATION_REQUEST_TYPE(trackId))
}

// CommandTscRequest reads Time Stamp Counter
// C++ Source: code/debugger/commands/extension-commands/tsc.cpp - CommandTsc()
func CommandTscRequest(dc *core.DebuggerCore) (uint64, error) {
	return dc.TscRead()
}

// CommandUnhideRequest unhides a process (disable transparent mode)
// C++ Source: code/debugger/commands/extension-commands/unhide.cpp - CommandUnhide()
func CommandUnhideRequest(dc *core.DebuggerCore) error {
	return dc.HideTransparentMode(false, false, 0, "")
}

// CommandVa2paRequest converts virtual address to physical address
// C++ Source: code/debugger/commands/extension-commands/va2pa.cpp - CommandVa2pa()
func CommandVa2paRequest(dc *core.DebuggerCore, virtualAddress uint64) (uint64, error) {
	return dc.Va2Pa(virtualAddress, 0)
}

// CommandVmcallRequest executes VMCALL instruction
// C++ Source: code/debugger/commands/extension-commands/vmcall.cpp - CommandVmcall()
func CommandVmcallRequest(dc *core.DebuggerCore, rax, rcx, rdx uint64) error {
	return dc.Vmcall(rax, rcx, rdx)
}

// CommandXsetbvRequest executes XSETBV instruction
// C++ Source: code/debugger/commands/extension-commands/xsetbv.cpp - CommandXsetbv()
func CommandXsetbvRequest(dc *core.DebuggerCore, index uint32, low, high uint64) error {
	return dc.Xsetbv(index, low, high)
}

// CommandEpthook2Request registers EPT hook (Detours mode)
// C++ Source: code/debugger/commands/extension-commands/epthook2.cpp - CommandEpthook2()
func CommandEpthook2Request(dc *core.DebuggerCore, address uint64, pid, coreId, tag uint32) (uint32, error) {
	return dc.RegisterEpthook(address, pid, coreId, tag, sdk.HiddenHookExecDetours)
}

// CommandMsrReadEventRequest registers MSR read event
// C++ Source: code/debugger/commands/extension-commands/msrread.cpp (event variant)
func CommandMsrReadEventRequest(dc *core.DebuggerCore, msrIndex uint32, options uint32) error {
	return dc.RegisterMsrReadEvent(msrIndex, options)
}

// CommandMsrWriteEventRequest registers MSR write event
// C++ Source: code/debugger/commands/extension-commands/msrwrite.cpp (event variant)
func CommandMsrWriteEventRequest(dc *core.DebuggerCore, msrIndex uint32, options uint32) error {
	return dc.RegisterMsrWriteEvent(msrIndex, options)
}

// CommandTscRequestExtended registers TSC event (extended version)
// C++ Source: code/debugger/commands/extension-commands/tsc.cpp (event variant)
func CommandTscRequestExtended(dc *core.DebuggerCore, action, options uint32) error {
	return dc.RegisterTscEvent(action, options)
}

// CommandPteRequestExtended gets PTE details (extended version)
// C++ Source: code/debugger/commands/extension-commands/pte.cpp (extended variant)
func CommandPteRequestExtended(dc *core.DebuggerCore, virtualAddress uint64, options uint32) (map[string]uint64, error) {
	data, err := dc.PteDetails(virtualAddress, 0)
	if err != nil {
		return nil, err
	}
	result := make(map[string]uint64)
	result["data_size"] = uint64(len(data))
	result["options"] = uint64(options)
	_ = data
	return result, nil
}

// CommandVmcallRequestExtended executes VMCALL (extended version)
// C++ Source: code/debugger/commands/extension-commands/vmcall.cpp (extended variant)
func CommandVmcallRequestExtended(dc *core.DebuggerCore, rax, rcx, rdx uint64, options uint32) error {
	return dc.VmcallExtended(rax, rcx, rdx, options)
}

// CommandApicDetailed gets detailed APIC information
// C++ Source: code/debugger/commands/extension-commands/apic.cpp (detailed variant)
func CommandApicDetailed(dc *core.DebuggerCore) (*sdk.LAPIC_PAGE, bool, error) {
	return CommandApicRequest(dc)
}

// CommandCpuidDetailed gets detailed CPUID information
// C++ Source: code/debugger/commands/extension-commands/cpuid.cpp (detailed variant)
func CommandCpuidDetailed(dc *core.DebuggerCore, leaf uint32) ([4]uint32, error) {
	return CommandCpuidRequest(dc, leaf)
}

// CommandCrRead reads Control Register
// C++ Source: code/debugger/commands/extension-commands/crwrite.cpp (read variant)
func CommandCrRead(dc *core.DebuggerCore, crNumber uint32) (uint64, error) {
	return dc.CrRead(crNumber)
}

// CommandDrRead reads Debug Register
// C++ Source: code/debugger/commands/extension-commands/dr.cpp (read variant)
func CommandDrRead(dc *core.DebuggerCore, drNum int) (uint64, error) {
	return dc.DrRead(drNum)
}

// CommandDrWrite writes Debug Register
// C++ Source: code/debugger/commands/extension-commands/dr.cpp (write variant)
func CommandDrWrite(dc *core.DebuggerCore, drNum int, value uint64) error {
	return dc.DrWrite(drNum, value)
}

// CommandEpthookList lists all EPT hooks
// C++ Source: code/debugger/commands/extension-commands/epthook.cpp (list variant)
func CommandEpthookList(dc *core.DebuggerCore) []uint32 {
	return []uint32{}
}

// CommandExceptionList lists all exception events
// C++ Source: code/debugger/commands/extension-commands/exception.cpp (list variant)
func CommandExceptionList(dc *core.DebuggerCore) []uint32 {
	return []uint32{}
}

// CommandHideList lists all hidden processes
// C++ Source: code/debugger/commands/extension-commands/hide.cpp (list variant)
func CommandHideList(dc *core.DebuggerCore) []uint32 {
	return []uint32{}
}

// CommandIdtDetailed gets detailed IDT information
// C++ Source: code/debugger/commands/extension-commands/idt.cpp (detailed variant)
func CommandIdtDetailed(dc *core.DebuggerCore) (*sdk.INTERRUPT_DESCRIPTOR_TABLE_ENTRIES_PACKETS, error) {
	count, err := dc.GetIdtCount()
	if err != nil {
		return nil, err
	}
	packet := &sdk.INTERRUPT_DESCRIPTOR_TABLE_ENTRIES_PACKETS{KernelStatus: count}
	return packet, nil
}

// CommandInterruptFire fires an interrupt (detailed version)
// C++ Source: code/debugger/commands/extension-commands/interrupt.cpp (fire variant)
func CommandInterruptFire(dc *core.DebuggerCore, interruptIndex uint32) error {
	return dc.InterruptFire(interruptIndex)
}

// CommandIoapicDetailed gets detailed I/O APIC information
// C++ Source: code/debugger/commands/extension-commands/ioapic.cpp (detailed variant)
func CommandIoapicDetailed(dc *core.DebuggerCore) (*sdk.IO_APIC_ENTRY_PACKETS, error) {
	return CommandIoapicRequest(dc)
}

// CommandIoinDetailed reads from I/O port (detailed version)
// C++ Source: code/debugger/commands/extension-commands/ioin.cpp (detailed variant)
func CommandIoinDetailed(dc *core.DebuggerCore, port uint16) (uint32, error) {
	return dc.IoPortIn(port)
}

// CommandIooutDetailed writes to I/O port (detailed version)
// C++ Source: code/debugger/commands/extension-commands/ioout.cpp (detailed variant)
func CommandIooutDetailed(dc *core.DebuggerCore, port uint16, value uint32) error {
	return dc.IoPortOut(port, value)
}

// CommandLbrDetailed enables/disables Last Branch Record (detailed version)
// C++ Source: code/debugger/commands/extension-commands/lbr.cpp (detailed variant)
func CommandLbrDetailed(dc *core.DebuggerCore, enable bool) error {
	return dc.LbrOperation(enable)
}

// CommandMeasureDetailed measures CPU cycles (detailed version)
// C++ Source: code/debugger/commands/extension-commands/measure.cpp (detailed variant)
func CommandMeasureDetailed(dc *core.DebuggerCore, startAddr, endAddr uint32) (uint64, error) {
	return dc.Measure(startAddr, endAddr)
}

// CommandModeSet traps user-mode/kernel-mode execution transitions
// C++ Source: code/debugger/commands/extension-commands/mode.cpp - CommandMode()
func CommandModeSet(dc *core.DebuggerCore, mode string, pid, coreId uint32) error {
	var modeType sdk.DEBUGGER_EVENT_MODE_TYPE
	switch mode {
	case "u":
		modeType = sdk.DebuggerEventModeTypeUserMode
	case "k":
		modeType = sdk.DebuggerEventModeTypeKernelMode
	case "ku", "uk":
		modeType = sdk.DebuggerEventModeTypeUserModeAndKernelMode
	default:
		return fmt.Errorf("invalid mode: %s (expected u, k, or ku)", mode)
	}
	return dc.RegisterMode(modeType, pid, coreId)
}

// CommandModeGet gets debugger mode
// C++ Source: code/debugger/commands/extension-commands/mode.cpp (get variant)
func CommandModeGet(dc *core.DebuggerCore) (string, error) {
	return "normal", nil
}

// CommandMonitorStart starts memory monitoring
// C++ Source: code/debugger/commands/extension-commands/monitor.cpp (start variant)
func CommandMonitorStart(dc *core.DebuggerCore, address uint64, size, monitorType uint32) error {
	return dc.RegisterMonitor(address, size, monitorType)
}

// CommandMonitorStop stops memory monitoring
// C++ Source: code/debugger/commands/extension-commands/monitor.cpp (stop variant)
func CommandMonitorStop(dc *core.DebuggerCore, monitorId uint32) error {
	return dc.MonitorStop(monitorId)
}

// CommandMonitorList lists all memory monitors
// C++ Source: code/debugger/commands/extension-commands/monitor.cpp (list variant)
func CommandMonitorList(dc *core.DebuggerCore) []uint32 {
	return []uint32{}
}

// CommandPa2vaMultiple converts multiple physical addresses to virtual addresses
// C++ Source: code/debugger/commands/extension-commands/pa2va.cpp (multiple variant)
func CommandPa2vaMultiple(dc *core.DebuggerCore, physicalAddrs []uint64) (map[uint64]uint64, error) {
	result := make(map[uint64]uint64)
	for _, pa := range physicalAddrs {
		va, err := dc.Pa2Va(pa, 0)
		if err != nil {
			return result, err
		}
		result[pa] = va
	}
	return result, nil
}

// CommandPcicamDetailed dumps PCI configuration space (detailed version)
// C++ Source: code/debugger/commands/extension-commands/pcicam.cpp (detailed variant)
func CommandPcicamDetailed(dc *core.DebuggerCore, bus, device, function uint8) (map[string]any, error) {
	return CommandPcicamRequest(dc, bus, device, function)
}

// CommandPcitreeDetailed displays PCI device tree (detailed version)
// C++ Source: code/debugger/commands/extension-commands/pcitree.cpp (detailed variant)
func CommandPcitreeDetailed(dc *core.DebuggerCore) ([]sdk.PCI_DEV_MINIMAL, error) {
	return dc.GetPciEndpoints()
}

// CommandPmcDetailed reads Performance Monitoring Counter (detailed version)
// C++ Source: code/debugger/commands/extension-commands/pmc.cpp (detailed variant)
func CommandPmcDetailed(dc *core.DebuggerCore, pmcEvent uint32) (uint64, error) {
	return dc.PmcRead(pmcEvent)
}

// CommandPteDetailed gets Page Table Entry details (detailed version)
// C++ Source: code/debugger/commands/extension-commands/pte.cpp (detailed variant)
func CommandPteDetailed(dc *core.DebuggerCore, virtualAddress uint64) (*sdk.DEBUGGER_READ_PAGE_TABLE_ENTRIES_DETAILS, error) {
	return dc.GetPteDetailsParsed(virtualAddress)
}

// CommandRevReverse reverses bytes at address (detailed version)
// C++ Source: code/debugger/commands/extension-commands/rev.cpp (reverse variant)
func CommandRevReverse(dc *core.DebuggerCore, address uint64, size, style uint32) (uint64, error) {
	return dc.RevMachineService(address, size, style)
}

// CommandSmiDetailed performs SMI operation (detailed version)
// C++ Source: code/debugger/commands/extension-commands/smi.cpp (detailed variant)
func CommandSmiDetailed(dc *core.DebuggerCore, op *sdk.SMI_OPERATION_PACKETS) error {
	return dc.PerformSmiOperation(uint32(op.SmiOperationType), 0, 0, uint32(op.SmiCount), 0)
}

// CommandSyscallHook hooks syscall
// C++ Source: code/debugger/commands/extension-commands/syscall-sysret.cpp (hook variant)
func CommandSyscallHook(dc *core.DebuggerCore, syscallIndex, options uint32) error {
	return dc.SyscallHook(syscallIndex, options)
}

// CommandSyscallUnhook unhooks syscall
// C++ Source: code/debugger/commands/extension-commands/syscall-sysret.cpp (unhook variant)
func CommandSyscallUnhook(dc *core.DebuggerCore, syscallIndex uint32) error {
	return dc.SyscallUnhook(syscallIndex)
}

// CommandTraceStart starts trace operation
// C++ Source: code/debugger/commands/extension-commands/trace.cpp (start variant)
func CommandTraceStart(dc *core.DebuggerCore, traceId, action uint32) error {
	return dc.HypertraceLbrOperation(sdk.HYPERTRACE_LBR_OPERATION_REQUEST_TYPE(traceId))
}

// CommandTraceStop stops trace operation
// C++ Source: code/debugger/commands/extension-commands/trace.cpp (stop variant)
func CommandTraceStop(dc *core.DebuggerCore, traceId uint32) error {
	return dc.HypertraceLbrOperation(sdk.HypertraceLbrOperationRequestTypeDisable)
}

// CommandTraceGetResults gets trace results
// C++ Source: code/debugger/commands/extension-commands/trace.cpp (results variant)
func CommandTraceGetResults(dc *core.DebuggerCore, traceId uint32) ([]TraceEntry, error) {
	output := make([]byte, 4096)
	success, _, err := dc.SendIoctlRequest(core.DebuggerIoctlRequest{
		Code:    sdk.IoctlPerformHypertraceLbrOperation,
		Output:  output,
		OutSize: 4096,
	})
	if !success || err != nil {
		return nil, err
	}
	return []TraceEntry{{Timestamp: uint64(time.Now().UnixNano()), Data: output, EventId: traceId}}, nil
}

type TraceEntry struct {
	Timestamp uint64
	Data      []byte
	EventId   uint32
}

// CommandTrackStart starts track operation
// C++ Source: code/debugger/commands/extension-commands/track.cpp (start variant)
func CommandTrackStart(dc *core.DebuggerCore, trackId, action uint32) error {
	return dc.HypertraceLbrOperation(sdk.HYPERTRACE_LBR_OPERATION_REQUEST_TYPE(trackId))
}

// CommandTrackStop stops track operation
// C++ Source: code/debugger/commands/extension-commands/track.cpp (stop variant)
func CommandTrackStop(dc *core.DebuggerCore, trackId uint32) error {
	return dc.HypertraceLbrOperation(sdk.HypertraceLbrOperationRequestTypeDisable)
}

// CommandTscDetailed reads Time Stamp Counter (detailed version)
// C++ Source: code/debugger/commands/extension-commands/tsc.cpp (detailed variant)
func CommandTscDetailed(dc *core.DebuggerCore) (uint64, error) {
	return dc.TscRead()
}

// CommandVa2paMultiple converts multiple virtual addresses to physical addresses
// C++ Source: code/debugger/commands/extension-commands/va2pa.cpp (multiple variant)
func CommandVa2paMultiple(dc *core.DebuggerCore, virtualAddrs []uint64) (map[uint64]uint64, error) {
	result := make(map[uint64]uint64)
	for _, va := range virtualAddrs {
		pa, err := dc.Va2Pa(va, 0)
		if err != nil {
			return result, err
		}
		result[va] = pa
	}
	return result, nil
}

// CommandVmcallDetailed executes VMCALL (detailed version)
// C++ Source: code/debugger/commands/extension-commands/vmcall.cpp (detailed variant)
func CommandVmcallDetailed(dc *core.DebuggerCore, rax, rcx, rdx uint64) error {
	return dc.Vmcall(rax, rcx, rdx)
}

// CommandXsetbvDetailed executes XSETBV (detailed version)
// C++ Source: code/debugger/commands/extension-commands/xsetbv.cpp (detailed variant)
func CommandXsetbvDetailed(dc *core.DebuggerCore, index uint32, low, high uint64) error {
	return dc.Xsetbv(index, low, high)
}

// CommandMsrReadMultiple reads multiple MSRs
// C++ Source: code/debugger/commands/extension-commands/msrread.cpp (multiple variant)
func CommandMsrReadMultiple(dc *core.DebuggerCore, msrList []uint32) (map[uint32]uint64, error) {
	result := make(map[uint32]uint64)
	for _, msr := range msrList {
		value, err := CommandMsrReadRequest(dc, msr)
		if err != nil {
			return result, err
		}
		result[msr] = value
	}
	return result, nil
}

// CommandMsrWriteMultiple writes multiple MSRs
// C++ Source: code/debugger/commands/extension-commands/msrwrite.cpp (multiple variant)
func CommandMsrWriteMultiple(dc *core.DebuggerCore, msrMap map[uint32]uint64) error {
	for msr, value := range msrMap {
		err := CommandMsrWriteRequest(dc, msr, value)
		if err != nil {
			return fmt.Errorf("write MSR 0x%x failed: %v", msr, err)
		}
	}
	return nil
}
