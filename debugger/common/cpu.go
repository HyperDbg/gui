package common

import (
	"golang.org/x/sys/windows"
)

const (
	CPUIDVendorStringEAX = 0
	CPUIDFeatureInfoEAX  = 1
	CPUIDAddrWidth       = 0x80000008
)

func CpuId(eaxInput uint32) (eax, ebx, ecx, edx uint32)
func RdTsc() (retLo, retHi uint32)

func CpuReadVendorString() [13]byte {
	_, ebx, ecx, edx := CpuId(CPUIDVendorStringEAX)
	var result [13]byte
	copy(result[0:4], uint32ToBytesLe(ebx))
	copy(result[4:8], uint32ToBytesLe(edx))
	copy(result[8:12], uint32ToBytesLe(ecx))
	result[12] = 0
	return result
}

func VmxSupportDetection() bool {
	_, _, ecx, _ := CpuId(CPUIDFeatureInfoEAX)
	vmxBit := (ecx >> 5) & 1
	return vmxBit == 1
}

func uint32ToBytesLe(val uint32) []byte {
	return []byte{
		byte(val),
		byte(val >> 8),
		byte(val >> 16),
		byte(val >> 24),
	}
}

// Getx86VirtualAddressWidth returns the processor's virtual address width
func Getx86VirtualAddressWidth() uint32 {
	eax, _, _, _ := CpuId(CPUIDAddrWidth)
	return (eax >> 8) & 0xff
}

// SetPrivilegeForToken enables or disables a privilege for the current process token
func SetPrivilegeForToken(privilegeName string, enable bool) bool {
	var token windows.Token
	currentProcess, err := windows.GetCurrentProcess()
	if err != nil {
		return false
	}

	err = windows.OpenProcessToken(currentProcess, windows.TOKEN_ADJUST_PRIVILEGES|windows.TOKEN_QUERY, &token)
	if err != nil {
		return false
	}
	defer token.Close()

	var luid windows.LUID
	privName, err := windows.UTF16PtrFromString(privilegeName)
	if err != nil {
		return false
	}

	err = windows.LookupPrivilegeValue(nil, privName, &luid)
	if err != nil {
		return false
	}

	var attr uint32
	if enable {
		attr = windows.SE_PRIVILEGE_ENABLED
	} else {
		attr = 0
	}

	privileges := windows.Tokenprivileges{
		PrivilegeCount: 1,
		Privileges: [1]windows.LUIDAndAttributes{
			{Luid: luid, Attributes: attr},
		},
	}

	err = windows.AdjustTokenPrivileges(token, false, &privileges, 0, nil, nil)
	if err != nil {
		return false
	}

	return true
}
