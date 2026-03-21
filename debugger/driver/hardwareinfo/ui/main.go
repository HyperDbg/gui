package main

// func main() {
// 	app.Run("cpuInfo", func(w *unison.Window) {
// 		w.Content().AddChild(LayoutCpuInfo())
// 	})
// }
//
// func LayoutCpuInfo() unison.Paneler {
// 	type (
// 		Data0 struct {
// 			Arg uint32
// 			EAX uint32
// 			EBX uint32
// 			ECX uint32
// 			EDX uint32
// 		}
// 		Data1 struct {
// 			Arg uint32
// 			EAX uint32
// 			EBX uint32
// 			ECX uint32
// 			EDX uint32
// 		}
// 	)
//
// 	panel := unison.NewPanel()
// 	panel.SetLayout(&unison.FlexLayout{Columns: 2})
//
// 	eax, ebx, ecx, edx := hardwareIndo.CpuidLow(0, 0)
// 	view0, kvPanel0 := widget.NewStructView(Data0{
// 		Arg: 0,
// 		EAX: eax,
// 		EBX: ebx,
// 		ECX: ecx,
// 		EDX: edx,
// 	}, func(data Data0) (values []widget.CellData) {
// 		return []widget.CellData{
// 			{Text: fmt.Sprintf("%08X", data.Arg)},
// 			{Text: fmt.Sprintf("%08X", data.EAX)},
// 			{Text: fmt.Sprintf("%08X", data.EBX)},
// 			{Text: fmt.Sprintf("%08X", data.ECX)},
// 			{Text: fmt.Sprintf("%08X", data.EDX)},
// 		}
// 	})
//
// 	eax, ebx, ecx, edx = hardwareIndo.CpuidLow(0, 1)
// 	view1, kvPanel1 := widget.NewStructView(Data0{
// 		Arg: 1,
// 		EAX: eax,
// 		EBX: ebx,
// 		ECX: ecx,
// 		EDX: edx,
// 	}, func(data Data0) (values []widget.CellData) {
// 		return []widget.CellData{
// 			{Text: fmt.Sprintf("%08X", data.Arg)},
// 			{Text: fmt.Sprintf("%08X", data.EAX)},
// 			{Text: fmt.Sprintf("%08X", data.EBX)},
// 			{Text: fmt.Sprintf("%08X", data.ECX)},
// 			{Text: fmt.Sprintf("%08X", data.EDX)},
// 		}
// 	})
// 	panel.AddChild(view0)
// 	// panel.AddChild(kvPanel0)
// 	panel.AddChild(view1)
// 	// panel.AddChild(kvPanel1)
// 	kvPanel0 = kvPanel0
// 	kvPanel1 = kvPanel1
// 	return panel
// }
