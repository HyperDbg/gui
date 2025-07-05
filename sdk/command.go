package sdk

type (
	Command struct {
		Name        string   `json:"Name"`
		Description string   `json:"Description"`
		Syntax      []string `json:"Syntax"`
		Examples    []string `json:"Examples"`
		Notes       []any    `json:"Notes"`
		FullName    string   `json:"FullName"`
	}
	Commands struct {
		Debugging []Command
		Extension []Command
		Hwdbg     []Command
		Meta      []Command
	}
)

/*
func LayoutCommands() ux.Widget {
	return ux.NewTableScroll(Command{}, ux.TableContext[Command]{

		MarshalRow: func(node *ux.Node[Command]) (cells []ux.CellData) {
			if node.Container() {
				node.Data.MethodName = node.SumChildren()
			}
			return []ux.CellData{
				{Text: node.Data.MethodName},
				{Text: node.Data.Cmd},
				{Text: fmt.Sprint(node.Data.Args)},
				{Text: node.Data.DoFunc},
				{Text: node.Data.Usage},
				{Text: fmt.Sprint(node.Data.Demo)},
			}
		},
		UnmarshalRow: func(node *ux.Node[Command], values []string) {
			mylog.Todo("UnmarshalRow")
		},
		SelectionChangedCallback: func(root *ux.Node[Command]) {
			mylog.Todo("SelectionChangedCallback")
		},
		SetRootRowsCallBack: func(root *ux.Node[Command]) {
			// m := orderedmap.New(InvalidCommandsKind, []Command{})
			// m.Set(DebuggingKind, commands)
			containerNodes := ux.NewContainerNodes[Command](InvalidCommandsKind.Keys())
			root.SetChildren(containerNodes)
			for _, node := range containerNodes {
				switch strings.TrimSuffix(node.Type, ux.ContainerKeyPostfix) {
				case DebuggingKind.String():
					for _, s := range InvalidDebuggingCommandsKind.Keys() {
						node.AddChildByData(Command{
							MethodName: s,
							Cmd:        s,
							Args:       nil,
							Usage:      "",
							Demo:       nil,
							DoFunc:     "Interpreter(StringToBytePointer(" + strconv.Quote(s) + "))",
						})
					}
				case ExtensionKind.String():
					for _, s := range InvalidExtensionCommandsKind.Keys() {
						node.AddChildByData(Command{
							MethodName: s,
							Cmd:        s,
							Args:       nil,
							Usage:      "",
							Demo:       nil,
							DoFunc:     "Interpreter(StringToBytePointer(" + strconv.Quote(s) + "))",
						})
					}
				case HwdbgKind.String():
					for _, s := range InvalidHwdbgCommandsKind.Keys() {
						node.AddChildByData(Command{
							MethodName: s,
							Cmd:        s,
							Args:       nil,
							Usage:      "",
							Demo:       nil,
							DoFunc:     "Interpreter(StringToBytePointer(" + strconv.Quote(s) + "))",
						})
					}
				case MetaKind.String():
					for _, s := range InvalidMetaCommandsKind.Keys() {
						node.AddChildByData(Command{
							MethodName: s,
							Cmd:        s,
							Args:       nil,
							Usage:      "",
							Demo:       nil,
							DoFunc:     "Interpreter(StringToBytePointer(" + strconv.Quote(s) + "))",
						})
					}
				default:
				}
			}
		},
		JsonName:   "commands",
		IsDocument: false,
	})
}
*/
