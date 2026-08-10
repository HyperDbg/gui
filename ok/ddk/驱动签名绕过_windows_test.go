package ddk

import (
	"testing"

	"github.com/ddkwork/golibrary/std/mylog"
)

func TestFindCiOptionsRVAFromDisk(t *testing.T) {
	f := NewKernelModuleFinder()
	rt := NewRTCore64()
	km := NewKernelMemory(rt)
	dse := NewDSEBypass(f, km)

	rva := mylog.Check2(dse.FindCiOptionsRVAFromDisk())

	t.Logf("g_CiOptions RVA: 0x%X", rva)
}

func TestParseCiConfigFlags(t *testing.T) {
	f := NewKernelModuleFinder()
	rt := NewRTCore64()
	km := NewKernelMemory(rt)
	dse := NewDSEBypass(f, km)

	flags := dse.ParseCiConfigFlags(CI_OPT_IN_UM | CI_OPT_IN_KM)
	t.Logf("CI_OPT_IN_UM|CI_OPT_IN_KM flags: %v", flags)

	flags2 := dse.ParseCiConfigFlags(CI_OPTIONS_DISABLED)
	t.Logf("CI_OPTIONS_DISABLED flags: %v", flags2)

	flags3 := dse.ParseCiConfigFlags(CI_OPT_IN_KM | CI_OPT_HVCI_ACTIVE)
	t.Logf("DSE+HVCI flags: %v", flags3)
}

func TestRTCore64_DSEBypass_ReadCiOptions(t *testing.T) {
	rt := loadRTCore64(t)
	defer rt.Unload()
	km := NewKernelMemory(rt)
	f := NewKernelModuleFinder()
	dse := NewDSEBypass(f, km)

	addr := dse.FindCiOptions()

	t.Logf("g_CiOptions address: %s", mylog.Hex(addr))

	val := dse.ReadCiOptions()

	t.Logf("g_CiOptions value: 0x%08X", val)
	t.Logf("DSE enabled: %v", dse.DSEEnabled())
	t.Logf("HVCI active: %v", dse.HVCIActive())

	flags := dse.ParseCiConfigFlags(val)
	t.Logf("CI config flags: %v", flags)
}
