package pages

import (
	"fmt"

	"gioui.org/layout"
	"github.com/ddkwork/ux/widget/logview"
)

type LogPage struct {
	logView *logview.LogView
	buffer  []string
}

func NewLog() *LogPage {
	return &LogPage{
		logView: logview.New(),
		buffer:  make([]string, 0),
	}
}

func (p *LogPage) Layout() layout.Widget {
	return p.logView.Layout
}

func (p *LogPage) Println(a ...any) {
	msg := fmt.Sprint(a...)
	p.buffer = append(p.buffer, msg)
}

func (p *LogPage) Printf(format string, a ...any) {
	msg := fmt.Sprintf(format, a...)
	p.buffer = append(p.buffer, msg)
}

func (p *LogPage) Print(a ...any) {
	msg := fmt.Sprint(a...)
	p.buffer = append(p.buffer, msg)
}
