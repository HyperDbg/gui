// Package symbolparser resolves symbol names to addresses and back, using the
// Windows Debug Help API (DbgHelp.dll) via pure syscalls (no cgo).
//
// This mirrors the C++ symbol-parser/ module but replaces the DIA/COM
// implementation with DbgHelp, which is shipped with Windows and supports
// the same Microsoft Symbol Server download flow. The 32/64-bit PDB
// distinction is handled by SymLoadModuleEx's SLMFLAG_XXX flags.
//
// The primary consumer is the api.Debugger.SymbolResolver interface, which
// lets EptHook accept "ntdll!RtlAllocateHeap" style symbol strings.
package symbolparser

import (
	"context"
	"fmt"
	"sync"
	"unsafe"

	"golang.org/x/sys/windows"
)

// Resolver is the symbol-resolution interface consumed by the api layer.
// Implementations must be goroutine-safe; DbgHelp itself is not, so the
// default implementation holds a mutex around every call.
type Resolver interface {
	// Init opens a DbgHelp context with the given symbol search path (e.g.
	// "srv*c:\\symbols*https://msdl.microsoft.com/download/symbols").
	Init(ctx context.Context, sympath string) error

	// LoadModule loads a PE module's symbols. base is the load address (as
	// reported by the loader or the debugger). Returns the module's symbol
	// base on success.
	LoadModule(ctx context.Context, imagePath string, base uint64) (uint64, error)

	// FromName resolves "module!symbol" or "symbol" to an address.
	FromName(ctx context.Context, name string) (uint64, error)

	// FromAddr resolves an address to "module!symbol+offset" (offset is 0
	// when addr is exactly at a symbol).
	FromAddr(ctx context.Context, addr uint64) (name string, offset uint64, err error)

	// Close releases the DbgHelp context.
	Close() error
}

// dbghelp is the default Resolver implementation backed by DbgHelp.dll.
type dbghelp struct {
	mu     sync.Mutex
	handle windows.Handle // SymInitialize handle (current process pseudo-handle)
	loaded bool
}

// New creates a DbgHelp-backed Resolver.
func New() Resolver {
	return &dbghelp{}
}

// DbgHelp.dll function names.
const (
	procSymInitializeW      = "SymInitializeW"
	procSymSetSearchPathW   = "SymSetSearchPathW"
	procSymCleanup          = "SymCleanup"
	procSymLoadModuleExW    = "SymLoadModuleExW"
	procSymFromNameW        = "SymFromNameW"
	procSymFromAddrW        = "SymFromAddrW"
	procSymGetModuleInfoW64 = "SymGetModuleInfoW64"
)

// Load flags / values.
const (
	slmflagVirtual = 0x00000001
	dbgModuleSize  = 0x10000 // sizeof(IMAGEHLP_MODULEW64)
)

var dbghelpDll = windows.NewLazySystemDLL("dbghelp.dll")

// SYMBOL_INFOW is the variable-length struct returned by SymFromNameW /
// SymFromAddrW. We allocate a buffer large enough for a 1KB symbol name.
const symbolInfoBufSize = unsafe.Sizeof(symbolInfo{}) + 1024*2 // 2KB name (UTF-16)

// symbolInfo mirrors the Win32 SYMBOL_INFOW layout. The SizeOfImage field
// is at the end before the variable-length Name[] array.
type symbolInfo struct {
	SizeOfStruct uint32
	TypeIndex    uint32
	Reserved     [2]uint64
	Index        uint32
	Size         uint32
	ModBase      uint64
	Flags        uint32
	Value        uint64
	Address      uint64
	Register     uint32
	Scope        uint32
	Tag          uint32
	NameLen      uint32
	MaxNameLen   uint32
	// Name [MaxNameLen]uint16 follows
}

// imagehlpModuleW64 mirrors IMAGEHLP_MODULEW64 (used by SymGetModuleInfoW64).
type imagehlpModuleW64 struct {
	SizeOfStruct    uint32
	BaseOfImage     uint64
	ImageSize       uint32
	TimeDateStamp   uint32
	CheckSum        uint32
	NumSyms         uint32
	SymType         uint32
	ModuleName      [32]uint16
	ImageName       [256]uint16
	LoadedImageName [256]uint16
	LoadedPdbName   [256]uint16
	CVSig           uint32
	CV              [780]uint16 // MAX_PATH*780 + padding
	CVData          uint32      // pointer-sized offset (unused here)
	PDBSig          uint32
	PDBSig70        [16]byte // GUID
	PDBAge          uint32
	PDBUnmatched    uint32
	ImageUnmatched  uint32
	PdbTypeMismatch uint32
	PrivateBuild    [256]uint16
	Reserved        [256]uint16
}

func (d *dbghelp) Init(ctx context.Context, sympath string) error {
	d.mu.Lock()
	defer d.mu.Unlock()
	if d.loaded {
		return fmt.Errorf("symbolparser: already initialised")
	}
	// Use the current process pseudo-handle. DbgHelp scopes symbol state to
	// this handle; we never call SymRegisterCallback64 here (no progress UI).
	d.handle = windows.CurrentProcess()
	proc := dbghelpDll.NewProc(procSymInitializeW)
	r1, _, err := proc.Call(
		uintptr(d.handle),
		0, // sympath passed via SymSetSearchPathW below for clarity
		0, // fInvadeProcess = FALSE (we load modules explicitly)
	)
	if r1 == 0 {
		return fmt.Errorf("SymInitializeW failed: %w", err)
	}
	if sympath != "" {
		sp, e := windows.UTF16PtrFromString(sympath)
		if e != nil {
			return e
		}
		proc2 := dbghelpDll.NewProc(procSymSetSearchPathW)
		r2, _, err2 := proc2.Call(uintptr(d.handle), uintptr(unsafe.Pointer(sp)))
		if r2 == 0 {
			return fmt.Errorf("SymSetSearchPathW failed: %w", err2)
		}
	}
	d.loaded = true
	return nil
}

func (d *dbghelp) LoadModule(ctx context.Context, imagePath string, base uint64) (uint64, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	if !d.loaded {
		return 0, fmt.Errorf("symbolparser: not initialised")
	}
	img, err := windows.UTF16PtrFromString(imagePath)
	if err != nil {
		return 0, err
	}
	proc := dbghelpDll.NewProc(procSymLoadModuleExW)
	r1, _, err := proc.Call(
		uintptr(d.handle),
		0,                            // hFile (NULL = load from image path)
		0,                            // ImageName (NULL = use imagePath)
		uintptr(unsafe.Pointer(img)), // ModuleName (path used for symbol lookup)
		uintptr(base),                // BaseOfDll
		0,                            // DllSize (0 = look up from PE header)
		0,                            // Data
		0,                            // Flags (SLMFLAG_NONE)
	)
	if r1 == 0 {
		return 0, fmt.Errorf("SymLoadModuleExW failed for %q: %w", imagePath, err)
	}
	return uint64(r1), nil
}

func (d *dbghelp) FromName(ctx context.Context, name string) (uint64, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	if !d.loaded {
		return 0, fmt.Errorf("symbolparser: not initialised")
	}
	nameW, err := windows.UTF16PtrFromString(name)
	if err != nil {
		return 0, err
	}
	// Allocate a buffer for the variable-length SYMBOL_INFOW.
	var buf [symbolInfoBufSize]byte
	si := (*symbolInfo)(unsafe.Pointer(&buf[0]))
	si.SizeOfStruct = uint32(unsafe.Sizeof(symbolInfo{}))
	si.MaxNameLen = 1024

	proc := dbghelpDll.NewProc(procSymFromNameW)
	r1, _, err := proc.Call(
		uintptr(d.handle),
		uintptr(unsafe.Pointer(nameW)),
		uintptr(unsafe.Pointer(si)),
	)
	if r1 == 0 {
		return 0, fmt.Errorf("SymFromNameW(%q) failed: %w", name, err)
	}
	return si.Address, nil
}

func (d *dbghelp) FromAddr(ctx context.Context, addr uint64) (string, uint64, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	if !d.loaded {
		return "", 0, fmt.Errorf("symbolparser: not initialised")
	}
	var buf [symbolInfoBufSize]byte
	si := (*symbolInfo)(unsafe.Pointer(&buf[0]))
	si.SizeOfStruct = uint32(unsafe.Sizeof(symbolInfo{}))
	si.MaxNameLen = 1024

	var displacement uint64
	proc := dbghelpDll.NewProc(procSymFromAddrW)
	r1, _, err := proc.Call(
		uintptr(d.handle),
		uintptr(addr),
		uintptr(unsafe.Pointer(&displacement)),
		uintptr(unsafe.Pointer(si)),
	)
	if r1 == 0 {
		return "", 0, fmt.Errorf("SymFromAddrW(0x%x) failed: %w", addr, err)
	}
	// Read the variable-length Name (UTF-16, null-terminated).
	nameStart := unsafe.Pointer(&buf[unsafe.Sizeof(symbolInfo{})])
	name := windows.UTF16PtrToString((*uint16)(nameStart))
	return name, displacement, nil
}

func (d *dbghelp) Close() error {
	d.mu.Lock()
	defer d.mu.Unlock()
	if !d.loaded {
		return nil
	}
	proc := dbghelpDll.NewProc(procSymCleanup)
	_, _, _ = proc.Call(uintptr(d.handle))
	d.loaded = false
	return nil
}
