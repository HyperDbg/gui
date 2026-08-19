package pages

import (
	"testing"
)

func TestEventsPage_ParseAddr(t *testing.T) {
	p := NewEvents(nil)

	cases := []struct {
		input string
		want  uint64
	}{
		{"0x10000", 0x10000},
		{"0X10000", 0x10000},
		{"10000", 0x10000},
		{"0xDEADBEEF", 0xDEADBEEF},
		{"deadbeef", 0xDEADBEEF},
		{"0", 0},
		{"invalid", 0},
		{"", 0},
		{"0xFFFFFFFFFFFFFFFF", 0xFFFFFFFFFFFFFFFF},
	}

	for _, tc := range cases {
		p.SetAddrInputForTest(tc.input)
		got := p.parseAddr()
		if got != tc.want {
			t.Errorf("parseAddr(%q) = 0x%X, want 0x%X", tc.input, got, tc.want)
		}
	}
}

func TestEventsPage_ParsePID(t *testing.T) {
	p := NewEvents(nil)

	cases := []struct {
		input string
		want  uint32
	}{
		{"0x100", 256},
		{"100", 256}, // 十进制回退
		{"0", 0},
		{"0xFF", 255},
		{"invalid", 0},
		{"", 0},
	}

	for _, tc := range cases {
		p.SetPIDInputForTest(tc.input)
		got := p.parsePID()
		if got != tc.want {
			t.Errorf("parsePID(%q) = %d, want %d", tc.input, got, tc.want)
		}
	}
}

func TestEventsPage_ParseTag(t *testing.T) {
	p := NewEvents(nil)

	cases := []struct {
		input string
		want  uint64
	}{
		{"0x1A", 0x1A},
		{"1A", 0x1A},
		{"0", 0},
		{"1234", 0x1234}, // 十六进制优先
		{"invalid", 0},
		{"", 0},
	}

	for _, tc := range cases {
		p.SetTagInputForTest(tc.input)
		got := p.parseTag()
		if got != tc.want {
			t.Errorf("parseTag(%q) = 0x%X, want 0x%X", tc.input, got, tc.want)
		}
	}
}

func TestEventsPage_AddHook(t *testing.T) {
	p := NewEvents(nil)

	p.AddHookForTest(HookRecord{
		Tag: 1, Type: "EptHook", Address: 0x1000, PID: 0, Enabled: true,
	})
	p.AddHookForTest(HookRecord{
		Tag: 2, Type: "SyscallHook", Address: 0, PID: 0, Enabled: true,
	})

	hooks := p.HooksForTest()
	if len(hooks) != 2 {
		t.Fatalf("hooks count = %d, want 2", len(hooks))
	}
	if hooks[0].Tag != 1 || hooks[0].Type != "EptHook" {
		t.Errorf("hooks[0] = %+v", hooks[0])
	}
	if hooks[1].Tag != 2 || hooks[1].Type != "SyscallHook" {
		t.Errorf("hooks[1] = %+v", hooks[1])
	}
}

func TestEventsPage_DisableHook(t *testing.T) {
	p := NewEvents(nil)
	p.AddHookForTest(HookRecord{Tag: 10, Type: "EptHook", Enabled: true})

	p.SetHookEnabledForTest(10, false)

	hooks := p.HooksForTest()
	if len(hooks) != 1 {
		t.Fatalf("hooks count = %d, want 1", len(hooks))
	}
	if hooks[0].Enabled {
		t.Error("hook should be disabled")
	}
}

func TestEventsPage_EnableHook(t *testing.T) {
	p := NewEvents(nil)
	p.AddHookForTest(HookRecord{Tag: 20, Type: "EptHook", Enabled: false})

	p.SetHookEnabledForTest(20, true)

	hooks := p.HooksForTest()
	if !hooks[0].Enabled {
		t.Error("hook should be enabled")
	}
}

func TestEventsPage_RemoveHook(t *testing.T) {
	p := NewEvents(nil)
	p.AddHookForTest(HookRecord{Tag: 1, Type: "EptHook", Enabled: true})
	p.AddHookForTest(HookRecord{Tag: 2, Type: "SyscallHook", Enabled: true})

	p.RemoveHookForTest(1)

	hooks := p.HooksForTest()
	if len(hooks) != 1 {
		t.Fatalf("hooks count after remove = %d, want 1", len(hooks))
	}
	if hooks[0].Tag != 2 {
		t.Errorf("remaining hook tag = %d, want 2", hooks[0].Tag)
	}
}

func TestEventsPage_RemoveHookNotFound(t *testing.T) {
	p := NewEvents(nil)
	p.AddHookForTest(HookRecord{Tag: 1, Type: "EptHook", Enabled: true})

	p.RemoveHookForTest(999)

	hooks := p.HooksForTest()
	if len(hooks) != 1 {
		t.Errorf("hooks count = %d, want 1 (remove non-existent should be no-op)", len(hooks))
	}
}

func TestEventsPage_RefreshListEmpty(t *testing.T) {
	p := NewEvents(nil)
	// 不应 panic
	p.RefreshListForTest()

	hooks := p.HooksForTest()
	if len(hooks) != 0 {
		t.Errorf("hooks count = %d, want 0", len(hooks))
	}
}
