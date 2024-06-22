package HPRDBGCTRL

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

//how to test?
//1 go to D:\workspace\workspace\branch\gui\ux\bin\debug
//run cmd "go test ." or make main pkg and cody this test code into main func
//need export all function to test,now only 1-7 function can be export
// more error found in test ...

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

/*
GOROOT=C:\Program Files\Go #gosetup
GOPATH=C:\Users\Admin\go #gosetup
"C:\Program Files\Go\bin\go.exe" test -c -o C:\Users\Admin\AppData\Local\JetBrains\GoLand2024.1\tmp\GoLand\___TestHyperDbgVmxSupportDetection_in_github_com_ddkwork_hyperdbgui_ux_bin_debug.test.exe github.com/ddkwork/hyperdbgui/ux/bin/debug #gosetup
"C:\Program Files\Go\bin\go.exe" tool test2json -t C:\Users\Admin\AppData\Local\JetBrains\GoLand2024.1\tmp\GoLand\___TestHyperDbgVmxSupportDetection_in_github_com_ddkwork_hyperdbgui_ux_bin_debug.test.exe -test.v -test.paniconexit0 -test.run ^\QTestHyperDbgVmxSupportDetection\E$ #gosetup
2024-06-23 06:03:10    Trace ->  --------- title --------- │ ------------------ info ------------------ //runtime.doInit1+0xec C:/Program Files/Go/src/runtime/proc.go:7176
=== RUN   TestHyperDbgVmxSupportDetection
[]
--- PASS: TestHyperDbgVmxSupportDetection (0.00s)
PASS
current processor vendor is : GenuineIntel
virtualization technology is vt-x
vmx operation is supported by your processor
err, CreateFile failed (2)
handle of the driver not found, probably the driver is not loaded. Did you use 'load' command?

Process finished with the exit code 0
*/
