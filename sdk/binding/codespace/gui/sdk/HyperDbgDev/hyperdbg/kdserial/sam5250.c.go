package kdserial



type (
	Sam5250 interface {
		Sam5250SetBaud() (ok bool) //col:35
		Sam5250GetByte() (ok bool) //col:61
		Sam5250PutByte() (ok bool) //col:98
	}
	sam5250 struct{}
)

func NewSam5250() Sam5250 { return &sam5250{} }

func (s *sam5250) Sam5250SetBaud() (ok bool) { //col:35












































	return true
}


func (s *sam5250) Sam5250GetByte() (ok bool) { //col:61






























	return true
}


func (s *sam5250) Sam5250PutByte() (ok bool) { //col:98












































	return true
}


