package sdk

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ddkwork/golibrary/stream/orderedmap"

	"github.com/ddkwork/golibrary/stream"

	"github.com/ddkwork/golibrary/mylog"
)

// readLines reads a whole file into memory and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file := mylog.Check2(os.Open(path))

	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// writeLines writes the lines to the given file.
func writeLines(lines []string, path string) error {
	file := mylog.Check2(os.Create(path))

	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}

// extractMacros extracts macros from a file, handling multi-line macros.
func extractMacros(lines []string) *orderedmap.OrderedMap[string, string] {
	macros := orderedmap.New[string, string]()
	var macroName string
	var macroValue strings.Builder
	inMacro := false

	for _, line := range lines {
		if strings.HasPrefix(line, "#define") {
			if inMacro {
				// Finish the previous macro
				macros.Set(macroName, macroValue.String())
			}
			inMacro = true
			parts := strings.Fields(line)
			if len(parts) > 1 {
				macroName = parts[1]
				macroValue.Reset()
				if len(parts) > 2 {
					macroValue.WriteString(strings.Join(parts[2:], " "))
				}
			}
		} else if inMacro {
			if strings.HasPrefix(line, " ") || strings.HasPrefix(line, "\t") {
				// Continuation of the current macro
				macroValue.WriteString(" ")
				macroValue.WriteString(strings.TrimSpace(line))
			} else {
				// Finish the current macro
				macros.Set(macroName, macroValue.String())
				inMacro = false
			}
		}
	}

	if inMacro {
		macros.Set(macroName, macroValue.String())
	}

	return macros
}

// normalizeMacro normalizes a macro definition by removing extra spaces.
func normalizeMacro(macro string) string {
	parts := strings.Fields(macro)
	return strings.Join(parts, " ")
}

func main() {
	headerFile := "merged_headers.h"
	macroFile := "macros.log"
	outputFile := "newMacros.log"

	headerLines := mylog.Check2(readLines(headerFile))

	macroLines := mylog.Check2(readLines(macroFile))

	headerMacros := extractMacros(headerLines)
	dmMacros := extractMacros(macroLines)
	mylog.Trace(headerFile, headerMacros.Len())
	mylog.Trace(macroFile, dmMacros.Len())

	// Filter dmMacros based on headerMacros
	for _, kv := range headerMacros.List() {
		normalizedHeaderValue := normalizeMacro(kv.Value)
		normalizedHeaderValue = strings.ReplaceAll(normalizedHeaderValue, "//todo", "")
		if dmValue, exists := dmMacros.Get(kv.Key); exists {
			normalizedDmValue := normalizeMacro(dmValue)

			switch {
			case strings.Contains(normalizedHeaderValue, "\\"):
				if normalizedDmValue == "0x4859504552444247" {
					mylog.Todo("debug")
				}
				split := strings.Split(normalizedHeaderValue, "\\")
				keep := false
				for _, s := range split {
					s = strings.TrimSpace(s)
					if len(s) < 2 {
						continue
					}
					if strings.Contains(s, "//") {
						before, _, found := strings.Cut(s, "//")
						if found {
							s = before
							s = strings.TrimSuffix(s, " ")
							if normalizedDmValue == s {
								keep = true
								break
							}
						}

					}
					if strings.Contains(normalizedDmValue, s) {
						keep = true
						break
					}
				}
				if !keep {
					mylog.Warning("delete macro", kv.Key)
					dmMacros.Delete(kv.Key)
				}
			default:
				if normalizedDmValue != normalizedHeaderValue {
					mylog.Warning("delete macro", kv.Key)
					dmMacros.Delete(kv.Key)
				}
			}

		}
	}

	var filteredLines []string
	for _, line := range macroLines {
		if strings.HasPrefix(line, "#define") {
			parts := strings.Fields(line)
			if len(parts) > 1 {
				_, exists := dmMacros.Get(parts[1])
				if exists {
					filteredLines = append(filteredLines, line)
				}
			}
		} else {
			filteredLines = append(filteredLines, line)
		}
	}
	mylog.Check(writeLines(filteredLines, outputFile))

	if macroCount(headerFile) != macroCount(outputFile) {
		mylog.Check("macro count mismatch")
	}

	fmt.Println("Macro definitions filtered successfully.")
}

func macroCount(path string) int {
	return strings.Count(stream.NewBuffer(path).String(), "#define ")
}
