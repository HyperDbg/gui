package core
type (
	Termination interface {
		TerminateExternalInterruptEvent() (ok bool)       //col:92
		TerminateHiddenHookReadAndWriteEvent() (ok bool)  //col:125
		TerminateHiddenHookExecCcEvent() (ok bool)        //col:151
		TerminateHiddenHookExecDetoursEvent() (ok bool)   //col:179
		TerminateRdmsrExecutionEvent() (ok bool)          //col:258
		TerminateWrmsrExecutionEvent() (ok bool)          //col:337
		TerminateExceptionEvent() (ok bool)               //col:417
		TerminateInInstructionExecutionEvent() (ok bool)  //col:503
		TerminateOutInstructionExecutionEvent() (ok bool) //col:589
		TerminateVmcallExecutionEvent() (ok bool)         //col:627
		TerminateCpuidExecutionEvent() (ok bool)          //col:665
		TerminateTscEvent() (ok bool)                     //col:744
		TerminatePmcEvent() (ok bool)                     //col:823
		TerminateControlRegistersEvent() (ok bool)        //col:902
		TerminateDebugRegistersEvent() (ok bool)          //col:981
		TerminateSyscallHookEferEvent() (ok bool)         //col:1067
		TerminateSysretHookEferEvent() (ok bool)          //col:1153
	}
)
func NewTermination() { return &termination{} }
func (t *termination) TerminateExternalInterruptEvent() (ok bool) { //col:92
	return true
}

func (t *termination) TerminateHiddenHookReadAndWriteEvent() (ok bool) { //col:125
	return true
}

func (t *termination) TerminateHiddenHookExecCcEvent() (ok bool) { //col:151
	return true
}

func (t *termination) TerminateHiddenHookExecDetoursEvent() (ok bool) { //col:179
	return true
}

func (t *termination) TerminateRdmsrExecutionEvent() (ok bool) { //col:258
	return true
}

func (t *termination) TerminateWrmsrExecutionEvent() (ok bool) { //col:337
	return true
}

func (t *termination) TerminateExceptionEvent() (ok bool) { //col:417
	return true
}

func (t *termination) TerminateInInstructionExecutionEvent() (ok bool) { //col:503
	return true
}

func (t *termination) TerminateOutInstructionExecutionEvent() (ok bool) { //col:589
	return true
}

func (t *termination) TerminateVmcallExecutionEvent() (ok bool) { //col:627
	return true
}

func (t *termination) TerminateCpuidExecutionEvent() (ok bool) { //col:665
	return true
}

func (t *termination) TerminateTscEvent() (ok bool) { //col:744
	return true
}

func (t *termination) TerminatePmcEvent() (ok bool) { //col:823
	return true
}

func (t *termination) TerminateControlRegistersEvent() (ok bool) { //col:902
	return true
}

func (t *termination) TerminateDebugRegistersEvent() (ok bool) { //col:981
	return true
}

func (t *termination) TerminateSyscallHookEferEvent() (ok bool) { //col:1067
	return true
}

func (t *termination) TerminateSysretHookEferEvent() (ok bool) { //col:1153
	return true
}

