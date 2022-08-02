package core

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\debugger\core\debugger.cpp.back

type (
	Debugger interface {
		ShowErrorMessage() (ok bool) //col:248
		InterpretScript() (ok bool)  //col:740
	}
	debugger struct{}
)

func NewDebugger() Debugger { return &debugger{} }

func (d *debugger) ShowErrorMessage() (ok bool) { //col:248
	/*
	   ShowErrorMessage(UINT32 Error)

	   	{
	   	    switch (Error)
	   	    {
	   	    case DEBUGGER_ERROR_TAG_NOT_EXISTS:
	   	        ShowMessages("err, tag not found (%x)\n",
	   	                     Error);
	   	        ShowMessages("err, invalid action type (%x)\n",
	   	                     Error);
	   	        ShowMessages("err, action buffer size is zero (%x)\n",
	   	                     Error);
	   	        ShowMessages("err, the event type is invalid (%x)\n",
	   	                     Error);
	   	        ShowMessages("err, unable to create event (%x)\n",
	   	                     Error);
	   	        ShowMessages("err, invalid address (%x)\n",
	   	                     Error);
	   	        ShowMessages("err, invalid core id (%x)\n",
	   	                     Error);
	   	        ShowMessages("err, exception index exceed first 32 entries (%x)\n",
	   	                     Error);
	   	        ShowMessages("err, interrupt index is not valid (%x)\n",
	   	                     Error);
	   	        ShowMessages("err, unable to hide or unhide debugger (%x)\n",
	   	                     Error);
	   	        ShowMessages("err, debugger already unhide (%x)\n",
	   	                     Error);
	   	        ShowMessages("err, edit memeory request has invalid parameters (%x)\n",
	   	                     Error);
	   	        ShowMessages("err, edit memory request has invalid address based on "
	   	                     "current process layout, the address might be valid but not "
	   	                     "present in the ram (%x)\n",
	   	                     Error);
	   	        ShowMessages("err, edit memory request has invalid address based on other "
	   	                     "process layout (%x)\n",
	   	                     Error);
	   	        ShowMessages("err, modify event with invalid tag (%x)\n",
	   	                     Error);
	   	        ShowMessages("err, modify event with invalid type of action (%x)\n",
	   	                     Error);
	   	        ShowMessages("err, invalid parameter passsed to stepping core. (%x)\n",
	   	                     Error);
	   	        ShowMessages(
	   	            "err, the target thread not found or the thread is disabled (%x)\n",
	   	            Error);
	   	        ShowMessages("err, invalid baud rate (%x)\n",
	   	                     Error);
	   	        ShowMessages("err, invalid serial port (%x)\n",
	   	                     Error);
	   	        ShowMessages("err, invalid core selected in switching cores (%x)\n",
	   	                     Error);
	   	        ShowMessages("err, unable to switch to the new process (%x)\n",
	   	                     Error);
	   	        ShowMessages("err, unable to run the script on remote debuggee (%x)\n",
	   	                     Error);
	   	        ShowMessages("err, invalid register number (%x)\n",
	   	                     Error);
	   	        ShowMessages("err, maximum number of breakpoints are used, you need to "
	   	                     "send an ioctl to re-allocate new pre-allocated buffers or "
	   	                     "configure HyperDbg to pre-allocated more buffers by "
	   	                     "configuring MAXIMUM_BREAKPOINTS_WITHOUT_CONTINUE (%x)\n",
	   	                     Error);
	   	        ShowMessages("err, breakpoint already exists on target address, you can "
	   	                     "use 'bl' command to view a list of breakpoints (%x)\n",
	   	                     Error);
	   	        ShowMessages("err, breakpoint id is invalid (%x)\n",
	   	                     Error);
	   	        ShowMessages("err, breakpoint already disabled (%x)\n",
	   	                     Error);
	   	        ShowMessages("err, breakpoint already enabled (%x)\n",
	   	                     Error);
	   	        ShowMessages("err, the memory type is invalid (%x)\n",
	   	                     Error);
	   	        ShowMessages("err, the process id is invalid, make sure to enter the "
	   	                     "process id in hex format, or if you want to use it in decimal "
	   	                     "format, add '0n' prefix to the number (%x)\n",
	   	                     Error);
	   	        ShowMessages("err, the event is not applied (%x)\n",
	   	                     Error);
	   	        ShowMessages("err, either the process id or the _EPROCESS is invalid or "
	   	                     "cannot get the details based on the provided parameters (%x)\n",
	   	                     Error);
	   	        ShowMessages("err, either the thread id, _ETHREAD, or _EPROCESS is invalid or "
	   	                     "cannot get the details based on the provided parameters (%x)\n",
	   	                     Error);
	   	        ShowMessages("err, the maximum breakpoint for a single page is hit, "
	   	                     "you cannot apply more breakpoints in this page (%x)\n",
	   	                     Error);
	   	        ShowMessages("err, the pre-allocated buffer is empty, usually this buffer will be "
	   	                     "filled at the next IOCTL when the debugger is continued (%x)\n"
	   	                     "please visit the documentation for the 'prealloc' command or use "
	   	                     "'.help prealloc' to to reserve more pre-allocated pools\n",
	   	                     Error);
	   	        ShowMessages("err, could not convert 2MB large page to 4KB pages (%x)\n",
	   	                     Error);
	   	        ShowMessages("err, failed to get the PML1 entry of the target address (%x)\n",
	   	                     Error);
	   	        ShowMessages("err, the page modification is not applied, make sure that you don't "
	   	                     "put multiple EPT Hooks or Monitors on a single page (%x)\n",
	   	                     Error);
	   	        ShowMessages("err, could not build the EPT hook (%x)\n",
	   	                     Error);
	   	        ShowMessages("err, could not find the allocation type (%x)\n",
	   	                     Error);
	   	        ShowMessages("err, invalid index specified for test query command (%x)\n",
	   	                     Error);
	   	        ShowMessages("err, unable to attach to the target process (%x)\n",
	   	                     Error);
	   	        ShowMessages("err, unable to remove hooks as the entrypoint of user-mode "
	   	                     "process is not reached yet (%x)\n",
	   	                     Error);
	   	        ShowMessages("err, unable to remove hooks (%x)\n",
	   	                     Error);
	   	        ShowMessages("err, the routines for getting the PEB is not correctly "
	   	                     "initialized (%x)\n",
	   	                     Error);
	   	        ShowMessages("err, unable to detect whether the process was 32-bit "
	   	                     "or 64-bit (%x)\n",
	   	                     Error);
	   	        ShowMessages("err, unable to kill the process (%x)\n",
	   	                     Error);
	   	        ShowMessages("err, invalid thread debugging token (%x)\n",
	   	                     Error);
	   	        ShowMessages("err, unable to pause the threads of the process (%x)\n",
	   	                     Error);
	   	        ShowMessages("err, the user debugger is already attached to this "
	   	                     "process, please use the '.switch' command to switch "
	   	                     "to this process (%x)\n",
	   	                     Error);
	   	        ShowMessages("err, the user debugger is not already attached to "
	   	                     "the process (%x)\n",
	   	                     Error);
	   	        ShowMessages("err, the user debugger is not able to detach from "
	   	                     "this process as there are paused threads in the "
	   	                     "target process, please make sure to remove all "
	   	                     "the events and continue the target process, then "
	   	                     "perform the detach again (%x)\n",
	   	                     Error);
	   	        ShowMessages("err, unable to switch to the process id or thread id "
	   	                     "as the target process id or thread id is not found in "
	   	                     "the attached threads list, please view the list of "
	   	                     "processes and threads by using the '.switch list' command (%x)\n",
	   	                     Error);
	   	        ShowMessages("err, unable to switch to the process as the process doesn't "
	   	                     "contain an active intercepted thread (%x)\n",
	   	                     Error);
	   	        ShowMessages("err, unable to get user-mode modules of the process (%x)\n",
	   	                     Error);
	   	        ShowMessages("err, unable to get the callstack (%x)\n",
	   	                     Error);
	   	        ShowMessages("err, unable to query count of processes or threads (%x)\n",
	   	                     Error);
	   	        ShowMessages("err, error not found (%x)\n",
	   	                     Error);

	   DebuggerGetNtoskrnlBase()

	   	{
	   	    NTSTATUS             Status                  = STATUS_UNSUCCESSFUL;
	   	    UINT64               NtoskrnlBase            = NULL;
	   	    PRTL_PROCESS_MODULES Modules                 = NULL;
	   	    ULONG                SysModuleInfoBufferSize = 0;
	   	    Status = NtQuerySystemInformation(SystemModuleInformation, NULL, NULL, &SysModuleInfoBufferSize);
	   	    Modules = (PRTL_PROCESS_MODULES)malloc(SysModuleInfoBufferSize);
	   	    NtQuerySystemInformation(SystemModuleInformation, Modules, SysModuleInfoBufferSize, NULL);
	   	    for (int i = 0; i < Modules->NumberOfModules; i++)
	   	    {
	   	        if (!strcmp((const char *)Modules->Modules[i].FullPathName +
	   	                        Modules->Modules[i].OffsetToFileName,
	   	                    "ntoskrnl.exe"))
	   	        {
	   	            NtoskrnlBase = (UINT64)Modules->Modules[i].ImageBase;
	   	            break;
	   	        }
	   	    }
	   	    free(Modules);

	   DebuggerPauseDebuggee()

	   	{
	   	    BOOLEAN                        StatusIoctl    = 0;
	   	    ULONG                          ReturnedLength = 0;
	   	    DEBUGGER_PAUSE_PACKET_RECEIVED PauseRequest   = {0};
	   	    StatusIoctl =
	   	        DeviceIoControl(g_DeviceHandle,
	   	                        IOCTL_PAUSE_PACKET_RECEIVED,
	   	                        &PauseRequest,
	   	                        SIZEOF_DEBUGGER_PAUSE_PACKET_RECEIVED,
	   	                        &PauseRequest,
	   	                        SIZEOF_DEBUGGER_PAUSE_PACKET_RECEIVED,
	   	                        &ReturnedLength,
	   	                        NULL
	   	        );
	   	    if (!StatusIoctl)
	   	    {
	   	        ShowMessages("ioctl failed with code 0x%x\n", GetLastError());
	   	    if (PauseRequest.Result == DEBUGGER_OPERATION_WAS_SUCCESSFULL)
	   	    {
	   	        return TRUE;
	   	    }
	   	    else
	   	    {
	   	        ShowErrorMessage(PauseRequest.Result);

	   IsConnectedToAnyInstanceOfDebuggerOrDebuggee()

	   	{
	   	    if (g_DeviceHandle)
	   	    {
	   	        ShowMessages("err, the current system is already connected to the local "
	   	                     "debugging, use '.disconnect' to disconnect\n");
	   	    else if (g_IsConnectedToRemoteDebuggee)
	   	    {
	   	        ShowMessages("err, the current system is already connected to remote "
	   	                     "machine (debuggee), use '.disconnect' to disconnect from the "
	   	                     "remote machine\n");
	   	    else if (g_IsConnectedToRemoteDebugger)
	   	    {
	   	        ShowMessages("err, the current system is already connected to remote "
	   	                     "machine (debugger), use '.disconnect' to disconnect from the "
	   	                     "remote machine from debugger\n");
	   	    else if (g_IsSerialConnectedToRemoteDebuggee)
	   	    {
	   	        ShowMessages(
	   	            "err, the current system is already connected to remote "
	   	            "machine (debuggee), use '.debug close' to disconnect from the "
	   	            "remote machine\n");
	   	    else if (g_IsSerialConnectedToRemoteDebugger)
	   	    {
	   	        ShowMessages(
	   	            "err, the current system is already connected to remote "
	   	            "machine (debugger), use '.debug close' to disconnect from the "
	   	            "remote machine from debugger\n");

	   IsTagExist(UINT64 Tag)

	   	{
	   	    PLIST_ENTRY                    TempList      = 0;
	   	    PDEBUGGER_GENERAL_EVENT_DETAIL CommandDetail = {0};
	   	    if (!g_EventTraceInitialized)
	   	    {
	   	        return FALSE;
	   	    }
	   	    TempList = &g_EventTrace;
	   	    while (&g_EventTrace != TempList->Blink)
	   	    {
	   	        TempList = TempList->Blink;
	   	        CommandDetail = CONTAINING_RECORD(TempList, DEBUGGER_GENERAL_EVENT_DETAIL, CommandsEventList);
	   	        if (CommandDetail->Tag == Tag ||
	   	            Tag == DEBUGGER_MODIFY_EVENTS_APPLY_TO_ALL_TAG)
	   	        {
	   	            return TRUE;
	   	        }
	   	    }
	   	    return FALSE;
	   	}
	*/
	return true
}

func (d *debugger) InterpretScript() (ok bool) { //col:740
	/*
	   InterpretScript(vector<string> * SplittedCommand,

	   	vector<string> * SplittedCommandCaseSensitive,
	   	PBOOLEAN         ScriptSyntaxErrors,
	   	PUINT64          BufferAddress,
	   	PUINT32          BufferLength,
	   	PUINT32          Pointer,
	   	PUINT64          ScriptCodeBuffer)

	   	{
	   	    BOOLEAN        IsTextVisited = FALSE;
	   	    BOOLEAN        IsInState     = FALSE;
	   	    BOOLEAN        IsEnded       = FALSE;
	   	    string         AppendedFinalBuffer;
	   	    vector<string> SaveBuffer;
	   	    vector<int>    IndexesToRemove;
	   	    UCHAR *        FinalBuffer;
	   	    int            Index            = 0;
	   	    int            NewIndexToRemove = 0;
	   	    int            OpenBracket      = 0;
	   	    size_t         CountOfOpenBrackets;
	   	    size_t         CountOfCloseBrackets;
	   	    UINT32         IndexInCommandCaseSensitive          = 0;
	   	    vector<string> SplittedCommandCaseSensitiveInstance = *SplittedCommandCaseSensitive;
	   	    string         TempStr;
	   	    *ScriptSyntaxErrors = TRUE;
	   	    for (auto Section : *SplittedCommand)
	   	    {
	   	        IndexInCommandCaseSensitive++;
	   	        Index++;
	   	        if (IsInState)
	   	        {
	   	            if (OpenBracket == 0 && Section.find('{') == string::npos)
	   	            {
	   	                if (!Section.compare("}"))
	   	                {
	   	                    IndexesToRemove.push_back(Index);
	   	                if (HasEnding(Section, "}"))
	   	                {
	   	                    IndexesToRemove.push_back(Index);
	   	                    TempStr = SplittedCommandCaseSensitiveInstance.at(IndexInCommandCaseSensitive - 1);
	   	                    SaveBuffer.emplace_back(TempStr.begin(), TempStr.begin() + TempStr.size() - 1);
	   	            IndexesToRemove.push_back(Index);
	   	            if (Section.find('{') != string::npos)
	   	            {
	   	                size_t CountOfBrackets = count(Section.begin(), Section.end(), '{');
	   	            if (Section.find('}') != string::npos)
	   	            {
	   	                size_t CountOfBrackets = count(Section.begin(), Section.end(), '}');
	   	                if (OpenBracket < 0)
	   	                {
	   	                    OpenBracket = 0;
	   	                    IsEnded     = TRUE;
	   	                    TempStr     = SplittedCommandCaseSensitiveInstance.at(IndexInCommandCaseSensitive - 1);
	   	                    SaveBuffer.emplace_back(TempStr.begin(), TempStr.begin() + TempStr.size() - 1);
	   	            SaveBuffer.push_back(SplittedCommandCaseSensitiveInstance.at(IndexInCommandCaseSensitive - 1));
	   	        if (IsTextVisited && !Section.compare("{"))
	   	        {
	   	            IndexesToRemove.push_back(Index);
	   	        if (IsTextVisited && Section.rfind('{', 0) == 0)
	   	        {
	   	            if (Section.find('{') != string::npos)
	   	            {
	   	                size_t CountOfBrackets = count(Section.begin(), Section.end(), '{');
	   	            if (OpenBracket == 0 && HasEnding(Section, "}"))
	   	            {
	   	                IndexesToRemove.push_back(Index);
	   	                TempStr = SplittedCommandCaseSensitiveInstance.at(IndexInCommandCaseSensitive - 1).erase(0, 1);
	   	                SaveBuffer.emplace_back(TempStr.begin(), TempStr.begin() + TempStr.size() - 1);
	   	            if (Section.find('}') != string::npos)
	   	            {
	   	                size_t CountOfBrackets = count(Section.begin(), Section.end(), '}');
	   	                if (OpenBracket < 0)
	   	                {
	   	                    OpenBracket = 0;
	   	                    IsEnded     = TRUE;
	   	                    TempStr     = SplittedCommandCaseSensitiveInstance.at(IndexInCommandCaseSensitive - 1);
	   	                    SaveBuffer.emplace_back(TempStr.begin(), TempStr.begin() + TempStr.size() - 1);
	   	            IndexesToRemove.push_back(Index);
	   	            SaveBuffer.push_back(SplittedCommandCaseSensitiveInstance.at(IndexInCommandCaseSensitive - 1).erase(0, 1));
	   	        if (!Section.compare("script"))
	   	        {
	   	            IndexesToRemove.push_back(Index);
	   	        if (!Section.compare("script{"))
	   	        {
	   	            IndexesToRemove.push_back(Index);
	   	        if (Section.rfind("script{", 0) == 0)
	   	        {
	   	            IndexesToRemove.push_back(Index);
	   	            CountOfOpenBrackets  = count(Section.begin(), Section.end(), '{');
	   	            CountOfCloseBrackets = count(Section.begin(), Section.end(), '}');
	   	            if (Section.find('{') != string::npos)
	   	            {
	   	                OpenBracket += CountOfOpenBrackets - 1;
	   	            }
	   	            if (CountOfOpenBrackets == CountOfCloseBrackets ||
	   	                (OpenBracket == 0 && HasEnding(Section, "}")))
	   	            {
	   	                TempStr = SplittedCommandCaseSensitiveInstance.at(IndexInCommandCaseSensitive - 1).erase(0, 7);
	   	                SaveBuffer.emplace_back(TempStr.begin(), TempStr.begin() + TempStr.size() - 1);
	   	                SaveBuffer.push_back(SplittedCommandCaseSensitiveInstance.at(IndexInCommandCaseSensitive - 1).erase(0, 7));
	   	                if (Section.find('}') != string::npos)
	   	                {
	   	                    size_t CountOfBrackets = count(Section.begin(), Section.end(), '}');
	   	                    if (OpenBracket < 0)
	   	                    {
	   	                        OpenBracket = 0;
	   	                        IsEnded     = TRUE;
	   	                    }
	   	                }
	   	                continue;
	   	            }
	   	        }
	   	    }
	   	    if (SaveBuffer.size() == 0)
	   	    {
	   	        return FALSE;
	   	    }
	   	    if (OpenBracket != 0)
	   	    {
	   	        return FALSE;
	   	    }
	   	    if (!IsEnded)
	   	    {
	   	        return FALSE;
	   	    }
	   	    for (auto Section : SaveBuffer)
	   	    {
	   	        AppendedFinalBuffer.append(Section);
	   	        AppendedFinalBuffer.append(" ");
	   	    if (AppendedFinalBuffer.rfind("file:", 0) == 0)
	   	    {
	   	        std::ifstream     t(AppendedFinalBuffer.erase(0, 5).c_str());
	   	        buffer << t.rdbuf();
	   	        AppendedFinalBuffer = buffer.str();
	   	        if (AppendedFinalBuffer.empty())
	   	        {
	   	            ShowMessages("err, either script file is not found or it's empty\n");
	   	        ScriptEngineParseWrapper((char *)AppendedFinalBuffer.c_str(), TRUE);
	   	    if (CodeBuffer == NULL)
	   	    {
	   	        *ScriptSyntaxErrors = TRUE;
	   	        return TRUE;
	   	    }
	   	    else
	   	    {
	   	        *ScriptSyntaxErrors = FALSE;
	   	    }
	   	    *BufferAddress    = ScriptEngineWrapperGetHead(CodeBuffer);
	   	    *BufferLength     = ScriptEngineWrapperGetSize(CodeBuffer);
	   	    *Pointer          = ScriptEngineWrapperGetPointer(CodeBuffer);
	   	    *ScriptCodeBuffer = (UINT64)CodeBuffer;
	   	    NewIndexToRemove = 0;
	   	    for (auto IndexToRemove : IndexesToRemove)
	   	    {
	   	        NewIndexToRemove++;
	   	        SplittedCommand->erase(SplittedCommand->begin() +
	   	                               (IndexToRemove - NewIndexToRemove));
	   	        SplittedCommandCaseSensitive->erase(SplittedCommandCaseSensitive->begin() +
	   	                                            (IndexToRemove - NewIndexToRemove));

	   InterpretConditionsAndCodes(vector<string> * SplittedCommand,

	   	vector<string> * SplittedCommandCaseSensitive,
	   	BOOLEAN          IsConditionBuffer,
	   	PUINT64          BufferAddress,
	   	PUINT32          BufferLength)

	   	{
	   	    BOOLEAN        IsTextVisited = FALSE;
	   	    BOOLEAN        IsInState     = FALSE;
	   	    BOOLEAN        IsEnded       = FALSE;
	   	    string         Temp;
	   	    string         TempStr;
	   	    string         AppendedFinalBuffer;
	   	    vector<string> SaveBuffer;
	   	    vector<CHAR>   ParsedBytes;
	   	    vector<int>    IndexesToRemove;
	   	    UCHAR *        FinalBuffer;
	   	    int            NewIndexToRemove = 0;
	   	    int            Index            = 0;
	   	    for (auto Section : *SplittedCommand)
	   	    {
	   	        Index++;
	   	        if (IsInState)
	   	        {
	   	            if (!Section.compare("}"))
	   	            {
	   	                IndexesToRemove.push_back(Index);
	   	            if (HasEnding(Section, "}"))
	   	            {
	   	                IndexesToRemove.push_back(Index);
	   	                SaveBuffer.emplace_back(Section.begin(), Section.begin() + Section.size() - 1);
	   	            IndexesToRemove.push_back(Index);
	   	            SaveBuffer.push_back(Section);
	   	        if (IsTextVisited && !Section.compare("{"))
	   	        {
	   	            IndexesToRemove.push_back(Index);
	   	        if (IsTextVisited && Section.rfind('{', 0) == 0)
	   	        {
	   	            if (HasEnding(Section, "}"))
	   	            {
	   	                IndexesToRemove.push_back(Index);
	   	                TempStr = Section.erase(0, 1);
	   	                SaveBuffer.emplace_back(TempStr.begin(), TempStr.begin() + TempStr.size() - 1);
	   	            IndexesToRemove.push_back(Index);
	   	            SaveBuffer.push_back(Section.erase(0, 1));
	   	        if (IsConditionBuffer)
	   	        {
	   	            if (!Section.compare("condition"))
	   	            {
	   	                IndexesToRemove.push_back(Index);
	   	            if (!Section.compare("code"))
	   	            {
	   	                IndexesToRemove.push_back(Index);
	   	        if (IsConditionBuffer)
	   	        {
	   	            if (!Section.compare("condition{"))
	   	            {
	   	                IndexesToRemove.push_back(Index);
	   	            if (!Section.compare("code{"))
	   	            {
	   	                IndexesToRemove.push_back(Index);
	   	        if (IsConditionBuffer)
	   	        {
	   	            if (Section.rfind("condition{", 0) == 0)
	   	            {
	   	                IndexesToRemove.push_back(Index);
	   	                if (!HasEnding(Section, "}"))
	   	                {
	   	                    SaveBuffer.push_back(Section.erase(0, 10));
	   	                    TempStr = Section.erase(0, 10);
	   	                    SaveBuffer.emplace_back(TempStr.begin(), TempStr.begin() + TempStr.size() - 1);
	   	            if (Section.rfind("code{", 0) == 0)
	   	            {
	   	                IndexesToRemove.push_back(Index);
	   	                if (!HasEnding(Section, "}"))
	   	                {
	   	                    SaveBuffer.push_back(Section.erase(0, 5));
	   	                    TempStr = Section.erase(0, 5);
	   	                    SaveBuffer.emplace_back(TempStr.begin(), TempStr.begin() + TempStr.size() - 1);
	   	    if (SaveBuffer.size() == 0)
	   	    {
	   	        return FALSE;
	   	    }
	   	    if (!IsEnded)
	   	    {
	   	        return FALSE;
	   	    }
	   	    SaveBuffer.push_back("c3");
	   	    for (auto Section : SaveBuffer)
	   	    {
	   	        if (Section.rfind("0x", 0) == 0 || Section.rfind("0X", 0) == 0 ||
	   	            Section.rfind("\\x", 0) == 0 || Section.rfind("\\X", 0) == 0)
	   	        {
	   	            Temp = Section.erase(0, 2);
	   	        else if (Section.rfind('x', 0) == 0 || Section.rfind('X', 0) == 0)
	   	        {
	   	            Temp = Section.erase(0, 1);
	   	            Temp = std::move(Section);
	   	        ReplaceAll(Temp, "\\x", "");
	   	        if (Temp.size() % 2 != 0)
	   	        {
	   	            Temp.insert(0, 1, '0');
	   	        if (!IsHexNotation(Temp))
	   	        {
	   	            ShowMessages("please enter condition code in a hex notation\n");
	   	        AppendedFinalBuffer.append(Temp);
	   	    ParsedBytes = HexToBytes(AppendedFinalBuffer);
	   	    FinalBuffer = (unsigned char *)malloc(ParsedBytes.size());
	   	    std::copy(ParsedBytes.begin(), ParsedBytes.end(), FinalBuffer);
	   	    *BufferAddress = (UINT64)FinalBuffer;
	   	    *BufferLength  = ParsedBytes.size();
	   	    for (auto IndexToRemove : IndexesToRemove)
	   	    {
	   	        NewIndexToRemove++;
	   	        SplittedCommand->erase(SplittedCommand->begin() +
	   	                               (IndexToRemove - NewIndexToRemove));
	   	        SplittedCommandCaseSensitive->erase(SplittedCommandCaseSensitive->begin() +
	   	                                            (IndexToRemove - NewIndexToRemove));

	   InterpretOutput(vector<string> * SplittedCommand,

	   	vector<string> * SplittedCommandCaseSensitive,
	   	vector<string> & InputSources)

	   	{
	   	    BOOLEAN        IsTextVisited = FALSE;
	   	    BOOLEAN        IsInState     = FALSE;
	   	    BOOLEAN        IsEnded       = FALSE;
	   	    string         AppendedFinalBuffer;
	   	    vector<string> SaveBuffer;
	   	    vector<int>    IndexesToRemove;
	   	    string         Token;
	   	    string         TempStr;
	   	    int            NewIndexToRemove                     = 0;
	   	    int            Index                                = 0;
	   	    char           Delimiter                            = ',';
	   	    size_t         Pos                                  = 0;
	   	    vector<string> SplittedCommandCaseSensitiveInstance = *SplittedCommandCaseSensitive;
	   	    UINT32         IndexInCommandCaseSensitive          = 0;
	   	    for (auto Section : *SplittedCommand)
	   	    {
	   	        IndexInCommandCaseSensitive++;
	   	        Index++;
	   	        if (IsInState)
	   	        {
	   	            if (!Section.compare("}"))
	   	            {
	   	                IndexesToRemove.push_back(Index);
	   	            if (HasEnding(Section, "}"))
	   	            {
	   	                IndexesToRemove.push_back(Index);
	   	                TempStr = SplittedCommandCaseSensitiveInstance.at(IndexInCommandCaseSensitive - 1);
	   	                SaveBuffer.emplace_back(TempStr.begin(), TempStr.begin() + TempStr.size() - 1);
	   	            IndexesToRemove.push_back(Index);
	   	            SaveBuffer.push_back(SplittedCommandCaseSensitiveInstance.at(IndexInCommandCaseSensitive - 1));
	   	        if (IsTextVisited && !Section.compare("{"))
	   	        {
	   	            IndexesToRemove.push_back(Index);
	   	        if (IsTextVisited && Section.rfind('{', 0) == 0)
	   	        {
	   	            if (HasEnding(Section, "}"))
	   	            {
	   	                IndexesToRemove.push_back(Index);
	   	                TempStr = SplittedCommandCaseSensitiveInstance.at(IndexInCommandCaseSensitive - 1).erase(0, 1);
	   	                SaveBuffer.emplace_back(TempStr.begin(), TempStr.begin() + TempStr.size() - 1);
	   	            IndexesToRemove.push_back(Index);
	   	            SaveBuffer.push_back(SplittedCommandCaseSensitiveInstance.at(IndexInCommandCaseSensitive - 1).erase(0, 1));
	   	        if (!Section.compare("output"))
	   	        {
	   	            IndexesToRemove.push_back(Index);
	   	        if (!Section.compare("output{"))
	   	        {
	   	            IndexesToRemove.push_back(Index);
	   	        if (Section.rfind("output{", 0) == 0)
	   	        {
	   	            IndexesToRemove.push_back(Index);
	   	            if (!HasEnding(Section, "}"))
	   	            {
	   	                SaveBuffer.push_back(SplittedCommandCaseSensitiveInstance.at(IndexInCommandCaseSensitive - 1).erase(0, 7));
	   	                TempStr = SplittedCommandCaseSensitiveInstance.at(IndexInCommandCaseSensitive - 1).erase(0, 7);
	   	                SaveBuffer.emplace_back(TempStr.begin(), TempStr.begin() + TempStr.size() - 1);
	   	    if (SaveBuffer.size() == 0)
	   	    {
	   	        return FALSE;
	   	    }
	   	    if (!IsEnded)
	   	    {
	   	        return FALSE;
	   	    }
	   	    for (auto Section : SaveBuffer)
	   	    {
	   	        AppendedFinalBuffer.append(Section);
	   	        AppendedFinalBuffer.append(" ");
	   	    if (AppendedFinalBuffer.find(Delimiter) != std::string::npos)
	   	    {
	   	        while ((Pos = AppendedFinalBuffer.find(Delimiter)) != string::npos)
	   	        {
	   	            Token = AppendedFinalBuffer.substr(0, Pos);
	   	            Trim(Token);
	   	            if (!Token.empty())
	   	            {
	   	                InputSources.push_back(Token);
	   	            AppendedFinalBuffer.erase(0, Pos + sizeof(Delimiter) / sizeof(char));
	   	        if (!AppendedFinalBuffer.empty())
	   	        {
	   	            InputSources.push_back(AppendedFinalBuffer);
	   	        InputSources.push_back(AppendedFinalBuffer);
	   	    for (auto IndexToRemove : IndexesToRemove)
	   	    {
	   	        NewIndexToRemove++;
	   	        SplittedCommand->erase(SplittedCommand->begin() +
	   	                               (IndexToRemove - NewIndexToRemove));
	   	        SplittedCommandCaseSensitive->erase(SplittedCommandCaseSensitive->begin() +
	   	                                            (IndexToRemove - NewIndexToRemove));

	   SendEventToKernel(PDEBUGGER_GENERAL_EVENT_DETAIL Event,

	   	UINT32                         EventBufferLength)

	   	{
	   	    BOOL                                  Status;
	   	    ULONG                                 ReturnedLength;
	   	    DEBUGGER_EVENT_AND_ACTION_REG_BUFFER  ReturnedBuffer = {0};
	   	    PDEBUGGER_EVENT_AND_ACTION_REG_BUFFER TempRegResult;
	   	    if (g_IsSerialConnectedToRemoteDebuggee)
	   	    {
	   	        TempRegResult =
	   	            KdSendRegisterEventPacketToDebuggee(Event, EventBufferLength);
	   	        memcpy(&ReturnedBuffer, TempRegResult, sizeof(DEBUGGER_EVENT_AND_ACTION_REG_BUFFER));
	   	        AssertShowMessageReturnStmt(g_DeviceHandle, ASSERT_MESSAGE_DRIVER_NOT_LOADED, AssertReturnFalse);
	   	            DeviceIoControl(g_DeviceHandle,
	   	                            IOCTL_DEBUGGER_REGISTER_EVENT,
	   	                            Event,
	   	                            EventBufferLength,
	   	                            &ReturnedBuffer,
	   	                            sizeof(DEBUGGER_EVENT_AND_ACTION_REG_BUFFER),
	   	                            &ReturnedLength,
	   	                            NULL
	   	            );
	   	        if (!Status)
	   	        {
	   	            ShowMessages("ioctl failed with code 0x%x\n", GetLastError());
	   	    if (ReturnedBuffer.IsSuccessful && ReturnedBuffer.Error == 0)
	   	    {
	   	        if (!g_IsSerialConnectedToRemoteDebuggee &&
	   	            !g_IsSerialConnectedToRemoteDebugger &&
	   	            g_BreakPrintingOutput &&
	   	            g_AutoUnpause)
	   	        {
	   	            g_BreakPrintingOutput = FALSE;
	   	            if (g_IsConnectedToRemoteDebuggee)
	   	            {
	   	                RemoteConnectionSendCommand("g", strlen("g") + 1);
	   	            ShowMessages("\n");
	   	        if (ReturnedBuffer.Error != 0)
	   	        {
	   	            ShowErrorMessage(ReturnedBuffer.Error);

	   RegisterActionToEvent(PDEBUGGER_GENERAL_EVENT_DETAIL Event,

	   	PDEBUGGER_GENERAL_ACTION       ActionBreakToDebugger,
	   	UINT32                         ActionBreakToDebuggerLength,
	   	PDEBUGGER_GENERAL_ACTION       ActionCustomCode,
	   	UINT32                         ActionCustomCodeLength,
	   	PDEBUGGER_GENERAL_ACTION       ActionScript,
	   	UINT32                         ActionScriptLength)

	   	{
	   	    BOOL                                  Status;
	   	    ULONG                                 ReturnedLength;
	   	    DEBUGGER_EVENT_AND_ACTION_REG_BUFFER  ReturnedBuffer = {0};
	   	    PDEBUGGER_EVENT_AND_ACTION_REG_BUFFER TempAddingResult;
	   	    if (g_IsSerialConnectedToRemoteDebuggee)
	   	    {
	   	        if (ActionBreakToDebugger != NULL)
	   	        {
	   	            TempAddingResult = KdSendAddActionToEventPacketToDebuggee(
	   	                ActionBreakToDebugger,
	   	                ActionBreakToDebuggerLength);
	   	            memcpy(&ReturnedBuffer, TempAddingResult, sizeof(DEBUGGER_EVENT_AND_ACTION_REG_BUFFER));
	   	        if (ActionCustomCode != NULL)
	   	        {
	   	            TempAddingResult = KdSendAddActionToEventPacketToDebuggee(
	   	                ActionCustomCode,
	   	                ActionCustomCodeLength);
	   	            memcpy(&ReturnedBuffer, TempAddingResult, sizeof(DEBUGGER_EVENT_AND_ACTION_REG_BUFFER));
	   	        if (ActionScript != NULL)
	   	        {
	   	            TempAddingResult = KdSendAddActionToEventPacketToDebuggee(
	   	                ActionScript,
	   	                ActionScriptLength);
	   	            memcpy(&ReturnedBuffer, TempAddingResult, sizeof(DEBUGGER_EVENT_AND_ACTION_REG_BUFFER));
	   	        AssertShowMessageReturnStmt(g_DeviceHandle, ASSERT_MESSAGE_DRIVER_NOT_LOADED, AssertReturnFalse);
	   	        if (ActionBreakToDebugger != NULL)
	   	        {
	   	            Status = DeviceIoControl(
	   	                g_DeviceHandle,
	   	                IOCTL_DEBUGGER_ADD_ACTION_TO_EVENT,
	   	                ActionBreakToDebugger,
	   	                ActionBreakToDebuggerLength,
	   	                &ReturnedBuffer,
	   	                sizeof(DEBUGGER_EVENT_AND_ACTION_REG_BUFFER),
	   	                &ReturnedLength,
	   	                NULL
	   	            );
	   	            if (!Status)
	   	            {
	   	                ShowMessages("ioctl failed with code 0x%x\n", GetLastError());
	   	        if (ActionCustomCode != NULL)
	   	        {
	   	            Status = DeviceIoControl(
	   	                g_DeviceHandle,
	   	                IOCTL_DEBUGGER_ADD_ACTION_TO_EVENT,
	   	                ActionCustomCode,
	   	                ActionCustomCodeLength,
	   	                &ReturnedBuffer,
	   	                sizeof(DEBUGGER_EVENT_AND_ACTION_REG_BUFFER),
	   	                &ReturnedLength,
	   	                NULL
	   	            );
	   	            if (!Status)
	   	            {
	   	                ShowMessages("ioctl failed with code 0x%x\n", GetLastError());
	   	        if (ActionScript != NULL)
	   	        {
	   	            Status = DeviceIoControl(
	   	                g_DeviceHandle,
	   	                IOCTL_DEBUGGER_ADD_ACTION_TO_EVENT,
	   	                ActionScript,
	   	                ActionScriptLength,
	   	                &ReturnedBuffer,
	   	                sizeof(DEBUGGER_EVENT_AND_ACTION_REG_BUFFER),
	   	                &ReturnedLength,
	   	                NULL
	   	            );
	   	            if (!Status)
	   	            {
	   	                ShowMessages("ioctl failed with code 0x%x\n", GetLastError());
	   	    FreeEventsAndActionsMemory(NULL, ActionBreakToDebugger, ActionCustomCode, ActionScript);
	   	    InsertHeadList(&g_EventTrace, &(Event->CommandsEventList));

	   GetNewDebuggerEventTag()

	   	{
	   	    return g_EventTag++;
	   	}
	*/
	return true
}

