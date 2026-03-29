package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

type FunctionSignature struct {
	Name       string
	Params     string
	ReturnType string
	Line       int
	RawLine    string
}

type TypeDefinition struct {
	Name string
	Kind string
	Line int
}

type ConstantDefinition struct {
	Name  string
	Value string
	Line  int
}

type ExternFunction struct {
	Name       string
	File       string
	Line       int
	Params     string
	ReturnType string
	FoundInWdk bool
}

type TypeMismatch struct {
	FunctionName  string
	RustFile      string
	RustLine      int
	RustSignature string
	WdkSignature  string
	WdkLine       int
	MismatchType  string
}

type WdkBindings struct {
	Functions  map[string]FunctionSignature
	Types      map[string]TypeDefinition
	Constants  map[string]ConstantDefinition
	SourceFile string
}

type RustExterns struct {
	Functions map[string][]ExternFunction
	Files     []string
}

type ValidationReport struct {
	Mismatches     []TypeMismatch
	NotExported    []string
	MissingFromWdk []string
	Statistics     map[string]int
}

type WdkUsage struct {
	File         string
	Line         int
	Kind         string
	Name         string
	SourceType   string
	NeedsFix     bool
	SuggestedFix string
}

type FixReport struct {
	FilesScanned   int
	TotalUsages    int
	NeedsFix       int
	Usages         []WdkUsage
	FilesWithFixes map[string][]WdkUsage
}

type IUsageScanner interface {
	ScanProjectUsage(projectRoot string, excludeDirs []string) error
	GetFixReport() *FixReport
	GenerateUsageReport(outputPath string) error
	ApplyFixes(dryRun bool) error
}

type BindgenConfig struct {
	ProjectRoot    string
	WdkBindingsDir string
	RustSourceDir  string
	OutputDir      string
	Verbose        bool
}

type IWdkParser interface {
	ParseNtddkFunctions(path string) error
	ParseTypes(path string) error
	ParseConstants(path string) error
	GetFunctions() map[string]FunctionSignature
	GetTypes() map[string]TypeDefinition
	GetConstants() map[string]ConstantDefinition
}

type IRustScanner interface {
	ScanExterns(rustSourceDir string) error
	GetExternFunctions() map[string][]ExternFunction
	GetScannedFiles() []string
}

type IValidator interface {
	Validate() *ValidationReport
	CompareSignatures(name string, rustFunc ExternFunction, wdkSig FunctionSignature) *TypeMismatch
}

type IReportGenerator interface {
	GenerateFunctionsList(outputPath string) error
	GenerateFunctionsGoMap(outputPath string) error
	GenerateTypesGoMap(outputPath string) error
	GenerateConstantsGoMap(outputPath string) error
	GenerateValidationReport(outputPath string) error
	GenerateFixSuggestions(outputPath string) error
}

type IBindgen interface {
	IWdkParser
	IRustScanner
	IValidator
	IReportGenerator
	LoadWdkBindings(ntddkPath, typesPath, constantsPath string) error
	GetWdkBindings() *WdkBindings
	GetRustExterns() *RustExterns
	GetReport() *ValidationReport
}

type Bindgen struct {
	config    BindgenConfig
	wdk       *WdkBindings
	externs   *RustExterns
	report    *ValidationReport
	fixReport *FixReport
}

func NewBindgen(config BindgenConfig) *Bindgen {
	return &Bindgen{
		config:    config,
		wdk:       &WdkBindings{Functions: make(map[string]FunctionSignature)},
		externs:   &RustExterns{Functions: make(map[string][]ExternFunction)},
		report:    &ValidationReport{Statistics: make(map[string]int)},
		fixReport: &FixReport{FilesWithFixes: make(map[string][]WdkUsage)},
	}
}

func (b *Bindgen) LoadWdkBindings(ntddkPath, typesPath, constantsPath string) error {
	if err := b.parseNtddkFunctions(ntddkPath); err != nil {
		return fmt.Errorf("failed to parse ntddk: %w", err)
	}

	if err := b.parseTypes(typesPath); err != nil {
		return fmt.Errorf("failed to parse types: %w", err)
	}

	if err := b.parseConstants(constantsPath); err != nil {
		return fmt.Errorf("failed to parse constants: %w", err)
	}

	return nil
}

func (b *Bindgen) parseNtddkFunctions(path string) error {
	content, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	b.wdk.SourceFile = path

	lines := strings.Split(string(content), "\n")

	funcStartPattern := regexp.MustCompile(`pub fn ([A-Z][a-zA-Z0-9]+)\s*\(`)

	i := 0
	for i < len(lines) {
		line := lines[i]
		matches := funcStartPattern.FindStringSubmatch(line)
		if matches == nil {
			i++
			continue
		}

		name := matches[1]
		if !b.isWdkPrefix(name) {
			i++
			continue
		}

		fullDecl := line
		for !strings.Contains(fullDecl, ";") && i+1 < len(lines) {
			i++
			fullDecl += "\n" + lines[i]
		}

		params := ""
		returnType := "void"

		if idx := strings.Index(fullDecl, "("); idx != -1 {
			braceCount := 0
			start := idx
			for j := idx; j < len(fullDecl); j++ {
				if fullDecl[j] == '(' {
					braceCount++
				} else if fullDecl[j] == ')' {
					braceCount--
					if braceCount == 0 {
						params = fullDecl[start+1 : j]
						rest := strings.TrimSpace(fullDecl[j+1:])
						if strings.HasPrefix(rest, "->") {
							retPart := strings.TrimPrefix(rest, "->")
							retPart = strings.TrimSuffix(retPart, ";")
							returnType = strings.TrimSpace(retPart)
						}
						break
					}
				}
			}
		}

		b.wdk.Functions[name] = FunctionSignature{
			Name:       name,
			Params:     strings.TrimSpace(params),
			ReturnType: returnType,
			Line:       i + 1,
			RawLine:    strings.TrimSpace(fullDecl),
		}
		i++
	}

	return nil
}

func (b *Bindgen) parseTypes(path string) error {
	content, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	b.wdk.Types = make(map[string]TypeDefinition)

	structPattern := regexp.MustCompile(`^pub struct ([A-Z_][a-zA-Z0-9_]*)`)
	typePattern := regexp.MustCompile(`^pub type ([A-Z_][a-zA-Z0-9_]*)\s*=`)
	enumPattern := regexp.MustCompile(`^pub enum ([A-Z_][a-zA-Z0-9_]*)`)

	skipTypes := map[string]bool{
		"Type": true,
	}

	lines := strings.Split(string(content), "\n")
	for i, line := range lines {
		if matches := structPattern.FindStringSubmatch(line); matches != nil {
			if !skipTypes[matches[1]] {
				b.wdk.Types[matches[1]] = TypeDefinition{Name: matches[1], Kind: "struct", Line: i + 1}
			}
		}
		if matches := typePattern.FindStringSubmatch(line); matches != nil {
			if !skipTypes[matches[1]] {
				b.wdk.Types[matches[1]] = TypeDefinition{Name: matches[1], Kind: "type", Line: i + 1}
			}
		}
		if matches := enumPattern.FindStringSubmatch(line); matches != nil {
			if !skipTypes[matches[1]] {
				b.wdk.Types[matches[1]] = TypeDefinition{Name: matches[1], Kind: "enum", Line: i + 1}
			}
		}
	}

	return nil
}

func (b *Bindgen) parseConstants(path string) error {
	content, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	b.wdk.Constants = make(map[string]ConstantDefinition)

	constPattern := regexp.MustCompile(`pub const ([A-Z_][A-Z0-9_]*)\s*:`)

	lines := strings.Split(string(content), "\n")
	for i, line := range lines {
		if matches := constPattern.FindStringSubmatch(line); matches != nil {
			b.wdk.Constants[matches[1]] = ConstantDefinition{
				Name:  matches[1],
				Line:  i + 1,
				Value: strings.TrimSpace(line),
			}
		}
	}

	return nil
}

func (b *Bindgen) ScanRustExterns(rustSourceDir string) error {
	b.externs.Functions = make(map[string][]ExternFunction)
	b.externs.Files = nil

	externPattern := regexp.MustCompile(`extern\s+"C"\s*\{`)
	funcPattern := regexp.MustCompile(`fn\s+([A-Z][a-zA-Z0-9]+)\s*\(([^)]*)\)(?:\s*->\s*([^;{]+))?`)

	err := filepath.Walk(rustSourceDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() || filepath.Ext(path) != ".rs" {
			return nil
		}

		content, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		relPath, _ := filepath.Rel(b.config.ProjectRoot, path)
		b.externs.Files = append(b.externs.Files, relPath)

		lines := strings.Split(string(content), "\n")
		inExternBlock := false

		for i, line := range lines {
			if externPattern.MatchString(line) {
				inExternBlock = true
				continue
			}

			if inExternBlock {
				if strings.Contains(line, "}") {
					inExternBlock = false
					continue
				}

				if matches := funcPattern.FindStringSubmatch(line); matches != nil {
					name := matches[1]
					if !b.isWdkPrefix(name) {
						continue
					}

					params := strings.TrimSpace(matches[2])
					returnType := "void"
					if len(matches) > 3 && matches[3] != "" {
						returnType = strings.TrimSpace(matches[3])
					}

					_, foundInWdk := b.wdk.Functions[name]

					extern := ExternFunction{
						Name:       name,
						File:       relPath,
						Line:       i + 1,
						Params:     params,
						ReturnType: returnType,
						FoundInWdk: foundInWdk,
					}

					b.externs.Functions[name] = append(b.externs.Functions[name], extern)
				}
			}
		}

		return nil
	})

	return err
}

func (b *Bindgen) Validate() *ValidationReport {
	b.report.Mismatches = nil
	b.report.NotExported = nil
	b.report.MissingFromWdk = nil

	for name, funcs := range b.externs.Functions {
		wdkSig, exists := b.wdk.Functions[name]
		if !exists {
			b.report.NotExported = append(b.report.NotExported, name)
			continue
		}

		for _, rustFunc := range funcs {
			if mismatch := b.compareSignatures(name, rustFunc, wdkSig); mismatch != nil {
				b.report.Mismatches = append(b.report.Mismatches, *mismatch)
			}
		}
	}

	for name, funcs := range b.externs.Functions {
		if !funcs[0].FoundInWdk {
			b.report.MissingFromWdk = append(b.report.MissingFromWdk, name)
		}
	}

	sort.Strings(b.report.NotExported)
	sort.Strings(b.report.MissingFromWdk)

	b.report.Statistics["total_externs"] = len(b.externs.Functions)
	b.report.Statistics["total_wdk_functions"] = len(b.wdk.Functions)
	b.report.Statistics["total_wdk_types"] = len(b.wdk.Types)
	b.report.Statistics["total_wdk_constants"] = len(b.wdk.Constants)
	b.report.Statistics["mismatches"] = len(b.report.Mismatches)
	b.report.Statistics["not_exported"] = len(b.report.NotExported)
	b.report.Statistics["missing_from_wdk"] = len(b.report.MissingFromWdk)

	return b.report
}

func (b *Bindgen) compareSignatures(name string, rustFunc ExternFunction, wdkSig FunctionSignature) *TypeMismatch {
	rustParams := b.normalizeParams(rustFunc.Params)
	wdkParams := b.normalizeParams(wdkSig.Params)

	rustReturn := b.normalizeType(rustFunc.ReturnType)
	wdkReturn := b.normalizeType(wdkSig.ReturnType)

	var mismatches []string

	if rustReturn != wdkReturn {
		mismatches = append(mismatches, fmt.Sprintf("return type: rust=%s, wdk=%s", rustReturn, wdkReturn))
	}

	rustParamList := b.parseParamList(rustParams)
	wdkParamList := b.parseParamList(wdkParams)

	if len(rustParamList) != len(wdkParamList) {
		mismatches = append(mismatches, fmt.Sprintf("param count: rust=%d, wdk=%d", len(rustParamList), len(wdkParamList)))
	} else {
		for i, rp := range rustParamList {
			wp := wdkParamList[i]
			rustType := b.normalizeType(rp.Type)
			wdkType := b.normalizeType(wp.Type)

			if !b.typesCompatible(rustType, wdkType) {
				mismatches = append(mismatches, fmt.Sprintf("param[%d]: rust=%s, wdk=%s", i, rustType, wdkType))
			}
		}
	}

	if len(mismatches) == 0 {
		return nil
	}

	return &TypeMismatch{
		FunctionName:  name,
		RustFile:      rustFunc.File,
		RustLine:      rustFunc.Line,
		RustSignature: rustFunc.Params,
		WdkSignature:  wdkSig.Params,
		WdkLine:       wdkSig.Line,
		MismatchType:  strings.Join(mismatches, "; "),
	}
}

func (b *Bindgen) normalizeParams(params string) string {
	params = strings.TrimSpace(params)
	params = strings.ReplaceAll(params, "\n", " ")
	params = strings.ReplaceAll(params, "\t", " ")
	for strings.Contains(params, "  ") {
		params = strings.ReplaceAll(params, "  ", " ")
	}
	return params
}

func (b *Bindgen) normalizeType(t string) string {
	t = strings.TrimSpace(t)
	t = strings.ToLower(t)
	t = strings.TrimPrefix(t, "mut ")
	t = strings.TrimPrefix(t, "const ")
	t = strings.TrimSuffix(t, "_t")
	return t
}

func (b *Bindgen) parseParamList(params string) []struct{ Name, Type string } {
	if params == "" {
		return nil
	}

	var result []struct{ Name, Type string }

	parts := strings.Split(params, ",")
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}

		fields := strings.Fields(part)
		if len(fields) >= 2 {
			result = append(result, struct{ Name, Type string }{
				Name: fields[len(fields)-1],
				Type: strings.Join(fields[:len(fields)-1], " "),
			})
		} else if len(fields) == 1 {
			result = append(result, struct{ Name, Type string }{
				Type: fields[0],
			})
		}
	}

	return result
}

func (b *Bindgen) typesCompatible(rustType, wdkType string) bool {
	if rustType == wdkType {
		return true
	}

	pointerTypes := map[string][]string{
		"u64":  {"pvoid", "handle", "ulong64", "usize", "isize", "i64"},
		"u32":  {"ulong", "uint32", "dword"},
		"i32":  {"ntstatus", "long", "int32"},
		"bool": {"boolean"},
	}

	for rust, wdkList := range pointerTypes {
		if rustType == rust {
			for _, wdk := range wdkList {
				if wdkType == wdk || strings.Contains(wdkType, wdk) {
					return true
				}
			}
		}
	}

	if strings.HasSuffix(rustType, "*") && strings.HasSuffix(wdkType, "*") {
		return true
	}

	if strings.HasPrefix(rustType, "*") && strings.Contains(wdkType, "p") {
		return true
	}

	return false
}

func (b *Bindgen) isWdkPrefix(name string) bool {
	for _, prefix := range []string{"Ps", "Mm", "Io", "Ex", "Ob", "Ke", "Rtl", "Nt", "Dbg"} {
		if strings.HasPrefix(name, prefix) {
			return true
		}
	}
	return false
}

func (b *Bindgen) GenerateFunctionsList(outputPath string) error {
	var sb strings.Builder
	sb.WriteString("// Auto-generated from ntddk.rs - DO NOT EDIT\n\n")

	prefixes := make(map[string][]FunctionSignature)
	for name, sig := range b.wdk.Functions {
		prefix := name[:2]
		prefixes[prefix] = append(prefixes[prefix], sig)
	}

	var prefixList []string
	for p := range prefixes {
		prefixList = append(prefixList, p)
	}
	sort.Strings(prefixList)

	for _, prefix := range prefixList {
		funcs := prefixes[prefix]
		sort.Slice(funcs, func(i, j int) bool { return funcs[i].Name < funcs[j].Name })

		sb.WriteString(fmt.Sprintf("// === %s Functions (%d) ===\n", prefix, len(funcs)))
		for _, f := range funcs {
			sb.WriteString(fmt.Sprintf("// Line %d: pub fn %s(%s)", f.Line, f.Name, f.Params))
			if f.ReturnType != "void" {
				sb.WriteString(fmt.Sprintf(" -> %s", f.ReturnType))
			}
			sb.WriteString("\n")
		}
		sb.WriteString("\n")
	}

	return os.WriteFile(outputPath, []byte(sb.String()), 0644)
}

func (b *Bindgen) GenerateFunctionsGoMap(outputPath string) error {
	var sb strings.Builder
	sb.WriteString("// Auto-generated from ntddk.rs - DO NOT EDIT\n\n")
	sb.WriteString("package main\n\n")
	sb.WriteString("// WdkFunctions contains all function names from ntddk.rs\n")
	sb.WriteString("var WdkFunctions = map[string]bool{\n")

	var names []string
	for name := range b.wdk.Functions {
		names = append(names, name)
	}
	sort.Strings(names)

	for _, name := range names {
		sb.WriteString(fmt.Sprintf("\t%q: true,\n", name))
	}
	sb.WriteString("}\n")

	return os.WriteFile(outputPath, []byte(sb.String()), 0644)
}

func (b *Bindgen) GenerateTypesGoMap(outputPath string) error {
	var sb strings.Builder
	sb.WriteString("// Auto-generated from types.rs - DO NOT EDIT\n\n")
	sb.WriteString("package main\n\n")
	sb.WriteString("// WdkTypes contains all type names from types.rs\n")
	sb.WriteString("var WdkTypes = map[string]bool{\n")

	var names []string
	for name := range b.wdk.Types {
		names = append(names, name)
	}
	sort.Strings(names)

	for _, name := range names {
		sb.WriteString(fmt.Sprintf("\t%q: true,\n", name))
	}
	sb.WriteString("}\n")

	return os.WriteFile(outputPath, []byte(sb.String()), 0644)
}

func (b *Bindgen) GenerateConstantsGoMap(outputPath string) error {
	var sb strings.Builder
	sb.WriteString("// Auto-generated from constants.rs - DO NOT EDIT\n\n")
	sb.WriteString("package main\n\n")
	sb.WriteString("// WdkConstants contains all constant names from constants.rs\n")
	sb.WriteString("var WdkConstants = map[string]bool{\n")

	var names []string
	for name := range b.wdk.Constants {
		names = append(names, name)
	}
	sort.Strings(names)

	for _, name := range names {
		sb.WriteString(fmt.Sprintf("\t%q: true,\n", name))
	}
	sb.WriteString("}\n")

	return os.WriteFile(outputPath, []byte(sb.String()), 0644)
}

func (b *Bindgen) GenerateValidationReport(outputPath string) error {
	var sb strings.Builder
	sb.WriteString("// Auto-generated validation report - DO NOT EDIT\n\n")
	sb.WriteString("=== WDK BINDING VALIDATION REPORT ===\n\n")

	sb.WriteString("## STATISTICS\n\n")
	sb.WriteString(fmt.Sprintf("Total WDK Functions: %d\n", b.report.Statistics["total_wdk_functions"]))
	sb.WriteString(fmt.Sprintf("Total WDK Types: %d\n", b.report.Statistics["total_wdk_types"]))
	sb.WriteString(fmt.Sprintf("Total WDK Constants: %d\n", b.report.Statistics["total_wdk_constants"]))
	sb.WriteString(fmt.Sprintf("Total Rust Extern Functions: %d\n", b.report.Statistics["total_externs"]))
	sb.WriteString(fmt.Sprintf("Type Mismatches: %d\n", b.report.Statistics["mismatches"]))
	sb.WriteString(fmt.Sprintf("Not Exported by WDK: %d\n", b.report.Statistics["not_exported"]))
	sb.WriteString("\n")

	if len(b.report.Mismatches) > 0 {
		sb.WriteString("## TYPE MISMATCHES\n\n")
		for _, m := range b.report.Mismatches {
			sb.WriteString(fmt.Sprintf("%s:%d: %s\n", m.RustFile, m.RustLine, m.FunctionName))
			sb.WriteString(fmt.Sprintf("  %s\n", m.MismatchType))
			sb.WriteString(fmt.Sprintf("  Rust: %s\n", m.RustSignature))
			sb.WriteString(fmt.Sprintf("  WDK:  %s (line %d)\n\n", m.WdkSignature, m.WdkLine))
		}
	} else {
		sb.WriteString("## TYPE MISMATCHES\n\nNo type mismatches found.\n\n")
	}

	if len(b.report.NotExported) > 0 {
		sb.WriteString("## FUNCTIONS NOT EXPORTED BY WDK\n\n")
		sb.WriteString("// These functions are not exported by wdk-sys and need manual declaration\n")
		for _, name := range b.report.NotExported {
			funcs := b.externs.Functions[name]
			sb.WriteString(fmt.Sprintf("%s\n", name))
			for _, f := range funcs {
				sb.WriteString(fmt.Sprintf("  -> %s:%d\n", f.File, f.Line))
			}
		}
		sb.WriteString("\n")
	}

	return os.WriteFile(outputPath, []byte(sb.String()), 0644)
}

func (b *Bindgen) GenerateFixSuggestions(outputPath string) error {
	var sb strings.Builder
	sb.WriteString("// Auto-generated fix suggestions - DO NOT EDIT\n\n")
	sb.WriteString("=== SUGGESTED FIXES FOR RUST EXTERN DECLARATIONS ===\n\n")

	sb.WriteString("## Functions that should use wdk_sys::ntddk\n\n")

	var availableFuncs []string
	for name := range b.externs.Functions {
		if _, exists := b.wdk.Functions[name]; exists {
			availableFuncs = append(availableFuncs, name)
		}
	}
	sort.Strings(availableFuncs)

	for _, name := range availableFuncs {
		funcs := b.externs.Functions[name]
		sb.WriteString(fmt.Sprintf("### %s\n", name))
		sb.WriteString(fmt.Sprintf("// Found in WDK: use `use wdk_sys::ntddk::%s;`\n", name))
		sb.WriteString("// Files to update:\n")
		for _, f := range funcs {
			sb.WriteString(fmt.Sprintf("//   - %s:%d\n", f.File, f.Line))
		}
		sb.WriteString("\n")
	}

	sb.WriteString("## Functions NOT in WDK bindings (need manual declaration)\n\n")

	for _, name := range b.report.NotExported {
		funcs := b.externs.Functions[name]
		sb.WriteString(fmt.Sprintf("### %s\n", name))
		sb.WriteString("// Not exported by WDK - add to ntapi/mod.rs\n")
		for _, f := range funcs {
			sb.WriteString(fmt.Sprintf("//   - %s:%d\n", f.File, f.Line))
		}
		sb.WriteString("\n")
	}

	return os.WriteFile(outputPath, []byte(sb.String()), 0644)
}

func (b *Bindgen) GetWdkBindings() *WdkBindings {
	return b.wdk
}

func (b *Bindgen) GetRustExterns() *RustExterns {
	return b.externs
}

func (b *Bindgen) GetReport() *ValidationReport {
	return b.report
}

func (b *Bindgen) GenerateNtapiMod(outputDir string, notExportedFunctions []NotExportedFunc) error {
	if err := b.generateNtapiApi(outputDir, notExportedFunctions); err != nil {
		return err
	}
	if err := b.generateNtapiTypes(outputDir); err != nil {
		return err
	}
	if err := b.generateNtapiConstants(outputDir); err != nil {
		return err
	}
	return b.generateNtapiModRs(outputDir)
}

func (b *Bindgen) generateNtapiApi(outputDir string, notExportedFunctions []NotExportedFunc) error {
	var sb strings.Builder
	sb.WriteString("#![allow(non_snake_case)]\n")
	sb.WriteString("#![allow(dead_code)]\n\n")

	exportedFuncs := []string{}
	for name := range b.wdk.Functions {
		if !b.isNotExported(name, notExportedFunctions) {
			exportedFuncs = append(exportedFuncs, name)
		}
	}
	sort.Strings(exportedFuncs)

	sb.WriteString(fmt.Sprintf("// Exported functions: %d\n", len(exportedFuncs)))
	sb.WriteString(fmt.Sprintf("// Not exported functions: %d\n\n", len(notExportedFunctions)))

	sb.WriteString("use super::types::*;\n\n")

	sb.WriteString("pub mod exported {\n")
	sb.WriteString("    pub use wdk_sys::ntddk::{\n")

	for _, name := range exportedFuncs {
		sb.WriteString(fmt.Sprintf("        %s,\n", name))
	}
	sb.WriteString("    };\n")
	sb.WriteString("}\n\n")

	sb.WriteString("pub mod not_exported {\n")
	sb.WriteString("    use super::*;\n\n")
	sb.WriteString("    extern \"C\" {\n")

	for _, f := range notExportedFunctions {
		sb.WriteString(fmt.Sprintf("        pub fn %s(%s)", f.Name, f.Params))
		if f.ReturnType != "" && f.ReturnType != "void" {
			sb.WriteString(fmt.Sprintf(" -> %s", f.ReturnType))
		}
		sb.WriteString(";\n")
	}

	sb.WriteString("    }\n")
	sb.WriteString("}\n\n")

	sb.WriteString("pub use exported::*;\n")
	sb.WriteString("pub use not_exported::*;\n")

	return os.WriteFile(filepath.Join(outputDir, "api.rs"), []byte(sb.String()), 0644)
}

func (b *Bindgen) generateNtapiTypes(outputDir string) error {
	var sb strings.Builder
	sb.WriteString("#![allow(non_snake_case)]\n")
	sb.WriteString("#![allow(dead_code)]\n\n")

	exportedTypes := []string{}
	for name := range b.wdk.Types {
		if strings.HasPrefix(name, "_") {
			continue
		}
		exportedTypes = append(exportedTypes, name)
	}
	sort.Strings(exportedTypes)

	sb.WriteString(fmt.Sprintf("// Types: %d\n\n", len(exportedTypes)))

	sb.WriteString("pub use wdk_sys::{\n")

	for _, name := range exportedTypes {
		sb.WriteString(fmt.Sprintf("    %s,\n", name))
	}
	sb.WriteString("};\n")

	return os.WriteFile(filepath.Join(outputDir, "types.rs"), []byte(sb.String()), 0644)
}

func (b *Bindgen) generateNtapiConstants(outputDir string) error {
	var sb strings.Builder
	sb.WriteString("#![allow(non_snake_case)]\n")
	sb.WriteString("#![allow(dead_code)]\n\n")

	exportedConstants := []string{}
	for name := range b.wdk.Constants {
		if strings.HasPrefix(name, "_") {
			continue
		}
		exportedConstants = append(exportedConstants, name)
	}
	sort.Strings(exportedConstants)

	sb.WriteString(fmt.Sprintf("// Constants: %d\n\n", len(exportedConstants)))

	sb.WriteString("pub use wdk_sys::{\n")

	for _, name := range exportedConstants {
		sb.WriteString(fmt.Sprintf("    %s,\n", name))
	}
	sb.WriteString("};\n")

	return os.WriteFile(filepath.Join(outputDir, "constants.rs"), []byte(sb.String()), 0644)
}

func (b *Bindgen) generateNtapiModRs(outputDir string) error {
	var sb strings.Builder
	sb.WriteString("#![allow(non_snake_case)]\n")
	sb.WriteString("#![allow(dead_code)]\n\n")

	sb.WriteString("mod api;\n")
	sb.WriteString("mod types;\n")
	sb.WriteString("mod constants;\n\n")

	sb.WriteString("pub use api::*;\n")
	sb.WriteString("pub use types::*;\n")
	sb.WriteString("pub use constants::*;\n")

	return os.WriteFile(filepath.Join(outputDir, "mod.rs"), []byte(sb.String()), 0644)
}

func (b *Bindgen) isNotExported(name string, notExported []NotExportedFunc) bool {
	for _, f := range notExported {
		if f.Name == name {
			return true
		}
	}
	return false
}

type NotExportedFunc struct {
	Name       string
	Params     string
	ReturnType string
}

func (b *Bindgen) ScanProjectUsage(projectRoot string, excludeDirs []string) error {
	b.fixReport = &FixReport{
		FilesWithFixes: make(map[string][]WdkUsage),
	}

	excludeMap := make(map[string]bool)
	for _, dir := range excludeDirs {
		excludeMap[dir] = true
	}

	err := filepath.Walk(projectRoot, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			relPath, _ := filepath.Rel(projectRoot, path)
			if excludeMap[relPath] || strings.Contains(relPath, "target") || strings.Contains(relPath, "todo") {
				return filepath.SkipDir
			}
			return nil
		}

		if !strings.HasSuffix(path, ".rs") {
			return nil
		}

		if strings.Contains(path, "ntapi/") {
			return nil
		}

		b.scanFileForWdkUsage(path)
		return nil
	})

	if err != nil {
		return fmt.Errorf("failed to scan project: %w", err)
	}

	b.fixReport.TotalUsages = len(b.fixReport.Usages)
	b.fixReport.FilesScanned = len(b.fixReport.FilesWithFixes)
	for _, usages := range b.fixReport.FilesWithFixes {
		for _, u := range usages {
			if u.NeedsFix {
				b.fixReport.NeedsFix++
			}
		}
	}

	return nil
}

func (b *Bindgen) scanFileForWdkUsage(filePath string) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return
	}

	lines := strings.Split(string(content), "\n")

	externBlockPattern := regexp.MustCompile(`extern\s+"(C|system)"\s*\{`)
	fnDeclPattern := regexp.MustCompile(`fn\s+([A-Z][a-zA-Z0-9]+)\s*\(([^)]*)\)(?:\s*->\s*([^;{]+))?`)
	useWdkSysPattern := regexp.MustCompile(`use\s+wdk_sys::`)
	useCrateNtapiPattern := regexp.MustCompile(`use\s+crate::ntapi::`)

	inExternBlock := false

	for i, line := range lines {
		if externBlockPattern.MatchString(line) {
			inExternBlock = true
			continue
		}

		if inExternBlock && strings.TrimSpace(line) == "}" {
			inExternBlock = false
			continue
		}

		if inExternBlock {
			matches := fnDeclPattern.FindStringSubmatch(line)
			if matches != nil {
				name := matches[1]
				if b.isWdkPrefix(name) {
					usage := WdkUsage{
						File:         filePath,
						Line:         i + 1,
						Kind:         "function",
						Name:         name,
						SourceType:   "extern_block",
						NeedsFix:     true,
						SuggestedFix: fmt.Sprintf("use crate::ntapi::%s;", name),
					}
					b.fixReport.Usages = append(b.fixReport.Usages, usage)
					b.fixReport.FilesWithFixes[filePath] = append(b.fixReport.FilesWithFixes[filePath], usage)
				}
			}
		}

		if useWdkSysPattern.MatchString(line) {
			usage := WdkUsage{
				File:         filePath,
				Line:         i + 1,
				Kind:         "import",
				Name:         "wdk_sys",
				SourceType:   "use_wdk_sys",
				NeedsFix:     true,
				SuggestedFix: "use crate::ntapi::*;",
			}
			b.fixReport.Usages = append(b.fixReport.Usages, usage)
			b.fixReport.FilesWithFixes[filePath] = append(b.fixReport.FilesWithFixes[filePath], usage)
		}

		if useCrateNtapiPattern.MatchString(line) {
			usage := WdkUsage{
				File:       filePath,
				Line:       i + 1,
				Kind:       "import",
				Name:       "ntapi",
				SourceType: "use_crate_ntapi",
				NeedsFix:   false,
			}
			b.fixReport.Usages = append(b.fixReport.Usages, usage)
		}
	}
}

func (b *Bindgen) GetFixReport() *FixReport {
	return b.fixReport
}

func (b *Bindgen) GenerateUsageReport(outputPath string) error {
	var sb strings.Builder

	sb.WriteString("# WDK Usage Fix Report\n\n")
	sb.WriteString(fmt.Sprintf("Generated: %s\n\n", "now"))
	sb.WriteString("## Summary\n\n")
	sb.WriteString(fmt.Sprintf("- Files Scanned: %d\n", b.fixReport.FilesScanned))
	sb.WriteString(fmt.Sprintf("- Total Usages: %d\n", b.fixReport.TotalUsages))
	sb.WriteString(fmt.Sprintf("- Needs Fix: %d\n\n", b.fixReport.NeedsFix))

	sb.WriteString("## Files Needing Fixes\n\n")

	for file, usages := range b.fixReport.FilesWithFixes {
		needsFix := false
		for _, u := range usages {
			if u.NeedsFix {
				needsFix = true
				break
			}
		}
		if !needsFix {
			continue
		}

		relPath, _ := filepath.Rel(b.config.ProjectRoot, file)
		sb.WriteString(fmt.Sprintf("### %s\n\n", relPath))

		for _, u := range usages {
			if u.NeedsFix {
				sb.WriteString(fmt.Sprintf("- Line %d: %s `%s` (source: %s)\n", u.Line, u.Kind, u.Name, u.SourceType))
				sb.WriteString(fmt.Sprintf("  - Suggested fix: `%s`\n", u.SuggestedFix))
			}
		}
		sb.WriteString("\n")
	}

	sb.WriteString("## Detailed Usage List\n\n")

	for _, u := range b.fixReport.Usages {
		relPath, _ := filepath.Rel(b.config.ProjectRoot, u.File)
		sb.WriteString(fmt.Sprintf("- %s:%d - %s: %s (needs fix: %v)\n", relPath, u.Line, u.Kind, u.Name, u.NeedsFix))
	}

	return os.WriteFile(outputPath, []byte(sb.String()), 0644)
}

func (b *Bindgen) ApplyFixes(dryRun bool) error {
	if b.fixReport == nil {
		return fmt.Errorf("no fix report available, run ScanProjectUsage first")
	}

	fileModifications := make(map[string][]FileModification)

	for file, usages := range b.fixReport.FilesWithFixes {
		var mods []FileModification
		for _, u := range usages {
			if !u.NeedsFix {
				continue
			}

			switch u.SourceType {
			case "use_wdk_sys":
				mods = append(mods, FileModification{
					Line:     u.Line,
					Kind:     "replace_use",
					OldText:  "",
					NewText:  u.SuggestedFix,
					FuncName: "",
				})
			case "extern_block":
				mods = append(mods, FileModification{
					Line:     u.Line,
					Kind:     "add_use",
					OldText:  "",
					NewText:  u.SuggestedFix,
					FuncName: u.Name,
				})
			}
		}
		if len(mods) > 0 {
			fileModifications[file] = mods
		}
	}

	for file, mods := range fileModifications {
		if dryRun {
			relPath, _ := filepath.Rel(b.config.ProjectRoot, file)
			fmt.Printf("Would modify: %s\n", relPath)
			for _, m := range mods {
				fmt.Printf("  Line %d: %s\n", m.Line, m.NewText)
			}
			continue
		}

		err := b.applyFileModifications(file, mods)
		if err != nil {
			return fmt.Errorf("failed to modify %s: %w", file, err)
		}
	}

	return nil
}

type FileModification struct {
	Line     int
	Kind     string
	OldText  string
	NewText  string
	FuncName string
}

func (b *Bindgen) applyFileModifications(filePath string, mods []FileModification) error {
	content, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	lines := strings.Split(string(content), "\n")

	linesToDelete := make(map[int]bool)

	for i := 0; i < len(lines); i++ {
		line := lines[i]
		trimmed := strings.TrimSpace(line)

		if strings.HasPrefix(trimmed, "use wdk_sys::") {
			if strings.Contains(line, "{") && !strings.Contains(line, "};") {
				indent := ""
				for _, c := range line {
					if c == ' ' || c == '\t' {
						indent += string(c)
					} else {
						break
					}
				}
				for j := i + 1; j < len(lines); j++ {
					if strings.Contains(lines[j], "};") {
						for k := i + 1; k <= j; k++ {
							linesToDelete[k] = true
						}
						lines[i] = indent + "use crate::ntapi::*;"
						i = j
						break
					}
				}
			} else {
				indent := ""
				for _, c := range line {
					if c == ' ' || c == '\t' {
						indent += string(c)
					} else {
						break
					}
				}
				lines[i] = indent + "use crate::ntapi::*;"
			}
		}
	}

	externBlockRanges := b.findExternBlockRanges(lines, mods)
	for _, r := range externBlockRanges {
		for i := r.start; i <= r.end; i++ {
			linesToDelete[i] = true
		}
	}

	useStatements := make(map[string]bool)
	for _, mod := range mods {
		if mod.Kind == "add_use" {
			useStatements[mod.FuncName] = true
		}
	}

	var result []string
	var addedUses map[string]bool = make(map[string]bool)
	var lastUseEndLine int = -1
	var hasNtapiWildcard bool

	inMultilineUse := false
	for i, line := range lines {
		if linesToDelete[i] {
			continue
		}

		trimmed := strings.TrimSpace(line)

		if strings.HasPrefix(trimmed, "use ") {
			if strings.Contains(line, "{") && !strings.Contains(line, "};") {
				inMultilineUse = true
			}
			if !inMultilineUse {
				lastUseEndLine = len(result)
			}
		}

		if inMultilineUse && strings.Contains(line, "};") {
			inMultilineUse = false
			lastUseEndLine = len(result)
		}

		if trimmed == "use crate::ntapi::*;" {
			hasNtapiWildcard = true
		}

		result = append(result, line)
	}

	if hasNtapiWildcard {
		useStatements = make(map[string]bool)
	}

	var finalResult []string
	for i, line := range result {
		trimmed := strings.TrimSpace(line)

		if strings.HasPrefix(trimmed, "use crate::ntapi::") {
			parts := strings.Split(trimmed, "::")
			if len(parts) >= 3 {
				name := strings.TrimSuffix(parts[len(parts)-1], ";")
				addedUses[name] = true
			}
		}

		if i == lastUseEndLine && len(useStatements) > 0 {
			finalResult = append(finalResult, line)
			for name := range useStatements {
				if !addedUses[name] {
					finalResult = append(finalResult, fmt.Sprintf("use crate::ntapi::%s;", name))
				}
			}
		} else {
			finalResult = append(finalResult, line)
		}
	}

	newContent := strings.Join(finalResult, "\n")
	return os.WriteFile(filePath, []byte(newContent), 0644)
}

type LineRange struct {
	start int
	end   int
}

func (b *Bindgen) findMultilineUseRanges(lines []string) []LineRange {
	var ranges []LineRange

	for i := 0; i < len(lines); i++ {
		line := lines[i]
		if strings.Contains(line, "use wdk_sys::") {
			if strings.Contains(line, "{") && !strings.Contains(line, "}") {
				start := i
				for j := i + 1; j < len(lines); j++ {
					if strings.Contains(lines[j], "};") {
						ranges = append(ranges, LineRange{start: start, end: j})
						i = j
						break
					}
				}
			}
		}
	}

	return ranges
}

func (b *Bindgen) findExternBlockRanges(lines []string, mods []FileModification) []LineRange {
	var ranges []LineRange

	for i := 0; i < len(lines); i++ {
		line := lines[i]
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "extern \"C\"") && strings.HasSuffix(trimmed, "{") {
			braceCount := 0
			start := i
			hasWdkFunc := false

			for j := i; j < len(lines); j++ {
				for _, c := range lines[j] {
					if c == '{' {
						braceCount++
					} else if c == '}' {
						braceCount--
						if braceCount == 0 {
							for k := start + 1; k < j; k++ {
								fnLine := strings.TrimSpace(lines[k])
								if strings.HasPrefix(fnLine, "fn ") {
									idx := strings.Index(fnLine, "(")
									if idx > 3 {
										name := strings.TrimSpace(fnLine[3:idx])
										if b.isWdkPrefix(name) {
											hasWdkFunc = true
											break
										}
									}
								}
							}
							if hasWdkFunc {
								ranges = append(ranges, LineRange{start: start, end: j})
							}
							i = j
							break
						}
					}
				}
				if braceCount == 0 {
					break
				}
			}
		}
	}

	return ranges
}
