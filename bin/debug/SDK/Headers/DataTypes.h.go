// Code generated by gengo. DO NOT EDIT.
package HPRDBGCTRL

import (
	"unsafe"
	"github.com/can1357/gengo/gengort"
)

const GengoLibraryName = "HPRDBGCTRL"

var GengoLibrary = gengort.NewLibrary(GengoLibraryName)

// @brief Different levels of paging
type PagingLevel int32

const (
	PAGING_LEVEL_PAGE_TABLE                   PagingLevel = 0
	PAGING_LEVEL_PAGE_DIRECTORY               PagingLevel = 1
	PAGING_LEVEL_PAGE_DIRECTORY_POINTER_TABLE PagingLevel = 2
	PAGING_LEVEL_PAGE_MAP_LEVEL4              PagingLevel = 3
)

// @brief Inum of intentions for buffers (buffer tag)
type PoolAllocationIntention int32

const (
	TRACKING_HOOKED_PAGES                  PoolAllocationIntention = 0
	EXEC_TRAMPOLINE                        PoolAllocationIntention = 1
	SPLIT_2MB_PAGING_TO_4KB_PAGE           PoolAllocationIntention = 2
	DETOUR_HOOK_DETAILS                    PoolAllocationIntention = 3
	BREAKPOINT_DEFINITION_STRUCTURE        PoolAllocationIntention = 4
	PROCESS_THREAD_HOLDER                  PoolAllocationIntention = 5
	INSTANT_REGULAR_EVENT_BUFFER           PoolAllocationIntention = 6
	INSTANT_BIG_EVENT_BUFFER               PoolAllocationIntention = 7
	INSTANT_REGULAR_EVENT_ACTION_BUFFER    PoolAllocationIntention = 8
	INSTANT_BIG_EVENT_ACTION_BUFFER        PoolAllocationIntention = 9
	INSTANT_REGULAR_SAFE_BUFFER_FOR_EVENTS PoolAllocationIntention = 10
	INSTANT_BIG_SAFE_BUFFER_FOR_EVENTS     PoolAllocationIntention = 11
)

// ///////////////////////////////////////////////
type DebugRegisterType int32

const (
	BREAK_ON_INSTRUCTION_FETCH              DebugRegisterType = 0
	BREAK_ON_WRITE_ONLY                     DebugRegisterType = 1
	BREAK_ON_IO_READ_OR_WRITE_NOT_SUPPORTED DebugRegisterType = 2
	BREAK_ON_READ_AND_WRITE_BUT_NOT_FETCH   DebugRegisterType = 3
)

// ///////////////////////////////////////////////
type VmxExecutionMode int32

const (
	VMX_EXECUTION_MODE_NON_ROOT VmxExecutionMode = 0
	VMX_EXECUTION_MODE_ROOT     VmxExecutionMode = 1
)

// @brief Type of calling the event
type VmmCallbackEventCallingStageType int32

const (
	VMM_CALLBACK_CALLING_STAGE_INVALID_EVENT_EMULATION VmmCallbackEventCallingStageType = 0
	VMM_CALLBACK_CALLING_STAGE_PRE_EVENT_EMULATION     VmmCallbackEventCallingStageType = 1
	VMM_CALLBACK_CALLING_STAGE_POST_EVENT_EMULATION    VmmCallbackEventCallingStageType = 2
	VMM_CALLBACK_CALLING_STAGE_ALL_EVENT_EMULATION     VmmCallbackEventCallingStageType = 3
)

// @brief enum to query different process and thread interception mechanisms
type DebuggerThreadProcessTracing int32

const (
	DEBUGGER_THREAD_PROCESS_TRACING_INTERCEPT_CLOCK_INTERRUPTS_FOR_THREAD_CHANGE  DebuggerThreadProcessTracing = 0
	DEBUGGER_THREAD_PROCESS_TRACING_INTERCEPT_CLOCK_INTERRUPTS_FOR_PROCESS_CHANGE DebuggerThreadProcessTracing = 1
	DEBUGGER_THREAD_PROCESS_TRACING_INTERCEPT_CLOCK_DEBUG_REGISTER_INTERCEPTION   DebuggerThreadProcessTracing = 2
	DEBUGGER_THREAD_PROCESS_TRACING_INTERCEPT_CLOCK_WAITING_FOR_MOV_CR3_VM_EXITS  DebuggerThreadProcessTracing = 3
)

// @brief Type of transferring buffer between user-to-kernel
type NotifyType int32

const (
	IRP_BASED   NotifyType = 0
	EVENT_BASED NotifyType = 1
)

// @brief different type of memory addresses
type DebuggerHookMemoryType int32

const (
	DEBUGGER_MEMORY_HOOK_VIRTUAL_ADDRESS  DebuggerHookMemoryType = 0
	DEBUGGER_MEMORY_HOOK_PHYSICAL_ADDRESS DebuggerHookMemoryType = 1
)

type DebuggeeUserInputPacket struct {
	CommandLen           int32
	IgnoreFinishedSignal int32
	Result               int32
}
type DebuggeeEventAndActionHeaderForRemotePacket struct {
	Length int32
}
type DebuggerPausePacketReceived struct {
	Result int32
}
type DebuggerTriggeredEventDetails struct {
	Tag     int32
	Context int32
	Stage   VmmCallbackEventCallingStageType
}
type DebuggeeKdPausedPacket struct {
	Rip                    int32
	IsProcessorOn32BitMode int32
	IgnoreDisassembling    int32
	PausingReason          int32
	CurrentCore            int32
	EventTag               int32
	EventCallingStage      VmmCallbackEventCallingStageType
	Rflags                 int32
	InstructionBytesOnRip  int32
	ReadInstructionLen     int32
}
type DebuggeeUdPausedPacket struct {
	Rip                   int32
	ProcessDebuggingToken int32
	Is32Bit               int32
	PausingReason         int32
	ProcessId             int32
	ThreadId              int32
	Rflags                int32
	EventTag              int32
	EventCallingStage     VmmCallbackEventCallingStageType
	InstructionBytesOnRip int32
	ReadInstructionLen    int32
	GuestRegs             int32
}
type DebuggeeMessagePacket struct {
	OperationCode int32
	Message       int32
}
type RegisterNotifyBuffer struct {
	Type   NotifyType
	hEvent int32
}
type DirectVmcallParameters struct {
	OptionalParam1 int32
	OptionalParam2 int32
	OptionalParam3 int32
}
type EptHooksContext struct {
	HookingTag      int32
	PhysicalAddress int32
	VirtualAddress  int32
}
type EptHooksAddressDetailsForMemoryMonitor struct {
	StartAddress    int32
	EndAddress      int32
	SetHookForRead  int32
	SetHookForWrite int32
	SetHookForExec  int32
	MemoryType      DebuggerHookMemoryType
	Tag             int32
}
type EptHooksAddressDetailsForEpthook2 struct {
	TargetAddress int32
	HookFunction  int32
}
type EptSingleHookUnhookingDetails struct {
	CallerNeedsToRestoreEntryAndInvalidateEpt int32
	RemoveBreakpointInterception              int32
	PhysicalAddress                           int32
	OriginalEntry                             int32
}
type Anon376_9 struct {
	Raw [1]int32
}
type Anon378_5 struct {
	// [Bits 3:0] Segment type.
	Type int32
	// [Bit 4] S - Descriptor type (0 = system; 1 = code or data).
	DescriptorType int32
	// [Bits 6:5] DPL - Descriptor privilege level.
	DescriptorPrivilegeLevel int32
	// [Bit 7] P - Segment present.
	Present   int32
	Reserved1 int32
	// [Bit 12] AVL - Available for use by system software.
	AvailableBit int32
	// [Bit 13] Reserved (except for CS). L - 64-bit mode active (for CS only).
	LongMode int32
	// [Bit 14] D/B - Default operation size (0 = 16-bit segment; 1 = 32-bit segment).
	DefaultBig int32
	// [Bit 15] G - Granularity.
	Granularity int32
	// [Bit 16] Segment unusable (0 = usable; 1 = unusable).
	Unusable  int32
	Reserved2 int32
}
type VmxSegmentSelector struct {
	Selector   int32
	Attributes Anon376_9
	Limit      int32
	Base       int32
}
type _Int128T = any
type _Uint128T = any
type __NSConstantString = any
type SizeT = uint64
type _BuiltinMsVaList = *byte
type _BuiltinVaList = *byte

// @brief Callback type that can be used to be used
// as a custom ShowMessages function
type Callback = unsafe.Pointer

// @brief The structure of user-input packet in HyperDbg
type PdebuggeeUserInputPacket = *DebuggeeUserInputPacket

// @brief The structure of user-input packet in HyperDbg
type PdebuggeeEventAndActionHeaderForRemotePacket = *DebuggeeEventAndActionHeaderForRemotePacket

// @brief request to pause and halt the system
type PdebuggerPausePacketReceived = *DebuggerPausePacketReceived

// @brief The structure of detail of a triggered event in HyperDbg
//
// @details This structure is also used for transferring breakpoint ids, RIP as the context, etc.
type PdebuggerTriggeredEventDetails = *DebuggerTriggeredEventDetails

// @brief The structure of pausing packet in kHyperDbg
type PdebuggeeKdPausedPacket = *DebuggeeKdPausedPacket

// @brief The structure of pausing packet in uHyperDbg
type PdebuggeeUdPausedPacket = *DebuggeeUdPausedPacket

// @brief The structure of message packet in HyperDbg
type PdebuggeeMessagePacket = *DebuggeeMessagePacket

// @brief Used to register event for transferring buffer between user-to-kernel
type PregisterNotifyBuffer = *RegisterNotifyBuffer

// @brief Used for sending direct VMCALLs on the VMX root-mode
type PdirectVmcallParameters = *DirectVmcallParameters

// @brief Temporary $context used in some EPT hook commands
type PeptHooksContext = *EptHooksContext

// @brief Setting details for EPT Hooks (!monitor)
type PeptHooksAddressDetailsForMemoryMonitor = *EptHooksAddressDetailsForMemoryMonitor

// @brief Setting details for EPT Hooks (!epthook2)
type PeptHooksAddressDetailsForEpthook2 = *EptHooksAddressDetailsForEpthook2

// @brief Details of unhooking single EPT hooks
type PeptSingleHookUnhookingDetails = *EptSingleHookUnhookingDetails

// @brief Describe segment selector in VMX
//
// @details This structure is copied from ia32.h to the SDK to
// be used as a data type for functions
type VmxSegmentAccessRightsType = Anon376_9

// @brief Segment selector
type PvmxSegmentSelector = *VmxSegmentSelector

func (s Anon376_9) AsUInt() int32 {
	return gengort.ReadBitcast[int32](unsafe.Add(unsafe.Pointer(unsafe.SliceData(s.Raw[:])), 0))
}
func (s *Anon376_9) SetAsUInt(v int32) {
	gengort.WriteBitcast(unsafe.Add(unsafe.Pointer(unsafe.SliceData(s.Raw[:])), 0), v)
}

//  Gengo init function.
func init() {
	gengort.Validate((*DebuggeeUserInputPacket)(nil), 0xc, 0x4, "CommandLen", 0x0, "IgnoreFinishedSignal", 0x4, "Result", 0x8)
	gengort.Validate((*DebuggeeEventAndActionHeaderForRemotePacket)(nil), 0x4, 0x4, "Length", 0x0)
	gengort.Validate((*DebuggerPausePacketReceived)(nil), 0x4, 0x4, "Result", 0x0)
	gengort.Validate((*DebuggerTriggeredEventDetails)(nil), 0xc, 0x4, "Tag", 0x0, "Context", 0x4, "Stage", 0x8)
	gengort.Validate((*DebuggeeKdPausedPacket)(nil), 0x28, 0x4, "Rip", 0x0, "IsProcessorOn32BitMode", 0x4, "IgnoreDisassembling", 0x8, "PausingReason", 0xc, "CurrentCore", 0x10, "EventTag", 0x14, "EventCallingStage", 0x18, "Rflags", 0x1c, "InstructionBytesOnRip", 0x20, "ReadInstructionLen", 0x24)
	gengort.Validate((*DebuggeeUdPausedPacket)(nil), 0x30, 0x4, "Rip", 0x0, "ProcessDebuggingToken", 0x4, "Is32Bit", 0x8, "PausingReason", 0xc, "ProcessId", 0x10, "ThreadId", 0x14, "Rflags", 0x18, "EventTag", 0x1c, "EventCallingStage", 0x20, "InstructionBytesOnRip", 0x24, "ReadInstructionLen", 0x28, "GuestRegs", 0x2c)
	gengort.Validate((*DebuggeeMessagePacket)(nil), 0x8, 0x4, "OperationCode", 0x0, "Message", 0x4)
	gengort.Validate((*RegisterNotifyBuffer)(nil), 0x8, 0x4, "Type", 0x0, "hEvent", 0x4)
	gengort.Validate((*DirectVmcallParameters)(nil), 0xc, 0x4, "OptionalParam1", 0x0, "OptionalParam2", 0x4, "OptionalParam3", 0x8)
	gengort.Validate((*EptHooksContext)(nil), 0xc, 0x4, "HookingTag", 0x0, "PhysicalAddress", 0x4, "VirtualAddress", 0x8)
	gengort.Validate((*EptHooksAddressDetailsForMemoryMonitor)(nil), 0x1c, 0x4, "StartAddress", 0x0, "EndAddress", 0x4, "SetHookForRead", 0x8, "SetHookForWrite", 0xc, "SetHookForExec", 0x10, "MemoryType", 0x14, "Tag", 0x18)
	gengort.Validate((*EptHooksAddressDetailsForEpthook2)(nil), 0x8, 0x4, "TargetAddress", 0x0, "HookFunction", 0x4)
	gengort.Validate((*EptSingleHookUnhookingDetails)(nil), 0x10, 0x4, "CallerNeedsToRestoreEntryAndInvalidateEpt", 0x0, "RemoveBreakpointInterception", 0x4, "PhysicalAddress", 0x8, "OriginalEntry", 0xc)
	gengort.Validate((*Anon376_9)(nil), 0x4, 0x4)
	gengort.Validate((*Anon378_5)(nil), 0x2c, 0x4, "Type", 0x0, "DescriptorType", 0x4, "DescriptorPrivilegeLevel", 0x8, "Present", 0xc, "Reserved1", 0x10, "AvailableBit", 0x14, "LongMode", 0x18, "DefaultBig", 0x1c, "Granularity", 0x20, "Unusable", 0x24, "Reserved2", 0x28)
	gengort.Validate((*VmxSegmentSelector)(nil), 0x10, 0x4, "Selector", 0x0, "Attributes", 0x4, "Limit", 0x8, "Base", 0xc)
}