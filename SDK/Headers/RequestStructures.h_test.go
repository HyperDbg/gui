package Headers

import (
	"reflect"
	"testing"
)

type (
	o struct {
		VirtualAddress      uint64
		ProcessId           uint32
		Pml4eVirtualAddress uint64
		Pml4eValue          uint64
		PdpteVirtualAddress uint64
		PdpteValue          uint64
		PdeVirtualAddress   uint64
		PdeValue            uint64
		PteVirtualAddress   uint64
		PteValue            uint64
		KernelStatus        uint32
	}
	//PDEBUGGER_READ_PAGE_TABLE_ENTRIES_DETAILS_ *DEBUGGER_READ_PAGE_TABLE_ENTRIES_DETAILS
)

func TestName(t *testing.T) {
	println(SizeOf2(o{}))
	//println(SIZEOF_DEBUGGER_READ_PAGE_TABLE_ENTRIES_DETAILS)
}
func SizeOf2(object any) (size int) {
	fields := reflect.VisibleFields(reflect.TypeOf(object))
	for _, field := range fields {
		switch field.Type.Kind() {
		case reflect.Uint8:
			size += 1
		case reflect.Uint16:
			size += 2
		case reflect.Uint32:
			size += 4
		case reflect.Uint64:
			size += 8
		}
	}
	return
}
