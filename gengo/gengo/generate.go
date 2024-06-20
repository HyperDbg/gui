package gengo

import (
	"fmt"
	"github.com/ddkwork/golibrary/mylog"
	"go/token"
	"strconv"
	"strings"

	"github.com/can1357/gengo/clang"
	"github.com/dave/dst"
)

type deferred []func()

func (d *deferred) Go() {
	list := *d
	*d = nil
	for _, f := range list {
		f()
	}
}

func (d *deferred) Do(f func()) {
	*d = append(*d, f)
}

type initTypeValidation struct {
	ty     TypeRef
	size   int
	align  int
	fields []dst.Expr
}

func copyDecoration(d *dst.NodeDecs, node clang.Node) {
	comments := strings.TrimSpace(node.Comment())
	if len(comments) > 0 {
		for _, c := range strings.Split(comments, "\n") {
			d.Start.Append("// " + c)
		}
		d.Before = dst.NewLine
		d.After = dst.NewLine
	}
}

func (mod Module) EmitEnum(n *clang.EnumDecl) {
	// Add the typedef.
	var ty dst.Expr
	for _, c := range n.Inner {
		if c, ok := c.(*clang.EnumConstantDecl); ok {
			ty = mod.Parent.ConvertQualType(c.Type.QualType)
			break
		}
	}
	if ty == nil {
		ty = BuiltinInt.Ref()
	}
	tydef := mod.AddType(TcEnum, n.Name, ty)
	tydef.Rename(mod.Parent.NameType(n.Name))
	copyDecoration(&tydef.Decl.Decs.NodeDecs, n)

	// Add the enum constants.
	spec := []dst.Spec{}
	val := -1
	for _, c := range n.Inner {
		if c, ok := c.(*clang.EnumConstantDecl); ok {
			val += 1
			if ival := clang.First[clang.ValueNode](c); ival != nil {
				if cval, ok := ival.(clang.ConstValueNode); ok {
					val, _ = strconv.Atoi(cval.Value())
				} else {
					mylog.CheckIgnore(fmt.Sprintf("unhandled value node: %T", ival))
					continue
				}
			}

			var dec dst.ValueSpecDecorations
			copyDecoration(&dec.NodeDecs, c)

			spec = append(spec, &dst.ValueSpec{
				Names: []*dst.Ident{dst.NewIdent(mod.Parent.NameValue(c.Name))},
				Type:  tydef.Ref(),
				Values: []dst.Expr{
					&dst.BasicLit{
						Kind:  token.INT,
						Value: strconv.Itoa(val),
					},
				},
				Decs: dec,
			})
		}
	}
	mod.Decls = append(mod.Decls, &dst.GenDecl{
		Tok:   token.CONST,
		Specs: spec,
	})
}

func (mod Module) TryEmitNaturalStruct(n *clang.RecordDecl, layout *clang.RecordLayout, ty TypeRef, fl *dst.FieldList, vld *[]initTypeValidation) bool {
	// Can't define a natural type if it is a union.
	if n.TagUsed == "union" {
		return false
	}
	if mod.Parent.ForceSynthetic(n.Name) {
		return false
	}

	v := initTypeValidation{
		ty:     ty,
		size:   layout.Size,
		align:  layout.Align,
		fields: []dst.Expr{},
	}

	// Add the fields.
	for _, c := range layout.Fields {
		var field *dst.Field
		if c.Name == "" {
			field = &dst.Field{
				Names: []*dst.Ident{},
				Type:  mod.Parent.ConvertQualType(c.Type),
			}
		} else {
			name := mod.Parent.NameField(c.Name, n.Name)
			field = &dst.Field{
				Names: []*dst.Ident{dst.NewIdent(name)},
				Type:  mod.Parent.ConvertQualType(c.Type),
			}
			v.fields = append(v.fields,
				&dst.BasicLit{
					Kind:  token.STRING,
					Value: strconv.Quote(name),
				},
				&dst.BasicLit{
					Kind:  token.INT,
					Value: fmt.Sprintf("0x%x", c.Offset),
				},
			)
		}
		for _, f := range n.Children() {
			if f, ok := f.(*clang.FieldDecl); ok {
				if f.Name == c.Name {
					copyDecoration(&field.Decs.NodeDecs, f)
					break
				}
			}
		}
		fl.List = append(fl.List, field)
	}

	*vld = append(*vld, v)
	return true
}

func (mod Module) EmitSyntheticStruct(n *clang.RecordDecl, layout *clang.RecordLayout, ty TypeRef, fl *dst.FieldList, vld *[]initTypeValidation) {
	// Add byte[size] field.
	unitType := BuiltinByte
	unitCount := layout.Size
	if layout.Align >= 8 {
		unitType = BuiltinInt64
		unitCount /= 8
	} else if layout.Align >= 4 {
		unitType = BuiltinInt32
		unitCount /= 4
	} else if layout.Align >= 2 {
		unitType = BuiltinInt16
		unitCount /= 2
	}
	fl.List = append(fl.List, &dst.Field{
		Names: []*dst.Ident{dst.NewIdent("Raw")},
		Type: &dst.ArrayType{
			Len: &dst.BasicLit{
				Kind:  token.INT,
				Value: strconv.Itoa(unitCount),
			},
			Elt: unitType.Ref(),
		},
	})
	*vld = append(*vld, initTypeValidation{
		ty:    ty,
		size:  layout.Size,
		align: layout.Align,
	})

	sliceOffset := func(offset int) dst.Expr {
		return &dst.CallExpr{
			Fun: &dst.SelectorExpr{
				X:   &dst.Ident{Name: "unsafe"},
				Sel: &dst.Ident{Name: "Add"},
			},
			Args: []dst.Expr{
				&dst.CallExpr{
					Fun: &dst.SelectorExpr{
						X:   &dst.Ident{Name: "unsafe"},
						Sel: &dst.Ident{Name: "Pointer"},
					},
					Args: []dst.Expr{
						&dst.CallExpr{
							Fun: &dst.SelectorExpr{
								X:   &dst.Ident{Name: "unsafe"},
								Sel: &dst.Ident{Name: "SliceData"},
							},
							Args: []dst.Expr{
								&dst.SliceExpr{
									X: &dst.SelectorExpr{
										X:   dst.NewIdent("s"),
										Sel: dst.NewIdent("Raw"),
									},
								},
							},
						},
					},
				},
				&dst.BasicLit{
					Kind:  token.INT,
					Value: strconv.Itoa(offset),
				},
			},
		}
	}

	// Add accessors for each field.
	for _, c := range layout.Fields {
		name := mod.Parent.NameField(c.Name, n.Name)

		var funcDec dst.FuncDeclDecorations
		for _, f := range n.Children() {
			if f, ok := f.(*clang.FieldDecl); ok {
				if f.Name == c.Name {
					copyDecoration(&funcDec.NodeDecs, f)
					break
				}
			}
		}

		mod.Decls = append(mod.Decls, &dst.FuncDecl{
			Name: dst.NewIdent(mod.Parent.NameGetter(name)),
			Decs: funcDec,
			Recv: &dst.FieldList{
				List: []*dst.Field{
					{
						Names: []*dst.Ident{dst.NewIdent("s")},
						Type:  ty.Ref(),
					},
				},
			},
			Type: &dst.FuncType{
				Params: &dst.FieldList{},
				Results: &dst.FieldList{
					List: []*dst.Field{
						{
							Names: []*dst.Ident{},
							Type:  mod.Parent.ConvertQualType(c.Type),
						},
					},
				},
			},
			Body: &dst.BlockStmt{
				List: []dst.Stmt{
					&dst.ReturnStmt{
						Results: []dst.Expr{
							&dst.CallExpr{
								Fun: &dst.IndexExpr{
									X: &dst.SelectorExpr{
										X:   dst.NewIdent("gengort"),
										Sel: dst.NewIdent("ReadBitcast"),
									},
									Index: mod.Parent.ConvertQualType(c.Type),
								},
								Args: []dst.Expr{
									sliceOffset(c.Offset),
								},
							},
						},
					},
				},
			},
		})
		mod.Decls = append(mod.Decls, &dst.FuncDecl{
			Name: dst.NewIdent(mod.Parent.NameSetter(name)),
			Decs: funcDec,
			Recv: &dst.FieldList{
				List: []*dst.Field{
					{
						Names: []*dst.Ident{dst.NewIdent("s")},
						Type:  &dst.StarExpr{X: ty.Ref()},
					},
				},
			},
			Type: &dst.FuncType{
				Params: &dst.FieldList{
					List: []*dst.Field{
						{
							Names: []*dst.Ident{dst.NewIdent("v")},
							Type:  mod.Parent.ConvertQualType(c.Type),
						},
					},
				},
				Results: &dst.FieldList{},
			},
			Body: &dst.BlockStmt{
				List: []dst.Stmt{
					&dst.ExprStmt{
						X: &dst.CallExpr{
							Fun: &dst.SelectorExpr{
								X:   dst.NewIdent("gengort"),
								Sel: dst.NewIdent("WriteBitcast"),
							},
							Args: []dst.Expr{
								sliceOffset(c.Offset),
								dst.NewIdent("v"),
							},
						},
					},
				},
			},
		})
	}
}

func (mod Module) EmitStruct(n *clang.RecordDecl, layouts *clang.LayoutMap, deferred *deferred, vld *[]initTypeValidation) {
	var layout *clang.RecordLayout
	var defName string
	var qualDefName string
	if n.Name == "" {
		line := fmt.Sprintf(":%d:%d", n.Loc.Line, n.Loc.Col)
		for ty, rec := range layouts.Map {
			if strings.Contains(ty, line) {
				layout = rec
				break
			}
		}
		if layout == nil {
			line = fmt.Sprintf(":%d:", n.Loc.Line)
			for ty, rec := range layouts.Map {
				if strings.Contains(ty, line) {
					layout = rec
					break
				}
			}
		}
		defName = fmt.Sprintf("Anon%d_%d", n.Loc.Line, n.Loc.Col)
		qualDefName = n.TagUsed + " " + defName
	} else {
		defName = n.Name
		qualDefName = n.TagUsed + " " + defName
		layout = layouts.Map[qualDefName]
	}

	if layout == nil {
		fmt.Println("[WARN] No layout for record:", n.Name)
		return
	}

	// Add the struct.
	tc := TcStruct
	if n.TagUsed == "union" {
		tc = TcUnion
	}
	fieldList := new(dst.FieldList)
	str := mod.AddType(tc, defName, &dst.StructType{
		Fields: fieldList,
	})
	str.Rename(mod.Parent.NameType(defName))
	if qualDefName != layout.Type {
		mod.Parent.RemapType(layout.Type, str)
	}

	deferred.Do(func() {
		if !mod.TryEmitNaturalStruct(n, layout, str, fieldList, vld) {
			mod.EmitSyntheticStruct(n, layout, str, fieldList, vld)
		}
	})
}

const PRELOAD = true

func (mod Module) EmitFunction(n *clang.FunctionDecl) {
	// Add the import.
	// - var __imp_func = Import("func")
	//
	name := mod.Parent.NameFunc(n.Name)
	importName := "__imp_" + n.Name

	if PRELOAD {
		mod.Decls = append(mod.Decls, &dst.GenDecl{
			Tok: token.VAR,
			Specs: []dst.Spec{
				&dst.ValueSpec{
					Names: []*dst.Ident{dst.NewIdent(importName)},
					Type: &dst.SelectorExpr{
						X:   dst.NewIdent("gengort"),
						Sel: dst.NewIdent("PreloadProc"),
					},
				},
			},
		})
		mod.OnInit(&dst.AssignStmt{
			Lhs: []dst.Expr{
				dst.NewIdent(importName),
			},
			Tok: token.ASSIGN,
			Rhs: []dst.Expr{
				&dst.CallExpr{
					Fun: &dst.SelectorExpr{
						X:   dst.NewIdent("GengoLibrary"),
						Sel: dst.NewIdent("ImportNow"),
					},
					Args: []dst.Expr{
						&dst.BasicLit{
							Kind:  token.STRING,
							Value: strconv.Quote(n.MangledName),
						},
					},
				},
			},
		})
	} else {
		mod.Decls = append(mod.Decls, &dst.GenDecl{
			Tok: token.VAR,
			Specs: []dst.Spec{
				&dst.ValueSpec{
					Names: []*dst.Ident{dst.NewIdent(importName)},
					Values: []dst.Expr{
						&dst.CallExpr{
							Fun: &dst.SelectorExpr{
								X:   dst.NewIdent("GengoLibrary"),
								Sel: dst.NewIdent("Import"),
							},
							Args: []dst.Expr{
								&dst.BasicLit{
									Kind:  token.STRING,
									Value: strconv.Quote(n.MangledName),
								},
							},
						},
					},
				},
			},
		})
	}

	// Get the result type.
	var result, result2 dst.Expr
	{
		before, _, ok := strings.Cut(n.Type.QualType, "(")
		if !ok {
			panic("unhandled function type")
		}
		before = strings.TrimSpace(before)
		result = mod.Parent.ConvertQualType(before)
		if ident, ok := result.(*dst.Ident); ok {
			if ident.Name == "void" {
				result = nil
			}
		}
		if result != nil {
			result2 = mod.Parent.ConvertQualType(before)
		}
	}

	// Add the function.
	// - func Func(arg *T) {
	// - 	__imp_func.Call1(
	// - 		marshalsyscall(arg),
	// - 	)
	// - }
	typ := &dst.FuncType{
		Params:  &dst.FieldList{},
		Results: &dst.FieldList{},
	}
	decl := &dst.FuncDecl{
		Name: dst.NewIdent(name),
		Type: typ,
	}
	copyDecoration(&decl.Decs.NodeDecs, n)
	mod.Decls = append(mod.Decls, decl)
	callExpr := &dst.CallExpr{
		Args: []dst.Expr{
			&dst.CallExpr{
				Fun: &dst.SelectorExpr{
					X:   dst.NewIdent(importName),
					Sel: dst.NewIdent("Addr"),
				},
			},
		},
	}

	// Add the parameters.
	paramNodes := clang.All[*clang.ParmVarDecl](n)
	for _, p := range paramNodes {
		arg := mod.Parent.NameArg(p.Name, p.Type.QualType, n.Name)
		typ.Params.List = append(typ.Params.List, &dst.Field{
			Names: []*dst.Ident{dst.NewIdent(arg)},
			Type:  mod.Parent.ConvertQualType(p.Type.QualType),
		})
		callExpr.Args = append(callExpr.Args, &dst.CallExpr{
			Fun: &dst.SelectorExpr{
				X:   dst.NewIdent("gengort"),
				Sel: dst.NewIdent("MarshallSyscall"),
			},
			Args: []dst.Expr{
				&dst.Ident{Name: arg},
			},
		})
	}
	callExpr.Fun = &dst.SelectorExpr{
		X:   dst.NewIdent("gengort"),
		Sel: dst.NewIdent("CCall" + strconv.Itoa(len(callExpr.Args)-1)),
	}

	// Infer the method receiver.
	rcv, mnameNew := mod.Parent.InferMethod(n.Name)
	idx := -1
	if rcv != "" {
		for i, p := range paramNodes {
			if strings.Contains(p.Type.QualType, rcv) {
				idx = i
				break
			}
		}
		if idx != -1 {
			name = mod.Parent.NameFunc(mnameNew)
			decl.Name = dst.NewIdent(name)
			decl.Recv = &dst.FieldList{
				List: []*dst.Field{
					typ.Params.List[idx],
				},
			}
			typ.Params.List = append(typ.Params.List[:idx], typ.Params.List[idx+1:]...)
		}
	}

	// Add the result.
	if result != nil {
		typ.Results.List = append(typ.Results.List, &dst.Field{
			Names: []*dst.Ident{},
			Type:  result,
		})
		decl.Body = &dst.BlockStmt{
			List: []dst.Stmt{
				&dst.AssignStmt{
					Lhs: []dst.Expr{
						dst.NewIdent("__res"),
					},
					Tok: token.DEFINE,
					Rhs: []dst.Expr{
						callExpr,
					},
				},
				&dst.ReturnStmt{
					Results: []dst.Expr{
						&dst.CallExpr{
							Fun: &dst.IndexExpr{
								X: &dst.SelectorExpr{
									X:   dst.NewIdent("gengort"),
									Sel: dst.NewIdent("UnmarshallSyscall"),
								},
								Index: result2,
							},
							Args: []dst.Expr{
								dst.NewIdent("__res"),
							},
						},
					},
				},
			},
		}
		// func Func(arg *T) T {
		// __res := __imp_func.Call(
		//	...
		// )
		// return unmarshalsyscall[T](__res)

	} else {
		decl.Body = &dst.BlockStmt{
			List: []dst.Stmt{
				&dst.ExprStmt{X: callExpr},
			},
		}
		// func Func(arg *T) {
		// __imp_func.Call(
		//	...
		// )
	}
}

func (mod Module) EmitTypedef(n *clang.TypedefDecl) {
	// Resolve the type.
	var ty dst.Expr
	if inner := clang.First[clang.TypeNode](n); inner != nil {
		ty = mod.Parent.ConvertTypeExpr(inner)
	} else {
		ty = mod.Parent.ConvertQualType(n.Type.QualType)
	}

	// Check if the typedef is a simple alias.
	if ident, ok := ty.(*dst.Ident); ok {
		// Find the aliased type.
		if res, ok := mod.Parent.FindType(ident.Name); ok {
			// If there is a declaration:
			if res.Decl != nil {
				// If the aliased type looks like it's a dummy name for the sake of:
				// typedef struct Name_ {} Name;
				dummy := strings.HasSuffix(ident.Name, "_")
				dummy = dummy || strings.HasSuffix(ident.Name, "_s")
				dummy = dummy || strings.HasPrefix(ident.Name, "_")
				if dummy {
					// Rename the aliased type.
					res.Rename(mod.Parent.NameType(n.Name))
					mod.Parent.RemapType(n.Name, res)
					return
				}
			}
		}
	}

	// Add the typedef.
	tydef := mod.AddType(TcTypedef, n.Name, ty)
	tydef.Rename(mod.Parent.NameType(n.Name))
	copyDecoration(&tydef.Decl.Decs.NodeDecs, n)
}

func (mod Module) EmitFrom(ast clang.Node, layouts *clang.LayoutMap) {
	deferred := deferred{}
	validators := []initTypeValidation{}

	// Define enums.
	clang.Visit(ast, func(ed *clang.EnumDecl) bool {
		mod.EmitEnum(ed)
		return true
	})

	// Define structs and unions.
	clang.Visit(ast, func(rd *clang.RecordDecl) bool {
		if rd.CompleteDefinition {
			mod.EmitStruct(rd, layouts, &deferred, &validators)
		}
		return true
	})

	mylog.Todo("TypedefDecl cpp model not working")
	// Define typedefs.
	clang.Visit(ast, func(td *clang.TypedefDecl) bool {
		mod.EmitTypedef(td)
		return true
	})

	// Define functions.
	clang.Visit(ast, func(fd *clang.FunctionDecl) bool {
		mod.EmitFunction(fd)
		return true
	})

	// Run deferred functions.
	deferred.Go()

	// Validate types.
	validatorBody := []dst.Stmt{}
	for _, v := range validators {
		args := []dst.Expr{
			// (*T)(nil)
			&dst.CallExpr{
				Fun: &dst.ParenExpr{
					X: &dst.StarExpr{
						X: v.ty.Ref(),
					},
				},
				Args: []dst.Expr{
					dst.NewIdent("nil"),
				},
			},
			// 0x1234
			&dst.BasicLit{
				Kind:  token.INT,
				Value: fmt.Sprintf("0x%x", v.size),
			},
			// 0x1234
			&dst.BasicLit{
				Kind:  token.INT,
				Value: fmt.Sprintf("0x%x", v.align),
			},
		}
		args = append(args, v.fields...)

		validatorBody = append(validatorBody, &dst.ExprStmt{
			X: &dst.CallExpr{
				Fun: &dst.SelectorExpr{
					X:   dst.NewIdent("gengort"),
					Sel: dst.NewIdent("Validate"),
				},
				Args: args,
			},
		})
	}
	mod.OnInit(validatorBody...)
}
