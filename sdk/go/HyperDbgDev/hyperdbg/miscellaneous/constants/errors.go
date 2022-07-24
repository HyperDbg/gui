package constants

import (
	_ "embed"
)

//go:embed "D:\\codespace\\workspace\\src\\cppkit\\gui\\sdk\\HyperDbgDev\\hyperdbg\\miscellaneous\\constants\\errors.txt"
var errorsBuf string

type (
	Errors interface {
		//Fn() (ok bool)
	}
	errors struct{}
)

func Newerrors() Errors { return &errors{} }
