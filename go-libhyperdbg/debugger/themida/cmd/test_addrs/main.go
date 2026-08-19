package main

import (
	"fmt"
	"syscall"
	"unsafe"
)

var (
	kernel32   = syscall.NewLazyDLL("kernel32.dll")
	getModProc = kernel32.NewProc("GetModuleHandleW")
)

func main() {
	dlls := []string{"ntdll.dll", "kernel32.dll", "kernelbase.dll"}
	for _, dll := range dlls {
		basePtr, _, _ := getModProc.Call(uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(dll))))
		fmt.Printf("%s base: 0x%X\n", dll, uint64(basePtr))
	}

	// Check VirtualAlloc export - is it a forwarder?
	k32Base, _, _ := getModProc.Call(uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("kernel32.dll"))))
	kbBase, _, _ := getModProc.Call(uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("kernelbase.dll"))))

	// Read kernel32 export table to check VirtualAlloc
	k32VA := syscall.NewLazyDLL("kernel32.dll").NewProc("VirtualAlloc")
	kbVA := syscall.NewLazyDLL("kernelbase.dll").NewProc("VirtualAlloc")
	k32SE := syscall.NewLazyDLL("kernel32.dll").NewProc("SetEvent")
	kbSE := syscall.NewLazyDLL("kernelbase.dll").NewProc("SetEvent")

	fmt.Printf("kernel32!VirtualAlloc  = 0x%X (offset from base: 0x%X)\n", k32VA.Addr(), uint64(k32VA.Addr())-uint64(k32Base))
	fmt.Printf("kernelbase!VirtualAlloc = 0x%X (offset from base: 0x%X)\n", kbVA.Addr(), uint64(kbVA.Addr())-uint64(kbBase))
	fmt.Printf("kernel32!SetEvent  = 0x%X (offset from base: 0x%X)\n", k32SE.Addr(), uint64(k32SE.Addr())-uint64(k32Base))
	fmt.Printf("kernelbase!SetEvent = 0x%X (offset from base: 0x%X)\n", kbSE.Addr(), uint64(kbSE.Addr())-uint64(kbBase))

	// Try reading first bytes at these addresses
	for name, addr := range map[string]uintptr{"kernel32!VirtualAlloc": k32VA.Addr(), "kernelbase!VirtualAlloc": kbVA.Addr(), "kernel32!SetEvent": k32SE.Addr(), "kernelbase!SetEvent": kbSE.Addr()} {
		buf := (*[4]byte)(unsafe.Pointer(addr))
		fmt.Printf("%s at 0x%X: %02X %02X %02X %02X\n", name, addr, buf[0], buf[1], buf[2], buf[3])
	}
}
