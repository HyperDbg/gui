










#include "pch.h"




extern BOOLEAN g_IsConnectedToHyperDbgLocally;
extern BOOLEAN g_IsDebuggerModulesLoaded;
extern BOOLEAN g_IsSerialConnectedToRemoteDebuggee;
extern BOOLEAN g_IsSerialConnectedToRemoteDebugger;






VOID
CommandUnloadHelp()
{
    ShowMessages(
        "unload : unloads the kernel modules and uninstalls the drivers.\n\n");

    ShowMessages("syntax : \tunload [remove] [ModuleName (string)]\n");

    ShowMessages("\n");
    ShowMessages("\t\te.g : unload vmm\n");
    ShowMessages("\t\te.g : unload remove vmm\n");
}








VOID
CommandUnload(vector<string> SplittedCommand, string Command)
{
    if (SplittedCommand.size() != 2 && SplittedCommand.size() != 3)
    {
        ShowMessages("incorrect use of 'unload'\n\n");
        CommandUnloadHelp();
        return;
    }

    
    
    
    if ((SplittedCommand.size() == 2 && !SplittedCommand.at(1).compare("vmm")) ||
        (SplittedCommand.size() == 3 && !SplittedCommand.at(2).compare("vmm") &&
         !SplittedCommand.at(1).compare("remove")))
    {
        if (!g_IsConnectedToHyperDbgLocally)
        {
            ShowMessages("you're not connected to any instance of HyperDbg, did you "
                         "use '.connect'? \n");
            return;
        }

        
        
        
        if (g_IsSerialConnectedToRemoteDebuggee || g_IsSerialConnectedToRemoteDebugger)
        {
            ShowMessages("you're connected to a an instance of HyperDbg, please use "
                         "'.debug close' command\n");
            return;
        }

        if (g_IsDebuggerModulesLoaded)
        {
            HyperDbgUnloadVmm();
        }
        else
        {
            ShowMessages("there is nothing to unload\n");
        }

        
        
        
        if (!SplittedCommand.at(1).compare("remove"))
        {
            
            
            
            if (HyperDbgStopVmmDriver())
            {
                ShowMessages("err, failed to stop driver\n");
                return;
            }

            
            
            
            if (HyperDbgUninstallVmmDriver())
            {
                ShowMessages("err, failed to uninstall the driver\n");
                return;
            }

            ShowMessages("the driver is removed\n");
        }
    }
    else
    {
        
        
        
        ShowMessages("module not found, currently 'vmm' is the only available "
                     "module for HyperDbg\n");
    }
}
