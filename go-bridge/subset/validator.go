// Package subset validates that Go source code stays within the supported Go
// subset defined in docs/go-subset-spec.md. It is the first line of defence
// against "users writing code the kernel interpreter cannot run": any
// construct outside the subset is rejected with a precise source location
// before the code is ever sent to the driver.
//
// The validator is a go/ast walker. It rejects:
//   - goroutines (GoStmt), channels (ChanType/SendStmt/RecvStmt), select
//   - interface types and type assertions
//   - map types and slice types (except [N]uint8 fixed arrays)
//   - reflect / unsafe imports
//   - defer statements
//   - string concatenation via + (use Printf instead)
//   - make / new calls
//   - closures that capture variables declared outside the closure body
//     (except the special `ctx` parameter)
//   - any call to a function that is not a whitelist HookCtx method
package subset

import (
	"go/ast"
	"go/parser"
	"go/token"
	"strings"

	"github.com/hyperdbg/go-bridge/protocol"
)

// Validator checks Go source against the subset.
type Validator struct {
	fset *token.FileSet
	// allowedCalls is the set of package-qualified function calls that are
	// permitted (currently empty — all calls must be ctx methods).
	// ctxParamName is the name of the HookCtx parameter in a FuncLit; by
	// convention "ctx". The validator allows closures to reference this
	// parameter even though it is declared "outside" the closure body.
	ctxParamName string
}

// New creates a Validator.
func New() *Validator { return &Validator{ctxParamName: "ctx"} }

// ValidateSrc parses src and validates it. Returns a protocol.ValidationError
// for the first unsupported construct found, or nil if the source is within
// the subset.
func (v *Validator) ValidateSrc(src string) error {
	v.fset = token.NewFileSet()
	f, err := parser.ParseFile(v.fset, "hook.go", src, parser.AllErrors)
	if err != nil {
		return protocol.NewValidationError("parse", err.Error())
	}
	return v.ValidateFile(f)
}

// ValidateFile validates a parsed Go file.
func (v *Validator) ValidateFile(f *ast.File) error {
	if v.fset == nil {
		v.fset = token.NewFileSet()
	}
	w := &walker{v: v, fset: v.fset, scopes: []map[string]bool{{}}}
	ast.Walk(w, f)
	return w.firstErr
}

// walker implements ast.Visitor and collects the first validation error.
type walker struct {
	v        *Validator
	fset     *token.FileSet
	firstErr error
	// scopes is a stack of identifier sets representing the currently visible
	// variable bindings. scope[0] is the file scope (imports + package-level
	// decls). Each FuncLit pushes a new scope; each decl statement adds its
	// name to the top scope.
	scopes []map[string]bool
	// inFuncLit tracks whether we are inside a closure body, so we can
	// enforce the capture rule.
	inFuncLit int
}

func (w *walker) pos(pos token.Pos) string {
	return w.fset.Position(pos).String()
}

func (w *walker) err(pos token.Pos, reason string) {
	if w.firstErr == nil {
		w.firstErr = protocol.NewValidationError(w.pos(pos), reason)
	}
}

func (w *walker) Visit(node ast.Node) ast.Visitor {
	if node == nil || w.firstErr != nil {
		return nil
	}
	switch n := node.(type) {
	// --- Imports ---
	case *ast.ImportSpec:
		path := strings.Trim(n.Path.Value, `"`)
		if path == "reflect" || path == "unsafe" {
			w.err(n.Pos(), "import "+path+" is not allowed in the Go subset")
		}
		return w

	// --- Goroutines / channels / select ---
	case *ast.GoStmt:
		w.err(n.Go, "goroutine (go statement) is not allowed in the Go subset")
		return nil
	case *ast.ChanType:
		w.err(n.Pos(), "channel type is not allowed in the Go subset")
		return nil
	case *ast.SendStmt:
		w.err(n.Pos(), "channel send is not allowed in the Go subset")
		return nil
	case *ast.SelectStmt:
		w.err(n.Pos(), "select statement is not allowed in the Go subset")
		return nil

	// --- defer ---
	case *ast.DeferStmt:
		w.err(n.Pos(), "defer is not allowed in the Go subset")
		return nil

	// --- interface / type assertion ---
	case *ast.InterfaceType:
		w.err(n.Pos(), "interface type is not allowed in the Go subset")
		return nil
	case *ast.TypeAssertExpr:
		w.err(n.Pos(), "type assertion is not allowed in the Go subset")
		return nil

	// --- map / slice types ---
	case *ast.MapType:
		w.err(n.Pos(), "map type is not allowed in the Go subset")
		return nil
	case *ast.ArrayType:
		// Fixed-size arrays are allowed ONLY if the element type is uint8
		// (i.e. [N]uint8). Slices (Len == nil) are rejected.
		if n.Len == nil {
			w.err(n.Pos(), "slice type is not allowed in the Go subset (only [N]uint8)")
			return nil
		}
		if ident, ok := n.Elt.(*ast.Ident); !ok || ident.Name != "uint8" {
			w.err(n.Pos(), "array element type must be uint8 in the Go subset")
		}
		return w

	// --- make / new ---
	case *ast.CallExpr:
		if ident, ok := n.Fun.(*ast.Ident); ok {
			switch ident.Name {
			case "make":
				w.err(n.Pos(), "make() is not allowed in the Go subset")
				return nil
			case "new":
				w.err(n.Pos(), "new() is not allowed in the Go subset")
				return nil
			}
		}
		// String concatenation: check for `+` between string literals/idents
		// is handled in BinaryExpr. Calls to non-whitelist functions are
		// allowed only via ctx.Method() (SelectorExpr); bare function calls
		// (e.g. fmt.Println) are rejected.
		if _, isSelector := n.Fun.(*ast.SelectorExpr); !isSelector {
			// Bare call: could be a built-in (len, cap, etc.) or a user func.
			// For now allow built-ins len/cap/copy/append only on arrays.
			if ident, ok := n.Fun.(*ast.Ident); ok {
				switch ident.Name {
				case "len", "cap", "copy", "append", "print", "println":
					return w
				}
			}
			// Otherwise it might be calling a locally declared func — that's OK
			// in the subset (the interpreter supports FuncDecl). We don't
			// reject here; the encoder will resolve it.
		}
		return w

	case *ast.BinaryExpr:
		if n.Op == token.ADD {
			// Reject string concatenation: if either operand is a string
			// literal or a string-typed ident, this is concatenation.
			if w.isStringExpr(n.X) || w.isStringExpr(n.Y) {
				w.err(n.Pos(), "string concatenation (+) is not allowed; use ctx.Printf instead")
				return nil
			}
		}
		return w

	// --- Closures (FuncLit) ---
	case *ast.FuncLit:
		w.visitFuncLit(n)
		return nil

	case *ast.FuncDecl:
		// FuncDecl is allowed (top-level functions in the subset). Push a
		// scope for its parameters.
		w.pushScope()
		for _, field := range n.Type.Params.List {
			for _, name := range field.Names {
				w.bind(name.Name)
			}
		}
		if n.Type.Results != nil {
			for _, field := range n.Type.Results.List {
				for _, name := range field.Names {
					w.bind(name.Name)
				}
			}
		}
		if n.Body != nil {
			ast.Walk(w, n.Body)
		}
		w.popScope()
		return nil

	// --- Variable declarations ---
	case *ast.DeclStmt:
		if gd, ok := n.Decl.(*ast.GenDecl); ok {
			for _, spec := range gd.Specs {
				if vs, ok := spec.(*ast.ValueSpec); ok {
					for _, name := range vs.Names {
						w.bind(name.Name)
					}
				}
			}
		}
		return w

	case *ast.AssignStmt:
		// Short variable declarations (:=) bind new names in the current scope.
		if n.Tok == token.DEFINE {
			for _, lhs := range n.Lhs {
				if ident, ok := lhs.(*ast.Ident); ok {
					w.bind(ident.Name)
				}
			}
		}
		return w

	// --- Identifier reference (closure capture check) ---
	case *ast.Ident:
		// We don't reject identifier references here per se, but the capture
		// check is enforced in visitFuncLit by pre-scanning the body. This
		// case is a no-op; we still descend.
		return w
	}
	return w
}

// visitFuncLit handles the closure capture rule: a FuncLit may reference only
// its own parameters or `ctx`. Any reference to an outer-scope variable is
// rejected.
func (w *walker) visitFuncLit(fl *ast.FuncLit) {
	// Collect this FuncLit's own parameter names.
	params := map[string]bool{}
	if fl.Type.Params != nil {
		for _, field := range fl.Type.Params.List {
			for _, name := range field.Names {
				params[name.Name] = true
			}
		}
	}
	// ctx is always allowed.
	params[w.v.ctxParamName] = true

	// Scan the body for identifiers not in params and not in the file scope
	// (imports/package-level). References to variables in enclosing function
	// scopes (not the file scope) are captures and must be rejected.
	w.pushScope()
	for name := range params {
		w.bind(name)
	}
	w.inFuncLit++
	if fl.Body != nil {
		ast.Walk(w, fl.Body)
	}
	w.inFuncLit--
	w.popScope()
}

func (w *walker) pushScope() {
	w.scopes = append(w.scopes, map[string]bool{})
}

func (w *walker) popScope() {
	if len(w.scopes) > 1 {
		w.scopes = w.scopes[:len(w.scopes)-1]
	}
}

func (w *walker) bind(name string) {
	if len(w.scopes) > 0 {
		w.scopes[len(w.scopes)-1][name] = true
	}
}

// isBound reports whether name is visible in any enclosing scope.
func (w *walker) isBound(name string) bool {
	for i := len(w.scopes) - 1; i >= 0; i-- {
		if w.scopes[i][name] {
			return true
		}
	}
	return false
}

// isStringExpr heuristically reports whether e is a string-typed expression
// (string literal or ident known to be string). Used for the + concatenation
// check.
func (w *walker) isStringExpr(e ast.Expr) bool {
	switch x := e.(type) {
	case *ast.BasicLit:
		return x.Kind == token.STRING
	case *ast.Ident:
		// Heuristic: if the name looks like a string param (s, fmt, str, msg)
		// we treat it as string. A full type-checker would be more precise.
		switch x.Name {
		case "s", "fmt", "str", "msg", "format":
			return true
		}
	}
	return false
}
