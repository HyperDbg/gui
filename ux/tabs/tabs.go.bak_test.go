package tabs

import (
	"github.com/ddkwork/golibrary/mylog"
	"github.com/ddkwork/golibrary/pather"
	"github.com/ddkwork/golibrary/src/caseconv"
	"github.com/ddkwork/golibrary/src/stream"
	"github.com/ddkwork/golibrary/src/stream/tool"
	"go/format"
	"io/fs"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
)

func TestGen(t *testing.T) {
	mapPather := make(map[string]*pather.Object, 0)
	filepath.Walk("pageIco", func(path string, info fs.FileInfo, err error) error {
		ext := filepath.Ext(path)
		if ext == ".png" || ext == ".ico" {
			base := filepath.Base(path)
			e := filepath.Ext(base)
			base = base[:len(base)-len(e)]
			lower := caseconv.ToCamelToLower(base, false)
			upper := caseconv.ToCamelUpper(lower, false)
			path = strings.ReplaceAll(path, `\`, `/`)
			mylog.Warning(lower, path)
			mapPather[path] = pather.New(pather.Input{
				LowerFileName: lower,
				UpperFileName: upper,
				EmbedPrefix:   "png",
				ObjectSuffix:  "Button ",
				Path:          path,
			})
		}
		return err
	})
	GenKind(mapPather)
	GenToolbar(mapPather)
}
func GenKind(in map[string]*pather.Object) {
	s := stream.NewString(`
//Code generated from mapPather - DO NOT EDIT.

package tabs

import (
	_ "embed"
	"github.com/ddkwork/golibrary/skiaLib/widget"
	"github.com/richardwilkes/unison"
)

type kind byte

var name kind = 0

const (
`)

	s.WriteStringLn("invalidKind kind = iota")
	for _, v := range in {
		s.WriteStringLn(v.OutPut.Kind)
	}
	s.WriteStringLn(")")

	//gen fmt.Stringer()
	s.WriteStringLn(`
func (k kind) String() string {
	switch k {`)
	for _, v := range in {
		s.WriteStringLn("case " + v.OutPut.Kind + ":")
		s.WriteStringLn("return " + strconv.Quote(v.OutPut.Kind))
	}
	s.WriteStringLn(`
	}
	return "invalid bmp kind"
}
`)

	//gen *unison.Image
	s.WriteStringLn(`
func (k kind) Image() *unison.Image {
	switch k {`)
	for _, v := range in {
		s.WriteStringLn("case " + v.OutPut.Kind + ":")
		s.WriteStringLn("return " + "widget.MustImage(" + v.OutPut.Embed + ")")
	}
	s.WriteStringLn(`
	}
	return nil
}
`)

	//gen embed buffers
	s.WriteStringLn("var (")
	for k, v := range in {
		s.WriteStringLn("//go:embed " + k)
		s.WriteStringLn(v.OutPut.Embed + " []byte")
		s.NewLine()
	}
	s.WriteStringLn(")")
	source, err := format.Source(s.Bytes())
	if !mylog.Error(err) {
		tool.File().WriteTruncate("tmp/kind.go", s.String())
		return
	}
	tool.File().WriteTruncate("tmp/kind.go", source)
	//mylog.Json("", string(source))
}

func GenToolbar(in map[string]*pather.Object) {
	s := stream.NewString(`
//Code generated from mapPather - DO NOT EDIT.

package tabs

import (
	_ "embed"
	"github.com/ddkwork/golibrary/skiaLib/widget"
	"github.com/ddkwork/golibrary/skiaLib/widget/canvasobjectapi"
	"github.com/richardwilkes/unison"
)

type (
	Interface interface {
		canvasobjectapi.Interface
		Actions
	}
	Actions interface { 
`)
	for _, v := range in {
		s.WriteStringLn(v.OutPut.Method + "() *unison.Button")
	}
	s.WriteStringLn("}")

	s.WriteStringLn("object struct {")
	for _, v := range in {
		s.WriteStringLn(v.OutPut.Object + "*unison.Button")
	}

	s.WriteStringLn("})")
	s.WriteStringLn("func New() Interface { return &object{} }")

	//gen CanvasObject
	s.WriteStringLn(`
func (o *object) CanvasObject(window fyne.Window) fyne.CanvasObject {
	panel := unison.NewPanel()
	panel.SetLayout(&unison.FlowLayout{
		HSpacing: unison.StdHSpacing,
		VSpacing: unison.StdVSpacing,
	})`)
	for _, v := range in {
		//o.helpButton = widget.CreateImageButton(helpKind.Image(), helpKind.String(), panel)
		s.WriteString("o.")
		s.WriteString(v.OutPut.Object)
		s.WriteString("=")
		s.WriteString("widget.CreateImageButton(")
		s.WriteString(v.OutPut.Kind)
		s.WriteString(".Image(), ")
		s.WriteString(v.OutPut.Kind)
		s.WriteStringLn(".String(), panel)")
	}
	s.WriteStringLn(`return panel
	}
	`)

	//gen methods
	for _, v := range in {
		s.WriteString("func (o *object) ")
		s.WriteString(v.OutPut.Method)
		s.WriteString("() *unison.Button      { return o.")
		s.WriteString(v.OutPut.Object)
		s.WriteStringLn("}")
	}

	source, err := format.Source(s.Bytes())
	if !mylog.Error(err) {
		tool.File().WriteTruncate("tmp/tabs.go", s.String())
		return
	}
	tool.File().WriteTruncate("tmp/tabs.go", source)
	//mylog.Json("", string(source))
}
