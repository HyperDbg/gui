package api

import (
	"gioui.org/layout"
)

type Interface interface {
	Layout() layout.Widget
	Update() error
	Clear()
	Self() any
}
