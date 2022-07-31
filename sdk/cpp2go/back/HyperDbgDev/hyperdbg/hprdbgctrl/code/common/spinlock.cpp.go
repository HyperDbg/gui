package common
type (
Spinlock interface{
 * Based on my benchmarks, this simple implementation beats other ()(ok bool)//col:44
SpinlockLock()(ok bool)//col:77
SpinlockLockWithCustomWait()(ok bool)//col:111
SpinlockUnlock()(ok bool)//col:122
}

)
func NewSpinlock() { return & spinlock{} }
func (s *spinlock) * Based on my benchmarks, this simple implementation beats other ()(ok bool){//col:44
return true
}

func (s *spinlock)SpinlockLock()(ok bool){//col:77
return true
}

func (s *spinlock)SpinlockLockWithCustomWait()(ok bool){//col:111
return true
}

func (s *spinlock)SpinlockUnlock()(ok bool){//col:122
return true
}

