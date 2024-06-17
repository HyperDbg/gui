package control

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"testing"

	"github.com/ddkwork/golibrary/mylog"
)

// todo use cc pkg get ast, func struct var method object etc
func extractFunctionNamesFromFile(filename string) ([]string, error) {
	content := mylog.Check2(ioutil.ReadFile(filename))

	re := regexp.MustCompile(`(?m:^\s*[\w:<>]+\s+\*?\s*([\w:<>]+)\s*\(.*\))`)

	matches := re.FindAllStringSubmatch(string(content), -1)

	var functionNames []string
	for _, match := range matches {
		functionNames = append(functionNames, match[1])
	}

	return functionNames, nil
}

func TestName(t *testing.T) {
	filename := "D:\\workspace\\workspace\\branch\\gui\\control\\hprdbgctrl\\code\\common\\common.cpp"
	functionNames := mylog.Check2(extractFunctionNamesFromFile(filename))
	fmt.Println("Extracted function names:")
	for _, name := range functionNames {
		fmt.Println(name)
	}
}
