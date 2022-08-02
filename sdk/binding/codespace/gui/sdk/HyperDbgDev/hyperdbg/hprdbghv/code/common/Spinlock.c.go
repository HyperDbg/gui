package common



type (
	Spinlock interface {
		SpinlockTryLock() (ok bool)                    //col:21
		SpinlockInterlockedCompareExchange() (ok bool) //col:42
		SpinlockLockWithCustomWait() (ok bool)         //col:60
		SpinlockUnlock() (ok bool)                     //col:64
	}
	spinlock struct{}
)

func NewSpinlock() Spinlock { return &spinlock{} }

func (s *spinlock) SpinlockTryLock() (ok bool) { //col:21


























	return true
}


func (s *spinlock) SpinlockInterlockedCompareExchange() (ok bool) { //col:42

























	return true
}


func (s *spinlock) SpinlockLockWithCustomWait() (ok bool) { //col:60





















	return true
}


func (s *spinlock) SpinlockUnlock() (ok bool) { //col:64







	return true
}


