package meta-commands
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\debugger\commands\meta-commands\debug.cpp.back

type (
Debug interface{
CommandDebugHelp()(ok bool)//col:20
CommandDebugCheckBaudrate()(ok bool)//col:33
CommandDebugCheckComPort()(ok bool)//col:57
CommandDebug()(ok bool)//col:174
}
debug struct{}
)

func NewDebug()Debug{ return & debug{} }

func (d *debug)CommandDebugHelp()(ok bool){//col:20
/*CommandDebugHelp()
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
}*/
return true
}

func (d *debug)CommandDebugCheckBaudrate()(ok bool){//col:33
/*CommandDebugCheckBaudrate(DWORD Baudrate)
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
}*/
return true
}

func (d *debug)CommandDebugCheckComPort()(ok bool){//col:57
/*CommandDebugCheckComPort(const string & ComPort, UINT32 * Port)
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
}*/
return true
}

func (d *debug)CommandDebug()(ok bool){//col:174
/*CommandDebug(vector<string> SplittedCommand, string Command)
{
    UINT32 Baudrate;
    UINT32 Port;
    if (SplittedCommand.size() == 2 && !SplittedCommand.at(1).compare("close"))
    {
        if (!g_IsSerialConnectedToRemoteDebuggee &&
            !g_IsSerialConnectedToRemoteDebugger)
        {
            ShowMessages(
                "err, debugger is not attached to any instance of debuggee\n");
        }
        else
        {
            KdCloseConnection();
        }
        return;
    }
    else if (SplittedCommand.size() <= 3)
    {
        ShowMessages("incorrect use of '.debug'\n\n");
        CommandDebugHelp();
        return;
    }
    if (!SplittedCommand.at(1).compare("remote"))
    {
        if (!SplittedCommand.at(2).compare("serial"))
        {
            if (SplittedCommand.size() != 5)
            {
                ShowMessages("incorrect use of '.debug'\n\n");
                CommandDebugHelp();
                return;
            }
            if (!IsNumber(SplittedCommand.at(3)))
            {
                ShowMessages("unknown parameter '%s'\n\n",
                             SplittedCommand.at(3).c_str());
                CommandDebugHelp();
                return;
            }
            Baudrate = stoi(SplittedCommand.at(3));
            if (!CommandDebugCheckBaudrate(Baudrate))
            {
                ShowMessages("err, baud rate is invalid\n\n");
                CommandDebugHelp();
                return;
            }
            if (!CommandDebugCheckComPort(SplittedCommand.at(4), &Port))
            {
                ShowMessages("err, COM port is invalid\n\n");
                CommandDebugHelp();
                return;
            }
            KdPrepareAndConnectDebugPort(SplittedCommand.at(4).c_str(), Baudrate, Port, FALSE, FALSE);
        }
        else if (!SplittedCommand.at(2).compare("namedpipe"))
        {
            string Delimiter = "namedpipe";
            string Token     = Command.substr(
                Command.find(Delimiter) + Delimiter.size() + 1,
                Command.size());
            KdPrepareAndConnectDebugPort(Token.c_str(), NULL, NULL, FALSE, TRUE);
        }
        else
        {
            ShowMessages("unknown parameter '%s'\n\n", SplittedCommand.at(2).c_str());
            CommandDebugHelp();
            return;
        }
    }
    else if (!SplittedCommand.at(1).compare("prepare"))
    {
        if (SplittedCommand.size() != 5)
        {
            ShowMessages("incorrect use of '.debug'\n\n");
            CommandDebugHelp();
            return;
        }
        if (!SplittedCommand.at(2).compare("serial"))
        {
            if (!IsNumber(SplittedCommand.at(3)))
            {
                ShowMessages("unknown parameter '%s'\n\n",
                             SplittedCommand.at(3).c_str());
                CommandDebugHelp();
                return;
            }
            Baudrate = stoi(SplittedCommand.at(3));
            if (!CommandDebugCheckBaudrate(Baudrate))
            {
                ShowMessages("err, baud rate is invalid\n\n");
                CommandDebugHelp();
                return;
            }
            if (!CommandDebugCheckComPort(SplittedCommand.at(4), &Port))
            {
                ShowMessages("err, COM port is invalid\n\n");
                CommandDebugHelp();
                return;
            }
            KdPrepareAndConnectDebugPort(SplittedCommand.at(4).c_str(), Baudrate, Port, TRUE, FALSE);
        }
        else
        {
            ShowMessages("invalid parameter '%s'\n\n", SplittedCommand.at(2));
            CommandDebugHelp();
            return;
        }
    }
    else
    {
        ShowMessages("invalid parameter '%s'\n\n", SplittedCommand.at(1));
        CommandDebugHelp();
        return;
    }
}*/
return true
}



