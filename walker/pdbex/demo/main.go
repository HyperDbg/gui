package main

import (
	"fmt"
	"image"
	"sort"
	"sync"

	"gioui.org/layout"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"github.com/ddkwork/ux/app"
	"github.com/ddkwork/ux/wdk"
	"github.com/ddkwork/ux/widget/button"
	"github.com/ddkwork/ux/widget/input"
	"github.com/ddkwork/ux/widget/logview"
	"github.com/ddkwork/ux/widget/tab"
	"github.com/ddkwork/ux/widget/treetable"

	pdbex "github.com/ddkwork/HyperDbg/walker/pdbex"
)

type FuncRow struct {
	Name    string `table:"Name"`
	Address string `table:"Address"`
	Size    string `table:"Size"`
}

type SymbolRow struct {
	Name string `table:"Name"`
	Type string `table:"Type"`
	Size string `table:"Size"`
}

type PDBViewer struct {
	pdb        *pdbex.PDB
	pdbPath    string
	pathInput  *input.Input
	loadButton *button.Button
	log        *logview.LogView

	funcTable   *treetable.TreeTable[FuncRow]
	symbolTable *treetable.TreeTable[SymbolRow]
	loading     bool
	loadingMu   sync.Mutex

	pendingFuncs   []FuncRow
	pendingSymbols []SymbolRow
	dataReady      bool
	dataMu         sync.Mutex

	sourceInput  *input.Input
	sourceButton *button.Button
	sourceResult *input.Input
	funcResult   *input.Input

	tabs *tab.TabView
}

func NewPDBViewer() *PDBViewer {
	v := &PDBViewer{
		pdb:          pdbex.NewPDB(),
		pathInput:    input.FilledTextInput(),
		loadButton:   button.Filled(),
		log:          logview.New(),
		funcTable:    treetable.NewTreeTable[FuncRow](),
		symbolTable:  treetable.NewTreeTable[SymbolRow](),
		sourceInput:  input.CompactInput(),
		sourceButton: button.Outlined(),
		sourceResult: input.FilledTextArea(),
		funcResult:   input.FilledTextArea(),
	}

	v.pathInput.LabelText = "PDB Path"
	v.pathInput.Editor.SetText(`D:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\build\Release\hyperkd.pdb`)

	v.sourceInput.LabelText = "RVA (hex, e.g. 33CB0)"
	v.sourceResult.LabelText = "Source Location"
	v.funcResult.LabelText = "Function Info"

	v.initTables()

	return v
}

func (v *PDBViewer) initTables() {
	v.funcTable.AirTable.TableContext = treetable.TableContext{
		SetRootRowsCallBack: func() {},
		JsonName:            "Functions",
	}

	v.symbolTable.AirTable.TableContext = treetable.TableContext{
		SetRootRowsCallBack: func() {},
		JsonName:            "Symbols",
	}
}

func (v *PDBViewer) Update(gtx layout.Context) {
	if v.loadButton.Clicked(gtx) {
		go v.loadPDBAsync()
	}

	if v.sourceButton.Clicked(gtx) {
		v.lookupSource(gtx)
	}

	v.dataMu.Lock()
	if v.dataReady {
		v.dataReady = false
		funcRows := v.pendingFuncs
		symbolRows := v.pendingSymbols
		v.dataMu.Unlock()

		v.funcTable.Root().SetChildren(nil)
		for _, row := range funcRows {
			v.funcTable.Root().AddChild(v.funcTable.NewNode(row))
		}
		v.funcTable.AirTable.Refresh()

		v.symbolTable.Root().SetChildren(nil)
		for _, row := range symbolRows {
			v.symbolTable.Root().AddChild(v.symbolTable.NewNode(row))
		}
		v.symbolTable.AirTable.Refresh()
	} else {
		v.dataMu.Unlock()
	}
}

func (v *PDBViewer) loadPDBAsync() {
	v.loadingMu.Lock()
	if v.loading {
		v.loadingMu.Unlock()
		return
	}
	v.loading = true
	v.loadingMu.Unlock()

	defer func() {
		v.loadingMu.Lock()
		v.loading = false
		v.loadingMu.Unlock()
		app.RequestRedraw()
	}()

	v.pdbPath = v.pathInput.Editor.GetText()
	if v.pdbPath == "" {
		v.log.Error("Please enter PDB path")
		return
	}

	v.log.Infof("Loading PDB: %s", v.pdbPath)

	if err := v.pdb.Open(v.pdbPath); err != nil {
		v.log.Errorf("Failed to load PDB: %v", err)
		return
	}

	v.log.Infof("PDB loaded successfully")
	v.log.Infof("Architecture: %s", v.pdb.GetArchitectureString())

	functions := v.pdb.GetAllFunctions()
	v.log.Infof("Functions: %d", len(functions))

	var funcNames []string
	for name := range functions {
		funcNames = append(funcNames, name)
	}
	sort.Strings(funcNames)

	var funcRows []FuncRow
	for _, name := range funcNames {
		fn := functions[name]
		funcRows = append(funcRows, FuncRow{
			Name:    name,
			Address: fmt.Sprintf("0x%X", fn.Address),
			Size:    fmt.Sprintf("%d", fn.Size),
		})
	}

	udts := v.pdb.GetUDTs()
	v.log.Infof("UDTs: %d", len(udts))

	var symbolRows []SymbolRow
	for _, sym := range udts {
		symbolRows = append(symbolRows, SymbolRow{
			Name: sym.Name,
			Type: sym.Tag.String(),
			Size: fmt.Sprintf("%d", sym.Size),
		})
	}

	v.dataMu.Lock()
	v.pendingFuncs = funcRows
	v.pendingSymbols = symbolRows
	v.dataReady = true
	v.dataMu.Unlock()

	v.log.Info("PDB parsing completed")
}

func (v *PDBViewer) lookupSource(gtx layout.Context) {
	rvaStr := v.sourceInput.Editor.GetText()
	if rvaStr == "" {
		v.log.Error("Please enter RVA")
		return
	}

	var rva uint32
	if _, err := fmt.Sscanf(rvaStr, "%X", &rva); err != nil {
		if _, err := fmt.Sscanf(rvaStr, "%x", &rva); err != nil {
			v.log.Errorf("Invalid RVA format: %s", rvaStr)
			return
		}
	}

	fileName, lineNumber, ok := v.pdb.FindSourceLineByRVA(rva)
	if ok {
		v.sourceResult.Editor.SetText(fmt.Sprintf("File: %s\nLine: %d", fileName, lineNumber))
		v.log.Infof("Source: %s:%d", fileName, lineNumber)
	} else {
		v.sourceResult.Editor.SetText("Not found")
		v.log.Warn("Source location not found")
	}

	funcName, ok := v.pdb.GetFunctionByOffset(uint64(rva))
	if ok {
		v.funcResult.Editor.SetText(fmt.Sprintf("Function: %s", funcName))
		v.log.Infof("Function: %s", funcName)
	} else {
		v.funcResult.Editor.SetText("Function not found")
	}
}

func (v *PDBViewer) Layout(gtx layout.Context) layout.Dimensions {
	materialTheme := wdk.GetMaterialTheme(gtx)
	paint.FillShape(gtx.Ops, materialTheme.Scheme.Background.Color.AsNRGBA(),
		clip.Rect(image.Rect(0, 0, gtx.Constraints.Max.X, gtx.Constraints.Max.Y)).Op())

	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.UniformInset(unit.Dp(8)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
					layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
						return v.pathInput.Layout(gtx)
					}),
					layout.Rigid(layout.Spacer{Width: unit.Dp(8)}.Layout),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						label := "Load PDB"
						v.loadingMu.Lock()
						if v.loading {
							label = "Loading..."
						}
						v.loadingMu.Unlock()
						return v.loadButton.Layout(gtx, label)
					}),
				)
			})
		}),
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			return v.tabs.Layout(gtx)
		}),
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.UniformInset(unit.Dp(4)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return v.log.Layout(gtx)
			})
		}),
	)
}

func (v *PDBViewer) LayoutFunctions(gtx layout.Context) layout.Dimensions {
	return v.funcTable.AirTable.Layout(gtx)
}

func (v *PDBViewer) LayoutSymbols(gtx layout.Context) layout.Dimensions {
	return v.symbolTable.AirTable.Layout(gtx)
}

func (v *PDBViewer) LayoutSourceLookup(gtx layout.Context) layout.Dimensions {
	return layout.Flex{Axis: layout.Vertical}.Layout(gtx,
		layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return layout.UniformInset(unit.Dp(8)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return layout.Flex{Axis: layout.Horizontal, Alignment: layout.Middle}.Layout(gtx,
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return wdk.LabelM(gtx, "RVA (hex):")
					}),
					layout.Rigid(layout.Spacer{Width: unit.Dp(8)}.Layout),
					layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
						return v.sourceInput.Layout(gtx)
					}),
					layout.Rigid(layout.Spacer{Width: unit.Dp(8)}.Layout),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return v.sourceButton.Layout(gtx, "Lookup")
					}),
				)
			})
		}),
		layout.Rigid(layout.Spacer{Height: unit.Dp(8)}.Layout),
		layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
			return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
				layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
					return layout.UniformInset(unit.Dp(4)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						return v.sourceResult.Layout(gtx)
					})
				}),
				layout.Flexed(1, func(gtx layout.Context) layout.Dimensions {
					return layout.UniformInset(unit.Dp(4)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
						return v.funcResult.Layout(gtx)
					})
				}),
			)
		}),
	)
}

func main() {
	v := NewPDBViewer()

	v.tabs = tab.New(layout.Horizontal)
	v.tabs.AddTab(tab.NewTabItem("Functions", v.LayoutFunctions))
	v.tabs.AddTab(tab.NewTabItem("Symbols", v.LayoutSymbols))
	v.tabs.AddTab(tab.NewTabItem("Source Lookup", v.LayoutSourceLookup))

	app.Run("PDB Viewer", func(gtx layout.Context) {
		v.Update(gtx)
		v.Layout(gtx)
	})
}
