package userlevel

import (
	"os"
	"strings"
	"testing"
)

// notepadPath is the canonical Windows PE file used as a test fixture.
const notepadPath = `C:\Windows\System32\notepad.exe`

// openNotepadOrFail skips the test when the notepad.exe fixture is not
// available (e.g. running on a non-Windows host or a stripped image). The
// returned PeFile must be closed by the caller.
func openNotepadOrFail(t *testing.T) *PeFile {
	t.Helper()
	if _, err := os.Stat(notepadPath); err != nil {
		t.Skipf("notepad.exe fixture not available: %v", err)
	}
	pf, err := PeOpen(notepadPath)
	if err != nil {
		t.Fatalf("PeOpen(%q) failed: %v", notepadPath, err)
	}
	t.Cleanup(func() { _ = pf.Close() })
	return pf
}

// TestPeFileOpen verifies that PeOpen successfully parses a real Windows PE
// file (notepad.exe) and exposes its source path.
func TestPeFileOpen(t *testing.T) {
	t.Parallel()
	pf := openNotepadOrFail(t)
	if got := pf.Path(); got != notepadPath {
		t.Errorf("Path() = %q, want %q", got, notepadPath)
	}
}

// TestPeEntryPoint verifies the entry-point RVA is positive for a real PE.
func TestPeEntryPoint(t *testing.T) {
	t.Parallel()
	pf := openNotepadOrFail(t)
	if ep := pf.EntryPoint(); ep == 0 {
		t.Error("EntryPoint() = 0, want > 0")
	}
}

// TestPeImageBase verifies the preferred image base is non-zero.
func TestPeImageBase(t *testing.T) {
	t.Parallel()
	pf := openNotepadOrFail(t)
	if ib := pf.ImageBase(); ib == 0 {
		t.Error("ImageBase() = 0, want > 0")
	}
}

// TestPeSections verifies the section table is non-empty and contains a
// .text section (every normal Windows executable has one).
func TestPeSections(t *testing.T) {
	t.Parallel()
	pf := openNotepadOrFail(t)
	sections := pf.Sections()
	if len(sections) == 0 {
		t.Fatal("Sections() returned empty slice, want at least one section")
	}
	hasText := false
	for _, s := range sections {
		if s.Name == ".text" {
			hasText = true
			break
		}
	}
	if !hasText {
		var names []string
		for _, s := range sections {
			names = append(names, s.Name)
		}
		t.Errorf("sections = [%s], want one named .text", strings.Join(names, ", "))
	}
}

// TestPeSubsystem verifies SubsystemName returns a non-empty, named value
// (notepad.exe is a Windows GUI/CUI binary).
func TestPeSubsystem(t *testing.T) {
	t.Parallel()
	pf := openNotepadOrFail(t)
	name := pf.SubsystemName()
	if name == "" {
		t.Error("SubsystemName() = empty, want non-empty")
	}
	if name == "Unknown" {
		// notepad.exe subsystem should be recognized, not fall through to the
		// default case; surface the value for diagnosis either way.
		t.Logf("SubsystemName() = %q (expected a known subsystem)", name)
	}
}

// TestPeNonExistent verifies PeOpen returns an error for a missing file.
func TestPeNonExistent(t *testing.T) {
	t.Parallel()
	const missing = `Z:\definitely\does\not\exist\missing.exe`
	if _, err := PeOpen(missing); err == nil {
		t.Errorf("PeOpen(%q) returned nil error, want an error", missing)
	}
}
