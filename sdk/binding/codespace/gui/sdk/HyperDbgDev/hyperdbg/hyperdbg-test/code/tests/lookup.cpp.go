package tests



type (
	Lookup interface {
		TestListenerForResultsThread() (ok bool) //col:78
	}
	lookup struct{}
)

func NewLookup() Lookup { return &lookup{} }

func (l *lookup) TestListenerForResultsThread() (ok bool) { //col:78



















































































	return true
}


