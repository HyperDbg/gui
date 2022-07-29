package meta-commands
//back\HyperDbgDev\hyperdbg\hprdbgctrl\code\debugger\commands\meta-commands\connect.cpp.back

type (
Connect interface{
CommandConnectHelp()(ok bool)//col:42
CommandConnect()(ok bool)//col:143
}
)

func NewConnect() { return & connect{} }

func (c *connect)CommandConnectHelp()(ok bool){//col:42
/*CommandConnectHelp()
{
    ShowMessages(".connect : connects to a remote or local machine to start "
                 "debugging.\n\n");
    ShowMessages("syntax : \t.connect [local]\n");
    ShowMessages("syntax : \t.connect [Ip (string)] [Port (decimal)]\n");
    ShowMessages("\n");
    ShowMessages("\t\te.g : .connect local\n");
    ShowMessages("\t\te.g : .connect 192.168.1.5 50000\n");
}*/
return true
}

func (c *connect)CommandConnect()(ok bool){//col:143
/*CommandConnect(vector<string> SplittedCommand, string Command)
{
    string ip;
    string port;
    if (g_IsConnectedToHyperDbgLocally || g_IsConnectedToRemoteDebuggee ||
        g_IsConnectedToRemoteDebugger)
    {
        ShowMessages("you're connected to a debugger, please use '.disconnect' "
                     "command\n");
        return;
    }
    if (g_IsSerialConnectedToRemoteDebuggee ||
        g_IsSerialConnectedToRemoteDebugger)
    {
        ShowMessages("you're connected to a an instance of HyperDbg, please use "
                     "'.debug close' command\n");
        return;
    }
    if (SplittedCommand.size() == 1)
    {
        ShowMessages("incorrect use of '.connect'\n\n");
        CommandConnectHelp();
        return;
    }
    else if (SplittedCommand.at(1) == "local" && SplittedCommand.size() == 2)
    {
        ShowMessages("local debugging (vmi-mode)\n");
        g_IsConnectedToHyperDbgLocally = TRUE;
        return;
    }
    else if (SplittedCommand.size() == 3 || SplittedCommand.size() == 2)
    {
        ip = SplittedCommand.at(1);
        if (SplittedCommand.size() == 3)
        {
            port = SplittedCommand.at(2);
        }
        if (!ValidateIP(ip))
        {
            ShowMessages("incorrect ip address\n");
            return;
        }
        if (SplittedCommand.size() == 3)
        {
            if (!IsNumber(port) || stoi(port) > 65535 || stoi(port) < 0)
            {
                ShowMessages("incorrect port\n");
                return;
            }
            g_ServerIp   = ip;
            g_ServerPort = port;
            RemoteConnectionConnect(ip.c_str(), port.c_str());
        }
        else
        {
            g_ServerIp   = ip;
            g_ServerPort = DEFAULT_PORT;
            RemoteConnectionConnect(ip.c_str(), DEFAULT_PORT);
        }
    }
    else
    {
        ShowMessages("incorrect use of '.connect'\n\n");
        CommandConnectHelp();
        return;
    }
}*/
return true
}



