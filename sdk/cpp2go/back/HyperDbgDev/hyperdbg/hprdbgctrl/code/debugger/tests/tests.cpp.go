package tests


type (
Tests interface{
SetupTestName()(ok bool)//col:100
CreateProcessAndOpenPipeConnection()(ok bool)//col:304
CloseProcessAndClosePipeConnection()(ok bool)//col:326
}

)

func NewTests() { return & tests{} }

func (t *tests)SetupTestName()(ok bool){//col:100





























return true
}


func (t *tests)CreateProcessAndOpenPipeConnection()(ok bool){//col:304










































































































return true
}


func (t *tests)CloseProcessAndClosePipeConnection()(ok bool){//col:326








return true
}




