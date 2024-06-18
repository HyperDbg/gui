#include "pch.h"
extern BOOLEAN      g_IsSerialConnectedToRemoteDebugger;
extern std::wstring g_StartCommandPath;
extern std::wstring g_StartCommandPathAndArguments;
extern BOOLEAN      g_IsSerialConnectedToRemoteDebugger;
VOID CommandRevHelp(){
    ShowMessages("!rev : uses the reversing machine module in order to reconstruct the programmer/memory assumptions.\n\n");
    ShowMessages("syntax : \t!rev [config] [pid ProcessId (hex)]\n");
    ShowMessages("syntax : \t!rev [path Path (string)] [Parameters (string)]\n");
    ShowMessages("\n");
    ShowMessages("\t\te.g : !rev path c:\\reverse eng\\my_file.exe\n");
    ShowMessages("\t\te.g : !rev pattern\n");
    ShowMessages("\t\te.g : !rev reconstruct\n");
    ShowMessages("\t\te.g : !rev pattern pid 1c0\n");
    ShowMessages("\t\te.g : !rev reconstruct pid 1c0\n");
}
VOID CommandRev(vector<string> SplitCommand, string Command){
    REVERSING_MACHINE_RECONSTRUCT_MEMORY_REQUEST RevRequest         = {0};
    BOOLEAN                                      SetPid             = FALSE;
    UINT32                                       TargetPid          = NULL;
    BOOLEAN                                      IgnoreFirstCommand = TRUE;
    REVERSING_MACHINE_RECONSTRUCT_MEMORY_MODE    Mode               = REVERSING_MACHINE_RECONSTRUCT_MEMORY_MODE_UNKNOWN;
    REVERSING_MACHINE_RECONSTRUCT_MEMORY_TYPE    Type               = REVERSING_MACHINE_RECONSTRUCT_MEMORY_TYPE_UNKNOWN;
    for (auto Section : SplitCommand){
        if (!Section.compare("!rev") && IgnoreFirstCommand){
            IgnoreFirstCommand = FALSE;
            continue;
        }
        else if (!Section.compare("pid") && !SetPid){
            SetPid = TRUE;
        }
        else if (SetPid){
            if (!ConvertStringToUInt32(Section, &TargetPid)){
                ShowMessages("err, couldn't resolve error at '%s'\n\n",
                             Section.c_str());
                CommandRevHelp();
                return;
            }
            SetPid = FALSE;
        }
        else if (!Section.compare("pattern")){
            Type = REVERSING_MACHINE_RECONSTRUCT_MEMORY_TYPE_PATTERN;
        }
        else if (!Section.compare("reconstruct")){
            Type = REVERSING_MACHINE_RECONSTRUCT_MEMORY_TYPE_RECONSTRUCT;
        }
        else{
            ShowMessages("err, couldn't resolve error at '%s'\n\n",
                         Section.c_str());
            CommandRevHelp();
            return;
        }
    }
    if (SetPid){
        ShowMessages("err, please enter a valid process id in hex format, "
                     "or if you want to use it in decimal format, add '0n' "
                     "prefix to the number\n");
        return;
    }
    RevRequest.ProcessId = TargetPid;
    RevRequestService(&RevRequest);
}
