package main

import (
	"fmt"
	"syscall"
)

func main() {
	// 加载DLL
	dll := syscall.NewLazyDLL("../mydll/mydll.dll")
	if dll == nil {
		fmt.Println("Failed to load DLL")
		return
	}

	// 获取导出的函数
	add := dll.NewProc("Add")
	if add == nil {
		fmt.Println("Failed to find Add function")
		return
	}

	// 调用函数
	a, b := 3, 5
	ret, _, _ := add.Call(uintptr(a), uintptr(b))
	fmt.Printf("Result: %d\n", ret)
}
