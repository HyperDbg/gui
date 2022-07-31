package common
type (
	Spinlock interface {
		*Based on my benchmarks, this simple implementation beats other ()
(ok bool) //col:44
SpinlockLock()(ok bool)                       //col:77
SpinlockInterlockedCompareExchange()(ok bool) //col:116
SpinlockLockWithCustomWait()(ok bool) //col:150
SpinlockUnlock()(ok bool) //col:161
}

)
func NewSpinlock() { return &spinlock{} }
func (s *spinlock) * Based on my benchmarks, this simple implementation beats other ()(ok bool) { //col:44
	return true
}

func (s *spinlock) SpinlockLock() (ok bool) { //col:77
	return true
}

func (s *spinlock) SpinlockInterlockedCompareExchange() (ok bool) { //col:116
	return true
}

func (s *spinlock) SpinlockLockWithCustomWait() (ok bool) { //col:150
	return true
}

func (s *spinlock) SpinlockUnlock() (ok bool) { //col:161
	return true
}

