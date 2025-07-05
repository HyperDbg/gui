package bindlib

import (
	"fmt"
	"reflect"

	"github.com/ddkwork/golibrary/std/mylog"
)

func Validate(ptrToStruct any, size, align uintptr, fields ...any) {
	return
	rtype := reflect.TypeOf(ptrToStruct).Elem()

	// Validate size
	if size != rtype.Size() {
		mylog.CheckIgnore(fmt.Sprintf("Mismatching sizof(%s) %d, expected %d", rtype.Name(), rtype.Size(), size))
	}

	// Validate alignment
	if align != uintptr(rtype.Align()) {
		mylog.CheckIgnore(fmt.Sprintf("Mismatching alignof(%s) %d, expected %d", rtype.Name(), rtype.Align(), align))
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
			mylog.CheckIgnore(fmt.Sprintf("Mismatching offsetof(%s::%s): %d, expected %d", rtype.Name(), fieldName, field.Offset, fieldOffset))
		}
	}
}
