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
	Name   string
	Kind   string
	Line   int
	Source string
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
	ApplyFixes() error
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

	if err := b.parseTypes(ntddkPath); err != nil {
		return fmt.Errorf("failed to parse ntddk types: %w", err)
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
						if after, ok := strings.CutPrefix(rest, "->"); ok {
							retPart := after
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

	if b.wdk.Types == nil {
		b.wdk.Types = make(map[string]TypeDefinition)
	}

	source := "types"
	if strings.HasSuffix(path, "ntddk.rs") {
		source = "ntddk"
	}

	structPattern := regexp.MustCompile(`^pub struct ([A-Z_][a-zA-Z0-9_]*)`)
	typePattern := regexp.MustCompile(`^pub type ([A-Z_][a-zA-Z0-9_]*)\s*=`)
	enumPattern := regexp.MustCompile(`^pub enum ([A-Z_][a-zA-Z0-9_]*)`)
	useAsPattern := regexp.MustCompile(`^pub use self::_([A-Z_][a-zA-Z0-9_]*)::Type as ([A-Z_][a-zA-Z0-9_]*);`)
	modPattern := regexp.MustCompile(`^pub mod ([A-Z_][a-zA-Z0-9_]*)\s*\{`)

	skipTypes := map[string]bool{
		"Type": true,
	}

	lines := strings.Split(string(content), "\n")
	for i, line := range lines {
		if matches := structPattern.FindStringSubmatch(line); matches != nil {
			if !skipTypes[matches[1]] {
				b.wdk.Types[matches[1]] = TypeDefinition{Name: matches[1], Kind: "struct", Line: i + 1, Source: source}
			}
		}
		if matches := typePattern.FindStringSubmatch(line); matches != nil {
			if !skipTypes[matches[1]] {
				b.wdk.Types[matches[1]] = TypeDefinition{Name: matches[1], Kind: "type", Line: i + 1, Source: source}
			}
		}
		if matches := enumPattern.FindStringSubmatch(line); matches != nil {
			if !skipTypes[matches[1]] {
				b.wdk.Types[matches[1]] = TypeDefinition{Name: matches[1], Kind: "enum", Line: i + 1, Source: source}
			}
		}
		if matches := useAsPattern.FindStringSubmatch(line); matches != nil {
			if !skipTypes[matches[2]] {
				b.wdk.Types[matches[2]] = TypeDefinition{Name: matches[2], Kind: "type", Line: i + 1, Source: source}
			}
		}
		if matches := modPattern.FindStringSubmatch(line); matches != nil {
			if !skipTypes[matches[1]] {
				b.wdk.Types[matches[1]] = TypeDefinition{Name: matches[1], Kind: "mod", Line: i + 1, Source: source}
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

func (b *Bindgen) splitParams(params string) []string {
	var result []string
	var current strings.Builder
	braceCount := 0
	angleCount := 0
	squareCount := 0

	for _, c := range params {
		switch c {
		case '(':
			braceCount++
		case ')':
			braceCount--
		case '<':
			angleCount++
		case '>':
			angleCount--
		case '[':
			squareCount++
		case ']':
			squareCount--
		case ',':
			if braceCount == 0 && angleCount == 0 && squareCount == 0 {
				result = append(result, current.String())
				current.Reset()
				continue
			}
		}
		current.WriteRune(c)
	}

	if current.Len() > 0 {
		result = append(result, current.String())
	}

	return result
}

func (b *Bindgen) parseParamList(params string) []struct{ Name, Type string } {
	if params == "" {
		return nil
	}

	var result []struct{ Name, Type string }

	parts := b.splitParams(params)
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}

		if strings.Contains(part, ":") {
			colonIdx := strings.Index(part, ":")
			name := strings.TrimSpace(part[:colonIdx])
			typePart := strings.TrimSpace(part[colonIdx+1:])
			result = append(result, struct{ Name, Type string }{
				Name: name,
				Type: typePart,
			})
		} else {
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
	if strings.HasPrefix(name, "_") {
		return false
	}
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

	return os.WriteFile(outputPath, []byte(sb.String()), 0o644)
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

	return os.WriteFile(outputPath, []byte(sb.String()), 0o644)
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

	return os.WriteFile(outputPath, []byte(sb.String()), 0o644)
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

	return os.WriteFile(outputPath, []byte(sb.String()), 0o644)
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

	return os.WriteFile(outputPath, []byte(sb.String()), 0o644)
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

	sb.WriteString("// Functions NOT in WDK bindings (need manual declaration)\n\n")

	for _, name := range b.report.NotExported {
		funcs := b.externs.Functions[name]
		sb.WriteString(fmt.Sprintf("### %s\n", name))
		sb.WriteString("// Not exported by WDK - add to ntapi/mod.rs\n")
		for _, f := range funcs {
			sb.WriteString(fmt.Sprintf("//   - %s:%d\n", f.File, f.Line))
		}
		sb.WriteString("\n")
	}

	return os.WriteFile(outputPath, []byte(sb.String()), 0o644)
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

	sb.WriteString(fmt.Sprintf("// Exported functions: %d\n", len(exportedFuncs)))
	sb.WriteString("pub use wdk_sys::ntddk::{\n")

	for _, name := range exportedFuncs {
		sb.WriteString(fmt.Sprintf("    %s,\n", name))
	}
	sb.WriteString("};\n\n")

	sb.WriteString(fmt.Sprintf("// Not exported functions: %d\n", len(notExportedFunctions)))
	sb.WriteString("extern \"C\" {\n")

	for _, f := range notExportedFunctions {
		sb.WriteString(fmt.Sprintf("    pub fn %s(%s)", f.Name, f.Params))
		if f.ReturnType != "" && f.ReturnType != "void" {
			sb.WriteString(fmt.Sprintf(" -> %s", f.ReturnType))
		}
		sb.WriteString(";\n")
	}

	sb.WriteString("}\n")

	return os.WriteFile(filepath.Join(outputDir, "ntddk.rs"), []byte(sb.String()), 0o644)
}

func (b *Bindgen) generateNtapiTypes(outputDir string) error {
	var sb strings.Builder
	sb.WriteString("#![allow(non_snake_case)]\n")
	sb.WriteString("#![allow(dead_code)]\n\n")

	typesFromWdk := []string{}
	typesFromNtddk := []string{}
	for name, def := range b.wdk.Types {
		if def.Source == "ntddk" {
			typesFromNtddk = append(typesFromNtddk, name)
		} else {
			typesFromWdk = append(typesFromWdk, name)
		}
	}
	sort.Strings(typesFromWdk)
	sort.Strings(typesFromNtddk)

	sb.WriteString(fmt.Sprintf("// Types from wdk_sys: %d\n", len(typesFromWdk)))
	sb.WriteString(fmt.Sprintf("// Types from ntddk: %d\n\n", len(typesFromNtddk)))

	if len(typesFromWdk) > 0 {
		sb.WriteString("pub use wdk_sys::{\n")
		for _, name := range typesFromWdk {
			sb.WriteString(fmt.Sprintf("    %s,\n", name))
		}
		sb.WriteString("};\n\n")
	}

	if len(typesFromNtddk) > 0 {
		sb.WriteString("pub use wdk_sys::ntddk::{\n")
		for _, name := range typesFromNtddk {
			sb.WriteString(fmt.Sprintf("    %s,\n", name))
		}
		sb.WriteString("};\n")
	}

	return os.WriteFile(filepath.Join(outputDir, "types.rs"), []byte(sb.String()), 0o644)
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

	return os.WriteFile(filepath.Join(outputDir, "constants.rs"), []byte(sb.String()), 0o644)
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

	return os.WriteFile(outputPath, []byte(sb.String()), 0o644)
}

func (b *Bindgen) ApplyFixes() error {
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
		relPath, _ := filepath.Rel(b.config.ProjectRoot, file)
		fmt.Printf("Would modify: %s\n", relPath)
		for _, m := range mods {
			fmt.Printf("  Line %d: %s\n", m.Line, m.NewText)
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
				var indent strings.Builder
				for _, c := range line {
					if c == ' ' || c == '\t' {
						indent.WriteString(string(c))
					} else {
						break
					}
				}
				for j := i + 1; j < len(lines); j++ {
					if strings.Contains(lines[j], "};") {
						for k := i + 1; k <= j; k++ {
							linesToDelete[k] = true
						}
						lines[i] = indent.String() + "use crate::ntapi::*;"
						i = j
						break
					}
				}
			} else {
				var indent strings.Builder
				for _, c := range line {
					if c == ' ' || c == '\t' {
						indent.WriteString(string(c))
					} else {
						break
					}
				}
				lines[i] = indent.String() + "use crate::ntapi::*;"
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
	return os.WriteFile(filePath, []byte(newContent), 0o644)
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

type HookEntry struct {
	Name       string
	Params     string
	ReturnType string
	ParamList  []HookParam
}

type HookParam struct {
	Name string
	Type string
}

func (b *Bindgen) GenerateHookDatabase(rustOutputDir, goOutputDir string, notExportedFunctions []NotExportedFunc) error {
	exportedHooks := []HookEntry{}
	for name, sig := range b.wdk.Functions {
		if b.isNotExported(name, notExportedFunctions) {
			continue
		}
		rawParams := b.parseParamList(sig.Params)
		hookParams := make([]HookParam, len(rawParams))
		for i, p := range rawParams {
			hookParams[i] = HookParam{Name: p.Name, Type: p.Type}
		}
		hook := HookEntry{
			Name:       name,
			Params:     sig.Params,
			ReturnType: sig.ReturnType,
			ParamList:  hookParams,
		}
		exportedHooks = append(exportedHooks, hook)
	}
	sort.Slice(exportedHooks, func(i, j int) bool {
		return exportedHooks[i].Name < exportedHooks[j].Name
	})

	if err := b.generateRustHookDb(rustOutputDir, exportedHooks); err != nil {
		return err
	}

	if err := b.generateGoHookDb(goOutputDir, exportedHooks); err != nil {
		return err
	}

	return nil
}

func (b *Bindgen) generateRustHookDb(outputDir string, hooks []HookEntry) error {
	if err := b.generateRustHookEpt(outputDir, hooks); err != nil {
		return err
	}
	if err := b.generateRustHookInline(outputDir, hooks); err != nil {
		return err
	}
	return b.generateRustHookArgs(outputDir, hooks)
}

func (b *Bindgen) generateRustHookEpt(outputDir string, hooks []HookEntry) error {
	var sb strings.Builder
	sb.WriteString("#![allow(non_snake_case)]\n")
	sb.WriteString("#![allow(dead_code)]\n\n")

	sb.WriteString(fmt.Sprintf("// EPT Hook Database: %d functions\n", len(hooks)))
	sb.WriteString("// EPT hooks provide stealth by using Extended Page Tables\n\n")

	sb.WriteString("use alloc::string::String;\n")
	sb.WriteString("use alloc::vec::Vec;\n")
	sb.WriteString("use crate::hyperkd::hyperhv::hooks::hooks::*;\n")
	sb.WriteString("use crate::hyperkd::hyperhv::ProcessId;\n")
	sb.WriteString("use crate::ntapi::*;\n\n")

	sb.WriteString("pub struct EptHookDb {\n")
	sb.WriteString("    pub name: &'static str,\n")
	sb.WriteString("    pub address: u64,\n")
	sb.WriteString("    pub enabled: bool,\n")
	sb.WriteString("}\n\n")

	sb.WriteString("pub static EPT_HOOK_DATABASE: &[EptHookDb] = &[\n")
	for _, h := range hooks {
		sb.WriteString(fmt.Sprintf("    EptHookDb { name: \"%s\", address: 0, enabled: false },\n", h.Name))
	}
	sb.WriteString("];\n\n")

	sb.WriteString("pub fn install_ept_hook_by_name(name: &str, process_id: ProcessId) -> Result<(), HookError> {\n")
	sb.WriteString("    match name {\n")
	for _, h := range hooks {
		sb.WriteString(fmt.Sprintf("        \"%s\" => {\n", h.Name))
		sb.WriteString(fmt.Sprintf("            let addr = %s as u64;\n", h.Name))
		sb.WriteString("            unsafe { HOOK_CONTEXT.lock().set_ept_hook(addr, 0, process_id) }\n")
		sb.WriteString("        }\n")
	}
	sb.WriteString("        _ => Err(HookError::InvalidAddress),\n")
	sb.WriteString("    }\n")
	sb.WriteString("}\n\n")

	sb.WriteString("pub fn remove_ept_hook_by_name(name: &str) -> Result<(), HookError> {\n")
	sb.WriteString("    match name {\n")
	for _, h := range hooks {
		sb.WriteString(fmt.Sprintf("        \"%s\" => {\n", h.Name))
		sb.WriteString(fmt.Sprintf("            let addr = %s as u64;\n", h.Name))
		sb.WriteString("            unsafe { HOOK_CONTEXT.lock().remove_ept_hook(addr) }\n")
		sb.WriteString("        }\n")
	}
	sb.WriteString("        _ => Err(HookError::NotHooked),\n")
	sb.WriteString("    }\n")
	sb.WriteString("}\n")

	return os.WriteFile(filepath.Join(outputDir, "ept_hook.rs"), []byte(sb.String()), 0o644)
}

func (b *Bindgen) generateRustHookInline(outputDir string, hooks []HookEntry) error {
	var sb strings.Builder
	sb.WriteString("#![allow(non_snake_case)]\n")
	sb.WriteString("#![allow(dead_code)]\n\n")

	sb.WriteString(fmt.Sprintf("// Inline Hook Database: %d functions\n", len(hooks)))
	sb.WriteString("// Inline hooks modify code directly, faster but detectable\n\n")

	sb.WriteString("use alloc::string::String;\n")
	sb.WriteString("use crate::hyperkd::hyperhv::hooks::hooks::*;\n")
	sb.WriteString("use crate::ntapi::*;\n\n")

	sb.WriteString("pub struct InlineHookDb {\n")
	sb.WriteString("    pub name: &'static str,\n")
	sb.WriteString("    pub address: u64,\n")
	sb.WriteString("    pub trampoline: u64,\n")
	sb.WriteString("    pub enabled: bool,\n")
	sb.WriteString("}\n\n")

	sb.WriteString("pub static INLINE_HOOK_DATABASE: &[InlineHookDb] = &[\n")
	for _, h := range hooks {
		sb.WriteString(fmt.Sprintf("    InlineHookDb { name: \"%s\", address: 0, trampoline: 0, enabled: false },\n", h.Name))
	}
	sb.WriteString("];\n\n")

	sb.WriteString("pub fn install_inline_hook_by_name(name: &str, hook_handler: u64) -> Result<u64, HookError> {\n")
	sb.WriteString("    match name {\n")
	for _, h := range hooks {
		sb.WriteString(fmt.Sprintf("        \"%s\" => {\n", h.Name))
		sb.WriteString(fmt.Sprintf("            let addr = %s as u64;\n", h.Name))
		sb.WriteString("            unsafe { HOOK_CONTEXT.lock().set_inline_hook_with_trampoline(addr, hook_handler, 14) }\n")
		sb.WriteString("        }\n")
	}
	sb.WriteString("        _ => Err(HookError::InvalidAddress),\n")
	sb.WriteString("    }\n")
	sb.WriteString("}\n\n")

	sb.WriteString("pub fn remove_inline_hook_by_name(name: &str) -> Result<(), HookError> {\n")
	sb.WriteString("    match name {\n")
	for _, h := range hooks {
		sb.WriteString(fmt.Sprintf("        \"%s\" => {\n", h.Name))
		sb.WriteString(fmt.Sprintf("            let addr = %s as u64;\n", h.Name))
		sb.WriteString("            unsafe { HOOK_CONTEXT.lock().remove_inline_hook(addr) }\n")
		sb.WriteString("        }\n")
	}
	sb.WriteString("        _ => Err(HookError::NotHooked),\n")
	sb.WriteString("    }\n")
	sb.WriteString("}\n")

	return os.WriteFile(filepath.Join(outputDir, "inline_hook.rs"), []byte(sb.String()), 0o644)
}

func (b *Bindgen) generateRustHookArgs(outputDir string, hooks []HookEntry) error {
	var sb strings.Builder
	sb.WriteString("#![allow(non_snake_case)]\n")
	sb.WriteString("#![allow(dead_code)]\n\n")

	sb.WriteString(fmt.Sprintf("// Hook Args structs: %d functions\n", len(hooks)))
	sb.WriteString("// Auto-generated by bindgen - DO NOT EDIT\n\n")

	sb.WriteString("use crate::ntapi::*;\n\n")

	sb.WriteString("// Args structs for each hooked API\n")
	sb.WriteString("// All types are directly from wdk_sys via ntapi::types\n\n")
	for _, h := range hooks {
		if len(h.ParamList) == 0 {
			continue
		}
		sb.WriteString(fmt.Sprintf("/// Arguments for %s\n", h.Name))
		sb.WriteString("#[repr(C)]\n")
		sb.WriteString(fmt.Sprintf("pub struct %sArgs {\n", h.Name))
		for _, p := range h.ParamList {
			if p.Name == "" || p.Type == "..." {
				continue
			}
			sb.WriteString(fmt.Sprintf("    pub %s: %s,\n", p.Name, p.Type))
		}
		sb.WriteString("}\n\n")
	}

	sb.WriteString("// Args registry for runtime lookup\n")
	sb.WriteString("pub static HOOK_ARGS_REGISTRY: &[(&str, usize)] = &[\n")
	for _, h := range hooks {
		if len(h.ParamList) == 0 {
			continue
		}
		sb.WriteString(fmt.Sprintf("    (\"%s\", core::mem::size_of::< %sArgs>()),\n", h.Name, h.Name))
	}
	sb.WriteString("];\n\n")

	sb.WriteString("pub fn get_args_size(name: &str) -> Option<usize> {\n")
	sb.WriteString("    for (n, size) in HOOK_ARGS_REGISTRY {\n")
	sb.WriteString("        if *n == name {\n")
	sb.WriteString("            return Some(*size);\n")
	sb.WriteString("        }\n")
	sb.WriteString("    }\n")
	sb.WriteString("    None\n")
	sb.WriteString("}\n")

	return os.WriteFile(filepath.Join(outputDir, "hook_args.rs"), []byte(sb.String()), 0o644)
}

func (b *Bindgen) generateGoHookDb(outputPath string, hooks []HookEntry) error {
	var sb strings.Builder
	sb.WriteString("package debugger\n\n")
	sb.WriteString("import (\n")
	sb.WriteString("    \"encoding/json\"\n")
	sb.WriteString("    \"fmt\"\n")
	sb.WriteString("    \"reflect\"\n")
	sb.WriteString("    \"strings\"\n")
	sb.WriteString(")\n\n")

	sb.WriteString(fmt.Sprintf("// Hook Database: %d NT API functions\n", len(hooks)))
	sb.WriteString("// Auto-generated by bindgen - DO NOT EDIT\n\n")

	sb.WriteString("// HookType defines the hook method\n")
	sb.WriteString("type HookType int\n\n")
	sb.WriteString("const (\n")
	sb.WriteString("    HookTypeEPT HookType = iota    // EPT hook - stealth but slower\n")
	sb.WriteString("    HookTypeInline                  // Inline hook - fast but detectable\n")
	sb.WriteString(")\n\n")

	sb.WriteString("type HookInfo struct {\n")
	sb.WriteString("    Name       string      `json:\"name\"`\n")
	sb.WriteString("    Params     []HookParam `json:\"params\"`\n")
	sb.WriteString("    ReturnType string      `json:\"return_type\"`\n")
	sb.WriteString("    GoType     string      `json:\"go_type\"`     // Go return type\n")
	sb.WriteString("}\n\n")

	sb.WriteString("type HookParam struct {\n")
	sb.WriteString("    Name   string `json:\"name\"`\n")
	sb.WriteString("    Type   string `json:\"type\"`    // Rust type\n")
	sb.WriteString("    GoType string `json:\"go_type\"` // Go type for user\n")
	sb.WriteString("}\n\n")

	sb.WriteString("type HookRequest struct {\n")
	sb.WriteString("    Action      string   `json:\"action\"`\n")
	sb.WriteString("    HookType    HookType `json:\"hook_type\"`\n")
	sb.WriteString("    ApiName     string   `json:\"api_name\"`\n")
	sb.WriteString("    ProcessId   uint32   `json:\"process_id,omitempty\"`\n")
	sb.WriteString("    HookHandler uint64   `json:\"hook_handler,omitempty\"`\n")
	sb.WriteString("}\n\n")

	sb.WriteString("type HookResponse struct {\n")
	sb.WriteString("    Success    bool   `json:\"success\"`\n")
	sb.WriteString("    Message    string `json:\"message\"`\n")
	sb.WriteString("    Trampoline uint64 `json:\"trampoline,omitempty\"`\n")
	sb.WriteString("}\n\n")

	sb.WriteString("// Rust to Go type mapping\n")
	sb.WriteString("var rustToGoTypes = map[string]string{\n")
	sb.WriteString("    \"void\": \"\",\n")
	sb.WriteString("    \"NTSTATUS\": \"int32\",\n")
	sb.WriteString("    \"BOOLEAN\": \"bool\",\n")
	sb.WriteString("    \"ULONG\": \"uint32\",\n")
	sb.WriteString("    \"PULONG\": \"*uint32\",\n")
	sb.WriteString("    \"USHORT\": \"uint16\",\n")
	sb.WriteString("    \"PUSHORT\": \"*uint16\",\n")
	sb.WriteString("    \"UCHAR\": \"uint8\",\n")
	sb.WriteString("    \"PUCHAR\": \"*uint8\",\n")
	sb.WriteString("    \"LONG\": \"int32\",\n")
	sb.WriteString("    \"PLONG\": \"*int32\",\n")
	sb.WriteString("    \"SHORT\": \"int16\",\n")
	sb.WriteString("    \"PSHORT\": \"*int16\",\n")
	sb.WriteString("    \"CHAR\": \"int8\",\n")
	sb.WriteString("    \"PCHAR\": \"*int8\",\n")
	sb.WriteString("    \"SIZE_T\": \"uintptr\",\n")
	sb.WriteString("    \"PSIZE_T\": \"*uintptr\",\n")
	sb.WriteString("    \"ULONG_PTR\": \"uintptr\",\n")
	sb.WriteString("    \"PULONG_PTR\": \"*uintptr\",\n")
	sb.WriteString("    \"ULONG64\": \"uint64\",\n")
	sb.WriteString("    \"PULONG64\": \"*uint64\",\n")
	sb.WriteString("    \"ULONGLONG\": \"uint64\",\n")
	sb.WriteString("    \"LONGLONG\": \"int64\",\n")
	sb.WriteString("    \"PLONGLONG\": \"*int64\",\n")
	sb.WriteString("    \"DWORD\": \"uint32\",\n")
	sb.WriteString("    \"PDWORD\": \"*uint32\",\n")
	sb.WriteString("    \"DWORD64\": \"uint64\",\n")
	sb.WriteString("    \"PDWORD64\": \"*uint64\",\n")
	sb.WriteString("    \"UINT\": \"uint32\",\n")
	sb.WriteString("    \"PUINT\": \"*uint32\",\n")
	sb.WriteString("    \"INT\": \"int32\",\n")
	sb.WriteString("    \"PINT\": \"*int32\",\n")
	sb.WriteString("    \"UINT_PTR\": \"uintptr\",\n")
	sb.WriteString("    \"INT_PTR\": \"intptr\",\n")
	sb.WriteString("    \"HANDLE\": \"uintptr\",\n")
	sb.WriteString("    \"PHANDLE\": \"*uintptr\",\n")
	sb.WriteString("    \"PVOID\": \"uintptr\",\n")
	sb.WriteString("    \"PPVOID\": \"*uintptr\",\n")
	sb.WriteString("    \"LPVOID\": \"uintptr\",\n")
	sb.WriteString("    \"LPCVOID\": \"uintptr\",\n")
	sb.WriteString("    \"PCSTR\": \"uintptr\",\n")
	sb.WriteString("    \"PSTR\": \"uintptr\",\n")
	sb.WriteString("    \"PCWSTR\": \"uintptr\",\n")
	sb.WriteString("    \"PWSTR\": \"uintptr\",\n")
	sb.WriteString("    \"LPCSTR\": \"uintptr\",\n")
	sb.WriteString("    \"LPSTR\": \"uintptr\",\n")
	sb.WriteString("    \"LPCWSTR\": \"uintptr\",\n")
	sb.WriteString("    \"LPWSTR\": \"uintptr\",\n")
	sb.WriteString("    \"KIRQL\": \"uint8\",\n")
	sb.WriteString("    \"PKIRQL\": \"*uint8\",\n")
	sb.WriteString("    \"KPRIORITY\": \"int32\",\n")
	sb.WriteString("    \"KAFFINITY\": \"uintptr\",\n")
	sb.WriteString("    \"LARGE_INTEGER\": \"int64\",\n")
	sb.WriteString("    \"PLARGE_INTEGER\": \"*int64\",\n")
	sb.WriteString("    \"ULARGE_INTEGER\": \"uint64\",\n")
	sb.WriteString("    \"PULARGE_INTEGER\": \"*uint64\",\n")
	sb.WriteString("    \"GUID\": \"[16]byte\",\n")
	sb.WriteString("    \"PGUID\": \"*[16]byte\",\n")
	sb.WriteString("    \"PEPROCESS\": \"uintptr\",\n")
	sb.WriteString("    \"PETHREAD\": \"uintptr\",\n")
	sb.WriteString("    \"PKTHREAD\": \"uintptr\",\n")
	sb.WriteString("    \"PKPROCESS\": \"uintptr\",\n")
	sb.WriteString("    \"PDRIVER_OBJECT\": \"uintptr\",\n")
	sb.WriteString("    \"PDEVICE_OBJECT\": \"uintptr\",\n")
	sb.WriteString("    \"PIRP\": \"uintptr\",\n")
	sb.WriteString("    \"PFILE_OBJECT\": \"uintptr\",\n")
	sb.WriteString("    \"PSID\": \"uintptr\",\n")
	sb.WriteString("    \"PEX_RUNDOWN_REF\": \"*uintptr\",\n")
	sb.WriteString("    \"EX_SPIN_LOCK\": \"uintptr\",\n")
	sb.WriteString("    \"PEX_SPIN_LOCK\": \"*uintptr\",\n")
	sb.WriteString("    \"KSPIN_LOCK\": \"uintptr\",\n")
	sb.WriteString("    \"PKSPIN_LOCK\": \"*uintptr\",\n")
	sb.WriteString("    \"PERESOURCE\": \"uintptr\",\n")
	sb.WriteString("    \"PKSEMAPHORE\": \"uintptr\",\n")
	sb.WriteString("    \"PKEVENT\": \"uintptr\",\n")
	sb.WriteString("    \"PKMUTANT\": \"uintptr\",\n")
	sb.WriteString("    \"PKDPC\": \"uintptr\",\n")
	sb.WriteString("    \"PKTIMER\": \"uintptr\",\n")
	sb.WriteString("    \"PIO_STATUS_BLOCK\": \"uintptr\",\n")
	sb.WriteString("    \"POBJECT_ATTRIBUTES\": \"uintptr\",\n")
	sb.WriteString("    \"PUNICODE_STRING\": \"uintptr\",\n")
	sb.WriteString("    \"PANSI_STRING\": \"uintptr\",\n")
	sb.WriteString("    \"PCLIENT_ID\": \"uintptr\",\n")
	sb.WriteString("    \"PACL\": \"uintptr\",\n")
	sb.WriteString("    \"PSECURITY_DESCRIPTOR\": \"uintptr\",\n")
	sb.WriteString("    \"PLIST_ENTRY\": \"uintptr\",\n")
	sb.WriteString("    \"PSINGLE_LIST_ENTRY\": \"uintptr\",\n")
	sb.WriteString("    \"...\": \"interface{}\",\n")
	sb.WriteString("}\n\n")

	sb.WriteString("func rustTypeToGo(rustType string) string {\n")
	sb.WriteString("    rustType = trimPointerType(rustType)\n")
	sb.WriteString("    if goType, ok := rustToGoTypes[rustType]; ok {\n")
	sb.WriteString("        return goType\n")
	sb.WriteString("    }\n")
	sb.WriteString("    if strings.HasPrefix(rustType, \"P\") {\n")
	sb.WriteString("        return \"uintptr\"\n")
	sb.WriteString("    }\n")
	sb.WriteString("    return \"uintptr\"\n")
	sb.WriteString("}\n\n")

	sb.WriteString("func trimPointerType(t string) string {\n")
	sb.WriteString("    t = strings.TrimSpace(t)\n")
	sb.WriteString("    for strings.HasPrefix(t, \"*mut \") || strings.HasPrefix(t, \"*const \") {\n")
	sb.WriteString("        if strings.HasPrefix(t, \"*mut \") {\n")
	sb.WriteString("            t = strings.TrimPrefix(t, \"*mut \")\n")
	sb.WriteString("        } else {\n")
	sb.WriteString("            t = strings.TrimPrefix(t, \"*const \")\n")
	sb.WriteString("        }\n")
	sb.WriteString("    }\n")
	sb.WriteString("    return t\n")
	sb.WriteString("}\n\n")

	sb.WriteString(fmt.Sprintf("var HookDatabase = []HookInfo{\n"))
	for _, h := range hooks {
		goReturnType := rustTypeToGo(h.ReturnType)
		sb.WriteString("    {\n")
		sb.WriteString(fmt.Sprintf("        Name: \"%s\",\n", h.Name))
		sb.WriteString("        Params: []HookParam{\n")
		for _, p := range h.ParamList {
			goType := rustTypeToGo(p.Type)
			sb.WriteString(fmt.Sprintf("            {Name: \"%s\", Type: \"%s\", GoType: \"%s\"},\n", p.Name, p.Type, goType))
		}
		sb.WriteString("        },\n")
		sb.WriteString(fmt.Sprintf("        ReturnType: \"%s\",\n", h.ReturnType))
		sb.WriteString(fmt.Sprintf("        GoType: \"%s\",\n", goReturnType))
		sb.WriteString("    },\n")
	}
	sb.WriteString("}\n\n")

	sb.WriteString("// Auto-generated Args structs for each API\n")
	for _, h := range hooks {
		if len(h.ParamList) == 0 {
			continue
		}
		sb.WriteString(fmt.Sprintf("type %sArgs struct {\n", h.Name))
		for _, p := range h.ParamList {
			if p.Name == "" || p.Type == "..." {
				continue
			}
			goType := rustTypeToGo(p.Type)
			fieldName := strings.Title(p.Name)
			sb.WriteString(fmt.Sprintf("    %s %s `json:\"%s\"`\n", fieldName, goType, p.Name))
		}
		sb.WriteString("}\n\n")
	}

	sb.WriteString("// ArgsRegistry maps API names to their Args struct types\n")
	sb.WriteString("var ArgsRegistry = map[string]reflect.Type{\n")
	for _, h := range hooks {
		if len(h.ParamList) == 0 {
			continue
		}
		sb.WriteString(fmt.Sprintf("    \"%s\": reflect.TypeOf(%sArgs{}),\n", h.Name, h.Name))
	}
	sb.WriteString("}\n\n")

	sb.WriteString("func GetHookInfo(name string) *HookInfo {\n")
	sb.WriteString("    for i := range HookDatabase {\n")
	sb.WriteString("        if HookDatabase[i].Name == name {\n")
	sb.WriteString("            return &HookDatabase[i]\n")
	sb.WriteString("        }\n")
	sb.WriteString("    }\n")
	sb.WriteString("    return nil\n")
	sb.WriteString("}\n\n")

	sb.WriteString("func (p *Packet) InstallHook(apiName string, hookType HookType, processId uint32, hookHandler uint64) (*Response[HookResponse], error) {\n")
	sb.WriteString("    req := HookRequest{\n")
	sb.WriteString("        Action:      \"install_hook\",\n")
	sb.WriteString("        HookType:    hookType,\n")
	sb.WriteString("        ApiName:     apiName,\n")
	sb.WriteString("        ProcessId:   processId,\n")
	sb.WriteString("        HookHandler: hookHandler,\n")
	sb.WriteString("    }\n")
	sb.WriteString("    data, _ := json.Marshal(req)\n")
	sb.WriteString("    resp := SendReceive[HookResponse](p, data)\n")
	sb.WriteString("    if resp == nil || !resp.Success {\n")
	sb.WriteString("        return nil, fmt.Errorf(\"install hook failed: %s\", resp.Message)\n")
	sb.WriteString("    }\n")
	sb.WriteString("    return resp, nil\n")
	sb.WriteString("}\n\n")

	sb.WriteString("func (p *Packet) RemoveHook(apiName string, hookType HookType) error {\n")
	sb.WriteString("    req := HookRequest{\n")
	sb.WriteString("        Action:   \"remove_hook\",\n")
	sb.WriteString("        HookType: hookType,\n")
	sb.WriteString("        ApiName:  apiName,\n")
	sb.WriteString("    }\n")
	sb.WriteString("    data, _ := json.Marshal(req)\n")
	sb.WriteString("    resp := SendReceive[HookResponse](p, data)\n")
	sb.WriteString("    if resp == nil || !resp.Success {\n")
	sb.WriteString("        return fmt.Errorf(\"remove hook failed: %s\", resp.Message)\n")
	sb.WriteString("    }\n")
	sb.WriteString("    return nil\n")
	sb.WriteString("}\n\n")

	sb.WriteString("func (p *Packet) ListHooks() []HookInfo {\n")
	sb.WriteString("    return HookDatabase\n")
	sb.WriteString("}\n\n")

	sb.WriteString("// HookCallContext provides context for hook callback\n")
	sb.WriteString("type HookCallContext struct {\n")
	sb.WriteString("    ApiName    string\n")
	sb.WriteString("    ProcessId  uint32\n")
	sb.WriteString("    ThreadId   uint32\n")
	sb.WriteString("    Registers  map[string]uint64\n")
	sb.WriteString("    StackPtr   uint64\n")
	sb.WriteString("    RetAddr    uint64\n")
	sb.WriteString("}\n\n")

	sb.WriteString("// HookCallback is called when a hooked API is invoked\n")
	sb.WriteString("type HookCallback func(ctx *HookCallContext) error\n\n")

	sb.WriteString("// HookManager manages hook registration and callbacks\n")
	sb.WriteString("type HookManager struct {\n")
	sb.WriteString("    hooks     map[string]*HookEntry\n")
	sb.WriteString("    callbacks map[string]HookCallback\n")
	sb.WriteString("}\n\n")

	sb.WriteString("type HookEntry struct {\n")
	sb.WriteString("    Info       *HookInfo\n")
	sb.WriteString("    Type       HookType\n")
	sb.WriteString("    ProcessId  uint32\n")
	sb.WriteString("    Handler    uint64\n")
	sb.WriteString("    Callback   HookCallback\n")
	sb.WriteString("    Enabled    bool\n")
	sb.WriteString("}\n\n")

	sb.WriteString("func NewHookManager() *HookManager {\n")
	sb.WriteString("    return &HookManager{\n")
	sb.WriteString("        hooks:     make(map[string]*HookEntry),\n")
	sb.WriteString("        callbacks: make(map[string]HookCallback),\n")
	sb.WriteString("    }\n")
	sb.WriteString("}\n\n")

	sb.WriteString("func (m *HookManager) Register(apiName string, hookType HookType, processId uint32, callback HookCallback) error {\n")
	sb.WriteString("    info := GetHookInfo(apiName)\n")
	sb.WriteString("    if info == nil {\n")
	sb.WriteString("        return fmt.Errorf(\"unknown API: %s\", apiName)\n")
	sb.WriteString("    }\n")
	sb.WriteString("    m.hooks[apiName] = &HookEntry{\n")
	sb.WriteString("        Info:      info,\n")
	sb.WriteString("        Type:      hookType,\n")
	sb.WriteString("        ProcessId: processId,\n")
	sb.WriteString("        Callback:  callback,\n")
	sb.WriteString("        Enabled:   true,\n")
	sb.WriteString("    }\n")
	sb.WriteString("    return nil\n")
	sb.WriteString("}\n\n")

	sb.WriteString("func (m *HookManager) Unregister(apiName string) {\n")
	sb.WriteString("    delete(m.hooks, apiName)\n")
	sb.WriteString("    delete(m.callbacks, apiName)\n")
	sb.WriteString("}\n\n")

	sb.WriteString("func (m *HookManager) GetHook(apiName string) *HookEntry {\n")
	sb.WriteString("    return m.hooks[apiName]\n")
	sb.WriteString("}\n\n")

	sb.WriteString("func (m *HookManager) ListActiveHooks() []string {\n")
	sb.WriteString("    var result []string\n")
	sb.WriteString("    for name, entry := range m.hooks {\n")
	sb.WriteString("        if entry.Enabled {\n")
	sb.WriteString("            result = append(result, name)\n")
	sb.WriteString("        }\n")
	sb.WriteString("    }\n")
	sb.WriteString("    return result\n")
	sb.WriteString("}\n")

	return os.WriteFile(outputPath, []byte(sb.String()), 0o644)
}

func rustTypeToGo(rustType string) string {
	rustType = strings.TrimSpace(rustType)
	for strings.HasPrefix(rustType, "*mut ") || strings.HasPrefix(rustType, "*const ") {
		if after, ok := strings.CutPrefix(rustType, "*mut "); ok {
			rustType = after
		} else {
			rustType = strings.TrimPrefix(rustType, "*const ")
		}
	}
	rustType = strings.TrimSpace(rustType)

	typeMap := map[string]string{
		"void":                 "",
		"NTSTATUS":             "int32",
		"BOOLEAN":              "bool",
		"ULONG":                "uint32",
		"PULONG":               "*uint32",
		"USHORT":               "uint16",
		"PUSHORT":              "*uint16",
		"UCHAR":                "uint8",
		"PUCHAR":               "*uint8",
		"LONG":                 "int32",
		"PLONG":                "*int32",
		"SHORT":                "int16",
		"PSHORT":               "*int16",
		"CHAR":                 "int8",
		"PCHAR":                "*int8",
		"SIZE_T":               "uintptr",
		"PSIZE_T":              "*uintptr",
		"ULONG_PTR":            "uintptr",
		"PULONG_PTR":           "*uintptr",
		"ULONG64":              "uint64",
		"PULONG64":             "*uint64",
		"ULONGLONG":            "uint64",
		"LONGLONG":             "int64",
		"PLONGLONG":            "*int64",
		"DWORD":                "uint32",
		"PDWORD":               "*uint32",
		"DWORD64":              "uint64",
		"PDWORD64":             "*uint64",
		"UINT":                 "uint32",
		"PUINT":                "*uint32",
		"INT":                  "int32",
		"PINT":                 "*int32",
		"UINT_PTR":             "uintptr",
		"INT_PTR":              "intptr",
		"HANDLE":               "uintptr",
		"PHANDLE":              "*uintptr",
		"PVOID":                "uintptr",
		"PPVOID":               "*uintptr",
		"LPVOID":               "uintptr",
		"LPCVOID":              "uintptr",
		"PCSTR":                "uintptr",
		"PSTR":                 "uintptr",
		"PCWSTR":               "uintptr",
		"PWSTR":                "uintptr",
		"LPCSTR":               "uintptr",
		"LPSTR":                "uintptr",
		"LPCWSTR":              "uintptr",
		"LPWSTR":               "uintptr",
		"KIRQL":                "uint8",
		"PKIRQL":               "*uint8",
		"KPRIORITY":            "int32",
		"KAFFINITY":            "uintptr",
		"LARGE_INTEGER":        "int64",
		"PLARGE_INTEGER":       "*int64",
		"PEPROCESS":            "uintptr",
		"PETHREAD":             "uintptr",
		"PKTHREAD":             "uintptr",
		"PKPROCESS":            "uintptr",
		"PDRIVER_OBJECT":       "uintptr",
		"PDEVICE_OBJECT":       "uintptr",
		"PIRP":                 "uintptr",
		"PFILE_OBJECT":         "uintptr",
		"PSID":                 "uintptr",
		"PERESOURCE":           "uintptr",
		"PEX_SPIN_LOCK":        "*uintptr",
		"PKSEMAPHORE":          "uintptr",
		"PKEVENT":              "uintptr",
		"PKMUTANT":             "uintptr",
		"PKDPC":                "uintptr",
		"PKTIMER":              "uintptr",
		"PKSPIN_LOCK":          "*uintptr",
		"PIO_STATUS_BLOCK":     "uintptr",
		"POBJECT_ATTRIBUTES":   "uintptr",
		"PUNICODE_STRING":      "uintptr",
		"PANSI_STRING":         "uintptr",
		"PCLIENT_ID":           "uintptr",
		"PACL":                 "uintptr",
		"PSECURITY_DESCRIPTOR": "uintptr",
		"PLIST_ENTRY":          "uintptr",
		"PSINGLE_LIST_ENTRY":   "uintptr",
		"...":                  "interface{}",
	}

	if goType, ok := typeMap[rustType]; ok {
		return goType
	}

	if strings.HasPrefix(rustType, "P") && len(rustType) > 1 {
		return "uintptr"
	}

	return "uintptr"
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

func GenerateBindgen(projectRoot string) error {
	bindgenDir := filepath.Join(projectRoot, "cmd", "rustgen")
	rustSourceDir := filepath.Join(projectRoot, "rust-driver", "kd", "src")
	outputDir := filepath.Join(projectRoot, "rust-driver", "kd", "src", "generated")

	if err := os.MkdirAll(outputDir, 0o755); err != nil {
		return fmt.Errorf("failed to create output dir: %w", err)
	}

	config := BindgenConfig{
		ProjectRoot:    projectRoot,
		WdkBindingsDir: bindgenDir,
		RustSourceDir:  rustSourceDir,
		OutputDir:      outputDir,
		Verbose:        true,
	}

	bg := NewBindgen(config)

	ntddkPath := filepath.Join(bindgenDir, "ntddk.rs")
	typesPath := filepath.Join(bindgenDir, "types.rs")
	constantsPath := filepath.Join(bindgenDir, "constants.rs")

	if err := bg.LoadWdkBindings(ntddkPath, typesPath, constantsPath); err != nil {
		return fmt.Errorf("failed to load WDK bindings: %w", err)
	}

	wdk := bg.GetWdkBindings()
	fmt.Printf("  Loaded WDK bindings: %d functions, %d types, %d constants\n",
		len(wdk.Functions), len(wdk.Types), len(wdk.Constants))

	if err := bg.ScanRustExterns(rustSourceDir); err != nil {
		return fmt.Errorf("failed to scan Rust externs: %w", err)
	}

	externs := bg.GetRustExterns()
	fmt.Printf("  Scanned Rust externs: %d functions in %d files\n",
		len(externs.Functions), len(externs.Files))

	report := bg.Validate()
	fmt.Printf("  Validation: %d mismatches, %d not exported\n",
		report.Statistics["mismatches"], report.Statistics["not_exported"])

	notExportedFuncs := make([]NotExportedFunc, len(report.NotExported))
	for i, name := range report.NotExported {
		externs := bg.externs.Functions[name]
		if len(externs) > 0 {
			notExportedFuncs[i] = NotExportedFunc{
				Name:       name,
				Params:     externs[0].Params,
				ReturnType: externs[0].ReturnType,
			}
		}
	}

	if err := bg.generateNtapiApi(outputDir, notExportedFuncs); err != nil {
		return fmt.Errorf("failed to generate ntapi api: %w", err)
	}

	if err := bg.generateNtapiTypes(outputDir); err != nil {
		return fmt.Errorf("failed to generate ntapi types: %w", err)
	}

	if err := bg.generateNtapiConstants(outputDir); err != nil {
		return fmt.Errorf("failed to generate ntapi constants: %w", err)
	}

	if err := bg.GenerateFunctionsList(filepath.Join(outputDir, "ntddk_list.txt")); err != nil {
		return fmt.Errorf("failed to generate functions list: %w", err)
	}

	if err := bg.GenerateFunctionsGoMap(filepath.Join(outputDir, "ntddk.go")); err != nil {
		return fmt.Errorf("failed to generate functions map: %w", err)
	}

	if err := bg.GenerateTypesGoMap(filepath.Join(outputDir, "types.go")); err != nil {
		return fmt.Errorf("failed to generate types map: %w", err)
	}

	if err := bg.GenerateConstantsGoMap(filepath.Join(outputDir, "constants.go")); err != nil {
		return fmt.Errorf("failed to generate constants map: %w", err)
	}

	if err := bg.GenerateValidationReport(filepath.Join(outputDir, "validation_report.txt")); err != nil {
		return fmt.Errorf("failed to generate validation report: %w", err)
	}

	fmt.Printf("  Scanning for WDK usage...\n")
	excludeDirs := []string{"cmd", "target", "todo"}
	if err := bg.ScanProjectUsage(projectRoot, excludeDirs); err != nil {
		return fmt.Errorf("failed to scan project usage: %w", err)
	}

	if err := bg.GenerateUsageReport(filepath.Join(outputDir, "wdk_usage_report.txt")); err != nil {
		return fmt.Errorf("failed to generate usage report: %w", err)
	}

	fmt.Printf("  Generating hook database...\n")
	rustHookDir := filepath.Join(projectRoot, "rust-driver", "kd", "src", "generated")
	goHookFile := filepath.Join(projectRoot, "debugger", "hook_db.go")
	if err := bg.GenerateHookDatabase(rustHookDir, goHookFile, notExportedFuncs); err != nil {
		return fmt.Errorf("failed to generate hook database: %w", err)
	}

	fmt.Printf("  Generated files in: %s\n", outputDir)

	return nil
}
