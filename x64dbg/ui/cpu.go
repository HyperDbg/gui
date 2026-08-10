package ui

import (
	"encoding/hex"

	"gioui.org/layout"
	"github.com/ddkwork/ux/widget/codeeditor"
	"github.com/ddkwork/ux/widget/split"

	"github.com/ddkwork/x64dbg/debugger"
)

type Cpu struct {
	asmAndImmData split.Split
	spTop         split.Split
	spBottom      split.Split
	sp            split.Split
	hexEditor     layout.Widget
}

func NewCpu(dbg *debugger.Debugger) *Cpu {
	disasmLayout := dbg.GetDisassembly().Layout()
	immLayout := dbg.GetImm().Layout()
	regLayout := dbg.GetRegisters().Layout()
	stackLayout := dbg.GetStack().Layout()

	asmAndImmData := split.Split{
		Ratio:  0.9,
		Bar:    10,
		Axis:   layout.Vertical,
		First:  disasmLayout,
		Second: immLayout,
	}

	spTop := split.Split{
		Ratio:  0.7,
		Bar:    10,
		Axis:   layout.Horizontal,
		First:  asmAndImmData.Layout,
		Second: regLayout,
	}
	spBottom := split.Split{
		Ratio:  0.5,
		Bar:    10,
		Axis:   layout.Horizontal,
		First:  codeeditor.New(hex.Dump([]byte{}), "go").Layout,
		Second: stackLayout,
	}

	sp := split.Split{
		Ratio:  0.7,
		Bar:    10,
		Axis:   layout.Vertical,
		First:  spTop.Layout,
		Second: spBottom.Layout,
	}

	hexEditor := codeeditor.New(hex.Dump([]byte{}), "go").Layout

	cpu := &Cpu{
		asmAndImmData: asmAndImmData,
		spTop:         spTop,
		spBottom:      spBottom,
		sp:            sp,
		hexEditor:     hexEditor,
	}

	return cpu
}

func (c *Cpu) Layout(gtx layout.Context) layout.Dimensions {
	return c.sp.Layout(gtx)
}
