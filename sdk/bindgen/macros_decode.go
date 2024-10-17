package bindgen

import (
	"github.com/ddkwork/golibrary/stream"
	"strings"
)

func extractMacros(lines []string) *stream.OrderedMap[string, string] {
	macros := stream.NewOrderedMap("", "")
	macros.Set("PAGE_SIZE", "4096")
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
