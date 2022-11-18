package Headers

import "unsafe"

type (
	QWORD uint64
	//UINT64 uint64   //repeated typedef unsigned __int64   UINT64, *PUINT64;
	//PUINT64 *uint64
	DWORD    uint32
	BOOL     bool
	BYTE     byte
	WORD     uint16
	INT      int
	UINT     uint
	PUINT    *uint
	ULONG64  uint64
	PULONG64 *uint64
	DWORD64  uint64
	PDWORD64 *uint64
	CHAR     int8
	WCHAR    rune

	UCHAR byte
	USHOR uint16
	ULONG uint32

	BOOLEAN  bool
	PBOOLEAN *bool

	INT8    int8
	PINT8   uint8
	INT16   int16
	PINT16  *int16
	INT32   int32
	PINT32  *int32
	INT64   int64
	PINT64  *int64
	UINT8   uint8
	PUINT8  *uint8
	UINT16  uint16
	PUINT16 *uint16
	UINT32  uint32
	PUINT32 *uint32
	UINT64  uint64
	PUINT64 *uint64

	PVOID *unsafe.Pointer //todo test
)

const (
	FALSE = 0
	TRUE  = 1

	UPPER_56_BITS                  = 0xffffffffffffff00
	UPPER_48_BITS                  = 0xffffffffffff0000
	UPPER_32_BITS                  = 0xffffffff00000000
	LOWER_32_BITS                  = 0x00000000ffffffff
	LOWER_16_BITS                  = 0x000000000000ffff
	LOWER_8_BITS                   = 0x00000000000000ff
	SECOND_LOWER_8_BITS            = 0x000000000000ff00
	UPPER_48_BITS_AND_LOWER_8_BITS = 0xffffffffffff00ff
)

type (
	GUEST_REGS struct {
		rax uint64
		rcx uint64
		rdx uint64
		rbx uint64
		rsp uint64
		rbp uint64
		rsi uint64
		rdi uint64
		r8  uint64
		r9  uint64
		r10 uint64
		r11 uint64
		r12 uint64
		r13 uint64
		r14 uint64
		r15 uint64
	}
	PGUEST_REGS *GUEST_REGS

	GUEST_EXTRA_REGISTERS struct {
		CS     uint16
		DS     uint16
		FS     uint16
		GS     uint16
		ES     uint16
		SS     uint16
		RFLAGS uint64
		RIP    uint64
	}
	PGUEST_EXTRA_REGISTERS *GUEST_EXTRA_REGISTERS

	DEBUGGER_TRIGGERED_EVENT_DETAILS struct {
		Tag     uint64
		Context PVOID
	}
	PDEBUGGER_TRIGGERED_EVENT_DETAILS *DEBUGGER_TRIGGERED_EVENT_DETAILS

	SCRIPT_ENGINE_VARIABLES_LIST struct {
		TempList            *uint64
		GlobalVariablesList *uint64
		LocalVariablesList  *uint64
	}
	PSCRIPT_ENGINE_VARIABLES_LIST *SCRIPT_ENGINE_VARIABLES_LIST
)
