package misc

import (
	"strings"
	"testing"
)

// pci.ids sample (subset of the real format):
//
//	# comment line
//	0010  Test Vendor One
//	\t0001  Device One
//	\t\t0010 0001  Subsystem One
//	0020  Test Vendor Two
//	\t0002  Device Two
const samplePciIds = `# HyperDbg test pci.ids sample
0010  Test Vendor One
	0001  Device One
		0010 0001  Subsystem One
0020  Test Vendor Two
	0002  Device Two
`

func TestPciId_LoadAndLookup(t *testing.T) {
	db := NewPciIdDatabase()
	if err := db.Load(strings.NewReader(samplePciIds)); err != nil {
		t.Fatalf("Load failed: %v", err)
	}

	name, ok := db.LookUpVendor("0010")
	if !ok || name != "Test Vendor One" {
		t.Errorf("LookUpVendor(0010) = %q ok=%v, want %q true", name, ok, "Test Vendor One")
	}
	// Case-insensitive.
	if name, ok := db.LookUpVendor("0010"); !ok || name != "Test Vendor One" {
		t.Errorf("case-insensitive lookup failed: %q ok=%v", name, ok)
	}

	name, ok = db.LookUpDevice("0010", "0001")
	if !ok || name != "Device One" {
		t.Errorf("LookUpDevice(0010,0001) = %q ok=%v", name, ok)
	}

	name, ok = db.LookUpSubsystem("0010", "0001", "0010", "0001")
	if !ok || name != "Subsystem One" {
		t.Errorf("LookUpSubsystem = %q ok=%v", name, ok)
	}

	_, ok = db.LookUpVendor("FFFF")
	if ok {
		t.Error("LookUpVendor on missing id should return false")
	}
}

func TestPciId_StringRoundTrip(t *testing.T) {
	db := NewPciIdDatabase()
	if err := db.Load(strings.NewReader(samplePciIds)); err != nil {
		t.Fatalf("Load failed: %v", err)
	}
	out := db.String()
	if !strings.Contains(out, "Test Vendor One") || !strings.Contains(out, "Device One") {
		t.Errorf("String() missing expected content: %q", out)
	}
}
