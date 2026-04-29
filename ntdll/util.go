package ntdll

import (
	"reflect"
	"unsafe"
)

// CallWithExpandingBuffer is intended to more seamlessly dealing with
// API calls that fill a caller-provided buffer with data. Usually,
// these functions expect the address and current size of the buffer
// and a reference to which the required buffer size will be written.
//
// This function calls fn which encapsulates the actual API call using
// buf and resultLength. If fn returns an NtStatus value that
// indicates an insufficient buffer sizze, the buffer is resized and
// the call is retried. This happens until the call succeeds or fails
// for a non-buffer-related reason.
//
// The size of buf is adjusted according to resultLength.
//
// Example:
/*
var keyName string
var buf []byte
var requiredLength uint32
if st := CallWithExpandingBuffer(func() NtStatus {
	return NtQueryKey(h, KeyNameInformation, &buf[0], uint32(len(buf)), &requiredLength)
}, &buf, &requiredLength); st.IsError() {
	// handle non buffer-size-related error
} else {
	kbi := (*KeyNameInformationT)(unsafe.Pointer(&buf[0]))
	keyName = NewUnicodeStringFromBuffer(
		(*uint16)(unsafe.Pointer(&kbi.Name)),
		int(kbi.NameLength)).String()
}
*/
func CallWithExpandingBuffer(fn func() NtStatus, buf *[]byte, resultLength *uint32) NtStatus {
	if len(*buf) == 0 {
		*buf = make([]byte, 32, 32)
	}
	for {
		st := fn()
		if int(*resultLength) <= cap(*buf) {
			(*reflect.SliceHeader)(unsafe.Pointer(buf)).Len = int(*resultLength)
		} else {
			orig := buf
			*buf = make([]byte, int(*resultLength))
			copy(*buf, *orig)
		}
		if st == STATUS_BUFFER_OVERFLOW || st == STATUS_BUFFER_TOO_SMALL || st == STATUS_INFO_LENGTH_MISMATCH {
			continue
		} else {
			return st
		}
	}
}
