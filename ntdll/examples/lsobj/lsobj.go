// lsobj is a console-based WinObj-like tool
// See https://technet.microsoft.com/en-us/sysinternals/winobj
package main

import (
	"errors"
	"fmt"
	"os"
	"unsafe"

	"github.com/ddkwork/golibrary/std/mylog"

	"github.com/ddkwork/HyperDbg/ntdll"
)

var skipDir = errors.New("skip this directory")

type walkFunc func(string, string) error

func walk(entry string, fn walkFunc) error {
	var h ntdll.Handle
	if st := ntdll.NtOpenDirectoryObject(&h, ntdll.STANDARD_RIGHTS_READ|ntdll.DIRECTORY_QUERY,
		ntdll.NewObjectAttributes(entry, 0, 0, nil),
	); st != 0 {
		return st.Error()
	}
	defer ntdll.NtClose(h)
	var context uint32
	for {
		var buf [32768]byte
		var length uint32
		switch st := ntdll.NtQueryDirectoryObject(
			h,
			&buf[0],
			uint32(len(buf)),
			true,
			context == 0,
			&context,
			&length,
		); st {
		case ntdll.STATUS_SUCCESS:
		case ntdll.STATUS_NO_MORE_ENTRIES:
			return nil
		default:
			return st.Error()
		}
		odi := (*ntdll.ObjectDirectoryInformationT)(unsafe.Pointer(&buf[0]))
		var path string
		if entry == `\` {
			path = `\` + odi.Name.String()
		} else {
			path = entry + `\` + odi.Name.String()
		}
		switch typ := odi.TypeName.String(); typ {
		case "Directory":
			mylog.Check(walk(path, fn))
		default:
			switch e := fn(path, typ); e {
			case skipDir:
				return nil
			case nil:
				continue
			default:
				return e
			}
		}
	}
}

func main() {
	var arg string
	if len(os.Args) > 1 {
		arg = os.Args[1]
	} else {
		arg = "\\"
	}
	walk(arg, func(path, typ string) error {
		switch typ {
		case "SymbolicLink":
			var h ntdll.Handle
			if st := ntdll.NtOpenSymbolicLinkObject(
				&h,
				ntdll.STANDARD_RIGHTS_READ|ntdll.DIRECTORY_QUERY,
				ntdll.NewObjectAttributes(path, 0, 0, nil),
			); st != 0 {
				fmt.Printf("%s %s -> <NtOpenSymbolicLinkObject error %08x>\n", path, typ, st)
				return nil
			}
			defer ntdll.NtClose(h)
			target := ntdll.NewEmptyUnicodeString(1024)
			var length uint32
			if st := ntdll.NtQuerySymbolicLinkObject(
				h,
				target,
				&length,
			); st != 0 {
				fmt.Printf("%s %s -> <NtQuerySymbolicLinkObject error %08x>\n", path, typ, st)
				return nil
			}
			fmt.Printf("%s %s -> %s\n", path, typ, target)
		default:
			fmt.Printf("%s %s\n", path, typ)
		}
		return nil
	})
}
