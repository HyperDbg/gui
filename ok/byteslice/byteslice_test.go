package byteslice_test

import (
	"testing"
	"unsafe"

	"github.com/ddkwork/golibrary/byteslice"
)

func int8Slice(s string) []int8 {
	a := make([]int8, len(s))
	for i := range s {
		a[i] = int8(s[i])
	}
	return a
}

func TestSliceToString(t *testing.T) {
	tests := []struct {
		name string
		in   any
		want string
	}{
		{"byte hello", []byte("hello"), "hello"},
		{"byte nul_truncate", []byte("hello\x00world"), "hello"},
		{"byte empty", []byte(nil), ""},
		{"byte nul_only", []byte{0}, ""},
		{"int8 hello", int8Slice("hello"), "hello"},
		{"int8 nul_truncate", int8Slice("hello\x00world"), "hello"},
		{"int8 empty", []int8(nil), ""},
		{"int8 nul_only", []int8{0}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got string
			switch v := tt.in.(type) {
			case []byte:
				got = byteslice.ToString(v)
			case []int8:
				got = byteslice.ToString(v)
			}
			if got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}
}

func TestPtrToString(t *testing.T) {
	bHello := []byte("hello\x00")
	bNul := []byte{0}
	iHello := int8Slice("hello\x00")
	iNul := []int8{0}

	tests := []struct {
		name string
		in   any
		want string
	}{
		{"byte hello", &bHello[0], "hello"},
		{"byte nil", (*byte)(nil), ""},
		{"byte nul", &bNul[0], ""},
		{"int8 hello", &iHello[0], "hello"},
		{"int8 nil", (*int8)(nil), ""},
		{"int8 nul", &iNul[0], ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got string
			switch v := tt.in.(type) {
			case *byte:
				got = byteslice.PtrToString(v)
			case *int8:
				got = byteslice.PtrToString(v)
			}
			if got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}
}

func TestSliceFromString(t *testing.T) {
	t.Run("byte ok", func(t *testing.T) {
		got := byteslice.FromString[byte]("hello")
		want := []byte("hello\x00")
		if string(got) != string(want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("byte empty", func(t *testing.T) {
		got := byteslice.FromString[byte]("")
		if len(got) != 1 || got[0] != 0 {
			t.Errorf("got %v, want [0]", got)
		}
	})

	t.Run("int8 ok", func(t *testing.T) {
		got := byteslice.FromString[int8]("hello")
		want := int8Slice("hello\x00")
		for i := range got {
			if got[i] != want[i] {
				t.Fatalf("got[%d]=%d, want[%d]=%d", i, got[i], i, want[i])
			}
		}
	})
}

func TestPtrFromString(t *testing.T) {
	t.Run("byte roundtrip", func(t *testing.T) {
		p := byteslice.PtrFromString[byte]("hello")
		if got := byteslice.PtrToString(p); got != "hello" {
			t.Errorf("got %q, want %q", got, "hello")
		}
	})

	t.Run("int8 roundtrip", func(t *testing.T) {
		p := byteslice.PtrFromString[int8]("hello")
		if got := byteslice.PtrToString(p); got != "hello" {
			t.Errorf("got %q, want %q", got, "hello")
		}
	})
}

func TestRoundTrip(t *testing.T) {
	s := "Hello, 世界! 🌍"
	for _, name := range []string{"byte", "int8"} {
		t.Run(name, func(t *testing.T) {
			switch name {
			case "byte":
				slice := byteslice.FromString[byte](s)
				if got := byteslice.ToString(slice); got != s {
					t.Errorf("slice roundtrip got %q, want %q", got, s)
				}
				p := byteslice.PtrFromString[byte](s)
				if got := byteslice.PtrToString(p); got != s {
					t.Errorf("ptr roundtrip got %q, want %q", got, s)
				}
			case "int8":
				slice := byteslice.FromString[int8](s)
				if got := byteslice.ToString(slice); got != s {
					t.Errorf("slice roundtrip got %q, want %q", got, s)
				}
				p := byteslice.PtrFromString[int8](s)
				if got := byteslice.PtrToString(p); got != s {
					t.Errorf("ptr roundtrip got %q, want %q", got, s)
				}
			}
		})
	}
}

func TestFromStruct(t *testing.T) {
	type point struct{ X, Y uint32 }
	p := point{X: 0x01020304, Y: 0x05060708}

	b := byteslice.FromStruct(&p)
	if len(b) != int(unsafe.Sizeof(p)) {
		t.Fatalf("got len %d, want %d", len(b), unsafe.Sizeof(p))
	}

	var p2 point
	p2 = *byteslice.ToStruct[point](b)
	if p2 != p {
		t.Errorf("got %v, want %v", p2, p)
	}
}

func TestToStruct(t *testing.T) {
	type header struct {
		Magic   uint32
		Version uint16
		Flags   uint16
	}

	t.Run("valid", func(t *testing.T) {
		h := header{Magic: 0x4D454D49, Version: 1, Flags: 0xFF}
		b := byteslice.FromStruct(&h)
		got := byteslice.ToStruct[header](b)
		if got.Magic != h.Magic || got.Version != h.Version || got.Flags != h.Flags {
			t.Errorf("got %+v, want %+v", *got, h)
		}
	})

	t.Run("too small", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("expected panic")
			}
		}()
		b := make([]byte, 1)
		byteslice.ToStruct[header](b)
	})

	t.Run("nil", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("expected panic")
			}
		}()
		byteslice.ToStruct[header](nil)
	})
}

func TestFromSlice(t *testing.T) {
	t.Run("uint32", func(t *testing.T) {
		s := []uint32{0x01020304, 0x05060708}
		b := byteslice.FromAnySlice(s)
		if len(b) != len(s)*int(unsafe.Sizeof(s[0])) {
			t.Fatalf("got len %d, want %d", len(b), len(s)*int(unsafe.Sizeof(s[0])))
		}
		if len(b) != 8 {
			t.Fatalf("got len %d, want 8", len(b))
		}
	})

	t.Run("uint16", func(t *testing.T) {
		s := []uint16{0x0102, 0x0304}
		b := byteslice.FromAnySlice(s)
		if len(b) != 4 {
			t.Fatalf("got len %d, want 4", len(b))
		}
	})

	t.Run("empty", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Error("expected panic")
			}
		}()
		byteslice.FromAnySlice[uint32](nil)
	})

	t.Run("struct slice", func(t *testing.T) {
		type pair struct{ A, B uint16 }
		s := []pair{{1, 2}, {3, 4}}
		b := byteslice.FromAnySlice(s)
		if len(b) != len(s)*int(unsafe.Sizeof(pair{})) {
			t.Fatalf("got len %d, want %d", len(b), len(s)*int(unsafe.Sizeof(pair{})))
		}
	})
}

func TestPtrFromSlice(t *testing.T) {
	b := []byte{'h', 'e', 'l', 'l', 'o'}
	p := byteslice.PtrFromAnySlice[byte](b)
	if p == nil {
		t.Fatal("nil pointer")
	}
	if *p != 'h' {
		t.Fatalf("expected 'h', got %c", *p)
	}
	*p = 'H'
	if b[0] != 'H' {
		t.Fatal("not same underlying memory")
	}

	i8 := []int8{1, 2, 3}
	q := byteslice.PtrFromAnySlice[int8](i8)
	if *q != 1 {
		t.Fatalf("expected 1, got %d", *q)
	}
	*q = -1
	if i8[0] != -1 {
		t.Fatal("int8 not same memory")
	}
}
