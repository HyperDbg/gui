package kerneldump

import (
	"encoding/binary"
	"fmt"
)

type ExceptionRecord struct {
	ExceptionCode        uint32
	ExceptionFlags       uint32
	ExceptionRecord      uint64
	ExceptionAddress     uint64
	NumberParameters     uint32
	__unusedAlignment    uint32
	ExceptionInformation [15]uint64
}

func (kd *KernelDump) GetExceptionRecord() (*ExceptionRecord, error) {
	if kd.Header.ContextRecord == 0 {
		return nil, fmt.Errorf("no context record in dump")
	}

	offset := int(kd.Header.ContextRecord)
	if offset+512 > len(kd.Data) {
		return nil, fmt.Errorf("exception record exceeds file size")
	}

	record := &ExceptionRecord{
		ExceptionCode:    binary.LittleEndian.Uint32(kd.Data[offset : offset+4]),
		ExceptionFlags:   binary.LittleEndian.Uint32(kd.Data[offset+4 : offset+8]),
		ExceptionRecord:  binary.LittleEndian.Uint64(kd.Data[offset+8 : offset+16]),
		ExceptionAddress: binary.LittleEndian.Uint64(kd.Data[offset+16 : offset+24]),
		NumberParameters: binary.LittleEndian.Uint32(kd.Data[offset+24 : offset+28]),
	}

	for i := range 15 {
		paramOffset := offset + 32 + i*8
		if paramOffset+8 > len(kd.Data) {
			break
		}
		record.ExceptionInformation[i] = binary.LittleEndian.Uint64(kd.Data[paramOffset : paramOffset+8])
	}

	return record, nil
}

func (kd *KernelDump) GetExceptionCodeName(code uint32) string {
	codes := map[uint32]string{
		0x80000001: "EXCEPTION_ACCESS_VIOLATION",
		0x80000002: "EXCEPTION_ARRAY_BOUNDS_EXCEEDED",
		0x80000003: "EXCEPTION_BREAKPOINT",
		0x80000004: "EXCEPTION_DATATYPE_MISALIGNMENT",
		0x80000005: "EXCEPTION_FLT_DENORMAL_OPERAND",
		0x80000006: "EXCEPTION_FLT_DIVIDE_BY_ZERO",
		0x80000007: "EXCEPTION_FLT_INEXACT_RESULT",
		0x80000008: "EXCEPTION_FLT_INVALID_OPERATION",
		0x80000009: "EXCEPTION_FLT_OVERFLOW",
		0x8000000A: "EXCEPTION_FLT_STACK_CHECK",
		0x8000000B: "EXCEPTION_FLT_UNDERFLOW",
		0x8000000C: "EXCEPTION_INTEGER_DIVIDE_BY_ZERO",
		0x8000000D: "EXCEPTION_INTEGER_OVERFLOW",
		0x8000000E: "EXCEPTION_PRIV_INSTRUCTION",
		0x8000000F: "EXCEPTION_IN_PAGE_ERROR",
		0x80000010: "EXCEPTION_ILLEGAL_INSTRUCTION",
		0x80000011: "EXCEPTION_NONCONTINUABLE_EXCEPTION",
		0x80000012: "EXCEPTION_STACK_OVERFLOW",
		0x80000013: "EXCEPTION_INVALID_DISPOSITION",
		0x80000014: "EXCEPTION_GUARD_PAGE",
		0x80000015: "EXCEPTION_INVALID_HANDLE",
		0x8000001D: "EXCEPTION_POSSIBLE_DEADLOCK",
		0x8000001E: "EXCEPTION_INVALID_LOCK_SEQUENCE",
		0x80000026: "EXCEPTION_INVALID_OPCODE",
		0xC0000005: "STATUS_ACCESS_VIOLATION",
		0xC0000006: "STATUS_IN_PAGE_ERROR",
		0xC0000008: "STATUS_INVALID_HANDLE",
		0xC0000025: "STATUS_NO_MEMORY",
		0xC00000FD: "STATUS_STACK_BUFFER_OVERRUN",
	}

	if name, ok := codes[code]; ok {
		return name
	}
	return fmt.Sprintf("UNKNOWN_EXCEPTION_0x%08X", code)
}
