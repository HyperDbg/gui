package main

import (
	"fmt"
	"os"
	"syscall"
	"unsafe"
)

var (
	ole32            = syscall.NewLazyDLL("ole32.dll")
	oleaut32         = syscall.NewLazyDLL("oleaut32.dll")
	coInitialize     = ole32.NewProc("CoInitialize")
	coUninitialize   = ole32.NewProc("CoUninitialize")
	coCreateInstance = ole32.NewProc("CoCreateInstance")
	sysAllocString   = oleaut32.NewProc("SysAllocString")
	sysFreeString    = oleaut32.NewProc("SysFreeString")
)

func main0() {
	fmt.Println("Starting DIA test...")

	ret, _, _ := coInitialize.Call(0)
	fmt.Printf("CoInitialize returned: %d\n", ret)

	clsidDiaSource := syscall.GUID{
		Data1: 0xe6756135,
		Data2: 0x1e65,
		Data3: 0x4d17,
		Data4: [8]byte{0x85, 0x76, 0x61, 0x07, 0x61, 0x39, 0x8c, 0x3c},
	}

	iidIDataSource := syscall.GUID{
		Data1: 0x79f1bb5f,
		Data2: 0xb66e,
		Data3: 0x48e5,
		Data4: [8]byte{0xb6, 0xa9, 0x15, 0x45, 0xc3, 0x23, 0xca, 0x3d},
	}

	var dataSource uintptr
	ret, _, _ = coCreateInstance.Call(
		uintptr(unsafe.Pointer(&clsidDiaSource)),
		0,
		1,
		uintptr(unsafe.Pointer(&iidIDataSource)),
		uintptr(unsafe.Pointer(&dataSource)),
	)
	fmt.Printf("CoCreateInstance returned: 0x%08X, dataSource: 0x%X\n", ret, dataSource)

	if dataSource != 0 {
		pdbPath := `D:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\build\Release\hyperkd.pdb`
		if _, err := os.Stat(pdbPath); os.IsNotExist(err) {
			fmt.Printf("PDB file not found: %s\n", pdbPath)
		} else {
			utf16Path, _ := syscall.UTF16PtrFromString(pdbPath)
			bstr, _, _ := sysAllocString.Call(uintptr(unsafe.Pointer(utf16Path)))
			fmt.Printf("BSTR: 0x%X\n", bstr)
			defer sysFreeString.Call(bstr)

			vtablePtr := *(*uintptr)(unsafe.Pointer(dataSource))
			vtable := (*[20]uintptr)(unsafe.Pointer(vtablePtr))

			loadDataFromPdb := vtable[4]
			fmt.Printf("Calling loadDataFromPdb at 0x%X\n", loadDataFromPdb)
			ret, _, _ = syscall.SyscallN(loadDataFromPdb, dataSource, bstr)
			fmt.Printf("loadDataFromPdb returned: 0x%08X\n", ret)

			if ret == 0 {
				fmt.Println("Success! Opening session...")

				var session uintptr
				ret, _, _ = syscall.SyscallN(vtable[8], dataSource, uintptr(unsafe.Pointer(&session)))
				fmt.Printf("openSession returned: 0x%08X, session: 0x%X\n", ret, session)

				if session != 0 {
					sessionVtablePtr := *(*uintptr)(unsafe.Pointer(session))
					sessionVtable := (*[20]uintptr)(unsafe.Pointer(sessionVtablePtr))

					var global uintptr
					ret, _, _ = syscall.SyscallN(sessionVtable[5], session, uintptr(unsafe.Pointer(&global)))
					fmt.Printf("get_globalScope returned: 0x%08X, global: 0x%X\n", ret, global)

					if global != 0 {
						globalVtablePtr := *(*uintptr)(unsafe.Pointer(global))
						globalVtable := (*[100]uintptr)(unsafe.Pointer(globalVtablePtr))

						fmt.Println("Getting symTag (index 4)...")
						getSymTag := globalVtable[4]
						var symTag uint32
						ret, _, _ = syscall.SyscallN(getSymTag, global, uintptr(unsafe.Pointer(&symTag)))
						fmt.Printf("get_symTag returned: 0x%08X, symTag: %d\n", ret, symTag)

						fmt.Println("Getting name (index 5)...")
						getName := globalVtable[5]
						var nameBstr uintptr
						ret, _, _ = syscall.SyscallN(getName, global, uintptr(unsafe.Pointer(&nameBstr)))
						fmt.Printf("get_name returned: 0x%08X, nameBstr: 0x%X\n", ret, nameBstr)
						if nameBstr != 0 {
							name := bstrToString(nameBstr)
							fmt.Printf("Name: %s\n", name)
							sysFreeString.Call(nameBstr)
						}

						fmt.Println("Getting machine type (index 94)...")
						getMachineType := globalVtable[94]
						var machineType uint32
						ret, _, _ = syscall.SyscallN(getMachineType, global, uintptr(unsafe.Pointer(&machineType)))
						fmt.Printf("get_machineType returned: 0x%08X, machineType: 0x%04X\n", ret, machineType)

						fmt.Println("Done with basic test!")

						// Release global
						globalRelease := globalVtable[2]
						syscall.SyscallN(globalRelease, global)
					}

					// Release session
					sessionRelease := sessionVtable[2]
					syscall.SyscallN(sessionRelease, session)
				}
			}

			// Release dataSource
			dataSourceRelease := vtable[2]
			syscall.SyscallN(dataSourceRelease, dataSource)
		}
	}

	coUninitialize.Call()
	fmt.Println("Done!")
}

func bstrToString(bstr uintptr) string {
	if bstr == 0 {
		return ""
	}

	ptr := (*uint16)(unsafe.Pointer(bstr))
	length := 0
	for p := ptr; *p != 0; p = (*uint16)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 2)) {
		length++
	}

	if length == 0 {
		return ""
	}

	buf := make([]uint16, length)
	for i := 0; i < length; i++ {
		buf[i] = *(*uint16)(unsafe.Pointer(uintptr(unsafe.Pointer(ptr)) + uintptr(i*2)))
	}

	return syscall.UTF16ToString(buf)
}
