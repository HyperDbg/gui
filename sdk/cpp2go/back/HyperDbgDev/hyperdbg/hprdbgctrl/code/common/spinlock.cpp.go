package common
//back\HyperDbgDev\hyperdbg\hprdbgctrl\code\common\spinlock.cpp.back

type (
Spinlock interface{
SpinlockTryLock()(ok bool)//col:4
SpinlockLock()(ok bool)//col:23
SpinlockLockWithCustomWait()(ok bool)//col:42
SpinlockUnlock()(ok bool)//col:46
}
)

func NewSpinlock() { return & spinlock{} }

func (s *spinlock)SpinlockTryLock()(ok bool){//col:4
/*SpinlockTryLock(volatile LONG * Lock)
{
    return (!(*Lock) && !_interlockedbittestandset(Lock, 0));
}*/
return true
}

func (s *spinlock)SpinlockLock()(ok bool){//col:23
/*SpinlockLock(volatile LONG * Lock)
{
    unsigned wait = 1;
    while (!SpinlockTryLock(Lock))
    {
        for (unsigned i = 0; i < wait; ++i)
        {
            _mm_pause();
        }
        if (wait * 2 > MaxWait)
        {
            wait = MaxWait;
        }
        else
        {
            wait = wait * 2;
        }
    }
}*/
return true
}

func (s *spinlock)SpinlockLockWithCustomWait()(ok bool){//col:42
/*SpinlockLockWithCustomWait(volatile LONG * Lock, unsigned MaximumWait)
{
    unsigned wait = 1;
    while (!SpinlockTryLock(Lock))
    {
        for (unsigned i = 0; i < wait; ++i)
        {
            _mm_pause();
        }
        if (wait * 2 > MaximumWait)
        {
            wait = MaximumWait;
        }
        else
        {
            wait = wait * 2;
        }
    }
}*/
return true
}

func (s *spinlock)SpinlockUnlock()(ok bool){//col:46
/*SpinlockUnlock(volatile LONG * Lock)
{
    *Lock = 0;
}*/
return true
}



