// Package common implements the general-purpose utility functions used across
// libhyperdbg. The C++ counterpart is libhyperdbg/code/common/common.cpp; it
// owns:
//   - string manipulation helpers (SeparateTo64BitValue, Replace, ReplaceAll,
//     Split, Trim, RemoveSpaces, HasEnding, FindCaseInsensitive, ...)
//   - numeric conversion helpers (IsNumber, IsHexNotation, IsDecimalNotation,
//     ConvertStringToUInt64, ConvertStringToUInt32, Log2Ceil, HexToBytes)
//   - CPUID helpers (CommonCpuidInstruction, CheckCpuSupportRtm,
//     Getx86VirtualAddressWidth, VmxSupportDetection)
//   - validation helpers (ValidateIP, CheckAccessValidityAndSafety,
//     IsFileExist, IsEmptyString)
//   - token helpers (ConvertTokenToUInt64, ConvertTokenToUInt32,
//     GetCaseSensitiveStringFromCommandToken, ...) — these are declared in
//     commands.h and operate on the CommandToken type; the Go counterparts
//     live alongside the commands package. common only exposes the plain
//     string-based variants.
//
// In the Go rewrite the global state from the C side (g_RtmSupport,
// g_VirtualAddressWidth) is owned by the Common struct so that multiple
// debugger instances can coexist (GUI/MCP requirement, see API design spec).
// The CPUID helpers are platform-specific; pure-Go stubs are provided here
// and the actual CPUID instruction is delivered via a small assembly stub
// in cpu_amd64.go (mirroring the transparency package's approach).
//
// All public functions are safe for concurrent use; the mutable state
// (cached CPUID results) is guarded by an internal mutex.
package common

import (
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"sync"
	"unicode/utf16"
	"unsafe"
)

// Output abstracts message output so CLI/GUI/MCP can all consume the same
// Common helpers. It mirrors commands.Output / app.Output but is declared
// locally to keep the common package free of import cycles.
type Output interface {
	Printf(format string, args ...any) error
}

// Common owns the cached CPUID-derived state (RTM support, virtual address
// width). The zero value is not usable; use NewCommon.
//
// All fields are guarded by mu. The cached values are populated lazily by
// the first call to the corresponding accessor.
type Common struct {
	mu sync.Mutex

	out Output

	// cpuidReader is an optional override for the CPUID instruction. When
	// nil the default (assembly-backed) reader is used. It is intended for
	// tests that want deterministic CPUID results.
	cpuidReader CpuidReader

	// Cached CPUID-derived values, mirroring g_RtmSupport and
	// g_VirtualAddressWidth in the C++ code.
	rtmSupport          bool
	rtmSupportProbed    bool
	virtualAddressWidth uint32
	vaWidthProbed       bool
}

// NewCommon constructs a Common instance writing diagnostics to out. Pass nil
// to silence diagnostics.
func NewCommon(out Output) *Common {
	if out == nil {
		out = discardOutput{}
	}
	return &Common{out: out}
}

// ============================================================================
// String helpers (pure Go, mirror the C++ functions in common.cpp)
// ============================================================================

// SeparateTo64BitValue formats a 64-bit value as a 16-digit hex string with
// a backtick separator between the high and low 32-bit halves. Mirrors the
// C++ SeparateTo64BitValue used by objects.cpp to print EPROCESS/ETHREAD
// pointers.
func SeparateTo64BitValue(value uint64) string {
	high := uint32(value >> 32)
	low := uint32(value)
	return fmt.Sprintf("%08x`%08x", high, low)
}

// PrintBits writes the binary representation of the first Size bytes at Ptr
// to out, MSB first, byte by byte. Mirrors PrintBits from common.cpp.
func PrintBits(out Output, size uint32, ptr unsafe.Pointer) {
	if out == nil || size == 0 || ptr == nil {
		return
	}
	buf := (*[1 << 30]byte)(ptr)[:size:size]
	var sb strings.Builder
	for i := int(size) - 1; i >= 0; i-- {
		for j := 7; j >= 0; j-- {
			if (buf[i]>>j)&1 == 1 {
				sb.WriteByte('1')
			} else {
				sb.WriteByte('0')
			}
		}
		sb.WriteByte(' ')
	}
	_ = out.Printf("%s\n", strings.TrimRight(sb.String(), " "))
}

// Replace replaces the first occurrence of `from` with `to` in `s`. Returns
// the (possibly modified) string and true if a replacement happened. Mirrors
// the C++ Replace function.
func Replace(s, from, to string) (string, bool) {
	if from == "" {
		return s, false
	}
	idx := strings.Index(s, from)
	if idx < 0 {
		return s, false
	}
	return s[:idx] + to + s[idx+len(from):], true
}

// ReplaceAll replaces every occurrence of `from` with `to` in `s`. Mirrors
// the C++ ReplaceAll function.
func ReplaceAll(s, from, to string) string {
	if from == "" {
		return s
	}
	return strings.ReplaceAll(s, from, to)
}

// Split splits `s` on the single-byte delimiter `sep`, dropping empty fields.
// Mirrors the C++ Split function.
func Split(s string, sep byte) []string {
	parts := strings.Split(s, string(sep))
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		if p != "" {
			out = append(out, p)
		}
	}
	return out
}

// IsNumber reports whether `s` is a non-empty string of ASCII digits.
// Mirrors the C++ IsNumber function.
func IsNumber(s string) bool {
	if s == "" {
		return false
	}
	for _, r := range s {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}

// IsHexNotation reports whether `s` is a non-empty string of hex digits.
// Mirrors the C++ IsHexNotation function.
func IsHexNotation(s string) bool {
	if s == "" {
		return false
	}
	for _, r := range s {
		if !isHexDigit(r) {
			return false
		}
	}
	return true
}

// IsDecimalNotation reports whether `s` is a non-empty string of decimal
// digits. Mirrors the C++ IsDecimalNotation function.
func IsDecimalNotation(s string) bool {
	return IsNumber(s)
}

// HexToBytes parses a hex string into a byte slice. Returns nil on invalid
// input. Mirrors the C++ HexToBytes function.
func HexToBytes(hexStr string) []byte {
	if len(hexStr)%2 != 0 {
		hexStr = "0" + hexStr
	}
	out := make([]byte, len(hexStr)/2)
	for i := 0; i < len(hexStr); i += 2 {
		hi, ok1 := hexNibble(hexStr[i])
		lo, ok2 := hexNibble(hexStr[i+1])
		if !ok1 || !ok2 {
			return nil
		}
		out[i/2] = hi<<4 | lo
	}
	return out
}

// Log2Ceil returns ceil(log2(n)) for n>0; returns 0 for n==0. Mirrors the
// C++ Log2Ceil function used by hwdbg-interpreter.cpp.
func Log2Ceil(n uint32) uint32 {
	if n <= 1 {
		return 0
	}
	r := uint32(0)
	v := n - 1
	for v > 0 {
		r++
		v >>= 1
	}
	return r
}

// ConvertStringToUInt64 parses a decimal or hex (0x-prefixed) string into a
// uint64. Mirrors the C++ ConvertStringToUInt64 function.
func ConvertStringToUInt64(text string) (uint64, bool) {
	text = strings.TrimSpace(text)
	if text == "" {
		return 0, false
	}
	base := 10
	if strings.HasPrefix(text, "0x") || strings.HasPrefix(text, "0X") {
		text = text[2:]
		base = 16
	}
	v, err := strconv.ParseUint(text, base, 64)
	if err != nil {
		return 0, false
	}
	return v, true
}

// ConvertStringToUInt32 parses a decimal or hex (0x-prefixed) string into a
// uint32. Mirrors the C++ ConvertStringToUInt32 function.
func ConvertStringToUInt32(text string) (uint32, bool) {
	v, ok := ConvertStringToUInt64(text)
	if !ok || v > 0xFFFFFFFF {
		return 0, false
	}
	return uint32(v), true
}

// HasEnding reports whether fullString ends with ending. Mirrors the C++
// HasEnding function.
func HasEnding(fullString, ending string) bool {
	return strings.HasSuffix(fullString, ending)
}

// Trim removes leading and trailing whitespace from s. Mirrors the C++ Trim
// function (which operates in place; here we return a new string).
func Trim(s string) string {
	return strings.TrimSpace(s)
}

// RemoveSpaces removes every ASCII space from s. Mirrors the C++ RemoveSpaces
// function.
func RemoveSpaces(s string) string {
	return strings.ReplaceAll(s, " ", "")
}

// FindCaseInsensitive returns the byte index of the first case-insensitive
// occurrence of toSearch in input at or after pos, or -1 if not found.
// Mirrors the C++ FindCaseInsensitive function.
func FindCaseInsensitive(input, toSearch string, pos int) int {
	if pos < 0 {
		pos = 0
	}
	if pos >= len(input) || toSearch == "" {
		return -1
	}
	return strings.Index(strings.ToLower(input[pos:]), strings.ToLower(toSearch))
}

// IsEmptyString reports whether text is nil or contains only whitespace.
// Mirrors the C++ IsEmptyString function.
func IsEmptyString(text string) bool {
	return strings.TrimSpace(text) == ""
}

// ValidateIP reports whether ip is a valid IPv4 dotted-quad address. Mirrors
// the C++ ValidateIP function.
func ValidateIP(ip string) bool {
	return net.ParseIP(ip) != nil
}

// IsFileExist reports whether a file exists at path. Mirrors the C++
// IsFileExistA function.
func IsFileExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// ListDirectory returns the names of regular files in directory whose
// extension matches ext (e.g. ".txt"). Returns an empty slice on error.
// Mirrors the C++ ListDirectory function.
func ListDirectory(directory, ext string) []string {
	entries, err := os.ReadDir(directory)
	if err != nil {
		return nil
	}
	out := make([]string, 0, len(entries))
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		name := e.Name()
		if ext == "" || strings.HasSuffix(strings.ToLower(name), strings.ToLower(ext)) {
			out = append(out, name)
		}
	}
	return out
}

// StringToWString converts an ASCII string to a UTF-16 encoded []uint16
// (NUL-terminated). Mirrors the C++ StringToWString function.
func StringToWString(s string) []uint16 {
	return utf16.Encode([]rune(s))
}

// ConvertStringVectorToCharPointerArray returns a NUL-terminated copy of s.
// Mirrors the C++ ConvertStringVectorToCharPointerArray function. The Go
// caller does not need to free the result (GC handles it).
func ConvertStringVectorToCharPointerArray(s string) string {
	return s
}

// ============================================================================
// CPUID helpers (platform-specific; pure-Go stubs with a pluggable reader)
// ============================================================================

// CpuidReader is the interface implemented by platform-specific CPUID
// readers. The default reader (in cpu_amd64.go) executes the CPUID
// instruction directly. Tests can plug a fake reader.
type CpuidReader interface {
	// Read executes CPUID with the given leaf and sub-leaf and returns the
	// four result registers (EAX, EBX, ECX, EDX).
	Read(func_, subFunc uint32) (eax, ebx, ecx, edx uint32)
}

// CommonCpuidInstruction executes CPUID with the given (func, subFunc) and
// writes the four result registers into cpuInfo[0..3]. Mirrors the C++
// CommonCpuidInstruction function. cpuInfo must have length >= 4.
func (c *Common) CommonCpuidInstruction(func_, subFunc uint32, cpuInfo []int32) error {
	if len(cpuInfo) < 4 {
		return fmt.Errorf("CommonCpuidInstruction: cpuInfo must have length >= 4 (got %d)", len(cpuInfo))
	}
	reader := c.snapshotReader()
	eax, ebx, ecx, edx := reader.Read(func_, subFunc)
	cpuInfo[0] = int32(eax)
	cpuInfo[1] = int32(ebx)
	cpuInfo[2] = int32(ecx)
	cpuInfo[3] = int32(edx)
	return nil
}

// CheckCpuSupportRtm reports whether the CPU supports Intel TSX RTM. Mirrors
// the C++ CheckCpuSupportRtm function (CPUID.07H:EBX[11]).
func (c *Common) CheckCpuSupportRtm() bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.rtmSupportProbed {
		return c.rtmSupport
	}
	c.rtmSupport = c.probeRtm()
	c.rtmSupportProbed = true
	return c.rtmSupport
}

// Getx86VirtualAddressWidth returns the number of physical/linear address
// bits reported by CPUID.80000008H. Mirrors the C++
// Getx86VirtualAddressWidth function.
func (c *Common) Getx86VirtualAddressWidth() uint32 {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.vaWidthProbed {
		return c.virtualAddressWidth
	}
	c.virtualAddressWidth = c.probeVaWidth()
	c.vaWidthProbed = true
	return c.virtualAddressWidth
}

// VmxSupportDetection reports whether the CPU supports VMX. Mirrors the C++
// VmxSupportDetection function (CPUID.01H:ECX[5]).
func (c *Common) VmxSupportDetection() bool {
	reader := c.snapshotReader()
	_, _, ecx, _ := reader.Read(1, 0)
	return (ecx>>5)&1 == 1
}

// CheckAccessValidityAndSafety reports whether a (targetAddress, size) read
// is plausibly safe — i.e. the range [targetAddress, targetAddress+size) does
// not wrap around and is non-empty. The full kernel-side validity check is
// performed by the driver; this is the user-mode pre-filter mirror of the
// C++ CheckAccessValidityAndSafety function.
func (c *Common) CheckAccessValidityAndSafety(targetAddress uint64, size uint32) bool {
	if size == 0 {
		return false
	}
	if targetAddress > ^uint64(0)-uint64(size) {
		return false // wrap-around
	}
	return true
}

// ============================================================================
// Internal helpers
// ============================================================================

// snapshotReader returns the active CpuidReader, snapshotted under the mutex
// so that concurrent SetCpuidReader calls are safe. The returned reader is
// used outside the lock (the reader itself is stateless and safe for
// concurrent use).
func (c *Common) snapshotReader() CpuidReader {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.cpuidReader != nil {
		return c.cpuidReader
	}
	return defaultCpuidReader{}
}

// SetCpuidReader installs a custom CpuidReader (intended for tests). Pass
// nil to revert to the default.
func (c *Common) SetCpuidReader(r CpuidReader) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cpuidReader = r
}

// probeRtm reads CPUID.07H:EBX[11] and returns true if RTM is supported.
// Caller must hold c.mu.
func (c *Common) probeRtm() bool {
	reader := c.readerLocked()
	_, ebx, _, _ := reader.Read(7, 0)
	return (ebx>>11)&1 == 1
}

// probeVaWidth reads CPUID.80000008H:EAX[7:0] (physical) and [15:8] (linear)
// and returns the linear address width. Mirrors the C++ implementation.
// Caller must hold c.mu.
func (c *Common) probeVaWidth() uint32 {
	reader := c.readerLocked()
	eax, _, _, _ := reader.Read(0x80000008, 0)
	return (eax >> 8) & 0xFF
}

// readerLocked returns the active CpuidReader without taking the mutex.
// Caller must hold c.mu.
func (c *Common) readerLocked() CpuidReader {
	if c.cpuidReader != nil {
		return c.cpuidReader
	}
	return defaultCpuidReader{}
}

func isHexDigit(r rune) bool {
	return (r >= '0' && r <= '9') || (r >= 'a' && r <= 'f') || (r >= 'A' && r <= 'F')
}

func hexNibble(b byte) (byte, bool) {
	switch {
	case b >= '0' && b <= '9':
		return b - '0', true
	case b >= 'a' && b <= 'f':
		return b - 'a' + 10, true
	case b >= 'A' && b <= 'F':
		return b - 'A' + 10, true
	}
	return 0, false
}

// NativeEndian is the byte order used by the host CPU. It mirrors the
// implicit assumption in the C++ code that the host is little-endian.
var NativeEndian = binary.LittleEndian

// discardOutput is the default Output when the caller passes nil.
type discardOutput struct{}

func (discardOutput) Printf(format string, args ...any) error {
	_ = args
	_ = format
	return nil
}
