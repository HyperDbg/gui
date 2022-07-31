package syscall-hook
type (
	EferHook interface {
		SyscallHookConfigureEFER() (ok bool)  //col:104
		SyscallHookEmulateSYSCALL() (ok bool) //col:171
		SyscallHookEmulateSYSRET() (ok bool)  //col:217
		SyscallHookHandleUD() (ok bool)       //col:379
	}
)
func NewEferHook() { return &eferHook{} }
func (e *eferHook) SyscallHookConfigureEFER() (ok bool) { //col:104
	return true
}

func (e *eferHook) SyscallHookEmulateSYSCALL() (ok bool) { //col:171
	return true
}

func (e *eferHook) SyscallHookEmulateSYSRET() (ok bool) { //col:217
	return true
}

func (e *eferHook) SyscallHookHandleUD() (ok bool) { //col:379
	return true
}

