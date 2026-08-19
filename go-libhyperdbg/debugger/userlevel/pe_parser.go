// Package userlevel implements the user-mode debugging helpers under
// libhyperdbg/code/debugger/user-level/: pe-parser (PE inspection for the
// `pe` command), ud (user-debugger state machine), and user-listening
// (event listener for the debuggee's UI events).
//
// The PE parser wraps ok/pe-main (github.com/saferwall/pe) so callers do not
// import that package directly; the rest of go-libhyperdbg stays agnostic of
// the PE parsing library.
package userlevel

import (
	"fmt"
	"strings"

	"github.com/saferwall/pe"
)

// PeFile wraps a parsed pe.File. It is the Go counterpart of the C++
// PeParse+PeShowInfo helpers used by the `pe` command.
type PeFile struct {
	path string
	f    *pe.File
}

// PeOpen parses the PE at path. The caller must call Close when done.
func PeOpen(path string) (*PeFile, error) {
	f, err := pe.New(path, &pe.Options{Fast: false})
	if err != nil {
		return nil, fmt.Errorf("PeOpen(%q): %w", path, err)
	}
	if err := f.Parse(); err != nil {
		return nil, fmt.Errorf("PeOpen(%q): parse failed: %w", path, err)
	}
	return &PeFile{path: path, f: f}, nil
}

// Close releases the underlying PE file resources.
func (p *PeFile) Close() error {
	if p.f == nil {
		return nil
	}
	return p.f.Close()
}

// EntryPoint returns the RVA of the entry point, or 0 if not available.
func (p *PeFile) EntryPoint() uint32 {
	if p.f == nil {
		return 0
	}
	if p.f.Is64 {
		oh := p.f.NtHeader.OptionalHeader.(pe.ImageOptionalHeader64)
		return oh.AddressOfEntryPoint
	}
	oh := p.f.NtHeader.OptionalHeader.(pe.ImageOptionalHeader32)
	return oh.AddressOfEntryPoint
}

// ImageBase returns the preferred load address.
func (p *PeFile) ImageBase() uint64 {
	if p.f == nil {
		return 0
	}
	if p.f.Is64 {
		oh := p.f.NtHeader.OptionalHeader.(pe.ImageOptionalHeader64)
		return oh.ImageBase
	}
	oh := p.f.NtHeader.OptionalHeader.(pe.ImageOptionalHeader32)
	return uint64(oh.ImageBase)
}

// Is64Bit reports whether the PE targets 64-bit.
func (p *PeFile) Is64Bit() bool {
	if p.f == nil {
		return false
	}
	return p.f.Is64
}

// Is32Bit reports whether the PE targets 32-bit (WoW64).
func (p *PeFile) Is32Bit() bool {
	return !p.Is64Bit()
}

// SubsystemName returns a human-readable name for the PE subsystem.
// Mirrors PeGetSubsystemName in pe-parser.cpp.
func (p *PeFile) SubsystemName() string {
	if p.f == nil {
		return "Unknown"
	}
	var subsystem uint16
	if p.f.Is64 {
		oh := p.f.NtHeader.OptionalHeader.(pe.ImageOptionalHeader64)
		subsystem = uint16(oh.Subsystem)
	} else {
		oh := p.f.NtHeader.OptionalHeader.(pe.ImageOptionalHeader32)
		subsystem = uint16(oh.Subsystem)
	}
	switch subsystem {
	case 1: // IMAGE_SUBSYSTEM_NATIVE
		return "Device Driver (Native Windows Process)"
	case 2: // IMAGE_SUBSYSTEM_WINDOWS_GUI
		return "Windows GUI"
	case 3: // IMAGE_SUBSYSTEM_WINDOWS_CUI
		return "Windows CLI"
	case 9: // IMAGE_SUBSYSTEM_WINDOWS_CE_GUI
		return "Windows CE GUI"
	default:
		return "Unknown"
	}
}

// DllCharacteristicsNames returns a comma-separated list of DLL characteristic
// flag names set on the PE. Mirrors PeShowDllCharacteristics in pe-parser.cpp.
func (p *PeFile) DllCharacteristicsNames() string {
	if p.f == nil {
		return ""
	}
	var flags uint16
	if p.f.Is64 {
		oh := p.f.NtHeader.OptionalHeader.(pe.ImageOptionalHeader64)
		flags = uint16(oh.DllCharacteristics)
	} else {
		oh := p.f.NtHeader.OptionalHeader.(pe.ImageOptionalHeader32)
		flags = uint16(oh.DllCharacteristics)
	}
	type entry struct {
		bit  uint16
		name string
	}
	entries := []entry{
		{0x0020, "High Entropy VA"},
		{0x0040, "Dynamic Base"},
		{0x0080, "Force Integrity"},
		{0x0100, "NX Compatible"},
		{0x0200, "No Isolation"},
		{0x0400, "No SEH"},
		{0x0800, "No Bind"},
		{0x1000, "AppContainer"},
		{0x2000, "WDM Driver"},
		{0x4000, "Guard CF"},
		{0x8000, "Terminal Server Aware"},
	}
	var parts []string
	for _, e := range entries {
		if flags&e.bit == e.bit {
			parts = append(parts, e.name)
		}
	}
	if len(parts) == 0 {
		return "None"
	}
	return strings.Join(parts, ", ")
}

// Sections returns information about each section in the PE.
type SectionInfo struct {
	Name     string
	VirtAddr uint32
	VirtSize uint32
	RawSize  uint32
	Chars    uint32 // characteristics
}

// Sections lists the sections of the PE in the order they appear in the table.
func (p *PeFile) Sections() []SectionInfo {
	if p.f == nil {
		return nil
	}
	out := make([]SectionInfo, 0, len(p.f.Sections))
	for _, s := range p.f.Sections {
		out = append(out, SectionInfo{
			Name:     strings.TrimRight(string(s.Header.Name[:]), "\x00"),
			VirtAddr: s.Header.VirtualAddress,
			VirtSize: s.Header.VirtualSize,
			RawSize:  s.Header.SizeOfRawData,
			Chars:    s.Header.Characteristics,
		})
	}
	return out
}

// Imports returns the list of (dll, function) pairs imported by the PE.
type ImportInfo struct {
	Dll      string
	Function string
}

// Imports lists all imports. Returns nil for PE files without imports.
func (p *PeFile) Imports() []ImportInfo {
	if p.f == nil {
		return nil
	}
	out := make([]ImportInfo, 0)
	for _, imp := range p.f.Imports {
		for _, fn := range imp.Functions {
			out = append(out, ImportInfo{
				Dll:      imp.Name,
				Function: fn.Name,
			})
		}
	}
	return out
}

// Exports returns the list of exported function names.
func (p *PeFile) Exports() []string {
	if p.f == nil {
		return nil
	}
	out := make([]string, 0, len(p.f.Export.Functions))
	for _, fn := range p.f.Export.Functions {
		out = append(out, fn.Name)
	}
	return out
}

// Path returns the file path the PE was loaded from.
func (p *PeFile) Path() string { return p.path }
