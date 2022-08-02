package transparency



type (
	Transparency interface {
		TransparentModeRdtscDiffVmexit() (ok bool)         //col:30
		TransparentModeRdtscEmulationDetection() (ok bool) //col:46
	}
	transparency struct{}
)

func NewTransparency() Transparency { return &transparency{} }

func (t *transparency) TransparentModeRdtscDiffVmexit() (ok bool) { //col:30






































	return true
}


func (t *transparency) TransparentModeRdtscEmulationDetection() (ok bool) { //col:46




















	return true
}


