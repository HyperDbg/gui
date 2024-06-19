package Headers

import (
	"encoding/binary"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSizeof(t *testing.T) {
	assert.Equal(t, 11, binary.Size(DEBUGGER_REMOTE_PACKET{}))
}

func TestHIBYTE(t *testing.T) {
	v := uint32(0x11223344)
	assert.Equal(t, byte(0x11), HIBYTE(v))
	assert.Equal(t, uint16(0x1122), HIWORD(v))
	assert.Equal(t, byte(0x44), LOBYTE(v))
	assert.Equal(t, uint16(0x3344), LOWORD(v))
}
