package ntdll

import "syscall"

var modntdll = syscall.NewLazyDLL("ntdll")

type Handle syscall.Handle

func fromBool(b bool) (r uintptr) {
	if b {
		r = 1
	}
	return
}
