package tests
type (
HyperdbgTest interface{
main()(ok bool)//col:157
}

)
func NewHyperdbgTest() { return & hyperdbgTest{} }
func (h *hyperdbgTest)main()(ok bool){//col:157
return true
}

