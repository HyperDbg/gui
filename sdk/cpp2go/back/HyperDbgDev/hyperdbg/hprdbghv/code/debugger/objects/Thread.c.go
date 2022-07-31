package objects


type (
Thread interface{
ThreadHandleThreadChange()(ok bool)//col:46
ThreadSwitch()(ok bool)//col:110
ThreadShowList()(ok bool)//col:316
ThreadInterpretThread()(ok bool)//col:409
ThreadDetectChangeByDebugRegisterOnGs()(ok bool)//col:542
ThreadDetectChangeByInterceptingClockInterrupts()(ok bool)//col:582
ThreadEnableOrDisableThreadChangeMonitor()(ok bool)//col:611
ThreadQueryCount()(ok bool)//col:643
ThreadQueryList()(ok bool)//col:672
ThreadQueryDetails()(ok bool)//col:697
}

)

func NewThread() { return & thread{} }

func (t *thread)ThreadHandleThreadChange()(ok bool){//col:46










return true
}


func (t *thread)ThreadSwitch()(ok bool){//col:110
































return true
}


func (t *thread)ThreadShowList()(ok bool){//col:316








































































































return true
}


func (t *thread)ThreadInterpretThread()(ok bool){//col:409













































return true
}


func (t *thread)ThreadDetectChangeByDebugRegisterOnGs()(ok bool){//col:542






























return true
}


func (t *thread)ThreadDetectChangeByInterceptingClockInterrupts()(ok bool){//col:582














return true
}


func (t *thread)ThreadEnableOrDisableThreadChangeMonitor()(ok bool){//col:611













return true
}


func (t *thread)ThreadQueryCount()(ok bool){//col:643
















return true
}


func (t *thread)ThreadQueryList()(ok bool){//col:672












return true
}


func (t *thread)ThreadQueryDetails()(ok bool){//col:697












return true
}




