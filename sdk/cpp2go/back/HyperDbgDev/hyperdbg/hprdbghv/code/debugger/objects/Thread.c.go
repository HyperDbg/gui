package objects
type (
	Thread interface {
		ThreadHandleThreadChange() (ok bool)                        //col:46
		ThreadSwitch() (ok bool)                                    //col:109
		ThreadShowList() (ok bool)                                  //col:314
		ThreadInterpretThread() (ok bool)                           //col:406
		ThreadDetectChangeByDebugRegisterOnGs() (ok bool)           //col:538
		ThreadDetectChangeByInterceptingClockInterrupts() (ok bool) //col:577
		ThreadEnableOrDisableThreadChangeMonitor() (ok bool)        //col:605
		ThreadQueryCount() (ok bool)                                //col:636
		ThreadQueryList() (ok bool)                                 //col:664
		ThreadQueryDetails() (ok bool)                              //col:688
	}
)
func NewThread() { return &thread{} }
func (t *thread) ThreadHandleThreadChange() (ok bool) { //col:46
	return true
}

func (t *thread) ThreadSwitch() (ok bool) { //col:109
	return true
}

func (t *thread) ThreadShowList() (ok bool) { //col:314
	return true
}

func (t *thread) ThreadInterpretThread() (ok bool) { //col:406
	return true
}

func (t *thread) ThreadDetectChangeByDebugRegisterOnGs() (ok bool) { //col:538
	return true
}

func (t *thread) ThreadDetectChangeByInterceptingClockInterrupts() (ok bool) { //col:577
	return true
}

func (t *thread) ThreadEnableOrDisableThreadChangeMonitor() (ok bool) { //col:605
	return true
}

func (t *thread) ThreadQueryCount() (ok bool) { //col:636
	return true
}

func (t *thread) ThreadQueryList() (ok bool) { //col:664
	return true
}

func (t *thread) ThreadQueryDetails() (ok bool) { //col:688
	return true
}

