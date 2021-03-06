#pragma once
extern "C" {
extern bool inline AsmVmxSupportDetection();
__declspec(dllexport) int HyperDbgLoadVmm();
__declspec(dllexport) int HyperDbgUnloadVmm();
__declspec(dllexport) int HyperDbgInstallVmmDriver();
__declspec(dllexport) int HyperDbgUninstallVmmDriver();
__declspec(dllexport) int HyperDbgStopVmmDriver();
__declspec(dllexport) int HyperDbgInterpreter(char * Command);
__declspec(dllexport) void HyperDbgShowSignature();
__declspec(dllexport) void HyperdbgSetTextMessageCallback(Callback handler);
__declspec(dllexport) void HyperDbgScriptReadFileAndExecuteCommand(std::vector<std::string> & PathAndArgs);
__declspec(dllexport) bool HyperDbgContinuePreviousCommand();
__declspec(dllexport) bool HyperDbgCheckMultilineCommand(std::string & CurrentCommand, bool Reset);
}
