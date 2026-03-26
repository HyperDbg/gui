package main

import (
	"bytes"

	"github.com/ddkwork/HyperDbg/debugger/driver"
	"github.com/ddkwork/golibrary/std/mylog"
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

func main() {
	mylog.Call(func() {run()})
}

	func run () {
	p := driver.New(`D:\ux\examples\hypedbg\rust-driver\examples\sysdemo\sysdemo.sys`, "sysDemo", "\\\\.\\sysDemo")

	p.Install()
	p.Start()

	connected := p.IsConnected()
	mylog.Info(connected)

	testStrings := []string{"hello", "world", "test", "driver", "communication"}
	for _, send := range testStrings {
		p.Send(bytes.NewBufferString(send), IOCTL_SEND_DATA)
		receive := p.Receive(IOCTL_RECEIVE_DATA)
		mylog.Warning(send, receive.String())
	}

	p.Stop()
	p.Uninstall()
}
