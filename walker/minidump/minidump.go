package minidump

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

const (
	Signature     = "MDMP"
	ValidVersion  = 0xA793
	DirectorySize = 96
)

type Header struct {
	Signature          [4]byte
	Version            uint32
	NumberOfStreams    uint32
	StreamDirectoryRva uint32
	CheckSum           uint32
	TimeDateStamp      uint32
	Flags              uint64
}

type Directory struct {
	StreamType uint32
	DataSize   uint32
	Rva        uint32
}

const (
	StreamTypeUnused               = 0
	StreamTypeReserved0            = 1
	StreamTypeReserved1            = 2
	StreamTypeThreadList           = 3
	StreamTypeModuleList           = 4
	StreamTypeMemoryList           = 5
	StreamTypeException            = 6
	StreamTypeSystemInfo           = 7
	StreamTypeThreadExList         = 8
	StreamTypeMemory64List         = 9
	StreamTypeCommentStreamA       = 10
	StreamTypeCommentStreamW       = 11
	StreamTypeHandleData           = 12
	StreamTypeFunctionTable        = 13
	StreamTypeVariableStream       = 14
	StreamTypeTokenStream          = 15
	StreamTypeJavaScriptDataStream = 16
	StreamTypeSystemMemoryInfo     = 17
	StreamTypeProcessVmCounters    = 18
	StreamTypeIptTrace             = 19
	StreamTypeThreadInfoList       = 20
	StreamTypeHandleOperationList  = 21
	StreamTypeLastReservedStream   = 0xFFFF
)

type Minidump struct {
	Header      Header
	Directories []Directory
	Streams     map[uint32][]byte
}

func Parse(path string) (*Minidump, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	m := &Minidump{
		Streams: make(map[uint32][]byte),
	}

	if err := binary.Read(file, binary.LittleEndian, &m.Header); err != nil {
		return nil, fmt.Errorf("failed to read header: %w", err)
	}

	sig := string(m.Header.Signature[:])
	if sig != Signature {
		return nil, fmt.Errorf("invalid signature: %s (expected %s)", sig, Signature)
	}

	if m.Header.Version != ValidVersion {
		return nil, fmt.Errorf("invalid version: 0x%X (expected 0x%X)", m.Header.Version, ValidVersion)
	}

	m.Directories = make([]Directory, m.Header.NumberOfStreams)
	if err := binary.Read(file, binary.LittleEndian, &m.Directories); err != nil {
		return nil, fmt.Errorf("failed to read directories: %w", err)
	}

	for _, dir := range m.Directories {
		if dir.Rva == 0 || dir.DataSize == 0 {
			continue
		}

		data := make([]byte, dir.DataSize)
		if _, err := file.Seek(int64(dir.Rva), io.SeekStart); err != nil {
			return nil, fmt.Errorf("failed to seek to stream data: %w", err)
		}
		if _, err := io.ReadFull(file, data); err != nil {
			return nil, fmt.Errorf("failed to read stream data: %w", err)
		}

		m.Streams[dir.StreamType] = data
	}

	return m, nil
}

func (m *Minidump) GetStream(streamType uint32) ([]byte, bool) {
	data, ok := m.Streams[streamType]
	return data, ok
}

func (m *Minidump) String() string {
	return fmt.Sprintf("Minidump:\n  Signature: %s\n  Version: 0x%X\n  Streams: %d",
		string(m.Header.Signature[:]),
		m.Header.Version,
		m.Header.NumberOfStreams,
	)
}
