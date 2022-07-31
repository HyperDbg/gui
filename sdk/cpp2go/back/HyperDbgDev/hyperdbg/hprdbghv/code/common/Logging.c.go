package common
type (
	Logging interface {
		LogInitialize() (ok bool)    //col:93
		LogUnInitialize() (ok bool)  //col:120
		LogSendBuffer() (ok bool)    //col:341
		LogMarkAllAsRead() (ok bool) //col:478
		*or not ()
(ok bool) //col:699
* ()(ok bool)                              //col:751
LogPrepareAndSendMessageToQueue()(ok bool) //col:898
LogSendMessageToQueue()(ok bool) //col:1047
LogNotifyUsermodeCallback()(ok bool) //col:1162
LogRegisterIrpBasedNotification()(ok bool) //col:1278
LogRegisterEventBasedNotification()(ok bool) //col:1340
}

)
func NewLogging() { return &logging{} }
func (l *logging) LogInitialize() (ok bool) { //col:93
	return true
}

func (l *logging) LogUnInitialize() (ok bool) { //col:120
	return true
}

func (l *logging) LogSendBuffer() (ok bool) { //col:341
	return true
}

func (l *logging) LogMarkAllAsRead() (ok bool) { //col:478
	return true
}

func (l *logging) * or not ()(ok bool) { //col:699
	return true
}

func (l *logging) * ()(ok bool) { //col:751
	return true
}

func (l *logging) LogPrepareAndSendMessageToQueue() (ok bool) { //col:898
	return true
}

func (l *logging) LogSendMessageToQueue() (ok bool) { //col:1047
	return true
}

func (l *logging) LogNotifyUsermodeCallback() (ok bool) { //col:1162
	return true
}

func (l *logging) LogRegisterIrpBasedNotification() (ok bool) { //col:1278
	return true
}

func (l *logging) LogRegisterEventBasedNotification() (ok bool) { //col:1340
	return true
}

