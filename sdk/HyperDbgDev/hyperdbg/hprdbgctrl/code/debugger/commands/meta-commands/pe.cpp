#include "pch.h"
using namespace std;
VOID
CommandPeHelp() {
    ShowMessages(".pe : parses portable executable (PE) files and dump sections.\n\n");
    ShowMessages("syntax : \t.pe [header] [FilePath (string)]\n");
    ShowMessages("syntax : \t.pe [section] [SectionName (string)] [FilePath (string)]\n");
    ShowMessages("\n");
    ShowMessages("\t\te.g : .pe header c:\\reverse files\\myfile.exe\n");
    ShowMessages("\t\te.g : .pe section .text c:\\reverse files\\myfile.exe\n");
    ShowMessages("\t\te.g : .pe section .rdata c:\\reverse files\\myfile.exe\n");
}

VOID
CommandPe(vector<string> SplittedCommand, string Command) {
    BOOLEAN Is32Bit = FALSE;
    wstring Filepath;
    BOOLEAN ShowDumpOfSection = FALSE;
    if (SplittedCommand.size() <= 2) {
        ShowMessages("err, incorrect use of '.pe' command\n\n");
        CommandPeHelp();
        return;
    }
    if (!SplittedCommand.at(1).compare("section")) {
        if (SplittedCommand.size() == 3) {
            ShowMessages("please specify a valid PE file\n\n");
            CommandPeHelp();
            return;
        }
        ShowDumpOfSection = TRUE;
    } else if (!SplittedCommand.at(1).compare("header")) {
        ShowDumpOfSection = FALSE;
    } else {
        ShowMessages("err, couldn't resolve error at '%s'\n\n",
                     SplittedCommand.at(1).c_str());
        CommandPeHelp();
        return;
    }
    Trim(Command);
    Command.erase(0, 3);
    if (!ShowDumpOfSection) {
        Command.erase(0, 6 + 1);
    } else {
        Command.erase(0, 7 + 1);
        Command.erase(0, SplittedCommand.at(2).length() + 1);
    }
    Trim(Command);
    StringToWString(Filepath, Command);
    if (!PeIsPE32BitOr64Bit(Filepath.c_str(), &Is32Bit)) {
        return;
    }
    if (!ShowDumpOfSection) {
        PeShowSectionInformationAndDump(Filepath.c_str(), NULL, Is32Bit);
    } else {
        PeShowSectionInformationAndDump(Filepath.c_str(), SplittedCommand.at(2).c_str(), Is32Bit);
    }
}
