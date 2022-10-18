#pragma once
#include <dia2.h> 
class PDBCallback : public IDiaLoadCallback2 {
public:
  ULONG STDMETHODCALLTYPE AddRef() override { return m_RefCount++; }
  ULONG STDMETHODCALLTYPE Release() override {
    if ((--m_RefCount) == 0) {
      delete this;
    }
    return m_RefCount;
  }
  HRESULT STDMETHODCALLTYPE QueryInterface(REFIID Rid,
                                           void **Interface) override {
    if (Interface == nullptr) {
      return E_INVALIDARG;
    }
    if (Rid == __uuidof(IDiaLoadCallback2)) {
      *Interface = (IDiaLoadCallback2 *)this;
    } else if (Rid == __uuidof(IDiaLoadCallback)) {
      *Interface = (IDiaLoadCallback *)this;
    } else if (Rid == __uuidof(IUnknown)) {
      *Interface = (IUnknown *)this;
    } else {
      *Interface = nullptr;
    }
    if (*Interface != nullptr) {
      AddRef();
      return S_OK;
    }
    return E_NOINTERFACE;
  }
  HRESULT STDMETHODCALLTYPE NotifyDebugDir(
      BOOL fExecutable, DWORD cbData,
      BYTE data[]) override 
  {
    return S_OK;
  }
  HRESULT STDMETHODCALLTYPE NotifyOpenDBG(LPCOLESTR dbgPath,
                                          HRESULT resultCode) override {
    return S_OK;
  }
  HRESULT STDMETHODCALLTYPE NotifyOpenPDB(LPCOLESTR pdbPath,
                                          HRESULT resultCode) override {
    return S_OK;
  }
  HRESULT STDMETHODCALLTYPE RestrictRegistryAccess() override {
    return S_OK;
  }
  HRESULT STDMETHODCALLTYPE RestrictSymbolServerAccess() override {
    return S_OK;
  }
  HRESULT STDMETHODCALLTYPE RestrictOriginalPathAccess() override {
    return S_OK;
  }
  HRESULT STDMETHODCALLTYPE RestrictReferencePathAccess() override {
    return S_OK;
  }
  HRESULT STDMETHODCALLTYPE RestrictDBGAccess() override { return S_OK; }
  HRESULT STDMETHODCALLTYPE RestrictSystemRootAccess() override { return S_OK; }
private:
  volatile unsigned long m_RefCount = 0;
};
