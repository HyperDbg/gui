package main

import (
	"fmt"
	"io"
	"net"
	"time"

	"github.com/ddkwork/HyperDbg/debugger/driver"
)

func main() {
	fmt.Println("=== Step 1: Load Driver ===")
	p := driver.New("D:\\ux\\examples\\hypedbg\\rust-driver\\examples\\netdemo\\netdemo.sys", "netDemo", "\\\\.\\netDemo")

	p.Install()
	p.Start()

	fmt.Println("\n=== Step 2: TCP Client Test ===")
	addr := "127.0.0.1:9527"
	fmt.Printf("Connecting to %s...\n", addr)

	conn, err := net.DialTimeout("tcp", addr, 5*time.Second)
	if err != nil {
		fmt.Printf("Dial failed: %v\n", err)
	} else {
		defer conn.Close()
		fmt.Println("Connected!")

		msg := []byte("Hello from Go TCP client!")
		_, err = conn.Write(msg)
		if err != nil {
			fmt.Printf("Write failed: %v\n", err)
		} else {
			fmt.Printf("Sent: %s\n", msg)
		}

		buf := make([]byte, 1024)
		conn.SetReadDeadline(time.Now().Add(5 * time.Second))
		n, err := conn.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Printf("Read failed: %v\n", err)
		} else {
			fmt.Printf("Received: %s\n", buf[:n])
		}
	}

	fmt.Println("\n=== Step 3: Unload Driver ===")
	p.Stop()
	p.Uninstall()

	fmt.Println("\n=== Test Complete ===")
}
