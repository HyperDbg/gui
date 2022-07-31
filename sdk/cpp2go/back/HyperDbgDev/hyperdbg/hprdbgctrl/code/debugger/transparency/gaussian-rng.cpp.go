package transparency
type (
GaussianRng interface{
Median()(ok bool)//col:42
Average()(ok bool)//col:67
CalculateStandardDeviation()(ok bool)//col:88
MedianAbsoluteDeviationTest()(ok bool)//col:111
Randn()(ok bool)//col:147
GuassianGenerateRandom()(ok bool)//col:215
TestGaussianFromFile()(ok bool)//col:243
}

)
func NewGaussianRng() { return & gaussianRng{} }
func (g *gaussianRng)Median()(ok bool){//col:42
return true
}

func (g *gaussianRng)Average()(ok bool){//col:67
return true
}

func (g *gaussianRng)CalculateStandardDeviation()(ok bool){//col:88
return true
}

func (g *gaussianRng)MedianAbsoluteDeviationTest()(ok bool){//col:111
return true
}

func (g *gaussianRng)Randn()(ok bool){//col:147
return true
}

func (g *gaussianRng)GuassianGenerateRandom()(ok bool){//col:215
return true
}

func (g *gaussianRng)TestGaussianFromFile()(ok bool){//col:243
return true
}

