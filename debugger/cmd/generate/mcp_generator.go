package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"os/exec"
	"strings"
	"text/template"

	"github.com/ddkwork/golibrary/std/mylog"
)

type Method struct {
	Name           string
	Params         []Param
	Results        []string
	Comment        string
	HasReturnValue bool
	HasError       bool
}

type Param struct {
	Name string
	Type string
}

func toJSONSchema(goType string) string {
	switch {
	case strings.HasPrefix(goType, "[]"):
		return "array"
	case strings.HasPrefix(goType, "map["):
		return "object"
	case goType == "string":
		return "string"
	case goType == "int", goType == "int8", goType == "int16", goType == "int32", goType == "int64",
		goType == "uint", goType == "uint8", goType == "uint16", goType == "uint32", goType == "uint64",
		goType == "uintptr":
		return "number"
	case goType == "bool":
		return "boolean"
	default:
		return "string"
	}
}

// isBasicType 检查是否是基本类型
func isBasicType(goType string) bool {
	switch goType {
	case "string", "int", "int8", "int16", "int32", "int64",
		"uint", "uint8", "uint16", "uint32", "uint64", "uintptr",
		"bool", "[]byte", "float32", "float64":
		return true
	default:
		return false
	}
}

// isSliceType 检查是否是切片类型
func isSliceType(goType string) bool {
	return strings.HasPrefix(goType, "[]")
}

// isMapType 检查是否是map类型
func isMapType(goType string) bool {
	return strings.HasPrefix(goType, "map[")
}

func toGoType(goType string) string {
	// 基本类型保持原样
	if isBasicType(goType) {
		return goType
	}
	// 基本类型的切片保持原样
	if isSliceType(goType) {
		// 提取元素类型
		elemType := goType[2:]
		if isBasicType(elemType) {
			return goType
		}
	}
	// Map类型保持原样
	if isMapType(goType) {
		return goType
	}
	// 其他类型（包括指针和自定义类型）都用 any
	return "any"
}

var tmpl = template.Must(template.New("mcp").Funcs(template.FuncMap{
	"toJSONSchema": toJSONSchema,
	"toGoType":     toGoType,
	"jsonTag": func(name string) string {
		return "`json:\"" + name + "\"`"
	},
	"isPointer": func(paramType string) bool {
		return strings.HasPrefix(paramType, "*")
	},
	"pointerBaseType": func(paramType string) string {
		if strings.HasPrefix(paramType, "*") {
			return paramType[1:]
		}
		return paramType
	},
	"castParam": func(paramType string, paramName string) string {
		// 基本类型直接使用
		if isBasicType(paramType) {
			return "args." + paramName
		}

		// any 类型直接使用类型断言
		if paramType == "any" {
			return "args." + paramName
		}

		// interface{} 类型直接使用
		if paramType == "interface{}" {
			return "args." + paramName
		}

		// 切片类型
		if isSliceType(paramType) {
			// 提取元素类型
			elemType := paramType[2:]
			// 基本类型的切片可以直接类型断言
			if isBasicType(elemType) {
				return "args." + paramName + ".(" + paramType + ")"
			}
			// 复杂切片类型需要 JSON 转换
			// 对于包含 . 的类型（外部包），生成转换代码
			if strings.Contains(elemType, ".") {
				// 生成一个临时变量名
				return "mustConvert[" + paramType + "](args." + paramName + ")"
			}
			// 其他切片类型尝试类型断言
			return "args." + paramName + ".(" + paramType + ")"
		}

		// Map类型
		if isMapType(paramType) {
			return "args." + paramName + ".(" + paramType + ")"
		}

		// 函数类型，需要替换类型引用为 debugger 包
		if strings.HasPrefix(paramType, "func(") {
			// 将函数类型中的自定义类型添加 debugger. 前缀
			// 注意：只替换没有 debugger. 前缀的类型
			debuggerFuncType := paramType
			// 替换 *DebugEvent -> *debugger.DebugEvent
			debuggerFuncType = strings.ReplaceAll(debuggerFuncType, "*DebugEvent", "*debugger.DebugEvent")
			// 替换 DebugEvent（不是 *DebugEvent 或 debugger.DebugEvent 的情况）
			debuggerFuncType = strings.ReplaceAll(debuggerFuncType, "DebugEvent", "debugger.DebugEvent")
			// 修复双重前缀
			debuggerFuncType = strings.ReplaceAll(debuggerFuncType, "debugger.debugger.", "debugger.")
			return "args." + paramName + ".(" + debuggerFuncType + ")"
		}

		// 指针类型
		if strings.HasPrefix(paramType, "*") {
			baseType := paramType[1:]
			// 基本类型的指针
			if isBasicType(baseType) {
				return "args." + paramName + ".(*" + baseType + ")"
			}
			// 切片类型的指针
			if isSliceType(baseType) {
				return "args." + paramName + ".(*" + baseType + ")"
			}
			// Map类型的指针
			if isMapType(baseType) {
				return "args." + paramName + ".(*" + baseType + ")"
			}
			// 自定义类型的指针
			return "args." + paramName + ".(*debugger." + baseType + ")"
		}

		// 自定义类型（非指针）
		return "args." + paramName + ".(debugger." + paramType + ")"
	},
}).Parse(`// Code generated from {{.Interface}}. DO NOT EDIT.

package {{.Package}}

import (
	"context"
	"fmt"

	"github.com/ddkwork/HyperDbg/debugger"
	"github.com/ddkwork/HyperDbg/debugger/register"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type {{.ServerName}} struct {
	impl debugger.{{.Interface}}
}

func New{{.ServerName}}(impl debugger.{{.Interface}}) *{{.ServerName}} {
	return &{{.ServerName}}{impl: impl}
}

func (s *{{.ServerName}}) RegisterTools(server *mcp.Server) {
	{{range .Methods}}
	{{if .Comment}}// {{.Comment}}{{end}}
	mcp.AddTool(server, &mcp.Tool{
		Name:        "{{.Name}}",
		Description: "{{.Comment}}",
		InputSchema: map[string]any{
			"type": "object",
			"properties": map[string]any{
				{{range .Params}}
				"{{.Name}}": map[string]any{
					"type": "{{toJSONSchema .Type}}",
				},
				{{end}}
			},
		},
	}, s.handle{{.Name}})
	{{end}}
}

{{range .Methods}}
{{if .Params}}
type {{$.ServerName}}{{.Name}}Params struct {
	{{range .Params}}
	{{.Name}} {{toGoType .Type}} {{jsonTag .Name}}
	{{end}}
}
{{else}}
type {{$.ServerName}}{{.Name}}Params struct {}
{{end}}

func (s *{{$.ServerName}}) handle{{.Name}}(ctx context.Context, req *mcp.CallToolRequest, args {{$.ServerName}}{{.Name}}Params) (*mcp.CallToolResult, any, error) {
	{{if .HasReturnValue}}
	{{if .HasError}}
	result, err := s.impl.{{.Name}}(
		{{if .Params}}
		{{range $i, $p := .Params}}
		{{castParam .Type .Name}},
		{{end}}
		{{else}}
		{{end}}
	)
	if err != nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: fmt.Sprintf("Error: %v", err)},
			},
		}, nil, err
	}
	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: fmt.Sprintf("%v", result)},
		},
	}, result, nil
	{{else}}
	result := s.impl.{{.Name}}(
		{{if .Params}}
		{{range $i, $p := .Params}}
		{{castParam .Type .Name}},
		{{end}}
		{{else}}
		{{end}}
	)
	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: fmt.Sprintf("%v", result)},
		},
	}, result, nil
	{{end}}
	{{else}}
	{{if .HasError}}
	err := s.impl.{{.Name}}(
		{{if .Params}}
		{{range $i, $p := .Params}}
		{{castParam .Type .Name}},
		{{end}}
		{{else}}
		{{end}}
	)
	if err != nil {
		return &mcp.CallToolResult{
			Content: []mcp.Content{
				&mcp.TextContent{Text: fmt.Sprintf("Error: %v", err)},
			},
		}, nil, err
	}
	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: "OK"},
		},
	}, nil, nil
	{{else}}
	s.impl.{{.Name}}(
		{{if .Params}}
		{{range $i, $p := .Params}}
		{{castParam .Type .Name}},
		{{end}}
		{{else}}
		{{end}}
	)
	return &mcp.CallToolResult{
		Content: []mcp.Content{
			&mcp.TextContent{Text: "OK"},
		},
	}, nil, nil
	{{end}}
	{{end}}
}
{{end}}
`))

var helperTmpl = template.Must(template.New("helper").Parse(`// Code generated from interfaces.go. DO NOT EDIT.

package {{.Package}}

import (
	"encoding/json"
)

func getArg[T any](args map[string]any, key string) T {
	if args == nil {
		var zero T
		return zero
	}
	val, ok := args[key]
	if !ok {
		var zero T
		return zero
	}
	bytes, err := json.Marshal(val)
	if err != nil {
		var zero T
		return zero
	}
	var result T
	if err := json.Unmarshal(bytes, &result); err != nil {
		var zero T
		return zero
	}
	return result
}

// mustConvert 使用 JSON 序列化/反序列化将 any 转换为目标类型
func mustConvert[T any](v any) T {
	if v == nil {
		var zero T
		return zero
	}
	bytes, err := json.Marshal(v)
	if err != nil {
		var zero T
		return zero
	}
	var result T
	if err := json.Unmarshal(bytes, &result); err != nil {
		var zero T
		return zero
	}
	return result
}
`))

type InterfaceConfig struct {
	Interface  string
	ServerName string
	OutputPath string
}

// 排除的接口列表
var excludedInterfaces = map[string]bool{
	"Eventer": true,
}

func main() {
	interfacePath := "debugger/interfaces.go"

	configs := []InterfaceConfig{
		{
			Interface:  "KernelDebugger",
			ServerName: "KernelModeMCPServer",
			OutputPath: "debugger/mcp/kernel_mode_mcp.go",
		},
		{
			Interface:  "UserDebugger",
			ServerName: "UserModeMCPServer",
			OutputPath: "debugger/mcp/user_mode_mcp.go",
		},
	}

	mylog.Check(generateHelperFile())

	for _, config := range configs {
		mylog.Check(generateMCPServer(interfacePath, config))
	}

	fmt.Println("所有 MCP 服务器生成完成")
}

func generateHelperFile() error {
	var buf bytes.Buffer
	mylog.Check(helperTmpl.Execute(&buf, map[string]any{
		"Package": "mcp",
	}))

	mylog.Check(os.WriteFile("debugger/mcp/helper.go", buf.Bytes(), 0o644))

	fmtCmd := exec.Command("go", "run", "mvdan.cc/gofumpt@latest", "-w", "debugger/mcp/helper.go")
	fmtCmd.Dir = "."
	fmt.Printf("已生成 debugger/mcp/helper.go\n")
	return nil
}

func generateMCPServer(interfacePath string, config InterfaceConfig) error {
	fset := token.NewFileSet()
	node := mylog.Check2(parser.ParseFile(fset, interfacePath, nil, parser.ParseComments))

	methods := mylog.Check2(extractInterfaceMethods(node, config.Interface))

	fmt.Printf("找到接口: %s\n", config.Interface)
	fmt.Printf("找到 %d 个方法\n", len(methods))

	var buf bytes.Buffer
	mylog.Check(tmpl.Execute(&buf, map[string]any{
		"Package":    "mcp",
		"Methods":    methods,
		"Interface":  config.Interface,
		"ServerName": config.ServerName,
	}))

	mylog.Check(os.WriteFile(config.OutputPath, buf.Bytes(), 0o644))

	fmtCmd := exec.Command("go", "run", "mvdan.cc/gofumpt@latest", "-w", config.OutputPath)
	fmtCmd.Dir = "."
	// 运行 goimports 修复导入
	importsCmd := exec.Command("go", "run", "golang.org/x/tools/cmd/goimports@latest", "-w", config.OutputPath)
	importsCmd.Dir = "."

	fmt.Printf("已生成 %s\n", config.OutputPath)
	return nil
}

func extractInterfaceMethods(node *ast.File, ifaceName string) ([]Method, error) {
	var methods []Method

	// 递归提取接口方法
	mylog.Check(extractInterfaceMethodsRecursive(node, ifaceName, &methods))

	return methods, nil
}

func extractInterfaceMethodsRecursive(node *ast.File, ifaceName string, methods *[]Method) error {
	var found bool
	methodNames := make(map[string]bool)

	ast.Inspect(node, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.TypeSpec:
			if x.Name.Name == ifaceName {
				found = true
				iface, ok := x.Type.(*ast.InterfaceType)
				if ok {
					// 先处理嵌入的接口
					for _, field := range iface.Methods.List {
						if field.Names == nil {
							// 这是嵌入的接口
							if embeddedType, ok := field.Type.(*ast.Ident); ok {
								// 检查是否在排除列表中
								if !excludedInterfaces[embeddedType.Name] {
									// 递归提取嵌入接口的方法
									mylog.Check(extractInterfaceMethodsRecursive(node, embeddedType.Name, methods))
								}
							}
						}
					}

					// 再处理当前接口的方法
					for _, method := range iface.Methods.List {
						if len(method.Names) > 0 {
							methodName := method.Names[0].Name
							// 检查是否已经存在（避免重复）
							if !methodNames[methodName] {
								m := Method{}
								m.Name = methodName

								if method.Comment != nil {
									m.Comment = strings.TrimSpace(method.Comment.Text())
								}

								if method.Type != nil {
									if fnType, ok := method.Type.(*ast.FuncType); ok {
										for _, param := range fnType.Params.List {
											// 处理多个参数名共享同一类型的情况，如 func(a, b int)
											if len(param.Names) > 0 {
												for _, name := range param.Names {
													m.Params = append(m.Params, Param{
														Name: name.Name,
														Type: typeToString(param.Type),
													})
												}
											} else {
												// 匿名参数
												m.Params = append(m.Params, Param{
													Name: "arg",
													Type: typeToString(param.Type),
												})
											}
										}
										if fnType.Results != nil {
											for _, result := range fnType.Results.List {
												m.Results = append(m.Results, typeToString(result.Type))
											}
											// 检查是否有 error 返回值
											m.HasError = len(m.Results) > 0 && m.Results[len(m.Results)-1] == "error"
											// 检查是否有除了 error 之外的返回值
											if len(m.Results) > 1 {
												m.HasReturnValue = true
											} else if len(m.Results) == 1 && m.Results[0] != "error" {
												m.HasReturnValue = true
											}
											// 调试输出
											if m.Name == "IsConnected" {
												fmt.Printf("Debug IsConnected: Results=%v, HasReturnValue=%v, HasError=%v\n", m.Results, m.HasReturnValue, m.HasError)
											}
										}
									}
								}
								if m.Name != "" {
									*methods = append(*methods, m)
									methodNames[methodName] = true
								}
							}
						}
					}
				}
			}
		}
		return true
	})

	if !found {
		return fmt.Errorf("未找到接口: %s", ifaceName)
	}

	return nil
}

func typeToString(expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.Ident:
		return t.Name
	case *ast.ArrayType:
		if t.Len == nil {
			return "[]" + typeToString(t.Elt)
		}
	case *ast.MapType:
		return "map[" + typeToString(t.Key) + "]" + typeToString(t.Value)
	case *ast.StarExpr:
		return "*" + typeToString(t.X)
	case *ast.InterfaceType:
		return "interface{}"
	case *ast.StructType:
		return "struct{}"
	case *ast.FuncType:
		return funcTypeToString(t)
	case *ast.SelectorExpr:
		// 处理包名.类型名的情况，如 register.RegisterContext
		return typeToString(t.X) + "." + t.Sel.Name
	}
	return "any"
}

// funcTypeToString 将函数类型转换为字符串
func funcTypeToString(fn *ast.FuncType) string {
	var sb strings.Builder
	sb.WriteString("func(")
	if fn.Params != nil {
		for i, param := range fn.Params.List {
			if i > 0 {
				sb.WriteString(", ")
			}
			// 处理多个参数名共享同一类型
			if len(param.Names) > 0 {
				for j, name := range param.Names {
					if j > 0 {
						sb.WriteString(", ")
					}
					sb.WriteString(name.Name)
				}
				sb.WriteString(" ")
			}
			sb.WriteString(typeToString(param.Type))
		}
	}
	sb.WriteString(")")
	if fn.Results != nil {
		if len(fn.Results.List) == 1 {
			sb.WriteString(" ")
			sb.WriteString(typeToString(fn.Results.List[0].Type))
		} else if len(fn.Results.List) > 1 {
			sb.WriteString(" (")
			for i, result := range fn.Results.List {
				if i > 0 {
					sb.WriteString(", ")
				}
				sb.WriteString(typeToString(result.Type))
			}
			sb.WriteString(")")
		}
	}
	return sb.String()
}
