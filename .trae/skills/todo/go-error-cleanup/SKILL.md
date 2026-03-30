---
name: "go-error-cleanup"
description: "Removes error returns and replaces with mylog.Check/Check2/Call. Invoke when user wants to clean up error handling in Go files."
---

# Go Error Cleanup

This skill removes error returns and replaces error handling with mylog utilities across Go files in the workspace.

## What It Does

1. **Removes error return types** from function signatures
2. **Replaces error checks** with `mylog.Check()` for simple error handling
3. **Replaces error returns** with `mylog.Check2()` for functions returning values
4. **Replaces defer recover() blocks** with `mylog.Call()` wrapper
5. **Removes panic() calls** (except in mylog itself)
6. **Removes recover() calls**
7. **Removes unused imports** like `fmt` (if only used for error formatting)

## When to Invoke

- User wants to clean up error handling in Go code
- User mentions "remove error returns" or "use mylog.Check"
- User wants to apply the same pattern as in `provider.go` to other files
- User wants to eliminate error types, panic, and recover from Go files

## Transformation Rules

### 1. Function Signatures
```go
// Before
func foo() error
func foo() (int, error)
func foo() (string, error)

// After
func foo()
func foo() int
func foo() string
```

### 2. Interface Method Signatures
```go
// Before
type MyInterface interface {
    Foo() error
    Bar() (int, error)
}

// After
type MyInterface interface {
    Foo()
    Bar() int
}
```

### 3. Simple Error Checks
```go
// Before
err := someFunc()
if err != nil {
    return fmt.Errorf("failed: %w", err)
}

// After
mylog.Check(someFunc())
```

### 4. Error Returns with Values
```go
// Before
val, err := someFunc()
if err != nil {
    return err
}
return val

// After
val := mylog.Check2(someFunc())
return val
```

### 5. Defer Recover Blocks
```go
// Before
defer func() {
    if r := recover(); r != nil {
        if err, ok := r.(error); ok {
            // handle error
        }
    }
}()

// After
mylog.Call(func() {
    // code that might panic
})
```

### 6. Remove Error Returns
```go
// Before
func foo() error {
    err := bar()
    if err != nil {
        return err
    }
    return nil
}

// After
func foo() {
    mylog.Check(bar())
}
```

### 7. Wrapper Methods (Interface Delegation)
```go
// Before - Wrapper method that delegates to interface
func (s *MyStruct) SomeMethod() error {
    return s.interface.Method()
}

// After - Remove error return and use mylog.Check
func (s *MyStruct) SomeMethod() {
    mylog.Check(s.interface.Method())
}

// Before - Wrapper method returning value
func (s *MyStruct) GetValue() (int, error) {
    return s.interface.GetValue()
}

// After - Use mylog.Check2 for value-returning methods
func (s *MyStruct) GetValue() int {
    return mylog.Check2(s.interface.GetValue())
}
```

### 8. Already Processed Methods (No Double Wrapping)
```go
// WRONG - Don't wrap methods that already use mylog internally
func (s *MyStruct) SomeMethod() {
    mylog.Check(s.interface.Method())  // Method() already uses mylog.Check internally
}

// CORRECT - Just call the method directly
func (s *MyStruct) SomeMethod() {
    s.interface.Method()  // Method() already handles errors with mylog.Check
}

// Example: provider.go already processed
// SendBuffer and ReceiveBuffer already use mylog.Check internally
// So in packet.go, just call them directly:
func (p *Packet) EPTHook(address uint64, size uint32, hookType EPTHookType) {
    if !p.IsConnected() {
        slog.Warn("driver not available")
        return
    }
    buffer := make([]byte, 24)
    binary.LittleEndian.PutUint64(buffer[0:8], address)
    binary.LittleEndian.PutUint32(buffer[8:12], size)
    binary.LittleEndian.PutUint32(buffer[12:16], uint32(hookType))
    
    p.driver.SendBuffer(buffer, IoctlDebuggerRegisterEvent)  // Already uses mylog.Check
}
```

## Important Notes

1. **Automatic file type detection**: The skill should automatically analyze each file to determine if it's an interface definition, implementation, or wrapper
2. **Dependency-based processing order**: Process files based on their dependencies, not hardcoded filenames
3. **Interface definitions first**: Always process files that define interfaces before files that implement them
4. **Implementations before wrappers**: Process files with actual logic before files that just delegate to interfaces
5. **Avoid duplicate work**: By processing in dependency order, you avoid modifying the same method multiple times
6. **Recursive scanning**: Scan all subdirectories to find all Go files, not just the root directory
7. **Don't double-wrap with mylog**: If a called method already uses mylog.Check internally, don't wrap it again with mylog.Check
8. **Check called method signatures**: Before wrapping with mylog.Check, check if the called method still returns error

## Implementation Steps

1. **Scan workspace recursively** for all `.go` files (including subdirectories)
2. **Analyze file types and dependencies**:
   - Identify interface definition files (files that define `type X interface`)
   - Identify implementation files (files that implement interface methods with actual logic)
   - Identify wrapper files (files that delegate to other interfaces)
   - Build dependency graph to determine processing order
3. **Process interface definition files first** (files that define interfaces)
   - Remove error return types from interface method signatures
4. **Process implementation files** (files with actual logic, not just delegation)
   - These are files that implement interfaces with real code, not just calling other methods
   - Replace defer recover() with mylog.Call()
   - Replace error checks with mylog.Check() or mylog.Check2()
   - Remove error return types from function signatures
   - Remove panic() and recover() calls
   - Remove unused imports (fmt if only used for errors)
5. **Process wrapper files** (files that delegate to interfaces)
   - These are files that just call methods from other interfaces
   - Remove error returns from wrapper methods
   - Replace `return s.someInterface.Method()` with `mylog.Check(s.someInterface.Method())`
   - Replace `return s.someInterface.Method()` with `mylog.Check2(s.someInterface.Method())` for methods returning values
6. **Process remaining Go files** with error handling
7. **Verify changes** by checking for remaining error types
8. **Report summary** of changes made

## File Type Detection Rules

### Interface Definition Files
- Contain `type X interface { ... }` definitions
- No struct implementations
- Only method signatures

### Implementation Files
- Contain struct definitions with methods
- Methods contain actual logic (IOCTL calls, data processing, etc.)
- NOT just delegating to other interface methods

### Wrapper Files
- Contain struct definitions that embed interfaces
- Methods mostly just call methods from embedded interfaces
- Example: `func (s *Wrapper) Foo() { return s.embedded.Foo() }`

## Files to Skip

- `mylog/*.go` - Don't modify the mylog package itself
- Test files (`*_test.go`) - Tests may need explicit error handling
- Files that already use mylog.Check pattern

## Example Output

After processing, report:
- Number of files processed
- Number of error returns removed
- Number of mylog.Check/Check2/Call added
- Number of imports removed
- List of files modified
