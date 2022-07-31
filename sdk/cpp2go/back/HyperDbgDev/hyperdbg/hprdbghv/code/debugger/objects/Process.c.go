package objects
type (
	Process interface {
		ProcessHandleProcessChange() (ok bool)                       //col:44
		ProcessSwitch() (ok bool)                                    //col:124
		ProcessDetectChangeByInterceptingClockInterrupts() (ok bool) //col:163
		ProcessDetectChangeByMov2Cr3Vmexits() (ok bool)              //col:204
		ProcessEnableOrDisableThreadChangeMonitor() (ok bool)        //col:233
		ProcessCheckIfEprocessIsValid() (ok bool)                    //col:307
		ProcessShowList() (ok bool)                                  //col:493
		ProcessInterpretProcess() (ok bool)                          //col:585
		ProcessQueryCount() (ok bool)                                //col:616
		ProcessQueryList() (ok bool)                                 //col:644
		ProcessQueryDetails() (ok bool)                              //col:665
	}
)
func NewProcess() { return &process{} }
func (p *process) ProcessHandleProcessChange() (ok bool) { //col:44
	return true
}

func (p *process) ProcessSwitch() (ok bool) { //col:124
	return true
}

func (p *process) ProcessDetectChangeByInterceptingClockInterrupts() (ok bool) { //col:163
	return true
}

func (p *process) ProcessDetectChangeByMov2Cr3Vmexits() (ok bool) { //col:204
	return true
}

func (p *process) ProcessEnableOrDisableThreadChangeMonitor() (ok bool) { //col:233
	return true
}

func (p *process) ProcessCheckIfEprocessIsValid() (ok bool) { //col:307
	return true
}

func (p *process) ProcessShowList() (ok bool) { //col:493
	return true
}

func (p *process) ProcessInterpretProcess() (ok bool) { //col:585
	return true
}

func (p *process) ProcessQueryCount() (ok bool) { //col:616
	return true
}

func (p *process) ProcessQueryList() (ok bool) { //col:644
	return true
}

func (p *process) ProcessQueryDetails() (ok bool) { //col:665
	return true
}

