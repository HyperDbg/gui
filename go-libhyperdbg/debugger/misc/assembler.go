// misc/assembler.go wraps ok/keystone to provide a simple Assemble function
// for the `a` command. Mirrors the C++ Assemble helper in libhyperdbg.
package misc

import (
	"fmt"
	"unsafe"

	"github.com/ddkwork/keystone"
)

// AsmMode selects 16/32/64-bit assembly. Maps to keystone KsMode16/32/64.
type AsmMode uint32

const (
	AsmMode16 AsmMode = AsmMode(keystone.KsMode16)
	AsmMode32 AsmMode = AsmMode(keystone.KsMode32)
	AsmMode64 AsmMode = AsmMode(keystone.KsMode64)
)

// Assembler wraps a keystone.Keystone instance. Like Disassembler, it is
// single-goroutine; concurrent consumers should create their own.
type Assembler struct {
	k *keystone.Keystone
}

// NewAssembler creates an Assembler using the default keystone instance.
// The first call triggers keystone.dll extraction (see ok/keystone init()).
func NewAssembler() *Assembler {
	return &Assembler{k: &keystone.Keystone{}}
}

// Assemble compiles a single assembly statement (e.g. "nop", "mov rax, 1")
// to machine bytes. addr is the runtime address passed to keystone for
// relative-branch resolution.
func (a *Assembler) Assemble(mode AsmMode, addr uint64, statement string) ([]byte, error) {
	var ks *keystone.Ks_engine
	if err := a.k.KsOpen(keystone.KsArchX86, int32(mode), &ks); err != keystone.KsErrOk {
		return nil, fmt.Errorf("Assemble: KsOpen failed: %v", err)
	}
	defer a.k.KsClose(ks)

	// keystone expects a null-terminated C string. UTF16PtrFromString is not
	// suitable here; build a NUL-terminated byte slice instead.
	cstr := append([]byte(statement), 0)
	var encoding *uint8
	var encSize uintptr
	var statCount uintptr
	rc := a.k.KsAsm(ks,
		(*int8)(unsafe.Pointer(&cstr[0])),
		addr,
		&encoding, &encSize, &statCount)
	if rc != 0 {
		errStr := "unknown"
		if s := a.k.KsStrerror(keystone.Ks_err(rc)); s != nil {
			// KsStrerror returns *int8; convert to Go string via strlen.
			p := uintptr(unsafe.Pointer(s))
			if p != 0 {
				var n int
				for b := *(*byte)(unsafe.Pointer(p)); b != 0; n++ {
					_ = b
				}
				// simpler: just call the helper that reads until NUL
				errStr = cstrToString(s)
			}
		}
		return nil, fmt.Errorf("Assemble: KsAsm failed (rc=%d): %s", rc, errStr)
	}
	if encSize == 0 {
		return nil, fmt.Errorf("Assemble: empty encoding")
	}
	// Copy out before KsClose frees the buffer.
	out := make([]byte, int(encSize))
	copy(out, (*[1 << 30]byte)(unsafe.Pointer(encoding))[:encSize:encSize])
	return out, nil
}

// cstrToString converts a *int8 C string to a Go string.
func cstrToString(p *int8) string {
	if p == nil {
		return ""
	}
	var n int
	ptr := unsafe.Pointer(p)
	for *(*int8)(ptr) != 0 {
		n++
		ptr = unsafe.Pointer(uintptr(ptr) + 1)
	}
	buf := make([]byte, n)
	for i := 0; i < n; i++ {
		buf[i] = *(*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + uintptr(i)))
	}
	return string(buf)
}
