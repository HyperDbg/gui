package ui

import (
	"testing"

	"github.com/ddkwork/golibrary/std/safemap"
	"github.com/ddkwork/golibrary/std/stream"
)

func TestGenTab(t *testing.T) {
	m := safemap.NewOrdered[string, string](func(yield func(string, string) bool) {
		yield("cpu", "cpu")
		yield("PeView", "PeView")
		yield("log", "log")
		yield("notes", "notes")
		yield("breaks", "breaks")
		yield("memory", "memory")
		yield("stack", "stack")
		yield("seh", "seh")
		yield("script", "script")
		yield("symbol", "symbol")
		yield("source", "source")
		yield("References", "References")
		yield("thread", "thread")
		yield("handle", "handle")
		yield("trace", "trace")
		yield("ark", "ark")
	})
	stream.NewGeneratedFile().SetPackageName("ui").EnumTypes("tabPage", m)
}
