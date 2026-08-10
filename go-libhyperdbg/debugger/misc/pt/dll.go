// Package pt — dll.go owns the libipt.dll loader and the lazyProc table.
//
// We deliberately do NOT embed libipt.dll the way ok/zydis and ok/keystone do.
// libipt.dll ships with the HyperDbg driver package and is installed alongside
// the HyperDbg bootloader; the runtime process is expected to find it via the
// standard DLL search path. If libipt.dll is not on the path, every call will
// return a *LazyProc error at first invocation.
package pt

import (
	"unsafe"

	"golang.org/x/sys/windows"
)

// libipt holds the handle to libipt.dll. We use NewLazyDLL so the load is
// deferred until the first proc lookup; this means importing the pt package
// is cheap even when no PT decoding is performed.
var libipt = windows.NewLazyDLL("libipt.dll")

// procCache memoises *windows.LazyProc lookups by name. libipt exports a
// stable C ABI so the proc addresses never change for a given DLL load.
var procCache = make(map[string]*windows.LazyProc)

// getProc returns the *LazyProc for the named libipt export, caching the
// lookup. Callers should call this at call time (not init time) so that
// missing-symbol errors surface at the call site rather than at package
// import.
func getProc(name string) *windows.LazyProc {
	if p, ok := procCache[name]; ok {
		return p
	}
	p := libipt.NewProc(name)
	procCache[name] = p
	return p
}

// ---- libipt function bindings ----
//
// Each binding is a thin wrapper that performs the syscall and returns the
// raw int status from libipt. Callers convert non-negative status to success
// via [ptErrcode] (negative => error).
//
// Exported libipt functions used here (signatures from intel-pt.h):
//
//	struct pt_packet_decoder *
//	pt_pkt_alloc_decoder(const struct pt_config *config);
//
//	void
//	pt_pkt_free_decoder(struct pt_packet_decoder *decoder);
//
//	int
//	pt_pkt_sync_forward(struct pt_packet_decoder *decoder);
//
//	int
//	pt_pkt_sync_backward(struct pt_packet_decoder *decoder);
//
//	int
//	pt_pkt_sync_set(struct pt_packet_decoder *decoder, uint64_t offset);
//
//	int
//	pt_pkt_get_offset(const struct pt_packet_decoder *decoder,
//	                  uint64_t *offset);
//
//	int
//	pt_pkt_get_sync_offset(const struct pt_packet_decoder *decoder,
//	                       uint64_t *offset);
//
//	const struct pt_config *
//	pt_pkt_get_config(const struct pt_packet_decoder *decoder);
//
//	int
//	pt_pkt_next(struct pt_packet_decoder *decoder,
//	            struct pt_packet *packet, size_t size);
//
//	int
//	pt_cpu_errata(struct pt_errata *errata, const struct pt_cpu *cpu);
//
//	const char *
//	pt_errstr(enum pt_error_code err);
//
//	struct pt_version
//	pt_library_version(void);

// ptPktAllocDecoder wraps pt_pkt_alloc_decoder.
func ptPktAllocDecoder(config *PtConfig) uintptr {
	r, _, _ := getProc("pt_pkt_alloc_decoder").Call(uintptr(unsafe.Pointer(config)))
	return r
}

// ptPktFreeDecoder wraps pt_pkt_free_decoder.
func ptPktFreeDecoder(dec uintptr) {
	getProc("pt_pkt_free_decoder").Call(dec)
}

// ptPktSyncForward wraps pt_pkt_sync_forward. Returns libipt status.
func ptPktSyncForward(dec uintptr) int {
	r, _, _ := getProc("pt_pkt_sync_forward").Call(dec)
	return int(r)
}

// ptPktSyncBackward wraps pt_pkt_sync_backward. Returns libipt status.
func ptPktSyncBackward(dec uintptr) int {
	r, _, _ := getProc("pt_pkt_sync_backward").Call(dec)
	return int(r)
}

// ptPktSyncSet wraps pt_pkt_sync_set. Returns libipt status.
func ptPktSyncSet(dec uintptr, offset uint64) int {
	r, _, _ := getProc("pt_pkt_sync_set").Call(dec, uintptr(offset))
	return int(r)
}

// ptPktNext wraps pt_pkt_next. Returns libipt status (bytes consumed or
// negative error code).
func ptPktNext(dec uintptr, packet *PtPacket, size uintptr) int {
	r, _, _ := getProc("pt_pkt_next").Call(dec, uintptr(unsafe.Pointer(packet)), size)
	return int(r)
}

// ptPktGetOffset wraps pt_pkt_get_offset. Returns libipt status.
func ptPktGetOffset(dec uintptr, offset *uint64) int {
	r, _, _ := getProc("pt_pkt_get_offset").Call(dec, uintptr(unsafe.Pointer(offset)))
	return int(r)
}

// ptPktGetSyncOffset wraps pt_pkt_get_sync_offset. Returns libipt status.
func ptPktGetSyncOffset(dec uintptr, offset *uint64) int {
	r, _, _ := getProc("pt_pkt_get_sync_offset").Call(dec, uintptr(unsafe.Pointer(offset)))
	return int(r)
}

// ptCpuErrata wraps pt_cpu_errata. Returns libipt status.
func ptCpuErrata(errata *PtErrata, cpu *PtCpu) int {
	r, _, _ := getProc("pt_cpu_errata").Call(uintptr(unsafe.Pointer(errata)), uintptr(unsafe.Pointer(cpu)))
	return int(r)
}

// ptErrstr wraps pt_errstr. Returns a C string pointer (NUL-terminated).
func ptErrstr(err PtErrorCode) *byte {
	r, _, _ := getProc("pt_errstr").Call(uintptr(uint32(err)))
	return (*byte)(unsafe.Pointer(r))
}

// ptLibraryVersion wraps pt_library_version. The returned bytes mirror
// struct pt_version (16 bytes on x64: uint8 major, uint8 minor, uint16
// patch, uint32 build, const char *ext).
func ptLibraryVersion() (major, minor uint8, patch uint16, build uint32, ext string) {
	type ptVersion struct {
		Major uint8
		Minor uint8
		Patch uint16
		Build uint32
		Ext   *byte
	}
	r, _, _ := getProc("pt_library_version").Call()
	if r == 0 {
		return 0, 0, 0, 0, ""
	}
	v := (*ptVersion)(unsafe.Pointer(r))
	major, minor, patch, build = v.Major, v.Minor, v.Patch, v.Build
	if v.Ext != nil {
		// Convert NUL-terminated C string to Go string. We trust libipt
		// to keep the buffer alive for the lifetime of the DLL. Cap the
		// scan at 64 KiB to defend against a malformed pointer.
		var n int
		for {
			b := *(*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(v.Ext)) + uintptr(n)))
			if b == 0 || n >= 1<<16 {
				break
			}
			n++
		}
		buf := make([]byte, n)
		for i := 0; i < n; i++ {
			buf[i] = *(*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(v.Ext)) + uintptr(i)))
		}
		ext = string(buf)
	}
	return
}
