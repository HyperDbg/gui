package src



type (
	Formatter interface {
		void_ZydisFormatterBufferInit() (ok bool) //col:21
	}
	formatter struct{}
)

func NewFormatter() Formatter { return &formatter{} }

func (f *formatter) void_ZydisFormatterBufferInit() (ok bool) { //col:21




























	return true
}


