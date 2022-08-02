package transparency



type (
	GaussianRng interface {
		Median() (ok bool)                     //col:20
		Average() (ok bool)                    //col:30
		CalculateStandardDeviation() (ok bool) //col:89
	}
	gaussianRng struct{}
)

func NewGaussianRng() GaussianRng { return &gaussianRng{} }

func (g *gaussianRng) Median() (ok bool) { //col:20























	return true
}


func (g *gaussianRng) Average() (ok bool) { //col:30













	return true
}


func (g *gaussianRng) CalculateStandardDeviation() (ok bool) { //col:89




































































	return true
}


