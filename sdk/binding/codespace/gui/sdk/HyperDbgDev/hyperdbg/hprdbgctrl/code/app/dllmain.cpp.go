package app

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\app\dllmain.cpp.back

type (
	Dllmain interface {
		DllMain() (ok bool) //col:12
	}
	dllmain struct{}
)

func NewDllmain() Dllmain { return &dllmain{} }

func (d *dllmain) DllMain() (ok bool) { //col:12
	/*DllMain(HMODULE hModule, DWORD ul_reason_for_call, LPVOID lpReserved)
	  {
	      switch (ul_reason_for_call)
	      {
	      case DLL_PROCESS_ATTACH:
	      case DLL_THREAD_ATTACH:
	      case DLL_THREAD_DETACH:
	      case DLL_PROCESS_DETACH:
	          break;
	      }
	      return TRUE;
	  }*/
	return true
}
