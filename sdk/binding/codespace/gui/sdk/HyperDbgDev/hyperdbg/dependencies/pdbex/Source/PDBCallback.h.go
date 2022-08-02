package Source

type (
	PdbCallback interface {
		ULONG_STDMETHODCALLTYPE_AddRef() (ok bool)                        //col:83
		HRESULT_STDMETHODCALLTYPE_NotifyOpenDBG() (ok bool)               //col:122
		HRESULT_STDMETHODCALLTYPE_NotifyOpenPDB() (ok bool)               //col:155
		HRESULT_STDMETHODCALLTYPE_RestrictRegistryAccess() (ok bool)      //col:182
		HRESULT_STDMETHODCALLTYPE_RestrictSymbolServerAccess() (ok bool)  //col:205
		HRESULT_STDMETHODCALLTYPE_RestrictOriginalPathAccess() (ok bool)  //col:224
		HRESULT_STDMETHODCALLTYPE_RestrictReferencePathAccess() (ok bool) //col:239
		HRESULT_STDMETHODCALLTYPE_RestrictDBGAccess() (ok bool)           //col:250
		HRESULT_STDMETHODCALLTYPE_RestrictSystemRootAccess() (ok bool)    //col:257
	}
	pdbCallback struct{}
)

func NewPdbCallback() PdbCallback { return &pdbCallback{} }

func (p *pdbCallback) ULONG_STDMETHODCALLTYPE_AddRef() (ok bool) { //col:83

	return true
}

func (p *pdbCallback) HRESULT_STDMETHODCALLTYPE_NotifyOpenDBG() (ok bool) { //col:122

	return true
}

func (p *pdbCallback) HRESULT_STDMETHODCALLTYPE_NotifyOpenPDB() (ok bool) { //col:155

	return true
}

func (p *pdbCallback) HRESULT_STDMETHODCALLTYPE_RestrictRegistryAccess() (ok bool) { //col:182

	return true
}

func (p *pdbCallback) HRESULT_STDMETHODCALLTYPE_RestrictSymbolServerAccess() (ok bool) { //col:205

	return true
}

func (p *pdbCallback) HRESULT_STDMETHODCALLTYPE_RestrictOriginalPathAccess() (ok bool) { //col:224

	return true
}

func (p *pdbCallback) HRESULT_STDMETHODCALLTYPE_RestrictReferencePathAccess() (ok bool) { //col:239

	return true
}

func (p *pdbCallback) HRESULT_STDMETHODCALLTYPE_RestrictDBGAccess() (ok bool) { //col:250

	return true
}

func (p *pdbCallback) HRESULT_STDMETHODCALLTYPE_RestrictSystemRootAccess() (ok bool) { //col:257

	return true
}
