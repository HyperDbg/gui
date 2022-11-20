package Headers

import (
	"github.com/ddkwork/golibrary/mylog"
	"testing"
)

func TestSizeof(t *testing.T) {

}

func TestHIBYTE(t *testing.T) {
	v := 0x11223344
	mylog.Hex("HIBYTE", HIBYTE(uint32(v)))
	mylog.Hex("HIWORD", HIWORD(uint32(v)))
	mylog.Hex("LOBYTE", LOBYTE(uint32(v)))
	mylog.Hex("LOWORD", LOWORD(uint32(v)))
}
