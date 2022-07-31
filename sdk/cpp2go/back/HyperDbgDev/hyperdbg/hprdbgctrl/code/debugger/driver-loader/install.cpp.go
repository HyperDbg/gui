package driver-loader
type (
Install interface{
InstallDriver()(ok bool)//col:111
ManageDriver()(ok bool)//col:216
RemoveDriver()(ok bool)//col:275
StartDriver()(ok bool)//col:372
StopDriver()(ok bool)//col:429
SetupDriverName()(ok bool)//col:512
}

)
func NewInstall() { return & install{} }
func (i *install)InstallDriver()(ok bool){//col:111
return true
}

func (i *install)ManageDriver()(ok bool){//col:216
return true
}

func (i *install)RemoveDriver()(ok bool){//col:275
return true
}

func (i *install)StartDriver()(ok bool){//col:372
return true
}

func (i *install)StopDriver()(ok bool){//col:429
return true
}

func (i *install)SetupDriverName()(ok bool){//col:512
return true
}

