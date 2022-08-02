package examples

type (
	String interface {
		static_ZyanStatus_PerformBasicTests() (ok bool) //col:14
		int_main() (ok bool)                            //col:30
	}
	string struct{}
)

func NewString() String { return &string{} }

func (s *string) static_ZyanStatus_PerformBasicTests() (ok bool) { //col:14

	return true
}

func (s *string) int_main() (ok bool) { //col:30

	return true
}
