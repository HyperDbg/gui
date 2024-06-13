package ux

import (
	"cogentcore.org/core/texteditor"
	"github.com/ddkwork/golibrary/mylog"
	"github.com/richardwilkes/unison"
)

func LayoutLog(parent unison.Paneler) unison.Paneler {
	LogEditor := texteditor.NewEditor(parent) // todo 日志语法高亮
	logBuf := texteditor.NewBuf()
	logBuf.SetText([]byte(mylog.Body()))
	LogEditor.SetBuf(logBuf)
}
