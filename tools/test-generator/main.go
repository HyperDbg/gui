package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"

	"github.com/ddkwork/golibrary/std/mylog"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: test-generator <input-file> [output-file] [interface-name]")
		os.Exit(1)
	}

	inputFile := os.Args[1]

	outputFile := ""
	if len(os.Args) >= 3 {
		outputFile = os.Args[2]
	} else {
		baseName := strings.TrimSuffix(filepath.Base(inputFile), filepath.Ext(inputFile))
		outputFile = filepath.Join(filepath.Dir(inputFile), baseName+"_test.go")
	}

	interfaceName := "api"
	if len(os.Args) >= 4 {
		interfaceName = os.Args[3]
	}

	mylog.Check(generateTests(inputFile, outputFile, interfaceName))

	fmt.Printf("Successfully generated tests to %s\n", outputFile)
}

func generateTests(inputFile, outputFile, interfaceName string) error {
	fset := token.NewFileSet()
	node := mylog.Check2(parser.ParseFile(fset, inputFile, nil, parser.ParseComments))

	interfaceMethods := mylog.Check2(extractInterfaceMethods(node, interfaceName))

	if len(interfaceMethods) == 0 {
		return fmt.Errorf("no methods found in interface %s", interfaceName)
	}

	constructorInfo := findConstructor(node)
	fileContext := analyzeFileContext(node)
	packageName := filepath.Base(filepath.Dir(inputFile))
	testContent := generateTestContent(packageName, interfaceMethods, constructorInfo, fileContext)

	mylog.Check(os.WriteFile(outputFile, []byte(testContent), 0o644))

	return nil
}

func extractInterfaceMethods(node *ast.File, interfaceName string) ([]*ast.Field, error) {
	var methods []*ast.Field

	for _, decl := range node.Decls {
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok || genDecl.Tok != token.TYPE {
			continue
		}

		for _, spec := range genDecl.Specs {
			typeSpec, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}

			if typeSpec.Name.Name != interfaceName {
				continue
			}

			interfaceType, ok := typeSpec.Type.(*ast.InterfaceType)
			if !ok {
				continue
			}

			for _, method := range interfaceType.Methods.List {
				if isExported(method) {
					if _, ok := method.Type.(*ast.FuncType); ok {
						methodName := method.Names[0].Name
						if methodName != "New" && methodName != "Close" {
							methods = append(methods, method)
						}
					}
				}
			}
		}
	}

	return methods, nil
}

func isExported(field *ast.Field) bool {
	if len(field.Names) > 0 {
		return ast.IsExported(field.Names[0].Name)
	}
	return false
}

func generateTestContent(packageName string, methods []*ast.Field, constructorInfo *ConstructorInfo, fileContext *FileContext) string {
	var builder strings.Builder

	builder.WriteString(fmt.Sprintf("package %s\n\n", packageName))
	builder.WriteString("import (\n")
	builder.WriteString("\t\"testing\"\n")
	if fileContext.HasMylogImport {
		builder.WriteString("\t\"github.com/ddkwork/golibrary/std/mylog\"\n")
	}
	builder.WriteString(")\n\n")

	for _, method := range methods {
		builder.WriteString(generateTestMethod(method, constructorInfo, fileContext))
	}

	return builder.String()
}

type ConstructorInfo struct {
	Name       string
	ParamTypes []string
}

type FileContext struct {
	HasMylogImport bool
}

type MethodInfo struct {
	Name        string
	HasError    bool
	ReturnCount int
	ReturnTypes []string
	ParamCount  int
	ParamTypes  []string
}

func findConstructor(node *ast.File) *ConstructorInfo {
	for _, decl := range node.Decls {
		funcDecl, ok := decl.(*ast.FuncDecl)
		if !ok {
			continue
		}

		if funcDecl.Name.Name == "New" && funcDecl.Recv == nil {
			if funcDecl.Type.Results != nil && len(funcDecl.Type.Results.List) > 0 {
				info := &ConstructorInfo{
					Name: "New",
				}

				if funcDecl.Type.Params != nil {
					for _, param := range funcDecl.Type.Params.List {
						paramType := getTypeString(param.Type)
						if len(param.Names) == 0 {
							info.ParamTypes = append(info.ParamTypes, paramType)
						} else {
							for range param.Names {
								info.ParamTypes = append(info.ParamTypes, paramType)
							}
						}
					}
				}

				return info
			}
		}
	}

	return nil
}

func analyzeFileContext(node *ast.File) *FileContext {
	ctx := &FileContext{}

	for _, imp := range node.Imports {
		importPath := strings.Trim(imp.Path.Value, `"`)
		if strings.Contains(importPath, "github.com/ddkwork/golibrary/std/mylog") {
			ctx.HasMylogImport = true
			break
		}
	}

	return ctx
}

func generateTestMethod(method *ast.Field, constructorInfo *ConstructorInfo, fileContext *FileContext) string {
	methodName := method.Names[0].Name
	methodInfo := parseMethodInfo(method)
	var builder strings.Builder

	builder.WriteString(fmt.Sprintf("func TestProvider_%s(t *testing.T) {\n", methodName))

	if constructorInfo != nil {
		args := make([]string, len(constructorInfo.ParamTypes))
		for i, paramType := range constructorInfo.ParamTypes {
			args[i] = getDefaultValue(paramType)
		}

		if fileContext.HasMylogImport {
			builder.WriteString(fmt.Sprintf("\tp := mylog.Check2(%s(%s))\n", constructorInfo.Name, strings.Join(args, ", ")))
		} else {
			builder.WriteString(fmt.Sprintf("\tp, _ := %s(%s)\n", constructorInfo.Name, strings.Join(args, ", ")))
		}
	} else {
		builder.WriteString("\tp := &Provider{}\n")
	}

	if fileContext.HasMylogImport && methodInfo.HasError {
		checkCall := generateCheckCall(methodInfo, "p."+methodName)
		builder.WriteString(fmt.Sprintf("\t%s\n", checkCall))
	} else {
		if methodInfo.ReturnCount > 0 {
			builder.WriteString(fmt.Sprintf("\t%s\n", generateMethodCallWithReturns(methodInfo, "p."+methodName)))
		} else {
			builder.WriteString(fmt.Sprintf("\t%s\n", generateMethodCall(methodInfo, "p."+methodName)))
		}
	}

	builder.WriteString("}\n\n")

	return builder.String()
}

func generateMethodCall(methodInfo *MethodInfo, methodCall string) string {
	if methodInfo.ParamCount == 0 {
		return fmt.Sprintf("%s()", methodCall)
	}

	args := make([]string, methodInfo.ParamCount)
	for i, paramType := range methodInfo.ParamTypes {
		args[i] = getDefaultValue(paramType)
	}
	return fmt.Sprintf("%s(%s)", methodCall, strings.Join(args, ", "))
}

func parseMethodInfo(method *ast.Field) *MethodInfo {
	info := &MethodInfo{
		Name: method.Names[0].Name,
	}

	funcType := method.Type.(*ast.FuncType)

	if funcType.Params != nil {
		info.ParamCount = 0
		for _, param := range funcType.Params.List {
			typeStr := getTypeString(param.Type)
			if len(param.Names) == 0 {
				info.ParamTypes = append(info.ParamTypes, typeStr)
				info.ParamCount++
			} else {
				for range param.Names {
					info.ParamTypes = append(info.ParamTypes, typeStr)
					info.ParamCount++
				}
			}
		}
	}

	if funcType.Results != nil {
		info.ReturnCount = len(funcType.Results.List)
		for _, result := range funcType.Results.List {
			typeStr := getTypeString(result.Type)
			info.ReturnTypes = append(info.ReturnTypes, typeStr)
			if typeStr == "error" {
				info.HasError = true
			}
		}
	}

	return info
}

func generateCheckCall(methodInfo *MethodInfo, methodCall string) string {
	if methodInfo.ReturnCount == 1 && methodInfo.ReturnTypes[0] == "error" {
		if methodInfo.ParamCount == 0 {
			return fmt.Sprintf("mylog.Check(%s())", methodCall)
		}
		args := make([]string, methodInfo.ParamCount)
		for i, paramType := range methodInfo.ParamTypes {
			args[i] = getDefaultValue(paramType)
		}
		return fmt.Sprintf("mylog.Check(%s(%s))", methodCall, strings.Join(args, ", "))
	}

	if methodInfo.ReturnCount == 2 && methodInfo.HasError {
		if methodInfo.ParamCount == 0 {
			return fmt.Sprintf("mylog.Check2(%s())", methodCall)
		}
		args := make([]string, methodInfo.ParamCount)
		for i, paramType := range methodInfo.ParamTypes {
			args[i] = getDefaultValue(paramType)
		}
		return fmt.Sprintf("mylog.Check2(%s(%s))", methodCall, strings.Join(args, ", "))
	}

	if methodInfo.ReturnCount == 3 && methodInfo.HasError {
		if methodInfo.ParamCount == 0 {
			return fmt.Sprintf("mylog.Check3(%s())", methodCall)
		}
		args := make([]string, methodInfo.ParamCount)
		for i, paramType := range methodInfo.ParamTypes {
			args[i] = getDefaultValue(paramType)
		}
		return fmt.Sprintf("mylog.Check3(%s(%s))", methodCall, strings.Join(args, ", "))
	}

	if methodInfo.ReturnCount == 4 && methodInfo.HasError {
		if methodInfo.ParamCount == 0 {
			return fmt.Sprintf("mylog.Check4(%s())", methodCall)
		}
		args := make([]string, methodInfo.ParamCount)
		for i, paramType := range methodInfo.ParamTypes {
			args[i] = getDefaultValue(paramType)
		}
		return fmt.Sprintf("mylog.Check4(%s(%s))", methodCall, strings.Join(args, ", "))
	}

	if methodInfo.ReturnCount == 5 && methodInfo.HasError {
		if methodInfo.ParamCount == 0 {
			return fmt.Sprintf("mylog.Check5(%s())", methodCall)
		}
		args := make([]string, methodInfo.ParamCount)
		for i, paramType := range methodInfo.ParamTypes {
			args[i] = getDefaultValue(paramType)
		}
		return fmt.Sprintf("mylog.Check5(%s(%s))", methodCall, strings.Join(args, ", "))
	}

	if methodInfo.ReturnCount == 6 && methodInfo.HasError {
		if methodInfo.ParamCount == 0 {
			return fmt.Sprintf("mylog.Check6(%s())", methodCall)
		}
		args := make([]string, methodInfo.ParamCount)
		for i, paramType := range methodInfo.ParamTypes {
			args[i] = getDefaultValue(paramType)
		}
		return fmt.Sprintf("mylog.Check6(%s(%s))", methodCall, strings.Join(args, ", "))
	}

	if methodInfo.ReturnCount == 7 && methodInfo.HasError {
		if methodInfo.ParamCount == 0 {
			return fmt.Sprintf("mylog.Check7(%s())", methodCall)
		}
		args := make([]string, methodInfo.ParamCount)
		for i, paramType := range methodInfo.ParamTypes {
			args[i] = getDefaultValue(paramType)
		}
		return fmt.Sprintf("mylog.Check7(%s(%s))", methodCall, strings.Join(args, ", "))
	}

	return generateMethodCall(methodInfo, methodCall)
}

func generateMethodCallWithReturns(methodInfo *MethodInfo, methodCall string) string {
	if methodInfo.ReturnCount == 0 {
		return generateMethodCall(methodInfo, methodCall)
	}

	vars := make([]string, methodInfo.ReturnCount)
	for i := range vars {
		vars[i] = "_"
	}

	if methodInfo.ParamCount == 0 {
		return fmt.Sprintf("%s, %s()", strings.Join(vars, ", "), methodCall)
	}

	args := make([]string, methodInfo.ParamCount)
	for i, paramType := range methodInfo.ParamTypes {
		args[i] = getDefaultValue(paramType)
	}
	return fmt.Sprintf("%s, %s(%s)", strings.Join(vars, ", "), methodCall, strings.Join(args, ", "))
}

func getDefaultValue(typeStr string) string {
	switch typeStr {
	case "string":
		return `""`
	case "int", "int8", "int16", "int32", "int64":
		return "0"
	case "uint", "uint8", "uint16", "uint32", "uint64":
		return "0"
	case "float32", "float64":
		return "0.0"
	case "bool":
		return "false"
	default:
		if strings.HasPrefix(typeStr, "[]") {
			return "nil"
		}
		if strings.HasPrefix(typeStr, "*") {
			return "nil"
		}
		if strings.HasPrefix(typeStr, "map[") {
			return "nil"
		}
		return "nil"
	}
}

func getTypeString(expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.Ident:
		return t.Name
	case *ast.SelectorExpr:
		return fmt.Sprintf("%s.%s", t.X, t.Sel)
	case *ast.ArrayType:
		return fmt.Sprintf("[]%s", getTypeString(t.Elt))
	case *ast.StarExpr:
		return fmt.Sprintf("*%s", getTypeString(t.X))
	case *ast.MapType:
		return fmt.Sprintf("map[%s]%s", getTypeString(t.Key), getTypeString(t.Value))
	default:
		return "interface{}"
	}
}
