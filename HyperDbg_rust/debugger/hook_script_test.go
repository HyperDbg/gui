package debugger

import (
	"strings"
	"testing"
)

func TestExtractClosureBody_SimpleAssignment(t *testing.T) {
	fn := func(ctx *HookContext) {
		ctx.Return = 0
	}

	code := ExtractClosureBody(fn)
	if code == "" {
		t.Error("Expected non-empty code")
		return
	}

	if !strings.Contains(code, "ctx.Return = 0") {
		t.Errorf("Expected 'ctx.Return = 0' in code, got: %s", code)
	}
}

func TestExtractClosureBody_TypeAssertion(t *testing.T) {
	fn := func(ctx *HookContext) {
		args := ctx.Args.(*NtDeviceIoControlFileArgs)
		_ = args
	}

	code := ExtractClosureBody(fn)
	if code == "" {
		t.Error("Expected non-empty code")
		return
	}

	if !strings.Contains(code, "ctx.Args") {
		t.Errorf("Expected 'ctx.Args' in code, got: %s", code)
	}
}

func TestExtractClosureBody_MemoryWrite(t *testing.T) {
	fn := func(ctx *HookContext) {
		args := ctx.Args.(*NtDeviceIoControlFileArgs)
		if args.IoControlCode == 0x12345678 {
			args.WriteOutputBytes([]byte{0x00, 0x11, 0x22})
		}
	}

	code := ExtractClosureBody(fn)
	if code == "" {
		t.Error("Expected non-empty code")
		return
	}

	if !strings.Contains(code, "IoControlCode") {
		t.Errorf("Expected 'IoControlCode' in code, got: %s", code)
	}
}

func TestExtractClosureBody_SwitchStatement(t *testing.T) {
	fn := func(ctx *HookContext) {
		args := ctx.Args.(*NtDeviceIoControlFileArgs)
		switch args.IoControlCode {
		case 0x0007c088:
			ctx.Return = 0
		case 0x002d1400:
			ctx.Return = 1
		}
	}

	code := ExtractClosureBody(fn)
	if code == "" {
		t.Error("Expected non-empty code")
		return
	}

	if !strings.Contains(code, "switch") {
		t.Errorf("Expected 'switch' in code, got: %s", code)
	}
}

func TestExtractClosureBody_MultipleStatements(t *testing.T) {
	fn := func(ctx *HookContext) {
		args := ctx.Args.(*NtDeviceIoControlFileArgs)
		args.WriteOutputBytes([]byte{1, 2, 3})
		args.WriteOutputString("test")
		ctx.Return = 0
	}

	code := ExtractClosureBody(fn)
	if code == "" {
		t.Error("Expected non-empty code")
		return
	}

	if !strings.Contains(code, "WriteOutputBytes") {
		t.Errorf("Expected 'WriteOutputBytes' in code, got: %s", code)
	}
	if !strings.Contains(code, "WriteOutputString") {
		t.Errorf("Expected 'WriteOutputString' in code, got: %s", code)
	}
}

func TestExtractClosureBody_NilHandler(t *testing.T) {
	code := ExtractClosureBody(nil)
	if code != "" {
		t.Errorf("Expected empty code for nil handler, got: %s", code)
	}
}

func TestExtractClosureBodyRaw(t *testing.T) {
	fn := func(ctx *HookContext) {
		args := ctx.Args.(*NtDeviceIoControlFileArgs)
		ctx.Return = 0
	}

	code := ExtractClosureBodyRaw(fn)
	if code == "" {
		t.Error("Expected non-empty code")
		return
	}

	if !strings.Contains(code, "ctx.Return") {
		t.Errorf("Expected 'ctx.Return' in code, got: %s", code)
	}
}

func TestSourceCache(t *testing.T) {
	ClearSourceCache()

	fn := func(ctx *HookContext) {
		ctx.Return = 0
	}

	code1 := ExtractClosureBody(fn)
	code2 := ExtractClosureBody(fn)

	if code1 != code2 {
		t.Error("Expected same code from cache")
	}
}

func TestClearSourceCache(t *testing.T) {
	fn := func(ctx *HookContext) {
		ctx.Return = 0
	}

	_ = ExtractClosureBody(fn)
	ClearSourceCache()

	if len(globalSourceCache.files) != 0 {
		t.Error("Expected empty cache after clear")
	}
}
