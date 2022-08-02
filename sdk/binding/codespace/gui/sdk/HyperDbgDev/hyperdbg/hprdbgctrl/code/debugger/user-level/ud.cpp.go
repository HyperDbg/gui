package user_level



type (
	Ud interface {
		UdInitializeUserDebugger() (ok bool)       //col:35
		UdRemoveActiveDebuggingProcess() (ok bool) //col:39
	}
	ud struct{}
)

func NewUd() Ud { return &ud{} }

func (u *ud) UdInitializeUserDebugger() (ok bool) { //col:35











































	return true
}


func (u *ud) UdRemoveActiveDebuggingProcess() (ok bool) { //col:39







	return true
}


