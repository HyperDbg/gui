package debugging-commands


type (
Test interface{
CommandTestHelp()(ok bool)//col:35
CommandTestPerformKernelTestsIoctl()(ok bool)//col:91
CommandTestPerformTest()(ok bool)//col:199
CommandTestInVmiMode()(ok bool)//col:266
CommandTestQueryState()(ok bool)//col:288
CommandTest()(ok bool)//col:321
}
















)

func NewTest() { return & test{} }

func (t *test)CommandTestHelp()(ok bool){//col:35









return true
}

















func (t *test)CommandTestPerformKernelTestsIoctl()(ok bool){//col:91























return true
}

















func (t *test)CommandTestPerformTest()(ok bool){//col:199

























































return true
}

















func (t *test)CommandTestInVmiMode()(ok bool){//col:266


























return true
}

















func (t *test)CommandTestQueryState()(ok bool){//col:288










return true
}

















func (t *test)CommandTest()(ok bool){//col:321

















return true
}



















