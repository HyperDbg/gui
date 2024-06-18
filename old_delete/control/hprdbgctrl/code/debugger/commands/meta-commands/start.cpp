#include "pch.h"
extern std::wstring g_StartCommandPath;
extern std::wstring g_StartCommandPathAndArguments;
extern BOOLEAN      g_IsSerialConnectedToRemoteDebugger;
VOID CommandStartHelp(){
    ShowMessages(".start : runs a user-mode process.\n\n");
    ShowMessages("syntax : \t.start [path Path (string)] [Parameters (string)]\n");
    ShowMessages("\n");
    ShowMessages("\t\te.g : .start path c:\\reverse eng\\my_file.exe\n");
}
VOID CommandStart(vector<string> SplitCommand, string Command){
    vector<string> PathAndArgs;
    string         Arguments = "";
#if ActivateUserModeDebugger == FALSE
    if (!g_IsSerialConnectedToRemoteDebugger){
        ShowMessages("the user-mode debugger in VMI Mode is still in the beta version and not stable. "
                     "we decided to exclude it from this release and release it in future versions. "
                     "if you want to test the user-mode debugger in VMI Mode, you should build "
                     "HyperDbg with special instructions. But starting processes is fully supported "
                     "in the Debugger Mode.\n"
                     "(it's not recommended to use it in VMI Mode yet!)\n");
        return;
    }
#endif 
    if (SplitCommand.size() <= 2){
        ShowMessages("incorrect use of the '.start'\n\n");
        CommandStartHelp();
        return;
    }
    if (!SplitCommand.at(1).compare("path")){
        Trim(Command);
        Command.erase(0, SplitCommand.at(0).size());
        Command.erase(0, 4 + 1);
        Trim(Command);
        SplitPathAndArgs(PathAndArgs, Command);
        StringToWString(g_StartCommandPath, PathAndArgs.at(0));
        if (PathAndArgs.size() != 1){
            for (auto item : PathAndArgs){
                Arguments += item + " ";
            }
            Arguments.pop_back();
            StringToWString(g_StartCommandPathAndArguments, Arguments);
        }
    }
    else{
        ShowMessages("err, couldn't resolve error at '%s'\n\n",
                     SplitCommand.at(1).c_str());
        CommandStartHelp();
        return;
    }
    if (Arguments.empty()){
        UdAttachToProcess(NULL,
                          g_StartCommandPath.c_str(),
                          NULL,
                          FALSE);
    }
    else{
        UdAttachToProcess(NULL,
                          g_StartCommandPath.c_str(),
                          (WCHAR *)g_StartCommandPathAndArguments.c_str(),
                          FALSE);
    }
}
