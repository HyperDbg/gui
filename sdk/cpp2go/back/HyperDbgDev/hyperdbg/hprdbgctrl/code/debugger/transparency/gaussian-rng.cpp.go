package transparency


type (
GaussianRng interface{
Median()(ok bool)//col:42
Average()(ok bool)//col:68
CalculateStandardDeviation()(ok bool)//col:90
MedianAbsoluteDeviationTest()(ok bool)//col:114
Randn()(ok bool)//col:151
GuassianGenerateRandom()(ok bool)//col:220
TestGaussianFromFile()(ok bool)//col:249
}






)

func NewGaussianRng() { return & gaussianRng{} }

func (g *gaussianRng)Median()(ok bool){//col:42



















return true
}







func (g *gaussianRng)Average()(ok bool){//col:68










return true
}







func (g *gaussianRng)CalculateStandardDeviation()(ok bool){//col:90









return true
}







func (g *gaussianRng)MedianAbsoluteDeviationTest()(ok bool){//col:114












return true
}







func (g *gaussianRng)Randn()(ok bool){//col:151






















return true
}







func (g *gaussianRng)GuassianGenerateRandom()(ok bool){//col:220































return true
}







func (g *gaussianRng)TestGaussianFromFile()(ok bool){//col:249


















return true
}









