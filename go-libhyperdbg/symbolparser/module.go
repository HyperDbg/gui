// Package symbolparser — module.go
//
// ModuleCache tracks loaded PE modules and their PDB associations, mirroring
// the C++ symbol-parser's g_LoadedModules vector plus the
// SymGetModuleBaseFromSearchMask / SymUnloadModuleSymbol helpers.
//
// The cache is goroutine-safe (sync.RWMutex) and holds no global state — every
// instance owns its own module table, so a GUI/MCP layer can run multiple
// debuggers in parallel without cross-talk. Symbol auto-loading, when a
// Resolver is bound, mirrors the C++ SymLoadFileSymbol → SymLoadModule64 path.
package symbolparser

import (
	"fmt"
	"path/filepath"
	"strings"
	"sync"
)

// PE machine type constants (subset of IMAGE_FILE_MACHINE_*). Kept local so
// the package does not pull in ok/pe-main just for the constants.
const (
	MachineTypeUnknown uint16 = 0x0000
	MachineTypeI386    uint16 = 0x014c
	MachineTypeAMD64   uint16 = 0x8664
	MachineTypeARM     uint16 = 0x01c0
	MachineTypeARM64   uint16 = 0xaa64
	MachineTypeIA64    uint16 = 0x0200
)

// Module is the Go counterpart of the C++ SYMBOL_LOADED_MODULE_DETAILS
// struct. Field layout is not C-ABI-critical (the C struct uses fixed-size
// char arrays for path/name strings); we use Go strings instead. The
// BaseAddress/ModuleBase split mirrors the C struct: BaseAddress is the
// virtual address where the module is loaded in the debuggee; ModuleBase is
// the symbol-session base returned by DbgHelp's SymLoadModuleEx (used as the
// key for SymUnloadModule64).
type Module struct {
	// FilePath is the PE image path on disk (or the PDB path if the caller
	// registered a standalone PDB). Mirrors PdbFilePath in the C struct when
	// the input is a .pdb.
	FilePath string

	// PdbFilePath is the resolved PDB file path used for symbol loading. For
	// a PE input this is typically the same as FilePath (DbgHelp resolves the
	// PDB via the PE debug directory); for a standalone PDB input it is the
	// PDB path.
	PdbFilePath string

	// BaseAddress is the virtual base address where the module is loaded in
	// the debuggee. Mirrors SYMBOL_LOADED_MODULE_DETAILS.BaseAddress.
	BaseAddress uint64

	// ModuleBase is the symbol-session base returned by SymLoadModuleEx (or
	// equivalent). Mirrors SYMBOL_LOADED_MODULE_DETAILS.ModuleBase. Used as
	// the handle for SymUnloadModule64. Zero when no Resolver is bound.
	ModuleBase uint64

	// ModuleName is the lowercase base name of the module without extension
	// (e.g. "ntdll", "kernelbase"). Mirrors SYMBOL_LOADED_MODULE_DETAILS.ModuleName.
	ModuleName string

	// AltName is the alternative lookup name. For ntoskrnl.exe/ntkrnlmp.exe
	// it is "nt"; for WOW64 PDBs with a 'w' prefix (e.g. wntdll.pdb) it is
	// the unprefixed module name. Mirrors SYMBOL_LOADED_MODULE_DETAILS.ModuleAlternativeName.
	AltName string

	// Size is the module's in-memory size in bytes (SizeOfImage). Zero means
	// unknown — GetByAddr then matches by base address only. The C struct
	// does not track size; we add it so address→module lookup is exact.
	Size uint64

	// MachineType is the PE IMAGE_FILE_HEADER.Machine value (see
	// MachineType* constants). Used by the WOW64 logic to pick 32/64-bit
	// PDBs. Zero means unknown.
	MachineType uint16

	// Is32Bit reports whether the module is a 32-bit image loaded under
	// WOW64. Set by the caller or by LoadPdbForProcess.
	Is32Bit bool
}

// ModuleCache is a goroutine-safe registry of loaded modules. It replaces the
// C++ global g_LoadedModules vector and its g_IsLoadedModulesInitialized
// guard. Lookups by name and by address are supported; the optional Resolver
// is invoked on Add to load symbols via DbgHelp (matching the C++ flow where
// SymLoadFileSymbol both records the module and calls SymLoadModule64).
type ModuleCache struct {
	mu sync.RWMutex

	// modules is keyed by lowercase ModuleName.
	modules map[string]*Module

	// altNames is keyed by lowercase AltName. The value is the same *Module
	// as in modules, so List() does not double-count.
	altNames map[string]*Module

	// resolver, if non-nil, is called from Add to load symbols via DbgHelp.
	// It mirrors the C++ SymLoadModule64 call inside SymLoadFileSymbol.
	resolver Resolver
}

// NewModuleCache returns an empty cache. Pass a non-nil resolver to enable
// automatic symbol loading in Add(); pass nil to track modules only (the
// caller is then responsible for calling Resolver.LoadModule separately).
func NewModuleCache(resolver Resolver) *ModuleCache {
	return &ModuleCache{
		modules:  make(map[string]*Module),
		altNames: make(map[string]*Module),
		resolver: resolver,
	}
}

// Add registers a module and, when a Resolver is bound, loads its symbols via
// DbgHelp. path may be a PE image or a standalone .pdb file. base is the
// virtual base address where the module is loaded in the debuggee.
//
// The module name is derived from the file's base name (without extension),
// lowercased — matching the C++ _splitpath + tolower loop in
// SymLoadFileSymbol. Alternative names are detected for ntoskrnl-family
// kernels (alt = "nt") and WOW64 'w'-prefixed PDBs (alt = unprefixed name),
// matching SymCheckNtoskrnlPrefix and SymCheckAndRemoveWow64Prefix.
//
// Returns the registered Module. Adding a module with an existing name
// overwrites the previous entry (matching the C++ behaviour where the same
// module can be reloaded).
func (c *ModuleCache) Add(path string, base uint64) (*Module, error) {
	name := moduleBaseName(path)
	alt := detectAlternativeName(path, name)

	m := &Module{
		FilePath:    path,
		PdbFilePath: path,
		BaseAddress: base,
		ModuleName:  name,
		AltName:     alt,
	}

	if c.resolver != nil {
		// Auto-load symbols via DbgHelp. Callers needing cancellation
		// should call Resolver.LoadModule directly and then
		// ModuleCache.Register to record the module.
		mb, err := c.resolver.LoadModule(path, base)
		if err != nil {
			return nil, fmt.Errorf("symbolparser: load module %q: %w", path, err)
		}
		m.ModuleBase = mb
	}

	c.mu.Lock()
	defer c.mu.Unlock()
	// Remove any stale alt-name entry pointing at the previous module
	// registered under this name, to avoid dangling alt references.
	if old, ok := c.modules[name]; ok && old.AltName != "" {
		delete(c.altNames, old.AltName)
	}
	c.modules[name] = m
	if alt != "" {
		c.altNames[alt] = m
	}
	return m, nil
}

// Register records a Module in the cache without invoking the Resolver. Use
// this when the caller has already loaded symbols (e.g. via a prior
// Resolver.LoadModule call) and just wants the cache to know about the
// module. Mirrors the tail of C++ SymLoadFileSymbol (the g_LoadedModules
// push_back) without the SymLoadModule64 half.
func (c *ModuleCache) Register(m *Module) {
	if m == nil || m.ModuleName == "" {
		return
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	if old, ok := c.modules[m.ModuleName]; ok && old.AltName != "" {
		delete(c.altNames, old.AltName)
	}
	c.modules[m.ModuleName] = m
	if m.AltName != "" {
		c.altNames[m.AltName] = m
	}
}

// Get looks up a module by name (case-insensitive). Both the primary
// ModuleName and the AltName are matched, mirroring the C++
// SymGetModuleBaseFromSearchMask two-pass scan. Returns (nil, false) when
// the cache is empty or the name is unknown.
func (c *ModuleCache) Get(name string) (*Module, bool) {
	if name == "" {
		return nil, false
	}
	key := strings.ToLower(name)
	c.mu.RLock()
	defer c.mu.RUnlock()
	if m, ok := c.modules[key]; ok {
		return m, true
	}
	if m, ok := c.altNames[key]; ok {
		return m, true
	}
	return nil, false
}

// GetByAddr returns the module whose [BaseAddress, BaseAddress+Size) range
// contains addr. When no module has a known Size, falls back to the module
// with the largest BaseAddress <= addr (the C++ code does not implement
// address→module lookup at all; this is a Go-side convenience for the `u`
// disassembler and stack-trace symbolisation).
func (c *ModuleCache) GetByAddr(addr uint64) (*Module, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	var bestBase *Module
	for _, m := range c.modules {
		if m.BaseAddress == 0 {
			continue
		}
		if addr < m.BaseAddress {
			continue
		}
		if m.Size > 0 && addr < m.BaseAddress+m.Size {
			// Exact range match — return immediately.
			return m, true
		}
		if bestBase == nil || m.BaseAddress > bestBase.BaseAddress {
			bestBase = m
		}
	}
	if bestBase != nil {
		return bestBase, true
	}
	return nil, false
}

// Remove drops a module from the cache (both primary and alt-name entries).
// It does NOT call SymUnloadModule64 — the caller is responsible for
// unloading symbols via the Resolver if needed. This mirrors the
// g_LoadedModules.erase half of C++ SymUnloadModuleSymbol; the
// SymUnloadModule64 half lives in the Resolver.
func (c *ModuleCache) Remove(name string) {
	key := strings.ToLower(name)
	c.mu.Lock()
	defer c.mu.Unlock()
	m, ok := c.modules[key]
	if !ok {
		// Maybe it was registered under the alt name.
		if alt, ok2 := c.altNames[key]; ok2 {
			m = alt
		} else {
			return
		}
	}
	delete(c.modules, m.ModuleName)
	if m.AltName != "" {
		delete(c.altNames, m.AltName)
	}
}

// List returns a snapshot of all registered modules. Order is unspecified
// (map iteration order); callers that need a stable order should sort the
// slice. Mirrors iterating g_LoadedModules in C++.
func (c *ModuleCache) List() []*Module {
	c.mu.RLock()
	defer c.mu.RUnlock()
	out := make([]*Module, 0, len(c.modules))
	for _, m := range c.modules {
		out = append(out, m)
	}
	return out
}

// SetResolver swaps the Resolver used by Add for symbol auto-loading. Pass
// nil to disable auto-loading. Existing cached modules are untouched. This
// is the Go counterpart of C++ SymInit() (which sets up DbgHelp) called
// lazily before SymLoadFileSymbol.
func (c *ModuleCache) SetResolver(r Resolver) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.resolver = r
}

// moduleBaseName extracts the file's base name without extension and
// lowercases it, mirroring the C++ _splitpath + tolower loop in
// SymLoadFileSymbol. Example: "C:\\Windows\\System32\\ntdll.dll" → "ntdll".
func moduleBaseName(path string) string {
	if path == "" {
		return ""
	}
	base := filepath.Base(path)
	// Strip the extension (filepath.Ext returns "" if none).
	if ext := filepath.Ext(base); ext != "" {
		base = strings.TrimSuffix(base, ext)
	}
	return strings.ToLower(base)
}

// detectAlternativeName mirrors the C++ SymCheckNtoskrnlPrefix and
// SymCheckAndRemoveWow64Prefix helpers. For ntoskrnl-family PDBs it returns
// "nt"; for WOW64 'w'-prefixed PDBs loaded from system32 it returns the
// unprefixed module name. Returns "" when no alternative name applies.
func detectAlternativeName(path, moduleName string) string {
	// ntoskrnl-family check (SymCheckNtoskrnlPrefix). The kernel PDBs are
	// ntkrnlmp.pdb / ntoskrnl.pdb / ntkrpamp.pdb / ntkrnlpa.pdb — for those
	// the alternative name is "nt".
	switch moduleName {
	case "ntkrnlmp", "ntoskrnl", "ntkrpamp", "ntkrnlpa":
		return "nt"
	}

	// WOW64 'w'-prefix check (SymCheckAndRemoveWow64Prefix). If the PDB name
	// starts with 'w' and the module path lives under ":\windows\system32",
	// the alternative name is the module name without the 'w' prefix. This
	// covers 32-bit system DLLs whose PDBs were named with a leading 'w' to
	// avoid colliding with the 64-bit PDB in the same symbol store.
	lowerPath := strings.ToLower(path)
	if strings.HasPrefix(moduleName, "w") && strings.Contains(lowerPath, `:\windows\system32`) {
		return moduleName[1:]
	}
	return ""
}
