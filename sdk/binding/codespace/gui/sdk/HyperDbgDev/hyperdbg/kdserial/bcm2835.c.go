package kdserial



type (
	Bcm2835 interface {
		Bcm2835RxReady() (ok bool) //col:34
		Bcm2835GetByte() (ok bool) //col:51
		Bcm2835PutByte() (ok bool) //col:90
	}
	bcm2835 struct{}
)

func NewBcm2835() Bcm2835 { return &bcm2835{} }

func (b *bcm2835) Bcm2835RxReady() (ok bool) { //col:34











































	return true
}


func (b *bcm2835) Bcm2835GetByte() (ok bool) { //col:51





















	return true
}


func (b *bcm2835) Bcm2835PutByte() (ok bool) { //col:90














































	return true
}


