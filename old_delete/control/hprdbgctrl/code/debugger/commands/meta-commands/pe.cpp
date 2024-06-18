#include "pch.h"
using namespace std;
VOID CommandPeHelp(){
    ShowMessages(".pe : parses portable executable (PE) files and dump sections.\n\n");
    ShowMessages("syntax : \t.pe [header] [FilePath (string)]\n");
    ShowMessages("syntax : \t.pe [section] [SectionName (string)] [FilePath (string)]\n");
    ShowMessages("\n");
    ShowMessages("\t\te.g : .pe header c:\\reverse files\\myfile.exe\n");
    ShowMessages("\t\te.g : .pe section .text c:\\reverse files\\myfile.exe\n");
    ShowMessages("\t\te.g : .pe section .rdata c:\\reverse files\\myfile.exe\n");
}
VOID CommandPe(vector<string> SplitCommand, string Command){
    BOOLEAN Is32Bit = FALSE;
    wstring Filepath;
    BOOLEAN ShowDumpOfSection = FALSE;
    if (SplitCommand.size() <= 2){
        ShowMessages("err, incorrect use of the '.pe' command\n\n");
        CommandPeHelp();
        return;
    }
    if (!SplitCommand.at(1).compare("section")){
        if (SplitCommand.size() == 3){
            ShowMessages("please specify a valid PE file\n\n");
            CommandPeHelp();
            return;
        }
        ShowDumpOfSection = TRUE;
    }
    else if (!SplitCommand.at(1).compare("header")){
        ShowDumpOfSection = FALSE;
    }
    else{
        ShowMessages("err, couldn't resolve error at '%s'\n\n",
                     SplitCommand.at(1).c_str());
        CommandPeHelp();
        return;
    }
    Trim(Command);
    Command.erase(0, SplitCommand.at(0).size());
    if (!ShowDumpOfSection){
        Command.erase(0, 6 + 1);
    }
    else{
        Command.erase(0, 7 + 1);
        Command.erase(0, SplitCommand.at(2).size() + 1);
    }
    Trim(Command);
    StringToWString(Filepath, Command);
    if (!PeIsPE32BitOr64Bit(Filepath.c_str(), &Is32Bit)){
        return;
    }
    if (!ShowDumpOfSection){
        PeShowSectionInformationAndDump(Filepath.c_str(), NULL, Is32Bit);
    }
    else{
        PeShowSectionInformationAndDump(Filepath.c_str(), SplitCommand.at(2).c_str(), Is32Bit);
    }
}
