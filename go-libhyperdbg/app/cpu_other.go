//go:build !amd64

package app

// readCpuidVendorString is a stub for non-amd64 platforms. The CPUID
// instruction is x86-specific; on other architectures there is no vendor
// string to read. Returns the empty string so GetProcessorVendor maps to
// ProcessorVendorOthers.
func readCpuidVendorString() string {
	return ""
}
