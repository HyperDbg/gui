// Package symbolparser — dia.go
//
// DiaResolver is a DIA (msdia140.dll) backed PDB resolver that reuses
// ok/pdbex to load and query PDB files via the COM IDiaDataSource +
// IDiaSession interfaces. It mirrors the per-PDB functionality that the C++
// symbol-parser obtains through DbgHelp's SymLoadModule64 + SymFromName /
// SymFromAddr — but goes straight to DIA, bypassing DbgHelp, so it works
// for PDBs that are not registered with any process symbol context (e.g. a
// PDB downloaded on the side for offline symbolisation).
//
// All DLL/COM calls go through windows.NewLazyDLL + LazyProc.Call /
// syscall.NewLazyDLL — no cgo. The DIA COM plumbing lives inside ok/pdbex
// (see dia_windows.go: CoCreateInstance → IDiaDataSource::loadDataFromPdb
// → openSession → enumerateSymbols / enumerateFunctions), which we import
// as github.com/ddkwork/pdbex.
package symbolparser

import (
	"errors"
	"fmt"
	"strings"
	"sync"

	"github.com/ddkwork/pdbex"
)

// DiaResolver loads a single PDB via DIA and resolves symbol names ↔ RVAs.
// It is goroutine-safe: pdbex.PDB guards its own state with a RWMutex, and
// we add a separate state mutex for LoadPdb/Close transitions so concurrent
// Get* calls don't race with a reload.
//
// Addresses returned by GetSymbolByName are RVAs (relative virtual addresses
// from the PDB), not absolute virtual addresses — the caller adds the
// module's BaseAddress to get the runtime VA. This matches what DIA stores
// (IDiaSymbol::get_relativeVirtualAddress) and avoids needing to know the
// load base at PDB-load time.
type DiaResolver struct {
	// state guards the loaded flag and pdb pointer so concurrent
	// LoadPdb/Close/Get* calls don't race. pdbex.PDB has its own internal
	// lock for symbol-map reads, so we don't need to hold state during
	// GetSymbolByName/GetSymbolByAddr themselves.
	state sync.RWMutex

	// pdb is the underlying ok/pdbex PDB instance. Nil when no PDB is
	// loaded (or after Close).
	pdb *pdbex.PDB

	// machineType is the PE machine type of the loaded PDB, as reported by
	// DIA. Zero when no PDB is loaded. Used by the WOW64 logic to pick
	// 32-bit vs 64-bit PDBs.
	machineType uint16

	// basePath, when non-empty, is the on-disk path of the loaded PDB. Used
	// for diagnostics and to mirror pdbex.PDB.GetPath without forcing callers
	// to import pdbex directly.
	basePath string
}

// NewDiaResolver creates an empty DIA-backed resolver. Call LoadPdb to bind
// a PDB file before invoking GetSymbolByName / GetSymbolByAddr. The
// underlying msdia140.dll / symsrv.dll are extracted from ok/pdbex's
// //go:embed on first use (see ok/pdbex/embed.go and dia_windows.go's
// loadMsdia140).
func NewDiaResolver() *DiaResolver {
	return &DiaResolver{}
}

// LoadPdb opens pdbPath via DIA and enumerates its symbols/functions. The
// machineType parameter is recorded for the WOW64 selection logic but does
// not override the type DIA reports from the PDB itself — the two must
// agree; a mismatch returns an error so callers don't accidentally load a
// 32-bit PDB under a 64-bit machine-type context.
//
// Calling LoadPdb on an already-loaded resolver replaces the previous PDB
// (the previous PDB is closed first). This mirrors the C++ flow where
// SymLoadModule64 on the same path reloads symbols.
func (d *DiaResolver) LoadPdb(pdbPath string, machineType uint16) error {
	d.state.Lock()
	defer d.state.Unlock()

	// Close any previously loaded PDB before opening a new one.
	if d.pdb != nil {
		d.pdb.Close()
		d.pdb = nil
		d.machineType = 0
		d.basePath = ""
	}

	p := pdbex.NewPDB()
	if err := p.Open(pdbPath); err != nil {
		return fmt.Errorf("symbolparser: DIA open %q: %w", pdbPath, err)
	}

	// Verify the machine type reported by DIA matches the caller's
	// expectation. machineType == 0 means "unknown / don't care" (the
	// caller hasn't been told yet), in which case we accept DIA's value.
	reported := p.GetMachineType()
	if machineType != 0 && reported != 0 && machineType != reported {
		p.Close()
		return fmt.Errorf("symbolparser: machine type mismatch for %q: expected 0x%04x, PDB reports 0x%04x",
			pdbPath, machineType, reported)
	}

	d.pdb = p
	d.machineType = reported
	d.basePath = pdbPath
	return nil
}

// GetSymbolByName resolves a symbol (function or data) name to its RVA. The
// returned address is relative to the module base — add Module.BaseAddress
// to get the absolute virtual address. Returns an error if no PDB is loaded
// or the name is unknown.
//
// Mirrors the address half of C++ SymConvertNameToAddress (which uses
// SymFromName). The "module!symbol" syntax is handled by ModuleCache.Get +
// DiaResolver, not here — pass the bare symbol name.
func (d *DiaResolver) GetSymbolByName(name string) (uint64, error) {
	if name == "" {
		return 0, errors.New("symbolparser: empty symbol name")
	}
	d.state.RLock()
	pdb := d.pdb
	d.state.RUnlock()
	if pdb == nil {
		return 0, errors.New("symbolparser: no PDB loaded")
	}

	// Try function lookup first (functions are the common case for hooks).
	if fn, ok := pdb.GetFunctionInfo(name); ok && fn.Address != 0 {
		return fn.Address, nil
	}
	// Fall back to the generic symbol map (covers data symbols, UDTs,
	// enums — though only data symbols have a meaningful address).
	if sym, ok := pdb.GetSymbolByName(name); ok {
		// pdbex.Symbol does not carry an RVA directly; for data symbols the
		// RVA is exposed via the DIA session's get_relativeVirtualAddress.
		// pdbex exposes function RVAs via FunctionInfo.Address; for non-
		// function symbols we cannot recover the RVA without extending
		// pdbex. Return an error to flag the gap rather than silently
		// returning 0.
		_ = sym
		return 0, fmt.Errorf("symbolparser: %q is not a function symbol (data-symbol RVA not exposed by pdbex)", name)
	}
	return 0, fmt.Errorf("symbolparser: symbol %q not found", name)
}

// GetSymbolByAddr resolves an RVA to (symbol name, offset within the
// symbol). The offset is 0 when addr lands exactly on a symbol start.
// addr is interpreted as an RVA — subtract Module.BaseAddress from an
// absolute VA before passing it here. Returns an error if no PDB is loaded
// or no function covers the address.
//
// Mirrors C++ SymFromAddr (DbgHelp) but operates on the DIA function
// table instead. Note: pdbex's GetFunctionByOffset performs a linear scan
// with the largest-base-less-than heuristic — for hot paths callers should
// cache results.
func (d *DiaResolver) GetSymbolByAddr(addr uint64) (string, uint64, error) {
	d.state.RLock()
	pdb := d.pdb
	d.state.RUnlock()
	if pdb == nil {
		return "", 0, errors.New("symbolparser: no PDB loaded")
	}

	name, ok := pdb.GetFunctionByOffset(addr)
	if !ok {
		return "", 0, fmt.Errorf("symbolparser: no function at RVA 0x%x", addr)
	}

	// Recover the function's start address so we can compute the
	// displacement. pdbex's GetFunctionByOffset only returns the name, so
	// we re-look up the FunctionInfo to get the start RVA.
	offset := uint64(0)
	if fn, ok := pdb.GetFunctionInfo(name); ok && fn.Address <= addr {
		offset = addr - fn.Address
	}
	return name, offset, nil
}

// Close releases the underlying DIA session and PDB symbol maps. Idempotent;
// calling Close on an already-closed resolver is a no-op. After Close, the
// resolver can be reused by calling LoadPdb again.
func (d *DiaResolver) Close() error {
	d.state.Lock()
	defer d.state.Unlock()
	if d.pdb != nil {
		d.pdb.Close()
		d.pdb = nil
		d.machineType = 0
		d.basePath = ""
	}
	return nil
}

// IsLoaded reports whether a PDB is currently loaded. Safe for concurrent
// use.
func (d *DiaResolver) IsLoaded() bool {
	d.state.RLock()
	defer d.state.RUnlock()
	return d.pdb != nil
}

// MachineType returns the PE machine type of the loaded PDB (as reported by
// DIA), or 0 if no PDB is loaded. See MachineType* constants.
func (d *DiaResolver) MachineType() uint16 {
	d.state.RLock()
	defer d.state.RUnlock()
	return d.machineType
}

// Path returns the on-disk path of the loaded PDB, or "" if none is loaded.
func (d *DiaResolver) Path() string {
	d.state.RLock()
	defer d.state.RUnlock()
	return d.basePath
}

// ResolveModuleSymbol combines a ModuleCache lookup with a DiaResolver
// name resolution. Given "ntdll!RtlAllocateHeap" (or just "RtlAllocateHeap"
// when the module is unambiguous), it returns the absolute virtual address.
// This mirrors the C++ SymConvertNameToAddress flow: split on '!', look up
// the module by name, then look up the symbol within the module's PDB.
func (d *DiaResolver) ResolveModuleSymbol(cache *ModuleCache, qualified string) (uint64, error) {
	if qualified == "" {
		return 0, errors.New("symbolparser: empty symbol")
	}

	// Split "module!symbol" — if there's no '!', assume the symbol is in
	// the "nt" module (matching C++ SymConvertNameToAddress default).
	var moduleName, symName string
	if idx := strings.IndexByte(qualified, '!'); idx >= 0 {
		moduleName = strings.ToLower(qualified[:idx])
		symName = qualified[idx+1:]
	} else {
		moduleName = "nt"
		symName = qualified
	}
	if symName == "" {
		return 0, fmt.Errorf("symbolparser: empty symbol name in %q", qualified)
	}

	mod, ok := cache.Get(moduleName)
	if !ok {
		return 0, fmt.Errorf("symbolparser: module %q not in cache", moduleName)
	}

	rva, err := d.GetSymbolByName(symName)
	if err != nil {
		return 0, err
	}
	return mod.BaseAddress + rva, nil
}
