package Source
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\dependencies\pdbex\Source\PDBCallback.h.back

type (
PdbCallback interface{
		ULONG_STDMETHODCALLTYPE_AddRef()(ok bool)//col:87
		ULONG_STDMETHODCALLTYPE_Release()(ok bool)//col:170
			if_()(ok bool)//col:251
		HRESULT_STDMETHODCALLTYPE_QueryInterface()(ok bool)//col:326
			if_()(ok bool)//col:399
			if_()(ok bool)//col:468
				PtrInterface_=_()(ok bool)//col:535
			else_if_()(ok bool)//col:600
				PtrInterface_=_()(ok bool)//col:663
			else_if_()(ok bool)//col:724
				PtrInterface_=_()(ok bool)//col:783
			if_()(ok bool)//col:836
				AddRef()(ok bool)//col:887
		HRESULT_STDMETHODCALLTYPE_NotifyDebugDir()(ok bool)//col:933
		HRESULT_STDMETHODCALLTYPE_NotifyOpenDBG()(ok bool)//col:972
		HRESULT_STDMETHODCALLTYPE_NotifyOpenPDB()(ok bool)//col:1005
		HRESULT_STDMETHODCALLTYPE_RestrictRegistryAccess()(ok bool)//col:1032
		HRESULT_STDMETHODCALLTYPE_RestrictSymbolServerAccess()(ok bool)//col:1055
		HRESULT_STDMETHODCALLTYPE_RestrictOriginalPathAccess()(ok bool)//col:1074
		HRESULT_STDMETHODCALLTYPE_RestrictReferencePathAccess()(ok bool)//col:1089
		HRESULT_STDMETHODCALLTYPE_RestrictDBGAccess()(ok bool)//col:1100
		HRESULT_STDMETHODCALLTYPE_RestrictSystemRootAccess()(ok bool)//col:1107
}
pdbCallback struct{}
)

func NewPdbCallback()PdbCallback{ return & pdbCallback{} }

func (p *pdbCallback)		ULONG_STDMETHODCALLTYPE_AddRef()(ok bool){//col:87
/*		ULONG STDMETHODCALLTYPE AddRef() override
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
				return S_OK;
			}
			return E_NOINTERFACE;
		}
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
};*/
return true
}

func (p *pdbCallback)		ULONG_STDMETHODCALLTYPE_Release()(ok bool){//col:170
/*		ULONG STDMETHODCALLTYPE Release() override
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
				return S_OK;
			}
			return E_NOINTERFACE;
		}
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
};*/
return true
}

func (p *pdbCallback)			if_()(ok bool){//col:251
/*			if ((--m_RefCount) == 0)
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
				return S_OK;
			}
			return E_NOINTERFACE;
		}
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
};*/
return true
}

func (p *pdbCallback)		HRESULT_STDMETHODCALLTYPE_QueryInterface()(ok bool){//col:326
/*		HRESULT STDMETHODCALLTYPE QueryInterface(REFIID Rid, void **Interface) override
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
				return S_OK;
			}
			return E_NOINTERFACE;
		}
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
};*/
return true
}

func (p *pdbCallback)			if_()(ok bool){//col:399
/*			if (Interface == nullptr)
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
				return S_OK;
			}
			return E_NOINTERFACE;
		}
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
};*/
return true
}

func (p *pdbCallback)			if_()(ok bool){//col:468
/*			if (Rid == __uuidof(IDiaLoadCallback2))
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
				return S_OK;
			}
			return E_NOINTERFACE;
		}
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
};*/
return true
}

func (p *pdbCallback)				PtrInterface_=_()(ok bool){//col:535
/*				*Interface = (IDiaLoadCallback2 *)this;
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
				return S_OK;
			}
			return E_NOINTERFACE;
		}
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
};*/
return true
}

func (p *pdbCallback)			else_if_()(ok bool){//col:600
/*			else if (Rid == __uuidof(IDiaLoadCallback))
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
				return S_OK;
			}
			return E_NOINTERFACE;
		}
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
};*/
return true
}

func (p *pdbCallback)				PtrInterface_=_()(ok bool){//col:663
/*				*Interface = (IDiaLoadCallback *)this;
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
				return S_OK;
			}
			return E_NOINTERFACE;
		}
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
};*/
return true
}

func (p *pdbCallback)			else_if_()(ok bool){//col:724
/*			else if (Rid == __uuidof(IUnknown))
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
				return S_OK;
			}
			return E_NOINTERFACE;
		}
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
};*/
return true
}

func (p *pdbCallback)				PtrInterface_=_()(ok bool){//col:783
/*				*Interface = (IUnknown *)this;
			}
			else
			{
				*Interface = nullptr;
			}
			if (*Interface != nullptr)
			{
				AddRef();
				return S_OK;
			}
			return E_NOINTERFACE;
		}
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
};*/
return true
}

func (p *pdbCallback)			if_()(ok bool){//col:836
/*			if (*Interface != nullptr)
			{
				AddRef();
				return S_OK;
			}
			return E_NOINTERFACE;
		}
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
};*/
return true
}

func (p *pdbCallback)				AddRef()(ok bool){//col:887
/*				AddRef();
				return S_OK;
			}
			return E_NOINTERFACE;
		}
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
};*/
return true
}

func (p *pdbCallback)		HRESULT_STDMETHODCALLTYPE_NotifyDebugDir()(ok bool){//col:933
/*		HRESULT STDMETHODCALLTYPE NotifyDebugDir(
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
};*/
return true
}

func (p *pdbCallback)		HRESULT_STDMETHODCALLTYPE_NotifyOpenDBG()(ok bool){//col:972
/*		HRESULT STDMETHODCALLTYPE NotifyOpenDBG(
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
};*/
return true
}

func (p *pdbCallback)		HRESULT_STDMETHODCALLTYPE_NotifyOpenPDB()(ok bool){//col:1005
/*		HRESULT STDMETHODCALLTYPE NotifyOpenPDB(
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
};*/
return true
}

func (p *pdbCallback)		HRESULT_STDMETHODCALLTYPE_RestrictRegistryAccess()(ok bool){//col:1032
/*		HRESULT STDMETHODCALLTYPE RestrictRegistryAccess() override
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
};*/
return true
}

func (p *pdbCallback)		HRESULT_STDMETHODCALLTYPE_RestrictSymbolServerAccess()(ok bool){//col:1055
/*		HRESULT STDMETHODCALLTYPE RestrictSymbolServerAccess() override
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
};*/
return true
}

func (p *pdbCallback)		HRESULT_STDMETHODCALLTYPE_RestrictOriginalPathAccess()(ok bool){//col:1074
/*		HRESULT STDMETHODCALLTYPE RestrictOriginalPathAccess() override
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
};*/
return true
}

func (p *pdbCallback)		HRESULT_STDMETHODCALLTYPE_RestrictReferencePathAccess()(ok bool){//col:1089
/*		HRESULT STDMETHODCALLTYPE RestrictReferencePathAccess() override
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
};*/
return true
}

func (p *pdbCallback)		HRESULT_STDMETHODCALLTYPE_RestrictDBGAccess()(ok bool){//col:1100
/*		HRESULT STDMETHODCALLTYPE RestrictDBGAccess() override
		{
			return S_OK;
		}
		HRESULT STDMETHODCALLTYPE RestrictSystemRootAccess() override
		{
			return S_OK;
		}
	private:
		volatile unsigned long m_RefCount = 0;
};*/
return true
}

func (p *pdbCallback)		HRESULT_STDMETHODCALLTYPE_RestrictSystemRootAccess()(ok bool){//col:1107
/*		HRESULT STDMETHODCALLTYPE RestrictSystemRootAccess() override
		{
			return S_OK;
		}
	private:
		volatile unsigned long m_RefCount = 0;
};*/
return true
}



