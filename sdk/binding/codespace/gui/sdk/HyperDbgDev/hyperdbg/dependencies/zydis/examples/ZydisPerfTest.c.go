package examples

type (
	ZydisPerfTest interface {
		static_void_StartCounter() (ok bool) //col:16
		static_void_StartCounter() (ok bool) //col:27
		static_void_StartCounter() (ok bool) //col:37
	}
	zydisPerfTest struct{}
)

func NewZydisPerfTest() ZydisPerfTest { return &zydisPerfTest{} }

func (z *zydisPerfTest) static_void_StartCounter() (ok bool) { //col:16

	return true
}

func (z *zydisPerfTest) static_void_StartCounter() (ok bool) { //col:27

	return true
}

func (z *zydisPerfTest) static_void_StartCounter() (ok bool) { //col:37

	return true
}
