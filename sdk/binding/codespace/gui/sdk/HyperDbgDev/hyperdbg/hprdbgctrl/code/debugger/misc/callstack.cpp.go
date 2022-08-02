package misc



type (
	Callstack interface {
		CallstackReturnAddressToCallingAddress() (ok bool) //col:74
	}
	callstack struct{}
)

func NewCallstack() Callstack { return &callstack{} }

func (c *callstack) CallstackReturnAddressToCallingAddress() (ok bool) { //col:74













































































	return true
}


