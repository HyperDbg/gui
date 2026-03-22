---
name: "test-generator"
description: "Generates unit test skeletons from Go interface definitions. Invoke when user wants to create test files with empty test functions based on interface methods."
---

# Unit Test Generator

This skill generates unit test skeleton files from Go interface definitions.

## When to Use

Invoke this skill when:
- User wants to generate unit test templates from a Go file containing interface definitions
- User needs test functions for all exported methods in an interface
- User wants test functions in the same order as defined in the interface
- User is unsatisfied with IDE-generated test templates

## How It Works

1. **Parse Input File**: Reads the specified Go source file
2. **Extract Interface**: Finds the interface definition (e.g., `type api interface { ... }`)
3. **Identify Exported Methods**: Extracts all method signatures that start with uppercase letters
4. **Generate Test Functions**: Creates empty test functions for each exported method
5. **Determine Package Name**: Uses the basename of the input file's directory as the package name
6. **Write Output**: Saves the generated test file to the specified output path

## Usage

### Command Format

```
test-generator <input-file> <output-file> [interface-name]
```

### Parameters

| Parameter | Required | Description |
|-----------|----------|-------------|
| `input-file` | Yes | Path to the Go source file containing the interface definition |
| `output-file` | Yes | Path where the generated test file should be saved |
| `interface-name` | No | Name of the interface to parse (default: `api`) |

### Example

```bash
# Generate tests from provider.go
test-generator d:\path\to\provider.go d:\path\to\provider_test.go

# Generate tests from a specific interface
test-generator d:\path\to\provider.go d:\path\to\provider_test.go api
```

## Generated Test Structure

For each exported method in the interface, the generator creates:

```go
func TestProvider_<MethodName>(t *testing.T) {
    type fields struct {
        // Fields from the Provider struct
    }
    type args struct {
        // Method parameters
    }
    tests := []struct {
        name    string
        fields  fields
        args    args
        want    <return-type>
        wantErr bool
    }{
        // TODO: Add test cases.
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            p := &Provider{
                // Initialize fields
            }
            got, err := p.<MethodName>(tt.args...)
            if (err != nil) != tt.wantErr {
                t.Errorf("<MethodName>() error = %v, wantErr %v", err, tt.wantErr)
                return
            }
            if !reflect.DeepEqual(got, tt.want) {
                t.Errorf("<MethodName>() got = %v, want %v", got, tt.want)
            }
        })
    }
}
```

## Features

- **Preserves Interface Order**: Test functions are generated in the same order as methods appear in the interface
- **Automatic Package Detection**: Uses the directory basename as the package name
- **Complete Test Structure**: Includes fields, args, and return value placeholders
- **TODO Markers**: Clearly marks where test cases should be added
- **Error Handling Support**: Automatically includes `wantErr` field for error-returning methods

## Implementation Notes

The generator:
1. Uses `go/parser` to parse the Go source file
2. Traverses the AST to find interface definitions
3. Filters for exported methods (uppercase first letter)
4. Generates appropriate test function signatures based on method signatures
5. Handles multiple return values and error types
