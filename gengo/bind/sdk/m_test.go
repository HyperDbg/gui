package sdk

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ddkwork/golibrary/mylog"
)

func TestName(t *testing.T) {
	mylog.Call(main)
}

func TestName2(t *testing.T) {
	assert.True(t, strings.Contains("((BUILD_MONTH_IS_OCT || BUILD_MONTH_IS_NOV || BUILD_MONTH_IS_DEC) ? '1' : '0')", "((BUILD_M"))
}
