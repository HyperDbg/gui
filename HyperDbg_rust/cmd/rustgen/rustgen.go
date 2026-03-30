package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"maps"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

const (
	GeneratedDir = "rust-driver/kd/src/generated"
)

var enumTypes = []string{
	"MessageType", "BreakpointType", "DebugState", "LogLevel",
	"EventType", "ExceptionCode", "MemoryAccessType", "VmxExitReason",
	"DebuggerEventType", "HookType",
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

type GeneratedFile struct {
	Filename string
	Content  string
}

type FileMapping struct {
	GoFile   string
	RustFile string
	Types    []string
	Enums    []string
	Aliases  []string
}

var fileMappings = []FileMapping{
	{GoFile: "common.go", RustFile: "common.rs", Types: []string{"ProcessId", "ThreadId", "Address", "PhysicalAddress", "ProcessInfo", "ThreadInfo", "ModuleInfo", "BreakpointInfo", "MemoryRegion", "CallStackFrame", "SymbolInfo", "VmxCapabilities", "Instruction", "FunctionInfo"}},
	{GoFile: "register.go", RustFile: "register.rs", Types: []string{"RegisterState"}},
	{GoFile: "event.go", RustFile: "event.rs", Enums: []string{"EventType", "DebuggerEventType"}, Types: []string{"EventHeader", "DebuggerEvent"}},
	{GoFile: "event_breakpoint.go", RustFile: "event_breakpoint.rs", Enums: []string{"BreakpointType"}, Types: []string{"BreakpointEvent"}},
	{GoFile: "event_exception.go", RustFile: "event_exception.rs", Enums: []string{"ExceptionCode"}, Types: []string{"ExceptionEvent"}},
	{GoFile: "event_memory.go", RustFile: "event_memory.rs", Enums: []string{"MemoryAccessType"}, Types: []string{"MemoryAccessEvent"}},
	{GoFile: "event_syscall.go", RustFile: "event_syscall.rs", Types: []string{"SyscallEvent"}},
	{GoFile: "event_process.go", RustFile: "event_process.rs", Types: []string{"ProcessEvent"}},
	{GoFile: "event_thread.go", RustFile: "event_thread.rs", Types: []string{"ThreadEvent"}},
	{GoFile: "event_module.go", RustFile: "event_module.rs", Types: []string{"ModuleEvent"}},
	{GoFile: "event_debug_print.go", RustFile: "event_debug_print.rs", Enums: []string{"LogLevel"}, Types: []string{"DebugPrintEvent"}},
	{GoFile: "event_vmx.go", RustFile: "event_vmx.rs", Enums: []string{"VmxExitReason"}, Types: []string{"VmxExitEvent"}},
	{GoFile: "event_trap.go", RustFile: "event_trap.rs", Types: []string{"TrapEvent"}},
	{GoFile: "event_hidden_hook.go", RustFile: "event_hidden_hook.rs", Types: []string{"HiddenHookEvent"}},
	{GoFile: "event_cpuid.go", RustFile: "event_cpuid.rs", Types: []string{"CpuidEvent"}},
	{GoFile: "event_tsc.go", RustFile: "event_tsc.rs", Types: []string{"TscEvent"}},
	{GoFile: "event_cr_access.go", RustFile: "event_cr_access.rs", Types: []string{"CrAccessEvent"}},
	{GoFile: "event_dr_access.go", RustFile: "event_dr_access.rs", Types: []string{"DrAccessEvent"}},
	{GoFile: "event_io_port.go", RustFile: "event_io_port.rs", Types: []string{"IoPortEvent"}},
	{GoFile: "event_msr.go", RustFile: "event_msr.rs", Types: []string{"MsrEvent"}},
	{GoFile: "hook_types.go", RustFile: "hook_types.rs", Types: []string{"HookFilter"}},
	{GoFile: "hook_db.go", RustFile: "hook_db.rs", Enums: []string{"HookType"}, Types: []string{"HookInfo", "HookParam"}},
	{GoFile: "hook_script_compiler.go", RustFile: "hook_script_compiler.rs", Enums: []string{"HookOpType"}, Types: []string{"HookOperation"}},
	{GoFile: "event_ept_violation.go", RustFile: "event_ept_violation.rs", Types: []string{"EptViolationEvent"}},
}

func GenerateRustTypes(projectRoot string) error {
	outputDir := filepath.Join(projectRoot, GeneratedDir)
	os.RemoveAll(outputDir)
	if err := os.MkdirAll(outputDir, 0o755); err != nil {
		return fmt.Errorf("failed to create output dir: %w", err)
	}

	allFiles := parseAllGoFiles(projectRoot)

	apiMethods := extractAPIMethods(allFiles["interfaces.go"], allFiles["packet.go"])

	generateTypesSplit(projectRoot, allFiles)
	generateIoctlCodes(projectRoot, allFiles)
	generateHandlersSplit(projectRoot, apiMethods)

	generateMod(projectRoot, fileMappings)

	return nil
}

func parseAllGoFiles(projectRoot string) map[string]*ast.File {
	files := make(map[string]*ast.File)

	goFiles, _ := filepath.Glob(filepath.Join(projectRoot, "debugger", "*.go"))
	for _, path := range goFiles {
		name := filepath.Base(path)
		files[name] = parseFile(path)
	}

	return files
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

	if interfacesFile == nil || packetFile == nil {
		return methods
	}

	debuggerMethods := extractDebuggerInterfaceMethods(interfacesFile)
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

func extractDebuggerInterfaceMethods(file *ast.File) []MethodInfo {
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

		var findJSONMarshal func(n ast.Node) *ast.CallExpr
		findJSONMarshal = func(n ast.Node) *ast.CallExpr {
			switch node := n.(type) {
			case *ast.CallExpr:
				if isJSONMarshalCall(node) {
					return node
				}
				for _, arg := range node.Args {
					if found := findJSONMarshal(arg); found != nil {
						return found
					}
				}
			case *ast.AssignStmt:
				for _, expr := range node.Rhs {
					if found := findJSONMarshal(expr); found != nil {
						return found
					}
				}
			}
			return nil
		}

		ast.Inspect(fn, func(n ast.Node) bool {
			if call := findJSONMarshal(n); call != nil {
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

	var findLiteral func(expr ast.Expr) *ast.CompositeLit
	findLiteral = func(expr ast.Expr) *ast.CompositeLit {
		switch e := expr.(type) {
		case *ast.CompositeLit:
			return e
		case *ast.CallExpr:
			if isJSONMarshalCall(e) {
				if len(e.Args) > 0 {
					return findLiteral(e.Args[0])
				}
			}
			for _, arg := range e.Args {
				if lit := findLiteral(arg); lit != nil {
					return lit
				}
			}
		}
		return nil
	}

	return findLiteral(call.Args[0])
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
		"Start", "Stop", "IsConnected", "GetState",
		"RegisterCallback", "GetEvent", "WaitForEvent", "GetConnectedDrivers",
		"WaitForDriver", "ExecuteScript", "ExecuteScriptWithContext",
	}
	return !slices.Contains(nonAPIMethods, name)
}

func getNonAPIMethodComment(name string) string {
	comments := map[string]string{
		"Start":                    "// Connection management - handled by driver lifecycle",
		"Stop":                     "// Connection management - handled by driver lifecycle",
		"IsConnected":              "// Connection management - not applicable in kernel",
		"GetState":                 "// State management - handled separately",
		"RegisterCallback":         "// Event handling - use EventQueue + emit_* instead",
		"GetEvent":                 "// Event handling - use EventQueue + emit_* instead",
		"WaitForEvent":             "// Event handling - use EventQueue + emit_* instead",
		"GetConnectedDrivers":      "// Driver management - user-space only",
		"WaitForDriver":            "// Driver management - user-space only",
		"ExecuteScript":            "// Script execution - user-space only",
		"ExecuteScriptWithContext": "// Script execution - user-space only",
	}
	return comments[name]
}

func extractReturnType(m MethodInfo) string {
	if m.Results == nil || len(m.Results.List) == 0 {
		return "Empty"
	}

	for _, field := range m.Results.List {
		typeStr := exprToString(field.Type)

		if after, ok := strings.CutPrefix(typeStr, "[]"); ok {
			inner := after
			if inner == "byte" {
				return "Vec<u8>"
			}
			return "Vec<" + goTypeToRust(inner) + ">"
		}

		if after, ok := strings.CutPrefix(typeStr, "*"); ok {
			inner := after
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

func generateTypesSplit(projectRoot string, allFiles map[string]*ast.File) {
	typesDir := filepath.Join(projectRoot, GeneratedDir)
	if err := os.MkdirAll(typesDir, 0o755); err != nil {
		panic(err)
	}

	allConstants := make(map[string][]ConstantDef)
	allEnums := make(map[string]EnumDef)
	allStructs := make(map[string]StructDef)
	allAliases := make(map[string]TypeAliasDef)

	for _, node := range allFiles {
		constants, enums, structs, aliases := extractAllFromNode(node)
		maps.Copy(allConstants, constants)
		maps.Copy(allEnums, enums)
		maps.Copy(allStructs, structs)
		maps.Copy(allAliases, aliases)
	}

	enumValues := extractEnumValuesAll(allConstants, allEnums)

	for _, mapping := range fileMappings {
		var buf bytes.Buffer
		buf.WriteString("// Auto-generated by rustgen - DO NOT EDIT\n")
		buf.WriteString("// Source: " + mapping.GoFile + "\n\n")
		buf.WriteString("#![allow(non_snake_case)]\n")
		buf.WriteString("#![allow(unused_imports)]\n\n")
		buf.WriteString("extern crate alloc;\n\n")
		buf.WriteString("use alloc::string::String;\n")
		buf.WriteString("use alloc::vec::Vec;\n")
		buf.WriteString("use serde::{Serialize, Deserialize};\n\n")
		buf.WriteString("use super::*;\n\n")

		hasContent := false

		for _, enumName := range mapping.Enums {
			if e, ok := allEnums[enumName]; ok {
				hasContent = true
				writeEnum(&buf, e, enumValues[e.Name])
			}
		}

		for _, typeName := range mapping.Types {
			if strings.Contains(typeName, "Id") || strings.Contains(typeName, "Address") {
				if a, ok := allAliases[typeName]; ok {
					hasContent = true
					writeTypeAlias(&buf, a)
				}
			}
			if s, ok := allStructs[typeName]; ok {
				hasContent = true
				writeStruct(&buf, s)
			}
		}

		if hasContent {
			outputPath := filepath.Join(typesDir, mapping.RustFile)
			if err := os.WriteFile(outputPath, buf.Bytes(), 0o644); err != nil {
				panic(err)
			}
			fmt.Printf("Generated %s\n", outputPath)
		}
	}

	generateResponseTypes(projectRoot)
}

func extractAllFromNode(node *ast.File) (map[string][]ConstantDef, map[string]EnumDef, map[string]StructDef, map[string]TypeAliasDef) {
	constants := make(map[string][]ConstantDef)
	enums := make(map[string]EnumDef)
	structs := make(map[string]StructDef)
	aliases := make(map[string]TypeAliasDef)

	for _, decl := range node.Decls {
		switch d := decl.(type) {
		case *ast.GenDecl:
			switch d.Tok {
			case token.CONST:
				var currentEnum string
				var iotaCounter int
				var hasIota bool
				var iotaOffset int
				for _, spec := range d.Specs {
					vs := spec.(*ast.ValueSpec)
					if len(vs.Names) > 0 && vs.Type != nil {
						typeStr := exprToString(vs.Type)
						if slices.Contains(enumTypes, typeStr) {
							currentEnum = typeStr
							iotaCounter = 0
							hasIota = false
							iotaOffset = 0
						}
					}
					for i, name := range vs.Names {
						val := ""
						if i < len(vs.Values) {
							val = exprToString(vs.Values[i])
						}
						if strings.HasPrefix(val, "iota") {
							hasIota = true
							if strings.Contains(val, "+") {
								parts := strings.Split(val, "+")
								if len(parts) == 2 {
									if n, err := strconv.Atoi(strings.TrimSpace(parts[1])); err == nil {
										iotaOffset = n
									}
								}
							} else if strings.Contains(val, "-") {
								parts := strings.Split(val, "-")
								if len(parts) == 2 {
									if n, err := strconv.Atoi(strings.TrimSpace(parts[1])); err == nil {
										iotaOffset = -n
									}
								}
							}
						}
						if hasIota && (val == "" || strings.HasPrefix(val, "iota")) {
							val = fmt.Sprintf("%d", iotaCounter+iotaOffset)
						}
						if name.Name != "_" && currentEnum != "" {
							constants[currentEnum] = append(constants[currentEnum], ConstantDef{
								Name:  name.Name,
								Value: val,
							})
						}
						if hasIota {
							iotaCounter++
						}
					}
				}
			case token.TYPE:
				for _, spec := range d.Specs {
					ts := spec.(*ast.TypeSpec)
					switch t := ts.Type.(type) {
					case *ast.Ident:
						if slices.Contains(enumTypes, ts.Name.Name) {
							enums[ts.Name.Name] = EnumDef{
								Name: ts.Name.Name,
								Type: t.Name,
							}
						} else if !ts.Assign.IsValid() {
							enums[ts.Name.Name] = EnumDef{
								Name: ts.Name.Name,
								Type: t.Name,
							}
						} else {
							aliases[ts.Name.Name] = TypeAliasDef{
								Name:    ts.Name.Name,
								Aliased: t.Name,
							}
						}
					case *ast.SelectorExpr:
						if ts.Assign.IsValid() {
							aliases[ts.Name.Name] = TypeAliasDef{
								Name:    ts.Name.Name,
								Aliased: t.Sel.Name,
							}
						}
					case *ast.StructType:
						structs[ts.Name.Name] = StructDef{
							Name:   ts.Name.Name,
							Fields: extractFields(t),
						}
					}
				}
			}
		}
	}

	return constants, enums, structs, aliases
}

func extractEnumValuesAll(allConstants map[string][]ConstantDef, allEnums map[string]EnumDef) map[string][]EnumValue {
	result := make(map[string][]EnumValue)
	for enumName := range allEnums {
		prefix := getEnumPrefix(enumName)
		for _, consts := range allConstants {
			for _, c := range consts {
				if strings.HasPrefix(c.Name, prefix) {
					result[enumName] = append(result[enumName], EnumValue{
						Name:  c.Name,
						Value: c.Value,
					})
				}
			}
		}
	}
	return result
}

func generateMod(projectRoot string, mappings []FileMapping) {
	var buf bytes.Buffer
	buf.WriteString("// Auto-generated by rustgen - DO NOT EDIT\n\n")

	for _, m := range mappings {
		buf.WriteString(fmt.Sprintf("mod %s;\n", strings.TrimSuffix(m.RustFile, ".rs")))
	}
	buf.WriteString("mod response;\n")
	buf.WriteString("mod router;\n")
	buf.WriteString("mod emit;\n")
	buf.WriteString("mod ioctl;\n")
	buf.WriteString("mod ntddk;\n")
	buf.WriteString("mod types;\n")
	buf.WriteString("mod constants;\n\n")

	buf.WriteString("// Re-export all types\n")
	buf.WriteString("pub use common::*;\n")
	buf.WriteString("pub use register::*;\n")
	buf.WriteString("pub use event::*;\n")
	buf.WriteString("pub use event_breakpoint::*;\n")
	buf.WriteString("pub use event_exception::*;\n")
	buf.WriteString("pub use event_memory::*;\n")
	buf.WriteString("pub use event_syscall::*;\n")
	buf.WriteString("pub use event_process::*;\n")
	buf.WriteString("pub use event_thread::*;\n")
	buf.WriteString("pub use event_module::*;\n")
	buf.WriteString("pub use event_debug_print::*;\n")
	buf.WriteString("pub use event_vmx::*;\n")
	buf.WriteString("pub use event_trap::*;\n")
	buf.WriteString("pub use event_hidden_hook::*;\n")
	buf.WriteString("pub use event_cpuid::*;\n")
	buf.WriteString("pub use event_tsc::*;\n")
	buf.WriteString("pub use event_cr_access::*;\n")
	buf.WriteString("pub use event_dr_access::*;\n")
	buf.WriteString("pub use event_io_port::*;\n")
	buf.WriteString("pub use event_msr::*;\n")
	buf.WriteString("pub use event_ept_violation::*;\n")
	buf.WriteString("pub use response::*;\n")
	buf.WriteString("pub use router::{Request, Response, Empty, DebuggerApi, dispatch_api, NoOpDebugger};\n")
	buf.WriteString("pub use emit::{EventQueue, emit_event};\n")
	buf.WriteString("pub use emit::*;\n")
	buf.WriteString("pub use ioctl::*;\n")
	buf.WriteString("pub use ntddk::*;\n")
	buf.WriteString("pub use types::*;\n")
	buf.WriteString("pub use constants::*;\n")

	outputPath := filepath.Join(projectRoot, GeneratedDir, "mod.rs")
	if err := os.WriteFile(outputPath, buf.Bytes(), 0o644); err != nil {
		panic(err)
	}
	fmt.Printf("Generated %s\n", outputPath)
}

func generateResponseTypes(projectRoot string) {
	var buf bytes.Buffer
	buf.WriteString("// Auto-generated by rustgen - DO NOT EDIT\n\n")
	buf.WriteString("#![allow(unused_imports)]\n\n")
	buf.WriteString("extern crate alloc;\n\n")
	buf.WriteString("use alloc::string::String;\n")
	buf.WriteString("use serde::{Deserialize, Serialize};\n\n")
	buf.WriteString("use super::*;\n\n")

	buf.WriteString(writeResponseStruct())

	outputPath := filepath.Join(projectRoot, GeneratedDir, "response.rs")
	if err := os.WriteFile(outputPath, buf.Bytes(), 0o644); err != nil {
		panic(err)
	}
	fmt.Printf("Generated %s\n", outputPath)
}

func generateIoctlCodes(projectRoot string, allFiles map[string]*ast.File) {
	ioctlFile := allFiles["ioctl.go"]
	if ioctlFile == nil {
		return
	}

	ioctlCodes := extractIoctlConstants(ioctlFile)
	if len(ioctlCodes) == 0 {
		return
	}

	var buf bytes.Buffer
	buf.WriteString("// Auto-generated by rustgen - DO NOT EDIT\n")
	buf.WriteString("// Source: ioctl.go\n\n")
	buf.WriteString("#![allow(non_snake_case)]\n")
	buf.WriteString("#![allow(dead_code)]\n\n")

	buf.WriteString("/// Re-export WDK constants for IOCTL (from wdk_sys)\n")
	buf.WriteString("pub use wdk_sys::{FILE_DEVICE_UNKNOWN, METHOD_BUFFERED, FILE_ANY_ACCESS};\n\n")

	buf.WriteString("/// CTL_CODE macro: (DeviceType << 16) | (Access << 14) | (Function << 2) | Method\n")
	buf.WriteString("pub const fn CTL_CODE(device_type: u32, function: u32, method: u32, access: u32) -> u32 {\n")
	buf.WriteString("    (device_type << 16) | (access << 14) | (function << 2) | method\n")
	buf.WriteString("}\n\n")

	buf.WriteString("/// IOCTL function codes (0x800 base)\n")
	for _, code := range ioctlCodes {
		buf.WriteString(fmt.Sprintf("pub const %s_FUNCTION: u32 = 0x%X;\n", code.Name, code.FunctionCode))
	}
	buf.WriteString("\n")

	buf.WriteString("/// IOCTL control codes for HyperDbg driver communication\n")
	buf.WriteString("/// Calculated using CTL_CODE(FILE_DEVICE_UNKNOWN, function, METHOD_BUFFERED, FILE_ANY_ACCESS)\n")
	for _, code := range ioctlCodes {
		buf.WriteString(fmt.Sprintf("pub const %s: u32 = CTL_CODE(FILE_DEVICE_UNKNOWN, %s_FUNCTION, METHOD_BUFFERED, FILE_ANY_ACCESS);\n", code.Name, code.Name))
	}
	buf.WriteString("\n")

	buf.WriteString("/// IOCTL code names for display\n")
	buf.WriteString("pub fn ioctl_name(code: u32) -> Option<&'static str> {\n")
	buf.WriteString("    match code {\n")
	for _, code := range ioctlCodes {
		buf.WriteString(fmt.Sprintf("        %s => Some(\"%s\"),\n", code.Name, code.Name))
	}
	buf.WriteString("        _ => None,\n")
	buf.WriteString("    }\n")
	buf.WriteString("}\n")

	outputPath := filepath.Join(projectRoot, GeneratedDir, "ioctl.rs")
	if err := os.WriteFile(outputPath, buf.Bytes(), 0o644); err != nil {
		panic(err)
	}
	fmt.Printf("Generated %s\n", outputPath)
}

type IoctlConst struct {
	Name         string
	FunctionCode uint32
	Value        uint32
}

func extractIoctlConstants(file *ast.File) []IoctlConst {
	var codes []IoctlConst

	const (
		FILE_DEVICE_UNKNOWN = 0x00000022
		FILE_ANY_ACCESS     = 0
		METHOD_BUFFERED     = 0
	)

	ctlCode := func(function uint32) uint32 {
		return (FILE_DEVICE_UNKNOWN << 16) | (FILE_ANY_ACCESS << 14) | (function << 2) | METHOD_BUFFERED
	}

	for _, decl := range file.Decls {
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok || genDecl.Tok != token.CONST {
			continue
		}

		var functionValue uint32
		var foundIota bool
		iotaIndex := 0

		for _, spec := range genDecl.Specs {
			vs, ok := spec.(*ast.ValueSpec)
			if !ok {
				continue
			}

			for i, name := range vs.Names {
				if !strings.HasPrefix(name.Name, "IOCTL_") {
					continue
				}

				var value uint32
				if i < len(vs.Values) {
					valStr := exprToString(vs.Values[i])
					if strings.Contains(valStr, "iota") {
						foundIota = true
						if strings.Contains(valStr, "+") {
							parts := strings.Split(valStr, "+")
							if len(parts) == 2 {
								if n, err := strconv.ParseUint(strings.TrimSpace(parts[1]), 0, 32); err == nil {
									functionValue = uint32(n)
								}
							}
						} else {
							functionValue = 0
						}
						value = ctlCode(functionValue + uint32(iotaIndex))
						iotaIndex++
					} else if foundIota {
						value = ctlCode(functionValue + uint32(iotaIndex))
						iotaIndex++
					} else if n, err := strconv.ParseUint(valStr, 0, 32); err == nil {
						value = uint32(n)
					}
				} else if foundIota {
					value = ctlCode(functionValue + uint32(iotaIndex))
					iotaIndex++
				}

				codes = append(codes, IoctlConst{
					Name:         name.Name,
					FunctionCode: functionValue + uint32(iotaIndex-1),
					Value:        value,
				})
			}
		}
	}

	return codes
}

func generateHandlersSplit(projectRoot string, apiMethods []APIMethod) {
	generateHandlersRouter(projectRoot, apiMethods)
	generateHandlersEmit(projectRoot)
}

func generateHandlersRouter(projectRoot string, apiMethods []APIMethod) {
	var buf bytes.Buffer

	buf.WriteString("// Auto-generated by rustgen - DO NOT EDIT\n")
	buf.WriteString("// Source: interfaces.go + packet.go (auto-scanned)\n\n")
	buf.WriteString("#![allow(non_snake_case)]\n\n")
	buf.WriteString("extern crate alloc;\n\n")
	buf.WriteString("use alloc::string::String;\n")
	buf.WriteString("use alloc::vec::Vec;\n")
	buf.WriteString("use alloc::format;\n")
	buf.WriteString("use serde::{Serialize, Deserialize};\n\n")
	buf.WriteString("use crate::common::types::*;\n\n")

	buf.WriteString("// Request structure for API calls\n")
	buf.WriteString("#[derive(Serialize, Deserialize, Debug)]\n")
	buf.WriteString("pub struct Request {\n")
	buf.WriteString("    pub action: String,\n")
	buf.WriteString("    #[serde(skip_serializing_if = \"Option::is_none\")]\n")
	buf.WriteString("    pub process_id: Option<u32>,\n")
	buf.WriteString("    #[serde(skip_serializing_if = \"Option::is_none\")]\n")
	buf.WriteString("    pub address: Option<u64>,\n")
	buf.WriteString("    #[serde(skip_serializing_if = \"Option::is_none\")]\n")
	buf.WriteString("    pub breakpoint_id: Option<u64>,\n")
	buf.WriteString("    #[serde(skip_serializing_if = \"Option::is_none\")]\n")
	buf.WriteString("    pub size: Option<u32>,\n")
	buf.WriteString("    #[serde(skip_serializing_if = \"Option::is_none\")]\n")
	buf.WriteString("    pub r#type: Option<i32>,\n")
	buf.WriteString("    #[serde(skip_serializing_if = \"Option::is_none\")]\n")
	buf.WriteString("    pub bp_type: Option<u32>,\n")
	buf.WriteString("    #[serde(skip_serializing_if = \"Option::is_none\")]\n")
	buf.WriteString("    pub data: Option<Vec<u8>>,\n")
	buf.WriteString("    #[serde(skip_serializing_if = \"Option::is_none\")]\n")
	buf.WriteString("    pub regs: Option<RegisterState>,\n")
	buf.WriteString("    #[serde(skip_serializing_if = \"Option::is_none\")]\n")
	buf.WriteString("    pub api_name: Option<String>,\n")
	buf.WriteString("    #[serde(skip_serializing_if = \"Option::is_none\")]\n")
	buf.WriteString("    pub hook_type: Option<u32>,\n")
	buf.WriteString("    #[serde(skip_serializing_if = \"Option::is_none\")]\n")
	buf.WriteString("    pub filter: Option<HookFilter>,\n")
	buf.WriteString("    #[serde(skip_serializing_if = \"Option::is_none\")]\n")
	buf.WriteString("    pub code: Option<String>,\n")
	buf.WriteString("}\n\n")

	buf.WriteString("// Response structure for API calls\n")
	buf.WriteString("#[derive(Serialize, Deserialize, Debug)]\n")
	buf.WriteString("pub struct Response<T: Serialize> {\n")
	buf.WriteString("    pub success: bool,\n")
	buf.WriteString("    pub message: String,\n")
	buf.WriteString("    #[serde(skip_serializing_if = \"Option::is_none\")]\n")
	buf.WriteString("    pub data: Option<T>,\n")
	buf.WriteString("}\n\n")

	buf.WriteString("// Empty type for responses without data\n")
	buf.WriteString("#[derive(Serialize, Deserialize, Debug, Clone, Default)]\n")
	buf.WriteString("pub struct Empty {}\n\n")

	buf.WriteString("// API Handler trait for implementing debugger operations\n")
	buf.WriteString("pub trait DebuggerApi {\n")
	for _, m := range apiMethods {
		buf.WriteString(fmt.Sprintf("    fn %s(&mut self, req: &Request) -> Result<%s, String>;\n", rustSafeName(toSnakeCase(m.Name)), m.ReturnType))
	}
	buf.WriteString("}\n\n")

	buf.WriteString("// Non-API methods (excluded from trait):\n")
	buf.WriteString("// - Start, Stop, IsConnected: Connection management - handled by driver lifecycle\n")
	buf.WriteString("// - GetState, Ping, Status: State/health checks - handled separately\n")
	buf.WriteString("// - RegisterCallback, GetEvent, WaitForEvent: Event handling - use EventQueue + emit_* instead\n")
	buf.WriteString("// - GetConnectedDrivers, WaitForDriver: Driver management - user-space only\n")
	buf.WriteString("// - ExecuteScript, ExecuteScriptWithContext: Script execution - user-space only\n\n")

	buf.WriteString("// Generic API dispatcher - takes raw JSON bytes and returns JSON response\n")
	buf.WriteString("pub fn dispatch_api<T: DebuggerApi>(api: &mut T, body: &[u8]) -> Vec<u8> {\n")
	buf.WriteString("    let req: Result<Request, _> = serde_json::from_slice(body);\n")
	buf.WriteString("    \n")
	buf.WriteString("    match req {\n")
	buf.WriteString("        Ok(req) => {\n")
	buf.WriteString("            match req.action.as_str() {\n")
	for _, m := range apiMethods {
		buf.WriteString(fmt.Sprintf("                \"%s\" => {\n", m.Action))
		buf.WriteString(fmt.Sprintf("                    match api.%s(&req) {\n", rustSafeName(toSnakeCase(m.Name))))
		buf.WriteString("                        Ok(data) => success_response(data),\n")
		buf.WriteString("                        Err(e) => error_response(&e),\n")
		buf.WriteString("                    }\n")
		buf.WriteString("                }\n")
	}
	buf.WriteString("                _ => {\n")
	buf.WriteString("                    error_response(&format!(\"Unknown action: {}\", req.action))\n")
	buf.WriteString("                }\n")
	buf.WriteString("            }\n")
	buf.WriteString("        }\n")
	buf.WriteString("        Err(e) => {\n")
	buf.WriteString("            error_response(&format!(\"Invalid request: {:?}\", e))\n")
	buf.WriteString("        }\n")
	buf.WriteString("    }\n")
	buf.WriteString("}\n\n")

	buf.WriteString("// Response helpers\n")
	buf.WriteString("fn success_response<T: serde::Serialize>(data: T) -> Vec<u8> {\n")
	buf.WriteString("    let resp = Response {\n")
	buf.WriteString("        success: true,\n")
	buf.WriteString("        message: String::from(\"OK\"),\n")
	buf.WriteString("        data: Some(data),\n")
	buf.WriteString("    };\n")
	buf.WriteString("    serde_json::to_vec(&resp).unwrap_or_default()\n")
	buf.WriteString("}\n\n")

	buf.WriteString("fn error_response(msg: &str) -> Vec<u8> {\n")
	buf.WriteString("    let resp = Response::<()> {\n")
	buf.WriteString("        success: false,\n")
	buf.WriteString("        message: String::from(msg),\n")
	buf.WriteString("        data: None,\n")
	buf.WriteString("    };\n")
	buf.WriteString("    serde_json::to_vec(&resp).unwrap_or_default()\n")
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

	buf.WriteString("// Default test implementation with sample data\n")
	buf.WriteString("pub struct NoOpDebugger;\n\n")

	buf.WriteString("impl DebuggerApi for NoOpDebugger {\n")
	for _, m := range apiMethods {
		buf.WriteString(fmt.Sprintf("    fn %s(&mut self, _req: &Request) -> Result<%s, String> {\n", rustSafeName(toSnakeCase(m.Name)), m.ReturnType))
		switch m.ReturnType {
		case "Empty":
			buf.WriteString("        Ok(Empty {})\n")
		case "Vec<u8>":
			buf.WriteString("        Ok(alloc::vec![0x90, 0x90, 0x90, 0x90])\n")
		case "RegisterState":
			buf.WriteString("        Ok(RegisterState {\n")
			buf.WriteString("            rax: 0x12345678,\n")
			buf.WriteString("            rbx: 0x87654321,\n")
			buf.WriteString("            rcx: 0x11111111,\n")
			buf.WriteString("            rdx: 0x22222222,\n")
			buf.WriteString("            rsi: 0x33333333,\n")
			buf.WriteString("            rdi: 0x44444444,\n")
			buf.WriteString("            rbp: 0x55555555,\n")
			buf.WriteString("            rsp: 0x7FFE0000,\n")
			buf.WriteString("            r8: 0x66666666,\n")
			buf.WriteString("            r9: 0x77777777,\n")
			buf.WriteString("            r10: 0x88888888,\n")
			buf.WriteString("            r11: 0x99999999,\n")
			buf.WriteString("            r12: 0xAAAAAAAA,\n")
			buf.WriteString("            r13: 0xBBBBBBBB,\n")
			buf.WriteString("            r14: 0xCCCCCCCC,\n")
			buf.WriteString("            r15: 0xDDDDDDDD,\n")
			buf.WriteString("            rip: 0x7FFE1234,\n")
			buf.WriteString("            rflags: 0x246,\n")
			buf.WriteString("            cr0: 0x80000000,\n")
			buf.WriteString("            cr2: 0,\n")
			buf.WriteString("            cr3: 0x1A2B3C4D,\n")
			buf.WriteString("            cr4: 0x406F8,\n")
			buf.WriteString("            dr0: 0,\n")
			buf.WriteString("            dr1: 0,\n")
			buf.WriteString("            dr2: 0,\n")
			buf.WriteString("            dr3: 0,\n")
			buf.WriteString("            dr6: 0xFFFF0FF0,\n")
			buf.WriteString("            dr7: 0x400,\n")
			buf.WriteString("            gdtr: 0xFFFFF80000000000,\n")
			buf.WriteString("            gsbase: 0xFFFFF80012340000,\n")
			buf.WriteString("            fsbase: 0,\n")
			buf.WriteString("        })\n")
		case "Vec<ProcessInfo>":
			buf.WriteString("        Ok(alloc::vec![\n")
			buf.WriteString("            ProcessInfo { process_id: 4, image_name: Some(String::from(\"System\")), base_address: Some(String::from(\"0xFFFFF80000000000\")), size: 0x100000, cr3: 0x1A000 },\n")
			buf.WriteString("            ProcessInfo { process_id: 1234, image_name: Some(String::from(\"notepad.exe\")), base_address: Some(String::from(\"0x7FFE00000000\")), size: 0x2000000, cr3: 0x2B000 },\n")
			buf.WriteString("        ])\n")
		case "Vec<ThreadInfo>":
			buf.WriteString("        Ok(alloc::vec![\n")
			buf.WriteString("            ThreadInfo { thread_id: 5678, process_id: 1234, start_address: Some(String::from(\"0x7FFE12340000\")), teb: Some(String::from(\"0x7FFE56780000\")), state: 2 },\n")
			buf.WriteString("        ])\n")
		case "Vec<ModuleInfo>":
			buf.WriteString("        Ok(alloc::vec![\n")
			buf.WriteString("            ModuleInfo { base_address: Some(String::from(\"0x7FFE00000000\")), size: 0x1000000, name: Some(String::from(\"ntdll.dll\")), path: Some(String::from(\"C:\\\\Windows\\\\System32\\\\ntdll.dll\")) },\n")
			buf.WriteString("            ModuleInfo { base_address: Some(String::from(\"0x7FFE10000000\")), size: 0x2000000, name: Some(String::from(\"kernel32.dll\")), path: Some(String::from(\"C:\\\\Windows\\\\System32\\\\kernel32.dll\")) },\n")
			buf.WriteString("        ])\n")
		default:
			if strings.HasPrefix(m.ReturnType, "Vec<") {
				buf.WriteString("        Ok(Vec::new())\n")
			} else {
				buf.WriteString("        Ok(Default::default())\n")
			}
		}
		buf.WriteString("    }\n\n")
	}
	buf.WriteString("}\n")

	outputPath := filepath.Join(projectRoot, GeneratedDir, "router.rs")
	if err := os.WriteFile(outputPath, buf.Bytes(), 0o644); err != nil {
		panic(err)
	}
	fmt.Printf("Generated %s\n", outputPath)
}

func generateHandlersEmit(projectRoot string) {
	emitCode := `// Auto-generated by rustgen - DO NOT EDIT

extern crate alloc;

use alloc::string::String;
use alloc::vec::Vec;
use alloc::format;
use crate::common::types::*;

fn addr_to_string(addr: u64) -> Option<String> {
    Some(format!("0x{:016X}", addr))
}

fn args_to_option(args: [u64; 8]) -> Option<Vec<u64>> {
    Some(args.to_vec())
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
        address: addr_to_string(address),
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
        address: addr_to_string(address),
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
        virtual_address: addr_to_string(virtual_address),
        physical_address: addr_to_string(physical_address),
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
        arguments: args_to_option(arguments),
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
        message: Some(message),
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
        guest_rip: addr_to_string(guest_rip),
        guest_rsp: addr_to_string(guest_rsp),
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
        address: addr_to_string(address),
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
    let event_type = match hook_type {
        MemoryAccessType::Read => EventType::HiddenHookRead,
        MemoryAccessType::Write => EventType::HiddenHookWrite,
        _ => EventType::HiddenHookExec,
    };
    emit_event(DebuggerEvent::HiddenHook(HiddenHookEvent {
        header: EventHeader {
            event_type,
            process_id,
            thread_id,
            core_id,
            timestamp: 0,
        },
        hook_address: addr_to_string(hook_address),
        hook_type,
        data: Some(data),
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
    let event_type = if is_write { EventType::MsrWrite } else { EventType::MsrRead };
    emit_event(DebuggerEvent::Msr(MsrEvent {
        header: EventHeader {
            event_type,
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
        guest_physical_address: addr_to_string(guest_physical_address),
        guest_virtual_address: addr_to_string(guest_virtual_address),
        read,
        write,
        execute,
        readable,
        writable,
        executable,
    }));
}
`

	outputPath := filepath.Join(projectRoot, GeneratedDir, "emit.rs")
	if err := os.WriteFile(outputPath, []byte(emitCode), 0o644); err != nil {
		panic(err)
	}
	fmt.Printf("Generated %s\n", outputPath)
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
	case *ast.BinaryExpr:
		return exprToString(e.X) + " " + e.Op.String() + " " + exprToString(e.Y)
	default:
		return fmt.Sprintf("%T", e)
	}
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
	case "HookType":
		return "HookType"
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
	firstValue := ""
	for i, v := range values {
		name := v.Name
		if strings.HasPrefix(name, prefix) {
			name = name[len(prefix):]
		}
		val := v.Value
		if val == "" || val == "iota" {
			val = fmt.Sprintf("%d", i)
		}
		if i == 0 {
			firstValue = name
		}
		buf.WriteString(fmt.Sprintf("    %s = %s,\n", name, val))
	}
	buf.WriteString("}\n\n")

	if firstValue != "" {
		buf.WriteString(fmt.Sprintf("impl Default for %s {\n", e.Name))
		buf.WriteString(fmt.Sprintf("    fn default() -> Self {\n"))
		buf.WriteString(fmt.Sprintf("        %s::%s\n", e.Name, firstValue))
		buf.WriteString("    }\n")
		buf.WriteString("}\n\n")
	}
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
		if f.Format == "hex" || strings.Contains(f.Name, "Address") || f.Type == "Address" || f.Type == "PhysicalAddress" {
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

pub fn event_type_from_name(name: &str) -> Option<EventType> {
    match name {
        "Breakpoint" => Some(EventType::Breakpoint),
        "Exception" => Some(EventType::Exception),
        "MemoryAccess" => Some(EventType::MemoryAccess),
        "SyscallEntry" => Some(EventType::SyscallEntry),
        "SyscallExit" => Some(EventType::SyscallExit),
        "ProcessCreate" => Some(EventType::ProcessCreate),
        "ProcessExit" => Some(EventType::ProcessExit),
        "ThreadCreate" => Some(EventType::ThreadCreate),
        "ThreadExit" => Some(EventType::ThreadExit),
        "ModuleLoad" => Some(EventType::ModuleLoad),
        "ModuleUnload" => Some(EventType::ModuleUnload),
        "DebugPrint" => Some(EventType::DebugPrint),
        "VmxExit" => Some(EventType::VmxExit),
        "Trap" => Some(EventType::Trap),
        "HiddenHookExec" => Some(EventType::HiddenHookExec),
        "HiddenHookRead" => Some(EventType::HiddenHookRead),
        "HiddenHookWrite" => Some(EventType::HiddenHookWrite),
        "Cpuid" => Some(EventType::Cpuid),
        "Tsc" => Some(EventType::Tsc),
        "Pmc" => Some(EventType::Pmc),
        "Interrupt" => Some(EventType::Interrupt),
        "ExceptionBitmap" => Some(EventType::ExceptionBitmap),
        "CrAccess" => Some(EventType::CrAccess),
        "DrAccess" => Some(EventType::DrAccess),
        "IoPort" => Some(EventType::IoPort),
        "MsrRead" => Some(EventType::MsrRead),
        "MsrWrite" => Some(EventType::MsrWrite),
        "EptViolation" => Some(EventType::EptViolation),
        "Vmcalled" => Some(EventType::Vmcalled),
        _ => None,
    }
}

pub fn event_type_name(t: EventType) -> &'static str {
    match t {
        EventType::Breakpoint => "Breakpoint",
        EventType::Exception => "Exception",
        EventType::MemoryAccess => "MemoryAccess",
        EventType::SyscallEntry => "SyscallEntry",
        EventType::SyscallExit => "SyscallExit",
        EventType::ProcessCreate => "ProcessCreate",
        EventType::ProcessExit => "ProcessExit",
        EventType::ThreadCreate => "ThreadCreate",
        EventType::ThreadExit => "ThreadExit",
        EventType::ModuleLoad => "ModuleLoad",
        EventType::ModuleUnload => "ModuleUnload",
        EventType::DebugPrint => "DebugPrint",
        EventType::VmxExit => "VmxExit",
        EventType::Trap => "Trap",
        EventType::HiddenHookExec => "HiddenHookExec",
        EventType::HiddenHookRead => "HiddenHookRead",
        EventType::HiddenHookWrite => "HiddenHookWrite",
        EventType::Cpuid => "Cpuid",
        EventType::Tsc => "Tsc",
        EventType::Pmc => "Pmc",
        EventType::Interrupt => "Interrupt",
        EventType::ExceptionBitmap => "ExceptionBitmap",
        EventType::CrAccess => "CrAccess",
        EventType::DrAccess => "DrAccess",
        EventType::IoPort => "IoPort",
        EventType::MsrRead => "MsrRead",
        EventType::MsrWrite => "MsrWrite",
        EventType::EptViolation => "EptViolation",
        EventType::Vmcalled => "Vmcalled",
    }
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

func rustSafeName(name string) string {
	keywords := map[string]bool{
		"as": true, "break": true, "const": true, "continue": true, "crate": true,
		"else": true, "enum": true, "extern": true, "false": true, "fn": true,
		"for": true, "if": true, "impl": true, "in": true, "let": true, "loop": true,
		"match": true, "mod": true, "move": true, "mut": true, "pub": true, "ref": true,
		"return": true, "self": true, "Self": true, "static": true, "struct": true,
		"super": true, "trait": true, "true": true, "type": true, "unsafe": true,
		"use": true, "where": true, "while": true, "async": true, "await": true,
		"dyn": true,
	}
	if keywords[name] {
		return name + "_"
	}
	return name
}
