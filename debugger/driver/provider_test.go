package driver

import (
	"bytes"
	"testing"

	"github.com/ddkwork/golibrary/std/mylog"
	"github.com/stretchr/testify/assert"
)

func CTL_CODE(deviceType, function, method, access uint32) uint32 {
	return ((deviceType) << 16) | ((access) << 14) | ((function) << 2) | (method)
}

const (
	FILE_DEVICE_UNKNOWN = 0x00000022
	METHOD_BUFFERED     = 0
	FILE_ANY_ACCESS     = 0
)

var (
	IOCTL_SEND_DATA    = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x800, METHOD_BUFFERED, FILE_ANY_ACCESS)
	IOCTL_RECEIVE_DATA = CTL_CODE(FILE_DEVICE_UNKNOWN, 0x801, METHOD_BUFFERED, FILE_ANY_ACCESS)
)

func TestProvider(t *testing.T) {
	mylog.Call(func() {
		p := New("sysDemo.sys", "sysDemo", "\\\\.\\sysDemo")

		p.Install()
		p.Start()

		connected := p.IsConnected()
		mylog.Info("IsConnected", connected)

		testStrings := []string{"hello", "world", "test", "driver", "communication"}
		for _, send := range testStrings {
			p.Send(bytes.NewBufferString(send), IOCTL_SEND_DATA)
			receive := p.Receive(IOCTL_RECEIVE_DATA)
			mylog.Warning(send, receive.String())
			switch send {
			case "hello":
				assert.Equal(t, "received data by user hello", receive.String())
			case "world":
				assert.Equal(t, "received data by user world", receive.String())
			case "test":
				assert.Equal(t, "received data by user test", receive.String())
			case "driver":
				assert.Equal(t, "received data by user driver", receive.String())
			case "communication":
				assert.Equal(t, "received data by user communication", receive.String())
			default:
				panic(send)
			}
		}

		p.Stop()
		p.Uninstall()
	})
}
