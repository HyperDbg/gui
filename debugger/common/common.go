package common

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func SeparateTo64BitValue(value uint64) string {
	return fmt.Sprintf("%08x`%08x", value>>32, value&0xFFFFFFFF)
}

func PrintBits(size uint32, ptr []byte) {
	for i := int(size) - 1; i >= 0; i-- {
		for j := 7; j >= 0; j-- {
			fmt.Printf("%d", (ptr[i]>>j)&1)
		}
		fmt.Printf(" ")
	}
}

func ReplaceAll(s, from, to string) string {
	return strings.ReplaceAll(s, from, to)
}

func Split(s string, c rune) []string {
	var result []string
	var buf strings.Builder
	for _, ch := range s {
		if ch != c {
			buf.WriteRune(ch)
		} else if buf.Len() > 0 {
			result = append(result, buf.String())
			buf.Reset()
		}
	}
	if buf.Len() > 0 {
		result = append(result, buf.String())
	}
	return result
}

func IsNumber(s string) bool {
	if len(s) == 0 {
		return false
	}
	for _, ch := range s {
		if ch < '0' || ch > '9' {
			return false
		}
	}
	return true
}

func IsHexNotation(s string) bool {
	if len(s) == 0 {
		return false
	}
	for _, ch := range s {
		if !unicode.Is(unicode.Hex_Digit, ch) {
			return false
		}
	}
	return true
}

func IsDecimalNotation(s string) bool {
	if len(s) == 0 {
		return false
	}
	for _, ch := range s {
		if !unicode.IsDigit(ch) {
			return false
		}
	}
	return true
}

func HexToBytes(hex string) []byte {
	var bytes []byte
	for i := 0; i+1 < len(hex); i += 2 {
		val, _ := strconv.ParseInt(hex[i:i+2], 16, 8)
		bytes = append(bytes, byte(val))
	}
	return bytes
}

type ConversionResult struct {
	Value uint64
	Ok    bool
}

func SetPrivilege(privilegeName string, enable bool) bool {
	return SetPrivilegeForToken(privilegeName, enable)
}

func IsFileExist(fileName string) bool {
	_, err := os.Stat(fileName)
	return err == nil
}

func IsEmptyString(text string) bool {
	if len(text) == 0 {
		return true
	}
	for _, ch := range text {
		if ch != ' ' && ch != '\t' && ch != '\n' {
			return false
		}
	}
	return true
}

const ConfigFileName = "hyperdbg.config"

func GetConfigFilePath() string {
	exePath, err := os.Executable()
	if err != nil {
		return ""
	}
	dir := filepath.Dir(exePath)
	return filepath.Join(dir, ConfigFileName)
}

func ListDirectory(directory, extension string) []string {
	var result []string
	pattern := filepath.Join(directory, extension)
	matches, err := filepath.Glob(pattern)
	if err != nil {
		return result
	}
	return matches
}

func StringToWString(s string) []uint16 {
	var result []uint16
	for _, r := range s {
		result = append(result, uint16(r))
	}
	result = append(result, 0)
	return result
}

func FindCaseInsensitive(input, toSearch string, pos int) int {
	lowerInput := strings.ToLower(input)
	lowerSearch := strings.ToLower(toSearch)
	idx := strings.Index(lowerInput[pos:], lowerSearch)
	if idx == -1 {
		return -1
	}
	return pos + idx
}

func ConvertStringVectorToCharPointerArray(input []string) []string {
	return input
}

func CheckCpuSupportRtm() bool {
	_, ebx2, _, _ := CpuId(7)
	// Check RTM (ebx bit 11) and HLE (ebx bit 4) from CPUID leaf 7
	return (ebx2>>11)&1 == 1 && (ebx2>>4)&1 == 1
}

const MaxCanonicalAddressWidth = 48

func CheckAddressCanonicality(vAddr uint64) (isCanonical bool, isKernelAddress bool) {
	addrWidth := uint64(MaxCanonicalAddressWidth)
	maxVirtualAddrLowHalf := (uint64(1) << (addrWidth - 1)) - 1
	minVirtualAddressHighHalf := ^maxVirtualAddrLowHalf

	if vAddr > maxVirtualAddrLowHalf && vAddr < minVirtualAddressHighHalf {
		return false, false
	}

	if minVirtualAddressHighHalf < vAddr {
		return true, true
	}
	return true, false
}

func CheckAddressValidityUsingTsx(address uint64) bool {
	// TSX validation not available in Go user-mode
	// Use safe memory access instead
	return false
}

const PageSize = 4096

func CheckAccessValidityAndSafety(targetAddress uint64, size uint32) bool {
	isCanonical, isKernelAddress := CheckAddressCanonicality(targetAddress)
	if !isCanonical {
		return false
	}
	if isKernelAddress {
		return false
	}

	// Cross-page boundary check
	offsetWithinPage := targetAddress + uint64(size) - (targetAddress &^ (PageSize - 1))
	if offsetWithinPage > PageSize {
		remaining := size
		addr := targetAddress
		for remaining > 0 {
			readSize := uint32((addr+PageSize)&^(PageSize-1) - addr)
			if readSize == PageSize && remaining < PageSize {
				readSize = remaining
			}
			// Cannot truly validate in Go user-mode without kernel support
			remaining -= readSize
			addr += uint64(readSize)
		}
	}
	return true
}

func HasEnding(fullString, ending string) bool {
	return strings.HasSuffix(fullString, ending)
}

func ValidateIP(ip string) bool {
	parts := Split(ip, '.')
	if len(parts) != 4 {
		return false
	}
	for _, part := range parts {
		if !IsNumber(part) {
			return false
		}
		val, err := strconv.Atoi(part)
		if err != nil || val < 0 || val > 255 {
			return false
		}
	}
	return true
}

func Ltrim(s string) string {
	return strings.TrimLeftFunc(s, unicode.IsSpace)
}

func Rtrim(s string) string {
	return strings.TrimRightFunc(s, unicode.IsSpace)
}

func Trim(s string) string {
	return strings.TrimSpace(s)
}

func RemoveSpaces(s string) string {
	return strings.ReplaceAll(s, " ", "")
}

func GetLowerStringFromCommandToken(token string) string {
	return strings.ToLower(token)
}

func ConvertStringToUInt32(text string) ConversionResult32 {
	isDecimal := false

	text = strings.TrimSpace(text)
	switch {
	case strings.HasPrefix(text, "0x") || strings.HasPrefix(text, "0X"):
		text = text[2:]
	case strings.HasPrefix(text, "\\x") || strings.HasPrefix(text, "\\X"):
		text = text[2:]
	case strings.HasPrefix(text, "x") || strings.HasPrefix(text, "X"):
		text = text[1:]
	case strings.HasPrefix(text, "0n") || strings.HasPrefix(text, "0N"):
		text = text[2:]
		isDecimal = true
	case strings.HasPrefix(text, "\\n") || strings.HasPrefix(text, "\\N"):
		text = text[2:]
		isDecimal = true
	case strings.HasPrefix(text, "n") || strings.HasPrefix(text, "N"):
		text = text[1:]
		isDecimal = true
	}

	text = strings.ReplaceAll(text, "`", "")
	text = strings.TrimSpace(text)

	if len(text) == 0 {
		return ConversionResult32{Ok: false}
	}

	if isDecimal {
		if !IsDecimalNotation(text) {
			return ConversionResult32{Ok: false}
		}
		val, err := strconv.ParseUint(text, 10, 32)
		if err != nil {
			return ConversionResult32{Ok: false}
		}
		return ConversionResult32{Value: uint32(val), Ok: true}
	} else {
		if !IsHexNotation(text) {
			return ConversionResult32{Ok: false}
		}
		val, err := strconv.ParseUint(text, 16, 32)
		if err != nil {
			return ConversionResult32{Ok: false}
		}
		return ConversionResult32{Value: uint32(val), Ok: true}
	}
}

type ConversionResult32 struct {
	Value uint32
	Ok    bool
}

func Log2Ceil(n uint32) uint32 {
	if n == 0 {
		return 0
	}
	n--
	var log2Floor uint32
	for n = n >> 1; n > 0; n = n >> 1 {
		log2Floor++
	}
	return log2Floor + 1
}

func ConvertTokenToUInt64(token string) ConversionResult {
	return ConvertStringToUInt64(token)
}

func ConvertTokenToUInt32(token string) ConversionResult32 {
	return ConvertStringToUInt32(token)
}

func IsValidIp4(ip string) bool {
	pattern := `^((25[0-5]|2[0-4]\d|[01]?\d\d?)\.){3}(25[0-5]|2[0-4]\d|[01]?\d\d?)$`
	matched, _ := regexp.MatchString(pattern, ip)
	return matched
}

func ConvertStringToUInt64(text string) ConversionResult {
	isDecimal := false

	text = strings.TrimSpace(text)
	switch {
	case strings.HasPrefix(text, "0x") || strings.HasPrefix(text, "0X"):
		text = text[2:]
	case strings.HasPrefix(text, "\\x") || strings.HasPrefix(text, "\\X"):
		text = text[2:]
	case strings.HasPrefix(text, "x") || strings.HasPrefix(text, "X"):
		text = text[1:]
	case strings.HasPrefix(text, "0n") || strings.HasPrefix(text, "0N"):
		text = text[2:]
		isDecimal = true
	case strings.HasPrefix(text, "\\n") || strings.HasPrefix(text, "\\N"):
		text = text[2:]
		isDecimal = true
	case strings.HasPrefix(text, "n") || strings.HasPrefix(text, "N"):
		text = text[1:]
		isDecimal = true
	}

	text = strings.ReplaceAll(text, "`", "")
	text = strings.TrimSpace(text)

	if len(text) == 0 {
		return ConversionResult{Ok: false}
	}

	if isDecimal {
		if !IsDecimalNotation(text) {
			return ConversionResult{Ok: false}
		}
		val, err := strconv.ParseUint(text, 10, 64)
		if err != nil {
			return ConversionResult{Ok: false}
		}
		return ConversionResult{Value: val, Ok: true}
	} else {
		if !IsHexNotation(text) {
			return ConversionResult{Ok: false}
		}
		val, err := strconv.ParseUint(text, 16, 64)
		if err != nil {
			return ConversionResult{Ok: false}
		}
		return ConversionResult{Value: val, Ok: true}
	}
}
