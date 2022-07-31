package syscall-hook
type (
	SsdtHook interface {
		SyscallHookGetKernelBase() (ok bool)      //col:75
		SyscallHookFindSsdt() (ok bool)           //col:149
		SyscallHookGetFunctionAddress() (ok bool) //col:199
		NtCreateFileHook() (ok bool)              //col:264
		SyscallHookTest() (ok bool)               //col:302
	}
)
func NewSsdtHook() { return &ssdtHook{} }
func (s *ssdtHook) SyscallHookGetKernelBase() (ok bool) { //col:75
	return true
}

func (s *ssdtHook) SyscallHookFindSsdt() (ok bool) { //col:149
	return true
}

func (s *ssdtHook) SyscallHookGetFunctionAddress() (ok bool) { //col:199
	return true
}

func (s *ssdtHook) NtCreateFileHook() (ok bool) { //col:264
	return true
}

func (s *ssdtHook) SyscallHookTest() (ok bool) { //col:302
	return true
}

