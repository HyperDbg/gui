// Package core — control.go
//
// Real implementations of:
//   - event management commands (events c|d|e <tag>) — Clear/Disable/Enable
//   - control/query "!" commands (!hide/!unhide/!va2pa/!pa2va/!pte/!idt/!rev)
//   - debugger utility commands (prealloc/preactivate/flush)
//
// Each method packs the corresponding SDK struct and sends the IOCTL that
// the C++ libhyperdbg sends for the same command, then interprets the
// kernel's KernelStatus field. The api layer (api/commands_extension.go,
// api/commands_debugging.go) delegates to these methods.
//
// References (C++):
//
//	events: debugging-commands/events.cpp
//	!hide/!unhide: extension-commands/hide.cpp + HyperDbgDisableTransparentMode
//	!va2pa/!pa2va: extension-commands/va2pa.cpp / pa2va.cpp
//	!pte: extension-commands/pte.cpp
//	!idt: extension-commands/idt.cpp
//	!rev: extension-commands/rev.cpp + RevRequestService
//	prealloc: debugging-commands/prealloc.cpp
//	preactivate: debugging-commands/preactivate.cpp
//	flush: debugging-commands/flush.cpp
package core

import (
	"fmt"

	"github.com/ddkwork/golibrary/byteslice"
	"github.com/ddkwork/hyperdbgsdk"
	"github.com/hyperdbg/go-libhyperdbg/debugger/comm"
)

// debuggerOperationWasSuccessful is DEBUGGER_OPERATION_WAS_SUCCESSFUL
// (ErrorCodes.h:23) — the kernel writes this into KernelStatus when an
// IOCTL completed without error.
const debuggerOperationWasSuccessful uint32 = 0xFFFFFFFF

// ----------------------------------------------------------------
// Event management — mirrors events.cpp 'e'/'d'/'c' subcommands.
// ----------------------------------------------------------------

// modifyEvent sends IOCTL_CODE_DEBUGGER_MODIFY_EVENTS with the given
// action (Enable/Disable/Clear) for the given tag. tag==0 with Clear
// means "clear all events" (used by UnloadVMM).
func (d *Debugger) modifyEvent(tag uint64, action hyperdbgsdk.DEBUGGER_MODIFY_EVENTS_TYPE) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.device == nil {
		return fmt.Errorf("modifyEvent: not connected")
	}
	req := hyperdbgsdk.DEBUGGER_MODIFY_EVENTS{
		Tag:          tag,
		TypeOfAction: action,
	}
	reqBuf := byteslice.FromStruct(&req)
	var dummy [256]byte
	if _, err := d.device.Ioctl(comm.IOCTL_CODE_DEBUGGER_MODIFY_EVENTS, reqBuf, dummy[:]); err != nil {
		return fmt.Errorf("modifyEvent(%v, tag=%d): IOCTL failed: %w", action, tag, err)
	}
	return nil
}

// ClearEvent removes the event with the given tag (events c <tag>).
// The kernel frees the event and any attached actions.
func (d *Debugger) ClearEvent(tag uint64) error {
	return d.modifyEvent(tag, hyperdbgsdk.DebuggerModifyEventsClear)
}

// DisableEvent temporarily disables the event with the given tag
// (events d <tag>). The event configuration is preserved so it can be
// re-enabled with EnableEvent.
func (d *Debugger) DisableEvent(tag uint64) error {
	return d.modifyEvent(tag, hyperdbgsdk.DebuggerModifyEventsDisable)
}

// EnableEvent re-enables a previously disabled event (events e <tag>).
func (d *Debugger) EnableEvent(tag uint64) error {
	return d.modifyEvent(tag, hyperdbgsdk.DebuggerModifyEventsEnable)
}

// ----------------------------------------------------------------
// Transparency — !hide / !unhide
// ----------------------------------------------------------------

// Hide enables transparent mode for the given process so that HyperDbg
// evades common anti-debugging / anti-hypervisor checks. If pid==0 the
// currently attached debuggee is used (when one exists).
//
// C++: hide.cpp — HyperDbgEnableTransparentMode sends
// IOCTL_DEBUGGER_HIDE_AND_UNHIDE_TO_TRANSPARENT_THE_DEBUGGER with
// IsHide=TRUE and the process id / name. The Go API only supports pid
// (not process name) because process-name resolution requires the PE
// helpers that live in the C++ libhyperdbg.
func (d *Debugger) Hide(pid uint32) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.state < StateVmmLoaded {
		return fmt.Errorf("Hide: VMM not loaded")
	}
	req := hyperdbgsdk.DEBUGGER_HIDE_AND_TRANSPARENT_DEBUGGER_MODE{
		IsHide:                               true,
		TrueIfProcessIdAndFalseIfProcessName: true,
		ProcId:                               pid,
	}
	reqBuf := byteslice.FromStruct(&req)
	if _, err := d.device.Ioctl(comm.IOCTL_CODE_DEBUGGER_HIDE_AND_UNHIDE_TO_TRANSPARENT_THE_DEBUGGER,
		reqBuf, reqBuf); err != nil {
		return fmt.Errorf("Hide: IOCTL failed: %w", err)
	}
	req = *byteslice.ToStruct[hyperdbgsdk.DEBUGGER_HIDE_AND_TRANSPARENT_DEBUGGER_MODE](reqBuf)
	if req.KernelStatus != debuggerOperationWasSuccessful {
		return fmt.Errorf("Hide: kernel rejected, status=0x%08X", req.KernelStatus)
	}
	return nil
}

// Unhide disables transparent mode (the inverse of Hide).
//
// C++: unhide.cpp sends the same IOCTL with IsHide=FALSE.
func (d *Debugger) Unhide() error {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.state < StateVmmLoaded {
		return fmt.Errorf("Unhide: VMM not loaded")
	}
	req := hyperdbgsdk.DEBUGGER_HIDE_AND_TRANSPARENT_DEBUGGER_MODE{
		IsHide: false,
	}
	reqBuf := byteslice.FromStruct(&req)
	if _, err := d.device.Ioctl(comm.IOCTL_CODE_DEBUGGER_HIDE_AND_UNHIDE_TO_TRANSPARENT_THE_DEBUGGER,
		reqBuf, reqBuf); err != nil {
		return fmt.Errorf("Unhide: IOCTL failed: %w", err)
	}
	req = *byteslice.ToStruct[hyperdbgsdk.DEBUGGER_HIDE_AND_TRANSPARENT_DEBUGGER_MODE](reqBuf)
	if req.KernelStatus != debuggerOperationWasSuccessful {
		return fmt.Errorf("Unhide: kernel rejected, status=0x%08X", req.KernelStatus)
	}
	return nil
}

// ----------------------------------------------------------------
// Address translation — !va2pa / !pa2va
// ----------------------------------------------------------------

// Va2Pa converts a virtual address to a physical address for the given
// process. If pid==0 the debugger process itself is used (System when
// kernel-debugging).
//
// C++: va2pa.cpp — sends IOCTL_DEBUGGER_VA2PA_AND_PA2VA_COMMANDS with
// IsVirtual2Physical=TRUE.
func (d *Debugger) VirtualToPhysical(va uint64, pid uint32) (uint64, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.state < StateVmmLoaded {
		return 0, fmt.Errorf("Va2Pa: VMM not loaded")
	}
	req := hyperdbgsdk.DEBUGGER_VA2PA_AND_PA2VA_COMMANDS{
		VirtualAddress:     va,
		ProcessId:          pid,
		IsVirtual2Physical: true,
	}
	reqBuf := byteslice.FromStruct(&req)
	if _, err := d.device.Ioctl(comm.IOCTL_CODE_DEBUGGER_VA2PA_AND_PA2VA_COMMANDS,
		reqBuf, reqBuf); err != nil {
		return 0, fmt.Errorf("Va2Pa: IOCTL failed: %w", err)
	}
	req = *byteslice.ToStruct[hyperdbgsdk.DEBUGGER_VA2PA_AND_PA2VA_COMMANDS](reqBuf)
	if req.KernelStatus != debuggerOperationWasSuccessful {
		return 0, fmt.Errorf("Va2Pa: kernel rejected, status=0x%08X", req.KernelStatus)
	}
	return req.PhysicalAddress, nil
}

// Pa2Va converts a physical address to a virtual address for the given
// process (the inverse of Va2Pa).
//
// C++: pa2va.cpp — same IOCTL with IsVirtual2Physical=FALSE.
func (d *Debugger) PhysicalToVirtual(pa uint64, pid uint32) (uint64, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.state < StateVmmLoaded {
		return 0, fmt.Errorf("Pa2Va: VMM not loaded")
	}
	req := hyperdbgsdk.DEBUGGER_VA2PA_AND_PA2VA_COMMANDS{
		PhysicalAddress:    pa,
		ProcessId:          pid,
		IsVirtual2Physical: false,
	}
	reqBuf := byteslice.FromStruct(&req)
	if _, err := d.device.Ioctl(comm.IOCTL_CODE_DEBUGGER_VA2PA_AND_PA2VA_COMMANDS,
		reqBuf, reqBuf); err != nil {
		return 0, fmt.Errorf("Pa2Va: IOCTL failed: %w", err)
	}
	req = *byteslice.ToStruct[hyperdbgsdk.DEBUGGER_VA2PA_AND_PA2VA_COMMANDS](reqBuf)
	if req.KernelStatus != debuggerOperationWasSuccessful {
		return 0, fmt.Errorf("Pa2Va: kernel rejected, status=0x%08X", req.KernelStatus)
	}
	return req.VirtualAddress, nil
}

// ----------------------------------------------------------------
// Page table — !pte
// ----------------------------------------------------------------

// Pte reads the page-table entry chain for the given virtual address and
// returns the full PTE details (PML4E/PDPTE/PDE/PTE virtual addresses and
// values). pid==0 means the debugger process.
//
// C++: pte.cpp — sends IOCTL_DEBUGGER_READ_PAGE_TABLE_ENTRIES_DETAILS.
func (d *Debugger) PageTableEntry(va uint64, pid uint32) (hyperdbgsdk.DEBUGGER_READ_PAGE_TABLE_ENTRIES_DETAILS, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.state < StateVmmLoaded {
		return hyperdbgsdk.DEBUGGER_READ_PAGE_TABLE_ENTRIES_DETAILS{}, fmt.Errorf("Pte: VMM not loaded")
	}
	req := hyperdbgsdk.DEBUGGER_READ_PAGE_TABLE_ENTRIES_DETAILS{
		VirtualAddress: va,
		ProcessId:      pid,
	}
	reqBuf := byteslice.FromStruct(&req)
	if _, err := d.device.Ioctl(comm.IOCTL_CODE_DEBUGGER_READ_PAGE_TABLE_ENTRIES_DETAILS,
		reqBuf, reqBuf); err != nil {
		return hyperdbgsdk.DEBUGGER_READ_PAGE_TABLE_ENTRIES_DETAILS{}, fmt.Errorf("Pte: IOCTL failed: %w", err)
	}
	req = *byteslice.ToStruct[hyperdbgsdk.DEBUGGER_READ_PAGE_TABLE_ENTRIES_DETAILS](reqBuf)
	return req, nil
}

// ----------------------------------------------------------------
// IDT — !idt
// ----------------------------------------------------------------

// Idt reads all 256 IDT entries and returns them. The kernel writes
// KernelStatus into the first field.
//
// C++: idt.cpp — sends IOCTL_QUERY_IDT_ENTRY.
func (d *Debugger) InterruptDescriptorTable() (hyperdbgsdk.INTERRUPT_DESCRIPTOR_TABLE_ENTRIES_PACKETS, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.state < StateVmmLoaded {
		return hyperdbgsdk.INTERRUPT_DESCRIPTOR_TABLE_ENTRIES_PACKETS{}, fmt.Errorf("Idt: VMM not loaded")
	}
	req := hyperdbgsdk.INTERRUPT_DESCRIPTOR_TABLE_ENTRIES_PACKETS{}
	reqBuf := byteslice.FromStruct(&req)
	if _, err := d.device.Ioctl(comm.IOCTL_CODE_QUERY_IDT_ENTRY,
		reqBuf, reqBuf); err != nil {
		return hyperdbgsdk.INTERRUPT_DESCRIPTOR_TABLE_ENTRIES_PACKETS{}, fmt.Errorf("Idt: IOCTL failed: %w", err)
	}
	req = *byteslice.ToStruct[hyperdbgsdk.INTERRUPT_DESCRIPTOR_TABLE_ENTRIES_PACKETS](reqBuf)
	return req, nil
}

// ----------------------------------------------------------------
// Reversing machine — !rev
// ----------------------------------------------------------------

// Rev triggers the reversing-machine memory reconstruction for the given
// process. mode selects user/kernel; typ selects reconstruct/pattern.
//
// C++: rev.cpp → RevRequestService sends
// IOCTL_REQUEST_REV_MACHINE_SERVICE.
func (d *Debugger) Revision(pid uint32, mode hyperdbgsdk.REVERSING_MACHINE_RECONSTRUCT_MEMORY_MODE, typ hyperdbgsdk.REVERSING_MACHINE_RECONSTRUCT_MEMORY_TYPE) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.state < StateVmmLoaded {
		return fmt.Errorf("Rev: VMM not loaded")
	}
	req := hyperdbgsdk.REVERSING_MACHINE_RECONSTRUCT_MEMORY_REQUEST{
		ProcessId: pid,
		Mode:      mode,
		Type:      typ,
	}
	reqBuf := byteslice.FromStruct(&req)
	if _, err := d.device.Ioctl(comm.IOCTL_CODE_REQUEST_REV_MACHINE_SERVICE,
		reqBuf, reqBuf); err != nil {
		return fmt.Errorf("Rev: IOCTL failed: %w", err)
	}
	req = *byteslice.ToStruct[hyperdbgsdk.REVERSING_MACHINE_RECONSTRUCT_MEMORY_REQUEST](reqBuf)
	if req.KernelStatus != debuggerOperationWasSuccessful {
		return fmt.Errorf("Rev: kernel rejected, status=0x%08X", req.KernelStatus)
	}
	return nil
}

// ----------------------------------------------------------------
// Debugger utility — prealloc / preactivate / flush
// ----------------------------------------------------------------

// Prealloc reserves count pre-allocated kernel pools of the given type
// (e.g. DebuggerPreallocCommandTypeEpthook2). Pre-allocation avoids
// pool-allocation failures in VMX-root where paged pool is unavailable.
//
// C++: prealloc.cpp — sends IOCTL_RESERVE_PRE_ALLOCATED_POOLS.
func (d *Debugger) Prealloc(typ hyperdbgsdk.DEBUGGER_PREALLOC_COMMAND_TYPE, count uint32) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.state < StateVmmLoaded {
		return fmt.Errorf("Prealloc: VMM not loaded")
	}
	req := hyperdbgsdk.DEBUGGER_PREALLOC_COMMAND{
		Type:  typ,
		Count: count,
	}
	reqBuf := byteslice.FromStruct(&req)
	if _, err := d.device.Ioctl(comm.IOCTL_CODE_RESERVE_PRE_ALLOCATED_POOLS,
		reqBuf, reqBuf); err != nil {
		return fmt.Errorf("Prealloc: IOCTL failed: %w", err)
	}
	req = *byteslice.ToStruct[hyperdbgsdk.DEBUGGER_PREALLOC_COMMAND](reqBuf)
	if req.KernelStatus != debuggerOperationWasSuccessful {
		return fmt.Errorf("Prealloc: kernel rejected, status=0x%08X", req.KernelStatus)
	}
	return nil
}

// Preactivate pre-activates the given functionality (currently only
// mode-change traps) so the kernel arms the corresponding VMX control
// before the event is registered.
//
// C++: preactivate.cpp — sends IOCTL_PREACTIVATE_FUNCTIONALITY.
func (d *Debugger) Preactivate(typ hyperdbgsdk.DEBUGGER_PREACTIVATE_COMMAND_TYPE) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.state < StateVmmLoaded {
		return fmt.Errorf("Preactivate: VMM not loaded")
	}
	req := hyperdbgsdk.DEBUGGER_PREACTIVATE_COMMAND{
		Type: typ,
	}
	reqBuf := byteslice.FromStruct(&req)
	if _, err := d.device.Ioctl(comm.IOCTL_CODE_PREACTIVATE_FUNCTIONALITY,
		reqBuf, reqBuf); err != nil {
		return fmt.Errorf("Preactivate: IOCTL failed: %w", err)
	}
	req = *byteslice.ToStruct[hyperdbgsdk.DEBUGGER_PREACTIVATE_COMMAND](reqBuf)
	if req.KernelStatus != debuggerOperationWasSuccessful {
		return fmt.Errorf("Preactivate: kernel rejected, status=0x%08X", req.KernelStatus)
	}
	return nil
}

// Flush flushes all pending kernel logging buffers to user mode so the
// message pump can drain them. Useful after disabling short-circuiting
// or before reading a log file synchronously.
//
// C++: flush.cpp — sends IOCTL_DEBUGGER_FLUSH_LOGGING_BUFFERS.
func (d *Debugger) Flush() error {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.device == nil {
		return fmt.Errorf("Flush: not connected")
	}
	if _, err := d.device.Ioctl(comm.IOCTL_CODE_DEBUGGER_FLUSH_LOGGING_BUFFERS,
		nil, nil); err != nil {
		return fmt.Errorf("Flush: IOCTL failed: %w", err)
	}
	return nil
}
