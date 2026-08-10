package subset

import (
	"strings"
	"testing"

	"github.com/hyperdbg/go-bridge/protocol"
)

// TestValidate_Accepted verifies that supported Go subset source passes
// validation without error.
func TestValidate_Accepted(t *testing.T) {
	cases := []struct {
		name string
		src  string
	}{
		{
			"simple_assign_and_call",
			`package hook
func hook(ctx *HookCtx) {
	ret := ctx.StackReadQword(0) & 0xFFFFFFFF
	ctx.Printf("RAH ret=%x", ret)
}`,
		},
		{
			"if_else",
			`package hook
func hook(ctx *HookCtx) {
	x := ctx.GetIP()
	if x > 0x1000 {
		ctx.Break()
	} else {
		ctx.Continue()
	}
}`,
		},
		{
			"for_loop",
			`package hook
func hook(ctx *HookCtx) {
	var i uint64 = 0
	for i < 10 {
		ctx.Printf("i=%d", i)
		i = i + 1
	}
}`,
		},
		{
			"var_decl",
			`package hook
func hook(ctx *HookCtx) {
	var addr uint64 = 0x00c10000
	var count uint32 = 64
	_ = addr
	_ = count
}`,
		},
		{
			"bitwise_ops",
			`package hook
func hook(ctx *HookCtx) {
	x := ctx.StackReadQword(0)
	y := x & 0xFFFF
	z := y | 0x1000
	w := z ^ 0xFF
	_ = w
}`,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			v := New()
			if err := v.ValidateSrc(c.src); err != nil {
				t.Errorf("expected acceptance, got error: %v", err)
			}
		})
	}
}

// TestValidate_Rejected verifies that unsupported constructs are rejected
// with a ValidationError.
func TestValidate_Rejected(t *testing.T) {
	cases := []struct {
		name        string
		src         string
		errContains string
	}{
		{
			"goroutine",
			`package hook
func hook(ctx *HookCtx) {
	go ctx.Break()
}`,
			"goroutine",
		},
		{
			"channel",
			`package hook
func hook(ctx *HookCtx) {
	ch := make(chan int)
	_ = ch
}`,
			"make",
		},
		{
			"defer",
			`package hook
func hook(ctx *HookCtx) {
	defer ctx.Break()
}`,
			"defer",
		},
		{
			"map_type",
			`package hook
func hook(ctx *HookCtx) {
	m := map[string]uint64{}
	_ = m
}`,
			"map",
		},
		{
			"interface_type",
			`package hook
type Iface interface { Foo() }
func hook(ctx *HookCtx) {
	_ = ctx
}`,
			"interface",
		},
		{
			"reflect_import",
			`package hook
import "reflect"
func hook(ctx *HookCtx) {
	_ = reflect.TypeOf
}`,
			"reflect",
		},
		{
			"unsafe_import",
			`package hook
import "unsafe"
func hook(ctx *HookCtx) {
	_ = unsafe.Pointer
}`,
			"unsafe",
		},
		{
			"slice_type",
			`package hook
func hook(ctx *HookCtx) {
	s := []byte{1, 2, 3}
	_ = s
}`,
			"slice",
		},
		{
			"new_builtin",
			`package hook
func hook(ctx *HookCtx) {
	p := new(uint64)
	_ = p
}`,
			"new",
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			v := New()
			err := v.ValidateSrc(c.src)
			if err == nil {
				t.Errorf("expected rejection, got nil error")
				return
			}
			if !strings.Contains(err.Error(), c.errContains) {
				t.Errorf("error %q does not contain %q", err.Error(), c.errContains)
			}
			if !protocol.IsSubsetUnsupported(err) {
				t.Errorf("expected ValidationError, got %T", err)
			}
		})
	}
}
