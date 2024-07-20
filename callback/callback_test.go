package callback

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"

	"github.com/ddkwork/app/bindgen/clang"
	"github.com/ddkwork/app/bindgen/gengo"
	"github.com/ddkwork/golibrary/mylog"
)

func TestDemoDll(t *testing.T) {
	pkg := gengo.NewPackage("callback")
	path := "src/callback.h"
	mylog.Check(pkg.Transform("callback", &clang.Options{
		Sources:          []string{path},
		AdditionalParams: []string{},
	}),
	)
	mylog.Check(pkg.WriteToDir("."))

	pfn := func(msg *byte) {
		fmt.Println("Received data:", BytePointerToString(msg))
	}
	SetTextMessageCallback(unsafe.Pointer(reflect.ValueOf(pfn).Pointer()))
	ShowMessages(nil)
}

func BytePointerToString(ptr *byte) string {
	var bytes []byte
	for *ptr != 0 {
		bytes = append(bytes, *ptr)
		ptr = (*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(ptr)) + 1))
	}
	return string(bytes)
}
