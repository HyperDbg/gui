package app

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\app\hprdbgctrl.cpp.back

type (
	Hprdbgctrl interface {
		HyperdbgSetTextMessageCallback() (ok bool) //col:4
		ShowMessages() (ok bool)                   //col:195
		HyperDbgStopVmmDriver() (ok bool)          //col:210
	}
	hprdbgctrl struct{}
)

func NewHprdbgctrl() Hprdbgctrl { return &hprdbgctrl{} }

func (h *hprdbgctrl) HyperdbgSetTextMessageCallback() (ok bool) { //col:4
	/*
	   HyperdbgSetTextMessageCallback(Callback handler)

	   	{
	   	    g_MessageHandler = handler;
	   	}
	*/
	return true
}

func (h *hprdbgctrl) ShowMessages() (ok bool) { //col:195
	/*
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
	   	        else if (g_IsSerialConnectedToRemoteDebugger)
	   	        {
	   	            KdSendUsermodePrints(TempMessage, sprintfresult);
	   	        if (g_LogOpened)
	   	        {
	   	            LogopenSaveToFile(TempMessage);
	   	        if (g_MessageHandler != NULL)
	   	        {
	   	            g_MessageHandler(TempMessage);

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
	   	        else if (ErrorNum == ERROR_GEN_FAILURE)
	   	        {
	   	            ShowMessages("err, a device attached to the system is not functioning\n"
	   	                         "vmx feature might be disabled from BIOS or VBS/HVCI is active\n");
	   	            ShowMessages("err, CreateFile failed with (%x)\n", ErrorNum);
	   	    char * OutputBuffer = (char *)malloc(UsermodeBufferSize);
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
	   	                    if (g_BreakPrintingOutput)
	   	                    {
	   	                        continue;
	   	                    }
	   	                    ShowMessages("%s", OutputBuffer + sizeof(UINT32));
	   	                    if (g_BreakPrintingOutput)
	   	                    {
	   	                        continue;
	   	                    }
	   	                    ShowMessages("%s", OutputBuffer + sizeof(UINT32));
	   	                    if (g_BreakPrintingOutput)
	   	                    {
	   	                        continue;
	   	                    }
	   	                    ShowMessages("%s", OutputBuffer + sizeof(UINT32));
	   	                    KdCloseConnection();
	   	                    KdHandleUserInputInDebuggee((DEBUGGEE_USER_INPUT_PACKET *)(OutputBuffer + sizeof(UINT32)));
	   	                    KdRegisterEventInDebuggee(
	   	                        (PDEBUGGER_GENERAL_EVENT_DETAIL)(OutputBuffer + sizeof(UINT32)),
	   	                        ReturnedLength);
	   	                    KdAddActionToEventInDebuggee(
	   	                        (PDEBUGGER_GENERAL_ACTION)(OutputBuffer + sizeof(UINT32)),
	   	                        ReturnedLength);
	   	                    KdSendModifyEventInDebuggee(
	   	                        (PDEBUGGER_MODIFY_EVENTS)(OutputBuffer + sizeof(UINT32)));
	   	                    SetEvent(g_IsDriverLoadedSuccessfully);
	   	                    KdReloadSymbolsInDebuggee(TRUE,
	   	                                              ((PDEBUGGEE_SYMBOL_REQUEST_PACKET)(OutputBuffer + sizeof(UINT32)))->ProcessId);
	   	                    UdHandleUserDebuggerPausing(
	   	                        (PDEBUGGEE_UD_PAUSED_PACKET)(OutputBuffer + sizeof(UINT32)));
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
	   	                    if (!OutputSourceFound)
	   	                    {
	   	                        ShowMessages("%s", OutputBuffer + sizeof(UINT32));
	   	                free(OutputBuffer);
	   	                if (!CloseHandle(Handle))
	   	                {
	   	                    ShowMessages("err, closing handle 0x%x\n", GetLastError());
	   	    catch (const std::exception &)
	   	    {
	   	        ShowMessages(" Exception !\n");
	   	    free(OutputBuffer);
	   	    if (!CloseHandle(Handle))
	   	    {
	   	        ShowMessages("err, closing handle 0x%x\n", GetLastError());

	   ThreadFunc(void * data)

	   	{
	   	    ReadIrpBasedBuffer();

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
	*/
	return true
}

func (h *hprdbgctrl) HyperDbgStopVmmDriver() (ok bool) { //col:210
	/*
	   HyperDbgStopVmmDriver()

	   	{
	   	    return HyperdbgStopDriver(VMM_DRIVER_NAME);

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
	*/
	return true
}

