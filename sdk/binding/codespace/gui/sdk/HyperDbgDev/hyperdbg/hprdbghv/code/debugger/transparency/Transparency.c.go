package transparency



type (
	Transparency interface {
		TransparentGetRand() (ok bool) //col:14
		TransparentLog() (ok bool)     //col:25
		TransparentSqrt() (ok bool)    //col:45
	}
	transparency struct{}
)

func NewTransparency() Transparency { return &transparency{} }

func (t *transparency) TransparentGetRand() (ok bool) { //col:14



















	return true
}


func (t *transparency) TransparentLog() (ok bool) { //col:25














	return true
}


func (t *transparency) TransparentSqrt() (ok bool) { //col:45























	return true
}


