package code
type (
PseudoRegisters interface{
ScriptEnginePseudoRegGetTid()(ok bool)//col:34
ScriptEnginePseudoRegGetCore()(ok bool)//col:51
ScriptEnginePseudoRegGetPid()(ok bool)//col:68
ScriptEnginePseudoRegGetPname()(ok bool)//col:117
ScriptEnginePseudoRegGetProc()(ok bool)//col:134
ScriptEnginePseudoRegGetThread()(ok bool)//col:151
ScriptEnginePseudoRegGetPeb()(ok bool)//col:282
ScriptEnginePseudoRegGetTeb()(ok bool)//col:299
ScriptEnginePseudoRegGetIp()(ok bool)//col:319
ScriptEnginePseudoRegGetBuffer()(ok bool)//col:340
ScriptEnginePseudoRegGetEventTag()(ok bool)//col:358
ScriptEnginePseudoRegGetEventId()(ok bool)//col:376
}

)
func NewPseudoRegisters() { return & pseudoRegisters{} }
func (p *pseudoRegisters)ScriptEnginePseudoRegGetTid()(ok bool){//col:34
return true
}

func (p *pseudoRegisters)ScriptEnginePseudoRegGetCore()(ok bool){//col:51
return true
}

func (p *pseudoRegisters)ScriptEnginePseudoRegGetPid()(ok bool){//col:68
return true
}

func (p *pseudoRegisters)ScriptEnginePseudoRegGetPname()(ok bool){//col:117
return true
}

func (p *pseudoRegisters)ScriptEnginePseudoRegGetProc()(ok bool){//col:134
return true
}

func (p *pseudoRegisters)ScriptEnginePseudoRegGetThread()(ok bool){//col:151
return true
}

func (p *pseudoRegisters)ScriptEnginePseudoRegGetPeb()(ok bool){//col:282
return true
}

func (p *pseudoRegisters)ScriptEnginePseudoRegGetTeb()(ok bool){//col:299
return true
}

func (p *pseudoRegisters)ScriptEnginePseudoRegGetIp()(ok bool){//col:319
return true
}

func (p *pseudoRegisters)ScriptEnginePseudoRegGetBuffer()(ok bool){//col:340
return true
}

func (p *pseudoRegisters)ScriptEnginePseudoRegGetEventTag()(ok bool){//col:358
return true
}

func (p *pseudoRegisters)ScriptEnginePseudoRegGetEventId()(ok bool){//col:376
return true
}

