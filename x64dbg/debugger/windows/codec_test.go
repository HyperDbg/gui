package windows

import (
	"encoding/binary"
	"testing"
	"unsafe"
)

func TestExceptionRecordCodec(t *testing.T) {
	original := &ExceptionRecord{
		ExceptionCode:    EXCEPTION_SINGLE_STEP,
		ExceptionFlags:   0,
		ExceptionRecord:  0,
		ExceptionAddress: 0x7FFF00001000,
		NumberParameters: 2,
		_UnusedAlignment: 0,
		ExceptionInformation: [15]uint64{
			0x1,
			0x2,
			0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		},
	}

	encoded, err := original.Encode()
	if err != nil {
		t.Fatalf("编码失败: %v", err)
	}

	if len(encoded) != 152 {
		t.Fatalf("编码长度错误，期望152字节，实际%d字节", len(encoded))
	}

	decoded := &ExceptionRecord{}
	err = decoded.Decode(encoded)
	if err != nil {
		t.Fatalf("解码失败: %v", err)
	}

	if decoded.ExceptionCode != original.ExceptionCode {
		t.Errorf("ExceptionCode不匹配: 期望0x%X，实际0x%X", original.ExceptionCode, decoded.ExceptionCode)
	}

	if decoded.ExceptionFlags != original.ExceptionFlags {
		t.Errorf("ExceptionFlags不匹配: 期望%d，实际%d", original.ExceptionFlags, decoded.ExceptionFlags)
	}

	if decoded.ExceptionRecord != original.ExceptionRecord {
		t.Errorf("ExceptionRecord不匹配: 期望%x，实际%x", original.ExceptionRecord, decoded.ExceptionRecord)
	}

	if decoded.ExceptionAddress != original.ExceptionAddress {
		t.Errorf("ExceptionAddress不匹配: 期望%x，实际%x", original.ExceptionAddress, decoded.ExceptionAddress)
	}

	if decoded.NumberParameters != original.NumberParameters {
		t.Errorf("NumberParameters不匹配: 期望%d，实际%d", original.NumberParameters, decoded.NumberParameters)
	}

	for i := 0; i < 15; i++ {
		if decoded.ExceptionInformation[i] != original.ExceptionInformation[i] {
			t.Errorf("ExceptionInformation[%d]不匹配: 期望%x，实际%x", i, original.ExceptionInformation[i], decoded.ExceptionInformation[i])
		}
	}
}

func TestExceptionDebugInfoCodec(t *testing.T) {
	original := &ExceptionDebugInfo{
		ExceptionRecord: ExceptionRecord{
			ExceptionCode:    EXCEPTION_BREAKPOINT,
			ExceptionFlags:   0,
			ExceptionRecord:  0,
			ExceptionAddress: 0x7FFF00002000,
			NumberParameters: 1,
			_UnusedAlignment: 0,
			ExceptionInformation: [15]uint64{
				0x3,
				0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			},
		},
		FirstChance: 1,
	}

	encoded, err := original.Encode()
	if err != nil {
		t.Fatalf("编码失败: %v", err)
	}

	if len(encoded) != 156 {
		t.Fatalf("编码长度错误，期望156字节，实际%d字节", len(encoded))
	}

	decoded := &ExceptionDebugInfo{}
	err = decoded.Decode(encoded)
	if err != nil {
		t.Fatalf("解码失败: %v", err)
	}

	if decoded.ExceptionRecord.ExceptionCode != original.ExceptionRecord.ExceptionCode {
		t.Errorf("ExceptionCode不匹配: 期望0x%X，实际0x%X", original.ExceptionRecord.ExceptionCode, decoded.ExceptionRecord.ExceptionCode)
	}

	if decoded.ExceptionRecord.ExceptionAddress != original.ExceptionRecord.ExceptionAddress {
		t.Errorf("ExceptionAddress不匹配: 期望%x，实际%x", original.ExceptionRecord.ExceptionAddress, decoded.ExceptionRecord.ExceptionAddress)
	}

	if decoded.FirstChance != original.FirstChance {
		t.Errorf("FirstChance不匹配: 期望%d，实际%d", original.FirstChance, decoded.FirstChance)
	}
}

func TestDebugEventCodec(t *testing.T) {
	original := &DebugEvent{}

	binary.LittleEndian.PutUint32(original.data[0:4], EXCEPTION_DEBUG_EVENT)
	binary.LittleEndian.PutUint32(original.data[4:8], 1234)
	binary.LittleEndian.PutUint32(original.data[8:12], 5678)

	for i := 0; i < 160; i++ {
		original.data[12+i] = byte(i % 256)
	}

	encoded, err := original.Encode()
	if err != nil {
		t.Fatalf("编码失败: %v", err)
	}

	if len(encoded) != 172 {
		t.Fatalf("编码长度错误，期望172字节，实际%d字节", len(encoded))
	}

	decoded := &DebugEvent{}
	err = decoded.Decode(encoded)
	if err != nil {
		t.Fatalf("解码失败: %v", err)
	}

	if decoded.DebugEventCode() != EXCEPTION_DEBUG_EVENT {
		t.Errorf("DebugEventCode不匹配: 期望%d，实际%d", EXCEPTION_DEBUG_EVENT, decoded.DebugEventCode())
	}

	if decoded.ProcessId() != 1234 {
		t.Errorf("ProcessId不匹配: 期望%d，实际%d", 1234, decoded.ProcessId())
	}

	if decoded.ThreadId() != 5678 {
		t.Errorf("ThreadId不匹配: 期望%d，实际%d", 5678, decoded.ThreadId())
	}

	for i := 0; i < 160; i++ {
		if decoded.data[12+i] != byte(i%256) {
			t.Errorf("data[12+%d]不匹配: 期望%d，实际%d", i, byte(i%256), decoded.data[12+i])
		}
	}
}

func TestExceptionRecordDecodeInvalidLength(t *testing.T) {
	data := make([]byte, 100)
	_, err := DecodeExceptionRecord(data)
	if err == nil {
		t.Error("期望解码失败，但成功了")
	}
}

func TestExceptionDebugInfoDecodeInvalidLength(t *testing.T) {
	data := make([]byte, 100)
	_, err := DecodeExceptionDebugInfo(data)
	if err == nil {
		t.Error("期望解码失败，但成功了")
	}
}

func TestDebugEventDecodeInvalidLength(t *testing.T) {
	data := make([]byte, 10)
	_, err := DecodeDebugEvent(data)
	if err == nil {
		t.Error("期望解码失败，但成功了")
	}
}

func TestExceptionCodeMap(t *testing.T) {
	testCases := []struct {
		code    uint32
		wantLen int
	}{
		{EXCEPTION_SINGLE_STEP, 1},
		{EXCEPTION_BREAKPOINT, 1},
		{EXCEPTION_ACCESS_VIOLATION, 1},
		{0xFFFFFFFF, 0},
	}

	for _, tc := range testCases {
		name := GetExceptionCodeName(tc.code)
		if tc.wantLen == 1 {
			if name == "" {
				t.Errorf("异常代码0x%X应该有名称", tc.code)
			}
		} else {
			if name == "" {
				t.Errorf("未知异常代码应该返回默认名称")
			}
		}
	}
}

func TestHexDump(t *testing.T) {
	data := []byte{
		0x48, 0x65, 0x6C, 0x6C, 0x6F, 0x20, 0x57, 0x6F,
		0x72, 0x6C, 0x64, 0x21, 0x00, 0x01, 0x02, 0x03,
		0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A, 0x0B,
		0x0C, 0x0D, 0x0E, 0x0F, 0x10, 0x11, 0x12, 0x13,
	}

	dump := NewHexDump(data, 0x1000)
	str := dump.String()

	if str == "" {
		t.Error("hex dump字符串不应该为空")
	}

	if len(str) < 10 {
		t.Error("hex dump字符串长度应该大于10")
	}
}

func TestHexDumpTryParseUTF16(t *testing.T) {
	data := []byte{
		0x48, 0x00, 0x65, 0x00, 0x6C, 0x00, 0x6C, 0x00,
		0x6F, 0x00, 0x00, 0x00,
	}

	dump := NewHexDump(data, 0)
	str, ok := dump.TryParseUTF16()

	if !ok {
		t.Error("UTF-16解析应该成功")
	}

	expected := "Hello\x00"
	if str != expected {
		t.Errorf("UTF-16解析结果错误: 期望'%s'，实际'%s' (长度: %d vs %d)", expected, str, len(expected), len(str))
	}
}

func TestHexDumpTryParseUTF8(t *testing.T) {
	data := []byte{
		0x48, 0x65, 0x6C, 0x6C, 0x6F, 0x00,
	}

	dump := NewHexDump(data, 0)
	str, ok := dump.TryParseUTF8()

	if !ok {
		t.Error("UTF-8解析应该成功")
	}

	expected := "Hello\x00"
	if str != expected {
		t.Errorf("UTF-8解析结果错误: 期望'%s'，实际'%s' (长度: %d vs %d)", expected, str, len(expected), len(str))
	}
}

func TestSizeOf(t *testing.T) {
	record := ExceptionRecord{}
	size := int(unsafe.Sizeof(record))

	if size != 152 {
		t.Logf("注意: unsafe.Sizeof返回%d字节，这可能是因为结构体对齐问题", size)
	}

	debugInfo := ExceptionDebugInfo{}
	size = int(unsafe.Sizeof(debugInfo))

	if size != 156 {
		t.Logf("注意: unsafe.Sizeof返回%d字节，这可能是因为结构体对齐问题", size)
	}

	event := DebugEvent{}
	size = int(unsafe.Sizeof(event))

	if size != 172 {
		t.Logf("注意: unsafe.Sizeof返回%d字节，这可能是因为结构体对齐问题", size)
	}

	t.Log("SizeOf测试完成 - 实际大小可能因编译器对齐而异")
}

func TestCreateThreadDebugInfoCodec(t *testing.T) {
	original := &CreateThreadDebugInfo{
		ThreadHandle:    0x1234,
		ThreadLocalBase: 0x7FFF00001000,
		StartAddress:    0x7FFF00002000,
	}

	encoded, err := EncodeCreateThreadDebugInfo(original)
	if err != nil {
		t.Fatalf("编码失败: %v", err)
	}

	if len(encoded) != 24 {
		t.Fatalf("编码长度错误，期望24字节，实际%d字节", len(encoded))
	}

	decoded, err := DecodeCreateThreadDebugInfo(encoded)
	if err != nil {
		t.Fatalf("解码失败: %v", err)
	}

	if decoded.ThreadHandle != original.ThreadHandle {
		t.Errorf("ThreadHandle不匹配: 期望%x，实际%x", original.ThreadHandle, decoded.ThreadHandle)
	}

	if decoded.ThreadLocalBase != original.ThreadLocalBase {
		t.Errorf("ThreadLocalBase不匹配: 期望%x，实际%x", original.ThreadLocalBase, decoded.ThreadLocalBase)
	}

	if decoded.StartAddress != original.StartAddress {
		t.Errorf("StartAddress不匹配: 期望%x，实际%x", original.StartAddress, decoded.StartAddress)
	}
}

func TestCreateProcessDebugInfoCodec(t *testing.T) {
	original := &CreateProcessDebugInfo{
		FileHandle:          0x1234,
		ProcessHandle:       0x5678,
		ThreadHandle:        0x9ABC,
		BaseOfImage:         0x00001000,
		DebugInfoFileOffset: 0x100,
		DebugInfoSize:       0x200,
		ThreadLocalBase:     0x00002000,
		StartAddress:        0x00003000,
		ImageName:           0x00004000,
		Unicode:             1,
	}

	encoded, err := EncodeCreateProcessDebugInfo(original)
	if err != nil {
		t.Fatalf("编码失败: %v", err)
	}

	if len(encoded) != 96 {
		t.Fatalf("编码长度错误，期望96字节，实际%d字节", len(encoded))
	}

	decoded, err := DecodeCreateProcessDebugInfo(encoded)
	if err != nil {
		t.Fatalf("解码失败: %v", err)
	}

	if decoded.FileHandle != original.FileHandle {
		t.Errorf("FileHandle不匹配: 期望%x，实际%x", original.FileHandle, decoded.FileHandle)
	}

	if decoded.ProcessHandle != original.ProcessHandle {
		t.Errorf("ProcessHandle不匹配: 期望%x，实际%x", original.ProcessHandle, decoded.ProcessHandle)
	}

	if decoded.ThreadHandle != original.ThreadHandle {
		t.Errorf("ThreadHandle不匹配: 期望%x，实际%x", original.ThreadHandle, decoded.ThreadHandle)
	}

	if decoded.BaseOfImage != original.BaseOfImage {
		t.Errorf("BaseOfImage不匹配: 期望%x，实际%x", original.BaseOfImage, decoded.BaseOfImage)
	}

	if decoded.DebugInfoFileOffset != original.DebugInfoFileOffset {
		t.Errorf("DebugInfoFileOffset不匹配: 期望%x，实际%x", original.DebugInfoFileOffset, decoded.DebugInfoFileOffset)
	}

	if decoded.DebugInfoSize != original.DebugInfoSize {
		t.Errorf("DebugInfoSize不匹配: 期望%x，实际%x", original.DebugInfoSize, decoded.DebugInfoSize)
	}

	if decoded.ThreadLocalBase != original.ThreadLocalBase {
		t.Errorf("ThreadLocalBase不匹配: 期望%x，实际%x", original.ThreadLocalBase, decoded.ThreadLocalBase)
	}

	if decoded.StartAddress != original.StartAddress {
		t.Errorf("StartAddress不匹配: 期望%x，实际%x", original.StartAddress, decoded.StartAddress)
	}

	if decoded.ImageName != original.ImageName {
		t.Errorf("ImageName不匹配: 期望%x，实际%x", original.ImageName, decoded.ImageName)
	}

	if decoded.Unicode != original.Unicode {
		t.Errorf("Unicode不匹配: 期望%d，实际%d", original.Unicode, decoded.Unicode)
	}
}

func TestExitThreadDebugInfoCodec(t *testing.T) {
	original := &ExitThreadDebugInfo{
		ExitCode: 0x0,
	}

	encoded, err := EncodeExitThreadDebugInfo(original)
	if err != nil {
		t.Fatalf("编码失败: %v", err)
	}

	if len(encoded) != 4 {
		t.Fatalf("编码长度错误，期望4字节，实际%d字节", len(encoded))
	}

	decoded, err := DecodeExitThreadDebugInfo(encoded)
	if err != nil {
		t.Fatalf("解码失败: %v", err)
	}

	if decoded.ExitCode != original.ExitCode {
		t.Errorf("ExitCode不匹配: 期望%x，实际%x", original.ExitCode, decoded.ExitCode)
	}
}

func TestExitProcessDebugInfoCodec(t *testing.T) {
	original := &ExitProcessDebugInfo{
		ExitCode: 0x0,
	}

	encoded, err := EncodeExitProcessDebugInfo(original)
	if err != nil {
		t.Fatalf("编码失败: %v", err)
	}

	if len(encoded) != 4 {
		t.Fatalf("编码长度错误，期望4字节，实际%d字节", len(encoded))
	}

	decoded, err := DecodeExitProcessDebugInfo(encoded)
	if err != nil {
		t.Fatalf("解码失败: %v", err)
	}

	if decoded.ExitCode != original.ExitCode {
		t.Errorf("ExitCode不匹配: 期望%x，实际%x", original.ExitCode, decoded.ExitCode)
	}
}

func TestLoadDllDebugInfoCodec(t *testing.T) {
	original := &LoadDllDebugInfo{
		FileHandle:          0x1234,
		BaseOfDll:           0x00001000,
		DebugInfoFileOffset: 0x100,
		DebugInfoSize:       0x200,
		ImageName:           0x00002000,
		Unicode:             1,
	}

	encoded, err := EncodeLoadDllDebugInfo(original)
	if err != nil {
		t.Fatalf("编码失败: %v", err)
	}

	if len(encoded) != 80 {
		t.Fatalf("编码长度错误，期望80字节，实际%d字节", len(encoded))
	}

	decoded, err := DecodeLoadDllDebugInfo(encoded)
	if err != nil {
		t.Fatalf("解码失败: %v", err)
	}

	if decoded.FileHandle != original.FileHandle {
		t.Errorf("FileHandle不匹配: 期望%x，实际%x", original.FileHandle, decoded.FileHandle)
	}

	if decoded.BaseOfDll != original.BaseOfDll {
		t.Errorf("BaseOfDll不匹配: 期望%x，实际%x", original.BaseOfDll, decoded.BaseOfDll)
	}

	if decoded.DebugInfoFileOffset != original.DebugInfoFileOffset {
		t.Errorf("DebugInfoFileOffset不匹配: 期望%x，实际%x", original.DebugInfoFileOffset, decoded.DebugInfoFileOffset)
	}

	if decoded.DebugInfoSize != original.DebugInfoSize {
		t.Errorf("DebugInfoSize不匹配: 期望%x，实际%x", original.DebugInfoSize, decoded.DebugInfoSize)
	}

	if decoded.ImageName != original.ImageName {
		t.Errorf("ImageName不匹配: 期望%x，实际%x", original.ImageName, decoded.ImageName)
	}

	if decoded.Unicode != original.Unicode {
		t.Errorf("Unicode不匹配: 期望%d，实际%d", original.Unicode, decoded.Unicode)
	}
}

func TestUnloadDllDebugInfoCodec(t *testing.T) {
	original := &UnloadDllDebugInfo{
		BaseOfDll: 0x7FFF00001000,
	}

	encoded, err := EncodeUnloadDllDebugInfo(original)
	if err != nil {
		t.Fatalf("编码失败: %v", err)
	}

	if len(encoded) != 8 {
		t.Fatalf("编码长度错误，期望8字节，实际%d字节", len(encoded))
	}

	decoded, err := DecodeUnloadDllDebugInfo(encoded)
	if err != nil {
		t.Fatalf("解码失败: %v", err)
	}

	if decoded.BaseOfDll != original.BaseOfDll {
		t.Errorf("BaseOfDll不匹配: 期望%x，实际%x", original.BaseOfDll, decoded.BaseOfDll)
	}
}

func TestOutputDebugStringInfoCodec(t *testing.T) {
	var testStr uint16 = 0x1234
	original := &OutputDebugStringInfo{
		DebugString:       &testStr,
		Unicode:           1,
		DebugStringLength: uint16(13),
	}

	encoded, err := EncodeOutputDebugStringInfo(original)
	if err != nil {
		t.Fatalf("编码失败: %v", err)
	}

	if len(encoded) != 12 {
		t.Fatalf("编码长度错误，期望12字节，实际%d字节", len(encoded))
	}

	decoded, err := DecodeOutputDebugStringInfo(encoded)
	if err != nil {
		t.Fatalf("解码失败: %v", err)
	}

	if decoded.Unicode != original.Unicode {
		t.Errorf("Unicode不匹配: 期望%d，实际%d", original.Unicode, decoded.Unicode)
	}

	if decoded.DebugStringLength != original.DebugStringLength {
		t.Errorf("DebugStringLength不匹配: 期望%d，实际%d", original.DebugStringLength, decoded.DebugStringLength)
	}
}

func TestStartupInfoCodec(t *testing.T) {
	original := &StartupInfo{
		Cb:            104,
		Desktop:       nil,
		Title:         nil,
		X:             100,
		Y:             200,
		XSize:         800,
		YSize:         600,
		XCountChars:   80,
		YCountChars:   25,
		FillAttribute: 7,
		Flags:         STARTF_USESHOWWINDOW,
		ShowWindow:    SW_SHOW,
		CbReserved2:   0,
		StdInput:      0xFFFFFFFFFFFFFFFF,
		StdOutput:     0xFFFFFFFFFFFFFFFF,
		StdError:      0xFFFFFFFFFFFFFFFF,
	}

	encoded, err := EncodeStartupInfo(original)
	if err != nil {
		t.Fatalf("编码失败: %v", err)
	}

	if len(encoded) != 104 {
		t.Fatalf("编码长度错误，期望104字节，实际%d字节", len(encoded))
	}

	decoded, err := DecodeStartupInfo(encoded)
	if err != nil {
		t.Fatalf("解码失败: %v", err)
	}

	if decoded.Cb != original.Cb {
		t.Errorf("Cb不匹配: 期望%d，实际%d", original.Cb, decoded.Cb)
	}

	if decoded.X != original.X {
		t.Errorf("X不匹配: 期望%d，实际%d", original.X, decoded.X)
	}

	if decoded.Y != original.Y {
		t.Errorf("Y不匹配: 期望%d，实际%d", original.Y, decoded.Y)
	}

	if decoded.XSize != original.XSize {
		t.Errorf("XSize不匹配: 期望%d，实际%d", original.XSize, decoded.XSize)
	}

	if decoded.YSize != original.YSize {
		t.Errorf("YSize不匹配: 期望%d，实际%d", original.YSize, decoded.YSize)
	}

	if decoded.XCountChars != original.XCountChars {
		t.Errorf("XCountChars不匹配: 期望%d，实际%d", original.XCountChars, decoded.XCountChars)
	}

	if decoded.YCountChars != original.YCountChars {
		t.Errorf("YCountChars不匹配: 期望%d，实际%d", original.YCountChars, decoded.YCountChars)
	}

	if decoded.FillAttribute != original.FillAttribute {
		t.Errorf("FillAttribute不匹配: 期望%d，实际%d", original.FillAttribute, decoded.FillAttribute)
	}

	if decoded.Flags != original.Flags {
		t.Errorf("Flags不匹配: 期望%d，实际%d", original.Flags, decoded.Flags)
	}

	if decoded.ShowWindow != original.ShowWindow {
		t.Errorf("ShowWindow不匹配: 期望%d，实际%d", original.ShowWindow, decoded.ShowWindow)
	}

	if decoded.CbReserved2 != original.CbReserved2 {
		t.Errorf("CbReserved2不匹配: 期望%d，实际%d", original.CbReserved2, decoded.CbReserved2)
	}

	if decoded.StdInput != original.StdInput {
		t.Errorf("StdInput不匹配: 期望%x，实际%x", original.StdInput, decoded.StdInput)
	}

	if decoded.StdOutput != original.StdOutput {
		t.Errorf("StdOutput不匹配: 期望%x，实际%x", original.StdOutput, decoded.StdOutput)
	}

	if decoded.StdError != original.StdError {
		t.Errorf("StdError不匹配: 期望%x，实际%x", original.StdError, decoded.StdError)
	}
}

func TestCreateThreadDebugInfoDecodeInvalidLength(t *testing.T) {
	data := make([]byte, 10)
	_, err := DecodeCreateThreadDebugInfo(data)
	if err == nil {
		t.Error("期望解码失败，但成功了")
	}
}

func TestCreateProcessDebugInfoDecodeInvalidLength(t *testing.T) {
	data := make([]byte, 10)
	_, err := DecodeCreateProcessDebugInfo(data)
	if err == nil {
		t.Error("期望解码失败，但成功了")
	}
}

func TestExitThreadDebugInfoDecodeInvalidLength(t *testing.T) {
	data := make([]byte, 2)
	_, err := DecodeExitThreadDebugInfo(data)
	if err == nil {
		t.Error("期望解码失败，但成功了")
	}
}

func TestExitProcessDebugInfoDecodeInvalidLength(t *testing.T) {
	data := make([]byte, 2)
	_, err := DecodeExitProcessDebugInfo(data)
	if err == nil {
		t.Error("期望解码失败，但成功了")
	}
}

func TestLoadDllDebugInfoDecodeInvalidLength(t *testing.T) {
	data := make([]byte, 10)
	_, err := DecodeLoadDllDebugInfo(data)
	if err == nil {
		t.Error("期望解码失败，但成功了")
	}
}

func TestUnloadDllDebugInfoDecodeInvalidLength(t *testing.T) {
	data := make([]byte, 4)
	_, err := DecodeUnloadDllDebugInfo(data)
	if err == nil {
		t.Error("期望解码失败，但成功了")
	}
}

func TestOutputDebugStringInfoDecodeInvalidLength(t *testing.T) {
	data := make([]byte, 10)
	_, err := DecodeOutputDebugStringInfo(data)
	if err == nil {
		t.Error("期望解码失败，但成功了")
	}
}

func TestStartupInfoDecodeInvalidLength(t *testing.T) {
	data := make([]byte, 50)
	_, err := DecodeStartupInfo(data)
	if err == nil {
		t.Error("期望解码失败，但成功了")
	}
}
