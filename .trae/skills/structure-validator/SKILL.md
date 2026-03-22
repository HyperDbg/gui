---
name: "structure-validator"
description: "Validates structure layouts for kernel communication to prevent BSOD. Invoke when implementing new structures or modifying IOCTL communication code."
---

# Structure Validator for Kernel Communication

This skill ensures all structures used in kernel communication are properly validated to prevent BSOD (Blue Screen of Death) caused by memory layout mismatches.

## Overview

When communicating with Windows kernel drivers via IOCTL, structure layout must match exactly between Go and C/C++. Any mismatch in size, field offsets, or padding can cause system crashes.

## Validation Process

### Step 1: Identify Structures

Analyze all structures used in:
- `debugger/user_debugger.go`
- `debugger/kernel_debugger.go`
- `debugger/packet.go`

For each structure, identify:
- Expected size (from C/C++ source)
- Field offsets
- Padding requirements

### Step 2: Create Unit Tests

In `debugger/structures_test.go`, use the `t.Run` pattern:

```go
func TestAll(t *testing.T) {
    t.Run("StructureName", func(t *testing.T) {
        layout := StructLayout{
            Name:         "StructureName",
            ExpectedSize: 72,
            ExpectedFields: map[string]uintptr{
                "Field1": 0,
                "Field2": 8,
                // ...
            },
        }
        verifyStruct[StructureName](t, layout)
        verifySerialize(t, layout, StructureName{
            // test values
        })
    })
}
```

### Step 3: Implement Validation Interface

Each structure should implement a validation interface:

```go
type Validator interface {
    Validate() error
}

type Serializer interface {
    ExpectedSize() uintptr
    ExpectedSerSize() uintptr
}
```

Example implementation:

```go
func (s *DebuggerAttachDetachUserModeProcess) Validate() error {
    if s.ProcessId == 0 {
        return errors.New("ProcessId cannot be zero")
    }
    if s.ThreadId == 0 {
        return errors.New("ThreadId cannot be zero")
    }
    return nil
}

func (s *DebuggerAttachDetachUserModeProcess) ExpectedSize() uintptr {
    return 72
}

func (s *DebuggerAttachDetachUserModeProcess) ExpectedSerSize() uintptr {
    return 72
}
```

### Step 4: Validate Before Sending

**CRITICAL**: Every IOCTL send method MUST call the validation interface before sending data to kernel.

#### Required Validation Pattern

Before any `SendReceive` or `Send` call, you MUST:

1. **Validate the structure** using the `Validate()` method
2. **Verify the size** using `ExpectedSize()` method
3. **Log validation failures** and return early to prevent BSOD

```go
func (s *UserDebug) SetBreakpoint(address uint64) {
    // ... input validation ...

    bpPacket := DebuggeeBpPacket{
        Address: address,
        Pid:     s.activeProcess.ProcessId,
        Tid:     s.activeProcess.ThreadId,
        // ...
    }

    // MANDATORY: Validate structure before sending
    if err := bpPacket.Validate(); err != nil {
        mylog.Warning("断点包验证失败", err)
        return
    }

    // MANDATORY: Verify size matches expected
    if unsafe.Sizeof(bpPacket) != bpPacket.ExpectedSize() {
        mylog.Warning("断点包大小不匹配", "got", unsafe.Sizeof(bpPacket), "want", bpPacket.ExpectedSize())
        return
    }

    // Now safe to send to kernel
    buffer := new(bytes.Buffer)
    if err := binary.Write(buffer, binary.LittleEndian, &bpPacket); err != nil {
        mylog.Warning("序列化失败", err)
        return
    }

    response := s.packeter.DriverProvider.SendReceive(buffer, IoctlSetBreakpointUserDebugger)
    // ... handle response ...
}
```

#### Validation Checklist for Every Send Method

- [ ] Call `packet.Validate()` and check error
- [ ] Call `unsafe.Sizeof(packet) == packet.ExpectedSize()`
- [ ] Return early on validation failure
- [ ] Log meaningful error messages

#### Example: Attach Process Validation

```go
func (s *UserDebug) StartProcess(path string) {
    // ... process creation ...

    attachReq := DebuggerAttachDetachUserModeProcess{
        IsStartingNewProcess: 1,
        ProcessId:            procInfo.ProcessId,
        ThreadId:             procInfo.ThreadId,
        Action:               DebuggerAttachDetachUserModeProcessActionAttach,
    }

    // MANDATORY validation
    if err := attachReq.Validate(); err != nil {
        mylog.Warning("attach请求验证失败", err)
        return
    }

    if unsafe.Sizeof(attachReq) != attachReq.ExpectedSize() {
        mylog.Warning("attach请求大小不匹配", "got", unsafe.Sizeof(attachReq), "want", attachReq.ExpectedSize())
        return
    }

    // Safe to send
    buffer := new(bytes.Buffer)
    binary.Write(buffer, binary.LittleEndian, &attachReq)
    response := s.packeter.DriverProvider.SendReceive(buffer, IoctlDebuggerAttachDetachUserModeProcess)
    // ...
}
```

## Structure Padding Rules

### Windows x64 Alignment

| Type | Alignment | Size |
|------|-----------|------|
| uint8 | 1 | 1 |
| uint16 | 2 | 2 |
| uint32 | 4 | 4 |
| uint64 | 8 | 8 |
| bool | 1 | 1 |
| arrays | element alignment | element size * count |

### Padding Example

```go
type Example struct {
    A uint8    // offset 0, size 1
    _ [3]byte  // padding 3 bytes for alignment
    B uint32   // offset 4, size 4
    C uint64   // offset 8, size 8
    D uint8    // offset 16, size 1
    _ [7]byte  // padding 7 bytes for total size multiple of 8
}
// Total: 24 bytes
```

## Common BSOD Causes

1. **Size mismatch**: Go struct size != C struct size
2. **Offset mismatch**: Field at wrong position
3. **Missing padding**: Alignment not preserved
4. **Invalid values**: Zero pointers, invalid IDs
5. **Buffer overflow**: Output buffer too small

## Testing Checklist

- [ ] Structure size matches C definition
- [ ] All field offsets verified
- [ ] Serialization/deserialization tested
- [ ] Validation method implemented
- [ ] Pre-send validation added to all IOCTL calls
- [ ] Edge cases tested (zero values, max values)

## Files to Update

1. `debugger/structures.go` - Add validation methods
2. `debugger/structures_test.go` - Add unit tests
3. `debugger/user_debugger.go` - Add pre-send validation
4. `debugger/kernel_debugger.go` - Add pre-send validation
5. `debugger/packet.go` - Add pre-send validation
