# HyperDbg Code Generator

统一代码生成工具，用于生成 Rust 驱动代码、WDK 绑定验证和 MCP 服务器。

## 文件结构

```
rustgen/
├── main.go          # 主入口，调用所有生成器
├── rustgen.go       # Rust 类型生成器
├── wdkgen.go        # WDK 绑定验证器
├── mcpgen.go        # MCP 服务器生成器
├── bindgen_test.go  # 测试文件
├── ntddk.rs         # WDK 函数定义 (输入)
├── types.rs         # WDK 类型定义 (输入)
├── constants.rs     # WDK 常量定义 (输入)
└── README.md        # 本文档
```

## 使用方法

```bash
cd cmd/rustgen
go build -o rustgen.exe .
./rustgen.exe
```

**注意**: 每次运行前会自动删除 `rust-driver/kd/src/generated/` 目录中的旧文件。

## 生成器说明

### 1. Rust 类型生成器 (rustgen.go)

**功能**: 从 Go 代码生成 Rust 类型和处理器

**入口函数**: `GenerateRustTypes(projectRoot string) error`

**生成内容**:
- `rust-driver/kd/src/generated/mod.rs` - 统一模块入口
- `rust-driver/kd/src/generated/common.rs` - 通用类型
- `rust-driver/kd/src/generated/register.rs` - 寄存器类型
- `rust-driver/kd/src/generated/event_*.rs` - 事件类型
- `rust-driver/kd/src/generated/router.rs` - API 路由器
- `rust-driver/kd/src/generated/emit.rs` - 事件发射器
- `rust-driver/kd/src/generated/ioctl.rs` - IOCTL 代码

**数据流**:
```
debugger/*.go -> AST 解析 -> 类型提取 -> Rust 代码生成
```

### 2. WDK 绑定验证器 (wdkgen.go)

**功能**: 验证 Rust 代码中的 WDK extern 声明与 wdk-sys 绑定一致性

**入口函数**: `GenerateBindgen(projectRoot string) error`

**生成内容**:
- `rust-driver/kd/src/generated/functions_list.txt` - 函数列表
- `rust-driver/kd/src/generated/ntddk.go` - Go 函数映射
- `rust-driver/kd/src/generated/types.go` - Go 类型映射
- `rust-driver/kd/src/generated/constants.go` - Go 常量映射
- `rust-driver/kd/src/generated/validation_report.txt` - 验证报告

**数据流**:
```
ntddk.rs/types.rs/constants.rs -> WDK 绑定解析
rust-driver/kd/src/*.rs -> Rust extern 扫描
-> 签名比对 -> 验证报告
```

### 3. MCP 服务器生成器 (mcpgen.go)

**功能**: 从 Debugger 接口生成 MCP 服务器代码

**入口函数**: `GenerateMCP(projectRoot string) error`

**生成内容**:
- `cmd/mcp/mcp.go` - MCP 服务器实现

**数据流**:
```
debugger/interfaces.go -> 接口解析 -> MCP 服务器生成
```

## 输入文件

| 文件 | 用途 | 使用者 |
|------|------|--------|
| `ntddk.rs` | WDK 函数签名 | wdkgen.go |
| `types.rs` | WDK 类型定义 | wdkgen.go |
| `constants.rs` | WDK 常量定义 | wdkgen.go |
| `debugger/*.go` | Go 类型定义 | rustgen.go, mcpgen.go |

## 类型定义

### rustgen.go
```go
type APIMethod struct { ... }      // API 方法信息
type APIParam struct { ... }       // API 参数
type MethodInfo struct { ... }     // 接口方法信息
type ConstantDef struct { ... }    // 常量定义
type EnumDef struct { ... }        // 枚举定义
type StructDef struct { ... }      // 结构体定义
type TypeAliasDef struct { ... }   // 类型别名
```

### wdkgen.go
```go
type FunctionSignature struct { ... }  // WDK 函数签名
type TypeDefinition struct { ... }     // WDK 类型定义
type ExternFunction struct { ... }     // Rust extern 函数
type ValidationReport struct { ... }   // 验证报告
type Bindgen struct { ... }            // 主验证器
```

### mcpgen.go
```go
type Method struct { ... }             // MCP 方法
type Param struct { ... }              // MCP 参数
type InterfaceConfig struct { ... }    // 接口配置
```

## 已知问题

### 重复代码

以下函数存在重复或相似实现：

| 函数 | 位置 | 说明 |
|------|------|------|
| `exprToString` | rustgen.go:1612 | AST 表达式转字符串 |
| `typeToString` | mcpgen.go:448 | AST 类型转字符串 (相似功能) |

**建议**: 可以统一到 `main.go` 作为公共工具函数。

## 测试

```bash
go test -v ./...
```

主要测试用例：
- `TestLoadWdkBindings` - WDK 绑定加载
- `TestScanRustExterns` - Rust extern 扫描
- `TestValidate` - 签名验证
- `TestGenerateAll` - 完整生成流程
