package walker

import (
	"fmt"
	"time"

	"github.com/ddkwork/HyperDbg_rust/protocol"
)

const (
	SyscallNtAllocateVirtualMemory = 0x12
	SyscallNtDeviceIoControlFile   = 0x07
	SMART_RCV_DRIVE_DATA           = 0x0007c088
	IOCTL_DISK_GET_DRIVE_GEOMETRY  = 0x00070000
	IOCTL_STORAGE_QUERY_PROPERTY   = 0x002d1400
	IOCTL_CPUID                    = 0x00220000
	AFD_CONNECT                    = 0x12007
)

type Walker interface {
	Maiter(event *protocol.SyscallEvent, do func())
	Hook(event *protocol.SyscallEvent, do func())
}

type ProcessBehaviorLogger struct {
	TargetPid        uint32
	SyscallCount     int
	SysretCount      int
	KmemAllocCount   int
	CpuidCount       int
	NetworkConnCount int
	MemAllocCount    int
}

func NewProcessBehaviorLogger(targetPid uint32) Walker {
	return &ProcessBehaviorLogger{TargetPid: targetPid}
}

func (l *ProcessBehaviorLogger) Maiter(event *protocol.SyscallEvent, do func()) {
	if event == nil || (l.TargetPid != 0 && event.ProcessID != l.TargetPid) {
		return
	}
	l.SyscallCount++
	do()
}

func (l *ProcessBehaviorLogger) Hook(event *protocol.SyscallEvent, do func()) {
}

func (l *ProcessBehaviorLogger) Summary() string {
	return fmt.Sprintf(
		`Process Behavior Summary:
  Syscalls:         %d
  Sysrets:          %d
  CPUID:            %d
  Memory Allocs:    %d
  Network Conns:    %d`,
		l.SyscallCount, l.SysretCount, l.CpuidCount, l.MemAllocCount, l.NetworkConnCount,
	)
}

type MemoryAllocationMonitor struct {
	TargetPid   uint32
	AllocCount  int
	TotalSize   uint64
	Allocations []MemoryAllocationInfo
}

type MemoryAllocationInfo struct {
	ProcessId     uint32
	Address       uint64
	Size          uint64
	RequestedSize uint64
}

func NewMemoryAllocationMonitor(targetPid uint32) Walker {
	return &MemoryAllocationMonitor{
		TargetPid:   targetPid,
		Allocations: make([]MemoryAllocationInfo, 0),
	}
}

func (m *MemoryAllocationMonitor) Maiter(event *protocol.SyscallEvent, do func()) {
	if event == nil || (m.TargetPid != 0 && event.ProcessID != m.TargetPid) {
		return
	}
	if event.RAX != SyscallNtAllocateVirtualMemory {
		return
	}
	m.AllocCount++
	alloc := MemoryAllocationInfo{
		ProcessId:     event.ProcessID,
		Address:       event.RDX,
		Size:          event.Options.OptionalParam1,
		RequestedSize: event.R9,
	}
	m.Allocations = append(m.Allocations, alloc)
	m.TotalSize += alloc.RequestedSize
	do()
}

func (m *MemoryAllocationMonitor) Hook(event *protocol.SyscallEvent, do func()) {
}

func (m *MemoryAllocationMonitor) Summary() string {
	return fmt.Sprintf(
		`Memory Allocation Summary:
  Total Allocations: %d
  Total Size:        %d bytes (%.2f MB)
  Average Size:      %.2f bytes`,
		m.AllocCount, m.TotalSize, float64(m.TotalSize)/1024/1024,
		float64(m.TotalSize)/float64(max(m.AllocCount, 1)),
	)
}

type ComprehensiveMonitor struct {
	TargetPid    uint32
	SyscallCount map[uint64]int
	NetworkConns []NetworkConnectionInfo
	MemAllocs    []MemoryAllocationInfo
	IoctlCalls   map[uint32]int
}

type NetworkConnectionInfo struct {
	ProcessId   uint32
	ProcessName string
	IPAddress   string
	Port        uint16
}

func NewComprehensiveMonitor(targetPid uint32) Walker {
	return &ComprehensiveMonitor{
		TargetPid:    targetPid,
		SyscallCount: make(map[uint64]int),
		NetworkConns: make([]NetworkConnectionInfo, 0),
		MemAllocs:    make([]MemoryAllocationInfo, 0),
		IoctlCalls:   make(map[uint32]int),
	}
}

func (m *ComprehensiveMonitor) Maiter(event *protocol.SyscallEvent, do func()) {
	if event == nil || (m.TargetPid != 0 && event.ProcessID != m.TargetPid) {
		return
	}
	syscallNum := event.RAX
	m.SyscallCount[syscallNum]++
	do()
}

func (m *ComprehensiveMonitor) Hook(event *protocol.SyscallEvent, do func()) {
}

func (m *ComprehensiveMonitor) Summary() string {
	return fmt.Sprintf(
		`Comprehensive Monitor Summary:
  Total Syscall Types:  %d
  Total Memory Allocs:  %d
  Total Network Conns:  %d
  Total IOCTL Types:    %d`,
		len(m.SyscallCount), len(m.MemAllocs), len(m.NetworkConns), len(m.IoctlCalls),
	)
}

type DFIRMonitor struct {
	TargetPid    uint32
	Events       []DFIREvent
	SyscallCount int
	CpuidCount   int
}

type DFIREvent struct {
	Timestamp int64
	EventType string
	ProcessId uint32
	ThreadId  uint32
	Data      map[string]any
}

func NewDFIRMonitor(targetPid uint32) Walker {
	return &DFIRMonitor{
		TargetPid: targetPid,
		Events:    make([]DFIREvent, 0),
	}
}

func (d *DFIRMonitor) Maiter(event *protocol.SyscallEvent, do func()) {
	if event == nil {
		return
	}
	d.SyscallCount++
	data := map[string]any{
		"syscall": event.RAX,
		"rcx":     event.RCX,
		"rdx":     event.RDX,
		"r8":      event.R8,
		"r9":      event.R9,
	}
	d.RecordEvent("syscall", event, data)
	do()
}

func (d *DFIRMonitor) Hook(event *protocol.SyscallEvent, do func()) {
}

func (d *DFIRMonitor) RecordEvent(eventType string, debugEvent *protocol.SyscallEvent, data map[string]any) {
	if debugEvent == nil || (d.TargetPid != 0 && debugEvent.ProcessID != d.TargetPid) {
		return
	}
	evt := DFIREvent{
		Timestamp: time.Now().Unix(),
		EventType: eventType,
		ProcessId: debugEvent.ProcessID,
		ThreadId:  debugEvent.ThreadID,
		Data:      data,
	}
	d.Events = append(d.Events, evt)
}

func (d *DFIRMonitor) Summary() string {
	eventTypes := make(map[string]int)
	for _, evt := range d.Events {
		eventTypes[evt.EventType]++
	}
	return fmt.Sprintf(
		`DFIR Monitor Summary:
  Total Events:     %d
  Syscall Events:   %d
  CPUID Events:     %d
  Event Types:      %v`,
		len(d.Events), d.SyscallCount, d.CpuidCount, eventTypes,
	)
}

type NtDeviceIoControlFileInfo struct {
	ProcessId         uint32
	ProcessName       string
	FileHandle        uint64
	IoControlCode     uint32
	InputBuffer       uint64
	InputBufferLength uint32
	OutputBuffer      uint64
	OutputBufferLen   uint32
	RAX               uint64
	RCX               uint64
	RDX               uint64
	R8                uint64
	R9                uint64
	R10               uint64
	RSP               uint64
}

func (info *NtDeviceIoControlFileInfo) String() string {
	return fmt.Sprintf(
		"PID: %d | Handle: 0x%x | IOCTL: 0x%08X | InBuf: 0x%x (len=%d) | OutBuf: 0x%x (len=%d)",
		info.ProcessId, info.FileHandle, info.IoControlCode,
		info.InputBuffer, info.InputBufferLength, info.OutputBuffer, info.OutputBufferLen,
	)
}

func ParseNtDeviceIoControlFileEvent(event *protocol.SyscallEvent) *NtDeviceIoControlFileInfo {
	if event == nil {
		return nil
	}
	return &NtDeviceIoControlFileInfo{
		ProcessId:     event.ProcessID,
		FileHandle:    event.RCX,
		IoControlCode: uint32(event.R10 >> 32),
		InputBuffer:   event.R8,
		OutputBuffer:  event.R9,
		RAX:           event.RAX,
		RCX:           event.RCX,
		RDX:           event.RDX,
		R8:            event.R8,
		R9:            event.R9,
		R10:           event.R10,
		RSP:           event.RSP,
	}
}

func IoControlCodeToString(code uint32) string {
	deviceType := (code >> 16) & 0xFFFF
	function := (code >> 2) & 0xFFF
	access := (code >> 14) & 0x3
	method := code & 0x3

	deviceNames := map[uint32]string{
		0x0001: "FILE_DEVICE_BEEP", 0x0002: "FILE_DEVICE_CD_ROM", 0x0003: "FILE_DEVICE_CD_ROM_FILE_SYSTEM",
		0x0004: "FILE_DEVICE_CONTROLLER", 0x0005: "FILE_DEVICE_DATALINK", 0x0006: "FILE_DEVICE_DFS",
		0x0007: "FILE_DEVICE_DISK", 0x0008: "FILE_DEVICE_DISK_FILE_SYSTEM", 0x0009: "FILE_DEVICE_FILE_SYSTEM",
		0x000A: "FILE_DEVICE_INPORT_PORT", 0x000B: "FILE_DEVICE_KEYBOARD", 0x000C: "FILE_DEVICE_MAILSLOT",
		0x000D: "FILE_DEVICE_MIDI_IN", 0x000E: "FILE_DEVICE_MIDI_OUT", 0x000F: "FILE_DEVICE_MOUSE",
		0x0010: "FILE_DEVICE_MULTI_UNC_PROVIDER", 0x0011: "FILE_DEVICE_NAMED_PIPE", 0x0012: "FILE_DEVICE_NETWORK",
		0x0013: "FILE_DEVICE_NETWORK_BROWSER", 0x0014: "FILE_DEVICE_NETWORK_FILE_SYSTEM", 0x0015: "FILE_DEVICE_NULL",
		0x0016: "FILE_DEVICE_PARALLEL_PORT", 0x0017: "FILE_DEVICE_PHYSICAL_NETCARD", 0x0018: "FILE_DEVICE_PRINTER",
		0x0019: "FILE_DEVICE_SCANNER", 0x001A: "FILE_DEVICE_SERIAL_MOUSE_PORT", 0x001B: "FILE_DEVICE_SERIAL_PORT",
		0x001C: "FILE_DEVICE_SCREEN", 0x001D: "FILE_DEVICE_SOUND", 0x001E: "FILE_DEVICE_STREAMS",
		0x001F: "FILE_DEVICE_TAPE", 0x0020: "FILE_DEVICE_TAPE_FILE_SYSTEM", 0x0021: "FILE_DEVICE_TRANSPORT",
		0x0022: "FILE_DEVICE_UNKNOWN", 0x0023: "FILE_DEVICE_VIDEO", 0x0024: "FILE_DEVICE_VIRTUAL_DISK",
		0x0025: "FILE_DEVICE_WAVE_IN", 0x0026: "FILE_DEVICE_WAVE_OUT", 0x0027: "FILE_DEVICE_8042_PORT",
		0x0028: "FILE_DEVICE_NETWORK_REDIRECTOR", 0x0029: "FILE_DEVICE_BATTERY", 0x002A: "FILE_DEVICE_BUS_EXTENDER",
		0x002B: "FILE_DEVICE_MODEM", 0x002C: "FILE_DEVICE_VDM", 0x002D: "FILE_DEVICE_MASS_STORAGE",
		0x002E: "FILE_DEVICE_SMB", 0x002F: "FILE_DEVICE_KS", 0x0030: "FILE_DEVICE_CHANGER",
		0x0031: "FILE_DEVICE_SMARTCARD", 0x0032: "FILE_DEVICE_ACPI", 0x0033: "FILE_DEVICE_DVD",
		0x0034: "FILE_DEVICE_FULLSCREEN_VIDEO", 0x0035: "FILE_DEVICE_DFS_FILE_SYSTEM",
		0x0036: "FILE_DEVICE_DFS_VOLUME", 0x0037: "FILE_DEVICE_SERENUM", 0x0038: "FILE_DEVICE_TERMSRV",
		0x0039: "FILE_DEVICE_KSEC", 0x003A: "FILE_DEVICE_FIPS", 0x003B: "FILE_DEVICE_INFINIBAND",
	}

	methodNames := map[uint32]string{
		0: "METHOD_BUFFERED", 1: "METHOD_IN_DIRECT", 2: "METHOD_OUT_DIRECT", 3: "METHOD_NEITHER",
	}

	accessNames := map[uint32]string{
		0: "FILE_ANY_ACCESS", 1: "FILE_READ_ACCESS", 2: "FILE_WRITE_ACCESS", 3: "FILE_READ_WRITE_ACCESS",
	}

	deviceName := deviceNames[deviceType]
	if deviceName == "" {
		deviceName = fmt.Sprintf("DEVICE_0x%04X", deviceType)
	}

	return fmt.Sprintf("%s:0x%03X %s %s (0x%08X)",
		deviceName, function, accessNames[access], methodNames[method], code)
}

type HardwareHookCallback struct {
	hookInfo *HookInfo
}

type HookInfo struct {
	SsdNumber SsdInfo
	Cpu1      CpuInfo
}

type SsdInfo struct {
	Mod    StringData
	Serial StringData
}

type StringData struct {
	Data [20]byte
}

func (s *StringData) String() string {
	end := 0
	for i, b := range s.Data {
		if b == 0 {
			end = i
			break
		}
	}
	return string(s.Data[:end])
}

type CpuInfo struct {
	Eax uint32
	Ecx uint32
	Edx uint32
}

func NewHardwareHookCallback() Walker {
	return &HardwareHookCallback{hookInfo: MyNotBookHook()}
}

func MyNotBookHook() *HookInfo {
	return &HookInfo{
		SsdNumber: SsdInfo{
			Mod:    StringData{Data: [20]byte{0x53, 0x61, 0x6D, 0x73, 0x75, 0x6E, 0x67, 0x20, 0x39, 0x38, 0x30, 0x20, 0x50, 0x52, 0x4F, 0x20, 0x35, 0x31, 0x32, 0x47}},
			Serial: StringData{Data: [20]byte{0x53, 0x36, 0x32, 0x4E, 0x58, 0x4D, 0x30, 0x31, 0x30, 0x34, 0x30, 0x33, 0x30, 0x37, 0x38, 0x30, 0x30, 0x30, 0x00, 0x00}},
		},
		Cpu1: CpuInfo{
			Eax: 0x000306A9,
			Ecx: 0x00000000,
			Edx: 0x00000000,
		},
	}
}

func (h *HardwareHookCallback) Maiter(event *protocol.SyscallEvent, do func()) {
}

func (h *HardwareHookCallback) Hook(event *protocol.SyscallEvent, do func()) {
	if event == nil {
		return
	}
	info := ParseNtDeviceIoControlFileEvent(event)
	if info == nil {
		return
	}
	switch info.IoControlCode {
	case SMART_RCV_DRIVE_DATA, IOCTL_STORAGE_QUERY_PROPERTY, IOCTL_CPUID:
		do()
	}
}

func (h *HardwareHookCallback) GetHookInfo() *HookInfo {
	return h.hookInfo
}

type IopXxxControlFileMonitor struct {
	targetIoctlCodes map[uint32]string
	eventCount       int
	matchedCount     int
	hookInfo         *HookInfo
}

func NewIopXxxControlFileMonitor() Walker {
	return &IopXxxControlFileMonitor{
		targetIoctlCodes: map[uint32]string{
			SMART_RCV_DRIVE_DATA:          "SMART_RCV_DRIVE_DATA",
			IOCTL_DISK_GET_DRIVE_GEOMETRY: "IOCTL_DISK_GET_DRIVE_GEOMETRY",
			IOCTL_STORAGE_QUERY_PROPERTY:  "IOCTL_STORAGE_QUERY_PROPERTY",
			IOCTL_CPUID:                   "IOCTL_CPUID",
			AFD_CONNECT:                   "AFD_CONNECT",
		},
		hookInfo: MyNotBookHook(),
	}
}

func (m *IopXxxControlFileMonitor) Maiter(event *protocol.SyscallEvent, do func()) {
	if event == nil {
		return
	}
	m.eventCount++
	do()
}

func (m *IopXxxControlFileMonitor) Hook(event *protocol.SyscallEvent, do func()) {
	if event == nil {
		return
	}
	info := ParseNtDeviceIoControlFileEvent(event)
	if info == nil {
		return
	}
	if _, found := m.targetIoctlCodes[info.IoControlCode]; found {
		m.matchedCount++
		do()
	}
}

func (m *IopXxxControlFileMonitor) Summary() string {
	return fmt.Sprintf("Total events: %d, Matched: %d", m.eventCount, m.matchedCount)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
