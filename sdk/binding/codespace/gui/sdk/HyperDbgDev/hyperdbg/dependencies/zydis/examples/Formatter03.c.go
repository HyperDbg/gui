package examples

type (
	Formatter03 interface {
		static_void_DisassembleBuffer() (ok bool) //col:30
	}
	formatter03 struct{}
)

func NewFormatter03() Formatter03 { return &formatter03{} }

func (f *formatter03) static_void_DisassembleBuffer() (ok bool) { //col:30

	return true
}
