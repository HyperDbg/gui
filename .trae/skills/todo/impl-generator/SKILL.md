---
name: "impl-generator"
description: "Generates interface method implementations for structs. Invoke when user wants to implement interface methods in a struct, automatically detecting missing methods and signature mismatches."
---

# Interface Implementation Generator

This skill generates interface method implementations for Go structs.

## When to Use

Invoke this skill when:
- User wants to implement interface methods in a struct
- User needs to check which interface methods are missing
- User wants to ensure method signatures match the interface
- User needs to organize methods according to interface order

## How It Works

1. **Parse Input File**: Reads the specified Go source file
2. **Find Struct**: Locates the target struct by name
3. **Identify Interface**: Finds the interface the struct should implement
4. **Compare Methods**: Checks which methods are implemented and which are missing
5. **Detect Mismatches**: Identifies methods with signature mismatches
6. **Generate Implementations**: Creates method skeletons for missing methods
7. **Sort Methods**: Orders methods according to interface definition
8. **Preserve Existing**: Keeps existing implementations intact

## Usage

### Command Format

```
impl-generator <input-file> <output-file> <struct-name> <interface-name>
```

### Parameters

| Parameter | Required | Description |
|-----------|----------|-------------|
| `input-file` | Yes | Path to the Go source file containing the struct |
| `output-file` | Yes | Path where the generated file should be saved |
| `struct-name` | Yes | Name of the struct to implement interface methods for |
| `interface-name` | Yes | Name of the interface to implement |

### Example

```bash
# Generate implementations for Provider struct implementing api interface
impl-generator d:\path\to\provider.go d:\path\to\provider_impl.go Provider api
```

## Features

- **Preserves Existing Code**: Never overwrites existing method implementations
- **Signature Validation**: Detects and reports signature mismatches with TODO comments
- **Interface Order**: Organizes methods in the same order as defined in the interface
- **Method Detection**: Automatically identifies which methods are already implemented
- **Missing Methods**: Generates skeletons for unimplemented methods
- **Error Handling**: Adds TODO comments for methods that need attention

## Generated Method Structure

For each missing interface method, the generator creates:

```go
// TODO: Implement <MethodName>
func (s *<StructName>) <MethodName>(<params>) (<returns>) {
    // TODO: Implement this method
    return <zero-values>
}
```

For methods with signature mismatches:

```go
// TODO: Signature mismatch with interface. Expected: <expected-signature>
func (s *<StructName>) <MethodName>(<current-params>) (<current-returns>) {
    // TODO: Update method signature to match interface
    return <zero-values>
}
```

## Implementation Notes

The generator:
1. Uses `go/parser` to parse the Go source file
2. Traverses the AST to find struct and interface definitions
3. Compares method signatures between struct and interface
4. Generates appropriate method implementations
5. Preserves all existing code and comments
6. Adds clear TODO markers for manual review
