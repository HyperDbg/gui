package misc

import "testing"

// TestCallstack_Rel32 verifies that a 0xE8 (CALL rel32) preceding the
// return address is detected as a 5-byte call. The E8 opcode sits 5 bytes
// before the return address, i.e. at b[n-5] in the 7-byte prefix.
func TestCallstack_Rel32(t *testing.T) {
	// Layout: b[0..6] are the 7 bytes immediately before the return address.
	// For E8 cd (5-byte), b[n-5]=b[2]=0xE8, b[3..6] = 4 displacement bytes.
	pre := []byte{0xAA, 0xAA, 0xE8, 0x00, 0x00, 0x00, 0x00}
	n, ok := CallstackReturnAddressToCallingAddress(pre)
	if !ok {
		t.Fatal("expected match for E8 cd")
	}
	if n != 5 {
		t.Errorf("got len=%d, want 5", n)
	}
	caller, ok := CallerSite(0x401000, pre)
	if !ok {
		t.Fatal("CallerSite returned false")
	}
	if caller != 0x401000-5 {
		t.Errorf("caller=0x%x, want 0x%x", caller, 0x401000-5)
	}
}

// TestCallstack_FarCall verifies 0x9A (far call ptr16:32, 7-byte).
func TestCallstack_FarCall(t *testing.T) {
	// 9A at b[n-7]=b[0].
	pre := []byte{0x9A, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	n, ok := CallstackReturnAddressToCallingAddress(pre)
	if !ok || n != 7 {
		t.Errorf("far call: ok=%v n=%d, want ok=true n=7", ok, n)
	}
}

// TestCallstack_Indirect6 verifies FF /2 with ModR/M=0x15 (6-byte form).
// Layout: FF at b[n-6]=b[1], ModR/M at b[n-5]=b[2].
func TestCallstack_Indirect6(t *testing.T) {
	pre := []byte{0xAA, 0xFF, 0x15, 0xA0, 0xA5, 0x48, 0x76}
	n, ok := CallstackReturnAddressToCallingAddress(pre)
	if !ok || n != 6 {
		t.Errorf("indirect6: ok=%v n=%d, want ok=true n=6", ok, n)
	}
}

// TestCallstack_Indirect2 verifies FF /2 with ModR/M=0xD0 (2-byte form, call
// rax). Layout: FF at b[n-2]=b[5], ModR/M at b[n-1]=b[6].
func TestCallstack_Indirect2(t *testing.T) {
	pre := []byte{0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xFF, 0xD0}
	n, ok := CallstackReturnAddressToCallingAddress(pre)
	if !ok || n != 2 {
		t.Errorf("indirect2: ok=%v n=%d, want ok=true n=2", ok, n)
	}
}

// TestCallstack_NoMatch verifies that random bytes return false.
func TestCallstack_NoMatch(t *testing.T) {
	pre := []byte{0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA, 0xAA}
	if _, ok := CallstackReturnAddressToCallingAddress(pre); ok {
		t.Error("expected no match for all-0xAA prefix")
	}
}

// TestCallstack_ShortBuffer verifies the buffer-length guard.
func TestCallstack_ShortBuffer(t *testing.T) {
	pre := []byte{0xE8, 0x00, 0x00}
	if _, ok := CallstackReturnAddressToCallingAddress(pre); ok {
		t.Error("expected false for buffer < 7 bytes")
	}
}
