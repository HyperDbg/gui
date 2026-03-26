package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"os/exec"
	"path/filepath"
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

func isSliceType(goType string) bool {
	return strings.HasPrefix(goType, "[]")
}

func isMapType(goType string) bool {
	return strings.HasPrefix(goType, "map[")
}

func toGoType(goType string) string {
	if isBasicType(goType) {
		return goType
	}
	if isSliceType(goType) {
		elemType := goType[2:]
		if isBasicType(elemType) {
			return goType
		}
	}
	if isMapType(goType) {
		return goType
	}
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
		if isBasicType(paramType) {
			return "args." + paramName
		}

		if paramType == "any" || paramType == "interface{}" {
			return "args." + paramName
		}

		if isSliceType(paramType) {
			elemType := paramType[2:]
			if isBasicType(elemType) {
				return "args." + paramName + ".(" + paramType + ")"
			}
			if strings.Contains(elemType, ".") {
				return "mustConvert[" + paramType + "](args." + paramName + ")"
			}
			return "args." + paramName + ".([]hyperdbgrust." + elemType + ")"
		}

		if isMapType(paramType) {
			return "args." + paramName + ".(" + paramType + ")"
		}

		if strings.HasPrefix(paramType, "func(") {
			return "args." + paramName + ".(" + paramType + ")"
		}

		if strings.Contains(paramType, ".") {
			return "args." + paramName + ".(" + paramType + ")"
		}

		if strings.HasPrefix(paramType, "*") {
			baseType := paramType[1:]
			if isBasicType(baseType) {
				return "args." + paramName + ".(*" + baseType + ")"
			}
			if isSliceType(baseType) {
				return "args." + paramName + ".(*" + baseType + ")"
			}
			if isMapType(baseType) {
				return "args." + paramName + ".(*" + baseType + ")"
			}
			if strings.Contains(baseType, ".") {
				return "args." + paramName + ".(*" + baseType + ")"
			}
			return "args." + paramName + ".(*hyperdbgrust." + baseType + ")"
		}

		return "args." + paramName + ".(hyperdbgrust." + paramType + ")"
	},
}).Parse(`// Code generated from {{.Interface}}. DO NOT EDIT.

package main

import (
	"context"
	"fmt"
	"time"

	"github.com/ddkwork/HyperDbg_rust"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type {{.ServerName}} struct {
	impl hyperdbgrust.{{.Interface}}
}

func New{{.ServerName}}(impl hyperdbgrust.{{.Interface}}) *{{.ServerName}} {
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

func main() {
	impl := hyperdbgrust.NewPacket()
	server := New{{.ServerName}}(impl)

	mcpServer := mcp.NewServer(&mcp.Implementation{
		Name:    "hyperdbg-mcp",
		Version: "1.0.0",
	}, nil)
	server.RegisterTools(mcpServer)

	ctx := context.Background()
	if err := mcpServer.Run(ctx, &mcp.StdioTransport{}); err != nil {
		panic(err)
	}
}
`))

type InterfaceConfig struct {
	Interface  string
	ServerName string
	OutputPath string
}

var excludedInterfaces = map[string]bool{}

func main() {
	interfacePath := "interfaces.go"

	configs := []InterfaceConfig{
		{
			Interface:  "Debugger",
			ServerName: "DebuggerMCPServer",
			OutputPath: "cmd/mcp/mcp.go",
		},
	}

	for _, config := range configs {
		mylog.Check(generateMCPServer(interfacePath, config))
	}

	fmt.Println("MCP 服务器生成完成")
}

func generateMCPServer(interfacePath string, config InterfaceConfig) error {
	fset := token.NewFileSet()
	node := mylog.Check2(parser.ParseFile(fset, interfacePath, nil, parser.ParseComments))

	methods := mylog.Check2(extractInterfaceMethods(node, config.Interface))

	fmt.Printf("找到接口: %s\n", config.Interface)
	fmt.Printf("找到 %d 个方法\n", len(methods))

	dir := filepath.Dir(config.OutputPath)
	mylog.Check(os.MkdirAll(dir, 0o755))

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
	importsCmd := exec.Command("go", "run", "golang.org/x/tools/cmd/goimports@latest", "-w", config.OutputPath)
	importsCmd.Dir = "."

	fmt.Printf("已生成 %s\n", config.OutputPath)
	return nil
}

func extractInterfaceMethods(node *ast.File, ifaceName string) ([]Method, error) {
	var methods []Method
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
					for _, field := range iface.Methods.List {
						if field.Names == nil {
							if embeddedType, ok := field.Type.(*ast.Ident); ok {
								if !excludedInterfaces[embeddedType.Name] {
									mylog.Check(extractInterfaceMethodsRecursive(node, embeddedType.Name, methods))
								}
							}
						}
					}

					for _, method := range iface.Methods.List {
						if len(method.Names) > 0 {
							methodName := method.Names[0].Name
							if !methodNames[methodName] {
								m := Method{}
								m.Name = methodName

								if method.Comment != nil {
									m.Comment = strings.TrimSpace(method.Comment.Text())
								}

								if method.Type != nil {
									if fnType, ok := method.Type.(*ast.FuncType); ok {
										for _, param := range fnType.Params.List {
											if len(param.Names) > 0 {
												for _, name := range param.Names {
													m.Params = append(m.Params, Param{
														Name: name.Name,
														Type: typeToString(param.Type),
													})
												}
											} else {
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
											m.HasError = len(m.Results) > 0 && m.Results[len(m.Results)-1] == "error"
											if len(m.Results) > 1 {
												m.HasReturnValue = true
											} else if len(m.Results) == 1 && m.Results[0] != "error" {
												m.HasReturnValue = true
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
		return typeToString(t.X) + "." + t.Sel.Name
	}
	return "any"
}

func funcTypeToString(fn *ast.FuncType) string {
	var sb strings.Builder
	sb.WriteString("func(")
	if fn.Params != nil {
		for i, param := range fn.Params.List {
			if i > 0 {
				sb.WriteString(", ")
			}
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
