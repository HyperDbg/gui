//go:build ignore

// Simple tool that tries to understand simple function definitions as
// used in describing the Windwos Native API.
//
// Comments in Go files that begin with "func:" are parsed as a C
// function prototype.
//
// Comments in Go files that begin with "type:" are parsed as a C
// type definitions.
//
// Comments in Go files that begin with "enum:" are parsed as a C
// enum definitions.
package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/format"
	"go/parser"
	"go/token"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"text/scanner"

	"github.com/ddkwork/golibrary/std/mylog"
)

type State int

const (
	StateInit State = iota
	StateParam
	StateMember
	StateExit
)

type Direction int

const (
	DirectionUnspecified Direction = iota
	DirectionIn
	DirectionInOpt
	DirectionOut
	DirectionOutOpt
	DirectionInOut
	DirectionInOutOpt
	DirectionReserved
)

type FunctionParameterDefinition struct {
	Direction
	Type string
	Name string
}

type FunctionDefinition struct {
	Type   string
	Name   string
	Params []FunctionParameterDefinition
}

type StructMemberDefinition struct {
	Name  string
	Type  string
	Array int
}

type StructDefinition struct {
	Name, OriginalName string
	Members            []StructMemberDefinition
}

var translation = map[string]string{
	"NTSTATUS":       "NtStatus",
	"HANDLE":         "Handle",
	"VOID":           "byte",
	"SIZE_T":         "uintptr",
	"LONGLONG":       "int64",
	"ULONGLONG":      "uint64",
	"ULONG":          "uint32",
	"LONG":           "int32",
	"DWORD":          "uint32",
	"USHORT":         "uint16",
	"SHORT":          "int16",
	"WSTR":           "uint16",
	"WCHAR":          "uint16",
	"UWORD":          "uint16",
	"WORD":           "uint16",
	"UCHAR":          "byte",
	"CCHAR":          "byte",
	"CHAR":           "byte",
	"BYTE":           "byte",
	"BOOLEAN":        "bool",
	"LARGE_INTEGER":  "int64",
	"ULARGE_INTEGER": "uint64",
}

func translate(from string) (to string) {
	var ok bool
	if to, ok = translation[from]; ok {
		return
	}
	// recognize pointer types
	if from[0] == 'P' &&
		!strings.HasPrefix(from, "PUBLIC_") &&
		!strings.HasPrefix(from, "PEB") &&
		!strings.HasPrefix(from, "PS_POST_PROCESS") &&
		!strings.HasPrefix(from, "PROCESS") {
		if rest := translate(from[1:]); rest != "" {
			return "*" + rest
		}
	}
	if strings.HasSuffix(from, "_PTR") {
		if rest := translate(from[:len(from)-4]); rest != "" {
			return "*" + rest
		}
	}
	if strings.Contains(from, "_") || strings.HasSuffix(from, "ID") {
		words := strings.Split(from, "_")
		for _, w := range words {
			to = to + strings.Title(strings.ToLower(w))
		}
		// This avoids collisions between constants and type definitions, e.g.
		// KeyBasicInformation (constant) vs. KEY_BASIC_INFORMATION (struct)
		if strings.HasSuffix(from, "_INFORMATION") ||
			strings.HasSuffix(from, "_TYPE") ||
			strings.HasPrefix(from, "TOKEN_") {
			to += "T"
		}
	} else {
		return strings.Title(strings.ToLower(from))
	}
	return
}

// ParseFunctionDefinition is a half-arsed parser for C function
// definitions as found in MSDN documentation for NTDLL member
// function.
//
// Example (adapted from
// https://msdn.microsoft.com/en-us/library/windows/hardware/ff567029(v=vs.85).aspx):
//
// NTSTATUS NtOpenSection(
//
//	_Out_ PHANDLE            SectionHandle,
//	_In_  ACCESS_MASK        DesiredAccess,
//	_In_  POBJECT_ATTRIBUTES ObjectAttributes
//
// );
func ParseFunctionDefinition(rd io.Reader) (*FunctionDefinition, error) {
	s := &scanner.Scanner{Mode: scanner.GoTokens}
	s.Init(rd)
	state := StateInit
	var f FunctionDefinition
	tokens := make([]string, 0)
	for {
		r := s.Scan()
		if r == scanner.EOF {
			break
		}
		switch state {
		case StateInit:
			switch r {
			case '(':
				if len(tokens) < 2 {
					return nil, fmt.Errorf("function definition needs at least a name and a type")
				}
				f.Name = tokens[0]
				var typ string
				for _, t := range tokens[1:] {
					switch t {
					// ignore WINAPI calling convention
					case "WINAPI":
					default:
						typ = t
					}
				}
				f.Type = translate(typ)
				if f.Type == "" {
					return nil, fmt.Errorf("did not find translation type for %s", typ)
				}
				tokens = tokens[0:0]
				state = StateParam
			case scanner.Ident:
				tokens = append([]string{s.TokenText()}, tokens...)
			default:
				return nil, fmt.Errorf("parse error: got %s", scanner.TokenString(r))
			}
		case StateParam:
			switch r {
			case scanner.Ident:
				tokens = append([]string{s.TokenText()}, tokens...)
			case ',', ')':
				if len(tokens) < 2 {
					return nil, fmt.Errorf("function parameter needs at least a name and a type")
				}

				p := FunctionParameterDefinition{Name: tokens[0]}

				var typ string
				for _, t := range tokens[1:] {
					switch t {
					case "_In_":
						p.Direction = DirectionIn
					case "_In_opt_":
						p.Direction = DirectionInOpt
					case "_Out_":
						p.Direction = DirectionOut
					case "_Out_opt_":
						p.Direction = DirectionOutOpt
					case "_Inout_":
						p.Direction = DirectionInOut
					case "_Inout_opt_":
						p.Direction = DirectionInOutOpt
					case "_Reserved_":
						p.Direction = DirectionReserved
					default:
						typ = t
					}
				}
				p.Type = translate(typ)
				if p.Type == "" {
					return nil, fmt.Errorf("did not find translation type for %s", typ)
				}
				f.Params = append(f.Params, p)
				tokens = tokens[0:0]
				if r == ')' {
					state = StateExit
				}
			}
		case StateExit:
			switch r {
			case ';':
				state = StateInit
				break
			}
		}
	}
	if state != StateInit {
		return nil, fmt.Errorf("parse error: wrong state %d", state)
	}
	return &f, nil
}

// ParseStructDefinition is a half-arsed parser for C struct
// definitions as found in MSDN documentation for NTDLL / Windows
// Driver types.
//
// Example:
//
//	typedef struct _OBJECT_ATTRIBUTES {
//	  ULONG           Length;
//	  HANDLE          RootDirectory;
//	  PUNICODE_STRING ObjectName;
//	  ULONG           Attributes;
//	  PVOID           SecurityDescriptor;
//	  PVOID           SecurityQualityOfService;
//	} OBJECT_ATTRIBUTES, *POBJECT_ATTRIBUTES;
func ParseStructDefinition(rd io.Reader) (*StructDefinition, error) {
	s := &scanner.Scanner{Mode: scanner.GoTokens}
	s.Init(rd)
	state := StateInit
	var sd StructDefinition
	for {
		r := s.Scan()
		if r == scanner.EOF {
			break
		}
		switch state {
		case StateInit:
			e := fmt.Errorf("expecting typedef struct _X {")
			if r != scanner.Ident || s.TokenText() != "typedef" {
				return nil, e
			}
			if r = s.Scan(); r != scanner.Ident || s.TokenText() != "struct" {
				return nil, e
			}
			if r = s.Scan(); r != scanner.Ident || !strings.HasPrefix(s.TokenText(), "_") {
				return nil, e
			}
			sd.OriginalName = s.TokenText()[1:]
			sd.Name = translate(sd.OriginalName)
			if sd.Name == "" {
				return nil, fmt.Errorf("did not find translation type for %s", sd.OriginalName)
			}
			if r = s.Scan(); r != '{' {
				return nil, e
			}
			state = StateMember
		case StateMember:
			if r == '}' {
				state = StateExit
				continue
			}
			if r != scanner.Ident {
				return nil, fmt.Errorf("%s: expecting type, got '%s'", sd.OriginalName, s.TokenText())
			}
			typ := translate(s.TokenText())
			if typ == "" {
				return nil, fmt.Errorf("did not find translation type for %s", s.TokenText())
			}
			m := StructMemberDefinition{Type: typ}
			if r = s.Scan(); r != scanner.Ident {
				return nil, fmt.Errorf("%s: expecting type / name pair, got '%s'", sd.OriginalName, s.TokenText())
			}
			m.Name = s.TokenText()
			if r = s.Scan(); r == '[' {
				if r = s.Scan(); r == scanner.Int {
					n := mylog.Check2(strconv.Atoi(s.TokenText()))
				} else if r == scanner.Ident && s.TokenText() == "ANYSIZE_ARRAY" {
					m.Type = m.Type
					m.Array = 1
				} else {
					return nil, fmt.Errorf("%s: expecting number after '[', got '%s'", sd.OriginalName, s.TokenText())
				}
				if r = s.Scan(); r != ']' {
					return nil, fmt.Errorf("%s: expecting ']', got '%s'", sd.OriginalName, s.TokenText())
				}
				r = s.Scan()
			} else if r != ';' {
				return nil, fmt.Errorf("%s: expecting semicolon after type / name pair, got '%s'", sd.OriginalName, s.TokenText())
			}
			sd.Members = append(sd.Members, m)
		case StateExit:
			if r == ';' {
				state = StateInit
				break
			}
			var typ string
			if r == '*' {
				typ = "*"
				r = s.Scan()
			}
			if r != scanner.Ident {
				return nil, fmt.Errorf("expecting type def, got '%s'", r)
			}
			typ += s.TokenText()
			if strings.HasPrefix(typ, "*P") {
				typ = typ[2:]
			}
			if typ != sd.OriginalName && typ[0] == 'I' {
				typ = typ[1:]
			}
			if typ != sd.OriginalName {
				return nil, fmt.Errorf("expecting type %s, got %s", sd.OriginalName, typ)
			}
			if s.Peek() == ',' {
				s.Scan()
			}
		}
	}
	if state != StateInit {
		return nil, fmt.Errorf("parse error: wrong state %d", state)
	}
	return &sd, nil
}

type EnumMemberDefinition struct {
	Name  string
	Value int
}

type EnumDefinition struct {
	Name, OriginalName string
	Members            []EnumMemberDefinition
}

// ParseEnumDefinition is a half-arsed parser for C enum definitions
// as found in MSDN documentation for NTDLL / Windows Driver types.
//
// Example:
//
//	typedef enum _KEY_VALUE_INFORMATION_CLASS {
//	  KeyValueBasicInformation           = 0,
//	  KeyValueFullInformation,
//	  KeyValuePartialInformation,
//	  KeyValueFullInformationAlign64,
//	  KeyValuePartialInformationAlign64,
//	  MaxKeyValueInfoClass
//	} KEY_VALUE_INFORMATION_CLASS;
func ParseEnumDefinition(rd io.Reader) (*EnumDefinition, error) {
	s := &scanner.Scanner{Mode: scanner.GoTokens}
	s.Init(rd)
	state := StateInit
	var ed EnumDefinition
	lastval := -1
	for {
		r := s.Scan()
		if r == scanner.EOF {
			break
		}
		switch state {
		case StateInit:
			e := fmt.Errorf("expecting typedef enum _X {")
			if r != scanner.Ident || s.TokenText() != "typedef" {
				return nil, e
			}
			if r = s.Scan(); r != scanner.Ident || s.TokenText() != "enum" {
				return nil, e
			}
			if r = s.Scan(); r != scanner.Ident || !strings.HasPrefix(s.TokenText(), "_") {
				return nil, e
			}
			ed.OriginalName = s.TokenText()[1:]
			ed.Name = translate(ed.OriginalName)
			if r = s.Scan(); r != '{' {
				return nil, e
			}
			state = StateMember
		case StateMember:
			if r == '}' {
				state = StateExit
				continue
			}
			if r != scanner.Ident {
				return nil, fmt.Errorf("%s: expecting name, got '%s'", ed.OriginalName, s.TokenText())
			}
			m := EnumMemberDefinition{Name: s.TokenText()}
			if r = s.Scan(); r == '=' {
				if r = s.Scan(); r != scanner.Int {
					return nil, fmt.Errorf("%s: expecting int, got '%s'", ed.OriginalName, s.TokenText())
				}
				v := mylog.Check2(strconv.Atoi(s.TokenText()))

				lastval = v
				m.Value = v
				r = s.Scan()
			} else if r == ',' || r == '}' {
				lastval++
				m.Value = lastval
			} else {
				return nil, fmt.Errorf("%s: expecting enum value or ellipsis, got '%s'", ed.OriginalName, s.TokenText())
			}
			if r == ',' || r == '}' {
				ed.Members = append(ed.Members, m)
			} else {
				return nil, fmt.Errorf("%s: Expecting ',' or '}', got '%s'", ed.OriginalName, s.TokenText())
			}
			if r == '}' {
				state = StateExit
			}
		case StateExit:
			if r == ';' {
				state = StateInit
				break
			}
			var typ string
			if r == '*' {
				typ = "*"
				r = s.Scan()
			}
			if r != scanner.Ident {
				return nil, fmt.Errorf("expecting type def, got '%s'", r)
			}
			typ += s.TokenText()
			if strings.HasPrefix(typ, "*P") {
				typ = typ[2:]
			}
			if typ != ed.OriginalName && typ[1:] == ed.OriginalName {
				typ = typ[1:]
			}
			if typ != ed.OriginalName {
				return nil, fmt.Errorf("expecting type %s, got %s", ed.OriginalName, typ)
			}
			if s.Peek() == ',' {
				s.Scan()
			}
		}
	}
	return &ed, nil
}

func main() {
	var outfile string
	var functions []FunctionDefinition
	var structs []StructDefinition
	var enums []EnumDefinition
	flag.StringVar(&outfile, "output", "", "output file")
	flag.Parse()
	if flag.NArg() != 1 {
		log.Fatal("pass exactly one file")
	}
	infile := filepath.Clean(flag.Arg(0))
	fset := token.NewFileSet()
	if outfile == "" {
		ext := filepath.Ext(infile)
		outfile = infile[:len(infile)-len(ext)] + "_generated" + ext
	}
	f := mylog.Check2(os.Open(infile))

	file := mylog.Check2(parser.ParseFile(fset, "", f, parser.ParseComments))

	for _, cgroup := range file.Comments {
		comment := strings.TrimSpace(cgroup.Text())
		if strings.HasPrefix(comment, "func:") {
			comment = comment[5:]
			f := mylog.Check2(ParseFunctionDefinition(strings.NewReader(comment)))

			functions = append(functions, *f)
		} else if strings.HasPrefix(comment, "type:") {
			comment = comment[5:]
			s := mylog.Check2(ParseStructDefinition(strings.NewReader(comment)))

			structs = append(structs, *s)
		} else if strings.HasPrefix(comment, "enum:") {
			comment = comment[5:]
			e := mylog.Check2(ParseEnumDefinition(strings.NewReader(comment)))

			enums = append(enums, *e)
		}
	}
	f.Close()
	buf := &bytes.Buffer{}
	fmt.Fprintf(buf, `
// This file was autogenerated using go run mkcode.go %s
// DO NOT EDIT.

package %s

`, strings.Join(os.Args[1:], " "), file.Name)

funcs:
	for _, function := range functions {
		for _, param := range function.Params {
			if param.Type[0] == '*' {
				fmt.Fprintln(buf, `import "unsafe"`)
				break funcs
			}
		}
	}
structs:
	for _, st := range structs {
		for _, m := range st.Members {
			if m.Array == 1 && !strings.HasPrefix(m.Name, "Reserved") {
				fmt.Fprintln(buf, `import "reflect"`)
				break structs
			}
		}
	}

	for _, enum := range enums {
		fmt.Fprintf(buf, `
// The %[1]s constants have been derived from the %[2]s enum definition.
type %[1]s uint32; const (
`,
			enum.Name, enum.OriginalName)
		for i, member := range enum.Members {
			if i == 0 {
				fmt.Fprintf(buf, "%s %s = %d\n", member.Name, enum.Name, member.Value)
			} else {
				fmt.Fprintf(buf, "%s = %d\n", member.Name, member.Value)
			}
		}
		fmt.Fprint(buf, ")\n")
	}

	if len(functions) != 0 {
		fmt.Fprintln(buf, "var (")
		for _, function := range functions {
			fmt.Fprintf(buf, `proc%[1]s = modntdll.NewProc("%[1]s")`, function.Name)
			fmt.Fprintln(buf)
		}
		fmt.Fprintln(buf, ")")
	}

	for _, st := range structs {
		fmt.Fprintf(buf, `
// %[1]s has been derived from the %[2]s struct definition.
type %[1]s struct {
`,
			st.Name, st.OriginalName)
		for _, m := range st.Members {
			if m.Array > 0 {
				fmt.Fprintf(buf, "%s [%d]%s\n", m.Name, m.Array, m.Type)
			} else {
				fmt.Fprintf(buf, "%s %s\n", m.Name, m.Type)
			}
		}
		fmt.Fprint(buf, "}\n\n")
		for _, m := range st.Members {
			if m.Array == 1 && !strings.HasPrefix(m.Name, "Reserved") {
				fmt.Fprintf(buf, `
// %[2]sSlice returns a slice over the elements of %[1]s.%[2]s.
//
// Beware: The data is not copied out of %[1]s. The size can usually be taken from an other member of the struct (%[1]s).
func (t *%[1]s) %[2]sSlice(size int) []%[3]s {
	s := []%[3]s{}
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&s))
	hdr.Data = uintptr(unsafe.Pointer(&t.%[2]s[0]))
	hdr.Len = size
	hdr.Cap = size
	return s
}

// Set%[2]sSlice copies s into the memory at %[1]s.%[2]s. 
//
// Beware: No bounds check is performed. Another member of the struct (%[1]s) usually has to be set to the array size.
func (t *%[1]s) Set%[2]sSlice(s []%[3]s) {
	s1 := []%[3]s{}
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&s1))
	hdr.Data = uintptr(unsafe.Pointer(&t.%[2]s[0]))
	hdr.Len = len(s)
	hdr.Cap = len(s)
	copy(s1, s)
}
`, st.Name, m.Name, m.Type)
			}
		}
	}

	for _, function := range functions {
		var plist, alist []string
		var out, inout, opt, reserved, unknown []string
		var comment string
		for _, param := range function.Params {
			switch param.Direction {
			case DirectionIn: // default, don't mention
			case DirectionInOpt:
				opt = append(opt, param.Name)
			case DirectionOut:
				out = append(out, param.Name)
			case DirectionOutOpt:
				out = append(out, param.Name)
				opt = append(opt, param.Name)
			case DirectionInOut:
				inout = append(inout, param.Name)
			case DirectionInOutOpt:
				inout = append(inout, param.Name)
				opt = append(opt, param.Name)
			case DirectionReserved:
				reserved = append(reserved, param.Name)
			default:
				unknown = append(unknown, param.Name)
			}
			alist = append(alist, fmt.Sprintf("%s %s", param.Name, param.Type))
			if param.Type == "bool" {
				plist = append(plist, fmt.Sprintf("fromBool(%s)", param.Name))
			} else if param.Type[0] == '*' {
				plist = append(plist, fmt.Sprintf("uintptr(unsafe.Pointer(%s))", param.Name))
			} else {
				plist = append(plist, fmt.Sprintf("uintptr(%s)", param.Name))
			}
			// alist = append(alist, fmt.Sprintf("p%d %s", n, param.Type))
			// plist = append(plist, "p"+strconv.Itoa(n))
		}
		if len(out) > 0 {
			comment = fmt.Sprintf("// OUT-parameter: %s.\n", strings.Join(out, ", "))
		}
		if len(inout) > 0 {
			comment += fmt.Sprintf("// INOUT-parameter: %s.\n", strings.Join(inout, ", "))
		}
		if len(reserved) > 0 {
			comment += fmt.Sprintf("// RESERVED-parameter: %s.\n", strings.Join(reserved, ", "))
		}
		if len(opt) > 0 {
			comment += fmt.Sprintf("// *OPT-parameter: %s.\n", strings.Join(opt, ", "))
		}
		if len(unknown) > 0 {
			comment += fmt.Sprintf("// unknown-parameter: %s.\n", strings.Join(unknown, ", "))
		}

		returnExpression := function.Type + "(r0)"
		if function.Type == "bool" {
			returnExpression = "r0 != 0"
		}
		fmt.Fprintf(buf, `
%[6]sfunc %[1]s(%[2]s) %[3]s {
	r0, _, _ := proc%[1]s.Call(%[4]s)
	return %[5]s
}

`, function.Name, "\n"+strings.Join(alist, ",\n")+",\n", function.Type, strings.Join(plist, ",\n"), returnExpression, comment)
	}

	f = mylog.Check2(os.Create(outfile))

	data := mylog.Check2(format.Source(buf.Bytes()))

	f.Write(data)
	f.Close()
}
