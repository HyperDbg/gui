package sdk

import (
	os "os"
	unsafe "unsafe"
)

type DWORD = uint64
type BOOL = int32
type BYTE = uint8
type WORD = uint16
type FLOAT = float32
type PFLOAT = *float32
type INT = int32
type UINT = uint32
type PUINT = *uint32
type PBOOL = *int32
type LPBOOL = *int32
type PBYTE = *uint8
type LPBYTE = *uint8
type PINT = *int32
type LPINT = *int32
type PWORD = *uint16
type LPWORD = *uint16
type LPLONG = *int64
type PDWORD = *uint64
type LPDWORD = *uint64
type LPVOID = unsafe.Pointer
type PVOID = unsafe.Pointer
type LPCVOID = unsafe.Pointer
type ULONG = uint64
type PULONG = *uint64
type USHORT = uint16
type PUSHORT = *uint16
type UCHAR = uint8
type PUCHAR = *uint8
type CHAR = int8
type SHORT = int16
type LONG = int64
type QWORD = uint64
type UINT64 = uint64
type PUINT64 = *uint64
type ULONG64 = uint64
type PULONG64 = *uint64
type DWORD64 = uint64
type PDWORD64 = *uint64
type BOOLEAN = uint8
type PBOOLEAN = *uint8
type INT8 = int8
type PINT8 = *int8
type INT16 = int16
type PINT16 = *int16
type INT32 = int32
type PINT32 = *int32
type INT64 = int64
type PINT64 = *int64
type UINT8 = uint8
type PUINT8 = *uint8
type UINT16 = uint16
type PUINT16 = *uint16
type UINT32 = uint32
type PUINT32 = *uint32
type wchar_t = *int16
type WCHAR = *int16
type struct__LIST_ENTRY struct {
	Flink *struct__LIST_ENTRY
	Blink *struct__LIST_ENTRY
}
type LIST_ENTRY = struct__LIST_ENTRY
type PLIST_ENTRY = *struct__LIST_ENTRY
type struct_GUEST_REGS struct {
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
type GUEST_REGS = struct_GUEST_REGS
type PGUEST_REGS = *struct_GUEST_REGS
type struct_GUEST_EXTRA_REGISTERS struct {
	CS     uint16
	DS     uint16
	FS     uint16
	GS     uint16
	ES     uint16
	SS     uint16
	RFLAGS uint64
	RIP    uint64
}
type GUEST_EXTRA_REGISTERS = struct_GUEST_EXTRA_REGISTERS
type PGUEST_EXTRA_REGISTERS = *struct_GUEST_EXTRA_REGISTERS
type struct__DEBUGGER_TRIGGERED_EVENT_DETAILS struct {
	Tag     uint64
	Context unsafe.Pointer
}
type DEBUGGER_TRIGGERED_EVENT_DETAILS = struct__DEBUGGER_TRIGGERED_EVENT_DETAILS
type PDEBUGGER_TRIGGERED_EVENT_DETAILS = *struct__DEBUGGER_TRIGGERED_EVENT_DETAILS
type struct__SCRIPT_ENGINE_VARIABLES_LIST struct {
	TempList            *uint64
	GlobalVariablesList *uint64
	LocalVariablesList  *uint64
}
type SCRIPT_ENGINE_VARIABLES_LIST = struct__SCRIPT_ENGINE_VARIABLES_LIST
type PSCRIPT_ENGINE_VARIABLES_LIST = *struct__SCRIPT_ENGINE_VARIABLES_LIST

var BuildDateTime [20]uint8 = [20]uint8{uint8(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(7))))), uint8(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(8))))), uint8(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(9))))), uint8(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(10))))), uint8('-'), uint8(func() int32 {
	if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(0))))) == 'O' || int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(0))))) == 'N' || int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(0))))) == 'D' {
		return '1'
	} else {
		return '0'
	}
}()), uint8(func() int32 {
	if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(0))))) == 'J' && int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(1))))) == 'a' && int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(2))))) == 'n' {
		return '1'
	} else {
		return func() int32 {
			if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(0))))) == 'F' {
				return '2'
			} else {
				return func() int32 {
					if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(0))))) == 'M' && int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(1))))) == 'a' && int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(2))))) == 'r' {
						return '3'
					} else {
						return func() int32 {
							if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(0))))) == 'A' && int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(1))))) == 'p' {
								return '4'
							} else {
								return func() int32 {
									if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(0))))) == 'M' && int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(1))))) == 'a' && int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(2))))) == 'y' {
										return '5'
									} else {
										return func() int32 {
											if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(0))))) == 'J' && int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(1))))) == 'u' && int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(2))))) == 'n' {
												return '6'
											} else {
												return func() int32 {
													if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(0))))) == 'J' && int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(1))))) == 'u' && int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(2))))) == 'l' {
														return '7'
													} else {
														return func() int32 {
															if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(0))))) == 'A' && int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(1))))) == 'u' {
																return '8'
															} else {
																return func() int32 {
																	if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(0))))) == 'S' {
																		return '9'
																	} else {
																		return func() int32 {
																			if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(0))))) == 'O' {
																				return '0'
																			} else {
																				return func() int32 {
																					if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(0))))) == 'N' {
																						return '1'
																					} else {
																						return func() int32 {
																							if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(0))))) == 'D' {
																								return '2'
																							} else {
																								return '?'
																							}
																						}()
																					}
																				}()
																			}
																		}()
																	}
																}()
															}
														}()
													}
												}()
											}
										}()
									}
								}()
							}
						}()
					}
				}()
			}
		}()
	}
}()), uint8('-'), uint8(func() int32 {
	if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(4))))) >= '0' {
		return int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(4)))))
	} else {
		return '0'
	}
}()), uint8(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(5))))), uint8(' '), uint8(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[9]int8{'1', '1', ':', '4', '7', ':', '1', '0', '\x00'})))) + uintptr(int32(0))))), uint8(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[9]int8{'1', '1', ':', '4', '7', ':', '1', '0', '\x00'})))) + uintptr(int32(1))))), uint8(':'), uint8(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[9]int8{'1', '1', ':', '4', '7', ':', '1', '0', '\x00'})))) + uintptr(int32(3))))), uint8(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[9]int8{'1', '1', ':', '4', '7', ':', '1', '0', '\x00'})))) + uintptr(int32(4))))), uint8(':'), uint8(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[9]int8{'1', '1', ':', '4', '7', ':', '1', '0', '\x00'})))) + uintptr(int32(6))))), uint8(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[9]int8{'1', '1', ':', '4', '7', ':', '1', '0', '\x00'})))) + uintptr(int32(7))))), uint8('\x00')}
var CompleteVersion [7]uint8 = [7]uint8{uint8('v'), uint8(48), uint8('.'), uint8(50), uint8('.'), uint8(48), uint8('\x00')}
var BuildVersion [9]uint8 = [9]uint8{uint8(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(7))))), uint8(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(8))))), uint8(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(9))))), uint8(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(10))))), uint8(func() int32 {
	if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(0))))) == 'O' || int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(0))))) == 'N' || int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(0))))) == 'D' {
		return '1'
	} else {
		return '0'
	}
}()), uint8(func() int32 {
	if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(0))))) == 'J' && int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(1))))) == 'a' && int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(2))))) == 'n' {
		return '1'
	} else {
		return func() int32 {
			if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(0))))) == 'F' {
				return '2'
			} else {
				return func() int32 {
					if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(0))))) == 'M' && int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(1))))) == 'a' && int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(2))))) == 'r' {
						return '3'
					} else {
						return func() int32 {
							if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(0))))) == 'A' && int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(1))))) == 'p' {
								return '4'
							} else {
								return func() int32 {
									if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(0))))) == 'M' && int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(1))))) == 'a' && int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(2))))) == 'y' {
										return '5'
									} else {
										return func() int32 {
											if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(0))))) == 'J' && int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(1))))) == 'u' && int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(2))))) == 'n' {
												return '6'
											} else {
												return func() int32 {
													if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(0))))) == 'J' && int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(1))))) == 'u' && int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(2))))) == 'l' {
														return '7'
													} else {
														return func() int32 {
															if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(0))))) == 'A' && int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(1))))) == 'u' {
																return '8'
															} else {
																return func() int32 {
																	if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(0))))) == 'S' {
																		return '9'
																	} else {
																		return func() int32 {
																			if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(0))))) == 'O' {
																				return '0'
																			} else {
																				return func() int32 {
																					if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(0))))) == 'N' {
																						return '1'
																					} else {
																						return func() int32 {
																							if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(0))))) == 'D' {
																								return '2'
																							} else {
																								return '?'
																							}
																						}()
																					}
																				}()
																			}
																		}()
																	}
																}()
															}
														}()
													}
												}()
											}
										}()
									}
								}()
							}
						}()
					}
				}()
			}
		}()
	}
}()), uint8(func() int32 {
	if int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(4))))) >= '0' {
		return int32(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(4)))))
	} else {
		return '0'
	}
}()), uint8(*(*int8)(unsafe.Pointer(uintptr(unsafe.Pointer((*int8)(unsafe.Pointer(&[12]int8{'J', 'u', 'l', ' ', '2', '1', ' ', '2', '0', '2', '2', '\x00'})))) + uintptr(int32(5))))), uint8('\x00')}

const (
	DEBUGGEE_PAUSING_REASON_NOT_PAUSED                            int32 = int32(0)
	DEBUGGEE_PAUSING_REASON_PAUSE_WITHOUT_DISASM                  int32 = 1
	DEBUGGEE_PAUSING_REASON_REQUEST_FROM_DEBUGGER                 int32 = 2
	DEBUGGEE_PAUSING_REASON_DEBUGGEE_STEPPED                      int32 = 3
	DEBUGGEE_PAUSING_REASON_DEBUGGEE_SOFTWARE_BREAKPOINT_HIT      int32 = 4
	DEBUGGEE_PAUSING_REASON_DEBUGGEE_HARDWARE_DEBUG_REGISTER_HIT  int32 = 5
	DEBUGGEE_PAUSING_REASON_DEBUGGEE_CORE_SWITCHED                int32 = 6
	DEBUGGEE_PAUSING_REASON_DEBUGGEE_PROCESS_SWITCHED             int32 = 7
	DEBUGGEE_PAUSING_REASON_DEBUGGEE_THREAD_SWITCHED              int32 = 8
	DEBUGGEE_PAUSING_REASON_DEBUGGEE_COMMAND_EXECUTION_FINISHED   int32 = 9
	DEBUGGEE_PAUSING_REASON_DEBUGGEE_EVENT_TRIGGERED              int32 = 10
	DEBUGGEE_PAUSING_REASON_DEBUGGEE_ENTRY_POINT_REACHED          int32 = 11
	DEBUGGEE_PAUSING_REASON_DEBUGGEE_GENERAL_DEBUG_BREAK          int32 = 12
	DEBUGGEE_PAUSING_REASON_DEBUGGEE_GENERAL_THREAD_INTERCEPTED   int32 = 13
	DEBUGGEE_PAUSING_REASON_HARDWARE_BASED_DEBUGGEE_GENERAL_BREAK int32 = 14
)

type DEBUGGEE_PAUSING_REASON = int32

const (
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_USER_MODE_PAUSE                            int32 = int32(1)
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_USER_MODE_DO_NOT_READ_ANY_PACKET           int32 = 2
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_MODE_STEP                         int32 = 3
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_MODE_CONTINUE                     int32 = 4
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_MODE_CLOSE_AND_UNLOAD_DEBUGGEE    int32 = 5
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_MODE_CHANGE_CORE                  int32 = 6
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_MODE_FLUSH_BUFFERS                int32 = 7
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_MODE_CALLSTACK                    int32 = 8
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_MODE_TEST_QUERY                   int32 = 9
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_MODE_CHANGE_PROCESS               int32 = 10
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_MODE_CHANGE_THREAD                int32 = 11
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_RUN_SCRIPT                        int32 = 12
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_USER_INPUT_BUFFER                 int32 = 13
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_SEARCH_QUERY                      int32 = 14
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_REGISTER_EVENT                    int32 = 15
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_ADD_ACTION_TO_EVENT               int32 = 16
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_QUERY_AND_MODIFY_EVENT            int32 = 17
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_READ_REGISTERS                    int32 = 18
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_READ_MEMORY                       int32 = 19
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_EDIT_MEMORY                       int32 = 20
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_BP                                int32 = 21
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_LIST_OR_MODIFY_BREAKPOINTS        int32 = 22
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_SYMBOL_RELOAD                     int32 = 23
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_QUERY_PA2VA_AND_VA2PA             int32 = 24
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_ON_VMX_ROOT_SYMBOL_QUERY_PTE                  int32 = 25
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_NO_ACTION                                     int32 = int32(0)
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_STARTED                              int32 = 1
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_LOGGING_MECHANISM                    int32 = 2
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_PAUSED_AND_CURRENT_INSTRUCTION       int32 = 3
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_CHANGING_CORE              int32 = 4
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_CHANGING_PROCESS           int32 = 5
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_CHANGING_THREAD            int32 = 6
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_RUNNING_SCRIPT             int32 = 7
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_FORMATS                    int32 = 8
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_FLUSH                      int32 = 9
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_CALLSTACK                  int32 = 10
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_TEST_QUERY                    int32 = 11
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_REGISTERING_EVENT          int32 = 12
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_ADDING_ACTION_TO_EVENT     int32 = 13
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_QUERY_AND_MODIFY_EVENT     int32 = 14
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_READING_REGISTERS          int32 = 15
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_READING_MEMORY             int32 = 16
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_EDITING_MEMORY             int32 = 17
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_BP                         int32 = 18
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_LIST_OR_MODIFY_BREAKPOINTS int32 = 19
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_UPDATE_SYMBOL_INFO                   int32 = 20
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RELOAD_SYMBOL_FINISHED               int32 = 21
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RELOAD_SEARCH_QUERY                  int32 = 22
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_PTE                        int32 = 23
	DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION_DEBUGGEE_RESULT_OF_VA2PA_AND_PA2VA            int32 = 24
)

type DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION = int32

const (
	DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGER_TO_DEBUGGEE_EXECUTE_ON_VMX_ROOT  int32 = int32(1)
	DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGER_TO_DEBUGGEE_EXECUTE_ON_USER_MODE int32 = 2
	DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGEE_TO_DEBUGGER                      int32 = 3
	DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGER_TO_DEBUGGEE_HARDWARE_LEVEL       int32 = int32(1)
	DEBUGGER_REMOTE_PACKET_TYPE_DEBUGGEE_TO_DEBUGGER_HARDWARE_LEVEL       int32 = 2
)

type DEBUGGER_REMOTE_PACKET_TYPE = int32
type struct__DEBUGGER_REMOTE_PACKET struct {
	Checksum                   uint8
	Indicator                  uint64
	TypeOfThePacket            int32
	RequestedActionOfThePacket int32
}
type DEBUGGER_REMOTE_PACKET = struct__DEBUGGER_REMOTE_PACKET
type PDEBUGGER_REMOTE_PACKET = *struct__DEBUGGER_REMOTE_PACKET
type Callback = func(*int8) int32
type struct__DEBUGGEE_USER_INPUT_PACKET struct {
	CommandLen           uint32
	IgnoreFinishedSignal uint8
	Result               uint32
}
type DEBUGGEE_USER_INPUT_PACKET = struct__DEBUGGEE_USER_INPUT_PACKET
type PDEBUGGEE_USER_INPUT_PACKET = *struct__DEBUGGEE_USER_INPUT_PACKET
type struct__DEBUGGEE_EVENT_AND_ACTION_HEADER_FOR_REMOTE_PACKET struct {
	Length uint32
}
type DEBUGGEE_EVENT_AND_ACTION_HEADER_FOR_REMOTE_PACKET = struct__DEBUGGEE_EVENT_AND_ACTION_HEADER_FOR_REMOTE_PACKET
type PDEBUGGEE_EVENT_AND_ACTION_HEADER_FOR_REMOTE_PACKET = *struct__DEBUGGEE_EVENT_AND_ACTION_HEADER_FOR_REMOTE_PACKET
type struct__DEBUGGER_PAUSE_PACKET_RECEIVED struct {
	Result uint32
}
type DEBUGGER_PAUSE_PACKET_RECEIVED = struct__DEBUGGER_PAUSE_PACKET_RECEIVED
type PDEBUGGER_PAUSE_PACKET_RECEIVED = *struct__DEBUGGER_PAUSE_PACKET_RECEIVED
type struct__DEBUGGEE_KD_PAUSED_PACKET struct {
	Rip                   uint64
	Is32BitAddress        uint8
	PausingReason         int32
	CurrentCore           uint64
	EventTag              uint64
	Rflags                uint64
	InstructionBytesOnRip [16]uint8
	ReadInstructionLen    uint16
}
type DEBUGGEE_KD_PAUSED_PACKET = struct__DEBUGGEE_KD_PAUSED_PACKET
type PDEBUGGEE_KD_PAUSED_PACKET = *struct__DEBUGGEE_KD_PAUSED_PACKET
type struct__DEBUGGEE_UD_PAUSED_PACKET struct {
	Rip                   uint64
	ProcessDebuggingToken uint64
	Is32Bit               uint8
	PausingReason         int32
	ProcessId             uint32
	ThreadId              uint32
	EventTag              uint64
	Rflags                uint64
	InstructionBytesOnRip [16]uint8
	ReadInstructionLen    uint16
	GuestRegs             struct_GUEST_REGS
}
type DEBUGGEE_UD_PAUSED_PACKET = struct__DEBUGGEE_UD_PAUSED_PACKET
type PDEBUGGEE_UD_PAUSED_PACKET = *struct__DEBUGGEE_UD_PAUSED_PACKET
type struct__DEBUGGEE_MESSAGE_PACKET struct {
	OperationCode uint32
	Message       [4096]int8
}
type DEBUGGEE_MESSAGE_PACKET = struct__DEBUGGEE_MESSAGE_PACKET
type PDEBUGGEE_MESSAGE_PACKET = *struct__DEBUGGEE_MESSAGE_PACKET

const (
	HIDDEN_HOOK_READ_AND_WRITE   int32 = 0
	HIDDEN_HOOK_READ             int32 = 1
	HIDDEN_HOOK_WRITE            int32 = 2
	HIDDEN_HOOK_EXEC_DETOURS     int32 = 3
	HIDDEN_HOOK_EXEC_CC          int32 = 4
	SYSCALL_HOOK_EFER_SYSCALL    int32 = 5
	SYSCALL_HOOK_EFER_SYSRET     int32 = 6
	CPUID_INSTRUCTION_EXECUTION  int32 = 7
	RDMSR_INSTRUCTION_EXECUTION  int32 = 8
	WRMSR_INSTRUCTION_EXECUTION  int32 = 9
	IN_INSTRUCTION_EXECUTION     int32 = 10
	OUT_INSTRUCTION_EXECUTION    int32 = 11
	EXCEPTION_OCCURRED           int32 = 12
	EXTERNAL_INTERRUPT_OCCURRED  int32 = 13
	DEBUG_REGISTERS_ACCESSED     int32 = 14
	TSC_INSTRUCTION_EXECUTION    int32 = 15
	PMC_INSTRUCTION_EXECUTION    int32 = 16
	VMCALL_INSTRUCTION_EXECUTION int32 = 17
	CONTROL_REGISTER_MODIFIED    int32 = 18
)

type DEBUGGER_EVENT_TYPE_ENUM = int32

const (
	BREAK_TO_DEBUGGER int32 = 0
	RUN_SCRIPT        int32 = 1
	RUN_CUSTOM_CODE   int32 = 2
)

type DEBUGGER_EVENT_ACTION_TYPE_ENUM = int32

const (
	DEBUGGER_EVENT_SYSCALL_SYSRET_SAFE_ACCESS_MEMORY int32 = int32(0)
	DEBUGGER_EVENT_SYSCALL_SYSRET_HANDLE_ALL_UD      int32 = int32(1)
)

type DEBUGGER_EVENT_SYSCALL_SYSRET_TYPE = int32

const (
	DEBUGGER_MODIFY_EVENTS_QUERY_STATE int32 = 0
	DEBUGGER_MODIFY_EVENTS_ENABLE      int32 = 1
	DEBUGGER_MODIFY_EVENTS_DISABLE     int32 = 2
	DEBUGGER_MODIFY_EVENTS_CLEAR       int32 = 3
)

type DEBUGGER_MODIFY_EVENTS_TYPE = int32
type struct__DEBUGGER_MODIFY_EVENTS struct {
	Tag          uint64
	KernelStatus uint64
	TypeOfAction int32
	IsEnabled    uint8
}
type DEBUGGER_MODIFY_EVENTS = struct__DEBUGGER_MODIFY_EVENTS
type PDEBUGGER_MODIFY_EVENTS = *struct__DEBUGGER_MODIFY_EVENTS
type struct__DEBUGGER_READ_PAGE_TABLE_ENTRIES_DETAILS struct {
	VirtualAddress      uint64
	ProcessId           uint32
	Pml4eVirtualAddress uint64
	Pml4eValue          uint64
	PdpteVirtualAddress uint64
	PdpteValue          uint64
	PdeVirtualAddress   uint64
	PdeValue            uint64
	PteVirtualAddress   uint64
	PteValue            uint64
	KernelStatus        uint32
}
type DEBUGGER_READ_PAGE_TABLE_ENTRIES_DETAILS = struct__DEBUGGER_READ_PAGE_TABLE_ENTRIES_DETAILS
type PDEBUGGER_READ_PAGE_TABLE_ENTRIES_DETAILS = *struct__DEBUGGER_READ_PAGE_TABLE_ENTRIES_DETAILS
type struct__DEBUGGER_VA2PA_AND_PA2VA_COMMANDS struct {
	VirtualAddress     uint64
	PhysicalAddress    uint64
	ProcessId          uint32
	IsVirtual2Physical uint8
	KernelStatus       uint32
}
type DEBUGGER_VA2PA_AND_PA2VA_COMMANDS = struct__DEBUGGER_VA2PA_AND_PA2VA_COMMANDS
type PDEBUGGER_VA2PA_AND_PA2VA_COMMANDS = *struct__DEBUGGER_VA2PA_AND_PA2VA_COMMANDS
type struct__DEBUGGER_DT_COMMAND_OPTIONS struct {
	TypeName             *int8
	SizeOfTypeName       uint64
	Address              uint64
	IsStruct             uint8
	BufferAddress        unsafe.Pointer
	TargetPid            uint32
	AdditionalParameters *int8
}
type DEBUGGER_DT_COMMAND_OPTIONS = struct__DEBUGGER_DT_COMMAND_OPTIONS
type PDEBUGGER_DT_COMMAND_OPTIONS = *struct__DEBUGGER_DT_COMMAND_OPTIONS

const (
	DEBUGGER_PREALLOC_COMMAND_TYPE_MONITOR             int32 = 0
	DEBUGGER_PREALLOC_COMMAND_TYPE_THREAD_INTERCEPTION int32 = 1
)

type DEBUGGER_PREALLOC_COMMAND_TYPE = int32
type struct__DEBUGGER_PREALLOC_COMMAND struct {
	Type         int32
	Count        uint64
	KernelStatus uint32
}
type DEBUGGER_PREALLOC_COMMAND = struct__DEBUGGER_PREALLOC_COMMAND
type PDEBUGGER_PREALLOC_COMMAND = *struct__DEBUGGER_PREALLOC_COMMAND

const (
	READ_FROM_KERNEL   int32 = 0
	READ_FROM_VMX_ROOT int32 = 1
)

type DEBUGGER_READ_READING_TYPE = int32

const (
	DEBUGGER_READ_PHYSICAL_ADDRESS int32 = 0
	DEBUGGER_READ_VIRTUAL_ADDRESS  int32 = 1
)

type DEBUGGER_READ_MEMORY_TYPE = int32

const (
	DEBUGGER_SHOW_COMMAND_DT            int32 = int32(1)
	DEBUGGER_SHOW_COMMAND_DISASSEMBLE64 int32 = 2
	DEBUGGER_SHOW_COMMAND_DISASSEMBLE32 int32 = 3
	DEBUGGER_SHOW_COMMAND_DB            int32 = 4
	DEBUGGER_SHOW_COMMAND_DC            int32 = 5
	DEBUGGER_SHOW_COMMAND_DQ            int32 = 6
	DEBUGGER_SHOW_COMMAND_DD            int32 = 7
)

type DEBUGGER_SHOW_MEMORY_STYLE = int32
type struct__DEBUGGER_READ_MEMORY struct {
	Pid          uint32
	Address      uint64
	Size         uint32
	MemoryType   int32
	ReadingType  int32
	DtDetails    *struct__DEBUGGER_DT_COMMAND_OPTIONS
	Style        int32
	ReturnLength uint32
	KernelStatus uint32
}
type DEBUGGER_READ_MEMORY = struct__DEBUGGER_READ_MEMORY
type PDEBUGGER_READ_MEMORY = *struct__DEBUGGER_READ_MEMORY
type struct__DEBUGGER_FLUSH_LOGGING_BUFFERS struct {
	KernelStatus                               uint32
	CountOfMessagesThatSetAsReadFromVmxRoot    uint32
	CountOfMessagesThatSetAsReadFromVmxNonRoot uint32
}
type DEBUGGER_FLUSH_LOGGING_BUFFERS = struct__DEBUGGER_FLUSH_LOGGING_BUFFERS
type PDEBUGGER_FLUSH_LOGGING_BUFFERS = *struct__DEBUGGER_FLUSH_LOGGING_BUFFERS
type struct__DEBUGGER_DEBUGGER_TEST_QUERY_BUFFER struct {
	RequestIndex uint32
	KernelStatus uint32
}
type DEBUGGER_DEBUGGER_TEST_QUERY_BUFFER = struct__DEBUGGER_DEBUGGER_TEST_QUERY_BUFFER
type PDEBUGGER_DEBUGGER_TEST_QUERY_BUFFER = *struct__DEBUGGER_DEBUGGER_TEST_QUERY_BUFFER
type struct__DEBUGGER_PERFORM_KERNEL_TESTS struct {
	KernelStatus uint32
}
type DEBUGGER_PERFORM_KERNEL_TESTS = struct__DEBUGGER_PERFORM_KERNEL_TESTS
type PDEBUGGER_PERFORM_KERNEL_TESTS = *struct__DEBUGGER_PERFORM_KERNEL_TESTS
type struct__DEBUGGER_SEND_COMMAND_EXECUTION_FINISHED_SIGNAL struct {
	KernelStatus uint32
}
type DEBUGGER_SEND_COMMAND_EXECUTION_FINISHED_SIGNAL = struct__DEBUGGER_SEND_COMMAND_EXECUTION_FINISHED_SIGNAL
type PDEBUGGER_SEND_COMMAND_EXECUTION_FINISHED_SIGNAL = *struct__DEBUGGER_SEND_COMMAND_EXECUTION_FINISHED_SIGNAL
type struct__DEBUGGEE_KERNEL_AND_USER_TEST_INFORMATION struct {
	Value uint64
	Tag   [32]int8
}
type DEBUGGEE_KERNEL_AND_USER_TEST_INFORMATION = struct__DEBUGGEE_KERNEL_AND_USER_TEST_INFORMATION
type PDEBUGGEE_KERNEL_AND_USER_TEST_INFORMATION = *struct__DEBUGGEE_KERNEL_AND_USER_TEST_INFORMATION
type struct__DEBUGGEE_SEND_GENERAL_PACKET_FROM_DEBUGGEE_TO_DEBUGGER struct {
	RequestedAction       int32
	LengthOfBuffer        uint32
	PauseDebuggeeWhenSent uint8
	KernelResult          uint32
}
type DEBUGGEE_SEND_GENERAL_PACKET_FROM_DEBUGGEE_TO_DEBUGGER = struct__DEBUGGEE_SEND_GENERAL_PACKET_FROM_DEBUGGEE_TO_DEBUGGER
type PDEBUGGEE_SEND_GENERAL_PACKET_FROM_DEBUGGEE_TO_DEBUGGER = *struct__DEBUGGEE_SEND_GENERAL_PACKET_FROM_DEBUGGEE_TO_DEBUGGER
type struct__DEBUGGER_SEND_USERMODE_MESSAGES_TO_DEBUGGER struct {
	KernelStatus uint32
	Length       uint32
}
type DEBUGGER_SEND_USERMODE_MESSAGES_TO_DEBUGGER = struct__DEBUGGER_SEND_USERMODE_MESSAGES_TO_DEBUGGER
type PDEBUGGER_SEND_USERMODE_MESSAGES_TO_DEBUGGER = *struct__DEBUGGER_SEND_USERMODE_MESSAGES_TO_DEBUGGER

const (
	DEBUGGER_MSR_READ  int32 = 0
	DEBUGGER_MSR_WRITE int32 = 1
)

type DEBUGGER_MSR_ACTION_TYPE = int32
type struct__DEBUGGER_READ_AND_WRITE_ON_MSR struct {
	Msr        uint64
	CoreNumber uint32
	ActionType int32
	Value      uint64
}
type DEBUGGER_READ_AND_WRITE_ON_MSR = struct__DEBUGGER_READ_AND_WRITE_ON_MSR
type PDEBUGGER_READ_AND_WRITE_ON_MSR = *struct__DEBUGGER_READ_AND_WRITE_ON_MSR

const (
	EDIT_PHYSICAL_MEMORY int32 = 0
	EDIT_VIRTUAL_MEMORY  int32 = 1
)

type DEBUGGER_EDIT_MEMORY_TYPE = int32

const (
	EDIT_BYTE  int32 = 0
	EDIT_DWORD int32 = 1
	EDIT_QWORD int32 = 2
)

type DEBUGGER_EDIT_MEMORY_BYTE_SIZE = int32
type struct__DEBUGGER_EDIT_MEMORY struct {
	Result             uint32
	Address            uint64
	ProcessId          uint32
	MemoryType         int32
	ByteSize           int32
	CountOf64Chunks    uint32
	FinalStructureSize uint32
	KernelStatus       uint32
}
type DEBUGGER_EDIT_MEMORY = struct__DEBUGGER_EDIT_MEMORY
type PDEBUGGER_EDIT_MEMORY = *struct__DEBUGGER_EDIT_MEMORY

const (
	SEARCH_PHYSICAL_MEMORY              int32 = 0
	SEARCH_VIRTUAL_MEMORY               int32 = 1
	SEARCH_PHYSICAL_FROM_VIRTUAL_MEMORY int32 = 2
)

type DEBUGGER_SEARCH_MEMORY_TYPE = int32

const (
	SEARCH_BYTE  int32 = 0
	SEARCH_DWORD int32 = 1
	SEARCH_QWORD int32 = 2
)

type DEBUGGER_SEARCH_MEMORY_BYTE_SIZE = int32
type struct__DEBUGGER_SEARCH_MEMORY struct {
	Address            uint64
	Length             uint64
	ProcessId          uint32
	MemoryType         int32
	ByteSize           int32
	CountOf64Chunks    uint32
	FinalStructureSize uint32
}
type DEBUGGER_SEARCH_MEMORY = struct__DEBUGGER_SEARCH_MEMORY
type PDEBUGGER_SEARCH_MEMORY = *struct__DEBUGGER_SEARCH_MEMORY
type struct__DEBUGGER_HIDE_AND_TRANSPARENT_DEBUGGER_MODE struct {
	IsHide                               uint8
	CpuidAverage                         uint64
	CpuidStandardDeviation               uint64
	CpuidMedian                          uint64
	RdtscAverage                         uint64
	RdtscStandardDeviation               uint64
	RdtscMedian                          uint64
	TrueIfProcessIdAndFalseIfProcessName uint8
	ProcId                               uint32
	LengthOfProcessName                  uint32
	KernelStatus                         uint64
}
type DEBUGGER_HIDE_AND_TRANSPARENT_DEBUGGER_MODE = struct__DEBUGGER_HIDE_AND_TRANSPARENT_DEBUGGER_MODE
type PDEBUGGER_HIDE_AND_TRANSPARENT_DEBUGGER_MODE = *struct__DEBUGGER_HIDE_AND_TRANSPARENT_DEBUGGER_MODE
type struct__DEBUGGER_PREPARE_DEBUGGEE struct {
	PortAddress         uint32
	Baudrate            uint32
	NtoskrnlBaseAddress uint64
	Result              uint32
	OsName              [256]int8
}
type DEBUGGER_PREPARE_DEBUGGEE = struct__DEBUGGER_PREPARE_DEBUGGEE
type PDEBUGGER_PREPARE_DEBUGGEE = *struct__DEBUGGER_PREPARE_DEBUGGEE
type struct__DEBUGGEE_CHANGE_CORE_PACKET struct {
	NewCore uint32
	Result  uint32
}
type DEBUGGEE_CHANGE_CORE_PACKET = struct__DEBUGGEE_CHANGE_CORE_PACKET
type PDEBUGGEE_CHANGE_CORE_PACKET = *struct__DEBUGGEE_CHANGE_CORE_PACKET

const (
	DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS_ACTION_ATTACH                                  int32 = 0
	DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS_ACTION_DETACH                                  int32 = 1
	DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS_ACTION_REMOVE_HOOKS                            int32 = 2
	DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS_ACTION_KILL_PROCESS                            int32 = 3
	DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS_ACTION_PAUSE_PROCESS                           int32 = 4
	DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS_ACTION_SWITCH_BY_PROCESS_OR_THREAD             int32 = 5
	DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS_ACTION_QUERY_COUNT_OF_ACTIVE_DEBUGGING_THREADS int32 = 6
)

type DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS_ACTION_TYPE = int32
type struct__DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS struct {
	IsStartingNewProcess                      uint8
	ProcessId                                 uint32
	ThreadId                                  uint32
	Is32Bit                                   uint8
	IsPaused                                  uint8
	Action                                    int32
	CountOfActiveDebuggingThreadsAndProcesses uint32
	Token                                     uint64
	Result                                    uint64
}
type DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS = struct__DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS
type PDEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS = *struct__DEBUGGER_ATTACH_DETACH_USER_MODE_PROCESS

const (
	DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_QUERY_PROCESS_COUNT   int32 = int32(1)
	DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_QUERY_THREAD_COUNT    int32 = int32(2)
	DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_QUERY_PROCESS_LIST    int32 = int32(3)
	DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_QUERY_THREAD_LIST     int32 = int32(4)
	DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_QUERY_CURRENT_PROCESS int32 = int32(5)
	DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_QUERY_CURRENT_THREAD  int32 = int32(6)
)

type DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_TYPES = int32

const (
	DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_ACTION_SHOW_INSTANTLY     int32 = int32(1)
	DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_ACTION_QUERY_COUNT        int32 = int32(2)
	DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_ACTION_QUERY_SAVE_DETAILS int32 = int32(3)
)

type DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS_ACTIONS = int32
type struct__DEBUGGEE_PROCESS_LIST_NEEDED_DETAILS struct {
	PsActiveProcessHead      uint64
	ImageFileNameOffset      uint64
	UniquePidOffset          uint64
	ActiveProcessLinksOffset uint64
}
type DEBUGGEE_PROCESS_LIST_NEEDED_DETAILS = struct__DEBUGGEE_PROCESS_LIST_NEEDED_DETAILS
type PDEBUGGEE_PROCESS_LIST_NEEDED_DETAILS = *struct__DEBUGGEE_PROCESS_LIST_NEEDED_DETAILS
type struct__DEBUGGEE_THREAD_LIST_NEEDED_DETAILS struct {
	ThreadListHeadOffset     uint32
	ThreadListEntryOffset    uint32
	CidOffset                uint32
	PsActiveProcessHead      uint64
	ActiveProcessLinksOffset uint64
	Process                  uint64
}
type DEBUGGEE_THREAD_LIST_NEEDED_DETAILS = struct__DEBUGGEE_THREAD_LIST_NEEDED_DETAILS
type PDEBUGGEE_THREAD_LIST_NEEDED_DETAILS = *struct__DEBUGGEE_THREAD_LIST_NEEDED_DETAILS
type struct__DEBUGGEE_PROCESS_LIST_DETAILS_ENTRY struct {
	Eprocess      uint64
	Pid           uint32
	Cr3           uint64
	ImageFileName [16]uint8
}
type DEBUGGEE_PROCESS_LIST_DETAILS_ENTRY = struct__DEBUGGEE_PROCESS_LIST_DETAILS_ENTRY
type PDEBUGGEE_PROCESS_LIST_DETAILS_ENTRY = *struct__DEBUGGEE_PROCESS_LIST_DETAILS_ENTRY
type struct__DEBUGGEE_THREAD_LIST_DETAILS_ENTRY struct {
	Eprocess      uint64
	Ethread       uint64
	Pid           uint64
	Tid           uint64
	ImageFileName [16]uint8
}
type DEBUGGEE_THREAD_LIST_DETAILS_ENTRY = struct__DEBUGGEE_THREAD_LIST_DETAILS_ENTRY
type PDEBUGGEE_THREAD_LIST_DETAILS_ENTRY = *struct__DEBUGGEE_THREAD_LIST_DETAILS_ENTRY
type struct__DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS struct {
	ProcessListNeededDetails struct__DEBUGGEE_PROCESS_LIST_NEEDED_DETAILS
	ThreadListNeededDetails  struct__DEBUGGEE_THREAD_LIST_NEEDED_DETAILS
	QueryType                int32
	QueryAction              int32
	Count                    uint32
	Result                   uint64
}
type DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS = struct__DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS
type PDEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS = *struct__DEBUGGER_QUERY_ACTIVE_PROCESSES_OR_THREADS
type struct__DEBUGGER_SINGLE_CALLSTACK_FRAME struct {
	IsStackAddressValid   uint8
	IsValidAddress        uint8
	IsExecutable          uint8
	Value                 uint64
	InstructionBytesOnRip [7]uint8
}
type DEBUGGER_SINGLE_CALLSTACK_FRAME = struct__DEBUGGER_SINGLE_CALLSTACK_FRAME
type PDEBUGGER_SINGLE_CALLSTACK_FRAME = *struct__DEBUGGER_SINGLE_CALLSTACK_FRAME

const (
	DEBUGGER_CALLSTACK_DISPLAY_METHOD_WITHOUT_PARAMS int32 = 0
	DEBUGGER_CALLSTACK_DISPLAY_METHOD_WITH_PARAMS    int32 = 1
)

type DEBUGGER_CALLSTACK_DISPLAY_METHOD = int32
type struct__DEBUGGER_CALLSTACK_REQUEST struct {
	Is32Bit       uint8
	KernelStatus  uint32
	DisplayMethod int32
	Size          uint32
	FrameCount    uint32
	BaseAddress   uint64
	BufferSize    uint64
}
type DEBUGGER_CALLSTACK_REQUEST = struct__DEBUGGER_CALLSTACK_REQUEST
type PDEBUGGER_CALLSTACK_REQUEST = *struct__DEBUGGER_CALLSTACK_REQUEST
type struct__USERMODE_DEBUGGING_THREAD_OR_PROCESS_STATE_DETAILS struct {
	ProcessId uint32
	ThreadId  uint32
	IsProcess uint8
}
type USERMODE_DEBUGGING_THREAD_OR_PROCESS_STATE_DETAILS = struct__USERMODE_DEBUGGING_THREAD_OR_PROCESS_STATE_DETAILS
type PUSERMODE_DEBUGGING_THREAD_OR_PROCESS_STATE_DETAILS = *struct__USERMODE_DEBUGGING_THREAD_OR_PROCESS_STATE_DETAILS
type struct__DEBUGGER_EVENT_ACTION_RUN_SCRIPT_CONFIGURATION struct {
	ScriptBuffer                uint64
	ScriptLength                uint32
	ScriptPointer               uint32
	OptionalRequestedBufferSize uint32
}
type DEBUGGER_EVENT_ACTION_RUN_SCRIPT_CONFIGURATION = struct__DEBUGGER_EVENT_ACTION_RUN_SCRIPT_CONFIGURATION
type PDEBUGGER_EVENT_ACTION_RUN_SCRIPT_CONFIGURATION = *struct__DEBUGGER_EVENT_ACTION_RUN_SCRIPT_CONFIGURATION
type struct__DEBUGGER_EVENT_REQUEST_BUFFER struct {
	EnabledRequestBuffer uint8
	RequestBufferSize    uint32
	RequstBufferAddress  uint64
}
type DEBUGGER_EVENT_REQUEST_BUFFER = struct__DEBUGGER_EVENT_REQUEST_BUFFER
type PDEBUGGER_EVENT_REQUEST_BUFFER = *struct__DEBUGGER_EVENT_REQUEST_BUFFER
type struct__DEBUGGER_EVENT_REQUEST_CUSTOM_CODE struct {
	CustomCodeBufferSize        uint32
	CustomCodeBufferAddress     unsafe.Pointer
	OptionalRequestedBufferSize uint32
}
type DEBUGGER_EVENT_REQUEST_CUSTOM_CODE = struct__DEBUGGER_EVENT_REQUEST_CUSTOM_CODE
type PDEBUGGER_EVENT_REQUEST_CUSTOM_CODE = *struct__DEBUGGER_EVENT_REQUEST_CUSTOM_CODE

const (
	DEBUGGER_UD_COMMAND_ACTION_TYPE_NONE         int32 = int32(0)
	DEBUGGER_UD_COMMAND_ACTION_TYPE_PAUSE        int32 = 1
	DEBUGGER_UD_COMMAND_ACTION_TYPE_CONTINUE     int32 = 2
	DEBUGGER_UD_COMMAND_ACTION_TYPE_REGULAR_STEP int32 = 3
)

type DEBUGGER_UD_COMMAND_ACTION_TYPE = int32
type struct__DEBUGGER_UD_COMMAND_ACTION struct {
	ActionType     int32
	OptionalParam1 uint64
	OptionalParam2 uint64
	OptionalParam3 uint64
	OptionalParam4 uint64
}
type DEBUGGER_UD_COMMAND_ACTION = struct__DEBUGGER_UD_COMMAND_ACTION
type PDEBUGGER_UD_COMMAND_ACTION = *struct__DEBUGGER_UD_COMMAND_ACTION
type struct__DEBUGGER_UD_COMMAND_PACKET struct {
	UdAction                    struct__DEBUGGER_UD_COMMAND_ACTION
	ProcessDebuggingDetailToken uint64
	TargetThreadId              uint32
	ApplyToAllPausedThreads     uint8
	Result                      uint32
}
type DEBUGGER_UD_COMMAND_PACKET = struct__DEBUGGER_UD_COMMAND_PACKET
type PDEBUGGER_UD_COMMAND_PACKET = *struct__DEBUGGER_UD_COMMAND_PACKET

const (
	DEBUGGEE_DETAILS_AND_SWITCH_PROCESS_GET_PROCESS_DETAILS int32 = 0
	DEBUGGEE_DETAILS_AND_SWITCH_PROCESS_GET_PROCESS_LIST    int32 = 1
	DEBUGGEE_DETAILS_AND_SWITCH_PROCESS_PERFORM_SWITCH      int32 = 2
)

type DEBUGGEE_DETAILS_AND_SWITCH_PROCESS_TYPE = int32
type struct__DEBUGGEE_DETAILS_AND_SWITCH_PROCESS_PACKET struct {
	ActionType            int32
	ProcessId             uint32
	Process               uint64
	IsSwitchByClkIntr     uint8
	ProcessName           [16]uint8
	ProcessListSymDetails struct__DEBUGGEE_PROCESS_LIST_NEEDED_DETAILS
	Result                uint32
}
type DEBUGGEE_DETAILS_AND_SWITCH_PROCESS_PACKET = struct__DEBUGGEE_DETAILS_AND_SWITCH_PROCESS_PACKET
type PDEBUGGEE_DETAILS_AND_SWITCH_PROCESS_PACKET = *struct__DEBUGGEE_DETAILS_AND_SWITCH_PROCESS_PACKET

const (
	DEBUGGEE_DETAILS_AND_SWITCH_THREAD_PERFORM_SWITCH     int32 = 0
	DEBUGGEE_DETAILS_AND_SWITCH_THREAD_GET_THREAD_DETAILS int32 = 1
	DEBUGGEE_DETAILS_AND_SWITCH_THREAD_GET_THREAD_LIST    int32 = 2
)

type DEBUGGEE_DETAILS_AND_SWITCH_THREAD_TYPE = int32
type struct__DEBUGGEE_DETAILS_AND_SWITCH_THREAD_PACKET struct {
	ActionType            int32
	ThreadId              uint32
	ProcessId             uint32
	Thread                uint64
	Process               uint64
	CheckByClockInterrupt uint8
	ProcessName           [16]uint8
	ThreadListSymDetails  struct__DEBUGGEE_THREAD_LIST_NEEDED_DETAILS
	Result                uint32
}
type DEBUGGEE_DETAILS_AND_SWITCH_THREAD_PACKET = struct__DEBUGGEE_DETAILS_AND_SWITCH_THREAD_PACKET
type PDEBUGGEE_DETAILS_AND_SWITCH_THREAD_PACKET = *struct__DEBUGGEE_DETAILS_AND_SWITCH_THREAD_PACKET

const (
	DEBUGGER_REMOTE_STEPPING_REQUEST_STEP_OVER               int32 = 0
	DEBUGGER_REMOTE_STEPPING_REQUEST_STEP_IN                 int32 = 1
	DEBUGGER_REMOTE_STEPPING_REQUEST_INSTRUMENTATION_STEP_IN int32 = 2
)

type DEBUGGER_REMOTE_STEPPING_REQUEST = int32
type struct__DEBUGGEE_STEP_PACKET struct {
	StepType                  int32
	IsCurrentInstructionACall uint8
	CallLength                uint32
}
type DEBUGGEE_STEP_PACKET = struct__DEBUGGEE_STEP_PACKET
type PDEBUGGEE_STEP_PACKET = *struct__DEBUGGEE_STEP_PACKET
type struct__DEBUGGEE_FORMATS_PACKET struct {
	Value  uint64
	Result uint32
}
type DEBUGGEE_FORMATS_PACKET = struct__DEBUGGEE_FORMATS_PACKET
type PDEBUGGEE_FORMATS_PACKET = *struct__DEBUGGEE_FORMATS_PACKET
type struct__DEBUGGEE_SYMBOL_REQUEST_PACKET struct {
	ProcessId uint32
}
type DEBUGGEE_SYMBOL_REQUEST_PACKET = struct__DEBUGGEE_SYMBOL_REQUEST_PACKET
type PDEBUGGEE_SYMBOL_REQUEST_PACKET = *struct__DEBUGGEE_SYMBOL_REQUEST_PACKET
type struct__DEBUGGEE_BP_PACKET struct {
	Address uint64
	Pid     uint32
	Tid     uint32
	Core    uint32
	Result  uint32
}
type DEBUGGEE_BP_PACKET = struct__DEBUGGEE_BP_PACKET
type PDEBUGGEE_BP_PACKET = *struct__DEBUGGEE_BP_PACKET
type struct__DEBUGGEE_BP_DESCRIPTOR struct {
	BreakpointId           uint64
	BreakpointsList        struct__LIST_ENTRY
	Enabled                uint8
	Address                uint64
	PhysAddress            uint64
	Pid                    uint32
	Tid                    uint32
	Core                   uint32
	InstructionLength      uint16
	PreviousByte           uint8
	SetRflagsIFBitOnMtf    uint8
	AvoidReApplyBreakpoint uint8
}
type DEBUGGEE_BP_DESCRIPTOR = struct__DEBUGGEE_BP_DESCRIPTOR
type PDEBUGGEE_BP_DESCRIPTOR = *struct__DEBUGGEE_BP_DESCRIPTOR

const (
	DEBUGGEE_BREAKPOINT_MODIFICATION_REQUEST_LIST_BREAKPOINTS int32 = 0
	DEBUGGEE_BREAKPOINT_MODIFICATION_REQUEST_ENABLE           int32 = 1
	DEBUGGEE_BREAKPOINT_MODIFICATION_REQUEST_DISABLE          int32 = 2
	DEBUGGEE_BREAKPOINT_MODIFICATION_REQUEST_CLEAR            int32 = 3
)

type DEBUGGEE_BREAKPOINT_MODIFICATION_REQUEST = int32
type struct__DEBUGGEE_BP_LIST_OR_MODIFY_PACKET struct {
	BreakpointId uint64
	Request      int32
	Result       uint32
}
type DEBUGGEE_BP_LIST_OR_MODIFY_PACKET = struct__DEBUGGEE_BP_LIST_OR_MODIFY_PACKET
type PDEBUGGEE_BP_LIST_OR_MODIFY_PACKET = *struct__DEBUGGEE_BP_LIST_OR_MODIFY_PACKET

const (
	DEBUGGER_CONDITIONAL_JUMP_STATUS_ERROR                int32 = int32(0)
	DEBUGGER_CONDITIONAL_JUMP_STATUS_NOT_CONDITIONAL_JUMP int32 = 1
	DEBUGGER_CONDITIONAL_JUMP_STATUS_JUMP_IS_TAKEN        int32 = 2
	DEBUGGER_CONDITIONAL_JUMP_STATUS_JUMP_IS_NOT_TAKEN    int32 = 3
)

type DEBUGGER_CONDITIONAL_JUMP_STATUS = int32
type struct__DEBUGGEE_SCRIPT_PACKET struct {
	ScriptBufferSize    uint32
	ScriptBufferPointer uint32
	IsFormat            uint8
	Result              uint32
}
type DEBUGGEE_SCRIPT_PACKET = struct__DEBUGGEE_SCRIPT_PACKET
type PDEBUGGEE_SCRIPT_PACKET = *struct__DEBUGGEE_SCRIPT_PACKET
type struct__DEBUGGEE_RESULT_OF_SEARCH_PACKET struct {
	CountOfResults uint32
	Result         uint32
}
type DEBUGGEE_RESULT_OF_SEARCH_PACKET = struct__DEBUGGEE_RESULT_OF_SEARCH_PACKET
type PDEBUGGEE_RESULT_OF_SEARCH_PACKET = *struct__DEBUGGEE_RESULT_OF_SEARCH_PACKET
type struct__DEBUGGEE_REGISTER_READ_DESCRIPTION struct {
	RegisterID   uint32
	Value        uint64
	KernelStatus uint32
}
type DEBUGGEE_REGISTER_READ_DESCRIPTION = struct__DEBUGGEE_REGISTER_READ_DESCRIPTION
type PDEBUGGEE_REGISTER_READ_DESCRIPTION = *struct__DEBUGGEE_REGISTER_READ_DESCRIPTION
type struct__MODULE_SYMBOL_DETAIL struct {
	IsSymbolDetailsFound   uint8
	IsLocalSymbolPath      uint8
	IsSymbolPDBAvaliable   uint8
	IsUserMode             uint8
	BaseAddress            uint64
	FilePath               [260]int8
	ModuleSymbolPath       [260]int8
	ModuleSymbolGuidAndAge [60]int8
}
type MODULE_SYMBOL_DETAIL = struct__MODULE_SYMBOL_DETAIL
type PMODULE_SYMBOL_DETAIL = *struct__MODULE_SYMBOL_DETAIL
type struct__USERMODE_LOADED_MODULE_SYMBOLS struct {
	BaseAddress uint64
	Entrypoint  uint64
	FilePath    [260]*int16
}
type USERMODE_LOADED_MODULE_SYMBOLS = struct__USERMODE_LOADED_MODULE_SYMBOLS
type PUSERMODE_LOADED_MODULE_SYMBOLS = *struct__USERMODE_LOADED_MODULE_SYMBOLS
type struct__USERMODE_LOADED_MODULE_DETAILS struct {
	ProcessId        uint32
	OnlyCountModules uint8
	ModulesCount     uint32
	Result           uint32
}
type USERMODE_LOADED_MODULE_DETAILS = struct__USERMODE_LOADED_MODULE_DETAILS
type PUSERMODE_LOADED_MODULE_DETAILS = *struct__USERMODE_LOADED_MODULE_DETAILS

var SymbolMapCallback func(uint64, *int8, *int8, uint32) int32

type struct__DEBUGGER_UPDATE_SYMBOL_TABLE struct {
	TotalSymbols       uint32
	CurrentSymbolIndex uint32
	SymbolDetailPacket struct__MODULE_SYMBOL_DETAIL
}
type DEBUGGER_UPDATE_SYMBOL_TABLE = struct__DEBUGGER_UPDATE_SYMBOL_TABLE
type PDEBUGGER_UPDATE_SYMBOL_TABLE = *struct__DEBUGGER_UPDATE_SYMBOL_TABLE
type struct__DEBUGGEE_SYMBOL_UPDATE_RESULT struct {
	KernelStatus uint64
}
type DEBUGGEE_SYMBOL_UPDATE_RESULT = struct__DEBUGGEE_SYMBOL_UPDATE_RESULT
type PDEBUGGEE_SYMBOL_UPDATE_RESULT = *struct__DEBUGGEE_SYMBOL_UPDATE_RESULT

func _cgo_main() int32 {
	return int32(0)
}
func main() {
	os.Exit(int(_cgo_main()))
}
