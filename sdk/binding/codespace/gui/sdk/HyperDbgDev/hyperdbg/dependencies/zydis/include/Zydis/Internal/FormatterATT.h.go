package Internal

type (
	FormatterAtt interface {
		ZyanStatus_ZydisFormatterATTFormatInstruction() (ok bool) //col:55
	}
	formatterAtt struct{}
)

func NewFormatterAtt() FormatterAtt { return &formatterAtt{} }

func (f *formatterAtt) ZyanStatus_ZydisFormatterATTFormatInstruction() (ok bool) { //col:55

	return true
}
