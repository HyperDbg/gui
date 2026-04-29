---
name: "bindgen-binding"
description: "Generates Go bindings from C headers and creates DLL bindings with unit tests. Invoke when user wants to add a new C library binding (WinDivert, Everything-SDK, zydis, keystone etc.) or when TestGenerate needs to be run for bindgen."
---

# Bindgen Binding Workflow

This skill handles the complete workflow of generating Go bindings from C header files using the custom bindgen tool (modernc/cc based).

## When to Invoke

- User wants to add a **new C library binding** (any DLL/header-based native library)
- User asks to run `TestGenerate` or regenerate bindings
- Generated code has compilation errors (enum overflow, missing imports)
- Unit tests for generated bindings need fixing or running
- User mentions "绑定" (binding), "bindgen", "生成代码", "DLL 绑定"

## Prerequisites

### System Environment Variables (Required, set via registry)

The following **must** be permanently set in the system environment (registry):

| Variable | Example Value |
|----------|--------------|
| `CC` | `E:\Program Files\Microsoft Visual Studio\2022\BuildTools\VC\Tools\MSVC\14.44.35207\bin\Hostx64\x64\cl.exe` |
| `CMAKE_C_COMPILER` | same as CC |
| `INCLUDE` | MSVC include + Windows SDK include paths (semicolon-separated) |
| `LIB` | MSVC lib + Windows SDK lib paths (semicolon-separated) |
| `VCToolsInstallDir` | `E:\Program Files\Microsoft Visual Studio\2022\BuildTools\VC\Tools\MSVC\14.44.35207\` |
| `VSINSTALLDIR` | `E:\Program Files\Microsoft Visual Studio\2022\BuildTools\` |
| `WDKContentRoot` | `E:\Program Files\Windows Kits\10\` |
| `WindowsTargetPlatformVersion` | `10.0.28000.0` |

**No fallback values are used.** All paths come from these environment variables.

### Architecture Notes

- **Do NOT use `cc.NewConfig()`** — it internally calls the compiler with GCC-style flags (`-dM`, `-E`, `-v`) which **MSVC cl.exe does not support**
- Instead, manually construct `*cc.Config` using `cc.NewABI()` + `findClExe()` + `loadMSVCExtraTypes()`
- **Predefined source** (`<predefined>`) must contain only `#define` macro definitions
- **Type definitions** (`typedef`, `struct`) go in a separate `<msvc_types>` source
- **Builtin source** (`<builtin>`) requires `__SIZE_TYPE__`, `__WCHAR_TYPE__`, `__PTRDIFF_TYPE__` etc. defined in Predefined

## Complete Workflow

### Step 1: Add BindgenConfig Entry

Edit `dep/bindgen/generate_test.go` in the `TestGenerate` function's config slice:

```go
{
    HeadersDir:  "project/<project-dir>/clone/include",
    OutputDir:   "project/<project-dir>",
    PackageName: "<go-package-name>",
    HeaderOrder: []string{
        "header1.h",
        "header2.h",
    },
    SingleFile: true,
    BindDll:    true,
    DllName:    "<dll-filename>.dll",
    DllFuncFilter: func(name string) bool {
        return strings.HasPrefix(name, "<prefix>")
    },
},
```

**Key fields:**
- `HeadersDir`: Path to directory containing `.h` files (must exist). Use `clone` subdirectory for source-controlled headers
- `OutputDir`: Where generated `.go` files are written
- `PackageName`: Go package name for generated code
- `ModuleName`: Optional custom module path. Defaults to `github.com/ddkwork/<PackageName>` if empty
- `HeaderOrder`: Ordered list of headers (order matters for dependencies)
- `BindDll`: If true, generates `dll.go` with embedded DLL + LazyDLL loading
- `DllName`: DLL filename for embedding
- `DllFuncFilter`: Optional function to select which exported functions to bind
- `ExtraIncludeDirs`: Additional include directories for header resolution
- `RecurseHeaders`: If true, recursively scan subdirectories for headers
- `SingleFile`: If true, output all generated code into a single file
- `Predefined`: Extra `#define` macros and declarations appended to the predefined source (use backtick raw strings for multi-line)
- `SkipMSVCTypes`: If true, skip `<msvc_types>` source injection
- `ExtraConstants`: Additional Go constants to inject into generated code

### Step 2: Clone Directory Convention

Each project uses a `clone` subdirectory for source-controlled C headers. This allows upgrading the upstream library without modifying generated bindings:

```
project/<name>/clone/        ← git submodule or copy of upstream source
project/<name>/clone/include/ ← headers for binding generation
project/<name>/              ← generated Go bindings output here
```

**IMPORTANT:** Never modify files inside `clone/` directories. If headers cause parsing errors (e.g., undefined `stderr`, `fprintf`), fix them via the `Predefined` field in `BindgenConfig` instead.

Example — XED headers reference `stderr`/`fprintf`/`abort` via `api_check` macros:
```go
Predefined: `
#define XED_ENCODER
#define XED_ENCODE_ORDER_MAX_OPERANDS 5
#define XED_ENCODE_ORDER_MAX_ENTRIES 35
#define stderr ((void*)0)
int fprintf(void*, const char*, ...);
int fflush(void*);
void abort(void);
`,
```

### Step 3: Build Native Library (if needed)

Some projects (zydis, keystone) require building a DLL from source before binding generation. Each project provides a `build.cmd`/`build.bat` script and a `CMakeLists.txt` that uses the `ewdk.cmake` infrastructure.

**Build command** (run from project directory):

```bash
# zydis
cd dep/bindgen/project/zydis && build.cmd

# keystone
cd dep/bindgen/project/keystone && build.bat
```

**What `build.cmd` does internally:**
```bash
cmake -B build -G "Ninja" -DCMAKE_BUILD_TYPE=Release . && cmake --build build --config Release
```

**Prerequisites for building:**
- System environment variables must be set in the registry (see `系统环境变量.log`): `CC`, `CMAKE_C_COMPILER`, `INCLUDE`, `LIB`, `WDKContentRoot`, `VCToolsInstallDir`, `WindowsTargetPlatformVersion`
- Ninja must be available in `PATH`
- `ewdk.cmake` must be in the project directory — it provides `um_dll()` function for UCRT/WDK integration

**CMakeLists.txt structure:**
1. Include `ewdk.cmake` from project root: `include(${CMAKE_CURRENT_SOURCE_DIR}/ewdk.cmake)`
2. Use `um_dll(<name> SHARED ${SourceFiles})` — NOT `add_library()` — for proper UCRT/WDK DLL linking
3. Source files reference `clone/` subdirectory: `file(GLOB SourceFiles "clone/<lib>/src/*.c")`
4. `ewdk.cmake` auto-configures include paths, link paths, and compile definitions from registry env vars

**Skipped projects:** x64dbg (explicitly skipped per user request)

### Step 4: Run Generator

```bash
cd dep/bindgen && go test -v -run TestGenerate/<name> -count=1 -timeout 300s .
```

To run all bindings:
```bash
cd dep/bindgen && go test -v -run TestGenerate -count=1 -timeout 300s .
```

**Expected output:**
- `msenv: loaded N include paths`
- `Processing: header1.h` per header
- `✅ All binding generations completed!`
- If `BindDll: true`: DLL binding code generated

### Step 5: Verify Generated Code

After generation, check these common issues that the **bindgen auto-handles**:

1. **Enum negative values** → Auto-detects `-` prefix, generates `int32` instead of `uint32`
2. **syscall import** → Auto-detects callback types, adds `"syscall"` import

### Step 6: Create/Edit Unit Tests

Test file location: `<OutputDir>/dll_test.go`

**Standard test pattern** (use package-local types, NOT external packages):

```go
package <pkgname>

import (
    "testing"
    "unsafe"
)

func Test<Feature>(t *testing.T) {
    z := &<StructType>{}
}
```

**Critical rules:**
- Use local type names from generated package
- For `*uintptr` parameters: `(*uintptr)(unsafe.Pointer(&slice[0]))`
- For string results from C: `unsafe.String((*byte)(unsafe.Pointer(ptr)), len)`

### Step 7: Run Tests

```bash
cd dep/bindgen/project/<name> && go test -count=1 -v -timeout 120s .
```

### Step 8: Full Verification Pipeline (Mandatory After Any Generator Change)

After modifying `generate_test.go` (e.g., fixing macro filtering, removing regex, changing constant generation logic), **always** run this complete verification sequence:

```bash
# 1. Regenerate all bindings
cd dep/bindgen && go test -v -run TestGenerate -count=1 -timeout 300s .

# 2. Run ALL project tests (comprehensive validation)
go test -v ./debugger/ -count=1                                          # HyperDbg debugger integration
go test -v ./dep/bindgen/project/Everything-SDK/ -count=1                 # Everything-SDK stub
go test -v ./dep/bindgen/project/keystone/ -count=1                       # Keystone assembler
go test -v ./dep/bindgen/project/xed/ -count=1                            # XED encoder/decoder
go test -v ./dep/bindgen/project/zydis/ -count=1                          # Zydis disassembler
```

**All 6 commands must pass (exit code 0)** before committing. This is the authoritative validation that the bindgen generator produces correct output for all bound projects.

| Test | What it validates | Location |
|------|-------------------|----------|
| `TestGenerate` | Code generation itself (no panics, files written) | `dep/bindgen/generate_test.go` |
| `debugger_test.go` | HyperDbg SDK constants/structs usable in debugger | `debugger/debugger_test.go` |
| `Everything-SDK` | Everything SDK stub compiles and types work | `dep/bindgen/project/Everything-SDK/` |
| `keystone dll_test.go` | Keystone DLL binding: open/close, asm x86/arm64 | `dep/bindgen/project/keystone/dll_test.go` |
| `xed dll_test.go` | XED binding: decode NOP/MOV/ADD, encode round-trip, assemble | `dep/bindgen/project/xed/dll_test.go` |
| `zydis dll_test.go` | Zydis binding: disassemble instructions | `dep/bindgen/project/zydis/dll_test.go` |

**No regex in generator**: The bindgen uses only string operations (`strings`, `strconv`, manual char-by-char parsing). All `regexp` usage has been removed from `generate_test.go`. Helper functions: `isWordChar()`, `hasWord()`, `replaceWord()`, `extractIdentifiers()`, `containsDigit()`, `parseSizeofType()`, etc.

## Core Source Architecture

### msenv.go — Environment Configuration

```
findClExe()          → os.Getenv("CC")                          // cl.exe path
findEWDKIncludePaths() → os.Getenv("INCLUDE") + SDK/headers      // include dirs
loadMSVCExtraTypes()  → #define macros (__SIZE_TYPE__, _MSC_VER, etc.)
loadMSVCTypes()       → typedef declarations (UINT64, LIST_ENTRY, etc.)
```

### generate_test.go — Config Construction

```
newMSVCConfig(t):
  1. cc.NewABI(goos, goarch)          // ABI info (no compiler probe)
  2. findClExe()                      // from env CC
  3. loadMSVCExtraTypes()             // predefined macros
  4. findEWDKIncludePaths()           // from env INCLUDE
  → returns *cc.Config{ABI, CC, Predefined, IncludePaths, SysIncludePaths}

Translate() source order:
  1. <predefined>   = cfg.Predefined + bc.Predefined  (#define macros only)
  2. <builtin>      = cc.Builtin                       (requires __SIZE_TYPE__ etc.)
  3. <msvc_types>    = loadMSVCTypes()                 (typedef/struct definitions)
  4..N header files  = HeaderOrder contents
```

### BindgenConfig struct fields

```go
type BindgenConfig struct {
    HeadersDir       string
    OutputDir        string
    PackageName      string
    ModuleName       string
    HeaderOrder      []string
    BindDll          bool
    DllName          string
    DllFuncFilter    func(name string) bool
    ExtraIncludeDirs []string
    RecurseHeaders   bool
    SingleFile       bool
    Predefined       string
    SkipMSVCTypes    bool
    ExtraConstants   map[string]extraConst
}
```

### parse_test.go — Parse Testing

Uses same `newTestConfig(t)` pattern as generate_test.go.

## Current Bound Projects

| Project | HeadersDir | OutputDir | Package | DLL | SingleFile | Status |
|---------|-----------|-----------|---------|-----|------------|--------|
| hyperdbg | project/hyperdbg/clone/SDK | project/hyperdbg/sdk | sdk | No | Yes | ✅ Active |
| zydis | project/zydis/clone/zydis/include | project/zydis | zydis | zydis.dll | Yes | ✅ Active |
| ipmrec | project/ARImpRec/clone | project/ARImpRec | ipmrec | ARImpRec.dll | No | ✅ Active |
| keystone | project/keystone/clone/keystone/include/keystone | project/keystone | keystone | keystone.dll | No | ✅ Active |
| windivert | project/WinDivert/clone/include | project/WinDivert | windivert | WinDivert64.dll | No | ✅ Active |
| everything | project/Everything-SDK/clone/include | project/Everything-SDK | everything | Everything64.dll | No | ✅ Active |
| xed | project/xed/clone/include | project/xed | xed | xed.dll | Yes | ✅ Active |

### XED Special Configuration

XED requires extra `Predefined` macros and `ExtraIncludeDirs` because:
- `xed-build-defines.h` enables `XED_API_CHECK_ENABLED` which expands `api_check` macros referencing `stderr`/`fprintf`/`abort`
- Solution: define `stderr` as `((void*)0)` and declare `fprintf`/`fflush`/`abort` in `Predefined`
- Extra include: `project/xed/clone/xed/include/public/xed` for internal headers

## File Reference

- Environment config: [msenv.go](file:///d:/New/New%20folder/hypedbg/dep/bindgen/msenv.go) — `findClExe()`, `findEWDKIncludePaths()`, `loadMSVCExtraTypes()`, `loadMSVCTypes()`
- Generator entry: [generate_test.go](file:///d:/New/New%20folder/hypedbg/dep/bindgen/generate_test.go) — `TestGenerate()` at line ~148, `newMSVCConfig()` at line ~18, `processBindgenConfig()` at line ~288
- Parse tests: [parse_test.go](file:///d:/New/New%20folder/hypedbg/dep/bindgen/parse_test.go) — `newTestConfig()`, `testParse()`, `TestRunAllParse()`
