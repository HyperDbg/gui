package ark

import (
	"testing"

	"github.com/ddkwork/golibrary/std/safemap"
	"github.com/ddkwork/golibrary/std/stream"
)

func TestName(t *testing.T) {
	g := stream.NewGeneratedFile()
	m := new(safemap.M[string, string])
	m.Set("kernelTables", "kernelTables")
	m.Set("explorer", "explorer")
	m.Set("taskManager", "taskManager")
	m.Set("driverTool", "driverTool")
	m.Set("registryEditor", "registryEditor")
	m.Set("hardwareMonitor", "hardwareMonitor")
	m.Set("hardwareHook", "hardwareHook")
	m.Set("randomHook", "randomHook")
	m.Set("environmentEditor", "environmentEditor")
	m.Set("vstart", "vstart")
	m.Set("crypt", "crypt")
	m.Set("packer", "packer")
	g.EnumTypes("arks", m)
}
