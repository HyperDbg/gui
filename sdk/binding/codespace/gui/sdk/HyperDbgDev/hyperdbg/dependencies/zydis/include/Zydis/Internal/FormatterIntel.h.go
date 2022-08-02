package Internal

type (
	FormatterIntel interface {
		ZyanStatus_ZydisFormatterIntelFormatInstruction() (ok bool) //col:59
		________________ZYAN_DEFINE_STRING_VIEW() (ok bool)         //col:131
	}
	formatterIntel struct{}
)

func NewFormatterIntel() FormatterIntel { return &formatterIntel{} }

func (f *formatterIntel) ZyanStatus_ZydisFormatterIntelFormatInstruction() (ok bool) { //col:59

	return true
}

func (f *formatterIntel) ________________ZYAN_DEFINE_STRING_VIEW() (ok bool) { //col:131

	return true
}
