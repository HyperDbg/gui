package Headers

import (
	"fmt"
	"github.com/ddkwork/golibrary/mylog"
	"reflect"
	"testing"
)

func TestSizeof(t *testing.T) {
	//type (
	//	DEBUGGER_REMOTE_PACKET struct {
	//		Checksum                   byte
	//		Indicator                  uint64
	//		TypeOfThePacket            DEBUGGER_REMOTE_PACKET_TYPE
	//		RequestedActionOfThePacket DEBUGGER_REMOTE_PACKET_REQUESTED_ACTION
	//	}
	//	PDEBUGGER_REMOTE_PACKET *DEBUGGER_REMOTE_PACKET
	//)

	// c++ 1+8+1+1= 11 bytes, but go can not 1 byte align
	//unsafe.Sizeof(unsafe.Pointer(DEBUGGER_REMOTE_PACKET))
	debuggerRemotePacket := DEBUGGER_REMOTE_PACKET{
		Checksum:                   0,
		Indicator:                  0,
		TypeOfThePacket:            0,
		RequestedActionOfThePacket: 0,
	}
	println(SizeOf(debuggerRemotePacket))
	fmt.Println(DEBUGGER_REMOTE_PACKET{})
}

func SizeOf(object any) (size int) {
	fields := reflect.VisibleFields(reflect.TypeOf(object))
	for _, field := range fields {
		switch field.Type.Kind() {
		case reflect.Uint8:
			size += 1
		case reflect.Uint64:
			size += 8
			//todo case all type
		}
	}
	return
}

func TestHIBYTE(t *testing.T) {
	v := uint32(0x11223344)
	assert := mylog.Assert(t)
	assert.Equal(byte(0x11), HIBYTE(v))
	assert.Equal(uint16(0x1122), HIWORD(v))
	assert.Equal(byte(0x44), LOBYTE(v))
	assert.Equal(uint16(0x3344), LOWORD(v))
}
