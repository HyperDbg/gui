package main

import (
	"testing"

	"github.com/ddkwork/golibrary/std/mylog"
	"github.com/ddkwork/golibrary/std/stream"
)

func TestFixError(t *testing.T) {
	mylog.Call(func() {
		stream.Fmt(".")
		stream.Fix(".")
	})
}
