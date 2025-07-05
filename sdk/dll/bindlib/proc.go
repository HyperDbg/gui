package bindlib

import (
	"sync/atomic"

	"github.com/ddkwork/golibrary/std/mylog"
	"github.com/ebitengine/purego"
)

const INVALID_PROC = ^uintptr(0)

type Proc struct {
	library *Library
	name    string
	cache   atomic.Uintptr
}

func (lp *Proc) addrSlow() uintptr {
	proc := lp.cache.Load()
	if proc == 0 {
		lib := mylog.Check2(lp.library.Get())
		proc = lib.Lookup(lp.name)
		if proc == 0 {
			proc = INVALID_PROC
		}
		lp.cache.Store(proc)
	}
	return proc
}

func (lp *Proc) Addr() uintptr {
	proc := lp.cache.Load()
	if proc == 0 {
		proc = lp.addrSlow()
	}
	return proc
}

func CCall0(proc uintptr) uintptr {
	r1, _, _ := purego.SyscallN(proc)
	return r1
}

func CCall1(proc uintptr, a uintptr) uintptr {
	r1, _, _ := purego.SyscallN(proc, a)
	return r1
}

func CCall2(proc uintptr, a, b uintptr) uintptr {
	r1, _, _ := purego.SyscallN(proc, a, b)
	return r1
}

func CCall3(proc uintptr, a, b, c uintptr) uintptr {
	r1, _, _ := purego.SyscallN(proc, a, b, c)
	return r1
}

func CCall4(proc uintptr, a, b, c, d uintptr) uintptr {
	r1, _, _ := purego.SyscallN(proc, a, b, c)
	return r1
}

func CCall5(proc uintptr, a, b, c, d, e uintptr) uintptr {
	r1, _, _ := purego.SyscallN(proc, a, b, c, d, e)
	return r1
}

func CCall6(proc uintptr, a, b, c, d, e, f uintptr) uintptr {
	r1, _, _ := purego.SyscallN(proc, a, b, c, d, e)
	return r1
}

func CCall7(proc uintptr, a, b, c, d, e, f, g uintptr) uintptr {
	r1, _, _ := purego.SyscallN(proc, a, b, c, d, e, f, g)
	return r1
}

func CCall8(proc uintptr, a, b, c, d, e, f, g, h uintptr) uintptr {
	r1, _, _ := purego.SyscallN(proc, a, b, c, d, e, f, g, h)
	return r1
}

func CCall9(proc uintptr, a, b, c, d, e, f, g, h, i uintptr) uintptr {
	r1, _, _ := purego.SyscallN(proc, a, b, c, d, e, f, g, h, i)
	return r1
}

func CCall10(proc uintptr, a, b, c, d, e, f, g, h, i, j uintptr) uintptr {
	r1, _, _ := purego.SyscallN(proc, a, b, c, d, e, f, g, h, i)
	return r1
}

func CCall11(proc uintptr, a, b, c, d, e, f, g, h, i, j, k uintptr) uintptr {
	r1, _, _ := purego.SyscallN(proc, a, b, c, d, e, f, g, h, i)
	return r1
}

func CCall12(proc uintptr, a, b, c, d, e, f, g, h, i, j, k, l uintptr) uintptr {
	r1, _, _ := purego.SyscallN(proc, a, b, c, d, e, f, g, h, i, j, k, l)
	return r1
}

func (lp *Proc) Call0() (r1 uintptr) {
	return CCall0(lp.Addr())
}

func (lp *Proc) Call1(a uintptr) (r1 uintptr) {
	return CCall1(lp.Addr(), a)
}

func (lp *Proc) Call2(a, b uintptr) (r1 uintptr) {
	return CCall2(lp.Addr(), a, b)
}

func (lp *Proc) Call3(a, b, c uintptr) (r1 uintptr) {
	return CCall3(lp.Addr(), a, b, c)
}

func (lp *Proc) Call4(a, b, c, d uintptr) (r1 uintptr) {
	return CCall4(lp.Addr(), a, b, c, d)
}

func (lp *Proc) Call5(a, b, c, d, e uintptr) (r1 uintptr) {
	return CCall5(lp.Addr(), a, b, c, d, e)
}

func (lp *Proc) Call6(a, b, c, d, e, f uintptr) (r1 uintptr) {
	return CCall6(lp.Addr(), a, b, c, d, e, f)
}

func (lp *Proc) Call7(a, b, c, d, e, f, g uintptr) (r1 uintptr) {
	return CCall7(lp.Addr(), a, b, c, d, e, f, g)
}

func (lp *Proc) Call8(a, b, c, d, e, f, g, h uintptr) (r1 uintptr) {
	return CCall8(lp.Addr(), a, b, c, d, e, f, g, h)
}

func (lp *Proc) Call9(a, b, c, d, e, f, g, h, i uintptr) (r1 uintptr) {
	return CCall9(lp.Addr(), a, b, c, d, e, f, g, h, i)
}

func (lp *Proc) Call10(a, b, c, d, e, f, g, h, i, j uintptr) (r1 uintptr) {
	return CCall10(lp.Addr(), a, b, c, d, e, f, g, h, i, j)
}

func (lp *Proc) Call11(a, b, c, d, e, f, g, h, i, j, k uintptr) (r1 uintptr) {
	return CCall11(lp.Addr(), a, b, c, d, e, f, g, h, i, j, k)
}

func (lp *Proc) Call12(a, b, c, d, e, f, g, h, i, j, k, l uintptr) (r1 uintptr) {
	return CCall12(lp.Addr(), a, b, c, d, e, f, g, h, i, j, k, l)
}

type PreloadProc uintptr

func (p PreloadProc) Addr() uintptr { return uintptr(p) }

func (lp PreloadProc) Call0() (r1 uintptr) {
	return CCall0(lp.Addr())
}

func (lp PreloadProc) Call1(a uintptr) (r1 uintptr) {
	return CCall1(lp.Addr(), a)
}

func (lp PreloadProc) Call2(a, b uintptr) (r1 uintptr) {
	return CCall2(lp.Addr(), a, b)
}

func (lp PreloadProc) Call3(a, b, c uintptr) (r1 uintptr) {
	return CCall3(lp.Addr(), a, b, c)
}

func (lp PreloadProc) Call4(a, b, c, d uintptr) (r1 uintptr) {
	return CCall4(lp.Addr(), a, b, c, d)
}

func (lp PreloadProc) Call5(a, b, c, d, e uintptr) (r1 uintptr) {
	return CCall5(lp.Addr(), a, b, c, d, e)
}

func (lp PreloadProc) Call6(a, b, c, d, e, f uintptr) (r1 uintptr) {
	return CCall6(lp.Addr(), a, b, c, d, e, f)
}

func (lp PreloadProc) Call7(a, b, c, d, e, f, g uintptr) (r1 uintptr) {
	return CCall7(lp.Addr(), a, b, c, d, e, f, g)
}

func (lp PreloadProc) Call8(a, b, c, d, e, f, g, h uintptr) (r1 uintptr) {
	return CCall8(lp.Addr(), a, b, c, d, e, f, g, h)
}

func (lp PreloadProc) Call9(a, b, c, d, e, f, g, h, i uintptr) (r1 uintptr) {
	return CCall9(lp.Addr(), a, b, c, d, e, f, g, h, i)
}

func (lp PreloadProc) Call10(a, b, c, d, e, f, g, h, i, j uintptr) (r1 uintptr) {
	return CCall10(lp.Addr(), a, b, c, d, e, f, g, h, i, j)
}

func (lp PreloadProc) Call11(a, b, c, d, e, f, g, h, i, j, k uintptr) (r1 uintptr) {
	return CCall11(lp.Addr(), a, b, c, d, e, f, g, h, i, j, k)
}

func (lp PreloadProc) Call12(a, b, c, d, e, f, g, h, i, j, k, l uintptr) (r1 uintptr) {
	return CCall12(lp.Addr(), a, b, c, d, e, f, g, h, i, j, k, l)
}
