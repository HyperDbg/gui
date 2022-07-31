package core


type (
Termination interface{
TerminateExternalInterruptEvent()(ok bool)//col:92
TerminateHiddenHookReadAndWriteEvent()(ok bool)//col:126
TerminateHiddenHookExecCcEvent()(ok bool)//col:153
TerminateHiddenHookExecDetoursEvent()(ok bool)//col:182
TerminateRdmsrExecutionEvent()(ok bool)//col:262
TerminateWrmsrExecutionEvent()(ok bool)//col:342
TerminateExceptionEvent()(ok bool)//col:423
TerminateInInstructionExecutionEvent()(ok bool)//col:510
TerminateOutInstructionExecutionEvent()(ok bool)//col:597
TerminateVmcallExecutionEvent()(ok bool)//col:636
TerminateCpuidExecutionEvent()(ok bool)//col:675
TerminateTscEvent()(ok bool)//col:755
TerminatePmcEvent()(ok bool)//col:835
TerminateControlRegistersEvent()(ok bool)//col:915
TerminateDebugRegistersEvent()(ok bool)//col:995
TerminateSyscallHookEferEvent()(ok bool)//col:1082
TerminateSysretHookEferEvent()(ok bool)//col:1169
}






)

func NewTermination() { return & termination{} }

func (t *termination)TerminateExternalInterruptEvent()(ok bool){//col:92





























return true
}







func (t *termination)TerminateHiddenHookReadAndWriteEvent()(ok bool){//col:126












return true
}







func (t *termination)TerminateHiddenHookExecCcEvent()(ok bool){//col:153




return true
}







func (t *termination)TerminateHiddenHookExecDetoursEvent()(ok bool){//col:182




return true
}







func (t *termination)TerminateRdmsrExecutionEvent()(ok bool){//col:262





























return true
}







func (t *termination)TerminateWrmsrExecutionEvent()(ok bool){//col:342





























return true
}







func (t *termination)TerminateExceptionEvent()(ok bool){//col:423





























return true
}







func (t *termination)TerminateInInstructionExecutionEvent()(ok bool){//col:510






























return true
}







func (t *termination)TerminateOutInstructionExecutionEvent()(ok bool){//col:597






























return true
}







func (t *termination)TerminateVmcallExecutionEvent()(ok bool){//col:636











return true
}







func (t *termination)TerminateCpuidExecutionEvent()(ok bool){//col:675











return true
}







func (t *termination)TerminateTscEvent()(ok bool){//col:755





























return true
}







func (t *termination)TerminatePmcEvent()(ok bool){//col:835





























return true
}







func (t *termination)TerminateControlRegistersEvent()(ok bool){//col:915





























return true
}







func (t *termination)TerminateDebugRegistersEvent()(ok bool){//col:995





























return true
}







func (t *termination)TerminateSyscallHookEferEvent()(ok bool){//col:1082






























return true
}







func (t *termination)TerminateSysretHookEferEvent()(ok bool){//col:1169






























return true
}









