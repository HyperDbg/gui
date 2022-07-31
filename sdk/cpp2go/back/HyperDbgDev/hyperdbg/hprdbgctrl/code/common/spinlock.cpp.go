package common


type (
Spinlock interface{
SpinlockTryLock()(ok bool)//col:44
SpinlockLock()(ok bool)//col:78
SpinlockLockWithCustomWait()(ok bool)//col:113
SpinlockUnlock()(ok bool)//col:125
}

)

func NewSpinlock() { return & spinlock{} }

func (s *spinlock)SpinlockTryLock()(ok bool){//col:44




return true
}


func (s *spinlock)SpinlockLock()(ok bool){//col:78



















return true
}


func (s *spinlock)SpinlockLockWithCustomWait()(ok bool){//col:113



















return true
}


func (s *spinlock)SpinlockUnlock()(ok bool){//col:125




return true
}




