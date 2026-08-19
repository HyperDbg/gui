// Package pt wraps the Intel Processor Trace decoder library (libipt) using
// pure syscalls against libipt.dll — no cgo.
//
// This file owns the Go-side mirror of the C types declared in
// HyperDbg/hyperdbg/dependencies/libipt/intel-pt.h. Struct layouts (field
// order, padding, alignment) match the MSVC x64 ABI so they can be passed
// to libipt.dll by pointer.
package pt

import "unsafe"

// PtCpuVendor — enum pt_cpu_vendor.
//
// Per the task spec the NewDecoder API takes a uint8 vendor; we expose the
// constants as uint8 to match. The on-the-wire struct field (PtCpu.Vendor)
// is uint32 because the C enum is 4 bytes wide.
type PtCpuVendor uint8

const (
	// pcv_unknown.
	PtCpuVendorUnknown PtCpuVendor = 0
	// pcv_intel.
	PtCpuVendorIntel PtCpuVendor = 1
	// Note: libipt's enum pt_cpu_vendor has no amd variant; the constant is
	// reserved here for callers that want a placeholder for non-Intel silicon
	// (libipt will treat it as unknown).
	PtCpuVendorAmd PtCpuVendor = 2
)

// Packet type constants — enum pt_packet_type.
type PtPacketType uint32

const (
	PptInvalid PtPacketType = iota
	PptUnknown
	PptPad
	PptPsb
	PptPsbend
	PptFup
	PptTip
	PptTipPge
	PptTipPgd
	PptTnt8
	PptTnt64
	PptMode
	PptPip
	PptVmcs
	PptCbr
	PptTsc
	PptTma
	PptMtc
	PptCyc
	PptStop
	PptOvf
	PptMnt
	PptExstop
	PptMwait
	PptPwre
	PptPwrx
	PptPtw
	PptCfe
	PptEvd
	PptTrig
)

// Error code constants — enum pt_error_code.
//
// libipt returns negative status codes; the absolute value is the enum value
// below. Use [ptErrcode] to convert a status return.
type PtErrorCode int32

const (
	PteOk              PtErrorCode = 0
	PteInternal        PtErrorCode = 1
	PteInvalid         PtErrorCode = 2
	PteNoSync          PtErrorCode = 3
	PteBadOpc          PtErrorCode = 4
	PteBadPacket       PtErrorCode = 5
	PteBadContext      PtErrorCode = 6
	PteEos             PtErrorCode = 7
	PteBadQuery        PtErrorCode = 8
	PteNoMem           PtErrorCode = 9
	PteBadConfig       PtErrorCode = 10
	PteNoIp            PtErrorCode = 11
	PteIpSuppressed    PtErrorCode = 12
	PteNoMap           PtErrorCode = 13
	PteBadInsn         PtErrorCode = 14
	PteNoTime          PtErrorCode = 15
	PteNoCbr           PtErrorCode = 16
	PteBadImage        PtErrorCode = 17
	PteBadLock         PtErrorCode = 18
	PteNotSupported    PtErrorCode = 19
	PteRetstackEmpty   PtErrorCode = 20
	PteBadRetcomp      PtErrorCode = 21
	PteBadStatusUpdate PtErrorCode = 22
	PteNoEnable        PtErrorCode = 23
	PteEventIgnored    PtErrorCode = 24
	PteOverflow        PtErrorCode = 25
	PteBadFile         PtErrorCode = 26
	PteBadCpu          PtErrorCode = 27
)

// IP compression — enum pt_ip_compression.
type PtIpCompression uint32

const (
	PtIpcSuppressed PtIpCompression = 0x0
	PtIpcUpdate16   PtIpCompression = 0x01
	PtIpcUpdate32   PtIpCompression = 0x02
	PtIpcSext48     PtIpCompression = 0x03
	PtIpcUpdate48   PtIpCompression = 0x04
	PtIpcFull       PtIpCompression = 0x06
)

// Execution mode — enum pt_exec_mode.
type PtExecMode uint32

const (
	PtemUnknown PtExecMode = iota
	Ptem16Bit
	Ptem32Bit
	Ptem64Bit
)

// Mode leaf — enum pt_mode_leaf.
type PtModeLeaf uint32

const (
	PtMolExec PtModeLeaf = 0x00
	PtMolTsx  PtModeLeaf = 0x20
)

// Decoder status flags — enum pt_status_flag.
const (
	PtsEventPending uint32 = 1 << 0
	PtsIpSuppressed uint32 = 1 << 1
	PtsEos          uint32 = 1 << 2
)

// PtCpu mirrors struct pt_cpu (8 bytes, alignment 4).
//
//   - vendor  : enum pt_cpu_vendor (uint32)
//   - family  : uint16
//   - model   : uint8
//   - stepping: uint8
type PtCpu struct {
	Vendor   uint32
	Family   uint16
	Model    uint8
	Stepping uint8
}

// PtErrata mirrors struct pt_errata (64 bytes, alignment 4).
//
// The C struct packs 10 single-bit fields into one uint32_t followed by
// uint32_t reserved[15]. We treat the bitfield word as opaque.
type PtErrata struct {
	Bits     uint32
	Reserved [15]uint32
}

// PtConfFlags mirrors struct pt_conf_flags (16 bytes, alignment 4).
//
// The C type is a union of decoder-variant bitfield structs with a
// uint32_t reserved[4] alternative. We expose the raw 16 bytes as a
// [4]uint32 to preserve 4-byte alignment.
type PtConfFlags struct {
	Variant [4]uint32
}

// PtConfAddrFilter mirrors struct pt_conf_addr_filter (136 bytes, alignment 8).
type PtConfAddrFilter struct {
	Config   uint64 // union of addr_cfg / ctl bitfields
	Addr0A   uint64
	Addr0B   uint64
	Addr1A   uint64
	Addr1B   uint64
	Addr2A   uint64
	Addr2B   uint64
	Addr3A   uint64
	Addr3B   uint64
	Reserved [8]uint64
}

// PtConfig mirrors struct pt_config (280 bytes, alignment 8).
//
// Layout (offsets in bytes):
//
//	  0  Size             uintptr
//	  8  Begin            *byte
//	 16  End              *byte
//	 24  DecodeCallback   uintptr  (function pointer)
//	 32  DecodeContext    unsafe.Pointer
//	 40  Cpu              PtCpu        (8 bytes)
//	 48  Errata           PtErrata     (64 bytes)
//	112  Cpuid0x15Eax     uint32
//	116  Cpuid0x15Ebx     uint32
//	120  MtcFreq          uint8
//	121  NomFreq          uint8
//	124  Flags            PtConfFlags  (16 bytes, 4-byte aligned)
//	144  AddrFilter       PtConfAddrFilter (136 bytes, 8-byte aligned)
//	280  (end)
type PtConfig struct {
	Size           uintptr
	Begin          *byte
	End            *byte
	DecodeCallback uintptr
	DecodeContext  unsafe.Pointer
	Cpu            PtCpu
	Errata         PtErrata
	Cpuid0x15Eax   uint32
	Cpuid0x15Ebx   uint32
	MtcFreq        uint8
	NomFreq        uint8
	Flags          PtConfFlags
	AddrFilter     PtConfAddrFilter
}

// PtAsid mirrors struct pt_asid (24 bytes, alignment 8).
type PtAsid struct {
	Size uintptr
	Cr3  uint64
	Vmcs uint64
}

// PtAsidNoCr3 / PtAsidNoVmcs mirror the pt_asid_no_cr3 / pt_asid_no_vmcs
// sentinels from intel-pt.h.
const (
	PtAsidNoCr3  uint64 = 0xffffffffffffffff
	PtAsidNoVmcs uint64 = 0xffffffffffffffff
)

// AsidInit initializes asid with size and sentinel cr3/vmcs values —
// mirrors the static inline pt_asid_init().
func AsidInit(asid *PtAsid) {
	asid.Size = unsafe.Sizeof(*asid)
	asid.Cr3 = PtAsidNoCr3
	asid.Vmcs = PtAsidNoVmcs
}

// ---- Packet payload structs (each 16 bytes to match the C union) ----

// PtPacketTnt mirrors struct pt_packet_tnt (16 bytes inside the union).
type PtPacketTnt struct {
	BitSize uint8
	_       [7]byte
	Payload uint64
}

// PtPacketIp mirrors struct pt_packet_ip (16 bytes inside the union).
type PtPacketIp struct {
	Ipc PtIpCompression
	_   [4]byte
	Ip  uint64
}

// PtPacketModeExec mirrors struct pt_packet_mode_exec.
type PtPacketModeExec struct {
	Bits uint32 // csl:1 | csd:1 | iflag:1 packed
}

// PtPacketModeTsx mirrors struct pt_packet_mode_tsx.
type PtPacketModeTsx struct {
	Bits uint32 // intx:1 | abrt:1 packed
}

// PtPacketMode mirrors struct pt_packet_mode (8 bytes inside the union).
type PtPacketMode struct {
	Leaf PtModeLeaf
	Bits uint32
}

// PtPacketPip mirrors struct pt_packet_pip (16 bytes inside the union).
type PtPacketPip struct {
	Cr3 uint64
	Nr  uint32 // bitfield, but exposed as uint32 storage
	_   [4]byte
}

// PtPacketTsc mirrors struct pt_packet_tsc (16 bytes inside the union).
type PtPacketTsc struct {
	Tsc uint64
	_   [8]byte
}

// PtPacketCbr mirrors struct pt_packet_cbr (16 bytes inside the union).
type PtPacketCbr struct {
	Ratio uint8
	_     [15]byte
}

// PtPacketTma mirrors struct pt_packet_tma (16 bytes inside the union).
type PtPacketTma struct {
	Ctc uint16
	Fc  uint16
	_   [12]byte
}

// PtPacketMtc mirrors struct pt_packet_mtc (16 bytes inside the union).
type PtPacketMtc struct {
	Ctc uint8
	_   [15]byte
}

// PtPacketCyc mirrors struct pt_packet_cyc (16 bytes inside the union).
type PtPacketCyc struct {
	Value uint64
	_     [8]byte
}

// PtPacketVmcs mirrors struct pt_packet_vmcs (16 bytes inside the union).
type PtPacketVmcs struct {
	Base uint64
	_    [8]byte
}

// PtPacketMnt mirrors struct pt_packet_mnt (16 bytes inside the union).
type PtPacketMnt struct {
	Payload uint64
	_       [8]byte
}

// PtPacketExstop mirrors struct pt_packet_exstop (16 bytes inside the union).
type PtPacketExstop struct {
	Bits uint32 // ip:1 packed
	_    [12]byte
}

// PtPacketMwait mirrors struct pt_packet_mwait (16 bytes inside the union).
type PtPacketMwait struct {
	Hints uint32
	Ext   uint32
	_     [8]byte
}

// PtPacketPwre mirrors struct pt_packet_pwre (16 bytes inside the union).
type PtPacketPwre struct {
	State    uint8
	SubState uint8
	_        [2]byte
	Bits     uint32 // hw:1 packed
	_        [8]byte
}

// PtPacketPwrx mirrors struct pt_packet_pwrx (16 bytes inside the union).
type PtPacketPwrx struct {
	Last    uint8
	Deepest uint8
	_       [2]byte
	Bits    uint32 // interrupt:1 | store:1 | autonomous:1 packed
	_       [8]byte
}

// PtPacketPtw mirrors struct pt_packet_ptw (16 bytes inside the union).
type PtPacketPtw struct {
	Payload uint64
	Plc     uint8
	_       [3]byte
	Bits    uint32 // ip:1 packed
}

// PtPacketUnknown mirrors struct pt_packet_unknown (16 bytes inside the union).
type PtPacketUnknown struct {
	Packet *byte
	Priv   unsafe.Pointer
}

// PtPacket mirrors struct pt_packet (24 bytes, alignment 8).
//
// Layout:
//
//	 0  Type    uint32  (enum pt_packet_type)
//	 4  Size    uint8
//	 5  _       [3]byte padding (payload union has 8-byte alignment)
//	 8  Payload [16]byte (raw union of all pt_packet_* payload structs)
//	24  (end)
type PtPacket struct {
	Type    PtPacketType
	Size    uint8
	_       [3]byte
	Payload [16]byte
}

// Packet is an alias for PtPacket so callers can refer to it as either name.
type Packet = PtPacket

// ---- Payload accessors ----
//
// Each accessor reinterprets the 16-byte Payload buffer as the corresponding
// C payload struct. The caller is responsible for picking the right accessor
// based on [PtPacket.Type].

// Tnt returns the payload interpreted as a TNT-8/TNT-64 packet.
func (p *PtPacket) Tnt() PtPacketTnt {
	return *(*PtPacketTnt)(unsafe.Pointer(&p.Payload[0]))
}

// Ip returns the payload interpreted as a TIP/FUP/TIP.PGE/TIP.PGD packet.
func (p *PtPacket) Ip() PtPacketIp {
	return *(*PtPacketIp)(unsafe.Pointer(&p.Payload[0]))
}

// Mode returns the payload interpreted as a MODE packet.
func (p *PtPacket) Mode() PtPacketMode {
	return *(*PtPacketMode)(unsafe.Pointer(&p.Payload[0]))
}

// Pip returns the payload interpreted as a PIP packet.
func (p *PtPacket) Pip() PtPacketPip {
	return *(*PtPacketPip)(unsafe.Pointer(&p.Payload[0]))
}

// Tsc returns the payload interpreted as a TSC packet.
func (p *PtPacket) Tsc() PtPacketTsc {
	return *(*PtPacketTsc)(unsafe.Pointer(&p.Payload[0]))
}

// Cbr returns the payload interpreted as a CBR packet.
func (p *PtPacket) Cbr() PtPacketCbr {
	return *(*PtPacketCbr)(unsafe.Pointer(&p.Payload[0]))
}

// Tma returns the payload interpreted as a TMA packet.
func (p *PtPacket) Tma() PtPacketTma {
	return *(*PtPacketTma)(unsafe.Pointer(&p.Payload[0]))
}

// Mtc returns the payload interpreted as an MTC packet.
func (p *PtPacket) Mtc() PtPacketMtc {
	return *(*PtPacketMtc)(unsafe.Pointer(&p.Payload[0]))
}

// Cyc returns the payload interpreted as a CYC packet.
func (p *PtPacket) Cyc() PtPacketCyc {
	return *(*PtPacketCyc)(unsafe.Pointer(&p.Payload[0]))
}

// Vmcs returns the payload interpreted as a VMCS packet.
func (p *PtPacket) Vmcs() PtPacketVmcs {
	return *(*PtPacketVmcs)(unsafe.Pointer(&p.Payload[0]))
}

// Mnt returns the payload interpreted as an MNT packet.
func (p *PtPacket) Mnt() PtPacketMnt {
	return *(*PtPacketMnt)(unsafe.Pointer(&p.Payload[0]))
}

// Exstop returns the payload interpreted as an EXSTOP packet.
func (p *PtPacket) Exstop() PtPacketExstop {
	return *(*PtPacketExstop)(unsafe.Pointer(&p.Payload[0]))
}

// Mwait returns the payload interpreted as an MWAIT packet.
func (p *PtPacket) Mwait() PtPacketMwait {
	return *(*PtPacketMwait)(unsafe.Pointer(&p.Payload[0]))
}

// Pwre returns the payload interpreted as a PWRE packet.
func (p *PtPacket) Pwre() PtPacketPwre {
	return *(*PtPacketPwre)(unsafe.Pointer(&p.Payload[0]))
}

// Pwrx returns the payload interpreted as a PWRX packet.
func (p *PtPacket) Pwrx() PtPacketPwrx {
	return *(*PtPacketPwrx)(unsafe.Pointer(&p.Payload[0]))
}

// Ptw returns the payload interpreted as a PTW packet.
func (p *PtPacket) Ptw() PtPacketPtw {
	return *(*PtPacketPtw)(unsafe.Pointer(&p.Payload[0]))
}

// Unknown returns the payload interpreted as an UNKNOWN packet (callback
// payload).
func (p *PtPacket) Unknown() PtPacketUnknown {
	return *(*PtPacketUnknown)(unsafe.Pointer(&p.Payload[0]))
}
