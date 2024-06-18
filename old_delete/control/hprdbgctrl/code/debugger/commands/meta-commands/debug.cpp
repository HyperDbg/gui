#include "pch.h"
extern HANDLE  g_SerialListeningThreadHandle;
extern HANDLE  g_SerialRemoteComPortHandle;
extern BOOLEAN g_IsSerialConnectedToRemoteDebuggee;
extern BOOLEAN g_IsSerialConnectedToRemoteDebugger;
extern BOOLEAN g_IsDebuggeeRunning;
VOID CommandDebugHelp(){
    ShowMessages(
        ".debug : debugs a target machine or makes this machine a debuggee.\n\n");
    ShowMessages(
        "syntax : \t.debug [remote] [serial|namedpipe] [Baudrate (decimal)] [Address (string)]\n");
    ShowMessages(
        "syntax : \t.debug [prepare] [serial] [Baudrate (decimal)] [Address (string)]\n");
    ShowMessages("syntax : \t.debug [close]\n");
    ShowMessages("\n");
    ShowMessages("\t\te.g : .debug remote serial 115200 com2\n");
    ShowMessages("\t\te.g : .debug remote namedpipe \\\\.\\pipe\\HyperDbgPipe\n");
    ShowMessages("\t\te.g : .debug prepare serial 115200 com1\n");
    ShowMessages("\t\te.g : .debug prepare serial 115200 com2\n");
    ShowMessages("\t\te.g : .debug close\n");
    ShowMessages(
        "\nvalid baud rates (decimal) : 110, 300, 600, 1200, 2400, 4800, 9600, "
        "14400, 19200, 38400, 56000, 57600, 115200, 128000, 256000\n");
    ShowMessages("valid COM ports : COM1, COM2, COM3, COM4 \n");
}
BOOLEAN CommandDebugCheckBaudrate(DWORD Baudrate){
    if (Baudrate == CBR_110 || Baudrate == CBR_300 || Baudrate == CBR_600 ||
        Baudrate == CBR_1200 || Baudrate == CBR_2400 || Baudrate == CBR_4800 ||
        Baudrate == CBR_9600 || Baudrate == CBR_14400 || Baudrate == CBR_19200 ||
        Baudrate == CBR_38400 || Baudrate == CBR_56000 || Baudrate == CBR_57600 ||
        Baudrate == CBR_115200 || Baudrate == CBR_128000 ||
        Baudrate == CBR_256000){
        return TRUE;
    }
    return FALSE;
}
BOOLEAN CommandDebugCheckComPort(const string & ComPort, UINT32 * Port){
    if (!ComPort.compare("com1")){
        *Port = COM1_PORT;
        return TRUE;
    }
    else if (!ComPort.compare("com2")){
        *Port = COM2_PORT;
        return TRUE;
    }
    else if (!ComPort.compare("com3")){
        *Port = COM3_PORT;
        return TRUE;
    }
    else if (!ComPort.compare("com4")){
        *Port = COM4_PORT;
        return TRUE;
    }
    return FALSE;
}
VOID CommandDebug(vector<string> SplitCommand, string Command){
    UINT32 Baudrate;
    UINT32 Port;
    if (SplitCommand.size() == 2 && !SplitCommand.at(1).compare("close")){
        if (g_IsSerialConnectedToRemoteDebuggee){
            KdCloseConnection();
        }
        else{
            ShowMessages(
                "err, debugger is not attached to any instance of debuggee\n");
        }
        return;
    }
    else if (SplitCommand.size() <= 3){
        ShowMessages("incorrect use of the '.debug'\n\n");
        CommandDebugHelp();
        return;
    }
    if (!SplitCommand.at(1).compare("remote")){
        if (!SplitCommand.at(2).compare("serial")){
            if (SplitCommand.size() != 5){
                ShowMessages("incorrect use of the '.debug'\n\n");
                CommandDebugHelp();
                return;
            }
            if (!IsNumber(SplitCommand.at(3))){
                ShowMessages("unknown parameter '%s'\n\n",
                             SplitCommand.at(3).c_str());
                CommandDebugHelp();
                return;
            }
            Baudrate = stoi(SplitCommand.at(3));
            if (!CommandDebugCheckBaudrate(Baudrate)){
                ShowMessages("err, baud rate is invalid\n\n");
                CommandDebugHelp();
                return;
            }
            if (!CommandDebugCheckComPort(SplitCommand.at(4), &Port)){
                ShowMessages("err, COM port is invalid\n\n");
                CommandDebugHelp();
                return;
            }
            KdPrepareAndConnectDebugPort(SplitCommand.at(4).c_str(), Baudrate, Port, FALSE, FALSE);
        }
        else if (!SplitCommand.at(2).compare("namedpipe")){
            string Delimiter = "namedpipe";
            string Token     = Command.substr(
                Command.find(Delimiter) + Delimiter.size() + 1,
                Command.size());
            KdPrepareAndConnectDebugPort(Token.c_str(), NULL, NULL, FALSE, TRUE);
        }
        else{
            ShowMessages("unknown parameter '%s'\n\n", SplitCommand.at(2).c_str());
            CommandDebugHelp();
            return;
        }
    }
    else if (!SplitCommand.at(1).compare("prepare")){
        if (SplitCommand.size() != 5){
            ShowMessages("incorrect use of the '.debug'\n\n");
            CommandDebugHelp();
            return;
        }
        if (!SplitCommand.at(2).compare("serial")){
            if (!IsNumber(SplitCommand.at(3))){
                ShowMessages("unknown parameter '%s'\n\n",
                             SplitCommand.at(3).c_str());
                CommandDebugHelp();
                return;
            }
            Baudrate = stoi(SplitCommand.at(3));
            if (!CommandDebugCheckBaudrate(Baudrate)){
                ShowMessages("err, baud rate is invalid\n\n");
                CommandDebugHelp();
                return;
            }
            if (!CommandDebugCheckComPort(SplitCommand.at(4), &Port)){
                ShowMessages("err, COM port is invalid\n\n");
                CommandDebugHelp();
                return;
            }
            KdPrepareAndConnectDebugPort(SplitCommand.at(4).c_str(), Baudrate, Port, TRUE, FALSE);
        }
        else{
            ShowMessages("invalid parameter '%s'\n\n", SplitCommand.at(2));
            CommandDebugHelp();
            return;
        }
    }
    else{
        ShowMessages("invalid parameter '%s'\n\n", SplitCommand.at(1));
        CommandDebugHelp();
        return;
    }
}
