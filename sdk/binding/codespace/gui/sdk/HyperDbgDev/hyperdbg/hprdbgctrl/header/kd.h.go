package header



type (
	Kd interface {
		____HKeyHolder() (ok bool) //col:11
	}
	kd struct{}
)

func NewKd() Kd { return &kd{} }

func (k *kd) ____HKeyHolder() (ok bool) { //col:11













	return true
}


