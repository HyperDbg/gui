package extensioncommands

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\debugger\\commands\\extension-commands\\syscall-sysret.cpp"
var syscallSysretBuf string

type (
	SyscallSysret interface {
		//Fn() (ok bool)
	}
	syscallSysret struct{}
)

func NewsyscallSysret() SyscallSysret { return &syscallSysret{} }
