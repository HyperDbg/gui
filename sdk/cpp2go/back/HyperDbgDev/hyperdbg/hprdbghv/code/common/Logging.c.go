package common


type (
Logging interface{
LogInitialize()(ok bool)//col:93
LogUnInitialize()(ok bool)//col:121
LogSendBuffer()(ok bool)//col:343
LogMarkAllAsRead()(ok bool)//col:481
LogReadBuffer()(ok bool)//col:703
LogCheckForNewMessage()(ok bool)//col:756
LogPrepareAndSendMessageToQueue()(ok bool)//col:904
LogSendMessageToQueue()(ok bool)//col:1054
LogNotifyUsermodeCallback()(ok bool)//col:1170
LogRegisterIrpBasedNotification()(ok bool)//col:1287
LogRegisterEventBasedNotification()(ok bool)//col:1350
}
















)

func NewLogging() { return & logging{} }

func (l *logging)LogInitialize()(ok bool){//col:93





























return true
}

















func (l *logging)LogUnInitialize()(ok bool){//col:121










return true
}

















func (l *logging)LogSendBuffer()(ok bool){//col:343































































































return true
}

















func (l *logging)LogMarkAllAsRead()(ok bool){//col:481























































return true
}

















func (l *logging)LogReadBuffer()(ok bool){//col:703










































































































return true
}

















func (l *logging)LogCheckForNewMessage()(ok bool){//col:756



























return true
}

















func (l *logging)LogPrepareAndSendMessageToQueue()(ok bool){//col:904





















































return true
}

















func (l *logging)LogSendMessageToQueue()(ok bool){//col:1054












































































return true
}

















func (l *logging)LogNotifyUsermodeCallback()(ok bool){//col:1170





























































return true
}

















func (l *logging)LogRegisterIrpBasedNotification()(ok bool){//col:1287


















































return true
}

















func (l *logging)LogRegisterEventBasedNotification()(ok bool){//col:1350































return true
}



















