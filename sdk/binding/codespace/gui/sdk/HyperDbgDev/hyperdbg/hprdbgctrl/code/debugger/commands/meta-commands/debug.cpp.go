package meta_commands

//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\debugger\commands\meta-commands\debug.cpp.back

type (
	Debug interface {
		CommandDebugHelp() (ok bool)         //col:32
		CommandDebugCheckComPort() (ok bool) //col:56
	}
	debug struct{}
)

func NewDebug() Debug { return &debug{} }

func (d *debug) CommandDebugHelp() (ok bool) { //col:32
	/*
	   CommandDebugHelp()

	   	{
	   	    ShowMessages(
	   	        ".debug : debugs a target machine or makes this machine a debuggee.\n\n");
	   	    ShowMessages(
	   	        "syntax : \t.debug [remote] [serial|namedpipe] [Baudrate (decimal)] [Address (string)]\n");
	   	    ShowMessages(
	   	        "syntax : \t.debug [prepare] [serial] [Baudrate (decimal)] [Address (string)]\n");
	   	    ShowMessages("syntax : \t.debug [close]\n");
	   	    ShowMessages("\n");
	   	    ShowMessages("\t\te.g : .debug remote serial 115200 com3\n");
	   	    ShowMessages("\t\te.g : .debug remote namedpipe \\\\.\\pipe\\HyperDbgPipe\n");
	   	    ShowMessages("\t\te.g : .debug prepare serial 115200 com1\n");
	   	    ShowMessages("\t\te.g : .debug prepare serial 115200 com2\n");
	   	    ShowMessages("\t\te.g : .debug close\n");
	   	    ShowMessages(
	   	        "\nvalid baud rates (decimal) : 110, 300, 600, 1200, 2400, 4800, 9600, "
	   	        "14400, 19200, 38400, 56000, 57600, 115200, 128000, 256000\n");
	   	    ShowMessages("valid COM ports : COM1, COM2, COM3, COM4 \n");

	   CommandDebugCheckBaudrate(DWORD Baudrate)

	   	{
	   	    if (Baudrate == CBR_110 || Baudrate == CBR_300 || Baudrate == CBR_600 ||
	   	        Baudrate == CBR_1200 || Baudrate == CBR_2400 || Baudrate == CBR_4800 ||
	   	        Baudrate == CBR_9600 || Baudrate == CBR_14400 || Baudrate == CBR_19200 ||
	   	        Baudrate == CBR_38400 || Baudrate == CBR_56000 || Baudrate == CBR_57600 ||
	   	        Baudrate == CBR_115200 || Baudrate == CBR_128000 ||
	   	        Baudrate == CBR_256000)
	   	    {
	   	        return TRUE;
	   	    }
	   	    return FALSE;
	   	}
	*/
	return true
}

func (d *debug) CommandDebugCheckComPort() (ok bool) { //col:56
	/*
	   CommandDebugCheckComPort(const string & ComPort, UINT32 * Port)

	   	{
	   	    if (!ComPort.compare("com1"))
	   	    {
	   	        *Port = COM1_PORT;
	   	        return TRUE;
	   	    }
	   	    else if (!ComPort.compare("com2"))
	   	    {
	   	        *Port = COM2_PORT;
	   	        return TRUE;
	   	    }
	   	    else if (!ComPort.compare("com3"))
	   	    {
	   	        *Port = COM3_PORT;
	   	        return TRUE;
	   	    }
	   	    else if (!ComPort.compare("com4"))
	   	    {
	   	        *Port = COM4_PORT;
	   	        return TRUE;
	   	    }
	   	    return FALSE;
	   	}
	*/
	return true
}

