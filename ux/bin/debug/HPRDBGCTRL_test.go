package HPRDBGCTRL

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/ddkwork/golibrary/stream"

	"github.com/stretchr/testify/assert"
)

func TestSdk(t *testing.T) {
	stream.CopyFile("hprdbgkd.sys", filepath.Join("C:\\Windows\\System32\\drivers\\hprdbgkd.sys"))

	assert.True(t, HyperDbgVmxSupportDetection()) // ok
	s := make([]byte, 0, 100)
	HyperDbgReadVendorString(s) // todo bug
	fmt.Println(s)

	HyperDbgInstallVmmDriver()
	// HyperDbgStartVmmDriver()

	// missing connection local or start ?
	HyperDbgLoadVmm()

	// unloading vmm driver before need wait 1 second?
	// HyperDbgStopVmmDriver()
	HyperDbgUnloadVmm()
	HyperDbgUninstallVmmDriver()
}
