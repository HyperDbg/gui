package Source

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\pdbex\Source\PDBCallback.h.back

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
	/*
				ULONG STDMETHODCALLTYPE AddRef() override
				{
					return m_RefCount++;
				}
				ULONG STDMETHODCALLTYPE Release() override
				{
					if ((--m_RefCount) == 0)
					{
						delete this;
					}
					return m_RefCount;
				}
				HRESULT STDMETHODCALLTYPE QueryInterface(REFIID Rid, void **Interface) override
				{
					if (Interface == nullptr)
					{
						return E_INVALIDARG;
					}
					if (Rid == __uuidof(IDiaLoadCallback2))
					{
						*Interface = (IDiaLoadCallback2 *)this;
					}
					else if (Rid == __uuidof(IDiaLoadCallback))
					{
						*Interface = (IDiaLoadCallback *)this;
					}
					else if (Rid == __uuidof(IUnknown))
					{
						*Interface = (IUnknown *)this;
					}
					else
					{
						*Interface = nullptr;
					}
					if (*Interface != nullptr)
					{
						AddRef();
				HRESULT STDMETHODCALLTYPE NotifyDebugDir(
					BOOL fExecutable,
					DWORD cbData,
					BYTE data[]) override
				{
					return S_OK;
				}
				HRESULT STDMETHODCALLTYPE NotifyOpenDBG(
					LPCOLESTR dbgPath,
					HRESULT resultCode) override
				{
					return S_OK;
				}
				HRESULT STDMETHODCALLTYPE NotifyOpenPDB(
					LPCOLESTR pdbPath,
					HRESULT resultCode) override
				{
					return S_OK;
				}
				HRESULT STDMETHODCALLTYPE RestrictRegistryAccess() override
				{
					return S_OK;
				}
				HRESULT STDMETHODCALLTYPE RestrictSymbolServerAccess() override
				{
					return S_OK;
				}
				HRESULT STDMETHODCALLTYPE RestrictOriginalPathAccess() override
				{
					return S_OK;
				}
				HRESULT STDMETHODCALLTYPE RestrictReferencePathAccess() override
				{
					return S_OK;
				}
				HRESULT STDMETHODCALLTYPE RestrictDBGAccess() override
				{
					return S_OK;
				}
				HRESULT STDMETHODCALLTYPE RestrictSystemRootAccess() override
				{
					return S_OK;
				}
			private:
				volatile unsigned long m_RefCount = 0;
		};
	*/
	return true
}

func (p *pdbCallback) HRESULT_STDMETHODCALLTYPE_NotifyOpenDBG() (ok bool) { //col:122
	/*
				HRESULT STDMETHODCALLTYPE NotifyOpenDBG(
					LPCOLESTR dbgPath,
					HRESULT resultCode) override
				{
					return S_OK;
				}
				HRESULT STDMETHODCALLTYPE NotifyOpenPDB(
					LPCOLESTR pdbPath,
					HRESULT resultCode) override
				{
					return S_OK;
				}
				HRESULT STDMETHODCALLTYPE RestrictRegistryAccess() override
				{
					return S_OK;
				}
				HRESULT STDMETHODCALLTYPE RestrictSymbolServerAccess() override
				{
					return S_OK;
				}
				HRESULT STDMETHODCALLTYPE RestrictOriginalPathAccess() override
				{
					return S_OK;
				}
				HRESULT STDMETHODCALLTYPE RestrictReferencePathAccess() override
				{
					return S_OK;
				}
				HRESULT STDMETHODCALLTYPE RestrictDBGAccess() override
				{
					return S_OK;
				}
				HRESULT STDMETHODCALLTYPE RestrictSystemRootAccess() override
				{
					return S_OK;
				}
			private:
				volatile unsigned long m_RefCount = 0;
		};
	*/
	return true
}

func (p *pdbCallback) HRESULT_STDMETHODCALLTYPE_NotifyOpenPDB() (ok bool) { //col:155
	/*
				HRESULT STDMETHODCALLTYPE NotifyOpenPDB(
					LPCOLESTR pdbPath,
					HRESULT resultCode) override
				{
					return S_OK;
				}
				HRESULT STDMETHODCALLTYPE RestrictRegistryAccess() override
				{
					return S_OK;
				}
				HRESULT STDMETHODCALLTYPE RestrictSymbolServerAccess() override
				{
					return S_OK;
				}
				HRESULT STDMETHODCALLTYPE RestrictOriginalPathAccess() override
				{
					return S_OK;
				}
				HRESULT STDMETHODCALLTYPE RestrictReferencePathAccess() override
				{
					return S_OK;
				}
				HRESULT STDMETHODCALLTYPE RestrictDBGAccess() override
				{
					return S_OK;
				}
				HRESULT STDMETHODCALLTYPE RestrictSystemRootAccess() override
				{
					return S_OK;
				}
			private:
				volatile unsigned long m_RefCount = 0;
		};
	*/
	return true
}

func (p *pdbCallback) HRESULT_STDMETHODCALLTYPE_RestrictRegistryAccess() (ok bool) { //col:182
	/*
				HRESULT STDMETHODCALLTYPE RestrictRegistryAccess() override
				{
					return S_OK;
				}
				HRESULT STDMETHODCALLTYPE RestrictSymbolServerAccess() override
				{
					return S_OK;
				}
				HRESULT STDMETHODCALLTYPE RestrictOriginalPathAccess() override
				{
					return S_OK;
				}
				HRESULT STDMETHODCALLTYPE RestrictReferencePathAccess() override
				{
					return S_OK;
				}
				HRESULT STDMETHODCALLTYPE RestrictDBGAccess() override
				{
					return S_OK;
				}
				HRESULT STDMETHODCALLTYPE RestrictSystemRootAccess() override
				{
					return S_OK;
				}
			private:
				volatile unsigned long m_RefCount = 0;
		};
	*/
	return true
}

func (p *pdbCallback) HRESULT_STDMETHODCALLTYPE_RestrictSymbolServerAccess() (ok bool) { //col:205
	/*
				HRESULT STDMETHODCALLTYPE RestrictSymbolServerAccess() override
				{
					return S_OK;
				}
				HRESULT STDMETHODCALLTYPE RestrictOriginalPathAccess() override
				{
					return S_OK;
				}
				HRESULT STDMETHODCALLTYPE RestrictReferencePathAccess() override
				{
					return S_OK;
				}
				HRESULT STDMETHODCALLTYPE RestrictDBGAccess() override
				{
					return S_OK;
				}
				HRESULT STDMETHODCALLTYPE RestrictSystemRootAccess() override
				{
					return S_OK;
				}
			private:
				volatile unsigned long m_RefCount = 0;
		};
	*/
	return true
}

func (p *pdbCallback) HRESULT_STDMETHODCALLTYPE_RestrictOriginalPathAccess() (ok bool) { //col:224
	/*
				HRESULT STDMETHODCALLTYPE RestrictOriginalPathAccess() override
				{
					return S_OK;
				}
				HRESULT STDMETHODCALLTYPE RestrictReferencePathAccess() override
				{
					return S_OK;
				}
				HRESULT STDMETHODCALLTYPE RestrictDBGAccess() override
				{
					return S_OK;
				}
				HRESULT STDMETHODCALLTYPE RestrictSystemRootAccess() override
				{
					return S_OK;
				}
			private:
				volatile unsigned long m_RefCount = 0;
		};
	*/
	return true
}

func (p *pdbCallback) HRESULT_STDMETHODCALLTYPE_RestrictReferencePathAccess() (ok bool) { //col:239
	/*
				HRESULT STDMETHODCALLTYPE RestrictReferencePathAccess() override
				{
					return S_OK;
				}
				HRESULT STDMETHODCALLTYPE RestrictDBGAccess() override
				{
					return S_OK;
				}
				HRESULT STDMETHODCALLTYPE RestrictSystemRootAccess() override
				{
					return S_OK;
				}
			private:
				volatile unsigned long m_RefCount = 0;
		};
	*/
	return true
}

func (p *pdbCallback) HRESULT_STDMETHODCALLTYPE_RestrictDBGAccess() (ok bool) { //col:250
	/*
				HRESULT STDMETHODCALLTYPE RestrictDBGAccess() override
				{
					return S_OK;
				}
				HRESULT STDMETHODCALLTYPE RestrictSystemRootAccess() override
				{
					return S_OK;
				}
			private:
				volatile unsigned long m_RefCount = 0;
		};
	*/
	return true
}

func (p *pdbCallback) HRESULT_STDMETHODCALLTYPE_RestrictSystemRootAccess() (ok bool) { //col:257
	/*
				HRESULT STDMETHODCALLTYPE RestrictSystemRootAccess() override
				{
					return S_OK;
				}
			private:
				volatile unsigned long m_RefCount = 0;
		};
	*/
	return true
}

