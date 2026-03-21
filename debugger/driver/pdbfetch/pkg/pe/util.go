package pe

import (
	"regexp"
)

// This will set a maximum length of a string to be retrieved from the file.
// It's there to prevent loading massive amounts of data from memory mapped
// files. Strings longer than 1MB should be rather rare.
const (
	MAX_STRING_LENGTH = 0x100000 // 2^20
)

var invalidImportName = []byte("<invalid>")

func MaxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func MinInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func MaxInt32(x, y int32) int32 {
	if x > y {
		return x
	}
	return y
}

func MinInt32(x, y int32) int32 {
	if x < y {
		return x
	}
	return y
}

func MaxUInt32(x, y uint32) uint32 {
	if x > y {
		return x
	}
	return y
}

func MinUInt32(x, y uint32) uint32 {
	if x < y {
		return x
	}
	return y
}

// Returns whether this value is a power of 2
func PowerOfTwo(val uint32) bool {
	ret := (val != 0) && (val&(val-1)) == 0x0
	return ret
}

// Helper functions to align numbers

func AlignDownUInt32(x, align uint32) uint32 {
	return x & ^(align - 1)
}

func AlignUpUInt32(x, align uint32) uint32 {
	if (x & (align - 1)) != 0 {
		return AlignDownUInt32(x, align) + align
	}
	return x
}

func AlignDownUInt64(x, align uint64) uint64 {
	return x & ^(align - 1)
}

func AlignUpUInt64(x, align uint64) uint64 {
	if (x & (align - 1)) != 0 {
		return AlignDownUInt64(x, align) + align
	}
	return x
}

func MemsetRepeat(a []byte, v byte) {
	if len(a) == 0 {
		return
	}
	a[0] = v
	for bp := 1; bp < len(a); bp *= 2 {
		copy(a[bp:], a[:bp])
	}
}

// Check if a imported name uses the valid accepted characters expected in mangled
// function names. If the symbol's characters don't fall within this charset
// we will assume the name is invalid.
var validFuncNameRegex = regexp.MustCompile(`^[\pL\pN_\?@$\(\)]+$`)

func validFuncName(name []byte) bool {
	return validFuncNameRegex.Match(name)
}

// Valid FAT32 8.3 short filename characters according to:
//
//	http://en.wikipedia.org/wiki/8.3_filename
//
// This will help decide whether DLL ASCII names are likely
// to be valid or otherwise corrupt data
//
// The filename length is not checked because the DLLs filename
// can be longer that the 8.3.
var validDOSNameRegex = regexp.MustCompile("^[\\pL\\pN!//$%&'\\(\\)`\\-@^_\\{\\}~+,.;=\\[\\]]+$")

func validDosFilename(name []byte) bool {
	return validDOSNameRegex.Match(name)
}
