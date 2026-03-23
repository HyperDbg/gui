package minidump

import (
	"encoding/binary"
	"fmt"
	"io"
)

type ExceptionRecord struct {
	ExceptionCode        uint32
	ExceptionFlags       uint32
	ExceptionRecord      uint64
	ExceptionAddress     uint64
	NumberParameters     uint32
	UnusedAlignment      uint32
	ExceptionInformation [15]uint64
}

type Exception struct {
	ThreadId        uint32
	UnusedAlignment uint32
	ExceptionRecord ExceptionRecord
	ThreadContext   Context
}

type Context struct {
	ContextFlags uint32
	Dr0          uint64
	Dr1          uint64
	Dr2          uint64
	Dr3          uint64
	Dr6          uint64
	Dr7          uint64
	FloatSave    [512]byte
	Gs           uint64
	Fs           uint64
	Es           uint64
	Ds           uint64
	Rax          uint64
	Rbx          uint64
	Rcx          uint64
	Rdx          uint64
	Rsp          uint64
	Rbp          uint64
	Rsi          uint64
	Rdi          uint64
	R8           uint64
	R9           uint64
	R10          uint64
	R11          uint64
	R12          uint64
	R13          uint64
	R14          uint64
	R15          uint64
	Rip          uint64
	Flgs         uint32
	Ss           uint16
	Cs           uint16
	DsPadding    uint16
	EspPadding   uint16
}

const (
	CONTEXT_AMD64 = 0x00100000
)

func (m *Minidump) GetException() (*Exception, error) {
	data, ok := m.GetStream(StreamTypeException)
	if !ok {
		return nil, fmt.Errorf("exception stream not found")
	}

	var exc Exception
	if err := binary.Read(newBytesReader(data), binary.LittleEndian, &exc); err != nil {
		return nil, fmt.Errorf("failed to read exception: %w", err)
	}

	return &exc, nil
}

type Thread struct {
	ThreadId      uint32
	SuspendCount  uint32
	PriorityClass uint32
	Priority      uint32
	Teb           uint64
	Stack         MemoryDescriptor
	ThreadContext Context
}

type MemoryDescriptor struct {
	StartOfMemoryRange uint64
	DataSize           uint64
}

func (m *Minidump) GetThreads() ([]Thread, error) {
	data, ok := m.GetStream(StreamTypeThreadList)
	if !ok {
		return nil, fmt.Errorf("thread list stream not found")
	}

	reader := newBytesReader(data)

	var numThreads uint32
	if err := binary.Read(reader, binary.LittleEndian, &numThreads); err != nil {
		return nil, fmt.Errorf("failed to read thread count: %w", err)
	}

	threads := make([]Thread, numThreads)
	for i := uint32(0); i < numThreads; i++ {
		if err := binary.Read(reader, binary.LittleEndian, &threads[i]); err != nil {
			return nil, fmt.Errorf("failed to read thread %d: %w", i, err)
		}
	}

	return threads, nil
}

type Module struct {
	BaseOfImage   uint64
	SizeOfImage   uint32
	CheckSum      uint32
	TimeDateStamp uint32
	ModuleNameRva uint32
	VersionInfo   ImageDataDirectory
	CvRecord      ImageDataDirectory
	MiscRecord    ImageDataDirectory
	Reserved0     uint64
	Reserved1     uint64
}

type ImageDataDirectory struct {
	DataSize uint32
	Rva      uint32
}

func (m *Minidump) GetModules() ([]Module, []string, error) {
	data, ok := m.GetStream(StreamTypeModuleList)
	if !ok {
		return nil, nil, fmt.Errorf("module list stream not found")
	}

	reader := newBytesReader(data)

	var numModules uint32
	if err := binary.Read(reader, binary.LittleEndian, &numModules); err != nil {
		return nil, nil, fmt.Errorf("failed to read module count: %w", err)
	}

	modules := make([]Module, numModules)
	for i := uint32(0); i < numModules; i++ {
		if err := binary.Read(reader, binary.LittleEndian, &modules[i]); err != nil {
			return nil, nil, fmt.Errorf("failed to read module %d: %w", i, err)
		}
	}

	names := make([]string, numModules)
	for i, mod := range modules {
		if mod.ModuleNameRva != 0 {
			nameData, ok := m.readRvaString(mod.ModuleNameRva)
			if ok {
				names[i] = nameData
			}
		}
	}

	return modules, names, nil
}

type SystemInfo struct {
	ProcessorArchitecture uint16
	ProcessorLevel        uint16
	ProcessorRevision     uint16
	Reserved0             uint8
	NumberOfProcessors    uint8
	ProductType           uint8
	MajorVersion          uint8
	MinorVersion          uint8
	BuildNumber           uint16
	PlatformId            uint32
	CSDVersionRva         uint32
	Reserved1             uint16
	Flgs                  uint16
	Reserved2             [3]uint32
}

func (m *Minidump) GetSystemInfo() (*SystemInfo, error) {
	data, ok := m.GetStream(StreamTypeSystemInfo)
	if !ok {
		return nil, fmt.Errorf("system info stream not found")
	}

	var sysInfo SystemInfo
	if err := binary.Read(newBytesReader(data), binary.LittleEndian, &sysInfo); err != nil {
		return nil, fmt.Errorf("failed to read system info: %w", err)
	}

	return &sysInfo, nil
}

type Memory64List struct {
	NumberOfMemoryRanges uint64
	BaseRva              uint64
	MemoryRanges         []MemoryDescriptor64
}

type MemoryDescriptor64 struct {
	StartOfMemoryRange uint64
	DataSize           uint64
}

func (m *Minidump) GetMemory64List() (*Memory64List, error) {
	data, ok := m.GetStream(StreamTypeMemory64List)
	if !ok {
		return nil, fmt.Errorf("memory64 list stream not found")
	}

	reader := newBytesReader(data)

	var list Memory64List
	if err := binary.Read(reader, binary.LittleEndian, &list.NumberOfMemoryRanges); err != nil {
		return nil, fmt.Errorf("failed to read memory range count: %w", err)
	}
	if err := binary.Read(reader, binary.LittleEndian, &list.BaseRva); err != nil {
		return nil, fmt.Errorf("failed to read base RVA: %w", err)
	}

	list.MemoryRanges = make([]MemoryDescriptor64, list.NumberOfMemoryRanges)
	for i := uint64(0); i < list.NumberOfMemoryRanges; i++ {
		if err := binary.Read(reader, binary.LittleEndian, &list.MemoryRanges[i]); err != nil {
			return nil, fmt.Errorf("failed to read memory range %d: %w", i, err)
		}
	}

	return &list, nil
}

func (m *Minidump) readRvaString(rva uint32) (string, bool) {
	for _, dir := range m.Directories {
		if dir.StreamType == StreamTypeModuleList {
			offset := int(rva - dir.Rva)
			if offset >= 0 && offset < len(m.Streams[dir.StreamType]) {
				data := m.Streams[dir.StreamType][offset:]
				end := 0
				for i, b := range data {
					if b == 0 {
						end = i
						break
					}
				}
				if end > 0 {
					return string(data[:end]), true
				}
			}
		}
	}
	return "", false
}

type bytesReader struct {
	data []byte
	pos  int
}

func newBytesReader(data []byte) *bytesReader {
	return &bytesReader{data: data}
}

func (r *bytesReader) Read(p []byte) (n int, err error) {
	if r.pos >= len(r.data) {
		return 0, io.EOF
	}
	n = copy(p, r.data[r.pos:])
	r.pos += n
	return n, nil
}
