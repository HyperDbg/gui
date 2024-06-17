#include "pch.h"
extern BOOLEAN                  g_IsSerialConnectedToRemoteDebuggee;
extern ACTIVE_DEBUGGING_PROCESS g_ActiveProcessDebuggingState;
VOID CommandSymHelp(){
    ShowMessages(".sym : performs the symbol actions.\n\n");
    ShowMessages("syntax : \t.sym [table]\n");
    ShowMessages("syntax : \t.sym [reload] [pid ProcessId (hex)]\n");
    ShowMessages("syntax : \t.sym [download]\n");
    ShowMessages("syntax : \t.sym [load]\n");
    ShowMessages("syntax : \t.sym [unload]\n");
    ShowMessages("syntax : \t.sym [add] [base Address (hex)] [path Path (string)]\n");
    ShowMessages("\n");
    ShowMessages("\t\te.g : .sym table\n");
    ShowMessages("\t\te.g : .sym reload\n");
    ShowMessages("\t\te.g : .sym load\n");
    ShowMessages("\t\te.g : .sym download\n");
    ShowMessages("\t\te.g : .sym add base fffff8077356000 path c:\\symbols\\my_dll.pdb\n");
    ShowMessages("\t\te.g : .sym unload\n");
}
VOID CommandSym(vector<string> SplitCommand, string Command){
    UINT64 BaseAddress   = NULL;
    UINT32 UserProcessId = NULL;
    if (SplitCommand.size() == 1){
        ShowMessages("incorrect use of the '.sym'\n\n");
        CommandSymHelp();
        return;
    }
    if (!SplitCommand.at(1).compare("table")){
        if (SplitCommand.size() != 2){
            ShowMessages("incorrect use of the '.sym'\n\n");
            CommandSymHelp();
            return;
        }
        SymbolBuildAndShowSymbolTable();
    }
    else if (!SplitCommand.at(1).compare("load") || !SplitCommand.at(1).compare("download")){
        if (SplitCommand.size() != 2){
            ShowMessages("incorrect use of the '.sym'\n\n");
            CommandSymHelp();
            return;
        }
        if (!SplitCommand.at(1).compare("load")){
            SymbolLoadOrDownloadSymbols(FALSE, FALSE);
        }
        else if (!SplitCommand.at(1).compare("download")){
            SymbolLoadOrDownloadSymbols(TRUE, FALSE);
        }
    }
    else if (!SplitCommand.at(1).compare("reload")){
        if (SplitCommand.size() != 2 && SplitCommand.size() != 4){
            ShowMessages("incorrect use of the '.sym'\n\n");
            CommandSymHelp();
            return;
        }
        if (SplitCommand.size() == 4){
            if (!SplitCommand.at(2).compare("pid")){
                if (!ConvertStringToUInt32(SplitCommand.at(3), &UserProcessId)){
                    ShowMessages("err, couldn't resolve error at '%s'\n\n",
                                 SplitCommand.at(3).c_str());
                    CommandSymHelp();
                    return;
                }
            }
            else{
                ShowMessages("incorrect use of the '.sym'\n\n");
                CommandSymHelp();
                return;
            }
        }
        if (g_IsSerialConnectedToRemoteDebuggee){
            SymbolReloadSymbolTableInDebuggerMode(UserProcessId);
        }
        else{
            if (UserProcessId == NULL){
                if (g_ActiveProcessDebuggingState.IsActive){
                    UserProcessId = g_ActiveProcessDebuggingState.ProcessId;
                }
                else{
                    UserProcessId = GetCurrentProcessId();
                }
            }
            if (SymbolLocalReload(UserProcessId)){
                ShowMessages("symbol table updated successfully\n");
            }
        }
    }
    else if (!SplitCommand.at(1).compare("unload")){
        if (SplitCommand.size() != 2){
            ShowMessages("incorrect use of the '.sym'\n\n");
            CommandSymHelp();
            return;
        }
        ScriptEngineUnloadAllSymbolsWrapper();
    }
    else if (!SplitCommand.at(1).compare("add")){
        if (SplitCommand.size() < 6){
            ShowMessages("incorrect use of the '.sym'\n\n");
            CommandSymHelp();
            return;
        }
        if (!SplitCommand.at(2).compare("base")){
            string Delimiter = "";
            string PathToPdb = "";
            if (!ConvertStringToUInt64(SplitCommand.at(3), &BaseAddress)){
                ShowMessages("please add a valid hex address to be used as the base address\n\n");
                CommandSymHelp();
                return;
            }
            if (SplitCommand.at(4).compare("path")){
                ShowMessages("incorrect use of the '.sym'\n\n");
                CommandSymHelp();
                return;
            }
            Delimiter = "path ";
            PathToPdb = Command.substr(Command.find(Delimiter) + 5, Command.size());
            if (!IsFileExistA(PathToPdb.c_str())){
                ShowMessages("pdb file not found\n");
                return;
            }
            ShowMessages("loading module symbol at '%s'\n", PathToPdb.c_str());
            ScriptEngineLoadFileSymbolWrapper(BaseAddress, PathToPdb.c_str(), NULL);
        }
        else{
            ShowMessages("incorrect use of the '.sym'\n\n");
            CommandSymHelp();
            return;
        }
    }
    else{
        ShowMessages("unknown parameter at '%s'\n\n", SplitCommand.at(1).c_str());
        CommandSymHelp();
        return;
    }
}
