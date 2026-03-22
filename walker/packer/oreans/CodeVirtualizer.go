//go:build windows

package oreans

import (
	"fmt"

	"github.com/oreans/virtualizersdk"
)

func main_() {
	string1 := "This is string1"
	virtualizersdk.Macro(virtualizersdk.TIGER_BLACK_START)
	string2 := "This is string2"
	string3 := string1 + string2
	fmt.Println(string3)
	// resp := mylog.Check2(http.GetMust("http://google.com"))

	virtualizersdk.Macro(virtualizersdk.TIGER_BLACK_END)
	string4 := string2 + "...and string 4"
	fmt.Println(string4)
}
