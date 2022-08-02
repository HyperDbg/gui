package common



type (
	Spinlock interface {
		SpinlockTryLock() (ok bool)            //col:21
		SpinlockLockWithCustomWait() (ok bool) //col:39
		SpinlockUnlock() (ok bool)             //col:43
	}
	spinlock struct{}
)

func NewSpinlock() Spinlock { return &spinlock{} }

func (s *spinlock) SpinlockTryLock() (ok bool) { //col:21


























	return true
}


func (s *spinlock) SpinlockLockWithCustomWait() (ok bool) { //col:39





















	return true
}


func (s *spinlock) SpinlockUnlock() (ok bool) { //col:43







	return true
}


