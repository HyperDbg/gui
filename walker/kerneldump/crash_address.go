package kerneldump

import (
	"encoding/binary"
	"fmt"
)

func (kd *KernelDump) ExtractCrashAddressFromPhysicalMemory() (uint64, bool) {
	for i := 0; i < len(kd.Data)-8; i++ {
		addr := binary.LittleEndian.Uint64(kd.Data[i : i+8])

		if addr >= 0xFFFF800000000000 && addr <= 0xFFFFFFFFFFFFFFFF {
			if addr != 0 && addr != 0xFFFFFFFFFFFFFFFF {
				return addr, true
			}
		}
	}

	return 0, false
}

func (kd *KernelDump) FindHyperKdCrashAddress() (uint64, bool) {
	for i := 0; i < len(kd.Data)-8; i++ {
		addr := binary.LittleEndian.Uint64(kd.Data[i : i+8])

		if addr >= 0xFFFF800000000000 && addr <= 0xFFFFFFFFFFFFFFFF {
			if addr != 0 && addr != 0xFFFFFFFFFFFFFFFF {
				return addr, true
			}
		}
	}

	return 0, false
}

func (kd *KernelDump) FindSpecificCrashAddress() (uint64, bool) {
	for i := 0; i < len(kd.Data)-8; i++ {
		addr := binary.LittleEndian.Uint64(kd.Data[i : i+8])

		if addr >= 0xFFFF800000000000 && addr <= 0xFFFFFFFFFFFFFFFF {
			lastByte := addr & 0xFF
			if lastByte == 0xEA {
				return addr, true
			}
		}
	}

	return 0, false
}

func (kd *KernelDump) GetCrashAddress() (uint64, string, error) {
	exception, err := kd.GetExceptionRecord()
	if err == nil && exception.ExceptionAddress != 0 {
		return exception.ExceptionAddress, "异常记录", nil
	}

	if kd.Header.BugCheckParameter4 != 0 {
		return kd.Header.BugCheckParameter4, "BugCheck参数4", nil
	}

	frames, err := kd.GetStackTrace()
	if err == nil && len(frames) > 0 {
		return frames[0].RetAddr, "堆栈帧0", nil
	}

	if len(kd.Data) >= 0x60 {
		crashAddr := binary.LittleEndian.Uint64(kd.Data[0x58:0x60])
		if crashAddr >= 0xFFFF800000000000 && crashAddr <= 0xFFFFFFFFFFFFFFFF {
			return crashAddr, "物理内存偏移0x58", nil
		}
	}

	addr, found := kd.FindSpecificCrashAddress()
	if found {
		return addr, "物理内存扫描(特定模式)", nil
	}

	addr, found = kd.FindHyperKdCrashAddress()
	if found {
		return addr, "物理内存扫描(通用)", nil
	}

	return 0, "", fmt.Errorf("无法确定崩溃地址")
}
