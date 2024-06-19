package clang

import (
	"fmt"
	"testing"

	"github.com/ddkwork/golibrary/mylog"
)

func TestRecordLayout_UnmarshalString(t *testing.T) {
	after := "[sizeof=32, align=8]"
	Size := 0
	Align := 0
	mylog.Check2(fmt.Sscanf(after, "[sizeof=%d, align=%d]", &Size, &Align))
}
