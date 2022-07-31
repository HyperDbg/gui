package user-level
type (
UserListening interface{
UdHandleUserDebuggerPausing()(ok bool)//col:113
}

)
func NewUserListening() { return & userListening{} }
func (u *userListening)UdHandleUserDebuggerPausing()(ok bool){//col:113
return true
}

