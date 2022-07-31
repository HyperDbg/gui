package cpp2go

import (
	"fmt"
	"github.com/ddkwork/librarygo/src/Comment"
	"github.com/ddkwork/librarygo/src/caseconv"
	"github.com/ddkwork/librarygo/src/mycheck"
	"github.com/ddkwork/librarygo/src/mylog"
	"github.com/ddkwork/librarygo/src/stream"
	"github.com/ddkwork/librarygo/src/stream/tool"
	"github.com/goplus/c2go"
	"github.com/goplus/c2go/cl"
	"github.com/goplus/c2go/clang/preprocessor"
	"github.com/goplus/gox"
	"go/types"
	"io/fs"
	"os"
	"path/filepath"
	"runtime/debug"
	"strings"
)

type (
	Interface interface {
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
		ext:        []string{".c", ".cpp", ".inc", ".h", ".c", ".asm"},
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
func (o *object) Dst(s string) *object {
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
	backPath := o.backPath(path)
	for _, e := range o.ext {
		if ext == e || o.expandExt(path) != "" {
			body, err := os.ReadFile(path)
			o.Check(err)
			cppPath := backPath + ".back"
			goPath := backPath + ".go"
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
	if !Comment.New().DeleteKepSpace(o.backPath(o.src)) {
		panic("!Comment.New().DeleteKepSpace(backPath)")
	}
	return o
}
func (o *object) fmtComment(col int) string { return " //col:" + fmt.Sprint(col) }
func (o *object) handleDefineBlock(col int, lines ...string) string {
	fnClean := func(line string) string {
		line = strings.ReplaceAll(line, `\`, ``)
		line = strings.ReplaceAll(line, `#define`, ``)
		line = strings.ReplaceAll(line, `sizeof(UINT32)`, `unsafe.Sizeof(uint32(0))`)
		line = strings.TrimSpace(line)
		return line
	}
	if len(lines) == 1 {
		line := lines[0]
		line = fnClean(line)
		split := strings.Split(line, " ")
		key := split[0]
		value := split[1:]
		return key + " = " + strings.Join(value, " ") + o.fmtComment(col)
	}
	for i, line := range lines {
		line = fnClean(line)
		lines[i] = line
	}
	key := lines[0]
	value := lines[1:]
	return key + " = " + strings.Join(value, " ") + o.fmtComment(col)
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
func (o *object) GetInterfaceName(path string) string {
	base := filepath.Base(path)
	if strings.Contains(base, `~`) {
		base = strings.ReplaceAll(base, `~`, `unknown`)
	}
	if strings.Contains(base, `-`) {
		base = strings.ReplaceAll(base, `-`, `_`)
	}
	if strings.Contains(base, `switch`) {
		base = strings.ReplaceAll(base, `switch`, `switchA`)
	}
	split := strings.Split(base, ".")
	return caseconv.ToCamelUpper(split[0], false)
}
func Run(path, pkg string) {
	abs, err := filepath.Abs(path)
	if !mycheck.Error(err) {
		return
	}
	cl.SetDebug(cl.DbgFlagAll)
	preprocessor.SetDebug(preprocessor.DbgFlagAll)
	gox.SetDebug(gox.DbgFlagInstruction) // | gox.DbgFlagMatch)
	c2go.Run(pkg, abs, c2go.FlagDumpJson, nil)
}
func (o *object) Convert() *object {
	for i, cpp := range o.back {
		if !strings.Contains(cpp.path, "Headers\\Events.h.back") {
			//continue //test
		}
		mylog.Info("current file path", cpp.path)
		lines, ok := tool.File().ToLines(cpp.body)
		if !ok {
			panic("ToLines")
		}
		b := stream.New()
		dir := filepath.Dir(cpp.path)
		pkgName := filepath.Base(dir)
		b.WriteStringLn("package " + pkgName)
		b.WriteStringLn("//" + cpp.path + "\n")

		define := o.GetDefine(lines)
		if define != "" {
			b.WriteStringLn(define)
			//mylog.Json("define ==> const", define)
		}

		enum := o.GetEnum(lines)
		if enum != "" {
			b.WriteStringLn(enum)
			//mylog.Json("enum ==> const", enum)
		}

		//Struct := o.GetStruct(lines)//todo need fix bug
		//if Struct != "" {
		//	b.WriteStringLn(Struct)
		//	//mylog.Json("Struct ==> struct", Struct)
		//}

		method := o.GetMethod(lines, o.GetInterfaceName(cpp.path))
		if method != "" {
			b.WriteStringLn(method)
			//mylog.Json("method ==> func", method)
		}

		//extern BOOLEAN g_IsSerialConnectedToRemoteDebuggee; //todo

		//mylog.Json("cpp ==> go", b.String())
		if !tool.File().WriteTruncate(o.goPath[i].path, b.String()) {
			panic("cpp ==> go error")
		}
		b.Reset()
	}
	return o
}
func (o *object) Check(err error) {
	if err != nil {
		debug.PrintStack()
		panic(err.Error())
	}
}
func (o *object) HandleEnumBlock(col int, lines ...string) string {
	type (
		enumType struct {
			name    string
			Type    string
			value   string
			comment string
		}
	)
	enums := make([]enumType, 0)
	var enum enumType
	for i, blockLine := range lines {
		switch {
		case strings.Contains(blockLine, "typedef enum "):
			enum.Type = strings.TrimPrefix(blockLine, `typedef enum _`) //todo check more
			continue
		case strings.Contains(blockLine, "{"):
			continue
		case blockLine == "":
			continue
		case strings.Contains(blockLine, "}"):
			continue
		}
		blockLine = strings.ReplaceAll(blockLine, ",", "")
		enum.comment = o.fmtComment(col + i)
		enum.name = blockLine
		enums = append(enums, enum)
	}
	tmp := stream.New()
	tmp.WriteStringLn("type " + enums[0].name + " uint32")
	tmp.WriteStringLn("const(")
	for i, enum := range enums {
		if !strings.Contains(enum.name, "=") {
			enum.value = fmt.Sprint(i + 1)
			tmp.WriteStringLn(strings.Join([]string{enum.name, enum.Type, "=", enum.value, enum.comment}, " "))
		} else {
			split := strings.Split(enum.name, "=")
			enum.name = split[0]
			enum.value = split[1]
			tmp.WriteStringLn(strings.Join([]string{enum.name, enum.Type, "=", enum.value, enum.comment}, " "))
		}
	}
	tmp.WriteStringLn(")\n")
	return tmp.String()
}
func (o *object) GetEnum(lines []string) string {
	b := stream.New()
	enum := make([]string, 0)
	for i, line := range lines {
		if strings.Contains(line, `typedef enum`) {
			col := i + 1
			block := lines[i:]
			blockBody := make([]string, 0)
			for _, s := range block {
				blockBody = append(blockBody, s)
				if strings.Contains(s, `}`) {
					enum = append(enum, o.HandleEnumBlock(col, blockBody...))
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
	return b.String()
}
func (o *object) bindGoType(cppType string) (goType string) {
	switch cppType {
	case "UINT64":
		return types.Typ[types.Uint64].Name()
	case "BOOLEAN":
		return types.Typ[types.Bool].Name()

	default:
		return cppType
	}
}
func (o *object) HandleStructBlock(col int, lines ...string) string {
	type (
		structType struct {
			name     string
			elemType string
			elemName string
			comment  string
		}
	)
	fields := make([]structType, 0)
	var Struct structType

	fnGetWord := func(line string) (word, next string) {
		//if strings.Contains(line, "MaxCpuidInputValue") { //debug
		//mylog.Info("line", line)
		//}
		space := strings.TrimSpace(line)
		index := strings.Index(space, " ")
		if index < 0 {
			word = space
			return
		}
		word = space[:index] //
		next = space[index:]
		return
	}
	for i, blockLine := range lines {
		switch {
		case strings.Contains(blockLine, "typedef struct "):
			Struct.name = strings.TrimPrefix(blockLine, `typedef struct _`) //todo check more
			continue
		case strings.Contains(blockLine, "{"):
			continue
		case blockLine == "":
			continue
		case o.isComment(blockLine):
			continue
		case strings.Contains(blockLine, "}"):
			continue
		}
		if strings.Contains(blockLine, ";") {
			word, next := fnGetWord(blockLine)
			if word[len(word)-1] != ';' {
				Struct.elemType = word
				word, next = fnGetWord(next)
				word = strings.ReplaceAll(word, ";", "")
				Struct.elemName = word
			} else { //not ' ' space,it's a nested field
				Struct.elemType = "byte" //? see  TypeOfAction;
				Struct.elemName = word
				word = strings.ReplaceAll(word, ";", "")
			}
		} else {
			//nested type
			blockLine = strings.TrimSpace(blockLine)
			Struct.elemType = blockLine
			camel := caseconv.ToCamelUpper(blockLine, false) //todo test
			camel = strings.TrimSpace(camel)
			camel = strings.TrimSuffix(camel, " ")
			Struct.elemName = camel
		}
		Struct.comment = o.fmtComment(col + i)
		fields = append(fields, Struct)
	}
	for i, field := range fields {
		fields[i].elemType = o.bindGoType(field.elemType)
	}
	if len(fields) == 0 {
		//mylog.Struct(fields)
		return "" //todo check
		panic("len(fields)==0")
	}
	tmp := stream.New()
	tmp.WriteStringLn("type " + fields[0].name + " struct{")
	for _, field := range fields {
		tmp.WriteStringLn(strings.Join([]string{field.elemName, field.elemType}, " "))
	}
	tmp.WriteStringLn("}\n")
	return tmp.String()
}
func (o *object) GetStruct(lines []string) string {
	b := stream.New()
	Struct := make([]string, 0)
	for i, line := range lines {
		if strings.Contains(line, `typedef struct`) {
			col := i + 1
			block := lines[i:]
			blockBody := make([]string, 0)
			for _, s := range block {
				blockBody = append(blockBody, s)
				if strings.Contains(s, `}`) {
					Struct = append(Struct, o.HandleStructBlock(col, blockBody...))
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
	return b.String()
}
func (o *object) isComment(line string) (ok bool) {
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
func (o *object) GetMethod(lines []string, InterfaceName string) string {
	fnIsApi := func(line string) (ok bool) {
		if o.isComment(line) {
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
					if !o.isComment(s) {
						methodBody = append(methodBody, s)
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
	objectName = strings.TrimSpace(objectName)
	objectName = strings.TrimRight(objectName, "")
	b.WriteStringLn("type (\n" + ApiName + " interface{")
	for _, method := range methods {
		b.WriteStringLn(method.api)
	}
	b.WriteStringLn("}\n)\n")
	b.WriteStringLn(`func New` + InterfaceName + `() { return & ` + objectName + `{} }` + "\n")

	ReceiverName := string(objectName[0])
	for _, method := range methods {
		split := strings.Split(method.api, `//`)
		b.WriteStringLn("func (" + ReceiverName + " *" + objectName + ")" + split[0] + "{//" + split[1])
		b.WriteStringLn(method.body)
		b.WriteStringLn("return true")
		b.WriteStringLn("}\n")
	}
	return b.String()
}
func (o *object) Var() *object {
	//TODO implement me
	panic("implement me")
}
func (o *object) Format() *object {
	//TODO implement me
	panic("implement me")
}