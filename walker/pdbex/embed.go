package pdbex

import _ "embed"

//go:embed test_dia/dia/bin/amd64/msdia140.dll
var msdia140x64 []byte

//go:embed test_dia/dia/bin/arm/msdia140.dll
var msdia140arm []byte

//go:embed test_dia/dia/bin/arm64/msdia140.dll
var msdia140arm64 []byte

//go:embed test_dia/dia/bin/msdia140.dll
var msdia140x86 []byte

//go:embed test_dia/dia/bin/amd64/symsrv.dll
var symsrvX64 []byte

//go:embed test_dia/dia/bin/symsrv.dll
var symsrvX86 []byte

//go:embed test_dia/dia/bin/arm64/symsrv.dll
var symsrvArm64 []byte
