package debugging-commands


type (
Eval interface{
CommandEvalHelp()(ok bool)//col:36
CommandEvalCheckTestcase()(ok bool)//col:187
CommandEval()(ok bool)//col:301
}

)

func NewEval() { return & eval{} }

func (e *eval)CommandEvalHelp()(ok bool){//col:36






return true
}


func (e *eval)CommandEvalCheckTestcase()(ok bool){//col:187























































































return true
}


func (e *eval)CommandEval()(ok bool){//col:301


















































return true
}




