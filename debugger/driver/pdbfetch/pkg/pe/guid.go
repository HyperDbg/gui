package pe

import (
	"encoding"
	"encoding/binary"
	"errors"
	"fmt"
	"strconv"

	"github.com/ddkwork/golibrary/std/mylog"
)

// GUID represents a GUID/UUID. It has the same structure as
// golang.org/x/sys/windows.GUID, without the need for golang.org/x/sys/windows
// as a dependency to allow compilation on linux. It is also so that it can
// be used with functions expecting that type. It is defined as its own type
// so that stringification and marshaling can be supported. The representation
// matches that used by native Windows code.
type GUID struct {
	Data1 uint32
	Data2 uint16
	Data3 uint16
	Data4 [8]byte
}

var (
	_ = (encoding.TextMarshaler)(GUID{})
	_ = (encoding.TextUnmarshaler)(&GUID{})
)

func fromArray(b [16]byte, order binary.ByteOrder) GUID {
	var g GUID
	g.Data1 = order.Uint32(b[0:4])
	g.Data2 = order.Uint16(b[4:6])
	g.Data3 = order.Uint16(b[6:8])
	copy(g.Data4[:], b[8:16])
	return g
}

func (g GUID) toArray(order binary.ByteOrder) [16]byte {
	b := [16]byte{}
	order.PutUint32(b[0:4], g.Data1)
	order.PutUint16(b[4:6], g.Data2)
	order.PutUint16(b[6:8], g.Data3)
	copy(b[8:16], g.Data4[:])
	return b
}

// FromArray constructs a GUID from a big-endian encoding array of 16 bytes.
func GuidFromArray(b [16]byte) GUID {
	return fromArray(b, binary.BigEndian)
}

// ToArray returns an array of 16 bytes representing the GUID in big-endian
// encoding.
func (g GUID) ToArray() [16]byte {
	return g.toArray(binary.BigEndian)
}

// FromWindowsArray constructs a GUID from a Windows encoding array of bytes.
func GuidFromWindowsArray(b [16]byte) GUID {
	return fromArray(b, binary.LittleEndian)
}

// ToWindowsArray returns an array of 16 bytes representing the GUID in Windows encoding.
func (g GUID) ToWindowsArray() [16]byte {
	return g.toArray(binary.LittleEndian)
}

// ToString returns a string representation of the value of this instance of the Guid structure.
// A single format specifier can be used to indicate how to format the value of this Guid.
// The format parameter can be "N", "D", "B", "P", or "X". If format is an empty string (""), "D" is used.
func (g GUID) ToString(format string) (string, error) {
	if format == "" || format == "D" {
		return fmt.Sprintf(
			"%08x-%04x-%04x-%04x-%012x",
			g.Data1,
			g.Data2,
			g.Data3,
			g.Data4[:2],
			g.Data4[2:]), nil
	}
	if format == "N" {
		return fmt.Sprintf(
			"%08x%04x%04x%04x%012x",
			g.Data1,
			g.Data2,
			g.Data3,
			g.Data4[:2],
			g.Data4[2:]), nil
	}
	if format == "B" {
		return fmt.Sprintf(
			"{%08x-%04x-%04x-%04x-%012x}",
			g.Data1,
			g.Data2,
			g.Data3,
			g.Data4[:2],
			g.Data4[2:]), nil
	}
	if format == "P" {
		return fmt.Sprintf(
			"(%08x-%04x-%04x-%04x-%012x)",
			g.Data1,
			g.Data2,
			g.Data3,
			g.Data4[:2],
			g.Data4[2:]), nil
	}
	if format == "X" {
		return fmt.Sprintf(
			"(%08x-%04x-%04x-%04x-%012x)",
			g.Data1,
			g.Data2,
			g.Data3,
			g.Data4[:2],
			g.Data4[2:]), nil
	}

	return "", errors.New("invalid format specified")
}

func (g GUID) String() string {
	guidStr, _ := g.ToString("")
	return guidStr
}

// FromString parses a string containing a GUID and returns the GUID. The only
// format currently supported is the `xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx`
// format.
func GuidFromString(s string) (GUID, error) {
	if len(s) != 36 {
		return GUID{}, fmt.Errorf("invalid GUID %q", s)
	}
	if s[8] != '-' || s[13] != '-' || s[18] != '-' || s[23] != '-' {
		return GUID{}, fmt.Errorf("invalid GUID %q", s)
	}

	var g GUID

	data1 := mylog.Check2(strconv.ParseUint(s[0:8], 16, 32))

	g.Data1 = uint32(data1)

	data2 := mylog.Check2(strconv.ParseUint(s[9:13], 16, 16))

	g.Data2 = uint16(data2)

	data3 := mylog.Check2(strconv.ParseUint(s[14:18], 16, 16))

	g.Data3 = uint16(data3)

	for i, x := range []int{19, 21, 24, 26, 28, 30, 32, 34} {
		v := mylog.Check2(strconv.ParseUint(s[x:x+2], 16, 8))

		g.Data4[i] = uint8(v)
	}

	return g, nil
}

// MarshalText returns the textual representation of the GUID.
func (g GUID) MarshalText() ([]byte, error) {
	return []byte(g.String()), nil
}

// UnmarshalText takes the textual representation of a GUID, and unmarhals it
// into this GUID.
func (g *GUID) UnmarshalText(text []byte) error {
	g2 := mylog.Check2(GuidFromString(string(text)))

	*g = g2
	return nil
}
