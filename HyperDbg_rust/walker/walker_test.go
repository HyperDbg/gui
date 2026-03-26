package walker

import (
	"testing"

	"github.com/ddkwork/HyperDbg_rust/protocol"
)

func TestWalker(t *testing.T) {
	t.Run("IoctlCodeToString", func(t *testing.T) {
		ioctlCodes := map[uint32]string{
			0x806: "IOCTL_DEBUGGER_REGISTER_EVENT",
			0x807: "IOCTL_DEBUGGER_ADD_ACTION_TO_EVENT",
			0x808: "IOCTL_DEBUGGER_MODIFY_EVENTS",
		}
		for code, name := range ioctlCodes {
			t.Logf("0x%04X: %s", code, name)
		}
	})

	t.Run("SyscallNumberLookup", func(t *testing.T) {
		syscalls := map[uint32]string{
			0x00: "NtAllocateVirtualMemory",
			0x07: "NtDeviceIoControlFile",
			0x12: "NtAllocateVirtualMemory",
		}
		for num, name := range syscalls {
			t.Logf("0x%02X: %s", num, name)
		}
	})

	t.Run("IoControlCodeParsing", func(t *testing.T) {
		testCodes := []uint32{0x00220000, 0x80002000, 0x80002004}
		for _, code := range testCodes {
			t.Logf("0x%08X -> %s", code, IoControlCodeToString(code))
		}
	})

	t.Run("ProcessBehaviorLogger", func(t *testing.T) {
		targetPid := uint32(0)
		walker := NewProcessBehaviorLogger(targetPid)
		logger := walker.(*ProcessBehaviorLogger)
		t.Logf("Process Behavior Logger initialized")
		t.Logf("Target PID: %d (0 means all processes)", targetPid)

		event := &protocol.SyscallEvent{
			ProcessID:     1234,
			ThreadID:      5678,
			CoreID:        0,
			SyscallNumber: 0x12,
			RAX:           0x12,
			RCX:           0x1000,
			RDX:           0x2000,
			R8:            0x3000,
			R9:            0x4000,
			R10:           0x5000,
			RSP:           0x6000,
		}

		for i := 0; i < 10; i++ {
			logger.Maiter(event, func() {
				t.Logf("method:syscall, pid: %x, tid: %x, syscall: %x",
					event.ProcessID, event.ThreadID, event.RAX)
			})
		}

		t.Log(logger.Summary())
	})

	t.Run("MemoryAllocationMonitor", func(t *testing.T) {
		targetPid := uint32(0)
		walker := NewMemoryAllocationMonitor(targetPid)
		monitor := walker.(*MemoryAllocationMonitor)
		t.Logf("Memory Allocation Monitor initialized")

		event := &protocol.SyscallEvent{
			ProcessID:     1234,
			ThreadID:      5678,
			CoreID:        0,
			SyscallNumber: SyscallNtAllocateVirtualMemory,
			RAX:           SyscallNtAllocateVirtualMemory,
			RCX:           0x1000,
			RDX:           0x7ffe0000,
			R8:            0x1000,
			R9:            0x1000,
			R10:           0x5000,
			RSP:           0x6000,
			Options: protocol.DebugEventOptions{
				OptionalParam1: 0x1000,
			},
		}

		for i := 0; i < 5; i++ {
			monitor.Maiter(event, func() {
				t.Logf("[SYSCALL] NtAllocateVirtualMemory, pid: %x | size: %x",
					event.ProcessID, event.R9)
			})
		}

		t.Log(monitor.Summary())
	})

	t.Run("ComprehensiveMonitor", func(t *testing.T) {
		targetPid := uint32(0)
		walker := NewComprehensiveMonitor(targetPid)
		monitor := walker.(*ComprehensiveMonitor)
		t.Logf("Comprehensive Monitor initialized")

		syscalls := []uint64{0x12, 0x07, 0x18, 0x19, 0x1a}

		for i, syscallNum := range syscalls {
			event := &protocol.SyscallEvent{
				ProcessID:     1234,
				ThreadID:      5678,
				CoreID:        0,
				SyscallNumber: syscallNum,
				RAX:           syscallNum,
				RCX:           0x1000,
				RDX:           0x2000,
				R8:            0x3000,
				R9:            0x4000,
				R10:           0x5000,
				RSP:           0x6000,
			}
			monitor.Maiter(event, func() {
				t.Logf("syscall: %x, pid: %x", event.RAX, event.ProcessID)
			})
			if i%2 == 0 {
				monitor.Maiter(event, func() {})
			}
		}

		t.Log(monitor.Summary())
	})

	t.Run("DFIRMonitor", func(t *testing.T) {
		targetPid := uint32(0)
		walker := NewDFIRMonitor(targetPid)
		monitor := walker.(*DFIRMonitor)
		t.Logf("DFIR Monitor initialized")

		for i := 0; i < 10; i++ {
			event := &protocol.SyscallEvent{
				ProcessID:     1234,
				ThreadID:      5678,
				CoreID:        0,
				SyscallNumber: uint64(i),
				RAX:           uint64(i),
				RCX:           0x1000,
				RDX:           0x2000,
				R8:            0x3000,
				R9:            0x4000,
				R10:           0x5000,
				RSP:           0x6000,
			}
			monitor.Maiter(event, func() {
				t.Logf("method:syscall, pid: %x, tid: %x, syscall: %x",
					event.ProcessID, event.ThreadID, event.RAX)
			})
		}

		t.Log(monitor.Summary())
	})

	t.Run("NtDeviceIoControlFileMonitor", func(t *testing.T) {
		t.Logf("Starting NtDeviceIoControlFile monitor")

		ioctlCodes := []uint32{
			SMART_RCV_DRIVE_DATA,
			IOCTL_DISK_GET_DRIVE_GEOMETRY,
			IOCTL_STORAGE_QUERY_PROPERTY,
			IOCTL_CPUID,
			AFD_CONNECT,
		}

		eventCount := 0
		for _, ioctlCode := range ioctlCodes {
			event := &protocol.SyscallEvent{
				ProcessID:     1234,
				ThreadID:      5678,
				CoreID:        0,
				SyscallNumber: SyscallNtDeviceIoControlFile,
				RAX:           SyscallNtDeviceIoControlFile,
				RCX:           0x1000,
				RDX:           0x2000,
				R8:            0x3000,
				R9:            0x4000,
				R10:           uint64(ioctlCode) << 32,
				RSP:           0x6000,
			}
			eventCount++
			info := ParseNtDeviceIoControlFileEvent(event)
			if info != nil {
				t.Logf("[%d] %s", eventCount, info.String())
				t.Logf("  IOCTL: %s", IoControlCodeToString(info.IoControlCode))
			}
		}

		t.Logf("Monitor finished, total events: %d", eventCount)
	})

	t.Run("HardwareHookCallback", func(t *testing.T) {
		t.Logf("Starting NtDeviceIoControlFile hook with hardware spoofing")

		walker := NewHardwareHookCallback()
		hookCallback := walker.(*HardwareHookCallback)

		ioctlCodes := []uint32{
			SMART_RCV_DRIVE_DATA,
			IOCTL_STORAGE_QUERY_PROPERTY,
			IOCTL_CPUID,
			0x00000001,
		}

		eventCount := 0
		for _, ioctlCode := range ioctlCodes {
			event := &protocol.SyscallEvent{
				ProcessID:     1234,
				ThreadID:      5678,
				CoreID:        0,
				SyscallNumber: SyscallNtDeviceIoControlFile,
				RAX:           SyscallNtDeviceIoControlFile,
				RCX:           0x1000,
				RDX:           0x2000,
				R8:            0x3000,
				R9:            0x4000,
				R10:           uint64(ioctlCode) << 32,
				RSP:           0x6000,
			}
			eventCount++
			hookCallback.Hook(event, func() {
				info := ParseNtDeviceIoControlFileEvent(event)
				if info != nil {
					t.Logf("[%d] PID: %d | IOCTL: %s",
						eventCount, info.ProcessId, IoControlCodeToString(info.IoControlCode))
				}
			})
		}

		t.Logf("Hook finished")
		t.Logf("Total events: %d", eventCount)
	})

	t.Run("IopXxxControlFileMonitor", func(t *testing.T) {
		walker := NewIopXxxControlFileMonitor()
		monitor := walker.(*IopXxxControlFileMonitor)
		t.Logf("IopXxxControlFile Monitor initialized")
		t.Logf("Target IOCTL codes:")
		for code, name := range monitor.targetIoctlCodes {
			t.Logf("  0x%08X: %s", code, name)
		}

		ioctlCodes := []uint32{
			SMART_RCV_DRIVE_DATA,
			IOCTL_DISK_GET_DRIVE_GEOMETRY,
			IOCTL_STORAGE_QUERY_PROPERTY,
			IOCTL_CPUID,
			AFD_CONNECT,
			0x00000001,
		}

		for _, ioctlCode := range ioctlCodes {
			event := &protocol.SyscallEvent{
				ProcessID:     1234,
				ThreadID:      5678,
				CoreID:        0,
				SyscallNumber: SyscallNtDeviceIoControlFile,
				RAX:           SyscallNtDeviceIoControlFile,
				RCX:           0x1000,
				RDX:           0x2000,
				R8:            0x3000,
				R9:            0x4000,
				R10:           uint64(ioctlCode) << 32,
				RSP:           0x6000,
			}
			monitor.Maiter(event, func() {
				monitor.Hook(event, func() {
					info := ParseNtDeviceIoControlFileEvent(event)
					if info != nil {
						t.Logf("Matched IOCTL: %s", IoControlCodeToString(info.IoControlCode))
					}
				})
			})
		}

		t.Logf("Monitor finished")
		t.Log(monitor.Summary())
	})

	t.Run("SSDSpoofTest", func(t *testing.T) {
		walker := NewHardwareHookCallback()
		hookCallback := walker.(*HardwareHookCallback)
		hookInfo := hookCallback.GetHookInfo()
		t.Logf("SSD Spoof Configuration:")
		t.Logf("  Model: %s", hookInfo.SsdNumber.Mod.String())
		t.Logf("  Serial: %s", hookInfo.SsdNumber.Serial.String())

		event := &protocol.SyscallEvent{
			ProcessID:     1234,
			ThreadID:      5678,
			CoreID:        0,
			SyscallNumber: SyscallNtDeviceIoControlFile,
			RAX:           SyscallNtDeviceIoControlFile,
			RCX:           0x1000,
			RDX:           0x2000,
			R8:            0x3000,
			R9:            0x4000,
			R10:           uint64(SMART_RCV_DRIVE_DATA) << 32,
			RSP:           0x6000,
		}

		smartCount := 0
		for i := 0; i < 3; i++ {
			info := ParseNtDeviceIoControlFileEvent(event)
			if info == nil {
				continue
			}
			if info.IoControlCode == SMART_RCV_DRIVE_DATA {
				smartCount++
				t.Logf("[SMART %d] PID: %d | Handle: 0x%x",
					smartCount, info.ProcessId, info.FileHandle)
				t.Logf("  Would spoof to: Model=%s, Serial=%s",
					hookInfo.SsdNumber.Mod.String(),
					hookInfo.SsdNumber.Serial.String())
			}
		}

		t.Logf("SSD spoof test finished, SMART calls: %d", smartCount)
	})

	t.Run("CPUSpoofTest", func(t *testing.T) {
		walker := NewHardwareHookCallback()
		hookCallback := walker.(*HardwareHookCallback)
		hookInfo := hookCallback.GetHookInfo()
		t.Logf("CPU Spoof Configuration:")
		t.Logf("  EAX: 0x%08X", hookInfo.Cpu1.Eax)
		t.Logf("  ECX: 0x%08X", hookInfo.Cpu1.Ecx)
		t.Logf("  EDX: 0x%08X", hookInfo.Cpu1.Edx)

		event := &protocol.SyscallEvent{
			ProcessID:     1234,
			ThreadID:      5678,
			CoreID:        0,
			SyscallNumber: SyscallNtDeviceIoControlFile,
			RAX:           SyscallNtDeviceIoControlFile,
			RCX:           0x1000,
			RDX:           0x2000,
			R8:            0x3000,
			R9:            0x4000,
			R10:           uint64(IOCTL_CPUID) << 32,
			RSP:           0x6000,
		}

		cpuidCount := 0
		for i := 0; i < 3; i++ {
			info := ParseNtDeviceIoControlFileEvent(event)
			if info == nil {
				continue
			}
			if info.IoControlCode == IOCTL_CPUID {
				cpuidCount++
				t.Logf("[CPUID %d] PID: %d | Handle: 0x%x",
					cpuidCount, info.ProcessId, info.FileHandle)
				t.Logf("  Would spoof to: EAX=0x%08X, ECX=0x%08X, EDX=0x%08X",
					hookInfo.Cpu1.Eax, hookInfo.Cpu1.Ecx, hookInfo.Cpu1.Edx)
			}
		}

		t.Logf("CPU spoof test finished, CPUID calls: %d", cpuidCount)
	})

	t.Run("NetworkConnectionMonitor", func(t *testing.T) {
		t.Logf("Starting network connection monitor")
		t.Logf("Monitoring AFD_CONNECT (0x%08X) for network connections", AFD_CONNECT)

		event := &protocol.SyscallEvent{
			ProcessID:     1234,
			ThreadID:      5678,
			CoreID:        0,
			SyscallNumber: SyscallNtDeviceIoControlFile,
			RAX:           SyscallNtDeviceIoControlFile,
			RCX:           0x1000,
			RDX:           0x2000,
			R8:            0x3000,
			R9:            0x4000,
			R10:           uint64(AFD_CONNECT) << 32,
			RSP:           0x6000,
		}

		connectionCount := 0
		for i := 0; i < 5; i++ {
			info := ParseNtDeviceIoControlFileEvent(event)
			if info == nil {
				continue
			}
			if info.IoControlCode == AFD_CONNECT {
				connectionCount++
				t.Logf("[%d] PID: %d | AFD_CONNECT detected",
					connectionCount, info.ProcessId)
			}
		}

		t.Logf("Network monitor finished, connections: %d", connectionCount)
	})
}
