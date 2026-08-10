//go:build !amd64

package transparency

// Non-amd64 stubs. HyperDbg's VMM driver is amd64-only, so on any other
// architecture the transparency measurement helpers return zeros (no RDTSC /
// CPUID available). The package still compiles, allowing cross-platform tooling
// (go vet, go doc, etc.) to run. The IOCTL path (HideDebugger / UnhideDebugger)
// does not depend on these and still works on any architecture that can open
// the device — though in practice the device only exists on amd64 Windows.

func rdtsc() uint64 { return 0 }

func cpuidLeaf(leaf uint32) (eax, ebx, ecx, edx uint32) {
	return 0, 0, 0, 0
}

func cpuidLeaf0() [4]uint32 { return [4]uint32{} }
