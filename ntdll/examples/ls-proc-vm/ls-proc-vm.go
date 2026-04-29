package main

import (
	"flag"
	"fmt"
	"os"
	"unsafe"

	"github.com/ddkwork/HyperDbg/ntdll"
)

func query_vm(ph ntdll.Handle, base uintptr) (*ntdll.MemoryBasicInformationT, ntdll.NtStatus) {
	var mbi ntdll.MemoryBasicInformationT
	s := ntdll.NtQueryVirtualMemory(
		ph,
		(*uint8)(unsafe.Pointer(base)),
		ntdll.MemoryBasicInformation,
		(*uint8)(unsafe.Pointer(&mbi)),
		unsafe.Sizeof(mbi),
		nil,
	)
	if s.IsError() {
		return nil, s
	}
	return &mbi, s
}

func query_name(ph ntdll.Handle, base uintptr) (string, ntdll.NtStatus) {
	buf := [1024]byte{}
	s := ntdll.NtQueryVirtualMemory(
		ph,
		(*uint8)(unsafe.Pointer(base)),
		ntdll.MemoryMappedFilenameInformation,
		&buf[0],
		1024,
		nil,
	)
	if s.IsError() {
		return "", s
	}
	return (*ntdll.UnicodeString)(unsafe.Pointer(&buf[0])).String(), s
}

type memProtect uint32

func (p memProtect) String() string {
	r := []byte("----")
	if p&ntdll.PAGE_NOACCESS != 0 {
		r[3] = '!'
	}
	if p&ntdll.PAGE_READONLY != 0 {
		r[0] = 'r'
	}
	if p&ntdll.PAGE_READWRITE != 0 {
		r[0] = 'r'
		r[1] = 'w'
	}
	if p&ntdll.PAGE_WRITECOPY != 0 {
		r[1] = 'w'
		r[3] = 'c'
	}
	if p&ntdll.PAGE_EXECUTE != 0 {
		r[2] = 'x'
	}
	if p&ntdll.PAGE_EXECUTE_READ != 0 {
		r[0] = 'r'
		r[2] = 'x'
	}
	if p&ntdll.PAGE_EXECUTE_READWRITE != 0 {
		r[0] = 'r'
		r[1] = 'w'
		r[2] = 'x'
	}
	if p&ntdll.PAGE_EXECUTE_WRITECOPY != 0 {
		r[1] = 'w'
		r[2] = 'x'
		r[3] = 'c'
	}
	if p&ntdll.PAGE_GUARD != 0 {
		r[3] = 'G'
	}
	if p&ntdll.PAGE_NOCACHE != 0 {
		r[3] = 'C'
	}
	return string(r)
}

type memState uint32

func (s memState) String() string {
	r := []byte("---")
	if s&ntdll.MEM_COMMIT != 0 {
		r[0] = 'c'
	}
	if s&ntdll.MEM_RESERVE != 0 {
		r[1] = 'r'
	}
	if s&ntdll.MEM_FREE != 0 {
		r[2] = 'f'
	}
	return string(r)
}

type memType uint32

func (s memType) String() string {
	r := []byte("---")
	if s&ntdll.MEM_MAPPED != 0 {
		r[0] = 'm'
	}
	if s&ntdll.MEM_PRIVATE != 0 {
		r[1] = 'p'
	}
	if s&ntdll.MEM_IMAGE != 0 {
		r[2] = 'i'
	}
	return string(r)
}

func main() {
	var pid uint
	flag.UintVar(&pid, "pid", 0, "target process id")
	flag.Parse()
	if pid == 0 {
		fmt.Printf("usage: %s -pid <pid>\n", os.Args[0])
		os.Exit(1)
	}

	var ph ntdll.Handle
	oa := ntdll.NewObjectAttributes("", 0, 0, nil)
	client := ntdll.ClientId{ntdll.Handle(pid), 0}
	s := ntdll.NtOpenProcess(&ph, ntdll.PROCESS_QUERY_INFORMATION, oa, &client)
	if s.IsError() {
		fmt.Printf("NtOpenProcess: %s", s)
		os.Exit(1)
	}

	addr := uintptr(0)
	for {
		mbi, s := query_vm(ph, addr)
		if s.IsError() {
			if s != ntdll.STATUS_INVALID_PARAMETER {
				fmt.Printf("NtQueryVirtualMemory: %08x: %s", addr, s.Error())
			}
			break
		}
		name, _ := query_name(ph, addr)
		if mbi.AllocationBase != nil {
			fmt.Printf("%12x %12x + %12x ty:%s st:%s pr:%s %s\n",
				uintptr(unsafe.Pointer(mbi.AllocationBase)),
				uintptr(unsafe.Pointer(mbi.BaseAddress)),
				mbi.RegionSize,
				memType(mbi.Type),
				memState(mbi.State),
				memProtect(mbi.Protect),
				name,
			)
		}
		addr = uintptr(unsafe.Pointer(mbi.BaseAddress)) + mbi.RegionSize
	}
}
