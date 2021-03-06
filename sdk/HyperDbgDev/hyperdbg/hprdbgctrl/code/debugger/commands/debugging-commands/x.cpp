#include "pch.h"
VOID
CommandXHelp() {
    ShowMessages("x : searches and shows symbols (wildcard) and corresponding addresses.\n\n");
    ShowMessages("syntax : \tx [Module!Name (wildcard string)]\n");
    ShowMessages("\n");
    ShowMessages("\t\te.g : x nt!ExAllocatePoolWithTag \n");
    ShowMessages("\t\te.g : x nt!ExAllocatePool* \n");
    ShowMessages("\t\te.g : x nt!* \n");
}

VOID
CommandX(vector<string> SplittedCommand, string Command) {
    if (SplittedCommand.size() == 1) {
        ShowMessages("incorrect use of 'x'\n\n");
        CommandXHelp();
        return;
    }
    Trim(Command);
    Command.erase(0, 1);
    Trim(Command);
    ScriptEngineSearchSymbolForMaskWrapper(Command.c_str());
}
