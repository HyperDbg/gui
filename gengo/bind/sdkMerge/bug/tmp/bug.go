// Code generated by gengo. DO NOT EDIT.
package bug

import (
	"unsafe"
	"github.com/can1357/gengo/gengort"
)

const GengoLibraryName = "bug"

var GengoLibrary = gengort.NewLibrary(GengoLibraryName)

type Cr3Type struct {
	Anon46_6
}
type Anon46_6 struct {
	Raw [1]int64
}
type Anon50_10 struct {
	Pcid            Uint64
	PageFrameNumber Uint64
	Reserved1       Uint64
	Reserved_2      Uint64
	PcidInvalidate  Uint64
}
type Anon61_9 struct {
	Raw [1]int32
}
type Anon63_5 struct {
	// [Bits 3:0] Segment type.
	Type Uint32
	// [Bit 4] S - Descriptor type (0 = system; 1 = code or data).
	DescriptorType Uint32
	// [Bits 6:5] DPL - Descriptor privilege level.
	DescriptorPrivilegeLevel Uint32
	// [Bit 7] P - Segment present.
	Present   Uint32
	Reserved1 Uint32
	// [Bit 12] AVL - Available for use by system software.
	AvailableBit Uint32
	// [Bit 13] Reserved (except for CS). L - 64-bit mode active (for CS only).
	LongMode Uint32
	// [Bit 14] D/B - Default operation size (0 = 16-bit segment; 1 = 32-bit segment).
	DefaultBig Uint32
	// [Bit 15] G - Granularity.
	Granularity Uint32
	// [Bit 16] Segment unusable (0 = usable; 1 = unusable).
	Unusable  Uint32
	Reserved2 Uint32
}
type XedImmdisS struct {
	CurrentlyUsedSpace  uint32
	MaxAllocatedSpace   uint32
	Present             int32
	ImmediateIsUnsigned int32
}
type _Int128T = any
type _Uint128T = any
type __NSConstantString = any
type SizeT = uint64
type _BuiltinMsVaList = *byte
type _BuiltinVaList = *byte
type Qword = uint64
type Uint64 = uint64
type Puint64 = *uint64
type Dword = uint64
type Bool = int32
type Byte = uint8
type Word = uint16
type Int = int32
type Uint = uint32
type Puint = *uint32
type Ulong64 = uint64
type Pulong64 = *uint64
type Dword64 = uint64
type Pdword64 = *uint64
type Char = byte
type Wchar = int32
type Uchar = uint8
type Ushort = uint16
type Ulong = uint64
type Boolean = Uchar
type Pboolean = *Boolean
type Int8 = int8
type Pint8 = *int8
type Int16 = int16
type Pint16 = *int16
type Int32 = int32
type Pint32 = *int32
type Int64 = int64
type Pint64 = *int64
type Uint8 = uint8
type Puint8 = *uint8
type Uint16 = uint16
type Puint16 = *uint16
type Uint32 = uint32
type Puint32 = *uint32
//type Uint64 = uint64
//type Puint64 = *uint64
type Pcr3Type = *Cr3Type
type VmxSegmentAccessRightsType = any
type XedImmdisT = XedImmdisS

func (s Anon46_6) Flags() Uint64 {
	return gengort.ReadBitcast[Uint64](unsafe.Add(unsafe.Pointer(unsafe.SliceData(s.Raw[:])), 0))
}
func (s *Anon46_6) SetFlags(v Uint64) {
	gengort.WriteBitcast(unsafe.Add(unsafe.Pointer(unsafe.SliceData(s.Raw[:])), 0), v)
}
func (s Anon46_6) Fields() Anon50_10 {
	return gengort.ReadBitcast[Anon50_10](unsafe.Add(unsafe.Pointer(unsafe.SliceData(s.Raw[:])), 0))
}
func (s *Anon46_6) SetFields(v Anon50_10) {
	gengort.WriteBitcast(unsafe.Add(unsafe.Pointer(unsafe.SliceData(s.Raw[:])), 0), v)
}
func (s Anon61_9)Get () Anon63_5 {
	return gengort.ReadBitcast[Anon63_5](unsafe.Add(unsafe.Pointer(unsafe.SliceData(s.Raw[:])), 0))
}
func (s *Anon61_9) Set(v Anon63_5) {
	gengort.WriteBitcast(unsafe.Add(unsafe.Pointer(unsafe.SliceData(s.Raw[:])), 0), v)
}
func (s Anon61_9) AsUInt() Uint32 {
	return gengort.ReadBitcast[Uint32](unsafe.Add(unsafe.Pointer(unsafe.SliceData(s.Raw[:])), 0))
}
func (s *Anon61_9) SetAsUInt(v Uint32) {
	gengort.WriteBitcast(unsafe.Add(unsafe.Pointer(unsafe.SliceData(s.Raw[:])), 0), v)
}

//  Gengo init function.
func init() {
	gengort.Validate((*Cr3Type)(nil), 0x8, 0x8)
	gengort.Validate((*Anon46_6)(nil), 0x8, 0x8)
	gengort.Validate((*Anon50_10)(nil), 0x8, 0x8, "Pcid", 0xc, "PageFrameNumber", 0x30, "Reserved1", 0x3c, "Reserved_2", 0x3f, "PcidInvalidate", 0x40)
	gengort.Validate((*Anon61_9)(nil), 0x4, 0x4)
	gengort.Validate((*Anon63_5)(nil), 0x4, 0x4, "Type", 0x4, "DescriptorType", 0x5, "DescriptorPrivilegeLevel", 0x7, "Present", 0x8, "Reserved1", 0xc, "AvailableBit", 0xd, "LongMode", 0xe, "DefaultBig", 0xf, "Granularity", 0x10, "Unusable", 0x11, "Reserved2", 0x20)
	gengort.Validate((*XedImmdisS)(nil), 0x4, 0x4, "CurrentlyUsedSpace", 0x4, "MaxAllocatedSpace", 0x8, "Present", 0x9, "ImmediateIsUnsigned", 0xa)
}
