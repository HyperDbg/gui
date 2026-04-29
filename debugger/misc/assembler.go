package misc

import (
	"fmt"
	"regexp"
	"strings"
)

type AssembleData struct {
	AsmRaw    string
	ParsedAsm string
	Opcode    []byte
	IsValid   bool
	Error     string
}

type Assembler struct {
	symbolResolver func(expr string) (uint64, bool)
}

func NewAssembler() *Assembler {
	return &Assembler{}
}

func (asmblr *Assembler) SetSymbolResolver(resolver func(expr string) (uint64, bool)) {
	asmblr.symbolResolver = resolver
}

func (asmblr *Assembler) Assemble(asmCode string) *AssembleData {
	result := &AssembleData{
		AsmRaw: asmCode,
	}

	parsed := asmblr.parseAssemblyData(asmCode)
	result.ParsedAsm = parsed

	opcode, err := asmblr.assembleToBytes(parsed)
	if err != nil {
		result.IsValid = false
		result.Error = err.Error()
		return result
	}

	result.Opcode = opcode
	result.IsValid = true
	return result
}

func (asmblr *Assembler) parseAssemblyData(rawAsm string) string {
	rawAsm = strings.ReplaceAll(rawAsm, "\n", " ")

	multipleSpaces := regexp.MustCompile(" +")
	rawAsm = multipleSpaces.ReplaceAllString(rawAsm, " ")

	instructions := strings.Split(rawAsm, ";")

	var processedInstructions []string
	for _, instruction := range instructions {
		instruction = strings.TrimSpace(instruction)
		if len(instruction) == 0 {
			continue
		}

		processed := asmblr.processSymbolsInInstruction(instruction)
		processedInstructions = append(processedInstructions, processed)
	}

	return strings.Join(processedInstructions, ";")
}

func (asmblr *Assembler) processSymbolsInInstruction(instruction string) string {
	result := instruction

	for {
		startIdx := strings.Index(result, "<")
		if startIdx == -1 {
			break
		}

		endIdx := strings.Index(result[startIdx:], ">")
		if endIdx == -1 {
			break
		}

		endIdx += startIdx
		expr := result[startIdx+1 : endIdx]

		var addr uint64
		var found bool

		if asmblr.symbolResolver != nil {
			addr, found = asmblr.symbolResolver(expr)
		}

		if !found || addr == 0 {
			fmt.Printf("err, failed to resolve the symbol [ %s ].\n", expr)
			result = result[:startIdx] + result[endIdx+1:]
			continue
		}

		addrStr := fmt.Sprintf("0x%x", addr)
		result = result[:startIdx] + addrStr + result[endIdx+1:]
	}

	return result
}

func (asmblr *Assembler) assembleToBytes(assembly string) ([]byte, error) {
	return nil, nil
}

func (asmblr *Assembler) Disassemble(opcode []byte, baseAddress uint64) ([]string, error) {
	return nil, nil
}
