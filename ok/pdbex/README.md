# PDB Parser (pdbex)

Go implementation of PDB parser using Windows DIA SDK for symbol extraction and type reconstruction.

## Features

- **PDB File Parsing**: Open and parse PDB debug symbol files
- **Symbol Extraction**: Extract functions, UDTs, enums, typedefs, etc.
- **Source Line Lookup**: Find source file and line number by address (RVA/VA) - useful for BSOD analysis
- **Header Reconstruction**: Generate compilable C headers from PDB types
- **PE Integration**: Map PE file offsets to symbol names

## Usage

### Basic Usage

```go
package main

import (
    "fmt"
    pdbex "github.com/ddkwork/HyperDbg/walker/pdbex"
)

func main() {
    pdb := pdbex.NewPDB()
    defer pdb.Close()

    // Open PDB file
    if err := pdb.Open("driver.pdb"); err != nil {
        panic(err)
    }

    // Get all functions
    functions := pdb.GetAllFunctions()
    for name, fn := range functions {
        fmt.Printf("%s: 0x%X (size: %d)\n", name, fn.Address, fn.Size)
    }

    // Find source line by RVA (useful for BSOD analysis)
    fileName, lineNumber, ok := pdb.FindSourceLineByRVA(0x1000)
    if ok {
        fmt.Printf("Source: %s:%d\n", fileName, lineNumber)
    }

    // Get symbol by name
    sym, ok := pdb.GetSymbolByName("_DEBUGGER_EVENT")
    if ok {
        fmt.Printf("Symbol: %s, Size: %d\n", sym.Name, sym.Size)
    }
}
```

### PE Integration

```go
// Create PE integration for address mapping
integration, err := pdbex.NewPEIntegration("driver.pdb", "driver.sys")
if err != nil {
    panic(err)
}
defer integration.Close()

// Get function name by RVA
name, ok := integration.GetFunctionNameByRVA(0x1000)
if ok {
    fmt.Printf("Function at 0x1000: %s\n", name)
}
```

### Header Reconstruction

```go
settings := pdbex.DefaultSettings()
settings.PrintHeader = true
settings.ShowOffsets = true

reconstructor := pdbex.NewHeaderReconstructor(settings)
header := reconstructor.ReconstructAll(pdb)
fmt.Println(header)
```

## API Reference

### PDB Methods

| Method | Description |
|--------|-------------|
| `Open(path string) error` | Open a PDB file |
| `Close()` | Close and release resources |
| `IsOpened() bool` | Check if PDB is opened |
| `GetPath() string` | Get PDB file path |
| `GetMachineType() uint16` | Get target machine type |
| `GetArchitectureString() string` | Get architecture name (i386, AMD64, ARM64, etc.) |
| `GetAllFunctions() map[string]*FunctionInfo` | Get all functions |
| `GetFunctionByOffset(offset uint64) (string, bool)` | Find function by address |
| `GetSymbolByName(name string) (*Symbol, bool)` | Get symbol by name |
| `GetUDTs() []*Symbol` | Get all UDT (struct/union/class) symbols |
| `GetEnums() []*Symbol` | Get all enum symbols |
| `FindSourceLineByRVA(rva uint32) (string, uint32, bool)` | Find source file and line by RVA |
| `FindSourceLineByVA(va uint64) (string, uint32, bool)` | Find source file and line by VA |
| `DumpSymbol(name string) (string, error)` | Dump symbol details |

## Files

| File | Description |
|------|-------------|
| `pdbex.go` | Main PDB parser interface |
| `dia_windows.go` | Windows DIA SDK bindings |
| `dia_vtable_gen.go` | Generated vtable constants |
| `symbol.go` | Symbol data structures |
| `reconstructor.go` | Header reconstruction |
| `pe_integration.go` | PE file integration |
| `source_mapper.go` | Source location mapping |
| `embed.go` | Embedded DLL resources |
| `pdbex_test.go` | Unit tests |

## DIA SDK

This package uses Windows DIA SDK (msdia140.dll) for PDB parsing. The DLLs are embedded and extracted at runtime:

- `test_dia/dia/bin/amd64/msdia140.dll` - x64
- `test_dia/dia/bin/arm64/msdia140.dll` - ARM64
- `test_dia/dia/bin/msdia140.dll` - x86

### Updating DIA SDK

Run the vtable generator to sync DIA SDK files and regenerate constants:

```bash
go run ./walker/pdbex/test_dia/count_vtable.go
```

## Requirements

- Windows OS (DIA SDK requirement)
- Go 1.21+
- `github.com/saferwall/pe` for PE parsing

## BSOD Analysis Example

```go
// Analyze crash address from BSOD dump
pdb := pdbex.NewPDB()
defer pdb.Close()

pdb.Open("hyperkd.pdb")

// Crash address (example: from !analyze output)
crashRVA := uint32(0x33CB0)

// Get source location
fileName, lineNumber, ok := pdb.FindSourceLineByRVA(crashRVA)
if ok {
    fmt.Printf("Crash location: %s:%d\n", fileName, lineNumber)
}

// Get function name
funcName, ok := pdb.GetFunctionByOffset(uint64(crashRVA))
if ok {
    fmt.Printf("Function: %s\n", funcName)
}
```

## License

MIT
