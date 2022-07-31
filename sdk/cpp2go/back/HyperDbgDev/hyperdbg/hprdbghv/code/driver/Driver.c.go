package driver


type (
Driver interface{
DriverEntry()(ok bool)//col:101
DrvUnload()(ok bool)//col:174
DrvCreate()(ok bool)//col:300
DrvRead()(ok bool)//col:320
DrvWrite()(ok bool)//col:340
DrvClose()(ok bool)//col:365
DrvUnsupported()(ok bool)//col:385
}

)

func NewDriver() { return & driver{} }

func (d *driver)DriverEntry()(ok bool){//col:101
















































return true
}


func (d *driver)DrvUnload()(ok bool){//col:174


































return true
}


func (d *driver)DrvCreate()(ok bool){//col:300



















































return true
}


func (d *driver)DrvRead()(ok bool){//col:320








return true
}


func (d *driver)DrvWrite()(ok bool){//col:340








return true
}


func (d *driver)DrvClose()(ok bool){//col:365








return true
}


func (d *driver)DrvUnsupported()(ok bool){//col:385








return true
}




