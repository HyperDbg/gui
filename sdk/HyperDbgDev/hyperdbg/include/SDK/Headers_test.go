package SDK

import (
	_ "embed"
	"github.com/ddkwork/librarygo/src/myc2go/windef"
	"github.com/ddkwork/librarygo/src/mycheck"
	"github.com/ddkwork/librarygo/src/stream"
	"os"
	"strings"
	"testing"
)

func TestName(t *testing.T) {
	paths := []string{
		"./Headers/BasicTypes.h",
		"./Headers./Constants.h",
		"./Headers/ErrorCodes.h",
		"./Headers/Connection.h",
		"./Headers/Datatypes.h",
		"./Headers/Ioctls.h",
		"./Headers/Events.h",
		"./Headers/RequestStructures.h",
		"./Headers/Symbols.h",
	}
	buffer := stream.New()
	for _, path := range paths {
		b, err := os.ReadFile(path)
		if !mycheck.Error(err) {
			return
		}
		buffer.WriteBytesLn(b)
	}
	all := strings.ReplaceAll(buffer.String(), "#pragma once", "")
	mycheck.Assert(t).True(windef.Creat(all, true))
}
