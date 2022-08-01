package common
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbghv\code\common\Spinlock.c.back

type (
Spinlock interface{
SpinlockTryLock()(ok bool)//col:4
SpinlockLock()(ok bool)//col:23
SpinlockInterlockedCompareExchange()(ok bool)//col:45
SpinlockLockWithCustomWait()(ok bool)//col:64
SpinlockUnlock()(ok bool)//col:68
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

func (s *spinlock)SpinlockInterlockedCompareExchange()(ok bool){//col:45
/*SpinlockInterlockedCompareExchange(
    LONG volatile * Destination,
    LONG            Exchange,
    LONG            Comperand)
{
    unsigned wait = 1;
    while (InterlockedCompareExchange(Destination, Exchange, Comperand) != Comperand)
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

func (s *spinlock)SpinlockLockWithCustomWait()(ok bool){//col:64
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

func (s *spinlock)SpinlockUnlock()(ok bool){//col:68
/*SpinlockUnlock(volatile LONG * Lock)
{
    *Lock = 0;
}*/
return true
}



