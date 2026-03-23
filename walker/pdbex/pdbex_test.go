package pdbex

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

var (
	testPDBPath = `D:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\build\Release\hyperkd.pdb`
	testPEPath  = `D:\笔记本\p\ux\examples\hypedbg\doc\cpp\HyperDbgUnified\build\Release\hyperkd.sys`
)

func skipIfNoPDB(t *testing.T) {
	if _, err := os.Stat(testPDBPath); os.IsNotExist(err) {
		t.Skip("PDB file not found, skipping test")
	}
}

func skipIfNoPE(t *testing.T) {
	if _, err := os.Stat(testPEPath); os.IsNotExist(err) {
		t.Skip("PE file not found, skipping test")
	}
}

func TestPDBOpen(t *testing.T) {
	skipIfNoPDB(t)

	pdb := NewPDB()
	defer pdb.Close()

	err := pdb.Open(testPDBPath)
	if err != nil {
		t.Fatalf("Failed to open PDB: %v", err)
	}

	if !pdb.IsOpened() {
		t.Error("PDB should be opened")
	}

	t.Logf("PDB path: %s", pdb.GetPath())
	t.Logf("Architecture: %s", pdb.GetArchitectureString())
	t.Logf("Machine type: 0x%04X", pdb.GetMachineType())
}

func TestPDBGetSymbols(t *testing.T) {
	skipIfNoPDB(t)

	pdb := NewPDB()
	defer pdb.Close()

	if err := pdb.Open(testPDBPath); err != nil {
		t.Fatalf("Failed to open PDB: %v", err)
	}

	symbols := pdb.GetSortedSymbols()
	if len(symbols) == 0 {
		t.Error("No symbols found in PDB")
	}

	t.Logf("Total symbols: %d", len(symbols))

	count := 0
	for _, sym := range symbols {
		if count >= 10 {
			break
		}
		t.Logf("Symbol: %s (Tag: %s, Size: %d)", sym.Name, sym.Tag, sym.Size)
		count++
	}
}

func TestPDBGetUDTs(t *testing.T) {
	skipIfNoPDB(t)

	pdb := NewPDB()
	defer pdb.Close()

	if err := pdb.Open(testPDBPath); err != nil {
		t.Fatalf("Failed to open PDB: %v", err)
	}

	udts := pdb.GetUDTs()
	if len(udts) == 0 {
		t.Error("No UDTs found in PDB")
	}

	t.Logf("Total UDTs: %d", len(udts))

	count := 0
	for _, udt := range udts {
		if count >= 5 {
			break
		}
		if udt.Udt != nil {
			t.Logf("UDT: %s (Kind: %s, Fields: %d, Size: %d)",
				udt.Name, udt.Udt.Kind, udt.Udt.FieldCount, udt.Size)
			count++
		}
	}
}

func TestPDBGetEnums(t *testing.T) {
	skipIfNoPDB(t)

	pdb := NewPDB()
	defer pdb.Close()

	if err := pdb.Open(testPDBPath); err != nil {
		t.Fatalf("Failed to open PDB: %v", err)
	}

	enums := pdb.GetEnums()
	if len(enums) == 0 {
		t.Skip("No enums found in PDB")
	}

	t.Logf("Total enums: %d", len(enums))

	count := 0
	for _, enum := range enums {
		if count >= 5 {
			break
		}
		if enum.Enum != nil {
			t.Logf("Enum: %s (Fields: %d)", enum.Name, enum.Enum.FieldCount)
			count++
		}
	}
}

func TestPDBGetFunctions(t *testing.T) {
	skipIfNoPDB(t)

	pdb := NewPDB()
	defer pdb.Close()

	if err := pdb.Open(testPDBPath); err != nil {
		t.Fatalf("Failed to open PDB: %v", err)
	}

	functions := pdb.GetAllFunctions()
	if len(functions) == 0 {
		t.Error("No functions found in PDB")
	}

	t.Logf("Total functions: %d", len(functions))

	count := 0
	for name, fn := range functions {
		if count >= 10 {
			break
		}
		t.Logf("Function: %s (Address: 0x%X, Size: %d)", name, fn.Address, fn.Size)
		count++
	}
}

func TestPDBGetSymbolByName(t *testing.T) {
	skipIfNoPDB(t)

	pdb := NewPDB()
	defer pdb.Close()

	if err := pdb.Open(testPDBPath); err != nil {
		t.Fatalf("Failed to open PDB: %v", err)
	}

	udts := pdb.GetUDTs()
	if len(udts) == 0 {
		t.Skip("No UDTs to test")
	}

	testSym := udts[0]
	sym, ok := pdb.GetSymbolByName(testSym.Name)
	if !ok {
		t.Fatalf("Failed to get symbol by name: %s", testSym.Name)
	}

	if sym.Name != testSym.Name {
		t.Errorf("Symbol name mismatch: got %s, want %s", sym.Name, testSym.Name)
	}

	t.Logf("Found symbol: %s (Tag: %s, Size: %d)", sym.Name, sym.Tag, sym.Size)
}

func TestPDBGetFunctionByOffset(t *testing.T) {
	skipIfNoPDB(t)

	pdb := NewPDB()
	defer pdb.Close()

	if err := pdb.Open(testPDBPath); err != nil {
		t.Fatalf("Failed to open PDB: %v", err)
	}

	functions := pdb.GetAllFunctions()
	if len(functions) == 0 {
		t.Skip("No functions to test")
	}

	var testFn *FunctionInfo
	for _, fn := range functions {
		if fn.Address > 0 {
			testFn = fn
			break
		}
	}

	if testFn == nil {
		t.Skip("No function with valid address found")
	}

	name, ok := pdb.GetFunctionByOffset(testFn.Address)
	if !ok {
		t.Fatalf("Failed to get function by offset: 0x%X", testFn.Address)
	}

	t.Logf("Function at offset 0x%X: %s", testFn.Address, name)
}

func TestHeaderReconstruction(t *testing.T) {
	skipIfNoPDB(t)

	pdb := NewPDB()
	defer pdb.Close()

	if err := pdb.Open(testPDBPath); err != nil {
		t.Fatalf("Failed to open PDB: %v", err)
	}

	settings := DefaultSettings()
	settings.PrintHeader = true
	settings.PrintDeclarations = false
	settings.PrintFunctions = false

	reconstructor := NewHeaderReconstructor(settings)

	udts := pdb.GetUDTs()
	if len(udts) == 0 {
		t.Skip("No UDTs to reconstruct")
	}

	testSym := udts[0]
	header, err := reconstructor.ReconstructSymbol(pdb, testSym.Name)
	if err != nil {
		t.Fatalf("Failed to reconstruct symbol: %v", err)
	}

	t.Logf("Reconstructed header for %s:", testSym.Name)
	t.Log(header)
}

func TestHeaderReconstructionAll(t *testing.T) {
	skipIfNoPDB(t)

	pdb := NewPDB()
	defer pdb.Close()

	if err := pdb.Open(testPDBPath); err != nil {
		t.Fatalf("Failed to open PDB: %v", err)
	}

	settings := DefaultSettings()
	settings.PrintHeader = true
	settings.PrintDeclarations = false
	settings.PrintFunctions = false

	reconstructor := NewHeaderReconstructor(settings)

	header := reconstructor.ReconstructAll(pdb)

	lines := strings.Split(header, "\n")
	displayLines := min(len(lines), 50)

	t.Logf("Reconstructed header (first %d lines):", displayLines)
	for i := range displayLines {
		t.Log(lines[i])
	}

	t.Logf("Total lines: %d", len(lines))
}

func TestPEFileOpen(t *testing.T) {
	skipIfNoPE(t)

	pe, err := NewPEFile(testPEPath)
	if err != nil {
		t.Fatalf("Failed to open PE file: %v", err)
	}

	t.Logf("PE path: %s", testPEPath)
	t.Logf("Image base: 0x%X", pe.GetImageBase())

	sections := pe.GetSections()
	t.Logf("Number of sections: %d", len(sections))
	for _, sec := range sections {
		t.Logf("  Section: %s (VA: 0x%X, Size: 0x%X)", sec.Name, sec.VirtualAddress, sec.VirtualSize)
	}

	exports := pe.GetExports()
	t.Logf("Number of exports: %d", len(exports))
	for i, exp := range exports {
		if i >= 10 {
			break
		}
		t.Logf("  Export: %s (Address: 0x%X)", exp.Name, exp.Address)
	}
}

func TestPEIntegration(t *testing.T) {
	skipIfNoPDB(t)
	skipIfNoPE(t)

	integration, err := NewPEIntegration(testPDBPath, testPEPath)
	if err != nil {
		t.Fatalf("Failed to create PE integration: %v", err)
	}
	defer integration.Close()

	t.Logf("PDB path: %s", integration.GetPDB().GetPath())
	t.Logf("PE image base: 0x%X", integration.GetPE().GetImageBase())

	functions := integration.GetAllFunctionsWithAddresses()
	if len(functions) == 0 {
		t.Error("No functions found")
	}

	t.Logf("Total functions: %d", len(functions))

	count := 0
	for _, fn := range functions {
		if count >= 10 {
			break
		}
		t.Logf("Function: %s (RVA: 0x%X, VA: 0x%X, Size: %d)",
			fn.Name, fn.RVA, fn.VA, fn.Size)
		count++
	}
}

func TestPEIntegrationGetFunctionByRVA(t *testing.T) {
	skipIfNoPDB(t)
	skipIfNoPE(t)

	integration, err := NewPEIntegration(testPDBPath, testPEPath)
	if err != nil {
		t.Fatalf("Failed to create PE integration: %v", err)
	}
	defer integration.Close()

	functions := integration.GetAllFunctionsWithAddresses()
	if len(functions) == 0 {
		t.Skip("No functions to test")
	}

	testFn := functions[0]

	name, ok := integration.GetFunctionNameByRVA(uint32(testFn.RVA))
	if !ok {
		t.Fatalf("Failed to get function by RVA: 0x%X", testFn.RVA)
	}

	t.Logf("Function at RVA 0x%X: %s", testFn.RVA, name)

	fn := integration.FindFunctionContainingRVA(uint32(testFn.RVA))
	if fn == nil {
		t.Fatalf("Failed to find function containing RVA: 0x%X", testFn.RVA)
	}

	t.Logf("Found function: %s (RVA: 0x%X, Size: %d)", fn.Name, fn.RVA, fn.Size)
}

func TestPDBDumpSymbol(t *testing.T) {
	skipIfNoPDB(t)

	pdb := NewPDB()
	defer pdb.Close()

	if err := pdb.Open(testPDBPath); err != nil {
		t.Fatalf("Failed to open PDB: %v", err)
	}

	udts := pdb.GetUDTs()
	if len(udts) == 0 {
		t.Skip("No UDTs to dump")
	}

	testSym := udts[0]
	dump, err := pdb.DumpSymbol(testSym.Name)
	if err != nil {
		t.Fatalf("Failed to dump symbol: %v", err)
	}

	t.Logf("Dump of symbol %s:", testSym.Name)
	t.Log(dump)
}

func TestSymbolIsUnnamed(t *testing.T) {
	tests := []struct {
		name     string
		expected bool
	}{
		{"", true},
		{"<anonymous-tag>", true},
		{"<unnamed-tag>", true},
		{"__unnamed", true},
		{"MyStruct", false},
		{"_MyStruct", false},
	}

	for _, tt := range tests {
		sym := &Symbol{Name: tt.name}
		result := sym.IsUnnamed()
		if result != tt.expected {
			t.Errorf("IsUnnamed(%q) = %v, want %v", tt.name, result, tt.expected)
		}
	}
}

func TestSymTagString(t *testing.T) {
	tests := []struct {
		tag      SymTag
		expected string
	}{
		{SymTagUDT, "UDT"},
		{SymTagEnumType, "Enum"},
		{SymTagFunction, "Function"},
		{SymTagPointerType, "PointerType"},
		{SymTagArrayType, "ArrayType"},
		{SymTagTypedef, "Typedef"},
		{SymTag(999), "Unknown"},
	}

	for _, tt := range tests {
		result := tt.tag.String()
		if result != tt.expected {
			t.Errorf("SymTag(%d).String() = %q, want %q", tt.tag, result, tt.expected)
		}
	}
}

func TestBasicTypeString(t *testing.T) {
	tests := []struct {
		baseType  BasicType
		size      uint32
		useStdInt bool
		expected  string
	}{
		{btVoid, 0, false, "void"},
		{btChar, 1, false, "char"},
		{btInt, 4, false, "int"},
		{btInt, 4, true, "int32_t"},
		{btUInt, 4, false, "unsigned int"},
		{btUInt, 4, true, "uint32_t"},
		{btFloat, 4, false, "float"},
		{btFloat, 8, false, "double"},
		{btBool, 1, false, "bool"},
	}

	for _, tt := range tests {
		result := GetBasicTypeString(tt.baseType, tt.size, tt.useStdInt)
		if result != tt.expected {
			t.Errorf("GetBasicTypeString(%d, %d, %v) = %q, want %q",
				tt.baseType, tt.size, tt.useStdInt, result, tt.expected)
		}
	}
}

func TestUdtKindString(t *testing.T) {
	tests := []struct {
		kind     UdtKind
		expected string
	}{
		{UdtStruct, "struct"},
		{UdtClass, "class"},
		{UdtUnion, "union"},
		{UdtKind(999), "unknown"},
	}

	for _, tt := range tests {
		result := tt.kind.String()
		if result != tt.expected {
			t.Errorf("UdtKind(%d).String() = %q, want %q", tt.kind, result, tt.expected)
		}
	}
}

func TestCallingConventionString(t *testing.T) {
	tests := []struct {
		callConv CallingConvention
		expected string
	}{
		{CallConvNearC, "__cdecl"},
		{CallConvNearStd, "__stdcall"},
		{CallConvThisCall, "__thiscall"},
		{CallConvNearFast, "__fastcall"},
		{CallConvNearVector, "__vectorcall"},
		{CallingConvention(999), "unknown"},
	}

	for _, tt := range tests {
		result := tt.callConv.String()
		if result != tt.expected {
			t.Errorf("CallingConvention(%d).String() = %q, want %q", tt.callConv, result, tt.expected)
		}
	}
}

func TestPDBWithNonExistentFile(t *testing.T) {
	pdb := NewPDB()
	defer pdb.Close()

	err := pdb.Open(filepath.Join(os.TempDir(), "nonexistent.pdb"))
	if err == nil {
		t.Error("Expected error for non-existent PDB file")
	}
}

func TestPEWithNonExistentFile(t *testing.T) {
	_, err := NewPEFile(filepath.Join(os.TempDir(), "nonexistent.sys"))
	if err == nil {
		t.Error("Expected error for non-existent PE file")
	}
}

func TestFindSourceLineByRVA(t *testing.T) {
	skipIfNoPDB(t)

	pdb := NewPDB()
	defer pdb.Close()

	if err := pdb.Open(testPDBPath); err != nil {
		t.Fatalf("Failed to open PDB: %v", err)
	}

	functions := pdb.GetAllFunctions()
	if len(functions) == 0 {
		t.Skip("No functions to test")
	}

	var testFn *FunctionInfo
	for _, fn := range functions {
		if fn.Address > 0 && fn.Size > 0 {
			testFn = fn
			break
		}
	}

	if testFn == nil {
		t.Skip("No function with valid address found")
	}

	rva := uint32(testFn.Address)
	fileName, lineNumber, ok := pdb.FindSourceLineByRVA(rva)
	t.Logf("RVA 0x%X: fileName=%q, lineNumber=%d, ok=%v", rva, fileName, lineNumber, ok)

	if ok && fileName != "" {
		t.Logf("Found source: %s:%d", fileName, lineNumber)
	}
}
