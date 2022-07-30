package cpp2go

import (
	"github.com/ddkwork/librarygo/src/mylog"
	"github.com/ddkwork/librarygo/src/stream/tool"
	"io/fs"
	"os"
	"path/filepath"
	"runtime/debug"
	"strings"
)

/*
c++   自己解析方法
h ast 批量
c c2go
mod重写和单元测试
*/
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
func (o *object) Convert() *object {
	for _, cpp := range o.back {
		if strings.Contains(cpp.path, "Ioctls.h.back") {
			println(cpp.path)
			println(cpp.body)
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

func (o *object) Define() *object {
	//TODO implement me
	panic("implement me")
}

func (o *object) Strcut() *object {
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
