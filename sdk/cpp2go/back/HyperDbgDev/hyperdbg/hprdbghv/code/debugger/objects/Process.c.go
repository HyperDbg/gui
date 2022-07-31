package objects


type (
Process interface{
ProcessHandleProcessChange()(ok bool)//col:44
ProcessSwitch()(ok bool)//col:125
ProcessDetectChangeByInterceptingClockInterrupts()(ok bool)//col:165
ProcessDetectChangeByMov2Cr3Vmexits()(ok bool)//col:207
ProcessEnableOrDisableThreadChangeMonitor()(ok bool)//col:237
ProcessCheckIfEprocessIsValid()(ok bool)//col:312
ProcessShowList()(ok bool)//col:499
ProcessInterpretProcess()(ok bool)//col:592
ProcessQueryCount()(ok bool)//col:624
ProcessQueryList()(ok bool)//col:653
ProcessQueryDetails()(ok bool)//col:675
}

)

func NewProcess() { return & process{} }

func (p *process)ProcessHandleProcessChange()(ok bool){//col:44










return true
}


func (p *process)ProcessSwitch()(ok bool){//col:125











































return true
}


func (p *process)ProcessDetectChangeByInterceptingClockInterrupts()(ok bool){//col:165














return true
}


func (p *process)ProcessDetectChangeByMov2Cr3Vmexits()(ok bool){//col:207














return true
}


func (p *process)ProcessEnableOrDisableThreadChangeMonitor()(ok bool){//col:237













return true
}


func (p *process)ProcessCheckIfEprocessIsValid()(ok bool){//col:312































return true
}


func (p *process)ProcessShowList()(ok bool){//col:499




























































































return true
}


func (p *process)ProcessInterpretProcess()(ok bool){//col:592











































return true
}


func (p *process)ProcessQueryCount()(ok bool){//col:624
















return true
}


func (p *process)ProcessQueryList()(ok bool){//col:653












return true
}


func (p *process)ProcessQueryDetails()(ok bool){//col:675










return true
}




