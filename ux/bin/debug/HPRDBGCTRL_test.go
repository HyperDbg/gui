package HPRDBGCTRL

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHyperDbgVmxSupportDetection(t *testing.T) {
	assert.True(t, HyperDbgVmxSupportDetection())
	s := make([]byte, 0, 100)
	HyperDbgReadVendorString(s)
	fmt.Println(s)

	HyperDbgUninstallVmmDriver()
	HyperDbgLoadVmm()

	// HyperDbgStopVmmDriver()

	HyperDbgUnloadVmm()
	HyperDbgUninstallVmmDriver()
}
