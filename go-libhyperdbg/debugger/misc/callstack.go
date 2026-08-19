// Package misc — callstack.go
//
// Implements CallstackReturnAddressToCallingAddress, the heuristic that walks
// back from a return address to find the calling CALL instruction. Used by
// the `k` (stack backtrace) command and by the break-control logic that
// decides whether a breakpoint hit is "user-set" or "stepping".
//
// Mirrors libhyperdbg/code/debugger/misc/callstack.cpp.
package misc

// CallstackReturnAddressToCallingAddress inspects the bytes immediately
// *before* ReturnAddress and, if they look like one of the x86 CALL encodings,
// returns the call instruction length (so the caller can subtract it from
// ReturnAddress to find the call site). Returns (0, false) when no known CALL
// encoding matches.
//
// preReturnBytes must contain at least 7 bytes that precede the return
// address (i.e. preReturnBytes[len-1] is the byte immediately before the
// return address). The C version indexes as ReturnAddress[-N]; here that
// becomes b[len-N].
func CallstackReturnAddressToCallingAddress(preReturnBytes []byte) (callLen uint32, ok bool) {
	n := len(preReturnBytes)
	if n < 7 {
		return 0, false
	}
	b := preReturnBytes
	const rmMask = 0xF8

	// 9A cp - CALL ptr16:32 (7-byte)
	if b[n-7] == 0x9A {
		return 7, true
	}
	// E8 cd - CALL rel32 (5-byte)
	if b[n-5] == 0xE8 {
		return 5, true
	}

	// FF /2 — near indirect call. The ModR/M byte position depends on the
	// total instruction length; check every form exactly as the C does.

	// 7-byte: FF [ModR/M=0x94|0x9C] [SIB] [4-byte disp]
	if b[n-7] == 0xFF && (b[n-6] == 0x94 || b[n-6] == 0x9C) {
		return 7, true
	}

	// 6-byte (a): FF [ModR/M] [4-byte disp]
	//   ModR/M mask 0xF8 in {0x90, 0x98} but not 0x94/0x9C
	if b[n-6] == 0xFF &&
		((b[n-5]&rmMask) == 0x90 || (b[n-5]&rmMask) == 0x98) &&
		(b[n-5] != 0x94 && b[n-5] != 0x9C) {
		return 6, true
	}
	// 6-byte (b): FF [ModR/M=0x15|0x1D] [4-byte disp]
	if b[n-6] == 0xFF && (b[n-5] == 0x15 || b[n-5] == 0x1D) {
		return 6, true
	}

	// 4-byte: FF [ModR/M=0x54|0x5C] [SIB] [1-byte disp]
	if b[n-4] == 0xFF && (b[n-3] == 0x54 || b[n-3] == 0x5C) {
		return 4, true
	}

	// 3-byte (a): FF [ModR/M] [1-byte disp]
	//   ModR/M mask 0xF8 in {0x50, 0x58} but not 0x54/0x5C
	if b[n-3] == 0xFF &&
		((b[n-2]&rmMask) == 0x50 || (b[n-2]&rmMask) == 0x58) &&
		(b[n-2] != 0x54 && b[n-2] != 0x5C) {
		return 3, true
	}
	// 3-byte (b): FF [ModR/M=0x14|0x1C] [SIB]
	if b[n-3] == 0xFF && (b[n-2] == 0x14 || b[n-2] == 0x1C) {
		return 3, true
	}

	// 2-byte (a): FF [ModR/M]
	//   ModR/M mask 0xF8 in {0xD0, 0xD8}
	if b[n-2] == 0xFF &&
		((b[n-1]&rmMask) == 0xD0 || (b[n-1]&rmMask) == 0xD8) {
		return 2, true
	}
	// 2-byte (b): FF [ModR/M]
	//   ModR/M mask 0xF8 in {0x10, 0x18} but not in {0x14, 0x15, 0x1C, 0x1D}
	if b[n-2] == 0xFF &&
		((b[n-1]&rmMask) == 0x10 || (b[n-1]&rmMask) == 0x18) &&
		(b[n-1] != 0x14 && b[n-1] != 0x15 && b[n-1] != 0x1C && b[n-1] != 0x1D) {
		return 2, true
	}

	return 0, false
}

// CallerSite returns the address of the CALL instruction that produced the
// given return address, or (0, false) if no recognised CALL pattern appears
// in the bytes preceding it. preReturnBytes must be the 7 bytes immediately
// before the return address.
func CallerSite(returnAddress uint64, preReturnBytes []byte) (uint64, bool) {
	n, ok := CallstackReturnAddressToCallingAddress(preReturnBytes)
	if !ok {
		return 0, false
	}
	return returnAddress - uint64(n), true
}
