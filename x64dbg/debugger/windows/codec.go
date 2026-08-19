package windows

import (
	"encoding/binary"
	"fmt"
	"unsafe"
)

// Encoder 编码器接口
type Encoder interface {
	Encode() ([]byte, error)
}

// Decoder 解码器接口
type Decoder interface {
	Decode(data []byte) error
}

// EncodeExceptionRecord 编码ExceptionRecord为字节数组
func EncodeExceptionRecord(record *ExceptionRecord) ([]byte, error) {
	buf := make([]byte, 152)

	binary.LittleEndian.PutUint32(buf[0:4], record.ExceptionCode)
	binary.LittleEndian.PutUint32(buf[4:8], record.ExceptionFlags)
	binary.LittleEndian.PutUint64(buf[8:16], record.ExceptionRecord)
	binary.LittleEndian.PutUint64(buf[16:24], record.ExceptionAddress)
	binary.LittleEndian.PutUint32(buf[24:28], record.NumberParameters)
	binary.LittleEndian.PutUint32(buf[28:32], record._UnusedAlignment)

	for i := 0; i < 15; i++ {
		binary.LittleEndian.PutUint64(buf[32+i*8:40+i*8], record.ExceptionInformation[i])
	}

	return buf, nil
}

// DecodeExceptionRecord 从字节数组解码ExceptionRecord
func DecodeExceptionRecord(data []byte) (*ExceptionRecord, error) {
	if len(data) < 152 {
		return nil, fmt.Errorf("数据长度不足，需要152字节，实际%d字节", len(data))
	}

	record := &ExceptionRecord{
		ExceptionCode:    binary.LittleEndian.Uint32(data[0:4]),
		ExceptionFlags:   binary.LittleEndian.Uint32(data[4:8]),
		ExceptionRecord:  binary.LittleEndian.Uint64(data[8:16]),
		ExceptionAddress: binary.LittleEndian.Uint64(data[16:24]),
		NumberParameters: binary.LittleEndian.Uint32(data[24:28]),
		_UnusedAlignment: binary.LittleEndian.Uint32(data[28:32]),
	}

	for i := 0; i < 15; i++ {
		record.ExceptionInformation[i] = binary.LittleEndian.Uint64(data[32+i*8 : 40+i*8])
	}

	return record, nil
}

// EncodeExceptionDebugInfo 编码ExceptionDebugInfo为字节数组
func EncodeExceptionDebugInfo(info *ExceptionDebugInfo) ([]byte, error) {
	recordData, err := EncodeExceptionRecord(&info.ExceptionRecord)
	if err != nil {
		return nil, err
	}

	buf := make([]byte, 156)
	copy(buf[:152], recordData)
	binary.LittleEndian.PutUint32(buf[152:156], info.FirstChance)

	return buf, nil
}

// DecodeExceptionDebugInfo 从字节数组解码ExceptionDebugInfo
func DecodeExceptionDebugInfo(data []byte) (*ExceptionDebugInfo, error) {
	if len(data) < 156 {
		return nil, fmt.Errorf("数据长度不足，需要156字节，实际%d字节", len(data))
	}

	record, err := DecodeExceptionRecord(data[:152])
	if err != nil {
		return nil, err
	}

	return &ExceptionDebugInfo{
		ExceptionRecord: *record,
		FirstChance:     binary.LittleEndian.Uint32(data[152:156]),
	}, nil
}

// EncodeCreateThreadDebugInfo 编码CreateThreadDebugInfo为字节数组
func EncodeCreateThreadDebugInfo(info *CreateThreadDebugInfo) ([]byte, error) {
	buf := make([]byte, 24)

	binary.LittleEndian.PutUint64(buf[0:8], uint64(info.ThreadHandle))
	binary.LittleEndian.PutUint64(buf[8:16], info.ThreadLocalBase)
	binary.LittleEndian.PutUint64(buf[16:24], info.StartAddress)

	return buf, nil
}

// DecodeCreateThreadDebugInfo 从字节数组解码CreateThreadDebugInfo
func DecodeCreateThreadDebugInfo(data []byte) (*CreateThreadDebugInfo, error) {
	if len(data) < 24 {
		return nil, fmt.Errorf("数据长度不足，需要24字节，实际%d字节", len(data))
	}

	return &CreateThreadDebugInfo{
		ThreadHandle:    Handle(binary.LittleEndian.Uint64(data[0:8])),
		ThreadLocalBase: binary.LittleEndian.Uint64(data[8:16]),
		StartAddress:    binary.LittleEndian.Uint64(data[16:24]),
	}, nil
}

// EncodeCreateProcessDebugInfo 编码CreateProcessDebugInfo为字节数组
func EncodeCreateProcessDebugInfo(info *CreateProcessDebugInfo) ([]byte, error) {
	buf := make([]byte, 96)

	binary.LittleEndian.PutUint32(buf[0:4], info.FileHandle)
	binary.LittleEndian.PutUint32(buf[8:12], info.ProcessHandle)
	binary.LittleEndian.PutUint32(buf[16:20], info.ThreadHandle)
	binary.LittleEndian.PutUint32(buf[24:28], info.BaseOfImage)
	binary.LittleEndian.PutUint32(buf[28:32], info.DebugInfoFileOffset)
	binary.LittleEndian.PutUint32(buf[32:36], info.DebugInfoSize)
	binary.LittleEndian.PutUint32(buf[36:40], info.ThreadLocalBase)
	binary.LittleEndian.PutUint32(buf[40:44], info.StartAddress)
	binary.LittleEndian.PutUint32(buf[44:48], info.ImageName)
	binary.LittleEndian.PutUint16(buf[52:54], info.Unicode)

	return buf, nil
}

// DecodeCreateProcessDebugInfo 从字节数组解码CreateProcessDebugInfo
func DecodeCreateProcessDebugInfo(data []byte) (*CreateProcessDebugInfo, error) {
	if len(data) < 96 {
		return nil, fmt.Errorf("数据长度不足，需要96字节，实际%d字节", len(data))
	}

	return &CreateProcessDebugInfo{
		FileHandle:          binary.LittleEndian.Uint32(data[0:4]),
		ProcessHandle:       binary.LittleEndian.Uint32(data[8:12]),
		ThreadHandle:        binary.LittleEndian.Uint32(data[16:20]),
		BaseOfImage:         binary.LittleEndian.Uint32(data[24:28]),
		DebugInfoFileOffset: binary.LittleEndian.Uint32(data[28:32]),
		DebugInfoSize:       binary.LittleEndian.Uint32(data[32:36]),
		ThreadLocalBase:     binary.LittleEndian.Uint32(data[36:40]),
		StartAddress:        binary.LittleEndian.Uint32(data[40:44]),
		ImageName:           binary.LittleEndian.Uint32(data[44:48]),
		Unicode:             binary.LittleEndian.Uint16(data[52:54]),
	}, nil
}

// EncodeExitThreadDebugInfo 编码ExitThreadDebugInfo为字节数组
func EncodeExitThreadDebugInfo(info *ExitThreadDebugInfo) ([]byte, error) {
	buf := make([]byte, 4)
	binary.LittleEndian.PutUint32(buf[0:4], info.ExitCode)
	return buf, nil
}

// DecodeExitThreadDebugInfo 从字节数组解码ExitThreadDebugInfo
func DecodeExitThreadDebugInfo(data []byte) (*ExitThreadDebugInfo, error) {
	if len(data) < 4 {
		return nil, fmt.Errorf("数据长度不足，需要4字节，实际%d字节", len(data))
	}

	return &ExitThreadDebugInfo{
		ExitCode: binary.LittleEndian.Uint32(data[0:4]),
	}, nil
}

// EncodeExitProcessDebugInfo 编码ExitProcessDebugInfo为字节数组
func EncodeExitProcessDebugInfo(info *ExitProcessDebugInfo) ([]byte, error) {
	buf := make([]byte, 4)
	binary.LittleEndian.PutUint32(buf[0:4], info.ExitCode)
	return buf, nil
}

// DecodeExitProcessDebugInfo 从字节数组解码ExitProcessDebugInfo
func DecodeExitProcessDebugInfo(data []byte) (*ExitProcessDebugInfo, error) {
	if len(data) < 4 {
		return nil, fmt.Errorf("数据长度不足，需要4字节，实际%d字节", len(data))
	}

	return &ExitProcessDebugInfo{
		ExitCode: binary.LittleEndian.Uint32(data[0:4]),
	}, nil
}

// EncodeLoadDllDebugInfo 编码LoadDllDebugInfo为字节数组
func EncodeLoadDllDebugInfo(info *LoadDllDebugInfo) ([]byte, error) {
	buf := make([]byte, 80)

	binary.LittleEndian.PutUint64(buf[0:8], uint64(info.FileHandle))
	binary.LittleEndian.PutUint64(buf[8:16], info.BaseOfDll)
	binary.LittleEndian.PutUint32(buf[16:20], info.DebugInfoFileOffset)
	binary.LittleEndian.PutUint32(buf[20:24], info.DebugInfoSize)
	binary.LittleEndian.PutUint16(buf[24:26], info.ImageName)
	binary.LittleEndian.PutUint16(buf[26:28], info.Unicode)

	return buf, nil
}

// DecodeLoadDllDebugInfo 从字节数组解码LoadDllDebugInfo
func DecodeLoadDllDebugInfo(data []byte) (*LoadDllDebugInfo, error) {
	if len(data) < 80 {
		return nil, fmt.Errorf("数据长度不足，需要80字节，实际%d字节", len(data))
	}

	return &LoadDllDebugInfo{
		FileHandle:          Handle(binary.LittleEndian.Uint64(data[0:8])),
		BaseOfDll:           binary.LittleEndian.Uint64(data[8:16]),
		DebugInfoFileOffset: binary.LittleEndian.Uint32(data[16:20]),
		DebugInfoSize:       binary.LittleEndian.Uint32(data[20:24]),
		ImageName:           binary.LittleEndian.Uint16(data[24:26]),
		Unicode:             binary.LittleEndian.Uint16(data[26:28]),
	}, nil
}

// EncodeUnloadDllDebugInfo 编码UnloadDllDebugInfo为字节数组
func EncodeUnloadDllDebugInfo(info *UnloadDllDebugInfo) ([]byte, error) {
	buf := make([]byte, 8)
	binary.LittleEndian.PutUint64(buf[0:8], info.BaseOfDll)
	return buf, nil
}

// DecodeUnloadDllDebugInfo 从字节数组解码UnloadDllDebugInfo
func DecodeUnloadDllDebugInfo(data []byte) (*UnloadDllDebugInfo, error) {
	if len(data) < 8 {
		return nil, fmt.Errorf("数据长度不足，需要8字节，实际%d字节", len(data))
	}

	return &UnloadDllDebugInfo{
		BaseOfDll: binary.LittleEndian.Uint64(data[0:8]),
	}, nil
}

// EncodeOutputDebugStringInfo 编码OutputDebugStringInfo为字节数组
func EncodeOutputDebugStringInfo(info *OutputDebugStringInfo) ([]byte, error) {
	buf := make([]byte, 12)

	binary.LittleEndian.PutUint64(buf[0:8], uint64(uintptr(unsafe.Pointer(info.DebugString))))
	binary.LittleEndian.PutUint16(buf[8:10], info.Unicode)
	binary.LittleEndian.PutUint16(buf[10:12], info.DebugStringLength)

	return buf, nil
}

// DecodeOutputDebugStringInfo 从字节数组解码OutputDebugStringInfo
func DecodeOutputDebugStringInfo(data []byte) (*OutputDebugStringInfo, error) {
	if len(data) < 12 {
		return nil, fmt.Errorf("数据长度不足，需要12字节，实际%d字节", len(data))
	}

	return &OutputDebugStringInfo{
		DebugString:       (*uint16)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(data[0:8])))),
		Unicode:           binary.LittleEndian.Uint16(data[8:10]),
		DebugStringLength: binary.LittleEndian.Uint16(data[10:12]),
	}, nil
}

// EncodeM128A 编码M128A为字节数组
func EncodeM128A(m128a *M128A) ([]byte, error) {
	buf := make([]byte, 16)

	binary.LittleEndian.PutUint64(buf[0:8], m128a.Low)
	binary.LittleEndian.PutUint64(buf[8:16], uint64(m128a.High))

	return buf, nil
}

// DecodeM128A 从字节数组解码M128A
func DecodeM128A(data []byte) (*M128A, error) {
	if len(data) < 16 {
		return nil, fmt.Errorf("数据长度不足，需要16字节，实际%d字节", len(data))
	}

	return &M128A{
		Low:  binary.LittleEndian.Uint64(data[0:8]),
		High: int64(binary.LittleEndian.Uint64(data[8:16])),
	}, nil
}

// EncodeFloatingSaveArea 编码FloatingSaveArea为字节数组
func EncodeFloatingSaveArea(area *FloatingSaveArea) ([]byte, error) {
	buf := make([]byte, 112)

	binary.LittleEndian.PutUint32(buf[0:4], area.ControlWord)
	binary.LittleEndian.PutUint32(buf[4:8], area.StatusWord)
	binary.LittleEndian.PutUint32(buf[8:12], area.TagWord)
	binary.LittleEndian.PutUint32(buf[12:16], area.ErrorOffset)
	binary.LittleEndian.PutUint32(buf[16:20], area.ErrorSelector)
	binary.LittleEndian.PutUint32(buf[20:24], area.DataOffset)
	binary.LittleEndian.PutUint32(buf[24:28], area.DataSelector)
	copy(buf[28:108], area.RegisterArea[:])
	binary.LittleEndian.PutUint32(buf[108:112], area.Cr0NpxState)

	return buf, nil
}

// DecodeFloatingSaveArea 从字节数组解码FloatingSaveArea
func DecodeFloatingSaveArea(data []byte) (*FloatingSaveArea, error) {
	if len(data) < 112 {
		return nil, fmt.Errorf("数据长度不足，需要112字节，实际%d字节", len(data))
	}

	area := &FloatingSaveArea{
		ControlWord:   binary.LittleEndian.Uint32(data[0:4]),
		StatusWord:    binary.LittleEndian.Uint32(data[4:8]),
		TagWord:       binary.LittleEndian.Uint32(data[8:12]),
		ErrorOffset:   binary.LittleEndian.Uint32(data[12:16]),
		ErrorSelector: binary.LittleEndian.Uint32(data[16:20]),
		DataOffset:    binary.LittleEndian.Uint32(data[20:24]),
		DataSelector:  binary.LittleEndian.Uint32(data[24:28]),
		Cr0NpxState:   binary.LittleEndian.Uint32(data[108:112]),
	}
	copy(area.RegisterArea[:], data[28:108])

	return area, nil
}

// EncodeContext 编码Context为字节数组
func EncodeContext(ctx *Context) ([]byte, error) {
	buf := make([]byte, 1232)

	binary.LittleEndian.PutUint64(buf[0:8], ctx.P1Home)
	binary.LittleEndian.PutUint64(buf[8:16], ctx.P2Home)
	binary.LittleEndian.PutUint64(buf[16:24], ctx.P3Home)
	binary.LittleEndian.PutUint64(buf[24:32], ctx.P4Home)
	binary.LittleEndian.PutUint64(buf[32:40], ctx.P5Home)
	binary.LittleEndian.PutUint64(buf[40:48], ctx.P6Home)
	binary.LittleEndian.PutUint32(buf[48:52], ctx.ContextFlags)
	binary.LittleEndian.PutUint32(buf[52:56], ctx.MxCsr)
	binary.LittleEndian.PutUint16(buf[56:58], ctx.SegCs)
	binary.LittleEndian.PutUint16(buf[58:60], ctx.SegDs)
	binary.LittleEndian.PutUint16(buf[60:62], ctx.SegEs)
	binary.LittleEndian.PutUint16(buf[62:64], ctx.SegFs)
	binary.LittleEndian.PutUint16(buf[64:66], ctx.SegGs)
	binary.LittleEndian.PutUint16(buf[66:68], ctx.SegSs)
	binary.LittleEndian.PutUint32(buf[68:72], ctx.EFlags)

	for i := 0; i < 8; i++ {
		offset := 72 + i*8
		switch i {
		case 0:
			binary.LittleEndian.PutUint64(buf[offset:offset+8], ctx.Dr0)
		case 1:
			binary.LittleEndian.PutUint64(buf[offset:offset+8], ctx.Dr1)
		case 2:
			binary.LittleEndian.PutUint64(buf[offset:offset+8], ctx.Dr2)
		case 3:
			binary.LittleEndian.PutUint64(buf[offset:offset+8], ctx.Dr3)
		case 4:
			binary.LittleEndian.PutUint64(buf[offset:offset+8], ctx.Dr6)
		case 5:
			binary.LittleEndian.PutUint64(buf[offset:offset+8], ctx.Dr7)
		}
	}

	regs := []uint64{
		ctx.Rax, ctx.Rcx, ctx.Rdx, ctx.Rbx, ctx.Rsp, ctx.Rbp, ctx.Rsi, ctx.Rdi,
		ctx.R8, ctx.R9, ctx.R10, ctx.R11, ctx.R12, ctx.R13, ctx.R14, ctx.R15, ctx.Rip,
	}
	for i, reg := range regs {
		offset := 136 + i*8
		binary.LittleEndian.PutUint64(buf[offset:offset+8], reg)
	}

	fltSaveData, err := EncodeFloatingSaveArea(&ctx.FltSave)
	if err != nil {
		return nil, err
	}
	copy(buf[272:384], fltSaveData)

	for i := 0; i < 26; i++ {
		offset := 384 + i*16
		m128aData, err := EncodeM128A(&ctx.VectorRegister[i])
		if err != nil {
			return nil, err
		}
		copy(buf[offset:offset+16], m128aData)
	}

	binary.LittleEndian.PutUint64(buf[800:808], ctx.VectorControl)
	binary.LittleEndian.PutUint64(buf[808:816], ctx.DebugControl)
	binary.LittleEndian.PutUint64(buf[816:824], ctx.LastBranchToRip)
	binary.LittleEndian.PutUint64(buf[824:832], ctx.LastBranchFromRip)
	binary.LittleEndian.PutUint64(buf[832:840], ctx.LastExceptionToRip)
	binary.LittleEndian.PutUint64(buf[840:848], ctx.LastExceptionFromRip)

	return buf, nil
}

// DecodeContext 从字节数组解码Context
func DecodeContext(data []byte) (*Context, error) {
	if len(data) < 848 {
		return nil, fmt.Errorf("数据长度不足，需要848字节，实际%d字节", len(data))
	}

	ctx := &Context{
		P1Home:       binary.LittleEndian.Uint64(data[0:8]),
		P2Home:       binary.LittleEndian.Uint64(data[8:16]),
		P3Home:       binary.LittleEndian.Uint64(data[16:24]),
		P4Home:       binary.LittleEndian.Uint64(data[24:32]),
		P5Home:       binary.LittleEndian.Uint64(data[32:40]),
		P6Home:       binary.LittleEndian.Uint64(data[40:48]),
		ContextFlags: binary.LittleEndian.Uint32(data[48:52]),
		MxCsr:        binary.LittleEndian.Uint32(data[52:56]),
		SegCs:        binary.LittleEndian.Uint16(data[56:58]),
		SegDs:        binary.LittleEndian.Uint16(data[58:60]),
		SegEs:        binary.LittleEndian.Uint16(data[60:62]),
		SegFs:        binary.LittleEndian.Uint16(data[62:64]),
		SegGs:        binary.LittleEndian.Uint16(data[64:66]),
		SegSs:        binary.LittleEndian.Uint16(data[66:68]),
		EFlags:       binary.LittleEndian.Uint32(data[68:72]),
		Dr0:          binary.LittleEndian.Uint64(data[72:80]),
		Dr1:          binary.LittleEndian.Uint64(data[80:88]),
		Dr2:          binary.LittleEndian.Uint64(data[88:96]),
		Dr3:          binary.LittleEndian.Uint64(data[96:104]),
		Dr6:          binary.LittleEndian.Uint64(data[104:112]),
		Dr7:          binary.LittleEndian.Uint64(data[112:120]),
	}

	regs := []*uint64{
		&ctx.Rax, &ctx.Rcx, &ctx.Rdx, &ctx.Rbx, &ctx.Rsp, &ctx.Rbp, &ctx.Rsi, &ctx.Rdi,
		&ctx.R8, &ctx.R9, &ctx.R10, &ctx.R11, &ctx.R12, &ctx.R13, &ctx.R14, &ctx.R15, &ctx.Rip,
	}
	for i, reg := range regs {
		offset := 136 + i*8
		*reg = binary.LittleEndian.Uint64(data[offset : offset+8])
	}

	fltSave, err := DecodeFloatingSaveArea(data[272:384])
	if err != nil {
		return nil, err
	}
	ctx.FltSave = *fltSave

	for i := 0; i < 26; i++ {
		offset := 384 + i*16
		m128a, err := DecodeM128A(data[offset : offset+16])
		if err != nil {
			return nil, err
		}
		ctx.VectorRegister[i] = *m128a
	}

	ctx.VectorControl = binary.LittleEndian.Uint64(data[800:808])
	ctx.DebugControl = binary.LittleEndian.Uint64(data[808:816])
	ctx.LastBranchToRip = binary.LittleEndian.Uint64(data[816:824])
	ctx.LastBranchFromRip = binary.LittleEndian.Uint64(data[824:832])
	ctx.LastExceptionToRip = binary.LittleEndian.Uint64(data[832:840])
	ctx.LastExceptionFromRip = binary.LittleEndian.Uint64(data[840:848])

	return ctx, nil
}

// EncodeStartupInfo 编码StartupInfo为字节数组
func EncodeStartupInfo(info *StartupInfo) ([]byte, error) {
	buf := make([]byte, 104)

	binary.LittleEndian.PutUint32(buf[0:4], info.Cb)
	binary.LittleEndian.PutUint64(buf[8:16], uint64(uintptr(unsafe.Pointer(info.Desktop))))
	binary.LittleEndian.PutUint64(buf[16:24], uint64(uintptr(unsafe.Pointer(info.Title))))
	binary.LittleEndian.PutUint32(buf[24:28], info.X)
	binary.LittleEndian.PutUint32(buf[28:32], info.Y)
	binary.LittleEndian.PutUint32(buf[32:36], info.XSize)
	binary.LittleEndian.PutUint32(buf[36:40], info.YSize)
	binary.LittleEndian.PutUint32(buf[40:44], info.XCountChars)
	binary.LittleEndian.PutUint32(buf[44:48], info.YCountChars)
	binary.LittleEndian.PutUint32(buf[48:52], info.FillAttribute)
	binary.LittleEndian.PutUint32(buf[52:56], info.Flags)
	binary.LittleEndian.PutUint16(buf[56:58], info.ShowWindow)
	binary.LittleEndian.PutUint16(buf[58:60], info.CbReserved2)
	binary.LittleEndian.PutUint64(buf[64:72], uint64(info.StdInput))
	binary.LittleEndian.PutUint64(buf[72:80], uint64(info.StdOutput))
	binary.LittleEndian.PutUint64(buf[80:88], uint64(info.StdError))

	return buf, nil
}

// DecodeStartupInfo 从字节数组解码StartupInfo
func DecodeStartupInfo(data []byte) (*StartupInfo, error) {
	if len(data) < 104 {
		return nil, fmt.Errorf("数据长度不足，需要104字节，实际%d字节", len(data))
	}

	return &StartupInfo{
		Cb:            binary.LittleEndian.Uint32(data[0:4]),
		Desktop:       (*uint16)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(data[8:16])))),
		Title:         (*uint16)(unsafe.Pointer(uintptr(binary.LittleEndian.Uint64(data[16:24])))),
		X:             binary.LittleEndian.Uint32(data[24:28]),
		Y:             binary.LittleEndian.Uint32(data[28:32]),
		XSize:         binary.LittleEndian.Uint32(data[32:36]),
		YSize:         binary.LittleEndian.Uint32(data[36:40]),
		XCountChars:   binary.LittleEndian.Uint32(data[40:44]),
		YCountChars:   binary.LittleEndian.Uint32(data[44:48]),
		FillAttribute: binary.LittleEndian.Uint32(data[48:52]),
		Flags:         binary.LittleEndian.Uint32(data[52:56]),
		ShowWindow:    binary.LittleEndian.Uint16(data[56:58]),
		CbReserved2:   binary.LittleEndian.Uint16(data[58:60]),
		StdInput:      Handle(binary.LittleEndian.Uint64(data[64:72])),
		StdOutput:     Handle(binary.LittleEndian.Uint64(data[72:80])),
		StdError:      Handle(binary.LittleEndian.Uint64(data[80:88])),
	}, nil
}

// EncodeProcessInformation 编码ProcessInformation为字节数组
func EncodeProcessInformation(info *ProcessInformation) ([]byte, error) {
	buf := make([]byte, 24)

	binary.LittleEndian.PutUint64(buf[0:8], uint64(info.Process))
	binary.LittleEndian.PutUint64(buf[8:16], uint64(info.Thread))
	binary.LittleEndian.PutUint32(buf[16:20], info.ProcessId)
	binary.LittleEndian.PutUint32(buf[20:24], info.ThreadId)

	return buf, nil
}

// DecodeProcessInformation 从字节数组解码ProcessInformation
func DecodeProcessInformation(data []byte) (*ProcessInformation, error) {
	if len(data) < 24 {
		return nil, fmt.Errorf("数据长度不足，需要24字节，实际%d字节", len(data))
	}

	return &ProcessInformation{
		Process:   Handle(binary.LittleEndian.Uint64(data[0:8])),
		Thread:    Handle(binary.LittleEndian.Uint64(data[8:16])),
		ProcessId: binary.LittleEndian.Uint32(data[16:20]),
		ThreadId:  binary.LittleEndian.Uint32(data[20:24]),
	}, nil
}

// EncodeMemoryBasicInformation 编码MemoryBasicInformation为字节数组
func EncodeMemoryBasicInformation(info *MemoryBasicInformation) ([]byte, error) {
	buf := make([]byte, 48)

	binary.LittleEndian.PutUint64(buf[0:8], info.BaseAddress)
	binary.LittleEndian.PutUint64(buf[8:16], info.AllocationBase)
	binary.LittleEndian.PutUint32(buf[16:20], info.AllocationProtect)
	binary.LittleEndian.PutUint64(buf[24:32], info.RegionSize)
	binary.LittleEndian.PutUint32(buf[32:36], info.State)
	binary.LittleEndian.PutUint32(buf[36:40], info.Protect)
	binary.LittleEndian.PutUint32(buf[40:44], info.Type)

	return buf, nil
}

// DecodeMemoryBasicInformation 从字节数组解码MemoryBasicInformation
func DecodeMemoryBasicInformation(data []byte) (*MemoryBasicInformation, error) {
	if len(data) < 48 {
		return nil, fmt.Errorf("数据长度不足，需要48字节，实际%d字节", len(data))
	}

	return &MemoryBasicInformation{
		BaseAddress:       binary.LittleEndian.Uint64(data[0:8]),
		AllocationBase:    binary.LittleEndian.Uint64(data[8:16]),
		AllocationProtect: binary.LittleEndian.Uint32(data[16:20]),
		RegionSize:        binary.LittleEndian.Uint64(data[24:32]),
		State:             binary.LittleEndian.Uint32(data[32:36]),
		Protect:           binary.LittleEndian.Uint32(data[36:40]),
		Type:              binary.LittleEndian.Uint32(data[40:44]),
	}, nil
}

// EncodeDebugEvent 编码DebugEvent为字节数组
func EncodeDebugEvent(event *DebugEvent) ([]byte, error) {
	buf := make([]byte, 172)

	binary.LittleEndian.PutUint32(buf[0:4], event.DebugEventCode())
	binary.LittleEndian.PutUint32(buf[4:8], event.ProcessId())
	binary.LittleEndian.PutUint32(buf[8:12], event.ThreadId())

	copy(buf[12:172], event.data[12:172])

	return buf, nil
}

// DecodeDebugEvent 从字节数组解码DebugEvent
func DecodeDebugEvent(data []byte) (*DebugEvent, error) {
	if len(data) < 172 {
		return nil, fmt.Errorf("数据长度不足，需要172字节，实际%d字节", len(data))
	}

	event := &DebugEvent{}
	copy(event.data[:], data)

	return event, nil
}

// Encode 编码ExceptionRecord为字节数组
func (r *ExceptionRecord) Encode() ([]byte, error) {
	return EncodeExceptionRecord(r)
}

// Decode 从字节数组解码ExceptionRecord
func (r *ExceptionRecord) Decode(data []byte) error {
	record, err := DecodeExceptionRecord(data)
	if err != nil {
		return err
	}
	*r = *record
	return nil
}

// Encode 编码ExceptionDebugInfo为字节数组
func (e *ExceptionDebugInfo) Encode() ([]byte, error) {
	return EncodeExceptionDebugInfo(e)
}

// Decode 从字节数组解码ExceptionDebugInfo
func (e *ExceptionDebugInfo) Decode(data []byte) error {
	info, err := DecodeExceptionDebugInfo(data)
	if err != nil {
		return err
	}
	*e = *info
	return nil
}

// Encode 编码CreateThreadDebugInfo为字节数组
func (e *CreateThreadDebugInfo) Encode() ([]byte, error) {
	return EncodeCreateThreadDebugInfo(e)
}

// Decode 从字节数组解码CreateThreadDebugInfo
func (e *CreateThreadDebugInfo) Decode(data []byte) error {
	info, err := DecodeCreateThreadDebugInfo(data)
	if err != nil {
		return err
	}
	*e = *info
	return nil
}

// Encode 编码CreateProcessDebugInfo为字节数组
func (e *CreateProcessDebugInfo) Encode() ([]byte, error) {
	return EncodeCreateProcessDebugInfo(e)
}

// Decode 从字节数组解码CreateProcessDebugInfo
func (e *CreateProcessDebugInfo) Decode(data []byte) error {
	info, err := DecodeCreateProcessDebugInfo(data)
	if err != nil {
		return err
	}
	*e = *info
	return nil
}

// Encode 编码ExitThreadDebugInfo为字节数组
func (e *ExitThreadDebugInfo) Encode() ([]byte, error) {
	return EncodeExitThreadDebugInfo(e)
}

// Decode 从字节数组解码ExitThreadDebugInfo
func (e *ExitThreadDebugInfo) Decode(data []byte) error {
	info, err := DecodeExitThreadDebugInfo(data)
	if err != nil {
		return err
	}
	*e = *info
	return nil
}

// Encode 编码ExitProcessDebugInfo为字节数组
func (e *ExitProcessDebugInfo) Encode() ([]byte, error) {
	return EncodeExitProcessDebugInfo(e)
}

// Decode 从字节数组解码ExitProcessDebugInfo
func (e *ExitProcessDebugInfo) Decode(data []byte) error {
	info, err := DecodeExitProcessDebugInfo(data)
	if err != nil {
		return err
	}
	*e = *info
	return nil
}

// Encode 编码LoadDllDebugInfo为字节数组
func (e *LoadDllDebugInfo) Encode() ([]byte, error) {
	return EncodeLoadDllDebugInfo(e)
}

// Decode 从字节数组解码LoadDllDebugInfo
func (e *LoadDllDebugInfo) Decode(data []byte) error {
	info, err := DecodeLoadDllDebugInfo(data)
	if err != nil {
		return err
	}
	*e = *info
	return nil
}

// Encode 编码UnloadDllDebugInfo为字节数组
func (e *UnloadDllDebugInfo) Encode() ([]byte, error) {
	return EncodeUnloadDllDebugInfo(e)
}

// Decode 从字节数组解码UnloadDllDebugInfo
func (e *UnloadDllDebugInfo) Decode(data []byte) error {
	info, err := DecodeUnloadDllDebugInfo(data)
	if err != nil {
		return err
	}
	*e = *info
	return nil
}

// Encode 编码OutputDebugStringInfo为字节数组
func (e *OutputDebugStringInfo) Encode() ([]byte, error) {
	return EncodeOutputDebugStringInfo(e)
}

// Decode 从字节数组解码OutputDebugStringInfo
func (e *OutputDebugStringInfo) Decode(data []byte) error {
	info, err := DecodeOutputDebugStringInfo(data)
	if err != nil {
		return err
	}
	*e = *info
	return nil
}

// Encode 编码M128A为字节数组
func (e *M128A) Encode() ([]byte, error) {
	return EncodeM128A(e)
}

// Decode 从字节数组解码M128A
func (e *M128A) Decode(data []byte) error {
	info, err := DecodeM128A(data)
	if err != nil {
		return err
	}
	*e = *info
	return nil
}

// Encode 编码FloatingSaveArea为字节数组
func (e *FloatingSaveArea) Encode() ([]byte, error) {
	return EncodeFloatingSaveArea(e)
}

// Decode 从字节数组解码FloatingSaveArea
func (e *FloatingSaveArea) Decode(data []byte) error {
	info, err := DecodeFloatingSaveArea(data)
	if err != nil {
		return err
	}
	*e = *info
	return nil
}

// Encode 编码Context为字节数组
func (e *Context) Encode() ([]byte, error) {
	return EncodeContext(e)
}

// Decode 从字节数组解码Context
func (e *Context) Decode(data []byte) error {
	info, err := DecodeContext(data)
	if err != nil {
		return err
	}
	*e = *info
	return nil
}

// Encode 编码StartupInfo为字节数组
func (e *StartupInfo) Encode() ([]byte, error) {
	return EncodeStartupInfo(e)
}

// Decode 从字节数组解码StartupInfo
func (e *StartupInfo) Decode(data []byte) error {
	info, err := DecodeStartupInfo(data)
	if err != nil {
		return err
	}
	*e = *info
	return nil
}

// Encode 编码ProcessInformation为字节数组
func (e *ProcessInformation) Encode() ([]byte, error) {
	return EncodeProcessInformation(e)
}

// Decode 从字节数组解码ProcessInformation
func (e *ProcessInformation) Decode(data []byte) error {
	info, err := DecodeProcessInformation(data)
	if err != nil {
		return err
	}
	*e = *info
	return nil
}

// Encode 编码MemoryBasicInformation为字节数组
func (e *MemoryBasicInformation) Encode() ([]byte, error) {
	return EncodeMemoryBasicInformation(e)
}

// Decode 从字节数组解码MemoryBasicInformation
func (e *MemoryBasicInformation) Decode(data []byte) error {
	info, err := DecodeMemoryBasicInformation(data)
	if err != nil {
		return err
	}
	*e = *info
	return nil
}

// Encode 编码DebugEvent为字节数组
func (e *DebugEvent) Encode() ([]byte, error) {
	return EncodeDebugEvent(e)
}

// Decode 从字节数组解码DebugEvent
func (e *DebugEvent) Decode(data []byte) error {
	event, err := DecodeDebugEvent(data)
	if err != nil {
		return err
	}
	*e = *event
	return nil
}

// SizeOf 返回结构体的大小
func SizeOf(v interface{}) int {
	return int(unsafe.Sizeof(v))
}
