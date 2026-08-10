package comm

import (
	"context"
	"testing"
)

// TestCTL_CODE verifies the runtime CTL_CODE function matches the Windows SDK
// macro: (DeviceType << 16) | (Access << 14) | (Function << 2) | Method.
func TestCTL_CODE(t *testing.T) {
	cases := []struct {
		deviceType, function, method, access uint32
		want                                 uint32
	}{
		// FILE_DEVICE_UNKNOWN=0x22, METHOD_BUFFERED=0, FILE_ANY_ACCESS=0
		// IOCTL_INIT_VMM: function = 0x800 + 0x01 = 0x801
		{0x22, 0x801, 0, 0, 0x222004},
		// IOCTL_TERMINATE_VMX: function = 0x800 + 0x200 + 0x01 = 0xA01
		{0x22, 0xA01, 0, 0, 0x222804},
		// A non-HyperDbg sanity check: FILE_DEVICE_DISK=0x07, function=0x100,
		// METHOD_BUFFERED, FILE_READ_ACCESS=1
		{0x07, 0x100, 0, 1, (0x07 << 16) | (1 << 14) | (0x100 << 2)},
	}
	for _, c := range cases {
		got := CTL_CODE(c.deviceType, c.function, c.method, c.access)
		if got != c.want {
			t.Errorf("CTL_CODE(%#x, %#x, %#x, %#x) = %#x, want %#x",
				c.deviceType, c.function, c.method, c.access, got, c.want)
		}
	}
}

// TestIoctlConstants verifies the pre-computed IOCTL_CODE_* constants match
// the expected values derived from Ioctls.h. Each expected value is computed
// as 0x220000 | (function << 2) where function is the IOCTL_START_CODE base
// plus the documented offset.
func TestIoctlConstants(t *testing.T) {
	cases := []struct {
		name   string
		got    uint32
		offset uint32 // function offset from IOCTL_START_CODE (0x800)
	}{
		// Basic group (base = 0x800)
		{"IOCTL_CODE_INIT_VMM", IOCTL_CODE_INIT_VMM, 0x800 + 0x01},
		{"IOCTL_CODE_INIT_HYPERTRACE", IOCTL_CODE_INIT_HYPERTRACE, 0x800 + 0x02},
		{"IOCTL_CODE_REGISTER_EVENT", IOCTL_CODE_REGISTER_EVENT, 0x800 + 0x03},

		// VMM group (base = 0xA00)
		{"IOCTL_CODE_TERMINATE_VMX", IOCTL_CODE_TERMINATE_VMX, 0xA00 + 0x01},
		{"IOCTL_CODE_DEBUGGER_READ_MEMORY", IOCTL_CODE_DEBUGGER_READ_MEMORY, 0xA00 + 0x02},
		{"IOCTL_CODE_DEBUGGER_REGISTER_EVENT", IOCTL_CODE_DEBUGGER_REGISTER_EVENT, 0xA00 + 0x05},
		{"IOCTL_CODE_DEBUGGER_ADD_ACTION_TO_EVENT", IOCTL_CODE_DEBUGGER_ADD_ACTION_TO_EVENT, 0xA00 + 0x06},
		{"IOCTL_CODE_DEBUGGER_MODIFY_EVENTS", IOCTL_CODE_DEBUGGER_MODIFY_EVENTS, 0xA00 + 0x0b},
		{"IOCTL_CODE_DEBUGGER_FLUSH_LOGGING_BUFFERS", IOCTL_CODE_DEBUGGER_FLUSH_LOGGING_BUFFERS, 0xA00 + 0x0c},
		{"IOCTL_CODE_PREPARE_DEBUGGEE", IOCTL_CODE_PREPARE_DEBUGGEE, 0xA00 + 0x0f},
		{"IOCTL_CODE_PAUSE_PACKET_RECEIVED", IOCTL_CODE_PAUSE_PACKET_RECEIVED, 0xA00 + 0x10},
		{"IOCTL_CODE_DEBUGGER_CPUID", IOCTL_CODE_DEBUGGER_CPUID, 0xA00 + 0x27},

		// HyperTrace group (base = 0xB00)
		{"IOCTL_CODE_PERFORM_HYPERTRACE_UNLOAD", IOCTL_CODE_PERFORM_HYPERTRACE_UNLOAD, 0xB00 + 0x01},
	}
	for _, c := range cases {
		want := uint32(0x220000) | (c.offset << 2)
		if c.got != want {
			t.Errorf("%s = %#x, want %#x", c.name, c.got, want)
		}
	}
}

// TestDevicePaths verifies the device/pipe name constants are non-empty and
// use the expected Win32 path prefix.
func TestDevicePaths(t *testing.T) {
	if DeviceName != `\\.\HyperDbgDebuggerDevice` {
		t.Errorf("DeviceName = %q, want %q", DeviceName, `\\.\HyperDbgDebuggerDevice`)
	}
	for _, p := range []string{HyperDbgPipe, HyperDbgOutputPipe, HyperDbgTestsPipe} {
		if p == "" {
			t.Error("pipe path is empty")
		}
	}
}

// TestOpenDevice_NoDriver verifies that opening the device fails (not panics)
// when the VMM driver is not loaded. This needs no admin rights — a missing
// device simply returns an error from CreateFile.
func TestOpenDevice_NoDriver(t *testing.T) {
	_, err := OpenDefault(context.Background())
	if err == nil {
		// The device exists, meaning the driver is loaded. That's fine — the
		// test only asserts Open returns an error OR succeeds without panic.
		t.Log("device opened successfully; driver appears to be loaded")
	}
	// We don't fail on success because the driver might be loaded in a dev
	// environment. The key invariant: no panic, and an error is *expected*
	// when the driver is absent.
}

// TestIoctlRoundTrip is the Phase A acceptance integration test for the comm
// layer. It requires the VMM driver to be loaded (device present) and admin
// rights.
//
// Flow: open device → send IOCTL_DEBUGGER_FLUSH_LOGGING_BUFFERS (a benign
// IOCTL that flushes kernel logs; it returns success even with no active
// debugging session) → close.
//
// To run: load hyperhv.sys, then `go test -run TestIoctlRoundTrip` from an
// elevated shell.
func TestIoctlRoundTrip(t *testing.T) {
	ctx := context.Background()
	dev, err := OpenDefault(ctx)
	if err != nil {
		t.Skipf("device not available (driver not loaded?): %v", err)
	}
	t.Cleanup(func() { _ = dev.Close() })

	// IOCTL_DEBUGGER_FLUSH_LOGGING_BUFFERS takes no input/output payload in
	// the common case; a zero-length call should still return without error.
	n, err := dev.Ioctl(ctx, IOCTL_CODE_DEBUGGER_FLUSH_LOGGING_BUFFERS, nil, nil)
	if err != nil {
		t.Fatalf("FlushLoggingBuffers IOCTL failed: %v", err)
	}
	t.Logf("FlushLoggingBuffers returned %d bytes", n)
}
