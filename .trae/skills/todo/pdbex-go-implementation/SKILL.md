---
name: "pdbex-go-implementation"
description: "Guide for implementing Go version of pdbex PDB parser. Invoke when developing PDB parsing, symbol extraction, or type reconstruction features."
---

# PDB Parser Go Implementation Guide

Complete guide for implementing a Go version of the pdbex PDB parser tool.

## Overview

This skill documents the process of porting the C++ pdbex tool to Go, enabling:
1. PDB file parsing and symbol extraction
2. Structure/union/enum reconstruction to compilable C headers
3. Function name lookup by offset from PE files
4. Source code location mapping

## Source Reference

**Original C++ Implementation:**
`d:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\HyperDbg\hyperdbg\dependencies\pdbex\`

**Key Source Files:**
- `Source/PDB.h` - Main PDB class interface
- `Source/PDB.cpp` - PDB parsing implementation using DIA SDK
- `Source/PDBExtractor.h/cpp` - Command-line interface and extraction logic
- `Source/PDBHeaderReconstructor.h/cpp` - Header reconstruction
- `Source/PDBSymbolVisitor.h` - Symbol visitor pattern
- `Source/UdtFieldDefinition.h` - UDT field definitions

## Target Package Location

`d:\笔记本\p\ux\examples\hypedbg\walker\pdbex\`

## Implementation Plan

### Phase 1: Core Data Structures

Create Go equivalents of the C++ symbol structures:

```go
package pdbex

type SymTagEnum int

const (
    SymTagNull SymTagEnum = iota
    SymTagExe
    SymTagCompiland
    SymTagCompilandDetails
    SymTagCompilandEnv
    SymTagFunction
    SymTagBlock
    SymTagData
    SymTagAnnotation
    SymTagLabel
    SymTagPublicSymbol
    SymTagUDT
    SymTagEnum
    SymTagFunctionType
    SymTagPointerType
    SymTagArrayType
    SymTagBaseType
    SymTagTypedef
    SymTagBaseClass
    SymTagFriend
    SymTagFunctionArgType
    SymTagFuncDebugStart
    SymTagFuncDebugEnd
    SymTagUsingNamespace
    SymTagVTableShape
    SymTagVTable
    SymTagCustom
    SymTagThunk
    SymTagCustomType
    SymTagManagedType
    SymTagDimension
)

type UdtKind int

const (
    UdtStruct UdtKind = iota
    UdtClass
    UdtUnion
)

type BasicType int

const (
    btNoType BasicType = iota
    btVoid
    btChar
    btWChar
    btInt
    btUInt
    btFloat
    btBCD
    btBool
    btLong
    btULong
    btCurrency
    btDate
    btVariant
    btComplex
    btBitfield
    btBSTR
    btHresult
    btChar16
    btChar32
)

type Symbol struct {
    Tag        SymTagEnum
    DataKind   DataKind
    BaseType   BasicType
    TypeId     uint32
    Size       uint32
    IsConst    bool
    IsVolatile bool
    Name       string
    
    // Union-like fields based on Tag
    Enum     *SymbolEnum
    Typedef  *SymbolTypedef
    Pointer  *SymbolPointer
    Array    *SymbolArray
    Function *SymbolFunction
    Udt      *SymbolUdt
}

type SymbolEnumField struct {
    Name   string
    Value  interface{}
    Parent *Symbol
}

type SymbolEnum struct {
    FieldCount uint32
    Fields     []*SymbolEnumField
}

type SymbolUdtField struct {
    Name       string
    Type       *Symbol
    Offset     uint32
    Bits       uint32
    BitPosition uint32
    Parent     *Symbol
}

type SymbolUdt struct {
    Kind       UdtKind
    FieldCount uint32
    Fields     []*SymbolUdtField
}

type SymbolPointer struct {
    Type        *Symbol
    IsReference bool
}

type SymbolArray struct {
    ElementType *Symbol
    ElementCount uint32
}

type SymbolFunction struct {
    ReturnType       *Symbol
    CallingConvention CallingConvention
    ArgumentCount    uint32
    Arguments        []*Symbol
}

type SymbolTypedef struct {
    Type *Symbol
}
```

### Phase 2: PDB Parser Interface

```go
type PDB struct {
    path        string
    machineType uint16
    language    uint32
    symbols     map[uint32]*Symbol
    symbolNames map[string]*Symbol
    functions   map[string]uint64
}

func NewPDB() *PDB

func (p *PDB) Open(path string) error

func (p *PDB) Close()

func (p *PDB) IsOpened() bool

func (p *PDB) GetPath() string

func (p *PDB) GetMachineType() uint16

func (p *PDB) GetSymbolByName(name string) (*Symbol, bool)

func (p *PDB) GetSymbolByTypeId(typeId uint32) (*Symbol, bool)

func (p *PDB) GetAllSymbols() map[uint32]*Symbol

func (p *PDB) GetAllFunctions() map[string]uint64

func (p *PDB) GetFunctionByOffset(offset uint64) (string, bool)
```

### Phase 3: Header Reconstruction

```go
type HeaderReconstructor struct {
    settings *Settings
    output   strings.Builder
}

type Settings struct {
    MemberStructExpansion ExpansionType
    CreatePaddingMembers  bool
    ShowOffsets           bool
    MicrosoftTypedefs     bool
    AllowBitFieldsInUnion bool
    AllowAnonymousDataTypes bool
    UseStdInt             bool
    SymbolPrefix          string
    SymbolSuffix          string
    AnonymousUnionPrefix  string
    AnonymousStructPrefix string
}

type ExpansionType int

const (
    ExpansionNone ExpansionType = iota
    ExpansionInlineUnnamed
    ExpansionInlineAll
)

func (r *HeaderReconstructor) Reconstruct(symbol *Symbol) string

func (r *HeaderReconstructor) ReconstructAll(pdb *PDB) string
```

### Phase 4: PE File Integration

```go
type PEIntegration struct {
    pdb    *PDB
    pePath string
}

func NewPEIntegration(pdbPath, pePath string) (*PEIntegration, error)

func (p *PEIntegration) GetFunctionNameByOffset(offset uint64) (string, bool)

func (p *PEIntegration) GetSourceLocation(offset uint64) (string, uint32, bool)
```

## Implementation Approach

### Option A: Pure Go with debug/dwarf (Limited PDB Support)

Go's standard library has limited PDB support. Use `golang.org/x/debug/dwarf` for DWARF, but PDB requires Windows-specific APIs.

### Option B: CGO with DIA SDK (Recommended)

Use CGO to interface with Windows DIA SDK (msdia140.dll):

```go
/*
#cgo LDFLAGS: -lole32 -loleaut32
#include <windows.h>
#include <dia2.h>
#include <atlcomcli.h>
*/
import "C"
```

### Option C: Use Existing Go PDB Libraries

Consider using or extending:
- `github.com/microsoft/go-winio/pkg/guid`
- `github.com/0xrawsec/golang-win32`
- Custom implementation based on PDB format specification

## Test Requirements

### Test PDB File
`D:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\build\Release\hyperkd.pdb`

### Test PE File
`D:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\build\Release\hyperkd.sys`

### Test Cases

```go
func TestPDBOpen(t *testing.T) {
    pdb := NewPDB()
    err := pdb.Open("hyperkd.pdb")
    require.NoError(t, err)
    defer pdb.Close()
    assert.True(t, pdb.IsOpened())
}

func TestGetSymbolByName(t *testing.T) {
    pdb := NewPDB()
    pdb.Open("hyperkd.pdb")
    defer pdb.Close()
    
    symbol, ok := pdb.GetSymbolByName("_DEBUGGER_EVENT_AND_ACTION_RESULT")
    require.True(t, ok)
    assert.Equal(t, SymTagUDT, symbol.Tag)
}

func TestGetFunctionByOffset(t *testing.T) {
    pdb := NewPDB()
    pdb.Open("hyperkd.pdb")
    defer pdb.Close()
    
    pe := NewPEIntegration(pdb, "hyperkd.sys")
    name, ok := pe.GetFunctionNameByOffset(0x1000)
    require.True(t, ok)
    assert.NotEmpty(t, name)
}

func TestHeaderReconstruction(t *testing.T) {
    pdb := NewPDB()
    pdb.Open("hyperkd.pdb")
    defer pdb.Close()
    
    reconstructor := NewHeaderReconstructor(DefaultSettings())
    header := reconstructor.ReconstructAll(pdb)
    
    // Verify header is compilable
    assert.Contains(t, header, "typedef struct")
}
```

## Source Code Location Mapping

For each function/method, add comments with C++ source path:

```go
// Symbol represents a debug symbol from PDB.
// Source: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/dependencies/pdbex/Source/PDB.h:247-276
type Symbol struct {
    // ...
}

// GetSymbolByName retrieves a symbol by its name.
// Source: doc/cpp/HyperDbgUnified/HyperDbg/hyperdbg/dependencies/pdbex/Source/PDB.cpp:453-458
func (p *PDB) GetSymbolByName(name string) (*Symbol, bool) {
    // ...
}
```

## Package Structure

```
walker/pdbex/
├── pdbex.go              # Main PDB parser
├── symbol.go             # Symbol data structures
├── reconstructor.go      # Header reconstruction
├── pe_integration.go     # PE file integration
├── dia/                  # DIA SDK bindings (CGO)
│   ├── dia.go
│   ├── dia_windows.go
│   └── types.go
├── pdbex_test.go         # Unit tests
└── testdata/             # Test PDB files
```

## Dependencies

- Windows OS (DIA SDK requirement)
- msdia140.dll (Visual Studio component)
- Go 1.21+
- `github.com/saferwall/pe` for PE parsing and export table inspection

## DIA SDK DLL Management

### x86 and x64 Versions

msdia140.dll comes in two architectures:
- **x64**: `msdia140.x64.dll` - For 64-bit PDB files
- **x86**: `msdia140.x86.dll` - For 32-bit PDB files

Both DLLs are embedded in the package using Go's `embed` package:

```
walker/pdbex/
├── msdia140.x64.dll    # 64-bit DIA SDK
├── msdia140.x86.dll    # 32-bit DIA SDK
├── embed.go            # Embed declarations
└── dia_windows.go      # DIA loading logic
```

### Embed Declaration (embed.go)

```go
package pdbex

import _ "embed"

//go:embed msdia140.x64.dll
var msdia140x64 []byte

//go:embed msdia140.x86.dll
var msdia140x86 []byte
```

### DLL Loading Strategy

1. Extract embedded DLL to temporary directory
2. Load using `LoadLibrary` syscall
3. Get `DllGetClassObject` export address
4. Create IDiaDataSource instance via ClassFactory

### DIA SDK Exports (via saferwall/pe)

Using `github.com/saferwall/pe` to inspect msdia140.dll exports:

```go
file, _ := pe.New("msdia140.dll", &pe.Options{})
file.Parse()

for _, exp := range file.Export.Functions {
    fmt.Printf("%s (Ordinal: %d, RVA: 0x%X)\n", 
        exp.Name, exp.Ordinal, exp.FunctionRVA)
}
```

Exported functions:
- `DllCanUnloadNow` (Ordinal: 1)
- `DllGetClassObject` (Ordinal: 2) - Main entry point for DIA
- `DllRegisterServer` (Ordinal: 3)
- `DllUnregisterServer` (Ordinal: 4)
- `VSDllRegisterServer` (Ordinal: 5)
- `VSDllUnregisterServer` (Ordinal: 6)

### CLSID for IDiaDataSource

```go
clsidDiaSource := syscall.GUID{
    Data1: 0xe6756135,
    Data2: 0x1e65,
    Data3: 0x4d72,
    Data4: [8]byte{0x8a, 0x7d, 0x40, 0x41, 0x07, 0x20, 0x8f, 0x2f},
}
```

## Progress Tracking

- [x] Phase 1: Core data structures (`symbol.go`)
- [x] Phase 2: PDB parser interface (`pdbex.go`, `dia_windows.go`)
- [x] Phase 3: Header reconstruction (`reconstructor.go`)
- [x] Phase 4: PE file integration (`pe_integration.go`)
- [x] Phase 5: Unit tests with hyperkd.pdb (`pdbex_test.go`)
- [x] Phase 6: Source location mapping (`source_mapper.go`)
- [x] Phase 7: Documentation

## Implementation Status

### Completed Files

| File | Description | Status |
|------|-------------|--------|
| `symbol.go` | Core data structures (SymTag, Symbol, UdtKind, BasicType, etc.) | ✅ Complete |
| `pdbex.go` | PDB parser interface with symbol/function management | ✅ Complete |
| `dia_windows.go` | Windows DIA SDK bindings using syscall | ✅ Complete |
| `reconstructor.go` | Header reconstruction with settings | ✅ Complete |
| `pe_integration.go` | PE file parsing and function offset lookup | ✅ Complete |
| `source_mapper.go` | Source location mapping and annotation | ✅ Complete |
| `pdbex_test.go` | Unit tests for all components | ✅ Complete |

### Key Implementation Details

1. **DIA SDK Integration**: Uses Windows syscall to interface with DIA SDK (msdia140.dll) for PDB parsing
2. **Symbol Types**: Supports UDT, Enum, Typedef, Pointer, Array, Function symbols
3. **Header Generation**: Generates compilable C headers with configurable settings
4. **PE Integration**: Parses PE files and maps function offsets to symbol names
5. **Source Mapping**: Tracks source file locations for functions and symbols

### Known Limitations

- DIA SDK requires Windows and Visual Studio components
- Source location extraction depends on PDB containing source information
- Some complex C++ constructs may not fully reconstruct

## References

- [PDB Format Specification](https://github.com/Microsoft/microsoft-pdb)
- [DIA SDK Documentation](https://docs.microsoft.com/en-us/visualstudio/debugger/debug-interface-access/debug-interface-access-sdk)
- [Original pdbex Repository](https://github.com/wbenny/pdbex)
