package gengort

import (
	"fmt"
	"reflect"
)

func Validate(ptrToStruct any, size, align uintptr, fields ...any) {
	rtype := reflect.TypeOf(ptrToStruct).Elem()

	// Validate size
	if size != rtype.Size() {
		panic(fmt.Sprintf("Mismatching sizof(%s) 0x%d, expected 0x%x", rtype.Name(), rtype.Size(), size))
	}

	// Validate alignment
	if align != uintptr(rtype.Align()) {
		panic(fmt.Sprintf("Mismatching alignof(%s) 0x%d, expected 0x%x", rtype.Name(), rtype.Align(), align))
	}

	// Validate fields
	for i := 0; i < len(fields); i += 2 {
		fieldName := fields[i].(string)
		fieldOffset := reflect.ValueOf(fields[i+1]).Int()
		field, ok := rtype.FieldByName(fieldName)
		if !ok {
			panic(fmt.Sprintf("Field %s not found", fieldName))
		}
		if field.Offset != uintptr(fieldOffset) {
			panic(fmt.Sprintf("Mismatching offsetof(%s::%s): 0x%x, expected 0x%x", rtype.Name(), fieldName, field.Offset, fieldOffset))
		}
	}
}
