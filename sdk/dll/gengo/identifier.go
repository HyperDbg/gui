package gengo

import "github.com/dave/dst"

type Identifier interface {
	String() string
	Ref() *dst.Ident
	Rename(string)
}

type TrackedIdentifier struct {
	Name string
	Refs []*dst.Ident
}

func (t *TrackedIdentifier) Ref() *dst.Ident {
	i := dst.NewIdent(t.Name)
	t.Refs = append(t.Refs, i)
	return i
}

func (t *TrackedIdentifier) String() string {
	return t.Name
}

func (t *TrackedIdentifier) Rename(k string) {
	t.Name = k
	for _, r := range t.Refs {
		r.Name = k
	}
}

type RelaxedIdentifier string

func (r RelaxedIdentifier) Ref() *dst.Ident {
	return dst.NewIdent(string(r))
}

func (r RelaxedIdentifier) String() string {
	return string(r)
}

func (r RelaxedIdentifier) Rename(k string) {
	panic("cannot rename a relaxed identifier")
}
