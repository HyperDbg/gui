package user_level



type (
	Ud interface {
		UdInitializeUserDebugger() (ok bool) //col:27
		UdStepInstructions() (ok bool)       //col:40
	}
	ud struct{}
)

func NewUd() Ud { return &ud{} }

func (u *ud) UdInitializeUserDebugger() (ok bool) { //col:27




































	return true
}


func (u *ud) UdStepInstructions() (ok bool) { //col:40

















	return true
}


