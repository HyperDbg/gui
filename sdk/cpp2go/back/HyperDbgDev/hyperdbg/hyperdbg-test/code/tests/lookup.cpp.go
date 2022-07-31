package tests
type (
Lookup interface{
TestListenerForResultsThread()(ok bool)//col:81
TestCreateLookupTable()(ok bool)//col:224
}

)
func NewLookup() { return & lookup{} }
func (l *lookup)TestListenerForResultsThread()(ok bool){//col:81
return true
}

func (l *lookup)TestCreateLookupTable()(ok bool){//col:224
return true
}

