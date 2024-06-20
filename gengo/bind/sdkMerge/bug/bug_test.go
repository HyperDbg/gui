package bug

import (
	"testing"

	"github.com/can1357/gengo/clang"
	"github.com/can1357/gengo/gengo"
	"github.com/ddkwork/golibrary/mylog"
)

func TestBug(t *testing.T) {
	pkg := gengo.NewPackage("bug")
	path := "bug.hpp"
	mylog.Check(pkg.Transform("bug", &clang.Options{
		Sources:          []string{path},
		AdditionalParams: []string{},
	}),
	)
	mylog.Check(pkg.WriteToDir("./tmp"))
}
