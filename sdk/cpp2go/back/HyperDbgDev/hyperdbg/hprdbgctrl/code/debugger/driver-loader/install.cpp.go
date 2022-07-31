package driver-loader


type (
Install interface{
InstallDriver()(ok bool)//col:111
ManageDriver()(ok bool)//col:217
RemoveDriver()(ok bool)//col:277
StartDriver()(ok bool)//col:375
StopDriver()(ok bool)//col:433
SetupDriverName()(ok bool)//col:517
}

)

func NewInstall() { return & install{} }

func (i *install)InstallDriver()(ok bool){//col:111







































return true
}


func (i *install)ManageDriver()(ok bool){//col:217












































return true
}


func (i *install)RemoveDriver()(ok bool){//col:277

























return true
}


func (i *install)StartDriver()(ok bool){//col:375














































return true
}


func (i *install)StopDriver()(ok bool){//col:433


























return true
}


func (i *install)SetupDriverName()(ok bool){//col:517






























return true
}




