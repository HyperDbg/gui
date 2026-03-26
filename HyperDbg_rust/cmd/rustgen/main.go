package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"slices"
	"strings"
)

func getProjectRoot() string {
	cwd, _ := os.Getwd()
	if filepath.Base(cwd) == "rustgen" {
		return filepath.Dir(filepath.Dir(cwd))
	}
	return cwd
}

const (
	RustOutputDir      = "rust-driver"
	TypesOutputFile    = "types_gen.rs"
	HandlersOutputFile = "handlers_gen.rs"
)

var enumTypes = []string{"MessageType", "BreakpointType", "DebugState"}

type APIMethod struct {
	Action      string
	Name        string
	Params      []APIParam
	ReturnType  string
	Description string
}

type APIParam struct {
	Name   string
	Type   string
	Format string
}

func main() {
	projectRoot := getProjectRoot()

	typesPath := filepath.Join(projectRoot, "types.go")
	interfacesPath := filepath.Join(projectRoot, "interfaces.go")
	packetPath := filepath.Join(projectRoot, "packet.go")

	typesFile := parseFile(typesPath)
	interfacesFile := parseFile(interfacesPath)
	packetFile := parseFile(packetPath)

	apiMethods := extractAPIMethods(interfacesFile, packetFile)

	generateModels(projectRoot, typesFile)
	generateHandlers(projectRoot, apiMethods)
}

func parseFile(path string) *ast.File {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
	if err != nil {
		panic(fmt.Errorf("failed to parse %s: %w", path, err))
	}
	return node
}

func extractAPIMethods(interfacesFile, packetFile *ast.File) []APIMethod {
	var methods []APIMethod

	debuggerMethods := extractInterfaceMethods(interfacesFile)
	actionMap := extractActionMap(packetFile)

	for _, m := range debuggerMethods {
		if !isAPIMethod(m.Name) {
			continue
		}

		action := actionMap[m.Name]
		if action == "" {
			action = toSnakeCase(m.Name)
		}

		returnType := extractReturnType(m)
		params := extractParams(m)

		methods = append(methods, APIMethod{
			Name:       m.Name,
			Action:     action,
			Params:     params,
			ReturnType: returnType,
		})
	}

	return methods
}

type MethodInfo struct {
	Name    string
	Params  []*ast.Field
	Results *ast.FieldList
}

func extractInterfaceMethods(file *ast.File) []MethodInfo {
	var methods []MethodInfo

	for _, decl := range file.Decls {
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok || genDecl.Tok != token.TYPE {
			continue
		}

		for _, spec := range genDecl.Specs {
			ts, ok := spec.(*ast.TypeSpec)
			if !ok || ts.Name.Name != "Debugger" {
				continue
			}

			iface, ok := ts.Type.(*ast.InterfaceType)
			if !ok {
				continue
			}

			for _, method := range iface.Methods.List {
				if len(method.Names) == 0 {
					continue
				}

				ft, ok := method.Type.(*ast.FuncType)
				if !ok {
					continue
				}

				methods = append(methods, MethodInfo{
					Name:    method.Names[0].Name,
					Params:  ft.Params.List,
					Results: ft.Results,
				})
			}
		}
	}

	return methods
}

func extractActionMap(file *ast.File) map[string]string {
	actionMap := make(map[string]string)

	for _, decl := range file.Decls {
		fn, ok := decl.(*ast.FuncDecl)
		if !ok || fn.Recv == nil || len(fn.Recv.List) == 0 {
			continue
		}

		methodName := fn.Name.Name

		ast.Inspect(fn, func(n ast.Node) bool {
			call, ok := n.(*ast.CallExpr)
			if !ok {
				return true
			}

			if isJSONMarshalCall(call) {
				lit := findMapLiteral(call)
				if lit != nil {
					action := extractActionFromMap(lit)
					if action != "" {
						actionMap[methodName] = action
					}
				}
			}

			return true
		})
	}

	return actionMap
}

func isJSONMarshalCall(call *ast.CallExpr) bool {
	sel, ok := call.Fun.(*ast.SelectorExpr)
	if !ok {
		return false
	}
	return sel.Sel.Name == "Marshal"
}

func findMapLiteral(call *ast.CallExpr) *ast.CompositeLit {
	if len(call.Args) < 1 {
		return nil
	}

	if call.Args[0] != nil {
		if check, ok := call.Args[0].(*ast.CallExpr); ok {
			if len(check.Args) > 0 {
				if lit, ok := check.Args[0].(*ast.CompositeLit); ok {
					return lit
				}
			}
		}

		if lit, ok := call.Args[0].(*ast.CompositeLit); ok {
			return lit
		}
	}

	return nil
}

func extractActionFromMap(lit *ast.CompositeLit) string {
	for _, elt := range lit.Elts {
		kv, ok := elt.(*ast.KeyValueExpr)
		if !ok {
			continue
		}

		key, ok := kv.Key.(*ast.BasicLit)
		if !ok || key.Value != `"action"` {
			continue
		}

		switch v := kv.Value.(type) {
		case *ast.BasicLit:
			return strings.Trim(v.Value, `"`)
		}
	}

	return ""
}

func isAPIMethod(name string) bool {
	nonAPIMethods := []string{
		"Start", "Stop", "IsConnected", "GetState", "Ping", "Status",
		"RegisterCallback", "GetEvent", "WaitForEvent", "GetConnectedDrivers",
		"WaitForDriver", "ExecuteScript", "ExecuteScriptWithContext",
	}
	return !slices.Contains(nonAPIMethods, name)
}

func extractReturnType(m MethodInfo) string {
	if m.Results == nil || len(m.Results.List) == 0 {
		return "Empty"
	}

	for _, field := range m.Results.List {
		typeStr := exprToString(field.Type)

		if strings.HasPrefix(typeStr, "[]") {
			inner := strings.TrimPrefix(typeStr, "[]")
			if inner == "byte" {
				return "Vec<u8>"
			}
			return "Vec<" + goTypeToRust(inner) + ">"
		}

		if strings.HasPrefix(typeStr, "*") {
			inner := strings.TrimPrefix(typeStr, "*")
			return goTypeToRust(inner)
		}

		if typeStr == "error" {
			continue
		}

		return goTypeToRust(typeStr)
	}

	return "Empty"
}

func extractParams(m MethodInfo) []APIParam {
	var params []APIParam

	for _, field := range m.Params {
		typeStr := exprToString(field.Type)

		for _, name := range field.Names {
			format := "int"
			if strings.Contains(strings.ToLower(name.Name), "address") ||
				strings.Contains(strings.ToLower(name.Name), "id") ||
				strings.Contains(strings.ToLower(name.Name), "breakpoint") {
				if typeStr == "uint64" {
					format = "hex"
				}
			}
			if strings.Contains(strings.ToLower(name.Name), "process") ||
				strings.Contains(strings.ToLower(name.Name), "thread") ||
				strings.Contains(strings.ToLower(name.Name), "size") {
				format = "dec"
			}

			params = append(params, APIParam{
				Name:   toSnakeCase(name.Name),
				Type:   goTypeToRust(typeStr),
				Format: format,
			})
		}
	}

	return params
}

func generateModels(projectRoot string, node *ast.File) {
	var buf bytes.Buffer
	buf.WriteString("// Auto-generated by rustgen - DO NOT EDIT\n")
	buf.WriteString("// Source: types.go\n\n")
	buf.WriteString("#![allow(non_snake_case)]\n\n")
	buf.WriteString("extern crate alloc;\n\n")
	buf.WriteString("use alloc::string::String;\n")
	buf.WriteString("use alloc::vec::Vec;\n")
	buf.WriteString("use serde::{Deserialize, Serialize};\n\n")

	var constants []ConstantDef
	var enums []EnumDef
	var structs []StructDef

	for _, decl := range node.Decls {
		switch d := decl.(type) {
		case *ast.GenDecl:
			switch d.Tok {
			case token.CONST:
				for _, spec := range d.Specs {
					vs := spec.(*ast.ValueSpec)
					for i, name := range vs.Names {
						if name.Name == "_" {
							continue
						}
						val := ""
						if i < len(vs.Values) {
							val = exprToString(vs.Values[i])
						}
						constants = append(constants, ConstantDef{
							Name:  name.Name,
							Value: val,
						})
					}
				}
			case token.TYPE:
				for _, spec := range d.Specs {
					ts := spec.(*ast.TypeSpec)
					switch t := ts.Type.(type) {
					case *ast.Ident:
						if slices.Contains(enumTypes, ts.Name.Name) {
							enums = append(enums, EnumDef{
								Name: ts.Name.Name,
								Type: t.Name,
							})
						}
					case *ast.StructType:
						structs = append(structs, StructDef{
							Name:   ts.Name.Name,
							Fields: extractFields(t),
						})
					}
				}
			}
		}
	}

	enumValues := extractEnumValues(constants, enums)
	for _, e := range enums {
		writeEnum(&buf, e, enumValues[e.Name])
	}

	for _, s := range structs {
		writeStruct(&buf, s)
	}

	buf.WriteString(writeResponseStruct())

	outputPath := filepath.Join(projectRoot, RustOutputDir, TypesOutputFile)
	if err := os.MkdirAll(filepath.Dir(outputPath), 0755); err != nil {
		panic(err)
	}
	if err := os.WriteFile(outputPath, buf.Bytes(), 0644); err != nil {
		panic(err)
	}

	fmt.Printf("Generated %s\n", outputPath)
}

func generateHandlers(projectRoot string, apiMethods []APIMethod) {
	var buf bytes.Buffer

	buf.WriteString("// Auto-generated by rustgen - DO NOT EDIT\n")
	buf.WriteString("// Source: interfaces.go + packet.go (auto-scanned)\n\n")
	buf.WriteString("#![allow(non_snake_case)]\n\n")
	buf.WriteString("extern crate alloc;\n\n")
	buf.WriteString("use alloc::string::String;\n")
	buf.WriteString("use alloc::vec::Vec;\n")
	buf.WriteString("use alloc::format;\n\n")
	buf.WriteString("use crate::util::{Request, parse_hex_string, parse_dec_string};\n")
	buf.WriteString("use crate::types_gen::*;\n")
	buf.WriteString("use net::{ResponseWriter, Request as HttpRequest, log_info, log_error, log_success};\n\n")

	buf.WriteString("// API Handler trait for implementing debugger operations\n")
	buf.WriteString("pub trait DebuggerApi {\n")
	for _, m := range apiMethods {
		buf.WriteString(fmt.Sprintf("    fn %s(&mut self, req: &Request) -> Result<%s, String>;\n", toSnakeCase(m.Name), m.ReturnType))
	}
	buf.WriteString("}\n\n")

	buf.WriteString("// Generic API dispatcher\n")
	buf.WriteString("pub unsafe fn dispatch_api<T: DebuggerApi>(api: &mut T, w: *mut ResponseWriter, r: *mut HttpRequest) {\n")
	buf.WriteString("    let body = core::slice::from_raw_parts((*r).Body, (*r).BodyLength);\n")
	buf.WriteString("    let req: Result<Request, _> = serde_json::from_slice(body);\n")
	buf.WriteString("    \n")
	buf.WriteString("    match req {\n")
	buf.WriteString("        Ok(req) => {\n")
	buf.WriteString("            log_info!(\"API request: {}\", req.action);\n")
	buf.WriteString("            \n")
	buf.WriteString("            match req.action.as_str() {\n")
	for _, m := range apiMethods {
		buf.WriteString(fmt.Sprintf("                \"%s\" => {\n", m.Action))
		buf.WriteString(fmt.Sprintf("                    match api.%s(&req) {\n", toSnakeCase(m.Name)))
		buf.WriteString("                        Ok(data) => write_success_response(w, data),\n")
		buf.WriteString("                        Err(e) => write_error_response(w, &e),\n")
		buf.WriteString("                    }\n")
		buf.WriteString("                }\n")
	}
	buf.WriteString("                _ => {\n")
	buf.WriteString("                    write_error_response(w, &format!(\"Unknown action: {}\", req.action));\n")
	buf.WriteString("                }\n")
	buf.WriteString("            }\n")
	buf.WriteString("        }\n")
	buf.WriteString("        Err(e) => {\n")
	buf.WriteString("            write_error_response(w, &format!(\"Invalid request: {:?}\", e));\n")
	buf.WriteString("        }\n")
	buf.WriteString("    }\n")
	buf.WriteString("}\n\n")

	buf.WriteString("// Response helpers\n")
	buf.WriteString("unsafe fn write_success_response<T: serde::Serialize>(w: *mut ResponseWriter, data: T) {\n")
	buf.WriteString("    let resp = Response {\n")
	buf.WriteString("        success: true,\n")
	buf.WriteString("        message: String::from(\"OK\"),\n")
	buf.WriteString("        data: Some(data),\n")
	buf.WriteString("    };\n")
	buf.WriteString("    (*w).WriteJSON(&resp);\n")
	buf.WriteString("}\n\n")

	buf.WriteString("unsafe fn write_error_response(w: *mut ResponseWriter, msg: &str) {\n")
	buf.WriteString("    let resp = Response::<()> {\n")
	buf.WriteString("        success: false,\n")
	buf.WriteString("        message: String::from(msg),\n")
	buf.WriteString("        data: None,\n")
	buf.WriteString("    };\n")
	buf.WriteString("    (*w).WriteJSON(&resp);\n")
	buf.WriteString("}\n\n")

	buf.WriteString("// Request parameter helpers\n")
	buf.WriteString("impl Request {\n")
	buf.WriteString("    pub fn get_size(&self) -> Option<u32> {\n")
	buf.WriteString("        self.size\n")
	buf.WriteString("    }\n\n")
	buf.WriteString("    pub fn get_type(&self) -> Option<i32> {\n")
	buf.WriteString("        self.r#type\n")
	buf.WriteString("    }\n\n")
	buf.WriteString("    pub fn get_data(&self) -> Option<&Vec<u8>> {\n")
	buf.WriteString("        self.data.as_ref()\n")
	buf.WriteString("    }\n\n")
	buf.WriteString("    pub fn get_regs(&self) -> Option<&RegisterState> {\n")
	buf.WriteString("        self.regs.as_ref()\n")
	buf.WriteString("    }\n")
	buf.WriteString("}\n\n")

	buf.WriteString("// Default no-op implementation for testing\n")
	buf.WriteString("pub struct NoOpDebugger;\n\n")

	buf.WriteString("impl DebuggerApi for NoOpDebugger {\n")
	for _, m := range apiMethods {
		buf.WriteString(fmt.Sprintf("    fn %s(&mut self, _req: &Request) -> Result<%s, String> {\n", toSnakeCase(m.Name), m.ReturnType))
		buf.WriteString(fmt.Sprintf("        log_info!(\"%s (no-op)\");\n", m.Name))
		if m.ReturnType == "Empty" {
			buf.WriteString("        Ok(())\n")
		} else if m.ReturnType == "Vec<u8>" {
			buf.WriteString("        Ok(Vec::new())\n")
		} else if m.ReturnType == "RegisterState" {
			buf.WriteString("        Ok(RegisterState::default())\n")
		} else if strings.HasPrefix(m.ReturnType, "Vec<") {
			buf.WriteString("        Ok(Vec::new())\n")
		} else {
			buf.WriteString("        Ok(Default::default())\n")
		}
		buf.WriteString("    }\n\n")
	}
	buf.WriteString("}\n")

	outputPath := filepath.Join(projectRoot, RustOutputDir, HandlersOutputFile)
	if err := os.MkdirAll(filepath.Dir(outputPath), 0755); err != nil {
		panic(err)
	}
	if err := os.WriteFile(outputPath, buf.Bytes(), 0644); err != nil {
		panic(err)
	}

	fmt.Printf("Generated %s\n", outputPath)
	fmt.Printf("\nExtracted %d API methods:\n", len(apiMethods))
	for _, m := range apiMethods {
		fmt.Printf("  - %s -> %s\n", m.Name, m.Action)
	}
}

type ConstantDef struct {
	Name  string
	Value string
}

type EnumDef struct {
	Name string
	Type string
}

type StructDef struct {
	Name   string
	Fields []FieldDef
}

type FieldDef struct {
	Name    string
	Type    string
	JSONTag string
	Format  string
}

func extractFields(st *ast.StructType) []FieldDef {
	var fields []FieldDef
	for _, field := range st.Fields.List {
		for _, name := range field.Names {
			f := FieldDef{
				Name: name.Name,
				Type: exprToString(field.Type),
			}
			if field.Tag != nil {
				tag := field.Tag.Value
				tag = strings.Trim(tag, "`")
				if jsonTag := reflect.StructTag(tag).Get("json"); jsonTag != "" {
					f.JSONTag = strings.Split(jsonTag, ",")[0]
				}
				if format := reflect.StructTag(tag).Get("format"); format != "" {
					f.Format = format
				}
			}
			if f.JSONTag == "" {
				f.JSONTag = toSnakeCase(f.Name)
			}
			fields = append(fields, f)
		}
	}
	return fields
}

func exprToString(expr ast.Expr) string {
	switch e := expr.(type) {
	case *ast.Ident:
		return e.Name
	case *ast.BasicLit:
		return e.Value
	case *ast.SelectorExpr:
		return exprToString(e.X) + "." + e.Sel.Name
	case *ast.CallExpr:
		return exprToString(e.Fun) + "(...)"
	case *ast.ArrayType:
		return "[]" + exprToString(e.Elt)
	case *ast.StarExpr:
		return "*" + exprToString(e.X)
	case *ast.MapType:
		return "map[" + exprToString(e.Key) + "]" + exprToString(e.Value)
	case *ast.InterfaceType:
		return "interface{}"
	default:
		return fmt.Sprintf("%T", e)
	}
}

func extractEnumValues(constants []ConstantDef, enums []EnumDef) map[string][]EnumValue {
	result := make(map[string][]EnumValue)
	for _, e := range enums {
		prefix := getEnumPrefix(e.Name)
		for _, c := range constants {
			if strings.HasPrefix(c.Name, prefix) {
				result[e.Name] = append(result[e.Name], EnumValue{
					Name:  c.Name,
					Value: c.Value,
				})
			}
		}
	}
	return result
}

func getEnumPrefix(enumName string) string {
	switch enumName {
	case "MessageType":
		return "MsgType"
	case "BreakpointType":
		return "Breakpoint"
	case "DebugState":
		return "State"
	}
	return enumName
}

type EnumValue struct {
	Name  string
	Value string
}

func writeEnum(buf *bytes.Buffer, e EnumDef, values []EnumValue) {
	rustType := goTypeToRust(e.Type)
	buf.WriteString(fmt.Sprintf("#[derive(Debug, Clone, Copy, PartialEq, Eq, Serialize, Deserialize)]\n"))
	buf.WriteString(fmt.Sprintf("#[repr(%s)]\n", rustType))
	buf.WriteString(fmt.Sprintf("pub enum %s {\n", e.Name))
	prefix := getEnumPrefix(e.Name)
	for _, v := range values {
		name := v.Name
		if strings.HasPrefix(name, prefix) {
			name = name[len(prefix):]
		}
		buf.WriteString(fmt.Sprintf("    %s = %s,\n", name, v.Value))
	}
	buf.WriteString("}\n\n")
}

func writeStruct(buf *bytes.Buffer, s StructDef) {
	if s.Name == "Empty" || s.Name == "ResponseType" || s.Name == "EventCallback" || s.Name == "Response" {
		return
	}
	if strings.HasPrefix(s.Name, "Message") {
		return
	}

	buf.WriteString(fmt.Sprintf("#[derive(Serialize, Deserialize, Debug, Clone, Default)]\n"))
	buf.WriteString(fmt.Sprintf("pub struct %s {\n", s.Name))
	for _, f := range s.Fields {
		rustType := goTypeToRust(f.Type)
		fieldName := toSnakeCase(f.Name)
		if fieldName == "type" {
			fieldName = "bp_type"
		}
		if f.Format == "hex" || strings.Contains(f.Name, "Address") {
			rustType = "String"
		}
		buf.WriteString(fmt.Sprintf("    #[serde(rename = \"%s\")]\n", f.JSONTag))
		if rustType == "String" || strings.HasPrefix(rustType, "Vec") || strings.HasPrefix(rustType, "Option") {
			buf.WriteString(fmt.Sprintf("    #[serde(skip_serializing_if = \"Option::is_none\")]\n"))
			buf.WriteString(fmt.Sprintf("    pub %s: Option<%s>,\n", fieldName, rustType))
		} else {
			buf.WriteString(fmt.Sprintf("    pub %s: %s,\n", fieldName, rustType))
		}
	}
	buf.WriteString("}\n\n")
}

func writeResponseStruct() string {
	return `#[derive(Serialize, Deserialize, Debug)]
pub struct Response<T: Serialize> {
    pub success: bool,
    pub message: String,
    #[serde(skip_serializing_if = "Option::is_none")]
    pub data: Option<T>,
}

pub type EmptyResponse = Response<()>;
pub type Empty = ();
`
}

func goTypeToRust(goType string) string {
	switch goType {
	case "uint8":
		return "u8"
	case "uint16":
		return "u16"
	case "uint32":
		return "u32"
	case "uint64":
		return "u64"
	case "int8":
		return "i8"
	case "int16":
		return "i16"
	case "int32":
		return "i32"
	case "int64":
		return "i64"
	case "int":
		return "i32"
	case "uint":
		return "u32"
	case "uintptr":
		return "usize"
	case "float32":
		return "f32"
	case "float64":
		return "f64"
	case "string":
		return "String"
	case "bool":
		return "bool"
	case "[]byte":
		return "Vec<u8>"
	case "MessageType":
		return "MessageType"
	case "BreakpointType":
		return "BreakpointType"
	case "DebugState":
		return "DebugState"
	default:
		if strings.HasPrefix(goType, "[]") {
			inner := goType[2:]
			return "Vec<" + goTypeToRust(inner) + ">"
		}
		if strings.HasPrefix(goType, "*") {
			inner := goType[1:]
			return "Option<" + goTypeToRust(inner) + ">"
		}
		if strings.HasPrefix(goType, "map[") {
			return "std::collections::HashMap<String, serde_json::Value>"
		}
		return goType
	}
}

var snakeCaseRegex = regexp.MustCompile("([a-z0-9])([A-Z])")

func toSnakeCase(s string) string {
	snake := snakeCaseRegex.ReplaceAllString(s, "${1}_${2}")
	return strings.ToLower(snake)
}
