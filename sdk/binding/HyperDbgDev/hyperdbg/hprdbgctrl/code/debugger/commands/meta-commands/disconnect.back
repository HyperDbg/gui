










#include "pch.h"




extern BOOLEAN g_IsConnectedToHyperDbgLocally;
extern BOOLEAN g_IsConnectedToRemoteDebuggee;
extern HANDLE  g_RemoteDebuggeeListeningThread;
extern HANDLE  g_EndOfMessageReceivedEvent;






VOID
CommandDisconnectHelp()
{
    ShowMessages(".disconnect : disconnects from a debugging session (it won't "
                 "unload the modules).\n\n");

    ShowMessages("syntax : \t.disconnect \n");
}








VOID
CommandDisconnect(vector<string> SplittedCommand, string Command)
{
    if (SplittedCommand.size() != 1)
    {
        ShowMessages("incorrect use of '.disconnect'\n\n");
        CommandDisconnectHelp();
        return;
    }

    if (!g_IsConnectedToHyperDbgLocally && !g_IsConnectedToRemoteDebuggee)
    {
        ShowMessages("you're not connected to any instance of HyperDbg, did you "
                     "use '.connect'? \n");
        return;
    }

    
    
    
    
    if (g_IsConnectedToHyperDbgLocally && g_DeviceHandle)
    {
        ShowMessages("you cannot disconnect in local debugging while the "
                     "driver is still loaded. please use 'unload' command before "
                     "disconnecting from the current instance of debugger\n");
        return;
    }

    
    
    
    g_IsConnectedToHyperDbgLocally = FALSE;

    
    
    
    if (g_IsConnectedToRemoteDebuggee)
    {
        
        
        
        
        TerminateThread(g_RemoteDebuggeeListeningThread, 0);
        CloseHandle(g_RemoteDebuggeeListeningThread);
        CloseHandle(g_EndOfMessageReceivedEvent);
        g_EndOfMessageReceivedEvent = NULL;

        RemoteConnectionCloseTheConnectionWithDebuggee();

        g_IsConnectedToRemoteDebuggee = FALSE;
    }

    ShowMessages("disconnected successfully\n");
}
