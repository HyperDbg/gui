#include "pch.h"
VOID
CommandClearScreenHelp() {
    ShowMessages(".cls : clears the screen.\n\n");
    ShowMessages("syntax : \t.cls\n");
}

VOID
CommandClearScreen(vector<string> SplittedCommand, string Command) {
    system("cls");
}
