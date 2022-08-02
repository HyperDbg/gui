package tests

type (
	ArgParse interface {
		auto_cvt_string_view() (ok bool) //col:8
	}
	argParse struct{}
)

func NewArgParse() ArgParse { return &argParse{} }

func (a *argParse) auto_cvt_string_view() (ok bool) { //col:8

	return true
}
