package common
//back\HyperDbgDev\hyperdbg\hprdbgctrl\code\common\spinlock.cpp.back

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
/* * Based on my benchmarks, this simple implementation beats other (often
 * more complex) spinlock implementations - such as queue spinlocks, ticket
 * spinlocks, MCS locks.  The only difference between this implementation
 * and completely naive spinlock is the "backoff".
 * 
 * Also, benefit of this implementation is that we can use it with
 * STL lock guards, e.g.: std::lock_guard.
 * 
 * Look here for more information:
 * 
 * 
 * 
#include "pch.h"
 * 
static unsigned MaxWait = 65536;
 * 
BOOLEAN
SpinlockTryLock(volatile LONG * Lock)
{
    return (!(*Lock) && !_interlockedbittestandset(Lock, 0));
}*/
return true
}

func (s *spinlock)SpinlockLock()(ok bool){//col:77
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

func (s *spinlock)SpinlockLockWithCustomWait()(ok bool){//col:111
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

func (s *spinlock)SpinlockUnlock()(ok bool){//col:122
/*SpinlockUnlock(volatile LONG * Lock)
{
    *Lock = 0;
}*/
return true
}



