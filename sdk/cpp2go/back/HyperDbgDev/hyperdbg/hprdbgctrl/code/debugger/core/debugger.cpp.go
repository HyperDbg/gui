package core
//back\HyperDbgDev\hyperdbg\hprdbgctrl\code\debugger\core\debugger.cpp.back

type (
Debugger interface{
ShowErrorMessage()(ok bool)//col:370
DebuggerGetNtoskrnlBase()(ok bool)//col:408
 * then when it returns true the debuggee is not paused anymore ()(ok bool)//col:458
IsConnectedToAnyInstanceOfDebuggerOrDebuggee()(ok bool)//col:509
IsTagExist()(ok bool)//col:548
 * successful ()(ok bool)//col:1034
 * successful ()(ok bool)//col:1412
 * input that needs to be considered for this event ()(ok bool)//col:1687
SendEventToKernel()(ok bool)//col:1803
RegisterActionToEvent()(ok bool)//col:2005
GetNewDebuggerEventTag()(ok bool)//col:2017
FreeEventsAndActionsMemory()(ok bool)//col:2053
InterpretGeneralEventAndActionsFields()(ok bool)//col:3002
}
)

func NewDebugger() { return & debugger{} }

func (d *debugger)ShowErrorMessage()(ok bool){//col:370
/*ShowErrorMessage(UINT32 Error)
{
    switch (Error)
    {
    case DEBUGGER_ERROR_TAG_NOT_EXISTS:
        ShowMessages("err, tag not found (%x)\n",
                     Error);
        break;
    case DEBUGGER_ERROR_INVALID_ACTION_TYPE:
        ShowMessages("err, invalid action type (%x)\n",
                     Error);
        break;
    case DEBUGGER_ERROR_ACTION_BUFFER_SIZE_IS_ZERO:
        ShowMessages("err, action buffer size is zero (%x)\n",
                     Error);
        break;
    case DEBUGGER_ERROR_EVENT_TYPE_IS_INVALID:
        ShowMessages("err, the event type is invalid (%x)\n",
                     Error);
        break;
    case DEBUGGER_ERROR_UNABLE_TO_CREATE_EVENT:
        ShowMessages("err, unable to create event (%x)\n",
                     Error);
        break;
    case DEBUGGER_ERROR_INVALID_ADDRESS:
        ShowMessages("err, invalid address (%x)\n",
                     Error);
        break;
    case DEBUGGER_ERROR_INVALID_CORE_ID:
        ShowMessages("err, invalid core id (%x)\n",
                     Error);
        break;
    case DEBUGGER_ERROR_EXCEPTION_INDEX_EXCEED_FIRST_32_ENTRIES:
        ShowMessages("err, exception index exceed first 32 entries (%x)\n",
                     Error);
        break;
    case DEBUGGER_ERROR_INTERRUPT_INDEX_IS_NOT_VALID:
        ShowMessages("err, interrupt index is not valid (%x)\n",
                     Error);
        break;
    case DEBUGGER_ERROR_UNABLE_TO_HIDE_OR_UNHIDE_DEBUGGER:
        ShowMessages("err, unable to hide or unhide debugger (%x)\n",
                     Error);
        break;
    case DEBUGGER_ERROR_DEBUGGER_ALREADY_UHIDE:
        ShowMessages("err, debugger already unhide (%x)\n",
                     Error);
        break;
    case DEBUGGER_ERROR_EDIT_MEMORY_STATUS_INVALID_PARAMETER:
        ShowMessages("err, edit memeory request has invalid parameters (%x)\n",
                     Error);
        break;
    case DEBUGGER_ERROR_EDIT_MEMORY_STATUS_INVALID_ADDRESS_BASED_ON_CURRENT_PROCESS:
        ShowMessages("err, edit memory request has invalid address based on "
                     "current process layout, the address might be valid but not "
                     "present in the ram (%x)\n",
                     Error);
        break;
    case DEBUGGER_ERROR_EDIT_MEMORY_STATUS_INVALID_ADDRESS_BASED_ON_OTHER_PROCESS:
        ShowMessages("err, edit memory request has invalid address based on other "
                     "process layout (%x)\n",
                     Error);
        break;
    case DEBUGGER_ERROR_MODIFY_EVENTS_INVALID_TAG:
        ShowMessages("err, modify event with invalid tag (%x)\n",
                     Error);
        break;
    case DEBUGGER_ERROR_MODIFY_EVENTS_INVALID_TYPE_OF_ACTION:
        ShowMessages("err, modify event with invalid type of action (%x)\n",
                     Error);
        break;
    case DEBUGGER_ERROR_STEPPING_INVALID_PARAMETER:
        ShowMessages("err, invalid parameter passsed to stepping core. (%x)\n",
                     Error);
        break;
    case DEBUGGER_ERROR_STEPPINGS_EITHER_THREAD_NOT_FOUND_OR_DISABLED:
        ShowMessages(
            "err, the target thread not found or the thread is disabled (%x)\n",
            Error);
        break;
    case DEBUGGER_ERROR_PREPARING_DEBUGGEE_INVALID_BAUDRATE:
        ShowMessages("err, invalid baud rate (%x)\n",
                     Error);
        break;
    case DEBUGGER_ERROR_PREPARING_DEBUGGEE_INVALID_SERIAL_PORT:
        ShowMessages("err, invalid serial port (%x)\n",
                     Error);
        break;
    case DEBUGGER_ERROR_PREPARING_DEBUGGEE_INVALID_CORE_IN_REMOTE_DEBUGGE:
        ShowMessages("err, invalid core selected in switching cores (%x)\n",
                     Error);
        break;
    case DEBUGGER_ERROR_PREPARING_DEBUGGEE_UNABLE_TO_SWITCH_TO_NEW_PROCESS:
        ShowMessages("err, unable to switch to the new process (%x)\n",
                     Error);
        break;
    case DEBUGGER_ERROR_PREPARING_DEBUGGEE_TO_RUN_SCRIPT:
        ShowMessages("err, unable to run the script on remote debuggee (%x)\n",
                     Error);
        break;
    case DEBUGGER_ERROR_INVALID_REGISTER_NUMBER:
        ShowMessages("err, invalid register number (%x)\n",
                     Error);
        break;
    case DEBUGGER_ERROR_MAXIMUM_BREAKPOINT_WITHOUT_CONTINUE:
        ShowMessages("err, maximum number of breakpoints are used, you need to "
                     "send an ioctl to re-allocate new pre-allocated buffers or "
                     "configure HyperDbg to pre-allocated more buffers by "
                     "configuring MAXIMUM_BREAKPOINTS_WITHOUT_CONTINUE (%x)\n",
                     Error);
        break;
    case DEBUGGER_ERROR_BREAKPOINT_ALREADY_EXISTS_ON_THE_ADDRESS:
        ShowMessages("err, breakpoint already exists on target address, you can "
                     "use 'bl' command to view a list of breakpoints (%x)\n",
                     Error);
        break;
    case DEBUGGER_ERROR_BREAKPOINT_ID_NOT_FOUND:
        ShowMessages("err, breakpoint id is invalid (%x)\n",
                     Error);
        break;
    case DEBUGGER_ERROR_BREAKPOINT_ALREADY_DISABLED:
        ShowMessages("err, breakpoint already disabled (%x)\n",
                     Error);
        break;
    case DEBUGGER_ERROR_BREAKPOINT_ALREADY_ENABLED:
        ShowMessages("err, breakpoint already enabled (%x)\n",
                     Error);
        break;
    case DEBUGGER_ERROR_MEMORY_TYPE_INVALID:
        ShowMessages("err, the memory type is invalid (%x)\n",
                     Error);
        break;
    case DEBUGGER_ERROR_INVALID_PROCESS_ID:
        ShowMessages("err, the process id is invalid, make sure to enter the "
                     "process id in hex format, or if you want to use it in decimal "
                     "format, add '0n' prefix to the number (%x)\n",
                     Error);
        break;
    case DEBUGGER_ERROR_EVENT_IS_NOT_APPLIED:
        ShowMessages("err, the event is not applied (%x)\n",
                     Error);
        break;
    case DEBUGGER_ERROR_DETAILS_OR_SWITCH_PROCESS_INVALID_PARAMETER:
        ShowMessages("err, either the process id or the _EPROCESS is invalid or "
                     "cannot get the details based on the provided parameters (%x)\n",
                     Error);
        break;
    case DEBUGGER_ERROR_DETAILS_OR_SWITCH_THREAD_INVALID_PARAMETER:
        ShowMessages("err, either the thread id, _ETHREAD, or _EPROCESS is invalid or "
                     "cannot get the details based on the provided parameters (%x)\n",
                     Error);
        break;
    case DEBUGGER_ERROR_MAXIMUM_BREAKPOINT_FOR_A_SINGLE_PAGE_IS_HIT:
        ShowMessages("err, the maximum breakpoint for a single page is hit, "
                     "you cannot apply more breakpoints in this page (%x)\n",
                     Error);
        break;
    case DEBUGGER_ERROR_PRE_ALLOCATED_BUFFER_IS_EMPTY:
        ShowMessages("err, the pre-allocated buffer is empty, usually this buffer will be "
                     "filled at the next IOCTL when the debugger is continued (%x)\n"
                     "please visit the documentation for the 'prealloc' command or use "
                     "'.help prealloc' to to reserve more pre-allocated pools\n",
                     Error);
        break;
    case DEBUGGER_ERROR_EPT_COULD_NOT_SPLIT_THE_LARGE_PAGE_TO_4KB_PAGES:
        ShowMessages("err, could not convert 2MB large page to 4KB pages (%x)\n",
                     Error);
        break;
    case DEBUGGER_ERROR_EPT_FAILED_TO_GET_PML1_ENTRY_OF_TARGET_ADDRESS:
        ShowMessages("err, failed to get the PML1 entry of the target address (%x)\n",
                     Error);
        break;
    case DEBUGGER_ERROR_EPT_MULTIPLE_HOOKS_IN_A_SINGLE_PAGE:
        ShowMessages("err, the page modification is not applied, make sure that you don't "
                     "put multiple EPT Hooks or Monitors on a single page (%x)\n",
                     Error);
        break;
    case DEBUGGER_ERROR_COULD_NOT_BUILD_THE_EPT_HOOK:
        ShowMessages("err, could not build the EPT hook (%x)\n",
                     Error);
        break;
    case DEBUGGER_ERROR_COULD_NOT_FIND_ALLOCATION_TYPE:
        ShowMessages("err, could not find the allocation type (%x)\n",
                     Error);
        break;
    case DEBUGGER_ERROR_INVALID_TEST_QUERY_INDEX:
        ShowMessages("err, invalid index specified for test query command (%x)\n",
                     Error);
        break;
    case DEBUGGER_ERROR_UNABLE_TO_ATTACH_TO_TARGET_USER_MODE_PROCESS:
        ShowMessages("err, unable to attach to the target process (%x)\n",
                     Error);
        break;
    case DEBUGGER_ERROR_UNABLE_TO_REMOVE_HOOKS_ENTRYPOINT_NOT_REACHED:
        ShowMessages("err, unable to remove hooks as the entrypoint of user-mode "
                     "process is not reached yet (%x)\n",
                     Error);
        break;
    case DEBUGGER_ERROR_UNABLE_TO_REMOVE_HOOKS:
        ShowMessages("err, unable to remove hooks (%x)\n",
                     Error);
        break;
    case DEBUGGER_ERROR_FUNCTIONS_FOR_INITIALIZING_PEB_ADDRESSES_ARE_NOT_INITIALIZED:
        ShowMessages("err, the routines for getting the PEB is not correctly "
                     "initialized (%x)\n",
                     Error);
        break;
    case DEBUGGER_ERROR_UNABLE_TO_DETECT_32_BIT_OR_64_BIT_PROCESS:
        ShowMessages("err, unable to detect whether the process was 32-bit "
                     "or 64-bit (%x)\n",
                     Error);
        break;
    case DEBUGGER_ERROR_UNABLE_TO_KILL_THE_PROCESS:
        ShowMessages("err, unable to kill the process (%x)\n",
                     Error);
        break;
    case DEBUGGER_ERROR_INVALID_THREAD_DEBUGGING_TOKEN:
        ShowMessages("err, invalid thread debugging token (%x)\n",
                     Error);
        break;
    case DEBUGGER_ERROR_UNABLE_TO_PAUSE_THE_PROCESS_THREADS:
        ShowMessages("err, unable to pause the threads of the process (%x)\n",
                     Error);
        break;
    case DEBUGGER_ERROR_UNABLE_TO_ATTACH_TO_AN_ALREADY_ATTACHED_PROCESS:
        ShowMessages("err, the user debugger is already attached to this "
                     "process, please use the '.switch' command to switch "
                     "to this process (%x)\n",
                     Error);
        break;
    case DEBUGGER_ERROR_THE_USER_DEBUGGER_NOT_ATTACHED_TO_THE_PROCESS:
        ShowMessages("err, the user debugger is not already attached to "
                     "the process (%x)\n",
                     Error);
        break;
    case DEBUGGER_ERROR_UNABLE_TO_DETACH_AS_THERE_ARE_PAUSED_THREADS:
        ShowMessages("err, the user debugger is not able to detach from "
                     "this process as there are paused threads in the "
                     "target process, please make sure to remove all "
                     "the events and continue the target process, then "
                     "perform the detach again (%x)\n",
                     Error);
        break;
    case DEBUGGER_ERROR_UNABLE_TO_SWITCH_PROCESS_ID_OR_THREAD_ID_IS_INVALID:
        ShowMessages("err, unable to switch to the process id or thread id "
                     "as the target process id or thread id is not found in "
                     "the attached threads list, please view the list of "
                     "processes and threads by using the '.switch list' command (%x)\n",
                     Error);
        break;
    case DEBUGGER_ERROR_UNABLE_TO_SWITCH_THERE_IS_NO_THREAD_ON_THE_PROCESS:
        ShowMessages("err, unable to switch to the process as the process doesn't "
                     "contain an active intercepted thread (%x)\n",
                     Error);
        break;
    case DEBUGGER_ERROR_UNABLE_TO_GET_MODULES_OF_THE_PROCESS:
        ShowMessages("err, unable to get user-mode modules of the process (%x)\n",
                     Error);
        break;
    case DEBUGGER_ERROR_UNABLE_TO_GET_CALLSTACK:
        ShowMessages("err, unable to get the callstack (%x)\n",
                     Error);
        break;
    case DEBUGGER_ERROR_UNABLE_TO_QUERY_COUNT_OF_PROCESSES_OR_THREADS:
        ShowMessages("err, unable to query count of processes or threads (%x)\n",
                     Error);
        break;
    default:
        ShowMessages("err, error not found (%x)\n",
                     Error);
        return FALSE;
        break;
    }
    return TRUE;
}*/
return true
}

func (d *debugger)DebuggerGetNtoskrnlBase()(ok bool){//col:408
/*DebuggerGetNtoskrnlBase()
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
    return NtoskrnlBase;
}*/
return true
}

func (d *debugger) * then when it returns true the debuggee is not paused anymore ()(ok bool){//col:458
/* * then when it returns true the debuggee is not paused anymore (continued)
BOOLEAN
DebuggerPauseDebuggee()
{
    BOOLEAN                        StatusIoctl    = 0;
    ULONG                          ReturnedLength = 0;
    DEBUGGER_PAUSE_PACKET_RECEIVED PauseRequest   = {0};
    StatusIoctl =
        );
    if (!StatusIoctl)
    {
        ShowMessages("ioctl failed with code 0x%x\n", GetLastError());
        return FALSE;
    }
    if (PauseRequest.Result == DEBUGGER_OPERATION_WAS_SUCCESSFULL)
    {
        return TRUE;
    }
    else
    {
        ShowErrorMessage(PauseRequest.Result);
        return FALSE;
    }
    return FALSE;
}*/
return true
}

func (d *debugger)IsConnectedToAnyInstanceOfDebuggerOrDebuggee()(ok bool){//col:509
/*IsConnectedToAnyInstanceOfDebuggerOrDebuggee()
{
    if (g_DeviceHandle)
    {
        ShowMessages("err, the current system is already connected to the local "
                     "debugging, use '.disconnect' to disconnect\n");
        return TRUE;
    }
    else if (g_IsConnectedToRemoteDebuggee)
    {
        ShowMessages("err, the current system is already connected to remote "
                     "machine (debuggee), use '.disconnect' to disconnect from the "
                     "remote machine\n");
        return TRUE;
    }
    else if (g_IsConnectedToRemoteDebugger)
    {
        ShowMessages("err, the current system is already connected to remote "
                     "machine (debugger), use '.disconnect' to disconnect from the "
                     "remote machine from debugger\n");
        return TRUE;
    }
    else if (g_IsSerialConnectedToRemoteDebuggee)
    {
        ShowMessages(
            "err, the current system is already connected to remote "
            "machine (debuggee), use '.debug close' to disconnect from the "
            "remote machine\n");
        return TRUE;
    }
    else if (g_IsSerialConnectedToRemoteDebugger)
    {
        ShowMessages(
            "err, the current system is already connected to remote "
            "machine (debugger), use '.debug close' to disconnect from the "
            "remote machine from debugger\n");
        return TRUE;
    }
    return FALSE;
}*/
return true
}

func (d *debugger)IsTagExist()(ok bool){//col:548
/*IsTagExist(UINT64 Tag)
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
}*/
return true
}

func (d *debugger) * successful ()(ok bool){//col:1034
/* * successful (false)
BOOLEAN
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
                    IsEnded = TRUE;
                    break;
                }
                if (HasEnding(Section, "}"))
                {
                    IndexesToRemove.push_back(Index);
                    TempStr = SplittedCommandCaseSensitiveInstance.at(IndexInCommandCaseSensitive - 1);
                    SaveBuffer.emplace_back(TempStr.begin(), TempStr.begin() + TempStr.size() - 1);
                    IsEnded = TRUE;
                    break;
                }
            }
            IndexesToRemove.push_back(Index);
            if (Section.find('{') != string::npos)
            {
                size_t CountOfBrackets = count(Section.begin(), Section.end(), '{');
                OpenBracket += CountOfBrackets;
            }
            if (Section.find('}') != string::npos)
            {
                size_t CountOfBrackets = count(Section.begin(), Section.end(), '}');
                OpenBracket -= CountOfBrackets;
                if (OpenBracket < 0)
                {
                    OpenBracket = 0;
                    IsEnded     = TRUE;
                    TempStr     = SplittedCommandCaseSensitiveInstance.at(IndexInCommandCaseSensitive - 1);
                    SaveBuffer.emplace_back(TempStr.begin(), TempStr.begin() + TempStr.size() - 1);
                    break;
                }
            }
            SaveBuffer.push_back(SplittedCommandCaseSensitiveInstance.at(IndexInCommandCaseSensitive - 1));
            continue;
        }
        if (IsTextVisited && !Section.compare("{"))
        {
            IndexesToRemove.push_back(Index);
            IsInState = TRUE;
            continue;
        }
        if (IsTextVisited && Section.rfind('{', 0) == 0)
        {
            if (Section.find('{') != string::npos)
            {
                size_t CountOfBrackets = count(Section.begin(), Section.end(), '{');
                OpenBracket += CountOfBrackets - 1;
            }
            if (OpenBracket == 0 && HasEnding(Section, "}"))
            {
                IndexesToRemove.push_back(Index);
                TempStr = SplittedCommandCaseSensitiveInstance.at(IndexInCommandCaseSensitive - 1).erase(0, 1);
                SaveBuffer.emplace_back(TempStr.begin(), TempStr.begin() + TempStr.size() - 1);
                IsEnded = TRUE;
                break;
            }
            if (Section.find('}') != string::npos)
            {
                size_t CountOfBrackets = count(Section.begin(), Section.end(), '}');
                OpenBracket -= CountOfBrackets;
                if (OpenBracket < 0)
                {
                    OpenBracket = 0;
                    IsEnded     = TRUE;
                    TempStr     = SplittedCommandCaseSensitiveInstance.at(IndexInCommandCaseSensitive - 1);
                    SaveBuffer.emplace_back(TempStr.begin(), TempStr.begin() + TempStr.size() - 1);
                    break;
                }
            }
            IndexesToRemove.push_back(Index);
            SaveBuffer.push_back(SplittedCommandCaseSensitiveInstance.at(IndexInCommandCaseSensitive - 1).erase(0, 1));
            IsInState = TRUE;
            continue;
        }
        if (!Section.compare("script"))
        {
            IndexesToRemove.push_back(Index);
            IsTextVisited = TRUE;
            continue;
        }
        if (!Section.compare("script{"))
        {
            IndexesToRemove.push_back(Index);
            IsTextVisited = TRUE;
            IsInState     = TRUE;
            continue;
        }
        if (Section.rfind("script{", 0) == 0)
        {
            IndexesToRemove.push_back(Index);
            IsTextVisited        = TRUE;
            IsInState            = TRUE;
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
                IsEnded     = TRUE;
                OpenBracket = 0;
                break;
            }
            else
            {
                SaveBuffer.push_back(SplittedCommandCaseSensitiveInstance.at(IndexInCommandCaseSensitive - 1).erase(0, 7));
                if (Section.find('}') != string::npos)
                {
                    size_t CountOfBrackets = count(Section.begin(), Section.end(), '}');
                    OpenBracket -= CountOfBrackets;
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
    }
    if (AppendedFinalBuffer.rfind("file:", 0) == 0)
    {
        std::ifstream     t(AppendedFinalBuffer.erase(0, 5).c_str());
        std::stringstream buffer;
        buffer << t.rdbuf();
        AppendedFinalBuffer = buffer.str();
        if (AppendedFinalBuffer.empty())
        {
            ShowMessages("err, either script file is not found or it's empty\n");
            *ScriptSyntaxErrors = TRUE;
            return TRUE;
        }
    }
    PVOID CodeBuffer =
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
    }
    return TRUE;
}*/
return true
}

func (d *debugger) * successful ()(ok bool){//col:1412
/* * successful (false)
BOOLEAN
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
                IsEnded = TRUE;
                break;
            }
            if (HasEnding(Section, "}"))
            {
                IndexesToRemove.push_back(Index);
                SaveBuffer.emplace_back(Section.begin(), Section.begin() + Section.size() - 1);
                IsEnded = TRUE;
                break;
            }
            IndexesToRemove.push_back(Index);
            SaveBuffer.push_back(Section);
            continue;
        }
        if (IsTextVisited && !Section.compare("{"))
        {
            IndexesToRemove.push_back(Index);
            IsInState = TRUE;
            continue;
        }
        if (IsTextVisited && Section.rfind('{', 0) == 0)
        {
            if (HasEnding(Section, "}"))
            {
                IndexesToRemove.push_back(Index);
                TempStr = Section.erase(0, 1);
                SaveBuffer.emplace_back(TempStr.begin(), TempStr.begin() + TempStr.size() - 1);
                IsEnded = TRUE;
                break;
            }
            IndexesToRemove.push_back(Index);
            SaveBuffer.push_back(Section.erase(0, 1));
            IsInState = TRUE;
            continue;
        }
        if (IsConditionBuffer)
        {
            if (!Section.compare("condition"))
            {
                IndexesToRemove.push_back(Index);
                IsTextVisited = TRUE;
                continue;
            }
        }
        else
        {
            if (!Section.compare("code"))
            {
                IndexesToRemove.push_back(Index);
                IsTextVisited = TRUE;
                continue;
            }
        }
        if (IsConditionBuffer)
        {
            if (!Section.compare("condition{"))
            {
                IndexesToRemove.push_back(Index);
                IsTextVisited = TRUE;
                IsInState     = TRUE;
                continue;
            }
        }
        else
        {
            if (!Section.compare("code{"))
            {
                IndexesToRemove.push_back(Index);
                IsTextVisited = TRUE;
                IsInState     = TRUE;
                continue;
            }
        }
        if (IsConditionBuffer)
        {
            if (Section.rfind("condition{", 0) == 0)
            {
                IndexesToRemove.push_back(Index);
                IsTextVisited = TRUE;
                IsInState     = TRUE;
                if (!HasEnding(Section, "}"))
                {
                    SaveBuffer.push_back(Section.erase(0, 10));
                    continue;
                }
                else
                {
                    TempStr = Section.erase(0, 10);
                    SaveBuffer.emplace_back(TempStr.begin(), TempStr.begin() + TempStr.size() - 1);
                    IsEnded = TRUE;
                    break;
                }
            }
        }
        else
        {
            if (Section.rfind("code{", 0) == 0)
            {
                IndexesToRemove.push_back(Index);
                IsTextVisited = TRUE;
                IsInState     = TRUE;
                if (!HasEnding(Section, "}"))
                {
                    SaveBuffer.push_back(Section.erase(0, 5));
                    continue;
                }
                else
                {
                    TempStr = Section.erase(0, 5);
                    SaveBuffer.emplace_back(TempStr.begin(), TempStr.begin() + TempStr.size() - 1);
                    IsEnded = TRUE;
                    break;
                }
            }
        }
    }
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
        }
        else if (Section.rfind('x', 0) == 0 || Section.rfind('X', 0) == 0)
        {
            Temp = Section.erase(0, 1);
        }
        else
        {
            Temp = std::move(Section);
        }
        ReplaceAll(Temp, "\\x", "");
        if (Temp.size() % 2 != 0)
        {
            Temp.insert(0, 1, '0');
        }
        if (!IsHexNotation(Temp))
        {
            ShowMessages("please enter condition code in a hex notation\n");
            return FALSE;
        }
        AppendedFinalBuffer.append(Temp);
    }
    ParsedBytes = HexToBytes(AppendedFinalBuffer);
    FinalBuffer = (unsigned char *)malloc(ParsedBytes.size());
    std::copy(ParsedBytes.begin(), ParsedBytes.end(), FinalBuffer);
    *BufferAddress = (UINT64)FinalBuffer;
    *BufferLength  = ParsedBytes.size();
    NewIndexToRemove = 0;
    for (auto IndexToRemove : IndexesToRemove)
    {
        NewIndexToRemove++;
        SplittedCommand->erase(SplittedCommand->begin() +
                               (IndexToRemove - NewIndexToRemove));
        SplittedCommandCaseSensitive->erase(SplittedCommandCaseSensitive->begin() +
                                            (IndexToRemove - NewIndexToRemove));
    }
    return TRUE;
}*/
return true
}

func (d *debugger) * input that needs to be considered for this event ()(ok bool){//col:1687
/* * input that needs to be considered for this event (other than just printing)
 * like sending over network, save to file, and send over a namedpipe
 *
 * by space case sensitive
 * it
 * successful (false)
BOOLEAN
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
                IsEnded = TRUE;
                break;
            }
            if (HasEnding(Section, "}"))
            {
                IndexesToRemove.push_back(Index);
                TempStr = SplittedCommandCaseSensitiveInstance.at(IndexInCommandCaseSensitive - 1);
                SaveBuffer.emplace_back(TempStr.begin(), TempStr.begin() + TempStr.size() - 1);
                IsEnded = TRUE;
                break;
            }
            IndexesToRemove.push_back(Index);
            SaveBuffer.push_back(SplittedCommandCaseSensitiveInstance.at(IndexInCommandCaseSensitive - 1));
            continue;
        }
        if (IsTextVisited && !Section.compare("{"))
        {
            IndexesToRemove.push_back(Index);
            IsInState = TRUE;
            continue;
        }
        if (IsTextVisited && Section.rfind('{', 0) == 0)
        {
            if (HasEnding(Section, "}"))
            {
                IndexesToRemove.push_back(Index);
                TempStr = SplittedCommandCaseSensitiveInstance.at(IndexInCommandCaseSensitive - 1).erase(0, 1);
                SaveBuffer.emplace_back(TempStr.begin(), TempStr.begin() + TempStr.size() - 1);
                IsEnded = TRUE;
                break;
            }
            IndexesToRemove.push_back(Index);
            SaveBuffer.push_back(SplittedCommandCaseSensitiveInstance.at(IndexInCommandCaseSensitive - 1).erase(0, 1));
            IsInState = TRUE;
            continue;
        }
        if (!Section.compare("output"))
        {
            IndexesToRemove.push_back(Index);
            IsTextVisited = TRUE;
            continue;
        }
        if (!Section.compare("output{"))
        {
            IndexesToRemove.push_back(Index);
            IsTextVisited = TRUE;
            IsInState     = TRUE;
            continue;
        }
        if (Section.rfind("output{", 0) == 0)
        {
            IndexesToRemove.push_back(Index);
            IsTextVisited = TRUE;
            IsInState     = TRUE;
            if (!HasEnding(Section, "}"))
            {
                SaveBuffer.push_back(SplittedCommandCaseSensitiveInstance.at(IndexInCommandCaseSensitive - 1).erase(0, 7));
                continue;
            }
            else
            {
                TempStr = SplittedCommandCaseSensitiveInstance.at(IndexInCommandCaseSensitive - 1).erase(0, 7);
                SaveBuffer.emplace_back(TempStr.begin(), TempStr.begin() + TempStr.size() - 1);
                IsEnded = TRUE;
                break;
            }
        }
    }
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
    }
    if (AppendedFinalBuffer.find(Delimiter) != std::string::npos)
    {
        while ((Pos = AppendedFinalBuffer.find(Delimiter)) != string::npos)
        {
            Token = AppendedFinalBuffer.substr(0, Pos);
            Trim(Token);
            if (!Token.empty())
            {
                InputSources.push_back(Token);
            }
            AppendedFinalBuffer.erase(0, Pos + sizeof(Delimiter) / sizeof(char));
        }
        if (!AppendedFinalBuffer.empty())
        {
            InputSources.push_back(AppendedFinalBuffer);
        }
    }
    else
    {
        InputSources.push_back(AppendedFinalBuffer);
    }
    NewIndexToRemove = 0;
    for (auto IndexToRemove : IndexesToRemove)
    {
        NewIndexToRemove++;
        SplittedCommand->erase(SplittedCommand->begin() +
                               (IndexToRemove - NewIndexToRemove));
        SplittedCommandCaseSensitive->erase(SplittedCommandCaseSensitive->begin() +
                                            (IndexToRemove - NewIndexToRemove));
    }
    return TRUE;
}*/
return true
}

func (d *debugger)SendEventToKernel()(ok bool){//col:1803
/*SendEventToKernel(PDEBUGGER_GENERAL_EVENT_DETAIL Event,
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
    }
    else
    {
        AssertShowMessageReturnStmt(g_DeviceHandle, ASSERT_MESSAGE_DRIVER_NOT_LOADED, AssertReturnFalse);
        Status =
            );
        if (!Status)
        {
            ShowMessages("ioctl failed with code 0x%x\n", GetLastError());
            return FALSE;
        }
    }
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
            }
            ShowMessages("\n");
        }
    }
    else
    {
        if (ReturnedBuffer.Error != 0)
        {
            ShowErrorMessage(ReturnedBuffer.Error);
        }
        return FALSE;
    }
    return TRUE;
}*/
return true
}

func (d *debugger)RegisterActionToEvent()(ok bool){//col:2005
/*RegisterActionToEvent(PDEBUGGER_GENERAL_EVENT_DETAIL Event,
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
        }
        if (ActionCustomCode != NULL)
        {
            TempAddingResult = KdSendAddActionToEventPacketToDebuggee(
                ActionCustomCode,
                ActionCustomCodeLength);
            memcpy(&ReturnedBuffer, TempAddingResult, sizeof(DEBUGGER_EVENT_AND_ACTION_REG_BUFFER));
        }
        if (ActionScript != NULL)
        {
            TempAddingResult = KdSendAddActionToEventPacketToDebuggee(
                ActionScript,
                ActionScriptLength);
            memcpy(&ReturnedBuffer, TempAddingResult, sizeof(DEBUGGER_EVENT_AND_ACTION_REG_BUFFER));
        }
    }
    else
    {
        AssertShowMessageReturnStmt(g_DeviceHandle, ASSERT_MESSAGE_DRIVER_NOT_LOADED, AssertReturnFalse);
        if (ActionBreakToDebugger != NULL)
        {
            Status = DeviceIoControl(
            );
            if (!Status)
            {
                ShowMessages("ioctl failed with code 0x%x\n", GetLastError());
                return FALSE;
            }
        }
        if (ActionCustomCode != NULL)
        {
            Status = DeviceIoControl(
            );
            if (!Status)
            {
                ShowMessages("ioctl failed with code 0x%x\n", GetLastError());
                return FALSE;
            }
        }
        if (ActionScript != NULL)
        {
            Status = DeviceIoControl(
            );
            if (!Status)
            {
                ShowMessages("ioctl failed with code 0x%x\n", GetLastError());
                return FALSE;
            }
        }
    }
    FreeEventsAndActionsMemory(NULL, ActionBreakToDebugger, ActionCustomCode, ActionScript);
    InsertHeadList(&g_EventTrace, &(Event->CommandsEventList));
    return TRUE;
}*/
return true
}

func (d *debugger)GetNewDebuggerEventTag()(ok bool){//col:2017
/*GetNewDebuggerEventTag()
{
    return g_EventTag++;
}*/
return true
}

func (d *debugger)FreeEventsAndActionsMemory()(ok bool){//col:2053
/*FreeEventsAndActionsMemory(PDEBUGGER_GENERAL_EVENT_DETAIL Event,
                           PDEBUGGER_GENERAL_ACTION       ActionBreakToDebugger,
                           PDEBUGGER_GENERAL_ACTION       ActionCustomCode,
                           PDEBUGGER_GENERAL_ACTION       ActionScript)
{
    if (Event != NULL)
    {
        if (Event->CommandStringBuffer != NULL)
        {
            free(Event->CommandStringBuffer);
        }
        free(Event);
    }
    if (ActionBreakToDebugger != NULL)
    {
        free(ActionBreakToDebugger);
    }
    if (ActionCustomCode != NULL)
    {
        free(ActionCustomCode);
    }
    if (ActionScript != NULL)
    {
        free(ActionScript);
    }
}*/
return true
}

func (d *debugger)InterpretGeneralEventAndActionsFields()(ok bool){//col:3002
/*InterpretGeneralEventAndActionsFields(
    vector<string> *                    SplittedCommand,
    vector<string> *                    SplittedCommandCaseSensitive,
    DEBUGGER_EVENT_TYPE_ENUM            EventType,
    PDEBUGGER_GENERAL_EVENT_DETAIL *    EventDetailsToFill,
    PUINT32                             EventBufferLength,
    PDEBUGGER_GENERAL_ACTION *          ActionDetailsToFillBreakToDebugger,
    PUINT32                             ActionBufferLengthBreakToDebugger,
    PDEBUGGER_GENERAL_ACTION *          ActionDetailsToFillCustomCode,
    PUINT32                             ActionBufferLengthCustomCode,
    PDEBUGGER_GENERAL_ACTION *          ActionDetailsToFillScript,
    PUINT32                             ActionBufferLengthScript,
    PDEBUGGER_EVENT_PARSING_ERROR_CAUSE ReasonForErrorInParsing)
{
    BOOLEAN                        Result                         = FALSE;
    PDEBUGGER_GENERAL_EVENT_DETAIL TempEvent                      = NULL;
    PDEBUGGER_GENERAL_ACTION       TempActionBreak                = NULL;
    PDEBUGGER_GENERAL_ACTION       TempActionScript               = NULL;
    PDEBUGGER_GENERAL_ACTION       TempActionCustomCode           = NULL;
    UINT32                         LengthOfCustomCodeActionBuffer = 0;
    UINT32                         LengthOfScriptActionBuffer     = 0;
    UINT32                         LengthOfBreakActionBuffer      = 0;
    UINT64                         ConditionBufferAddress;
    UINT32                         ConditionBufferLength = 0;
    vector<string>                 ListOfOutputSources;
    UINT64                         CodeBufferAddress;
    UINT32                         CodeBufferLength = 0;
    UINT64                         ScriptBufferAddress;
    UINT64                         ScriptCodeBuffer     = 0;
    BOOLEAN                        HasScriptSyntaxError = 0;
    UINT32                         ScriptBufferLength   = 0;
    UINT32                         ScriptBufferPointer  = 0;
    UINT32                         LengthOfEventBuffer  = 0;
    string                         CommandString;
    BOOLEAN                        HasConditionBuffer              = FALSE;
    BOOLEAN                        HasOutputPath                   = FALSE;
    BOOLEAN                        HasCodeBuffer                   = FALSE;
    BOOLEAN                        HasScript                       = FALSE;
    BOOLEAN                        IsNextCommandPid                = FALSE;
    BOOLEAN                        IsNextCommandCoreId             = FALSE;
    BOOLEAN                        IsNextCommandBufferSize         = FALSE;
    BOOLEAN                        IsNextCommandImmediateMessaging = FALSE;
    BOOLEAN                        ImmediateMessagePassing         = UseImmediateMessagingByDefaultOnEvents;
    UINT32                         CoreId;
    UINT32                         ProcessId;
    UINT32                         IndexOfValidSourceTags;
    UINT32                         RequestBuffer = 0;
    PLIST_ENTRY                    TempList;
    BOOLEAN                        OutputSourceFound;
    vector<int>                    IndexesToRemove;
    vector<UINT64>                 ListOfValidSourceTags;
    int                            NewIndexToRemove = 0;
    int                            Index            = 0;
    for (auto Section : *SplittedCommandCaseSensitive)
    {
        CommandString.append(Section);
        CommandString.append(" ");
    }
    UINT64 BufferOfCommandStringLength = CommandString.size() + 1;
    PVOID BufferOfCommandString = malloc(BufferOfCommandStringLength);
    RtlZeroMemory(BufferOfCommandString, BufferOfCommandStringLength);
    memcpy(BufferOfCommandString, CommandString.c_str(), CommandString.size());
    if (!InterpretConditionsAndCodes(SplittedCommand, SplittedCommandCaseSensitive, TRUE, &ConditionBufferAddress, &ConditionBufferLength))
    {
        HasConditionBuffer = FALSE;
    }
    else
    {
        HasConditionBuffer = TRUE;
    ShowMessages(
        "\n========================= Condition =========================\n");
    ShowMessages(
        "\nUINT64  DebuggerCheckForCondition(PGUEST_REGS Regs_RCX, PVOID "
        "Context_RDX)\n{\n");
    HyperDbgDisassembler64((unsigned char *)ConditionBufferAddress, 0x0,
                           ConditionBufferLength);
    ShowMessages("}\n\n");
    ShowMessages(
        "=============================================================\n");
    }
    if (!InterpretConditionsAndCodes(SplittedCommand, SplittedCommandCaseSensitive, FALSE, &CodeBufferAddress, &CodeBufferLength))
    {
        HasCodeBuffer = FALSE;
    }
    else
    {
        HasCodeBuffer = TRUE;
    ShowMessages(
        "\n=========================    Code    =========================\n");
    ShowMessages("\nPVOID DebuggerRunCustomCodeFunc(PVOID "
                 "PreAllocatedBufferAddress_RCX, "
                 "PGUEST_REGS Regs_RDX, PVOID Context_R8)\n{\n");
    HyperDbgDisassembler64((unsigned char *)CodeBufferAddress, 0x0,
                           CodeBufferLength);
    ShowMessages("}\n\n");
    ShowMessages(
        "=============================================================\n");
    }
    if (!InterpretScript(SplittedCommand,
                         SplittedCommandCaseSensitive,
                         &HasScriptSyntaxError,
                         &ScriptBufferAddress,
                         &ScriptBufferLength,
                         &ScriptBufferPointer,
                         &ScriptCodeBuffer))
    {
        HasScript = FALSE;
    }
    else
    {
        if (HasScriptSyntaxError)
        {
            free(BufferOfCommandString);
            *ReasonForErrorInParsing = DEBUGGER_EVENT_PARSING_ERROR_CAUSE_SCRIPT_SYNTAX_ERROR;
            return FALSE;
        }
        HasScript = TRUE;
    }
    if (!InterpretOutput(SplittedCommand, SplittedCommandCaseSensitive, ListOfOutputSources))
    {
        HasOutputPath = FALSE;
    }
    else
    {
        if (ListOfOutputSources.size() == 0)
        {
            free(BufferOfCommandString);
            ShowMessages("err, no input found\n");
            *ReasonForErrorInParsing = DEBUGGER_EVENT_PARSING_ERROR_CAUSE_NO_INPUT;
            return FALSE;
        }
        if (ListOfOutputSources.size() >
            DebuggerOutputSourceMaximumRemoteSourceForSingleEvent)
        {
            free(BufferOfCommandString);
            ShowMessages(
                "err, based on this build of HyperDbg, the maximum input sources for "
                "a single event is 0x%x sources but you entered 0x%x sources\n",
                DebuggerOutputSourceMaximumRemoteSourceForSingleEvent,
                ListOfOutputSources.size());
            *ReasonForErrorInParsing = DEBUGGER_EVENT_PARSING_ERROR_CAUSE_MAXIMUM_INPUT_REACHED;
            return FALSE;
        }
        if (!g_OutputSourcesInitialized)
        {
            free(BufferOfCommandString);
            ShowMessages("err, the name you entered, not found. Did you use "
                         "'output' commmand to create it?\n");
            *ReasonForErrorInParsing = DEBUGGER_EVENT_PARSING_ERROR_CAUSE_OUTPUT_NAME_NOT_FOUND;
            return FALSE;
        }
        for (auto item : ListOfOutputSources)
        {
            TempList          = 0;
            OutputSourceFound = FALSE;
            TempList = &g_OutputSources;
            while (&g_OutputSources != TempList->Flink)
            {
                TempList = TempList->Flink;
                PDEBUGGER_EVENT_FORWARDING CurrentOutputSourceDetails =
                    CONTAINING_RECORD(TempList, DEBUGGER_EVENT_FORWARDING, OutputSourcesList);
                if (strcmp(CurrentOutputSourceDetails->Name,
                           RemoveSpaces(item).c_str()) == 0)
                {
                    if (CurrentOutputSourceDetails->State == EVENT_FORWARDING_CLOSED)
                    {
                        free(BufferOfCommandString);
                        ShowMessages("err, output source already closed\n");
                        *ReasonForErrorInParsing = DEBUGGER_EVENT_PARSING_ERROR_CAUSE_OUTPUT_SOURCE_ALREADY_CLOSED;
                        return FALSE;
                    }
                    OutputSourceFound = TRUE;
                    ListOfValidSourceTags.push_back(
                        CurrentOutputSourceDetails->OutputUniqueTag);
                    break;
                }
            }
            if (!OutputSourceFound)
            {
                free(BufferOfCommandString);
                ShowMessages("err, the name you entered, not found. Did you use "
                             "'output' commmand to create it?\n");
                *ReasonForErrorInParsing = DEBUGGER_EVENT_PARSING_ERROR_CAUSE_OUTPUT_NAME_NOT_FOUND;
                return FALSE;
            }
        }
        HasOutputPath = TRUE;
    }
  Layout of Buffer :
   ________________________________
  |                                |
  |  DEBUGGER_GENERAL_EVENT_DETAIL |
  |                                |
  |________________________________|
  |                                |
  |       Condition Buffer         |
  |                                |
  |________________________________|
   ________________________________
  |                                |
  |     DEBUGGER_GENERAL_ACTION    |
  |                                |
  |________________________________|
  |                                |
  |     Condition Custom Code      |
  |       or Script Buffer         |
  |________________________________|
    LengthOfEventBuffer =
        sizeof(DEBUGGER_GENERAL_EVENT_DETAIL) + ConditionBufferLength;
    TempEvent = (PDEBUGGER_GENERAL_EVENT_DETAIL)malloc(LengthOfEventBuffer);
    RtlZeroMemory(TempEvent, LengthOfEventBuffer);
    if (TempEvent == NULL)
    {
        ShowMessages("err, allocation error\n");
        *ReasonForErrorInParsing = DEBUGGER_EVENT_PARSING_ERROR_CAUSE_ALLOCATION_ERROR;
        goto ReturnWithError;
    }
    TempEvent->IsEnabled = TRUE;
    TempEvent->Tag = GetNewDebuggerEventTag();
    TempEvent->CoreId = DEBUGGER_EVENT_APPLY_TO_ALL_CORES;
    if (g_ActiveProcessDebuggingState.IsActive)
    {
        ShowMessages("notice: as you're debugging a user-mode application, "
                     "this event will only trigger on your current debugging process "
                     "(pid:%x). If you want the event from the entire system, "
                     "add 'pid all' to the event\n",
                     g_ActiveProcessDebuggingState.ProcessId);
        TempEvent->ProcessId = g_ActiveProcessDebuggingState.ProcessId;
    }
    else
    {
        TempEvent->ProcessId = DEBUGGER_EVENT_APPLY_TO_ALL_PROCESSES;
    }
    TempEvent->EventType = EventType;
    TempEvent->CreationTime = time(0);
    TempEvent->CommandStringBuffer = BufferOfCommandString;
    if (HasConditionBuffer)
    {
        memcpy((PVOID)((UINT64)TempEvent + sizeof(DEBUGGER_GENERAL_EVENT_DETAIL)),
               (PVOID)ConditionBufferAddress,
               ConditionBufferLength);
        TempEvent->ConditionBufferSize = ConditionBufferLength;
    }
    if (HasCodeBuffer)
    {
        LengthOfCustomCodeActionBuffer =
            sizeof(DEBUGGER_GENERAL_ACTION) + CodeBufferLength;
        TempActionCustomCode =
            (PDEBUGGER_GENERAL_ACTION)malloc(LengthOfCustomCodeActionBuffer);
        RtlZeroMemory(TempActionCustomCode, LengthOfCustomCodeActionBuffer);
        memcpy(
            (PVOID)((UINT64)TempActionCustomCode + sizeof(DEBUGGER_GENERAL_ACTION)),
            (PVOID)CodeBufferAddress,
            CodeBufferLength);
        TempActionCustomCode->EventTag = TempEvent->Tag;
        TempActionCustomCode->ActionType = RUN_CUSTOM_CODE;
        TempActionCustomCode->CustomCodeBufferSize = CodeBufferLength;
        TempEvent->CountOfActions = TempEvent->CountOfActions + 1;
    }
    if (HasScript)
    {
        LengthOfScriptActionBuffer =
            sizeof(DEBUGGER_GENERAL_ACTION) + ScriptBufferLength;
        TempActionScript =
            (PDEBUGGER_GENERAL_ACTION)malloc(LengthOfScriptActionBuffer);
        RtlZeroMemory(TempActionScript, LengthOfScriptActionBuffer);
        memcpy((PVOID)((UINT64)TempActionScript + sizeof(DEBUGGER_GENERAL_ACTION)),
               (PVOID)ScriptBufferAddress,
               ScriptBufferLength);
        TempActionScript->EventTag = TempEvent->Tag;
        TempActionScript->ActionType = RUN_SCRIPT;
        TempActionScript->ScriptBufferSize    = ScriptBufferLength;
        TempActionScript->ScriptBufferPointer = ScriptBufferPointer;
        TempEvent->CountOfActions = TempEvent->CountOfActions + 1;
        ScriptEngineWrapperRemoveSymbolBuffer((PVOID)ScriptCodeBuffer);
    }
    if (!HasCodeBuffer && !HasScript)
    {
        LengthOfBreakActionBuffer = sizeof(DEBUGGER_GENERAL_ACTION);
        TempActionBreak =
            (PDEBUGGER_GENERAL_ACTION)malloc(LengthOfBreakActionBuffer);
        RtlZeroMemory(TempActionBreak, LengthOfBreakActionBuffer);
        TempActionBreak->EventTag = TempEvent->Tag;
        TempActionBreak->ActionType = BREAK_TO_DEBUGGER;
        TempEvent->CountOfActions = TempEvent->CountOfActions + 1;
    }
    for (auto Section : *SplittedCommand)
    {
        Index++;
        if (IsNextCommandBufferSize)
        {
            if (!ConvertStringToUInt32(Section, &RequestBuffer))
            {
                ShowMessages("err, buffer size is invalid\n");
                *ReasonForErrorInParsing = DEBUGGER_EVENT_PARSING_ERROR_CAUSE_FORMAT_ERROR;
                goto ReturnWithError;
            }
            else
            {
                if (TempActionBreak != NULL)
                {
                    TempActionBreak->PreAllocatedBuffer = RequestBuffer;
                }
                if (TempActionScript != NULL)
                {
                    TempActionScript->PreAllocatedBuffer = RequestBuffer;
                }
                if (TempActionCustomCode != NULL)
                {
                    TempActionCustomCode->PreAllocatedBuffer = RequestBuffer;
                }
            }
            IsNextCommandBufferSize = FALSE;
            IndexesToRemove.push_back(Index);
            continue;
        }
        if (IsNextCommandImmediateMessaging)
        {
            if (!Section.compare("yes"))
            {
                ImmediateMessagePassing = TRUE;
            }
            else if (!Section.compare("no"))
            {
                ImmediateMessagePassing = FALSE;
            }
            else
            {
                ShowMessages("err, immediate messaging token is invalid\n");
                *ReasonForErrorInParsing = DEBUGGER_EVENT_PARSING_ERROR_CAUSE_FORMAT_ERROR;
                goto ReturnWithError;
            }
            IsNextCommandImmediateMessaging = FALSE;
            IndexesToRemove.push_back(Index);
            continue;
        }
        if (IsNextCommandPid)
        {
            if (!Section.compare("all"))
            {
                TempEvent->ProcessId = DEBUGGER_EVENT_APPLY_TO_ALL_PROCESSES;
            }
            else if (!ConvertStringToUInt32(Section, &ProcessId))
            {
                ShowMessages("err, pid is invalid\n");
                *ReasonForErrorInParsing = DEBUGGER_EVENT_PARSING_ERROR_CAUSE_FORMAT_ERROR;
                goto ReturnWithError;
            }
            else
            {
                TempEvent->ProcessId = ProcessId;
            }
            IsNextCommandPid = FALSE;
            IndexesToRemove.push_back(Index);
            continue;
        }
        if (IsNextCommandCoreId)
        {
            if (!ConvertStringToUInt32(Section, &CoreId))
            {
                ShowMessages("err, core id is invalid\n");
                *ReasonForErrorInParsing = DEBUGGER_EVENT_PARSING_ERROR_CAUSE_FORMAT_ERROR;
                goto ReturnWithError;
            }
            else
            {
                TempEvent->CoreId = CoreId;
            }
            IsNextCommandCoreId = FALSE;
            IndexesToRemove.push_back(Index);
            continue;
        }
        if (!Section.compare("pid"))
        {
            IsNextCommandPid = TRUE;
            IndexesToRemove.push_back(Index);
            continue;
        }
        if (!Section.compare("core"))
        {
            IsNextCommandCoreId = TRUE;
            IndexesToRemove.push_back(Index);
            continue;
        }
        if (!Section.compare("imm"))
        {
            IsNextCommandImmediateMessaging = TRUE;
            IndexesToRemove.push_back(Index);
            continue;
        }
        if (!Section.compare("buffer"))
        {
            IsNextCommandBufferSize = TRUE;
            IndexesToRemove.push_back(Index);
            continue;
        }
    }
    if (IsNextCommandCoreId)
    {
        ShowMessages("err, please specify a value for 'core'\n");
        *ReasonForErrorInParsing = DEBUGGER_EVENT_PARSING_ERROR_CAUSE_FORMAT_ERROR;
        goto ReturnWithError;
    }
    if (IsNextCommandPid)
    {
        ShowMessages("err, please specify a value for 'pid'\n");
        *ReasonForErrorInParsing = DEBUGGER_EVENT_PARSING_ERROR_CAUSE_FORMAT_ERROR;
        goto ReturnWithError;
    }
    if (IsNextCommandBufferSize)
    {
        ShowMessages("err, please specify a value for 'buffer'\n");
        *ReasonForErrorInParsing = DEBUGGER_EVENT_PARSING_ERROR_CAUSE_FORMAT_ERROR;
        goto ReturnWithError;
    }
    if (!g_IsSerialConnectedToRemoteDebuggee && TempActionBreak != NULL)
    {
        ShowMessages(
            "err, it's not possible to break to the debugger in VMI Mode. "
            "You should operate in Debugger Mode to break and get the "
            "full control of the system. Still, you can use 'script' and run "
            "'custom code' in your local debugging (VMI Mode)\n");
        *ReasonForErrorInParsing = DEBUGGER_EVENT_PARSING_ERROR_CAUSE_ATTEMPT_TO_BREAK_ON_VMI_MODE;
        goto ReturnWithError;
    }
    if (!ImmediateMessagePassing && HasOutputPath)
    {
        ShowMessages("err, non-immediate message passing is not supported in "
                     "'output-forwarding mode'\n");
        *ReasonForErrorInParsing = DEBUGGER_EVENT_PARSING_ERROR_CAUSE_IMMEDIATE_MESSAGING_IN_EVENT_FORWARDING_MODE;
        goto ReturnWithError;
    }
    if (TempActionBreak != NULL)
    {
        TempActionBreak->ImmediateMessagePassing = ImmediateMessagePassing;
    }
    if (TempActionScript != NULL)
    {
        TempActionScript->ImmediateMessagePassing = ImmediateMessagePassing;
    }
    if (TempActionCustomCode != NULL)
    {
        TempActionCustomCode->ImmediateMessagePassing = ImmediateMessagePassing;
    }
    IndexOfValidSourceTags = 0;
    for (auto item : ListOfValidSourceTags)
    {
        TempEvent->OutputSourceTags[IndexOfValidSourceTags] = item;
        IndexOfValidSourceTags++;
    }
    if (HasOutputPath)
    {
        TempEvent->HasCustomOutput = TRUE;
    }
    *EventDetailsToFill = TempEvent;
    *EventBufferLength  = LengthOfEventBuffer;
    if (TempActionBreak != NULL)
    {
        *ActionDetailsToFillBreakToDebugger = TempActionBreak;
        *ActionBufferLengthBreakToDebugger  = LengthOfBreakActionBuffer;
    }
    if (TempActionScript != NULL)
    {
        *ActionDetailsToFillScript = TempActionScript;
        *ActionBufferLengthScript  = LengthOfScriptActionBuffer;
    }
    if (TempActionCustomCode != NULL)
    {
        *ActionDetailsToFillCustomCode = TempActionCustomCode;
        *ActionBufferLengthCustomCode  = LengthOfCustomCodeActionBuffer;
    }
    for (auto IndexToRemove : IndexesToRemove)
    {
        NewIndexToRemove++;
        SplittedCommand->erase(SplittedCommand->begin() +
                               (IndexToRemove - NewIndexToRemove));
        SplittedCommandCaseSensitive->erase(SplittedCommandCaseSensitive->begin() +
                                            (IndexToRemove - NewIndexToRemove));
    }
    if (!g_EventTraceInitialized)
    {
        InitializeListHead(&g_EventTrace);
        g_EventTraceInitialized = TRUE;
    }
    *ReasonForErrorInParsing = DEBUGGER_EVENT_PARSING_ERROR_CAUSE_SUCCESSFUL_NO_ERROR;
    return TRUE;
ReturnWithError:
    if (BufferOfCommandString)
    {
        free(BufferOfCommandString);
    }
    if (TempEvent)
    {
        free(TempEvent);
    }
    if (TempActionBreak != NULL)
    {
        free(TempActionBreak);
    }
    if (TempActionScript != NULL)
    {
        free(TempActionScript);
    }
    if (TempActionCustomCode != NULL)
    {
        free(TempActionCustomCode);
    }
    return FALSE;
}*/
return true
}



