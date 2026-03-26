package main

import (
	"bufio"
	"fmt"
	"net"
	"time"

	"github.com/ddkwork/HyperDbg/debugger/driver"
)

func main() {
	fmt.Println("=== Step 1: Load Driver ===")
	p := driver.NewWithOptions(`D:\ux\examples\hypedbg\rust-driver\examples\netdemo\netdemo.sys`, "netdemo", "\\\\.\\netdemo", true)

	p.Install()
	p.Start()

	fmt.Println("\n=== Step 2: TCP Client Test (IPv6 Port 40007) ===")
	addr := "[::1]:40007"
	fmt.Printf("Connecting to %s...\n", addr)

	conn, err := net.DialTimeout("tcp", addr, 5*time.Second)
	if err != nil {
		fmt.Printf("Dial fai///、、、、、、、、、、、、、、、、、、、、、///、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、、led: %v\n", err)
		p.Stop()
		p.Uninstall()
		return
	}
	defer conn.Close()
	fmt.Println("Connected!")

	testMessages := []string{
		"Hello, WSK!",
		"PING",
		"Test message 123",
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ",
	}

	for _, msg := range testMessages {
		data := []byte(msg + "\n")
		_, err = conn.Write(data)
		if err != nil {
			fmt.Printf("Send failed: %v\n", err)
			continue
		}
		fmt.Printf("Sent: %s", msg)

		reader := bufio.NewReader(conn)
		respData, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Read failed: %v\n", err)
			continue
		}
		fmt.Printf("Echo: %s\n\n", respData[:len(respData)-1])
	}

	fmt.Println("\n=== Step 3: Unload Driver ===")
	p.Stop()
	p.Uninstall()

	fmt.Println("\n=== Test Complete ===")
}
