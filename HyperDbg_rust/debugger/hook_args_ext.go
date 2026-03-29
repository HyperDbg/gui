package debugger

import (
	"encoding/binary"
	"unsafe"
)

func (a *NtDeviceIoControlFileArgs) WriteOutputString(s string) {
	if a.OutputBuffer == 0 {
		return
	}
	buf := []byte(s)
	for i, b := range buf {
		if i >= int(a.OutputBufferLength) {
			break
		}
		*(*byte)(unsafe.Pointer(a.OutputBuffer + uintptr(i))) = b
	}
}

func (a *NtDeviceIoControlFileArgs) WriteOutputBytes(data []byte) {
	if a.OutputBuffer == 0 {
		return
	}
	for i, b := range data {
		if i >= int(a.OutputBufferLength) {
			break
		}
		*(*byte)(unsafe.Pointer(a.OutputBuffer + uintptr(i))) = b
	}
}

func (a *NtDeviceIoControlFileArgs) WriteDiskInfo(model, serial string) {
	if a.OutputBuffer == 0 || a.OutputBufferLength < 512 {
		return
	}

	modelBytes := []byte(model)
	serialBytes := []byte(serial)

	modelPadded := make([]byte, 40)
	serialPadded := make([]byte, 20)
	copy(modelPadded, modelBytes)
	copy(serialPadded, serialBytes)

	modelSwapped := make([]byte, 40)
	serialSwapped := make([]byte, 20)
	for i := 0; i < 40; i += 2 {
		if i+1 < len(modelPadded) {
			modelSwapped[i] = modelPadded[i+1]
			modelSwapped[i+1] = modelPadded[i]
		}
	}
	for i := 0; i < 20; i += 2 {
		if i+1 < len(serialPadded) {
			serialSwapped[i] = serialPadded[i+1]
			serialSwapped[i+1] = serialPadded[i]
		}
	}

	serialOffset := uintptr(10)
	modelOffset := uintptr(27)

	for i, b := range serialSwapped {
		if serialOffset+uintptr(i) < a.OutputBuffer+uintptr(a.OutputBufferLength) {
			*(*byte)(unsafe.Pointer(a.OutputBuffer + serialOffset + uintptr(i))) = b
		}
	}

	for i, b := range modelSwapped {
		if modelOffset+uintptr(i) < a.OutputBuffer+uintptr(a.OutputBufferLength) {
			*(*byte)(unsafe.Pointer(a.OutputBuffer + modelOffset + uintptr(i))) = b
		}
	}
}

func (a *NtDeviceIoControlFileArgs) WriteOutputUint32(v uint32) {
	if a.OutputBuffer == 0 || a.OutputBufferLength < 4 {
		return
	}
	*(*uint32)(unsafe.Pointer(a.OutputBuffer)) = v
}

func (a *NtDeviceIoControlFileArgs) WriteOutputUint64(v uint64) {
	if a.OutputBuffer == 0 || a.OutputBufferLength < 8 {
		return
	}
	*(*uint64)(unsafe.Pointer(a.OutputBuffer)) = v
}

func (a *NtDeviceIoControlFileArgs) ReadInputBytes() []byte {
	if a.InputBuffer == 0 || a.InputBufferLength == 0 {
		return nil
	}
	buf := make([]byte, a.InputBufferLength)
	for i := uint32(0); i < a.InputBufferLength; i++ {
		buf[i] = *(*byte)(unsafe.Pointer(a.InputBuffer + uintptr(i)))
	}
	return buf
}

func (a *NtDeviceIoControlFileArgs) ReadInputString() string {
	if a.InputBuffer == 0 || a.InputBufferLength == 0 {
		return ""
	}
	buf := make([]byte, 0, a.InputBufferLength)
	for i := uint32(0); i < a.InputBufferLength; i++ {
		b := *(*byte)(unsafe.Pointer(a.InputBuffer + uintptr(i)))
		if b == 0 {
			break
		}
		buf = append(buf, b)
	}
	return string(buf)
}

type IopXxxControlFileArgs struct {
	FileHandle        uintptr `json:"FileHandle"`
	IoControlCode     uint32  `json:"IoControlCode"`
	InputBuffer       uintptr `json:"InputBuffer"`
	InputBufferLength uint32  `json:"InputBufferLength"`
	OutputBuffer      uintptr `json:"OutputBuffer"`
	OutputBufferLength uint32 `json:"OutputBufferLength"`
}

func (a *IopXxxControlFileArgs) WriteOutputString(s string) {
	if a.OutputBuffer == 0 {
		return
	}
	buf := []byte(s)
	for i, b := range buf {
		if i >= int(a.OutputBufferLength) {
			break
		}
		*(*byte)(unsafe.Pointer(a.OutputBuffer + uintptr(i))) = b
	}
}

func (a *IopXxxControlFileArgs) WriteOutputBytes(data []byte) {
	if a.OutputBuffer == 0 {
		return
	}
	for i, b := range data {
		if i >= int(a.OutputBufferLength) {
			break
		}
		*(*byte)(unsafe.Pointer(a.OutputBuffer + uintptr(i))) = b
	}
}

type CpuidArgs struct {
	Leaf uint32 `json:"Leaf"`
	SubLeaf uint32 `json:"SubLeaf"`
	EAX uint32 `json:"EAX"`
	EBX uint32 `json:"EBX"`
	ECX uint32 `json:"ECX"`
	EDX uint32 `json:"EDX"`
}

type NtReadFileArgs struct {
	FileHandle           uintptr `json:"FileHandle"`
	Event                uintptr `json:"Event"`
	ApcRoutine           uintptr `json:"ApcRoutine"`
	ApcContext           uintptr `json:"ApcContext"`
	IoStatusBlock        uintptr `json:"IoStatusBlock"`
	Buffer               uintptr `json:"Buffer"`
	Length               uint32  `json:"Length"`
	ByteOffset           uintptr `json:"ByteOffset"`
	Key                  uintptr `json:"Key"`
}

func (a *NtReadFileArgs) ReadBytes() []byte {
	if a.Buffer == 0 || a.Length == 0 {
		return nil
	}
	buf := make([]byte, a.Length)
	for i := uint32(0); i < a.Length; i++ {
		buf[i] = *(*byte)(unsafe.Pointer(a.Buffer + uintptr(i)))
	}
	return buf
}

func (a *NtReadFileArgs) WriteBytes(data []byte) {
	if a.Buffer == 0 {
		return
	}
	for i, b := range data {
		if i >= int(a.Length) {
			break
		}
		*(*byte)(unsafe.Pointer(a.Buffer + uintptr(i))) = b
	}
}

type NtQuerySystemInformationArgs struct {
	SystemInformationClass uint32  `json:"SystemInformationClass"`
	SystemInformation      uintptr `json:"SystemInformation"`
	SystemInformationLength uint32  `json:"SystemInformationLength"`
	ReturnLength           uintptr `json:"ReturnLength"`
}

func (a *NtQuerySystemInformationArgs) WriteBytes(data []byte) {
	if a.SystemInformation == 0 {
		return
	}
	for i, b := range data {
		if i >= int(a.SystemInformationLength) {
			break
		}
		*(*byte)(unsafe.Pointer(a.SystemInformation + uintptr(i))) = b
	}
}

func WriteMemoryU32(addr uintptr, v uint32) {
	if addr == 0 {
		return
	}
	*(*uint32)(unsafe.Pointer(addr)) = v
}

func WriteMemoryU64(addr uintptr, v uint64) {
	if addr == 0 {
		return
	}
	*(*uint64)(unsafe.Pointer(addr)) = v
}

func WriteMemoryBytes(addr uintptr, data []byte) {
	if addr == 0 {
		return
	}
	for i, b := range data {
		*(*byte)(unsafe.Pointer(addr + uintptr(i))) = b
	}
}

func ReadMemoryBytes(addr uintptr, length uint32) []byte {
	if addr == 0 || length == 0 {
		return nil
	}
	buf := make([]byte, length)
	for i := uint32(0); i < length; i++ {
		buf[i] = *(*byte)(unsafe.Pointer(addr + uintptr(i)))
	}
	return buf
}

func WriteMemoryString(addr uintptr, s string) {
	if addr == 0 {
		return
	}
	buf := []byte(s)
	for i, b := range buf {
		*(*byte)(unsafe.Pointer(addr + uintptr(i))) = b
	}
}

func ReadMemoryString(addr uintptr, maxLen uint32) string {
	if addr == 0 || maxLen == 0 {
		return ""
	}
	buf := make([]byte, 0, maxLen)
	for i := uint32(0); i < maxLen; i++ {
		b := *(*byte)(unsafe.Pointer(addr + uintptr(i)))
		if b == 0 {
			break
		}
		buf = append(buf, b)
	}
	return string(buf)
}

func LittleEndianPutUint16(b []byte, v uint16) {
	binary.LittleEndian.PutUint16(b, v)
}

func LittleEndianPutUint32(b []byte, v uint32) {
	binary.LittleEndian.PutUint32(b, v)
}

func LittleEndianPutUint64(b []byte, v uint64) {
	binary.LittleEndian.PutUint64(b, v)
}

func LittleEndianUint16(b []byte) uint16 {
	return binary.LittleEndian.Uint16(b)
}

func LittleEndianUint32(b []byte) uint32 {
	return binary.LittleEndian.Uint32(b)
}

func LittleEndianUint64(b []byte) uint64 {
	return binary.LittleEndian.Uint64(b)
}
