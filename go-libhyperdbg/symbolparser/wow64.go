// Package symbolparser — wow64.go
//
// WOW64 helpers that bridge 32-bit debuggee processes and 64-bit DbgHelp
// symbol loading. Mirrors the C++ symbol-parser flow where
// SymConvertWow64CompatibilityPaths rewrites system32→syswow64 paths and
// MODULE_SYMBOL_DETAIL.Is32Bit picks the right PDB architecture; here we
// expose the same selection via IsWow64Process + the SymLoadModuleExW
// SLMFLAG_ALT_INDEX flag (DbgHelp's documented switch for the 32-bit
// alternate symbol index).
//
// All Win32 calls go through windows.NewLazyDLL + LazyProc.Call — no cgo.
package symbolparser

import (
	"errors"
	"fmt"
	"unsafe"

	"golang.org/x/sys/windows"
)

// SYMLOADMODULEFLAGS values accepted by SymLoadModuleExW's Flags parameter.
// These match the Microsoft dbghelp.h enum. SLMFLAG_ALT_INDEX is the
// documented switch for loading a 32-bit PDB into a 64-bit DbgHelp context
// (i.e. WOW64); the task-spec names SLMFLAG_32BIT/SLMFLAG_64BIT, so we expose
// those as aliases for callers that prefer the architecture-named form.
const (
	slmFlagNone      = 0x00000000 // No flag (default 64-bit behaviour).
	slmFlagVirtual   = 0x00000002 // Module is virtual (loaded in memory only).
	slmFlagAltIndex  = 0x00000004 // Use the alternate (32-bit) symbol index.
	slmFlagNoSymbols = 0x00000008 // Load image but do not load symbols.

	// SLMFLAG_32BIT / SLMFLAG_64BIT — task-spec aliases. SLMFLAG_32BIT
	// selects the 32-bit alternate symbol index (equivalent to
	// SLMFLAG_ALT_INDEX); SLMFLAG_64BIT is the default no-flag behaviour.
	// These are not Win32 constants — they are the package-level
	// abstraction the task asks for ("封装 SLMFLAG_32BIT / SLMFLAG_64BIT").
	SLMFLAG_32BIT = slmFlagAltIndex
	SLMFLAG_64BIT = slmFlagNone
)

// kernel32 procs used here (module-level LazyProc so we don't re-resolve on
// every call). The dbghelp procs are already declared in parser.go.
var (
	kernel32Dll        = windows.NewLazySystemDLL("kernel32.dll")
	procIsWow64Process = kernel32Dll.NewProc("IsWow64Process")
)

// IsWow64Process reports whether hProcess is a 32-bit WOW64 process running
// under a 64-bit host. Wraps kernel32!IsWow64Process. Returns false (with no
// error) when running on a 32-bit host — there are no WOW64 processes there.
//
// Used by LoadPdbForProcess to pick the right SLMFLAG_* when loading a
// module's symbols, and by callers that need to mirror the C++
// MODULE_SYMBOL_DETAIL.Is32Bit field.
func IsWow64Process(hProcess uintptr) (bool, error) {
	if hProcess == 0 {
		return false, errors.New("symbolparser: IsWow64Process: nil process handle")
	}
	var wow64 int32
	r1, _, err := procIsWow64Process.Call(hProcess, uintptr(unsafe.Pointer(&wow64)))
	if r1 == 0 {
		return false, fmt.Errorf("IsWow64Process failed: %w", err)
	}
	return wow64 != 0, nil
}

// LoadPdbForProcess loads moduleName's symbols into the current process's
// DbgHelp context, auto-selecting the 32-bit or 64-bit symbol index based on
// processHandle's WOW64 status. Returns the SymLoadModuleExW-reported module
// base (the symbol-session handle used as the key for SymUnloadModule64).
//
// processHandle is the OS process handle of the debuggee — it is only used
// to call IsWow64Process and pick SLMFLAG_32BIT vs SLMFLAG_64BIT. It is NOT
// passed to SymLoadModuleExW (DbgHelp scopes symbols to the *current*
// process, mirroring C++ SymLoadFileSymbol which always passes
// GetCurrentProcess() to SymLoadModule64). moduleName is the image path
// (e.g. "C:\\Windows\\System32\\ntdll.dll") — DbgHelp follows the PE debug
// directory to locate the PDB.
//
// Prerequisite: the caller must have initialised DbgHelp on the current
// process via Resolver.Init (which calls SymInitialize). Calling
// LoadPdbForProcess without Init will cause SymLoadModuleExW to fail.
//
// The base address is passed as 0 — DbgHelp resolves it from the PE header.
// If you have an explicit load base (e.g. from the loader or a memory
// dump), use LoadPdbForProcessAtBase instead.
func LoadPdbForProcess(processHandle uintptr, moduleName string) (uint64, error) {
	return loadPdbForProcessInternal(processHandle, moduleName, 0)
}

// LoadPdbForProcessAtBase is like LoadPdbForProcess but lets the caller
// supply an explicit module base address (useful when loading a PDB for a
// module that's been relocated, or when working from a memory snapshot
// rather than the live process). base is the virtual address where the
// module is loaded in the debuggee.
func LoadPdbForProcessAtBase(processHandle uintptr, moduleName string, base uint64) (uint64, error) {
	return loadPdbForProcessInternal(processHandle, moduleName, base)
}

// loadPdbForProcessInternal is the shared core of LoadPdbForProcess and
// LoadPdbForProcessAtBase. It picks the SLMFLAG based on WOW64 status and
// calls SymLoadModuleExW.
func loadPdbForProcessInternal(processHandle uintptr, moduleName string, base uint64) (uint64, error) {
	if moduleName == "" {
		return 0, errors.New("symbolparser: LoadPdbForProcess: empty moduleName")
	}

	// Pick the flag: 32-bit debuggee → alternate symbol index, 64-bit →
	// default. A WOW64 failure (e.g. running on 32-bit host where
	// IsWow64Process is still callable) is treated as 64-bit, matching the
	// C++ fallback in SymbolInitLoad where Is32Bit defaults to FALSE.
	flags := uintptr(SLMFLAG_64BIT)
	if processHandle != 0 {
		is32, err := IsWow64Process(processHandle)
		if err == nil && is32 {
			flags = uintptr(SLMFLAG_32BIT)
		}
	}

	img, err := windows.UTF16PtrFromString(moduleName)
	if err != nil {
		return 0, fmt.Errorf("symbolparser: UTF16 for %q: %w", moduleName, err)
	}

	// Use the current process as the DbgHelp handle. This mirrors C++
	// SymLoadFileSymbol which always passes GetCurrentProcess() to
	// SymLoadModule64; the caller is responsible for SymInitialize'ing the
	// current process via Resolver.Init first.
	hProcess := windows.CurrentProcess()
	proc := dbghelpDll.NewProc(procSymLoadModuleExW)
	r1, _, err := proc.Call(
		uintptr(hProcess),
		0,                            // hFile (NULL = load from image path)
		0,                            // ImageName (NULL = use ModuleName)
		uintptr(unsafe.Pointer(img)), // ModuleName (path used for symbol lookup)
		uintptr(base),                // BaseOfDll
		0,                            // DllSize (0 = look up from PE header)
		0,                            // Data
		flags,                        // Flags (SLMFLAG_32BIT or SLMFLAG_64BIT)
	)
	if r1 == 0 {
		return 0, fmt.Errorf("SymLoadModuleExW failed for %q (flags=0x%x): %w", moduleName, flags, err)
	}
	return uint64(r1), nil
}
