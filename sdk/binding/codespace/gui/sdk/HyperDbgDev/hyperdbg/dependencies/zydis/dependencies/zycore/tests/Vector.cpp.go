package tests

type (
	Vector interface {
		____void_SetUp() (ok bool)                //col:28
		static_ZyanStatus_FreeZyanU16() (ok bool) //col:33
	}
	vector struct{}
)

func NewVector() Vector { return &vector{} }

func (v *vector) ____void_SetUp() (ok bool) { //col:28

	return true
}

func (v *vector) static_ZyanStatus_FreeZyanU16() (ok bool) { //col:33

	return true
}
