package kdserial


const(
CLOCK_RATE = 3125000UL //col:24
)

type (
Apm88xxxx interface{
Uart16550InitializePortCommon()(ok bool)//col:103
Apm88xxxxSetBaud()(ok bool)//col:141
}
















)

func NewApm88xxxx() { return & apm88xxxx{} }

func (a *apm88xxxx)Uart16550InitializePortCommon()(ok bool){//col:103








































return true
}

















func (a *apm88xxxx)Apm88xxxxSetBaud()(ok bool){//col:141










return true
}



















