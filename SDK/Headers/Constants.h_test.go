package Headers

import (
	"github.com/ddkwork/golibrary/mylog"
	"testing"
)

func TestSizeof(t *testing.T) {

}

func TestHIBYTE(t *testing.T) {
	v := uint32(0x11223344)
	assert := mylog.Assert(t)
	assert.Equal(byte(0x11), HIBYTE(v))
	assert.Equal(uint16(0x1122), HIWORD(v))
	assert.Equal(byte(0x44), LOBYTE(v))
	assert.Equal(uint16(0x3344), LOWORD(v))
}
