package hook

import (
	"testing"

	"github.com/ddkwork/HyperDbg/debugger/driver/hardwareinfo"
	"github.com/ddkwork/golibrary/std/mylog"
)

func Test_hardware(t *testing.T) {
	// t.Skip()
	h := hardwareinfo.New()
	// if !h.SsdInfo.GetMust() { // todo bug cpu pkg init
	//	return
	// }
	if !h.CpuInfo.Get() {
		return
	}
	if !h.MacInfo.Get() {
		return
	}
	mylog.Struct(h)
}
