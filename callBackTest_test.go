package main

import (
	"testing"

	"github.com/ddkwork/golibrary/mylog"
)

type callBackTest3 struct {
	name     string
	callback func(arg string)
}

func (c *callBackTest3) Name(callback func(arg string)) string {
	callback("")
	return c.name
}

func Test_callBackTest3_Name(t *testing.T) {
	c := &callBackTest3{
		name:     "abc",
		callback: nil,
	}
	c.callback = func(arg string) {
		arg = "costomized"
		// See here, we should pass in a memory variable requested by go,
		// and write the return string of the command function to the variable memory from go after the callback of c. At this point,
		// the buffer returned by the command function should be synchronized with the memory of the variable from go regardless of whether or not it is returned.
		// this should make the out-of-synchronization problem disappear
		c.name += arg //
		mylog.Info("callback called with arg: %s", arg, "", c.name)
	}
	println(c.Name(c.callback))
}
