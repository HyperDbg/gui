package HPRDBGCTRL

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHyperDbgVmxSupportDetection(t *testing.T) {
	assert.True(t, HyperDbgVmxSupportDetection()) // ok
	s := make([]byte, 0, 100)
	HyperDbgReadVendorString(s) // todo bug
	fmt.Println(s)

	// missing connection local or start ?
	HyperDbgLoadVmm()

	//unloading vmm driver before need wait 1 second?
	// HyperDbgStopVmmDriver()
	HyperDbgUnloadVmm()
	HyperDbgUninstallVmmDriver()
}
