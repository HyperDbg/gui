#include <Windows.h>
#include <string>
#include <conio.h>
#include <iostream>
#include <vector>
#include "SDK/HyperDbgSdk.h"
#include "SDK/Imports/HyperDbgCtrlImports.h"
using namespace std;
#pragma comment(lib, "HPRDBGCTRL.lib")
int
main(int argc, char * argv[]) {
    bool           ExitFromDebugger = false;
    string         PreviousCommand;
    bool           Reset = false;
    vector<string> Args;
    printf("HyperDbg Debugger [version: %s, build: %s]\n", CompleteVersion, BuildVersion);
    printf("Please visit https://docs.hyperdbg.org for more information...\n");
    printf("HyperDbg is released under the GNU Public License v3 (GPLv3).\n\n");
    if (argc != 1) {
        if (!strcmp(argv[1], "--script")) {
            for (size_t i = 2; i < argc; i++) {
                std::string TempStr(argv[i]);
                Args.push_back(TempStr);
            }
            if (!Args.empty()) {
                HyperDbgScriptReadFileAndExecuteCommand(Args);
                printf("\n");
            } else {
                printf("err, invalid command line options passed to the HyperDbg!\n");
                return 1;
            }
        } else {
            printf("err, invalid command line options passed to the HyperDbg!\n");
            return 1;
        }
    }
    while (!ExitFromDebugger) {
        HyperDbgShowSignature();
        string CurrentCommand = "";
        Reset                 = true;
    GetMultiLinecCommand:
        string TempCommand = "";
        getline(cin, TempCommand);
        if (cin.fail() || cin.eof()) {
            cin.clear(); // reset cin state
            printf("\n\n");
            continue;
        }
        if (HyperDbgCheckMultilineCommand(TempCommand, Reset) == true) {
            Reset = false;
            CurrentCommand += TempCommand + "\n";
            printf("> ");
            goto GetMultiLinecCommand;
        } else {
            Reset = true;
            CurrentCommand += TempCommand;
        }
        if (!CurrentCommand.compare("") &&
            HyperDbgContinuePreviousCommand()) {
            CurrentCommand = PreviousCommand;
        } else {
            PreviousCommand = CurrentCommand;
        }
        int CommandExecutionResult = HyperDbgInterpreter((char *)CurrentCommand.c_str());
        if (CommandExecutionResult == 1) {
            ExitFromDebugger = true;
        }
        if (CommandExecutionResult != 2) {
            printf("\n");
        }
    }
    return 0;
}
