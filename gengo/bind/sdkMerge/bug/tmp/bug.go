// Code generated by gengo. DO NOT EDIT.
package bug

import (
	"unsafe"

	"github.com/can1357/gengo/gengort"
)

const GengoLibraryName = "bug"

var GengoLibrary = gengort.NewLibrary(GengoLibraryName)

type Cr3Type struct {
	Anon44_6
}
type Anon44_6 struct {
	Raw [1]int64
}
type Anon48_10 struct {
	Pcid            Uint64
	PageFrameNumber Uint64
	Reserved1       Uint64
	Reserved_2      Uint64
	PcidInvalidate  Uint64
}
type (
	_Int128T           = any
	_Uint128T          = any
	__NSConstantString = any
	SizeT              = uint64
	_BuiltinMsVaList   = *byte
	_BuiltinVaList     = *byte
	Qword              = uint64
	Uint64             = uint64
	Puint64            = *uint64
	Dword              = uint64
	Bool               = int32
	Byte               = uint8
	Word               = uint16
	Int                = int32
	Uint               = uint32
	Puint              = *uint32
	Ulong64            = uint64
	Pulong64           = *uint64
	Dword64            = uint64
	Pdword64           = *uint64
	Char               = byte
	Wchar              = int16
	Uchar              = uint8
	Ushort             = uint16
	Ulong              = uint64
	Boolean            = Uchar
	Pboolean           = *Boolean
	Int8               = int8
	Pint8              = *int8
	Int16              = int16
	Pint16             = *int16
	Int32              = int32
	Pint32             = *int32
	Int64              = int64
	Pint64             = *int64
	Uint8              = uint8
	Puint8             = *uint8
	Uint16             = uint16
	Puint16            = *uint16
	Uint32             = uint32
	Puint32            = *uint32
	Uint64             = uint64
	Puint64            = *uint64
	Pcr3Type           = *Cr3Type
)

func (s Anon44_6) Flags() Uint64 {
	return gengort.ReadBitcast[Uint64](unsafe.Add(unsafe.Pointer(unsafe.SliceData(s.Raw[:])), 0))
}

func (s *Anon44_6) SetFlags(v Uint64) {
	gengort.WriteBitcast(unsafe.Add(unsafe.Pointer(unsafe.SliceData(s.Raw[:])), 0), v)
}

func (s Anon44_6) Fields() Anon48_10 {
	return gengort.ReadBitcast[Anon48_10](unsafe.Add(unsafe.Pointer(unsafe.SliceData(s.Raw[:])), 0))
}

func (s *Anon44_6) SetFields(v Anon48_10) {
	gengort.WriteBitcast(unsafe.Add(unsafe.Pointer(unsafe.SliceData(s.Raw[:])), 0), v)
}

// Gengo init function.
func init() {
	gengort.Validate((*Cr3Type)(nil), 0x8, 0x8)
	gengort.Validate((*Anon44_6)(nil), 0x8, 0x8)
	gengort.Validate((*Anon48_10)(nil), 0x8, 0x8, "Pcid", 0xc, "PageFrameNumber", 0x30, "Reserved1", 0x3c, "Reserved_2", 0x3f, "PcidInvalidate", 0x40)
}
