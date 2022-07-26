package common

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\hprdbgctrl\\code\\common\\spinlock.cpp"
var spinlockBuf string

type (
	Spinlock interface {
		//Fn() (ok bool)
	}
	spinlock struct{}
)

func Newspinlock() Spinlock { return &spinlock{} }
