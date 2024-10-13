package ark

import (
	"testing"

	"github.com/ddkwork/golibrary/stream"
)

func TestName(t *testing.T) {
	g := stream.NewGeneratedFile()
	g.Enum("arks", []string{
		"kernelTables",
		"explorer",
		"taskManager",
		"driverTool",
		"registryEditor",
		"hardwareMonitor",
		"hardwareHook",
		"randomHook",
		"environmentEditor",
		"vstart",
		"crypt",
		"packer",
	}, nil)
}
