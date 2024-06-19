package ARImpRec

import (
	"testing"

	"github.com/can1357/gengo/clang"
	"github.com/can1357/gengo/gengo"
	"github.com/ddkwork/golibrary/mylog"
)

func TestGetProcName(t *testing.T) {
	GetProcName()
}

func TestARImpRec(t *testing.T) {
	t.Skip("bug to more")
	pkg := gengo.NewPackage("ARImpRec")
	mylog.Check(pkg.Transform("ARImpRec", &clang.Options{
		Sources:          []string{"ARImpRec.h"},
		AdditionalParams: []string{},
	}))
	mylog.Check(pkg.WriteToDir("."))
}
