package kdserial
const(
CLOCK_RATE =   2995200ul //col:23
COM_EFR =      0x2 //col:24
COM_MDR1 =     0x8 //col:25
LC_DATA_SIZE = 0x03 //col:26
)
type (
Omap interface{
Copyright ()(ok bool)//col:146
OmapSetBaud()(ok bool)//col:250
}

)
func NewOmap() { return & omap{} }
func (o *omap)Copyright ()(ok bool){//col:146
return true
}

func (o *omap)OmapSetBaud()(ok bool){//col:250
return true
}

