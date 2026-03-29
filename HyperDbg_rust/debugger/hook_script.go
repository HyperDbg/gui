package debugger

import (
	"bufio"
	"os"
	"reflect"
	"runtime"
	"strings"
	"sync"
)

type sourceCache struct {
	mu    sync.RWMutex
	files map[string][]string
}

var globalSourceCache = &sourceCache{
	files: make(map[string][]string),
}

func (sc *sourceCache) getLines(file string) []string {
	sc.mu.RLock()
	lines, ok := sc.files[file]
	sc.mu.RUnlock()

	if ok {
		return lines
	}

	sc.mu.Lock()
	defer sc.mu.Unlock()

	if lines, ok := sc.files[file]; ok {
		return lines
	}

	f, err := os.Open(file)
	if err != nil {
		return nil
	}
	defer f.Close()

	var result []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	if len(result) > 0 {
		sc.files[file] = result
	}

	return result
}

func ExtractClosureBody(fn HookScriptHandler) string {
	v := reflect.ValueOf(fn)
	if v.Kind() != reflect.Func {
		return ""
	}

	fpc := runtime.FuncForPC(v.Pointer())
	if fpc == nil {
		return ""
	}

	file, startLine := fpc.FileLine(v.Pointer())
	if file == "" {
		return ""
	}

	lines := globalSourceCache.getLines(file)
	if len(lines) == 0 {
		return ""
	}

	if startLine < 1 || startLine > len(lines) {
		return ""
	}

	var bodyLines []string
	braceCount := 0
	foundOpen := false
	inClosure := false

	for i := startLine - 1; i < len(lines); i++ {
		line := lines[i]
		trimmed := strings.TrimSpace(line)

		if !inClosure {
			if strings.Contains(line, "func(") || strings.Contains(line, "func (") {
				inClosure = true
			} else {
				continue
			}
		}

		if !foundOpen {
			if idx := strings.Index(line, "{"); idx >= 0 {
				foundOpen = true
				braceCount = strings.Count(line[idx:], "{") - strings.Count(line[idx:], "}")
				afterBrace := strings.TrimSpace(line[idx+1:])
				if afterBrace != "" && afterBrace != "}" {
					bodyLines = append(bodyLines, afterBrace)
				}
				if braceCount <= 0 {
					break
				}
			}
		} else {
			braceCount += strings.Count(line, "{")
			braceCount -= strings.Count(line, "}")

			if braceCount <= 0 {
				if idx := strings.LastIndex(line, "}"); idx >= 0 {
					beforeBrace := strings.TrimSpace(line[:idx])
					if beforeBrace != "" {
						bodyLines = append(bodyLines, beforeBrace)
					}
				}
				break
			}
			bodyLines = append(bodyLines, trimmed)
		}
	}

	return strings.Join(bodyLines, "\n")
}

func ExtractClosureBodyRaw(fn HookScriptHandler) string {
	v := reflect.ValueOf(fn)
	if v.Kind() != reflect.Func {
		return ""
	}

	fpc := runtime.FuncForPC(v.Pointer())
	if fpc == nil {
		return ""
	}

	file, startLine := fpc.FileLine(v.Pointer())
	if file == "" {
		return ""
	}

	lines := globalSourceCache.getLines(file)
	if len(lines) == 0 {
		return ""
	}

	if startLine < 1 || startLine > len(lines) {
		return ""
	}

	var bodyLines []string
	braceCount := 0
	foundOpen := false
	inClosure := false

	for i := startLine - 1; i < len(lines); i++ {
		line := lines[i]

		if !inClosure {
			if strings.Contains(line, "func(") || strings.Contains(line, "func (") {
				inClosure = true
			} else {
				continue
			}
		}

		if !foundOpen {
			if idx := strings.Index(line, "{"); idx >= 0 {
				foundOpen = true
				braceCount = strings.Count(line[idx:], "{") - strings.Count(line[idx:], "}")
				afterBrace := line[idx+1:]
				bodyLines = append(bodyLines, afterBrace)
				if braceCount <= 0 {
					break
				}
			}
		} else {
			braceCount += strings.Count(line, "{")
			braceCount -= strings.Count(line, "}")

			if braceCount <= 0 {
				if idx := strings.LastIndex(line, "}"); idx >= 0 {
					bodyLines = append(bodyLines, line[:idx])
				}
				break
			}
			bodyLines = append(bodyLines, line)
		}
	}

	return strings.Join(bodyLines, "\n")
}

func ClearSourceCache() {
	globalSourceCache.mu.Lock()
	defer globalSourceCache.mu.Unlock()
	globalSourceCache.files = make(map[string][]string)
}
