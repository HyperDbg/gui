package syscall-hook


type (
SsdtHook interface{
SyscallHookGetKernelBase()(ok bool)//col:75
SyscallHookFindSsdt()(ok bool)//col:150
SyscallHookGetFunctionAddress()(ok bool)//col:201
NtCreateFileHook()(ok bool)//col:267
SyscallHookTest()(ok bool)//col:306
}

)

func NewSsdtHook() { return & ssdtHook{} }

func (s *ssdtHook)SyscallHookGetKernelBase()(ok bool){//col:75














































return true
}


func (s *ssdtHook)SyscallHookFindSsdt()(ok bool){//col:150










































return true
}


func (s *ssdtHook)SyscallHookGetFunctionAddress()(ok bool){//col:201






























return true
}


func (s *ssdtHook)NtCreateFileHook()(ok bool){//col:267









































return true
}


func (s *ssdtHook)SyscallHookTest()(ok bool){//col:306














return true
}




