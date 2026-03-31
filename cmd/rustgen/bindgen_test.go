package main

import (
	"fmt"
	"os"
	"path/filepath"
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
