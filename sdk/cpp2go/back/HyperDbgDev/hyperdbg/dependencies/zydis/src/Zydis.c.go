package src


type (
Zydis interface{
ZyanU64 ZydisGetVersion()(ok bool)//col:36
ZyanStatus ZydisIsFeatureEnabled()(ok bool)//col:72
}

)

func NewZydis() { return & zydis{} }

func (z *zydis)ZyanU64 ZydisGetVersion()(ok bool){//col:36




return true
}


func (z *zydis)ZyanStatus ZydisIsFeatureEnabled()(ok bool){//col:72
































return true
}




