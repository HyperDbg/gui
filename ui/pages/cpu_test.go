package pages

import (
	"testing"
)

func TestCpuPage_ParseAddr(t *testing.T) {
	c := NewCpu(nil, nil)

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
		// CpuPage.parseAddr 在解析失败时返回 0x10000（默认地址）
		{"invalid", 0x10000},
		{"", 0x10000},
	}

	for _, tc := range cases {
		c.SetAddrInputForTest(tc.input)
		got := c.parseAddr()
		if got != tc.want {
			t.Errorf("parseAddr(%q) = 0x%X, want 0x%X", tc.input, got, tc.want)
		}
	}
}

func TestParseRIP(t *testing.T) {
	cases := []struct {
		name string
		text string
		want uint64
	}{
		{"empty", "", 0},
		{"rip_uppercase_eq", "RIP=000000007FF41234 RFL=00000202", 0x7FF41234},
		{"rip_lowercase_eq", "rip=7ff41234", 0x7FF41234},
		{"rip_colon", "RIP: 7FF41234", 0x7FF41234},
		{"rip_with_0x", "RIP=0x7FF41234", 0x7FF41234},
		{"rip_multiline", "RAX=0000\nRIP=000000007FF41234\nRSP=0000", 0x7FF41234},
		{"rip_no_match", "RAX=0000 RSP=0000", 0},
		{"rip_garbage", "RIP=garbage", 0},
		{"rip_trailing_comma", "RIP=7FF41234,RAX=0", 0x7FF41234},
		{"rip_with_space", "RIP = 7FF41234", 0x7FF41234},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := parseRIP(tc.text)
			if got != tc.want {
				t.Errorf("parseRIP(%q) = 0x%X, want 0x%X", tc.text, got, tc.want)
			}
		})
	}
}
