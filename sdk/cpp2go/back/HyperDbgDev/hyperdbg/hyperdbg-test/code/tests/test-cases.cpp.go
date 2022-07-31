package tests


type (
TestCases interface{
TestCase()(ok bool)//col:50
}






)

func NewTestCases() { return & testCases{} }

func (t *testCases)TestCase()(ok bool){//col:50




















return true
}









