package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func getProjectRoot(t *testing.T) string {
	t.Helper()
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get cwd: %v", err)
	}
	return filepath.Join(cwd, "..", "..")
}

func getBindgenDir(t *testing.T) string {
	t.Helper()
	return filepath.Join(getProjectRoot(t), "cmd", "bindgen")
}

func TestLoadWdkBindings(t *testing.T) {
	bindgenDir := getBindgenDir(t)

	config := BindgenConfig{
		ProjectRoot:    getProjectRoot(t),
		WdkBindingsDir: bindgenDir,
		OutputDir:      bindgenDir,
		Verbose:        true,
	}

	bg := NewBindgen(config)

	err := bg.LoadWdkBindings(
		filepath.Join(bindgenDir, "ntddk.rs"),
		filepath.Join(bindgenDir, "types.rs"),
		filepath.Join(bindgenDir, "constants.rs"),
	)
	if err != nil {
		t.Fatalf("Failed to load WDK bindings: %v", err)
	}

	wdk := bg.GetWdkBindings()

	t.Logf("Loaded WDK bindings:")
	t.Logf("  Functions: %d", len(wdk.Functions))
	t.Logf("  Types: %d", len(wdk.Types))
	t.Logf("  Constants: %d", len(wdk.Constants))

	if len(wdk.Functions) == 0 {
		t.Error("No functions loaded from ntddk.rs")
	}

	for _, prefix := range []string{"Ps", "Mm", "Ke", "Ob"} {
		count := 0
		for name := range wdk.Functions {
			if len(name) >= 2 && name[:2] == prefix {
				count++
			}
		}
		t.Logf("  %s functions: %d", prefix, count)
	}
}

func TestScanRustExterns(t *testing.T) {
	projectRoot := getProjectRoot(t)
	bindgenDir := getBindgenDir(t)

	config := BindgenConfig{
		ProjectRoot:    projectRoot,
		WdkBindingsDir: bindgenDir,
		RustSourceDir:  filepath.Join(projectRoot, "rust-driver", "kd", "src"),
		OutputDir:      bindgenDir,
		Verbose:        true,
	}

	bg := NewBindgen(config)

	err := bg.LoadWdkBindings(
		filepath.Join(bindgenDir, "ntddk.rs"),
		filepath.Join(bindgenDir, "types.rs"),
		filepath.Join(bindgenDir, "constants.rs"),
	)
	if err != nil {
		t.Fatalf("Failed to load WDK bindings: %v", err)
	}

	err = bg.ScanRustExterns(config.RustSourceDir)
	if err != nil {
		t.Fatalf("Failed to scan Rust externs: %v", err)
	}

	externs := bg.GetRustExterns()

	t.Logf("Scanned Rust extern functions:")
	t.Logf("  Total: %d", len(externs.Functions))
	t.Logf("  Files scanned: %d", len(externs.Files))

	missingCount := 0
	for _, funcs := range externs.Functions {
		if !funcs[0].FoundInWdk {
			missingCount++
		}
	}
	t.Logf("  Missing from WDK: %d", missingCount)
}

func TestValidate(t *testing.T) {
	projectRoot := getProjectRoot(t)
	bindgenDir := getBindgenDir(t)

	config := BindgenConfig{
		ProjectRoot:    projectRoot,
		WdkBindingsDir: bindgenDir,
		RustSourceDir:  filepath.Join(projectRoot, "rust-driver", "kd", "src"),
		OutputDir:      bindgenDir,
		Verbose:        true,
	}

	bg := NewBindgen(config)

	err := bg.LoadWdkBindings(
		filepath.Join(bindgenDir, "ntddk.rs"),
		filepath.Join(bindgenDir, "types.rs"),
		filepath.Join(bindgenDir, "constants.rs"),
	)
	if err != nil {
		t.Fatalf("Failed to load WDK bindings: %v", err)
	}

	err = bg.ScanRustExterns(config.RustSourceDir)
	if err != nil {
		t.Fatalf("Failed to scan Rust externs: %v", err)
	}

	report := bg.Validate()

	t.Logf("Validation report:")
	t.Logf("  Total WDK Functions: %d", report.Statistics["total_wdk_functions"])
	t.Logf("  Total WDK Types: %d", report.Statistics["total_wdk_types"])
	t.Logf("  Total WDK Constants: %d", report.Statistics["total_wdk_constants"])
	t.Logf("  Total Rust Externs: %d", report.Statistics["total_externs"])
	t.Logf("  Type Mismatches: %d", report.Statistics["mismatches"])
	t.Logf("  Not Exported by WDK: %d", report.Statistics["not_exported"])

	if len(report.Mismatches) > 0 {
		t.Logf("\n=== TYPE MISMATCHES DETECTED ===")
		for _, m := range report.Mismatches {
			t.Logf("  %s:%d: %s - %s", m.RustFile, m.RustLine, m.FunctionName, m.MismatchType)
		}
	}

	if len(report.NotExported) > 0 {
		t.Logf("\n=== FUNCTIONS NOT EXPORTED BY WDK ===")
		for _, name := range report.NotExported {
			funcs := bg.GetRustExterns().Functions[name]
			t.Logf("  %s (%s)", name, funcs[0].File)
		}
	}
}

func TestGenerateAll(t *testing.T) {
	projectRoot := getProjectRoot(t)
	bindgenDir := getBindgenDir(t)

	config := BindgenConfig{
		ProjectRoot:    projectRoot,
		WdkBindingsDir: bindgenDir,
		RustSourceDir:  filepath.Join(projectRoot, "rust-driver", "kd", "src"),
		OutputDir:      bindgenDir,
		Verbose:        true,
	}

	bg := NewBindgen(config)

	err := bg.LoadWdkBindings(
		filepath.Join(bindgenDir, "ntddk.rs"),
		filepath.Join(bindgenDir, "types.rs"),
		filepath.Join(bindgenDir, "constants.rs"),
	)
	if err != nil {
		t.Fatalf("Failed to load WDK bindings: %v", err)
	}

	err = bg.ScanRustExterns(config.RustSourceDir)
	if err != nil {
		t.Fatalf("Failed to scan Rust externs: %v", err)
	}

	report := bg.Validate()

	t.Logf("Generating output files...")

	if err := bg.GenerateFunctionsList(filepath.Join(bindgenDir, "functions_list.txt")); err != nil {
		t.Fatalf("Failed to generate functions list: %v", err)
	}
	t.Logf("  - functions_list.txt")

	if err := bg.GenerateFunctionsGoMap(filepath.Join(bindgenDir, "functions_gen.go")); err != nil {
		t.Fatalf("Failed to generate functions map: %v", err)
	}
	t.Logf("  - functions_gen.go")

	if err := bg.GenerateTypesGoMap(filepath.Join(bindgenDir, "types_gen.go")); err != nil {
		t.Fatalf("Failed to generate types map: %v", err)
	}
	t.Logf("  - types_gen.go")

	if err := bg.GenerateConstantsGoMap(filepath.Join(bindgenDir, "constants_gen.go")); err != nil {
		t.Fatalf("Failed to generate constants map: %v", err)
	}
	t.Logf("  - constants_gen.go")

	if err := bg.GenerateValidationReport(filepath.Join(bindgenDir, "validation_report.txt")); err != nil {
		t.Fatalf("Failed to generate validation report: %v", err)
	}
	t.Logf("  - validation_report.txt")

	if err := bg.GenerateFixSuggestions(filepath.Join(bindgenDir, "fix_suggestions.txt")); err != nil {
		t.Fatalf("Failed to generate fix suggestions: %v", err)
	}
	t.Logf("  - fix_suggestions.txt")

	t.Logf("\nSummary:")
	t.Logf("  WDK Functions: %d", report.Statistics["total_wdk_functions"])
	t.Logf("  WDK Types: %d", report.Statistics["total_wdk_types"])
	t.Logf("  WDK Constants: %d", report.Statistics["total_wdk_constants"])
	t.Logf("  Rust Externs: %d", report.Statistics["total_externs"])
	t.Logf("  Type Mismatches: %d", report.Statistics["mismatches"])
	t.Logf("  Not Exported: %d", report.Statistics["not_exported"])
}

func TestPrintWdkFunctionSignatures(t *testing.T) {
	bindgenDir := getBindgenDir(t)

	config := BindgenConfig{
		ProjectRoot:    getProjectRoot(t),
		WdkBindingsDir: bindgenDir,
		OutputDir:      bindgenDir,
	}

	bg := NewBindgen(config)

	err := bg.LoadWdkBindings(
		filepath.Join(bindgenDir, "ntddk.rs"),
		filepath.Join(bindgenDir, "types.rs"),
		filepath.Join(bindgenDir, "constants.rs"),
	)
	if err != nil {
		t.Fatalf("Failed to load WDK bindings: %v", err)
	}

	wdk := bg.GetWdkBindings()

	t.Logf("\n=== Sample WDK Function Signatures ===\n")

	samples := []string{
		"PsLookupProcessByProcessId",
		"PsGetCurrentProcessId",
		"KeStackAttachProcess",
		"MmGetVirtualForPhysical",
		"ObfDereferenceObject",
	}

	for _, name := range samples {
		if sig, exists := wdk.Functions[name]; exists {
			t.Logf("%s:", name)
			t.Logf("  Line %d: pub fn %s(%s)", sig.Line, sig.Name, sig.Params)
			if sig.ReturnType != "void" {
				t.Logf("  -> %s", sig.ReturnType)
			}
			t.Logf("")
		} else {
			t.Logf("%s: NOT FOUND", name)
		}
	}
}

func TestFindMissingTypes(t *testing.T) {
	projectRoot := getProjectRoot(t)
	bindgenDir := getBindgenDir(t)

	config := BindgenConfig{
		ProjectRoot:    projectRoot,
		WdkBindingsDir: bindgenDir,
		RustSourceDir:  filepath.Join(projectRoot, "rust-driver", "kd", "src"),
		OutputDir:      bindgenDir,
	}

	bg := NewBindgen(config)

	err := bg.LoadWdkBindings(
		filepath.Join(bindgenDir, "ntddk.rs"),
		filepath.Join(bindgenDir, "types.rs"),
		filepath.Join(bindgenDir, "constants.rs"),
	)
	if err != nil {
		t.Fatalf("Failed to load WDK bindings: %v", err)
	}

	wdk := bg.GetWdkBindings()

	typesUsed := make(map[string][]string)

	for name, sig := range wdk.Functions {
		extractTypes(sig.Params, typesUsed, name)
		if sig.ReturnType != "void" {
			typesUsed[sig.ReturnType] = append(typesUsed[sig.ReturnType], name)
		}
	}

	missingTypes := make([]string, 0)
	for typeName := range typesUsed {
		if _, exists := wdk.Types[typeName]; !exists {
			missingTypes = append(missingTypes, typeName)
		}
	}

	t.Logf("\n=== Types used in WDK functions but not defined ===\n")
	if len(missingTypes) > 0 {
		for _, typeName := range missingTypes {
			t.Logf("  %s (used by: %v)", typeName, typesUsed[typeName][:min(3, len(typesUsed[typeName]))])
		}
	} else {
		t.Logf("  All types are defined")
	}
}

func extractTypes(params string, types map[string][]string, funcName string) {
	words := []string{}
	current := ""
	inGeneric := false

	for _, c := range params {
		switch c {
		case '<':
			inGeneric = true
			current += string(c)
		case '>':
			inGeneric = false
			current += string(c)
		case ' ', ',', '(', ')', '*', '&':
			if !inGeneric && current != "" {
				words = append(words, current)
				current = ""
			} else {
				current += string(c)
			}
		default:
			current += string(c)
		}
	}
	if current != "" {
		words = append(words, current)
	}

	for _, word := range words {
		if len(word) >= 2 && word[0] == 'P' && word[1] >= 'A' && word[1] <= 'Z' {
			types[word] = append(types[word], funcName)
		}
		if len(word) >= 2 && word[0] >= 'A' && word[0] <= 'Z' && word[1] >= 'A' && word[1] <= 'Z' {
			types[word] = append(types[word], funcName)
		}
	}
}

func TestMain(m *testing.M) {
	fmt.Println("=== Bindgen Test Suite ===")
	fmt.Println("Testing WDK binding validation and Rust code analysis")
	fmt.Println()
	os.Exit(m.Run())
}

func TestGenerateNtapiMod(t *testing.T) {
	projectRoot := getProjectRoot(t)
	bindgenDir := getBindgenDir(t)

	config := BindgenConfig{
		ProjectRoot:    projectRoot,
		WdkBindingsDir: bindgenDir,
		RustSourceDir:  filepath.Join(projectRoot, "rust-driver", "kd", "src"),
		OutputDir:      bindgenDir,
		Verbose:        true,
	}

	bg := NewBindgen(config)

	err := bg.LoadWdkBindings(
		filepath.Join(bindgenDir, "ntddk.rs"),
		filepath.Join(bindgenDir, "types.rs"),
		filepath.Join(bindgenDir, "constants.rs"),
	)
	if err != nil {
		t.Fatalf("Failed to load WDK bindings: %v", err)
	}

	notExported := []NotExportedFunc{
		{
			Name:       "KeGenericCallDpc",
			Params:     "DpcRoutine: unsafe extern \"C\" fn(*mut core::ffi::c_void, *mut core::ffi::c_void, *mut core::ffi::c_void), Context: *mut core::ffi::c_void",
			ReturnType: "",
		},
		{
			Name:       "KeInsertQueueDpc",
			Params:     "Dpc: PRKDPC, SystemArgument1: PVOID, SystemArgument2: PVOID",
			ReturnType: "BOOLEAN",
		},
		{
			Name:       "KeSignalCallDpcDone",
			Params:     "",
			ReturnType: "",
		},
		{
			Name:       "KeSignalCallDpcSynchronize",
			Params:     "",
			ReturnType: "",
		},
		{
			Name:       "ObDereferenceObject",
			Params:     "Object: PVOID",
			ReturnType: "",
		},
		{
			Name:       "ObOpenObjectByPointer",
			Params:     "Object: PVOID, DesiredAccess: u32, PassedAccessState: PVOID, AccessMode: i32, ObjectType: PVOID, HandleAttributes: u32, ProcessHandle: HANDLE",
			ReturnType: "NTSTATUS",
		},
		{
			Name:       "PsGetCurrentProcess",
			Params:     "",
			ReturnType: "PEPROCESS",
		},
		{
			Name:       "PsGetCurrentThread",
			Params:     "",
			ReturnType: "PETHREAD",
		},
		{
			Name:       "PsGetNextProcess",
			Params:     "Process: PEPROCESS",
			ReturnType: "PEPROCESS",
		},
		{
			Name:       "PsGetNextProcessThread",
			Params:     "Thread: PETHREAD, Process: PEPROCESS",
			ReturnType: "PETHREAD",
		},
		{
			Name:       "PsGetProcessImageFileName",
			Params:     "Process: PEPROCESS",
			ReturnType: "*mut u8",
		},
		{
			Name:       "PsGetProcessPeb",
			Params:     "Process: PEPROCESS",
			ReturnType: "PVOID",
		},
		{
			Name:       "PsGetProcessSectionBaseAddress",
			Params:     "Process: PEPROCESS",
			ReturnType: "PVOID",
		},
		{
			Name:       "PsGetProcessWow64Process",
			Params:     "Process: PEPROCESS",
			ReturnType: "PVOID",
		},
		{
			Name:       "PsGetContextThread",
			Params:     "Thread: PETHREAD, Context: *mut core::ffi::c_void",
			ReturnType: "NTSTATUS",
		},
		{
			Name:       "PsSetContextThread",
			Params:     "Thread: PETHREAD, Context: *mut core::ffi::c_void",
			ReturnType: "NTSTATUS",
		},
		{
			Name:       "PsSuspendProcess",
			Params:     "Process: PEPROCESS",
			ReturnType: "NTSTATUS",
		},
		{
			Name:       "PsResumeProcess",
			Params:     "Process: PEPROCESS",
			ReturnType: "NTSTATUS",
		},
		{
			Name:       "PsSuspendThread",
			Params:     "Thread: PETHREAD, PreviousSuspendCount: *mut ULONG",
			ReturnType: "NTSTATUS",
		},
		{
			Name:       "PsResumeThread",
			Params:     "Thread: PETHREAD, PreviousSuspendCount: *mut ULONG",
			ReturnType: "NTSTATUS",
		},
		{
			Name:       "RtlCopyUnicodeString",
			Params:     "DestinationString: PVOID, SourceString: PVOID",
			ReturnType: "",
		},
		{
			Name:       "RtlPcToFileHeader",
			Params:     "PcValue: PVOID, BaseOfImage: *mut PVOID",
			ReturnType: "PVOID",
		},
	}

	ntapiDir := filepath.Join(projectRoot, "rust-driver", "kd", "src", "ntapi")

	err = bg.GenerateNtapiMod(ntapiDir, notExported)
	if err != nil {
		t.Fatalf("Failed to generate ntapi files: %v", err)
	}

	t.Logf("Generated ntapi files in: %s", ntapiDir)
	t.Logf("Exported functions: %d", len(bg.GetWdkBindings().Functions)-len(notExported))
	t.Logf("Not exported functions: %d", len(notExported))

	sig, ok := bg.GetWdkBindings().Functions["NtDeviceIoControlFile"]
	if !ok {
		t.Error("NtDeviceIoControlFile not found")
	} else {
		t.Logf("NtDeviceIoControlFile params: %s", sig.Params)
		t.Logf("NtDeviceIoControlFile return: %s", sig.ReturnType)
		if sig.ReturnType != "NTSTATUS" {
			t.Errorf("Expected return type NTSTATUS, got %s", sig.ReturnType)
		}
	}
}

func TestScanProjectUsage(t *testing.T) {
	projectRoot := getProjectRoot(t)
	bindgenDir := getBindgenDir(t)

	config := BindgenConfig{
		ProjectRoot:    projectRoot,
		WdkBindingsDir: bindgenDir,
		RustSourceDir:  filepath.Join(projectRoot, "rust-driver", "kd", "src"),
		OutputDir:      bindgenDir,
		Verbose:        true,
	}

	bg := NewBindgen(config)

	err := bg.LoadWdkBindings(
		filepath.Join(bindgenDir, "ntddk.rs"),
		filepath.Join(bindgenDir, "types.rs"),
		filepath.Join(bindgenDir, "constants.rs"),
	)
	if err != nil {
		t.Fatalf("Failed to load WDK bindings: %v", err)
	}

	excludeDirs := []string{"cmd/bindgen", "target", "todo"}
	err = bg.ScanProjectUsage(projectRoot, excludeDirs)
	if err != nil {
		t.Fatalf("Failed to scan project: %v", err)
	}

	report := bg.GetFixReport()
	t.Logf("Files scanned: %d", report.FilesScanned)
	t.Logf("Total usages: %d", report.TotalUsages)
	t.Logf("Needs fix: %d", report.NeedsFix)

	reportPath := filepath.Join(bindgenDir, "wdk_usage_report.md")
	err = bg.GenerateUsageReport(reportPath)
	if err != nil {
		t.Fatalf("Failed to generate usage report: %v", err)
	}

	t.Logf("Generated usage report: %s", reportPath)

	if report.NeedsFix > 0 {
		t.Logf("\n=== Files needing fixes ===")
		for file, usages := range report.FilesWithFixes {
			for _, u := range usages {
				if u.NeedsFix {
					relPath, _ := filepath.Rel(projectRoot, file)
					t.Logf("  %s:%d - %s: %s", relPath, u.Line, u.Kind, u.Name)
				}
			}
		}
	}
}

func TestApplyFixesDryRun(t *testing.T) {
	projectRoot := getProjectRoot(t)
	bindgenDir := getBindgenDir(t)

	config := BindgenConfig{
		ProjectRoot:    projectRoot,
		WdkBindingsDir: bindgenDir,
		RustSourceDir:  filepath.Join(projectRoot, "rust-driver", "kd", "src"),
		OutputDir:      bindgenDir,
		Verbose:        true,
	}

	bg := NewBindgen(config)

	err := bg.LoadWdkBindings(
		filepath.Join(bindgenDir, "ntddk.rs"),
		filepath.Join(bindgenDir, "types.rs"),
		filepath.Join(bindgenDir, "constants.rs"),
	)
	if err != nil {
		t.Fatalf("Failed to load WDK bindings: %v", err)
	}

	excludeDirs := []string{"cmd/bindgen", "target", "todo"}
	err = bg.ScanProjectUsage(projectRoot, excludeDirs)
	if err != nil {
		t.Fatalf("Failed to scan project: %v", err)
	}

	err = bg.ApplyFixes(false)
	if err != nil {
		t.Fatalf("Failed to apply fixes: %v", err)
	}

	t.Log("Fixes applied successfully")
}

func TestApplyFixesMultiLineUse(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name: "single_line_use_wdk_sys",
			input: `use wdk_sys::*;

fn test() {
    let a = 1;
}`,
			expected: `use crate::ntapi::*;

fn test() {
    let a = 1;
}`,
		},
		{
			name: "multi_line_use_wdk_sys",
			input: `#[cfg(feature = "standalone-driver")]
use wdk_sys::ntddk::{
    IoDeleteDevice,
    IofCompleteRequest,
};

fn test() {
    unsafe {
        IoDeleteDevice(device);
    }
}`,
			expected: `#[cfg(feature = "standalone-driver")]
use crate::ntapi::*;

fn test() {
    unsafe {
        IoDeleteDevice(device);
    }
}`,
		},
		{
			name: "extern_block_with_wdk_functions",
			input: `use wdk_sys::*;

extern "C" {
    fn PsGetCurrentProcess() -> PEPROCESS;
    fn PsGetNextProcess(Process: PEPROCESS) -> PEPROCESS;
}

fn test() {
    unsafe {
        let proc = PsGetCurrentProcess();
    }
}`,
			expected: `use crate::ntapi::*;


fn test() {
    unsafe {
        let proc = PsGetCurrentProcess();
    }
}`,
		},
		{
			name: "use_after_multiline_use_crate",
			input: `use core::sync::atomic::{AtomicBool, Ordering};

use crate::hyperkd::hyperhv::state::{
    EPT_PML2_ENTRY,
    VMM_EPT_PML4E_COUNT,
};

fn test() {
    let a = 1;
}`,
			expected: `use core::sync::atomic::{AtomicBool, Ordering};

use crate::hyperkd::hyperhv::state::{
    EPT_PML2_ENTRY,
    VMM_EPT_PML4E_COUNT,
};

fn test() {
    let a = 1;
}`,
		},
		{
			name: "use_wdk_sys_after_multiline_use_crate",
			input: `use core::sync::atomic::{AtomicBool, Ordering};

use crate::hyperkd::hyperhv::state::{
    EPT_PML2_ENTRY,
};

use wdk_sys::*;

fn test() {
    let a = 1;
}`,
			expected: `use core::sync::atomic::{AtomicBool, Ordering};

use crate::hyperkd::hyperhv::state::{
    EPT_PML2_ENTRY,
};

use crate::ntapi::*;

fn test() {
    let a = 1;
}`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			bg := NewBindgen(BindgenConfig{})

			bg.fixReport = &FixReport{
				FilesWithFixes: make(map[string][]WdkUsage),
			}

			lines := strings.Split(tc.input, "\n")
			for i, line := range lines {
				if strings.Contains(line, "use wdk_sys::") {
					bg.fixReport.Usages = append(bg.fixReport.Usages, WdkUsage{
						File:         "test.rs",
						Line:         i + 1,
						Kind:         "import",
						Name:         "wdk_sys",
						SourceType:   "use_wdk_sys",
						NeedsFix:     true,
						SuggestedFix: "use crate::ntapi::*;",
					})
					bg.fixReport.FilesWithFixes["test.rs"] = append(bg.fixReport.FilesWithFixes["test.rs"], WdkUsage{
						File:         "test.rs",
						Line:         i + 1,
						Kind:         "import",
						Name:         "wdk_sys",
						SourceType:   "use_wdk_sys",
						NeedsFix:     true,
						SuggestedFix: "use crate::ntapi::*;",
					})
				}
				if strings.Contains(line, "fn PsGetCurrentProcess") || strings.Contains(line, "fn PsGetNextProcess") {
					name := "PsGetCurrentProcess"
					if strings.Contains(line, "PsGetNextProcess") {
						name = "PsGetNextProcess"
					}
					bg.fixReport.Usages = append(bg.fixReport.Usages, WdkUsage{
						File:         "test.rs",
						Line:         i + 1,
						Kind:         "function",
						Name:         name,
						SourceType:   "extern_block",
						NeedsFix:     true,
						SuggestedFix: fmt.Sprintf("use crate::ntapi::%s;", name),
					})
					bg.fixReport.FilesWithFixes["test.rs"] = append(bg.fixReport.FilesWithFixes["test.rs"], WdkUsage{
						File:         "test.rs",
						Line:         i + 1,
						Kind:         "function",
						Name:         name,
						SourceType:   "extern_block",
						NeedsFix:     true,
						SuggestedFix: fmt.Sprintf("use crate::ntapi::%s;", name),
					})
				}
			}

			bg.fixReport.TotalUsages = len(bg.fixReport.Usages)
			bg.fixReport.FilesScanned = 1

			tmpFile := filepath.Join(t.TempDir(), "test.rs")
			err := os.WriteFile(tmpFile, []byte(tc.input), 0o644)
			if err != nil {
				t.Fatalf("Failed to write temp file: %v", err)
			}

			bg.fixReport.FilesWithFixes[tmpFile] = bg.fixReport.FilesWithFixes["test.rs"]
			delete(bg.fixReport.FilesWithFixes, "test.rs")

			err = bg.ApplyFixes(false)
			if err != nil {
				t.Fatalf("Failed to apply fixes: %v", err)
			}

			result, err := os.ReadFile(tmpFile)
			if err != nil {
				t.Fatalf("Failed to read result: %v", err)
			}

			if string(result) != tc.expected {
				t.Errorf("Expected:\n%s\n\nGot:\n%s", tc.expected, string(result))
			}
		})
	}
}

func TestGenerateHookDatabase(t *testing.T) {
	projectRoot := getProjectRoot(t)
	bindgenDir := getBindgenDir(t)

	config := BindgenConfig{
		ProjectRoot:    projectRoot,
		WdkBindingsDir: bindgenDir,
		RustSourceDir:  filepath.Join(projectRoot, "rust-driver", "kd", "src"),
		OutputDir:      bindgenDir,
		Verbose:        true,
	}

	bg := NewBindgen(config)

	err := bg.LoadWdkBindings(
		filepath.Join(bindgenDir, "ntddk.rs"),
		filepath.Join(bindgenDir, "types.rs"),
		filepath.Join(bindgenDir, "constants.rs"),
	)
	if err != nil {
		t.Fatalf("Failed to load WDK bindings: %v", err)
	}

	notExported := []NotExportedFunc{
		{Name: "KeGenericCallDpc", Params: "DpcRoutine: unsafe extern \"C\" fn(*mut core::ffi::c_void, *mut core::ffi::c_void, *mut core::ffi::c_void), Context: *mut core::ffi::c_void", ReturnType: ""},
		{Name: "KeInitializeProcess", Params: "Process: PKPROCESS, Affinity: KAFFINITY, BaseThread: ULONG, Quantum: UCHAR, ThreadQuantumRatio: UCHAR", ReturnType: ""},
		{Name: "KeReadyThread", Params: "Thread: PKTHREAD", ReturnType: ""},
		{Name: "KiInitializeThread", Params: "Thread: PKTHREAD, Process: PKPROCESS, StartContext: PVOID, SystemRoutine: PKSYSTEM_ROUTINE, UserRoutine: PKUSER_ROUTINE, Teb: PVOID, Environment: PVOID", ReturnType: ""},
		{Name: "KeSetBasePriorityThread", Params: "Thread: PKTHREAD, Increment: KPRIORITY", ReturnType: "KPRIORITY"},
		{Name: "KeSetPriorityThread", Params: "Thread: PKTHREAD, Priority: KPRIORITY", ReturnType: "KPRIORITY"},
		{Name: "KeSetDisableBoostThread", Params: "Thread: PKTHREAD, Disable: BOOLEAN", ReturnType: "BOOLEAN"},
		{Name: "KeSetIdealProcessorThread", Params: "Thread: PKTHREAD, IdealProcessor: CCHAR", ReturnType: "CCHAR"},
		{Name: "KeSetAffinityThread", Params: "Thread: PKTHREAD, Affinity: KAFFINITY", ReturnType: "KAFFINITY"},
		{Name: "KeDelayExecutionThread", Params: "WaitMode: KPROCESSOR_MODE, Alertable: BOOLEAN, Interval: PLARGE_INTEGER", ReturnType: "NTSTATUS"},
		{Name: "KeTestAlertThread", Params: "WaitMode: KPROCESSOR_MODE", ReturnType: "BOOLEAN"},
		{Name: "KeForceResumeThread", Params: "Thread: PKTHREAD", ReturnType: "ULONG"},
		{Name: "KeSuspendThread", Params: "Thread: PKTHREAD", ReturnType: "ULONG"},
		{Name: "KeResumeThread", Params: "Thread: PKTHREAD", ReturnType: "ULONG"},
		{Name: "KeAlertResumeThread", Params: "Thread: PKTHREAD, PreviousCount: PULONG", ReturnType: "ULONG"},
		{Name: "KeAlertThread", Params: "Thread: PKTHREAD, AlertMode: KPROCESSOR_MODE", ReturnType: "BOOLEAN"},
		{Name: "KeWaitForSingleObject", Params: "Object: PVOID, WaitReason: KWAIT_REASON, WaitMode: KPROCESSOR_MODE, Alertable: BOOLEAN, Timeout: PLARGE_INTEGER", ReturnType: "NTSTATUS"},
		{Name: "KeWaitForMultipleObjects", Params: "Count: ULONG, Object: PVOID[], WaitType: WAIT_TYPE, WaitReason: KWAIT_REASON, WaitMode: KPROCESSOR_MODE, Alertable: BOOLEAN, Timeout: PLARGE_INTEGER, WaitBlockArray: PKWAIT_BLOCK", ReturnType: "NTSTATUS"},
		{Name: "KeWaitForMutexObject", Params: "Mutex: PVOID, WaitReason: KWAIT_REASON, WaitMode: KPROCESSOR_MODE, Alertable: BOOLEAN, Timeout: PLARGE_INTEGER", ReturnType: "NTSTATUS"},
		{Name: "KeReleaseMutant", Params: "Mutant: PKMUTANT, Increment: KPRIORITY, Abandoned: BOOLEAN, Wait: BOOLEAN", ReturnType: "LONG"},
		{Name: "KeReleaseSemaphore", Params: "Semaphore: PKSEMAPHORE, Increment: KPRIORITY, Adjustment: LONG, Wait: BOOLEAN", ReturnType: "LONG"},
		{Name: "KeSetEvent", Params: "Event: PKEVENT, Increment: KPRIORITY, Wait: BOOLEAN", ReturnType: "LONG"},
	}

	rustHookDir := filepath.Join(projectRoot, "rust-driver", "kd", "src", "hyperkd", "hyperhv", "hooks", "hook_db")
	goHookFile := filepath.Join(projectRoot, "debugger", "hook_db_gen.go")

	err = bg.GenerateHookDatabase(rustHookDir, goHookFile, notExported)
	if err != nil {
		t.Fatalf("Failed to generate hook database: %v", err)
	}

	t.Logf("Generated Rust hook database in: %s", rustHookDir)
	t.Logf("Generated Go hook database: %s", goHookFile)
	t.Logf("Total hooks: %d", len(bg.GetWdkBindings().Functions)-len(notExported))

	ls, _ := os.ReadDir(rustHookDir)
	for _, f := range ls {
		t.Logf("  - %s", f.Name())
	}
}
