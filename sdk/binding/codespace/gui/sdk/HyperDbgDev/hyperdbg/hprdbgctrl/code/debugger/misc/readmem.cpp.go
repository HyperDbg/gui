package misc

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\debugger\misc\readmem.cpp.back

type (
	Readmem interface {
		HyperDbgReadMemoryAndDisassemble() (ok bool) //col:85
		ShowMemoryCommandDB() (ok bool)              //col:122
		ShowMemoryCommandDC() (ok bool)              //col:160
		ShowMemoryCommandDD() (ok bool)              //col:185
		ShowMemoryCommandDQ() (ok bool)              //col:211
	}
	readmem struct{}
)

func NewReadmem() Readmem { return &readmem{} }

func (r *readmem) HyperDbgReadMemoryAndDisassemble() (ok bool) { //col:85
	/*
	   HyperDbgReadMemoryAndDisassemble(DEBUGGER_SHOW_MEMORY_STYLE   Style,

	   	UINT64                       Address,
	   	DEBUGGER_READ_MEMORY_TYPE    MemoryType,
	   	DEBUGGER_READ_READING_TYPE   ReadingType,
	   	UINT32                       Pid,
	   	UINT32                       Size,
	   	PDEBUGGER_DT_COMMAND_OPTIONS DtDetails)

	   	{
	   	    BOOL                 Status;
	   	    ULONG                ReturnedLength;
	   	    DEBUGGER_READ_MEMORY ReadMem = {0};
	   	    UINT32               OperationCode;
	   	    CHAR                 Character;
	   	    ReadMem.Address     = Address;
	   	    ReadMem.Pid         = Pid;
	   	    ReadMem.Size        = Size;
	   	    ReadMem.MemoryType  = MemoryType;
	   	    ReadMem.ReadingType = ReadingType;
	   	    ReadMem.Style       = Style;
	   	    ReadMem.DtDetails   = DtDetails;
	   	    if (g_IsSerialConnectedToRemoteDebuggee)
	   	    {
	   	        KdSendReadMemoryPacketToDebuggee(&ReadMem);
	   	        return;
	   	    }
	   	    AssertShowMessageReturnStmt(g_DeviceHandle, ASSERT_MESSAGE_DRIVER_NOT_LOADED, AssertReturn);
	   	    unsigned char * OutputBuffer = (unsigned char *)malloc(Size * sizeof(unsigned char));
	   	    ZeroMemory(OutputBuffer, Size * sizeof(unsigned char));
	   	    Status = DeviceIoControl(g_DeviceHandle,
	   	                             IOCTL_DEBUGGER_READ_MEMORY,
	   	                             &ReadMem,
	   	                             SIZEOF_DEBUGGER_READ_MEMORY,
	   	                             OutputBuffer,
	   	                             Size,
	   	                             &ReturnedLength,
	   	                             NULL
	   	    );
	   	    if (!Status)
	   	    {
	   	        ShowMessages("ioctl failed with code 0x%x\n", GetLastError());
	   	        free(OutputBuffer);
	   	        return;
	   	    }
	   	    switch (Style)
	   	    {
	   	    case DEBUGGER_SHOW_COMMAND_DT:
	   	        if (Size == ReturnedLength)
	   	        {
	   	            ScriptEngineShowDataBasedOnSymbolTypesWrapper(DtDetails->TypeName,
	   	                                                          Address,
	   	                                                          FALSE,
	   	                                                          OutputBuffer,
	   	                                                          DtDetails->AdditionalParameters);
	   	        }
	   	        else if (ReturnedLength == 0)
	   	        {
	   	            ShowMessages("err, invalid address");
	   	        }
	   	        else
	   	        {
	   	            ShowMessages("err, invalid address or memory is smaller than the structure size");
	   	        }
	   	        break;
	   	    case DEBUGGER_SHOW_COMMAND_DB:
	   	        ShowMemoryCommandDB(OutputBuffer, Size, Address, MemoryType, ReturnedLength);
	   	        break;
	   	    case DEBUGGER_SHOW_COMMAND_DC:
	   	        ShowMemoryCommandDC(OutputBuffer, Size, Address, MemoryType, ReturnedLength);
	   	        break;
	   	    case DEBUGGER_SHOW_COMMAND_DD:
	   	        ShowMemoryCommandDD(OutputBuffer, Size, Address, MemoryType, ReturnedLength);
	   	        break;
	   	    case DEBUGGER_SHOW_COMMAND_DQ:
	   	        ShowMemoryCommandDQ(OutputBuffer, Size, Address, MemoryType, ReturnedLength);
	   	        break;
	   	    case DEBUGGER_SHOW_COMMAND_DISASSEMBLE64:
	   	        HyperDbgDisassembler64(OutputBuffer, Address, ReturnedLength, 0, FALSE, NULL);
	   	        break;
	   	    case DEBUGGER_SHOW_COMMAND_DISASSEMBLE32:
	   	        HyperDbgDisassembler32(OutputBuffer, Address, ReturnedLength, 0, FALSE, NULL);
	   	        break;
	   	    }
	   	    free(OutputBuffer);
	   	    ShowMessages("\n");
	   	}
	*/
	return true
}

func (r *readmem) ShowMemoryCommandDB() (ok bool) { //col:122
	/*
	   ShowMemoryCommandDB(unsigned char * OutputBuffer, UINT Size, UINT64 Address, DEBUGGER_READ_MEMORY_TYPE MemoryType, UINT64 Length)

	   	{
	   	    unsigned int Character;
	   	    for (int i = 0; i < Size; i += 16)
	   	    {
	   	        if (MemoryType == DEBUGGER_READ_PHYSICAL_ADDRESS)
	   	        {
	   	            ShowMessages("#\t");
	   	        }
	   	        ShowMessages("%s  ", SeparateTo64BitValue((UINT64)(Address + i)).c_str());
	   	        for (size_t j = 0; j < 16; j++)
	   	        {
	   	            if (i + j >= Length)
	   	            {
	   	                ShowMessages("?? ");
	   	            }
	   	            else
	   	            {
	   	                ShowMessages("%02X ", OutputBuffer[i + j]);
	   	            }
	   	        }
	   	        ShowMessages(" ");
	   	        for (size_t j = 0; j < 16; j++)
	   	        {
	   	            Character = (OutputBuffer[i + j]);
	   	            if (isprint(Character))
	   	            {
	   	                ShowMessages("%c", Character);
	   	            }
	   	            else
	   	            {
	   	                ShowMessages(".");
	   	            }
	   	        }
	   	        ShowMessages("\n");
	   	    }
	   	}
	*/
	return true
}

func (r *readmem) ShowMemoryCommandDC() (ok bool) { //col:160
	/*
	   ShowMemoryCommandDC(unsigned char * OutputBuffer, UINT Size, UINT64 Address, DEBUGGER_READ_MEMORY_TYPE MemoryType, UINT64 Length)

	   	{
	   	    unsigned int Character;
	   	    for (int i = 0; i < Size; i += 16)
	   	    {
	   	        if (MemoryType == DEBUGGER_READ_PHYSICAL_ADDRESS)
	   	        {
	   	            ShowMessages("#\t");
	   	        }
	   	        ShowMessages("%s  ", SeparateTo64BitValue((UINT64)(Address + i)).c_str());
	   	        for (size_t j = 0; j < 16; j += 4)
	   	        {
	   	            if (i + j >= Length)
	   	            {
	   	                ShowMessages("???????? ");
	   	            }
	   	            else
	   	            {
	   	                UINT32 OutputBufferVar = *((UINT32 *)&OutputBuffer[i + j]);
	   	                ShowMessages("%08X ", OutputBufferVar);
	   	            }
	   	        }
	   	        ShowMessages(" ");
	   	        for (size_t j = 0; j < 16; j++)
	   	        {
	   	            Character = (OutputBuffer[i + j]);
	   	            if (isprint(Character))
	   	            {
	   	                ShowMessages("%c", Character);
	   	            }
	   	            else
	   	            {
	   	                ShowMessages(".");
	   	            }
	   	        }
	   	        ShowMessages("\n");
	   	    }
	   	}
	*/
	return true
}

func (r *readmem) ShowMemoryCommandDD() (ok bool) { //col:185
	/*
	   ShowMemoryCommandDD(unsigned char * OutputBuffer, UINT Size, UINT64 Address, DEBUGGER_READ_MEMORY_TYPE MemoryType, UINT64 Length)

	   	{
	   	    CHAR Character;
	   	    for (int i = 0; i < Size; i += 16)
	   	    {
	   	        if (MemoryType == DEBUGGER_READ_PHYSICAL_ADDRESS)
	   	        {
	   	            ShowMessages("#\t");
	   	        }
	   	        ShowMessages("%s  ", SeparateTo64BitValue((UINT64)(Address + i)).c_str());
	   	        for (size_t j = 0; j < 16; j += 4)
	   	        {
	   	            if (i + j >= Length)
	   	            {
	   	                ShowMessages("???????? ");
	   	            }
	   	            else
	   	            {
	   	                UINT32 OutputBufferVar = *((UINT32 *)&OutputBuffer[i + j]);
	   	                ShowMessages("%08X ", OutputBufferVar);
	   	            }
	   	        }
	   	        ShowMessages("\n");
	   	    }
	   	}
	*/
	return true
}

func (r *readmem) ShowMemoryCommandDQ() (ok bool) { //col:211
	/*
	   ShowMemoryCommandDQ(unsigned char * OutputBuffer, UINT Size, UINT64 Address, DEBUGGER_READ_MEMORY_TYPE MemoryType, UINT64 Length)

	   	{
	   	    for (int i = 0; i < Size; i += 16)
	   	    {
	   	        if (MemoryType == DEBUGGER_READ_PHYSICAL_ADDRESS)
	   	        {
	   	            ShowMessages("#\t");
	   	        }
	   	        ShowMessages("%s  ", SeparateTo64BitValue((UINT64)(Address + i)).c_str());
	   	        for (size_t j = 0; j < 16; j += 8)
	   	        {
	   	            if (i + j >= Length)
	   	            {
	   	                ShowMessages("???????? ");
	   	            }
	   	            else
	   	            {
	   	                UINT32 OutputBufferVar = *((UINT32 *)&OutputBuffer[i + j + 4]);
	   	                ShowMessages("%08X`", OutputBufferVar);
	   	                OutputBufferVar = *((UINT32 *)&OutputBuffer[i + j]);
	   	                ShowMessages("%08X ", OutputBufferVar);
	   	            }
	   	        }
	   	        ShowMessages("\n");
	   	    }
	   	}
	*/
	return true
}

