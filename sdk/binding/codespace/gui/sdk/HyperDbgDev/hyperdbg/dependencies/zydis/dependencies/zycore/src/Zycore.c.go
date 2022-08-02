package src

type (
	Zycore interface {
		ZyanU64_ZycoreGetVersion() (ok bool) //col:4
	}
	zycore struct{}
)

func NewZycore() Zycore { return &zycore{} }

func (z *zycore) ZyanU64_ZycoreGetVersion() (ok bool) { //col:4

	return true
}
