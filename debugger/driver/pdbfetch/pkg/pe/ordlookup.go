package pe

import (
	"fmt"
	"strings"
)

var OrdNames = map[string]map[uint64]string{
	"ws2_32.dll":   Ws232OrdNames,
	"wsock32.dll":  Ws232OrdNames,
	"oleaut32.dll": Oleaut32OrdNames,
}

func OrdLookup(libname string, ord uint64, makeName bool) string {
	names, ok := OrdNames[strings.ToLower(libname)]
	if ok {
		if name, ok := names[ord]; ok {
			return name
		}
	}
	if makeName {
		return fmt.Sprintf("ord%d", ord)
	} else {
		return ""
	}
}
