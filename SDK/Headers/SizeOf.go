package Headers

import "reflect"

func SizeOf(object any) (size int) {
	fields := reflect.VisibleFields(reflect.TypeOf(object))
	for _, field := range fields {
		switch field.Type.Kind() {
		case reflect.Uint8:
			size += 1
		case reflect.Uint64:
			size += 8
			//todo case all type
		}
	}
	return
}
