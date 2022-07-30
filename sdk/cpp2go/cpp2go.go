package cpp2go

import (
	"fmt"
	"github.com/ddkwork/librarygo/src/mylog"
	"github.com/ddkwork/librarygo/src/stream/tool"
	"io/fs"
	"os"
	"path/filepath"
	"runtime/debug"
	"strings"
)

type (
	Interface interface {

		//Ext() //include txt when status code dir
	}
	pathBody struct {
		path string
		body string
	}
	expandPath struct {
		path string
		ext  string
	}
	object struct {
		src        string
		dst        string
		col        string
		noExtPath  string
		ext        []string
		define     []string
		back       []pathBody
		goPath     []pathBody
		expandPath expandPath
	}
)

func New() Interface { return NewObj() }
func NewObj() *object {
	return &object{
		src:        "",
		dst:        "",
		col:        "",
		noExtPath:  "",
		ext:        make([]string, 0),
		back:       make([]pathBody, 0),
		goPath:     make([]pathBody, 0),
		expandPath: expandPath{},
	}
}

func (o *object) ExpandPath(no, ext string) *object {
	o.expandPath = expandPath{
		path: no,
		ext:  ext,
	}
	return o
}

func (o *object) Src(s string) *object {
	o.src = s
	return o
}

func (o *object) Dsc(s string) *object {
	o.dst = s
	return o
}

func (o *object) expandExt(path string) (e string) {
	if strings.Contains(path, o.expandPath.path) {
		return filepath.Ext(path)
	}
	return ""
}

func (o *object) backPath(path string) string {
	slash := filepath.ToSlash(path)
	_, after, found := strings.Cut(slash, `/`)
	if !found {
		panic("!found")
	}
	return filepath.Join(o.dst, after)
}
func (o *object) backCpp(path string) *object {
	ext := filepath.Ext(path)
	for _, e := range o.ext {
		if ext == e || o.expandExt(path) != "" {
			body, err := os.ReadFile(path)
			o.Check(err)
			cppPath := o.backPath(path) + ".back"
			goPath := o.backPath(path) + ".go"
			o.back = append(o.back, pathBody{
				path: cppPath,
				body: string(body),
			})
			tool.File().WriteTruncate(cppPath, body)
			o.goPath = append(o.goPath, pathBody{
				path: goPath,
				body: "",
			})
		}
	}
	return o
}
func (o *object) Back() *object {
	mylog.Info("src", o.src)
	o.Check(filepath.Walk(o.src, func(path string, info fs.FileInfo, err error) error {
		o.backCpp(path)
		return err
	}))
	return o
}
func (o *object) handleDefineBlock(col int, lines ...string) string {
	fnClean := func(line string) string {
		line = strings.ReplaceAll(line, `\`, ``)
		line = strings.ReplaceAll(line, `#define`, ``)
		line = strings.ReplaceAll(line, `sizeof(UINT32)`, `unsafe.Sizeof(uint32(0))`)
		line = strings.TrimSpace(line)
		return line
	}
	notes := " //col:" + fmt.Sprint(col)
	if len(lines) == 1 {
		line := lines[0]
		line = fnClean(line)
		split := strings.Split(line, " ")
		key := split[0]
		value := split[1:]
		return key + " = " + strings.Join(value, " ") + notes
	}
	for i, line := range lines {
		line = fnClean(line)
		lines[i] = line
	}
	key := lines[0]
	value := lines[1:]
	return key + " = " + strings.Join(value, " ") + notes
}
func (o *object) Convert() *object {
	for _, cpp := range o.back {
		if strings.Contains(cpp.path, "Constants.h.back") {
			mylog.Info("current file path", cpp.path)
			//println(cpp.body)
			lines, ok := tool.File().ToLines(cpp.body)
			if !ok {
				panic("ToLines")
			}
			for i, line := range lines {
				block := make([]string, 0)
				if strings.Contains(line, `#define`) {
					col := i + 1
					//mylog.Info(fmt.Sprint(col), line)
					if !strings.Contains(line, `\`) {
						o.define = append(o.define, o.handleDefineBlock(col, line))
					} else {
						start := lines[i:]
						for _, s := range start {
							block = append(block, s)
							if !strings.Contains(s, `\`) {
								l := o.handleDefineBlock(col, block...)
								//mylog.Trace(fmt.Sprint(col), l)
								o.define = append(o.define, l)
								break
							}
						}
					}
				}
			}
			for _, s := range o.define {
				mylog.Trace("", s)
			}
		}
	}
	return o
}
func (o *object) Check(err error) {
	if err != nil {
		debug.PrintStack()
		panic(err.Error())
	}
}

func (o *object) Enum() *object {
	//TODO implement me
	panic("implement me")
}

func (o *object) Struct() *object {
	//TODO implement me
	panic("implement me")
}

func (o *object) Var() *object {
	//TODO implement me
	panic("implement me")
}

func (o *object) Method() *object {
	//TODO implement me
	panic("implement me")
}

func (o *object) Type() *object {
	//TODO implement me
	panic("implement me")
}

func (o *object) Format() *object {
	//TODO implement me
	panic("implement me")
}

func (o *object) Ext(e []string) *object {
	o.ext = e
	return o
}
