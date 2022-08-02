package app



type (
	Dllmain interface {
		DllMain() (ok bool) //col:12
	}
	dllmain struct{}
)

func NewDllmain() Dllmain { return &dllmain{} }

func (d *dllmain) DllMain() (ok bool) { //col:12















	return true
}


