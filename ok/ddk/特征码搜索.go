package ddk

import (
	"strings"
)

type Pattern struct {
	Bytes    []byte
	Masks    []byte
	FirstIdx int
}

func ParsePattern(pattern string) *Pattern {
	clean := strings.ReplaceAll(pattern, " ", "")
	length := len(clean) / 2
	p := &Pattern{
		Bytes:    make([]byte, length),
		Masks:    make([]byte, length),
		FirstIdx: length,
	}

	for i := range length {
		chunk := clean[i*2 : (i+1)*2]
		var mask, value byte

		switch {
		case chunk == "??":
			mask, value = 0xFF, 0x00
		case chunk[0] == '?':
			mask = 0xF0
			value = parseHex(chunk[1])
		case chunk[1] == '?':
			mask = 0x0F
			value = parseHex(chunk[0]) << 4
		default:
			mask = 0x00
			value = (parseHex(chunk[0]) << 4) | parseHex(chunk[1])
		}

		if mask != 0xFF && p.FirstIdx == length {
			p.FirstIdx = i
		}

		p.Bytes[i] = value
		p.Masks[i] = mask
	}
	return p
}

func parseHex(c byte) byte {
	if c >= '0' && c <= '9' {
		return c - '0'
	}
	if c >= 'A' && c <= 'F' {
		return c - 'A' + 10
	}
	if c >= 'a' && c <= 'f' {
		return c - 'a' + 10
	}
	return 0
}

func (p *Pattern) SearchMemoryChunked(mem []byte, chunkSize int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 0; i < len(mem); i += chunkSize {
			end := min(i+chunkSize, len(mem))
			chunk := mem[i:end]
			for offset := range p.searchChunk(chunk) {
				ch <- i + offset
			}
		}
	}()
	return ch
}

func (p *Pattern) SearchMemory(mem []byte) []int {
	var matches []int
	for offset := range p.SearchMemoryChunked(mem, 4096) {
		matches = append(matches, offset)
	}
	return matches
}

func (p *Pattern) searchChunk(chunk []byte) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		length := len(p.Bytes)
		for i := 0; i <= len(chunk)-length; i++ {
			firstCheck := i + p.FirstIdx
			if firstCheck >= len(chunk) {
				break
			}

			masked := chunk[firstCheck] &^ p.Masks[p.FirstIdx]
			if masked != p.Bytes[p.FirstIdx] {
				continue
			}

			match := true
			for j := range length {
				pos := i + j
				if pos >= len(chunk) {
					match = false
					break
				}
				if p.Masks[j] != 0xFF && (chunk[pos]&^p.Masks[j]) != p.Bytes[j] {
					match = false
					break
				}
			}

			if match {
				ch <- i
			}
		}
	}()
	return ch
}
