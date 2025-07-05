package hookApiMgr

import "github.com/ddkwork/golibrary/std/stream"

type (
	cpu1 struct { // todo merge to hardware pkg add get  and set for hook
		eax int64
		ecx int64
		edx int64
	}
	ssdNumber struct {
		mod    *stream.Buffer // 40
		serial *stream.Buffer // 20
	}
	mac0 struct {
		addr *stream.Buffer // 8
	}
	hookInfo struct {
		cpu1      cpu1
		ssdNumber ssdNumber
		mac0      mac0
	}
)

func MyNotBookHook() *hookInfo {
	return &hookInfo{
		cpu1: cpu1{
			eax: 0x000306A9,
			ecx: 0x3DBAE3BF,
			edx: 0xBFEBFBFF,
		},
		ssdNumber: ssdNumber{
			mod:    stream.NewBuffer("Hitachi HTS545050A7E380"), // tigo SSD
			serial: stream.NewBuffer("TEA55A3Q2NTK8R"),          // TA1591503892
		},
		mac0: mac0{
			addr: stream.NewHexString("D4BED97952EB"), // D4BED96EFE43 修过笔记本
		},
	}
}

func OtherOne() *hookInfo {
	return &hookInfo{}
}
