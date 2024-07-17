package main

import "C"

//export WriteStringToFile
func WriteStringToFile(str *C.char) {
	goStr := C.GoString(str)
	// 这里可以写入文件，为了简单起见，我们只是打印出来
	println(goStr)
}

func main() {}
