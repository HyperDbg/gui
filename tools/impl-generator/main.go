package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
	"path/filepath"
	"strings"

	"github.com/ddkwork/golibrary/std/mylog"
)

type MethodInfo struct {
	Field           *ast.Field
	SourceInterface string
	IsEmbedded      bool
	Index           int
}

func main() {
	if len(os.Args) < 5 {
		fmt.Println("Usage:")
		fmt.Println("  impl-generator <input-file> <output-file> <struct-name> <interface-name>")
		fmt.Println("  impl-generator --update-comments <input-file> <output-file> <struct-name> <interface-name>")
		fmt.Println("  impl-generator --check-implementation <input-file> <output-file> <struct-name> <interface-name>")
		os.Exit(1)
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]
	structName := os.Args[3]
	interfaceName := os.Args[4]

	updateComments := false
	checkImplementation := false
	if inputFile == "--update-comments" {
		updateComments = true
		inputFile = os.Args[2]
		outputFile = os.Args[3]
		structName = os.Args[4]
		interfaceName = os.Args[5]
	} else if inputFile == "--check-implementation" {
		checkImplementation = true
		inputFile = os.Args[2]
		outputFile = os.Args[3]
		structName = os.Args[4]
		interfaceName = os.Args[5]
	}

	if updateComments {
		mylog.Check(updateMethodComments(inputFile, outputFile, structName, interfaceName))
	} else if checkImplementation {
		mylog.Check(checkImplementationAndComment(inputFile, outputFile, structName, interfaceName))
	} else {
		mylog.Check(generateImplementations(inputFile, outputFile, structName, interfaceName))
	}

	fmt.Printf("Successfully generated to %s\n", outputFile)
}

func generateImplementations(inputFile, outputFile, structName, interfaceName string) error {
	fset := token.NewFileSet()
	node := mylog.Check2(parser.ParseFile(fset, inputFile, nil, parser.ParseComments))

	interfaceMethods := mylog.Check2(extractInterfaceMethods(node, interfaceName))

	if len(interfaceMethods) == 0 {
		return fmt.Errorf("no methods found in interface %s", interfaceName)
	}

	structMethods := mylog.Check2(extractStructMethods(node, structName))

	packageName := filepath.Base(filepath.Dir(inputFile))
	content := generateImplContent(packageName, structName, interfaceMethods, structMethods)

	mylog.Check(os.WriteFile(outputFile, []byte(content), 0o644))

	return nil
}

func extractInterfaceMethods(node *ast.File, interfaceName string) ([]MethodInfo, error) {
	methods := make([]MethodInfo, 0)

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

			extractedMethods := extractMethodsFromInterface(interfaceType, node, interfaceName)
			methods = append(methods, extractedMethods...)
		}
	}

	return methods, nil
}

func extractMethodsFromInterface(interfaceType *ast.InterfaceType, node *ast.File, currentInterface string) []MethodInfo {
	methods := make([]MethodInfo, 0)

	for i, method := range interfaceType.Methods.List {
		if len(method.Names) > 0 && isExported(method) {
			if _, ok := method.Type.(*ast.FuncType); ok {
				methods = append(methods, MethodInfo{
					Field:           method,
					SourceInterface: currentInterface,
					IsEmbedded:      false,
					Index:           i + 1,
				})
			}
		} else if len(method.Names) == 0 {
			if ident, ok := method.Type.(*ast.Ident); ok {
				embeddedInterfaceName := ident.Name
				embeddedMethods := mylog.Check2(extractInterfaceMethods(node, embeddedInterfaceName))
				for _, embeddedMethod := range embeddedMethods {
					embeddedMethod.IsEmbedded = true
					methods = append(methods, embeddedMethod)
				}
			}
		}
	}

	return methods
}

func extractStructMethods(node *ast.File, structName string) (map[string]*ast.FuncDecl, error) {
	methods := make(map[string]*ast.FuncDecl)

	for _, decl := range node.Decls {
		funcDecl, ok := decl.(*ast.FuncDecl)
		if !ok {
			continue
		}

		if funcDecl.Recv == nil || len(funcDecl.Recv.List) == 0 {
			continue
		}

		recv := funcDecl.Recv.List[0]
		starExpr, ok := recv.Type.(*ast.StarExpr)
		if !ok {
			continue
		}

		ident, ok := starExpr.X.(*ast.Ident)
		if !ok || ident.Name != structName {
			continue
		}

		methods[funcDecl.Name.Name] = funcDecl
	}

	return methods, nil
}

func isExported(field *ast.Field) bool {
	if len(field.Names) > 0 {
		return ast.IsExported(field.Names[0].Name)
	}
	return false
}

func generateImplContent(packageName, structName string, interfaceMethods []MethodInfo, structMethods map[string]*ast.FuncDecl) string {
	var builder strings.Builder

	builder.WriteString(fmt.Sprintf("package %s\n\n", packageName))
	builder.WriteString("// Generated interface implementations\n")
	builder.WriteString("// TODO: Review and implement the methods below\n\n")

	for _, method := range interfaceMethods {
		methodName := method.Field.Names[0].Name
		builder.WriteString(generateMethod(structName, methodName, method, structMethods[methodName]))
	}

	return builder.String()
}

func generateMethod(structName, methodName string, interfaceMethod MethodInfo, structMethod *ast.FuncDecl) string {
	var builder strings.Builder

	builder.WriteString(fmt.Sprintf("// 来自接口: %s 第%d个签名\n", interfaceMethod.SourceInterface, interfaceMethod.Index))

	if structMethod != nil {
		signatureMatch := compareSignatures(interfaceMethod.Field, structMethod)
		if signatureMatch {
			return ""
		}

		builder.WriteString(fmt.Sprintf("// TODO: Signature mismatch with interface %s\n", methodName))
		builder.WriteString(fmt.Sprintf("// Expected: %s\n", getMethodSignature(interfaceMethod.Field)))
		builder.WriteString(fmt.Sprintf("// Current: %s\n", getMethodSignature(structMethod.Type)))
	} else {
		if interfaceMethod.IsEmbedded {
			builder.WriteString(fmt.Sprintf("// TODO: Implement %s (from %s)\n", methodName, interfaceMethod.SourceInterface))
		} else {
			builder.WriteString(fmt.Sprintf("// TODO: Implement %s\n", methodName))
		}
	}

	builder.WriteString(extractMethodComments(interfaceMethod.Field))
	builder.WriteString(fmt.Sprintf("func (s *%s) %s", structName, methodName))

	funcType := interfaceMethod.Field.Type.(*ast.FuncType)

	if funcType.Params != nil && len(funcType.Params.List) > 0 {
		builder.WriteString("(")
		for i, param := range funcType.Params.List {
			if i > 0 {
				builder.WriteString(", ")
			}
			for j, name := range param.Names {
				if j > 0 {
					builder.WriteString(", ")
				}
				builder.WriteString(fmt.Sprintf("%s %s", name, getTypeString(param.Type)))
			}
		}
		builder.WriteString(")")
	}

	if funcType.Results != nil && len(funcType.Results.List) > 0 {
		builder.WriteString(" (")
		for i, result := range funcType.Results.List {
			if i > 0 {
				builder.WriteString(", ")
			}
			builder.WriteString(getTypeString(result.Type))
		}
		builder.WriteString(")")
	}

	if interfaceMethod.IsEmbedded {
		var paramList []string
		if funcType.Params != nil && len(funcType.Params.List) > 0 {
			for _, param := range funcType.Params.List {
				for _, name := range param.Names {
					paramList = append(paramList, name.Name)
				}
			}
		}

		fieldName := strings.ToLower(interfaceMethod.SourceInterface)
		builder.WriteString(fmt.Sprintf(" { return s.%s(%s) }\n\n", fieldName, strings.Join(paramList, ", ")))
	} else {
		builder.WriteString(" {\n")
		builder.WriteString("\t// TODO: Implement this method\n")

		if funcType.Results != nil && len(funcType.Results.List) > 0 {
			returnValues := make([]string, 0, len(funcType.Results.List))
			for _, result := range funcType.Results.List {
				returnValues = append(returnValues, getZeroValue(result.Type))
			}
			builder.WriteString(fmt.Sprintf("\treturn %s\n", strings.Join(returnValues, ", ")))
		}

		builder.WriteString("}\n\n")
	}

	return builder.String()
}

func checkImplementationAndComment(inputFile, outputFile, structName, interfaceName string) error {
	fset := token.NewFileSet()

	interfaceNode := mylog.Check2(parser.ParseFile(fset, inputFile, nil, parser.ParseComments))

	var implNode *ast.File
	mylog.Check2(os.Stat(outputFile))
	implNode = mylog.Check2(parser.ParseFile(fset, outputFile, nil, parser.ParseComments))

	interfaceMethods := mylog.Check2(extractInterfaceMethods(interfaceNode, interfaceName))

	if len(interfaceMethods) == 0 {
		return fmt.Errorf("no methods found in interface %s", interfaceName)
	}

	content := generateCheckedImplementationContent(implNode, structName, interfaceMethods, fset, outputFile)

	mylog.Check(os.WriteFile(outputFile, []byte(content), 0o644))

	return nil
}

func generateCheckedImplementationContent(node *ast.File, structName string, interfaceMethods []MethodInfo, fset *token.FileSet, outputPath string) string {
	var builder strings.Builder

	builder.WriteString(fmt.Sprintf("package %s\n\n", node.Name.Name))

	implementedMethods := make(map[string]*ast.FuncDecl)
	var misplacedMethods []string

	for _, decl := range node.Decls {
		if genDecl, ok := decl.(*ast.GenDecl); ok {
			if genDecl.Tok == token.IMPORT {
				builder.WriteString("import (\n")
				for _, spec := range genDecl.Specs {
					importSpec := spec.(*ast.ImportSpec)
					builder.WriteString(fmt.Sprintf("\t%s\n", importSpec.Path.Value))
				}
				builder.WriteString(")\n\n")
			} else if genDecl.Tok == token.TYPE {
				var buf bytes.Buffer
				if e := (format.Node(&buf, fset, genDecl)); e == nil {
					builder.WriteString(buf.String())
					builder.WriteString("\n")
				}
			} else if genDecl.Tok == token.CONST || genDecl.Tok == token.VAR {
				var buf bytes.Buffer
				if e := (format.Node(&buf, fset, genDecl)); e == nil {
					builder.WriteString(buf.String())
					builder.WriteString("\n")
				}
			} else {
				var buf bytes.Buffer
				if e := (format.Node(&buf, fset, genDecl)); e == nil {
					builder.WriteString(buf.String())
					builder.WriteString("\n")
				}
			}
		} else if funcDecl, ok := decl.(*ast.FuncDecl); ok {
			if funcDecl.Recv == nil || len(funcDecl.Recv.List) == 0 {
				var buf bytes.Buffer
				if e := (format.Node(&buf, fset, funcDecl)); e == nil {
					builder.WriteString(buf.String())
					builder.WriteString("\n")
				}
			} else {
				recv := funcDecl.Recv.List[0]
				if starExpr, ok := recv.Type.(*ast.StarExpr); ok {
					if ident, ok := starExpr.X.(*ast.Ident); ok && ident.Name == structName {
						methodName := funcDecl.Name.Name
						if isInInterface(methodName, interfaceMethods) {
							implementedMethods[methodName] = funcDecl
						} else {
							if !strings.HasSuffix(methodName, "_skip") {
								var buf bytes.Buffer
								mylog.Check(format.Node(&buf, fset, funcDecl))
								misplacedMethods = append(misplacedMethods, buf.String())
								builder.WriteString("// TODO: 实现错了，需要移动\n")
								builder.WriteString("/*\n")
								builder.WriteString(buf.String())
								builder.WriteString("*/\n\n")
							}
						}
					}
				}
			}
		}
	}

	for _, method := range interfaceMethods {
		methodName := method.Field.Names[0].Name

		if funcDecl, exists := implementedMethods[methodName]; exists {
			hasIndexComment := false
			if funcDecl.Doc != nil {
				docText := funcDecl.Doc.Text()
				if strings.Contains(docText, "来自接口") {
					hasIndexComment = true
				}
			}

			if !hasIndexComment {
				builder.WriteString(fmt.Sprintf("// 来自接口: %s 第%d个签名\n", method.SourceInterface, method.Index))
			}

			var buf bytes.Buffer
			if e := (format.Node(&buf, fset, funcDecl)); e == nil {
				builder.WriteString(buf.String())
				builder.WriteString("\n")
			}
		} else {
			builder.WriteString(fmt.Sprintf("// 来自接口: %s 第%d个签名\n", method.SourceInterface, method.Index))
			if method.IsEmbedded {
				builder.WriteString(fmt.Sprintf("// TODO: Implement %s (from %s)\n", methodName, method.SourceInterface))
			} else {
				builder.WriteString(fmt.Sprintf("// TODO: Implement %s\n", methodName))
			}
			builder.WriteString(extractMethodComments(method.Field))
			builder.WriteString(fmt.Sprintf("func (s *%s) %s", structName, methodName))

			funcType := method.Field.Type.(*ast.FuncType)

			if funcType.Params != nil && len(funcType.Params.List) > 0 {
				builder.WriteString("(")
				for i, param := range funcType.Params.List {
					if i > 0 {
						builder.WriteString(", ")
					}
					for j, name := range param.Names {
						if j > 0 {
							builder.WriteString(", ")
						}
						builder.WriteString(fmt.Sprintf("%s %s", name, getTypeString(param.Type)))
					}
				}
				builder.WriteString(")")
			} else {
				builder.WriteString("()")
			}

			if funcType.Results != nil && len(funcType.Results.List) > 0 {
				builder.WriteString(" (")
				for i, result := range funcType.Results.List {
					if i > 0 {
						builder.WriteString(", ")
					}
					builder.WriteString(getTypeString(result.Type))
				}
				builder.WriteString(")")
			}

			if method.IsEmbedded {
				builder.WriteString(" { return s.")
				builder.WriteString(strings.ToLower(method.SourceInterface))
				builder.WriteString(".")
				builder.WriteString(methodName)
				builder.WriteString("(")
				if funcType.Params != nil && len(funcType.Params.List) > 0 {
					for i, param := range funcType.Params.List {
						if i > 0 {
							builder.WriteString(", ")
						}
						for j, name := range param.Names {
							if j > 0 {
								builder.WriteString(", ")
							}
							builder.WriteString(name.Name)
						}
					}
				}
				builder.WriteString(") }\n\n")
			} else {
				if funcType.Results != nil && len(funcType.Results.List) > 0 {
					builder.WriteString(" {\n\tpanic(\"not implemented\")\n}\n\n")
				} else {
					builder.WriteString(" {\n\t// TODO: implement\n}\n\n")
				}
			}
		}
	}

	if len(misplacedMethods) > 0 {
		badFileBuilder := strings.Builder{}
		badFileBuilder.WriteString(fmt.Sprintf("package %s\n\n", node.Name.Name))
		badFileBuilder.WriteString("// ==================== 以下方法实现位置错误，需要移动 ====================\n\n")

		for _, methodCode := range misplacedMethods {
			badFileBuilder.WriteString("// TODO: 实现错了，需要移动\n")
			badFileBuilder.WriteString(methodCode)
			badFileBuilder.WriteString("\n")
		}

		badFileName := filepath.Join(filepath.Dir(outputPath), filepath.Base(outputPath[:len(outputPath)-3])+"_bad.go")
		mylog.Check(os.WriteFile(badFileName, []byte(badFileBuilder.String()), 0o644))
	}

	return builder.String()
}

func isInInterface(methodName string, interfaceMethods []MethodInfo) bool {
	for _, method := range interfaceMethods {
		if method.Field.Names[0].Name == methodName {
			return true
		}
	}
	return false
}

func getExprString(expr ast.Expr) string {
	var buf bytes.Buffer
	mylog.Check(printer.Fprint(&buf, token.NewFileSet(), expr))
	return buf.String()
}

func getStmtString(stmt ast.Stmt, fset *token.FileSet) string {
	var buf bytes.Buffer
	mylog.Check(printer.Fprint(&buf, fset, stmt))
	return buf.String()
}

func compareSignatures(interfaceMethod *ast.Field, structMethod *ast.FuncDecl) bool {
	ifaceType := interfaceMethod.Type.(*ast.FuncType)
	structType := structMethod.Type

	return signaturesEqual(ifaceType, structType)
}

func signaturesEqual(a, b ast.Expr) bool {
	aFunc, ok1 := a.(*ast.FuncType)
	bFunc, ok2 := b.(*ast.FuncType)

	if !ok1 || !ok2 {
		return false
	}

	if !paramsEqual(aFunc.Params, bFunc.Params) {
		return false
	}

	if !resultsEqual(aFunc.Results, bFunc.Results) {
		return false
	}

	return true
}

func paramsEqual(a, b *ast.FieldList) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if len(a.List) != len(b.List) {
		return false
	}

	for i := range a.List {
		if !typesEqual(a.List[i].Type, b.List[i].Type) {
			return false
		}
	}

	return true
}

func resultsEqual(a, b *ast.FieldList) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if len(a.List) != len(b.List) {
		return false
	}

	for i := range a.List {
		if !typesEqual(a.List[i].Type, b.List[i].Type) {
			return false
		}
	}

	return true
}

func typesEqual(a, b ast.Expr) bool {
	aStr := getTypeString(a)
	bStr := getTypeString(b)
	return aStr == bStr
}

func getMethodSignature(method any) string {
	var builder strings.Builder

	switch m := method.(type) {
	case *ast.Field:
		builder.WriteString(m.Names[0].Name)
		funcType := m.Type.(*ast.FuncType)
		builder.WriteString(getFuncTypeString(funcType))
	case *ast.FuncType:
		builder.WriteString(getFuncTypeString(m))
	}

	return builder.String()
}

func getFuncTypeString(funcType *ast.FuncType) string {
	var builder strings.Builder

	if funcType.Params == nil || len(funcType.Params.List) == 0 {
		builder.WriteString("()")
	} else {
		builder.WriteString("(")
		for i, param := range funcType.Params.List {
			if i > 0 {
				builder.WriteString(", ")
			}
			for j, name := range param.Names {
				if j > 0 {
					builder.WriteString(", ")
				}
				builder.WriteString(fmt.Sprintf("%s %s", name, getTypeString(param.Type)))
			}
		}
		builder.WriteString(")")
	}

	if funcType.Results != nil && len(funcType.Results.List) > 0 {
		if len(funcType.Results.List) == 1 {
			builder.WriteString(" ")
			builder.WriteString(getTypeString(funcType.Results.List[0].Type))
		} else {
			builder.WriteString(" (")
			for i, result := range funcType.Results.List {
				if i > 0 {
					builder.WriteString(", ")
				}
				builder.WriteString(getTypeString(result.Type))
			}
			builder.WriteString(")")
		}
	}

	return builder.String()
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
	case *ast.FuncType:
		return getFuncTypeString(t)
	default:
		return "interface{}"
	}
}

func getZeroValue(expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.Ident:
		switch t.Name {
		case "string":
			return `""`
		case "bool":
			return "false"
		case "int", "int8", "int16", "int32", "int64":
			return "0"
		case "uint", "uint8", "uint16", "uint32", "uint64":
			return "0"
		case "float32", "float64":
			return "0.0"
		case "error":
			return "nil"
		default:
			return fmt.Sprintf("%s{}", t.Name)
		}
	case *ast.ArrayType:
		return "nil"
	case *ast.StarExpr:
		return "nil"
	case *ast.MapType:
		return "nil"
	case *ast.InterfaceType:
		return "nil"
	default:
		return "nil"
	}
}

func extractMethodComments(field *ast.Field) string {
	if field.Doc == nil {
		return ""
	}

	var builder strings.Builder
	for _, comment := range field.Doc.List {
		text := comment.Text
		if !strings.HasPrefix(text, "//") {
			builder.WriteString(fmt.Sprintf("// %s\n", text))
		} else {
			builder.WriteString(fmt.Sprintf("%s\n", text))
		}
	}

	return builder.String()
}

func updateMethodComments(inputFile, outputFile, structName, interfaceName string) error {
	fset := token.NewFileSet()

	interfaceNode := mylog.Check2(parser.ParseFile(fset, inputFile, nil, parser.ParseComments))

	var implNode *ast.File
	if _, e := (os.Stat(outputFile)); os.IsNotExist(e) {
		implNode = &ast.File{
			Name:  ast.NewIdent(filepath.Base(filepath.Dir(inputFile))),
			Scope: ast.NewScope(nil),
		}
	} else {
		implNode = mylog.Check2(parser.ParseFile(fset, outputFile, nil, parser.ParseComments))
	}

	interfaceMethods := mylog.Check2(extractInterfaceMethods(interfaceNode, interfaceName))

	if len(interfaceMethods) == 0 {
		return fmt.Errorf("no methods found in interface %s", interfaceName)
	}

	packageName := filepath.Base(filepath.Dir(inputFile))
	content := generateUpdatedCommentsContent(packageName, structName, interfaceMethods, implNode)

	mylog.Check(os.WriteFile(outputFile, []byte(content), 0o644))

	return nil
}

func generateUpdatedCommentsContent(packageName, structName string, interfaceMethods []MethodInfo, implNode *ast.File) string {
	var builder strings.Builder

	builder.WriteString(fmt.Sprintf("package %s\n\n", packageName))
	builder.WriteString("// Updated method comments from interface\n")
	builder.WriteString("// TODO: Review updated comments and apply to implementation\n\n")

	existingMethods := make(map[string]*ast.FuncDecl)
	for _, decl := range implNode.Decls {
		if funcDecl, ok := decl.(*ast.FuncDecl); ok {
			if funcDecl.Recv != nil && len(funcDecl.Recv.List) > 0 {
				recv := funcDecl.Recv.List[0]
				if starExpr, ok := recv.Type.(*ast.StarExpr); ok {
					if ident, ok := starExpr.X.(*ast.Ident); ok && ident.Name == structName {
						existingMethods[funcDecl.Name.Name] = funcDecl
					}
				}
			}
		}
	}

	if len(existingMethods) == 0 {
		builder.WriteString(fmt.Sprintf("// No existing methods found for struct '%s'. Generating all methods with interface comments:\n\n", structName))
		for _, method := range interfaceMethods {
			methodName := method.Field.Names[0].Name
			builder.WriteString(generateMethodWithComments(structName, methodName, method))
		}
	} else {
		builder.WriteString(fmt.Sprintf("// Found %d existing methods for struct '%s'. Generating comment updates:\n\n", len(existingMethods), structName))
		for _, method := range interfaceMethods {
			methodName := method.Field.Names[0].Name
			if existingMethod, exists := existingMethods[methodName]; exists {
				builder.WriteString(generateCommentUpdateMethod(structName, methodName, method, existingMethod))
			}
		}
	}

	return builder.String()
}

func generateMethodWithComments(structName, methodName string, interfaceMethod MethodInfo) string {
	var builder strings.Builder

	builder.WriteString(fmt.Sprintf("// TODO: Implement %s", methodName))
	if interfaceMethod.IsEmbedded {
		builder.WriteString(fmt.Sprintf(" (from %s)", interfaceMethod.SourceInterface))
	}
	builder.WriteString("\n")
	builder.WriteString(extractMethodComments(interfaceMethod.Field))
	builder.WriteString(fmt.Sprintf("func (s *%s) %s", structName, methodName))

	funcType := interfaceMethod.Field.Type.(*ast.FuncType)

	if funcType.Params != nil && len(funcType.Params.List) > 0 {
		builder.WriteString("(")
		for i, param := range funcType.Params.List {
			if i > 0 {
				builder.WriteString(", ")
			}
			for j, name := range param.Names {
				if j > 0 {
					builder.WriteString(", ")
				}
				builder.WriteString(fmt.Sprintf("%s %s", name, getTypeString(param.Type)))
			}
		}
		builder.WriteString(")")
	}

	if funcType.Results != nil && len(funcType.Results.List) > 0 {
		builder.WriteString(" (")
		for i, result := range funcType.Results.List {
			if i > 0 {
				builder.WriteString(", ")
			}
			builder.WriteString(getTypeString(result.Type))
		}
		builder.WriteString(")")
	}

	builder.WriteString(" {\n")
	builder.WriteString("\t// TODO: Implement this method\n")

	if funcType.Results != nil && len(funcType.Results.List) > 0 {
		returnValues := make([]string, 0, len(funcType.Results.List))
		for _, result := range funcType.Results.List {
			returnValues = append(returnValues, getZeroValue(result.Type))
		}
		builder.WriteString(fmt.Sprintf("\treturn %s\n", strings.Join(returnValues, ", ")))
	}

	builder.WriteString("}\n\n")

	return builder.String()
}

func generateCommentUpdateMethod(structName, methodName string, interfaceMethod MethodInfo, existingMethod *ast.FuncDecl) string {
	var builder strings.Builder

	builder.WriteString(fmt.Sprintf("// TODO: Update %s", methodName))
	if interfaceMethod.IsEmbedded {
		builder.WriteString(fmt.Sprintf(" (from %s)", interfaceMethod.SourceInterface))
	}
	builder.WriteString("\n")
	builder.WriteString(extractMethodComments(interfaceMethod.Field))
	builder.WriteString(fmt.Sprintf("func (s *%s) %s", structName, methodName))
	builder.WriteString(getFuncTypeString(existingMethod.Type))
	builder.WriteString(" {\n")
	builder.WriteString("\t// TODO: Update implementation\n")
	builder.WriteString("}\n\n")

	return builder.String()
}
