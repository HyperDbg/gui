package debugging_commands

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\debugger\commands\debugging-commands\events.cpp.back

type (
	Events interface {
		CommandEventsHelp() (ok bool)                       //col:17
		CommandEvents() (ok bool)                           //col:72
		CommandEventQueryEventState() (ok bool)             //col:98
		CommandEventsShowEvents() (ok bool)                 //col:130
		CommandEventDisableEvent() (ok bool)                //col:155
		CommandEventEnableEvent() (ok bool)                 //col:180
		CommandEventClearEvent() (ok bool)                  //col:218
		CommandEventsClearAllEventsAndResetTags() (ok bool) //col:227
		CommandEventsHandleModifiedEvent() (ok bool)        //col:327
		CommandEventsModifyAndQueryEvents() (ok bool)       //col:377
	}
	events struct{}
)

func NewEvents() Events { return &events{} }

func (e *events) CommandEventsHelp() (ok bool) { //col:17
	/*
	   CommandEventsHelp()

	   	{
	   	    ShowMessages("events : shows active and disabled events\n");
	   	    ShowMessages("syntax : \tevents\n");
	   	    ShowMessages("syntax : \tevents [e|d|c all|EventNumber (hex)]\n");
	   	    ShowMessages("e : enable\n");
	   	    ShowMessages("d : disable\n");
	   	    ShowMessages("c : clear\n");
	   	    ShowMessages("note : If you specify 'all' then e, d, or c will be applied to "
	   	                 "all of the events.\n\n");
	   	    ShowMessages("\n");
	   	    ShowMessages("\te.g : events \n");
	   	    ShowMessages("\te.g : events e 12\n");
	   	    ShowMessages("\te.g : events d 10\n");
	   	    ShowMessages("\te.g : events c 10\n");
	   	    ShowMessages("\te.g : events c all\n");
	   	}
	*/
	return true
}

func (e *events) CommandEvents() (ok bool) { //col:72
	/*
	   CommandEvents(vector<string> SplittedCommand, string Command)

	   	{
	   	    DEBUGGER_MODIFY_EVENTS_TYPE RequestedAction;
	   	    UINT64                      RequestedTag;
	   	    if (SplittedCommand.size() != 1 && SplittedCommand.size() != 3)
	   	    {
	   	        ShowMessages("incorrect use of '%s'\n\n", SplittedCommand.at(0).c_str());
	   	        CommandEventsHelp();
	   	        return;
	   	    }
	   	    if (SplittedCommand.size() == 1)
	   	    {
	   	        if (!g_EventTraceInitialized)
	   	        {
	   	            ShowMessages("no active/disabled events \n");
	   	            return;
	   	        }
	   	        CommandEventsShowEvents();
	   	        return;
	   	    }
	   	    if (!SplittedCommand.at(1).compare("e"))
	   	    {
	   	        RequestedAction = DEBUGGER_MODIFY_EVENTS_ENABLE;
	   	    }
	   	    else if (!SplittedCommand.at(1).compare("d"))
	   	    {
	   	        RequestedAction = DEBUGGER_MODIFY_EVENTS_DISABLE;
	   	    }
	   	    else if (!SplittedCommand.at(1).compare("c"))
	   	    {
	   	        RequestedAction = DEBUGGER_MODIFY_EVENTS_CLEAR;
	   	    }
	   	    else
	   	    {
	   	        ShowMessages("incorrect use of '%s'\n\n", SplittedCommand.at(0).c_str());
	   	        CommandEventsHelp();
	   	        return;
	   	    }
	   	    if (!SplittedCommand.at(2).compare("all"))
	   	    {
	   	        RequestedTag = DEBUGGER_MODIFY_EVENTS_APPLY_TO_ALL_TAG;
	   	    }
	   	    else if (!ConvertStringToUInt64(SplittedCommand.at(2), &RequestedTag))
	   	    {
	   	        ShowMessages(
	   	            "please specify a correct hex value for tag id (event number)\n\n");
	   	        CommandEventsHelp();
	   	        return;
	   	    }
	   	    if (RequestedTag != DEBUGGER_MODIFY_EVENTS_APPLY_TO_ALL_TAG)
	   	    {
	   	        RequestedTag = RequestedTag + DebuggerEventTagStartSeed;
	   	    }
	   	    CommandEventsModifyAndQueryEvents(RequestedTag, RequestedAction);
	   	}
	*/
	return true
}

func (e *events) CommandEventQueryEventState() (ok bool) { //col:98
	/*
	   CommandEventQueryEventState(UINT64 Tag)

	   	{
	   	    BOOLEAN IsEnabled;
	   	    if (g_IsSerialConnectedToRemoteDebuggee)
	   	    {
	   	        if (KdSendEventQueryAndModifyPacketToDebuggee(
	   	                Tag,
	   	                DEBUGGER_MODIFY_EVENTS_QUERY_STATE,
	   	                &IsEnabled))
	   	        {
	   	            return IsEnabled;
	   	        }
	   	        else
	   	        {
	   	            ShowMessages("err, unable to get the state of the event\n");
	   	            return FALSE;
	   	        }
	   	    }
	   	    else
	   	    {
	   	        return CommandEventsModifyAndQueryEvents(
	   	            Tag,
	   	            DEBUGGER_MODIFY_EVENTS_QUERY_STATE);
	   	    }
	   	    return FALSE;
	   	}
	*/
	return true
}

func (e *events) CommandEventsShowEvents() (ok bool) { //col:130
	/*
	   CommandEventsShowEvents()

	   	{
	   	    PLIST_ENTRY TempList         = 0;
	   	    BOOLEAN     IsThereAnyEvents = FALSE;
	   	    TempList = &g_EventTrace;
	   	    while (&g_EventTrace != TempList->Blink)
	   	    {
	   	        TempList = TempList->Blink;
	   	        PDEBUGGER_GENERAL_EVENT_DETAIL CommandDetail = CONTAINING_RECORD(TempList, DEBUGGER_GENERAL_EVENT_DETAIL, CommandsEventList);
	   	        string                         CommandMessage((char *)CommandDetail->CommandStringBuffer);
	   	        ReplaceAll(CommandMessage, "\n", " ");
	   	        if (CommandMessage.length() > 70)
	   	        {
	   	            CommandMessage = CommandMessage.substr(0, 70);
	   	            CommandMessage += "...";
	   	        }
	   	        ShowMessages("%x\t(%s)\t    %s\n",
	   	                     CommandDetail->Tag - DebuggerEventTagStartSeed,
	   	                     CommandEventQueryEventState(CommandDetail->Tag)
	   	                         ? "enabled"
	   	                         : "disabled",
	   	                     CommandMessage.c_str());
	   	        if (!IsThereAnyEvents)
	   	        {
	   	            IsThereAnyEvents = TRUE;
	   	        }
	   	    }
	   	    if (!IsThereAnyEvents)
	   	    {
	   	        ShowMessages("no active/disabled events \n");
	   	    }
	   	}
	*/
	return true
}

func (e *events) CommandEventDisableEvent() (ok bool) { //col:155
	/*
	   CommandEventDisableEvent(UINT64 Tag)

	   	{
	   	    PLIST_ENTRY TempList = 0;
	   	    BOOLEAN     Result   = FALSE;
	   	    TempList = &g_EventTrace;
	   	    while (&g_EventTrace != TempList->Blink)
	   	    {
	   	        TempList = TempList->Blink;
	   	        PDEBUGGER_GENERAL_EVENT_DETAIL CommandDetail = CONTAINING_RECORD(TempList, DEBUGGER_GENERAL_EVENT_DETAIL, CommandsEventList);
	   	        if (CommandDetail->Tag == Tag ||
	   	            Tag == DEBUGGER_MODIFY_EVENTS_APPLY_TO_ALL_TAG)
	   	        {
	   	            CommandDetail->IsEnabled = FALSE;
	   	            if (!Result)
	   	            {
	   	                Result = TRUE;
	   	            }
	   	            if (Tag != DEBUGGER_MODIFY_EVENTS_APPLY_TO_ALL_TAG)
	   	            {
	   	                return TRUE;
	   	            }
	   	        }
	   	    }
	   	    return Result;
	   	}
	*/
	return true
}

func (e *events) CommandEventEnableEvent() (ok bool) { //col:180
	/*
	   CommandEventEnableEvent(UINT64 Tag)

	   	{
	   	    PLIST_ENTRY TempList = 0;
	   	    BOOLEAN     Result   = FALSE;
	   	    TempList = &g_EventTrace;
	   	    while (&g_EventTrace != TempList->Blink)
	   	    {
	   	        TempList = TempList->Blink;
	   	        PDEBUGGER_GENERAL_EVENT_DETAIL CommandDetail = CONTAINING_RECORD(TempList, DEBUGGER_GENERAL_EVENT_DETAIL, CommandsEventList);
	   	        if (CommandDetail->Tag == Tag ||
	   	            Tag == DEBUGGER_MODIFY_EVENTS_APPLY_TO_ALL_TAG)
	   	        {
	   	            CommandDetail->IsEnabled = TRUE;
	   	            if (!Result)
	   	            {
	   	                Result = TRUE;
	   	            }
	   	            if (Tag != DEBUGGER_MODIFY_EVENTS_APPLY_TO_ALL_TAG)
	   	            {
	   	                return TRUE;
	   	            }
	   	        }
	   	    }
	   	    return Result;
	   	}
	*/
	return true
}

func (e *events) CommandEventClearEvent() (ok bool) { //col:218
	/*
	   CommandEventClearEvent(UINT64 Tag)

	   	{
	   	    PLIST_ENTRY                    TempList         = 0;
	   	    BOOLEAN                        Result           = FALSE;
	   	    PDEBUGGER_GENERAL_EVENT_DETAIL TmpCommandDetail = NULL;
	   	    TempList = &g_EventTrace;
	   	    while (&g_EventTrace != TempList->Flink)
	   	    {
	   	        TempList = TempList->Flink;
	   	        PDEBUGGER_GENERAL_EVENT_DETAIL CommandDetail = CONTAINING_RECORD(TempList, DEBUGGER_GENERAL_EVENT_DETAIL, CommandsEventList);
	   	        if (CommandDetail->Tag == Tag || Tag == DEBUGGER_MODIFY_EVENTS_APPLY_TO_ALL_TAG)
	   	        {
	   	            free(CommandDetail->CommandStringBuffer);
	   	            if (!Result)
	   	            {
	   	                Result = TRUE;
	   	                TmpCommandDetail = CommandDetail;
	   	            }
	   	            if (Tag != DEBUGGER_MODIFY_EVENTS_APPLY_TO_ALL_TAG)
	   	            {
	   	                RemoveEntryList(&CommandDetail->CommandsEventList);
	   	                free(CommandDetail);
	   	                return TRUE;
	   	            }
	   	        }
	   	    }
	   	    if (Result && Tag == DEBUGGER_MODIFY_EVENTS_APPLY_TO_ALL_TAG)
	   	    {
	   	        while ((UINT64)TmpCommandDetail != (UINT64)&g_EventTrace.Flink)
	   	        {
	   	            PVOID Temp       = (PVOID)TmpCommandDetail;
	   	            TmpCommandDetail = (PDEBUGGER_GENERAL_EVENT_DETAIL)TmpCommandDetail->CommandsEventList.Flink;
	   	            free(Temp);
	   	        }
	   	        InitializeListHead(&g_EventTrace);
	   	    }
	   	    return Result;
	   	}
	*/
	return true
}

func (e *events) CommandEventsClearAllEventsAndResetTags() (ok bool) { //col:227
	/*
	   CommandEventsClearAllEventsAndResetTags()

	   	{
	   	    if (g_EventTraceInitialized)
	   	    {
	   	        CommandEventClearEvent(DEBUGGER_MODIFY_EVENTS_APPLY_TO_ALL_TAG);
	   	        g_EventTraceInitialized = FALSE;
	   	        g_EventTag = DebuggerEventTagStartSeed;
	   	    }
	   	}
	*/
	return true
}

func (e *events) CommandEventsHandleModifiedEvent() (ok bool) { //col:327
	/*
	   CommandEventsHandleModifiedEvent(

	   	UINT64                  Tag,
	   	PDEBUGGER_MODIFY_EVENTS ModifyEventRequest)

	   	{
	   	    if (ModifyEventRequest->KernelStatus == DEBUGGER_OPERATION_WAS_SUCCESSFULL)
	   	    {
	   	        if (ModifyEventRequest->TypeOfAction == DEBUGGER_MODIFY_EVENTS_ENABLE)
	   	        {
	   	            if (!CommandEventEnableEvent(Tag))
	   	            {
	   	                ShowMessages("err, the event was successfully, "
	   	                             "(enabled|disabled|cleared) but "
	   	                             "can't apply it to the user-mode structures\n");
	   	            }
	   	        }
	   	        else if (ModifyEventRequest->TypeOfAction ==
	   	                 DEBUGGER_MODIFY_EVENTS_DISABLE)
	   	        {
	   	            if (!CommandEventDisableEvent(Tag))
	   	            {
	   	                ShowMessages("err, the event was successfully, "
	   	                             "(enabled|disabled|cleared) but "
	   	                             "can't apply it to the user-mode structures\n");
	   	            }
	   	            else
	   	            {
	   	                if (g_BreakPrintingOutput)
	   	                {
	   	                    if (!g_IsConnectedToRemoteDebuggee &&
	   	                        !g_IsSerialConnectedToRemoteDebuggee &&
	   	                        !g_IsSerialConnectedToRemoteDebugger)
	   	                    {
	   	                        if (!g_AutoFlush)
	   	                        {
	   	                            ShowMessages(
	   	                                "auto-flush mode is disabled, if there is still "
	   	                                "messages or buffers in the kernel, you continue to see "
	   	                                "the messages when you run 'g' until the kernel "
	   	                                "buffers are empty. you can run 'settings autoflush "
	   	                                "on' and after disabling and clearing events, "
	   	                                "kernel buffers will be flushed automatically\n");
	   	                        }
	   	                        else
	   	                        {
	   	                            CommandFlushRequestFlush();
	   	                        }
	   	                    }
	   	                }
	   	            }
	   	        }
	   	        else if (ModifyEventRequest->TypeOfAction ==
	   	                 DEBUGGER_MODIFY_EVENTS_CLEAR)
	   	        {
	   	            if (!CommandEventClearEvent(Tag))
	   	            {
	   	                ShowMessages("err, the event was successfully, "
	   	                             "(enabled|disabled|cleared) but "
	   	                             "can't apply it to the user-mode structures\n");
	   	            }
	   	            else
	   	            {
	   	                if (g_BreakPrintingOutput)
	   	                {
	   	                    if (!g_IsConnectedToRemoteDebuggee)
	   	                    {
	   	                        if (!g_AutoFlush)
	   	                        {
	   	                            ShowMessages(
	   	                                "auto-flush mode is disabled, if there is still "
	   	                                "messages or buffers in the kernel, you continue to see "
	   	                                "the messages when you run 'g' until the kernel "
	   	                                "buffers are empty. you can run 'settings autoflush "
	   	                                "on' and after disabling and clearing events, "
	   	                                "kernel buffers will be flushed automatically\n");
	   	                        }
	   	                        else
	   	                        {
	   	                            CommandFlushRequestFlush();
	   	                        }
	   	                    }
	   	                }
	   	            }
	   	        }
	   	        else if (ModifyEventRequest->TypeOfAction ==
	   	                 DEBUGGER_MODIFY_EVENTS_QUERY_STATE)
	   	        {
	   	        }
	   	        else
	   	        {
	   	            ShowMessages(
	   	                "err, the event was successfully, (enabled|disabled|cleared) but "
	   	                "can't apply it to the user-mode structures\n");
	   	        }
	   	    }
	   	    else
	   	    {
	   	        ShowErrorMessage(ModifyEventRequest->KernelStatus);
	   	        return;
	   	    }
	   	}
	*/
	return true
}

func (e *events) CommandEventsModifyAndQueryEvents() (ok bool) { //col:377
	/*
	   CommandEventsModifyAndQueryEvents(UINT64                      Tag,

	   	DEBUGGER_MODIFY_EVENTS_TYPE TypeOfAction)

	   	{
	   	    BOOLEAN                Status;
	   	    ULONG                  ReturnedLength;
	   	    DEBUGGER_MODIFY_EVENTS ModifyEventRequest = {0};
	   	    if (!IsTagExist(Tag))
	   	    {
	   	        if (Tag == DEBUGGER_MODIFY_EVENTS_APPLY_TO_ALL_TAG)
	   	        {
	   	            ShowMessages("there is no event\n");
	   	        }
	   	        else
	   	        {
	   	            ShowMessages("err, tag id is invalid\n");
	   	        }
	   	        return FALSE;
	   	    }
	   	    if (g_IsSerialConnectedToRemoteDebuggee)
	   	    {
	   	        KdSendEventQueryAndModifyPacketToDebuggee(Tag, TypeOfAction, NULL);
	   	    }
	   	    else
	   	    {
	   	        AssertShowMessageReturnStmt(g_DeviceHandle, ASSERT_MESSAGE_DRIVER_NOT_LOADED, AssertReturnFalse);
	   	        ModifyEventRequest.Tag          = Tag;
	   	        ModifyEventRequest.TypeOfAction = TypeOfAction;
	   	        Status =
	   	            DeviceIoControl(g_DeviceHandle,
	   	                            IOCTL_DEBUGGER_MODIFY_EVENTS,
	   	                            &ModifyEventRequest,
	   	                            SIZEOF_DEBUGGER_MODIFY_EVENTS,
	   	                            &ModifyEventRequest,
	   	                            SIZEOF_DEBUGGER_MODIFY_EVENTS,
	   	                            &ReturnedLength,
	   	                            NULL
	   	            );
	   	        if (!Status)
	   	        {
	   	            ShowMessages("ioctl failed with code 0x%x\n", GetLastError());
	   	            return FALSE;
	   	        }
	   	        CommandEventsHandleModifiedEvent(Tag, &ModifyEventRequest);
	   	        if (TypeOfAction == DEBUGGER_MODIFY_EVENTS_QUERY_STATE)
	   	        {
	   	            return ModifyEventRequest.IsEnabled;
	   	        }
	   	    }
	   	    return TRUE;
	   	}
	*/
	return true
}

