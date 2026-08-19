// Package common — spinlock.go
//
// Implements the spinlock wrapper. The C++ counterpart is
// libhyperdbg/code/common/spinlock.cpp; it implements a custom backoff
// spinlock (derived from hvpp by Petr Benes) with three entry points:
//
//	SpinlockTryLock(Lock)             — non-blocking acquire
//	SpinlockLock(Lock)                — blocking acquire with exponential
//	                                    backoff capped at MaxWait (65536)
//	SpinlockLockWithCustomWait(Lock, MaxWait) — variant with a custom cap
//	SpinlockUnlock(Lock)              — release
//
// The C++ spinlock is intended for VMX-root mode where ordinary OS
// synchronisation primitives cannot be used. In the Go user-mode rewrite we
// do not need a custom spinlock: sync.Mutex provides the same semantics with
// better behaviour (it yields to the scheduler instead of busy-waiting, and
// it integrates with the runtime's deadlock detector). This file therefore
// exposes a Spinlock type that wraps sync.Mutex and preserves the C++ API
// shape (TryLock / Lock / LockWithCustomWait / Unlock) so that ports of
// C++ code that used SpinlockLock/Unlock can be mechanical.
//
// The backoff parameter (MaximumWait) is accepted for API compatibility but
// has no effect in the Go implementation — sync.Mutex does not expose a
// spin count.
package common

import "sync"

// MaxWait mirrors the C++ MaxWait constant (65536). It is the default upper
// bound on the backoff loop in the C++ implementation. In the Go wrapper it
// is retained only for API compatibility.
const MaxWait = 65536

// Spinlock wraps sync.Mutex to mirror the C++ SpinlockLock/Unlock API.
//
// The zero value is a usable, unlocked spinlock.
type Spinlock struct {
	mu sync.Mutex
}

// NewSpinlock returns a new unlocked Spinlock. Equivalent to taking the
// zero value; provided for explicitness.
func NewSpinlock() *Spinlock {
	return &Spinlock{}
}

// TryLock attempts to acquire the lock without blocking. Returns true if the
// lock was acquired. Mirrors the C++ SpinlockTryLock function.
func (s *Spinlock) TryLock() bool {
	return s.mu.TryLock()
}

// Lock acquires the lock, blocking until it is available. Mirrors the C++
// SpinlockLock function. The exponential backoff from the C++ implementation
// is replaced by sync.Mutex's built-in spinning + scheduler yield, which is
// strictly better in user mode.
func (s *Spinlock) Lock() {
	s.mu.Lock()
}

// LockWithCustomWait acquires the lock, blocking until it is available. The
// maximumWait parameter is accepted for API compatibility with the C++
// SpinlockLockWithCustomWait function but has no effect in the Go
// implementation (sync.Mutex does not expose a spin count).
func (s *Spinlock) LockWithCustomWait(maximumWait uint32) {
	_ = maximumWait // unused: see comment on Spinlock
	s.mu.Lock()
}

// Unlock releases the lock. Mirrors the C++ SpinlockUnlock function.
func (s *Spinlock) Unlock() {
	s.mu.Unlock()
}
