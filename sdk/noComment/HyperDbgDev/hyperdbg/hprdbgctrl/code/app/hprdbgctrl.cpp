#include "pch.h"
using namespace std;
extern HANDLE     g_DeviceHandle;
extern HANDLE     g_IsDriverLoadedSuccessfully;
extern BOOLEAN    g_IsVmxOffProcessStart;
extern Callback   g_MessageHandler;
extern TCHAR      g_DriverLocation[MAX_PATH];
extern LIST_ENTRY g_EventTrace;
extern BOOLEAN    g_LogOpened;
extern BOOLEAN    g_BreakPrintingOutput;
extern BOOLEAN    g_IsConnectedToRemoteDebugger;
extern BOOLEAN    g_OutputSourcesInitialized;
extern BOOLEAN    g_IsSerialConnectedToRemoteDebugger;
extern BOOLEAN    g_IsDebuggerModulesLoaded;
extern LIST_ENTRY g_OutputSources;
VOID
HyperdbgSetTextMessageCallback(Callback handler)
{
    g_MessageHandler = handler;
}

VOID
ShowMessages(const char * Fmt, ...)
{
    va_list ArgList;
    va_list Args;
    char    TempMessage[COMMUNICATION_BUFFER_SIZE + TCP_END_OF_BUFFER_CHARS_COUNT] = {0};
    if (g_MessageHandler == NULL && !g_IsConnectedToRemoteDebugger &&
        !g_IsSerialConnectedToRemoteDebugger)
    {
        va_start(Args, Fmt);
        vprintf(Fmt, Args);
        va_end(Args);
        if (!g_LogOpened)
        {
            return;
        }
    }
    va_start(ArgList, Fmt);
    int sprintfresult = vsprintf_s(TempMessage, Fmt, ArgList);
    va_end(ArgList);
    if (sprintfresult != -1)
    {
        if (g_IsConnectedToRemoteDebugger)
        {
            RemoteConnectionSendResultsToHost(TempMessage, sprintfresult);
        }
        else if (g_IsSerialConnectedToRemoteDebugger)
        {
            KdSendUsermodePrints(TempMessage, sprintfresult);
        }
        if (g_LogOpened)
        {
            LogopenSaveToFile(TempMessage);
        }
        if (g_MessageHandler != NULL)
        {
            g_MessageHandler(TempMessage);
        }
    }
}

#if !UseDbgPrintInsteadOfUsermodeMessageTracking
void
ReadIrpBasedBuffer()
{
    BOOL                   Status;
    ULONG                  ReturnedLength;
    REGISTER_NOTIFY_BUFFER RegisterEvent;
    UINT32                 OperationCode;
    DWORD                  ErrorNum;
    HANDLE                 Handle;
    BOOLEAN                OutputSourceFound;
    PLIST_ENTRY            TempList;
    RegisterEvent.hEvent = NULL;
    RegisterEvent.Type   = IRP_BASED;
    Handle = CreateFileA(
        "\\\\.\\HyperdbgHypervisorDevice",
        GENERIC_READ | GENERIC_WRITE,
        FILE_SHARE_READ | FILE_SHARE_WRITE,
        NULL, 
        OPEN_EXISTING,
        FILE_ATTRIBUTE_NORMAL | FILE_FLAG_OVERLAPPED,
        NULL); 
    if (Handle == INVALID_HANDLE_VALUE)
    {
        ErrorNum = GetLastError();
        if (ErrorNum == ERROR_ACCESS_DENIED)
        {
            ShowMessages("err, access denied\nare you sure you have administrator "
                         "rights?\n");
        }
        else if (ErrorNum == ERROR_GEN_FAILURE)
        {
            ShowMessages("err, a device attached to the system is not functioning\n"
                         "vmx feature might be disabled from BIOS or VBS/HVCI is active\n");
        }
        else
        {
            ShowMessages("err, CreateFile failed with (%x)\n", ErrorNum);
        }
        g_DeviceHandle = NULL;
        Handle         = NULL;
        return;
    }
    char * OutputBuffer = (char *)malloc(UsermodeBufferSize);
    try
    {
        while (TRUE)
        {
            if (!g_IsVmxOffProcessStart)
            {
                ZeroMemory(OutputBuffer, UsermodeBufferSize);
                Sleep(DefaultSpeedOfReadingKernelMessages); 
                Status = DeviceIoControl(
                    Handle,               
                    IOCTL_REGISTER_EVENT, 
                    &RegisterEvent,       
                    SIZEOF_REGISTER_EVENT *
                        2,              
                    OutputBuffer,       
                    UsermodeBufferSize, 
                    &ReturnedLength,    
                    NULL                
                );
                if (!Status)
                {
                    continue;
                }
                OperationCode = 0;
                memcpy(&OperationCode, OutputBuffer, sizeof(UINT32));
                switch (OperationCode)
                {
                case OPERATION_LOG_NON_IMMEDIATE_MESSAGE:
                    if (g_BreakPrintingOutput)
                    {
                        continue;
                    }
                    ShowMessages("%s", OutputBuffer + sizeof(UINT32));
                    break;
                case OPERATION_LOG_INFO_MESSAGE:
                    if (g_BreakPrintingOutput)
                    {
                        continue;
                    }
                    ShowMessages("%s", OutputBuffer + sizeof(UINT32));
                    break;
                case OPERATION_LOG_ERROR_MESSAGE:
                    if (g_BreakPrintingOutput)
                    {
                        continue;
                    }
                    ShowMessages("%s", OutputBuffer + sizeof(UINT32));
                    break;
                case OPERATION_LOG_WARNING_MESSAGE:
                    if (g_BreakPrintingOutput)
                    {
                        continue;
                    }
                    ShowMessages("%s", OutputBuffer + sizeof(UINT32));
                    break;
                case OPERATION_COMMAND_FROM_DEBUGGER_CLOSE_AND_UNLOAD_VMM:
                    KdCloseConnection();
                    break;
                case OPERATION_DEBUGGEE_USER_INPUT:
                    KdHandleUserInputInDebuggee((DEBUGGEE_USER_INPUT_PACKET *)(OutputBuffer + sizeof(UINT32)));
                    break;
                case OPERATION_DEBUGGEE_REGISTER_EVENT:
                    KdRegisterEventInDebuggee(
                        (PDEBUGGER_GENERAL_EVENT_DETAIL)(OutputBuffer + sizeof(UINT32)),
                        ReturnedLength);
                    break;
                case OPERATION_DEBUGGEE_ADD_ACTION_TO_EVENT:
                    KdAddActionToEventInDebuggee(
                        (PDEBUGGER_GENERAL_ACTION)(OutputBuffer + sizeof(UINT32)),
                        ReturnedLength);
                    break;
                case OPERATION_DEBUGGEE_CLEAR_EVENTS:
                    KdSendModifyEventInDebuggee(
                        (PDEBUGGER_MODIFY_EVENTS)(OutputBuffer + sizeof(UINT32)));
                    break;
                case OPERATION_HYPERVISOR_DRIVER_IS_SUCCESSFULLY_LOADED:
                    SetEvent(g_IsDriverLoadedSuccessfully);
                    break;
                case OPERATION_HYPERVISOR_DRIVER_END_OF_IRPS:
                    break;
                case OPERATION_COMMAND_FROM_DEBUGGER_RELOAD_SYMBOL:
                    KdReloadSymbolsInDebuggee(TRUE,
                                              ((PDEBUGGEE_SYMBOL_REQUEST_PACKET)(OutputBuffer + sizeof(UINT32)))->ProcessId);
                    break;
                case OPERATION_NOTIFICATION_FROM_USER_DEBUGGER_PAUSE:
                    UdHandleUserDebuggerPausing(
                        (PDEBUGGEE_UD_PAUSED_PACKET)(OutputBuffer + sizeof(UINT32)));
                    break;
                default:
                    if (g_BreakPrintingOutput)
                    {
                        continue;
                    }
                    OutputSourceFound = FALSE;
                    if (g_OutputSourcesInitialized)
                    {
                        TempList = &g_EventTrace;
                        while (&g_EventTrace != TempList->Blink)
                        {
                            TempList = TempList->Blink;
                            PDEBUGGER_GENERAL_EVENT_DETAIL EventDetail = CONTAINING_RECORD(
                                TempList,
                                DEBUGGER_GENERAL_EVENT_DETAIL,
                                CommandsEventList);
                            if (EventDetail->HasCustomOutput)
                            {
                                OutputSourceFound = TRUE;
                                if (!ForwardingPerformEventForwarding(
                                        EventDetail,
                                        OutputBuffer + sizeof(UINT32),
                                        ReturnedLength - sizeof(UINT32) + 1))
                                {
                                    ShowMessages("err, there was an error transferring the "
                                                 "message to the remote sources\n");
                                }
                                break;
                            }
                        }
                    }
                    if (!OutputSourceFound)
                    {
                        ShowMessages("%s", OutputBuffer + sizeof(UINT32));
                    }
                    break;
                }
            }
            else
            {
                free(OutputBuffer);
                if (!CloseHandle(Handle))
                {
                    ShowMessages("err, closing handle 0x%x\n", GetLastError());
                };
                return;
            }
        }
    }
    catch (const std::exception &)
    {
        ShowMessages(" Exception !\n");
    }
    free(OutputBuffer);
    if (!CloseHandle(Handle))
    {
        ShowMessages("err, closing handle 0x%x\n", GetLastError());
    };
}

DWORD WINAPI
ThreadFunc(void * data)
{
    ReadIrpBasedBuffer();
    return 0;
}

#endif
HPRDBGCTRL_API int
HyperDbgInstallVmmDriver()
{
    if (!SetupDriverName(g_DriverLocation, sizeof(g_DriverLocation)))
    {
        return 1;
    }
    if (!ManageDriver(VMM_DRIVER_NAME, g_DriverLocation, DRIVER_FUNC_INSTALL))
    {
        ShowMessages("unable to install driver\n");
        ManageDriver(VMM_DRIVER_NAME, g_DriverLocation, DRIVER_FUNC_REMOVE);
        return 1;
    }
    return 0;
}

int
HyperdbgStopDriver(LPCTSTR DriverName)
{
    if (g_DriverLocation[0] != (TCHAR)0 &&
        ManageDriver(DriverName, g_DriverLocation, DRIVER_FUNC_STOP))
    {
        return 0;
    }
    else
    {
        return 1;
    }
}

HPRDBGCTRL_API int
HyperDbgStopVmmDriver()
{
    return HyperdbgStopDriver(VMM_DRIVER_NAME);
}

int
HyperdbgUninstallDriver(LPCTSTR DriverName)
{
    if (g_DriverLocation[0] != (TCHAR)0 &&
        ManageDriver(DriverName, g_DriverLocation, DRIVER_FUNC_REMOVE))
    {
        return 0;
    }
    else
    {
        return 1;
    }
}

HPRDBGCTRL_API int
HyperDbgUninstallVmmDriver()
{
    return HyperdbgUninstallDriver(VMM_DRIVER_NAME);
}

HPRDBGCTRL_API int
HyperDbgLoadVmm()
{
    string CpuID;
    DWORD  ErrorNum;
    HANDLE hProcess;
    DWORD  ThreadId;
    if (g_DeviceHandle)
    {
        ShowMessages("handle of the driver found, if you use 'load' before, please "
                     "unload it using 'unload'\n");
        return 1;
    }
    CpuID = ReadVendorString();
    ShowMessages("current processor vendor is : %s\n", CpuID.c_str());
    if (CpuID == "GenuineIntel")
    {
        ShowMessages("virtualization technology is vt-x\n");
    }
    else
    {
        ShowMessages("this program is not designed to run in a non-VT-x "
                     "environemnt !\n");
        return 1;
    }
    if (VmxSupportDetection())
    {
        ShowMessages("vmx operation is supported by your processor\n");
    }
    else
    {
        ShowMessages("vmx operation is not supported by your processor\n");
        return 1;
    }
    g_IsVmxOffProcessStart = FALSE;
    g_DeviceHandle = CreateFileA(
        "\\\\.\\HyperdbgHypervisorDevice",
        GENERIC_READ | GENERIC_WRITE,
        FILE_SHARE_READ | FILE_SHARE_WRITE,
        NULL, 
        OPEN_EXISTING,
        FILE_ATTRIBUTE_NORMAL | FILE_FLAG_OVERLAPPED,
        NULL); 
    if (g_DeviceHandle == INVALID_HANDLE_VALUE)
    {
        ErrorNum = GetLastError();
        if (ErrorNum == ERROR_ACCESS_DENIED)
        {
            ShowMessages("err, access denied\nare you sure you have administrator "
                         "rights?\n");
        }
        else if (ErrorNum == ERROR_GEN_FAILURE)
        {
            ShowMessages("err, a device attached to the system is not functioning\n"
                         "vmx feature might be disabled from BIOS or VBS/HVCI is active\n");
        }
        else
        {
            ShowMessages("err, CreateFile failed (%x)\n", ErrorNum);
        }
        g_DeviceHandle = NULL;
        return 1;
    }
    InitializeListHead(&g_EventTrace);
#if !UseDbgPrintInsteadOfUsermodeMessageTracking
    HANDLE Thread = CreateThread(NULL, 0, ThreadFunc, NULL, 0, &ThreadId);
#endif
    return 0;
}

HPRDBGCTRL_API int
HyperDbgUnloadVmm()
{
    BOOL Status;
    AssertShowMessageReturnStmt(g_DeviceHandle, ASSERT_MESSAGE_DRIVER_NOT_LOADED, AssertReturnOne);
    ShowMessages("start terminating...\n");
    UdUninitializeUserDebugger();
    Status = DeviceIoControl(g_DeviceHandle,      
                             IOCTL_TERMINATE_VMX, 
                             NULL,                
                             0,                   
                             NULL,                
                             0,                   
                             NULL,                
                             NULL                 
    );
    if (!Status)
    {
        ShowMessages("ioctl failed with code 0x%x\n", GetLastError());
        return 1;
    }
    Status = DeviceIoControl(
        g_DeviceHandle,                                      
        IOCTL_RETURN_IRP_PENDING_PACKETS_AND_DISALLOW_IOCTL, 
        NULL,                                                
        0,                                                   
        NULL,                                                
        0,                                                   
        NULL,                                                
        NULL                                                 
    );
    if (!Status)
    {
        ShowMessages("ioctl failed with code 0x%x\n", GetLastError());
        return 1;
    }
    g_IsVmxOffProcessStart = TRUE;
    Sleep(1000); 
    if (!CloseHandle(g_DeviceHandle))
    {
        ShowMessages("err, closing handle 0x%x\n", GetLastError());
        return 1;
    };
    g_DeviceHandle = NULL;
    g_IsDebuggerModulesLoaded = FALSE;
    SymbolDeleteSymTable();
    ShowMessages("you're not on HyperDbg's hypervisor anymore!\n");
    return 0;
}

