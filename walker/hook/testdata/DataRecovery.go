package testdata

import "github.com/ddkwork/golibrary/std/stream"

type (
	Cpu1 struct {
		Eax int64
		Ecx int64
		Edx int64
	}
	SsdNumber struct {
		Mod    *stream.Buffer
		Serial *stream.Buffer
	}
	Mac0 struct {
		Addr *stream.Buffer
	}
	HookInfo struct {
		Cpu1      Cpu1
		SsdNumber SsdNumber
		Mac0      Mac0
	}
)

func MyNotBookHook() *HookInfo {
	return &HookInfo{
		Cpu1: Cpu1{
			Eax: 0x000306A9,
			Ecx: 0x3DBAE3BF,
			Edx: 0xBFEBFBFF,
		},
		SsdNumber: SsdNumber{
			Mod:    stream.NewBuffer("Hitachi HTS545050A7E380"),
			Serial: stream.NewBuffer("TEA55A3Q2NTK8R"),
		},
		Mac0: Mac0{
			Addr: stream.NewHexString("D4BED97952EB"),
		},
	}
}

func OtherOne() *HookInfo {
	return &HookInfo{}
}
