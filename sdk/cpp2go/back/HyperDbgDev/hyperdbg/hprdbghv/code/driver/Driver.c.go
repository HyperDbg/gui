package driver
type (
	Driver interface {
		DriverEntry() (ok bool)    //col:101
		DrvUnload() (ok bool)      //col:173
		DrvCreate() (ok bool)      //col:298
		DrvRead() (ok bool)        //col:317
		DrvWrite() (ok bool)       //col:336
		DrvClose() (ok bool)       //col:360
		DrvUnsupported() (ok bool) //col:379
	}
)
func NewDriver() { return &driver{} }
func (d *driver) DriverEntry() (ok bool) { //col:101
	return true
}

func (d *driver) DrvUnload() (ok bool) { //col:173
	return true
}

func (d *driver) DrvCreate() (ok bool) { //col:298
	return true
}

func (d *driver) DrvRead() (ok bool) { //col:317
	return true
}

func (d *driver) DrvWrite() (ok bool) { //col:336
	return true
}

func (d *driver) DrvClose() (ok bool) { //col:360
	return true
}

func (d *driver) DrvUnsupported() (ok bool) { //col:379
	return true
}

