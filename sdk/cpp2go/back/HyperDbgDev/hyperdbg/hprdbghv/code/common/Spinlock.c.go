package common


type (
Spinlock interface{
SpinlockTryLock()(ok bool)//col:44
SpinlockLock()(ok bool)//col:78
SpinlockInterlockedCompareExchange()(ok bool)//col:118
SpinlockLockWithCustomWait()(ok bool)//col:153
SpinlockUnlock()(ok bool)//col:165
}






)

func NewSpinlock() { return & spinlock{} }

func (s *spinlock)SpinlockTryLock()(ok bool){//col:44




return true
}







func (s *spinlock)SpinlockLock()(ok bool){//col:78



















return true
}







func (s *spinlock)SpinlockInterlockedCompareExchange()(ok bool){//col:118






















return true
}







func (s *spinlock)SpinlockLockWithCustomWait()(ok bool){//col:153



















return true
}







func (s *spinlock)SpinlockUnlock()(ok bool){//col:165




return true
}









