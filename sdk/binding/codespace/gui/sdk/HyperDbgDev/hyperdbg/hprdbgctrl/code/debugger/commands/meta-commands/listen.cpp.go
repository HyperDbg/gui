package meta-commands
//binding\codespace\gui\sdk\HyperDbgDev\hyperdbg\hprdbgctrl\code\debugger\commands\meta-commands\listen.cpp.back

type (
Listen interface{
CommandListenHelp()(ok bool)//col:12
CommandListen()(ok bool)//col:58
}
)

func NewListen() { return & listen{} }

func (l *listen)CommandListenHelp()(ok bool){//col:12
/*CommandListenHelp()
{
    ShowMessages(".listen : listens for a client to connect to HyperDbg (works as "
                 "a guest server).\n\n");
    ShowMessages("note : \tif you don't specify port then HyperDbg uses the "
                 "default port (%s)\n",
                 DEFAULT_PORT);
    ShowMessages("syntax : \t.listen [Port (decimal)]\n");
    ShowMessages("\n");
    ShowMessages("\t\te.g : .listen\n");
    ShowMessages("\t\te.g : .listen 50000\n");
}*/
return true
}

func (l *listen)CommandListen()(ok bool){//col:58
/*CommandListen(vector<string> SplittedCommand, string Command)
{
    string port;
    if (SplittedCommand.size() >= 3)
    {
        ShowMessages("incorrect use of '.listen'\n\n");
        CommandListenHelp();
        return;
    }
    if (g_IsConnectedToHyperDbgLocally || g_IsConnectedToRemoteDebuggee ||
        g_IsConnectedToRemoteDebugger)
    {
        ShowMessages("you're connected to a debugger, please use '.disconnect' "
                     "command\n");
        return;
    }
    if (g_IsSerialConnectedToRemoteDebuggee || g_IsSerialConnectedToRemoteDebugger)
    {
        ShowMessages("you're connected to a an instance of HyperDbg, please use "
                     "'.debug close' command\n");
        return;
    }
    if (SplittedCommand.size() == 1)
    {
        ShowMessages("listening on %s ...\n", DEFAULT_PORT);
        RemoteConnectionListen(DEFAULT_PORT);
        return;
    }
    else if (SplittedCommand.size() == 2)
    {
        port = SplittedCommand.at(1);
        if (!IsNumber(port) || stoi(port) > 65535 || stoi(port) < 0)
        {
            ShowMessages("incorrect port\n");
            return;
        }
        ShowMessages("listening on %s ...\n", port.c_str());
        RemoteConnectionListen(port.c_str());
    }
    else
    {
        ShowMessages("incorrect use of '.listen'\n\n");
        CommandListenHelp();
        return;
    }
}*/
return true
}



