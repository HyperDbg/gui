package windows

import (
	"encoding/hex"
	"fmt"
	"unicode/utf16"
)

// HexDump hex dump功能，用于协助解析数据结构
type HexDump struct {
	data    []byte
	address uint64
}

// NewHexDump 创建一个新的hex dump
func NewHexDump(data []byte, address uint64) *HexDump {
	return &HexDump{
		data:    data,
		address: address,
	}
}

// String 返回hex dump的字符串表示
func (h *HexDump) String() string {
	if h.address == 0 {
		return string(hex.Dump(h.data))
	}

	result := ""
	for i := 0; i < len(h.data); i += 16 {
		result += fmt.Sprintf("%08X  ", h.address+uint64(i))

		for j := 0; j < 16; j++ {
			if i+j < len(h.data) {
				result += fmt.Sprintf("%02x ", h.data[i+j])
			} else {
				result += "   "
			}
		}

		result += " |"

		for j := 0; j < 16; j++ {
			if i+j < len(h.data) {
				b := h.data[i+j]
				if b >= 32 && b <= 126 {
					result += string(b)
				} else {
					result += "."
				}
			} else {
				result += " "
			}
		}

		result += "|\n"
	}

	return result
}

// Print 打印hex dump
func (h *HexDump) Print() {
	fmt.Print(h.String())
}

// TryParseUTF16 尝试解析为UTF-16字符串
func (h *HexDump) TryParseUTF16() (string, bool) {
	if len(h.data) < 2 {
		return "", false
	}

	runes := make([]uint16, 0, len(h.data)/2)
	for i := 0; i < len(h.data)-1; i += 2 {
		r := uint16(h.data[i]) | uint16(h.data[i+1])<<8
		runes = append(runes, r)
	}

	str := string(utf16.Decode(runes))
	return str, true
}

// TryParseUTF8 尝试解析为UTF-8字符串
func (h *HexDump) TryParseUTF8() (string, bool) {
	str := string(h.data)
	return str, true
}

// TryParseString 尝试解析为字符串（先UTF-16，后UTF-8）
func (h *HexDump) TryParseString() (string, string, bool) {
	if utf16Str, ok := h.TryParseUTF16(); ok {
		return utf16Str, "", true
	}

	if utf8Str, ok := h.TryParseUTF8(); ok {
		return "", utf8Str, true
	}

	return "", "", false
}

// DumpBytes hex dump字节数组
func DumpBytes(data []byte, address uint64) {
	dump := NewHexDump(data, address)
	dump.Print()
}
