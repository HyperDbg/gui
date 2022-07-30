package cpp2go

import (
	"fmt"
	"github.com/ddkwork/librarygo/src/caseconv"
	"github.com/ddkwork/librarygo/src/mylog"
	"github.com/ddkwork/librarygo/src/stream"
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
	methodObj struct {
		api  string
		body string
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
func (o *object) GetDefine(lines []string) string {
	b := stream.New()
	b.WriteStringLn("const(")
	define := make([]string, 0)
	for i, line := range lines {
		block := make([]string, 0)
		if strings.Contains(line, `#define`) {
			col := i + 1
			//mylog.Info(fmt.Sprint(col), line)
			if !strings.Contains(line, `\`) {
				define = append(define, o.handleDefineBlock(col, line))
			} else {
				start := lines[i:]
				for _, s := range start {
					block = append(block, s)
					if !strings.Contains(s, `\`) {
						l := o.handleDefineBlock(col, block...)
						//mylog.Trace(fmt.Sprint(col), l)
						define = append(define, l)
						break
					}
				}
			}
		}
	}
	if len(define) == 0 {
		return ""
	}
	for _, s := range define {
		b.WriteStringLn(s)
	}
	b.WriteStringLn(")")
	return b.String()
}

func (o *object) Convert() *object {
	for _, cpp := range o.back {
		if strings.Contains(cpp.path, "bd.cpp.back") {
			mylog.Info("current file path", cpp.path)
			//println(cpp.body)
			lines, ok := tool.File().ToLines(cpp.body)
			if !ok {
				panic("ToLines")
			}
			b := stream.New()
			dir := filepath.Dir(cpp.path)
			pkgName := filepath.Base(dir) //todo rename
			b.WriteStringLn("package " + pkgName)
			b.WriteStringLn("//" + cpp.path)

			//define := o.GetDefine(lines)
			//if define != "" {
			//	b.WriteStringLn(define)
			//	mylog.Json("define ==> const", define)
			//}

			//enum := o.GetEnum(lines)
			//if enum != "" {
			//	b.WriteStringLn(enum)
			//	mylog.Json("enum ==> const", enum)
			//}

			//Struct := o.GetStruct(lines)
			//if Struct != "" {
			//	b.WriteStringLn(Struct)
			//	mylog.Json("Struct ==> struct", Struct)
			//}

			method := o.GetMethod(lines, pkgName)
			if method != "" {
				b.WriteStringLn(method)
				mylog.Json("method ==> func", method)
			}

			//extern BOOLEAN g_IsSerialConnectedToRemoteDebuggee;

			//mylog.Json("cpp ==> go", b.String())
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

func (o *object) GetEnum(lines []string) string {
	b := stream.New()
	b.WriteStringLn("const(") //todo
	enum := make([]string, 0)
	for i, line := range lines {
		if strings.Contains(line, `typedef enum`) {
			//col := i + 1
			block := lines[i:]
			for _, s := range block {
				if s != "" { //todo
					//enum = append(enum, o.handleDefineBlock(col, s))
					enum = append(enum, s)
				}
				if strings.Contains(s, `}`) {
					break
				}
			}
		}
	}
	if len(enum) == 0 {
		return ""
	}
	for _, s := range enum {
		b.WriteStringLn(s)
	}
	b.WriteStringLn(")")
	return b.String()
}

func (o *object) GetStruct(lines []string) string {
	b := stream.New()
	b.WriteStringLn("type xxx struct(")
	Struct := make([]string, 0)
	for i, line := range lines {
		//block := make([]string, 0)
		if strings.Contains(line, `typedef struct`) {
			//col := i + 1
			block := lines[i:]
			for _, s := range block {
				if s != "" { //todo
					//enum = append(enum, o.handleDefineBlock(col, s))
					Struct = append(Struct, s)
				}
				if strings.Contains(s, `}`) {
					break
				}
			}
		}
	}
	if len(Struct) == 0 {
		return ""
	}
	for _, s := range Struct {
		b.WriteStringLn(s)
	}
	b.WriteStringLn(")")
	return b.String()
}

func (o *object) GetMethod(lines []string, InterfaceName string) string {
	isNotes := func(line string) (ok bool) {
		switch {
		case strings.Contains(line, `/*`):
			return true
		case strings.Contains(line, `*/`):
			return true
		case strings.Contains(line, `//`):
			return true
		case strings.Contains(line, `@`):
			return true
		default:
			return
		}
	}
	fnIsApi := func(line string) (ok bool) {
		if isNotes(line) {
			return
		}
		return strings.Contains(line, `(`)
	}
	methods := make([]methodObj, 0)
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		if fnIsApi(line) {
			col := i + 1
			block := lines[i:]
			methodBody := make([]string, 0)
			for j, s := range block {
				if s != "" {
					methodBody = append(methodBody, s)
					if !isNotes(s) {
						if s[0] == '}' {
							api, _, found := strings.Cut(methodBody[0], "(")
							if !found {
								panic("api not found")
							}
							api += "()(ok bool)"
							api += "//col:" + fmt.Sprint(col+j)
							//mylog.Info(fmt.Sprint(col+j), s)
							body := `/*`
							body += strings.Join(methodBody, "\n")
							body += `*/`
							methods = append(methods, methodObj{
								api:  api,
								body: body,
							})
							i += j
							break
						}
					}
				}
			}
		}
	}
	if len(methods) == 0 {
		return ""
	}
	b := stream.New()
	ApiName := caseconv.ToCamelUpper(InterfaceName, false)
	objectName := caseconv.ToCamel(InterfaceName, false)
	b.WriteStringLn("type " + ApiName + " interface(")
	for _, method := range methods {
		b.WriteStringLn(method.api)
	}
	b.WriteStringLn(")")
	ReceiverName := string(objectName[0])
	for _, method := range methods {
		b.WriteStringLn("func (" + ReceiverName + " *" + objectName + ") (ok bool){")
		b.WriteStringLn(method.body)
		b.WriteStringLn("return true")
		b.WriteStringLn("}")
	}
	return b.String()
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
