//go:build amd64

package hardwareinfo

import (
	"encoding/binary"
	"fmt"

	"github.com/ddkwork/HyperDbg/debugger/driver/hardwareinfo/cpuid"
	"github.com/ddkwork/golibrary/std/stream"
)

func cpuid_low(arg1, arg2 uint32) (eax, ebx, ecx, edx uint32) // implemented in cpuidlow_amd64.s
func xgetbv_low(arg1 uint32) (eax, edx uint32)                // implemented in cpuidlow_amd64.s

func CpuidLow(arg1, arg2 uint32) (eax, ebx, ecx, edx uint32) {
	return cpuid_low(arg1, arg2)
}

type (
	Reg struct {
		Eax, Ebx, Ecx, Edx uint32
	}
	cpuInfo struct {
		Cpu0                 Reg
		Cpu1                 Reg
		Vendor               string
		ProcessorBrandString string
	}
)

func (c *cpuInfo) FormatCpu0() []string {
	return []string{
		fmt.Sprintf("%08X", c.Cpu0.Eax),
		fmt.Sprintf("%08X", c.Cpu0.Ebx),
		fmt.Sprintf("%08X", c.Cpu0.Ecx),
		fmt.Sprintf("%08X", c.Cpu0.Edx),
	}
}

func (c *cpuInfo) FormatCpu1() []string {
	return []string{
		fmt.Sprintf("%08X", c.Cpu1.Eax),
		fmt.Sprintf("%08X", c.Cpu1.Ebx),
		fmt.Sprintf("%08X", c.Cpu1.Ecx),
		fmt.Sprintf("%08X", c.Cpu1.Edx),
	}
}

func (c *cpuInfo) Get() (ok bool) {
	eax, ebx, ecx, edx := CpuidLow(0, 0)
	c.Cpu0 = Reg{
		Eax: eax,
		Ebx: ebx,
		Ecx: ecx,
		Edx: edx,
	}
	b := stream.NewBuffer("")
	b.Write(binary.LittleEndian.AppendUint32(nil, ebx))
	b.Write(binary.LittleEndian.AppendUint32(nil, edx))
	b.Write(binary.LittleEndian.AppendUint32(nil, ecx))
	c.Vendor = b.String()
	// mylog.Info("cpu vendor", b.String())

	eax, ebx, ecx, edx = CpuidLow(1, 0)
	c.Cpu1 = Reg{
		Eax: eax,
		Ebx: ebx,
		Ecx: ecx,
		Edx: edx,
	}
	// mylog.Hex("Eax", Eax)
	// mylog.Hex("ebx", ebx)
	// mylog.Hex("ecx", ecx)
	// mylog.Hex("edx", edx)
	// mylog.Info("ProcessorBrandString:", strings.TrimSpace(cpuid.ProcessorBrandString))
	c.ProcessorBrandString = cpuid.ProcessorBrandString
	return true
}
