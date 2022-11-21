package Headers

import "unsafe"

type ConstantsVar int

var BUFFER_HEADER = SizeOf(10) //todo add struct

var (
	MaxSerialPacketSize   = UsermodeBufferSize + SizeOf(DEBUGGER_REMOTE_PACKET{}) + SERIAL_END_OF_BUFFER_CHARS_COUNT
	LogBufferSize         = MaximumPacketsCapacity * (PacketChunkSize + SizeOf(BUFFER_HEADER))
	LogBufferSizePriority = MaximumPacketsCapacityPriority * (PacketChunkSize + SizeOf(BUFFER_HEADER))
)

const (
	VERSION_MAJOR ConstantsVar = 0
	VERSION_MINOR              = 2
	VERSION_PATCH              = 0

	MaximumPacketsCapacity         = 1000
	MaximumPacketsCapacityPriority = 10
	PacketChunkSize                = 4096 // PAGE_SIZE
	UsermodeBufferSize             = int(unsafe.Sizeof(uint32(0)) + PacketChunkSize + 1)

	DbgPrintLimitation                                    = 512
	DebuggerEventTagStartSeed                             = 0x1000000
	DebuggerThreadDebuggingTagStartSeed                   = 0x1000000
	DebuggerOutputSourceTagStartSeed                      = 0x1
	DebuggerOutputSourceMaximumRemoteSourceForSingleEvent = 0x5
	DebuggerScriptEngineMemcpyMovingBufferSize            = 64
	DEFAULT_PORT                                          = "50000"
	COMMUNICATION_BUFFER_SIZE                             = PacketChunkSize + 0x100
	OPERATION_MANDATORY_DEBUGGEE_BIT                      = (1 << 31)
	OPERATION_LOG_INFO_MESSAGE                            = 0x1
	OPERATION_LOG_WARNING_MESSAGE                         = 0x2
	OPERATION_LOG_ERROR_MESSAGE                           = 0x3
	OPERATION_LOG_NON_IMMEDIATE_MESSAGE                   = 0x4
	OPERATION_LOG_WITH_TAG                                = 0x5
	OPERATION_COMMAND_FROM_DEBUGGER_CLOSE_AND_UNLOAD_VMM  = 0x6 | OPERATION_MANDATORY_DEBUGGEE_BIT
	OPERATION_DEBUGGEE_USER_INPUT                         = 0x7 | OPERATION_MANDATORY_DEBUGGEE_BIT
	OPERATION_DEBUGGEE_REGISTER_EVENT                     = 0x8 | OPERATION_MANDATORY_DEBUGGEE_BIT
	OPERATION_DEBUGGEE_ADD_ACTION_TO_EVENT                = 0x9 | OPERATION_MANDATORY_DEBUGGEE_BIT
	OPERATION_DEBUGGEE_CLEAR_EVENTS                       = 0xa | OPERATION_MANDATORY_DEBUGGEE_BIT
	OPERATION_HYPERVISOR_DRIVER_IS_SUCCESSFULLY_LOADED    = 0xb | OPERATION_MANDATORY_DEBUGGEE_BIT
	OPERATION_HYPERVISOR_DRIVER_END_OF_IRPS               = 0xc | OPERATION_MANDATORY_DEBUGGEE_BIT
	OPERATION_COMMAND_FROM_DEBUGGER_RELOAD_SYMBOL         = 0xd | OPERATION_MANDATORY_DEBUGGEE_BIT
	OPERATION_NOTIFICATION_FROM_USER_DEBUGGER_PAUSE       = 0xe | OPERATION_MANDATORY_DEBUGGEE_BIT
	MAXIMUM_BREAKPOINTS_WITHOUT_CONTINUE                  = 50
	SERIAL_END_OF_BUFFER_CHARS_COUNT                      = 0x4
	SERIAL_END_OF_BUFFER_CHAR_1                           = 0x00
	SERIAL_END_OF_BUFFER_CHAR_2                           = 0x80
	SERIAL_END_OF_BUFFER_CHAR_3                           = 0xEE
	SERIAL_END_OF_BUFFER_CHAR_4                           = 0xFF
	TCP_END_OF_BUFFER_CHARS_COUNT                         = 0x4
	TCP_END_OF_BUFFER_CHAR_1                              = 0x10
	TCP_END_OF_BUFFER_CHAR_2                              = 0x20
	TCP_END_OF_BUFFER_CHAR_3                              = 0x33
	TCP_END_OF_BUFFER_CHAR_4                              = 0x44
	MAXIMUM_CHARACTER_FOR_OS_NAME                         = 256
	MAXIMUM_INSTR_SIZE                                    = 16
	MAXIMUM_CALL_INSTR_SIZE                               = 7
	MAXIMUM_SUPPORTED_SYMBOLS                             = 1000
	MAXIMUM_GUID_AND_AGE_SIZE                             = 60
	INDICATOR_OF_HYPERDBG_PACKET                          = 0x4859504552444247 // HYPERDBG = 0x4859504552444247
	MaximumSearchResults                                  = 0x1000
	X86_FLAGS_CF                                          = (1 << 0)
	X86_FLAGS_PF                                          = (1 << 2)
	X86_FLAGS_AF                                          = (1 << 4)
	X86_FLAGS_ZF                                          = (1 << 6)
	X86_FLAGS_SF                                          = (1 << 7)
	X86_FLAGS_TF                                          = (1 << 8)
	X86_FLAGS_IF                                          = (1 << 9)
	X86_FLAGS_DF                                          = (1 << 10)
	X86_FLAGS_OF                                          = (1 << 11)
	X86_FLAGS_STATUS_MASK                                 = (0xfff)
	X86_FLAGS_IOPL_MASK                                   = (3 << 12)
	X86_FLAGS_IOPL_SHIFT                                  = (12)
	X86_FLAGS_IOPL_SHIFT_2ND_BIT                          = (13)
	X86_FLAGS_NT                                          = (1 << 14)
	X86_FLAGS_RF                                          = (1 << 16)
	X86_FLAGS_VM                                          = (1 << 17)
	X86_FLAGS_AC                                          = (1 << 18)
	X86_FLAGS_VIF                                         = (1 << 19)
	X86_FLAGS_VIP                                         = (1 << 20)
	X86_FLAGS_ID                                          = (1 << 21)
	X86_FLAGS_RESERVED_ONES                               = 0x2
	X86_FLAGS_RESERVED                                    = 0xffc0802a
	X86_FLAGS_RESERVED_BITS                               = 0xffc38028
	X86_FLAGS_FIXED                                       = 0x00000002
	MAX_TEMP_COUNT                                        = 128
	MAX_VAR_COUNT                                         = 512
	MAX_FUNCTION_NAME_LENGTH                              = 32
	DEBUGGER_MODIFY_EVENTS_APPLY_TO_ALL_TAG               = 0xffffffffffffffff
	DISASSEMBLY_MAXIMUM_DISTANCE_FROM_OBJECT_NAME         = 0xffff
	DEBUGGER_READ_AND_WRITE_ON_MSR_APPLY_ALL_CORES        = 0xffffffff
	DEBUGGER_DEBUGGEE_IS_RUNNING_NO_CORE                  = 0xffffffff
	DEBUGGER_EVENT_APPLY_TO_ALL_CORES                     = 0xffffffff
	DEBUGGER_EVENT_APPLY_TO_ALL_PROCESSES                 = 0xffffffff
	DEBUGGER_EVENT_MSR_READ_OR_WRITE_ALL_MSRS             = 0xffffffff
	DEBUGGER_EVENT_EXCEPTIONS_ALL_FIRST_32_ENTRIES        = 0xffffffff
	DEBUGGER_EVENT_SYSCALL_ALL_SYSRET_OR_SYSCALLS         = 0xffffffff
	DEBUGGER_EVENT_ALL_IO_PORTS                           = 0xffffffff
	DEBUGGEE_BP_APPLY_TO_ALL_CORES                        = 0xffffffff
	DEBUGGEE_BP_APPLY_TO_ALL_PROCESSES                    = 0xffffffff
	DEBUGGEE_BP_APPLY_TO_ALL_THREADS                      = 0xffffffff
	DEBUGGEE_SHOW_ALL_REGISTERS                           = 0xffffffff
)

func LOWORD(l uint32) uint16 { return uint16(l) }
func LOBYTE(l uint32) uint8  { return (byte(l)) }
func HIWORD(l uint32) uint16 { return uint16(l >> 16) }
func HIBYTE(l uint32) uint8  { return byte(l >> 24) }
