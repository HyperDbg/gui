package debugger

import (
	"bytes"
	"encoding/binary"
	"reflect"
	"testing"
	"unsafe"

	"github.com/ddkwork/golibrary/std/mylog"
)

type StructLayout struct {
	Name            string
	Size            uintptr
	ExpectedSize    uintptr
	ExpectedSerSize uintptr
	FieldOffsets    map[string]uintptr
	ExpectedFields  map[string]uintptr
}

func TestAll(t *testing.T) {
	t.Run("DebuggerAttachDetachUserModeProcess", func(t *testing.T) {
		layout := StructLayout{
			Name:         "DebuggerAttachDetachUserModeProcess",
			ExpectedSize: 72,
			ExpectedFields: map[string]uintptr{
				"IsStartingNewProcess":            0,
				"ProcessId":                       4,
				"ThreadId":                        8,
				"CheckCallbackAtFirstInstruction": 12,
				"Is32Bit":                         13,
				"Rip":                             16,
				"InstructionBytesOnRip":           24,
				"SizeOfInstruction":               40,
				"IsPaused":                        44,
				"Action":                          48,
				"CountOfActiveDebuggingThreadsAndProcesses": 52,
				"Token":  56,
				"Result": 64,
			},
		}
		verifyStruct[DebuggerAttachDetachUserModeProcess](t, layout)
		verifySerialize(t, layout, DebuggerAttachDetachUserModeProcess{
			IsStartingNewProcess:            1,
			ProcessId:                       1234,
			ThreadId:                        5678,
			CheckCallbackAtFirstInstruction: 0,
			Is32Bit:                         0,
			Action:                          DebuggerAttachDetachUserModeProcessActionAttach,
			Token:                           0x1234567890ABCDEF,
			Result:                          0,
		})
	})

	t.Run("DebuggeeBpPacket", func(t *testing.T) {
		layout := StructLayout{
			Name:            "DebuggeeBpPacket",
			ExpectedSize:    32,
			ExpectedSerSize: 32,
			ExpectedFields: map[string]uintptr{
				"Address":           0,
				"Pid":               8,
				"Tid":               12,
				"Core":              16,
				"RemoveAfterHit":    20,
				"CheckForCallbacks": 21,
				"Result":            24,
			},
		}
		verifyStruct[DebuggeeBpPacket](t, layout)
		verifySerialize(t, layout, DebuggeeBpPacket{
			Address:           0x140001000,
			Pid:               1234,
			Tid:               5678,
			Core:              0,
			RemoveAfterHit:    0,
			CheckForCallbacks: 0,
			Result:            0,
		})
	})

	t.Run("DebuggeeBpListOrModifyPacket", func(t *testing.T) {
		layout := StructLayout{
			Name:            "DebuggeeBpListOrModifyPacket",
			ExpectedSize:    16,
			ExpectedSerSize: 16,
			ExpectedFields: map[string]uintptr{
				"BreakpointId": 0,
				"Request":      8,
				"Result":       12,
			},
		}
		verifyStruct[DebuggeeBpListOrModifyPacket](t, layout)
		verifySerialize(t, layout, DebuggeeBpListOrModifyPacket{
			BreakpointId: 1,
			Request:      DebuggeeBreakpointModificationRequestClear,
			Result:       0,
		})
	})

	t.Run("DebuggerUdCommandAction", func(t *testing.T) {
		layout := StructLayout{
			Name:            "DebuggerUdCommandAction",
			ExpectedSize:    40,
			ExpectedSerSize: 40,
			ExpectedFields: map[string]uintptr{
				"ActionType":     0,
				"OptionalParam1": 8,
				"OptionalParam2": 16,
				"OptionalParam3": 24,
				"OptionalParam4": 32,
			},
		}
		verifyStruct[DebuggerUdCommandAction](t, layout)
	})

	t.Run("DebuggerUdCommandPacket", func(t *testing.T) {
		layout := StructLayout{
			Name:            "DebuggerUdCommandPacket",
			ExpectedSize:    64,
			ExpectedSerSize: 64,
			ExpectedFields: map[string]uintptr{
				"UdAction":                    0,
				"ProcessDebuggingDetailToken": 40,
				"TargetThreadId":              48,
				"ApplyToAllPausedThreads":     52,
				"WaitForEventCompletion":      53,
				"Result":                      56,
			},
		}
		verifyStruct[DebuggerUdCommandPacket](t, layout)
	})

	t.Run("DebuggerReadMemory", func(t *testing.T) {
		layout := StructLayout{
			Name:            "DebuggerReadMemory",
			ExpectedSize:    48,
			ExpectedSerSize: 48,
			ExpectedFields: map[string]uintptr{
				"Pid":            0,
				"Address":        8,
				"Size":           16,
				"GetAddressMode": 20,
				"AddressMode":    24,
				"MemoryType":     28,
				"ReadingType":    32,
				"ReturnLength":   36,
				"KernelStatus":   40,
			},
		}
		verifyStruct[DebuggerReadMemory](t, layout)
	})

	t.Run("DebuggerVa2paAndPa2vaCommands", func(t *testing.T) {
		layout := StructLayout{
			Name:            "DebuggerVa2paAndPa2vaCommands",
			ExpectedSize:    32,
			ExpectedSerSize: 32,
			ExpectedFields: map[string]uintptr{
				"VirtualAddress":     0,
				"PhysicalAddress":    8,
				"ProcessId":          16,
				"IsVirtual2Physical": 20,
				"KernelStatus":       24,
			},
		}
		verifyStruct[DebuggerVa2paAndPa2vaCommands](t, layout)
	})

	t.Run("DebuggerGeneralEventDetails", func(t *testing.T) {
		layout := StructLayout{
			Name:            "DebuggerGeneralEventDetails",
			ExpectedSize:    24,
			ExpectedSerSize: 24,
			ExpectedFields: map[string]uintptr{
				"Tag":       0,
				"ProcessId": 8,
				"IsEnabled": 12,
				"Type":      16,
			},
		}
		verifyStruct[DebuggerGeneralEventDetails](t, layout)
	})

	t.Run("DebuggerEventAction", func(t *testing.T) {
		layout := StructLayout{
			Name:            "DebuggerEventAction",
			ExpectedSize:    56,
			ExpectedSerSize: 45,
			ExpectedFields: map[string]uintptr{
				"ActionType":                0,
				"ActionTag":                 8,
				"SendTheResultsImmediately": 16,
				"OptionalParam1":            24,
				"OptionalParam2":            32,
				"OptionalParam3":            40,
				"OptionalParam4":            48,
			},
		}
		verifyStruct[DebuggerEventAction](t, layout)
	})

	t.Run("DebuggerModifyEventRequest", func(t *testing.T) {
		layout := StructLayout{
			Name:            "DebuggerModifyEventRequest",
			ExpectedSize:    24,
			ExpectedSerSize: 17,
			ExpectedFields: map[string]uintptr{
				"Tag":       0,
				"IsEnabled": 8,
				"Type":      12,
				"ProcessId": 16,
			},
		}
		verifyStruct[DebuggerModifyEventRequest](t, layout)
	})

	t.Run("DebuggerEventAndActionResult", func(t *testing.T) {
		layout := StructLayout{
			Name:            "DebuggerEventAndActionResult",
			ExpectedSize:    8,
			ExpectedSerSize: 5,
			ExpectedFields: map[string]uintptr{
				"IsSuccessful": 0,
				"Error":        4,
			},
		}
		verifyStruct[DebuggerEventAndActionResult](t, layout)
	})

	t.Run("DebuggeeKdPausedPacket", func(t *testing.T) {
		layout := StructLayout{
			Name:         "DebuggeeKdPausedPacket",
			ExpectedSize: 72,
			ExpectedFields: map[string]uintptr{
				"Rip":                    0,
				"IsProcessorOn32BitMode": 8,
				"IgnoreDisassembling":    9,
				"PausingReason":          12,
				"CurrentCore":            16,
				"EventTag":               24,
				"EventCallingStage":      32,
				"Rflags":                 40,
				"InstructionBytesOnRip":  48,
				"ReadInstructionLen":     64,
			},
		}
		verifyStruct[DebuggeeKdPausedPacket](t, layout)
	})

	t.Run("DebuggeeDetailsAndSwitchProcessPacket", func(t *testing.T) {
		layout := StructLayout{
			Name:         "DebuggeeDetailsAndSwitchProcessPacket",
			ExpectedSize: 72,
			ExpectedFields: map[string]uintptr{
				"ActionType":            0,
				"ProcessId":             4,
				"Process":               8,
				"IsSwitchByClkIntr":     16,
				"ProcessName":           17,
				"ProcessListSymDetails": 40,
				"Result":                64,
			},
		}
		verifyStruct[DebuggeeDetailsAndSwitchProcessPacket](t, layout)
	})

	t.Run("DebuggeeDetailsAndSwitchThreadPacket", func(t *testing.T) {
		layout := StructLayout{
			Name:         "DebuggeeDetailsAndSwitchThreadPacket",
			ExpectedSize: 56,
			ExpectedFields: map[string]uintptr{
				"ActionType":  0,
				"ThreadId":    4,
				"Thread":      8,
				"ProcessId":   16,
				"Process":     24,
				"ProcessName": 32,
				"Result":      48,
			},
		}
		verifyStruct[DebuggeeDetailsAndSwitchThreadPacket](t, layout)
	})

	t.Run("DebuggeeRegisterReadDescription", func(t *testing.T) {
		layout := StructLayout{
			Name:         "DebuggeeRegisterReadDescription",
			ExpectedSize: 24,
			ExpectedFields: map[string]uintptr{
				"RegisterId":   0,
				"Value":        8,
				"KernelStatus": 16,
			},
		}
		verifyStruct[DebuggeeRegisterReadDescription](t, layout)
	})

	t.Run("PageTableEntries", func(t *testing.T) {
		layout := StructLayout{
			Name:         "PageTableEntries",
			ExpectedSize: 80,
			ExpectedFields: map[string]uintptr{
				"VirtualAddress":      0,
				"ProcessId":           8,
				"Pml4eVirtualAddress": 16,
				"Pml4eValue":          24,
				"PdpteVirtualAddress": 32,
				"PdpteValue":          40,
				"PdeVirtualAddress":   48,
				"PdeValue":            56,
				"PteVirtualAddress":   64,
				"PteValue":            72,
			},
		}
		verifyStruct[PageTableEntries](t, layout)
	})

	t.Run("DebuggerReadPageTableEntriesDetails", func(t *testing.T) {
		layout := StructLayout{
			Name:         "DebuggerReadPageTableEntriesDetails",
			ExpectedSize: 88,
			ExpectedFields: map[string]uintptr{
				"VirtualAddress":      0,
				"ProcessId":           8,
				"Pml4eVirtualAddress": 16,
				"Pml4eValue":          24,
				"PdpteVirtualAddress": 32,
				"PdpteValue":          40,
				"PdeVirtualAddress":   48,
				"PdeValue":            56,
				"PteVirtualAddress":   64,
				"PteValue":            72,
				"KernelStatus":        80,
			},
		}
		verifyStruct[DebuggerReadPageTableEntriesDetails](t, layout)
	})

	t.Run("CoreInfo", func(t *testing.T) {
		layout := StructLayout{
			Name:         "CoreInfo",
			ExpectedSize: 16,
			ExpectedFields: map[string]uintptr{
				"CoreID":             0,
				"CurrentCore":        4,
				"IsHalted":           5,
				"InstructionPointer": 8,
			},
		}
		verifyStruct[CoreInfo](t, layout)
	})

	t.Run("VersionInfo", func(t *testing.T) {
		layout := StructLayout{
			Name:         "VersionInfo",
			ExpectedSize: 20,
			ExpectedFields: map[string]uintptr{
				"Major":   0,
				"Minor":   4,
				"Patch":   8,
				"Build":   12,
				"IsBeta":  16,
				"IsDebug": 17,
			},
		}
		verifyStruct[VersionInfo](t, layout)
	})

	t.Run("UserModeProcessDetails", func(t *testing.T) {
		layout := StructLayout{
			Name:         "UserModeProcessDetails",
			ExpectedSize: 48,
			ExpectedFields: map[string]uintptr{
				"Token":                       0,
				"Enabled":                     8,
				"ActiveThreadId":              12,
				"EntrypointOfMainModule":      16,
				"BaseAddressOfMainModule":     24,
				"Eprocess":                    32,
				"ProcessId":                   40,
				"Is32Bit":                     44,
				"IsOnTheStartingPhase":        45,
				"IsOnThreadInterceptingPhase": 46,
			},
		}
		verifyStruct[UserModeProcessDetails](t, layout)
	})

	t.Run("DebuggerSetBreakpointUserDebugger", func(t *testing.T) {
		layout := StructLayout{
			Name:         "DebuggerSetBreakpointUserDebugger",
			ExpectedSize: 40,
			ExpectedFields: map[string]uintptr{
				"ProcessId":                   0,
				"ThreadId":                    4,
				"BreakpointType":              8,
				"Address":                     16,
				"ProcessDebuggingDetailToken": 24,
				"Result":                      32,
			},
		}
		verifyStruct[DebuggerSetBreakpointUserDebugger](t, layout)
	})

	t.Run("DebuggerGetUserModeModuleDetailsRequest", func(t *testing.T) {
		layout := StructLayout{
			Name:         "DebuggerGetUserModeModuleDetailsRequest",
			ExpectedSize: 16,
			ExpectedFields: map[string]uintptr{
				"ProcessDebuggingDetailToken": 0,
			},
		}
		verifyStruct[DebuggerGetUserModeModuleDetailsRequest](t, layout)
	})

	t.Run("DebuggerModuleDetailResponse", func(t *testing.T) {
		layout := StructLayout{
			Name:         "DebuggerModuleDetailResponse",
			ExpectedSize: 16,
			ExpectedFields: map[string]uintptr{
				"BaseAddress": 0,
				"Size":        8,
				"Is32Bit":     12,
			},
		}
		verifyStruct[DebuggerModuleDetailResponse](t, layout)
	})
}

func verifyStruct[T any](t *testing.T, layout StructLayout) {
	var zero T
	typ := reflect.TypeOf(zero)
	layout.Size = unsafe.Sizeof(zero)

	t.Logf("%s size: %d bytes", layout.Name, layout.Size)

	if layout.Size != layout.ExpectedSize {
		t.Errorf("%s 大小不匹配: got %d, want %d", layout.Name, layout.Size, layout.ExpectedSize)
	}

	layout.FieldOffsets = make(map[string]uintptr)
	for field := range typ.Fields() {
		if field.IsExported() {
			offset := field.Offset
			layout.FieldOffsets[field.Name] = offset
			if expected, ok := layout.ExpectedFields[field.Name]; ok {
				if offset != expected {
					t.Errorf("%s.%s 偏移不匹配: got %d, want %d", layout.Name, field.Name, offset, expected)
				} else {
					t.Logf("%s.%s 偏移: %d ✓", layout.Name, field.Name, offset)
				}
			}
		}
	}
}

func verifySerialize[T any](t *testing.T, layout StructLayout, value T) {
	buffer := new(bytes.Buffer)
	mylog.Check(binary.Write(buffer, binary.LittleEndian, &value))

	t.Logf("%s 序列化后大小: %d bytes", layout.Name, buffer.Len())
	t.Logf("%s 序列化后数据: %x", layout.Name, buffer.Bytes())

	expectedSerSize := layout.ExpectedSerSize
	if expectedSerSize == 0 {
		expectedSerSize = layout.ExpectedSize
	}
	if buffer.Len() != int(expectedSerSize) {
		t.Errorf("%s 序列化后大小不匹配: got %d, want %d", layout.Name, buffer.Len(), expectedSerSize)
	}

	var decoded T
	mylog.Check(binary.Read(bytes.NewBuffer(buffer.Bytes()), binary.LittleEndian, &decoded))

	decodedValue := reflect.ValueOf(decoded)
	originalValue := reflect.ValueOf(value)
	for i := 0; i < decodedValue.NumField(); i++ {
		field := decodedValue.Type().Field(i)
		if field.IsExported() {
			decodedField := decodedValue.Field(i)
			originalField := originalValue.Field(i)
			if !reflect.DeepEqual(decodedField.Interface(), originalField.Interface()) {
				t.Errorf("%s.%s 不匹配: got %v, want %v", layout.Name, field.Name, decodedField.Interface(), originalField.Interface())
			}
		}
	}
}
