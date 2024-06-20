package libdemo

import (
	"testing"

	"github.com/can1357/gengo/clang"
	"github.com/can1357/gengo/gengo"
	"github.com/ddkwork/golibrary/mylog"
)

func TestHello(t *testing.T) {
	Hello()
}

func TestDemoDll(t *testing.T) {
	pkg := gengo.NewPackage("libdemo")
	path := "cpp\\library.h"
	mylog.Check(pkg.Transform("libdemo", &clang.Options{
		Sources:          []string{path},
		AdditionalParams: []string{},
	}),
	)
	mylog.Check(pkg.WriteToDir("."))
}
