package ux

import (
	"cogentcore.org/core/gi"
	"cogentcore.org/core/texteditor"
	"github.com/ddkwork/golibrary/mylog"
)

func pageLog(parent *gi.Frame) {
	LogEditor := texteditor.NewEditor(parent) // todo 日志语法高亮
	logBuf := texteditor.NewBuf()
	logBuf.SetText([]byte(mylog.Body()))
	LogEditor.SetBuf(logBuf)
}
