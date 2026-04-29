package main

import (
	"bytes"
	"fmt"
	"os"
	"unsafe"

	"github.com/ddkwork/HyperDbg/ntdll"
)

func main() {
	buf := make([]byte, 1024)
	var rlen uint32
	if st := ntdll.CallWithExpandingBuffer(func() ntdll.NtStatus {
		return ntdll.NtQuerySystemInformation(
			ntdll.SystemModuleInformation,
			&buf[0],
			uint32(len(buf)),
			&rlen,
		)
	}, &buf, &rlen); st.IsError() {
		fmt.Fprintf(os.Stderr, "NtQuerySystemInformation: %s, len=%d", st.Error(), rlen)
		os.Exit(1)
	}
	mi := (*ntdll.SystemModuleInformationT)(unsafe.Pointer(&buf[0]))
	fmt.Println("idx map_base img_base         img_size flags    load init cnt  name")
	fmt.Println("--- -------- ---------------- -------- -------- ---- ---- ---- ----")
	for i, module := range mi.GetModules() {
		nlen := bytes.Index(module.Name[:], []byte{0})
		name := string(module.Name[:nlen])
		fmt.Printf("%3d %08x %016x %08x %08x %04x %04x %04x %s\n", i,
			module.MappedBase, module.ImageBase, module.ImageSize,
			module.Flags,
			module.LoadOrderIndex, module.InitOrderIndex, module.LoadCount,
			name)
	}
}
