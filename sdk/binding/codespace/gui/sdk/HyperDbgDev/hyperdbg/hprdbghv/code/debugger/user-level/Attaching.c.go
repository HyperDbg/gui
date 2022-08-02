package user_level



type (
	Attaching interface {
		AttachingInitialize() (ok bool)                                 //col:55
		AttachingFindProcessDebuggingDetailsByProcessId() (ok bool)     //col:66
		AttachingFindProcessDebuggingDetailsInStartingPhase() (ok bool) //col:77
		AttachingRemoveAndFreeAllProcessDebuggingDetails() (ok bool)    //col:111
		AttachingReachedToProcessEntrypoint() (ok bool)                 //col:243
		AttachingPerformAttachToProcess() (ok bool)                     //col:379
		AttachingPauseProcess() (ok bool)                               //col:392
		AttachingKillProcess() (ok bool)                                //col:473
	}
	attaching struct{}
)

func NewAttaching() Attaching { return &attaching{} }

func (a *attaching) AttachingInitialize() (ok bool) { //col:55































































	return true
}


func (a *attaching) AttachingFindProcessDebuggingDetailsByProcessId() (ok bool) { //col:66














	return true
}


func (a *attaching) AttachingFindProcessDebuggingDetailsInStartingPhase() (ok bool) { //col:77














	return true
}


func (a *attaching) AttachingRemoveAndFreeAllProcessDebuggingDetails() (ok bool) { //col:111









































	return true
}


func (a *attaching) AttachingReachedToProcessEntrypoint() (ok bool) { //col:243
















































































































































	return true
}


func (a *attaching) AttachingPerformAttachToProcess() (ok bool) { //col:379















































































































































	return true
}


func (a *attaching) AttachingPauseProcess() (ok bool) { //col:392
















	return true
}


func (a *attaching) AttachingKillProcess() (ok bool) { //col:473


























































































	return true
}


