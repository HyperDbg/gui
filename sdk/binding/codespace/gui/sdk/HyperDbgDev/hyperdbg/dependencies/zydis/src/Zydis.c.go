package src



type (
	Zydis interface {
		ZyanU64_ZydisGetVersion() (ok bool)          //col:4
		ZyanStatus_ZydisIsFeatureEnabled() (ok bool) //col:36
	}
	zydis struct{}
)

func NewZydis() Zydis { return &zydis{} }

func (z *zydis) ZyanU64_ZydisGetVersion() (ok bool) { //col:4







	return true
}


func (z *zydis) ZyanStatus_ZydisIsFeatureEnabled() (ok bool) { //col:36



























































	return true
}


