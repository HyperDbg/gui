
# Zydis Bindings for Go [![GoDoc](https://pkg.go.dev/badge/go.jpap.org/zydis.svg)](https://pkg.go.dev/go.jpap.org/zydis)

# Import

```go
import "go.jpap.org/zydis"
```
# Overview

Zydis is a Go wrapper for the fast and lightweight Zydis x86/x86-64
disassembler library, found at http://zydis.re/.  This package provides
bindings via cgo and is considered a "complete" wrapper of the Zydis API,
ready for production use.

It was created because the pure-Go disassembler, found at
https://godoc.org/golang.org/x/arch/x86/x86asm, is significantly lacking in
x86-64 support.  Decoding x86* is complex business, and it was more
straightforward to write these bindings instead of digging deep into the
pure-Go package.

## Requires Git-LFS
**This package uses Git LFS found to store a precompiled version of the Zydis
library (see below), so please make sure you have it installed before
getting this package.**  Learn about Git LFS at https://git-lfs.github.com/.

## Sample Code
See the file `cmd/demo.go`.

## Upgrading the Zydis Library
The Zydis library is packaged as a static syso object file so that this
package is go gettable.  Precompiled macOS (amd64, arm64), Linux (amd64,
arm64), and Windows (amd64, 386) binaries are provided.

Use the Makefile in the `lib/` folder to upgrade to a newer version,
rebuild, or add support for another platform.  The default Makefile target
clones the Zydis repo and its submodule, performs the build, and creates the
syso files for Go linkage under macOS with suitable cross-compilers
installed.




