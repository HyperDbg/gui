package Source


type (
PdbCallback interface{
		ULONG STDMETHODCALLTYPE AddRef()(ok bool)//col:120
}

)

func NewPdbCallback() { return & pdbCallback{} }

func (p *pdbCallback)		ULONG STDMETHODCALLTYPE AddRef()(ok bool){//col:120






















































































return true
}




