package header
type (
Routines interface{
TestCase()(ok bool)//col:26
}

)
func NewRoutines() { return & routines{} }
func (r *routines)TestCase()(ok bool){//col:26
return true
}

