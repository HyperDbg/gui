package cpp2go

import (
	"fmt"
	"github.com/ddkwork/librarygo/src/Comment"
	"github.com/ddkwork/librarygo/src/caseconv"
	"github.com/ddkwork/librarygo/src/cppBlock"
	"github.com/ddkwork/librarygo/src/mylog"
	"github.com/ddkwork/librarygo/src/stream"
	"github.com/ddkwork/librarygo/src/stream/tool"
	"go/types"
	"io/fs"
	"os"
	"path/filepath"
	"runtime/debug"
	"strconv"
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
		src           string
		dst           string
		col           string
		expandPathExt string
		ext           []string
		back          []pathBody
		goPath        []pathBody
		expandPath    expandPath
	}
)

func New() Interface { return NewObj() }
func NewObj() *object {
	return &object{
		src:           "",
		dst:           "",
		col:           "",
		expandPathExt: "",
		ext:           []string{".c", ".cpp", ".h", ".c", ".asm"}, //".inc", too big
		back:          make([]pathBody, 0),
		goPath:        make([]pathBody, 0),
		expandPath:    expandPath{},
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
	before, after, found := strings.Cut(slash, `/`)
	if !found {
		return before
	}
	return filepath.Join(o.dst, after)
}
func (o *object) backCpp(path string) *object {
	ext := filepath.Ext(path)
	backPath := o.backPath(path)
	for _, e := range o.ext {
		if ext == e || o.expandExt(path) != "" {
			body, err := os.ReadFile(path)
			Check(err)
			cppPath := backPath + ".back"
			goPath := backPath + ".go"
			o.back = append(o.back, pathBody{path: cppPath, body: ""})
			tool.File().WriteTruncate(cppPath, body)
			o.goPath = append(o.goPath, pathBody{path: goPath, body: ""})
		}
	}
	return o
}
func (o *object) Back() *object {
	mylog.Info("src", o.src)
	Check(filepath.Walk(o.src, func(path string, info fs.FileInfo, err error) error {
		o.backCpp(path)
		return err
	}))
	if !Comment.New().DeleteKeepNewLine(o.backPath(o.src)) {
		panic("!Comment.New().DeleteKeepNewLine(backPath)")
	}
	for i, body := range o.back {
		b, err := os.ReadFile(body.path)
		Check(err)
		o.back[i].body = string(b)
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
func (o *object) GetDefine(lines cppBlock.Lines) string {
	isCTL_CODE := false
	for _, line := range lines {
		if line.Line != "" {
			if strings.Contains(line.Line, "CTL_CODE") {
				isCTL_CODE = true
				break
			}
		}
	}
	b := stream.New()
	if isCTL_CODE {
		s := `
func CTL_CODE(deviceType, function, methods, access uint32) uint32 {
	return ((deviceType) << 16) | ((access) << 14) | ((function) << 2) | (methods)
}

const (
	FILE_DEVICE_UNKNOWN = windef.FILE_DEVICE_UNKNOWN
	METHOD_BUFFERED     = windef.METHOD_BUFFERED
	FILE_ANY_ACCESS     = windef.FILE_ANY_ACCESS
)
`
		b.WriteStringLn(s)
		b.WriteStringLn("var(")
	} else {
		b.WriteStringLn("const(")
	}
	define := make([]string, 0)
	for i, line := range lines {
		block := make([]string, 0)
		if strings.Contains(line.Line, `#define`) {
			col := i + 1
			//mylog.Info(fmt.Sprint(col), line)
			if !strings.Contains(line.Line, `\`) {
				define = append(define, o.handleDefineBlock(col, line.Line))
			} else {
				start := lines[i:]
				for _, s := range start {
					block = append(block, s.Line)
					if !strings.Contains(s.Line, `\`) {
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

type (
	tmpInterface interface {
		removeComment(lines []string)
		read() []string
		write(lines []string)
		name() string
		finished()
		clean(b cppBlock.Lines)
		flush()
	}
	tmpObject struct{}
)

func newTmpObject() tmpInterface {
	return &tmpObject{}
}

func (t *tmpObject) clean(b cppBlock.Lines) {
	lines := t.read()
	for _, info := range b {
		for i2 := range lines {
			if info.Col == i2+1 {
				lines[i2] = ""
			}
		}
	}
	t.write(lines)
}

func (t *tmpObject) flush() { Check(os.Remove(t.name())) }
func (t *tmpObject) finished() {
	lines := t.read()
	for _, line := range lines {
		if line != "" {
			panic(t.name() + "not finished")
		}
	}
}
func (t *tmpObject) name() string { return "./tmp.bak" }
func (t *tmpObject) removeComment(lines []string) {
	for i, line := range lines {
		switch {
		case line == "#pragma once":
			lines[i] = ""
		case strings.Contains(line, `/*`) && strings.Contains(line, `*/`):
			start := strings.Index(line, `/*`)
			end := strings.LastIndex(line, `*/`)
			c := line[start-1 : end+2] //?
			dd := strconv.Quote(line)
			dd = dd //8-18 ? "        /* flags  */ 0, \\"
			line = strings.ReplaceAll(line, c, "")
			lines[i] = line //        /* flags  */ 0, \
		case strings.Contains(line, "//"):
			before, _, found := strings.Cut(line, "//")
			if !found {
				panic("// not found")
			}
			lines[i] = before
		}
	}
	t.write(lines)
}
func (t *tmpObject) read() []string {
	file, err := os.ReadFile(t.name())
	Check(err)
	lines, ok := tool.File().ToLines(file)
	if !ok {
		panic(t.name() + " ToLines")
	}
	return lines
}
func (t *tmpObject) write(lines []string) {
	if !tool.File().WriteTruncate(t.name(), stream.New().LinesToString(lines)) {
		panic(t.name())
	}
}

type (
	BlockObject struct {
		//todo class add ?
		externs cppBlock.Lines
		defines cppBlock.Lines
		structs cppBlock.Lines
		enums   cppBlock.Lines
		methods cppBlock.Lines
	}
)

func (o *object) Block(lines []string) (b BlockObject) {
	t := newTmpObject()
	t.removeComment(lines)
	//defer func() { t.finished() }()
	b.defines = cppBlock.FindDefine(t.read())
	t.clean(b.defines)

	b.externs = cppBlock.FindExtern(t.read())
	t.clean(b.externs)

	b.enums = cppBlock.FindEnum(t.read())
	t.clean(b.enums)

	b.structs = cppBlock.FindStruct(t.read())
	t.clean(b.structs)

	b.methods = cppBlock.FindMethod(t.read())
	t.clean(b.methods)
	return
}
func (o *object) Convert() *object {
	for i, cpp := range o.back {
		if !strings.Contains(cpp.path, "phnt\\ntpebteb.h.back") {
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

		block := o.Block(lines)

		define := o.GetDefine(block.defines)
		if define != "" {
			b.WriteStringLn(define)
			//mylog.Json("define ==> const", define)
		}

		enum := o.GetEnum(block.enums)
		if enum != "" {
			b.WriteStringLn(enum)
			//mylog.Json("enum ==> const", enum)
		}

		Struct := o.GetStruct(block.structs)
		if Struct != "" {
			b.WriteStringLn(Struct)
			//mylog.Json("Struct ==> struct", Struct)
		}

		method := o.GetMethod(block.methods, o.GetInterfaceName(cpp.path))
		if method != "" {
			b.WriteStringLn(method)
			//mylog.Json("methods ==> func", methods)
		}

		//mylog.Json("cpp ==> go", b.String())
		//source, err := format.Source(b.Bytes())
		//Check(err)
		if !tool.File().WriteTruncate(o.goPath[i].path, b.String()) {
			panic("cpp ==> go error")
		}
		b.Reset()
	}
	return o
}
func Check(err error) {
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
	//tmp.WriteStringLn("type " + enums[0].name + " uint32")
	tmp.WriteStringLn("const(")
	for i, enum := range enums {
		if !strings.Contains(enum.name, "=") {
			enum.value = fmt.Sprint(i + 1)
			tmp.WriteStringLn(strings.Join([]string{enum.name, "=", enum.value, enum.comment}, " "))
		} else {
			split := strings.Split(enum.name, "=")
			enum.name = split[0]
			enum.value = split[1]
			tmp.WriteStringLn(strings.Join([]string{enum.name, "=", enum.value, enum.comment}, " "))
		}
	}
	tmp.WriteStringLn(")\n")
	return tmp.String()
}
func (o *object) GetEnum(lines cppBlock.Lines) string {
	b := stream.New()
	enum := make([]string, 0)
	for i, line := range lines {
		if strings.Contains(line.Line, `typedef enum`) {
			col := i + 1
			block := lines[i:]
			blockBody := make([]string, 0)
			for _, s := range block {
				blockBody = append(blockBody, s.Line)
				if strings.Contains(s.Line, `}`) {
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
	case "UINT32":
		return types.Typ[types.Uint32].Name()
	case "time_t":
		return "time.Time"
	case "BOOLEAN":
		return types.Typ[types.Bool].Name()
	case "void":
		return "uintptr(0)"
	case "LIST_ENTRY":
		return "*list.List"
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
				word = strings.ReplaceAll(word, ";", "")
				Struct.elemName = word
			}
		} else {
			//nested type
			blockLine = strings.TrimSpace(blockLine)
			if blockLine == "" { //todo bug,need more test
				continue
			}
			Struct.elemType = blockLine
			//mylog.Trace(fmt.Sprint(col), blockLine)
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
func (o *object) GetStruct(lines cppBlock.Lines) string {
	b := stream.New()
	Struct := make([]string, 0)
	for i, line := range lines {
		if strings.Contains(line.Line, `typedef struct`) {
			col := i + 1
			block := lines[i:]
			blockBody := make([]string, 0)
			for _, s := range block {
				blockBody = append(blockBody, s.Line)
				if strings.Contains(s.Line, `}`) {
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
func (o *object) isComment(line string) (ok bool) { //todo delete
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
func (o *object) GetMethod(lines cppBlock.Lines, InterfaceName string) string {
	fnIsApi := func(line string) (ok bool) {
		if o.isComment(line) {
			return
		}
		switch {
		case strings.Contains(line, `#define`):
			return
		case strings.Contains(line, `typedef struct`):
			return
		case strings.Contains(line, `(`):
			return true
		default:
			return
		}
	}
	methods := make([]methodObj, 0)
	for i := 0; i < len(lines); i++ {
		line := lines[i]
		if fnIsApi(line.Line) {
			col := i + 1
			block := lines[i:]
			methodBody := make([]string, 0)
			for j, s := range block {
				if s.Line != "" {
					if !o.isComment(s.Line) {
						methodBody = append(methodBody, s.Line)
						if s.Line[0] == '}' {
							api, _, found := strings.Cut(methodBody[0], "(")
							if !found {
								panic("api not found")
							}
							api = strings.ReplaceAll(api, " ", "_")
							api = strings.ReplaceAll(api, "*", "Ptr")
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
