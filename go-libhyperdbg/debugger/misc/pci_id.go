// Package misc — pci_id.go
//
// Provides runtime access to the PCI ID database (vendor/device names,
// subsystem mappings). Mirrors libhyperdbg/code/debugger/misc/pci-id.cpp.
//
// The database is the canonical pci.ids file shipped with Linux distributions
// (format: vendor lines, device lines indented with a TAB, subsystem lines
// indented with two TABs). The Go version loads it once from a file path or
// embedded blob and exposes LookUp helpers used by the `pcitree` command.
package misc

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
)

// PciIdDatabase is the in-memory representation of the pci.ids file.
type PciIdDatabase struct {
	mu      sync.RWMutex
	vendors map[string]*PciVendor // keyed by 4-hex-digit vendor id (uppercase)
}

// PciVendor is one vendor entry.
type PciVendor struct {
	Id   string // 4 hex digits, uppercase
	Name string
	// Devices maps device-id (4 hex digits, uppercase) to device info.
	Devices map[string]*PciDevice
}

// PciDevice is one device entry under a vendor.
type PciDevice struct {
	Id   string
	Name string
	// Subsystems keyed by "subvendor:subdevice" (each 4 hex digits, uppercase).
	Subsystems map[string]*PciSubsystem
}

// PciSubsystem is one (subvendor, subdevice) pair under a device.
type PciSubsystem struct {
	Subvendor string
	Subdevice string
	Name      string
}

// NewPciIdDatabase returns an empty database.
func NewPciIdDatabase() *PciIdDatabase {
	return &PciIdDatabase{vendors: make(map[string]*PciVendor)}
}

// Load reads the pci.ids file from r and merges it into the database. Multiple
// Load calls accumulate (later entries override earlier ones with the same id).
func (db *PciIdDatabase) Load(r io.Reader) error {
	sc := bufio.NewScanner(r)
	sc.Buffer(make([]byte, 0, 64*1024), 1024*1024)
	var currentVendor *PciVendor
	var currentDevice *PciDevice
	db.mu.Lock()
	defer db.mu.Unlock()
	for sc.Scan() {
		line := sc.Text()
		// Strip trailing CR (in case of CRLF input on Windows).
		line = strings.TrimRight(line, "\r")
		// Skip comments and blanks.
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		switch {
		case strings.HasPrefix(line, "\t\t"):
			// Subsystem: \t\t subvendor subdevice  name
			fields := strings.SplitN(line[2:], "  ", 2)
			if len(fields) != 2 || currentDevice == nil {
				continue
			}
			ids := strings.Fields(fields[0])
			if len(ids) != 2 {
				continue
			}
			sub := &PciSubsystem{
				Subvendor: strings.ToUpper(ids[0]),
				Subdevice: strings.ToUpper(ids[1]),
				Name:      strings.TrimSpace(fields[1]),
			}
			if currentDevice.Subsystems == nil {
				currentDevice.Subsystems = make(map[string]*PciSubsystem)
			}
			currentDevice.Subsystems[sub.Subvendor+":"+sub.Subdevice] = sub
		case strings.HasPrefix(line, "\t"):
			// Device: \t devid  name
			fields := strings.SplitN(line[1:], "  ", 2)
			if len(fields) != 2 || currentVendor == nil {
				continue
			}
			id := strings.ToUpper(strings.TrimSpace(fields[0]))
			dev := &PciDevice{
				Id:         id,
				Name:       strings.TrimSpace(fields[1]),
				Subsystems: make(map[string]*PciSubsystem),
			}
			currentVendor.Devices[id] = dev
			currentDevice = dev
		default:
			// Vendor: vendorid  name
			fields := strings.SplitN(line, "  ", 2)
			if len(fields) != 2 {
				continue
			}
			id := strings.ToUpper(strings.TrimSpace(fields[0]))
			v := &PciVendor{
				Id:      id,
				Name:    strings.TrimSpace(fields[1]),
				Devices: make(map[string]*PciDevice),
			}
			db.vendors[id] = v
			currentVendor = v
			currentDevice = nil
		}
	}
	return sc.Err()
}

// LoadFile is a convenience that opens path and calls Load.
func (db *PciIdDatabase) LoadFile(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	return db.Load(f)
}

// LookUpVendor returns the vendor name for a 4-hex-digit vendor id (case-
// insensitive). Returns ("", false) when not found.
func (db *PciIdDatabase) LookUpVendor(vendorId string) (string, bool) {
	db.mu.RLock()
	defer db.mu.RUnlock()
	v, ok := db.vendors[strings.ToUpper(vendorId)]
	if !ok {
		return "", false
	}
	return v.Name, true
}

// LookUpDevice returns the device name for a (vendor, device) pair. Returns
// ("", false) when not found.
func (db *PciIdDatabase) LookUpDevice(vendorId, deviceId string) (string, bool) {
	db.mu.RLock()
	defer db.mu.RUnlock()
	v, ok := db.vendors[strings.ToUpper(vendorId)]
	if !ok {
		return "", false
	}
	d, ok := v.Devices[strings.ToUpper(deviceId)]
	if !ok {
		return "", false
	}
	return d.Name, true
}

// LookUpSubsystem returns the subsystem name for a (vendor, device, subvendor,
// subdevice) tuple.
func (db *PciIdDatabase) LookUpSubsystem(vendorId, deviceId, subvendorId, subdeviceId string) (string, bool) {
	db.mu.RLock()
	defer db.mu.RUnlock()
	v, ok := db.vendors[strings.ToUpper(vendorId)]
	if !ok {
		return "", false
	}
	d, ok := v.Devices[strings.ToUpper(deviceId)]
	if !ok {
		return "", false
	}
	s, ok := d.Subsystems[strings.ToUpper(subvendorId)+":"+strings.ToUpper(subdeviceId)]
	if !ok {
		return "", false
	}
	return s.Name, true
}

// String renders the database as the original pci.ids text. Useful for
// round-trip verification against the C-side emitter.
func (db *PciIdDatabase) String() string {
	db.mu.RLock()
	defer db.mu.RUnlock()
	var b strings.Builder
	for vid, v := range db.vendors {
		fmt.Fprintf(&b, "%s  %s\n", vid, v.Name)
		for did, d := range v.Devices {
			fmt.Fprintf(&b, "\t%s  %s\n", did, d.Name)
			for _, s := range d.Subsystems {
				fmt.Fprintf(&b, "\t\t%s %s  %s\n", s.Subvendor, s.Subdevice, s.Name)
			}
		}
	}
	return b.String()
}
