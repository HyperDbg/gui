package misc

import (
	"encoding/binary"

	"github.com/ddkwork/sdk"
)

type MemoryReadResult struct {
	Data        []byte
	BytesRead   uint32
	AddressMode sdk.DEBUGGER_READ_MEMORY_ADDRESS_MODE
	Success     bool
}

func FormatMemoryAsBytes(data []byte, groupSize int) [][]byte {
	var result [][]byte
	for i := 0; i < len(data); i += groupSize {
		end := min(i+groupSize, len(data))
		result = append(result, data[i:end])
	}
	return result
}

func FormatMemoryAsDwords(data []byte) []uint32 {
	var dwords []uint32
	for i := 0; i+3 < len(data); i += 4 {
		dwords = append(dwords, binary.LittleEndian.Uint32(data[i:i+4]))
	}
	return dwords
}

func FormatMemoryAsQwords(data []byte) []uint64 {
	var qwords []uint64
	for i := 0; i+7 < len(data); i += 8 {
		qwords = append(qwords, binary.LittleEndian.Uint64(data[i:i+8]))
	}
	return qwords
}

func FormatMemoryAsString(data []byte) string {
	var result []byte
	for _, b := range data {
		if b >= 0x20 && b < 0x7f {
			result = append(result, b)
		} else {
			result = append(result, '.')
		}
	}
	return string(result)
}

func CalculateChecksum(data []byte) byte {
	var sum byte
	for _, b := range data {
		sum += b
	}
	return sum
}

func CompareBuffers(buf1, buf2 []byte) bool {
	if len(buf1) != len(buf2) {
		return false
	}
	for i := range buf1 {
		if buf1[i] != buf2[i] {
			return false
		}
	}
	return true
}
