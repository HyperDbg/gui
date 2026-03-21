//go:build windows

package cve

import (
	"errors"
	"iter"
	"unsafe"

	"github.com/ddkwork/HyperDbg/ntdll"
	"golang.org/x/sys/windows"
)

// var (
//	dll                      = windows.NewLazySystemDLL("ntdll.dll")
//	procNtWriteVirtualMemory = dll.NewProc("NtWriteVirtualMemory")
//	procNtReadVirtualMemory  = dll.NewProc("NtReadVirtualMemory")
//	procNtFsControlFile      = dll.NewProc("NtFsControlFile")
// )

func NtQuerySystemInformation[T any](sysInfoClass int32) iter.Seq[T] {
	type provider struct {
		Count    uint32
		Children [1]T
	}
	var infos []T
	for size := uint32(128 * 1024); ; {
		b := make([]byte, size)
		err := windows.NtQuerySystemInformation(sysInfoClass, unsafe.Pointer(&b[0]), size, &size)
		switch {
		case errors.Is(err, windows.STATUS_INFO_LENGTH_MISMATCH):
			continue
		case err == nil:
			break
		default:
			break
		}
		m := (*provider)(unsafe.Pointer(&b[0]))
		infos = unsafe.Slice(&m.Children[0], int(m.Count))
		break
	}
	return func(yield func(T) bool) {
		for _, info := range infos {
			if !yield(info) {
				return
			}
		}
	}
}

// func NtWriteVirtualMemory(Dst *uintptr, Src *uintptr, Size uintptr) {
//	var size uintptr
//	ret, _, _ := procNtWriteVirtualMemory.Call(uintptr(windows.CurrentProcess()), uintptr(unsafe.Pointer(Dst)), uintptr(unsafe.Pointer(Src)), Size, uintptr(unsafe.Pointer(&size)))
//	Status := windows.NTStatus(ret)
//	if !errors.Is(Status, windows.STATUS_SUCCESS) {
//		mylog.Check(Status)
//	}
// }
// func NtReadVirtualMemory(Src *uintptr, Dst *uintptr, Size uintptr) {
//	var size uintptr
//	ret, _, _ := procNtReadVirtualMemory.Call(uintptr(windows.CurrentProcess()), uintptr(unsafe.Pointer(Src)), uintptr(unsafe.Pointer(Dst)), Size, uintptr(unsafe.Pointer(&size)))
//	Status := windows.NTStatus(ret)
//	if !errors.Is(Status, windows.STATUS_SUCCESS) {
//		mylog.Check(Status)
//	}
// }

func InitializeObjectAttributes(path string) ntdll.ObjectAttributes {
	objectAttr := ntdll.ObjectAttributes{
		ObjectName: ntdll.NewUnicodeString(path),
	}
	objectAttr.Length = uint32(unsafe.Sizeof(objectAttr))
	return objectAttr
}
