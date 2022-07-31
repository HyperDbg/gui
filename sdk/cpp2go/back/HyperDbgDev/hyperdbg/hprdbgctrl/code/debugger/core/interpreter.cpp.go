package core
type (
Interpreter interface{
HyperDbgInterpreter()(ok bool)//col:262
InterpreterRemoveComments()(ok bool)//col:319
HyperDbgShowSignature()(ok bool)//col:361
HyperDbgCheckMultilineCommand()(ok bool)//col:453
HyperDbgContinuePreviousCommand()(ok bool)//col:481
GetCommandAttributes()(ok bool)//col:515
InitializeDebugger()(ok bool)//col:565
InitializeCommandsDictionary()(ok bool)//col:821
}

)
func NewInterpreter() { return & interpreter{} }
func (i *interpreter)HyperDbgInterpreter()(ok bool){//col:262
return true
}

func (i *interpreter)InterpreterRemoveComments()(ok bool){//col:319
return true
}

func (i *interpreter)HyperDbgShowSignature()(ok bool){//col:361
return true
}

func (i *interpreter)HyperDbgCheckMultilineCommand()(ok bool){//col:453
return true
}

func (i *interpreter)HyperDbgContinuePreviousCommand()(ok bool){//col:481
return true
}

func (i *interpreter)GetCommandAttributes()(ok bool){//col:515
return true
}

func (i *interpreter)InitializeDebugger()(ok bool){//col:565
return true
}

func (i *interpreter)InitializeCommandsDictionary()(ok bool){//col:821
return true
}

