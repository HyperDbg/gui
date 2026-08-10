package byteslice

import (
	"slices"
	"strings"
	"unsafe"

	"github.com/ddkwork/golibrary/std/mylog"
)

type Type interface {
	~uint8 | ~int8
}

func FromString[T Type](s string) []T {
	if strings.IndexByte(s, 0) != -1 {
		mylog.Check("字符串包含NUL字节")
	}
	a := make([]T, len(s)+1)
	copy(unsafe.Slice((*byte)(unsafe.Pointer(unsafe.SliceData(a))), len(s)+1), s)
	return a
}

func ToString[T Type](s []T) string {
	if i := slices.Index(s, T(0)); i != -1 {
		s = s[:i]
	}
	return unsafe.String((*byte)(unsafe.Pointer(unsafe.SliceData(s))), len(s))
}

func PtrToString[T Type](p *T) string {
	if p == nil {
		mylog.Warning("传入空指针")
		return ""
	}
	n := 0
	for ptr := unsafe.Pointer(p); *(*T)(ptr) != 0; n++ {
		ptr = unsafe.Pointer(uintptr(ptr) + 1)
	}
	return unsafe.String((*byte)(unsafe.Pointer(p)), n)
}

func PtrFromString[T Type](s string) *T {
	return &FromString[T](s)[0]
}

func FromStruct[T any](s *T) []byte {
	return unsafe.Slice((*byte)(unsafe.Pointer(s)), unsafe.Sizeof(*s))
}

func ToStruct[T any](b []byte) *T {
	if len(b) < int(unsafe.Sizeof(new(T))) {
		mylog.Check("结构体大小验证失败")
	}
	return (*T)(unsafe.Pointer(unsafe.SliceData(b)))
}

func FromAnySlice[T any](s []T) []byte {
	if len(s) == 0 {
		mylog.Check("空buffer")
	}
	return unsafe.Slice((*byte)(unsafe.Pointer(unsafe.SliceData(s))), len(s)*int(unsafe.Sizeof(s[0])))
}

func PtrFromAnySlice[T any](s []T) *T {
	if len(s) == 0 {
		mylog.Check("空buffer")
	}
	return unsafe.SliceData(s)
}
