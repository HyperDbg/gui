#include "pch.h"
extern BOOLEAN                  g_IsSerialConnectedToRemoteDebuggee;
extern ACTIVE_DEBUGGING_PROCESS g_ActiveProcessDebuggingState;
HANDLE DumpFileHandle;
VOID CommandDumpHelp(){
    ShowMessages(".dump & !dump : saves memory context into a file.\n\n");
    ShowMessages("syntax : \t.dump [FromAddress (hex)] [ToAddress (hex)] [pid ProcessId (hex)] [path Path (string)]\n");
    ShowMessages("\nIf you want to dump physical memory then add '!' at the "
                 "start of the command\n\n");
    ShowMessages("\n");
    ShowMessages("\t\te.g : .dump 401000 40b000 path c:\\rev\\dump1.dmp\n");
    ShowMessages("\t\te.g : .dump 401000 40b000 pid 1c0 path c:\\rev\\desktop\\dump2.dmp\n");
    ShowMessages("\t\te.g : .dump fffff801deadb000 fffff801deade054 path c:\\rev\\dump3.dmp\n");
    ShowMessages("\t\te.g : .dump fffff801deadb000 fffff801deade054 path c:\\rev\\dump4.dmp\n");
    ShowMessages("\t\te.g : .dump 00007ff8349f2000 00007ff8349f8000 path c:\\rev\\dump5.dmp\n");
    ShowMessages("\t\te.g : .dump @rax+@rcx @rax+@rcx+1000 path c:\\rev\\dump6.dmp\n");
    ShowMessages("\t\te.g : !dump 1000 2100 path c:\\rev\\dump7.dmp\n");
}
VOID CommandDump(vector<string> SplitCommand, string Command){
    wstring                   Filepath;
    UINT32                    ActualLength;
    UINT32                    Iterator;
    UINT32                    Pid                 = 0;
    UINT32                    Length              = 0;
    UINT64                    StartAddress        = 0;
    UINT64                    EndAddress          = 0;
    BOOLEAN                   IsFirstCommand      = TRUE;
    BOOLEAN                   NextIsProcId        = FALSE;
    BOOLEAN                   NextIsPath          = FALSE;
    BOOLEAN                   IsTheFirstAddr      = FALSE;
    BOOLEAN                   IsTheSecondAddr     = FALSE;
    BOOLEAN                   IsDumpPathSpecified = FALSE;
    string                    FirstCommand        = SplitCommand.front();
    DEBUGGER_READ_MEMORY_TYPE MemoryType          = DEBUGGER_READ_VIRTUAL_ADDRESS;
    if (SplitCommand.size() <= 4){
        ShowMessages("err, incorrect use of the '.dump' command\n\n");
        CommandDumpHelp();
        return;
    }
    if (g_ActiveProcessDebuggingState.IsActive){
        Pid = g_ActiveProcessDebuggingState.ProcessId;
    }
    for (auto Section : SplitCommand){
        if (IsFirstCommand == TRUE){
            IsFirstCommand = FALSE;
            continue;
        }
        else if (NextIsProcId){
            if (!ConvertStringToUInt32(Section, &Pid)){
                ShowMessages("please specify a correct hex value for process id\n\n");
                CommandDumpHelp();
                return;
            }
            NextIsProcId = FALSE;
            continue;
        }
        else if (NextIsPath){
            StringToWString(Filepath, Section);
            IsDumpPathSpecified = TRUE;
            NextIsPath = FALSE;
        }
        else if (!Section.compare("pid")){
            NextIsProcId = TRUE;
            continue;
        }
        else if (!Section.compare("path")){
            NextIsPath = TRUE;
            continue;
        }
        else if (!IsTheFirstAddr && SymbolConvertNameOrExprToAddress(Section, &StartAddress)){
            IsTheFirstAddr = TRUE;
        }
        else if (!IsTheSecondAddr && SymbolConvertNameOrExprToAddress(Section, &EndAddress)){
            IsTheSecondAddr = TRUE;
        }
        else{
            ShowMessages("err, couldn't resolve error at '%s'\n\n",
                         Section.c_str());
            CommandDumpHelp();
            return;
        }
    }
    if (NextIsProcId){
        ShowMessages("please specify a correct hex value for process id\n\n");
        CommandDumpHelp();
        return;
    }
    if (NextIsPath || !IsDumpPathSpecified){
        ShowMessages("please specify a correct path for saving the dump\n\n");
        CommandDumpHelp();
        return;
    }
    if (!IsTheFirstAddr || !IsTheSecondAddr){
        ShowMessages("err, please specify the start and end address in hex format\n");
        return;
    }
    if (StartAddress >= EndAddress){
        ShowMessages("err, please note that the 'to' address should be greater than the 'from' address\n");
        return;
    }
    if (g_IsSerialConnectedToRemoteDebuggee && Pid != 0){
        ShowMessages(ASSERT_MESSAGE_CANNOT_SPECIFY_PID);
        return;
    }
    if (Pid == 0){
        Pid = GetCurrentProcessId();
    }
    if (!FirstCommand.compare("!dump")){
        MemoryType = DEBUGGER_READ_PHYSICAL_ADDRESS;
    }
    DumpFileHandle = CreateFileW(
        Filepath.c_str(),
        GENERIC_WRITE,
        0,
        NULL,
        CREATE_ALWAYS,
        FILE_ATTRIBUTE_NORMAL,
        NULL);
    if (DumpFileHandle == INVALID_HANDLE_VALUE){
        ShowMessages("err, unable to create or open the file\n");
        return;
    }
    Length = (UINT32)(EndAddress - StartAddress);
    ActualLength = NULL;
    Iterator     = Length / PAGE_SIZE;
    for (size_t i = 0; i <= Iterator; i++){
        UINT64 Address = StartAddress + (i * PAGE_SIZE);
        if (Length >= PAGE_SIZE){
            ActualLength = PAGE_SIZE;
        }
        else{
            ActualLength = Length;
        }
        Length -= ActualLength;
        if (ActualLength != 0){
            HyperDbgReadMemoryAndDisassemble(
                DEBUGGER_SHOW_COMMAND_DUMP,
                Address,
                MemoryType,
                READ_FROM_KERNEL,
                Pid,
                ActualLength,
                NULL);
        }
    }
    if (DumpFileHandle != NULL){
        CloseHandle(DumpFileHandle);
        DumpFileHandle = NULL;
    }
    ShowMessages("the dump file is saved at: %ls\n", Filepath.c_str());
}
VOID CommandDumpSaveIntoFile(PVOID Buffer, UINT32 Length){
    DWORD BytesWritten;
    if (DumpFileHandle == NULL){
        ShowMessages("err, invalid handle for saving the dump buffer is specified\n");
        return;
    }
    if (!WriteFile(DumpFileHandle, Buffer, Length, &BytesWritten, NULL)){
        ShowMessages("err, unable to write buffer into the dump\n");
        CloseHandle(DumpFileHandle);
        DumpFileHandle = NULL;
        return;
    }
}
