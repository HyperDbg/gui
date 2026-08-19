// Package pt — ipt.go owns the high-level Decoder type that drives libipt's
// packet decoder.
//
// The decoding loop mirrors PtHelperDecodeCorePackets in
// HyperDbg/hyperdbg/libhyperdbg/code/debugger/misc/pt-helper.cpp:
//
//	pt_config_init(&Config);
//	Config.begin = (UINT8 *)Buffer;
//	Config.end   = (UINT8 *)Buffer + Size;
//	Decoder = pt_pkt_alloc_decoder(&Config);
//	for (;;) {
//	    Status = pt_pkt_sync_forward(Decoder);
//	    if (Status < 0) break;
//	    for (;;) {
//	        struct pt_packet Packet;
//	        Status = pt_pkt_next(Decoder, &Packet, sizeof(Packet));
//	        if (Status < 0) break;
//	        ...handle Packet...
//	    }
//	}
//	pt_pkt_free_decoder(Decoder);
package pt

import (
	"errors"
	"fmt"
	"unsafe"
)

// Decoder wraps an opaque libipt pt_packet_decoder*. The zero value is not
// usable; obtain one via [NewDecoder].
//
// A Decoder is not safe for concurrent use — libipt is single-threaded per
// decoder instance. Callers needing parallel decode must allocate one
// Decoder per goroutine.
type Decoder struct {
	// handle is the opaque pt_packet_decoder* returned by
	// pt_pkt_alloc_decoder. We keep it as uintptr so the Go GC does not
	// mistake it for a pointer (it is heap memory owned by libipt).
	handle uintptr

	// buf is retained so the libipt-decoded trace buffer stays alive for
	// the lifetime of the Decoder (libipt stores raw begin/end pointers
	// into it, so it must not be moved or collected).
	buf []byte

	// config is the pt_config we passed to pt_pkt_alloc_decoder. Kept
	// around so callers can inspect it via [Decoder.Config].
	config PtConfig
}

// Named sentinel errors for the libipt status codes that the higher-level
// HyperDbg code actually inspects. Less common codes are surfaced via a
// generic [ptError] wrapping the numeric value.
var (
	ErrInternal        = errors.New("libipt: pte_internal")
	ErrInvalid         = errors.New("libipt: pte_invalid")
	ErrNoSync          = errors.New("libipt: pte_nosync (decoder out of sync)")
	ErrBadOpc          = errors.New("libipt: pte_bad_opc (unknown opcode)")
	ErrBadPacket       = errors.New("libipt: pte_bad_packet (unknown payload)")
	ErrBadContext      = errors.New("libipt: pte_bad_context")
	ErrEos             = errors.New("libipt: pte_eos (end of stream)")
	ErrBadQuery        = errors.New("libipt: pte_bad_query")
	ErrNoMem           = errors.New("libipt: pte_nomem")
	ErrBadConfig       = errors.New("libipt: pte_bad_config")
	ErrNoIp            = errors.New("libipt: pte_noip")
	ErrIpSuppressed    = errors.New("libipt: pte_ip_suppressed")
	ErrNoMap           = errors.New("libipt: pte_nomap")
	ErrBadInsn         = errors.New("libipt: pte_bad_insn")
	ErrNoTime          = errors.New("libipt: pte_no_time")
	ErrNoCbr           = errors.New("libipt: pte_no_cbr")
	ErrBadImage        = errors.New("libipt: pte_bad_image")
	ErrBadLock         = errors.New("libipt: pte_bad_lock")
	ErrNotSupported    = errors.New("libipt: pte_not_supported")
	ErrRetstackEmpty   = errors.New("libipt: pte_retstack_empty")
	ErrBadRetcomp      = errors.New("libipt: pte_bad_retcomp")
	ErrBadStatusUpdate = errors.New("libipt: pte_bad_status_update")
	ErrNoEnable        = errors.New("libipt: pte_no_enable")
	ErrEventIgnored    = errors.New("libipt: pte_event_ignored")
	ErrOverflow        = errors.New("libipt: pte_overflow")
	ErrBadFile         = errors.New("libipt: pte_bad_file")
	ErrBadCpu          = errors.New("libipt: pte_bad_cpu")
)

// ptError is the fallback error type for status codes without a named
// sentinel. It carries the raw libipt status code so callers can match on
// it if they want to handle an error code we didn't pre-declare.
type ptError struct {
	status int // the original (negative) status return from libipt
}

func (e *ptError) Error() string {
	return fmt.Sprintf("libipt: status=%d", e.status)
}

// ptErrcode converts a libipt status return into a Go error. A non-negative
// status means success and yields nil. A negative status is mapped to the
// matching sentinel error from the [Err*] set above, or to a *ptError when
// no sentinel exists.
func ptErrcode(status int) error {
	if status >= 0 {
		return nil
	}
	switch PtErrorCode(-status) {
	case PteInternal:
		return ErrInternal
	case PteInvalid:
		return ErrInvalid
	case PteNoSync:
		return ErrNoSync
	case PteBadOpc:
		return ErrBadOpc
	case PteBadPacket:
		return ErrBadPacket
	case PteBadContext:
		return ErrBadContext
	case PteEos:
		return ErrEos
	case PteBadQuery:
		return ErrBadQuery
	case PteNoMem:
		return ErrNoMem
	case PteBadConfig:
		return ErrBadConfig
	case PteNoIp:
		return ErrNoIp
	case PteIpSuppressed:
		return ErrIpSuppressed
	case PteNoMap:
		return ErrNoMap
	case PteBadInsn:
		return ErrBadInsn
	case PteNoTime:
		return ErrNoTime
	case PteNoCbr:
		return ErrNoCbr
	case PteBadImage:
		return ErrBadImage
	case PteBadLock:
		return ErrBadLock
	case PteNotSupported:
		return ErrNotSupported
	case PteRetstackEmpty:
		return ErrRetstackEmpty
	case PteBadRetcomp:
		return ErrBadRetcomp
	case PteBadStatusUpdate:
		return ErrBadStatusUpdate
	case PteNoEnable:
		return ErrNoEnable
	case PteEventIgnored:
		return ErrEventIgnored
	case PteOverflow:
		return ErrOverflow
	case PteBadFile:
		return ErrBadFile
	case PteBadCpu:
		return ErrBadCpu
	}
	return &ptError{status: status}
}

// configInit zero-initialises config and sets config.Size to sizeof(PtConfig)
// — mirrors the static inline pt_config_init() from intel-pt.h.
func configInit(config *PtConfig) {
	*config = PtConfig{}
	config.Size = unsafe.Sizeof(*config)
}

// NewDecoder allocates a libipt packet decoder over buf.
//
// cpuVendor selects the CPU vendor used for errata selection; pass
// [PtCpuVendorIntel] for Intel silicon and [PtCpuVendorUnknown] when the
// vendor is not known (libipt will skip vendor-specific errata).
//
// The caller must keep buf alive and unmodified for the lifetime of the
// returned Decoder (libipt stores raw begin/end pointers into it). The
// Decoder retains a reference to buf for this purpose.
//
// Mirrors the C sequence:
//
//	pt_config_init(&Config);
//	Config.begin = buf;
//	Config.end   = buf + len(buf);
//	Decoder = pt_pkt_alloc_decoder(&Config);
func NewDecoder(buf []byte, cpuVendor uint8) (*Decoder, error) {
	if len(buf) == 0 {
		return nil, errors.New("pt: NewDecoder: empty trace buffer")
	}

	var d Decoder
	d.buf = buf

	configInit(&d.config)
	d.config.Begin = &buf[0]
	d.config.End = &buf[len(buf)-1]
	d.config.Cpu.Vendor = uint32(cpuVendor)

	// pt_cpu_errata is optional; if it fails the decoder still works
	// (errata just enable workarounds for known silicon bugs). We ignore
	// the status here to match the C PtHelperDecodeCorePackets, which
	// never calls pt_cpu_errata either.
	_ = ptCpuErrata(&d.config.Errata, &d.config.Cpu)

	h := ptPktAllocDecoder(&d.config)
	if h == 0 {
		return nil, errors.New("pt: pt_pkt_alloc_decoder returned NULL")
	}
	d.handle = h
	return &d, nil
}

// Config returns a copy of the pt_config in use by the decoder. Modifications
// to the returned struct do not affect the decoder.
func (d *Decoder) Config() PtConfig { return d.config }

// Handle returns the raw libipt decoder handle. Exposed so power users can
// call libipt functions we have not wrapped yet (e.g. via their own
// syscall shim). Most callers should not need this.
func (d *Decoder) Handle() uintptr { return d.handle }

// SyncForward advances to the next PSB synchronization point in the forward
// direction. Returns nil on success or [ErrEos] when no further sync point
// is found (which signals the end of the trace for a forward-iteration loop).
//
// Mirrors pt_pkt_sync_forward().
func (d *Decoder) SyncForward() error {
	if d.handle == 0 {
		return ErrInvalid
	}
	return ptErrcode(ptPktSyncForward(d.handle))
}

// SyncBackward advances to the next PSB synchronization point in the
// backward direction. Useful for backward iteration over a trace.
//
// Mirrors pt_pkt_sync_backward().
func (d *Decoder) SyncBackward() error {
	if d.handle == 0 {
		return ErrInvalid
	}
	return ptErrcode(ptPktSyncBackward(d.handle))
}

// SyncSet hard-sets the decoder position to offset (must point at a PSB).
//
// Mirrors pt_pkt_sync_set().
func (d *Decoder) SyncSet(offset uint64) error {
	if d.handle == 0 {
		return ErrInvalid
	}
	return ptErrcode(ptPktSyncSet(d.handle, offset))
}

// Offset retrieves the current decoder position within the trace buffer.
//
// Mirrors pt_pkt_get_offset().
func (d *Decoder) Offset() (uint64, error) {
	if d.handle == 0 {
		return 0, ErrInvalid
	}
	var off uint64
	if err := ptErrcode(ptPktGetOffset(d.handle, &off)); err != nil {
		return 0, err
	}
	return off, nil
}

// SyncOffset retrieves the position of the last synchronization point.
//
// Mirrors pt_pkt_get_sync_offset().
func (d *Decoder) SyncOffset() (uint64, error) {
	if d.handle == 0 {
		return 0, ErrInvalid
	}
	var off uint64
	if err := ptErrcode(ptPktGetSyncOffset(d.handle, &off)); err != nil {
		return 0, err
	}
	return off, nil
}

// Next decodes the next packet and advances the decoder. It returns the
// number of bytes consumed alongside the packet, plus any error.
//
// Returns [ErrEos] when the decoder has reached the end of the trace
// stream — callers driving a forward-iteration loop should treat that as a
// termination signal rather than a hard error.
//
// Mirrors pt_pkt_next(decoder, &packet, sizeof(packet)).
func (d *Decoder) Next() (Packet, int, error) {
	if d.handle == 0 {
		return PtPacket{}, 0, ErrInvalid
	}
	var pkt PtPacket
	status := ptPktNext(d.handle, &pkt, unsafe.Sizeof(pkt))
	if err := ptErrcode(status); err != nil {
		return PtPacket{}, 0, err
	}
	return pkt, status, nil
}

// Close releases the underlying libipt decoder. It is safe to call Close on
// a Decoder that has already been closed (subsequent calls are no-ops).
//
// After Close returns the Decoder must not be used again.
func (d *Decoder) Close() {
	if d.handle == 0 {
		return
	}
	ptPktFreeDecoder(d.handle)
	d.handle = 0
}

// DecodeImage is a convenience function that drives the full packet-decode
// loop over ptBuf and returns every decoded packet. textStart/textEnd
// describe the .text section that was captured (mirrors
// PtHelperCaptureImage's outputs) — they are not consumed by libipt's
// packet decoder (only the instruction decoder needs an image) but are
// accepted here so the API mirrors the C PtHelperDecodeCorePackets entry
// point and so callers have a single place to plumb both pieces.
//
// The function loops:
//
//	for {
//	    if SyncForward() == ErrEos: break
//	    for {
//	        pkt, _, err := Next()
//	        if err != nil: break inner
//	        packets = append(packets, pkt)
//	    }
//	}
//
// returning the accumulated packets. If a non-[ErrEos] error is hit, the
// packets decoded so far are returned along with the error.
func DecodeImage(textStart, textEnd uint64, ptBuf []byte) ([]Packet, error) {
	if len(ptBuf) == 0 {
		return nil, errors.New("pt: DecodeImage: empty trace buffer")
	}
	// textStart/textEnd are accepted for API symmetry with the C side and
	// with the instruction-decoder path; the packet decoder doesn't read
	// them. Touch them here so static analysers don't flag them as unused.
	_ = textStart
	_ = textEnd

	dec, err := NewDecoder(ptBuf, uint8(PtCpuVendorIntel))
	if err != nil {
		return nil, err
	}
	defer dec.Close()

	var out []Packet
	for {
		if err := dec.SyncForward(); err != nil {
			if errors.Is(err, ErrEos) {
				return out, nil
			}
			return out, err
		}
		for {
			pkt, _, err := dec.Next()
			if err != nil {
				if errors.Is(err, ErrEos) {
					break
				}
				return out, err
			}
			out = append(out, pkt)
		}
	}
}

// Compile-time assertion that PtConfig and PtPacket have the sizes we
// documented (libipt reads config.Size and validates it; if our struct
// shrinks below sizeof(struct pt_config) pt_pkt_alloc_decoder will reject
// it).
var (
	_ = [1]struct{}{}[unsafe.Sizeof(PtConfig{})-280]
	_ = [1]struct{}{}[unsafe.Sizeof(PtPacket{})-24]
	_ = [1]struct{}{}[unsafe.Sizeof(PtAsid{})-24]
)
