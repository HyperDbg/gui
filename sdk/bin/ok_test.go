package bin

import (
	"fmt"
	"github.com/ddkwork/HyperDbg/sdk"
	"syscall"
	"testing"
	"unsafe"
)

func TestName(t *testing.T) {
	lib, err := syscall.LoadDLL("libhyperdbg.dll")
	if err != nil {
		fmt.Printf("Failed to load libhyperdbg.dll: %v\n", err)
		return
	}

	procSetTextMessageCallback, err := lib.FindProc("hyperdbg_u_set_text_message_callback")
	if err != nil {
		fmt.Printf("Failed to find hyperdbg_u_set_text_message_callback: %v\n", err)
		lib.Release()
		return
	}

	procInterpreter, err := lib.FindProc("hyperdbg_u_interpreter")
	if err != nil {
		fmt.Printf("Failed to find hyperdbg_u_interpreter: %v\n", err)
		lib.Release()
		return
	}

	callback := syscall.NewCallback(func(text *byte) int { //bug in bindgen library
		fmt.Printf("Test in the handler | ")
		fmt.Println("Received data:", sdk.BytePointerToString(text))
		return 0
	})

	_, _, _ = procSetTextMessageCallback.Call(callback)

	text := append([]byte("help !monitor"), 0) // 确保字符串以NULL结尾
	_, _, _ = procInterpreter.Call(uintptr(unsafe.Pointer(&text[0])))
	lib.Release()
}
