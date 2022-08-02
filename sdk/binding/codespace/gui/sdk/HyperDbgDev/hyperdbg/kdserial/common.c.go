package kdserial



type (
	Common interface {
		DllInitialize() (ok bool) //col:5
		DllUnload() (ok bool)     //col:9
	}
	common struct{}
)

func NewCommon() Common { return &common{} }

func (c *common) DllInitialize() (ok bool) { //col:5









	return true
}


func (c *common) DllUnload() (ok bool) { //col:9







	return true
}


