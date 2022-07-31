package tests
type (
Tests interface{
SetupTestName()(ok bool)//col:100
CreateProcessAndOpenPipeConnection()(ok bool)//col:303
CloseProcessAndClosePipeConnection()(ok bool)//col:324
}

)
func NewTests() { return & tests{} }
func (t *tests)SetupTestName()(ok bool){//col:100
return true
}

func (t *tests)CreateProcessAndOpenPipeConnection()(ok bool){//col:303
return true
}

func (t *tests)CloseProcessAndClosePipeConnection()(ok bool){//col:324
return true
}

