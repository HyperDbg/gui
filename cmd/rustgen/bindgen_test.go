package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func getProjectRootForTest(t *testing.T) string {
	t.Helper()
	cwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get cwd: %v", err)
	}
	return filepath.Join(cwd, "..", "..")
}

func getBindgenDir(t *testing.T) string {
	t.Helper()
	return filepath.Join(getProjectRootForTest(t), "cmd", "bindgen")
}

func TestLoadWdkBindings(t *testing.T) {
	bindgenDir := getBindgenDir(t)

	config := BindgenConfig{
		ProjectRoot:    getProjectRootForTest(t),
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
	projectRoot := getProjectRootForTest(t)
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
	projectRoot := getProjectRootForTest(t)
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

func TestMain(m *testing.M) {
	fmt.Println("=== Bindgen Test Suite ===")
	fmt.Println("Testing WDK binding validation and Rust code analysis")
	fmt.Println()
	os.Exit(m.Run())
}

func TestScanProjectUsage(t *testing.T) {
	projectRoot := getProjectRootForTest(t)
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
	projectRoot := getProjectRootForTest(t)
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

	err = bg.ApplyFixes()
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

			err = bg.ApplyFixes()
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
