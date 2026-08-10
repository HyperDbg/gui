# Go Subset + AST Binary Protocol Specification

This document is the contract between the **Go user-mode compiler**
(`go-bridge/ast`) and the **kernel C interpreter**
(`hyperkd/code/go-interp/interp.c`). Both sides MUST agree byte-for-byte with
this spec.

## 1. Go Subset Definition

### 1.1 Supported Syntax

| Category | Syntax |
|---|---|
| Variable declaration | `var x uint64 = 0`, `x := 1` |
| Types | `uint8/16/32/64`, `int8/16/32/64`, `bool`, `string` (read-only literal) |
| Arithmetic | `+ - * / %` |
| Bitwise | `& \| ^ << >> &^` |
| Comparison | `== != < > <= >=` |
| Logical | `&& \|\| !` |
| Control flow | `if/else`, `for` (three-clause), `break/continue/return` |
| Function declaration | `func name(args) (rets) { ... }` |
| Simple closure | `func(ctx *HookCtx) { ... }` (captures **only** ctx) |
| Whitelist calls | see §3 |
| Fixed-size array | `[N]uint8` |

### 1.2 Unsupported (VMX-root constraints)

`goroutine`, `channel`, `select`, `interface`, type assertion, `map`,
`slice` (except `[N]uint8`), `reflect`, `unsafe`, `defer`, string
concatenation (`+`), `make`, `new`, closures capturing external variables
(other than `ctx`).

### 1.3 Closure Capture Rule

A `FuncLit` node may reference **only** `ctx` or its own parameters. The
user-mode validator (`go-bridge/subset`) statically rejects any `FuncLit`
that references an identifier declared outside its body (excluding `ctx`).

## 2. AST Binary Protocol

### 2.1 Node Opcodes

Each node begins with a 1-byte opcode.

| Opcode | Node | Payload |
|---|---|---|
| `0x00` | Nil | (empty — used for absent Else) |
| `0x01` | Literal | 1 byte kind + N bytes value |
| `0x02` | Ident | 2 bytes name_id |
| `0x03` | BinaryExpr | 1 byte op + LHS + RHS |
| `0x04` | UnaryExpr | 1 byte op + Operand |
| `0x05` | CallExpr | 2 bytes func_id + 1 byte nargs + N children |
| `0x06` | SelectorExpr | X + 2 bytes field_id |
| `0x07` | AssignStmt | 1 byte op + LHS + RHS |
| `0x08` | IfStmt | Cond + Then + Else |
| `0x09` | ForStmt | Init + Cond + Post + Body |
| `0x0A` | BlockStmt | 2 bytes nstmts + N statements |
| `0x0B` | ReturnStmt | 1 byte nvals + N children |
| `0x0C` | FuncDecl | 2 bytes name_id + 1 byte nargs + N args + 1 byte nrets + N rets + Body |
| `0x0D` | FuncLit | 1 byte nargs + N args + 1 byte nrets + N rets + Body |
| `0x0E` | DeclStmt | 1 byte kind + 2 bytes name_id + Expr |
| `0x0F` | ArrayType | 4 bytes len + 1 byte elem_kind |
| `0x10` | CompositeLit | Type + 2 bytes nels + N elements |
| `0x11` | IndexExpr | X + Index |

### 2.2 Literal Kinds

| Kind | Value bytes | Go type |
|---|---|---|
| `0x01` | 1 | `uint8` |
| `0x02` | 2 | `uint16` |
| `0x03` | 4 | `uint32` |
| `0x04` | 8 | `uint64` |
| `0x05` | 1 | `int8` |
| `0x06` | 2 | `int16` |
| `0x07` | 4 | `int32` |
| `0x08` | 8 | `int64` |
| `0x09` | 1 | `bool` (0=false, 1=true) |
| `0x0A` | 2 bytes len + N | `string` |

### 2.3 Binary Operators

| Op | Symbol | | Op | Symbol |
|---|---|---|---|---|
| `0x01` | `+` | | `0x09` | `<<` |
| `0x02` | `-` | | `0x0A` | `>>` |
| `0x03` | `*` | | `0x0B` | `==` |
| `0x04` | `/` | | `0x0C` | `!=` |
| `0x05` | `%` | | `0x0D` | `<` |
| `0x06` | `&` | | `0x0E` | `>` |
| `0x07` | `\|` | | `0x0F` | `<=` |
| `0x08` | `^` | | `0x10` | `>=` |
| | | | `0x11` | `&&` |
| | | | `0x12` | `\|\|` |
| | | | `0x13` | `&^` |

### 2.4 Unary Operators

| Op | Symbol |
|---|---|
| `0x01` | `-` (negate) |
| `0x02` | `!` (logical not) |
| `0x03` | `^` (bitwise not) |

### 2.5 Assignment Operators

| Op | Symbol |
|---|---|
| `0x01` | `=` |
| `0x02` | `+=` |
| `0x03` | `-=` |
| `0x04` | `*=` |
| `0x05` | `/=` |
| `0x06` | `%=` |
| `0x07` | `&=` |
| `0x08` | `\|=` |
| `0x09` | `^=` |
| `0x0A` | `<<=` |
| `0x0B` | `>>=` |
| `0x0C` | `&^=` |

### 2.6 String Table

Appended after the AST node stream:

```
[4 bytes] count          — number of strings
For each string:
  [2 bytes] len           — string length (little-endian)
  [N bytes] data          — UTF-8 bytes (not null-terminated)
```

`name_id` and `field_id` values in Ident/SelectorExpr/FuncDecl nodes are
**indices** into this table.

## 3. Whitelist Functions

HookCtx methods exposed to the interpreter. `func_id` is the 2-byte value in
CallExpr when the callee is a SelectorExpr `ctx.Method`.

| func_id | Method | Signature |
|---|---|---|
| `0x0001` | `StackReadQword` | `(offset uint32) uint64` |
| `0x0002` | `StackReadDword` | `(offset uint32) uint32` |
| `0x0003` | `Reg` | `(name string) uint64` |
| `0x0004` | `SetReg` | `(name string, val uint64)` |
| `0x0005` | `ReadMem` | `(addr uint64, buf []byte)` |
| `0x0006` | `ReadMemQword` | `(addr uint64) uint64` |
| `0x0007` | `Printf` | `(fmt string, args ...uint64)` |
| `0x0008` | `Break` | `()` |
| `0x0009` | `Continue` | `()` |
| `0x000A` | `GetPid` | `() uint32` |
| `0x000B` | `GetTid` | `() uint32` |
| `0x000C` | `GetIP` | `() uint64` |

`Printf` supports only `%x`, `%d`, `%s`, `%#x`. Output goes to a kernel log
buffer (no direct I/O in VMX-root).

## 4. Cycle Budget

Every AST node execution increments a cycle counter. When the counter reaches
**1,000,000**, the interpreter aborts and returns `Continue` (the default
action). This prevents `for {}` from BSOD-ing the system.

## 5. Error Isolation

The interpreter runs inside `__try/__except` (SEH). Any fault (stack overflow,
type error, unknown node, memory access failure) is caught and converted to a
logged error + `Continue` return. **Script errors must never BSOD.**

## 6. Wire Format

A complete hook callback payload:

```
[N bytes] ast_nodes      — node stream (root opcode at position 0)
[4 bytes] str_count      — string table count
[...]    strings         — string table entries
```

The user-mode encoder (`go-bridge/ast`) produces this; the kernel decoder
(`ast_decode.c::AstDecode`) consumes it. The kernel starts reading the root
opcode at position 0 — there is no ast_size prefix.
