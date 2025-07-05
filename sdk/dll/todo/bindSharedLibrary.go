package main

import (
	"flag"
	"log"
	"os"
	"strings"
	"text/template"

	"modernc.org/cc/v4"
)

type param struct {
	Type string
	Ptr  string
	Name string
}

type funcDecl struct {
	Class      string
	Ret        string
	Ptr        string
	Name       string
	Parameters []param
}

func newFunc(d *cc.Declaration) *funcDecl {
	if d.Case != cc.DeclarationDecl {
		return nil
	}
	var fnClass string
	ds := d.DeclarationSpecifiers
	if ds.Case == cc.DeclarationSpecifiersStorage {
		fnClass = ds.StorageClassSpecifier.Token.SrcStr()
		ds = ds.DeclarationSpecifiers
	}
	if ds == nil || ds.Case != cc.DeclarationSpecifiersTypeSpec {
		return nil
	}
	retType := ds.TypeSpecifier.Token.SrcStr()

	if d.InitDeclaratorList == nil {
		return nil
	}
	id := d.InitDeclaratorList.InitDeclarator
	if id.Case != cc.InitDeclaratorDecl {
		return nil
	}
	decl := id.Declarator
	var retPtr string
	if decl.Pointer != nil {
		retPtr = decl.Pointer.Token.SrcStr()
	}
	dd := decl.DirectDeclarator
	if dd == nil || dd.Case != cc.DirectDeclaratorFuncParam {
		return nil
	}
	dd2 := dd.DirectDeclarator
	if dd2 == nil || dd2.Case != cc.DirectDeclaratorIdent {
		return nil
	}
	fnName := dd2.Token.SrcStr()

	var params []param
	if ptl := dd.ParameterTypeList; ptl != nil {
		if ptl.Case != cc.ParameterTypeListList {
			return nil
		}
		for t := ptl.ParameterList; t != nil; t = t.ParameterList {
			pdecl := t.ParameterDeclaration
			if pdecl.Case != cc.ParameterDeclarationDecl {
				return nil
			}
			ds := pdecl.DeclarationSpecifiers
			if ds == nil || ds.Case != cc.DeclarationSpecifiersTypeSpec {
				return nil
			}
			paramType := ds.TypeSpecifier.Token.SrcStr()
			d := pdecl.Declarator
			if d == nil {
				return nil
			}
			dd := d.DirectDeclarator
			if dd == nil || dd.Case != cc.DirectDeclaratorIdent {
				return nil
			}
			paramName := dd.Token.SrcStr()
			var paramPtr string
			if d.Pointer != nil {
				paramPtr = d.Pointer.Token.SrcStr()
			}
			params = append(params, param{
				Type: paramType,
				Ptr:  paramPtr,
				Name: paramName,
			})
		}
	}

	return &funcDecl{
		Class:      fnClass,
		Ret:        retType,
		Ptr:        retPtr,
		Name:       fnName,
		Parameters: params,
	}
}

type tplArg struct {
	Pkg   string
	Hdr   string
	CSrc  string
	Funcs []*funcDecl
}

var cTplSrc = `// GENERATED, DO NOT EDIT.

{{with .Hdr}}#include "{{.}}"{{- end}}

{{range .Funcs}}
{{.Ret}} {{.Ptr}}_{{.Name}}(
	CK_FUNCTION_LIST_PTR _funcs,
	{{- range $index, $el := .Parameters}}
	{{- if $index}},{{end}}
	{{$el.Type}} {{$el.Ptr}}{{$el.Name}}
	{{- end}}
) {
	return (*_funcs->{{.Name}})(
		{{- range $index, $el := .Parameters}}
		{{- if $index}},{{end}}
		{{$el.Name}}
		{{- end}}
	);
}
{{end}}
`

var goTplSrc = `package {{.Pkg}}

// GENERATED, DO NOT EDIT.

/*
#include "{{.Hdr}}"

{{range .Funcs}}
{{.Ret}} {{.Ptr}}_{{.Name}}(
	CK_FUNCTION_LIST_PTR _funcs,
	{{- range $index, $el := .Parameters}}
	{{- if $index}},{{end}}
	{{$el.Type}} {{$el.Ptr}}{{$el.Name}}
	{{- end}}
) {
	return (*_funcs->{{.Name}})(
		{{- range $index, $el := .Parameters}}
		{{- if $index}},{{end}}
		{{$el.Name}}
		{{- end}}
	);
}
{{end}}
*/
import "C"

type functionTable struct {
	t *C.CK_FUNCTION_LIST
}

{{range .Funcs}}
func (f *functionTable) {{mute_keywords .Name}}(
	{{- range .Parameters}}
	{{mute_keywords .Name}} {{if .Ptr}}*{{end}}C.{{.Type}},
	{{- end}}
) error {
	rv := C._{{.Name}}(
		f.t,
		{{- range .Parameters}}
		{{mute_keywords .Name}},
		{{- end}}
	);
	if rv != C.CKR_OK {
		return &Error{fnName: "{{.Name}}", code: rv}
	}
	return nil
}
{{end}}
`

var keywords = map[string]struct{}{
	"break":       {},
	"case":        {},
	"chan":        {},
	"const":       {},
	"continue":    {},
	"default":     {},
	"defer":       {},
	"else":        {},
	"fallthrough": {},
	"for":         {},
	"func":        {},
	"go":          {},
	"goto":        {},
	"if":          {},
	"import":      {},
	"interface":   {},
	"map":         {},
	"package":     {},
	"range":       {},
	"return":      {},
	"select":      {},
	"struct":      {},
	"switch":      {},
	"type":        {},
	"var":         {},
}

var (
	cTpl  = template.Must(template.New("c_template").Parse(cTplSrc))
	goTpl = template.Must(template.New("go_template").Funcs(template.FuncMap{
		"mute_keywords": func(src string) string {
			if _, ok := keywords[src]; ok {
				return "_" + src
			}
			return src
		},
	}).Parse(goTplSrc))
)

func main() {
	var (
		hdr   string
		cHdr  string
		cOut  string
		goOut string
		pkg   string
	)

	flag.StringVar(&hdr, "i", "", "Input header")
	flag.StringVar(&cHdr, "h", "", "C header")
	flag.StringVar(&cOut, "c", "", "C output")
	flag.StringVar(&goOut, "g", "", "Go output")
	flag.StringVar(&pkg, "p", "", "Go package")
	flag.Parse()

	ccCfg, _ := cc.NewConfig("", "")
	src := []cc.Source{
		{Name: "<predefined>", Value: ccCfg.Predefined},
		{Name: "<builtin>", Value: cc.Builtin},
		{Name: hdr},
	}
	ast, err := cc.Parse(ccCfg, src)
	if err != nil {
		log.Fatal(err)
	}

	var funcs []*funcDecl
	for tu := ast.TranslationUnit; tu != nil; tu = tu.TranslationUnit {
		if ed := tu.ExternalDeclaration; ed.Case == cc.ExternalDeclarationDecl {
			if fn := newFunc(ed.Declaration); fn != nil {
				if fn.Class == "extern" && strings.HasPrefix(fn.Name, "C_") && fn.Ret == "CK_RV" {
					funcs = append(funcs, fn)
				}
			}
		}
	}
	args := &tplArg{
		Hdr:   cHdr,
		Pkg:   pkg,
		CSrc:  cOut,
		Funcs: funcs,
	}

	if cOut != "" {
		fd, err := os.Create(cOut)
		if err != nil {
			log.Fatal(err)
		}
		defer fd.Close()

		if err := cTpl.Execute(fd, &args); err != nil {
			log.Fatal(err)
		}
		return
	}

	if goOut != "" {
		fd, err := os.Create(goOut)
		if err != nil {
			log.Fatal(err)
		}
		defer fd.Close()

		if err := goTpl.Execute(fd, &args); err != nil {
			log.Fatal(err)
		}
	}
}
