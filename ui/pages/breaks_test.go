package pages

import (
	"testing"
)

func TestBreaksPage_ParseAddr(t *testing.T) {
	p := NewBreaks(nil)

	cases := []struct {
		input string
		want  uint64
	}{
		{"0x10000", 0x10000},
		{"0X10000", 0x10000},
		{"10000", 0x10000},
		{"0xDEAD", 0xDEAD},
		{"dead", 0xDEAD},
		{"0", 0},
		{"invalid", 0},
		{"", 0},
	}

	for _, tc := range cases {
		p.SetAddrInputForTest(tc.input)
		got := p.parseAddr()
		if got != tc.want {
			t.Errorf("parseAddr(%q) = 0x%X, want 0x%X", tc.input, got, tc.want)
		}
	}
}

func TestBreaksPage_ParseTag(t *testing.T) {
	p := NewBreaks(nil)

	cases := []struct {
		input string
		want  uint64
	}{
		{"0x1A", 0x1A},
		{"1A", 0x1A},
		{"0", 0},
		{"1234", 0x1234},
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
