package Imports
//back\HyperDbgDev\hyperdbg\include\SDK\Imports\HyperDbgCtrlImports.h.back

type (
HyperDbgCtrlImports interface{
__declspec()(ok bool)//col:31
}
)

func NewHyperDbgCtrlImports() { return & hyperDbgCtrlImports{} }

func (h *hyperDbgCtrlImports)__declspec()(ok bool){//col:31
/*__declspec(dllimport) int HyperDbgLoadVmm();
__declspec(dllimport) int HyperDbgUnloadVmm();
__declspec(dllimport) int HyperDbgInstallVmmDriver();
__declspec(dllimport) int HyperDbgUninstallVmmDriver();
__declspec(dllimport) int HyperDbgStopVmmDriver();
__declspec(dllimport) int HyperDbgInterpreter(char * Command);
__declspec(dllimport) void HyperDbgShowSignature();
__declspec(dllimport) void HyperDbgSetTextMessageCallback(Callback handler);
__declspec(dllimport) void HyperDbgScriptReadFileAndExecuteCommand(std::vector<std::string> & PathAndArgs);
__declspec(dllimport) bool HyperDbgContinuePreviousCommand();
__declspec(dllimport) bool HyperDbgCheckMultilineCommand(std::string & CurrentCommand, bool Reset);
#ifdef __cplusplus
}*/
return true
}



