package code


type (
PseudoRegisters interface{
ScriptEnginePseudoRegGetTid()(ok bool)//col:34
ScriptEnginePseudoRegGetCore()(ok bool)//col:52
ScriptEnginePseudoRegGetPid()(ok bool)//col:70
ScriptEnginePseudoRegGetPname()(ok bool)//col:120
ScriptEnginePseudoRegGetProc()(ok bool)//col:138
ScriptEnginePseudoRegGetThread()(ok bool)//col:156
ScriptEnginePseudoRegGetPeb()(ok bool)//col:288
ScriptEnginePseudoRegGetTeb()(ok bool)//col:306
ScriptEnginePseudoRegGetIp()(ok bool)//col:327
ScriptEnginePseudoRegGetBuffer()(ok bool)//col:349
ScriptEnginePseudoRegGetEventTag()(ok bool)//col:368
ScriptEnginePseudoRegGetEventId()(ok bool)//col:387
}

)

func NewPseudoRegisters() { return & pseudoRegisters{} }

func (p *pseudoRegisters)ScriptEnginePseudoRegGetTid()(ok bool){//col:34







return true
}


func (p *pseudoRegisters)ScriptEnginePseudoRegGetCore()(ok bool){//col:52







return true
}


func (p *pseudoRegisters)ScriptEnginePseudoRegGetPid()(ok bool){//col:70







return true
}


func (p *pseudoRegisters)ScriptEnginePseudoRegGetPname()(ok bool){//col:120

























return true
}


func (p *pseudoRegisters)ScriptEnginePseudoRegGetProc()(ok bool){//col:138







return true
}


func (p *pseudoRegisters)ScriptEnginePseudoRegGetThread()(ok bool){//col:156







return true
}


func (p *pseudoRegisters)ScriptEnginePseudoRegGetPeb()(ok bool){//col:288
















































































return true
}


func (p *pseudoRegisters)ScriptEnginePseudoRegGetTeb()(ok bool){//col:306







return true
}


func (p *pseudoRegisters)ScriptEnginePseudoRegGetIp()(ok bool){//col:327







return true
}


func (p *pseudoRegisters)ScriptEnginePseudoRegGetBuffer()(ok bool){//col:349







return true
}


func (p *pseudoRegisters)ScriptEnginePseudoRegGetEventTag()(ok bool){//col:368







return true
}


func (p *pseudoRegisters)ScriptEnginePseudoRegGetEventId()(ok bool){//col:387







return true
}




