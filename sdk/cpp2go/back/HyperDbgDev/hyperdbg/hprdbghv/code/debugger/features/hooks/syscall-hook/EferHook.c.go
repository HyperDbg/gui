package syscall-hook


type (
EferHook interface{
SyscallHookConfigureEFER()(ok bool)//col:104
SyscallHookEmulateSYSCALL()(ok bool)//col:172
SyscallHookEmulateSYSRET()(ok bool)//col:219
SyscallHookHandleUD()(ok bool)//col:382
}

)

func NewEferHook() { return & eferHook{} }

func (e *eferHook)SyscallHookConfigureEFER()(ok bool){//col:104




























return true
}


func (e *eferHook)SyscallHookEmulateSYSCALL()(ok bool){//col:172























return true
}


func (e *eferHook)SyscallHookEmulateSYSRET()(ok bool){//col:219















return true
}


func (e *eferHook)SyscallHookHandleUD()(ok bool){//col:382




























































return true
}




