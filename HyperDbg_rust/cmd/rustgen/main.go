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

var enumTypes = []string{
	"MessageType", "BreakpointType", "DebugState", "LogLevel",
	"EventType", "ExceptionCode", "MemoryAccessType", "VmxExitReason",
	"DebuggerEventType",
}

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
	var typeAliases []TypeAliasDef

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
						} else if !ts.Assign.IsValid() {
							enums = append(enums, EnumDef{
								Name: ts.Name.Name,
								Type: t.Name,
							})
						} else {
							typeAliases = append(typeAliases, TypeAliasDef{
								Name:    ts.Name.Name,
								Aliased: t.Name,
							})
						}
					case *ast.SelectorExpr:
						if ts.Assign.IsValid() {
							typeAliases = append(typeAliases, TypeAliasDef{
								Name:    ts.Name.Name,
								Aliased: t.Sel.Name,
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

	for _, a := range typeAliases {
		writeTypeAlias(&buf, a)
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

type TypeAliasDef struct {
	Name    string
	Aliased string
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
	case "DebuggerEventType":
		return "DebuggerEvent"
	}
	return enumName
}

type EnumValue struct {
	Name  string
	Value string
}

func writeTypeAlias(buf *bytes.Buffer, a TypeAliasDef) {
	rustType := goTypeToRust(a.Aliased)
	buf.WriteString(fmt.Sprintf("pub type %s = %s;\n", a.Name, rustType))
}

func writeEnum(buf *bytes.Buffer, e EnumDef, values []EnumValue) {
	rustType := goTypeToRust(e.Type)
	buf.WriteString(fmt.Sprintf("#[derive(Debug, Clone, Copy, PartialEq, Eq, Serialize, Deserialize)]\n"))
	buf.WriteString(fmt.Sprintf("#[repr(%s)]\n", rustType))
	buf.WriteString(fmt.Sprintf("pub enum %s {\n", e.Name))
	prefix := getEnumPrefix(e.Name)
	for i, v := range values {
		name := v.Name
		if strings.HasPrefix(name, prefix) {
			name = name[len(prefix):]
		}
		val := v.Value
		if val == "" || val == "iota" {
			val = fmt.Sprintf("%d", i)
		}
		buf.WriteString(fmt.Sprintf("    %s = %s,\n", name, val))
	}
	buf.WriteString("}\n\n")
}

func writeStruct(buf *bytes.Buffer, s StructDef) {
	if s.Name == "Empty" || s.Name == "ResponseType" || s.Name == "EventCallback" || s.Name == "Response" || s.Name == "DebuggerEvent" {
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

#[derive(Clone, Debug, Serialize, Deserialize)]
pub enum DebuggerEvent {
    Breakpoint(BreakpointEvent),
    Exception(ExceptionEvent),
    MemoryAccess(MemoryAccessEvent),
    Syscall(SyscallEvent),
    ProcessCreate(ProcessEvent),
    ProcessExit(ProcessEvent),
    ThreadCreate(ThreadEvent),
    ThreadExit(ThreadEvent),
    ModuleLoad(ModuleEvent),
    ModuleUnload(ModuleEvent),
    DebugPrint(DebugPrintEvent),
    VmxExit(VmxExitEvent),
    Trap(TrapEvent),
    HiddenHook(HiddenHookEvent),
    Cpuid(CpuidEvent),
    Tsc(TscEvent),
    CrAccess(CrAccessEvent),
    DrAccess(DrAccessEvent),
    IoPort(IoPortEvent),
    Msr(MsrEvent),
    EptViolation(EptViolationEvent),
}

pub fn debugger_event_type_from_name(name: &str) -> Option<DebuggerEventType> {
    match name {
        "Breakpoint" => Some(DebuggerEventType::Breakpoint),
        "Exception" => Some(DebuggerEventType::Exception),
        "MemoryAccess" => Some(DebuggerEventType::MemoryAccess),
        "Syscall" => Some(DebuggerEventType::Syscall),
        "ProcessCreate" => Some(DebuggerEventType::ProcessCreate),
        "ProcessExit" => Some(DebuggerEventType::ProcessExit),
        "ThreadCreate" => Some(DebuggerEventType::ThreadCreate),
        "ThreadExit" => Some(DebuggerEventType::ThreadExit),
        "ModuleLoad" => Some(DebuggerEventType::ModuleLoad),
        "ModuleUnload" => Some(DebuggerEventType::ModuleUnload),
        "DebugPrint" => Some(DebuggerEventType::DebugPrint),
        "VmxExit" => Some(DebuggerEventType::VmxExit),
        "Trap" => Some(DebuggerEventType::Trap),
        "HiddenHook" => Some(DebuggerEventType::HiddenHook),
        "Cpuid" => Some(DebuggerEventType::Cpuid),
        "Tsc" => Some(DebuggerEventType::Tsc),
        "CrAccess" => Some(DebuggerEventType::CrAccess),
        "DrAccess" => Some(DebuggerEventType::DrAccess),
        "IoPort" => Some(DebuggerEventType::IoPort),
        "Msr" => Some(DebuggerEventType::Msr),
        "EptViolation" => Some(DebuggerEventType::EptViolation),
        _ => None,
    }
}

pub fn debugger_event_type_name(t: DebuggerEventType) -> &'static str {
    match t {
        DebuggerEventType::Breakpoint => "Breakpoint",
        DebuggerEventType::Exception => "Exception",
        DebuggerEventType::MemoryAccess => "MemoryAccess",
        DebuggerEventType::Syscall => "Syscall",
        DebuggerEventType::ProcessCreate => "ProcessCreate",
        DebuggerEventType::ProcessExit => "ProcessExit",
        DebuggerEventType::ThreadCreate => "ThreadCreate",
        DebuggerEventType::ThreadExit => "ThreadExit",
        DebuggerEventType::ModuleLoad => "ModuleLoad",
        DebuggerEventType::ModuleUnload => "ModuleUnload",
        DebuggerEventType::DebugPrint => "DebugPrint",
        DebuggerEventType::VmxExit => "VmxExit",
        DebuggerEventType::Trap => "Trap",
        DebuggerEventType::HiddenHook => "HiddenHook",
        DebuggerEventType::Cpuid => "Cpuid",
        DebuggerEventType::Tsc => "Tsc",
        DebuggerEventType::CrAccess => "CrAccess",
        DebuggerEventType::DrAccess => "DrAccess",
        DebuggerEventType::IoPort => "IoPort",
        DebuggerEventType::Msr => "Msr",
        DebuggerEventType::EptViolation => "EptViolation",
    }
}

pub unsafe fn emit_event(event: DebuggerEvent) {
    extern "C" {
        static mut EVENT_QUEUE: *mut EventQueue;
    }
    if !EVENT_QUEUE.is_null() {
        (*EVENT_QUEUE).push(event);
    }
}

pub struct EventQueue {
    events: Vec<DebuggerEvent>,
    max_size: usize,
}

impl EventQueue {
    pub fn new(max_size: usize) -> Self {
        Self {
            events: Vec::new(),
            max_size,
        }
    }

    pub fn push(&mut self, event: DebuggerEvent) {
        if self.events.len() >= self.max_size {
            self.events.remove(0);
        }
        self.events.push(event);
    }

    pub fn pop(&mut self) -> Option<DebuggerEvent> {
        self.events.pop()
    }

    pub fn len(&self) -> usize {
        self.events.len()
    }

    pub fn is_empty(&self) -> bool {
        self.events.is_empty()
    }
}

pub unsafe fn emit_breakpoint_event(
    process_id: ProcessId,
    thread_id: ThreadId,
    core_id: u32,
    address: Address,
    breakpoint_id: u64,
    registers: &RegisterState,
) {
    emit_event(DebuggerEvent::Breakpoint(BreakpointEvent {
        header: EventHeader {
            event_type: EventType::Breakpoint,
            process_id,
            thread_id,
            core_id,
            timestamp: 0,
        },
        address,
        breakpoint_id,
        registers: registers.clone(),
    }));
}

pub unsafe fn emit_exception_event(
    process_id: ProcessId,
    thread_id: ThreadId,
    core_id: u32,
    exception_code: ExceptionCode,
    address: Address,
    error_code: u64,
    registers: &RegisterState,
) {
    emit_event(DebuggerEvent::Exception(ExceptionEvent {
        header: EventHeader {
            event_type: EventType::Exception,
            process_id,
            thread_id,
            core_id,
            timestamp: 0,
        },
        exception_code,
        address,
        error_code,
        registers: registers.clone(),
    }));
}

pub unsafe fn emit_memory_access_event(
    process_id: ProcessId,
    thread_id: ThreadId,
    core_id: u32,
    virtual_address: Address,
    physical_address: PhysicalAddress,
    access_type: MemoryAccessType,
    size: u32,
    value: u64,
) {
    emit_event(DebuggerEvent::MemoryAccess(MemoryAccessEvent {
        header: EventHeader {
            event_type: EventType::MemoryAccess,
            process_id,
            thread_id,
            core_id,
            timestamp: 0,
        },
        virtual_address,
        physical_address,
        access_type,
        size,
        value,
    }));
}

pub unsafe fn emit_syscall_event(
    process_id: ProcessId,
    thread_id: ThreadId,
    core_id: u32,
    syscall_number: u32,
    arguments: [u64; 8],
    is_entry: bool,
) {
    emit_event(DebuggerEvent::Syscall(SyscallEvent {
        header: EventHeader {
            event_type: if is_entry { EventType::SyscallEntry } else { EventType::SyscallExit },
            process_id,
            thread_id,
            core_id,
            timestamp: 0,
        },
        syscall_number,
        arguments,
        return_value: 0,
        is_entry,
    }));
}

pub unsafe fn emit_process_event(
    process_id: ProcessId,
    thread_id: ThreadId,
    core_id: u32,
    process_info: ProcessInfo,
    parent_process_id: ProcessId,
    is_create: bool,
) {
    let event = if is_create {
        DebuggerEvent::ProcessCreate(ProcessEvent {
            header: EventHeader {
                event_type: EventType::ProcessCreate,
                process_id,
                thread_id,
                core_id,
                timestamp: 0,
            },
            process_info,
            parent_process_id,
            is_create,
        })
    } else {
        DebuggerEvent::ProcessExit(ProcessEvent {
            header: EventHeader {
                event_type: EventType::ProcessExit,
                process_id,
                thread_id,
                core_id,
                timestamp: 0,
            },
            process_info,
            parent_process_id,
            is_create,
        })
    };
    emit_event(event);
}

pub unsafe fn emit_thread_event(
    process_id: ProcessId,
    thread_id: ThreadId,
    core_id: u32,
    thread_info: ThreadInfo,
    is_create: bool,
) {
    let event = if is_create {
        DebuggerEvent::ThreadCreate(ThreadEvent {
            header: EventHeader {
                event_type: EventType::ThreadCreate,
                process_id,
                thread_id,
                core_id,
                timestamp: 0,
            },
            thread_info,
            is_create,
        })
    } else {
        DebuggerEvent::ThreadExit(ThreadEvent {
            header: EventHeader {
                event_type: EventType::ThreadExit,
                process_id,
                thread_id,
                core_id,
                timestamp: 0,
            },
            thread_info,
            is_create,
        })
    };
    emit_event(event);
}

pub unsafe fn emit_module_event(
    process_id: ProcessId,
    thread_id: ThreadId,
    core_id: u32,
    module_info: ModuleInfo,
    is_load: bool,
) {
    let event = if is_load {
        DebuggerEvent::ModuleLoad(ModuleEvent {
            header: EventHeader {
                event_type: EventType::ModuleLoad,
                process_id,
                thread_id,
                core_id,
                timestamp: 0,
            },
            module_info,
            is_load,
        })
    } else {
        DebuggerEvent::ModuleUnload(ModuleEvent {
            header: EventHeader {
                event_type: EventType::ModuleUnload,
                process_id,
                thread_id,
                core_id,
                timestamp: 0,
            },
            module_info,
            is_load,
        })
    };
    emit_event(event);
}

pub unsafe fn emit_debug_print_event(
    process_id: ProcessId,
    thread_id: ThreadId,
    core_id: u32,
    message: String,
    level: LogLevel,
) {
    emit_event(DebuggerEvent::DebugPrint(DebugPrintEvent {
        header: EventHeader {
            event_type: EventType::DebugPrint,
            process_id,
            thread_id,
            core_id,
            timestamp: 0,
        },
        message,
        level,
    }));
}

pub unsafe fn emit_vmx_exit_event(
    process_id: ProcessId,
    thread_id: ThreadId,
    core_id: u32,
    exit_reason: VmxExitReason,
    exit_qualification: u64,
    guest_rip: Address,
    guest_rsp: Address,
    instruction_length: u32,
) {
    emit_event(DebuggerEvent::VmxExit(VmxExitEvent {
        header: EventHeader {
            event_type: EventType::VmxExit,
            process_id,
            thread_id,
            core_id,
            timestamp: 0,
        },
        exit_reason,
        exit_qualification,
        guest_rip,
        guest_rsp,
        instruction_length,
    }));
}

pub unsafe fn emit_trap_event(
    process_id: ProcessId,
    thread_id: ThreadId,
    core_id: u32,
    trap_number: u32,
    error_code: u64,
    address: Address,
) {
    emit_event(DebuggerEvent::Trap(TrapEvent {
        header: EventHeader {
            event_type: EventType::Trap,
            process_id,
            thread_id,
            core_id,
            timestamp: 0,
        },
        trap_number,
        error_code,
        address,
    }));
}

pub unsafe fn emit_hidden_hook_event(
    process_id: ProcessId,
    thread_id: ThreadId,
    core_id: u32,
    hook_address: Address,
    hook_type: MemoryAccessType,
    data: Vec<u8>,
) {
    emit_event(DebuggerEvent::HiddenHook(HiddenHookEvent {
        header: EventHeader {
            event_type: EventType::HiddenHook,
            process_id,
            thread_id,
            core_id,
            timestamp: 0,
        },
        hook_address,
        hook_type,
        data,
    }));
}

pub unsafe fn emit_cpuid_event(
    process_id: ProcessId,
    thread_id: ThreadId,
    core_id: u32,
    leaf: u32,
    sub_leaf: u32,
    eax: u32,
    ebx: u32,
    ecx: u32,
    edx: u32,
) {
    emit_event(DebuggerEvent::Cpuid(CpuidEvent {
        header: EventHeader {
            event_type: EventType::Cpuid,
            process_id,
            thread_id,
            core_id,
            timestamp: 0,
        },
        leaf,
        sub_leaf,
        eax,
        ebx,
        ecx,
        edx,
    }));
}

pub unsafe fn emit_tsc_event(
    process_id: ProcessId,
    thread_id: ThreadId,
    core_id: u32,
    tsc_value: u64,
    rdtsc_exit: bool,
) {
    emit_event(DebuggerEvent::Tsc(TscEvent {
        header: EventHeader {
            event_type: EventType::Tsc,
            process_id,
            thread_id,
            core_id,
            timestamp: 0,
        },
        tsc_value,
        rdtsc_exit,
    }));
}

pub unsafe fn emit_cr_access_event(
    process_id: ProcessId,
    thread_id: ThreadId,
    core_id: u32,
    cr_number: u32,
    is_write: bool,
    old_value: u64,
    new_value: u64,
) {
    emit_event(DebuggerEvent::CrAccess(CrAccessEvent {
        header: EventHeader {
            event_type: EventType::CrAccess,
            process_id,
            thread_id,
            core_id,
            timestamp: 0,
        },
        cr_number,
        is_write,
        old_value,
        new_value,
    }));
}

pub unsafe fn emit_dr_access_event(
    process_id: ProcessId,
    thread_id: ThreadId,
    core_id: u32,
    dr_number: u32,
    is_write: bool,
    value: u64,
) {
    emit_event(DebuggerEvent::DrAccess(DrAccessEvent {
        header: EventHeader {
            event_type: EventType::DrAccess,
            process_id,
            thread_id,
            core_id,
            timestamp: 0,
        },
        dr_number,
        is_write,
        value,
    }));
}

pub unsafe fn emit_io_port_event(
    process_id: ProcessId,
    thread_id: ThreadId,
    core_id: u32,
    port: u16,
    size: u32,
    is_write: bool,
    value: u32,
) {
    emit_event(DebuggerEvent::IoPort(IoPortEvent {
        header: EventHeader {
            event_type: EventType::IoPort,
            process_id,
            thread_id,
            core_id,
            timestamp: 0,
        },
        port,
        size,
        is_write,
        value,
    }));
}

pub unsafe fn emit_msr_event(
    process_id: ProcessId,
    thread_id: ThreadId,
    core_id: u32,
    msr: u32,
    is_write: bool,
    value: u64,
) {
    emit_event(DebuggerEvent::Msr(MsrEvent {
        header: EventHeader {
            event_type: EventType::Msr,
            process_id,
            thread_id,
            core_id,
            timestamp: 0,
        },
        msr,
        is_write,
        value,
    }));
}

pub unsafe fn emit_ept_violation_event(
    process_id: ProcessId,
    thread_id: ThreadId,
    core_id: u32,
    guest_physical_address: PhysicalAddress,
    guest_virtual_address: Address,
    read: bool,
    write: bool,
    execute: bool,
    readable: bool,
    writable: bool,
    executable: bool,
) {
    emit_event(DebuggerEvent::EptViolation(EptViolationEvent {
        header: EventHeader {
            event_type: EventType::EptViolation,
            process_id,
            thread_id,
            core_id,
            timestamp: 0,
        },
        guest_physical_address,
        guest_virtual_address,
        read,
        write,
        execute,
        readable,
        writable,
        executable,
    }));
}
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
