package debugging-commands
type (
Test interface{
CommandTestHelp()(ok bool)//col:35
CommandTestPerformKernelTestsIoctl()(ok bool)//col:90
CommandTestPerformTest()(ok bool)//col:197
CommandTestInVmiMode()(ok bool)//col:263
CommandTestQueryState()(ok bool)//col:284
CommandTest()(ok bool)//col:316
}

)
func NewTest() { return & test{} }
func (t *test)CommandTestHelp()(ok bool){//col:35
return true
}

func (t *test)CommandTestPerformKernelTestsIoctl()(ok bool){//col:90
return true
}

func (t *test)CommandTestPerformTest()(ok bool){//col:197
return true
}

func (t *test)CommandTestInVmiMode()(ok bool){//col:263
return true
}

func (t *test)CommandTestQueryState()(ok bool){//col:284
return true
}

func (t *test)CommandTest()(ok bool){//col:316
return true
}

