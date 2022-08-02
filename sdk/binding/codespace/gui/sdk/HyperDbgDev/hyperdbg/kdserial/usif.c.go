package kdserial



type (
	Usif interface {
		UsifInitializePort() (ok bool) //col:18
		UsifSetBaud() (ok bool)        //col:30
		UsifGetByte() (ok bool)        //col:48
		UsifPutByte() (ok bool)        //col:84
	}
	usif struct{}
)

func NewUsif() Usif { return &usif{} }

func (u *usif) UsifInitializePort() (ok bool) { //col:18






















	return true
}


func (u *usif) UsifSetBaud() (ok bool) { //col:30
















	return true
}


func (u *usif) UsifGetByte() (ok bool) { //col:48






















	return true
}


func (u *usif) UsifPutByte() (ok bool) { //col:84











































	return true
}


